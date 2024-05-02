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
)

// IStakeRegistryStakeUpdate is an auto generated low-level Go binding around an user-defined struct.
type IStakeRegistryStakeUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	Stake                 *big.Int
}

// StakeRegistryMetaData contains all meta data concerning the StakeRegistry contract.
var StakeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegationManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"minimumStake\",\"type\":\"uint96\"}],\"name\":\"MinimumStakeForQuorumUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"name\":\"OperatorStakeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"QuorumCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyAddedToQuorum\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"StrategyMultiplierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyRemovedFromQuorum\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_WEIGHING_FUNCTION_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHTING_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"name\":\"_strategyParams\",\"type\":\"tuple[]\"}],\"name\":\"addStrategies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getCurrentStake\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getCurrentTotalStake\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getLatestStakeUpdate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StakeUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getStakeAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStakeAtBlockNumberAndIndex\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getStakeHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StakeUpdate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getStakeHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStakeUpdateAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StakeUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getStakeUpdateIndexAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getTotalStakeAtBlockNumberFromIndex\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getTotalStakeHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"getTotalStakeIndicesAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getTotalStakeUpdateAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StakeUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"minimumStake\",\"type\":\"uint96\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams[]\",\"name\":\"_strategyParams\",\"type\":\"tuple[]\"}],\"name\":\"initializeQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"minimumStakeForQuorum\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"strategyIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint96[]\",\"name\":\"newMultipliers\",\"type\":\"uint96[]\"}],\"name\":\"modifyStrategyParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"uint96[]\",\"name\":\"\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96[]\",\"name\":\"\",\"type\":\"uint96[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"indicesToRemove\",\"type\":\"uint256[]\"}],\"name\":\"removeStrategies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"minimumStake\",\"type\":\"uint96\"}],\"name\":\"setMinimumStakeForQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"strategiesPerQuorum\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"strategyParams\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"strategyParamsByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIStakeRegistry.StrategyParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"strategyParamsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"updateOperatorStake\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"weightOfOperatorForQuorum\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(StakeRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// MAXWEIGHINGFUNCTIONLENGTH is a free data retrieval call binding the contract method 0x7c172347.
//
// Solidity: function MAX_WEIGHING_FUNCTION_LENGTH() view returns(uint8)
func (_StakeRegistry *StakeRegistryCaller) MAXWEIGHINGFUNCTIONLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "MAX_WEIGHING_FUNCTION_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXWEIGHINGFUNCTIONLENGTH is a free data retrieval call binding the contract method 0x7c172347.
//
// Solidity: function MAX_WEIGHING_FUNCTION_LENGTH() view returns(uint8)
func (_StakeRegistry *StakeRegistrySession) MAXWEIGHINGFUNCTIONLENGTH() (uint8, error) {
	return _StakeRegistry.Contract.MAXWEIGHINGFUNCTIONLENGTH(&_StakeRegistry.CallOpts)
}

// MAXWEIGHINGFUNCTIONLENGTH is a free data retrieval call binding the contract method 0x7c172347.
//
// Solidity: function MAX_WEIGHING_FUNCTION_LENGTH() view returns(uint8)
func (_StakeRegistry *StakeRegistryCallerSession) MAXWEIGHINGFUNCTIONLENGTH() (uint8, error) {
	return _StakeRegistry.Contract.MAXWEIGHINGFUNCTIONLENGTH(&_StakeRegistry.CallOpts)
}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) WEIGHTINGDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "WEIGHTING_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _StakeRegistry.Contract.WEIGHTINGDIVISOR(&_StakeRegistry.CallOpts)
}

// WEIGHTINGDIVISOR is a free data retrieval call binding the contract method 0x5e5a6775.
//
// Solidity: function WEIGHTING_DIVISOR() view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) WEIGHTINGDIVISOR() (*big.Int, error) {
	return _StakeRegistry.Contract.WEIGHTINGDIVISOR(&_StakeRegistry.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StakeRegistry *StakeRegistryCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StakeRegistry *StakeRegistrySession) Delegation() (common.Address, error) {
	return _StakeRegistry.Contract.Delegation(&_StakeRegistry.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_StakeRegistry *StakeRegistryCallerSession) Delegation() (common.Address, error) {
	return _StakeRegistry.Contract.Delegation(&_StakeRegistry.CallOpts)
}

// GetCurrentStake is a free data retrieval call binding the contract method 0x5401ed27.
//
// Solidity: function getCurrentStake(bytes32 operatorId, uint8 quorumNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistryCaller) GetCurrentStake(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getCurrentStake", operatorId, quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStake is a free data retrieval call binding the contract method 0x5401ed27.
//
// Solidity: function getCurrentStake(bytes32 operatorId, uint8 quorumNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistrySession) GetCurrentStake(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetCurrentStake(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetCurrentStake is a free data retrieval call binding the contract method 0x5401ed27.
//
// Solidity: function getCurrentStake(bytes32 operatorId, uint8 quorumNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistryCallerSession) GetCurrentStake(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetCurrentStake(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetCurrentTotalStake is a free data retrieval call binding the contract method 0xd5eccc05.
//
// Solidity: function getCurrentTotalStake(uint8 quorumNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistryCaller) GetCurrentTotalStake(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getCurrentTotalStake", quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTotalStake is a free data retrieval call binding the contract method 0xd5eccc05.
//
// Solidity: function getCurrentTotalStake(uint8 quorumNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistrySession) GetCurrentTotalStake(quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetCurrentTotalStake(&_StakeRegistry.CallOpts, quorumNumber)
}

// GetCurrentTotalStake is a free data retrieval call binding the contract method 0xd5eccc05.
//
// Solidity: function getCurrentTotalStake(uint8 quorumNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistryCallerSession) GetCurrentTotalStake(quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetCurrentTotalStake(&_StakeRegistry.CallOpts, quorumNumber)
}

// GetLatestStakeUpdate is a free data retrieval call binding the contract method 0xf851e198.
//
// Solidity: function getLatestStakeUpdate(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistryCaller) GetLatestStakeUpdate(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) (IStakeRegistryStakeUpdate, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getLatestStakeUpdate", operatorId, quorumNumber)

	if err != nil {
		return *new(IStakeRegistryStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryStakeUpdate)).(*IStakeRegistryStakeUpdate)

	return out0, err

}

