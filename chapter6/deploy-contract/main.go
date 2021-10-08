package main

import (
  "fmt"
  "context"
  "math/big"
  "io/ioutil"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "contract/wrapper/eztoken"
)

func main(){
  fmt.Println("deploy contract demo")
  
  client,err := ethclient.Dial("http://localhost:8545")
  assert(err)
    
  //account
  fromPrvKeyHex := "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
  fromAddressHex := "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1"
    
  fromPrvKey,err := crypto.HexToECDSA(fromPrvKeyHex[2:])  
  assert(err)
  fromAddress := common.HexToAddress(fromAddressHex)
  
  nonce,err := client.PendingNonceAt(context.Background(),fromAddress)  
  assert(err)
    
  auth := bind.NewKeyedTransactor(fromPrvKey)
  auth.Nonce = big.NewInt(int64(nonce))
  auth.Value = big.NewInt(0)
  auth.GasLimit = uint64(3000000)
  auth.GasPrice = big.NewInt(2000000000)
   
  initialAmount := big.NewInt(1000000)
  tokenName := "HAPPY TOKEN"
  decimalUnits := uint8(0)
  tokenSymbol := "HAPY"
  address,tx,inst,err := eztoken.DeployEztoken(auth,client,initialAmount,tokenName,decimalUnits,tokenSymbol)
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

