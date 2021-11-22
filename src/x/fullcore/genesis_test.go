package fullcore_test

import (
	"testing"

	keepertest "github.com/obada-foundation/fullcore/testutil/keeper"
	"github.com/obada-foundation/fullcore/x/fullcore"
	"github.com/obada-foundation/fullcore/x/fullcore/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoreKeeper(t)
	fullcore.InitGenesis(ctx, *k, genesisState)
	got := fullcore.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