// GetLatestStakeUpdate is a free data retrieval call binding the contract method 0xf851e198.
//
// Solidity: function getLatestStakeUpdate(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistrySession) GetLatestStakeUpdate(operatorId [32]byte, quorumNumber uint8) (IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetLatestStakeUpdate(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetLatestStakeUpdate is a free data retrieval call binding the contract method 0xf851e198.
//
// Solidity: function getLatestStakeUpdate(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistryCallerSession) GetLatestStakeUpdate(operatorId [32]byte, quorumNumber uint8) (IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetLatestStakeUpdate(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeAtBlockNumber is a free data retrieval call binding the contract method 0xfa28c627.
//
// Solidity: function getStakeAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistryCaller) GetStakeAtBlockNumber(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getStakeAtBlockNumber", operatorId, quorumNumber, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeAtBlockNumber is a free data retrieval call binding the contract method 0xfa28c627.
//
// Solidity: function getStakeAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistrySession) GetStakeAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetStakeAtBlockNumber(&_StakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetStakeAtBlockNumber is a free data retrieval call binding the contract method 0xfa28c627.
//
// Solidity: function getStakeAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint96)
func (_StakeRegistry *StakeRegistryCallerSession) GetStakeAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (*big.Int, error) {
	return _StakeRegistry.Contract.GetStakeAtBlockNumber(&_StakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetStakeAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0xf2be94ae.
//
// Solidity: function getStakeAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index) view returns(uint96)
func (_StakeRegistry *StakeRegistryCaller) GetStakeAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, operatorId [32]byte, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getStakeAtBlockNumberAndIndex", quorumNumber, blockNumber, operatorId, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0xf2be94ae.
//
// Solidity: function getStakeAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index) view returns(uint96)
func (_StakeRegistry *StakeRegistrySession) GetStakeAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, operatorId [32]byte, index *big.Int) (*big.Int, error) {
	return _StakeRegistry.Contract.GetStakeAtBlockNumberAndIndex(&_StakeRegistry.CallOpts, quorumNumber, blockNumber, operatorId, index)
}

// GetStakeAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0xf2be94ae.
//
// Solidity: function getStakeAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, bytes32 operatorId, uint256 index) view returns(uint96)
func (_StakeRegistry *StakeRegistryCallerSession) GetStakeAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, operatorId [32]byte, index *big.Int) (*big.Int, error) {
	return _StakeRegistry.Contract.GetStakeAtBlockNumberAndIndex(&_StakeRegistry.CallOpts, quorumNumber, blockNumber, operatorId, index)
}

// GetStakeHistory is a free data retrieval call binding the contract method 0x2cd95940.
//
// Solidity: function getStakeHistory(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96)[])
func (_StakeRegistry *StakeRegistryCaller) GetStakeHistory(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) ([]IStakeRegistryStakeUpdate, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getStakeHistory", operatorId, quorumNumber)

	if err != nil {
		return *new([]IStakeRegistryStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakeRegistryStakeUpdate)).(*[]IStakeRegistryStakeUpdate)

	return out0, err

}

// GetStakeHistory is a free data retrieval call binding the contract method 0x2cd95940.
//
// Solidity: function getStakeHistory(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96)[])
func (_StakeRegistry *StakeRegistrySession) GetStakeHistory(operatorId [32]byte, quorumNumber uint8) ([]IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetStakeHistory(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeHistory is a free data retrieval call binding the contract method 0x2cd95940.
//
// Solidity: function getStakeHistory(bytes32 operatorId, uint8 quorumNumber) view returns((uint32,uint32,uint96)[])
func (_StakeRegistry *StakeRegistryCallerSession) GetStakeHistory(operatorId [32]byte, quorumNumber uint8) ([]IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetStakeHistory(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeHistoryLength is a free data retrieval call binding the contract method 0x4bd26e09.
//
// Solidity: function getStakeHistoryLength(bytes32 operatorId, uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetStakeHistoryLength(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getStakeHistoryLength", operatorId, quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeHistoryLength is a free data retrieval call binding the contract method 0x4bd26e09.
//
// Solidity: function getStakeHistoryLength(bytes32 operatorId, uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetStakeHistoryLength(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetStakeHistoryLength(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeHistoryLength is a free data retrieval call binding the contract method 0x4bd26e09.
//
// Solidity: function getStakeHistoryLength(bytes32 operatorId, uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetStakeHistoryLength(operatorId [32]byte, quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetStakeHistoryLength(&_StakeRegistry.CallOpts, operatorId, quorumNumber)
}

// GetStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xac6bfb03.
//
// Solidity: function getStakeUpdateAtIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistryCaller) GetStakeUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, operatorId [32]byte, index *big.Int) (IStakeRegistryStakeUpdate, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getStakeUpdateAtIndex", quorumNumber, operatorId, index)

	if err != nil {
		return *new(IStakeRegistryStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryStakeUpdate)).(*IStakeRegistryStakeUpdate)

	return out0, err

}

// GetStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xac6bfb03.
//
// Solidity: function getStakeUpdateAtIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistrySession) GetStakeUpdateAtIndex(quorumNumber uint8, operatorId [32]byte, index *big.Int) (IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetStakeUpdateAtIndex(&_StakeRegistry.CallOpts, quorumNumber, operatorId, index)
}

// GetStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xac6bfb03.
//
// Solidity: function getStakeUpdateAtIndex(uint8 quorumNumber, bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistryCallerSession) GetStakeUpdateAtIndex(quorumNumber uint8, operatorId [32]byte, index *big.Int) (IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetStakeUpdateAtIndex(&_StakeRegistry.CallOpts, quorumNumber, operatorId, index)
}

// GetStakeUpdateIndexAtBlockNumber is a free data retrieval call binding the contract method 0xdd9846b9.
//
// Solidity: function getStakeUpdateIndexAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_StakeRegistry *StakeRegistryCaller) GetStakeUpdateIndexAtBlockNumber(opts *bind.CallOpts, operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (uint32, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getStakeUpdateIndexAtBlockNumber", operatorId, quorumNumber, blockNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStakeUpdateIndexAtBlockNumber is a free data retrieval call binding the contract method 0xdd9846b9.
//
// Solidity: function getStakeUpdateIndexAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_StakeRegistry *StakeRegistrySession) GetStakeUpdateIndexAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (uint32, error) {
	return _StakeRegistry.Contract.GetStakeUpdateIndexAtBlockNumber(&_StakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetStakeUpdateIndexAtBlockNumber is a free data retrieval call binding the contract method 0xdd9846b9.
//
// Solidity: function getStakeUpdateIndexAtBlockNumber(bytes32 operatorId, uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_StakeRegistry *StakeRegistryCallerSession) GetStakeUpdateIndexAtBlockNumber(operatorId [32]byte, quorumNumber uint8, blockNumber uint32) (uint32, error) {
	return _StakeRegistry.Contract.GetStakeUpdateIndexAtBlockNumber(&_StakeRegistry.CallOpts, operatorId, quorumNumber, blockNumber)
}

// GetTotalStakeAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc8294c56.
//
// Solidity: function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(uint96)
func (_StakeRegistry *StakeRegistryCaller) GetTotalStakeAtBlockNumberFromIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getTotalStakeAtBlockNumberFromIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc8294c56.
//
// Solidity: function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(uint96)
func (_StakeRegistry *StakeRegistrySession) GetTotalStakeAtBlockNumberFromIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _StakeRegistry.Contract.GetTotalStakeAtBlockNumberFromIndex(&_StakeRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetTotalStakeAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc8294c56.
//
// Solidity: function getTotalStakeAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(uint96)
func (_StakeRegistry *StakeRegistryCallerSession) GetTotalStakeAtBlockNumberFromIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _StakeRegistry.Contract.GetTotalStakeAtBlockNumberFromIndex(&_StakeRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetTotalStakeHistoryLength is a free data retrieval call binding the contract method 0x0491b41c.
//
// Solidity: function getTotalStakeHistoryLength(uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) GetTotalStakeHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getTotalStakeHistoryLength", quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeHistoryLength is a free data retrieval call binding the contract method 0x0491b41c.
//
// Solidity: function getTotalStakeHistoryLength(uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) GetTotalStakeHistoryLength(quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetTotalStakeHistoryLength(&_StakeRegistry.CallOpts, quorumNumber)
}

// GetTotalStakeHistoryLength is a free data retrieval call binding the contract method 0x0491b41c.
//
// Solidity: function getTotalStakeHistoryLength(uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) GetTotalStakeHistoryLength(quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.GetTotalStakeHistoryLength(&_StakeRegistry.CallOpts, quorumNumber)
}

// GetTotalStakeIndicesAtBlockNumber is a free data retrieval call binding the contract method 0x81c07502.
//
// Solidity: function getTotalStakeIndicesAtBlockNumber(uint32 blockNumber, bytes quorumNumbers) view returns(uint32[])
func (_StakeRegistry *StakeRegistryCaller) GetTotalStakeIndicesAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, quorumNumbers []byte) ([]uint32, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getTotalStakeIndicesAtBlockNumber", blockNumber, quorumNumbers)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetTotalStakeIndicesAtBlockNumber is a free data retrieval call binding the contract method 0x81c07502.
//
// Solidity: function getTotalStakeIndicesAtBlockNumber(uint32 blockNumber, bytes quorumNumbers) view returns(uint32[])
func (_StakeRegistry *StakeRegistrySession) GetTotalStakeIndicesAtBlockNumber(blockNumber uint32, quorumNumbers []byte) ([]uint32, error) {
	return _StakeRegistry.Contract.GetTotalStakeIndicesAtBlockNumber(&_StakeRegistry.CallOpts, blockNumber, quorumNumbers)
}

// GetTotalStakeIndicesAtBlockNumber is a free data retrieval call binding the contract method 0x81c07502.
//
// Solidity: function getTotalStakeIndicesAtBlockNumber(uint32 blockNumber, bytes quorumNumbers) view returns(uint32[])
func (_StakeRegistry *StakeRegistryCallerSession) GetTotalStakeIndicesAtBlockNumber(blockNumber uint32, quorumNumbers []byte) ([]uint32, error) {
	return _StakeRegistry.Contract.GetTotalStakeIndicesAtBlockNumber(&_StakeRegistry.CallOpts, blockNumber, quorumNumbers)
}

// GetTotalStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xb6904b78.
//
// Solidity: function getTotalStakeUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistryCaller) GetTotalStakeUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IStakeRegistryStakeUpdate, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "getTotalStakeUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IStakeRegistryStakeUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryStakeUpdate)).(*IStakeRegistryStakeUpdate)

	return out0, err

}

// GetTotalStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xb6904b78.
//
// Solidity: function getTotalStakeUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistrySession) GetTotalStakeUpdateAtIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetTotalStakeUpdateAtIndex(&_StakeRegistry.CallOpts, quorumNumber, index)
}

// GetTotalStakeUpdateAtIndex is a free data retrieval call binding the contract method 0xb6904b78.
//
// Solidity: function getTotalStakeUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((uint32,uint32,uint96))
func (_StakeRegistry *StakeRegistryCallerSession) GetTotalStakeUpdateAtIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryStakeUpdate, error) {
	return _StakeRegistry.Contract.GetTotalStakeUpdateAtIndex(&_StakeRegistry.CallOpts, quorumNumber, index)
}

