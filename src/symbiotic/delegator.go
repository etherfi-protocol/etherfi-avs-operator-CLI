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

// DelegatorMetaData contains all meta data concerning the Delegator contract.
var DelegatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"networkRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vaultFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorVaultOptInService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operatorNetworkOptInService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorFactory\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"entityType\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateRoleHolder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedsMaxNetworkLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientHookGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingRoleHolders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNetwork\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSlasher\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressRoleHolder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"captureTimestamp\",\"type\":\"uint48\"}],\"name\":\"OnSlash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"SetHook\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMaxNetworkLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetNetworkLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"SetOperatorNetworkShares\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOOK_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOOK_RESERVE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HOOK_SET_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETWORK_LIMIT_SET_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NETWORK_REGISTRY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_NETWORK_OPT_IN_SERVICE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_NETWORK_SHARES_SET_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_VAULT_OPT_IN_SERVICE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"}],\"name\":\"maxNetworkLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"}],\"name\":\"networkLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"networkLimitAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"captureTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"operatorNetworkShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"operatorNetworkSharesAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook_\",\"type\":\"address\"}],\"name\":\"setHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"identifier\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMaxNetworkLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setNetworkLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"setOperatorNetworkShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"hints\",\"type\":\"bytes\"}],\"name\":\"stakeAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"staticDelegateCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"}],\"name\":\"totalOperatorNetworkShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnetwork\",\"type\":\"bytes32\"},{\"internalType\":\"uint48\",\"name\":\"timestamp\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"hint\",\"type\":\"bytes\"}],\"name\":\"totalOperatorNetworkSharesAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DelegatorABI is the input ABI used to generate the binding from.
// Deprecated: Use DelegatorMetaData.ABI instead.
var DelegatorABI = DelegatorMetaData.ABI

// Delegator is an auto generated Go binding around an Ethereum contract.
type Delegator struct {
	DelegatorCaller     // Read-only binding to the contract
	DelegatorTransactor // Write-only binding to the contract
	DelegatorFilterer   // Log filterer for contract events
}

// DelegatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegatorSession struct {
	Contract     *Delegator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegatorCallerSession struct {
	Contract *DelegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DelegatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegatorTransactorSession struct {
	Contract     *DelegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DelegatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelegatorRaw struct {
	Contract *Delegator // Generic contract binding to access the raw methods on
}

// DelegatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegatorCallerRaw struct {
	Contract *DelegatorCaller // Generic read-only contract binding to access the raw methods on
}

// DelegatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegatorTransactorRaw struct {
	Contract *DelegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegator creates a new instance of Delegator, bound to a specific deployed contract.
func NewDelegator(address common.Address, backend bind.ContractBackend) (*Delegator, error) {
	contract, err := bindDelegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Delegator{DelegatorCaller: DelegatorCaller{contract: contract}, DelegatorTransactor: DelegatorTransactor{contract: contract}, DelegatorFilterer: DelegatorFilterer{contract: contract}}, nil
}

// NewDelegatorCaller creates a new read-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorCaller(address common.Address, caller bind.ContractCaller) (*DelegatorCaller, error) {
	contract, err := bindDelegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorCaller{contract: contract}, nil
}

// NewDelegatorTransactor creates a new write-only instance of Delegator, bound to a specific deployed contract.
func NewDelegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*DelegatorTransactor, error) {
	contract, err := bindDelegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegatorTransactor{contract: contract}, nil
}

// NewDelegatorFilterer creates a new log filterer instance of Delegator, bound to a specific deployed contract.
func NewDelegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*DelegatorFilterer, error) {
	contract, err := bindDelegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegatorFilterer{contract: contract}, nil
}

// bindDelegator binds a generic wrapper to an already deployed contract.
func bindDelegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DelegatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.DelegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.DelegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Delegator *DelegatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Delegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Delegator *DelegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Delegator *DelegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Delegator.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Delegator *DelegatorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Delegator.Contract.DEFAULTADMINROLE(&_Delegator.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Delegator.Contract.DEFAULTADMINROLE(&_Delegator.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Delegator *DelegatorCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Delegator *DelegatorSession) FACTORY() (common.Address, error) {
	return _Delegator.Contract.FACTORY(&_Delegator.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Delegator *DelegatorCallerSession) FACTORY() (common.Address, error) {
	return _Delegator.Contract.FACTORY(&_Delegator.CallOpts)
}

// HOOKGASLIMIT is a free data retrieval call binding the contract method 0xff54740f.
//
// Solidity: function HOOK_GAS_LIMIT() view returns(uint256)
func (_Delegator *DelegatorCaller) HOOKGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "HOOK_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HOOKGASLIMIT is a free data retrieval call binding the contract method 0xff54740f.
//
// Solidity: function HOOK_GAS_LIMIT() view returns(uint256)
func (_Delegator *DelegatorSession) HOOKGASLIMIT() (*big.Int, error) {
	return _Delegator.Contract.HOOKGASLIMIT(&_Delegator.CallOpts)
}

// HOOKGASLIMIT is a free data retrieval call binding the contract method 0xff54740f.
//
// Solidity: function HOOK_GAS_LIMIT() view returns(uint256)
func (_Delegator *DelegatorCallerSession) HOOKGASLIMIT() (*big.Int, error) {
	return _Delegator.Contract.HOOKGASLIMIT(&_Delegator.CallOpts)
}

// HOOKRESERVE is a free data retrieval call binding the contract method 0x557cab44.
//
// Solidity: function HOOK_RESERVE() view returns(uint256)
func (_Delegator *DelegatorCaller) HOOKRESERVE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "HOOK_RESERVE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HOOKRESERVE is a free data retrieval call binding the contract method 0x557cab44.
//
// Solidity: function HOOK_RESERVE() view returns(uint256)
func (_Delegator *DelegatorSession) HOOKRESERVE() (*big.Int, error) {
	return _Delegator.Contract.HOOKRESERVE(&_Delegator.CallOpts)
}

// HOOKRESERVE is a free data retrieval call binding the contract method 0x557cab44.
//
// Solidity: function HOOK_RESERVE() view returns(uint256)
func (_Delegator *DelegatorCallerSession) HOOKRESERVE() (*big.Int, error) {
	return _Delegator.Contract.HOOKRESERVE(&_Delegator.CallOpts)
}

// HOOKSETROLE is a free data retrieval call binding the contract method 0x6679191e.
//
// Solidity: function HOOK_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCaller) HOOKSETROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "HOOK_SET_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HOOKSETROLE is a free data retrieval call binding the contract method 0x6679191e.
//
// Solidity: function HOOK_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorSession) HOOKSETROLE() ([32]byte, error) {
	return _Delegator.Contract.HOOKSETROLE(&_Delegator.CallOpts)
}

// HOOKSETROLE is a free data retrieval call binding the contract method 0x6679191e.
//
// Solidity: function HOOK_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCallerSession) HOOKSETROLE() ([32]byte, error) {
	return _Delegator.Contract.HOOKSETROLE(&_Delegator.CallOpts)
}

