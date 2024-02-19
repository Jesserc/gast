/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createRawCmd represents the createRaw command
var createRawCmd = &cobra.Command{
	Use:   "create-raw",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		signedMessageRaw, err := handleCreateRawTransaction(
			txRpcUrl,
			to,
			txData,
			privKey,
			wei,
			nonce,
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
	createRawCmd.Flags().StringVarP(&txRpcUrl, "url", "u", "", "RPC url")
	createRawCmd.Flags().StringVarP(&to, "to", "t", "", "recipient")
	createRawCmd.Flags().StringVarP(&txData, "data", "d", "", "data")
	createRawCmd.Flags().StringVarP(&privKey, "private-key", "p", "", "private key to sign transaction")
	createRawCmd.Flags().Uint64VarP(&wei, "wei", "w", 0, "wei")
	createRawCmd.Flags().Uint64VarP(&nonce, "nonce", "n", 0, "nonce")

	createRawCmd.MarkFlagRequired("url")
	createRawCmd.MarkFlagRequired("to")
	createRawCmd.MarkFlagRequired("data")
	createRawCmd.MarkFlagRequired("private-key")
	createRawCmd.MarkFlagRequired("wei")
	createRawCmd.MarkFlagRequired("nonce")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createRawCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createRawCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
