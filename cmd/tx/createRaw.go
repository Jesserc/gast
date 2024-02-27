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

// createRawCmd represents the createRaw command
var createRawCmd = &cobra.Command{
	Use:   "create-raw",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		signedMessageRaw, err := createRawTransaction(
			params.TxRpcUrlValue,
			params.ToValue,
			params.TxDataValue,
			params.PrivKeyValue,
			params.GasLimitValue,
			params.WeiValue,
		)
		if err != nil {
			fmt.Printf("%s%s%s\n", params.ColorRed, err, params.ColorReset)
			os.Exit(1)
		}
		fmt.Printf("%sraw signed message:%s %s\n", params.ColorGreen, params.ColorReset, signedMessageRaw)
	},
}

func init() {

	// Flags and configuration settings.
	createRawCmd.Flags().StringVarP(&params.TxRpcUrlValue, "url", "u", "", "RPC url")
	createRawCmd.Flags().StringVarP(&params.ToValue, "to", "t", "", "recipient")
	createRawCmd.Flags().StringVarP(&params.TxDataValue, "data", "d", "", "transaction data (optional)")
	createRawCmd.Flags().StringVarP(&params.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	createRawCmd.Flags().Uint64VarP(&params.GasLimitValue, "gas-limit", "l", 0, "max gas limit")
	createRawCmd.Flags().Uint64VarP(&params.WeiValue, "wei", "w", 0, "amount to send (optional)")
	createRawCmd.Flags().Uint64VarP(&params.NonceValue, "nonce", "n", 0, "transaction nonce")

	createRawCmd.MarkFlagRequired("url")
	createRawCmd.MarkFlagRequired("to")
	createRawCmd.MarkFlagRequired("private-key")
	createRawCmd.MarkFlagRequired("gas-limit")
}
