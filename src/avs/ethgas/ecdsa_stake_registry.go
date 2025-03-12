// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethgas

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

// Quorum is an auto generated low-level Go binding around an user-defined struct.
type Quorum struct {
	Strategies []StrategyParams
}

// StrategyParams is an auto generated low-level Go binding around an user-defined struct.
type StrategyParams struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ECDSAStakeRegistryMetaData contains all meta data concerning the ECDSAStakeRegistry contract.
var ECDSAStakeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegationManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientSignedStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQuorum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReferenceBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignedWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustUpdateAllOperators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_old\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_new\",\"type\":\"uint256\"}],\"name\":\"MinimumWeightUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_avs\",\"type\":\"address\"}],\"name\":\"OperatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_avs\",\"type\":\"address\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWeight\",\"type\":\"uint256\"}],\"name\":\"OperatorWeightUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structQuorum\",\"name\":\"_old\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structQuorum\",\"name\":\"_new\",\"type\":\"tuple\"}],\"name\":\"QuorumUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"updateBlock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSigningKey\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigningKey\",\"type\":\"address\"}],\"name\":\"SigningKeyUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_thresholdWeight\",\"type\":\"uint256\"}],\"name\":\"ThresholdWeightUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTotalWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalWeight\",\"type\":\"uint256\"}],\"name\":\"TotalWeightUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMinimumWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinimumWeight\",\"type\":\"uint256\"}],\"name\":\"UpdateMinimumWeight\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getLastCheckpointOperatorWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastCheckpointThresholdWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_blockNumber\",\"type\":\"uint32\"}],\"name\":\"getLastCheckpointThresholdWeightAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastCheckpointTotalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_blockNumber\",\"type\":\"uint32\"}],\"name\":\"getLastCheckpointTotalWeightAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getLastestOperatorSigningKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getOperatorSigningKeyAtBlock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorWeightAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_serviceManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_thresholdWeight\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"internalType\":\"structQuorum\",\"name\":\"_quorum\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signatureData\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"operatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"internalType\":\"structQuorum\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_signingKey\",\"type\":\"address\"}],\"name\":\"registerOperatorWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMinimumWeight\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"updateMinimumWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSigningKey\",\"type\":\"address\"}],\"name\":\"updateOperatorSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"updateOperators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[][]\",\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"updateOperatorsForQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"internalType\":\"structQuorum\",\"name\":\"_quorum\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"}],\"name\":\"updateQuorumConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_thresholdWeight\",\"type\":\"uint256\"}],\"name\":\"updateStakeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ECDSAStakeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAStakeRegistryMetaData.ABI instead.
var ECDSAStakeRegistryABI = ECDSAStakeRegistryMetaData.ABI

// ECDSAStakeRegistry is an auto generated Go binding around an Ethereum contract.
type ECDSAStakeRegistry struct {
	ECDSAStakeRegistryCaller     // Read-only binding to the contract
	ECDSAStakeRegistryTransactor // Write-only binding to the contract
	ECDSAStakeRegistryFilterer   // Log filterer for contract events
}

// ECDSAStakeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSAStakeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAStakeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSAStakeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAStakeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAStakeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAStakeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSAStakeRegistrySession struct {
	Contract     *ECDSAStakeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ECDSAStakeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSAStakeRegistryCallerSession struct {
	Contract *ECDSAStakeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ECDSAStakeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSAStakeRegistryTransactorSession struct {
	Contract     *ECDSAStakeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ECDSAStakeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSAStakeRegistryRaw struct {
	Contract *ECDSAStakeRegistry // Generic contract binding to access the raw methods on
}

// ECDSAStakeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSAStakeRegistryCallerRaw struct {
	Contract *ECDSAStakeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ECDSAStakeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSAStakeRegistryTransactorRaw struct {
	Contract *ECDSAStakeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSAStakeRegistry creates a new instance of ECDSAStakeRegistry, bound to a specific deployed contract.
func NewECDSAStakeRegistry(address common.Address, backend bind.ContractBackend) (*ECDSAStakeRegistry, error) {
	contract, err := bindECDSAStakeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistry{ECDSAStakeRegistryCaller: ECDSAStakeRegistryCaller{contract: contract}, ECDSAStakeRegistryTransactor: ECDSAStakeRegistryTransactor{contract: contract}, ECDSAStakeRegistryFilterer: ECDSAStakeRegistryFilterer{contract: contract}}, nil
}

// NewECDSAStakeRegistryCaller creates a new read-only instance of ECDSAStakeRegistry, bound to a specific deployed contract.
func NewECDSAStakeRegistryCaller(address common.Address, caller bind.ContractCaller) (*ECDSAStakeRegistryCaller, error) {
	contract, err := bindECDSAStakeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryCaller{contract: contract}, nil
}

// NewECDSAStakeRegistryTransactor creates a new write-only instance of ECDSAStakeRegistry, bound to a specific deployed contract.
func NewECDSAStakeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSAStakeRegistryTransactor, error) {
	contract, err := bindECDSAStakeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryTransactor{contract: contract}, nil
}

// NewECDSAStakeRegistryFilterer creates a new log filterer instance of ECDSAStakeRegistry, bound to a specific deployed contract.
func NewECDSAStakeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAStakeRegistryFilterer, error) {
	contract, err := bindECDSAStakeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryFilterer{contract: contract}, nil
}

// bindECDSAStakeRegistry binds a generic wrapper to an already deployed contract.
func bindECDSAStakeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSAStakeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAStakeRegistry *ECDSAStakeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAStakeRegistry.Contract.ECDSAStakeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAStakeRegistry *ECDSAStakeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.ECDSAStakeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAStakeRegistry *ECDSAStakeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.ECDSAStakeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSAStakeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetLastCheckpointOperatorWeight(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getLastCheckpointOperatorWeight", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetLastCheckpointOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointOperatorWeight(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// GetLastCheckpointOperatorWeight is a free data retrieval call binding the contract method 0x3b242e4a.
//
// Solidity: function getLastCheckpointOperatorWeight(address _operator) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetLastCheckpointOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointOperatorWeight(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetLastCheckpointThresholdWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getLastCheckpointThresholdWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetLastCheckpointThresholdWeight() (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointThresholdWeight(&_ECDSAStakeRegistry.CallOpts)
}

// GetLastCheckpointThresholdWeight is a free data retrieval call binding the contract method 0xb933fa74.
//
// Solidity: function getLastCheckpointThresholdWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetLastCheckpointThresholdWeight() (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointThresholdWeight(&_ECDSAStakeRegistry.CallOpts)
}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetLastCheckpointThresholdWeightAtBlock(opts *bind.CallOpts, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getLastCheckpointThresholdWeightAtBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetLastCheckpointThresholdWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointThresholdWeightAtBlock(&_ECDSAStakeRegistry.CallOpts, _blockNumber)
}

// GetLastCheckpointThresholdWeightAtBlock is a free data retrieval call binding the contract method 0x1e4cd85e.
//
// Solidity: function getLastCheckpointThresholdWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetLastCheckpointThresholdWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointThresholdWeightAtBlock(&_ECDSAStakeRegistry.CallOpts, _blockNumber)
}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetLastCheckpointTotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getLastCheckpointTotalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetLastCheckpointTotalWeight() (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointTotalWeight(&_ECDSAStakeRegistry.CallOpts)
}

// GetLastCheckpointTotalWeight is a free data retrieval call binding the contract method 0x314f3a49.
//
// Solidity: function getLastCheckpointTotalWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetLastCheckpointTotalWeight() (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointTotalWeight(&_ECDSAStakeRegistry.CallOpts)
}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetLastCheckpointTotalWeightAtBlock(opts *bind.CallOpts, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getLastCheckpointTotalWeightAtBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetLastCheckpointTotalWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointTotalWeightAtBlock(&_ECDSAStakeRegistry.CallOpts, _blockNumber)
}

// GetLastCheckpointTotalWeightAtBlock is a free data retrieval call binding the contract method 0x0dba3394.
//
// Solidity: function getLastCheckpointTotalWeightAtBlock(uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetLastCheckpointTotalWeightAtBlock(_blockNumber uint32) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetLastCheckpointTotalWeightAtBlock(&_ECDSAStakeRegistry.CallOpts, _blockNumber)
}

