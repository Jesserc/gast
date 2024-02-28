/*
Copyright © 2024 NAME HERE <raymondjesse713@gmail.com>
*/

package gasprice

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ctx = context.Background()

	// Define a map for chain IDs and corresponding network names
	networkNames = map[uint64]string{
		0x01:     "Ethereum Mainnet",
		0x05:     "Goerli Testnet",
		0xAA36A7: "Sepolia",
		0x89:     "Polygon Mainnet",
		0x13881:  "Polygon Mumbai Testnet",
		0x0A:     "Optimism Mainnet",
		0x1A4:    "Optimism Goerli Testnet",
		0xA4B1:   "Arbitrum One Mainnet",
		0x66EED:  "Arbitrum Goerli Testnet",
		0x2105:   "Base Mainnet",
		0xE708:   "Linea Mainnet",
		0x144:    "zkSync Mainnet",
	}
)

func GetCurrentGasPrice(rpcUrl string) (string, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return "", err
	}

	chainIdBigInt, err := client.ChainID(ctx)
	if err != nil {
		return "", err
	}

	chainId := chainIdBigInt.Uint64()

	// Retrieve the network name from the map and print
	if networkName, ok := networkNames[chainId]; ok {
		fmt.Printf("Retrieving Gas Price on %s\n", networkName)
	} else {
		fmt.Printf("Retrieving Gas Price on network with chain ID: 0x%x\n", chainId)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	return gasPrice.String(), nil
}
