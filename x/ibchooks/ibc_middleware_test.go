package ibchooks_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/MonCatCat/provenance/app"
	"github.com/MonCatCat/provenance/internal/pioconfig"
	testutil "github.com/MonCatCat/provenance/testutil/ibc"
	"github.com/MonCatCat/provenance/x/ibchooks"
	"github.com/MonCatCat/provenance/x/ibchooks/keeper"
	"github.com/MonCatCat/provenance/x/marker/types"
	markertypes "github.com/MonCatCat/provenance/x/marker/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdksim "github.com/cosmos/cosmos-sdk/simapp"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v6/modules/core/exported"
	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

var (
	_ ibchooks.Hooks = TestRecvOverrideHooks{}
	_ ibchooks.Hooks = TestRecvBeforeAfterHooks{}
)

type Status struct {
	OverrideRan bool
	BeforeRan   bool
	AfterRan    bool
}

// Recv
type TestRecvOverrideHooks struct{ Status *Status }

func (t TestRecvOverrideHooks) OnRecvPacketOverride(im ibchooks.IBCMiddleware, ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) ibcexported.Acknowledgement {
	t.Status.OverrideRan = true
	ack := im.App.OnRecvPacket(ctx, packet, relayer)
	return ack
}

type TestRecvBeforeAfterHooks struct{ Status *Status }

func (t TestRecvBeforeAfterHooks) OnRecvPacketBeforeHook(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress) {
	t.Status.BeforeRan = true
}

func (t TestRecvBeforeAfterHooks) OnRecvPacketAfterHook(ctx sdk.Context, packet channeltypes.Packet, relayer sdk.AccAddress, ack ibcexported.Acknowledgement) {
	t.Status.AfterRan = true
}

type HooksTestSuite struct {
	suite.Suite

	App         *app.App
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress

	coordinator *ibctesting.Coordinator

	chainA *testutil.TestChain
	chainB *testutil.TestChain

	path *ibctesting.Path
}

func SetupSimApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	pioconfig.SetProvenanceConfig(sdk.DefaultBondDenom, 0)
	db := dbm.NewMemDB()
	encCdc := app.MakeEncodingConfig()
	provenanceApp := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 5, encCdc, sdksim.EmptyAppOptions{})
	genesis := app.NewDefaultGenesisState(encCdc.Marshaler)
	return provenanceApp, genesis
}

func init() {
	ibctesting.DefaultTestingAppInit = SetupSimApp
}

func (suite *HooksTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 2)
	suite.chainA = &testutil.TestChain{
		TestChain: suite.coordinator.GetChain(ibctesting.GetChainID(1)),
	}
	suite.chainB = &testutil.TestChain{
		TestChain: suite.coordinator.GetChain(ibctesting.GetChainID(2)),
	}
	suite.path = NewTransferPath(suite.chainA, suite.chainB)
	suite.coordinator.Setup(suite.path)
}

func TestIBCHooksTestSuite(t *testing.T) {
	suite.Run(t, new(HooksTestSuite))
}

func NewTransferPath(chainA, chainB *testutil.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA.TestChain, chainB.TestChain)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointA.ChannelConfig.Version = transfertypes.Version
	path.EndpointB.ChannelConfig.Version = transfertypes.Version

	return path
}

