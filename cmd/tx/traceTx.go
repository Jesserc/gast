/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// traceTxCmd represents the trace command
var traceTxCmd = &cobra.Command{
	Use:   "trace",
	Short: "Retrieve and display the execution trace (path) of a given transaction hash",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rootTrace, err := TraceTx(gastParams.TxHashValue, gastParams.TxRpcUrlValue)
		if err != nil {
			log.Crit("Failed to trace transaction", "error", err)
		}
		printTrace(rootTrace, 0, false, "")
	},
}

func init() {
	// Flags and configuration settings.
	traceTxCmd.Flags().StringVar(&gastParams.TxHashValue, "hash", "", "Transaction hash to trace")
	traceTxCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url (optional, but must support Otterscan's ots_traceTransaction")

	// Mark flags required
	traceTxCmd.MarkFlagRequired("hash")
}
