package keeper

import (
	"cosmossdk.io/x/nft"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// InitGenesis initializes the capability module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	for _, class := range genState.Classes {
		if err := k.SaveClass(ctx, *class); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	genState := types.GenesisState{}

	classes := k.nk.GetClasses(ctx)

	for _, nftClass := range classes {
		genState.Classes = append(genState.Classes, &types.Class{
			Id:     nftClass.Id,
			Name:   nftClass.Name,
			Symbol: nftClass.Symbol,
			Uri:    nftClass.Uri,
		})
	}

	return &genState
}

// SaveClass saves a class to the store
func (k Keeper) SaveClass(ctx sdk.Context, class types.Class) error {
	if !k.nk.HasClass(ctx, class.Id) {
		return k.nk.SaveClass(ctx, nft.Class{
			Id:     class.Id,
			Name:   class.Name,
			Symbol: class.Symbol,
			Uri:    class.Uri,
		})
	}

	return nil
}
