package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	localNfts := make([]types.NFT, 0)

	for _, n := range nfts {
		// TODO: check why we have empty duplicates
		if len(n.Id) > 0 {
			localNfts = append(localNfts, types.NFT{
				ClassId: n.ClassId,
				Id:      n.Id,
				Uri:     n.Uri,
				UriHash: n.UriHash,
				Data:    n.Data,
			})
		}
	}

	return &types.QueryGetAllNftByOwnerResponse{NFT: localNfts}, nil
}
