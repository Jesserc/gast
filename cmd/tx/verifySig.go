/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"github.com/Jesserc/gast/cmd/tx/gastParams"
	"github.com/spf13/cobra"
)

// VerifySigCmd represents the verifySig command
var VerifySigCmd = &cobra.Command{
	Use:   "verifySig",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		handleVerifySig(gastParams.SigValue, gastParams.SigAddressValue, gastParams.SigHashValue)
	},
}

func init() {
	// Flags and configuration settings.
	VerifySigCmd.Flags().StringVarP(&gastParams.SigValue, "", "", "", "")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifySigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifySigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
