package keeper

import (
	"context"

	store "cosmossdk.io/core/store"
	"cosmossdk.io/x/nft"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// Keeper of the x/obit store
type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	nftKeeper    types.NftKeeper
}

// NewKeeper constructor for Keeper
func NewKeeper(cdc codec.BinaryCodec, sk store.KVStoreService, nk types.NftKeeper) Keeper {
	return Keeper{
		cdc:          cdc,
		storeService: sk,
		nftKeeper:    nk,
	}
}

// SaveClass saves a class to the store
func (k Keeper) SaveClass(ctx context.Context, class types.Class) error {
	if !k.nftKeeper.HasClass(ctx, class.Id) {
		return k.nftKeeper.SaveClass(ctx, nft.Class{
			Id:     class.Id,
			Name:   class.Name,
			Symbol: class.Symbol,
			Uri:    class.Uri,
		})
	}

	return nil
}