// NETWORKLIMITSETROLE is a free data retrieval call binding the contract method 0x7d24bb27.
//
// Solidity: function NETWORK_LIMIT_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCaller) NETWORKLIMITSETROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "NETWORK_LIMIT_SET_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NETWORKLIMITSETROLE is a free data retrieval call binding the contract method 0x7d24bb27.
//
// Solidity: function NETWORK_LIMIT_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorSession) NETWORKLIMITSETROLE() ([32]byte, error) {
	return _Delegator.Contract.NETWORKLIMITSETROLE(&_Delegator.CallOpts)
}

// NETWORKLIMITSETROLE is a free data retrieval call binding the contract method 0x7d24bb27.
//
// Solidity: function NETWORK_LIMIT_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCallerSession) NETWORKLIMITSETROLE() ([32]byte, error) {
	return _Delegator.Contract.NETWORKLIMITSETROLE(&_Delegator.CallOpts)
}

// NETWORKREGISTRY is a free data retrieval call binding the contract method 0xc0cd7c3e.
//
// Solidity: function NETWORK_REGISTRY() view returns(address)
func (_Delegator *DelegatorCaller) NETWORKREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "NETWORK_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NETWORKREGISTRY is a free data retrieval call binding the contract method 0xc0cd7c3e.
//
// Solidity: function NETWORK_REGISTRY() view returns(address)
func (_Delegator *DelegatorSession) NETWORKREGISTRY() (common.Address, error) {
	return _Delegator.Contract.NETWORKREGISTRY(&_Delegator.CallOpts)
}

// NETWORKREGISTRY is a free data retrieval call binding the contract method 0xc0cd7c3e.
//
// Solidity: function NETWORK_REGISTRY() view returns(address)
func (_Delegator *DelegatorCallerSession) NETWORKREGISTRY() (common.Address, error) {
	return _Delegator.Contract.NETWORKREGISTRY(&_Delegator.CallOpts)
}

// OPERATORNETWORKOPTINSERVICE is a free data retrieval call binding the contract method 0x1a80e500.
//
// Solidity: function OPERATOR_NETWORK_OPT_IN_SERVICE() view returns(address)
func (_Delegator *DelegatorCaller) OPERATORNETWORKOPTINSERVICE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "OPERATOR_NETWORK_OPT_IN_SERVICE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATORNETWORKOPTINSERVICE is a free data retrieval call binding the contract method 0x1a80e500.
//
// Solidity: function OPERATOR_NETWORK_OPT_IN_SERVICE() view returns(address)
func (_Delegator *DelegatorSession) OPERATORNETWORKOPTINSERVICE() (common.Address, error) {
	return _Delegator.Contract.OPERATORNETWORKOPTINSERVICE(&_Delegator.CallOpts)
}

// OPERATORNETWORKOPTINSERVICE is a free data retrieval call binding the contract method 0x1a80e500.
//
// Solidity: function OPERATOR_NETWORK_OPT_IN_SERVICE() view returns(address)
func (_Delegator *DelegatorCallerSession) OPERATORNETWORKOPTINSERVICE() (common.Address, error) {
	return _Delegator.Contract.OPERATORNETWORKOPTINSERVICE(&_Delegator.CallOpts)
}

// OPERATORNETWORKSHARESSETROLE is a free data retrieval call binding the contract method 0xe78eb6ae.
//
// Solidity: function OPERATOR_NETWORK_SHARES_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCaller) OPERATORNETWORKSHARESSETROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "OPERATOR_NETWORK_SHARES_SET_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORNETWORKSHARESSETROLE is a free data retrieval call binding the contract method 0xe78eb6ae.
//
// Solidity: function OPERATOR_NETWORK_SHARES_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorSession) OPERATORNETWORKSHARESSETROLE() ([32]byte, error) {
	return _Delegator.Contract.OPERATORNETWORKSHARESSETROLE(&_Delegator.CallOpts)
}

// OPERATORNETWORKSHARESSETROLE is a free data retrieval call binding the contract method 0xe78eb6ae.
//
// Solidity: function OPERATOR_NETWORK_SHARES_SET_ROLE() view returns(bytes32)
func (_Delegator *DelegatorCallerSession) OPERATORNETWORKSHARESSETROLE() ([32]byte, error) {
	return _Delegator.Contract.OPERATORNETWORKSHARESSETROLE(&_Delegator.CallOpts)
}

// OPERATORVAULTOPTINSERVICE is a free data retrieval call binding the contract method 0x128e5d82.
//
// Solidity: function OPERATOR_VAULT_OPT_IN_SERVICE() view returns(address)
func (_Delegator *DelegatorCaller) OPERATORVAULTOPTINSERVICE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "OPERATOR_VAULT_OPT_IN_SERVICE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATORVAULTOPTINSERVICE is a free data retrieval call binding the contract method 0x128e5d82.
//
// Solidity: function OPERATOR_VAULT_OPT_IN_SERVICE() view returns(address)
func (_Delegator *DelegatorSession) OPERATORVAULTOPTINSERVICE() (common.Address, error) {
	return _Delegator.Contract.OPERATORVAULTOPTINSERVICE(&_Delegator.CallOpts)
}

// OPERATORVAULTOPTINSERVICE is a free data retrieval call binding the contract method 0x128e5d82.
//
// Solidity: function OPERATOR_VAULT_OPT_IN_SERVICE() view returns(address)
func (_Delegator *DelegatorCallerSession) OPERATORVAULTOPTINSERVICE() (common.Address, error) {
	return _Delegator.Contract.OPERATORVAULTOPTINSERVICE(&_Delegator.CallOpts)
}

