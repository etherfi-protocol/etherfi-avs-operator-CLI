// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package symbiotic

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

// IBurnerRouterInitParams is an auto generated low-level Go binding around an user-defined struct.
type IBurnerRouterInitParams struct {
	Owner                    common.Address
	Collateral               common.Address
	Delay                    *big.Int
	GlobalReceiver           common.Address
	NetworkReceivers         []IBurnerRouterNetworkReceiver
	OperatorNetworkReceivers []IBurnerRouterOperatorNetworkReceiver
}

// IBurnerRouterNetworkReceiver is an auto generated low-level Go binding around an user-defined struct.
type IBurnerRouterNetworkReceiver struct {
	Network  common.Address
	Receiver common.Address
}

// IBurnerRouterOperatorNetworkReceiver is an auto generated low-level Go binding around an user-defined struct.
type IBurnerRouterOperatorNetworkReceiver struct {
	Network  common.Address
	Operator common.Address
	Receiver common.Address
}

// BurnerRouterMetaData contains all meta data concerning the BurnerRouter contract.
var BurnerRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateNetworkReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateOperatorNetworkReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReceiverSetEpochsDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotReady\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AcceptDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AcceptGlobalReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"}],\"name\":\"AcceptNetworkReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"AcceptOperatorNetworkReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"delay\",\"type\":\"uint48\"}],\"name\":\"SetDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetGlobalReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetNetworkReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetOperatorNetworkReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TriggerTransfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"}],\"name\":\"acceptNetworkReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"acceptOperatorNetworkReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"value\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"delay\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"globalReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structIBurnerRouter.NetworkReceiver[]\",\"name\":\"networkReceivers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"internalType\":\"structIBurnerRouter.OperatorNetworkReceiver[]\",\"name\":\"operatorNetworkReceivers\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBurnerRouter.InitParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"}],\"name\":\"networkReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"name\":\"onSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"operatorNetworkReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDelay\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"value\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGlobalReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"}],\"name\":\"pendingNetworkReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"pendingOperatorNetworkReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newDelay\",\"type\":\"uint48\"}],\"name\":\"setDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setGlobalReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setNetworkReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"network\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"setOperatorNetworkReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"triggerTransfer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BurnerRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnerRouterMetaData.ABI instead.
var BurnerRouterABI = BurnerRouterMetaData.ABI

// BurnerRouter is an auto generated Go binding around an Ethereum contract.
type BurnerRouter struct {
	BurnerRouterCaller     // Read-only binding to the contract
	BurnerRouterTransactor // Write-only binding to the contract
	BurnerRouterFilterer   // Log filterer for contract events
}

// BurnerRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnerRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnerRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnerRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnerRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnerRouterSession struct {
	Contract     *BurnerRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnerRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnerRouterCallerSession struct {
	Contract *BurnerRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BurnerRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnerRouterTransactorSession struct {
	Contract     *BurnerRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BurnerRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnerRouterRaw struct {
	Contract *BurnerRouter // Generic contract binding to access the raw methods on
}

// BurnerRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnerRouterCallerRaw struct {
	Contract *BurnerRouterCaller // Generic read-only contract binding to access the raw methods on
}

// BurnerRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnerRouterTransactorRaw struct {
	Contract *BurnerRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnerRouter creates a new instance of BurnerRouter, bound to a specific deployed contract.
func NewBurnerRouter(address common.Address, backend bind.ContractBackend) (*BurnerRouter, error) {
	contract, err := bindBurnerRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnerRouter{BurnerRouterCaller: BurnerRouterCaller{contract: contract}, BurnerRouterTransactor: BurnerRouterTransactor{contract: contract}, BurnerRouterFilterer: BurnerRouterFilterer{contract: contract}}, nil
}

// NewBurnerRouterCaller creates a new read-only instance of BurnerRouter, bound to a specific deployed contract.
func NewBurnerRouterCaller(address common.Address, caller bind.ContractCaller) (*BurnerRouterCaller, error) {
	contract, err := bindBurnerRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterCaller{contract: contract}, nil
}

// NewBurnerRouterTransactor creates a new write-only instance of BurnerRouter, bound to a specific deployed contract.
func NewBurnerRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnerRouterTransactor, error) {
	contract, err := bindBurnerRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterTransactor{contract: contract}, nil
}

// NewBurnerRouterFilterer creates a new log filterer instance of BurnerRouter, bound to a specific deployed contract.
func NewBurnerRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnerRouterFilterer, error) {
	contract, err := bindBurnerRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterFilterer{contract: contract}, nil
}

// bindBurnerRouter binds a generic wrapper to an already deployed contract.
func bindBurnerRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnerRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnerRouter *BurnerRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnerRouter.Contract.BurnerRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnerRouter *BurnerRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerRouter.Contract.BurnerRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnerRouter *BurnerRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnerRouter.Contract.BurnerRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnerRouter *BurnerRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnerRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnerRouter *BurnerRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnerRouter *BurnerRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnerRouter.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address receiver) view returns(uint256 amount)
func (_BurnerRouter *BurnerRouterCaller) BalanceOf(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "balanceOf", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address receiver) view returns(uint256 amount)
func (_BurnerRouter *BurnerRouterSession) BalanceOf(receiver common.Address) (*big.Int, error) {
	return _BurnerRouter.Contract.BalanceOf(&_BurnerRouter.CallOpts, receiver)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address receiver) view returns(uint256 amount)
