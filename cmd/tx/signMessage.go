/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SignMessageCmd represents the signMessage command
var SignMessageCmd = &cobra.Command{
	Use:   "sign-message",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signMessage called")
		if signRaw {
			handleSignMessageRaw(
				txRpcUrl,
				to,
				txData,
				privKey,
				wei,
				nonce,
			)
		} else if signHash {
			handleSignMessageHash(
				txRpcUrl,
				to,
				txData,
				privKey,
				wei,
				nonce,
			)
		}

		// fmt.Println(message)
		fmt.Println("sign message done")

		// to string, data string, privateKey string, wei uint64, nonce uint64
	},
}

func init() {
	// Flags and configuration settings.
	SignMessageCmd.Flags().BoolVarP(&signHash, "sign-hash", "s", false, "sign transaction and return signature hash")
	SignMessageCmd.Flags().BoolVarP(&signRaw, "sign-raw", "r", false, "sign transaction and return raw signature")

	SignMessageCmd.Flags().StringVarP(&txRpcUrl, "url", "u", "", "RPC url")
	SignMessageCmd.Flags().StringVarP(&to, "to", "t", "", "recipient")
	SignMessageCmd.Flags().StringVarP(&txData, "data", "d", "", "data")
	SignMessageCmd.Flags().StringVarP(&privKey, "private-key", "p", "", "private key to sign transaction")
	SignMessageCmd.Flags().Uint64VarP(&wei, "wei", "w", 0, "wei")
	SignMessageCmd.Flags().Uint64VarP(&nonce, "nonce", "n", 0, "nonce")

	SignMessageCmd.MarkFlagsOneRequired("sign-hash", "sign-raw")

	SignMessageCmd.MarkFlagRequired("url")
	SignMessageCmd.MarkFlagRequired("to")
	SignMessageCmd.MarkFlagRequired("data")
	SignMessageCmd.MarkFlagRequired("private-key")
	SignMessageCmd.MarkFlagRequired("wei")
	SignMessageCmd.MarkFlagRequired("nonce")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signMessageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signMessageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
