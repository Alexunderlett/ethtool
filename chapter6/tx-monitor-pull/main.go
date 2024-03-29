package main

import (
  "fmt"
  "context"
  "time"
  "math/big"
  "hubwiz.com/ethtool"
)

func main(){
  fmt.Println("tx monitor demo")
  go trigger()
  filter_monitor()  
}

func trigger(){
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err)

  ctx := context.Background()  
    
  accounts,err := client.EthAccounts(ctx)
  assert(err)

  ticker := time.Tick(5 * time.Second)  
  for range ticker {
    msg := map[string]interface{}{
      "from": accounts[0],
      "to": accounts[1],
      "value": big.NewInt(1000),
    }
    txid,err := client.EthSendTransaction(ctx,msg)
    assert(err)
    fmt.Println("trigger txid: ",txid.Hex())
  }  
}

func filter_monitor(){
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err)
    
  ctx := context.Background()
    
  fid,err := client.EthNewBlockFilter(ctx)  
  assert(err)
    
  timer := time.Tick(2 * time.Second)
  for range timer {
    hashes,err := client.EthGetFilterChanges(ctx,fid)
    assert(err)
    for _,hash := range hashes {
	  //fmt.Println("monitored block hash: ", hash.Hex())
      block,err := client.EthGetBlockByHash(ctx,hash)
      assert(err)
      for _, tx := range block.Transactions() {
        fmt.Println("-> txid: ",tx.Hash().Hex())
      }
    }
  }
}

func assert(err error){
  if err != nil {
    panic(err)
  }
}