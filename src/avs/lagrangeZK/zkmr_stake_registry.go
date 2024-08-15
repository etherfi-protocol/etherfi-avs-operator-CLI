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

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X *big.Int
	Y *big.Int
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

// ZKMRStakeRegistryMetaData contains all meta data concerning the ZKMRStakeRegistry contract.
var ZKMRStakeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQuorum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyHasBeenUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewOwnerIsZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoHandoverRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServiceManagerAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"BatchWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newShares\",\"type\":\"uint256\"}],\"name\":\"MinimumSharesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"}],\"name\":\"OperatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"}],\"name\":\"OperatorEvicted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"avs\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"OperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipHandoverRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structQuorum\",\"name\":\"oldQuorum\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structQuorum\",\"name\":\"newQuorum\",\"type\":\"tuple\"}],\"name\":\"QuorumUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"Whitelisted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"completeOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"evictOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegationManager_\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"internalType\":\"structQuorum\",\"name\":\"quorum_\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumShares_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"keyHasBeenUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"result\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"ownershipHandoverExpiresAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"internalType\":\"structQuorum\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestOwnershipHandover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"serviceManager_\",\"type\":\"address\"}],\"name\":\"setServiceManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"toggleWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMinimumShares\",\"type\":\"uint256\"}],\"name\":\"updateMinimumShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"updateOperatorKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structStrategyParams[]\",\"name\":\"strategies\",\"type\":\"tuple[]\"}],\"internalType\":\"structQuorum\",\"name\":\"quorum_\",\"type\":\"tuple\"}],\"name\":\"updateQuorumConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZKMRStakeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKMRStakeRegistryMetaData.ABI instead.
var ZKMRStakeRegistryABI = ZKMRStakeRegistryMetaData.ABI

// ZKMRStakeRegistry is an auto generated Go binding around an Ethereum contract.
type ZKMRStakeRegistry struct {
	ZKMRStakeRegistryCaller     // Read-only binding to the contract
	ZKMRStakeRegistryTransactor // Write-only binding to the contract
	ZKMRStakeRegistryFilterer   // Log filterer for contract events
}

// ZKMRStakeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZKMRStakeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMRStakeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKMRStakeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMRStakeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKMRStakeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKMRStakeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKMRStakeRegistrySession struct {
	Contract     *ZKMRStakeRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZKMRStakeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKMRStakeRegistryCallerSession struct {
	Contract *ZKMRStakeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZKMRStakeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKMRStakeRegistryTransactorSession struct {
	Contract     *ZKMRStakeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZKMRStakeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZKMRStakeRegistryRaw struct {
	Contract *ZKMRStakeRegistry // Generic contract binding to access the raw methods on
}

// ZKMRStakeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKMRStakeRegistryCallerRaw struct {
	Contract *ZKMRStakeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ZKMRStakeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKMRStakeRegistryTransactorRaw struct {
	Contract *ZKMRStakeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZKMRStakeRegistry creates a new instance of ZKMRStakeRegistry, bound to a specific deployed contract.
func NewZKMRStakeRegistry(address common.Address, backend bind.ContractBackend) (*ZKMRStakeRegistry, error) {
	contract, err := bindZKMRStakeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistry{ZKMRStakeRegistryCaller: ZKMRStakeRegistryCaller{contract: contract}, ZKMRStakeRegistryTransactor: ZKMRStakeRegistryTransactor{contract: contract}, ZKMRStakeRegistryFilterer: ZKMRStakeRegistryFilterer{contract: contract}}, nil
}

// NewZKMRStakeRegistryCaller creates a new read-only instance of ZKMRStakeRegistry, bound to a specific deployed contract.
func NewZKMRStakeRegistryCaller(address common.Address, caller bind.ContractCaller) (*ZKMRStakeRegistryCaller, error) {
	contract, err := bindZKMRStakeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryCaller{contract: contract}, nil
}

// NewZKMRStakeRegistryTransactor creates a new write-only instance of ZKMRStakeRegistry, bound to a specific deployed contract.
func NewZKMRStakeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZKMRStakeRegistryTransactor, error) {
	contract, err := bindZKMRStakeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryTransactor{contract: contract}, nil
}

// NewZKMRStakeRegistryFilterer creates a new log filterer instance of ZKMRStakeRegistry, bound to a specific deployed contract.
func NewZKMRStakeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZKMRStakeRegistryFilterer, error) {
	contract, err := bindZKMRStakeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryFilterer{contract: contract}, nil
}

// bindZKMRStakeRegistry binds a generic wrapper to an already deployed contract.
func bindZKMRStakeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKMRStakeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMRStakeRegistry *ZKMRStakeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMRStakeRegistry.Contract.ZKMRStakeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMRStakeRegistry *ZKMRStakeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.ZKMRStakeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMRStakeRegistry *ZKMRStakeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.ZKMRStakeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKMRStakeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.contract.Transact(opts, method, params...)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) DelegationManager() (common.Address, error) {
	return _ZKMRStakeRegistry.Contract.DelegationManager(&_ZKMRStakeRegistry.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) DelegationManager() (common.Address, error) {
	return _ZKMRStakeRegistry.Contract.DelegationManager(&_ZKMRStakeRegistry.CallOpts)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0xcd8e9cd3.
//
// Solidity: function getOperatorShares(address operator) view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) GetOperatorShares(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "getOperatorShares", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorShares is a free data retrieval call binding the contract method 0xcd8e9cd3.
//
// Solidity: function getOperatorShares(address operator) view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) GetOperatorShares(operator common.Address) (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.GetOperatorShares(&_ZKMRStakeRegistry.CallOpts, operator)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0xcd8e9cd3.
//
// Solidity: function getOperatorShares(address operator) view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) GetOperatorShares(operator common.Address) (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.GetOperatorShares(&_ZKMRStakeRegistry.CallOpts, operator)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address operator) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) IsRegistered(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "isRegistered", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address operator) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) IsRegistered(operator common.Address) (bool, error) {
	return _ZKMRStakeRegistry.Contract.IsRegistered(&_ZKMRStakeRegistry.CallOpts, operator)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address operator) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) IsRegistered(operator common.Address) (bool, error) {
	return _ZKMRStakeRegistry.Contract.IsRegistered(&_ZKMRStakeRegistry.CallOpts, operator)
}

