// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// LagrangeServiceMetaData contains all meta data concerning the LagrangeService contract.
var LagrangeServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"_committee\",\"type\":\"address\"},{\"internalType\":\"contractIStakeManager\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_avsDirectoryAddress\",\"type\":\"address\"},{\"internalType\":\"contractIVoteWeigher\",\"name\":\"_voteWeigher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"OperatorSubscribed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"OperatorUnsubscribed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[2][]\",\"name\":\"additionalBlsPubKeys\",\"type\":\"uint256[2][]\"}],\"name\":\"addBlsPubKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"name\":\"addOperatorsToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"contractIAVSDirectory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"committee\",\"outputs\":[{\"internalType\":\"contractILagrangeCommittee\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[2][]\",\"name\":\"blsPubKeys\",\"type\":\"uint256[2][]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"indices\",\"type\":\"uint32[]\"}],\"name\":\"removeBlsPubKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"}],\"name\":\"removeOperatorsFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"subscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"unsubscribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"uint256[2]\",\"name\":\"blsPubKey\",\"type\":\"uint256[2]\"}],\"name\":\"updateBlsPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSignAddress\",\"type\":\"address\"}],\"name\":\"updateSignAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteWeigher\",\"outputs\":[{\"internalType\":\"contractIVoteWeigher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LagrangeServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use LagrangeServiceMetaData.ABI instead.
var LagrangeServiceABI = LagrangeServiceMetaData.ABI

// LagrangeService is an auto generated Go binding around an Ethereum contract.
type LagrangeService struct {
	LagrangeServiceCaller     // Read-only binding to the contract
	LagrangeServiceTransactor // Write-only binding to the contract
	LagrangeServiceFilterer   // Log filterer for contract events
}

// LagrangeServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type LagrangeServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LagrangeServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LagrangeServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LagrangeServiceSession struct {
	Contract     *LagrangeService  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LagrangeServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LagrangeServiceCallerSession struct {
	Contract *LagrangeServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LagrangeServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LagrangeServiceTransactorSession struct {
	Contract     *LagrangeServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LagrangeServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type LagrangeServiceRaw struct {
	Contract *LagrangeService // Generic contract binding to access the raw methods on
}

// LagrangeServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LagrangeServiceCallerRaw struct {
	Contract *LagrangeServiceCaller // Generic read-only contract binding to access the raw methods on
}

// LagrangeServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LagrangeServiceTransactorRaw struct {
	Contract *LagrangeServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLagrangeService creates a new instance of LagrangeService, bound to a specific deployed contract.
func NewLagrangeService(address common.Address, backend bind.ContractBackend) (*LagrangeService, error) {
	contract, err := bindLagrangeService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LagrangeService{LagrangeServiceCaller: LagrangeServiceCaller{contract: contract}, LagrangeServiceTransactor: LagrangeServiceTransactor{contract: contract}, LagrangeServiceFilterer: LagrangeServiceFilterer{contract: contract}}, nil
}

// NewLagrangeServiceCaller creates a new read-only instance of LagrangeService, bound to a specific deployed contract.
func NewLagrangeServiceCaller(address common.Address, caller bind.ContractCaller) (*LagrangeServiceCaller, error) {
	contract, err := bindLagrangeService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceCaller{contract: contract}, nil
}

// NewLagrangeServiceTransactor creates a new write-only instance of LagrangeService, bound to a specific deployed contract.
func NewLagrangeServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*LagrangeServiceTransactor, error) {
	contract, err := bindLagrangeService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceTransactor{contract: contract}, nil
}

// NewLagrangeServiceFilterer creates a new log filterer instance of LagrangeService, bound to a specific deployed contract.
func NewLagrangeServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*LagrangeServiceFilterer, error) {
	contract, err := bindLagrangeService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceFilterer{contract: contract}, nil
}

