/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// getBlobCmd represents the getBlob command
var getBlobCmd = &cobra.Command{
	Use:   "get-blob",
	Short: "Get blob transaction data",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		blob, err := GetBlob(gastParams.BlockRootOrSlotNumberVar, gastParams.KZGCommitmentVar)
		if err != nil {
			log.Crit("Failed to get blob", "error", err)
		}

		fmt.Println("Blob data\n", blob)
	},
}

func init() {
	// Flags and configuration settings.
	getBlobCmd.Flags().StringVarP(&gastParams.BlockRootOrSlotNumberVar, "id", "i", "", "block root (32 bytes) or slot number")
	getBlobCmd.Flags().StringVarP(&gastParams.KZGCommitmentVar, "kzg-commitment", "k", "", "kzg commitment (48 bytes)")

	getBlobCmd.MarkFlagRequired("id")
	getBlobCmd.MarkFlagRequired("kzg-commitment")
}
