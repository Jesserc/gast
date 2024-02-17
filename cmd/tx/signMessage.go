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

		handleSignMessage(
			"https://sepolia.infura.io/v3/",
			"0x571B102323C3b8B8Afb30619Ac1d36d85359fb84",
			"eth signed message",
			"0x2843e08c0fa87258545656e44955aa2c6ca2ebb92fa65507e4e5728570d36662",
			0x9184e72a,
			0,
		)
		// fmt.Println(message)
		fmt.Println("sign message done")

		// to string, data string, privateKey string, wei uint64, nonce uint64
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signMessageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signMessageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
