package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func (k msgServer) UpdateNFT(goCtx context.Context, msg *types.MsgUpdateNFT) (*types.MsgUpdateNFTResponse, error) {
	resp := &types.MsgUpdateNFTResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	NFT, ok := k.nftKeeper.GetNFT(ctx, types.OBTClass, msg.Id)
	if !ok {
		return resp, fmt.Errorf("NFT %s doesn't exists", msg.Id)
	}

	ownerAddress := k.nftKeeper.GetOwner(ctx, types.OBTClass, msg.Id)
	if !ownerAddress.Equals(sdk.AccAddress(msg.Editor)) {
		return resp, fmt.Errorf("permission denied")
	}

	data, err := codectypes.NewAnyWithValue(msg.NFTData)
	if err != nil {
		return resp, err
	}

	NFT.Data = data
	if err := k.nftKeeper.Update(ctx, NFT); err != nil {
		return resp, err
	}

	resp.NFT = (*types.NFT)(&NFT)

	return resp, nil
}

func (k msgServer) TransferNFT(goCtx context.Context, msg *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	resp := &types.MsgTransferNFTResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Sender == msg.Receiver {
		return resp, errors.Wrap(errortypes.ErrInvalidAddress, "sender and receiver cannot be the same")
	}

	if err := k.nftKeeper.Transfer(ctx, types.OBTClass, msg.Id, sdk.AccAddress(msg.Receiver)); err != nil {
		return resp, err
	}

	return resp, nil
}

func (k msgServer) MintNFT(goCtx context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
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

	if err := k.nftKeeper.Mint(ctx, nftToken, sdk.AccAddress(msg.Creator)); err != nil {
		return resp, err
	}

	return &types.MsgMintNFTResponse{
		Id: msg.Id,
	}, nil
}

// UpdateUriHash updates the URI hash of an NFT
func (k msgServer) UpdateUriHash(goCtx context.Context, msg *types.MsgUpdateUriHash) (*types.MsgUpdateUriHashResponse, error) {
	resp := &types.MsgUpdateUriHashResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	NFT, ok := k.nftKeeper.GetNFT(ctx, types.OBTClass, msg.Id)
	if !ok {
		return resp, fmt.Errorf("NFT %s doesn't exists", msg.Id)
	}

	ownerAddress := k.nftKeeper.GetOwner(ctx, types.OBTClass, msg.Id)
	if !ownerAddress.Equals(sdk.AccAddress(msg.Editor)) {
		return resp, fmt.Errorf("permission denied")
	}

	NFT.UriHash = msg.UriHash

	if err := k.nftKeeper.Update(ctx, NFT); err != nil {
		return resp, err
	}

	return resp, nil
}
