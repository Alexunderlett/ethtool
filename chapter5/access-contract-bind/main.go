package main

import (
  "contract/wrapper/eztoken"
  "fmt"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "hubwiz.com/ethtool"
  "math/big"
)

func main(){
  fmt.Println("access contract demo")
  
  client,err := ethtool.Dial("https://rpcapi.rainbow.kim")
  assert(err)

  //addrHexBytes,err := ioutil.ReadFile("../contract/build/EzToken.addr")
  //fmt.Println("contract address: ",addrHexBytes)

  assert(err)
  contractAddress := common.HexToAddress("0x02e6a91543aC7139ED1c29b526b3fF03dC2C03d9") //代币地址
  fmt.Println("contract address: ",contractAddress.Hex())

  inst,err := eztoken.NewEztoken(contractAddress,client)
  assert(err)
  fmt.Println("inst: ",inst)

  credential,err := ethtool.HexToCredential("0xf06727dd488591cc5a6cdb3294c70be8397e955b3bb592a42b9fbcb0883be8c3")
  assert(err)

  txOpts := bind.NewKeyedTransactor(credential.PrivateKey)

  toAddress := common.HexToAddress("0xD0A567b968327f85122ea6430410D1aC269a6dfE")
  amount := big.NewInt(1_000_000_000)

  tx,err := inst.Transfer(txOpts,toAddress,amount)
  assert(err)
  fmt.Println("txid: ",tx.Hash().Hex())

  callOpts := &bind.CallOpts{
    From: credential.Address,
  }
  balance,err := inst.BalanceOf(callOpts, toAddress)
  assert(err)
  fmt.Println("balance: ", balance)
    
}

func assert(err error){
  if err != nil {
    panic(err)
  }
}