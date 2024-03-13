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
	Short: "Manage transactions, including creation, signing, submitting and tracing",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

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
	TxCmd.AddCommand(sendBlobTxCmd)
	TxCmd.AddCommand(sendCmd)
	TxCmd.AddCommand(getBlobCmd)
}
