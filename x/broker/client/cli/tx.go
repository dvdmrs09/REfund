package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/defund-labs/defund/x/broker/types"
	"github.com/spf13/cobra"
)

// GetTxCmd creates and returns the broker tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdAddLiquiditySource(),
		CmdAddConnectionBroker(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdAddLiquiditySource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-liquidity-source [broker-id] [pool-id]",
		Short: "Add a new liquidity source for the broker specified.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBrokerId := args[0]
			argPoolId := args[1]

			poolid, err := strconv.ParseUint(argPoolId, 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddLiquiditySource(
				clientCtx.GetFromAddress().String(),
				argBrokerId,
				poolid,
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

func CmdAddConnectionBroker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activate-broker [broker-id] [connection-id]",
		Short: "Activates a broker chain by adding the Connection ID to the broker. Can only activate inactive brokers.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBrokerId := args[0]
			argConnectionId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddConnectionBroker(
				clientCtx.GetFromAddress().String(),
				argBrokerId,
				argConnectionId,
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
