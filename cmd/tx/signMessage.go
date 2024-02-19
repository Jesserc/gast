/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// SignCmd represents the signMessage command
var SignCmd = &cobra.Command{
	Use:   "sign-message",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		signedMessageHash, err := handleSignMessageHash(
			txData,
			privKey,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("\nsigned message hash:", signedMessageHash)
	},
}

func init() {
	// Flags and configuration settings.
	// SignCmd.Flags().BoolVarP(&signHash, "sign", "s", false, "sign transaction and return signature hash")
	// SignCmd.Flags().BoolVarP(&createRaw, "create-raw", "r", false, "create raw transaction that can be propagated later")

	SignCmd.Flags().StringVarP(&txData, "data", "d", "", "message to sign")
	SignCmd.Flags().StringVarP(&privKey, "private-key", "p", "", "private key to sign transaction")

	SignCmd.MarkFlagRequired("data")
	SignCmd.MarkFlagRequired("private-key")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signMessageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signMessageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
