package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintObit = "mint_obit"

var _ sdk.Msg = &MsgMintObit{}

func NewMsgMintObit(creator, serialNumberHash, manufacturer, pn, ta, ownerDID string) *MsgMintObit {
	return &MsgMintObit{
		Creator:          creator,
		SerialNumberHash: serialNumberHash,
		Manufacturer:     manufacturer,
		PartNumber:       pn,
		TrustAnchor:      ta,
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
