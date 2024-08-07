// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package witnesschain

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

// WitnessChainOperatorRegistryMetaData contains all meta data concerning the WitnessChainOperatorRegistry contract.
var WitnessChainOperatorRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"OperatorSuspended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"operator\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"OperatorsWhiteListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"WatchtowerDeRegisteredFromOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"WatchtowerRegisteredToOperator\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operatorsList\",\"type\":\"address[]\"}],\"name\":\"addToOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateWatchtowerRegistrationMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkIsDelegatedOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtowerAddress\",\"type\":\"address\"}],\"name\":\"deRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableCheckIsDelegatedOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableCheckIsDelegatedOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllActiveOperators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasherAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isActiveOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"}],\"name\":\"isValidWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signedMessage\",\"type\":\"bytes\"}],\"name\":\"registerWatchtowerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManagerAddress\",\"type\":\"address\"}],\"name\":\"setDelegationManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasherAddress\",\"type\":\"address\"}],\"name\":\"setSlasherAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasherAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"suspend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// WitnessChainOperatorRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use WitnessChainOperatorRegistryMetaData.ABI instead.
var WitnessChainOperatorRegistryABI = WitnessChainOperatorRegistryMetaData.ABI

// WitnessChainOperatorRegistry is an auto generated Go binding around an Ethereum contract.
type WitnessChainOperatorRegistry struct {
	WitnessChainOperatorRegistryCaller     // Read-only binding to the contract
	WitnessChainOperatorRegistryTransactor // Write-only binding to the contract
	WitnessChainOperatorRegistryFilterer   // Log filterer for contract events
}

// WitnessChainOperatorRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WitnessChainOperatorRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessChainOperatorRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WitnessChainOperatorRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessChainOperatorRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WitnessChainOperatorRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessChainOperatorRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WitnessChainOperatorRegistrySession struct {
	Contract     *WitnessChainOperatorRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WitnessChainOperatorRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WitnessChainOperatorRegistryCallerSession struct {
	Contract *WitnessChainOperatorRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// WitnessChainOperatorRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WitnessChainOperatorRegistryTransactorSession struct {
	Contract     *WitnessChainOperatorRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// WitnessChainOperatorRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WitnessChainOperatorRegistryRaw struct {
	Contract *WitnessChainOperatorRegistry // Generic contract binding to access the raw methods on
}

// WitnessChainOperatorRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WitnessChainOperatorRegistryCallerRaw struct {
	Contract *WitnessChainOperatorRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// WitnessChainOperatorRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WitnessChainOperatorRegistryTransactorRaw struct {
	Contract *WitnessChainOperatorRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWitnessChainOperatorRegistry creates a new instance of WitnessChainOperatorRegistry, bound to a specific deployed contract.
func NewWitnessChainOperatorRegistry(address common.Address, backend bind.ContractBackend) (*WitnessChainOperatorRegistry, error) {
	contract, err := bindWitnessChainOperatorRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistry{WitnessChainOperatorRegistryCaller: WitnessChainOperatorRegistryCaller{contract: contract}, WitnessChainOperatorRegistryTransactor: WitnessChainOperatorRegistryTransactor{contract: contract}, WitnessChainOperatorRegistryFilterer: WitnessChainOperatorRegistryFilterer{contract: contract}}, nil
}

// NewWitnessChainOperatorRegistryCaller creates a new read-only instance of WitnessChainOperatorRegistry, bound to a specific deployed contract.
func NewWitnessChainOperatorRegistryCaller(address common.Address, caller bind.ContractCaller) (*WitnessChainOperatorRegistryCaller, error) {
	contract, err := bindWitnessChainOperatorRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryCaller{contract: contract}, nil
}

// NewWitnessChainOperatorRegistryTransactor creates a new write-only instance of WitnessChainOperatorRegistry, bound to a specific deployed contract.
func NewWitnessChainOperatorRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*WitnessChainOperatorRegistryTransactor, error) {
	contract, err := bindWitnessChainOperatorRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryTransactor{contract: contract}, nil
}

