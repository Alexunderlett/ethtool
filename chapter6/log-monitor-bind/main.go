package main

import (
  "fmt"
  "time"
  "io/ioutil"
  "math/big"
  "contract/wrapper/eztoken"
  "hubwiz.com/ethtool"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main(){
  fmt.Println("monitor contract log with wrapper")
  go trigger()
  monitor()
}

func trigger(){
  addrHexBytes,err := ioutil.ReadFile("../contract/build/EzToken.addr")
  assert(err)
  contractAddress := common.HexToAddress(string(addrHexBytes))  
  assert(err)
    
  client,err := ethtool.Dial("ws://localhost:8545")
  assert(err)
    
  inst,err := eztoken.NewEztoken(contractAddress,client)
  assert(err)
  //fmt.Println(inst)
  
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

func monitor(){
  addrHexBytes,err := ioutil.ReadFile("../contract/build/EzToken.addr")
  assert(err)
  contractAddress := common.HexToAddress(string(addrHexBytes))  
  assert(err)
    
  client,err := ethtool.Dial("ws://localhost:8545")
  assert(err)
    
  inst,err := eztoken.NewEztoken(contractAddress,client)
  assert(err)
 
  watchOpts := &bind.WatchOpts{}  
  events := make(chan *eztoken.EztokenTransfer)
  _from := []common.Address{}
  _to := []common.Address{}
  sub,err := inst.WatchTransfer(watchOpts,events,_from,_to)
  assert(err)
  //fmt.Println(sub) 
    
  for{
    select {
    case err := <- sub.Err():
      panic(err)
    case event := <- events:
      fmt.Println("captured:")
      fmt.Println("-> from: ",event.From.Hex())
      fmt.Println("-> to: ",event.To.Hex())
      fmt.Println("-> value:",event.Value)
    }
  }  
    
}

func assert(err error){
  if err != nil {
    panic(err)
  }
}