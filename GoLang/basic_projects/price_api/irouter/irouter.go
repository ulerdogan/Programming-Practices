// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package irouter

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IPancakeRouter01MetaData contains all meta data concerning the IPancakeRouter01 contract.
var IPancakeRouter01MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d06ca61f": "getAmountsOut(uint256,address[])",
	},
}

// IPancakeRouter01ABI is the input ABI used to generate the binding from.
// Deprecated: Use IPancakeRouter01MetaData.ABI instead.
var IPancakeRouter01ABI = IPancakeRouter01MetaData.ABI

// Deprecated: Use IPancakeRouter01MetaData.Sigs instead.
// IPancakeRouter01FuncSigs maps the 4-byte function signature to its string representation.
var IPancakeRouter01FuncSigs = IPancakeRouter01MetaData.Sigs

// IPancakeRouter01 is an auto generated Go binding around an Ethereum contract.
type IPancakeRouter01 struct {
	IPancakeRouter01Caller     // Read-only binding to the contract
	IPancakeRouter01Transactor // Write-only binding to the contract
	IPancakeRouter01Filterer   // Log filterer for contract events
}

// IPancakeRouter01Caller is an auto generated read-only Go binding around an Ethereum contract.
type IPancakeRouter01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IPancakeRouter01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPancakeRouter01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPancakeRouter01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPancakeRouter01Session struct {
	Contract     *IPancakeRouter01 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPancakeRouter01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPancakeRouter01CallerSession struct {
	Contract *IPancakeRouter01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IPancakeRouter01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPancakeRouter01TransactorSession struct {
	Contract     *IPancakeRouter01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IPancakeRouter01Raw is an auto generated low-level Go binding around an Ethereum contract.
type IPancakeRouter01Raw struct {
	Contract *IPancakeRouter01 // Generic contract binding to access the raw methods on
}

// IPancakeRouter01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPancakeRouter01CallerRaw struct {
	Contract *IPancakeRouter01Caller // Generic read-only contract binding to access the raw methods on
}

// IPancakeRouter01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPancakeRouter01TransactorRaw struct {
	Contract *IPancakeRouter01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIPancakeRouter01 creates a new instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01(address common.Address, backend bind.ContractBackend) (*IPancakeRouter01, error) {
	contract, err := bindIPancakeRouter01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01{IPancakeRouter01Caller: IPancakeRouter01Caller{contract: contract}, IPancakeRouter01Transactor: IPancakeRouter01Transactor{contract: contract}, IPancakeRouter01Filterer: IPancakeRouter01Filterer{contract: contract}}, nil
}

// NewIPancakeRouter01Caller creates a new read-only instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Caller(address common.Address, caller bind.ContractCaller) (*IPancakeRouter01Caller, error) {
	contract, err := bindIPancakeRouter01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Caller{contract: contract}, nil
}

// NewIPancakeRouter01Transactor creates a new write-only instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Transactor(address common.Address, transactor bind.ContractTransactor) (*IPancakeRouter01Transactor, error) {
	contract, err := bindIPancakeRouter01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Transactor{contract: contract}, nil
}

// NewIPancakeRouter01Filterer creates a new log filterer instance of IPancakeRouter01, bound to a specific deployed contract.
func NewIPancakeRouter01Filterer(address common.Address, filterer bind.ContractFilterer) (*IPancakeRouter01Filterer, error) {
	contract, err := bindIPancakeRouter01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPancakeRouter01Filterer{contract: contract}, nil
}

// bindIPancakeRouter01 binds a generic wrapper to an already deployed contract.
func bindIPancakeRouter01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPancakeRouter01ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter01.Contract.IPancakeRouter01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.IPancakeRouter01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter01 *IPancakeRouter01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.IPancakeRouter01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPancakeRouter01 *IPancakeRouter01CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPancakeRouter01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPancakeRouter01 *IPancakeRouter01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPancakeRouter01 *IPancakeRouter01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPancakeRouter01.Contract.contract.Transact(opts, method, params...)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Caller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _IPancakeRouter01.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01Session) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsOut(&_IPancakeRouter01.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_IPancakeRouter01 *IPancakeRouter01CallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _IPancakeRouter01.Contract.GetAmountsOut(&_IPancakeRouter01.CallOpts, amountIn, path)
}
