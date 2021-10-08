package main

import (
  "context"
  "fmt"
  "hubwiz.com/ethtool"
  "math/big"
)

func main(){
  fmt.Println("transaction receipt demo")
  
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err)
    
  ctx := context.Background()
    
  accounts,err := client.EthAccounts(ctx)
    
  msg := map[string]interface{}{
    "from": accounts[0],
    "to": accounts[1],
    "value": big.NewInt(1e9),
  }  
  //txid,err := client.EthSendTransaction(ctx,msg)
  //assert(err)
  //fmt.Println("txid: ",txid.Hex())
  //
  ////wait...
  //
  //receipt,err := client.EthGetTransactionReceipt(ctx,txid)
  //assert(err)
  //fmt.Printf("receipt: %+v\n",receipt)
  estimation,err := client.EthEstimateGas(ctx,msg)
  assert(err)
  fmt.Println("eth_estimageGas: ",estimation)

  gasPrice,err := client.EthGasPrice(ctx)
  assert(err)
  fmt.Println("gasPrice: ",gasPrice)
}

func assert(err error){
  if err != nil {
    panic(err)
  }
}