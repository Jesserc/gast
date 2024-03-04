/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"os"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// traceTxCmd represents the trace command
var traceTxCmd = &cobra.Command{
	Use:   "trace",
	Short: "Retrieves and displays the execution trace (path) of a given transaction hash",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rootTrace, err := TraceTx(gastParams.TxHashValue, gastParams.TxRpcUrlValue)
		if err != nil {
			log.Error("An error occurred", "err", err)
			os.Exit(1)
		}
		printTrace(rootTrace, 0, false, "")
	},
}

func init() {
	// Flags and configuration settings.
	traceTxCmd.Flags().StringVar(&gastParams.TxHashValue, "hash", "", "Transaction hash to trace")
	traceTxCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url (optional, but must support Otterscan's ots_traceTransaction")

	// Mark flags required
	traceTxCmd.MarkFlagRequired("hash")
}
