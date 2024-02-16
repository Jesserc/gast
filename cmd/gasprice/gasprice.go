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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	GaspriceCmd.Flags().StringVarP(&rpcUrl, "url", "u", "", "The rpc url for gas price")
	GaspriceCmd.Flags().BoolVarP(&eth, "eth", "e", false, "Use default Ethereum rpc url")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gaspriceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gaspriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
