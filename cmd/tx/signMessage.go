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

// SignCmd represents the signMessage command
var SignCmd = &cobra.Command{
	Use:   "sign-message",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		signedMessageHash, err := signMessage(
			gastParams.TxDataValue,
			gastParams.PrivKeyValue,
		)
		if err != nil {
			fmt.Printf("%s%s%s\n", gastParams.ColorRed, err.Error(), gastParams.ColorReset)
			os.Exit(1)
		}
		fmt.Printf("%ssigned message:%s\n %s\n", gastParams.ColorGreen, gastParams.ColorReset, signedMessageHash)
	},
}

func init() {
	// Flags and configuration settings.
	SignCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "message to sign")
	SignCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")

	// Mark flags required
	SignCmd.MarkFlagsRequiredTogether("data", "private-key")
}
