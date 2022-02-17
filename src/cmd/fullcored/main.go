package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/obada-foundation/fullcore/app"
	"github.com/obada-foundation/fullcore/cmd/fullcored/cmd"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.Name,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	rootCmd.AddCommand(cmd.NewTestnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}))
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
