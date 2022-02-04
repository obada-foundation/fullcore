package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/obada-foundation/fullcore/x/nft"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func (k msgServer) MintObit(goCtx context.Context, msg *types.MsgMintObit) (*types.MsgMintObitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	obit := types.Obit{
		Did: msg.Did,
	}

	if !k.nftKeeper.HasClass(ctx, "OBT") {
		k.nftKeeper.SaveClass(ctx, nft.Class{
			Id:     "OBT",
			Name:   "Obada network NFT Token",
			Symbol: "OBT",
			Uri:    "https://obada.io",
		})
	}

	nftToken := nft.NFT{
		ClassId: "OBT",
		Id:      obit.Did,
		Uri:     "http://google.com",
		UriHash: "",
		Data:    &codectypes.Any{},
	}

	err := k.nftKeeper.Mint(ctx, nftToken, sdk.AccAddress(msg.Creator))

	if err != nil {
		return &types.MsgMintObitResponse{}, err
	}

	return &types.MsgMintObitResponse{
		Did: obd.Id,
	}, nil
}
