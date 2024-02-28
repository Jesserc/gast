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

// VerifySigCmd represents the verifySig command
var VerifySigCmd = &cobra.Command{
	Use:   "verify-sig",
	Short: "Verifies the signature of a signed message (can be created with the sign-message command)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		isSigner, err := handleVerifySig(gastParams.SigValue, gastParams.SigAddressValue, gastParams.SigMsgValue)
		if err != nil {
			fmt.Printf("%s%s%s\n", gastParams.ColorRed, err, gastParams.ColorReset)
			os.Exit(1)
		}
		if isSigner {
			fmt.Printf("%s %ssigned%s \"%s\"\n", gastParams.SigAddressValue, gastParams.ColorGreen, gastParams.ColorReset, gastParams.SigMsgValue)
		} else {
			fmt.Printf("%s %sdid not signed%s \"%s\"\n", gastParams.SigAddressValue, gastParams.ColorRed, gastParams.ColorReset, gastParams.SigMsgValue)
			os.Exit(1)
		}
	},
}

func init() {
	// Flags and configuration settings
	VerifySigCmd.Flags().StringVarP(&gastParams.SigValue, "sig", "s", "", "Signed message to verify")
	VerifySigCmd.Flags().StringVarP(&gastParams.SigAddressValue, "address", "a", "", "Message signer address")
	VerifySigCmd.Flags().StringVarP(&gastParams.SigMsgValue, "msg", "m", "", "Original message")

	// Mark flags required
	VerifySigCmd.MarkFlagsRequiredTogether("sig", "address", "msg")
}
