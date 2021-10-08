package main

import (
  "fmt"
  "io/ioutil"
  "math/big"
  "context"
  "time"
  "hubwiz.com/ethtool"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum"
  "contract/wrapper/eztoken"
)

func main(){
  fmt.Println("log monitor demo")
  go trigger()  
  push_monitor()
}

func trigger(){
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err)
    
  data,err := ioutil.ReadFile("../contract/build/EzToken.addr") 
  assert(err)
  contractAddress := common.HexToAddress(string(data))
  inst,err := eztoken.NewEztoken(contractAddress,client)
  assert(err)  
        
  credential,err := ethtool.HexToCredential("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")  
  assert(err)
  txOpts := bind.NewKeyedTransactor(credential.PrivateKey)
    
  toAddress := common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0")
  amount := big.NewInt(100)
    
  timer := time.Tick(5 * time.Second)
  for range timer {
    tx,err := inst.Transfer(txOpts,toAddress,amount)  
    assert(err)
    fmt.Println("trigger txid: ",tx.Hash().Hex())  
  }
}

func push_monitor(){
  client,err := ethtool.Dial("ws://localhost:8545")
  assert(err)
  
  data,err := ioutil.ReadFile("../../chapter5/contract/build/EzToken.addr") 
  assert(err)
  contractAddress := common.HexToAddress(string(data))
  _ = contractAddress
    
  query := ethereum.FilterQuery{
//    Addresses: []common.Address{contractAddress},
  }  
  logs := make(chan types.Log)  
  sub,err := client.SubscribeFilterLogs(context.Background(),query,logs)  
  assert(err)
    
  for {
    select {
      case err := <- sub.Err():
        panic(err)
      case log := <- logs:
        fmt.Println("captured log:")
        fmt.Println("-> address: ",log.Address.Hex())
        fmt.Println("-> data: ",log.Data)
        fmt.Println("-> topics: ",log.Topics)
    }
  }  
}

func assert(err error){
  if err != nil {
    panic(err)
  }
}

