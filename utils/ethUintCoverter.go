package utils

import (
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/params"
)

func EthConversion(wei uint64, denomination string, precision int) (string, error) {
	weiValue := new(big.Int).SetUint64(wei)

	var value *big.Float
	var v string

	dLower := strings.ToLower(denomination)
	switch dLower {
	case "eth":
		value = new(big.Float).Quo(new(big.Float).SetInt(weiValue), new(big.Float).SetFloat64(params.Ether))
		v = value.Text('f', 18)
	case "gwei":
		value = new(big.Float).Quo(new(big.Float).SetInt(weiValue), new(big.Float).SetFloat64(params.GWei))
		v = value.Text('f', precision)
	case "wei":
		v = strconv.FormatUint(wei, 10)
	default:
		err := errors.New("denomination not supported: " + denomination)
		if err != nil {
			return "", err
		}
	}
	
	return v, nil
}