func (suite *HooksTestSuite) TestOnRecvPacketHooks() {
	var (
		trace    transfertypes.DenomTrace
		amount   sdk.Int
		receiver string
		status   Status
	)

	testCases := []struct {
		msg      string
		malleate func(*Status)
		expPass  bool
	}{
		{"override", func(status *Status) {
			suite.chainB.GetProvenanceApp().TransferStack.
				ICS4Middleware.Hooks = TestRecvOverrideHooks{Status: status}
		}, true},
		{"before and after", func(status *Status) {
			suite.chainB.GetProvenanceApp().TransferStack.
				ICS4Middleware.Hooks = TestRecvBeforeAfterHooks{Status: status}
		}, true},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.msg, func() {
			suite.SetupTest() // reset

			path := NewTransferPath(suite.chainA, suite.chainB)
			suite.coordinator.Setup(path)
			receiver = suite.chainB.SenderAccount.GetAddress().String() // must be explicitly changed in malleate
			status = Status{}

			amount = sdk.NewInt(100) // must be explicitly changed in malleate
			seq := uint64(1)

			trace = transfertypes.ParseDenomTrace(sdk.DefaultBondDenom)

			// send coin from chainA to chainB
			transferMsg := transfertypes.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.NewCoin(trace.IBCDenom(), amount), suite.chainA.SenderAccount.GetAddress().String(), receiver, clienttypes.NewHeight(1, 110), 0, "")
			_, err := suite.chainA.SendMsgs(transferMsg)
			suite.Require().NoError(err) // message committed

			tc.malleate(&status)

			data := transfertypes.NewFungibleTokenPacketData(trace.GetFullDenomPath(), amount.String(), suite.chainA.SenderAccount.GetAddress().String(), receiver, "")
			packet := channeltypes.NewPacket(data.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.NewHeight(1, 100), 0)

			ack := suite.chainB.GetProvenanceApp().TransferStack.
				OnRecvPacket(suite.chainB.GetContext(), packet, suite.chainA.SenderAccount.GetAddress())

			if tc.expPass {
				suite.Require().True(ack.Success())
			} else {
				suite.Require().False(ack.Success())
			}

			if _, ok := suite.chainB.GetProvenanceApp().TransferStack.
				ICS4Middleware.Hooks.(TestRecvOverrideHooks); ok {
				suite.Require().True(status.OverrideRan)
				suite.Require().False(status.BeforeRan)
				suite.Require().False(status.AfterRan)
			}

			if _, ok := suite.chainB.GetProvenanceApp().TransferStack.
				ICS4Middleware.Hooks.(TestRecvBeforeAfterHooks); ok {
				suite.Require().False(status.OverrideRan)
				suite.Require().True(status.BeforeRan)
				suite.Require().True(status.AfterRan)
			}
		})
	}
}

func (suite *HooksTestSuite) makeMockPacket(receiver, memo string, prevSequence uint64) channeltypes.Packet {
	packetData := transfertypes.FungibleTokenPacketData{
		Denom:    sdk.DefaultBondDenom,
		Amount:   "1",
		Sender:   suite.chainB.SenderAccount.GetAddress().String(),
		Receiver: receiver,
		Memo:     memo,
	}

	return channeltypes.NewPacket(
		packetData.GetBytes(),
		prevSequence+1,
		suite.path.EndpointB.ChannelConfig.PortID,
		suite.path.EndpointB.ChannelID,
		suite.path.EndpointA.ChannelConfig.PortID,
		suite.path.EndpointA.ChannelID,
		clienttypes.NewHeight(0, 100),
		0,
	)
}

func (suite *HooksTestSuite) receivePacket(receiver, memo string) []byte {
	return suite.receivePacketWithSequence(receiver, memo, 0)
}

func (suite *HooksTestSuite) receivePacketWithSequence(receiver, memo string, prevSequence uint64) []byte {
	channelCap := suite.chainB.GetChannelCapability(
		suite.path.EndpointB.ChannelConfig.PortID,
		suite.path.EndpointB.ChannelID)

	packet := suite.makeMockPacket(receiver, memo, prevSequence)

	seq, err := suite.chainB.GetProvenanceApp().HooksICS4Wrapper.SendPacket(
		suite.chainB.GetContext(),
		channelCap,
		packet.SourcePort,
		packet.SourceChannel,
		packet.TimeoutHeight,
		packet.TimeoutTimestamp,
		packet.Data)
	suite.Require().NoError(err, "IBC send failed. Expected success. %s", err)
	suite.Require().NotZero(seq, "IBC send expected positive sequence number, %d", seq)

	// Update both clients
	err = suite.path.EndpointB.UpdateClient()
	suite.Require().NoError(err)
	err = suite.path.EndpointA.UpdateClient()
	suite.Require().NoError(err)

	// recv in chain a
	res, err := suite.path.EndpointA.RecvPacketWithResult(packet)
	suite.Require().NoError(err)

	// get the ack from the chain a's response
	ack, err := ibctesting.ParseAckFromEvents(res.GetEvents())
	suite.Require().NoError(err)

	// manually send the acknowledgement to chain b
	err = suite.path.EndpointA.AcknowledgePacket(packet, ack)
	suite.Require().NoError(err)
	return ack
}

