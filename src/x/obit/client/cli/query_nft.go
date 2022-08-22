package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/obada-foundation/fullcore/x/obit/types"
	"github.com/spf13/cobra"
)

func CmdShowNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft [did]",
		Short: "shows single NFT by DID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			did := args[0]

			params := &types.QueryGetNftRequest{
				Did: did,
			}

			res, err := queryClient.GetNft(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowByOwner() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nfts-by-owner [owner-address]",
		Short: "shows all obits owned by specific address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			address := args[0]

			params := &types.QueryGetNftsByAddressRequest{
				Address: address,
			}

			res, err := queryClient.GetNftsByAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
