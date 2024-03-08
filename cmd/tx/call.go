/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

// callCmd represents the call command
var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Make view call to a contract",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("call command called") // TODO: implement logic in handleCall.go
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// callCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// callCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