func (suite *HooksTestSuite) TestRecvTransferWithMetadata() {
	// Setup contract
	codeID := suite.chainA.StoreContractEchoDirect(&suite.Suite)
	addr := suite.chainA.InstantiateContract(&suite.Suite, "{}", codeID)
	memo := fmt.Sprintf(`{"marker":{},"wasm":{"contract":"%s","msg":{"echo":{"msg":"test"}}}}`, addr)
	ackBytes := suite.receivePacket(addr.String(), memo)
	ackStr := string(ackBytes)
	fmt.Println(ackStr)
	var ack map[string]string // This can't be unmarshalled to Acknowledgement because it's fetched from the events
	err := json.Unmarshal(ackBytes, &ack)
	suite.Require().NoError(err)
	suite.Require().NotContains(ack, "error")
	suite.Require().Equal(ack["result"], "eyJjb250cmFjdF9yZXN1bHQiOiJkR2hwY3lCemFHOTFiR1FnWldOb2J3PT0iLCJpYmNfYWNrIjoiZXlKeVpYTjFiSFFpT2lKQlVUMDlJbjA9In0=", "Ack result must match")
}

// After successfully executing a wasm call, the contract should have the funds sent via IBC
func (suite *HooksTestSuite) TestFundsAreTransferredToTheContract() {
	// Setup contract
	codeID := suite.chainA.StoreContractEchoDirect(&suite.Suite)
	addr := suite.chainA.InstantiateContract(&suite.Suite, "{}", codeID)

	// Check that the contract has no funds
	localDenom := ibchooks.MustExtractDenomFromPacketOnRecv(suite.makeMockPacket("", "", 0))
	balance := suite.chainA.GetProvenanceApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)

	// Execute the contract via IBC
	ackBytes := suite.receivePacket(addr.String(), fmt.Sprintf(`{"marker":{},"wasm":{"contract":"%s","msg":{"echo":{"msg":"test"}}}}`, addr))
	ackStr := string(ackBytes)
	fmt.Println(ackStr)
	var ack map[string]string // This can't be unmarshalled to Acknowledgement because it's fetched from the events
	err := json.Unmarshal(ackBytes, &ack)
	suite.Require().NoError(err)
	suite.Require().NotContains(ack, "error")
	suite.Require().Equal(ack["result"], "eyJjb250cmFjdF9yZXN1bHQiOiJkR2hwY3lCemFHOTFiR1FnWldOb2J3PT0iLCJpYmNfYWNrIjoiZXlKeVpYTjFiSFFpT2lKQlVUMDlJbjA9In0=")

	// Check that the token has now been transferred to the contract
	balance = suite.chainA.GetProvenanceApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(1), balance.Amount)
}

// If the wasm call wails, the contract acknowledgement should be an error and the funds returned
func (suite *HooksTestSuite) TestFundsAreReturnedOnFailedContractExec() {
	// Setup contract
	codeID := suite.chainA.StoreContractEchoDirect(&suite.Suite)
	addr := suite.chainA.InstantiateContract(&suite.Suite, "{}", codeID)

	// Check that the contract has no funds
	localDenom := ibchooks.MustExtractDenomFromPacketOnRecv(suite.makeMockPacket("", "", 0))
	balance := suite.chainA.GetProvenanceApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)

	// Execute the contract via IBC with a message that the contract will reject
	ackBytes := suite.receivePacket(addr.String(), fmt.Sprintf(`{"marker":{},"wasm":{"contract":"%s","msg":{"not_echo":{"msg":"test"}}}}`, addr))
	ackStr := string(ackBytes)
	fmt.Println(ackStr)
	var ack map[string]string // This can't be unmarshalled to Acknowledgement because it's fetched from the events
	err := json.Unmarshal(ackBytes, &ack)
	suite.Require().NoError(err)
	suite.Require().Contains(ack, "error")

	// Check that the token has now been transferred to the contract
	balance = suite.chainA.GetProvenanceApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	fmt.Println(balance)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)
}

// After successfully executing a wasm call, the contract should have the funds sent via IBC
func (suite *HooksTestSuite) TestFundTracking() {
	// Setup contract
	codeID := suite.chainA.StoreContractCounterDirect(&suite.Suite)
	addr := suite.chainA.InstantiateContract(&suite.Suite, `{"count": 0}`, codeID)

	// Check that the contract has no funds
	localDenom := ibchooks.MustExtractDenomFromPacketOnRecv(suite.makeMockPacket("", "", 0))
	balance := suite.chainA.GetProvenanceApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(0), balance.Amount)

	memo := fmt.Sprintf(`{"marker":{},"wasm":{"contract":"%s","msg":{"increment":{}}}}`, addr)

	// Execute the contract via IBC
	suite.receivePacket(
		addr.String(), memo)

	prefix := sdk.GetConfig().GetBech32AccountAddrPrefix()
	senderLocalAcc, err := keeper.DeriveIntermediateSender("channel-0", suite.chainB.SenderAccount.GetAddress().String(), prefix)
	suite.Require().NoError(err)

	state := suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"count":0}`, state)

	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_total_funds": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"total_funds":[{"denom":"ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878","amount":"1"}]}`, state)

	suite.receivePacketWithSequence(
		addr.String(),
		memo, 1)

	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"count":1}`, state)

	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_total_funds": {"addr": "%s"}}`, senderLocalAcc)))
	suite.Require().Equal(`{"total_funds":[{"denom":"ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878","amount":"2"}]}`, state)

	// Check that the token has now been transferred to the contract
	balance = suite.chainA.GetProvenanceApp().BankKeeper.GetBalance(suite.chainA.GetContext(), addr, localDenom)
	suite.Require().Equal(sdk.NewInt(2), balance.Amount)
}

