package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEditMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update metadata [did]",
		Short: "Broadcast message edit-metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			nftData := types.NFTData{}

			msg := types.NewMsgEditMetadata(
				argDID,
				clientCtx.GetFromAddress().String(),
				nftData,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