// MinimumStakeForQuorum is a free data retrieval call binding the contract method 0xc46778a5.
//
// Solidity: function minimumStakeForQuorum(uint8 ) view returns(uint96)
func (_StakeRegistry *StakeRegistryCaller) MinimumStakeForQuorum(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "minimumStakeForQuorum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeForQuorum is a free data retrieval call binding the contract method 0xc46778a5.
//
// Solidity: function minimumStakeForQuorum(uint8 ) view returns(uint96)
func (_StakeRegistry *StakeRegistrySession) MinimumStakeForQuorum(arg0 uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.MinimumStakeForQuorum(&_StakeRegistry.CallOpts, arg0)
}

// MinimumStakeForQuorum is a free data retrieval call binding the contract method 0xc46778a5.
//
// Solidity: function minimumStakeForQuorum(uint8 ) view returns(uint96)
func (_StakeRegistry *StakeRegistryCallerSession) MinimumStakeForQuorum(arg0 uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.MinimumStakeForQuorum(&_StakeRegistry.CallOpts, arg0)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_StakeRegistry *StakeRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_StakeRegistry *StakeRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _StakeRegistry.Contract.RegistryCoordinator(&_StakeRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_StakeRegistry *StakeRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _StakeRegistry.Contract.RegistryCoordinator(&_StakeRegistry.CallOpts)
}

// StrategiesPerQuorum is a free data retrieval call binding the contract method 0x9f3ccf65.
//
// Solidity: function strategiesPerQuorum(uint8 , uint256 ) view returns(address)
func (_StakeRegistry *StakeRegistryCaller) StrategiesPerQuorum(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "strategiesPerQuorum", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategiesPerQuorum is a free data retrieval call binding the contract method 0x9f3ccf65.
//
// Solidity: function strategiesPerQuorum(uint8 , uint256 ) view returns(address)
func (_StakeRegistry *StakeRegistrySession) StrategiesPerQuorum(arg0 uint8, arg1 *big.Int) (common.Address, error) {
	return _StakeRegistry.Contract.StrategiesPerQuorum(&_StakeRegistry.CallOpts, arg0, arg1)
}

// StrategiesPerQuorum is a free data retrieval call binding the contract method 0x9f3ccf65.
//
// Solidity: function strategiesPerQuorum(uint8 , uint256 ) view returns(address)
func (_StakeRegistry *StakeRegistryCallerSession) StrategiesPerQuorum(arg0 uint8, arg1 *big.Int) (common.Address, error) {
	return _StakeRegistry.Contract.StrategiesPerQuorum(&_StakeRegistry.CallOpts, arg0, arg1)
}

// StrategyParams is a free data retrieval call binding the contract method 0x08732461.
//
// Solidity: function strategyParams(uint8 , uint256 ) view returns(address strategy, uint96 multiplier)
func (_StakeRegistry *StakeRegistryCaller) StrategyParams(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "strategyParams", arg0, arg1)

	outstruct := new(struct {
		Strategy   common.Address
		Multiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Strategy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Multiplier = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StrategyParams is a free data retrieval call binding the contract method 0x08732461.
//
// Solidity: function strategyParams(uint8 , uint256 ) view returns(address strategy, uint96 multiplier)
func (_StakeRegistry *StakeRegistrySession) StrategyParams(arg0 uint8, arg1 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	return _StakeRegistry.Contract.StrategyParams(&_StakeRegistry.CallOpts, arg0, arg1)
}

// StrategyParams is a free data retrieval call binding the contract method 0x08732461.
//
// Solidity: function strategyParams(uint8 , uint256 ) view returns(address strategy, uint96 multiplier)
func (_StakeRegistry *StakeRegistryCallerSession) StrategyParams(arg0 uint8, arg1 *big.Int) (struct {
	Strategy   common.Address
	Multiplier *big.Int
}, error) {
	return _StakeRegistry.Contract.StrategyParams(&_StakeRegistry.CallOpts, arg0, arg1)
}

// StrategyParamsByIndex is a free data retrieval call binding the contract method 0xadc804da.
//
// Solidity: function strategyParamsByIndex(uint8 quorumNumber, uint256 index) view returns((address,uint96))
func (_StakeRegistry *StakeRegistryCaller) StrategyParamsByIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IStakeRegistryStrategyParams, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "strategyParamsByIndex", quorumNumber, index)

	if err != nil {
		return *new(IStakeRegistryStrategyParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeRegistryStrategyParams)).(*IStakeRegistryStrategyParams)

	return out0, err

}

// StrategyParamsByIndex is a free data retrieval call binding the contract method 0xadc804da.
//
// Solidity: function strategyParamsByIndex(uint8 quorumNumber, uint256 index) view returns((address,uint96))
func (_StakeRegistry *StakeRegistrySession) StrategyParamsByIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryStrategyParams, error) {
	return _StakeRegistry.Contract.StrategyParamsByIndex(&_StakeRegistry.CallOpts, quorumNumber, index)
}

// StrategyParamsByIndex is a free data retrieval call binding the contract method 0xadc804da.
//
// Solidity: function strategyParamsByIndex(uint8 quorumNumber, uint256 index) view returns((address,uint96))
func (_StakeRegistry *StakeRegistryCallerSession) StrategyParamsByIndex(quorumNumber uint8, index *big.Int) (IStakeRegistryStrategyParams, error) {
	return _StakeRegistry.Contract.StrategyParamsByIndex(&_StakeRegistry.CallOpts, quorumNumber, index)
}

// StrategyParamsLength is a free data retrieval call binding the contract method 0x3ca5a5f5.
//
// Solidity: function strategyParamsLength(uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCaller) StrategyParamsLength(opts *bind.CallOpts, quorumNumber uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "strategyParamsLength", quorumNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyParamsLength is a free data retrieval call binding the contract method 0x3ca5a5f5.
//
// Solidity: function strategyParamsLength(uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistrySession) StrategyParamsLength(quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.StrategyParamsLength(&_StakeRegistry.CallOpts, quorumNumber)
}

// StrategyParamsLength is a free data retrieval call binding the contract method 0x3ca5a5f5.
//
// Solidity: function strategyParamsLength(uint8 quorumNumber) view returns(uint256)
func (_StakeRegistry *StakeRegistryCallerSession) StrategyParamsLength(quorumNumber uint8) (*big.Int, error) {
	return _StakeRegistry.Contract.StrategyParamsLength(&_StakeRegistry.CallOpts, quorumNumber)
}

// WeightOfOperatorForQuorum is a free data retrieval call binding the contract method 0x1f9b74e0.
//
// Solidity: function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) view returns(uint96)
func (_StakeRegistry *StakeRegistryCaller) WeightOfOperatorForQuorum(opts *bind.CallOpts, quorumNumber uint8, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeRegistry.contract.Call(opts, &out, "weightOfOperatorForQuorum", quorumNumber, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightOfOperatorForQuorum is a free data retrieval call binding the contract method 0x1f9b74e0.
//
// Solidity: function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) view returns(uint96)
func (_StakeRegistry *StakeRegistrySession) WeightOfOperatorForQuorum(quorumNumber uint8, operator common.Address) (*big.Int, error) {
	return _StakeRegistry.Contract.WeightOfOperatorForQuorum(&_StakeRegistry.CallOpts, quorumNumber, operator)
}

// WeightOfOperatorForQuorum is a free data retrieval call binding the contract method 0x1f9b74e0.
//
// Solidity: function weightOfOperatorForQuorum(uint8 quorumNumber, address operator) view returns(uint96)
func (_StakeRegistry *StakeRegistryCallerSession) WeightOfOperatorForQuorum(quorumNumber uint8, operator common.Address) (*big.Int, error) {
	return _StakeRegistry.Contract.WeightOfOperatorForQuorum(&_StakeRegistry.CallOpts, quorumNumber, operator)
}

// AddStrategies is a paid mutator transaction binding the contract method 0xc601527d.
//
// Solidity: function addStrategies(uint8 quorumNumber, (address,uint96)[] _strategyParams) returns()
func (_StakeRegistry *StakeRegistryTransactor) AddStrategies(opts *bind.TransactOpts, quorumNumber uint8, _strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "addStrategies", quorumNumber, _strategyParams)
}

// AddStrategies is a paid mutator transaction binding the contract method 0xc601527d.
//
// Solidity: function addStrategies(uint8 quorumNumber, (address,uint96)[] _strategyParams) returns()
func (_StakeRegistry *StakeRegistrySession) AddStrategies(quorumNumber uint8, _strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _StakeRegistry.Contract.AddStrategies(&_StakeRegistry.TransactOpts, quorumNumber, _strategyParams)
}

// AddStrategies is a paid mutator transaction binding the contract method 0xc601527d.
//
// Solidity: function addStrategies(uint8 quorumNumber, (address,uint96)[] _strategyParams) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) AddStrategies(quorumNumber uint8, _strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _StakeRegistry.Contract.AddStrategies(&_StakeRegistry.TransactOpts, quorumNumber, _strategyParams)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_StakeRegistry *StakeRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "deregisterOperator", operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_StakeRegistry *StakeRegistrySession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.DeregisterOperator(&_StakeRegistry.TransactOpts, operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.DeregisterOperator(&_StakeRegistry.TransactOpts, operatorId, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0xff694a77.
//
// Solidity: function initializeQuorum(uint8 quorumNumber, uint96 minimumStake, (address,uint96)[] _strategyParams) returns()
func (_StakeRegistry *StakeRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8, minimumStake *big.Int, _strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber, minimumStake, _strategyParams)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0xff694a77.
//
// Solidity: function initializeQuorum(uint8 quorumNumber, uint96 minimumStake, (address,uint96)[] _strategyParams) returns()
func (_StakeRegistry *StakeRegistrySession) InitializeQuorum(quorumNumber uint8, minimumStake *big.Int, _strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _StakeRegistry.Contract.InitializeQuorum(&_StakeRegistry.TransactOpts, quorumNumber, minimumStake, _strategyParams)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0xff694a77.
//
// Solidity: function initializeQuorum(uint8 quorumNumber, uint96 minimumStake, (address,uint96)[] _strategyParams) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) InitializeQuorum(quorumNumber uint8, minimumStake *big.Int, _strategyParams []IStakeRegistryStrategyParams) (*types.Transaction, error) {
	return _StakeRegistry.Contract.InitializeQuorum(&_StakeRegistry.TransactOpts, quorumNumber, minimumStake, _strategyParams)
}

// ModifyStrategyParams is a paid mutator transaction binding the contract method 0x20b66298.
//
// Solidity: function modifyStrategyParams(uint8 quorumNumber, uint256[] strategyIndices, uint96[] newMultipliers) returns()
func (_StakeRegistry *StakeRegistryTransactor) ModifyStrategyParams(opts *bind.TransactOpts, quorumNumber uint8, strategyIndices []*big.Int, newMultipliers []*big.Int) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "modifyStrategyParams", quorumNumber, strategyIndices, newMultipliers)
}