// bindLagrangeService binds a generic wrapper to an already deployed contract.
func bindLagrangeService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LagrangeServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LagrangeService *LagrangeServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LagrangeService.Contract.LagrangeServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LagrangeService *LagrangeServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeService.Contract.LagrangeServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LagrangeService *LagrangeServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LagrangeService.Contract.LagrangeServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LagrangeService *LagrangeServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LagrangeService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LagrangeService *LagrangeServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LagrangeService *LagrangeServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LagrangeService.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_LagrangeService *LagrangeServiceCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_LagrangeService *LagrangeServiceSession) AvsDirectory() (common.Address, error) {
	return _LagrangeService.Contract.AvsDirectory(&_LagrangeService.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_LagrangeService *LagrangeServiceCallerSession) AvsDirectory() (common.Address, error) {
	return _LagrangeService.Contract.AvsDirectory(&_LagrangeService.CallOpts)
}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_LagrangeService *LagrangeServiceCaller) Committee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "committee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_LagrangeService *LagrangeServiceSession) Committee() (common.Address, error) {
	return _LagrangeService.Contract.Committee(&_LagrangeService.CallOpts)
}

// Committee is a free data retrieval call binding the contract method 0xd864e740.
//
// Solidity: function committee() view returns(address)
func (_LagrangeService *LagrangeServiceCallerSession) Committee() (common.Address, error) {
	return _LagrangeService.Contract.Committee(&_LagrangeService.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_LagrangeService *LagrangeServiceCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_LagrangeService *LagrangeServiceSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _LagrangeService.Contract.GetOperatorRestakedStrategies(&_LagrangeService.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_LagrangeService *LagrangeServiceCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _LagrangeService.Contract.GetOperatorRestakedStrategies(&_LagrangeService.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_LagrangeService *LagrangeServiceCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_LagrangeService *LagrangeServiceSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _LagrangeService.Contract.GetRestakeableStrategies(&_LagrangeService.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_LagrangeService *LagrangeServiceCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _LagrangeService.Contract.GetRestakeableStrategies(&_LagrangeService.CallOpts)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_LagrangeService *LagrangeServiceCaller) OperatorWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "operatorWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_LagrangeService *LagrangeServiceSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _LagrangeService.Contract.OperatorWhitelist(&_LagrangeService.CallOpts, arg0)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_LagrangeService *LagrangeServiceCallerSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _LagrangeService.Contract.OperatorWhitelist(&_LagrangeService.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeService *LagrangeServiceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeService *LagrangeServiceSession) Owner() (common.Address, error) {
	return _LagrangeService.Contract.Owner(&_LagrangeService.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeService *LagrangeServiceCallerSession) Owner() (common.Address, error) {
	return _LagrangeService.Contract.Owner(&_LagrangeService.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_LagrangeService *LagrangeServiceCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_LagrangeService *LagrangeServiceSession) StakeManager() (common.Address, error) {
	return _LagrangeService.Contract.StakeManager(&_LagrangeService.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_LagrangeService *LagrangeServiceCallerSession) StakeManager() (common.Address, error) {
	return _LagrangeService.Contract.StakeManager(&_LagrangeService.CallOpts)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_LagrangeService *LagrangeServiceCaller) VoteWeigher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeService.contract.Call(opts, &out, "voteWeigher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_LagrangeService *LagrangeServiceSession) VoteWeigher() (common.Address, error) {
	return _LagrangeService.Contract.VoteWeigher(&_LagrangeService.CallOpts)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_LagrangeService *LagrangeServiceCallerSession) VoteWeigher() (common.Address, error) {
	return _LagrangeService.Contract.VoteWeigher(&_LagrangeService.CallOpts)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x992b098e.
//
// Solidity: function addBlsPubKeys(uint256[2][] additionalBlsPubKeys) returns()
func (_LagrangeService *LagrangeServiceTransactor) AddBlsPubKeys(opts *bind.TransactOpts, additionalBlsPubKeys [][2]*big.Int) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "addBlsPubKeys", additionalBlsPubKeys)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x992b098e.
//
// Solidity: function addBlsPubKeys(uint256[2][] additionalBlsPubKeys) returns()
func (_LagrangeService *LagrangeServiceSession) AddBlsPubKeys(additionalBlsPubKeys [][2]*big.Int) (*types.Transaction, error) {
	return _LagrangeService.Contract.AddBlsPubKeys(&_LagrangeService.TransactOpts, additionalBlsPubKeys)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0x992b098e.
//
// Solidity: function addBlsPubKeys(uint256[2][] additionalBlsPubKeys) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) AddBlsPubKeys(additionalBlsPubKeys [][2]*big.Int) (*types.Transaction, error) {
	return _LagrangeService.Contract.AddBlsPubKeys(&_LagrangeService.TransactOpts, additionalBlsPubKeys)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_LagrangeService *LagrangeServiceTransactor) AddOperatorsToWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "addOperatorsToWhitelist", operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_LagrangeService *LagrangeServiceSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.AddOperatorsToWhitelist(&_LagrangeService.TransactOpts, operators)
}

// AddOperatorsToWhitelist is a paid mutator transaction binding the contract method 0x1be4b9f7.
//
// Solidity: function addOperatorsToWhitelist(address[] operators) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) AddOperatorsToWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.AddOperatorsToWhitelist(&_LagrangeService.TransactOpts, operators)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_LagrangeService *LagrangeServiceTransactor) Deregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "deregister")
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_LagrangeService *LagrangeServiceSession) Deregister() (*types.Transaction, error) {
	return _LagrangeService.Contract.Deregister(&_LagrangeService.TransactOpts)
}

// Deregister is a paid mutator transaction binding the contract method 0xaff5edb1.
//
// Solidity: function deregister() returns()
func (_LagrangeService *LagrangeServiceTransactorSession) Deregister() (*types.Transaction, error) {
	return _LagrangeService.Contract.Deregister(&_LagrangeService.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LagrangeService *LagrangeServiceTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LagrangeService *LagrangeServiceSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.Initialize(&_LagrangeService.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.Initialize(&_LagrangeService.TransactOpts, initialOwner)
}

// Register is a paid mutator transaction binding the contract method 0x7a79ea33.
//
// Solidity: function register(address signAddress, uint256[2][] blsPubKeys, (bytes,bytes32,uint256) operatorSignature) returns()
func (_LagrangeService *LagrangeServiceTransactor) Register(opts *bind.TransactOpts, signAddress common.Address, blsPubKeys [][2]*big.Int, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "register", signAddress, blsPubKeys, operatorSignature)
}

// Register is a paid mutator transaction binding the contract method 0x7a79ea33.
//
// Solidity: function register(address signAddress, uint256[2][] blsPubKeys, (bytes,bytes32,uint256) operatorSignature) returns()
func (_LagrangeService *LagrangeServiceSession) Register(signAddress common.Address, blsPubKeys [][2]*big.Int, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _LagrangeService.Contract.Register(&_LagrangeService.TransactOpts, signAddress, blsPubKeys, operatorSignature)
}

// Register is a paid mutator transaction binding the contract method 0x7a79ea33.
//
// Solidity: function register(address signAddress, uint256[2][] blsPubKeys, (bytes,bytes32,uint256) operatorSignature) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) Register(signAddress common.Address, blsPubKeys [][2]*big.Int, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _LagrangeService.Contract.Register(&_LagrangeService.TransactOpts, signAddress, blsPubKeys, operatorSignature)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x5e332fe4.
//
// Solidity: function removeBlsPubKeys(uint32[] indices) returns()
func (_LagrangeService *LagrangeServiceTransactor) RemoveBlsPubKeys(opts *bind.TransactOpts, indices []uint32) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "removeBlsPubKeys", indices)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x5e332fe4.
//
// Solidity: function removeBlsPubKeys(uint32[] indices) returns()
func (_LagrangeService *LagrangeServiceSession) RemoveBlsPubKeys(indices []uint32) (*types.Transaction, error) {
	return _LagrangeService.Contract.RemoveBlsPubKeys(&_LagrangeService.TransactOpts, indices)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x5e332fe4.
//
// Solidity: function removeBlsPubKeys(uint32[] indices) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) RemoveBlsPubKeys(indices []uint32) (*types.Transaction, error) {
	return _LagrangeService.Contract.RemoveBlsPubKeys(&_LagrangeService.TransactOpts, indices)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_LagrangeService *LagrangeServiceTransactor) RemoveOperatorsFromWhitelist(opts *bind.TransactOpts, operators []common.Address) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "removeOperatorsFromWhitelist", operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_LagrangeService *LagrangeServiceSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.RemoveOperatorsFromWhitelist(&_LagrangeService.TransactOpts, operators)
}

// RemoveOperatorsFromWhitelist is a paid mutator transaction binding the contract method 0xd822af3c.
//
// Solidity: function removeOperatorsFromWhitelist(address[] operators) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) RemoveOperatorsFromWhitelist(operators []common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.RemoveOperatorsFromWhitelist(&_LagrangeService.TransactOpts, operators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeService *LagrangeServiceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeService *LagrangeServiceSession) RenounceOwnership() (*types.Transaction, error) {
	return _LagrangeService.Contract.RenounceOwnership(&_LagrangeService.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeService *LagrangeServiceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LagrangeService.Contract.RenounceOwnership(&_LagrangeService.TransactOpts)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_LagrangeService *LagrangeServiceTransactor) Subscribe(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "subscribe", chainID)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_LagrangeService *LagrangeServiceSession) Subscribe(chainID uint32) (*types.Transaction, error) {
	return _LagrangeService.Contract.Subscribe(&_LagrangeService.TransactOpts, chainID)
}

// Subscribe is a paid mutator transaction binding the contract method 0x2e94d67b.
//
// Solidity: function subscribe(uint32 chainID) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) Subscribe(chainID uint32) (*types.Transaction, error) {
	return _LagrangeService.Contract.Subscribe(&_LagrangeService.TransactOpts, chainID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeService *LagrangeServiceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeService *LagrangeServiceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.TransferOwnership(&_LagrangeService.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.TransferOwnership(&_LagrangeService.TransactOpts, newOwner)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_LagrangeService *LagrangeServiceTransactor) Unsubscribe(opts *bind.TransactOpts, chainID uint32) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "unsubscribe", chainID)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_LagrangeService *LagrangeServiceSession) Unsubscribe(chainID uint32) (*types.Transaction, error) {
	return _LagrangeService.Contract.Unsubscribe(&_LagrangeService.TransactOpts, chainID)
}

// Unsubscribe is a paid mutator transaction binding the contract method 0x0512d04c.
//
// Solidity: function unsubscribe(uint32 chainID) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) Unsubscribe(chainID uint32) (*types.Transaction, error) {
	return _LagrangeService.Contract.Unsubscribe(&_LagrangeService.TransactOpts, chainID)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_LagrangeService *LagrangeServiceTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_LagrangeService *LagrangeServiceSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _LagrangeService.Contract.UpdateAVSMetadataURI(&_LagrangeService.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _LagrangeService.Contract.UpdateAVSMetadataURI(&_LagrangeService.TransactOpts, _metadataURI)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc8d8e013.
//
// Solidity: function updateBlsPubKey(uint32 index, uint256[2] blsPubKey) returns()
func (_LagrangeService *LagrangeServiceTransactor) UpdateBlsPubKey(opts *bind.TransactOpts, index uint32, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "updateBlsPubKey", index, blsPubKey)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc8d8e013.
//
// Solidity: function updateBlsPubKey(uint32 index, uint256[2] blsPubKey) returns()
func (_LagrangeService *LagrangeServiceSession) UpdateBlsPubKey(index uint32, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _LagrangeService.Contract.UpdateBlsPubKey(&_LagrangeService.TransactOpts, index, blsPubKey)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0xc8d8e013.
//
// Solidity: function updateBlsPubKey(uint32 index, uint256[2] blsPubKey) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) UpdateBlsPubKey(index uint32, blsPubKey [2]*big.Int) (*types.Transaction, error) {
	return _LagrangeService.Contract.UpdateBlsPubKey(&_LagrangeService.TransactOpts, index, blsPubKey)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x5fa7306a.
//
// Solidity: function updateSignAddress(address newSignAddress) returns()
func (_LagrangeService *LagrangeServiceTransactor) UpdateSignAddress(opts *bind.TransactOpts, newSignAddress common.Address) (*types.Transaction, error) {
	return _LagrangeService.contract.Transact(opts, "updateSignAddress", newSignAddress)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x5fa7306a.
//
// Solidity: function updateSignAddress(address newSignAddress) returns()
func (_LagrangeService *LagrangeServiceSession) UpdateSignAddress(newSignAddress common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.UpdateSignAddress(&_LagrangeService.TransactOpts, newSignAddress)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x5fa7306a.
//
// Solidity: function updateSignAddress(address newSignAddress) returns()
func (_LagrangeService *LagrangeServiceTransactorSession) UpdateSignAddress(newSignAddress common.Address) (*types.Transaction, error) {
	return _LagrangeService.Contract.UpdateSignAddress(&_LagrangeService.TransactOpts, newSignAddress)
}

// LagrangeServiceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LagrangeService contract.
type LagrangeServiceInitializedIterator struct {
	Event *LagrangeServiceInitialized // Event containing the contract specifics and raw log

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
func (it *LagrangeServiceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeServiceInitialized)
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
		it.Event = new(LagrangeServiceInitialized)
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
func (it *LagrangeServiceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeServiceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeServiceInitialized represents a Initialized event raised by the LagrangeService contract.
type LagrangeServiceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LagrangeService *LagrangeServiceFilterer) FilterInitialized(opts *bind.FilterOpts) (*LagrangeServiceInitializedIterator, error) {

	logs, sub, err := _LagrangeService.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceInitializedIterator{contract: _LagrangeService.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LagrangeService *LagrangeServiceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LagrangeServiceInitialized) (event.Subscription, error) {

	logs, sub, err := _LagrangeService.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeServiceInitialized)
				if err := _LagrangeService.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LagrangeService *LagrangeServiceFilterer) ParseInitialized(log types.Log) (*LagrangeServiceInitialized, error) {
	event := new(LagrangeServiceInitialized)
	if err := _LagrangeService.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeServiceOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the LagrangeService contract.
type LagrangeServiceOperatorDeregisteredIterator struct {
	Event *LagrangeServiceOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *LagrangeServiceOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeServiceOperatorDeregistered)
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
		it.Event = new(LagrangeServiceOperatorDeregistered)
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
func (it *LagrangeServiceOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeServiceOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeServiceOperatorDeregistered represents a OperatorDeregistered event raised by the LagrangeService contract.
type LagrangeServiceOperatorDeregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_LagrangeService *LagrangeServiceFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*LagrangeServiceOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LagrangeService.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceOperatorDeregisteredIterator{contract: _LagrangeService.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_LagrangeService *LagrangeServiceFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *LagrangeServiceOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LagrangeService.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeServiceOperatorDeregistered)
				if err := _LagrangeService.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_LagrangeService *LagrangeServiceFilterer) ParseOperatorDeregistered(log types.Log) (*LagrangeServiceOperatorDeregistered, error) {
	event := new(LagrangeServiceOperatorDeregistered)
	if err := _LagrangeService.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeServiceOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the LagrangeService contract.
type LagrangeServiceOperatorRegisteredIterator struct {
	Event *LagrangeServiceOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *LagrangeServiceOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeServiceOperatorRegistered)
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
		it.Event = new(LagrangeServiceOperatorRegistered)
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
func (it *LagrangeServiceOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeServiceOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeServiceOperatorRegistered represents a OperatorRegistered event raised by the LagrangeService contract.
type LagrangeServiceOperatorRegistered struct {
	Operator        common.Address
	ServeUntilBlock uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32 serveUntilBlock)
func (_LagrangeService *LagrangeServiceFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*LagrangeServiceOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LagrangeService.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceOperatorRegisteredIterator{contract: _LagrangeService.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32 serveUntilBlock)
func (_LagrangeService *LagrangeServiceFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *LagrangeServiceOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LagrangeService.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeServiceOperatorRegistered)
				if err := _LagrangeService.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x3ed331d6c3431aecc422f169b89a3c24f9e23cef141e10631262a3fc865f513a.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32 serveUntilBlock)
func (_LagrangeService *LagrangeServiceFilterer) ParseOperatorRegistered(log types.Log) (*LagrangeServiceOperatorRegistered, error) {
	event := new(LagrangeServiceOperatorRegistered)
	if err := _LagrangeService.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeServiceOperatorSubscribedIterator is returned from FilterOperatorSubscribed and is used to iterate over the raw logs and unpacked data for OperatorSubscribed events raised by the LagrangeService contract.
type LagrangeServiceOperatorSubscribedIterator struct {
	Event *LagrangeServiceOperatorSubscribed // Event containing the contract specifics and raw log

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
func (it *LagrangeServiceOperatorSubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeServiceOperatorSubscribed)
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
		it.Event = new(LagrangeServiceOperatorSubscribed)
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
func (it *LagrangeServiceOperatorSubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeServiceOperatorSubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeServiceOperatorSubscribed represents a OperatorSubscribed event raised by the LagrangeService contract.
type LagrangeServiceOperatorSubscribed struct {
	Operator common.Address
	ChainID  uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSubscribed is a free log retrieval operation binding the contract event 0xa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc.
//
// Solidity: event OperatorSubscribed(address indexed operator, uint32 indexed chainID)
func (_LagrangeService *LagrangeServiceFilterer) FilterOperatorSubscribed(opts *bind.FilterOpts, operator []common.Address, chainID []uint32) (*LagrangeServiceOperatorSubscribedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _LagrangeService.contract.FilterLogs(opts, "OperatorSubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceOperatorSubscribedIterator{contract: _LagrangeService.contract, event: "OperatorSubscribed", logs: logs, sub: sub}, nil
}

// WatchOperatorSubscribed is a free log subscription operation binding the contract event 0xa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc.
//
// Solidity: event OperatorSubscribed(address indexed operator, uint32 indexed chainID)
func (_LagrangeService *LagrangeServiceFilterer) WatchOperatorSubscribed(opts *bind.WatchOpts, sink chan<- *LagrangeServiceOperatorSubscribed, operator []common.Address, chainID []uint32) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _LagrangeService.contract.WatchLogs(opts, "OperatorSubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeServiceOperatorSubscribed)
				if err := _LagrangeService.contract.UnpackLog(event, "OperatorSubscribed", log); err != nil {
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

// ParseOperatorSubscribed is a log parse operation binding the contract event 0xa7aebf234bd4af17c69e39555c1da417a31714b8efc90375f4019f3024bbf4bc.
//
// Solidity: event OperatorSubscribed(address indexed operator, uint32 indexed chainID)
func (_LagrangeService *LagrangeServiceFilterer) ParseOperatorSubscribed(log types.Log) (*LagrangeServiceOperatorSubscribed, error) {
	event := new(LagrangeServiceOperatorSubscribed)
	if err := _LagrangeService.contract.UnpackLog(event, "OperatorSubscribed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeServiceOperatorUnsubscribedIterator is returned from FilterOperatorUnsubscribed and is used to iterate over the raw logs and unpacked data for OperatorUnsubscribed events raised by the LagrangeService contract.
type LagrangeServiceOperatorUnsubscribedIterator struct {
	Event *LagrangeServiceOperatorUnsubscribed // Event containing the contract specifics and raw log

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
func (it *LagrangeServiceOperatorUnsubscribedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeServiceOperatorUnsubscribed)
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
		it.Event = new(LagrangeServiceOperatorUnsubscribed)
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
func (it *LagrangeServiceOperatorUnsubscribedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeServiceOperatorUnsubscribedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeServiceOperatorUnsubscribed represents a OperatorUnsubscribed event raised by the LagrangeService contract.
type LagrangeServiceOperatorUnsubscribed struct {
	Operator common.Address
	ChainID  uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnsubscribed is a free log retrieval operation binding the contract event 0xc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d8.
//
// Solidity: event OperatorUnsubscribed(address indexed operator, uint32 indexed chainID)
func (_LagrangeService *LagrangeServiceFilterer) FilterOperatorUnsubscribed(opts *bind.FilterOpts, operator []common.Address, chainID []uint32) (*LagrangeServiceOperatorUnsubscribedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _LagrangeService.contract.FilterLogs(opts, "OperatorUnsubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceOperatorUnsubscribedIterator{contract: _LagrangeService.contract, event: "OperatorUnsubscribed", logs: logs, sub: sub}, nil
}

// WatchOperatorUnsubscribed is a free log subscription operation binding the contract event 0xc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d8.
//
// Solidity: event OperatorUnsubscribed(address indexed operator, uint32 indexed chainID)
func (_LagrangeService *LagrangeServiceFilterer) WatchOperatorUnsubscribed(opts *bind.WatchOpts, sink chan<- *LagrangeServiceOperatorUnsubscribed, operator []common.Address, chainID []uint32) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}

	logs, sub, err := _LagrangeService.contract.WatchLogs(opts, "OperatorUnsubscribed", operatorRule, chainIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeServiceOperatorUnsubscribed)
				if err := _LagrangeService.contract.UnpackLog(event, "OperatorUnsubscribed", log); err != nil {
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

// ParseOperatorUnsubscribed is a log parse operation binding the contract event 0xc839611458a2fab2b8b94182f828cd8d886dc80a56b20f644bd651d3e128c6d8.
//
// Solidity: event OperatorUnsubscribed(address indexed operator, uint32 indexed chainID)
func (_LagrangeService *LagrangeServiceFilterer) ParseOperatorUnsubscribed(log types.Log) (*LagrangeServiceOperatorUnsubscribed, error) {
	event := new(LagrangeServiceOperatorUnsubscribed)
	if err := _LagrangeService.contract.UnpackLog(event, "OperatorUnsubscribed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeServiceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LagrangeService contract.
type LagrangeServiceOwnershipTransferredIterator struct {
	Event *LagrangeServiceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LagrangeServiceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeServiceOwnershipTransferred)
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
		it.Event = new(LagrangeServiceOwnershipTransferred)
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
func (it *LagrangeServiceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeServiceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeServiceOwnershipTransferred represents a OwnershipTransferred event raised by the LagrangeService contract.
type LagrangeServiceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LagrangeService *LagrangeServiceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LagrangeServiceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LagrangeService.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeServiceOwnershipTransferredIterator{contract: _LagrangeService.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LagrangeService *LagrangeServiceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LagrangeServiceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LagrangeService.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeServiceOwnershipTransferred)
				if err := _LagrangeService.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LagrangeService *LagrangeServiceFilterer) ParseOwnershipTransferred(log types.Log) (*LagrangeServiceOwnershipTransferred, error) {
	event := new(LagrangeServiceOwnershipTransferred)
	if err := _LagrangeService.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
