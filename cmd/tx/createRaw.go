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
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		signedMessageRaw, err := createRawTransaction(
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
	createRawCmd.Flags().Uint64VarP(&gastParams.GasLimitValue, "gas-limit", "l", 0, "max gas limit")
	createRawCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "amount to send (optional)")
	createRawCmd.Flags().Uint64VarP(&gastParams.NonceValue, "nonce", "n", 0, "transaction nonce")

	createRawCmd.MarkFlagRequired("url")
	createRawCmd.MarkFlagRequired("to")
	createRawCmd.MarkFlagRequired("private-key")
	createRawCmd.MarkFlagRequired("gas-limit")
}