// NewWitnessChainOperatorRegistryFilterer creates a new log filterer instance of WitnessChainOperatorRegistry, bound to a specific deployed contract.
func NewWitnessChainOperatorRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*WitnessChainOperatorRegistryFilterer, error) {
	contract, err := bindWitnessChainOperatorRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryFilterer{contract: contract}, nil
}

// bindWitnessChainOperatorRegistry binds a generic wrapper to an already deployed contract.
func bindWitnessChainOperatorRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WitnessChainOperatorRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WitnessChainOperatorRegistry.Contract.WitnessChainOperatorRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.WitnessChainOperatorRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.WitnessChainOperatorRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WitnessChainOperatorRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.contract.Transact(opts, method, params...)
}

// CalculateWatchtowerRegistrationMessageHash is a free data retrieval call binding the contract method 0x45adeda0.
//
// Solidity: function calculateWatchtowerRegistrationMessageHash(address operator, uint256 expiry) pure returns(bytes32)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) CalculateWatchtowerRegistrationMessageHash(opts *bind.CallOpts, operator common.Address, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "calculateWatchtowerRegistrationMessageHash", operator, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWatchtowerRegistrationMessageHash is a free data retrieval call binding the contract method 0x45adeda0.
//
// Solidity: function calculateWatchtowerRegistrationMessageHash(address operator, uint256 expiry) pure returns(bytes32)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) CalculateWatchtowerRegistrationMessageHash(operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _WitnessChainOperatorRegistry.Contract.CalculateWatchtowerRegistrationMessageHash(&_WitnessChainOperatorRegistry.CallOpts, operator, expiry)
}

// CalculateWatchtowerRegistrationMessageHash is a free data retrieval call binding the contract method 0x45adeda0.
//
// Solidity: function calculateWatchtowerRegistrationMessageHash(address operator, uint256 expiry) pure returns(bytes32)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) CalculateWatchtowerRegistrationMessageHash(operator common.Address, expiry *big.Int) ([32]byte, error) {
	return _WitnessChainOperatorRegistry.Contract.CalculateWatchtowerRegistrationMessageHash(&_WitnessChainOperatorRegistry.CallOpts, operator, expiry)
}

// CheckIsDelegatedOperator is a free data retrieval call binding the contract method 0x9053c5b3.
//
// Solidity: function checkIsDelegatedOperator() view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) CheckIsDelegatedOperator(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "checkIsDelegatedOperator")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIsDelegatedOperator is a free data retrieval call binding the contract method 0x9053c5b3.
//
// Solidity: function checkIsDelegatedOperator() view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) CheckIsDelegatedOperator() (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.CheckIsDelegatedOperator(&_WitnessChainOperatorRegistry.CallOpts)
}

// CheckIsDelegatedOperator is a free data retrieval call binding the contract method 0x9053c5b3.
//
// Solidity: function checkIsDelegatedOperator() view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) CheckIsDelegatedOperator() (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.CheckIsDelegatedOperator(&_WitnessChainOperatorRegistry.CallOpts)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) DelegationManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "delegationManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) DelegationManagerAddress() (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.DelegationManagerAddress(&_WitnessChainOperatorRegistry.CallOpts)
}

// DelegationManagerAddress is a free data retrieval call binding the contract method 0xb15e6689.
//
// Solidity: function delegationManagerAddress() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) DelegationManagerAddress() (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.DelegationManagerAddress(&_WitnessChainOperatorRegistry.CallOpts)
}

// GetAllActiveOperators is a free data retrieval call binding the contract method 0x5c468db3.
//
// Solidity: function getAllActiveOperators() view returns(address[])
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) GetAllActiveOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "getAllActiveOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllActiveOperators is a free data retrieval call binding the contract method 0x5c468db3.
//
// Solidity: function getAllActiveOperators() view returns(address[])
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) GetAllActiveOperators() ([]common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.GetAllActiveOperators(&_WitnessChainOperatorRegistry.CallOpts)
}

