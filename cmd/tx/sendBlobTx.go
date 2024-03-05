/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// sendBlobTxCmd represents the sendBlobTx command
var sendBlobTxCmd = &cobra.Command{
	Use:   "send-blob",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		blobTxHash := SendBlobTX(gastParams.TxRpcUrlValue, gastParams.TxDataValue, gastParams.PrivKeyValue, gastParams.ToValue, "gast/blob-tx/")
		log.Info("Successfully sent blob transaction", "hash", " "+blobTxHash)

	},
}

func init() {
	// Flags and configuration settings.
	sendBlobTxCmd.Flags().StringVarP(&gastParams.TxRpcUrlValue, "rpc-url", "u", "", "RPC url for transaction")
	sendBlobTxCmd.Flags().StringVarP(&gastParams.TxDataValue, "blob-data", "b", "", "blob data (hex or string)")
	sendBlobTxCmd.Flags().StringVarP(&gastParams.PrivKeyValue, "private-key", "p", "", "private key to sign transaction")
	sendBlobTxCmd.Flags().StringVarP(&gastParams.ToValue, "to", "t", "", "blob transaction recipient")
}
