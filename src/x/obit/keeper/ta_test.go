package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/obada-foundation/fullcore/testutil/keeper"
	"github.com/obada-foundation/fullcore/testutil/nullify"
	"github.com/obada-foundation/fullcore/x/obit/keeper"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/stretchr/testify/require"
)

func createNTa(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ta {
	items := make([]types.Ta, n)
	for i := range items {
		items[i].Id = keeper.AppendTa(ctx, items[i])
	}
	return items
}

func TestTaGet(t *testing.T) {
	keeper, ctx := keepertest.ObitKeeper(t)
	items := createNTa(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTa(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTaRemove(t *testing.T) {
	keeper, ctx := keepertest.ObitKeeper(t)
	items := createNTa(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTa(ctx, item.Id)
		_, found := keeper.GetTa(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTaGetAll(t *testing.T) {
	keeper, ctx := keepertest.ObitKeeper(t)
	items := createNTa(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTa(ctx)),
	)
}

func TestTaCount(t *testing.T) {
	keeper, ctx := keepertest.ObitKeeper(t)
	items := createNTa(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTaCount(ctx))
}
