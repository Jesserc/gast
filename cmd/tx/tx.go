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

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// txCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// txCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