func (_BurnerRouter *BurnerRouterCallerSession) BalanceOf(receiver common.Address) (*big.Int, error) {
	return _BurnerRouter.Contract.BalanceOf(&_BurnerRouter.CallOpts, receiver)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_BurnerRouter *BurnerRouterCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_BurnerRouter *BurnerRouterSession) Collateral() (common.Address, error) {
	return _BurnerRouter.Contract.Collateral(&_BurnerRouter.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_BurnerRouter *BurnerRouterCallerSession) Collateral() (common.Address, error) {
	return _BurnerRouter.Contract.Collateral(&_BurnerRouter.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint48 value)
func (_BurnerRouter *BurnerRouterCaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint48 value)
func (_BurnerRouter *BurnerRouterSession) Delay() (*big.Int, error) {
	return _BurnerRouter.Contract.Delay(&_BurnerRouter.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint48 value)
func (_BurnerRouter *BurnerRouterCallerSession) Delay() (*big.Int, error) {
	return _BurnerRouter.Contract.Delay(&_BurnerRouter.CallOpts)
}

// GlobalReceiver is a free data retrieval call binding the contract method 0x467aea20.
//
// Solidity: function globalReceiver() view returns(address value)
func (_BurnerRouter *BurnerRouterCaller) GlobalReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "globalReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalReceiver is a free data retrieval call binding the contract method 0x467aea20.
//
// Solidity: function globalReceiver() view returns(address value)
func (_BurnerRouter *BurnerRouterSession) GlobalReceiver() (common.Address, error) {
	return _BurnerRouter.Contract.GlobalReceiver(&_BurnerRouter.CallOpts)
}

// GlobalReceiver is a free data retrieval call binding the contract method 0x467aea20.
//
// Solidity: function globalReceiver() view returns(address value)
func (_BurnerRouter *BurnerRouterCallerSession) GlobalReceiver() (common.Address, error) {
	return _BurnerRouter.Contract.GlobalReceiver(&_BurnerRouter.CallOpts)
}

// LastBalance is a free data retrieval call binding the contract method 0x8f1c56bd.
//
// Solidity: function lastBalance() view returns(uint256)
func (_BurnerRouter *BurnerRouterCaller) LastBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "lastBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBalance is a free data retrieval call binding the contract method 0x8f1c56bd.
//
// Solidity: function lastBalance() view returns(uint256)
func (_BurnerRouter *BurnerRouterSession) LastBalance() (*big.Int, error) {
	return _BurnerRouter.Contract.LastBalance(&_BurnerRouter.CallOpts)
}

// LastBalance is a free data retrieval call binding the contract method 0x8f1c56bd.
//
// Solidity: function lastBalance() view returns(uint256)
func (_BurnerRouter *BurnerRouterCallerSession) LastBalance() (*big.Int, error) {
	return _BurnerRouter.Contract.LastBalance(&_BurnerRouter.CallOpts)
}

// NetworkReceiver is a free data retrieval call binding the contract method 0xae89186b.
//
// Solidity: function networkReceiver(address network) view returns(address value)
func (_BurnerRouter *BurnerRouterCaller) NetworkReceiver(opts *bind.CallOpts, network common.Address) (common.Address, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "networkReceiver", network)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkReceiver is a free data retrieval call binding the contract method 0xae89186b.
//
// Solidity: function networkReceiver(address network) view returns(address value)
func (_BurnerRouter *BurnerRouterSession) NetworkReceiver(network common.Address) (common.Address, error) {
	return _BurnerRouter.Contract.NetworkReceiver(&_BurnerRouter.CallOpts, network)
}

// NetworkReceiver is a free data retrieval call binding the contract method 0xae89186b.
//
// Solidity: function networkReceiver(address network) view returns(address value)
func (_BurnerRouter *BurnerRouterCallerSession) NetworkReceiver(network common.Address) (common.Address, error) {
	return _BurnerRouter.Contract.NetworkReceiver(&_BurnerRouter.CallOpts, network)
}

// OperatorNetworkReceiver is a free data retrieval call binding the contract method 0xd439351f.
//
// Solidity: function operatorNetworkReceiver(address network, address operator) view returns(address value)
func (_BurnerRouter *BurnerRouterCaller) OperatorNetworkReceiver(opts *bind.CallOpts, network common.Address, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "operatorNetworkReceiver", network, operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorNetworkReceiver is a free data retrieval call binding the contract method 0xd439351f.
//
// Solidity: function operatorNetworkReceiver(address network, address operator) view returns(address value)
func (_BurnerRouter *BurnerRouterSession) OperatorNetworkReceiver(network common.Address, operator common.Address) (common.Address, error) {
	return _BurnerRouter.Contract.OperatorNetworkReceiver(&_BurnerRouter.CallOpts, network, operator)
}

// OperatorNetworkReceiver is a free data retrieval call binding the contract method 0xd439351f.
//
// Solidity: function operatorNetworkReceiver(address network, address operator) view returns(address value)
func (_BurnerRouter *BurnerRouterCallerSession) OperatorNetworkReceiver(network common.Address, operator common.Address) (common.Address, error) {
	return _BurnerRouter.Contract.OperatorNetworkReceiver(&_BurnerRouter.CallOpts, network, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BurnerRouter *BurnerRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BurnerRouter *BurnerRouterSession) Owner() (common.Address, error) {
	return _BurnerRouter.Contract.Owner(&_BurnerRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BurnerRouter *BurnerRouterCallerSession) Owner() (common.Address, error) {
	return _BurnerRouter.Contract.Owner(&_BurnerRouter.CallOpts)
}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint48 value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCaller) PendingDelay(opts *bind.CallOpts) (struct {
	Value     *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "pendingDelay")

	outstruct := new(struct {
		Value     *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint48 value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterSession) PendingDelay() (struct {
	Value     *big.Int
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingDelay(&_BurnerRouter.CallOpts)
}

// PendingDelay is a free data retrieval call binding the contract method 0x4ca8f0ed.
//
// Solidity: function pendingDelay() view returns(uint48 value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCallerSession) PendingDelay() (struct {
	Value     *big.Int
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingDelay(&_BurnerRouter.CallOpts)
}

// PendingGlobalReceiver is a free data retrieval call binding the contract method 0x3cf966c9.
//
// Solidity: function pendingGlobalReceiver() view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCaller) PendingGlobalReceiver(opts *bind.CallOpts) (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "pendingGlobalReceiver")

	outstruct := new(struct {
		Value     common.Address
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingGlobalReceiver is a free data retrieval call binding the contract method 0x3cf966c9.
//
// Solidity: function pendingGlobalReceiver() view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterSession) PendingGlobalReceiver() (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingGlobalReceiver(&_BurnerRouter.CallOpts)
}

// PendingGlobalReceiver is a free data retrieval call binding the contract method 0x3cf966c9.
//
// Solidity: function pendingGlobalReceiver() view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCallerSession) PendingGlobalReceiver() (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingGlobalReceiver(&_BurnerRouter.CallOpts)
}

// PendingNetworkReceiver is a free data retrieval call binding the contract method 0x03321289.
//
// Solidity: function pendingNetworkReceiver(address network) view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCaller) PendingNetworkReceiver(opts *bind.CallOpts, network common.Address) (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "pendingNetworkReceiver", network)

	outstruct := new(struct {
		Value     common.Address
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingNetworkReceiver is a free data retrieval call binding the contract method 0x03321289.
//
// Solidity: function pendingNetworkReceiver(address network) view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterSession) PendingNetworkReceiver(network common.Address) (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingNetworkReceiver(&_BurnerRouter.CallOpts, network)
}

// PendingNetworkReceiver is a free data retrieval call binding the contract method 0x03321289.
//
// Solidity: function pendingNetworkReceiver(address network) view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCallerSession) PendingNetworkReceiver(network common.Address) (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingNetworkReceiver(&_BurnerRouter.CallOpts, network)
}

// PendingOperatorNetworkReceiver is a free data retrieval call binding the contract method 0x0760bac4.
//
// Solidity: function pendingOperatorNetworkReceiver(address network, address operator) view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCaller) PendingOperatorNetworkReceiver(opts *bind.CallOpts, network common.Address, operator common.Address) (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _BurnerRouter.contract.Call(opts, &out, "pendingOperatorNetworkReceiver", network, operator)

	outstruct := new(struct {
		Value     common.Address
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingOperatorNetworkReceiver is a free data retrieval call binding the contract method 0x0760bac4.
//
// Solidity: function pendingOperatorNetworkReceiver(address network, address operator) view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterSession) PendingOperatorNetworkReceiver(network common.Address, operator common.Address) (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingOperatorNetworkReceiver(&_BurnerRouter.CallOpts, network, operator)
}

// PendingOperatorNetworkReceiver is a free data retrieval call binding the contract method 0x0760bac4.
//
// Solidity: function pendingOperatorNetworkReceiver(address network, address operator) view returns(address value, uint48 timestamp)
func (_BurnerRouter *BurnerRouterCallerSession) PendingOperatorNetworkReceiver(network common.Address, operator common.Address) (struct {
	Value     common.Address
	Timestamp *big.Int
}, error) {
	return _BurnerRouter.Contract.PendingOperatorNetworkReceiver(&_BurnerRouter.CallOpts, network, operator)
}

// AcceptDelay is a paid mutator transaction binding the contract method 0xf6f371ee.
//
// Solidity: function acceptDelay() returns()
func (_BurnerRouter *BurnerRouterTransactor) AcceptDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "acceptDelay")
}

// AcceptDelay is a paid mutator transaction binding the contract method 0xf6f371ee.
//
// Solidity: function acceptDelay() returns()
func (_BurnerRouter *BurnerRouterSession) AcceptDelay() (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptDelay(&_BurnerRouter.TransactOpts)
}

// AcceptDelay is a paid mutator transaction binding the contract method 0xf6f371ee.
//
// Solidity: function acceptDelay() returns()
func (_BurnerRouter *BurnerRouterTransactorSession) AcceptDelay() (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptDelay(&_BurnerRouter.TransactOpts)
}

// AcceptGlobalReceiver is a paid mutator transaction binding the contract method 0x74df73dd.
//
// Solidity: function acceptGlobalReceiver() returns()
func (_BurnerRouter *BurnerRouterTransactor) AcceptGlobalReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "acceptGlobalReceiver")
}

// AcceptGlobalReceiver is a paid mutator transaction binding the contract method 0x74df73dd.
//
// Solidity: function acceptGlobalReceiver() returns()
func (_BurnerRouter *BurnerRouterSession) AcceptGlobalReceiver() (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptGlobalReceiver(&_BurnerRouter.TransactOpts)
}

// AcceptGlobalReceiver is a paid mutator transaction binding the contract method 0x74df73dd.
//
// Solidity: function acceptGlobalReceiver() returns()
func (_BurnerRouter *BurnerRouterTransactorSession) AcceptGlobalReceiver() (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptGlobalReceiver(&_BurnerRouter.TransactOpts)
}

// AcceptNetworkReceiver is a paid mutator transaction binding the contract method 0x0bcf996f.
//
// Solidity: function acceptNetworkReceiver(address network) returns()
func (_BurnerRouter *BurnerRouterTransactor) AcceptNetworkReceiver(opts *bind.TransactOpts, network common.Address) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "acceptNetworkReceiver", network)
}

// AcceptNetworkReceiver is a paid mutator transaction binding the contract method 0x0bcf996f.
//
// Solidity: function acceptNetworkReceiver(address network) returns()
func (_BurnerRouter *BurnerRouterSession) AcceptNetworkReceiver(network common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptNetworkReceiver(&_BurnerRouter.TransactOpts, network)
}

// AcceptNetworkReceiver is a paid mutator transaction binding the contract method 0x0bcf996f.
//
// Solidity: function acceptNetworkReceiver(address network) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) AcceptNetworkReceiver(network common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptNetworkReceiver(&_BurnerRouter.TransactOpts, network)
}

// AcceptOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x898dc787.
//
// Solidity: function acceptOperatorNetworkReceiver(address network, address operator) returns()
func (_BurnerRouter *BurnerRouterTransactor) AcceptOperatorNetworkReceiver(opts *bind.TransactOpts, network common.Address, operator common.Address) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "acceptOperatorNetworkReceiver", network, operator)
}

// AcceptOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x898dc787.
//
// Solidity: function acceptOperatorNetworkReceiver(address network, address operator) returns()
func (_BurnerRouter *BurnerRouterSession) AcceptOperatorNetworkReceiver(network common.Address, operator common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptOperatorNetworkReceiver(&_BurnerRouter.TransactOpts, network, operator)
}

// AcceptOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x898dc787.
//
// Solidity: function acceptOperatorNetworkReceiver(address network, address operator) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) AcceptOperatorNetworkReceiver(network common.Address, operator common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.AcceptOperatorNetworkReceiver(&_BurnerRouter.TransactOpts, network, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0x42f929df.
//
// Solidity: function initialize((address,address,uint48,address,(address,address)[],(address,address,address)[]) params) returns()
func (_BurnerRouter *BurnerRouterTransactor) Initialize(opts *bind.TransactOpts, params IBurnerRouterInitParams) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "initialize", params)
}

// Initialize is a paid mutator transaction binding the contract method 0x42f929df.
//
// Solidity: function initialize((address,address,uint48,address,(address,address)[],(address,address,address)[]) params) returns()
func (_BurnerRouter *BurnerRouterSession) Initialize(params IBurnerRouterInitParams) (*types.Transaction, error) {
	return _BurnerRouter.Contract.Initialize(&_BurnerRouter.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x42f929df.
//
// Solidity: function initialize((address,address,uint48,address,(address,address)[],(address,address,address)[]) params) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) Initialize(params IBurnerRouterInitParams) (*types.Transaction, error) {
	return _BurnerRouter.Contract.Initialize(&_BurnerRouter.TransactOpts, params)
}

// OnSlash is a paid mutator transaction binding the contract method 0x065c1e03.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 , uint48 ) returns()
func (_BurnerRouter *BurnerRouterTransactor) OnSlash(opts *bind.TransactOpts, subnetwork [32]byte, operator common.Address, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "onSlash", subnetwork, operator, arg2, arg3)
}

