package main

import (
  "fmt"
  "github.com/ethereum/go-ethereum/rpc"
  "github.com/ethereum/go-ethereum/common"
)

func main(){
  fmt.Println("rpc geth client demo")
  web3ClientVersion()
  web3Sha3()
}

func web3ClientVersion(){
  client,err := rpc.Dial("http://localhost:8545")
  assert(err);
  var result string
  err = client.Call(&result,"web3_clientVersion")  
  assert(err)
  fmt.Println("version:", result)
}

func web3Sha3(){
  client,err := rpc.Dial("http://localhost:8545")
  assert(err);
  var result string
  data := common.ToHex([]byte("hello,ethereum"))
  err = client.Call(&result,"web3_sha3",data)  
  assert(err)
  fmt.Println("keccak256 hash: ",result)
}


func assert(err error) {
  if err != nil {
	panic(err)
  }
}