// GetLastestOperatorSigningKey is a free data retrieval call binding the contract method 0xcdcd3581.
//
// Solidity: function getLastestOperatorSigningKey(address _operator) view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetLastestOperatorSigningKey(opts *bind.CallOpts, _operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getLastestOperatorSigningKey", _operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLastestOperatorSigningKey is a free data retrieval call binding the contract method 0xcdcd3581.
//
// Solidity: function getLastestOperatorSigningKey(address _operator) view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetLastestOperatorSigningKey(_operator common.Address) (common.Address, error) {
	return _ECDSAStakeRegistry.Contract.GetLastestOperatorSigningKey(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// GetLastestOperatorSigningKey is a free data retrieval call binding the contract method 0xcdcd3581.
//
// Solidity: function getLastestOperatorSigningKey(address _operator) view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetLastestOperatorSigningKey(_operator common.Address) (common.Address, error) {
	return _ECDSAStakeRegistry.Contract.GetLastestOperatorSigningKey(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// GetOperatorSigningKeyAtBlock is a free data retrieval call binding the contract method 0x5e1042e8.
//
// Solidity: function getOperatorSigningKeyAtBlock(address _operator, uint256 _blockNumber) view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetOperatorSigningKeyAtBlock(opts *bind.CallOpts, _operator common.Address, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getOperatorSigningKeyAtBlock", _operator, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorSigningKeyAtBlock is a free data retrieval call binding the contract method 0x5e1042e8.
//
// Solidity: function getOperatorSigningKeyAtBlock(address _operator, uint256 _blockNumber) view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetOperatorSigningKeyAtBlock(_operator common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _ECDSAStakeRegistry.Contract.GetOperatorSigningKeyAtBlock(&_ECDSAStakeRegistry.CallOpts, _operator, _blockNumber)
}

// GetOperatorSigningKeyAtBlock is a free data retrieval call binding the contract method 0x5e1042e8.
//
// Solidity: function getOperatorSigningKeyAtBlock(address _operator, uint256 _blockNumber) view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetOperatorSigningKeyAtBlock(_operator common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _ECDSAStakeRegistry.Contract.GetOperatorSigningKeyAtBlock(&_ECDSAStakeRegistry.CallOpts, _operator, _blockNumber)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetOperatorWeight(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getOperatorWeight", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetOperatorWeight(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// GetOperatorWeight is a free data retrieval call binding the contract method 0x98ec1ac9.
//
// Solidity: function getOperatorWeight(address _operator) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetOperatorWeight(_operator common.Address) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetOperatorWeight(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) GetOperatorWeightAtBlock(opts *bind.CallOpts, _operator common.Address, _blockNumber uint32) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "getOperatorWeightAtBlock", _operator, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) GetOperatorWeightAtBlock(_operator common.Address, _blockNumber uint32) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetOperatorWeightAtBlock(&_ECDSAStakeRegistry.CallOpts, _operator, _blockNumber)
}

// GetOperatorWeightAtBlock is a free data retrieval call binding the contract method 0x955f2d90.
//
// Solidity: function getOperatorWeightAtBlock(address _operator, uint32 _blockNumber) view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) GetOperatorWeightAtBlock(_operator common.Address, _blockNumber uint32) (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.GetOperatorWeightAtBlock(&_ECDSAStakeRegistry.CallOpts, _operator, _blockNumber)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) IsValidSignature(opts *bind.CallOpts, _dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "isValidSignature", _dataHash, _signatureData)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) IsValidSignature(_dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	return _ECDSAStakeRegistry.Contract.IsValidSignature(&_ECDSAStakeRegistry.CallOpts, _dataHash, _signatureData)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _dataHash, bytes _signatureData) view returns(bytes4)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) IsValidSignature(_dataHash [32]byte, _signatureData []byte) ([4]byte, error) {
	return _ECDSAStakeRegistry.Contract.IsValidSignature(&_ECDSAStakeRegistry.CallOpts, _dataHash, _signatureData)
}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) MinimumWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "minimumWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) MinimumWeight() (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.MinimumWeight(&_ECDSAStakeRegistry.CallOpts)
}