// KeyHasBeenUsed is a free data retrieval call binding the contract method 0x3e92d0bc.
//
// Solidity: function keyHasBeenUsed((uint256,uint256) publicKey) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) KeyHasBeenUsed(opts *bind.CallOpts, publicKey PublicKey) (bool, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "keyHasBeenUsed", publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyHasBeenUsed is a free data retrieval call binding the contract method 0x3e92d0bc.
//
// Solidity: function keyHasBeenUsed((uint256,uint256) publicKey) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) KeyHasBeenUsed(publicKey PublicKey) (bool, error) {
	return _ZKMRStakeRegistry.Contract.KeyHasBeenUsed(&_ZKMRStakeRegistry.CallOpts, publicKey)
}

// KeyHasBeenUsed is a free data retrieval call binding the contract method 0x3e92d0bc.
//
// Solidity: function keyHasBeenUsed((uint256,uint256) publicKey) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) KeyHasBeenUsed(publicKey PublicKey) (bool, error) {
	return _ZKMRStakeRegistry.Contract.KeyHasBeenUsed(&_ZKMRStakeRegistry.CallOpts, publicKey)
}

// MinimumShares is a free data retrieval call binding the contract method 0x8d78fd0b.
//
// Solidity: function minimumShares() view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) MinimumShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "minimumShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumShares is a free data retrieval call binding the contract method 0x8d78fd0b.
//
// Solidity: function minimumShares() view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) MinimumShares() (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.MinimumShares(&_ZKMRStakeRegistry.CallOpts)
}

