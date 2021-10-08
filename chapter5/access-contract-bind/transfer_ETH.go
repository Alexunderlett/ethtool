package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"hubwiz.com/ethtool"
	"io/ioutil"
	"math/big"
	"os"
)

func main() {
	filePtr, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	type (
		Config struct {
			Rpc        string   `json:"rpc"`
			PrivateKey string   `json:"privateKey"`
			Amount     float64     `json:"amount"`
			GasLimit   uint     `json:"gasLimit"`
			GasPrice  float64   `json:"gasPrice"`
			Address    []string `json:"address"`
		}
	)
	var config Config
	// 创建json解码器
	byteValue, _ := ioutil.ReadAll(filePtr)
	json.Unmarshal(byteValue, &config)

	credential, err := ethtool.HexToCredential(config.PrivateKey)
	assert(err)
	fmt.Println("from address: ", credential.Address.Hex())
	client, err := ethtool.Dial("https://rpcapi.rainbow.kim")
	assert(err)

	ctx := context.Background()

	chainid, err := client.NetVersion(ctx)
	assert(err)
	var length = len(config.Address)
	var lastnoce uint64
	for i := 0; i < length; i++ {
		to := common.HexToAddress(config.Address[i])
		fmt.Println("==================", "transfer to", config.Address[i], "================")
		nonce, err := client.EthGetTransactionCount(ctx, credential.Address, "pending")

		if lastnoce >= nonce {
			nonce = lastnoce + 1
		}
		lastnoce = nonce
		assert(err)
		fmt.Println("nonce: ", nonce)
		tx := types.NewTransaction(
			nonce,
			to,
			ethtool.ToWei(big.NewFloat(config.Amount), 1e18),
			uint64(config.GasLimit),
			ethtool.ToWei(big.NewFloat(config.GasPrice), 1e9),
			nil,
		)
		signedTx, err := credential.SignTx(tx, chainid)
		assert(err)
		fmt.Println("signedTx: ", signedTx)

		buf := new(bytes.Buffer)
		err = signedTx.EncodeRLP(buf)
		assert(err)
		txid, err := client.EthSendRawTransaction(ctx, buf.Bytes())
		assert(err)
		fmt.Println("raw tx id: ", txid.Hex())

		balance, err := client.EthGetBalance(ctx, to, "latest")
		assert(err)
		fmt.Println("balance received: ", balance)
	}
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
