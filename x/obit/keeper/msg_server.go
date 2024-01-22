package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	"cosmossdk.io/x/nft"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) UpdateNFT(goCtx context.Context, msg *types.MsgUpdateNFT) (*types.MsgUpdateNFTResponse, error) {
	resp := &types.MsgUpdateNFTResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	NFT, ok := ms.k.nftKeeper.GetNFT(ctx, types.OBTClass, msg.Id)
	if !ok {
		return resp, fmt.Errorf("NFT %s doesn't exists", msg.Id)
	}

	ownerAddress := ms.k.nftKeeper.GetOwner(ctx, types.OBTClass, msg.Id)
	if !ownerAddress.Equals(sdk.AccAddress(msg.Editor)) {
		return resp, fmt.Errorf("permission denied")
	}

	data, err := codectypes.NewAnyWithValue(msg.NftData)
	if err != nil {
		return resp, err
	}

	NFT.Data = data
	if err := ms.k.nftKeeper.Update(ctx, NFT); err != nil {
		return resp, err
	}

	resp.Nft = (*types.NFT)(&NFT)

	return resp, nil
}

func (ms msgServer) TransferNFT(goCtx context.Context, msg *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	resp := &types.MsgTransferNFTResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Sender == msg.Receiver {
		return resp, errors.Wrap(errortypes.ErrInvalidAddress, "sender and receiver cannot be the same")
	}

	if err := ms.k.nftKeeper.Transfer(ctx, types.OBTClass, msg.Id, sdk.AccAddress(msg.Receiver)); err != nil {
		return resp, err
	}

	return resp, nil
}

// BatchTransferNFT batch transfers NFTs
func (ms msgServer) BatchMintNFT(ctx context.Context, msg *types.MsgBatchMintNFT) (*types.MsgBatchMintNFTResponse, error) {
	batchNfts := make([]nft.NFT, 0, len(msg.GetNft()))

	for _, n := range msg.GetNft() {
		data, err := codectypes.NewAnyWithValue(&types.NFTData{
			Usn: n.Usn,
		})
		if err != nil {
			return nil, err
		}

		nftToken := nft.NFT{
			ClassId: types.OBTClass,
			Id:      n.Id,
			Uri:     n.Uri,
			UriHash: n.UriHash,
			Data:    data,
		}

		batchNfts = append(batchNfts, nftToken)

	}

	if err := ms.k.nftKeeper.BatchMint(ctx, batchNfts, sdk.AccAddress(msg.GetCreator())); err != nil {
		return nil, err
	}

	return &types.MsgBatchMintNFTResponse{}, nil
}

// MintNFT mints an NFT
func (ms msgServer) MintNFT(goCtx context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	resp := &types.MsgMintNFTResponse{}

	// check URI hash
	data, err := codectypes.NewAnyWithValue(&types.NFTData{
		Usn: msg.Usn,
	})
	if err != nil {
		return resp, err
	}

	nftToken := nft.NFT{
		ClassId: types.OBTClass,
		Id:      msg.Id,
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		Data:    data,
	}

	if err := ms.k.nftKeeper.Mint(ctx, nftToken, sdk.AccAddress(msg.Creator)); err != nil {
		return resp, err
	}

	return &types.MsgMintNFTResponse{
		Id: msg.Id,
	}, nil
}

// UpdateUriHash updates the URI hash of an NFT
func (ms msgServer) UpdateUriHash(goCtx context.Context, msg *types.MsgUpdateUriHash) (*types.MsgUpdateUriHashResponse, error) {
	resp := &types.MsgUpdateUriHashResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	NFT, ok := ms.k.nftKeeper.GetNFT(ctx, types.OBTClass, msg.Id)
	if !ok {
		return resp, fmt.Errorf("NFT %s doesn't exists", msg.Id)
	}

	ownerAddress := ms.k.nftKeeper.GetOwner(ctx, types.OBTClass, msg.Id)
	if !ownerAddress.Equals(sdk.AccAddress(msg.Editor)) {
		return resp, fmt.Errorf("permission denied")
	}

	NFT.UriHash = msg.UriHash

	if err := ms.k.nftKeeper.Update(ctx, NFT); err != nil {
		return resp, err
	}

	return resp, nil
}
