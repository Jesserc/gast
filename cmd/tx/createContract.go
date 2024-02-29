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

// createContractCmd represents the createContract command
var createContractCmd = &cobra.Command{
	Use:   "create-contract",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := CreateContract(
			gastParams.TxRpcUrlValue,
			gastParams.TxDataValue,
			gastParams.PrivKeyValue,
			gastParams.GasLimitValue,
			gastParams.WeiValue,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(d)
	},
}

func init() {
	// Flags and configuration settings.
	createContractCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")
	createContractCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "transaction data (optional)")
	createContractCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	createContractCmd.Flags().Uint64VarP(&gastParams.GasLimitValue, "gas-limit", "l", 0, "max gas limit")
	createContractCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "amount to send (optional)")

	// Mark flags required
	createContractCmd.MarkFlagsRequiredTogether("url", "private-key", "gas-limit")
}
