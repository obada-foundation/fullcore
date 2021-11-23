package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

var _ sdk.Msg = &MsgCreateObit{}

func NewMsgCreateObit(creator, serialNumberHash, manufacturer, partNumber, sig string) *MsgCreateObit {
	return &MsgCreateObit{
		Creator:          creator,
		Owner:            creator,
		SerialNumberHash: serialNumberHash,
		Manufacturer:     manufacturer,
		PartNumber:       partNumber,
		Signature:        sig,
	}
}

func (msg MsgCreateObit) Route() string {
	return RouterKey
}

func (msg MsgCreateObit) Type() string {
	return "CreateObit"
}

func (msg *MsgCreateObit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateObit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateObit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err := ValidateSignature(msg.Signature); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid signature (%s)", err)
	}

	return nil
}

func ValidateSignature(sig string) error {
	return errors.New(sig)
}