// ModifyStrategyParams is a paid mutator transaction binding the contract method 0x20b66298.
//
// Solidity: function modifyStrategyParams(uint8 quorumNumber, uint256[] strategyIndices, uint96[] newMultipliers) returns()
func (_StakeRegistry *StakeRegistrySession) ModifyStrategyParams(quorumNumber uint8, strategyIndices []*big.Int, newMultipliers []*big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.ModifyStrategyParams(&_StakeRegistry.TransactOpts, quorumNumber, strategyIndices, newMultipliers)
}

// ModifyStrategyParams is a paid mutator transaction binding the contract method 0x20b66298.
//
// Solidity: function modifyStrategyParams(uint8 quorumNumber, uint256[] strategyIndices, uint96[] newMultipliers) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) ModifyStrategyParams(quorumNumber uint8, strategyIndices []*big.Int, newMultipliers []*big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.ModifyStrategyParams(&_StakeRegistry.TransactOpts, quorumNumber, strategyIndices, newMultipliers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x25504777.
//
// Solidity: function registerOperator(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint96[], uint96[])
func (_StakeRegistry *StakeRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "registerOperator", operator, operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x25504777.
//
// Solidity: function registerOperator(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint96[], uint96[])
func (_StakeRegistry *StakeRegistrySession) RegisterOperator(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.RegisterOperator(&_StakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x25504777.
//
// Solidity: function registerOperator(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint96[], uint96[])
func (_StakeRegistry *StakeRegistryTransactorSession) RegisterOperator(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.RegisterOperator(&_StakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// RemoveStrategies is a paid mutator transaction binding the contract method 0x5f1f2d77.
//
// Solidity: function removeStrategies(uint8 quorumNumber, uint256[] indicesToRemove) returns()
func (_StakeRegistry *StakeRegistryTransactor) RemoveStrategies(opts *bind.TransactOpts, quorumNumber uint8, indicesToRemove []*big.Int) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "removeStrategies", quorumNumber, indicesToRemove)
}

// RemoveStrategies is a paid mutator transaction binding the contract method 0x5f1f2d77.
//
// Solidity: function removeStrategies(uint8 quorumNumber, uint256[] indicesToRemove) returns()
func (_StakeRegistry *StakeRegistrySession) RemoveStrategies(quorumNumber uint8, indicesToRemove []*big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.RemoveStrategies(&_StakeRegistry.TransactOpts, quorumNumber, indicesToRemove)
}

// RemoveStrategies is a paid mutator transaction binding the contract method 0x5f1f2d77.
//
// Solidity: function removeStrategies(uint8 quorumNumber, uint256[] indicesToRemove) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) RemoveStrategies(quorumNumber uint8, indicesToRemove []*big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.RemoveStrategies(&_StakeRegistry.TransactOpts, quorumNumber, indicesToRemove)
}

// SetMinimumStakeForQuorum is a paid mutator transaction binding the contract method 0xbc9a40c3.
//
// Solidity: function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) returns()
func (_StakeRegistry *StakeRegistryTransactor) SetMinimumStakeForQuorum(opts *bind.TransactOpts, quorumNumber uint8, minimumStake *big.Int) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "setMinimumStakeForQuorum", quorumNumber, minimumStake)
}

