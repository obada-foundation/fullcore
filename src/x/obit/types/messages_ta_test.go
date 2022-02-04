package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/obada-foundation/fullcore/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTa_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTa
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTa{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTa{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateTa_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTa
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTa{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTa{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteTa_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTa
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTa{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTa{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
