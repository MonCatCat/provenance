package simulation_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/MonCatCat/provenance/app"
	"github.com/MonCatCat/provenance/x/name/simulation"
	"github.com/MonCatCat/provenance/x/name/types"
)

func TestDecodeStore(t *testing.T) {
	cdc := app.MakeEncodingConfig().Marshaler
	dec := simulation.NewDecodeStore(cdc)

	testNameRecord := types.NewNameRecord("test", sdk.AccAddress{}, true)

	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: types.NameKeyPrefix, Value: cdc.MustMarshal(&testNameRecord)},
			{Key: types.AddressKeyPrefix, Value: cdc.MustMarshal(&testNameRecord)},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Name Record", fmt.Sprintf("Name: A:[%v], B:[%v]\n", testNameRecord, testNameRecord)},
		{"Address Cache", fmt.Sprintf("Addr: A:[%v], B:[%v]\n", testNameRecord, testNameRecord)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
