/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// estimateGasCmd represents the TryEstimateGas command
var estimateGasCmd = &cobra.Command{
	Use:   "estimate-gas",
	Short: "Estimate the gas required to execute a given transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		estimatedTxGas, err := TryEstimateGas(gastParams.TxRpcUrlValue, gastParams.FromValue, gastParams.ToValue, gastParams.TxDataValue, gastParams.WeiValue)
		if err != nil {
			log.Crit("Failed to estimate gas", "error", err)
		}
		fmt.Printf("Estimated gas: %s%d%s\n", gastParams.ColorGreen, estimatedTxGas, gastParams.ColorReset)
	},
}

func init() {
	// Flags and configuration settings.
	estimateGasCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url")
	estimateGasCmd.Flags().StringVarP(&gastParams.FromValue, "from", "f", "", "sender")
	estimateGasCmd.Flags().StringVarP(&gastParams.ToValue, "to", "t", "", "recipient")
	estimateGasCmd.Flags().StringVarP(&gastParams.TxDataValue, "data", "d", "", "data (optional)")
	estimateGasCmd.Flags().Uint64VarP(&gastParams.WeiValue, "wei", "w", 0, "wei (optional)")

	// Mark flags required
	estimateGasCmd.MarkFlagRequired("rpc-url")
	estimateGasCmd.MarkFlagRequired("from")
	estimateGasCmd.MarkFlagRequired("to")
	estimateGasCmd.MarkFlagsRequiredTogether("url", "from", "to")
}
