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

func CmdMintObit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-obit [serial-number-hash] [manufacturer] [pn] [ta] [owner-did]",
		Short: "Broadcast message mint-obit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSnh := args[0]
			argMan := args[1]
			argPn := args[2]
			argTa := args[3]
			argOwnerDid := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintObit(
				clientCtx.GetFromAddress().String(),
				argSnh,
				argMan,
				argPn,
				argTa,
				argOwnerDid,
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
