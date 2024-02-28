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

// traceTxCmd represents the traceTx command
var traceTxCmd = &cobra.Command{
	Use:   "trace",
	Short: "Retrieves and displays the execution trace (path) of a given transaction hash",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		rootTrace, err := handleTraceTx(gastParams.TxHashValue, gastParams.TxRpcUrlValue)
		if err != nil {
			fmt.Printf("%s%s%s\n", gastParams.ColorRed, err.Error(), gastParams.ColorReset)
			os.Exit(1)
		}
		printTrace(rootTrace, 0, false, "")
	},
}

func init() {
	// Flags and configuration settings.
	traceTxCmd.Flags().StringVar(&gastParams.TxHashValue, "hash", "", "Transaction hash to trace")
	traceTxCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")

	// Mark flags required
	traceTxCmd.MarkFlagsRequiredTogether("hash", "url")
}
