/*
Copyright © 2024 NAME HERE <raymondjesse713@gmail.com>
*/

package main

import (
	"os"

	"github.com/Jesserc/gast/cmd"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	cmd.Execute()
}

func init() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
}
