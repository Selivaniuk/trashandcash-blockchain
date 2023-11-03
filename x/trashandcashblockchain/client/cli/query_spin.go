package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"trashandcash-blockchain/x/trashandcashblockchain/types"
)

func CmdListSpin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-spin",
		Short: "list all spin",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// pageReq, err := client.ReadPageRequest(cmd.Flags())
			// if err != nil {
			// 	return err
			// }

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySpinsRequest{
				// Pagination: pageReq,
			}

			res, err := queryClient.Spins(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowSpin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-spin [id]",
		Short: "shows a spin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetSpinRequest{
				Id: id,
			}

			res, err := queryClient.Spin(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
