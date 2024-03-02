/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/Jesserc/gast/cmd/tx/gastParams"
	"github.com/Jesserc/gast/utils"
	"github.com/spf13/cobra"
)

// createContractCmd represents the createContract command
var createContractCmd = &cobra.Command{
	Use:   "create-contract",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		bytecode, err := utils.CompileSol(gastParams.DirValue)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		txReceipt, err := CreateContract(
			gastParams.TxRpcUrlValue,
			// "608060405234801561000f575f80fd5b506107e85f8190555061010d806100255f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c8063ef88a09214610038578063fd08921b14610054575b5f80fd5b610052600480360381019061004d91906100ba565b610072565b005b61005c61007b565b60405161006991906100f4565b60405180910390f35b805f8190555050565b5f8054905090565b5f80fd5b5f819050919050565b61009981610087565b81146100a3575f80fd5b50565b5f813590506100b481610090565b92915050565b5f602082840312156100cf576100ce610083565b5b5f6100dc848285016100a6565b91505092915050565b6100ee81610087565b82525050565b5f6020820190506101075f8301846100e5565b9291505056",
			bytecode,
			gastParams.PrivKeyValue,
			gastParams.GasLimitValue,
			gastParams.WeiValue,
		)
		// txReceipt, err := CreateContract(
		// 	gastParams.TxRpcUrlValue,
		// 	gastParams.TxDataValue,
		// 	gastParams.PrivKeyValue,
		// 	gastParams.GasLimitValue,
		// 	gastParams.WeiValue,
		// )
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

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

	},
}

func init() {
	// Flags and configuration settings.
	createContractCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")
	// createContractCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "transaction data (optional)")
	createContractCmd.Flags().StringVarP(&gastParams.DirValue, "dir", "d", "", "path to solidity code")
	createContractCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	createContractCmd.Flags().Uint64VarP(&gastParams.GasLimitValue, "gas-limit", "l", 0, "max gas limit")
	createContractCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "amount to send (optional)")

	// Mark flags required
	createContractCmd.MarkFlagsRequiredTogether("url", "private-key", "gas-limit", "dir")
}

// func compileSol(dir string) {
// 	var stdoutBuf, stderrBuf bytes.Buffer
// 	command := "solc" + " " + dir + " " + "--no-cbor-metadata" + " " + "--bin"
// 	cmd := exec.Command("sh", "-c", command)
// 	cmd.Stdout = &stdoutBuf
// 	cmd.Stderr = &stderrBuf
//
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Printf("Error: %s", err)
// 		log.Printf("stderr: %s", stderrBuf.String())
// 		return
// 	}
//
// 	output := stdoutBuf.String()
// 	fmt.Println(output)
//
// 	startIndex := strings.Index(output, "6080")
// 	fmt.Println(startIndex)
// 	if startIndex == -1 {
// 		// log.Println("Bytecode start sequence not found.")
// 		return
// 	}
//
// 	// Extract and print the bytecode.
// 	bytecode := output[startIndex:]
// 	fmt.Println(bytecode)
// }
