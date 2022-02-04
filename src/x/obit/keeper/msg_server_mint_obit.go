package keeper

import (
	"context"
	"errors"
	"log"
	"os"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	obadasdk "github.com/obada-foundation/sdkgo"

	"github.com/obada-foundation/fullcore/x/nft"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func (k msgServer) MintObit(goCtx context.Context, msg *types.MsgMintObit) (*types.MsgMintObitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	resp := &types.MsgMintObitResponse{}

	logger := log.New(os.Stdout, " :: OBADA SDK ::", 0)

	osdk, err := obadasdk.NewSdk(logger, true)
	if err != nil {
		return resp, err
	}

	// Validate trust anchor
	if len(msg.OwnerDid) > 0 {
		isCompliant, err := k.CheckCompliance(msg.TrustAnchor, msg.OwnerDid)
		if err != nil {
			return resp, err
		}

		if !isCompliant {
			return resp, errors.New("owner is not compliant")
		}
	}

	obt, err := osdk.NewObit(obadasdk.ObitDto{
		ObitIDDto: obadasdk.ObitIDDto{
			SerialNumberHash: msg.SerialNumberHash,
			Manufacturer:     msg.Manufacturer,
			PartNumber:       msg.PartNumber,
		},
		OwnerDid: msg.OwnerDid,
	})

	if err != nil {
		return resp, err
	}

	// Find a right place for provisioning NFT class
	if !k.nftKeeper.HasClass(ctx, "OBT") {
		k.nftKeeper.SaveClass(ctx, nft.Class{
			Id:     "OBT",
			Name:   "Obada network NFT Token",
			Symbol: "OBT",
			Uri:    "https://obada.io",
		})
	}

	did := obt.GetObitID()

	// check URI hash

	nftToken := nft.NFT{
		ClassId: "OBT",
		Id:      did.GetDid(),
		Uri:     "http://google.com",
		UriHash: "",
		Data:    &codectypes.Any{},
	}

	if err := k.nftKeeper.Mint(ctx, nftToken, sdk.AccAddress(msg.Creator)); err != nil {
		return resp, err
	}

	return &types.MsgMintObitResponse{
		Did: obt.GetObdDID().GetValue(),
	}, nil
}
