/*
Copyright © 2024 NAME HERE <raymondjesse713@gmail.com>
*/

package gasprice

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	eth, op, base, linea, arb, zkSync bool
	rpcUrl                            string
)

// GaspriceCmd represents the gasprice command
var GaspriceCmd = &cobra.Command{
	Use:   "gas-price",
	Short: "Get the current gas price",
	Long:  "Get the current gas price",
	Run: func(cmd *cobra.Command, args []string) {
		// Check if any flag is provided
		if !(eth || op || arb || base || linea || zkSync) && rpcUrl == "" {
			// If no flag is provided, print usage and return
			cmd.Usage()
			return
		}

		// If flags are provided, get gas price
		gPrice, err := fetchGasPrice()
		if err != nil {
			log.Printf("%v: run gast help gasprice for all commands\n", err)
			return
		}
		fmt.Printf("Current gas price: %v\n", gPrice)
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
	GaspriceCmd.Flags().StringVarP(&rpcUrl, "url", "u", "", "specify RPC url for gas price")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gaspriceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gaspriceCmd.Flags().BoolP("toggle", "handleTraceTx", false, "Help message for toggle")
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
