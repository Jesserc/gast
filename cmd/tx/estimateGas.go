/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	txRpcUrl, from, to, txData string
	wei                        uint64
)

// EstimateGasCmd represents the estimateGas command
var EstimateGasCmd = &cobra.Command{
	Use:   "estimate-gas",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// if txRpcUrl == "" || from == "" || to == "" || txData == "" || wei == 0 {
		// 	cmd.Usage()
		// 	return
		// }
		estimatedTxGas, err := HandleEstimateGas(txRpcUrl, from, to, txData, wei)
		if err != nil {
			fmt.Println(err) // TODO: log as error
			return
		}
		fmt.Println("Estimated gas:", estimatedTxGas)
	},
}

func init() {
	// Flags and configuration settings.
	EstimateGasCmd.Flags().StringVarP(&txRpcUrl, "url", "u", "", "RPC url")
	EstimateGasCmd.Flags().StringVarP(&from, "from", "f", "", "sender")
	EstimateGasCmd.Flags().StringVarP(&to, "to", "t", "", "recipient")
	EstimateGasCmd.Flags().StringVarP(&txData, "data", "d", "", "transaction data")
	EstimateGasCmd.Flags().Uint64VarP(&wei, "wei", "w", 0, "amount in wei")

	EstimateGasCmd.MarkFlagRequired("url")
	EstimateGasCmd.MarkFlagRequired("from")
	EstimateGasCmd.MarkFlagRequired("to")
	EstimateGasCmd.MarkFlagRequired("data")
	EstimateGasCmd.MarkFlagRequired("wei")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// estimateGasCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// estimateGasCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
