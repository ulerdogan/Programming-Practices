// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igotest

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

// IgotestMetaData contains all meta data concerning the Igotest contract.
var IgotestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IgotestABI is the input ABI used to generate the binding from.
// Deprecated: Use IgotestMetaData.ABI instead.
var IgotestABI = IgotestMetaData.ABI

// Igotest is an auto generated Go binding around an Ethereum contract.
type Igotest struct {
	IgotestCaller     // Read-only binding to the contract
	IgotestTransactor // Write-only binding to the contract
	IgotestFilterer   // Log filterer for contract events
}

// IgotestCaller is an auto generated read-only Go binding around an Ethereum contract.
type IgotestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IgotestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IgotestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IgotestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IgotestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IgotestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IgotestSession struct {
	Contract     *Igotest          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IgotestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IgotestCallerSession struct {
	Contract *IgotestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IgotestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IgotestTransactorSession struct {
	Contract     *IgotestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IgotestRaw is an auto generated low-level Go binding around an Ethereum contract.
type IgotestRaw struct {
	Contract *Igotest // Generic contract binding to access the raw methods on
}

// IgotestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IgotestCallerRaw struct {
	Contract *IgotestCaller // Generic read-only contract binding to access the raw methods on
}

// IgotestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IgotestTransactorRaw struct {
	Contract *IgotestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIgotest creates a new instance of Igotest, bound to a specific deployed contract.
func NewIgotest(address common.Address, backend bind.ContractBackend) (*Igotest, error) {
	contract, err := bindIgotest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Igotest{IgotestCaller: IgotestCaller{contract: contract}, IgotestTransactor: IgotestTransactor{contract: contract}, IgotestFilterer: IgotestFilterer{contract: contract}}, nil
}

// NewIgotestCaller creates a new read-only instance of Igotest, bound to a specific deployed contract.
func NewIgotestCaller(address common.Address, caller bind.ContractCaller) (*IgotestCaller, error) {
	contract, err := bindIgotest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IgotestCaller{contract: contract}, nil
}

// NewIgotestTransactor creates a new write-only instance of Igotest, bound to a specific deployed contract.
func NewIgotestTransactor(address common.Address, transactor bind.ContractTransactor) (*IgotestTransactor, error) {
	contract, err := bindIgotest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IgotestTransactor{contract: contract}, nil
}

// NewIgotestFilterer creates a new log filterer instance of Igotest, bound to a specific deployed contract.
func NewIgotestFilterer(address common.Address, filterer bind.ContractFilterer) (*IgotestFilterer, error) {
	contract, err := bindIgotest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IgotestFilterer{contract: contract}, nil
}

// bindIgotest binds a generic wrapper to an already deployed contract.
func bindIgotest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IgotestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Igotest *IgotestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Igotest.Contract.IgotestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Igotest *IgotestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Igotest.Contract.IgotestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Igotest *IgotestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Igotest.Contract.IgotestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Igotest *IgotestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Igotest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Igotest *IgotestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Igotest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Igotest *IgotestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Igotest.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Igotest *IgotestTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Igotest.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Igotest *IgotestSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Igotest.Contract.Mint(&_Igotest.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Igotest *IgotestTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Igotest.Contract.Mint(&_Igotest.TransactOpts, to)
}
