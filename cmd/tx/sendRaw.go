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
		txReceipt, err := sendRawTransaction(params.RawTx, params.TxRpcUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(txReceipt)
	},
}

func init() {
	// Flags and configuration settings.
	SendRawCmd.Flags().StringVarP(&params.RawTx, "raw-tx", "r", "", "raw transaction to send")
	SendRawCmd.Flags().StringVarP(&params.TxRpcUrl, "rpc-url", "u", "", "specify RPC url for transaction")

	SendRawCmd.MarkFlagRequired("raw-tx")
	SendRawCmd.MarkFlagRequired("rpc-url")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendRawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendRawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
