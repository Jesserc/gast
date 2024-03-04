/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// sendRawCmd represents the sendRaw command
var sendRawCmd = &cobra.Command{
	Use:   "send-raw",
	Short: "Submits a raw, signed transaction to the Ethereum network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		txReceipt, txDetails, err := SendRawTransaction(gastParams.RawTxValue, gastParams.TxRpcUrlValue)
		if err != nil {
			log.Error("An error occurred", "err", err)
			os.Exit(1)
		}
		// Print the entire JSON with the added fields
		fmt.Println(gastParams.ColorGreen, "Transaction details:", gastParams.ColorReset)
		fmt.Println(txDetails)
		fmt.Printf("%sReceipt:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt)

	},
}

func init() {
	// Flags and configuration settings.
	sendRawCmd.Flags().StringVarP(&gastParams.RawTxValue, "raw-tx", "r", "", "raw transaction to send")
	sendRawCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url for transaction")

	// Mark flags required
	sendRawCmd.MarkFlagsRequiredTogether("raw-tx", "rpc-url")
}
