package main

import (
  "bytes"
  "context"
  "fmt"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  "hubwiz.com/ethtool"
  "math/big"
)

func main(){
  credential,err := ethtool.HexToCredential("0xe872122c04df93040ede8996c0e738f35a0ea44e77642d97eb5c3deedbdd4201")
  assert(err)
  fmt.Println("from address: ",credential.Address.Hex())

  to := common.HexToAddress("0x40DE77004FA015c6D172D3F2C6b411a8db17eCfb")
  fmt.Println("to address: ", to.Hex())
      
  client,err := ethtool.Dial("https://rpc.xdaichain.com")
  assert(err)
    
  ctx := context.Background()
      
  chainid,err := client.NetVersion(ctx)
  assert(err)
  fmt.Println("chainid: ",chainid)
  
  nonce,err := client.EthGetTransactionCount(ctx,credential.Address,"pending")  
  assert(err)
  fmt.Println("nonce: ",nonce)
    
  tx := types.NewTransaction(
    nonce,
    to,
    big.NewInt(1*10e18),
    uint64(21000),
    big.NewInt(2e9),
    nil,
  )
  signedTx,err := credential.SignTx(tx,chainid)
  assert(err)
  fmt.Println("signedTx: ",signedTx)
  
  /*  
  err = client.SendTransaction(ctx,signedTx)
  assert(err)
  fmt.Println("raw tx id: ",signedTx.Hash().Hex())
  */
    
  buf := new(bytes.Buffer)
  err = signedTx.EncodeRLP(buf)
  assert(err)
  txid,err := client.EthSendRawTransaction(ctx,buf.Bytes())
  assert(err)
  fmt.Println("raw tx id: ",txid.Hex())
    
  balance,err := client.EthGetBalance(ctx,to,"latest")
  assert(err)
  fmt.Println("balance received: ",balance)
}


func assert(err error) {
  if err != nil {
    panic(err)
  }
}

