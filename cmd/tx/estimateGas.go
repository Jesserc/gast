/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/Jesserc/gast/cmd/tx/params"
	"github.com/spf13/cobra"
)

// EstimateGasCmd represents the estimateGas command
var EstimateGasCmd = &cobra.Command{
	Use:   "estimate-gas",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		estimatedTxGas, err := estimateGas(params.TxRpcUrlValue, params.FromValue, params.ToValue, params.TxDataValue, params.WeiValue)
		if err != nil {
			fmt.Printf("%s%s%s\n", params.ColorRed, err.Error(), params.ColorReset)
			os.Exit(1)
		}
		fmt.Printf("Estimated gas: %s%d%s\n", params.ColorGreen, estimatedTxGas, params.ColorReset)
	},
}

func init() {
	// Flags and configuration settings.
	EstimateGasCmd.Flags().StringVarP(&params.TxRpcUrlValue, "url", "u", "", "RPC url")
	EstimateGasCmd.Flags().StringVarP(&params.FromValue, "from", "f", "", "sender")
	EstimateGasCmd.Flags().StringVarP(&params.ToValue, "to", "t", "", "recipient")
	EstimateGasCmd.Flags().StringVarP(&params.TxDataValue, "data", "d", "", "data")
	EstimateGasCmd.Flags().Uint64VarP(&params.WeiValue, "wei", "w", 0, "wei")

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
	// estimateGasCmd.Flags().BoolP("toggle", "handleTraceTx", false, "Help message for toggle")
}
