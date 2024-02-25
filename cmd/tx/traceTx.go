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

// TraceTxCmd represents the traceTx command
var TraceTxCmd = &cobra.Command{
	Use:   "trace-tx",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("traceTx called")
		_, err := handleTraceTx(params.TxHash, params.TxRpcUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// fmt.Println(tracedTx)
	},
}

func init() {
	// Flags and configuration settings.
	TraceTxCmd.Flags().StringVar(&params.TxHash, "hash", "", "Transaction hash to trace")
	TraceTxCmd.Flags().StringVarP(&params.TxRpcUrl, "url", "u", "", "RPC url")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// traceTxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// traceTxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