// MinimumShares is a free data retrieval call binding the contract method 0x8d78fd0b.
//
// Solidity: function minimumShares() view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) MinimumShares() (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.MinimumShares(&_ZKMRStakeRegistry.CallOpts)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address operator) view returns(uint256 x, uint256 y)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) Operators(opts *bind.CallOpts, operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "operators", operator)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address operator) view returns(uint256 x, uint256 y)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) Operators(operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ZKMRStakeRegistry.Contract.Operators(&_ZKMRStakeRegistry.CallOpts, operator)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address operator) view returns(uint256 x, uint256 y)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) Operators(operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ZKMRStakeRegistry.Contract.Operators(&_ZKMRStakeRegistry.CallOpts, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) Owner() (common.Address, error) {
	return _ZKMRStakeRegistry.Contract.Owner(&_ZKMRStakeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) Owner() (common.Address, error) {
	return _ZKMRStakeRegistry.Contract.Owner(&_ZKMRStakeRegistry.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.OwnershipHandoverExpiresAt(&_ZKMRStakeRegistry.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.OwnershipHandoverExpiresAt(&_ZKMRStakeRegistry.CallOpts, pendingOwner)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) Quorum(opts *bind.CallOpts) (Quorum, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(Quorum), err
	}

	out0 := *abi.ConvertType(out[0], new(Quorum)).(*Quorum)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) Quorum() (Quorum, error) {
	return _ZKMRStakeRegistry.Contract.Quorum(&_ZKMRStakeRegistry.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(((address,uint96)[]))
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) Quorum() (Quorum, error) {
	return _ZKMRStakeRegistry.Contract.Quorum(&_ZKMRStakeRegistry.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) ServiceManager() (common.Address, error) {
	return _ZKMRStakeRegistry.Contract.ServiceManager(&_ZKMRStakeRegistry.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) ServiceManager() (common.Address, error) {
	return _ZKMRStakeRegistry.Contract.ServiceManager(&_ZKMRStakeRegistry.CallOpts)
}

// TotalOperators is a free data retrieval call binding the contract method 0x492ec79f.
//
// Solidity: function totalOperators() view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) TotalOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "totalOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOperators is a free data retrieval call binding the contract method 0x492ec79f.
//
// Solidity: function totalOperators() view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) TotalOperators() (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.TotalOperators(&_ZKMRStakeRegistry.CallOpts)
}

// TotalOperators is a free data retrieval call binding the contract method 0x492ec79f.
//
// Solidity: function totalOperators() view returns(uint256)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) TotalOperators() (*big.Int, error) {
	return _ZKMRStakeRegistry.Contract.TotalOperators(&_ZKMRStakeRegistry.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ZKMRStakeRegistry.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) Whitelist(arg0 common.Address) (bool, error) {
	return _ZKMRStakeRegistry.Contract.Whitelist(&_ZKMRStakeRegistry.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _ZKMRStakeRegistry.Contract.Whitelist(&_ZKMRStakeRegistry.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] addrs) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) AddToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "addToWhitelist", addrs)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] addrs) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) AddToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.AddToWhitelist(&_ZKMRStakeRegistry.TransactOpts, addrs)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] addrs) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) AddToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.AddToWhitelist(&_ZKMRStakeRegistry.TransactOpts, addrs)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.CancelOwnershipHandover(&_ZKMRStakeRegistry.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.CancelOwnershipHandover(&_ZKMRStakeRegistry.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.CompleteOwnershipHandover(&_ZKMRStakeRegistry.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.CompleteOwnershipHandover(&_ZKMRStakeRegistry.TransactOpts, pendingOwner)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "deregisterOperator")
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) DeregisterOperator() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.DeregisterOperator(&_ZKMRStakeRegistry.TransactOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x857dc190.
//
// Solidity: function deregisterOperator() returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) DeregisterOperator() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.DeregisterOperator(&_ZKMRStakeRegistry.TransactOpts)
}

// EvictOperator is a paid mutator transaction binding the contract method 0x01696902.
//
// Solidity: function evictOperator(address operator) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) EvictOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "evictOperator", operator)
}

// EvictOperator is a paid mutator transaction binding the contract method 0x01696902.
//
// Solidity: function evictOperator(address operator) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) EvictOperator(operator common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.EvictOperator(&_ZKMRStakeRegistry.TransactOpts, operator)
}

// EvictOperator is a paid mutator transaction binding the contract method 0x01696902.
//
// Solidity: function evictOperator(address operator) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) EvictOperator(operator common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.EvictOperator(&_ZKMRStakeRegistry.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e5e0ef.
//
// Solidity: function initialize(address delegationManager_, ((address,uint96)[]) quorum_, address owner_, uint256 minimumShares_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) Initialize(opts *bind.TransactOpts, delegationManager_ common.Address, quorum_ Quorum, owner_ common.Address, minimumShares_ *big.Int) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "initialize", delegationManager_, quorum_, owner_, minimumShares_)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e5e0ef.
//
// Solidity: function initialize(address delegationManager_, ((address,uint96)[]) quorum_, address owner_, uint256 minimumShares_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) Initialize(delegationManager_ common.Address, quorum_ Quorum, owner_ common.Address, minimumShares_ *big.Int) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.Initialize(&_ZKMRStakeRegistry.TransactOpts, delegationManager_, quorum_, owner_, minimumShares_)
}

