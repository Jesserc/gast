/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createContractCmd represents the createContract command
var createContractCmd = &cobra.Command{
	Use:   "createContract",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createContract called")
	},
}

func init() {

}
