package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	moduletestutil "github.com/obada-foundation/fullcore/testutil/module/codec"
	"github.com/obada-foundation/fullcore/x/obit"
	"github.com/obada-foundation/fullcore/x/obit/simulation"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/stretchr/testify/require"
)

// RandomizedGenState generates a random GenesisState for nft.
func TestRandomizedGenState(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(obit.AppModuleBasic{})

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          encCfg.Codec,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var obitGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &obitGenesis)

	require.Len(t, obitGenesis.Classes, len(simState.Accounts)-1)
}
