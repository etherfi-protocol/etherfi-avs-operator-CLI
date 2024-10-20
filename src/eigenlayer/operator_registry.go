// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eigenlayer

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

// OperatorRegistryMetaData contains all meta data concerning the OperatorRegistry contract.
var OperatorRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registeredWallet\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registeredWallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operatorMessage\",\"type\":\"string\"}],\"name\":\"Register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OperatorRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorRegistryMetaData.ABI instead.
var OperatorRegistryABI = OperatorRegistryMetaData.ABI

// OperatorRegistry is an auto generated Go binding around an Ethereum contract.
type OperatorRegistry struct {
	OperatorRegistryCaller     // Read-only binding to the contract
	OperatorRegistryTransactor // Write-only binding to the contract
	OperatorRegistryFilterer   // Log filterer for contract events
}

// OperatorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorRegistrySession struct {
	Contract     *OperatorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorRegistryCallerSession struct {
	Contract *OperatorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OperatorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorRegistryTransactorSession struct {
	Contract     *OperatorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OperatorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRegistryRaw struct {
	Contract *OperatorRegistry // Generic contract binding to access the raw methods on
}

// OperatorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorRegistryCallerRaw struct {
	Contract *OperatorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorRegistryTransactorRaw struct {
	Contract *OperatorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperatorRegistry creates a new instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistry(address common.Address, backend bind.ContractBackend) (*OperatorRegistry, error) {
	contract, err := bindOperatorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistry{OperatorRegistryCaller: OperatorRegistryCaller{contract: contract}, OperatorRegistryTransactor: OperatorRegistryTransactor{contract: contract}, OperatorRegistryFilterer: OperatorRegistryFilterer{contract: contract}}, nil
}

// NewOperatorRegistryCaller creates a new read-only instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryCaller(address common.Address, caller bind.ContractCaller) (*OperatorRegistryCaller, error) {
	contract, err := bindOperatorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryCaller{contract: contract}, nil
}

// NewOperatorRegistryTransactor creates a new write-only instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorRegistryTransactor, error) {
	contract, err := bindOperatorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryTransactor{contract: contract}, nil
}

// NewOperatorRegistryFilterer creates a new log filterer instance of OperatorRegistry, bound to a specific deployed contract.
func NewOperatorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorRegistryFilterer, error) {
	contract, err := bindOperatorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryFilterer{contract: contract}, nil
}

// bindOperatorRegistry binds a generic wrapper to an already deployed contract.
func bindOperatorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRegistry *OperatorRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRegistry.Contract.OperatorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRegistry *OperatorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.OperatorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRegistry *OperatorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.OperatorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OperatorRegistry *OperatorRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OperatorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OperatorRegistry *OperatorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OperatorRegistry *OperatorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.contract.Transact(opts, method, params...)
}

// Register is a paid mutator transaction binding the contract method 0x6ba0831d.
//
// Solidity: function Register(address registeredWallet, string operatorMessage) returns()
func (_OperatorRegistry *OperatorRegistryTransactor) Register(opts *bind.TransactOpts, registeredWallet common.Address, operatorMessage string) (*types.Transaction, error) {
	return _OperatorRegistry.contract.Transact(opts, "Register", registeredWallet, operatorMessage)
}

// Register is a paid mutator transaction binding the contract method 0x6ba0831d.
//
// Solidity: function Register(address registeredWallet, string operatorMessage) returns()
func (_OperatorRegistry *OperatorRegistrySession) Register(registeredWallet common.Address, operatorMessage string) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Register(&_OperatorRegistry.TransactOpts, registeredWallet, operatorMessage)
}

// Register is a paid mutator transaction binding the contract method 0x6ba0831d.
//
// Solidity: function Register(address registeredWallet, string operatorMessage) returns()
func (_OperatorRegistry *OperatorRegistryTransactorSession) Register(registeredWallet common.Address, operatorMessage string) (*types.Transaction, error) {
	return _OperatorRegistry.Contract.Register(&_OperatorRegistry.TransactOpts, registeredWallet, operatorMessage)
}

// OperatorRegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the OperatorRegistry contract.
type OperatorRegistryRegisteredIterator struct {
	Event *OperatorRegistryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OperatorRegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorRegistryRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OperatorRegistryRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OperatorRegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorRegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorRegistryRegistered represents a Registered event raised by the OperatorRegistry contract.
type OperatorRegistryRegistered struct {
	Operator         common.Address
	RegisteredWallet common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address indexed operator, address indexed registeredWallet)
func (_OperatorRegistry *OperatorRegistryFilterer) FilterRegistered(opts *bind.FilterOpts, operator []common.Address, registeredWallet []common.Address) (*OperatorRegistryRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var registeredWalletRule []interface{}
	for _, registeredWalletItem := range registeredWallet {
		registeredWalletRule = append(registeredWalletRule, registeredWalletItem)
	}

	logs, sub, err := _OperatorRegistry.contract.FilterLogs(opts, "Registered", operatorRule, registeredWalletRule)
	if err != nil {
		return nil, err
	}
	return &OperatorRegistryRegisteredIterator{contract: _OperatorRegistry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address indexed operator, address indexed registeredWallet)
func (_OperatorRegistry *OperatorRegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *OperatorRegistryRegistered, operator []common.Address, registeredWallet []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var registeredWalletRule []interface{}
	for _, registeredWalletItem := range registeredWallet {
		registeredWalletRule = append(registeredWalletRule, registeredWalletItem)
	}

	logs, sub, err := _OperatorRegistry.contract.WatchLogs(opts, "Registered", operatorRule, registeredWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorRegistryRegistered)
				if err := _OperatorRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistered is a log parse operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address indexed operator, address indexed registeredWallet)
func (_OperatorRegistry *OperatorRegistryFilterer) ParseRegistered(log types.Log) (*OperatorRegistryRegistered, error) {
	event := new(OperatorRegistryRegistered)
	if err := _OperatorRegistry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
