package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/provenance-io/provenance/internal/antewrapper"
	msgfeestypes "github.com/provenance-io/provenance/x/msgfees/types"
)

type MsgFeeInvoker struct {
	msgFeeKeeper   msgfeestypes.MsgFeesKeeper
	bankKeeper     bankkeeper.Keeper
	accountKeeper  msgfeestypes.AccountKeeper
	feegrantKeeper msgfeestypes.FeegrantKeeper
	txDecoder      sdk.TxDecoder
}

// NewMsgFeeInvoker concrete impl of how to charge Msg Based Fees
func NewMsgFeeInvoker(bankKeeper bankkeeper.Keeper, accountKeeper msgfeestypes.AccountKeeper,
	feegrantKeeper msgfeestypes.FeegrantKeeper, msgFeeKeeper msgfeestypes.MsgFeesKeeper, decoder sdk.TxDecoder) MsgFeeInvoker {
	return MsgFeeInvoker{
		msgFeeKeeper,
		bankKeeper,
		accountKeeper,
		feegrantKeeper,
		decoder,
	}
}

func (afd MsgFeeInvoker) Invoke(ctx sdk.Context, simulate bool) (sdk.Coins, sdk.Events, error) {
	chargedFees := sdk.Coins{}
	eventsToReturn := sdk.Events{}

	if len(ctx.TxBytes()) != 0 {
		tx, err := afd.txDecoder(ctx.TxBytes())
		if err != nil {
			panic(fmt.Errorf("error in MsgFeeInvoker.Invoke() while getting txBytes: %w", err))
		}

		feeTx, err := antewrapper.GetFeeTx(tx)
		if err != nil {
			// For provenance, should be a FeeTx since antehandler should enforce it,
			// but not adding complexity here.
			panic(err)
		}

		feeGasMeter, err := antewrapper.GetFeeGasMeter(ctx)
		if err != nil {
			// For provenance, should be a FeeGasMeter since antehandler should enforce it,
			// but not adding complexity here.
			panic(err)
		}

		// eat up the gas cost for charging fees. (This one is on us, Cheers!, mainly because we don't want to fail at this step, imo, but we can remove this is f necessary)
		ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

		consumedFees := feeGasMeter.FeeConsumed()
		if consumedFees.IsAnyNegative() {
			return nil, nil, sdkerrors.ErrInvalidCoins.Wrapf("consumed fees %v are negative, which should not be possible, aborting", chargedFees)
		}

		// this sweeps all extra fees too, 1. keeps current behavior 2. accounts for priority mempool
		baseFeeConsumed := feeGasMeter.BaseFeeConsumed()
		unchargedFees, _ := feeTx.GetFee().SafeSub(baseFeeConsumed...)

		deductFeesFrom, err := antewrapper.GetFeePayerUsingFeeGrant(ctx, afd.feegrantKeeper, feeTx, unchargedFees, tx.GetMsgs())
		if err != nil {
			return nil, nil, err
		}

		deductFeesFromAcc := afd.accountKeeper.GetAccount(ctx, deductFeesFrom)
		if deductFeesFromAcc == nil {
			return nil, nil, sdkerrors.ErrUnknownAddress.Wrapf("fee payer address: %q does not exist", deductFeesFrom)
		}

		// If there's fees left to collect, or there were consumed fees, deduct/distribute them now.
		if !unchargedFees.IsZero() || !consumedFees.IsZero() {
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			err = afd.msgFeeKeeper.DeductFeesDistributions(afd.bankKeeper, ctx, deductFeesFromAcc, unchargedFees, feeGasMeter.FeeConsumedDistributions())
			if err != nil {
				return nil, nil, err
			}
			eventsToReturn = append(eventsToReturn, ctx.EventManager().Events()...)
		}
		// the uncharged fees have now been charged.
		chargedFees = chargedFees.Add(unchargedFees...)

		// If there were msg based fees, add some events for them.
		if !consumedFees.IsZero() {
			// Add event with fee breakdown between additional fees and the rest.
			nonMsgFees := baseFeeConsumed.Add(chargedFees...).Sub(consumedFees...)
			eventsToReturn = append(eventsToReturn, sdk.NewEvent(sdk.EventTypeTx,
				sdk.NewAttribute(antewrapper.AttributeKeyAdditionalFee, consumedFees.String()),
				sdk.NewAttribute(antewrapper.AttributeKeyBaseFee, nonMsgFees.String()),
				sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String())))

			// Add event with a breakdown of those fees.
			msgFeesSummaryEvent, err := sdk.TypedEventToEvent(feeGasMeter.EventFeeSummary())
			if err != nil {
				return nil, nil, err
			}
			if len(msgFeesSummaryEvent.Attributes) > 0 {
				eventsToReturn = append(eventsToReturn, msgFeesSummaryEvent)
			}
		}
	}

	return chargedFees, eventsToReturn, nil
}

