package obit

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/obada-foundation/fullcore/testutil/sample"
	obitsimulation "github.com/obada-foundation/fullcore/x/obit/simulation"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = obitsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgMintObit = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintObit int = 100

	opWeightMsgCreateTa = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTa int = 100

	opWeightMsgUpdateTa = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTa int = 100

	opWeightMsgDeleteTa = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTa int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	obitGenesis := types.GenesisState{
		TaList: []types.Ta{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TaCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&obitGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMintObit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintObit, &weightMsgMintObit, nil,
		func(_ *rand.Rand) {
			weightMsgMintObit = defaultWeightMsgMintObit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintObit,
		obitsimulation.SimulateMsgMintObit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTa int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTa, &weightMsgCreateTa, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTa = defaultWeightMsgCreateTa
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTa,
		obitsimulation.SimulateMsgCreateTa(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTa int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTa, &weightMsgUpdateTa, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTa = defaultWeightMsgUpdateTa
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTa,
		obitsimulation.SimulateMsgUpdateTa(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTa int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTa, &weightMsgDeleteTa, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTa = defaultWeightMsgDeleteTa
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTa,
		obitsimulation.SimulateMsgDeleteTa(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
