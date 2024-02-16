/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package gasprice

import (
	"github.com/spf13/cobra"
)

var (
	rpcUrl string
	eth,
	op,
	base,
	linea,
	arb,
	zkSync bool
)

// GaspriceCmd represents the gasprice command
var GaspriceCmd = &cobra.Command{
	Use:   "gasprice",
	Short: "Get the current gas price",
	Long:  "Get the current gas price",
	Run: func(cmd *cobra.Command, args []string) {
		if rpcUrl != "" {
			handleRPCCommand(rpcUrl)
		} else if eth {
			handleRPCCommand("https://rpc.mevblocker.io")
		} else if op {
			handleRPCCommand("https://optimism.publicnode.com")
		} else if base {
			handleRPCCommand("https://base.llamarpc.com")
		} else if arb {
			handleRPCCommand("https://arbitrum.llamarpc.com")
		} else if linea {
			handleRPCCommand("https://1rpc.io/linea")
		} else if zkSync {
			handleRPCCommand("https://1rpc.io/zksync2-era")
		} else {
			cmd.Usage()
		}
	},
}

func init() {
	// Flags and configuration settings.
	GaspriceCmd.Flags().StringVarP(&rpcUrl, "url", "u", "", "The RPC url for gas price")
	GaspriceCmd.Flags().BoolVarP(&eth, "eth", "", false, "Use default Ethereum RPC url")
	GaspriceCmd.Flags().BoolVarP(&op, "op", "", false, "Use default Optimism RPC url")
	GaspriceCmd.Flags().BoolVarP(&arb, "arb", "", false, "Use default Arbitrum RPC url")
	GaspriceCmd.Flags().BoolVarP(&base, "base", "", false, "Use default Base RPC url")
	GaspriceCmd.Flags().BoolVarP(&linea, "linea", "", false, "Use default Linea RPC url")
	GaspriceCmd.Flags().BoolVarP(&zkSync, "zksync", "", false, "Use default zkSync RPC URL")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gaspriceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gaspriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
