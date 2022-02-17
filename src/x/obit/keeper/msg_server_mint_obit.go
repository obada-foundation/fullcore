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

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	resp := &types.MsgSendResponse{}

	ctx := sdk.UnwrapSDKContext(goCtx)

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return resp, err
	}

	if err := k.nftKeeper.Transfer(ctx, "OBT", msg.Did, receiver); err != nil {
		return resp, err
	}

	return resp, nil
}

func (k msgServer) MintObit(goCtx context.Context, msg *types.MsgMintObit) (*types.MsgMintObitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	resp := &types.MsgMintObitResponse{}

	logger := log.New(os.Stdout, " :: OBADA SDK ::", 0)

	osdk, err := obadasdk.NewSdk(logger, true)
	if err != nil {
		return resp, err
	}

	// Validate trust anchor
	if len(msg.TrustAnchorToken) > 0 {
		isCompliant, err := k.CheckCompliance(msg.TrustAnchorToken, msg.TrustAnchorToken)
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
		TrustAnchorToken: msg.TrustAnchorToken,
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

	checksum, err := obt.GetChecksum(nil)
	if err != nil {
		return resp, err
	}

	// check URI hash
	data, err := codectypes.NewAnyWithValue(&types.NFTData{
		Documents:        msg.Documents,
		TrustAnchorToken: msg.TrustAnchorToken,
		Checksum:         checksum.GetHash(),
		Usn:              did.GetUsn(),
	})
	if err != nil {
		return resp, err
	}

	nftToken := nft.NFT{
		ClassId: "OBT",
		Id:      did.GetDid(),
		Uri:     msg.Uri,
		UriHash: msg.UriHash,
		Data:    data,
	}

	if err := k.nftKeeper.Mint(ctx, nftToken, sdk.AccAddress(msg.Creator)); err != nil {
		return resp, err
	}

	return &types.MsgMintObitResponse{
		Did: did.GetDid(),
	}, nil
}
