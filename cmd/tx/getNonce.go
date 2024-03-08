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

// getNonceCmd represents the getNonce command
var getNonceCmd = &cobra.Command{
	Use:   "get-nonce",
	Short: "Get the transaction count of an account",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		currentNonce, nextNonce, err := GetNonce(gastParams.FromValue, gastParams.TxRpcUrlValue)
		if err != nil {
			log.Crit("Failed to get nonce", "Error", err)
		}
		fmt.Printf("%sCurrent nonce:%s %v, %sNext nonce:%s %v\n", gastParams.ColorGreen, gastParams.ColorReset, currentNonce, gastParams.ColorGreen, gastParams.ColorReset, nextNonce)
	},
}

func init() {
	// Flags and configuration settings.
	getNonceCmd.Flags().StringVarP(&gastParams.FromValue, "address", "a", "", "Address to get nonce")
	getNonceCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url")

	getNonceCmd.MarkFlagRequired("address")
	getNonceCmd.MarkFlagRequired("url")
	getNonceCmd.MarkFlagsRequiredTogether("address", "url")
}