// TYPE is a free data retrieval call binding the contract method 0xbb24fe8a.
//
// Solidity: function TYPE() view returns(uint64)
func (_Delegator *DelegatorCaller) TYPE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "TYPE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TYPE is a free data retrieval call binding the contract method 0xbb24fe8a.
//
// Solidity: function TYPE() view returns(uint64)
func (_Delegator *DelegatorSession) TYPE() (uint64, error) {
	return _Delegator.Contract.TYPE(&_Delegator.CallOpts)
}

// TYPE is a free data retrieval call binding the contract method 0xbb24fe8a.
//
// Solidity: function TYPE() view returns(uint64)
func (_Delegator *DelegatorCallerSession) TYPE() (uint64, error) {
	return _Delegator.Contract.TYPE(&_Delegator.CallOpts)
}

// VAULTFACTORY is a free data retrieval call binding the contract method 0x103f2907.
//
// Solidity: function VAULT_FACTORY() view returns(address)
func (_Delegator *DelegatorCaller) VAULTFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "VAULT_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VAULTFACTORY is a free data retrieval call binding the contract method 0x103f2907.
//
// Solidity: function VAULT_FACTORY() view returns(address)
func (_Delegator *DelegatorSession) VAULTFACTORY() (common.Address, error) {
	return _Delegator.Contract.VAULTFACTORY(&_Delegator.CallOpts)
}

