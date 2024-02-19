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
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	// Flags and configuration settings.
	TxCmd.AddCommand(EstimateGasCmd)
	TxCmd.AddCommand(SignCmd)
	TxCmd.AddCommand(createRawCmd)
}
