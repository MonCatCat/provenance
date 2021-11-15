package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/provenance-io/provenance/x/msgfees/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the msgfees MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateMsgBasedFee(goCtx context.Context, request *types.CreateMsgBasedFeeRequest) (*types.CreateMsgBasedFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Validate transaction message.
	err := request.ValidateBasic()

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	existing, err := k.GetMsgBasedFee(ctx, request.GetMsgBasedFee().MsgTypeUrl)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	if existing != nil {
		return nil, sdkerrors.Wrap(types.ErrMsgFeeAlreadyExists, err.Error())
	}

	k.SetMsgBasedFee(ctx, *request.MsgBasedFee)

	return &types.CreateMsgBasedFeeResponse{
		MsgBasedFee: request.MsgBasedFee,
	}, nil
}

func (k msgServer) CalculateMsgBasedFees(goCtx context.Context, request *types.CalculateFeePerMsgRequest) (*types.CalculateMsgBasedFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	additionalFees := sdk.Coins{}
	reqTx := *request.Tx
	txBytes, err := reqTx.Marshal()
	if err != nil {
		return nil, fmt.Errorf("tx not of sdk.Tx type")
	}
	gasInfo, result, err := k.simulateFunc(txBytes)
	if err != nil {
		return nil, err
	}
	ctx.Logger().Info("NOTICE: Gas Info: %v Result: %v", gasInfo, result)

	msgs := request.Tx.GetMsgs()
	for _, msg := range msgs {
		typeURL := sdk.MsgTypeURL(msg)
		msgFees, err := k.GetMsgBasedFee(ctx, typeURL)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}

		if msgFees == nil {
			continue
		}
		if msgFees.AdditionalFee.IsPositive() {
			additionalFees = additionalFees.Add(sdk.NewCoin(msgFees.AdditionalFee.Denom, msgFees.AdditionalFee.Amount))
		}
	}

	return &types.CalculateMsgBasedFeesResponse{
		FeeAmount: additionalFees,
	}, nil
}
