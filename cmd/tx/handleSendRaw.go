package transaction

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func sendRawTransaction(rawTx string) error {
	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return err
	}

	tx := new(types.Transaction)

	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return err
	}

	client, err := ethclient.Dial("https://goerli.gateway.tenderly.co")
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
	return nil
}
