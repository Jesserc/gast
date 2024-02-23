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
			params.TxRpcUrl,
			params.To,
			params.TxData,
			params.PrivKey,
			params.GasLimit,
			params.Wei,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("\nraw signed message:", signedMessageRaw)
	},
}

func init() {

	// Flags and configuration settings.
	createRawCmd.Flags().StringVarP(&params.TxRpcUrl, "url", "u", "", "RPC url")
	createRawCmd.Flags().StringVarP(&params.To, "to", "t", "", "recipient")
	createRawCmd.Flags().StringVarP(&params.TxData, "data", "d", "", "transaction data (optional)")
	createRawCmd.Flags().StringVarP(&params.PrivKey, "private-key", "p", "", "private key to sign transaction")
	createRawCmd.Flags().Uint64VarP(&params.GasLimit, "gas-limit", "l", 0, "max gas limit")
	createRawCmd.Flags().Uint64VarP(&params.Wei, "wei", "w", 0, "amount to send (optional)")
	createRawCmd.Flags().Uint64VarP(&params.Nonce, "nonce", "n", 0, "transaction nonce")

	createRawCmd.MarkFlagRequired("url")
	createRawCmd.MarkFlagRequired("to")
	createRawCmd.MarkFlagRequired("private-key")
	createRawCmd.MarkFlagRequired("gas-limit")
}
