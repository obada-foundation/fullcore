package keeper

import (
	"context"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/obada-foundation/fullcore/x/fullcore/types"
	"github.com/obada-foundation/sdkgo"
	"log"
	"os"
)

func (k msgServer) CreateObit(goCtx context.Context, msg *types.MsgCreateObit) (*types.MsgCreateObitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	dto := sdkgo.ObitIDDto{
		SerialNumberHash: msg.SerialNumberHash,
		PartNumber:       msg.PartNumber,
		Manufacturer:     msg.Manufacturer,
	}

	logger := log.New(os.Stdout, "", 0)

	obadaSdk, err := sdkgo.NewSdk(logger, false)
	if err != nil {
		return nil, err
	}

	obitID, err := obadaSdk.NewObitID(dto)
	if err != nil {
		return nil, err
	}

	DID := obitID.GetDid()

	getObit := k.GetObit(ctx, DID)

	if getObit.Did == DID {
		return nil, errors.New("such DID is already exists")
	}

	var obit = types.Obit{
		Did:              DID,
		Creator:          msg.Creator,
		Owner:            msg.Creator,
		SerialNumberHash: msg.SerialNumberHash,
		PartNumber:       msg.PartNumber,
		Manufacturer:     msg.PartNumber,
		Signature:        msg.Signature,
	}

	k.AppendObit(ctx, obit)

	return &types.MsgCreateObitResponse{
		Obit: &obit,
	}, nil
}
