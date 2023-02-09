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
		Classes: []*types.Class{
			{
				Id:     types.OBTClass,
				Name:   "Obada network NFT Token",
				Symbol: types.OBTClass,
				Uri:    "https://obada.io",
			},
		},
	}

	k, ctx := keepertest.ObitKeeper(t)
	obit.InitGenesis(ctx, *k, genesisState)
	got := obit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState, got)
}
