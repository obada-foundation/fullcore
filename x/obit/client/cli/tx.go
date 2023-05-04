package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/obada-foundation/fullcore/x/obit/types"
)

var (
	// DefaultRelativePacketTimeoutTimestamp relative packet timeout timestamp
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	// nolint:unused // for the future refactoring
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	// nolint:unused // for the future refactoring
	listSeparator = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdMintNFT())
	cmd.AddCommand(CmdUpdateNFT())
	// this line is used by starport scaffolding # 1

	return cmd
}
