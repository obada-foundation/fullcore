package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TaAll(c context.Context, req *types.QueryAllTaRequest) (*types.QueryAllTaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tas []types.Ta
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	taStore := prefix.NewStore(store, types.KeyPrefix(types.TaKey))

	pageRes, err := query.Paginate(taStore, req.Pagination, func(key []byte, value []byte) error {
		var ta types.Ta
		if err := k.cdc.Unmarshal(value, &ta); err != nil {
			return err
		}

		tas = append(tas, ta)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTaResponse{Ta: tas, Pagination: pageRes}, nil
}

func (k Keeper) Ta(c context.Context, req *types.QueryGetTaRequest) (*types.QueryGetTaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	ta, found := k.GetTa(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTaResponse{Ta: ta}, nil
}
