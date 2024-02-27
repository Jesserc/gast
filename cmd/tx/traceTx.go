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

// TraceTxCmd represents the traceTx command
var TraceTxCmd = &cobra.Command{
	Use:   "trace-tx",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("traceTx called")
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
	TraceTxCmd.Flags().StringVar(&gastParams.TxHashValue, "hash", "", "Transaction hash to trace")
	TraceTxCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "url", "u", "", "RPC url")

	// Mark flags required
	TraceTxCmd.MarkFlagsRequiredTogether("hash", "url")
}
