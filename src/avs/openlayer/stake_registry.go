// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package openlayer

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
	_ = abi.ConvertType
)

// StakeRegistryMetaData contains all meta data concerning the StakeRegistry contract.
var StakeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorSignAddr\",\"type\":\"address\"}],\"name\":\"updateOperatorSignAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeRegistryMetaData.ABI instead.
var StakeRegistryABI = StakeRegistryMetaData.ABI

// StakeRegistry is an auto generated Go binding around an Ethereum contract.
type StakeRegistry struct {
	StakeRegistryCaller     // Read-only binding to the contract
	StakeRegistryTransactor // Write-only binding to the contract
	StakeRegistryFilterer   // Log filterer for contract events
}

// StakeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeRegistrySession struct {
	Contract     *StakeRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeRegistryCallerSession struct {
	Contract *StakeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeRegistryTransactorSession struct {
	Contract     *StakeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRegistryRaw struct {
	Contract *StakeRegistry // Generic contract binding to access the raw methods on
}

// StakeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeRegistryCallerRaw struct {
	Contract *StakeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// StakeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeRegistryTransactorRaw struct {
	Contract *StakeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeRegistry creates a new instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistry(address common.Address, backend bind.ContractBackend) (*StakeRegistry, error) {
	contract, err := bindStakeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeRegistry{StakeRegistryCaller: StakeRegistryCaller{contract: contract}, StakeRegistryTransactor: StakeRegistryTransactor{contract: contract}, StakeRegistryFilterer: StakeRegistryFilterer{contract: contract}}, nil
}

// NewStakeRegistryCaller creates a new read-only instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistryCaller(address common.Address, caller bind.ContractCaller) (*StakeRegistryCaller, error) {
	contract, err := bindStakeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryCaller{contract: contract}, nil
}

// NewStakeRegistryTransactor creates a new write-only instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeRegistryTransactor, error) {
	contract, err := bindStakeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryTransactor{contract: contract}, nil
}

// NewStakeRegistryFilterer creates a new log filterer instance of StakeRegistry, bound to a specific deployed contract.
func NewStakeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeRegistryFilterer, error) {
	contract, err := bindStakeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryFilterer{contract: contract}, nil
}

// bindStakeRegistry binds a generic wrapper to an already deployed contract.
func bindStakeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeRegistry *StakeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeRegistry.Contract.StakeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeRegistry *StakeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRegistry.Contract.StakeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeRegistry *StakeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeRegistry.Contract.StakeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeRegistry *StakeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeRegistry *StakeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeRegistry *StakeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeRegistry.Contract.contract.Transact(opts, method, params...)
}

// UpdateOperatorSignAddr is a paid mutator transaction binding the contract method 0xbfed57b0.
//
// Solidity: function updateOperatorSignAddr(address operatorSignAddr) returns()
func (_StakeRegistry *StakeRegistryTransactor) UpdateOperatorSignAddr(opts *bind.TransactOpts, operatorSignAddr common.Address) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateOperatorSignAddr", operatorSignAddr)
}

// UpdateOperatorSignAddr is a paid mutator transaction binding the contract method 0xbfed57b0.
//
// Solidity: function updateOperatorSignAddr(address operatorSignAddr) returns()
func (_StakeRegistry *StakeRegistrySession) UpdateOperatorSignAddr(operatorSignAddr common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorSignAddr(&_StakeRegistry.TransactOpts, operatorSignAddr)
}

// UpdateOperatorSignAddr is a paid mutator transaction binding the contract method 0xbfed57b0.
//
// Solidity: function updateOperatorSignAddr(address operatorSignAddr) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateOperatorSignAddr(operatorSignAddr common.Address) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorSignAddr(&_StakeRegistry.TransactOpts, operatorSignAddr)
}
