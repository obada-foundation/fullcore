package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/spf13/cobra"
)

func CmdShowByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nfts-by-owner [owner-address]",
		Short: "shows all obits owned by specific address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			owner := args[0]

			params := &types.QueryGetAllNftByOwnerRequest{
				Owner: owner,
			}

			res, err := queryClient.GetAllNftByOwner(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
