package simulation_test

import (
	"math/rand"
	"testing"

	"github.com/MonCatCat/provenance/internal/pioconfig"
	"github.com/MonCatCat/provenance/x/msgfees/keeper"
	"github.com/MonCatCat/provenance/x/msgfees/simulation"
	"github.com/MonCatCat/provenance/x/msgfees/types"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	simapp "github.com/MonCatCat/provenance/app"
	simappparams "github.com/MonCatCat/provenance/app/params"
)

func TestProposalContents(t *testing.T) {
	app := simapp.Setup(t)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// initialize parameters
	s := rand.NewSource(1)
	r := rand.New(s)

	accounts := simtypes.RandomAccounts(r, 3)

	// execute ProposalContents function
	weightedProposalContent := simulation.ProposalContents(keeper.NewKeeper(app.AppCodec(), app.GetKey(types.ModuleName),
		app.GetSubspace(types.ModuleName), "", pioconfig.GetProvenanceConfig().FeeDenom, nil, nil, app.InterfaceRegistry()))
	require.Len(t, weightedProposalContent, 2)

	w0 := weightedProposalContent[0]

	// tests w0 interface:
	require.Equal(t, simulation.OpWeightAddMsgFeesProposal, w0.AppParamsKey())
	require.Equal(t, simappparams.DefaultWeightAddMsgFeeProposalContent, w0.DefaultWeight())

	content := w0.ContentSimulatorFn()(r, ctx, accounts)

	require.Equal(t, "fyzeOcbWwNbeHVIkPZBSpYuLyYggwexjxusrBqDOTtGTOWeLrQKjLxzIivHSlcxgdXhhuTSkuxKGLwQvuyNhYFmBZHeAerqyNEUz", content.GetDescription())
	require.Equal(t, "GqiQWIXnku", content.GetTitle())

	require.Equal(t, "msgfees", content.ProposalRoute())
	require.Equal(t, "AddMsgFee", content.ProposalType())
}
