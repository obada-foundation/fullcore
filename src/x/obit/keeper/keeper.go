package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc       codec.BinaryCodec
		storeKey  storetypes.StoreKey
		memKey    storetypes.StoreKey
		nftKeeper types.NftKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	nftKeeper types.NftKeeper,
) *Keeper {
	return &Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		memKey:    memKey,
		nftKeeper: nftKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SaveClass saves a class to the store
func (k Keeper) SaveClass(ctx sdk.Context, class types.Class) error {
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
