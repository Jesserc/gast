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
		cmd.Help()
	},
}

// TODO: work on command descriptions (also start with a case)
func init() {
	// Flags and configuration settings.
	TxCmd.AddCommand(estimateGasCmd)
	TxCmd.AddCommand(signCmd)
	TxCmd.AddCommand(createRawCmd)
	TxCmd.AddCommand(sendRawCmd)
	TxCmd.AddCommand(traceTxCmd)
	TxCmd.AddCommand(verifySigCmd)
	TxCmd.AddCommand(getNonceCmd)
	TxCmd.AddCommand(createContractCmd)
}
