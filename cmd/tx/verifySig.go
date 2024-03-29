/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"os"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// verifySigCmd represents the verifySig command
var verifySigCmd = &cobra.Command{
	Use:   "verify-sig",
	Short: "Verify the signature of a signed message (can be created with the sign-message command)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		isSigner, err := VerifySig(gastParams.SigValue, gastParams.SigAddressValue, gastParams.SigMsgValue)
		if err != nil {
			log.Crit("Failed to verify signature", "error", err)
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
	verifySigCmd.Flags().StringVarP(&gastParams.SigValue, "sig", "s", "", "Signed message to verify")
	verifySigCmd.Flags().StringVarP(&gastParams.SigAddressValue, "address", "a", "", "Message signer address")
	verifySigCmd.Flags().StringVarP(&gastParams.SigMsgValue, "msg", "m", "", "Original message")

	// Mark flags required
	verifySigCmd.MarkFlagRequired("sig")
	verifySigCmd.MarkFlagRequired("address")
	verifySigCmd.MarkFlagRequired("msg")
	verifySigCmd.MarkFlagsRequiredTogether("sig", "address", "msg")
}
