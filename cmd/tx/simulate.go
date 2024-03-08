/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

// simulateCmd represents the simulate command
var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Simulate an Ethereum transaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("simulate command called") // TODO: implement logic in handleSimulate.go
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// simulateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// simulateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