// GetAllActiveOperators is a free data retrieval call binding the contract method 0x5c468db3.
//
// Solidity: function getAllActiveOperators() view returns(address[])
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) GetAllActiveOperators() ([]common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.GetAllActiveOperators(&_WitnessChainOperatorRegistry.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) GetOperator(opts *bind.CallOpts, watchtower common.Address) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "getOperator", watchtower)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) GetOperator(watchtower common.Address) (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.GetOperator(&_WitnessChainOperatorRegistry.CallOpts, watchtower)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address watchtower) view returns(address operator)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) GetOperator(watchtower common.Address) (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.GetOperator(&_WitnessChainOperatorRegistry.CallOpts, watchtower)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) IsActiveOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "isActiveOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) IsActiveOperator(operator common.Address) (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.IsActiveOperator(&_WitnessChainOperatorRegistry.CallOpts, operator)
}

// IsActiveOperator is a free data retrieval call binding the contract method 0x3367cca5.
//
// Solidity: function isActiveOperator(address operator) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) IsActiveOperator(operator common.Address) (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.IsActiveOperator(&_WitnessChainOperatorRegistry.CallOpts, operator)
}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) IsValidWatchtower(opts *bind.CallOpts, watchtower common.Address) (bool, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "isValidWatchtower", watchtower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) IsValidWatchtower(watchtower common.Address) (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.IsValidWatchtower(&_WitnessChainOperatorRegistry.CallOpts, watchtower)
}

// IsValidWatchtower is a free data retrieval call binding the contract method 0x9a521382.
//
// Solidity: function isValidWatchtower(address watchtower) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) IsValidWatchtower(watchtower common.Address) (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.IsValidWatchtower(&_WitnessChainOperatorRegistry.CallOpts, watchtower)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address operator) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) IsWhitelisted(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "isWhitelisted", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address operator) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) IsWhitelisted(operator common.Address) (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.IsWhitelisted(&_WitnessChainOperatorRegistry.CallOpts, operator)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address operator) view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) IsWhitelisted(operator common.Address) (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.IsWhitelisted(&_WitnessChainOperatorRegistry.CallOpts, operator)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address ) view returns(address operatorAddress, bool isActive)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) OperatorDetails(opts *bind.CallOpts, arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsActive        bool
}, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "operatorDetails", arg0)

	outstruct := new(struct {
		OperatorAddress common.Address
		IsActive        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperatorAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address ) view returns(address operatorAddress, bool isActive)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) OperatorDetails(arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsActive        bool
}, error) {
	return _WitnessChainOperatorRegistry.Contract.OperatorDetails(&_WitnessChainOperatorRegistry.CallOpts, arg0)
}

