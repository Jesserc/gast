/*
Copyright © 2024 NAME HERE <raymondjesse713@gmail.com>
*/

package gasprice

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
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
		return "", fmt.Errorf("failed to dial RPC client: %s", err)
	}

	chainIdBigInt, err := client.ChainID(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get chain ID: %s", err)
	}

	chainId := chainIdBigInt.Uint64()

	// Retrieve the network name from the map and print
	if networkName, ok := networkNames[chainId]; ok {
		log.Info("Retrieving Gas Price", "network", networkName)
	} else {
		log.Info("Retrieving Gas Price", "network with chain ID", hexutil.EncodeUint64(chainId))
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to get gas price: %s", err)
	}

	return gasPrice.String(), nil
}

/*
func EthConversion(wei uint64, denomination string) (string, error) {
	weiValue := new(big.Int).SetUint64(wei)

	var value *big.Float
	var v string

	dLower := strings.ToLower(denomination)
	switch dLower {
	case "eth":
		value = new(big.Float).Quo(new(big.Float).SetInt(weiValue), new(big.Float).SetFloat64(params.Ether))
		v = value.Text('f', 18)
		v = strings.TrimRight(v, "0")
		v = strings.TrimRight(v, ".")
	case "gwei":
		value = new(big.Float).Quo(new(big.Float).SetInt(weiValue), new(big.Float).SetFloat64(params.GWei))
		v = value.Text('f', 9)
		v = strings.TrimRight(v, "0")
		v = strings.TrimRight(v, ".")
	case "wei":
		v = strconv.FormatUint(wei, 10)
	default:
		err := errors.New("denomination not supported: " + denomination)
		return "", err
	}

	return v, nil
}



func TestEthConversion(t *testing.T) {
	v, err := EthConversion(10e18, "eth")
	require.NoError(t, err, "error should be nil")
	require.Equal(t, "10", v, "10e18 Wei should equal 10 ETH")

	v, err = EthConversion(10e18, "gwei")
	require.NoError(t, err, "error should be nil")
	require.Equal(t, "10000000000", v, "10e18 Wei should equal 10000000000 Gwei")

	v, err = EthConversion(10e18, "wei")
	require.NoError(t, err, "error should be nil")
	require.Equal(t, "10000000000000000000", v, "10e18 Wei should equal 10000000000000000000 Wei")

	v, err = EthConversion(10e18, "invalid")
	require.Error(t, err, "error should not be nil")
	require.ErrorContains(t, err, "denomination not supported")
}

*/