// OnSlash is a paid mutator transaction binding the contract method 0x065c1e03.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 , uint48 ) returns()
func (_BurnerRouter *BurnerRouterSession) OnSlash(subnetwork [32]byte, operator common.Address, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _BurnerRouter.Contract.OnSlash(&_BurnerRouter.TransactOpts, subnetwork, operator, arg2, arg3)
}

// OnSlash is a paid mutator transaction binding the contract method 0x065c1e03.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 , uint48 ) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) OnSlash(subnetwork [32]byte, operator common.Address, arg2 *big.Int, arg3 *big.Int) (*types.Transaction, error) {
	return _BurnerRouter.Contract.OnSlash(&_BurnerRouter.TransactOpts, subnetwork, operator, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnerRouter *BurnerRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnerRouter *BurnerRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _BurnerRouter.Contract.RenounceOwnership(&_BurnerRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BurnerRouter *BurnerRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BurnerRouter.Contract.RenounceOwnership(&_BurnerRouter.TransactOpts)
}

// SetDelay is a paid mutator transaction binding the contract method 0x40868ce6.
//
// Solidity: function setDelay(uint48 newDelay) returns()
func (_BurnerRouter *BurnerRouterTransactor) SetDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "setDelay", newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x40868ce6.
//
// Solidity: function setDelay(uint48 newDelay) returns()
func (_BurnerRouter *BurnerRouterSession) SetDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetDelay(&_BurnerRouter.TransactOpts, newDelay)
}

// SetDelay is a paid mutator transaction binding the contract method 0x40868ce6.
//
// Solidity: function setDelay(uint48 newDelay) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) SetDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetDelay(&_BurnerRouter.TransactOpts, newDelay)
}

// SetGlobalReceiver is a paid mutator transaction binding the contract method 0xa472e384.
//
// Solidity: function setGlobalReceiver(address receiver) returns()
func (_BurnerRouter *BurnerRouterTransactor) SetGlobalReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "setGlobalReceiver", receiver)
}

// SetGlobalReceiver is a paid mutator transaction binding the contract method 0xa472e384.
//
// Solidity: function setGlobalReceiver(address receiver) returns()
func (_BurnerRouter *BurnerRouterSession) SetGlobalReceiver(receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetGlobalReceiver(&_BurnerRouter.TransactOpts, receiver)
}

// SetGlobalReceiver is a paid mutator transaction binding the contract method 0xa472e384.
//
// Solidity: function setGlobalReceiver(address receiver) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) SetGlobalReceiver(receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetGlobalReceiver(&_BurnerRouter.TransactOpts, receiver)
}

// SetNetworkReceiver is a paid mutator transaction binding the contract method 0xacea136b.
//
// Solidity: function setNetworkReceiver(address network, address receiver) returns()
func (_BurnerRouter *BurnerRouterTransactor) SetNetworkReceiver(opts *bind.TransactOpts, network common.Address, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "setNetworkReceiver", network, receiver)
}

// SetNetworkReceiver is a paid mutator transaction binding the contract method 0xacea136b.
//
// Solidity: function setNetworkReceiver(address network, address receiver) returns()
func (_BurnerRouter *BurnerRouterSession) SetNetworkReceiver(network common.Address, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetNetworkReceiver(&_BurnerRouter.TransactOpts, network, receiver)
}

// SetNetworkReceiver is a paid mutator transaction binding the contract method 0xacea136b.
//
// Solidity: function setNetworkReceiver(address network, address receiver) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) SetNetworkReceiver(network common.Address, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetNetworkReceiver(&_BurnerRouter.TransactOpts, network, receiver)
}

// SetOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x462dac19.
//
// Solidity: function setOperatorNetworkReceiver(address network, address operator, address receiver) returns()
func (_BurnerRouter *BurnerRouterTransactor) SetOperatorNetworkReceiver(opts *bind.TransactOpts, network common.Address, operator common.Address, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "setOperatorNetworkReceiver", network, operator, receiver)
}

// SetOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x462dac19.
//
// Solidity: function setOperatorNetworkReceiver(address network, address operator, address receiver) returns()
func (_BurnerRouter *BurnerRouterSession) SetOperatorNetworkReceiver(network common.Address, operator common.Address, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetOperatorNetworkReceiver(&_BurnerRouter.TransactOpts, network, operator, receiver)
}

// SetOperatorNetworkReceiver is a paid mutator transaction binding the contract method 0x462dac19.
//
// Solidity: function setOperatorNetworkReceiver(address network, address operator, address receiver) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) SetOperatorNetworkReceiver(network common.Address, operator common.Address, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.SetOperatorNetworkReceiver(&_BurnerRouter.TransactOpts, network, operator, receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BurnerRouter *BurnerRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BurnerRouter *BurnerRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.TransferOwnership(&_BurnerRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BurnerRouter *BurnerRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.TransferOwnership(&_BurnerRouter.TransactOpts, newOwner)
}

// TriggerTransfer is a paid mutator transaction binding the contract method 0xa51b90be.
//
// Solidity: function triggerTransfer(address receiver) returns(uint256 amount)
func (_BurnerRouter *BurnerRouterTransactor) TriggerTransfer(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.contract.Transact(opts, "triggerTransfer", receiver)
}

// TriggerTransfer is a paid mutator transaction binding the contract method 0xa51b90be.
//
// Solidity: function triggerTransfer(address receiver) returns(uint256 amount)
func (_BurnerRouter *BurnerRouterSession) TriggerTransfer(receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.TriggerTransfer(&_BurnerRouter.TransactOpts, receiver)
}