// OperatorDetails is a free data retrieval call binding the contract method 0xc5e480db.
//
// Solidity: function operatorDetails(address ) view returns(address operatorAddress, bool isActive)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) OperatorDetails(arg0 common.Address) (struct {
	OperatorAddress common.Address
	IsActive        bool
}, error) {
	return _WitnessChainOperatorRegistry.Contract.OperatorDetails(&_WitnessChainOperatorRegistry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) Owner() (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.Owner(&_WitnessChainOperatorRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) Owner() (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.Owner(&_WitnessChainOperatorRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) Paused() (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.Paused(&_WitnessChainOperatorRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) Paused() (bool, error) {
	return _WitnessChainOperatorRegistry.Contract.Paused(&_WitnessChainOperatorRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _WitnessChainOperatorRegistry.Contract.ProxiableUUID(&_WitnessChainOperatorRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WitnessChainOperatorRegistry.Contract.ProxiableUUID(&_WitnessChainOperatorRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCaller) SlasherAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainOperatorRegistry.contract.Call(opts, &out, "slasherAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) SlasherAddress() (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.SlasherAddress(&_WitnessChainOperatorRegistry.CallOpts)
}

// SlasherAddress is a free data retrieval call binding the contract method 0xd53c61bf.
//
// Solidity: function slasherAddress() view returns(address)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryCallerSession) SlasherAddress() (common.Address, error) {
	return _WitnessChainOperatorRegistry.Contract.SlasherAddress(&_WitnessChainOperatorRegistry.CallOpts)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) AddToOperatorWhitelist(opts *bind.TransactOpts, operatorsList []common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "addToOperatorWhitelist", operatorsList)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) AddToOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.AddToOperatorWhitelist(&_WitnessChainOperatorRegistry.TransactOpts, operatorsList)
}

// AddToOperatorWhitelist is a paid mutator transaction binding the contract method 0xc8525c3e.
//
// Solidity: function addToOperatorWhitelist(address[] operatorsList) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) AddToOperatorWhitelist(operatorsList []common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.AddToOperatorWhitelist(&_WitnessChainOperatorRegistry.TransactOpts, operatorsList)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address watchtowerAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) DeRegister(opts *bind.TransactOpts, watchtowerAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "deRegister", watchtowerAddress)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address watchtowerAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) DeRegister(watchtowerAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.DeRegister(&_WitnessChainOperatorRegistry.TransactOpts, watchtowerAddress)
}

// DeRegister is a paid mutator transaction binding the contract method 0x5b114af6.
//
// Solidity: function deRegister(address watchtowerAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) DeRegister(watchtowerAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.DeRegister(&_WitnessChainOperatorRegistry.TransactOpts, watchtowerAddress)
}

// DisableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x827a1bf5.
//
// Solidity: function disableCheckIsDelegatedOperator() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) DisableCheckIsDelegatedOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "disableCheckIsDelegatedOperator")
}

// DisableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x827a1bf5.
//
// Solidity: function disableCheckIsDelegatedOperator() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) DisableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.DisableCheckIsDelegatedOperator(&_WitnessChainOperatorRegistry.TransactOpts)
}

// DisableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x827a1bf5.
//
// Solidity: function disableCheckIsDelegatedOperator() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) DisableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.DisableCheckIsDelegatedOperator(&_WitnessChainOperatorRegistry.TransactOpts)
}

// EnableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x17fa0ce8.
//
// Solidity: function enableCheckIsDelegatedOperator() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) EnableCheckIsDelegatedOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "enableCheckIsDelegatedOperator")
}

// EnableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x17fa0ce8.
//
// Solidity: function enableCheckIsDelegatedOperator() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) EnableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.EnableCheckIsDelegatedOperator(&_WitnessChainOperatorRegistry.TransactOpts)
}

// EnableCheckIsDelegatedOperator is a paid mutator transaction binding the contract method 0x17fa0ce8.
//
// Solidity: function enableCheckIsDelegatedOperator() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) EnableCheckIsDelegatedOperator() (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.EnableCheckIsDelegatedOperator(&_WitnessChainOperatorRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) Initialize(opts *bind.TransactOpts, _delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "initialize", _delegationManagerAddress, _slasherAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) Initialize(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.Initialize(&_WitnessChainOperatorRegistry.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManagerAddress, address _slasherAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) Initialize(_delegationManagerAddress common.Address, _slasherAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.Initialize(&_WitnessChainOperatorRegistry.TransactOpts, _delegationManagerAddress, _slasherAddress)
}

// RegisterWatchtowerAsOperator is a paid mutator transaction binding the contract method 0x11d2c708.
//
// Solidity: function registerWatchtowerAsOperator(address watchtower, uint256 expiry, bytes signedMessage) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) RegisterWatchtowerAsOperator(opts *bind.TransactOpts, watchtower common.Address, expiry *big.Int, signedMessage []byte) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "registerWatchtowerAsOperator", watchtower, expiry, signedMessage)
}

