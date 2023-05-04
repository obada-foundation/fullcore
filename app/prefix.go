package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AccountAddressPrefix defines the prefix of account address
	AccountAddressPrefix = "obada"
)

var (
	// AccountPubKeyPrefix defines the prefix of account public key
	AccountPubKeyPrefix = AccountAddressPrefix + "pub"

	// ValidatorAddressPrefix defines the prefix of validator address
	ValidatorAddressPrefix = AccountAddressPrefix + "valoper"

	// ValidatorPubKeyPrefix defines the prefix of validator public key
	ValidatorPubKeyPrefix = AccountAddressPrefix + "valoperpub"

	// ConsNodeAddressPrefix defines the prefix of consensus node address
	ConsNodeAddressPrefix = AccountAddressPrefix + "valcons"

	// ConsNodePubKeyPrefix defines the prefix of consensus node public key
	ConsNodePubKeyPrefix = AccountAddressPrefix + "valconspub"
)

// SetConfig sets the global sdk config.
func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}