// SetMinimumStakeForQuorum is a paid mutator transaction binding the contract method 0xbc9a40c3.
//
// Solidity: function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) returns()
func (_StakeRegistry *StakeRegistrySession) SetMinimumStakeForQuorum(quorumNumber uint8, minimumStake *big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.SetMinimumStakeForQuorum(&_StakeRegistry.TransactOpts, quorumNumber, minimumStake)
}

// SetMinimumStakeForQuorum is a paid mutator transaction binding the contract method 0xbc9a40c3.
//
// Solidity: function setMinimumStakeForQuorum(uint8 quorumNumber, uint96 minimumStake) returns()
func (_StakeRegistry *StakeRegistryTransactorSession) SetMinimumStakeForQuorum(quorumNumber uint8, minimumStake *big.Int) (*types.Transaction, error) {
	return _StakeRegistry.Contract.SetMinimumStakeForQuorum(&_StakeRegistry.TransactOpts, quorumNumber, minimumStake)
}

// UpdateOperatorStake is a paid mutator transaction binding the contract method 0x66acfefe.
//
// Solidity: function updateOperatorStake(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint192)
func (_StakeRegistry *StakeRegistryTransactor) UpdateOperatorStake(opts *bind.TransactOpts, operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.contract.Transact(opts, "updateOperatorStake", operator, operatorId, quorumNumbers)
}