// Initialize is a paid mutator transaction binding the contract method 0x20e5e0ef.
//
// Solidity: function initialize(address delegationManager_, ((address,uint96)[]) quorum_, address owner_, uint256 minimumShares_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) Initialize(delegationManager_ common.Address, quorum_ Quorum, owner_ common.Address, minimumShares_ *big.Int) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.Initialize(&_ZKMRStakeRegistry.TransactOpts, delegationManager_, quorum_, owner_, minimumShares_)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xf8fdb096.
//
// Solidity: function registerOperator((uint256,uint256) publicKey, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, publicKey PublicKey, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "registerOperator", publicKey, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xf8fdb096.
//
// Solidity: function registerOperator((uint256,uint256) publicKey, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) RegisterOperator(publicKey PublicKey, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RegisterOperator(&_ZKMRStakeRegistry.TransactOpts, publicKey, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xf8fdb096.
//
// Solidity: function registerOperator((uint256,uint256) publicKey, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) RegisterOperator(publicKey PublicKey, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RegisterOperator(&_ZKMRStakeRegistry.TransactOpts, publicKey, operatorSignature)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] addrs) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "removeFromWhitelist", addrs)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] addrs) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) RemoveFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RemoveFromWhitelist(&_ZKMRStakeRegistry.TransactOpts, addrs)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] addrs) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) RemoveFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RemoveFromWhitelist(&_ZKMRStakeRegistry.TransactOpts, addrs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RenounceOwnership(&_ZKMRStakeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RenounceOwnership(&_ZKMRStakeRegistry.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RequestOwnershipHandover(&_ZKMRStakeRegistry.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.RequestOwnershipHandover(&_ZKMRStakeRegistry.TransactOpts)
}

// SetServiceManager is a paid mutator transaction binding the contract method 0x9b41bf23.
//
// Solidity: function setServiceManager(address serviceManager_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) SetServiceManager(opts *bind.TransactOpts, serviceManager_ common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "setServiceManager", serviceManager_)
}

// SetServiceManager is a paid mutator transaction binding the contract method 0x9b41bf23.
//
// Solidity: function setServiceManager(address serviceManager_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) SetServiceManager(serviceManager_ common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.SetServiceManager(&_ZKMRStakeRegistry.TransactOpts, serviceManager_)
}

// SetServiceManager is a paid mutator transaction binding the contract method 0x9b41bf23.
//
// Solidity: function setServiceManager(address serviceManager_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) SetServiceManager(serviceManager_ common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.SetServiceManager(&_ZKMRStakeRegistry.TransactOpts, serviceManager_)
}

// ToggleWhitelist is a paid mutator transaction binding the contract method 0x39393ac9.
//
// Solidity: function toggleWhitelist(address client) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) ToggleWhitelist(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "toggleWhitelist", client)
}

// ToggleWhitelist is a paid mutator transaction binding the contract method 0x39393ac9.
//
// Solidity: function toggleWhitelist(address client) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) ToggleWhitelist(client common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.ToggleWhitelist(&_ZKMRStakeRegistry.TransactOpts, client)
}

// ToggleWhitelist is a paid mutator transaction binding the contract method 0x39393ac9.
//
// Solidity: function toggleWhitelist(address client) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) ToggleWhitelist(client common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.ToggleWhitelist(&_ZKMRStakeRegistry.TransactOpts, client)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.TransferOwnership(&_ZKMRStakeRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.TransferOwnership(&_ZKMRStakeRegistry.TransactOpts, newOwner)
}

// UpdateMinimumShares is a paid mutator transaction binding the contract method 0x0e50e175.
//
// Solidity: function updateMinimumShares(uint256 newMinimumShares) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) UpdateMinimumShares(opts *bind.TransactOpts, newMinimumShares *big.Int) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "updateMinimumShares", newMinimumShares)
}

// UpdateMinimumShares is a paid mutator transaction binding the contract method 0x0e50e175.
//
// Solidity: function updateMinimumShares(uint256 newMinimumShares) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) UpdateMinimumShares(newMinimumShares *big.Int) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.UpdateMinimumShares(&_ZKMRStakeRegistry.TransactOpts, newMinimumShares)
}

// UpdateMinimumShares is a paid mutator transaction binding the contract method 0x0e50e175.
//
// Solidity: function updateMinimumShares(uint256 newMinimumShares) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) UpdateMinimumShares(newMinimumShares *big.Int) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.UpdateMinimumShares(&_ZKMRStakeRegistry.TransactOpts, newMinimumShares)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0xf6b35715.
//
// Solidity: function updateOperatorKey((uint256,uint256) publicKey) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) UpdateOperatorKey(opts *bind.TransactOpts, publicKey PublicKey) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "updateOperatorKey", publicKey)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0xf6b35715.
//
// Solidity: function updateOperatorKey((uint256,uint256) publicKey) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) UpdateOperatorKey(publicKey PublicKey) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.UpdateOperatorKey(&_ZKMRStakeRegistry.TransactOpts, publicKey)
}

