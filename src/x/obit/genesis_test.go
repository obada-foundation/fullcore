package obit_test

import (
	"testing"

	keepertest "github.com/obada-foundation/core/testutil/keeper"
	"github.com/obada-foundation/core/x/obit"
	"github.com/obada-foundation/core/x/obit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ObitKeeper(t)
	obit.InitGenesis(ctx, *k, genesisState)
	got := obit.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
