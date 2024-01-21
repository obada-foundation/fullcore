package keeper

import (
	store "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// Keeper of the x/obit store
type Keeper struct {
	cdc          codec.BinaryCodec
	storeService store.KVStoreService
	nk           types.NftKeeper
}
