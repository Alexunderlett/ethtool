package main
//http请求demo
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
	contractAddress := common.HexToAddress("0x02e6a91543aC7139ED1c29b526b3fF03dC2C03d9") //代币地址
	fmt.Println("contract address: ",contractAddress.Hex())

	inst,err := NewEztoken(contractAddress,client)
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
func (_Eztoken *EztokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Eztoken.contract.Transact(opts, "transfer", _to, _value)
}

const EztokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"name\":\"_tokenName\",\"type\":\"string\"},{\"name\":\"_decimalUnits\",\"type\":\"uint8\"},{\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"
