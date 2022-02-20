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

	if err := k.nftKeeper.Transfer(ctx, types.OBTClass, msg.Did, sdk.AccAddress(msg.Receiver)); err != nil {
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

	docs := make([]map[string]string, len(msg.Documents))
	for _, doc := range msg.Documents {
		d := make(map[string]string)
		d["name"] = doc.Name

		docs = append(docs, map[string]string{
			"name": doc.Name,
			"hash": doc.Hash,
			"uri":  doc.Uri,
		})
	}

	obt, err := osdk.NewObit(obadasdk.ObitDto{
		ObitIDDto: obadasdk.ObitIDDto{
			SerialNumberHash: msg.SerialNumberHash,
			Manufacturer:     msg.Manufacturer,
			PartNumber:       msg.PartNumber,
		},
		TrustAnchorToken: msg.TrustAnchorToken,
		Documents:        docs,
	})

	if err != nil {
		return resp, err
	}

	// Find a right place for provisioning NFT class
	if !k.nftKeeper.HasClass(ctx, types.OBTClass) {
		k.nftKeeper.SaveClass(ctx, nft.Class{
			Id:     types.OBTClass,
			Name:   "Obada network NFT Token",
			Symbol: types.OBTClass,
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
		ClassId: types.OBTClass,
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
