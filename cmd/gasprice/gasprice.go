/*
Copyright © 2024 NAME HERE <raymondjesse713@gmail.com>
*/

package gasprice

import (
	"fmt"

	"github.com/Jesserc/gast/cmd/gastParams"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

var (
	eth, op, base, linea, arb, zkSync bool
	rpcUrl                            string
)

// GaspriceCmd represents the gasprice command
var GaspriceCmd = &cobra.Command{
	Use:   "gas-price",
	Short: "Fetch the current gas price",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		gPrice, err := fetchGasPrice()
		if err != nil {
			log.Crit("Failed to get gas price", "error", err)
		}
		fmt.Printf("%ssuggested gas price: %s%v\n", gastParams.ColorGreen, gastParams.ColorReset, gPrice)
	},
}

func init() {
	// Flags and configuration settings.
	GaspriceCmd.Flags().BoolVarP(&eth, "eth", "", false, "Use default Ethereum RPC url")
	GaspriceCmd.Flags().BoolVarP(&op, "op", "", false, "Use default Optimism RPC url")
	GaspriceCmd.Flags().BoolVarP(&arb, "arb", "", false, "Use default Arbitrum RPC url")
	GaspriceCmd.Flags().BoolVarP(&base, "base", "", false, "Use default Base RPC url")
	GaspriceCmd.Flags().BoolVarP(&linea, "linea", "", false, "Use default Linea RPC url")
	GaspriceCmd.Flags().BoolVarP(&zkSync, "zksync", "", false, "Use default zkSync RPC URL")
	GaspriceCmd.Flags().StringVarP(&rpcUrl, "rpc-url", "u", "", "specify RPC url for gas price")

	GaspriceCmd.MarkFlagsOneRequired("eth", "op", "arb", "base", "linea", "zksync", "rpc-url")
}

func fetchGasPrice() (string, error) {
	var url string

	switch {
	case rpcUrl != "":
		url = rpcUrl
	case eth:
		url = "https://rpc.mevblocker.io"
	case op:
		url = "https://optimism.publicnode.com"
	case arb:
		url = "https://arbitrum.llamarpc.com"
	case base:
		url = "https://base.llamarpc.com"
	case linea:
		url = "https://1rpc.io/linea"
	case zkSync:
		url = "https://1rpc.io/zksync2-era"
	default:
		return "", fmt.Errorf("no network specified")
	}

	return GetCurrentGasPrice(url)
}
