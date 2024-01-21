package types

import (
	"cosmossdk.io/x/nft"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// NftKeeper defines the expected nft keeper
type NftKeeper interface {
	// Methods imported from nft should be defined here
	Update(ctx sdk.Context, token nft.NFT) error

	// GetOwner returns the owner of the specified NFT
	GetOwner(ctx sdk.Context, classID, nftID string) sdk.AccAddress

	// Mint mints a new NFT
	Mint(ctx sdk.Context, token nft.NFT, receiver sdk.AccAddress) error

	// Burn burns a NFT
	Transfer(ctx sdk.Context, classID string, nftID string, receiver sdk.AccAddress) error

	// GetNFTs returns all NFTs
	GetNFTsOfClassByOwner(ctx sdk.Context, classID string, owner sdk.AccAddress) (nfts []nft.NFT)

	// GetNFT returns the specified NFT
	GetNFT(ctx sdk.Context, classID, nftID string) (nft.NFT, bool)

	// GetClass returns the specified class
	SaveClass(ctx sdk.Context, class nft.Class) error

	// GetClass returns the specified class
	GetClasses(ctx sdk.Context) (classes []*nft.Class)

	// HasClass returns whether the specified class exists
	HasClass(ctx sdk.Context, classID string) bool
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}
