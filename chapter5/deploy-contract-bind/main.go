package main

import (
  "fmt"
  "math/big"
  "io/ioutil"
  "hubwiz.com/ethtool"
  "contract/wrapper/eztoken"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main(){
  fmt.Println("deploy contract demo")
  
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err)
    
  credential,err := ethtool.HexToCredential("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")  
  assert(err)

  txOpts := bind.NewKeyedTransactor(credential.PrivateKey)  
   
  tokenSupply := big.NewInt(1000000)
  tokenName := "HAPPY TOKEN"
  tokenDecimals := uint8(0)
  tokenSymbol := "HAPY"
  address,tx,inst,err := eztoken.DeployEztoken(txOpts,client,tokenSupply,tokenName,tokenDecimals,tokenSymbol)
  assert(err)
  fmt.Println("deployed at: ",address.Hex())
  fmt.Println("txid: ", tx.Hash().Hex())
  _ = inst
    
  fmt.Println("save deployed address...")  
  err = ioutil.WriteFile("../contract/build/EzToken.addr",[]byte(address.Hex()),0644)  
  assert(err)
    
  fmt.Println("done.")  
}

func assert(err error) {
  if err != nil {
    panic(err)
  }
}

