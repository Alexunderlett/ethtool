package main
//合约调用demo
import (

	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"hubwiz.com/ethtool"
	"math/big"
	"strings"
)

func main(){
	fmt.Println("access contract demo")

	client,err := ethtool.Dial("https://rpcapi.rainbow.kim")
	assert(err)

	//addrHexBytes,err := ioutil.ReadFile("../contract/build/EzToken.addr")
	//fmt.Println("contract address: ",addrHexBytes)

	assert(err)
	contractAddress := common.HexToAddress("0x76Db492d33dfE4f6D2EaD2a91D09EeCeEf0Db09F") //合约地址
	fmt.Println("contract address: ",contractAddress.Hex())

	inst,err := NewEztoken(contractAddress,client)
	assert(err)
	fmt.Println("inst: ",inst)

	credential,err := ethtool.HexToCredential("0xf06727dd488591cc5a6cdb3294c70be8397e955b3bb592a42b9fbcb0883be8c3")
	assert(err)

	txOpts := bind.NewKeyedTransactor(credential.PrivateKey)
	//
	//toAddress := common.HexToAddress("0xD0A567b968327f85122ea6430410D1aC269a6dfE")
	//amount := big.NewInt(1_000_000_000)
	//
	//tx,err := inst.Transfer(txOpts,toAddress,amount)
	////assert(err)
	////fmt.Println("txid: ",tx.Hash().Hex())
	//
	//callOpts := &bind.TransactOpts{
	//	From: credential.Address,
	//}
	//number,err := inst.Number(callOpts)
	//assert(err)
	//fmt.Println("number: ", number)
	number,err := inst.SetCompleted(txOpts,big.NewInt(1_000_000_000))
	assert(err)
	fmt.Println("number: ", number)
}

func assert(err error){
	if err != nil {
		panic(err)
	}
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

func (_Eztoken *EztokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eztoken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

func (_Eztoken *EztokenCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eztoken.contract.Call(opts, out, "number")
	return *ret0, err
}

func (_Eztoken *EztokenTransactor) SetCompleted(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Eztoken.contract.Transact(opts, "setCompleted", _value)
}

const EztokenABI = "[\n\t{\n\t\t\"constant\": false,\n\t\t\"inputs\": [\n\t\t\t{\n\t\t\t\t\"name\": \"completed\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"name\": \"setCompleted\",\n\t\t\"outputs\": [],\n\t\t\"payable\": false,\n\t\t\"stateMutability\": \"nonpayable\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"constant\": true,\n\t\t\"inputs\": [],\n\t\t\"name\": \"number\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"uint256\"\n\t\t\t}\n\t\t],\n\t\t\"payable\": false,\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t},\n\t{\n\t\t\"constant\": true,\n\t\t\"inputs\": [],\n\t\t\"name\": \"owner\",\n\t\t\"outputs\": [\n\t\t\t{\n\t\t\t\t\"name\": \"\",\n\t\t\t\t\"type\": \"address\"\n\t\t\t}\n\t\t],\n\t\t\"payable\": false,\n\t\t\"stateMutability\": \"view\",\n\t\t\"type\": \"function\"\n\t}\n]"