// TriggerTransfer is a paid mutator transaction binding the contract method 0xa51b90be.
//
// Solidity: function triggerTransfer(address receiver) returns(uint256 amount)
func (_BurnerRouter *BurnerRouterTransactorSession) TriggerTransfer(receiver common.Address) (*types.Transaction, error) {
	return _BurnerRouter.Contract.TriggerTransfer(&_BurnerRouter.TransactOpts, receiver)
}

// BurnerRouterAcceptDelayIterator is returned from FilterAcceptDelay and is used to iterate over the raw logs and unpacked data for AcceptDelay events raised by the BurnerRouter contract.
type BurnerRouterAcceptDelayIterator struct {
	Event *BurnerRouterAcceptDelay // Event containing the contract specifics and raw log

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
func (it *BurnerRouterAcceptDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterAcceptDelay)
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
		it.Event = new(BurnerRouterAcceptDelay)
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
func (it *BurnerRouterAcceptDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterAcceptDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterAcceptDelay represents a AcceptDelay event raised by the BurnerRouter contract.
type BurnerRouterAcceptDelay struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAcceptDelay is a free log retrieval operation binding the contract event 0x54b01918a30e934a38ff39572e9095d6c78b521b8efec12d15a1b485156257eb.
//
// Solidity: event AcceptDelay()
func (_BurnerRouter *BurnerRouterFilterer) FilterAcceptDelay(opts *bind.FilterOpts) (*BurnerRouterAcceptDelayIterator, error) {

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "AcceptDelay")
	if err != nil {
		return nil, err
	}
	return &BurnerRouterAcceptDelayIterator{contract: _BurnerRouter.contract, event: "AcceptDelay", logs: logs, sub: sub}, nil
}

// WatchAcceptDelay is a free log subscription operation binding the contract event 0x54b01918a30e934a38ff39572e9095d6c78b521b8efec12d15a1b485156257eb.
//
// Solidity: event AcceptDelay()
func (_BurnerRouter *BurnerRouterFilterer) WatchAcceptDelay(opts *bind.WatchOpts, sink chan<- *BurnerRouterAcceptDelay) (event.Subscription, error) {

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "AcceptDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterAcceptDelay)
				if err := _BurnerRouter.contract.UnpackLog(event, "AcceptDelay", log); err != nil {
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

// ParseAcceptDelay is a log parse operation binding the contract event 0x54b01918a30e934a38ff39572e9095d6c78b521b8efec12d15a1b485156257eb.
//
// Solidity: event AcceptDelay()
func (_BurnerRouter *BurnerRouterFilterer) ParseAcceptDelay(log types.Log) (*BurnerRouterAcceptDelay, error) {
	event := new(BurnerRouterAcceptDelay)
	if err := _BurnerRouter.contract.UnpackLog(event, "AcceptDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterAcceptGlobalReceiverIterator is returned from FilterAcceptGlobalReceiver and is used to iterate over the raw logs and unpacked data for AcceptGlobalReceiver events raised by the BurnerRouter contract.
type BurnerRouterAcceptGlobalReceiverIterator struct {
	Event *BurnerRouterAcceptGlobalReceiver // Event containing the contract specifics and raw log

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
func (it *BurnerRouterAcceptGlobalReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterAcceptGlobalReceiver)
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
		it.Event = new(BurnerRouterAcceptGlobalReceiver)
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
func (it *BurnerRouterAcceptGlobalReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterAcceptGlobalReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterAcceptGlobalReceiver represents a AcceptGlobalReceiver event raised by the BurnerRouter contract.
type BurnerRouterAcceptGlobalReceiver struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAcceptGlobalReceiver is a free log retrieval operation binding the contract event 0x74167a6969567de7e1730e9b22e87e4fe263e7fa04bec628436c424fc7bd6b8e.
//
// Solidity: event AcceptGlobalReceiver()
func (_BurnerRouter *BurnerRouterFilterer) FilterAcceptGlobalReceiver(opts *bind.FilterOpts) (*BurnerRouterAcceptGlobalReceiverIterator, error) {

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "AcceptGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return &BurnerRouterAcceptGlobalReceiverIterator{contract: _BurnerRouter.contract, event: "AcceptGlobalReceiver", logs: logs, sub: sub}, nil
}

// WatchAcceptGlobalReceiver is a free log subscription operation binding the contract event 0x74167a6969567de7e1730e9b22e87e4fe263e7fa04bec628436c424fc7bd6b8e.
//
// Solidity: event AcceptGlobalReceiver()
func (_BurnerRouter *BurnerRouterFilterer) WatchAcceptGlobalReceiver(opts *bind.WatchOpts, sink chan<- *BurnerRouterAcceptGlobalReceiver) (event.Subscription, error) {

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "AcceptGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterAcceptGlobalReceiver)
				if err := _BurnerRouter.contract.UnpackLog(event, "AcceptGlobalReceiver", log); err != nil {
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

// ParseAcceptGlobalReceiver is a log parse operation binding the contract event 0x74167a6969567de7e1730e9b22e87e4fe263e7fa04bec628436c424fc7bd6b8e.
//
// Solidity: event AcceptGlobalReceiver()
func (_BurnerRouter *BurnerRouterFilterer) ParseAcceptGlobalReceiver(log types.Log) (*BurnerRouterAcceptGlobalReceiver, error) {
	event := new(BurnerRouterAcceptGlobalReceiver)
	if err := _BurnerRouter.contract.UnpackLog(event, "AcceptGlobalReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterAcceptNetworkReceiverIterator is returned from FilterAcceptNetworkReceiver and is used to iterate over the raw logs and unpacked data for AcceptNetworkReceiver events raised by the BurnerRouter contract.
type BurnerRouterAcceptNetworkReceiverIterator struct {
	Event *BurnerRouterAcceptNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *BurnerRouterAcceptNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterAcceptNetworkReceiver)
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
		it.Event = new(BurnerRouterAcceptNetworkReceiver)
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
func (it *BurnerRouterAcceptNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterAcceptNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterAcceptNetworkReceiver represents a AcceptNetworkReceiver event raised by the BurnerRouter contract.
type BurnerRouterAcceptNetworkReceiver struct {
	Network common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAcceptNetworkReceiver is a free log retrieval operation binding the contract event 0x1a2023b9b05a5599a274f08b91afd34b22b21ea58b7ca66ef06897746db55b0f.
//
// Solidity: event AcceptNetworkReceiver(address indexed network)
func (_BurnerRouter *BurnerRouterFilterer) FilterAcceptNetworkReceiver(opts *bind.FilterOpts, network []common.Address) (*BurnerRouterAcceptNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "AcceptNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterAcceptNetworkReceiverIterator{contract: _BurnerRouter.contract, event: "AcceptNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchAcceptNetworkReceiver is a free log subscription operation binding the contract event 0x1a2023b9b05a5599a274f08b91afd34b22b21ea58b7ca66ef06897746db55b0f.
//
// Solidity: event AcceptNetworkReceiver(address indexed network)
func (_BurnerRouter *BurnerRouterFilterer) WatchAcceptNetworkReceiver(opts *bind.WatchOpts, sink chan<- *BurnerRouterAcceptNetworkReceiver, network []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "AcceptNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterAcceptNetworkReceiver)
				if err := _BurnerRouter.contract.UnpackLog(event, "AcceptNetworkReceiver", log); err != nil {
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

// ParseAcceptNetworkReceiver is a log parse operation binding the contract event 0x1a2023b9b05a5599a274f08b91afd34b22b21ea58b7ca66ef06897746db55b0f.
//
// Solidity: event AcceptNetworkReceiver(address indexed network)
func (_BurnerRouter *BurnerRouterFilterer) ParseAcceptNetworkReceiver(log types.Log) (*BurnerRouterAcceptNetworkReceiver, error) {
	event := new(BurnerRouterAcceptNetworkReceiver)
	if err := _BurnerRouter.contract.UnpackLog(event, "AcceptNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterAcceptOperatorNetworkReceiverIterator is returned from FilterAcceptOperatorNetworkReceiver and is used to iterate over the raw logs and unpacked data for AcceptOperatorNetworkReceiver events raised by the BurnerRouter contract.
type BurnerRouterAcceptOperatorNetworkReceiverIterator struct {
	Event *BurnerRouterAcceptOperatorNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *BurnerRouterAcceptOperatorNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterAcceptOperatorNetworkReceiver)
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
		it.Event = new(BurnerRouterAcceptOperatorNetworkReceiver)
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
func (it *BurnerRouterAcceptOperatorNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterAcceptOperatorNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterAcceptOperatorNetworkReceiver represents a AcceptOperatorNetworkReceiver event raised by the BurnerRouter contract.
type BurnerRouterAcceptOperatorNetworkReceiver struct {
	Network  common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptOperatorNetworkReceiver is a free log retrieval operation binding the contract event 0x1261e5a4e7d8e8b5c4b7a8205d04deb702f9aa1eec8959839252b0636c6e45ab.
//
// Solidity: event AcceptOperatorNetworkReceiver(address indexed network, address indexed operator)
func (_BurnerRouter *BurnerRouterFilterer) FilterAcceptOperatorNetworkReceiver(opts *bind.FilterOpts, network []common.Address, operator []common.Address) (*BurnerRouterAcceptOperatorNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "AcceptOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterAcceptOperatorNetworkReceiverIterator{contract: _BurnerRouter.contract, event: "AcceptOperatorNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchAcceptOperatorNetworkReceiver is a free log subscription operation binding the contract event 0x1261e5a4e7d8e8b5c4b7a8205d04deb702f9aa1eec8959839252b0636c6e45ab.
//
// Solidity: event AcceptOperatorNetworkReceiver(address indexed network, address indexed operator)
func (_BurnerRouter *BurnerRouterFilterer) WatchAcceptOperatorNetworkReceiver(opts *bind.WatchOpts, sink chan<- *BurnerRouterAcceptOperatorNetworkReceiver, network []common.Address, operator []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "AcceptOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterAcceptOperatorNetworkReceiver)
				if err := _BurnerRouter.contract.UnpackLog(event, "AcceptOperatorNetworkReceiver", log); err != nil {
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

// ParseAcceptOperatorNetworkReceiver is a log parse operation binding the contract event 0x1261e5a4e7d8e8b5c4b7a8205d04deb702f9aa1eec8959839252b0636c6e45ab.
//
// Solidity: event AcceptOperatorNetworkReceiver(address indexed network, address indexed operator)
func (_BurnerRouter *BurnerRouterFilterer) ParseAcceptOperatorNetworkReceiver(log types.Log) (*BurnerRouterAcceptOperatorNetworkReceiver, error) {
	event := new(BurnerRouterAcceptOperatorNetworkReceiver)
	if err := _BurnerRouter.contract.UnpackLog(event, "AcceptOperatorNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BurnerRouter contract.
type BurnerRouterInitializedIterator struct {
	Event *BurnerRouterInitialized // Event containing the contract specifics and raw log

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
func (it *BurnerRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterInitialized)
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
		it.Event = new(BurnerRouterInitialized)
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
func (it *BurnerRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterInitialized represents a Initialized event raised by the BurnerRouter contract.
type BurnerRouterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BurnerRouter *BurnerRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*BurnerRouterInitializedIterator, error) {

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BurnerRouterInitializedIterator{contract: _BurnerRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BurnerRouter *BurnerRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BurnerRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterInitialized)
				if err := _BurnerRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BurnerRouter *BurnerRouterFilterer) ParseInitialized(log types.Log) (*BurnerRouterInitialized, error) {
	event := new(BurnerRouterInitialized)
	if err := _BurnerRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BurnerRouter contract.
type BurnerRouterOwnershipTransferredIterator struct {
	Event *BurnerRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BurnerRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterOwnershipTransferred)
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
		it.Event = new(BurnerRouterOwnershipTransferred)
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
func (it *BurnerRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterOwnershipTransferred represents a OwnershipTransferred event raised by the BurnerRouter contract.
type BurnerRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BurnerRouter *BurnerRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BurnerRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterOwnershipTransferredIterator{contract: _BurnerRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BurnerRouter *BurnerRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnerRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterOwnershipTransferred)
				if err := _BurnerRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BurnerRouter *BurnerRouterFilterer) ParseOwnershipTransferred(log types.Log) (*BurnerRouterOwnershipTransferred, error) {
	event := new(BurnerRouterOwnershipTransferred)
	if err := _BurnerRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterSetDelayIterator is returned from FilterSetDelay and is used to iterate over the raw logs and unpacked data for SetDelay events raised by the BurnerRouter contract.
type BurnerRouterSetDelayIterator struct {
	Event *BurnerRouterSetDelay // Event containing the contract specifics and raw log

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
func (it *BurnerRouterSetDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterSetDelay)
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
		it.Event = new(BurnerRouterSetDelay)
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
func (it *BurnerRouterSetDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterSetDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterSetDelay represents a SetDelay event raised by the BurnerRouter contract.
type BurnerRouterSetDelay struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetDelay is a free log retrieval operation binding the contract event 0xc4694f5e679fbd4fa31b993053f3134c2857558c12fe87ce9ea6bf3b1ef21770.
//
// Solidity: event SetDelay(uint48 delay)
func (_BurnerRouter *BurnerRouterFilterer) FilterSetDelay(opts *bind.FilterOpts) (*BurnerRouterSetDelayIterator, error) {

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "SetDelay")
	if err != nil {
		return nil, err
	}
	return &BurnerRouterSetDelayIterator{contract: _BurnerRouter.contract, event: "SetDelay", logs: logs, sub: sub}, nil
}

// WatchSetDelay is a free log subscription operation binding the contract event 0xc4694f5e679fbd4fa31b993053f3134c2857558c12fe87ce9ea6bf3b1ef21770.
//
// Solidity: event SetDelay(uint48 delay)
func (_BurnerRouter *BurnerRouterFilterer) WatchSetDelay(opts *bind.WatchOpts, sink chan<- *BurnerRouterSetDelay) (event.Subscription, error) {

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "SetDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterSetDelay)
				if err := _BurnerRouter.contract.UnpackLog(event, "SetDelay", log); err != nil {
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

// ParseSetDelay is a log parse operation binding the contract event 0xc4694f5e679fbd4fa31b993053f3134c2857558c12fe87ce9ea6bf3b1ef21770.
//
// Solidity: event SetDelay(uint48 delay)
func (_BurnerRouter *BurnerRouterFilterer) ParseSetDelay(log types.Log) (*BurnerRouterSetDelay, error) {
	event := new(BurnerRouterSetDelay)
	if err := _BurnerRouter.contract.UnpackLog(event, "SetDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterSetGlobalReceiverIterator is returned from FilterSetGlobalReceiver and is used to iterate over the raw logs and unpacked data for SetGlobalReceiver events raised by the BurnerRouter contract.
type BurnerRouterSetGlobalReceiverIterator struct {
	Event *BurnerRouterSetGlobalReceiver // Event containing the contract specifics and raw log

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
func (it *BurnerRouterSetGlobalReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterSetGlobalReceiver)
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
		it.Event = new(BurnerRouterSetGlobalReceiver)
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
func (it *BurnerRouterSetGlobalReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterSetGlobalReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterSetGlobalReceiver represents a SetGlobalReceiver event raised by the BurnerRouter contract.
type BurnerRouterSetGlobalReceiver struct {
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalReceiver is a free log retrieval operation binding the contract event 0x81c31ea2c5656f89fc438850c31cc9b7ccd45beec811b7e0a71c64a98b61f7c5.
//
// Solidity: event SetGlobalReceiver(address receiver)
func (_BurnerRouter *BurnerRouterFilterer) FilterSetGlobalReceiver(opts *bind.FilterOpts) (*BurnerRouterSetGlobalReceiverIterator, error) {

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "SetGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return &BurnerRouterSetGlobalReceiverIterator{contract: _BurnerRouter.contract, event: "SetGlobalReceiver", logs: logs, sub: sub}, nil
}

// WatchSetGlobalReceiver is a free log subscription operation binding the contract event 0x81c31ea2c5656f89fc438850c31cc9b7ccd45beec811b7e0a71c64a98b61f7c5.
//
// Solidity: event SetGlobalReceiver(address receiver)
func (_BurnerRouter *BurnerRouterFilterer) WatchSetGlobalReceiver(opts *bind.WatchOpts, sink chan<- *BurnerRouterSetGlobalReceiver) (event.Subscription, error) {

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "SetGlobalReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterSetGlobalReceiver)
				if err := _BurnerRouter.contract.UnpackLog(event, "SetGlobalReceiver", log); err != nil {
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

// ParseSetGlobalReceiver is a log parse operation binding the contract event 0x81c31ea2c5656f89fc438850c31cc9b7ccd45beec811b7e0a71c64a98b61f7c5.
//
// Solidity: event SetGlobalReceiver(address receiver)
func (_BurnerRouter *BurnerRouterFilterer) ParseSetGlobalReceiver(log types.Log) (*BurnerRouterSetGlobalReceiver, error) {
	event := new(BurnerRouterSetGlobalReceiver)
	if err := _BurnerRouter.contract.UnpackLog(event, "SetGlobalReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterSetNetworkReceiverIterator is returned from FilterSetNetworkReceiver and is used to iterate over the raw logs and unpacked data for SetNetworkReceiver events raised by the BurnerRouter contract.
type BurnerRouterSetNetworkReceiverIterator struct {
	Event *BurnerRouterSetNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *BurnerRouterSetNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterSetNetworkReceiver)
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
		it.Event = new(BurnerRouterSetNetworkReceiver)
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
func (it *BurnerRouterSetNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterSetNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterSetNetworkReceiver represents a SetNetworkReceiver event raised by the BurnerRouter contract.
type BurnerRouterSetNetworkReceiver struct {
	Network  common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetNetworkReceiver is a free log retrieval operation binding the contract event 0xd324c14c83226723f8446d113edef5f1e51f1bcf8ac2a583ae5f5e7f27808f3f.
//
// Solidity: event SetNetworkReceiver(address indexed network, address receiver)
func (_BurnerRouter *BurnerRouterFilterer) FilterSetNetworkReceiver(opts *bind.FilterOpts, network []common.Address) (*BurnerRouterSetNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "SetNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterSetNetworkReceiverIterator{contract: _BurnerRouter.contract, event: "SetNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchSetNetworkReceiver is a free log subscription operation binding the contract event 0xd324c14c83226723f8446d113edef5f1e51f1bcf8ac2a583ae5f5e7f27808f3f.
//
// Solidity: event SetNetworkReceiver(address indexed network, address receiver)
func (_BurnerRouter *BurnerRouterFilterer) WatchSetNetworkReceiver(opts *bind.WatchOpts, sink chan<- *BurnerRouterSetNetworkReceiver, network []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "SetNetworkReceiver", networkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterSetNetworkReceiver)
				if err := _BurnerRouter.contract.UnpackLog(event, "SetNetworkReceiver", log); err != nil {
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

// ParseSetNetworkReceiver is a log parse operation binding the contract event 0xd324c14c83226723f8446d113edef5f1e51f1bcf8ac2a583ae5f5e7f27808f3f.
//
// Solidity: event SetNetworkReceiver(address indexed network, address receiver)
func (_BurnerRouter *BurnerRouterFilterer) ParseSetNetworkReceiver(log types.Log) (*BurnerRouterSetNetworkReceiver, error) {
	event := new(BurnerRouterSetNetworkReceiver)
	if err := _BurnerRouter.contract.UnpackLog(event, "SetNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterSetOperatorNetworkReceiverIterator is returned from FilterSetOperatorNetworkReceiver and is used to iterate over the raw logs and unpacked data for SetOperatorNetworkReceiver events raised by the BurnerRouter contract.
type BurnerRouterSetOperatorNetworkReceiverIterator struct {
	Event *BurnerRouterSetOperatorNetworkReceiver // Event containing the contract specifics and raw log

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
func (it *BurnerRouterSetOperatorNetworkReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterSetOperatorNetworkReceiver)
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
		it.Event = new(BurnerRouterSetOperatorNetworkReceiver)
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
func (it *BurnerRouterSetOperatorNetworkReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterSetOperatorNetworkReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterSetOperatorNetworkReceiver represents a SetOperatorNetworkReceiver event raised by the BurnerRouter contract.
type BurnerRouterSetOperatorNetworkReceiver struct {
	Network  common.Address
	Operator common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOperatorNetworkReceiver is a free log retrieval operation binding the contract event 0x3692549eb3eb5e4546a8e42a78f360aaa361c0faf3345292813dfdfbcef3c887.
//
// Solidity: event SetOperatorNetworkReceiver(address indexed network, address indexed operator, address receiver)
func (_BurnerRouter *BurnerRouterFilterer) FilterSetOperatorNetworkReceiver(opts *bind.FilterOpts, network []common.Address, operator []common.Address) (*BurnerRouterSetOperatorNetworkReceiverIterator, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "SetOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterSetOperatorNetworkReceiverIterator{contract: _BurnerRouter.contract, event: "SetOperatorNetworkReceiver", logs: logs, sub: sub}, nil
}

// WatchSetOperatorNetworkReceiver is a free log subscription operation binding the contract event 0x3692549eb3eb5e4546a8e42a78f360aaa361c0faf3345292813dfdfbcef3c887.
//
// Solidity: event SetOperatorNetworkReceiver(address indexed network, address indexed operator, address receiver)
func (_BurnerRouter *BurnerRouterFilterer) WatchSetOperatorNetworkReceiver(opts *bind.WatchOpts, sink chan<- *BurnerRouterSetOperatorNetworkReceiver, network []common.Address, operator []common.Address) (event.Subscription, error) {

	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "SetOperatorNetworkReceiver", networkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterSetOperatorNetworkReceiver)
				if err := _BurnerRouter.contract.UnpackLog(event, "SetOperatorNetworkReceiver", log); err != nil {
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

// ParseSetOperatorNetworkReceiver is a log parse operation binding the contract event 0x3692549eb3eb5e4546a8e42a78f360aaa361c0faf3345292813dfdfbcef3c887.
//
// Solidity: event SetOperatorNetworkReceiver(address indexed network, address indexed operator, address receiver)
func (_BurnerRouter *BurnerRouterFilterer) ParseSetOperatorNetworkReceiver(log types.Log) (*BurnerRouterSetOperatorNetworkReceiver, error) {
	event := new(BurnerRouterSetOperatorNetworkReceiver)
	if err := _BurnerRouter.contract.UnpackLog(event, "SetOperatorNetworkReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BurnerRouterTriggerTransferIterator is returned from FilterTriggerTransfer and is used to iterate over the raw logs and unpacked data for TriggerTransfer events raised by the BurnerRouter contract.
type BurnerRouterTriggerTransferIterator struct {
	Event *BurnerRouterTriggerTransfer // Event containing the contract specifics and raw log

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
func (it *BurnerRouterTriggerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnerRouterTriggerTransfer)
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
		it.Event = new(BurnerRouterTriggerTransfer)
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
func (it *BurnerRouterTriggerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnerRouterTriggerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnerRouterTriggerTransfer represents a TriggerTransfer event raised by the BurnerRouter contract.
type BurnerRouterTriggerTransfer struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTriggerTransfer is a free log retrieval operation binding the contract event 0xd5be285f1b0878becfe756e58f0cf3aa449bc4c406c2aae066f3a33d54e01ecf.
//
// Solidity: event TriggerTransfer(address indexed receiver, uint256 amount)
func (_BurnerRouter *BurnerRouterFilterer) FilterTriggerTransfer(opts *bind.FilterOpts, receiver []common.Address) (*BurnerRouterTriggerTransferIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BurnerRouter.contract.FilterLogs(opts, "TriggerTransfer", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BurnerRouterTriggerTransferIterator{contract: _BurnerRouter.contract, event: "TriggerTransfer", logs: logs, sub: sub}, nil
}

// WatchTriggerTransfer is a free log subscription operation binding the contract event 0xd5be285f1b0878becfe756e58f0cf3aa449bc4c406c2aae066f3a33d54e01ecf.
//
// Solidity: event TriggerTransfer(address indexed receiver, uint256 amount)
func (_BurnerRouter *BurnerRouterFilterer) WatchTriggerTransfer(opts *bind.WatchOpts, sink chan<- *BurnerRouterTriggerTransfer, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BurnerRouter.contract.WatchLogs(opts, "TriggerTransfer", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnerRouterTriggerTransfer)
				if err := _BurnerRouter.contract.UnpackLog(event, "TriggerTransfer", log); err != nil {
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

// ParseTriggerTransfer is a log parse operation binding the contract event 0xd5be285f1b0878becfe756e58f0cf3aa449bc4c406c2aae066f3a33d54e01ecf.
//
// Solidity: event TriggerTransfer(address indexed receiver, uint256 amount)
func (_BurnerRouter *BurnerRouterFilterer) ParseTriggerTransfer(log types.Log) (*BurnerRouterTriggerTransfer, error) {
	event := new(BurnerRouterTriggerTransfer)
	if err := _BurnerRouter.contract.UnpackLog(event, "TriggerTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
