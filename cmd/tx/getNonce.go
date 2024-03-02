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

// getNonceCmd represents the getNonce command
var getNonceCmd = &cobra.Command{
	Use:   "get-nonce",
	Short: "Get transaction count of an account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		currentNonce, nextNonce, err := GetNonce(gastParams.FromValue, gastParams.TxRpcUrlValue)
		if err != nil {
			fmt.Printf("%s%s%s\n", gastParams.ColorRed, err, gastParams.ColorReset)
			os.Exit(1)
		}

		fmt.Printf("%sCurrent nonce:%s %v, %sNext nonce:%s %v\n", gastParams.ColorGreen, gastParams.ColorReset, currentNonce, gastParams.ColorGreen, gastParams.ColorReset, nextNonce)
	},
}

func init() {
	// Flags and configuration settings.
	getNonceCmd.Flags().StringVarP(&gastParams.FromValue, "address", "a", "", "Address to get nonce")
	getNonceCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")

	getNonceCmd.MarkFlagsRequiredTogether("address", "url")
}
