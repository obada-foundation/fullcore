package obit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/obada-foundation/fullcore/x/obit/keeper"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the ta
	for _, elem := range genState.TaList {
		k.SetTa(ctx, elem)
	}

	// Set ta count
	k.SetTaCount(ctx, genState.TaCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TaList = k.GetAllTa(ctx)
	genesis.TaCount = k.GetTaCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