// RegisterWatchtowerAsOperator is a paid mutator transaction binding the contract method 0x11d2c708.
//
// Solidity: function registerWatchtowerAsOperator(address watchtower, uint256 expiry, bytes signedMessage) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) RegisterWatchtowerAsOperator(watchtower common.Address, expiry *big.Int, signedMessage []byte) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.RegisterWatchtowerAsOperator(&_WitnessChainOperatorRegistry.TransactOpts, watchtower, expiry, signedMessage)
}

// RegisterWatchtowerAsOperator is a paid mutator transaction binding the contract method 0x11d2c708.
//
// Solidity: function registerWatchtowerAsOperator(address watchtower, uint256 expiry, bytes signedMessage) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) RegisterWatchtowerAsOperator(watchtower common.Address, expiry *big.Int, signedMessage []byte) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.RegisterWatchtowerAsOperator(&_WitnessChainOperatorRegistry.TransactOpts, watchtower, expiry, signedMessage)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.RenounceOwnership(&_WitnessChainOperatorRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.RenounceOwnership(&_WitnessChainOperatorRegistry.TransactOpts)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) SetDelegationManagerAddress(opts *bind.TransactOpts, _delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "setDelegationManagerAddress", _delegationManagerAddress)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) SetDelegationManagerAddress(_delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.SetDelegationManagerAddress(&_WitnessChainOperatorRegistry.TransactOpts, _delegationManagerAddress)
}

// SetDelegationManagerAddress is a paid mutator transaction binding the contract method 0x7831ad20.
//
// Solidity: function setDelegationManagerAddress(address _delegationManagerAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) SetDelegationManagerAddress(_delegationManagerAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.SetDelegationManagerAddress(&_WitnessChainOperatorRegistry.TransactOpts, _delegationManagerAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) SetSlasherAddress(opts *bind.TransactOpts, _slasherAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "setSlasherAddress", _slasherAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) SetSlasherAddress(_slasherAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.SetSlasherAddress(&_WitnessChainOperatorRegistry.TransactOpts, _slasherAddress)
}

// SetSlasherAddress is a paid mutator transaction binding the contract method 0xaddd9cf5.
//
// Solidity: function setSlasherAddress(address _slasherAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) SetSlasherAddress(_slasherAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.SetSlasherAddress(&_WitnessChainOperatorRegistry.TransactOpts, _slasherAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) Suspend(opts *bind.TransactOpts, operatorAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "suspend", operatorAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) Suspend(operatorAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.Suspend(&_WitnessChainOperatorRegistry.TransactOpts, operatorAddress)
}

// Suspend is a paid mutator transaction binding the contract method 0x286781c7.
//
// Solidity: function suspend(address operatorAddress) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) Suspend(operatorAddress common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.Suspend(&_WitnessChainOperatorRegistry.TransactOpts, operatorAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.TransferOwnership(&_WitnessChainOperatorRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.TransferOwnership(&_WitnessChainOperatorRegistry.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.UpgradeTo(&_WitnessChainOperatorRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.UpgradeTo(&_WitnessChainOperatorRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.UpgradeToAndCall(&_WitnessChainOperatorRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessChainOperatorRegistry.Contract.UpgradeToAndCall(&_WitnessChainOperatorRegistry.TransactOpts, newImplementation, data)
}

// WitnessChainOperatorRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryAdminChangedIterator struct {
	Event *WitnessChainOperatorRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryAdminChanged)
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
		it.Event = new(WitnessChainOperatorRegistryAdminChanged)
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
func (it *WitnessChainOperatorRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryAdminChanged represents a AdminChanged event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryAdminChangedIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryAdminChangedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryAdminChanged)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseAdminChanged(log types.Log) (*WitnessChainOperatorRegistryAdminChanged, error) {
	event := new(WitnessChainOperatorRegistryAdminChanged)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryBeaconUpgradedIterator struct {
	Event *WitnessChainOperatorRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryBeaconUpgraded)
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
		it.Event = new(WitnessChainOperatorRegistryBeaconUpgraded)
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
func (it *WitnessChainOperatorRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WitnessChainOperatorRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryBeaconUpgradedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryBeaconUpgraded)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*WitnessChainOperatorRegistryBeaconUpgraded, error) {
	event := new(WitnessChainOperatorRegistryBeaconUpgraded)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryInitializedIterator struct {
	Event *WitnessChainOperatorRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryInitialized)
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
		it.Event = new(WitnessChainOperatorRegistryInitialized)
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
func (it *WitnessChainOperatorRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryInitialized represents a Initialized event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryInitializedIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryInitializedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryInitialized)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseInitialized(log types.Log) (*WitnessChainOperatorRegistryInitialized, error) {
	event := new(WitnessChainOperatorRegistryInitialized)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryOperatorSuspendedIterator is returned from FilterOperatorSuspended and is used to iterate over the raw logs and unpacked data for OperatorSuspended events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryOperatorSuspendedIterator struct {
	Event *WitnessChainOperatorRegistryOperatorSuspended // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryOperatorSuspendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryOperatorSuspended)
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
		it.Event = new(WitnessChainOperatorRegistryOperatorSuspended)
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
func (it *WitnessChainOperatorRegistryOperatorSuspendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryOperatorSuspendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryOperatorSuspended represents a OperatorSuspended event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryOperatorSuspended struct {
	Operator    common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSuspended is a free log retrieval operation binding the contract event 0x98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def.
//
// Solidity: event OperatorSuspended(address operator, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterOperatorSuspended(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryOperatorSuspendedIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "OperatorSuspended")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryOperatorSuspendedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "OperatorSuspended", logs: logs, sub: sub}, nil
}

// WatchOperatorSuspended is a free log subscription operation binding the contract event 0x98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def.
//
// Solidity: event OperatorSuspended(address operator, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchOperatorSuspended(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryOperatorSuspended) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "OperatorSuspended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryOperatorSuspended)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "OperatorSuspended", log); err != nil {
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

// ParseOperatorSuspended is a log parse operation binding the contract event 0x98ff055e829cd4bfba35f84d3b43877fd53084d2deebbb1cda57404f44df4def.
//
// Solidity: event OperatorSuspended(address operator, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseOperatorSuspended(log types.Log) (*WitnessChainOperatorRegistryOperatorSuspended, error) {
	event := new(WitnessChainOperatorRegistryOperatorSuspended)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "OperatorSuspended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryOperatorsWhiteListedIterator is returned from FilterOperatorsWhiteListed and is used to iterate over the raw logs and unpacked data for OperatorsWhiteListed events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryOperatorsWhiteListedIterator struct {
	Event *WitnessChainOperatorRegistryOperatorsWhiteListed // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryOperatorsWhiteListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryOperatorsWhiteListed)
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
		it.Event = new(WitnessChainOperatorRegistryOperatorsWhiteListed)
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
func (it *WitnessChainOperatorRegistryOperatorsWhiteListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryOperatorsWhiteListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryOperatorsWhiteListed represents a OperatorsWhiteListed event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryOperatorsWhiteListed struct {
	Operator    []common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorsWhiteListed is a free log retrieval operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterOperatorsWhiteListed(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryOperatorsWhiteListedIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "OperatorsWhiteListed")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryOperatorsWhiteListedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "OperatorsWhiteListed", logs: logs, sub: sub}, nil
}

// WatchOperatorsWhiteListed is a free log subscription operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchOperatorsWhiteListed(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryOperatorsWhiteListed) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "OperatorsWhiteListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryOperatorsWhiteListed)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "OperatorsWhiteListed", log); err != nil {
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

// ParseOperatorsWhiteListed is a log parse operation binding the contract event 0xdd7d462090c5f3ddd88b4c509269e35d5e148454dc3a9bd24812f78efc8f306f.
//
// Solidity: event OperatorsWhiteListed(address[] operator, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseOperatorsWhiteListed(log types.Log) (*WitnessChainOperatorRegistryOperatorsWhiteListed, error) {
	event := new(WitnessChainOperatorRegistryOperatorsWhiteListed)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "OperatorsWhiteListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryOwnershipTransferredIterator struct {
	Event *WitnessChainOperatorRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryOwnershipTransferred)
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
		it.Event = new(WitnessChainOperatorRegistryOwnershipTransferred)
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
func (it *WitnessChainOperatorRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WitnessChainOperatorRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryOwnershipTransferredIterator{contract: _WitnessChainOperatorRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryOwnershipTransferred)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*WitnessChainOperatorRegistryOwnershipTransferred, error) {
	event := new(WitnessChainOperatorRegistryOwnershipTransferred)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryPausedIterator struct {
	Event *WitnessChainOperatorRegistryPaused // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryPaused)
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
		it.Event = new(WitnessChainOperatorRegistryPaused)
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
func (it *WitnessChainOperatorRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryPaused represents a Paused event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryPausedIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryPausedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryPaused)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParsePaused(log types.Log) (*WitnessChainOperatorRegistryPaused, error) {
	event := new(WitnessChainOperatorRegistryPaused)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryUnpausedIterator struct {
	Event *WitnessChainOperatorRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryUnpaused)
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
		it.Event = new(WitnessChainOperatorRegistryUnpaused)
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
func (it *WitnessChainOperatorRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryUnpaused represents a Unpaused event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryUnpausedIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryUnpausedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryUnpaused)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseUnpaused(log types.Log) (*WitnessChainOperatorRegistryUnpaused, error) {
	event := new(WitnessChainOperatorRegistryUnpaused)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryUpgradedIterator struct {
	Event *WitnessChainOperatorRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryUpgraded)
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
		it.Event = new(WitnessChainOperatorRegistryUpgraded)
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
func (it *WitnessChainOperatorRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryUpgraded represents a Upgraded event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WitnessChainOperatorRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryUpgradedIterator{contract: _WitnessChainOperatorRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryUpgraded)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseUpgraded(log types.Log) (*WitnessChainOperatorRegistryUpgraded, error) {
	event := new(WitnessChainOperatorRegistryUpgraded)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperatorIterator is returned from FilterWatchtowerDeRegisteredFromOperator and is used to iterate over the raw logs and unpacked data for WatchtowerDeRegisteredFromOperator events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperatorIterator struct {
	Event *WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator)
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
		it.Event = new(WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator)
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
func (it *WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator represents a WatchtowerDeRegisteredFromOperator event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator struct {
	Operator    common.Address
	Watchtower  common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWatchtowerDeRegisteredFromOperator is a free log retrieval operation binding the contract event 0xecfc290ff8c3aac71e14aee07653f81e5aa316be4d2315ba2ec1bff9dc50cd79.
//
// Solidity: event WatchtowerDeRegisteredFromOperator(address operator, address watchtower, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterWatchtowerDeRegisteredFromOperator(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperatorIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "WatchtowerDeRegisteredFromOperator")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperatorIterator{contract: _WitnessChainOperatorRegistry.contract, event: "WatchtowerDeRegisteredFromOperator", logs: logs, sub: sub}, nil
}

// WatchWatchtowerDeRegisteredFromOperator is a free log subscription operation binding the contract event 0xecfc290ff8c3aac71e14aee07653f81e5aa316be4d2315ba2ec1bff9dc50cd79.
//
// Solidity: event WatchtowerDeRegisteredFromOperator(address operator, address watchtower, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchWatchtowerDeRegisteredFromOperator(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "WatchtowerDeRegisteredFromOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "WatchtowerDeRegisteredFromOperator", log); err != nil {
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

// ParseWatchtowerDeRegisteredFromOperator is a log parse operation binding the contract event 0xecfc290ff8c3aac71e14aee07653f81e5aa316be4d2315ba2ec1bff9dc50cd79.
//
// Solidity: event WatchtowerDeRegisteredFromOperator(address operator, address watchtower, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseWatchtowerDeRegisteredFromOperator(log types.Log) (*WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator, error) {
	event := new(WitnessChainOperatorRegistryWatchtowerDeRegisteredFromOperator)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "WatchtowerDeRegisteredFromOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainOperatorRegistryWatchtowerRegisteredToOperatorIterator is returned from FilterWatchtowerRegisteredToOperator and is used to iterate over the raw logs and unpacked data for WatchtowerRegisteredToOperator events raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryWatchtowerRegisteredToOperatorIterator struct {
	Event *WitnessChainOperatorRegistryWatchtowerRegisteredToOperator // Event containing the contract specifics and raw log

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
func (it *WitnessChainOperatorRegistryWatchtowerRegisteredToOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainOperatorRegistryWatchtowerRegisteredToOperator)
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
		it.Event = new(WitnessChainOperatorRegistryWatchtowerRegisteredToOperator)
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
func (it *WitnessChainOperatorRegistryWatchtowerRegisteredToOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainOperatorRegistryWatchtowerRegisteredToOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainOperatorRegistryWatchtowerRegisteredToOperator represents a WatchtowerRegisteredToOperator event raised by the WitnessChainOperatorRegistry contract.
type WitnessChainOperatorRegistryWatchtowerRegisteredToOperator struct {
	Operator    common.Address
	Watchtower  common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWatchtowerRegisteredToOperator is a free log retrieval operation binding the contract event 0x3d521ffddd8cbc2d603c1e7fe9af4b70adbbc43675a7281c5624d8053e526f06.
//
// Solidity: event WatchtowerRegisteredToOperator(address operator, address watchtower, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) FilterWatchtowerRegisteredToOperator(opts *bind.FilterOpts) (*WitnessChainOperatorRegistryWatchtowerRegisteredToOperatorIterator, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.FilterLogs(opts, "WatchtowerRegisteredToOperator")
	if err != nil {
		return nil, err
	}
	return &WitnessChainOperatorRegistryWatchtowerRegisteredToOperatorIterator{contract: _WitnessChainOperatorRegistry.contract, event: "WatchtowerRegisteredToOperator", logs: logs, sub: sub}, nil
}

// WatchWatchtowerRegisteredToOperator is a free log subscription operation binding the contract event 0x3d521ffddd8cbc2d603c1e7fe9af4b70adbbc43675a7281c5624d8053e526f06.
//
// Solidity: event WatchtowerRegisteredToOperator(address operator, address watchtower, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) WatchWatchtowerRegisteredToOperator(opts *bind.WatchOpts, sink chan<- *WitnessChainOperatorRegistryWatchtowerRegisteredToOperator) (event.Subscription, error) {

	logs, sub, err := _WitnessChainOperatorRegistry.contract.WatchLogs(opts, "WatchtowerRegisteredToOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainOperatorRegistryWatchtowerRegisteredToOperator)
				if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "WatchtowerRegisteredToOperator", log); err != nil {
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

// ParseWatchtowerRegisteredToOperator is a log parse operation binding the contract event 0x3d521ffddd8cbc2d603c1e7fe9af4b70adbbc43675a7281c5624d8053e526f06.
//
// Solidity: event WatchtowerRegisteredToOperator(address operator, address watchtower, uint256 blockNumber)
func (_WitnessChainOperatorRegistry *WitnessChainOperatorRegistryFilterer) ParseWatchtowerRegisteredToOperator(log types.Log) (*WitnessChainOperatorRegistryWatchtowerRegisteredToOperator, error) {
	event := new(WitnessChainOperatorRegistryWatchtowerRegisteredToOperator)
	if err := _WitnessChainOperatorRegistry.contract.UnpackLog(event, "WatchtowerRegisteredToOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
