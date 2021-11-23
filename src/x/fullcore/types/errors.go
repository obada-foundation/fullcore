package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/core module sentinel errors
var (
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidSignature = sdkerrors.Register(ModuleName, 1101, "invalid registry signature")
)