// custom MsgTransfer constructor that supports Memo
func NewMsgTransfer(
	token sdk.Coin, sender, receiver string, memo string,
) *transfertypes.MsgTransfer {
	return &transfertypes.MsgTransfer{
		SourcePort:       "transfer",
		SourceChannel:    "channel-0",
		Token:            token,
		Sender:           sender,
		Receiver:         receiver,
		TimeoutHeight:    clienttypes.NewHeight(0, 100),
		TimeoutTimestamp: 0,
		Memo:             memo,
	}
}

type Direction int64

const (
	AtoB Direction = iota
	BtoA
)

func (suite *HooksTestSuite) GetEndpoints(direction Direction) (sender *ibctesting.Endpoint, receiver *ibctesting.Endpoint) {
	switch direction {
	case AtoB:
		sender = suite.path.EndpointA
		receiver = suite.path.EndpointB
	case BtoA:
		sender = suite.path.EndpointB
		receiver = suite.path.EndpointA
	}
	return sender, receiver
}

func (suite *HooksTestSuite) RelayPacket(packet channeltypes.Packet, direction Direction) (*sdk.Result, []byte) {
	sender, receiver := suite.GetEndpoints(direction)

	err := receiver.UpdateClient()
	suite.Require().NoError(err)

	// receiver Receives
	receiveResult, err := receiver.RecvPacketWithResult(packet)
	suite.Require().NoError(err)

	ack, err := ibctesting.ParseAckFromEvents(receiveResult.GetEvents())
	suite.Require().NoError(err)

	// sender Acknowledges
	err = sender.AcknowledgePacket(packet, ack)
	suite.Require().NoError(err)

	err = sender.UpdateClient()
	suite.Require().NoError(err)
	err = receiver.UpdateClient()
	suite.Require().NoError(err)

	return receiveResult, ack
}

func (suite *HooksTestSuite) FullSend(msg sdk.Msg, direction Direction) (*sdk.Result, *sdk.Result, string, error) {
	var sender *testutil.TestChain
	switch direction {
	case AtoB:
		sender = suite.chainA
	case BtoA:
		sender = suite.chainB
	}
	sendResult, err := sender.SendMsgsNoCheck(msg)
	suite.Require().NoError(err)

	packet, err := ibctesting.ParsePacketFromEvents(sendResult.GetEvents())
	suite.Require().NoError(err)

	receiveResult, ack := suite.RelayPacket(packet, direction)

	return sendResult, receiveResult, string(ack), err
}

func (suite *HooksTestSuite) TestAcks() {
	codeID := suite.chainA.StoreContractCounterDirect(&suite.Suite)
	addr := suite.chainA.InstantiateContract(&suite.Suite, `{"count": 0}`, codeID)

	// Generate swap instructions for the contract
	callbackMemo := fmt.Sprintf(`{"ibc_callback":"%s"}`, addr)
	// Send IBC transfer with the memo with crosschain-swap instructions
	transferMsg := NewMsgTransfer(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)), suite.chainA.SenderAccount.GetAddress().String(), addr.String(), callbackMemo)
	suite.FullSend(transferMsg, AtoB)

	// The test contract will increment the counter for itself every time it receives an ack
	state := suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, addr)))
	suite.Require().Equal(`{"count":1}`, state)

	suite.FullSend(transferMsg, AtoB)
	state = suite.chainA.QueryContract(
		&suite.Suite, addr,
		[]byte(fmt.Sprintf(`{"get_count": {"addr": "%s"}}`, addr)))
	suite.Require().Equal(`{"count":2}`, state)

}

