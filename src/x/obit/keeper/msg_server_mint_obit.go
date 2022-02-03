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

	obd := nft.NFT{
		ClassId: "OBD",
		Id:      obit.Did,
		Uri:     "",
		UriHash: "",
		Data:    &codectypes.Any{},
	}

	logger := k.Logger(ctx)

	logger.Info("NFT before minting", obd)

	err := k.nftKeeper.Mint(ctx, obd, sdk.AccAddress(msg.Creator))

	if err != nil {
		return nil, err
	}

	logger.Info("NFT before minting", obd, err)

	return &types.MsgMintObitResponse{
		Did: obd.Id,
	}, nil
}
