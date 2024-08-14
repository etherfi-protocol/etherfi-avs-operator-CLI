// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lagrangezk

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ZKMRServiceManagerMetaData contains all meta data concerning the ZKMRServiceManager contract.
var ZKMRServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAVSDirectory\",\"name\":\"avsDirectory_\",\"type\":\"address\"},{\"internalType\":\"contractIZKMRStakeRegistry\",\"name\":\"stakeRegistry_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIZKMRStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZKMRServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKMRServiceManagerMetaData.ABI instead.
var ZKMRServiceManagerABI = ZKMRServiceManagerMetaData.ABI

// ZKMRServiceManager is an auto generated Go binding around an Ethereum contract.
type ZKMRServiceManager struct {
	ZKMRServiceManagerCaller     // Read-only binding to the contract
	ZKMRServiceManagerTransactor // Write-only binding to the contract
	ZKMRServiceManagerFilterer   // Log filterer for contract events
}

// ZKMRServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKMRServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMRServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKMRServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMRServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKMRServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMRServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKMRServiceManagerSession struct {
	Contract     *ZKMRServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZKMRServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKMRServiceManagerCallerSession struct {
	Contract *ZKMRServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ZKMRServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKMRServiceManagerTransactorSession struct {
	Contract     *ZKMRServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ZKMRServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKMRServiceManagerRaw struct {
	Contract *ZKMRServiceManager // Generic contract binding to access the raw methods on
}

// ZKMRServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKMRServiceManagerCallerRaw struct {
	Contract *ZKMRServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ZKMRServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKMRServiceManagerTransactorRaw struct {
	Contract *ZKMRServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKMRServiceManager creates a new instance of ZKMRServiceManager, bound to a specific deployed contract.
func NewZKMRServiceManager(address common.Address, backend bind.ContractBackend) (*ZKMRServiceManager, error) {
	contract, err := bindZKMRServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManager{ZKMRServiceManagerCaller: ZKMRServiceManagerCaller{contract: contract}, ZKMRServiceManagerTransactor: ZKMRServiceManagerTransactor{contract: contract}, ZKMRServiceManagerFilterer: ZKMRServiceManagerFilterer{contract: contract}}, nil
}

// NewZKMRServiceManagerCaller creates a new read-only instance of ZKMRServiceManager, bound to a specific deployed contract.
func NewZKMRServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ZKMRServiceManagerCaller, error) {
	contract, err := bindZKMRServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManagerCaller{contract: contract}, nil
}

// NewZKMRServiceManagerTransactor creates a new write-only instance of ZKMRServiceManager, bound to a specific deployed contract.
func NewZKMRServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKMRServiceManagerTransactor, error) {
	contract, err := bindZKMRServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManagerTransactor{contract: contract}, nil
}

// NewZKMRServiceManagerFilterer creates a new log filterer instance of ZKMRServiceManager, bound to a specific deployed contract.
func NewZKMRServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKMRServiceManagerFilterer, error) {
	contract, err := bindZKMRServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManagerFilterer{contract: contract}, nil
}

