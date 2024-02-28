/*
Copyright © 2024 NAME HERE <raymondjesse713@gmail.com>
*/

package transaction

import (
	"github.com/spf13/cobra"
)

// TxCmd represents the tx command
var TxCmd = &cobra.Command{
	Use:   "tx",
	Short: "Manages Ethereum transactions, including creation, signing, and tracing",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// TODO: work on command descriptions (also start with a case)
func init() {
	// Flags and configuration settings.
	TxCmd.AddCommand(EstimateGasCmd)
	TxCmd.AddCommand(SignCmd)
	TxCmd.AddCommand(createRawCmd)
	TxCmd.AddCommand(SendRawCmd)
	TxCmd.AddCommand(TraceTxCmd)
	TxCmd.AddCommand(VerifySigCmd)
}
