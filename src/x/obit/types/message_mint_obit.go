package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgMintObit = "mint_obit"
	TypeMsgSend     = "send_obit"
)

var _ sdk.Msg = &MsgMintObit{}

func NewMsgMintObit(creator, serialNumberHash, manufacturer, pn, obdDid, ownerDID string) *MsgMintObit {
	return &MsgMintObit{
		Creator:          creator,
		SerialNumberHash: serialNumberHash,
		Manufacturer:     manufacturer,
		PartNumber:       pn,
		ObdDid:           obdDid,
		OwnerDid:         ownerDID,
	}
}

func (msg *MsgMintObit) Route() string {
	return RouterKey
}

func (msg *MsgMintObit) Type() string {
	return TypeMsgMintObit
}

func (msg *MsgMintObit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintObit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintObit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgSend{}

func NewMsgSend(DID, sender, receiver string) *MsgSend {
	return &MsgSend{
		Did:      DID,
		Sender:   sender,
		Receiver: receiver,
	}
}

func (msg *MsgSend) Route() string {
	return RouterKey
}

func (msg *MsgSend) Type() string {
	return TypeMsgSend
}

func (msg *MsgSend) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}
	return nil
}
