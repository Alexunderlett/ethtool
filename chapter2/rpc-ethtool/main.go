package main

import (
  "fmt"
  "context"
  "hubwiz.com/ethtool"
)

func main(){
  fmt.Println("ethtool demo")
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err)
  
  accounts,err := client.EthAccounts(context.Background())
  assert(err)
  fmt.Println("accounts: ", accounts)
    
  balance,err := client.EthGetBalance(context.Background(),accounts[0],"latest")  
  assert(err)
  fmt.Println("account#0 balance: ",balance)
}

func assert(err error) {
  if err != nil {
	panic(err)
  }
}