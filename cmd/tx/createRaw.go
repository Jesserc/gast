/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// createRawCmd represents the createRaw command
var createRawCmd = &cobra.Command{
	Use:   "create-raw",
	Short: "Generate a raw, signed EIP-1559 transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rawTransaction, err := CreateRawTransaction(
			gastParams.TxRpcUrlValue,
			gastParams.ToValue,
			gastParams.TxDataValue,
			gastParams.PrivKeyValue,
			gastParams.GasLimitValue,
			gastParams.WeiValue,
		)
		if err != nil {
			log.Crit("Failed to create raw transaction", "error", err)
		}
		fmt.Println() // spacing
		fmt.Printf("%sraw signed message:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, rawTransaction)
	},
}

func init() {
	// Flags and configuration settings.
	createRawCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url")
	createRawCmd.Flags().StringVarP(&gastParams.ToValue, "to", "t", "", "recipient")
	createRawCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "transaction data (optional)")
	createRawCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	createRawCmd.Flags().Uint64VarP(&gastParams.GasLimitValue, "gas-limit", "l", 0, "transaction gas limit")
	createRawCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "amount to send (optional)")

	// Mark flags required
	createRawCmd.MarkFlagRequired("rpc-url")
	createRawCmd.MarkFlagRequired("private-key")
	createRawCmd.MarkFlagRequired("gas-limit")
	createRawCmd.MarkFlagsRequiredTogether("url", "private-key", "gas-limit")
}