// UpdateOperatorStake is a paid mutator transaction binding the contract method 0x66acfefe.
//
// Solidity: function updateOperatorStake(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint192)
func (_StakeRegistry *StakeRegistrySession) UpdateOperatorStake(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorStake(&_StakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// UpdateOperatorStake is a paid mutator transaction binding the contract method 0x66acfefe.
//
// Solidity: function updateOperatorStake(address operator, bytes32 operatorId, bytes quorumNumbers) returns(uint192)
func (_StakeRegistry *StakeRegistryTransactorSession) UpdateOperatorStake(operator common.Address, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _StakeRegistry.Contract.UpdateOperatorStake(&_StakeRegistry.TransactOpts, operator, operatorId, quorumNumbers)
}

// StakeRegistryMinimumStakeForQuorumUpdatedIterator is returned from FilterMinimumStakeForQuorumUpdated and is used to iterate over the raw logs and unpacked data for MinimumStakeForQuorumUpdated events raised by the StakeRegistry contract.
type StakeRegistryMinimumStakeForQuorumUpdatedIterator struct {
	Event *StakeRegistryMinimumStakeForQuorumUpdated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryMinimumStakeForQuorumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryMinimumStakeForQuorumUpdated)
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
		it.Event = new(StakeRegistryMinimumStakeForQuorumUpdated)
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
func (it *StakeRegistryMinimumStakeForQuorumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryMinimumStakeForQuorumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryMinimumStakeForQuorumUpdated represents a MinimumStakeForQuorumUpdated event raised by the StakeRegistry contract.
type StakeRegistryMinimumStakeForQuorumUpdated struct {
	QuorumNumber uint8
	MinimumStake *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinimumStakeForQuorumUpdated is a free log retrieval operation binding the contract event 0x26eecff2b70b0a71104ff4d940ba7162d23a95c248771fc487a7be17a596b3cf.
//
// Solidity: event MinimumStakeForQuorumUpdated(uint8 indexed quorumNumber, uint96 minimumStake)
func (_StakeRegistry *StakeRegistryFilterer) FilterMinimumStakeForQuorumUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*StakeRegistryMinimumStakeForQuorumUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "MinimumStakeForQuorumUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryMinimumStakeForQuorumUpdatedIterator{contract: _StakeRegistry.contract, event: "MinimumStakeForQuorumUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumStakeForQuorumUpdated is a free log subscription operation binding the contract event 0x26eecff2b70b0a71104ff4d940ba7162d23a95c248771fc487a7be17a596b3cf.
//
// Solidity: event MinimumStakeForQuorumUpdated(uint8 indexed quorumNumber, uint96 minimumStake)
func (_StakeRegistry *StakeRegistryFilterer) WatchMinimumStakeForQuorumUpdated(opts *bind.WatchOpts, sink chan<- *StakeRegistryMinimumStakeForQuorumUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "MinimumStakeForQuorumUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryMinimumStakeForQuorumUpdated)
				if err := _StakeRegistry.contract.UnpackLog(event, "MinimumStakeForQuorumUpdated", log); err != nil {
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

// ParseMinimumStakeForQuorumUpdated is a log parse operation binding the contract event 0x26eecff2b70b0a71104ff4d940ba7162d23a95c248771fc487a7be17a596b3cf.
//
// Solidity: event MinimumStakeForQuorumUpdated(uint8 indexed quorumNumber, uint96 minimumStake)
func (_StakeRegistry *StakeRegistryFilterer) ParseMinimumStakeForQuorumUpdated(log types.Log) (*StakeRegistryMinimumStakeForQuorumUpdated, error) {
	event := new(StakeRegistryMinimumStakeForQuorumUpdated)
	if err := _StakeRegistry.contract.UnpackLog(event, "MinimumStakeForQuorumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryOperatorStakeUpdateIterator is returned from FilterOperatorStakeUpdate and is used to iterate over the raw logs and unpacked data for OperatorStakeUpdate events raised by the StakeRegistry contract.
type StakeRegistryOperatorStakeUpdateIterator struct {
	Event *StakeRegistryOperatorStakeUpdate // Event containing the contract specifics and raw log

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
func (it *StakeRegistryOperatorStakeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryOperatorStakeUpdate)
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
		it.Event = new(StakeRegistryOperatorStakeUpdate)
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
func (it *StakeRegistryOperatorStakeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryOperatorStakeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryOperatorStakeUpdate represents a OperatorStakeUpdate event raised by the StakeRegistry contract.
type StakeRegistryOperatorStakeUpdate struct {
	OperatorId   [32]byte
	QuorumNumber uint8
	Stake        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOperatorStakeUpdate is a free log retrieval operation binding the contract event 0x2f527d527e95d8fe40aec55377743bb779087da3f6d0d08f12e36444da62327d.
//
// Solidity: event OperatorStakeUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint96 stake)
func (_StakeRegistry *StakeRegistryFilterer) FilterOperatorStakeUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*StakeRegistryOperatorStakeUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "OperatorStakeUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryOperatorStakeUpdateIterator{contract: _StakeRegistry.contract, event: "OperatorStakeUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorStakeUpdate is a free log subscription operation binding the contract event 0x2f527d527e95d8fe40aec55377743bb779087da3f6d0d08f12e36444da62327d.
//
// Solidity: event OperatorStakeUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint96 stake)
func (_StakeRegistry *StakeRegistryFilterer) WatchOperatorStakeUpdate(opts *bind.WatchOpts, sink chan<- *StakeRegistryOperatorStakeUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "OperatorStakeUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryOperatorStakeUpdate)
				if err := _StakeRegistry.contract.UnpackLog(event, "OperatorStakeUpdate", log); err != nil {
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

// ParseOperatorStakeUpdate is a log parse operation binding the contract event 0x2f527d527e95d8fe40aec55377743bb779087da3f6d0d08f12e36444da62327d.
//
// Solidity: event OperatorStakeUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint96 stake)
func (_StakeRegistry *StakeRegistryFilterer) ParseOperatorStakeUpdate(log types.Log) (*StakeRegistryOperatorStakeUpdate, error) {
	event := new(StakeRegistryOperatorStakeUpdate)
	if err := _StakeRegistry.contract.UnpackLog(event, "OperatorStakeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryQuorumCreatedIterator is returned from FilterQuorumCreated and is used to iterate over the raw logs and unpacked data for QuorumCreated events raised by the StakeRegistry contract.
type StakeRegistryQuorumCreatedIterator struct {
	Event *StakeRegistryQuorumCreated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryQuorumCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryQuorumCreated)
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
		it.Event = new(StakeRegistryQuorumCreated)
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
func (it *StakeRegistryQuorumCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryQuorumCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryQuorumCreated represents a QuorumCreated event raised by the StakeRegistry contract.
type StakeRegistryQuorumCreated struct {
	QuorumNumber uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQuorumCreated is a free log retrieval operation binding the contract event 0x831a9c86c45bb303caf3f064be2bc2b9fd4ecf19e47c4ac02a61e75dabfe55b4.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber)
func (_StakeRegistry *StakeRegistryFilterer) FilterQuorumCreated(opts *bind.FilterOpts, quorumNumber []uint8) (*StakeRegistryQuorumCreatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "QuorumCreated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryQuorumCreatedIterator{contract: _StakeRegistry.contract, event: "QuorumCreated", logs: logs, sub: sub}, nil
}

// WatchQuorumCreated is a free log subscription operation binding the contract event 0x831a9c86c45bb303caf3f064be2bc2b9fd4ecf19e47c4ac02a61e75dabfe55b4.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber)
func (_StakeRegistry *StakeRegistryFilterer) WatchQuorumCreated(opts *bind.WatchOpts, sink chan<- *StakeRegistryQuorumCreated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "QuorumCreated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryQuorumCreated)
				if err := _StakeRegistry.contract.UnpackLog(event, "QuorumCreated", log); err != nil {
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

// ParseQuorumCreated is a log parse operation binding the contract event 0x831a9c86c45bb303caf3f064be2bc2b9fd4ecf19e47c4ac02a61e75dabfe55b4.
//
// Solidity: event QuorumCreated(uint8 indexed quorumNumber)
func (_StakeRegistry *StakeRegistryFilterer) ParseQuorumCreated(log types.Log) (*StakeRegistryQuorumCreated, error) {
	event := new(StakeRegistryQuorumCreated)
	if err := _StakeRegistry.contract.UnpackLog(event, "QuorumCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryStrategyAddedToQuorumIterator is returned from FilterStrategyAddedToQuorum and is used to iterate over the raw logs and unpacked data for StrategyAddedToQuorum events raised by the StakeRegistry contract.
type StakeRegistryStrategyAddedToQuorumIterator struct {
	Event *StakeRegistryStrategyAddedToQuorum // Event containing the contract specifics and raw log

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
func (it *StakeRegistryStrategyAddedToQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryStrategyAddedToQuorum)
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
		it.Event = new(StakeRegistryStrategyAddedToQuorum)
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
func (it *StakeRegistryStrategyAddedToQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryStrategyAddedToQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryStrategyAddedToQuorum represents a StrategyAddedToQuorum event raised by the StakeRegistry contract.
type StakeRegistryStrategyAddedToQuorum struct {
	QuorumNumber uint8
	Strategy     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyAddedToQuorum is a free log retrieval operation binding the contract event 0x10565e56cacbf32eca267945f054fec02e59750032d113d3302182ad967f5404.
//
// Solidity: event StrategyAddedToQuorum(uint8 indexed quorumNumber, address strategy)
func (_StakeRegistry *StakeRegistryFilterer) FilterStrategyAddedToQuorum(opts *bind.FilterOpts, quorumNumber []uint8) (*StakeRegistryStrategyAddedToQuorumIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "StrategyAddedToQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryStrategyAddedToQuorumIterator{contract: _StakeRegistry.contract, event: "StrategyAddedToQuorum", logs: logs, sub: sub}, nil
}

// WatchStrategyAddedToQuorum is a free log subscription operation binding the contract event 0x10565e56cacbf32eca267945f054fec02e59750032d113d3302182ad967f5404.
//
// Solidity: event StrategyAddedToQuorum(uint8 indexed quorumNumber, address strategy)
func (_StakeRegistry *StakeRegistryFilterer) WatchStrategyAddedToQuorum(opts *bind.WatchOpts, sink chan<- *StakeRegistryStrategyAddedToQuorum, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "StrategyAddedToQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryStrategyAddedToQuorum)
				if err := _StakeRegistry.contract.UnpackLog(event, "StrategyAddedToQuorum", log); err != nil {
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

// ParseStrategyAddedToQuorum is a log parse operation binding the contract event 0x10565e56cacbf32eca267945f054fec02e59750032d113d3302182ad967f5404.
//
// Solidity: event StrategyAddedToQuorum(uint8 indexed quorumNumber, address strategy)
func (_StakeRegistry *StakeRegistryFilterer) ParseStrategyAddedToQuorum(log types.Log) (*StakeRegistryStrategyAddedToQuorum, error) {
	event := new(StakeRegistryStrategyAddedToQuorum)
	if err := _StakeRegistry.contract.UnpackLog(event, "StrategyAddedToQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryStrategyMultiplierUpdatedIterator is returned from FilterStrategyMultiplierUpdated and is used to iterate over the raw logs and unpacked data for StrategyMultiplierUpdated events raised by the StakeRegistry contract.
type StakeRegistryStrategyMultiplierUpdatedIterator struct {
	Event *StakeRegistryStrategyMultiplierUpdated // Event containing the contract specifics and raw log

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
func (it *StakeRegistryStrategyMultiplierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryStrategyMultiplierUpdated)
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
		it.Event = new(StakeRegistryStrategyMultiplierUpdated)
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
func (it *StakeRegistryStrategyMultiplierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryStrategyMultiplierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryStrategyMultiplierUpdated represents a StrategyMultiplierUpdated event raised by the StakeRegistry contract.
type StakeRegistryStrategyMultiplierUpdated struct {
	QuorumNumber uint8
	Strategy     common.Address
	Multiplier   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyMultiplierUpdated is a free log retrieval operation binding the contract event 0x11a5641322da1dff56a4b66eaac31ffa465295ece907cd163437793b4d009a75.
//
// Solidity: event StrategyMultiplierUpdated(uint8 indexed quorumNumber, address strategy, uint256 multiplier)
func (_StakeRegistry *StakeRegistryFilterer) FilterStrategyMultiplierUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*StakeRegistryStrategyMultiplierUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "StrategyMultiplierUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryStrategyMultiplierUpdatedIterator{contract: _StakeRegistry.contract, event: "StrategyMultiplierUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategyMultiplierUpdated is a free log subscription operation binding the contract event 0x11a5641322da1dff56a4b66eaac31ffa465295ece907cd163437793b4d009a75.
//
// Solidity: event StrategyMultiplierUpdated(uint8 indexed quorumNumber, address strategy, uint256 multiplier)
func (_StakeRegistry *StakeRegistryFilterer) WatchStrategyMultiplierUpdated(opts *bind.WatchOpts, sink chan<- *StakeRegistryStrategyMultiplierUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "StrategyMultiplierUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryStrategyMultiplierUpdated)
				if err := _StakeRegistry.contract.UnpackLog(event, "StrategyMultiplierUpdated", log); err != nil {
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

// ParseStrategyMultiplierUpdated is a log parse operation binding the contract event 0x11a5641322da1dff56a4b66eaac31ffa465295ece907cd163437793b4d009a75.
//
// Solidity: event StrategyMultiplierUpdated(uint8 indexed quorumNumber, address strategy, uint256 multiplier)
func (_StakeRegistry *StakeRegistryFilterer) ParseStrategyMultiplierUpdated(log types.Log) (*StakeRegistryStrategyMultiplierUpdated, error) {
	event := new(StakeRegistryStrategyMultiplierUpdated)
	if err := _StakeRegistry.contract.UnpackLog(event, "StrategyMultiplierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRegistryStrategyRemovedFromQuorumIterator is returned from FilterStrategyRemovedFromQuorum and is used to iterate over the raw logs and unpacked data for StrategyRemovedFromQuorum events raised by the StakeRegistry contract.
type StakeRegistryStrategyRemovedFromQuorumIterator struct {
	Event *StakeRegistryStrategyRemovedFromQuorum // Event containing the contract specifics and raw log

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
func (it *StakeRegistryStrategyRemovedFromQuorumIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRegistryStrategyRemovedFromQuorum)
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
		it.Event = new(StakeRegistryStrategyRemovedFromQuorum)
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
func (it *StakeRegistryStrategyRemovedFromQuorumIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRegistryStrategyRemovedFromQuorumIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRegistryStrategyRemovedFromQuorum represents a StrategyRemovedFromQuorum event raised by the StakeRegistry contract.
type StakeRegistryStrategyRemovedFromQuorum struct {
	QuorumNumber uint8
	Strategy     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemovedFromQuorum is a free log retrieval operation binding the contract event 0x31fa2e2cd280c9375e13ffcf3d81e2378100186e4058f8d3ddb690b82dcd31f7.
//
// Solidity: event StrategyRemovedFromQuorum(uint8 indexed quorumNumber, address strategy)
func (_StakeRegistry *StakeRegistryFilterer) FilterStrategyRemovedFromQuorum(opts *bind.FilterOpts, quorumNumber []uint8) (*StakeRegistryStrategyRemovedFromQuorumIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.FilterLogs(opts, "StrategyRemovedFromQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &StakeRegistryStrategyRemovedFromQuorumIterator{contract: _StakeRegistry.contract, event: "StrategyRemovedFromQuorum", logs: logs, sub: sub}, nil
}

// WatchStrategyRemovedFromQuorum is a free log subscription operation binding the contract event 0x31fa2e2cd280c9375e13ffcf3d81e2378100186e4058f8d3ddb690b82dcd31f7.
//
// Solidity: event StrategyRemovedFromQuorum(uint8 indexed quorumNumber, address strategy)
func (_StakeRegistry *StakeRegistryFilterer) WatchStrategyRemovedFromQuorum(opts *bind.WatchOpts, sink chan<- *StakeRegistryStrategyRemovedFromQuorum, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _StakeRegistry.contract.WatchLogs(opts, "StrategyRemovedFromQuorum", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRegistryStrategyRemovedFromQuorum)
				if err := _StakeRegistry.contract.UnpackLog(event, "StrategyRemovedFromQuorum", log); err != nil {
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

// ParseStrategyRemovedFromQuorum is a log parse operation binding the contract event 0x31fa2e2cd280c9375e13ffcf3d81e2378100186e4058f8d3ddb690b82dcd31f7.
//
// Solidity: event StrategyRemovedFromQuorum(uint8 indexed quorumNumber, address strategy)
func (_StakeRegistry *StakeRegistryFilterer) ParseStrategyRemovedFromQuorum(log types.Log) (*StakeRegistryStrategyRemovedFromQuorum, error) {
	event := new(StakeRegistryStrategyRemovedFromQuorum)
	if err := _StakeRegistry.contract.UnpackLog(event, "StrategyRemovedFromQuorum", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