func AggregateEvents(anteEvents []abci.Event, resultEvents []abci.Event) ([]abci.Event, []abci.Event) {
	if len(resultEvents) == 0 { // tx failed...fix fee event to have charged fee
		var err error
		var txFee sdk.Coins
		var feeIndex int
		var feeFound, spenderFound bool
		for i, event := range anteEvents {
			if !feeFound && event.Type == sdk.EventTypeTx && string(event.Attributes[0].Key) == sdk.AttributeKeyFee {
				feeFound = true
				feeIndex = i
			}
			// first spent coin event is the coin sent to fee module for tx
			if !spenderFound && event.Type == banktypes.EventTypeCoinSpent && string(event.Attributes[0].Key) == banktypes.AttributeKeySpender && len(anteEvents) >= i+3 {
				txFee, err = sdk.ParseCoinsNormalized(string(event.Attributes[1].Value))
				if err != nil {
					return nil, nil
				}
				spenderFound = true
			}
		}
		if feeFound && spenderFound {
			anteEvents[feeIndex].Attributes[0].Value = []byte(txFee.String())
		}
	}

	return anteEvents, resultEvents
}

// func AggregateEventsOld(resultEvents []abci.Event, feeEvents []abci.Event) ([]abci.Event, error) {
// 	var baseFee sdk.Coins
// 	var additionalFee sdk.Coins

// 	// 1.) Find amount in coins passed into the tx as a fee
// 	var txFee sdk.Coins
// 	var feeIndex int
// 	var feeFound, spenderFound bool
// 	for i, event := range resultEvents {
// 		if !feeFound && event.Type == sdk.EventTypeTx && string(event.Attributes[0].Key) == sdk.AttributeKeyFee {
// 			txFee, _ = sdk.ParseCoinsNormalized(string(event.Attributes[0].Value))
// 			feeFound = true
// 			feeIndex = i
// 		}
// 		if !spenderFound && event.Type == banktypes.EventTypeCoinSpent && string(event.Attributes[0].Key) == banktypes.AttributeKeySpender && len(resultEvents) >= i+3 {
// 			amount, _ := sdk.ParseCoinsNormalized(string(event.Attributes[0].Value))
// 			if !amount.IsEqual(txFee) { //
// 				hadOverage = true
// 			}
// 			spenderFound = true
// 		}
// 	}

// 	// 2.) Find any additional fees paid
// 	for _, event := range feeEvents {
// 		if event.Type == sdk.EventTypeTx && string(event.Attributes[0].Key) == antewrapper.AttributeKeyAdditionalFee {
// 			additionalFee, _ = sdk.ParseCoinsNormalized(string(event.Attributes[0].Value))
// 		}
// 	}

// 	// 3.) Calculate the total base fee
// 	baseFee = txFee.Sub(additionalFee...)

// 	// 4.)
// 	value := fmt.Sprintf("%v %v %v %v", txFee, feeIndex, baseFee, additionalFee)
// 	println(value)

// 	// var feemoduleAddress string
// 	// var hadOverage bool
// 	// var processedFeePay bool
// 	// // first coin_spent address is fee module account
// 	// for i, event := range resultEvents {
// 	// 	if event.Type == banktypes.EventTypeCoinSpent && string(event.Attributes[0].Key) == banktypes.AttributeKeySpender && len(resultEvents) >= i+3 {
// 	// 		feePayer := string(event.Attributes[0].Value)
// 	// 		amount, _ := sdk.ParseCoinsNormalized(string(event.Attributes[0].Value))
// 	// 		if !amount.IsEqual(txFee) { //
// 	// 			hadOverage = true
// 	// 		}
// 	// 	}
// 	// }

// 	// if hadOverage && resultEvents[feeIndex].Type == sdk.EventTypeTx && string(resultEvents[feeIndex].Attributes[0].Key) == sdk.AttributeKeyFee {
// 	// 	resultEvents[feeIndex].Attributes[0].Value = []byte(txFee.String())
// 	// }

// 	return append(resultEvents, feeEvents...), nil
// }
