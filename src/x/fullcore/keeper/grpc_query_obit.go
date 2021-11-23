package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/obada-foundation/fullcore/x/fullcore/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Obit(c context.Context, req *types.QueryGetObitRequest) (*types.QueryGetObitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var obit types.Obit
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObitKey))
	k.cdc.MustUnmarshal(store.Get(types.KeyPrefix(types.ObitKey+req.Did)), &obit)

	return &types.QueryGetObitResponse{Obit: &obit}, nil
}
