/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send EIP-1559 transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		txReceiptUrl := SendTransaction(
			gastParams.TxRpcUrlValue,
			gastParams.ToValue,
			gastParams.TxDataValue,
			gastParams.PrivKeyValue,
			gastParams.GasLimitValue,
			gastParams.WeiValue,
		)

		fmt.Printf("%sTx Receipt:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceiptUrl)
	},
}

func init() {
	// Flags and configuration settings.
	sendCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")
	sendCmd.Flags().StringVarP(&gastParams.ToValue, "to", "t", "", "recipient")
	sendCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "transaction data (optional)")
	sendCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	sendCmd.Flags().Uint64VarP(&gastParams.GasLimitValue, "gas-limit", "l", 0, "transaction gas limit")
	sendCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "amount to send (optional)")

	// Mark flags required
	sendCmd.MarkFlagsRequiredTogether("url", "private-key", "gas-limit")
}
