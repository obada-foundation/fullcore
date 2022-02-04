package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/obada-foundation/fullcore/x/nft"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAllNftByOwner(c context.Context, req *types.QueryGetAllNftByOwnerRequest) (*types.QueryGetAllNftByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	nfts := k.nftKeeper.GetNFTsOfClassByOwner(ctx, "OBT", sdk.AccAddress(req.Owner))

	return &types.QueryGetAllNftByOwnerResponse{Nft: nfts}, nil
}
