/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/Jesserc/gast/cmd/tx/gastParams"
	"github.com/spf13/cobra"
)

// createRawCmd represents the createRaw command
var createRawCmd = &cobra.Command{
	Use:   "create-raw",
	Short: "Generates a raw, unsigned EIP-1559 transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		signedMessageRaw, err := CreateRawTransaction(
			gastParams.TxRpcUrlValue,
			gastParams.ToValue,
			gastParams.TxDataValue,
			gastParams.PrivKeyValue,
			gastParams.GasLimitValue,
			gastParams.WeiValue,
		)
		if err != nil {
			fmt.Printf("%s%s%s\n", gastParams.ColorRed, err, gastParams.ColorReset)
			os.Exit(1)
		}
		fmt.Printf("%sraw signed message:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, signedMessageRaw)
	},
}

func init() {
	// Flags and configuration settings.
	createRawCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")
	createRawCmd.Flags().StringVarP(&gastParams.ToValue, "to", "t", "", "recipient")
	createRawCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "transaction data (optional)")
	createRawCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	createRawCmd.Flags().Uint64VarP(&gastParams.GasLimitValue, "gas-limit", "l", 0, "transaction gas limit")
	createRawCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "amount to send (optional)")

	// Mark flags required
	createRawCmd.MarkFlagsRequiredTogether("url", "private-key", "gas-limit")
}
