package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	nft "github.com/obada-foundation/fullcore/x/nft"
)

type NftKeeper interface {
	Mint(ctx sdk.Context, token nft.NFT, receiver sdk.AccAddress) error

	Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error

	GetNFTsOfClassByOwner(ctx sdk.Context, classID string, owner sdk.AccAddress) (nfts []nft.NFT)

	SaveClass(ctx sdk.Context, class nft.Class) error

	HasClass(ctx sdk.Context, classID string) bool
	// Methods imported from nft should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
