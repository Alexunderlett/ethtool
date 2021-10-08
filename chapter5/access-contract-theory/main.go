package main

import (
  "context"
  "fmt"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "hubwiz.com/ethtool"
  "io/ioutil"
  "math/big"
  "strings"
)

func main(){
  fmt.Println("access contract theory demo")
  
  abiBytes,err := ioutil.ReadFile("../contract/build/EzToken.abi")
  assert(err)
  tokenAbi,err := abi.JSON(strings.NewReader(string(abiBytes)))
  assert(err)
  
  addrBytes,err := ioutil.ReadFile("../contract/build/EzToken.addr")
  fmt.Println("contract11111: ", addrBytes)
  assert(err)
  contractAddress := common.HexToAddress(string(addrBytes))
  fmt.Println("contract address: ", contractAddress.Hex())
    
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err)
  
  ctx := context.Background()
    
  accounts,err := client.EthAccounts(ctx)
  assert(err)
    
  //transfer from account 0 => 1
    
  data,err := tokenAbi.Pack(
    "transfer",
    accounts[1],
    big.NewInt(100),
  )
  assert(err)
  msg := map[string]interface{}{
    "from": accounts[0],
    "to": contractAddress,
    "gas": big.NewInt(2000000),
    "data": common.ToHex(data),
  }
  txid,err := client.EthSendTransaction(ctx,msg)
  assert(err)
  fmt.Println("txid: ",txid.Hex())
   
    
  //balanceOf
  data,err = tokenAbi.Pack(
    "balanceOf",
    accounts[0],
  )
  msg = map[string]interface{}{
    "from": accounts[0],
    "to": contractAddress,
    "gas": big.NewInt(2000000),
    "data": common.ToHex(data),
  }
  ret,err := client.EthCall(ctx,msg)
  assert(err)
  fmt.Println("balance: ",ret)
  
  //abi decode balance
  balancePtr := new(*big.Int)
  err = tokenAbi.Unpack(balancePtr,"balanceOf",ret)
  assert(err)
  fmt.Println("balance decoded: ", *balancePtr)
}

func assert(err error) {
  if err != nil {
    panic(err)
  }
}