func (suite *HooksTestSuite) TestSendWithoutMemo() {
	chainA := suite.chainA.GetProvenanceApp()
	chainASenderAddress := suite.chainA.SenderAccount.GetAddress()
	chainBSenderAddress := suite.chainB.SenderAccount.GetAddress()

	// Sending a packet without memo to ensure that the ibc_callback middleware doesn't interfere with a regular send
	transferMsg := NewMsgTransfer(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)), chainASenderAddress.String(), chainBSenderAddress.String(), "")
	_, _, ack, err := suite.FullSend(transferMsg, AtoB)
	suite.Require().NoError(err, "FullSend()")
	suite.Require().Contains(ack, "result")
	prefixedDenom := transfertypes.GetDenomPrefix(suite.path.EndpointB.ChannelConfig.PortID, suite.path.EndpointB.ChannelID) + sdk.DefaultBondDenom
	denom := transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
	marker, err := suite.chainB.GetProvenanceApp().MarkerKeeper.GetMarkerByDenom(suite.chainB.GetContext(), denom)
	suite.Require().NoError(err, "GetMarkerByDenom()")
	suite.Require().Equal(marker.GetDenom(), denom)

	transferMsg = NewMsgTransfer(sdk.NewCoin(denom, sdk.NewInt(100)), chainBSenderAddress.String(), chainASenderAddress.String(), "")
	_, _, ack, err = suite.FullSend(transferMsg, BtoA)
	suite.Require().NoError(err, "FullSend()")
	suite.Require().Contains(ack, "result")
	stakeAddr := markertypes.MustGetMarkerAddress(sdk.DefaultBondDenom)
	marker, err = chainA.MarkerKeeper.GetMarker(suite.chainA.GetContext(), stakeAddr)
	suite.Require().NoError(err, "GetMarker()")
	suite.Require().Nil(marker, "Should not create marker on chain A")

	hotdogs := "hotdogs"
	rv := markertypes.NewMarkerAccount(
		chainA.AccountKeeper.NewAccountWithAddress(suite.chainA.GetContext(), markertypes.MustGetMarkerAddress(hotdogs)).(*authtypes.BaseAccount),
		sdk.NewInt64Coin(hotdogs, 10000),
		suite.chainA.SenderAccount.GetAddress(),
		[]types.AccessGrant{
			{Address: chainASenderAddress.String(), Permissions: markertypes.AccessList{markertypes.Access_Transfer, markertypes.Access_Withdraw}},
		},
		types.StatusProposed,
		types.MarkerType_RestrictedCoin,
		true,  // supply fixed
		true,  // allow gov
		false, // no force transfer
		[]string{},
	)

	chainA.MarkerKeeper.AddSetNetAssetValues(suite.chainA.GetContext(), rv, []types.NetAssetValue{types.NewNetAssetValue(sdk.NewInt64Coin(types.UsdDenom, int64(1000)), 1)}, markertypes.ModuleName)
	suite.Require().NoError(err, "chainA AddSetNetAssetValues()")
	err = chainA.MarkerKeeper.AddFinalizeAndActivateMarker(suite.chainA.GetContext(), rv)
	suite.Require().NoError(err, "chainA AddFinalizeAndActivateMarker()")
	err = chainA.MarkerKeeper.WithdrawCoins(suite.chainA.GetContext(), chainASenderAddress, chainASenderAddress, hotdogs, sdk.NewCoins(sdk.NewInt64Coin(hotdogs, 55)))
	suite.Require().NoError(err, "chainA WithdrawCoins()")

	transferMsg = NewMsgTransfer(sdk.NewCoin(hotdogs, sdk.NewInt(55)), chainASenderAddress.String(), chainBSenderAddress.String(), "")
	_, _, ack, err = suite.FullSend(transferMsg, AtoB)
	suite.Require().NoError(err, "AtoB FullSend()")
	suite.Require().Contains(ack, "result", "FullSend() ack check")
	prefixedDenom = transfertypes.GetDenomPrefix(suite.path.EndpointB.ChannelConfig.PortID, suite.path.EndpointB.ChannelID) + hotdogs
	denom = transfertypes.ParseDenomTrace(prefixedDenom).IBCDenom()
	marker, err = suite.chainB.GetProvenanceApp().MarkerKeeper.GetMarkerByDenom(suite.chainB.GetContext(), denom)
	suite.Require().NoError(err, "chainB GetMarkerByDenom()")
	suite.Require().Equal(marker.GetDenom(), denom)
	suite.Require().True(marker.HasAccess(chainASenderAddress.String(), markertypes.Access_Transfer), "ChainB Ibc marker should have transfer rights added")
}
