package ethclient

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestGetBlock(t *testing.T) {
	url := "http://18.168.16.120:32915"
	client, err := Dial(url)
	if err != nil {
		t.Errorf("Dial failed, err: %v", err)
	}
	if blk, err := client.BlockByNumber(context.Background(), big.NewInt(33653)); err != nil {
		t.Errorf("BlockByNumber failed, err: %v", err)
	} else {
		t.Logf("BlockByNumber success, block: %v", blk)
	}
}

func TestGetTx(t *testing.T) {
	hash := common.HexToHash("0xa5f84d5fdb38836f48e0cbf20db97e4649ba4d8d9a39e7042cdf4c303fd66fc5")
	url := "https://achain.bitheart.org"
	client, err := Dial(url)
	if err != nil {
		t.Errorf("Dial failed, err: %v", err)
	}
	if tx, _, err := client.TransactionByHash(context.Background(), hash); err != nil {
		t.Errorf("TransactionByHash failed, err: %v", err)
	} else {
		data, _ := tx.MarshalBinary()
		t.Log("raw tx :", common.Bytes2Hex(data))
	}
}
