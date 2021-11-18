package core_test

import (
	"testing"

	keepertest "github.com/obada-foundation/core/testutil/keeper"
	"github.com/obada-foundation/core/x/core"
	"github.com/obada-foundation/core/x/core/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoreKeeper(t)
	core.InitGenesis(ctx, *k, genesisState)
	got := core.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
