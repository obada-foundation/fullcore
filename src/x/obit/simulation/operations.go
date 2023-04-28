package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	simtestutil "github.com/obada-foundation/fullcore/testutil/sims"
	"github.com/obada-foundation/fullcore/x/obit/keeper"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

const (
	// OpWeightMsgMintNFT Simulation operation weights constants
	OpWeightMsgMintNFT = "op_weight_msg_mint_nft" //nolint:gosec

	// WeightMsgMintNFT operations weights
	WeightMsgMintNFT = 100
)

var TypeMsgMintNFT = sdk.MsgTypeURL(&types.MsgMintNFT{})

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	registry cdctypes.InterfaceRegistry,
	appParams simtypes.AppParams,
	cdc codec.JSONCodec,
	ak nft.AccountKeeper,
	bk nft.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var weightMsgMintNFT int

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgMintNFT,
			SimulateMsgMintNFT(codec.NewProtoCodec(registry), ak, bk, k),
		),
	}
}

// SimulateMsgMintNFT mints an NFTs with random values.
func SimulateMsgMintNFT(
	cdc *codec.ProtoCodec,
	ak nft.AccountKeeper,
	bk nft.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		creator, _ := simtypes.RandomAcc(r, accs)

		minterAcc := ak.GetAccount(ctx, creator.Address)

		spendableCoins := bk.SpendableCoins(ctx, creator.Address)

		fees, err := simtypes.RandomFees(r, ctx, spendableCoins)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgMintNFT, err.Error()), nil, err
		}

		msg := &types.MsgMintNFT{
			Creator: minterAcc.GetAddress().String(),
		}

		txCfg := tx.NewTxConfig(cdc, tx.DefaultSignModes)
		tx, err := simtestutil.GenSignedMockTx(
			r,
			txCfg,
			[]sdk.Msg{msg},
			fees,
			simtestutil.DefaultGenTxGas,
			chainID,
			[]uint64{minterAcc.GetAccountNumber()},
			[]uint64{minterAcc.GetSequence()},
			creator.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, TypeMsgMintNFT, "unable to generate mock tx"), nil, err
		}

		if _, _, err = app.SimDeliver(txCfg.TxEncoder(), tx); err != nil {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", cdc), nil, nil
	}
}