// MinimumWeight is a free data retrieval call binding the contract method 0x40bf2fb7.
//
// Solidity: function minimumWeight() view returns(uint256)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) MinimumWeight() (*big.Int, error) {
	return _ECDSAStakeRegistry.Contract.MinimumWeight(&_ECDSAStakeRegistry.CallOpts)
}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) OperatorRegistered(opts *bind.CallOpts, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "operatorRegistered", _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) OperatorRegistered(_operator common.Address) (bool, error) {
	return _ECDSAStakeRegistry.Contract.OperatorRegistered(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// OperatorRegistered is a free data retrieval call binding the contract method 0xec7fbb31.
//
// Solidity: function operatorRegistered(address _operator) view returns(bool)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) OperatorRegistered(_operator common.Address) (bool, error) {
	return _ECDSAStakeRegistry.Contract.OperatorRegistered(&_ECDSAStakeRegistry.CallOpts, _operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) Owner() (common.Address, error) {
	return _ECDSAStakeRegistry.Contract.Owner(&_ECDSAStakeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) Owner() (common.Address, error) {
	return _ECDSAStakeRegistry.Contract.Owner(&_ECDSAStakeRegistry.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCaller) Quorum(opts *bind.CallOpts) (Quorum, error) {
	var out []interface{}
	err := _ECDSAStakeRegistry.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(Quorum), err
	}

	out0 := *abi.ConvertType(out[0], new(Quorum)).(*Quorum)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) Quorum() (Quorum, error) {
	return _ECDSAStakeRegistry.Contract.Quorum(&_ECDSAStakeRegistry.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_ECDSAStakeRegistry *ECDSAStakeRegistryCallerSession) Quorum() (Quorum, error) {
	return _ECDSAStakeRegistry.Contract.Quorum(&_ECDSAStakeRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "deregisterOperator")
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) DeregisterOperator() (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.DeregisterOperator(&_ECDSAStakeRegistry.TransactOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) DeregisterOperator() (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.DeregisterOperator(&_ECDSAStakeRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) Initialize(opts *bind.TransactOpts, _serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "initialize", _serviceManager, _thresholdWeight, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) Initialize(_serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.Initialize(&_ECDSAStakeRegistry.TransactOpts, _serviceManager, _thresholdWeight, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0xab118995.
//
// Solidity: function initialize(address _serviceManager, uint256 _thresholdWeight, ((address,uint96)[]) _quorum) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) Initialize(_serviceManager common.Address, _thresholdWeight *big.Int, _quorum Quorum) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.Initialize(&_ECDSAStakeRegistry.TransactOpts, _serviceManager, _thresholdWeight, _quorum)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x3d5611f6.
//
// Solidity: function registerOperatorWithSignature((bytes,bytes32,uint256) _operatorSignature, address _signingKey) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) RegisterOperatorWithSignature(opts *bind.TransactOpts, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _signingKey common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "registerOperatorWithSignature", _operatorSignature, _signingKey)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x3d5611f6.
//
// Solidity: function registerOperatorWithSignature((bytes,bytes32,uint256) _operatorSignature, address _signingKey) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) RegisterOperatorWithSignature(_operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _signingKey common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.RegisterOperatorWithSignature(&_ECDSAStakeRegistry.TransactOpts, _operatorSignature, _signingKey)
}

// RegisterOperatorWithSignature is a paid mutator transaction binding the contract method 0x3d5611f6.
//
// Solidity: function registerOperatorWithSignature((bytes,bytes32,uint256) _operatorSignature, address _signingKey) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) RegisterOperatorWithSignature(_operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _signingKey common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.RegisterOperatorWithSignature(&_ECDSAStakeRegistry.TransactOpts, _operatorSignature, _signingKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.RenounceOwnership(&_ECDSAStakeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.RenounceOwnership(&_ECDSAStakeRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.TransferOwnership(&_ECDSAStakeRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.TransferOwnership(&_ECDSAStakeRegistry.TransactOpts, newOwner)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) UpdateMinimumWeight(opts *bind.TransactOpts, _newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "updateMinimumWeight", _newMinimumWeight, _operators)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) UpdateMinimumWeight(_newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateMinimumWeight(&_ECDSAStakeRegistry.TransactOpts, _newMinimumWeight, _operators)
}

// UpdateMinimumWeight is a paid mutator transaction binding the contract method 0x696255be.
//
// Solidity: function updateMinimumWeight(uint256 _newMinimumWeight, address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) UpdateMinimumWeight(_newMinimumWeight *big.Int, _operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateMinimumWeight(&_ECDSAStakeRegistry.TransactOpts, _newMinimumWeight, _operators)
}

// UpdateOperatorSigningKey is a paid mutator transaction binding the contract method 0x743c31f4.
//
// Solidity: function updateOperatorSigningKey(address _newSigningKey) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) UpdateOperatorSigningKey(opts *bind.TransactOpts, _newSigningKey common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "updateOperatorSigningKey", _newSigningKey)
}

// UpdateOperatorSigningKey is a paid mutator transaction binding the contract method 0x743c31f4.
//
// Solidity: function updateOperatorSigningKey(address _newSigningKey) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) UpdateOperatorSigningKey(_newSigningKey common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateOperatorSigningKey(&_ECDSAStakeRegistry.TransactOpts, _newSigningKey)
}

// UpdateOperatorSigningKey is a paid mutator transaction binding the contract method 0x743c31f4.
//
// Solidity: function updateOperatorSigningKey(address _newSigningKey) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) UpdateOperatorSigningKey(_newSigningKey common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateOperatorSigningKey(&_ECDSAStakeRegistry.TransactOpts, _newSigningKey)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) UpdateOperators(opts *bind.TransactOpts, _operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "updateOperators", _operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) UpdateOperators(_operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateOperators(&_ECDSAStakeRegistry.TransactOpts, _operators)
}

// UpdateOperators is a paid mutator transaction binding the contract method 0x00cf2ab5.
//
// Solidity: function updateOperators(address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) UpdateOperators(_operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateOperators(&_ECDSAStakeRegistry.TransactOpts, _operators)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, arg1)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateOperatorsForQuorum(&_ECDSAStakeRegistry.TransactOpts, operatorsPerQuorum, arg1)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes ) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, arg1 []byte) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateOperatorsForQuorum(&_ECDSAStakeRegistry.TransactOpts, operatorsPerQuorum, arg1)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) UpdateQuorumConfig(opts *bind.TransactOpts, _quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "updateQuorumConfig", _quorum, _operators)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) UpdateQuorumConfig(_quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateQuorumConfig(&_ECDSAStakeRegistry.TransactOpts, _quorum, _operators)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xdec5d1f6.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) _quorum, address[] _operators) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) UpdateQuorumConfig(_quorum Quorum, _operators []common.Address) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateQuorumConfig(&_ECDSAStakeRegistry.TransactOpts, _quorum, _operators)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactor) UpdateStakeThreshold(opts *bind.TransactOpts, _thresholdWeight *big.Int) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.contract.Transact(opts, "updateStakeThreshold", _thresholdWeight)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistrySession) UpdateStakeThreshold(_thresholdWeight *big.Int) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateStakeThreshold(&_ECDSAStakeRegistry.TransactOpts, _thresholdWeight)
}

// UpdateStakeThreshold is a paid mutator transaction binding the contract method 0x5ef53329.
//
// Solidity: function updateStakeThreshold(uint256 _thresholdWeight) returns()
func (_ECDSAStakeRegistry *ECDSAStakeRegistryTransactorSession) UpdateStakeThreshold(_thresholdWeight *big.Int) (*types.Transaction, error) {
	return _ECDSAStakeRegistry.Contract.UpdateStakeThreshold(&_ECDSAStakeRegistry.TransactOpts, _thresholdWeight)
}

// ECDSAStakeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryInitializedIterator struct {
	Event *ECDSAStakeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryInitialized)
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
		it.Event = new(ECDSAStakeRegistryInitialized)
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
func (it *ECDSAStakeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryInitialized represents a Initialized event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ECDSAStakeRegistryInitializedIterator, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryInitializedIterator{contract: _ECDSAStakeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryInitialized)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseInitialized(log types.Log) (*ECDSAStakeRegistryInitialized, error) {
	event := new(ECDSAStakeRegistryInitialized)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryMinimumWeightUpdatedIterator is returned from FilterMinimumWeightUpdated and is used to iterate over the raw logs and unpacked data for MinimumWeightUpdated events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryMinimumWeightUpdatedIterator struct {
	Event *ECDSAStakeRegistryMinimumWeightUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryMinimumWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryMinimumWeightUpdated)
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
		it.Event = new(ECDSAStakeRegistryMinimumWeightUpdated)
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
func (it *ECDSAStakeRegistryMinimumWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryMinimumWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryMinimumWeightUpdated represents a MinimumWeightUpdated event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryMinimumWeightUpdated struct {
	Old *big.Int
	New *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinimumWeightUpdated is a free log retrieval operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterMinimumWeightUpdated(opts *bind.FilterOpts) (*ECDSAStakeRegistryMinimumWeightUpdatedIterator, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "MinimumWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryMinimumWeightUpdatedIterator{contract: _ECDSAStakeRegistry.contract, event: "MinimumWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumWeightUpdated is a free log subscription operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchMinimumWeightUpdated(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryMinimumWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "MinimumWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryMinimumWeightUpdated)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "MinimumWeightUpdated", log); err != nil {
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

// ParseMinimumWeightUpdated is a log parse operation binding the contract event 0x713ca53b88d6eb63f5b1854cb8cbdd736ec51eda225e46791aa9298b0160648f.
//
// Solidity: event MinimumWeightUpdated(uint256 _old, uint256 _new)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseMinimumWeightUpdated(log types.Log) (*ECDSAStakeRegistryMinimumWeightUpdated, error) {
	event := new(ECDSAStakeRegistryMinimumWeightUpdated)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "MinimumWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOperatorDeregisteredIterator struct {
	Event *ECDSAStakeRegistryOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryOperatorDeregistered)
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
		it.Event = new(ECDSAStakeRegistryOperatorDeregistered)
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
func (it *ECDSAStakeRegistryOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryOperatorDeregistered represents a OperatorDeregistered event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOperatorDeregistered struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, _operator []common.Address, _avs []common.Address) (*ECDSAStakeRegistryOperatorDeregisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "OperatorDeregistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryOperatorDeregisteredIterator{contract: _ECDSAStakeRegistry.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryOperatorDeregistered, _operator []common.Address, _avs []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "OperatorDeregistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryOperatorDeregistered)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed _operator, address indexed _avs)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseOperatorDeregistered(log types.Log) (*ECDSAStakeRegistryOperatorDeregistered, error) {
	event := new(ECDSAStakeRegistryOperatorDeregistered)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOperatorRegisteredIterator struct {
	Event *ECDSAStakeRegistryOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryOperatorRegistered)
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
		it.Event = new(ECDSAStakeRegistryOperatorRegistered)
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
func (it *ECDSAStakeRegistryOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryOperatorRegistered represents a OperatorRegistered event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOperatorRegistered struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, _operator []common.Address, _avs []common.Address) (*ECDSAStakeRegistryOperatorRegisteredIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "OperatorRegistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryOperatorRegisteredIterator{contract: _ECDSAStakeRegistry.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryOperatorRegistered, _operator []common.Address, _avs []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}
	var _avsRule []interface{}
	for _, _avsItem := range _avs {
		_avsRule = append(_avsRule, _avsItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "OperatorRegistered", _operatorRule, _avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryOperatorRegistered)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed _operator, address indexed _avs)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseOperatorRegistered(log types.Log) (*ECDSAStakeRegistryOperatorRegistered, error) {
	event := new(ECDSAStakeRegistryOperatorRegistered)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryOperatorWeightUpdatedIterator is returned from FilterOperatorWeightUpdated and is used to iterate over the raw logs and unpacked data for OperatorWeightUpdated events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOperatorWeightUpdatedIterator struct {
	Event *ECDSAStakeRegistryOperatorWeightUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryOperatorWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryOperatorWeightUpdated)
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
		it.Event = new(ECDSAStakeRegistryOperatorWeightUpdated)
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
func (it *ECDSAStakeRegistryOperatorWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryOperatorWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryOperatorWeightUpdated represents a OperatorWeightUpdated event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOperatorWeightUpdated struct {
	Operator  common.Address
	OldWeight *big.Int
	NewWeight *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorWeightUpdated is a free log retrieval operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterOperatorWeightUpdated(opts *bind.FilterOpts, _operator []common.Address) (*ECDSAStakeRegistryOperatorWeightUpdatedIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "OperatorWeightUpdated", _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryOperatorWeightUpdatedIterator{contract: _ECDSAStakeRegistry.contract, event: "OperatorWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWeightUpdated is a free log subscription operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchOperatorWeightUpdated(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryOperatorWeightUpdated, _operator []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "OperatorWeightUpdated", _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryOperatorWeightUpdated)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OperatorWeightUpdated", log); err != nil {
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

// ParseOperatorWeightUpdated is a log parse operation binding the contract event 0x88770dc862e47a7ed586907857eb1b75e4c5ffc8b707c7ee10eb74d6885fe594.
//
// Solidity: event OperatorWeightUpdated(address indexed _operator, uint256 oldWeight, uint256 newWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseOperatorWeightUpdated(log types.Log) (*ECDSAStakeRegistryOperatorWeightUpdated, error) {
	event := new(ECDSAStakeRegistryOperatorWeightUpdated)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OperatorWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOwnershipTransferredIterator struct {
	Event *ECDSAStakeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryOwnershipTransferred)
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
		it.Event = new(ECDSAStakeRegistryOwnershipTransferred)
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
func (it *ECDSAStakeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ECDSAStakeRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryOwnershipTransferredIterator{contract: _ECDSAStakeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryOwnershipTransferred)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ECDSAStakeRegistryOwnershipTransferred, error) {
	event := new(ECDSAStakeRegistryOwnershipTransferred)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryQuorumUpdatedIterator is returned from FilterQuorumUpdated and is used to iterate over the raw logs and unpacked data for QuorumUpdated events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryQuorumUpdatedIterator struct {
	Event *ECDSAStakeRegistryQuorumUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryQuorumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryQuorumUpdated)
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
		it.Event = new(ECDSAStakeRegistryQuorumUpdated)
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
func (it *ECDSAStakeRegistryQuorumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryQuorumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryQuorumUpdated represents a QuorumUpdated event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryQuorumUpdated struct {
	Old Quorum
	New Quorum
	Raw types.Log // Blockchain specific contextual infos
}

// FilterQuorumUpdated is a free log retrieval operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterQuorumUpdated(opts *bind.FilterOpts) (*ECDSAStakeRegistryQuorumUpdatedIterator, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryQuorumUpdatedIterator{contract: _ECDSAStakeRegistry.contract, event: "QuorumUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumUpdated is a free log subscription operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchQuorumUpdated(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryQuorumUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryQuorumUpdated)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
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

// ParseQuorumUpdated is a log parse operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) _old, ((address,uint96)[]) _new)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseQuorumUpdated(log types.Log) (*ECDSAStakeRegistryQuorumUpdated, error) {
	event := new(ECDSAStakeRegistryQuorumUpdated)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistrySigningKeyUpdateIterator is returned from FilterSigningKeyUpdate and is used to iterate over the raw logs and unpacked data for SigningKeyUpdate events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistrySigningKeyUpdateIterator struct {
	Event *ECDSAStakeRegistrySigningKeyUpdate // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistrySigningKeyUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistrySigningKeyUpdate)
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
		it.Event = new(ECDSAStakeRegistrySigningKeyUpdate)
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
func (it *ECDSAStakeRegistrySigningKeyUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistrySigningKeyUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistrySigningKeyUpdate represents a SigningKeyUpdate event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistrySigningKeyUpdate struct {
	Operator      common.Address
	UpdateBlock   *big.Int
	NewSigningKey common.Address
	OldSigningKey common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyUpdate is a free log retrieval operation binding the contract event 0xd061168252f441733658f09e4d8f5b2d998ed4ef24a2bbfd6ceca52ea1315002.
//
// Solidity: event SigningKeyUpdate(address indexed operator, uint256 indexed updateBlock, address indexed newSigningKey, address oldSigningKey)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterSigningKeyUpdate(opts *bind.FilterOpts, operator []common.Address, updateBlock []*big.Int, newSigningKey []common.Address) (*ECDSAStakeRegistrySigningKeyUpdateIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var updateBlockRule []interface{}
	for _, updateBlockItem := range updateBlock {
		updateBlockRule = append(updateBlockRule, updateBlockItem)
	}
	var newSigningKeyRule []interface{}
	for _, newSigningKeyItem := range newSigningKey {
		newSigningKeyRule = append(newSigningKeyRule, newSigningKeyItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "SigningKeyUpdate", operatorRule, updateBlockRule, newSigningKeyRule)
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistrySigningKeyUpdateIterator{contract: _ECDSAStakeRegistry.contract, event: "SigningKeyUpdate", logs: logs, sub: sub}, nil
}

// WatchSigningKeyUpdate is a free log subscription operation binding the contract event 0xd061168252f441733658f09e4d8f5b2d998ed4ef24a2bbfd6ceca52ea1315002.
//
// Solidity: event SigningKeyUpdate(address indexed operator, uint256 indexed updateBlock, address indexed newSigningKey, address oldSigningKey)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchSigningKeyUpdate(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistrySigningKeyUpdate, operator []common.Address, updateBlock []*big.Int, newSigningKey []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var updateBlockRule []interface{}
	for _, updateBlockItem := range updateBlock {
		updateBlockRule = append(updateBlockRule, updateBlockItem)
	}
	var newSigningKeyRule []interface{}
	for _, newSigningKeyItem := range newSigningKey {
		newSigningKeyRule = append(newSigningKeyRule, newSigningKeyItem)
	}

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "SigningKeyUpdate", operatorRule, updateBlockRule, newSigningKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistrySigningKeyUpdate)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "SigningKeyUpdate", log); err != nil {
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

// ParseSigningKeyUpdate is a log parse operation binding the contract event 0xd061168252f441733658f09e4d8f5b2d998ed4ef24a2bbfd6ceca52ea1315002.
//
// Solidity: event SigningKeyUpdate(address indexed operator, uint256 indexed updateBlock, address indexed newSigningKey, address oldSigningKey)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseSigningKeyUpdate(log types.Log) (*ECDSAStakeRegistrySigningKeyUpdate, error) {
	event := new(ECDSAStakeRegistrySigningKeyUpdate)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "SigningKeyUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryThresholdWeightUpdatedIterator is returned from FilterThresholdWeightUpdated and is used to iterate over the raw logs and unpacked data for ThresholdWeightUpdated events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryThresholdWeightUpdatedIterator struct {
	Event *ECDSAStakeRegistryThresholdWeightUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryThresholdWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryThresholdWeightUpdated)
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
		it.Event = new(ECDSAStakeRegistryThresholdWeightUpdated)
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
func (it *ECDSAStakeRegistryThresholdWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryThresholdWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryThresholdWeightUpdated represents a ThresholdWeightUpdated event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryThresholdWeightUpdated struct {
	ThresholdWeight *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterThresholdWeightUpdated is a free log retrieval operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterThresholdWeightUpdated(opts *bind.FilterOpts) (*ECDSAStakeRegistryThresholdWeightUpdatedIterator, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "ThresholdWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryThresholdWeightUpdatedIterator{contract: _ECDSAStakeRegistry.contract, event: "ThresholdWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdWeightUpdated is a free log subscription operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchThresholdWeightUpdated(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryThresholdWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "ThresholdWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryThresholdWeightUpdated)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "ThresholdWeightUpdated", log); err != nil {
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

// ParseThresholdWeightUpdated is a log parse operation binding the contract event 0x9324f7e5a7c0288808a634ccde44b8e979676474b22e29ee9dd569b55e791a4b.
//
// Solidity: event ThresholdWeightUpdated(uint256 _thresholdWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseThresholdWeightUpdated(log types.Log) (*ECDSAStakeRegistryThresholdWeightUpdated, error) {
	event := new(ECDSAStakeRegistryThresholdWeightUpdated)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "ThresholdWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryTotalWeightUpdatedIterator is returned from FilterTotalWeightUpdated and is used to iterate over the raw logs and unpacked data for TotalWeightUpdated events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryTotalWeightUpdatedIterator struct {
	Event *ECDSAStakeRegistryTotalWeightUpdated // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryTotalWeightUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryTotalWeightUpdated)
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
		it.Event = new(ECDSAStakeRegistryTotalWeightUpdated)
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
func (it *ECDSAStakeRegistryTotalWeightUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryTotalWeightUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryTotalWeightUpdated represents a TotalWeightUpdated event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryTotalWeightUpdated struct {
	OldTotalWeight *big.Int
	NewTotalWeight *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalWeightUpdated is a free log retrieval operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterTotalWeightUpdated(opts *bind.FilterOpts) (*ECDSAStakeRegistryTotalWeightUpdatedIterator, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "TotalWeightUpdated")
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryTotalWeightUpdatedIterator{contract: _ECDSAStakeRegistry.contract, event: "TotalWeightUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalWeightUpdated is a free log subscription operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchTotalWeightUpdated(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryTotalWeightUpdated) (event.Subscription, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "TotalWeightUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryTotalWeightUpdated)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "TotalWeightUpdated", log); err != nil {
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

// ParseTotalWeightUpdated is a log parse operation binding the contract event 0x86dcf86b12dfeedea74ae9300dbdaa193bcce5809369c8177ea2f4eaaa65729b.
//
// Solidity: event TotalWeightUpdated(uint256 oldTotalWeight, uint256 newTotalWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseTotalWeightUpdated(log types.Log) (*ECDSAStakeRegistryTotalWeightUpdated, error) {
	event := new(ECDSAStakeRegistryTotalWeightUpdated)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "TotalWeightUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAStakeRegistryUpdateMinimumWeightIterator is returned from FilterUpdateMinimumWeight and is used to iterate over the raw logs and unpacked data for UpdateMinimumWeight events raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryUpdateMinimumWeightIterator struct {
	Event *ECDSAStakeRegistryUpdateMinimumWeight // Event containing the contract specifics and raw log

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
func (it *ECDSAStakeRegistryUpdateMinimumWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ECDSAStakeRegistryUpdateMinimumWeight)
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
		it.Event = new(ECDSAStakeRegistryUpdateMinimumWeight)
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
func (it *ECDSAStakeRegistryUpdateMinimumWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ECDSAStakeRegistryUpdateMinimumWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ECDSAStakeRegistryUpdateMinimumWeight represents a UpdateMinimumWeight event raised by the ECDSAStakeRegistry contract.
type ECDSAStakeRegistryUpdateMinimumWeight struct {
	OldMinimumWeight *big.Int
	NewMinimumWeight *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumWeight is a free log retrieval operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) FilterUpdateMinimumWeight(opts *bind.FilterOpts) (*ECDSAStakeRegistryUpdateMinimumWeightIterator, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.FilterLogs(opts, "UpdateMinimumWeight")
	if err != nil {
		return nil, err
	}
	return &ECDSAStakeRegistryUpdateMinimumWeightIterator{contract: _ECDSAStakeRegistry.contract, event: "UpdateMinimumWeight", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumWeight is a free log subscription operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) WatchUpdateMinimumWeight(opts *bind.WatchOpts, sink chan<- *ECDSAStakeRegistryUpdateMinimumWeight) (event.Subscription, error) {

	logs, sub, err := _ECDSAStakeRegistry.contract.WatchLogs(opts, "UpdateMinimumWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ECDSAStakeRegistryUpdateMinimumWeight)
				if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "UpdateMinimumWeight", log); err != nil {
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

// ParseUpdateMinimumWeight is a log parse operation binding the contract event 0x1ea42186b305fa37310450d9fb87ea1e8f0c7f447e771479e3b27634bfe84dc1.
//
// Solidity: event UpdateMinimumWeight(uint256 oldMinimumWeight, uint256 newMinimumWeight)
func (_ECDSAStakeRegistry *ECDSAStakeRegistryFilterer) ParseUpdateMinimumWeight(log types.Log) (*ECDSAStakeRegistryUpdateMinimumWeight, error) {
	event := new(ECDSAStakeRegistryUpdateMinimumWeight)
	if err := _ECDSAStakeRegistry.contract.UnpackLog(event, "UpdateMinimumWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
