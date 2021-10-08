package main

import (

	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"hubwiz.com/ethtool"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	result struct {
		Messages []string `json:"messages"`
		Signatures []string `json:"signatures"`
		Prices map[string]string `json:"prices"`
		Timestamp string `json:"timestamp"`
	}
)

func main() {
	fmt.Println("start")
	strCoin := [13]string{"BTC", "ETH", "XTZ", "DAI", "REP", "ZRX", "BAT", "KNC", "LINK", "COMP", "UNI", "GRT", "SNX"}
	var data result
	proxyAddr := "http://127.0.0.1:11000"
	url := "https://api.pro.coinbase.com/oracle"
	cli := NewHttpClient(proxyAddr)
	data2,_ := HttpGET(cli, url)
	fmt.Println(string(data2))
	json.Unmarshal(data2, &data)

	length := len(data.Messages)

	client,err := ethtool.Dial("https://rpcapi.rainbow.kim")
	assert(err)

	contractAddress := common.HexToAddress("0x00c4770D3Feb38ad07f879Abd96619FBdeb00520") //合约地址
	sourceAddress := common.HexToAddress("0xfCEAdAFab14d46e20144F48824d0C09B1a03F2BC") //source地址

	fmt.Println("contract address: ",contractAddress.Hex())

	inst,err := NewEztoken(contractAddress,client)
	assert(err)
	//导入私钥
	credential,err := ethtool.HexToCredential("0xf06727dd488591cc5a6cdb3294c70be8397e955b3bb592a42b9fbcb0883be8c3")
	assert(err)
	//打包请求源（账户，gas）
	txOpts := bind.NewKeyedTransactor(credential.PrivateKey)
	txOpts.GasLimit = 3000000

	//字符串转uint64
	intNum, _ := strconv.Atoi(data.Timestamp)
	timestamps := uint64(intNum)
	//循环喂价
	for i:=0; i<length; i++ {
		//fmt.Println(data.Prices[strCoin[i]])
		float,err := strconv.ParseFloat(data.Prices[strCoin[i]],64)
		inst.Puts(txOpts,sourceAddress,timestamps,strCoin[i],uint64(float*1_000_000_000_000))
		assert(err)
		fmt.Println("喂价成功: ", strCoin[i])
	}

}

//合约方法
func (_Eztoken *EztokenTransactor) Puts(opts *bind.TransactOpts,_source common.Address, _time uint64, _key string, _value uint64) (*types.Transaction, error) {
	return _Eztoken.contract.Transact(opts, "putInternal", _source, _time,_key,_value)
}

func assert(err error){
	if err != nil {
		panic(err)
	}
}

func NewHttpClient(proxyAddr string) *http.Client {
	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		return nil
	}

	netTransport := &http.Transport{
		//Proxy: http.ProxyFromEnvironment,
		Proxy: http.ProxyURL(proxy),
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*time.Duration(10))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * time.Duration(5),
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
}

func HttpGET(client *http.Client, url string) (body []byte, err error) {
	rsp, err := client.Get(url)
	if err != nil {
		return
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK || err != nil{
		err = fmt.Errorf("HTTP GET Code=%v, URI=%v, err=%v", rsp.StatusCode, url, err)
		return
	}
	fmt.Println(rsp.Body)
	return ioutil.ReadAll(rsp.Body)
}
type Eztoken struct {
	EztokenCaller     // Read-only binding to the contract
	EztokenTransactor // Write-only binding to the contract
	EztokenFilterer   // Log filterer for contract events
}
type EztokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EztokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EztokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EztokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EztokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

func NewEztoken(address common.Address, backend bind.ContractBackend) (*Eztoken, error) {
	contract, err := bindEztoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eztoken{EztokenCaller: EztokenCaller{contract: contract}, EztokenTransactor: EztokenTransactor{contract: contract}, EztokenFilterer: EztokenFilterer{contract: contract}}, nil
}

func bindEztoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EztokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

const EztokenABI = "[\n\t{\n\t\t\"inputs\": [],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"constructor\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"priorTimestamp\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"messageTimestamp\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint256\",\n\t\t\t\t\"name\": \"blockTimestamp\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"NotWritten\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes\",\n\t\t\t\t\"name\": \"message\",\n\t\t\t\t\"type\": \"bytes\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes\",\n\t\t\t\t\"name\": \"signature\",\n\t\t\t\t\"type\": \"bytes\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"put\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"string\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"string\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"source\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"timestamp\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"string\",\n\t\t\t\t\"name\": \"key\",\n\t\t\t\t\"type\": \"string\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"value\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"putInternal\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"string\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"string\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"anonymous\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"indexed\": true,\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"source\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"string\",\n\t\t\t\t\"name\": \"key\",\n\t\t\t\t\"type\": \"string\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"timestamp\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"indexed\": false,\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"value\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"Write\",\n\t\t\"type\": \"event\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"source\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"string\",\n\t\t\t\t\"name\": \"key\",\n\t\t\t\t\"type\": \"string\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"get\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"source\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"string\",\n\t\t\t\t\"name\": \"key\",\n\t\t\t\t\"type\": \"string\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"getPrice\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"uint64\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint64\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [],\n\t\t\"name\": \"owner\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes\",\n\t\t\t\t\"name\": \"message\",\n\t\t\t\t\"type\": \"bytes\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"internalType\": \"bytes\",\n\t\t\t\t\"name\": \"signature\",\n\t\t\t\t\"type\": \"bytes\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"source\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"internalType\": \"address\",\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"stateMutability\": \"pure\",\n\t\t\"type\": \"function\"\n\t}\n]"
