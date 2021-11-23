package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang-jwt/jwt"
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

	_, err := jwt.ParseRSAPublicKeyFromPEM([]byte("-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAueET5LUA849k4Z7IB0YP\nLj/aEsH0TBU6SHiW0FE6sA9ZsTHtmJhZiL7T4h7tw5HAO9HwU1BdbeaelyxJ6cBu\nBruV8jTjKa5eqXlyyJSzNGYj2EQx8LI/lExlZ0+m24vk5M45mdOn2O2XmZ6BLXy3\nmYVn+FzA5zkibNVR5lIuJVyfPQty9ThP1GMxtKTn5rBGWaRSYHu/UH/YCowJpKCf\nIh9QIyOtFyBU9dTbhdEVUNmy0TzouR0skZeT3j5LaqWoIWR9Z2YdRb3PCRmIouMT\ndZai+iPR9IFdvdPrciG1fxBNF2nGY6cn33zc5/JatDIYIRBP/rAOREKA55fxrdW7\n+stkCYWgdSsQnP3CHaCeGW9qfaOJZVd0AclVWL5DUCfVGpWL38AXKjBQn9nS8P9a\n+FjaDLWqmHduZoSlVC0ZECP0Fp5AjWsaD4Fow0/wDHViiezhTGYJWW2sgg3pMmJ1\njtPdbZdaYJv+de27xiT+q7sabT4lWDJaZ0/uMAFAQqwGZ+WfRHaCmeK7g0gS+iTn\nrMp7kNZ2GZCvW5BiuIgKB5baM+nEZVr/dWVhhwcB/d45Fq4Vvk9me9b6w/qNVGP7\nfoEIwu9RZLStTEHgCY/ABrGN9H+/DwT1wNEgDcfh/n8KwXXs/LVk2VSbkT40mne9\nJpqjxiSsrj6SiHl0EcUc9nMCAwEAAQ==\n-----END PUBLIC KEY-----"))
	if err != nil {
		return err
	}

	token, err := jwt.Parse(sig, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		pubkey, err := jwt.ParseRSAPublicKeyFromPEM([]byte("-----BEGIN PUBLIC KEY-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAueET5LUA849k4Z7IB0YP\nLj/aEsH0TBU6SHiW0FE6sA9ZsTHtmJhZiL7T4h7tw5HAO9HwU1BdbeaelyxJ6cBu\nBruV8jTjKa5eqXlyyJSzNGYj2EQx8LI/lExlZ0+m24vk5M45mdOn2O2XmZ6BLXy3\nmYVn+FzA5zkibNVR5lIuJVyfPQty9ThP1GMxtKTn5rBGWaRSYHu/UH/YCowJpKCf\nIh9QIyOtFyBU9dTbhdEVUNmy0TzouR0skZeT3j5LaqWoIWR9Z2YdRb3PCRmIouMT\ndZai+iPR9IFdvdPrciG1fxBNF2nGY6cn33zc5/JatDIYIRBP/rAOREKA55fxrdW7\n+stkCYWgdSsQnP3CHaCeGW9qfaOJZVd0AclVWL5DUCfVGpWL38AXKjBQn9nS8P9a\n+FjaDLWqmHduZoSlVC0ZECP0Fp5AjWsaD4Fow0/wDHViiezhTGYJWW2sgg3pMmJ1\njtPdbZdaYJv+de27xiT+q7sabT4lWDJaZ0/uMAFAQqwGZ+WfRHaCmeK7g0gS+iTn\nrMp7kNZ2GZCvW5BiuIgKB5baM+nEZVr/dWVhhwcB/d45Fq4Vvk9me9b6w/qNVGP7\nfoEIwu9RZLStTEHgCY/ABrGN9H+/DwT1wNEgDcfh/n8KwXXs/LVk2VSbkT40mne9\nJpqjxiSsrj6SiHl0EcUc9nMCAwEAAQ==\n-----END PUBLIC KEY-----"))
		if err != nil {
			return nil, err
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return pubkey, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	}

	return errors.New(sig)
}
