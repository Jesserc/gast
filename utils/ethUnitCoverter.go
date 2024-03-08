package utils

import (
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/params"
)

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
