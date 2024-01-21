package keeper

import (
	store "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

type (
	// Keeper defines the obit module's keeper. It is responsible for managing
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		nk           types.NftKeeper
	}
)

// NewKeeper constructor for Keeper
func NewKeeper(storeService store.KVStoreService, cdc codec.BinaryCodec, nftKeeper types.NftKeeper) *Keeper {
	return &Keeper{
		cdc:          cdc,
		storeService: storeService,
		nk:           nftKeeper,
	}
}
