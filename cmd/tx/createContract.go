/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/Jesserc/gast/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// createContractCmd represents the createContract command
var createContractCmd = &cobra.Command{
	Use:   "create-contract",
	Short: "Deploy Solidity contract",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		bytecode, err := utils.CompileSol(gastParams.DirValue)
		if err != nil {
			log.Crit("Failed to compile contract", "err", err)
		}

		txReceipt, err := CreateContract(
			gastParams.TxRpcUrlValue,
			bytecode,
			gastParams.PrivKeyValue,
			gastParams.GasLimitValue,
			gastParams.WeiValue,
		)
		if err != nil {
			log.Crit("Failed to deploy contract", "err", err)
		}
		if txReceipt != nil {
			fmt.Printf("\nTransaction Receipt:\n")
			fmt.Printf("%sTransaction Hash%s: %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.TransactionHash)
			fmt.Printf("%sBlock Hash:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.BlockHash)
			fmt.Printf("%sBlock Number:%s %v\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.BlockNumber)
			fmt.Printf("%sTransaction Index In Block%s: %v\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.TransactionIndex)
			fmt.Printf("%sType:%s %v\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.Type)
			if txReceipt.Status == 0 {
				fmt.Printf("%sStatus:%s %sFailed%s\n", gastParams.ColorGreen, gastParams.ColorReset, gastParams.ColorRed, gastParams.ColorReset)
			} else {
				fmt.Printf("%sStatus:%s Success\n", gastParams.ColorGreen, gastParams.ColorReset)
			}
			fmt.Printf("%sGas Used:%s %v\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.GasUsed)
			fmt.Printf("%sGas Price:%s %v Wei\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.EffectiveGasPrice)
			fmt.Printf("%sTransaction Fee:%s %v Wei\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.TransactionCost)
			fmt.Printf("%sDeployed Contract Address:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.ContractAddress)
			fmt.Printf("%sLogs:%s %+v\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.Logs)
			fmt.Printf("%sReceipt:%s %s\n", gastParams.ColorGreen, gastParams.ColorReset, txReceipt.TransactionURL)
		} else {
			return
		}
	},
}

func init() {
	// Flags and configuration settings.
	createContractCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url")
	createContractCmd.Flags().StringVarP(&gastParams.DirValue, "dir", "d", "", "path to solidity code")
	createContractCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	createContractCmd.Flags().Uint64VarP(&gastParams.GasLimitValue, "gas-limit", "l", 0, "max gas limit")
	createContractCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "amount to send (optional)")

	// Mark flags required
	createContractCmd.MarkFlagRequired("rpc-url")
	createContractCmd.MarkFlagRequired("private-key")
	createContractCmd.MarkFlagRequired("gas-limit")
	createContractCmd.MarkFlagRequired("dir")
	createContractCmd.MarkFlagsRequiredTogether("url", "private-key", "gas-limit", "dir")
}
