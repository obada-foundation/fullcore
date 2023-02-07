package keeper

import (
	"context"
	"fmt"
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetNFT(c context.Context, req *types.QueryGetNFTRequest) (*types.NFT, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	did, err := url.QueryUnescape(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	nft, ok := k.nftKeeper.GetNFT(ctx, types.OBTClass, did)
	if !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no %s with DID %s found", "OBT", did))
	}

	obt := types.NFT(nft)

	return &obt, nil

}

func (k Keeper) GetNFTByAddress(c context.Context, req *types.QueryGetAllNFTByAddressRequest) (*types.QueryGetAllNFTByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	nfts := k.nftKeeper.GetNFTsOfClassByOwner(ctx, "OBT", sdk.AccAddress(req.Address))

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

	return &types.QueryGetAllNFTByAddressResponse{NFT: localNfts}, nil
}
