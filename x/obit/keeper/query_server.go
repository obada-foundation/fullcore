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

var _ types.QueryServer = queryServer{}

type queryServer struct {
	k Keeper
}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

// GetNFT implements gRPC query handler method
func (qs queryServer) GetNFT(ctx context.Context, req *types.QueryGetNFTRequest) (*types.QueryGetNFTResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	did, err := url.QueryUnescape(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	nft, ok := qs.k.nftKeeper.GetNFT(ctx, types.OBTClass, did)
	if !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no %s with DID %s found", "OBT", did))
	}

	obt := types.NFT(nft)

	return &types.QueryGetNFTResponse{
		Nft: obt,
	}, nil

}

// GetNFTByAddress implements gRPC query handler method
func (qs queryServer) GetNFTByAddress(c context.Context, req *types.QueryGetNFTByAddressRequest) (*types.QueryGetNFTByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	nfts := qs.k.nftKeeper.GetNFTsOfClassByOwner(ctx, "OBT", sdk.AccAddress(req.Address))

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

	return &types.QueryGetNFTByAddressResponse{Nft: localNfts}, nil
}