// bindZKMRServiceManager binds a generic wrapper to an already deployed contract.
func bindZKMRServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKMRServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMRServiceManager *ZKMRServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMRServiceManager.Contract.ZKMRServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMRServiceManager *ZKMRServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.ZKMRServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMRServiceManager *ZKMRServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.ZKMRServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMRServiceManager *ZKMRServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMRServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ZKMRServiceManager *ZKMRServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMRServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ZKMRServiceManager *ZKMRServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ZKMRServiceManager.Contract.AvsDirectory(&_ZKMRServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ZKMRServiceManager *ZKMRServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ZKMRServiceManager.Contract.AvsDirectory(&_ZKMRServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address ) view returns(address[])
func (_ZKMRServiceManager *ZKMRServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, arg0 common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ZKMRServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", arg0)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address ) view returns(address[])
func (_ZKMRServiceManager *ZKMRServiceManagerSession) GetOperatorRestakedStrategies(arg0 common.Address) ([]common.Address, error) {
	return _ZKMRServiceManager.Contract.GetOperatorRestakedStrategies(&_ZKMRServiceManager.CallOpts, arg0)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address ) view returns(address[])
func (_ZKMRServiceManager *ZKMRServiceManagerCallerSession) GetOperatorRestakedStrategies(arg0 common.Address) ([]common.Address, error) {
	return _ZKMRServiceManager.Contract.GetOperatorRestakedStrategies(&_ZKMRServiceManager.CallOpts, arg0)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ZKMRServiceManager *ZKMRServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZKMRServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ZKMRServiceManager *ZKMRServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ZKMRServiceManager.Contract.GetRestakeableStrategies(&_ZKMRServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ZKMRServiceManager *ZKMRServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ZKMRServiceManager.Contract.GetRestakeableStrategies(&_ZKMRServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ZKMRServiceManager *ZKMRServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMRServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ZKMRServiceManager *ZKMRServiceManagerSession) Owner() (common.Address, error) {
	return _ZKMRServiceManager.Contract.Owner(&_ZKMRServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ZKMRServiceManager *ZKMRServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ZKMRServiceManager.Contract.Owner(&_ZKMRServiceManager.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ZKMRServiceManager *ZKMRServiceManagerCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZKMRServiceManager.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ZKMRServiceManager *ZKMRServiceManagerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _ZKMRServiceManager.Contract.OwnershipHandoverExpiresAt(&_ZKMRServiceManager.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ZKMRServiceManager *ZKMRServiceManagerCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _ZKMRServiceManager.Contract.OwnershipHandoverExpiresAt(&_ZKMRServiceManager.CallOpts, pendingOwner)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ZKMRServiceManager *ZKMRServiceManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMRServiceManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ZKMRServiceManager *ZKMRServiceManagerSession) StakeRegistry() (common.Address, error) {
	return _ZKMRServiceManager.Contract.StakeRegistry(&_ZKMRServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ZKMRServiceManager *ZKMRServiceManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ZKMRServiceManager.Contract.StakeRegistry(&_ZKMRServiceManager.CallOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.CancelOwnershipHandover(&_ZKMRServiceManager.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.CancelOwnershipHandover(&_ZKMRServiceManager.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.CompleteOwnershipHandover(&_ZKMRServiceManager.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.CompleteOwnershipHandover(&_ZKMRServiceManager.TransactOpts, pendingOwner)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.DeregisterOperatorFromAVS(&_ZKMRServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.DeregisterOperatorFromAVS(&_ZKMRServiceManager.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address avsDirectory_, address stakeRegistry_, address owner_) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) Initialize(opts *bind.TransactOpts, avsDirectory_ common.Address, stakeRegistry_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "initialize", avsDirectory_, stakeRegistry_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address avsDirectory_, address stakeRegistry_, address owner_) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) Initialize(avsDirectory_ common.Address, stakeRegistry_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.Initialize(&_ZKMRServiceManager.TransactOpts, avsDirectory_, stakeRegistry_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address avsDirectory_, address stakeRegistry_, address owner_) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) Initialize(avsDirectory_ common.Address, stakeRegistry_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.Initialize(&_ZKMRServiceManager.TransactOpts, avsDirectory_, stakeRegistry_, owner_)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.RegisterOperatorToAVS(&_ZKMRServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.RegisterOperatorToAVS(&_ZKMRServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.RenounceOwnership(&_ZKMRServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.RenounceOwnership(&_ZKMRServiceManager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.RequestOwnershipHandover(&_ZKMRServiceManager.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.RequestOwnershipHandover(&_ZKMRServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.TransferOwnership(&_ZKMRServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.TransferOwnership(&_ZKMRServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ZKMRServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.UpdateAVSMetadataURI(&_ZKMRServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ZKMRServiceManager *ZKMRServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ZKMRServiceManager.Contract.UpdateAVSMetadataURI(&_ZKMRServiceManager.TransactOpts, _metadataURI)
}

// ZKMRServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerInitializedIterator struct {
	Event *ZKMRServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ZKMRServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRServiceManagerInitialized)
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
		it.Event = new(ZKMRServiceManagerInitialized)
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
func (it *ZKMRServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRServiceManagerInitialized represents a Initialized event raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZKMRServiceManagerInitializedIterator, error) {

	logs, sub, err := _ZKMRServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManagerInitializedIterator{contract: _ZKMRServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZKMRServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ZKMRServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRServiceManagerInitialized)
				if err := _ZKMRServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) ParseInitialized(log types.Log) (*ZKMRServiceManagerInitialized, error) {
	event := new(ZKMRServiceManagerInitialized)
	if err := _ZKMRServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRServiceManagerOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerOwnershipHandoverCanceledIterator struct {
	Event *ZKMRServiceManagerOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *ZKMRServiceManagerOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRServiceManagerOwnershipHandoverCanceled)
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
		it.Event = new(ZKMRServiceManagerOwnershipHandoverCanceled)
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
func (it *ZKMRServiceManagerOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRServiceManagerOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRServiceManagerOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*ZKMRServiceManagerOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRServiceManager.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManagerOwnershipHandoverCanceledIterator{contract: _ZKMRServiceManager.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *ZKMRServiceManagerOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRServiceManager.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRServiceManagerOwnershipHandoverCanceled)
				if err := _ZKMRServiceManager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*ZKMRServiceManagerOwnershipHandoverCanceled, error) {
	event := new(ZKMRServiceManagerOwnershipHandoverCanceled)
	if err := _ZKMRServiceManager.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRServiceManagerOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerOwnershipHandoverRequestedIterator struct {
	Event *ZKMRServiceManagerOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *ZKMRServiceManagerOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRServiceManagerOwnershipHandoverRequested)
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
		it.Event = new(ZKMRServiceManagerOwnershipHandoverRequested)
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
func (it *ZKMRServiceManagerOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRServiceManagerOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRServiceManagerOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*ZKMRServiceManagerOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRServiceManager.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManagerOwnershipHandoverRequestedIterator{contract: _ZKMRServiceManager.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *ZKMRServiceManagerOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRServiceManager.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRServiceManagerOwnershipHandoverRequested)
				if err := _ZKMRServiceManager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) ParseOwnershipHandoverRequested(log types.Log) (*ZKMRServiceManagerOwnershipHandoverRequested, error) {
	event := new(ZKMRServiceManagerOwnershipHandoverRequested)
	if err := _ZKMRServiceManager.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerOwnershipTransferredIterator struct {
	Event *ZKMRServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZKMRServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRServiceManagerOwnershipTransferred)
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
		it.Event = new(ZKMRServiceManagerOwnershipTransferred)
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
func (it *ZKMRServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ZKMRServiceManager contract.
type ZKMRServiceManagerOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ZKMRServiceManagerOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKMRServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRServiceManagerOwnershipTransferredIterator{contract: _ZKMRServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZKMRServiceManagerOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKMRServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRServiceManagerOwnershipTransferred)
				if err := _ZKMRServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_ZKMRServiceManager *ZKMRServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ZKMRServiceManagerOwnershipTransferred, error) {
	event := new(ZKMRServiceManagerOwnershipTransferred)
	if err := _ZKMRServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
