package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/obada-foundation/fullcore/testutil/keeper"
	"github.com/obada-foundation/fullcore/x/obit/keeper"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ObitKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}