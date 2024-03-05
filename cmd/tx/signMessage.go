/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/spf13/cobra"
)

// signCmd represents the SignETHMessage command
var signCmd = &cobra.Command{
	Use:   "sign-message",
	Short: "Signs a given message with the private key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		signedMessageHash := SignETHMessage(
			gastParams.TxDataValue,
			gastParams.PrivKeyValue,
		)
		
		fmt.Printf("%ssigned message:%s\n %s\n", gastParams.ColorGreen, gastParams.ColorReset, signedMessageHash)
	},
}

func init() {
	// Flags and configuration settings.
	signCmd.Flags().StringVarP(&gastParams.TxDataValue, "message", "m", "", "message to sign")
	signCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")

	// Mark flags required
	signCmd.MarkFlagsRequiredTogether("message", "private-key")
}