// VAULTFACTORY is a free data retrieval call binding the contract method 0x103f2907.
//
// Solidity: function VAULT_FACTORY() view returns(address)
func (_Delegator *DelegatorCallerSession) VAULTFACTORY() (common.Address, error) {
	return _Delegator.Contract.VAULTFACTORY(&_Delegator.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(uint64)
func (_Delegator *DelegatorCaller) VERSION(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(uint64)
func (_Delegator *DelegatorSession) VERSION() (uint64, error) {
	return _Delegator.Contract.VERSION(&_Delegator.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() pure returns(uint64)
func (_Delegator *DelegatorCallerSession) VERSION() (uint64, error) {
	return _Delegator.Contract.VERSION(&_Delegator.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Delegator *DelegatorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Delegator *DelegatorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Delegator.Contract.GetRoleAdmin(&_Delegator.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Delegator *DelegatorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Delegator.Contract.GetRoleAdmin(&_Delegator.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Delegator *DelegatorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Delegator *DelegatorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Delegator.Contract.HasRole(&_Delegator.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Delegator *DelegatorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Delegator.Contract.HasRole(&_Delegator.CallOpts, role, account)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_Delegator *DelegatorCaller) Hook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "hook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_Delegator *DelegatorSession) Hook() (common.Address, error) {
	return _Delegator.Contract.Hook(&_Delegator.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_Delegator *DelegatorCallerSession) Hook() (common.Address, error) {
	return _Delegator.Contract.Hook(&_Delegator.CallOpts)
}

// MaxNetworkLimit is a free data retrieval call binding the contract method 0xd15b740e.
//
// Solidity: function maxNetworkLimit(bytes32 subnetwork) view returns(uint256 value)
func (_Delegator *DelegatorCaller) MaxNetworkLimit(opts *bind.CallOpts, subnetwork [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "maxNetworkLimit", subnetwork)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNetworkLimit is a free data retrieval call binding the contract method 0xd15b740e.
//
// Solidity: function maxNetworkLimit(bytes32 subnetwork) view returns(uint256 value)
func (_Delegator *DelegatorSession) MaxNetworkLimit(subnetwork [32]byte) (*big.Int, error) {
	return _Delegator.Contract.MaxNetworkLimit(&_Delegator.CallOpts, subnetwork)
}

// MaxNetworkLimit is a free data retrieval call binding the contract method 0xd15b740e.
//
// Solidity: function maxNetworkLimit(bytes32 subnetwork) view returns(uint256 value)
func (_Delegator *DelegatorCallerSession) MaxNetworkLimit(subnetwork [32]byte) (*big.Int, error) {
	return _Delegator.Contract.MaxNetworkLimit(&_Delegator.CallOpts, subnetwork)
}

// NetworkLimit is a free data retrieval call binding the contract method 0x3eb22c0f.
//
// Solidity: function networkLimit(bytes32 subnetwork) view returns(uint256)
func (_Delegator *DelegatorCaller) NetworkLimit(opts *bind.CallOpts, subnetwork [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "networkLimit", subnetwork)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NetworkLimit is a free data retrieval call binding the contract method 0x3eb22c0f.
//
// Solidity: function networkLimit(bytes32 subnetwork) view returns(uint256)
func (_Delegator *DelegatorSession) NetworkLimit(subnetwork [32]byte) (*big.Int, error) {
	return _Delegator.Contract.NetworkLimit(&_Delegator.CallOpts, subnetwork)
}

// NetworkLimit is a free data retrieval call binding the contract method 0x3eb22c0f.
//
// Solidity: function networkLimit(bytes32 subnetwork) view returns(uint256)
func (_Delegator *DelegatorCallerSession) NetworkLimit(subnetwork [32]byte) (*big.Int, error) {
	return _Delegator.Contract.NetworkLimit(&_Delegator.CallOpts, subnetwork)
}

// NetworkLimitAt is a free data retrieval call binding the contract method 0x5d32a1c9.
//
// Solidity: function networkLimitAt(bytes32 subnetwork, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorCaller) NetworkLimitAt(opts *bind.CallOpts, subnetwork [32]byte, timestamp *big.Int, hint []byte) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "networkLimitAt", subnetwork, timestamp, hint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NetworkLimitAt is a free data retrieval call binding the contract method 0x5d32a1c9.
//
// Solidity: function networkLimitAt(bytes32 subnetwork, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorSession) NetworkLimitAt(subnetwork [32]byte, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _Delegator.Contract.NetworkLimitAt(&_Delegator.CallOpts, subnetwork, timestamp, hint)
}

// NetworkLimitAt is a free data retrieval call binding the contract method 0x5d32a1c9.
//
// Solidity: function networkLimitAt(bytes32 subnetwork, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorCallerSession) NetworkLimitAt(subnetwork [32]byte, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _Delegator.Contract.NetworkLimitAt(&_Delegator.CallOpts, subnetwork, timestamp, hint)
}

// OperatorNetworkShares is a free data retrieval call binding the contract method 0x42c53e33.
//
// Solidity: function operatorNetworkShares(bytes32 subnetwork, address operator) view returns(uint256)
func (_Delegator *DelegatorCaller) OperatorNetworkShares(opts *bind.CallOpts, subnetwork [32]byte, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "operatorNetworkShares", subnetwork, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorNetworkShares is a free data retrieval call binding the contract method 0x42c53e33.
//
// Solidity: function operatorNetworkShares(bytes32 subnetwork, address operator) view returns(uint256)
func (_Delegator *DelegatorSession) OperatorNetworkShares(subnetwork [32]byte, operator common.Address) (*big.Int, error) {
	return _Delegator.Contract.OperatorNetworkShares(&_Delegator.CallOpts, subnetwork, operator)
}

// OperatorNetworkShares is a free data retrieval call binding the contract method 0x42c53e33.
//
// Solidity: function operatorNetworkShares(bytes32 subnetwork, address operator) view returns(uint256)
func (_Delegator *DelegatorCallerSession) OperatorNetworkShares(subnetwork [32]byte, operator common.Address) (*big.Int, error) {
	return _Delegator.Contract.OperatorNetworkShares(&_Delegator.CallOpts, subnetwork, operator)
}

// OperatorNetworkSharesAt is a free data retrieval call binding the contract method 0x1a7a7044.
//
// Solidity: function operatorNetworkSharesAt(bytes32 subnetwork, address operator, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorCaller) OperatorNetworkSharesAt(opts *bind.CallOpts, subnetwork [32]byte, operator common.Address, timestamp *big.Int, hint []byte) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "operatorNetworkSharesAt", subnetwork, operator, timestamp, hint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorNetworkSharesAt is a free data retrieval call binding the contract method 0x1a7a7044.
//
// Solidity: function operatorNetworkSharesAt(bytes32 subnetwork, address operator, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorSession) OperatorNetworkSharesAt(subnetwork [32]byte, operator common.Address, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _Delegator.Contract.OperatorNetworkSharesAt(&_Delegator.CallOpts, subnetwork, operator, timestamp, hint)
}

// OperatorNetworkSharesAt is a free data retrieval call binding the contract method 0x1a7a7044.
//
// Solidity: function operatorNetworkSharesAt(bytes32 subnetwork, address operator, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorCallerSession) OperatorNetworkSharesAt(subnetwork [32]byte, operator common.Address, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _Delegator.Contract.OperatorNetworkSharesAt(&_Delegator.CallOpts, subnetwork, operator, timestamp, hint)
}

// Stake is a free data retrieval call binding the contract method 0xfd4d447c.
//
// Solidity: function stake(bytes32 subnetwork, address operator) view returns(uint256)
func (_Delegator *DelegatorCaller) Stake(opts *bind.CallOpts, subnetwork [32]byte, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "stake", subnetwork, operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stake is a free data retrieval call binding the contract method 0xfd4d447c.
//
// Solidity: function stake(bytes32 subnetwork, address operator) view returns(uint256)
func (_Delegator *DelegatorSession) Stake(subnetwork [32]byte, operator common.Address) (*big.Int, error) {
	return _Delegator.Contract.Stake(&_Delegator.CallOpts, subnetwork, operator)
}

// Stake is a free data retrieval call binding the contract method 0xfd4d447c.
//
// Solidity: function stake(bytes32 subnetwork, address operator) view returns(uint256)
func (_Delegator *DelegatorCallerSession) Stake(subnetwork [32]byte, operator common.Address) (*big.Int, error) {
	return _Delegator.Contract.Stake(&_Delegator.CallOpts, subnetwork, operator)
}

// StakeAt is a free data retrieval call binding the contract method 0xe02f6937.
//
// Solidity: function stakeAt(bytes32 subnetwork, address operator, uint48 timestamp, bytes hints) view returns(uint256)
func (_Delegator *DelegatorCaller) StakeAt(opts *bind.CallOpts, subnetwork [32]byte, operator common.Address, timestamp *big.Int, hints []byte) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "stakeAt", subnetwork, operator, timestamp, hints)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAt is a free data retrieval call binding the contract method 0xe02f6937.
//
// Solidity: function stakeAt(bytes32 subnetwork, address operator, uint48 timestamp, bytes hints) view returns(uint256)
func (_Delegator *DelegatorSession) StakeAt(subnetwork [32]byte, operator common.Address, timestamp *big.Int, hints []byte) (*big.Int, error) {
	return _Delegator.Contract.StakeAt(&_Delegator.CallOpts, subnetwork, operator, timestamp, hints)
}

// StakeAt is a free data retrieval call binding the contract method 0xe02f6937.
//
// Solidity: function stakeAt(bytes32 subnetwork, address operator, uint48 timestamp, bytes hints) view returns(uint256)
func (_Delegator *DelegatorCallerSession) StakeAt(subnetwork [32]byte, operator common.Address, timestamp *big.Int, hints []byte) (*big.Int, error) {
	return _Delegator.Contract.StakeAt(&_Delegator.CallOpts, subnetwork, operator, timestamp, hints)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delegator *DelegatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delegator *DelegatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Delegator.Contract.SupportsInterface(&_Delegator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Delegator *DelegatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Delegator.Contract.SupportsInterface(&_Delegator.CallOpts, interfaceId)
}

// TotalOperatorNetworkShares is a free data retrieval call binding the contract method 0xc43dc03f.
//
// Solidity: function totalOperatorNetworkShares(bytes32 subnetwork) view returns(uint256)
func (_Delegator *DelegatorCaller) TotalOperatorNetworkShares(opts *bind.CallOpts, subnetwork [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "totalOperatorNetworkShares", subnetwork)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOperatorNetworkShares is a free data retrieval call binding the contract method 0xc43dc03f.
//
// Solidity: function totalOperatorNetworkShares(bytes32 subnetwork) view returns(uint256)
func (_Delegator *DelegatorSession) TotalOperatorNetworkShares(subnetwork [32]byte) (*big.Int, error) {
	return _Delegator.Contract.TotalOperatorNetworkShares(&_Delegator.CallOpts, subnetwork)
}

// TotalOperatorNetworkShares is a free data retrieval call binding the contract method 0xc43dc03f.
//
// Solidity: function totalOperatorNetworkShares(bytes32 subnetwork) view returns(uint256)
func (_Delegator *DelegatorCallerSession) TotalOperatorNetworkShares(subnetwork [32]byte) (*big.Int, error) {
	return _Delegator.Contract.TotalOperatorNetworkShares(&_Delegator.CallOpts, subnetwork)
}

// TotalOperatorNetworkSharesAt is a free data retrieval call binding the contract method 0x8b3f10b0.
//
// Solidity: function totalOperatorNetworkSharesAt(bytes32 subnetwork, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorCaller) TotalOperatorNetworkSharesAt(opts *bind.CallOpts, subnetwork [32]byte, timestamp *big.Int, hint []byte) (*big.Int, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "totalOperatorNetworkSharesAt", subnetwork, timestamp, hint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOperatorNetworkSharesAt is a free data retrieval call binding the contract method 0x8b3f10b0.
//
// Solidity: function totalOperatorNetworkSharesAt(bytes32 subnetwork, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorSession) TotalOperatorNetworkSharesAt(subnetwork [32]byte, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _Delegator.Contract.TotalOperatorNetworkSharesAt(&_Delegator.CallOpts, subnetwork, timestamp, hint)
}

// TotalOperatorNetworkSharesAt is a free data retrieval call binding the contract method 0x8b3f10b0.
//
// Solidity: function totalOperatorNetworkSharesAt(bytes32 subnetwork, uint48 timestamp, bytes hint) view returns(uint256)
func (_Delegator *DelegatorCallerSession) TotalOperatorNetworkSharesAt(subnetwork [32]byte, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _Delegator.Contract.TotalOperatorNetworkSharesAt(&_Delegator.CallOpts, subnetwork, timestamp, hint)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Delegator *DelegatorCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Delegator.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Delegator *DelegatorSession) Vault() (common.Address, error) {
	return _Delegator.Contract.Vault(&_Delegator.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Delegator *DelegatorCallerSession) Vault() (common.Address, error) {
	return _Delegator.Contract.Vault(&_Delegator.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Delegator *DelegatorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Delegator *DelegatorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.GrantRole(&_Delegator.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Delegator *DelegatorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.GrantRole(&_Delegator.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_Delegator *DelegatorTransactor) Initialize(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "initialize", data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_Delegator *DelegatorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.Initialize(&_Delegator.TransactOpts, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes data) returns()
func (_Delegator *DelegatorTransactorSession) Initialize(data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.Initialize(&_Delegator.TransactOpts, data)
}

// OnSlash is a paid mutator transaction binding the contract method 0xe49561ee.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp, bytes data) returns()
func (_Delegator *DelegatorTransactor) OnSlash(opts *bind.TransactOpts, subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "onSlash", subnetwork, operator, amount, captureTimestamp, data)
}

// OnSlash is a paid mutator transaction binding the contract method 0xe49561ee.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp, bytes data) returns()
func (_Delegator *DelegatorSession) OnSlash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.OnSlash(&_Delegator.TransactOpts, subnetwork, operator, amount, captureTimestamp, data)
}

// OnSlash is a paid mutator transaction binding the contract method 0xe49561ee.
//
// Solidity: function onSlash(bytes32 subnetwork, address operator, uint256 amount, uint48 captureTimestamp, bytes data) returns()
func (_Delegator *DelegatorTransactorSession) OnSlash(subnetwork [32]byte, operator common.Address, amount *big.Int, captureTimestamp *big.Int, data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.OnSlash(&_Delegator.TransactOpts, subnetwork, operator, amount, captureTimestamp, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Delegator *DelegatorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Delegator *DelegatorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.RenounceRole(&_Delegator.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Delegator *DelegatorTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.RenounceRole(&_Delegator.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Delegator *DelegatorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Delegator *DelegatorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.RevokeRole(&_Delegator.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Delegator *DelegatorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.RevokeRole(&_Delegator.TransactOpts, role, account)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address hook_) returns()
func (_Delegator *DelegatorTransactor) SetHook(opts *bind.TransactOpts, hook_ common.Address) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "setHook", hook_)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address hook_) returns()
func (_Delegator *DelegatorSession) SetHook(hook_ common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.SetHook(&_Delegator.TransactOpts, hook_)
}

// SetHook is a paid mutator transaction binding the contract method 0x3dfd3873.
//
// Solidity: function setHook(address hook_) returns()
func (_Delegator *DelegatorTransactorSession) SetHook(hook_ common.Address) (*types.Transaction, error) {
	return _Delegator.Contract.SetHook(&_Delegator.TransactOpts, hook_)
}

// SetMaxNetworkLimit is a paid mutator transaction binding the contract method 0x23f752d5.
//
// Solidity: function setMaxNetworkLimit(uint96 identifier, uint256 amount) returns()
func (_Delegator *DelegatorTransactor) SetMaxNetworkLimit(opts *bind.TransactOpts, identifier *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "setMaxNetworkLimit", identifier, amount)
}

// SetMaxNetworkLimit is a paid mutator transaction binding the contract method 0x23f752d5.
//
// Solidity: function setMaxNetworkLimit(uint96 identifier, uint256 amount) returns()
func (_Delegator *DelegatorSession) SetMaxNetworkLimit(identifier *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Delegator.Contract.SetMaxNetworkLimit(&_Delegator.TransactOpts, identifier, amount)
}

// SetMaxNetworkLimit is a paid mutator transaction binding the contract method 0x23f752d5.
//
// Solidity: function setMaxNetworkLimit(uint96 identifier, uint256 amount) returns()
func (_Delegator *DelegatorTransactorSession) SetMaxNetworkLimit(identifier *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Delegator.Contract.SetMaxNetworkLimit(&_Delegator.TransactOpts, identifier, amount)
}

// SetNetworkLimit is a paid mutator transaction binding the contract method 0x02145348.
//
// Solidity: function setNetworkLimit(bytes32 subnetwork, uint256 amount) returns()
func (_Delegator *DelegatorTransactor) SetNetworkLimit(opts *bind.TransactOpts, subnetwork [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "setNetworkLimit", subnetwork, amount)
}

// SetNetworkLimit is a paid mutator transaction binding the contract method 0x02145348.
//
// Solidity: function setNetworkLimit(bytes32 subnetwork, uint256 amount) returns()
func (_Delegator *DelegatorSession) SetNetworkLimit(subnetwork [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Delegator.Contract.SetNetworkLimit(&_Delegator.TransactOpts, subnetwork, amount)
}

// SetNetworkLimit is a paid mutator transaction binding the contract method 0x02145348.
//
// Solidity: function setNetworkLimit(bytes32 subnetwork, uint256 amount) returns()
func (_Delegator *DelegatorTransactorSession) SetNetworkLimit(subnetwork [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Delegator.Contract.SetNetworkLimit(&_Delegator.TransactOpts, subnetwork, amount)
}

// SetOperatorNetworkShares is a paid mutator transaction binding the contract method 0xa33bc287.
//
// Solidity: function setOperatorNetworkShares(bytes32 subnetwork, address operator, uint256 shares) returns()
func (_Delegator *DelegatorTransactor) SetOperatorNetworkShares(opts *bind.TransactOpts, subnetwork [32]byte, operator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "setOperatorNetworkShares", subnetwork, operator, shares)
}

// SetOperatorNetworkShares is a paid mutator transaction binding the contract method 0xa33bc287.
//
// Solidity: function setOperatorNetworkShares(bytes32 subnetwork, address operator, uint256 shares) returns()
func (_Delegator *DelegatorSession) SetOperatorNetworkShares(subnetwork [32]byte, operator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Delegator.Contract.SetOperatorNetworkShares(&_Delegator.TransactOpts, subnetwork, operator, shares)
}

// SetOperatorNetworkShares is a paid mutator transaction binding the contract method 0xa33bc287.
//
// Solidity: function setOperatorNetworkShares(bytes32 subnetwork, address operator, uint256 shares) returns()
func (_Delegator *DelegatorTransactorSession) SetOperatorNetworkShares(subnetwork [32]byte, operator common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Delegator.Contract.SetOperatorNetworkShares(&_Delegator.TransactOpts, subnetwork, operator, shares)
}

// StaticDelegateCall is a paid mutator transaction binding the contract method 0x9f86fd85.
//
// Solidity: function staticDelegateCall(address target, bytes data) returns()
func (_Delegator *DelegatorTransactor) StaticDelegateCall(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _Delegator.contract.Transact(opts, "staticDelegateCall", target, data)
}

// StaticDelegateCall is a paid mutator transaction binding the contract method 0x9f86fd85.
//
// Solidity: function staticDelegateCall(address target, bytes data) returns()
func (_Delegator *DelegatorSession) StaticDelegateCall(target common.Address, data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.StaticDelegateCall(&_Delegator.TransactOpts, target, data)
}

// StaticDelegateCall is a paid mutator transaction binding the contract method 0x9f86fd85.
//
// Solidity: function staticDelegateCall(address target, bytes data) returns()
func (_Delegator *DelegatorTransactorSession) StaticDelegateCall(target common.Address, data []byte) (*types.Transaction, error) {
	return _Delegator.Contract.StaticDelegateCall(&_Delegator.TransactOpts, target, data)
}

// DelegatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Delegator contract.
type DelegatorInitializedIterator struct {
	Event *DelegatorInitialized // Event containing the contract specifics and raw log

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
func (it *DelegatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorInitialized)
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
		it.Event = new(DelegatorInitialized)
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
func (it *DelegatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorInitialized represents a Initialized event raised by the Delegator contract.
type DelegatorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Delegator *DelegatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*DelegatorInitializedIterator, error) {

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DelegatorInitializedIterator{contract: _Delegator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Delegator *DelegatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DelegatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorInitialized)
				if err := _Delegator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Delegator *DelegatorFilterer) ParseInitialized(log types.Log) (*DelegatorInitialized, error) {
	event := new(DelegatorInitialized)
	if err := _Delegator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorOnSlashIterator is returned from FilterOnSlash and is used to iterate over the raw logs and unpacked data for OnSlash events raised by the Delegator contract.
type DelegatorOnSlashIterator struct {
	Event *DelegatorOnSlash // Event containing the contract specifics and raw log

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
func (it *DelegatorOnSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorOnSlash)
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
		it.Event = new(DelegatorOnSlash)
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
func (it *DelegatorOnSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorOnSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorOnSlash represents a OnSlash event raised by the Delegator contract.
type DelegatorOnSlash struct {
	Subnetwork       [32]byte
	Operator         common.Address
	Amount           *big.Int
	CaptureTimestamp *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnSlash is a free log retrieval operation binding the contract event 0x741a5de99085c0d660f3e4192217b0ffb0ea4e35a0480de48e857a4bc3ee36ed.
//
// Solidity: event OnSlash(bytes32 indexed subnetwork, address indexed operator, uint256 amount, uint48 captureTimestamp)
func (_Delegator *DelegatorFilterer) FilterOnSlash(opts *bind.FilterOpts, subnetwork [][32]byte, operator []common.Address) (*DelegatorOnSlashIterator, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "OnSlash", subnetworkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorOnSlashIterator{contract: _Delegator.contract, event: "OnSlash", logs: logs, sub: sub}, nil
}

// WatchOnSlash is a free log subscription operation binding the contract event 0x741a5de99085c0d660f3e4192217b0ffb0ea4e35a0480de48e857a4bc3ee36ed.
//
// Solidity: event OnSlash(bytes32 indexed subnetwork, address indexed operator, uint256 amount, uint48 captureTimestamp)
func (_Delegator *DelegatorFilterer) WatchOnSlash(opts *bind.WatchOpts, sink chan<- *DelegatorOnSlash, subnetwork [][32]byte, operator []common.Address) (event.Subscription, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "OnSlash", subnetworkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorOnSlash)
				if err := _Delegator.contract.UnpackLog(event, "OnSlash", log); err != nil {
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

// ParseOnSlash is a log parse operation binding the contract event 0x741a5de99085c0d660f3e4192217b0ffb0ea4e35a0480de48e857a4bc3ee36ed.
//
// Solidity: event OnSlash(bytes32 indexed subnetwork, address indexed operator, uint256 amount, uint48 captureTimestamp)
func (_Delegator *DelegatorFilterer) ParseOnSlash(log types.Log) (*DelegatorOnSlash, error) {
	event := new(DelegatorOnSlash)
	if err := _Delegator.contract.UnpackLog(event, "OnSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Delegator contract.
type DelegatorRoleAdminChangedIterator struct {
	Event *DelegatorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DelegatorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorRoleAdminChanged)
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
		it.Event = new(DelegatorRoleAdminChanged)
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
func (it *DelegatorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorRoleAdminChanged represents a RoleAdminChanged event raised by the Delegator contract.
type DelegatorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Delegator *DelegatorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DelegatorRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorRoleAdminChangedIterator{contract: _Delegator.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Delegator *DelegatorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DelegatorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorRoleAdminChanged)
				if err := _Delegator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Delegator *DelegatorFilterer) ParseRoleAdminChanged(log types.Log) (*DelegatorRoleAdminChanged, error) {
	event := new(DelegatorRoleAdminChanged)
	if err := _Delegator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Delegator contract.
type DelegatorRoleGrantedIterator struct {
	Event *DelegatorRoleGranted // Event containing the contract specifics and raw log

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
func (it *DelegatorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorRoleGranted)
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
		it.Event = new(DelegatorRoleGranted)
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
func (it *DelegatorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorRoleGranted represents a RoleGranted event raised by the Delegator contract.
type DelegatorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delegator *DelegatorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DelegatorRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorRoleGrantedIterator{contract: _Delegator.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delegator *DelegatorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DelegatorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorRoleGranted)
				if err := _Delegator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delegator *DelegatorFilterer) ParseRoleGranted(log types.Log) (*DelegatorRoleGranted, error) {
	event := new(DelegatorRoleGranted)
	if err := _Delegator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Delegator contract.
type DelegatorRoleRevokedIterator struct {
	Event *DelegatorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DelegatorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorRoleRevoked)
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
		it.Event = new(DelegatorRoleRevoked)
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
func (it *DelegatorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorRoleRevoked represents a RoleRevoked event raised by the Delegator contract.
type DelegatorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delegator *DelegatorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DelegatorRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorRoleRevokedIterator{contract: _Delegator.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delegator *DelegatorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DelegatorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorRoleRevoked)
				if err := _Delegator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Delegator *DelegatorFilterer) ParseRoleRevoked(log types.Log) (*DelegatorRoleRevoked, error) {
	event := new(DelegatorRoleRevoked)
	if err := _Delegator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorSetHookIterator is returned from FilterSetHook and is used to iterate over the raw logs and unpacked data for SetHook events raised by the Delegator contract.
type DelegatorSetHookIterator struct {
	Event *DelegatorSetHook // Event containing the contract specifics and raw log

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
func (it *DelegatorSetHookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorSetHook)
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
		it.Event = new(DelegatorSetHook)
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
func (it *DelegatorSetHookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorSetHookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorSetHook represents a SetHook event raised by the Delegator contract.
type DelegatorSetHook struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetHook is a free log retrieval operation binding the contract event 0x5bbb1d3ebb6a3ad2a0f17ff35e579a83af60604d1d3c2a4c83c62adecadf666d.
//
// Solidity: event SetHook(address indexed hook)
func (_Delegator *DelegatorFilterer) FilterSetHook(opts *bind.FilterOpts, hook []common.Address) (*DelegatorSetHookIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "SetHook", hookRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorSetHookIterator{contract: _Delegator.contract, event: "SetHook", logs: logs, sub: sub}, nil
}

// WatchSetHook is a free log subscription operation binding the contract event 0x5bbb1d3ebb6a3ad2a0f17ff35e579a83af60604d1d3c2a4c83c62adecadf666d.
//
// Solidity: event SetHook(address indexed hook)
func (_Delegator *DelegatorFilterer) WatchSetHook(opts *bind.WatchOpts, sink chan<- *DelegatorSetHook, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "SetHook", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorSetHook)
				if err := _Delegator.contract.UnpackLog(event, "SetHook", log); err != nil {
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

// ParseSetHook is a log parse operation binding the contract event 0x5bbb1d3ebb6a3ad2a0f17ff35e579a83af60604d1d3c2a4c83c62adecadf666d.
//
// Solidity: event SetHook(address indexed hook)
func (_Delegator *DelegatorFilterer) ParseSetHook(log types.Log) (*DelegatorSetHook, error) {
	event := new(DelegatorSetHook)
	if err := _Delegator.contract.UnpackLog(event, "SetHook", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorSetMaxNetworkLimitIterator is returned from FilterSetMaxNetworkLimit and is used to iterate over the raw logs and unpacked data for SetMaxNetworkLimit events raised by the Delegator contract.
type DelegatorSetMaxNetworkLimitIterator struct {
	Event *DelegatorSetMaxNetworkLimit // Event containing the contract specifics and raw log

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
func (it *DelegatorSetMaxNetworkLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorSetMaxNetworkLimit)
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
		it.Event = new(DelegatorSetMaxNetworkLimit)
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
func (it *DelegatorSetMaxNetworkLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorSetMaxNetworkLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorSetMaxNetworkLimit represents a SetMaxNetworkLimit event raised by the Delegator contract.
type DelegatorSetMaxNetworkLimit struct {
	Subnetwork [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMaxNetworkLimit is a free log retrieval operation binding the contract event 0xc67e7929681aa1bccd63f52b3799bf5805f3009f197db6fdf584b14f7fbf608c.
//
// Solidity: event SetMaxNetworkLimit(bytes32 indexed subnetwork, uint256 amount)
func (_Delegator *DelegatorFilterer) FilterSetMaxNetworkLimit(opts *bind.FilterOpts, subnetwork [][32]byte) (*DelegatorSetMaxNetworkLimitIterator, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "SetMaxNetworkLimit", subnetworkRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorSetMaxNetworkLimitIterator{contract: _Delegator.contract, event: "SetMaxNetworkLimit", logs: logs, sub: sub}, nil
}

// WatchSetMaxNetworkLimit is a free log subscription operation binding the contract event 0xc67e7929681aa1bccd63f52b3799bf5805f3009f197db6fdf584b14f7fbf608c.
//
// Solidity: event SetMaxNetworkLimit(bytes32 indexed subnetwork, uint256 amount)
func (_Delegator *DelegatorFilterer) WatchSetMaxNetworkLimit(opts *bind.WatchOpts, sink chan<- *DelegatorSetMaxNetworkLimit, subnetwork [][32]byte) (event.Subscription, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "SetMaxNetworkLimit", subnetworkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorSetMaxNetworkLimit)
				if err := _Delegator.contract.UnpackLog(event, "SetMaxNetworkLimit", log); err != nil {
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

// ParseSetMaxNetworkLimit is a log parse operation binding the contract event 0xc67e7929681aa1bccd63f52b3799bf5805f3009f197db6fdf584b14f7fbf608c.
//
// Solidity: event SetMaxNetworkLimit(bytes32 indexed subnetwork, uint256 amount)
func (_Delegator *DelegatorFilterer) ParseSetMaxNetworkLimit(log types.Log) (*DelegatorSetMaxNetworkLimit, error) {
	event := new(DelegatorSetMaxNetworkLimit)
	if err := _Delegator.contract.UnpackLog(event, "SetMaxNetworkLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorSetNetworkLimitIterator is returned from FilterSetNetworkLimit and is used to iterate over the raw logs and unpacked data for SetNetworkLimit events raised by the Delegator contract.
type DelegatorSetNetworkLimitIterator struct {
	Event *DelegatorSetNetworkLimit // Event containing the contract specifics and raw log

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
func (it *DelegatorSetNetworkLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorSetNetworkLimit)
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
		it.Event = new(DelegatorSetNetworkLimit)
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
func (it *DelegatorSetNetworkLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorSetNetworkLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorSetNetworkLimit represents a SetNetworkLimit event raised by the Delegator contract.
type DelegatorSetNetworkLimit struct {
	Subnetwork [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetNetworkLimit is a free log retrieval operation binding the contract event 0x00899d104fc3d8820bd96540612bcc5c448c8837b13b9f7faa43ad0728f0c14f.
//
// Solidity: event SetNetworkLimit(bytes32 indexed subnetwork, uint256 amount)
func (_Delegator *DelegatorFilterer) FilterSetNetworkLimit(opts *bind.FilterOpts, subnetwork [][32]byte) (*DelegatorSetNetworkLimitIterator, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "SetNetworkLimit", subnetworkRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorSetNetworkLimitIterator{contract: _Delegator.contract, event: "SetNetworkLimit", logs: logs, sub: sub}, nil
}

// WatchSetNetworkLimit is a free log subscription operation binding the contract event 0x00899d104fc3d8820bd96540612bcc5c448c8837b13b9f7faa43ad0728f0c14f.
//
// Solidity: event SetNetworkLimit(bytes32 indexed subnetwork, uint256 amount)
func (_Delegator *DelegatorFilterer) WatchSetNetworkLimit(opts *bind.WatchOpts, sink chan<- *DelegatorSetNetworkLimit, subnetwork [][32]byte) (event.Subscription, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "SetNetworkLimit", subnetworkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorSetNetworkLimit)
				if err := _Delegator.contract.UnpackLog(event, "SetNetworkLimit", log); err != nil {
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

// ParseSetNetworkLimit is a log parse operation binding the contract event 0x00899d104fc3d8820bd96540612bcc5c448c8837b13b9f7faa43ad0728f0c14f.
//
// Solidity: event SetNetworkLimit(bytes32 indexed subnetwork, uint256 amount)
func (_Delegator *DelegatorFilterer) ParseSetNetworkLimit(log types.Log) (*DelegatorSetNetworkLimit, error) {
	event := new(DelegatorSetNetworkLimit)
	if err := _Delegator.contract.UnpackLog(event, "SetNetworkLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelegatorSetOperatorNetworkSharesIterator is returned from FilterSetOperatorNetworkShares and is used to iterate over the raw logs and unpacked data for SetOperatorNetworkShares events raised by the Delegator contract.
type DelegatorSetOperatorNetworkSharesIterator struct {
	Event *DelegatorSetOperatorNetworkShares // Event containing the contract specifics and raw log

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
func (it *DelegatorSetOperatorNetworkSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelegatorSetOperatorNetworkShares)
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
		it.Event = new(DelegatorSetOperatorNetworkShares)
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
func (it *DelegatorSetOperatorNetworkSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelegatorSetOperatorNetworkSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelegatorSetOperatorNetworkShares represents a SetOperatorNetworkShares event raised by the Delegator contract.
type DelegatorSetOperatorNetworkShares struct {
	Subnetwork [32]byte
	Operator   common.Address
	Shares     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetOperatorNetworkShares is a free log retrieval operation binding the contract event 0x739a5a3ec0ff71e2386d0013deac5f44e0935a98def2e2a5ddf9a709518c8294.
//
// Solidity: event SetOperatorNetworkShares(bytes32 indexed subnetwork, address indexed operator, uint256 shares)
func (_Delegator *DelegatorFilterer) FilterSetOperatorNetworkShares(opts *bind.FilterOpts, subnetwork [][32]byte, operator []common.Address) (*DelegatorSetOperatorNetworkSharesIterator, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Delegator.contract.FilterLogs(opts, "SetOperatorNetworkShares", subnetworkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DelegatorSetOperatorNetworkSharesIterator{contract: _Delegator.contract, event: "SetOperatorNetworkShares", logs: logs, sub: sub}, nil
}

// WatchSetOperatorNetworkShares is a free log subscription operation binding the contract event 0x739a5a3ec0ff71e2386d0013deac5f44e0935a98def2e2a5ddf9a709518c8294.
//
// Solidity: event SetOperatorNetworkShares(bytes32 indexed subnetwork, address indexed operator, uint256 shares)
func (_Delegator *DelegatorFilterer) WatchSetOperatorNetworkShares(opts *bind.WatchOpts, sink chan<- *DelegatorSetOperatorNetworkShares, subnetwork [][32]byte, operator []common.Address) (event.Subscription, error) {

	var subnetworkRule []interface{}
	for _, subnetworkItem := range subnetwork {
		subnetworkRule = append(subnetworkRule, subnetworkItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Delegator.contract.WatchLogs(opts, "SetOperatorNetworkShares", subnetworkRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelegatorSetOperatorNetworkShares)
				if err := _Delegator.contract.UnpackLog(event, "SetOperatorNetworkShares", log); err != nil {
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

// ParseSetOperatorNetworkShares is a log parse operation binding the contract event 0x739a5a3ec0ff71e2386d0013deac5f44e0935a98def2e2a5ddf9a709518c8294.
//
// Solidity: event SetOperatorNetworkShares(bytes32 indexed subnetwork, address indexed operator, uint256 shares)
func (_Delegator *DelegatorFilterer) ParseSetOperatorNetworkShares(log types.Log) (*DelegatorSetOperatorNetworkShares, error) {
	event := new(DelegatorSetOperatorNetworkShares)
	if err := _Delegator.contract.UnpackLog(event, "SetOperatorNetworkShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
