package obit_test

import (
	"testing"

	keepertest "github.com/obada-foundation/fullcore/testutil/keeper"
	"github.com/obada-foundation/fullcore/testutil/nullify"
	"github.com/obada-foundation/fullcore/x/obit"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TaList: []types.Ta{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TaCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ObitKeeper(t)
	obit.InitGenesis(ctx, *k, genesisState)
	got := obit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaList, got.TaList)
	require.Equal(t, genesisState.TaCount, got.TaCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
