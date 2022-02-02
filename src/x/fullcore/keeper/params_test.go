package keeper_test

import (
	"testing"

	testkeeper "github.com/obada-foundation/fullcore/testutil/keeper"
	"github.com/obada-foundation/fullcore/x/fullcore/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FullcoreKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
