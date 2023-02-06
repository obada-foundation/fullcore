package keeper

import (
	"context"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func (k msgServer) UpdateNFT(goCtx context.Context, msg *types.MsgUpdateNFT) (*types.MsgUpdateNFTResponse, error) {
	resp := &types.MsgUpdateNFTResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	nft, ok := k.nftKeeper.GetNFT(ctx, types.OBTClass, msg.Id)
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

	nft.Data = data
	if err := k.nftKeeper.Update(ctx, nft); err != nil {
		return resp, err
	}

	resp.NFT = (*types.NFT)(&nft)

	return resp, nil
}

func (k msgServer) TransferNFT(goCtx context.Context, msg *types.MsgTransferNFT) (*types.MsgTransferNFTResponse, error) {
	resp := &types.MsgTransferNFTResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.nftKeeper.Transfer(ctx, types.OBTClass, msg.Id, sdk.AccAddress(msg.Receiver)); err != nil {
		return resp, err
	}

	return resp, nil
}

func (k msgServer) MintNFT(goCtx context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	resp := &types.MsgMintNFTResponse{}

	// Find a right place for provisioning NFT class
	if !k.nftKeeper.HasClass(ctx, types.OBTClass) {
		err := k.nftKeeper.SaveClass(ctx, nft.Class{
			Id:     types.OBTClass,
			Name:   "Obada network NFT Token",
			Symbol: types.OBTClass,
			Uri:    "https://obada.io",
		})
		if err != nil {
			return resp, err
		}
	}

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
