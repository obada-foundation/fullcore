package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/obada-foundation/fullcore/x/fullcore/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

func (k Keeper) ObitAll(c context.Context, req *types.QueryAllObitRequest) (*types.QueryAllObitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var obits []*types.Obit
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	postStore := prefix.NewStore(store, types.KeyPrefix(types.ObitKey))

	pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
		var obit types.Obit
		if err := k.cdc.Unmarshal(value, &obit); err != nil {
			return err
		}

		obits = append(obits, &obit)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllObitResponse{Obit: obits, Pagination: pageRes}, nil
}

func (k Keeper) GetObitCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObitCountKey))
	byteKey := types.KeyPrefix(types.ObitCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

func (k Keeper) SetObitCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObitCountKey))
	byteKey := types.KeyPrefix(types.ObitCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) AppendObit(ctx sdk.Context, obit types.Obit) uint64 {
	count := k.GetObitCount(ctx)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObitKey))
	key := types.KeyPrefix(types.ObitKey + obit.Did)
	value := k.cdc.MustMarshal(&obit)
	store.Set(key, value)

	// Update comment count
	k.SetObitCount(ctx, count+1)

	return count
}

func (k Keeper) GetObit(ctx sdk.Context, key string) types.Obit {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObitKey))
	var obit types.Obit
	k.cdc.MustUnmarshal(store.Get(types.KeyPrefix(types.ObitKey+key)), &obit)
	return obit
}

func (k Keeper) HasObit(ctx sdk.Context, did string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObitKey))
	return store.Has(types.KeyPrefix(types.ObitKey + did))
}

func (k Keeper) GetObitOwner(ctx sdk.Context, key string) string {
	return k.GetObit(ctx, key).Owner
}

func (k Keeper) GetAllObit(ctx sdk.Context) (msgs []types.Obit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObitKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ObitKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Obit
		k.cdc.MustUnmarshal(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
