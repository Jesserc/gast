/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/Jesserc/gast/cmd/tx/params"
	"github.com/spf13/cobra"
)

// SendRawCmd represents the sendRaw command
var SendRawCmd = &cobra.Command{
	Use:   "send-raw",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		txReceipt, err := sendRawTransaction(params.RawTxValue, params.TxRpcUrlValue)
		if err != nil {
			fmt.Printf("%s%s%s\n", params.ColorRed, err, params.ColorReset)
			os.Exit(1)
		}

		fmt.Printf("%sReceipt:%s %s\n", params.ColorGreen, params.ColorReset, txReceipt)
	},
}

func init() {
	// Flags and configuration settings.
	SendRawCmd.Flags().StringVarP(&params.RawTxValue, "raw-tx", "r", "", "raw transaction to send")
	SendRawCmd.Flags().StringVarP(&params.TxRpcUrlValue, "rpc-url", "u", "", "specify RPC url for transaction")

	SendRawCmd.MarkFlagRequired("raw-tx")
	SendRawCmd.MarkFlagRequired("rpc-url")
}
