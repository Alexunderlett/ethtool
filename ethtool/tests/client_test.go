package tests

import (
  "testing"
  "context"
  "math/big"
  "hubwiz.com/ethtool"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/core/types"
  _ "github.com/ethereum/go-ethereum"
)

func TestClient(t *testing.T){
  t.Skip("later test")
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err,t)
  t.Logf("%+v",client)
    
  ctx := context.Background()
    
  chainid,err := client.NetworkID(ctx)
  assert(err,t)
  t.Log("chainid: ",chainid)
    
  version,err := client.Web3ClientVersion(ctx)
  assert(err,t)
  t.Log("web3_clientVersion: ",version)
    
  hash,err := client.Web3Sha3(ctx,"0x68656c6c6f20776f726c64")  
  assert(err,t)
  t.Log("web3_sha3: ",hash)
    
  netversion,err := client.NetVersion(ctx)
  assert(err,t)
  t.Log("net_version: ",netversion)
    
  netlistening,err := client.NetListening(ctx)
  assert(err,t)
  t.Log("net_listening: ",netlistening)
    
  netpeercount,err := client.NetPeerCount(ctx)
  assert(err,t)
  t.Log("net_peerCount: ",netpeercount)
    
  protocolversion,err := client.EthProtocolVersion(ctx)
  assert(err,t)
  t.Log("eth_protocolVersion: ",protocolversion)
    
  coinbase,err := client.EthCoinbase(ctx)
  assert(err,t)
  t.Log("eth_coinbase: ",coinbase.Hex())
    
  mining,err := client.EthMining(ctx)
  assert(err,t)
  t.Log("eth_mining: ",mining)
    
  hashrate,err := client.EthHashrate(ctx)
  assert(err,t)
  t.Log("eth_hashrate: ", hashrate)
    
  gasprice,err := client.EthGasPrice(ctx)
  assert(err,t)
  t.Log("eth_gasPrice: ",gasprice)
    
  accounts,err := client.EthAccounts(ctx)
  assert(err,t)
  t.Log("eth_accounts: ",accounts)
    
  balance,err := client.EthGetBalance(ctx,accounts[0],nil)
  assert(err,t)
  t.Log("eth_getBalance: ",balance)
  
  /*  
  storage,err := client.EthGetStorageAt(ctx,accounts[0],common.HexToHash("0x0"),nil)
  assert(err,t)
  t.Log("eth_getStorageAt: ",storage)
  */
    
  txcount,err := client.EthGetTransactionCount(ctx,accounts[0],"latest")
  assert(err,t)
  t.Log("eth_getTransactionCount: ",txcount)
   
  /*  
  blocktxcount,err := client.EthGetBlockTransactionCountByHash(ctx,common.HexToHash("0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"))
  assert(err,t)
  t.Log("eth_getBlockTransactionCountByHash: ",blocktxcount)
  */
    
  blocktxcount2,err := client.EthGetBlockTransactionCountByNumber(ctx,0)
  assert(err,t)
  t.Log("eth_getBlockTransactionCountByNumber: ",blocktxcount2)
    
  //unclecount1
    
  unclecount2,err := client.EthGetUncleCountByBlockNumber(ctx,"latest")
  assert(err,t)
  t.Log("eth_getUncleCountByBlockNumber: ",unclecount2)
    
  code,err := client.EthGetCode(ctx,accounts[0],"latest")
  assert(err,t)
  t.Log("eth_getCode: ",code)
    
  signature,err := client.EthSign(ctx,accounts[0],[]byte("this is a demo"))
  assert(err,t)
  t.Log("eth_sign: ",signature)
    
  msg := map[string]interface{}{
    "from": accounts[0],
    "to": &accounts[1],
    "value": big.NewInt(10000),
  }
  txid,err := client.EthSendTransaction(ctx,msg)
  assert(err,t)
  t.Log("eth_sendTransaction: ",txid)
    
  /*  
  txid2,err := client.EthSendRawTransaction(ctx,[]byte{1,2,3,4,5})  
  assert(err,t)
  t.Log("eth_sendRawTransaction: ",txid2)
  */
  
  /*  
  callret,err := client.EthCall(ctx,stx)
  assert(err,t)
  t.Log("eth_call: ",callret)  
  */
    
  estimation,err := client.EthEstimateGas(ctx,msg)
  assert(err,t)
  t.Log("eth_estimageGas: ",estimation)
}

func TestCredential(t *testing.T){
  t.Skip("later test credential")
  c1,err := ethtool.NewCredential()
  assert(err,t)
  t.Log("credential: ",c1)
    
  c2,err := ethtool.HexToCredential("4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
  assert(err,t)
  t.Log("address: ", c2.Address.Hex())
  
  msg := "this is a demo"   
  signature,err := c2.Sign([]byte(msg))
  assert(err,t)
  t.Log("signature: ",signature)
    
  verified := c2.Verify([]byte(msg),signature)
  t.Log("verified: ",verified)
    
  tx := types.NewTransaction(
    uint64(1),
    common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"),
    big.NewInt(100000),
    uint64(21000),
    big.NewInt(2000000000),
    nil,
  )
  chainid3 := big.NewInt(111)
  signedTx,err := c2.SignTx(tx,chainid3)
  assert(err,t)
  t.Logf("signed tx: %+v\n",signedTx)
}

func TestRawTransaction(t *testing.T) {
  t.Skip("later review")
  client,err := ethtool.Dial("http://localhost:8545")
  assert(err,t)
  t.Log("client: ",client)
    
  credential,err := ethtool.HexToCredential("4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
  assert(err,t)
  t.Log("address: ", credential.Address.Hex())
  
  ctx := context.Background()
    
  nonce,err := client.PendingNonceAt(ctx,credential.Address)
  assert(err,t)
  t.Log("nonce: ",nonce)
    
  chainid,err := client.NetworkID(ctx)
  assert(err,t)
  t.Log("chainid: ",chainid)
    
  tx := types.NewTransaction(
    nonce,
    common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"),
    big.NewInt(100000),
    uint64(21000),
    big.NewInt(2000000000),
    nil,
  )
  signedTx,err := credential.SignTx(tx,chainid)  
  assert(err,t)
  t.Log("signed tx: ", signedTx)  
  
  err = client.SendTransaction(ctx,signedTx)
  assert(err,t)
  t.Log("txid: ", signedTx.Hash().Hex())
}

func TestValueUnit(t *testing.T) {
  value := big.NewInt(100000000000)
  t.Log("value in wei: ", value)
  t.Log("value in gwei: ", ethtool.FromWei(value,ethtool.Gwei))
    
  fvalue := big.NewFloat(123.56)
  t.Log("value in ether: ",fvalue)
  t.Log("value in wei: ", ethtool.ToWei(fvalue,ethtool.Ether))  
}

func assert(err error, t *testing.T) {
  if err != nil {
    t.Fatal(err)
  }
}
