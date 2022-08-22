package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

func (k msgServer) CreateTa(goCtx context.Context, msg *types.MsgCreateTa) (*types.MsgCreateTaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var ta = types.Ta{
		Creator: msg.Creator,
		Title:   msg.Title,
		Pubkey:  msg.Pubkey,
	}

	id := k.AppendTa(
		ctx,
		ta,
	)

	return &types.MsgCreateTaResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTa(goCtx context.Context, msg *types.MsgUpdateTa) (*types.MsgUpdateTaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var ta = types.Ta{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Pubkey:  msg.Pubkey,
	}

	// Checks that the element exists
	val, found := k.GetTa(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTa(ctx, ta)

	return &types.MsgUpdateTaResponse{}, nil
}

func (k msgServer) DeleteTa(goCtx context.Context, msg *types.MsgDeleteTa) (*types.MsgDeleteTaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTa(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTa(ctx, msg.Id)

	return &types.MsgDeleteTaResponse{}, nil
}
