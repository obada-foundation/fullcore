package fullcore_test

import (
	"testing"

	keepertest "github.com/obada-foundation/fullcore/testutil/keeper"
	"github.com/obada-foundation/fullcore/testutil/nullify"
	"github.com/obada-foundation/fullcore/x/fullcore"
	"github.com/obada-foundation/fullcore/x/fullcore/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FullcoreKeeper(t)
	fullcore.InitGenesis(ctx, *k, genesisState)
	got := fullcore.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
