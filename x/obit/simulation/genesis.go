package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// genClasses returns a slice of nft class.
func genClasses(r *rand.Rand, accounts []simtypes.Account) []*types.Class {
	classes := make([]*types.Class, len(accounts)-1)
	for i := 0; i < len(accounts)-1; i++ {
		classes[i] = &types.Class{
			Id:          simtypes.RandStringOfLength(r, 10),
			Name:        simtypes.RandStringOfLength(r, 10),
			Symbol:      simtypes.RandStringOfLength(r, 10),
			Description: simtypes.RandStringOfLength(r, 10),
			Uri:         simtypes.RandStringOfLength(r, 10),
		}
	}
	return classes
}

// RandomizedGenState generates a random GenesisState for nft.
func RandomizedGenState(simState *module.SimulationState) {
	var classes []*types.Class
	simState.AppParams.GetOrGenerate(
		simState.Cdc, "obit", &classes, simState.Rand,
		func(r *rand.Rand) { classes = genClasses(r, simState.Accounts) },
	)

	obitGenesis := &types.GenesisState{
		Classes: classes,
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(obitGenesis)
}
