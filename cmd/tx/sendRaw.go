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

// sendRawCmd represents the sendRaw command
var sendRawCmd = &cobra.Command{
	Use:   "send-raw",
	Short: "Submit a raw, signed transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		txReceiptUrl, txDetails, err := SendRawTransaction(gastParams.RawTxValue, gastParams.TxRpcUrlValue)
		if err != nil {
			log.Crit("Failed to send raw transaction", "error", err)
		}

		log.Info("Transaction details:")
		fmt.Println(txDetails)
		fmt.Printf("%sTx Receipt:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceiptUrl)

	},
}

func init() {
	// Flags and configuration settings.
	sendRawCmd.Flags().StringVarP(&gastParams.RawTxValue, "raw-tx", "r", "", "raw transaction to send")
	sendRawCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url for transaction")

	// Mark flags required
	sendRawCmd.MarkFlagsRequiredTogether("raw-tx", "rpc-url")
}