// UpdateOperatorKey is a paid mutator transaction binding the contract method 0xf6b35715.
//
// Solidity: function updateOperatorKey((uint256,uint256) publicKey) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) UpdateOperatorKey(publicKey PublicKey) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.UpdateOperatorKey(&_ZKMRStakeRegistry.TransactOpts, publicKey)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xe333c3ab.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) quorum_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactor) UpdateQuorumConfig(opts *bind.TransactOpts, quorum_ Quorum) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.contract.Transact(opts, "updateQuorumConfig", quorum_)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xe333c3ab.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) quorum_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistrySession) UpdateQuorumConfig(quorum_ Quorum) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.UpdateQuorumConfig(&_ZKMRStakeRegistry.TransactOpts, quorum_)
}

// UpdateQuorumConfig is a paid mutator transaction binding the contract method 0xe333c3ab.
//
// Solidity: function updateQuorumConfig(((address,uint96)[]) quorum_) returns()
func (_ZKMRStakeRegistry *ZKMRStakeRegistryTransactorSession) UpdateQuorumConfig(quorum_ Quorum) (*types.Transaction, error) {
	return _ZKMRStakeRegistry.Contract.UpdateQuorumConfig(&_ZKMRStakeRegistry.TransactOpts, quorum_)
}

// ZKMRStakeRegistryBatchWhitelistedIterator is returned from FilterBatchWhitelisted and is used to iterate over the raw logs and unpacked data for BatchWhitelisted events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryBatchWhitelistedIterator struct {
	Event *ZKMRStakeRegistryBatchWhitelisted // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryBatchWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryBatchWhitelisted)
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
		it.Event = new(ZKMRStakeRegistryBatchWhitelisted)
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
func (it *ZKMRStakeRegistryBatchWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryBatchWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryBatchWhitelisted represents a BatchWhitelisted event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryBatchWhitelisted struct {
	Addrs []common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBatchWhitelisted is a free log retrieval operation binding the contract event 0xe481b9d265afbda7824b6945408c17197396b0621a7e5418ac204f8c5abe6445.
//
// Solidity: event BatchWhitelisted(address[] addrs, bool value)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterBatchWhitelisted(opts *bind.FilterOpts) (*ZKMRStakeRegistryBatchWhitelistedIterator, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "BatchWhitelisted")
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryBatchWhitelistedIterator{contract: _ZKMRStakeRegistry.contract, event: "BatchWhitelisted", logs: logs, sub: sub}, nil
}

// WatchBatchWhitelisted is a free log subscription operation binding the contract event 0xe481b9d265afbda7824b6945408c17197396b0621a7e5418ac204f8c5abe6445.
//
// Solidity: event BatchWhitelisted(address[] addrs, bool value)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchBatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryBatchWhitelisted) (event.Subscription, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "BatchWhitelisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryBatchWhitelisted)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "BatchWhitelisted", log); err != nil {
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

// ParseBatchWhitelisted is a log parse operation binding the contract event 0xe481b9d265afbda7824b6945408c17197396b0621a7e5418ac204f8c5abe6445.
//
// Solidity: event BatchWhitelisted(address[] addrs, bool value)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseBatchWhitelisted(log types.Log) (*ZKMRStakeRegistryBatchWhitelisted, error) {
	event := new(ZKMRStakeRegistryBatchWhitelisted)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "BatchWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryInitializedIterator struct {
	Event *ZKMRStakeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryInitialized)
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
		it.Event = new(ZKMRStakeRegistryInitialized)
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
func (it *ZKMRStakeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryInitialized represents a Initialized event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ZKMRStakeRegistryInitializedIterator, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryInitializedIterator{contract: _ZKMRStakeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryInitialized)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseInitialized(log types.Log) (*ZKMRStakeRegistryInitialized, error) {
	event := new(ZKMRStakeRegistryInitialized)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryMinimumSharesUpdatedIterator is returned from FilterMinimumSharesUpdated and is used to iterate over the raw logs and unpacked data for MinimumSharesUpdated events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryMinimumSharesUpdatedIterator struct {
	Event *ZKMRStakeRegistryMinimumSharesUpdated // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryMinimumSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryMinimumSharesUpdated)
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
		it.Event = new(ZKMRStakeRegistryMinimumSharesUpdated)
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
func (it *ZKMRStakeRegistryMinimumSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryMinimumSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryMinimumSharesUpdated represents a MinimumSharesUpdated event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryMinimumSharesUpdated struct {
	OldShares *big.Int
	NewShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinimumSharesUpdated is a free log retrieval operation binding the contract event 0x5a347441ec3f4b995103b86e3f172999500a609ebed0757430ff8af75f2fbbfd.
//
// Solidity: event MinimumSharesUpdated(uint256 oldShares, uint256 newShares)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterMinimumSharesUpdated(opts *bind.FilterOpts) (*ZKMRStakeRegistryMinimumSharesUpdatedIterator, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "MinimumSharesUpdated")
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryMinimumSharesUpdatedIterator{contract: _ZKMRStakeRegistry.contract, event: "MinimumSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumSharesUpdated is a free log subscription operation binding the contract event 0x5a347441ec3f4b995103b86e3f172999500a609ebed0757430ff8af75f2fbbfd.
//
// Solidity: event MinimumSharesUpdated(uint256 oldShares, uint256 newShares)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchMinimumSharesUpdated(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryMinimumSharesUpdated) (event.Subscription, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "MinimumSharesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryMinimumSharesUpdated)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "MinimumSharesUpdated", log); err != nil {
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

// ParseMinimumSharesUpdated is a log parse operation binding the contract event 0x5a347441ec3f4b995103b86e3f172999500a609ebed0757430ff8af75f2fbbfd.
//
// Solidity: event MinimumSharesUpdated(uint256 oldShares, uint256 newShares)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseMinimumSharesUpdated(log types.Log) (*ZKMRStakeRegistryMinimumSharesUpdated, error) {
	event := new(ZKMRStakeRegistryMinimumSharesUpdated)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "MinimumSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorDeregisteredIterator struct {
	Event *ZKMRStakeRegistryOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryOperatorDeregistered)
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
		it.Event = new(ZKMRStakeRegistryOperatorDeregistered)
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
func (it *ZKMRStakeRegistryOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryOperatorDeregistered represents a OperatorDeregistered event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorDeregistered struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed operator, address indexed avs)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ZKMRStakeRegistryOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryOperatorDeregisteredIterator{contract: _ZKMRStakeRegistry.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x31e0adfec71bccee37b6e83a90c2fedb17d8f1693fee863c4771e7bfe2aed580.
//
// Solidity: event OperatorDeregistered(address indexed operator, address indexed avs)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryOperatorDeregistered, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryOperatorDeregistered)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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
// Solidity: event OperatorDeregistered(address indexed operator, address indexed avs)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseOperatorDeregistered(log types.Log) (*ZKMRStakeRegistryOperatorDeregistered, error) {
	event := new(ZKMRStakeRegistryOperatorDeregistered)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryOperatorEvictedIterator is returned from FilterOperatorEvicted and is used to iterate over the raw logs and unpacked data for OperatorEvicted events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorEvictedIterator struct {
	Event *ZKMRStakeRegistryOperatorEvicted // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryOperatorEvictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryOperatorEvicted)
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
		it.Event = new(ZKMRStakeRegistryOperatorEvicted)
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
func (it *ZKMRStakeRegistryOperatorEvictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryOperatorEvictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryOperatorEvicted represents a OperatorEvicted event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorEvicted struct {
	Operator common.Address
	Avs      common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorEvicted is a free log retrieval operation binding the contract event 0x0e870a9e10ce1e89472bdabc78c39947aebff13bcdfbae1759b4b9850755efcc.
//
// Solidity: event OperatorEvicted(address indexed operator, address indexed avs)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterOperatorEvicted(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ZKMRStakeRegistryOperatorEvictedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "OperatorEvicted", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryOperatorEvictedIterator{contract: _ZKMRStakeRegistry.contract, event: "OperatorEvicted", logs: logs, sub: sub}, nil
}

// WatchOperatorEvicted is a free log subscription operation binding the contract event 0x0e870a9e10ce1e89472bdabc78c39947aebff13bcdfbae1759b4b9850755efcc.
//
// Solidity: event OperatorEvicted(address indexed operator, address indexed avs)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchOperatorEvicted(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryOperatorEvicted, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "OperatorEvicted", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryOperatorEvicted)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorEvicted", log); err != nil {
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

// ParseOperatorEvicted is a log parse operation binding the contract event 0x0e870a9e10ce1e89472bdabc78c39947aebff13bcdfbae1759b4b9850755efcc.
//
// Solidity: event OperatorEvicted(address indexed operator, address indexed avs)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseOperatorEvicted(log types.Log) (*ZKMRStakeRegistryOperatorEvicted, error) {
	event := new(ZKMRStakeRegistryOperatorEvicted)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorEvicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorRegisteredIterator struct {
	Event *ZKMRStakeRegistryOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryOperatorRegistered)
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
		it.Event = new(ZKMRStakeRegistryOperatorRegistered)
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
func (it *ZKMRStakeRegistryOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryOperatorRegistered represents a OperatorRegistered event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorRegistered struct {
	Operator  common.Address
	Avs       common.Address
	PublicKey PublicKey
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4e7411f7b946b0536be9ba333b8250b3266710e75fdc438cc0b0be5e285eb42d.
//
// Solidity: event OperatorRegistered(address indexed operator, address indexed avs, (uint256,uint256) publicKey)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ZKMRStakeRegistryOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryOperatorRegisteredIterator{contract: _ZKMRStakeRegistry.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4e7411f7b946b0536be9ba333b8250b3266710e75fdc438cc0b0be5e285eb42d.
//
// Solidity: event OperatorRegistered(address indexed operator, address indexed avs, (uint256,uint256) publicKey)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryOperatorRegistered, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryOperatorRegistered)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x4e7411f7b946b0536be9ba333b8250b3266710e75fdc438cc0b0be5e285eb42d.
//
// Solidity: event OperatorRegistered(address indexed operator, address indexed avs, (uint256,uint256) publicKey)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseOperatorRegistered(log types.Log) (*ZKMRStakeRegistryOperatorRegistered, error) {
	event := new(ZKMRStakeRegistryOperatorRegistered)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryOperatorUpdatedIterator is returned from FilterOperatorUpdated and is used to iterate over the raw logs and unpacked data for OperatorUpdated events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorUpdatedIterator struct {
	Event *ZKMRStakeRegistryOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryOperatorUpdated)
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
		it.Event = new(ZKMRStakeRegistryOperatorUpdated)
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
func (it *ZKMRStakeRegistryOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryOperatorUpdated represents a OperatorUpdated event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOperatorUpdated struct {
	Operator  common.Address
	Avs       common.Address
	PublicKey PublicKey
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorUpdated is a free log retrieval operation binding the contract event 0x8947c0f26ae5515b92104474953e276cc5e976f5dc70d58b2e3f5857957a6ab9.
//
// Solidity: event OperatorUpdated(address indexed operator, address indexed avs, (uint256,uint256) publicKey)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterOperatorUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ZKMRStakeRegistryOperatorUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "OperatorUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryOperatorUpdatedIterator{contract: _ZKMRStakeRegistry.contract, event: "OperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorUpdated is a free log subscription operation binding the contract event 0x8947c0f26ae5515b92104474953e276cc5e976f5dc70d58b2e3f5857957a6ab9.
//
// Solidity: event OperatorUpdated(address indexed operator, address indexed avs, (uint256,uint256) publicKey)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchOperatorUpdated(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryOperatorUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "OperatorUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryOperatorUpdated)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
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

// ParseOperatorUpdated is a log parse operation binding the contract event 0x8947c0f26ae5515b92104474953e276cc5e976f5dc70d58b2e3f5857957a6ab9.
//
// Solidity: event OperatorUpdated(address indexed operator, address indexed avs, (uint256,uint256) publicKey)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseOperatorUpdated(log types.Log) (*ZKMRStakeRegistryOperatorUpdated, error) {
	event := new(ZKMRStakeRegistryOperatorUpdated)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOwnershipHandoverCanceledIterator struct {
	Event *ZKMRStakeRegistryOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryOwnershipHandoverCanceled)
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
		it.Event = new(ZKMRStakeRegistryOwnershipHandoverCanceled)
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
func (it *ZKMRStakeRegistryOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*ZKMRStakeRegistryOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryOwnershipHandoverCanceledIterator{contract: _ZKMRStakeRegistry.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryOwnershipHandoverCanceled)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*ZKMRStakeRegistryOwnershipHandoverCanceled, error) {
	event := new(ZKMRStakeRegistryOwnershipHandoverCanceled)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOwnershipHandoverRequestedIterator struct {
	Event *ZKMRStakeRegistryOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryOwnershipHandoverRequested)
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
		it.Event = new(ZKMRStakeRegistryOwnershipHandoverRequested)
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
func (it *ZKMRStakeRegistryOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*ZKMRStakeRegistryOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryOwnershipHandoverRequestedIterator{contract: _ZKMRStakeRegistry.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryOwnershipHandoverRequested)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseOwnershipHandoverRequested(log types.Log) (*ZKMRStakeRegistryOwnershipHandoverRequested, error) {
	event := new(ZKMRStakeRegistryOwnershipHandoverRequested)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOwnershipTransferredIterator struct {
	Event *ZKMRStakeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryOwnershipTransferred)
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
		it.Event = new(ZKMRStakeRegistryOwnershipTransferred)
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
func (it *ZKMRStakeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*ZKMRStakeRegistryOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryOwnershipTransferredIterator{contract: _ZKMRStakeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryOwnershipTransferred)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*ZKMRStakeRegistryOwnershipTransferred, error) {
	event := new(ZKMRStakeRegistryOwnershipTransferred)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryQuorumUpdatedIterator is returned from FilterQuorumUpdated and is used to iterate over the raw logs and unpacked data for QuorumUpdated events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryQuorumUpdatedIterator struct {
	Event *ZKMRStakeRegistryQuorumUpdated // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryQuorumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryQuorumUpdated)
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
		it.Event = new(ZKMRStakeRegistryQuorumUpdated)
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
func (it *ZKMRStakeRegistryQuorumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryQuorumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryQuorumUpdated represents a QuorumUpdated event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryQuorumUpdated struct {
	OldQuorum Quorum
	NewQuorum Quorum
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuorumUpdated is a free log retrieval operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) oldQuorum, ((address,uint96)[]) newQuorum)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterQuorumUpdated(opts *bind.FilterOpts) (*ZKMRStakeRegistryQuorumUpdatedIterator, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryQuorumUpdatedIterator{contract: _ZKMRStakeRegistry.contract, event: "QuorumUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumUpdated is a free log subscription operation binding the contract event 0x23aad4e61744ece164130aa415c1616e80136b0f0770e56589438b90b269265e.
//
// Solidity: event QuorumUpdated(((address,uint96)[]) oldQuorum, ((address,uint96)[]) newQuorum)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchQuorumUpdated(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryQuorumUpdated) (event.Subscription, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "QuorumUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryQuorumUpdated)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
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
// Solidity: event QuorumUpdated(((address,uint96)[]) oldQuorum, ((address,uint96)[]) newQuorum)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseQuorumUpdated(log types.Log) (*ZKMRStakeRegistryQuorumUpdated, error) {
	event := new(ZKMRStakeRegistryQuorumUpdated)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "QuorumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKMRStakeRegistryWhitelistedIterator is returned from FilterWhitelisted and is used to iterate over the raw logs and unpacked data for Whitelisted events raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryWhitelistedIterator struct {
	Event *ZKMRStakeRegistryWhitelisted // Event containing the contract specifics and raw log

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
func (it *ZKMRStakeRegistryWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKMRStakeRegistryWhitelisted)
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
		it.Event = new(ZKMRStakeRegistryWhitelisted)
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
func (it *ZKMRStakeRegistryWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKMRStakeRegistryWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKMRStakeRegistryWhitelisted represents a Whitelisted event raised by the ZKMRStakeRegistry contract.
type ZKMRStakeRegistryWhitelisted struct {
	Addr  common.Address
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelisted is a free log retrieval operation binding the contract event 0xa54714518c5d275fdcd3d2a461e4858e4e8cb04fb93cd0bca9d6d34115f26440.
//
// Solidity: event Whitelisted(address addr, bool value)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) FilterWhitelisted(opts *bind.FilterOpts) (*ZKMRStakeRegistryWhitelistedIterator, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.FilterLogs(opts, "Whitelisted")
	if err != nil {
		return nil, err
	}
	return &ZKMRStakeRegistryWhitelistedIterator{contract: _ZKMRStakeRegistry.contract, event: "Whitelisted", logs: logs, sub: sub}, nil
}

// WatchWhitelisted is a free log subscription operation binding the contract event 0xa54714518c5d275fdcd3d2a461e4858e4e8cb04fb93cd0bca9d6d34115f26440.
//
// Solidity: event Whitelisted(address addr, bool value)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) WatchWhitelisted(opts *bind.WatchOpts, sink chan<- *ZKMRStakeRegistryWhitelisted) (event.Subscription, error) {

	logs, sub, err := _ZKMRStakeRegistry.contract.WatchLogs(opts, "Whitelisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKMRStakeRegistryWhitelisted)
				if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "Whitelisted", log); err != nil {
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

// ParseWhitelisted is a log parse operation binding the contract event 0xa54714518c5d275fdcd3d2a461e4858e4e8cb04fb93cd0bca9d6d34115f26440.
//
// Solidity: event Whitelisted(address addr, bool value)
func (_ZKMRStakeRegistry *ZKMRStakeRegistryFilterer) ParseWhitelisted(log types.Log) (*ZKMRStakeRegistryWhitelisted, error) {
	event := new(ZKMRStakeRegistryWhitelisted)
	if err := _ZKMRStakeRegistry.contract.UnpackLog(event, "Whitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
