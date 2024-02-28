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

// SendRawCmd represents the sendRaw command
var SendRawCmd = &cobra.Command{
	Use:   "send-raw",
	Short: "Submits a raw, signed transaction to the Ethereum network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		txReceipt, err := sendRawTransaction(gastParams.RawTxValue, gastParams.TxRpcUrlValue)
		if err != nil {
			fmt.Printf("%s%s%s\n", gastParams.ColorRed, err, gastParams.ColorReset)
			os.Exit(1)
		}

		fmt.Printf("%sReceipt:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt)
	},
}

func init() {
	// Flags and configuration settings.
	SendRawCmd.Flags().StringVarP(&gastParams.RawTxValue, "raw-tx", "r", "", "raw transaction to send")
	SendRawCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "specify RPC url for transaction")

	// Mark flags required
	SendRawCmd.MarkFlagsRequiredTogether("raw-tx", "rpc-url")
}
