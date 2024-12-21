// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package predicate

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

// Task is an auto generated low-level Go binding around an user-defined struct.
type Task struct {
	TaskId               string
	MsgSender            common.Address
	Target               common.Address
	Value                *big.Int
	EncodedSigAndArgs    []byte
	PolicyID             string
	QuorumThresholdCount uint32
	ExpireByBlockNumber  *big.Int
}

// ServiceManagerMetaData contains all meta data concerning the ServiceManager contract.
var ServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ServiceManager__ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServiceManager__InvalidOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServiceManager__InvalidStrategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServiceManager__Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"avsDirectory\",\"type\":\"address\"}],\"name\":\"AVSDirectoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"AggregatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegationManager\",\"type\":\"address\"}],\"name\":\"DelegationManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policy\",\"type\":\"string\"}],\"name\":\"DeployedPolicy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"NonCompliantTask\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address[][]\",\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"OperatorsStakesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"}],\"name\":\"RemovedPolicy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"}],\"name\":\"SetPolicy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"socialGraphID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socialGraphConfig\",\"type\":\"string\"}],\"name\":\"SocialGraphDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeRegistry\",\"type\":\"address\"}],\"name\":\"StakeRegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"}],\"name\":\"TaskExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"quorumThresholdCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expireByBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signerAddresses\",\"type\":\"address[]\"}],\"name\":\"TaskValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"thresholdStake\",\"type\":\"uint256\"}],\"name\":\"ThresholdStakeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"addStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"clientToPolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_policyID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_policy\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_quorumThreshold\",\"type\":\"uint256\"}],\"name\":\"deployPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_socialGraphID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_socialGraphConfig\",\"type\":\"string\"}],\"name\":\"deploySocialGraph\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deployedPolicyIDs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployedPolicies\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSocialGraphIDs\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedSigAndArgs\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdCount\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"expireByBlockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structTask\",\"name\":\"_task\",\"type\":\"tuple\"}],\"name\":\"hashTaskWithExpiry\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"idToPolicy\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"idToSocialGraph\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_avsDirectory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_thresholdStake\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalStake\",\"type\":\"uint256\"},{\"internalType\":\"enumServiceManager.OperatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operatorSigningKey\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_policyID\",\"type\":\"string\"}],\"name\":\"removePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"removeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldSigningKey\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newSigningKey\",\"type\":\"address\"}],\"name\":\"rotatePredicateSigningKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avsDirectory\",\"type\":\"address\"}],\"name\":\"setAVSDirectory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManager\",\"type\":\"address\"}],\"name\":\"setDelegationManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"setMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_policyID\",\"type\":\"string\"}],\"name\":\"setPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeRegistry\",\"type\":\"address\"}],\"name\":\"setStakeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_thresholdStake\",\"type\":\"uint256\"}],\"name\":\"setThresholdStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signingKeyToOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"socialGraphIDs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"spentTaskIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"strategies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thresholdStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[][]\",\"name\":\"operatorsPerQuorum\",\"type\":\"address[][]\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"updateOperatorsForQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"taskId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedSigAndArgs\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"policyID\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdCount\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"expireByBlockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structTask\",\"name\":\"_task\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"signerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"validateSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ServiceManagerMetaData.ABI instead.
var ServiceManagerABI = ServiceManagerMetaData.ABI

// ServiceManager is an auto generated Go binding around an Ethereum contract.
type ServiceManager struct {
	ServiceManagerCaller     // Read-only binding to the contract
	ServiceManagerTransactor // Write-only binding to the contract
	ServiceManagerFilterer   // Log filterer for contract events
}

// ServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServiceManagerSession struct {
	Contract     *ServiceManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServiceManagerCallerSession struct {
	Contract *ServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServiceManagerTransactorSession struct {
	Contract     *ServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServiceManagerRaw struct {
	Contract *ServiceManager // Generic contract binding to access the raw methods on
}

// ServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServiceManagerCallerRaw struct {
	Contract *ServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServiceManagerTransactorRaw struct {
	Contract *ServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServiceManager creates a new instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManager(address common.Address, backend bind.ContractBackend) (*ServiceManager, error) {
	contract, err := bindServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServiceManager{ServiceManagerCaller: ServiceManagerCaller{contract: contract}, ServiceManagerTransactor: ServiceManagerTransactor{contract: contract}, ServiceManagerFilterer: ServiceManagerFilterer{contract: contract}}, nil
}

// NewServiceManagerCaller creates a new read-only instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ServiceManagerCaller, error) {
	contract, err := bindServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerCaller{contract: contract}, nil
}

// NewServiceManagerTransactor creates a new write-only instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceManagerTransactor, error) {
	contract, err := bindServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerTransactor{contract: contract}, nil
}

// NewServiceManagerFilterer creates a new log filterer instance of ServiceManager, bound to a specific deployed contract.
func NewServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceManagerFilterer, error) {
	contract, err := bindServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerFilterer{contract: contract}, nil
}

// bindServiceManager binds a generic wrapper to an already deployed contract.
func bindServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceManager *ServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceManager.Contract.ServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceManager *ServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceManager.Contract.ServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceManager *ServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceManager.Contract.ServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceManager *ServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceManager *ServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceManager *ServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceManager.Contract.contract.Transact(opts, method, params...)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ServiceManager *ServiceManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ServiceManager *ServiceManagerSession) Aggregator() (common.Address, error) {
	return _ServiceManager.Contract.Aggregator(&_ServiceManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ServiceManager *ServiceManagerCallerSession) Aggregator() (common.Address, error) {
	return _ServiceManager.Contract.Aggregator(&_ServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ServiceManager *ServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ServiceManager *ServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ServiceManager.Contract.AvsDirectory(&_ServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ServiceManager *ServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ServiceManager.Contract.AvsDirectory(&_ServiceManager.CallOpts)
}

// ClientToPolicy is a free data retrieval call binding the contract method 0x140a16bc.
//
// Solidity: function clientToPolicy(address , string ) view returns(bool)
func (_ServiceManager *ServiceManagerCaller) ClientToPolicy(opts *bind.CallOpts, arg0 common.Address, arg1 string) (bool, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "clientToPolicy", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClientToPolicy is a free data retrieval call binding the contract method 0x140a16bc.
//
// Solidity: function clientToPolicy(address , string ) view returns(bool)
func (_ServiceManager *ServiceManagerSession) ClientToPolicy(arg0 common.Address, arg1 string) (bool, error) {
	return _ServiceManager.Contract.ClientToPolicy(&_ServiceManager.CallOpts, arg0, arg1)
}

// ClientToPolicy is a free data retrieval call binding the contract method 0x140a16bc.
//
// Solidity: function clientToPolicy(address , string ) view returns(bool)
func (_ServiceManager *ServiceManagerCallerSession) ClientToPolicy(arg0 common.Address, arg1 string) (bool, error) {
	return _ServiceManager.Contract.ClientToPolicy(&_ServiceManager.CallOpts, arg0, arg1)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ServiceManager *ServiceManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ServiceManager *ServiceManagerSession) DelegationManager() (common.Address, error) {
	return _ServiceManager.Contract.DelegationManager(&_ServiceManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ServiceManager *ServiceManagerCallerSession) DelegationManager() (common.Address, error) {
	return _ServiceManager.Contract.DelegationManager(&_ServiceManager.CallOpts)
}

// DeployedPolicyIDs is a free data retrieval call binding the contract method 0xddb49ce1.
//
// Solidity: function deployedPolicyIDs(uint256 ) view returns(string)
func (_ServiceManager *ServiceManagerCaller) DeployedPolicyIDs(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "deployedPolicyIDs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DeployedPolicyIDs is a free data retrieval call binding the contract method 0xddb49ce1.
//
// Solidity: function deployedPolicyIDs(uint256 ) view returns(string)
func (_ServiceManager *ServiceManagerSession) DeployedPolicyIDs(arg0 *big.Int) (string, error) {
	return _ServiceManager.Contract.DeployedPolicyIDs(&_ServiceManager.CallOpts, arg0)
}

// DeployedPolicyIDs is a free data retrieval call binding the contract method 0xddb49ce1.
//
// Solidity: function deployedPolicyIDs(uint256 ) view returns(string)
func (_ServiceManager *ServiceManagerCallerSession) DeployedPolicyIDs(arg0 *big.Int) (string, error) {
	return _ServiceManager.Contract.DeployedPolicyIDs(&_ServiceManager.CallOpts, arg0)
}

// GetDeployedPolicies is a free data retrieval call binding the contract method 0x0b3ce015.
//
// Solidity: function getDeployedPolicies() view returns(string[])
func (_ServiceManager *ServiceManagerCaller) GetDeployedPolicies(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "getDeployedPolicies")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetDeployedPolicies is a free data retrieval call binding the contract method 0x0b3ce015.
//
// Solidity: function getDeployedPolicies() view returns(string[])
func (_ServiceManager *ServiceManagerSession) GetDeployedPolicies() ([]string, error) {
	return _ServiceManager.Contract.GetDeployedPolicies(&_ServiceManager.CallOpts)
}

// GetDeployedPolicies is a free data retrieval call binding the contract method 0x0b3ce015.
//
// Solidity: function getDeployedPolicies() view returns(string[])
func (_ServiceManager *ServiceManagerCallerSession) GetDeployedPolicies() ([]string, error) {
	return _ServiceManager.Contract.GetDeployedPolicies(&_ServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ServiceManager *ServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ServiceManager *ServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ServiceManager.Contract.GetOperatorRestakedStrategies(&_ServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ServiceManager *ServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ServiceManager.Contract.GetOperatorRestakedStrategies(&_ServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ServiceManager *ServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ServiceManager *ServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ServiceManager.Contract.GetRestakeableStrategies(&_ServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ServiceManager *ServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ServiceManager.Contract.GetRestakeableStrategies(&_ServiceManager.CallOpts)
}

// GetSocialGraphIDs is a free data retrieval call binding the contract method 0xd18a1325.
//
// Solidity: function getSocialGraphIDs() view returns(string[])
func (_ServiceManager *ServiceManagerCaller) GetSocialGraphIDs(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "getSocialGraphIDs")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetSocialGraphIDs is a free data retrieval call binding the contract method 0xd18a1325.
//
// Solidity: function getSocialGraphIDs() view returns(string[])
func (_ServiceManager *ServiceManagerSession) GetSocialGraphIDs() ([]string, error) {
	return _ServiceManager.Contract.GetSocialGraphIDs(&_ServiceManager.CallOpts)
}

// GetSocialGraphIDs is a free data retrieval call binding the contract method 0xd18a1325.
//
// Solidity: function getSocialGraphIDs() view returns(string[])
func (_ServiceManager *ServiceManagerCallerSession) GetSocialGraphIDs() ([]string, error) {
	return _ServiceManager.Contract.GetSocialGraphIDs(&_ServiceManager.CallOpts)
}

// HashTaskWithExpiry is a free data retrieval call binding the contract method 0x949528fc.
//
// Solidity: function hashTaskWithExpiry((string,address,address,uint256,bytes,string,uint32,uint256) _task) pure returns(bytes32)
func (_ServiceManager *ServiceManagerCaller) HashTaskWithExpiry(opts *bind.CallOpts, _task Task) ([32]byte, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "hashTaskWithExpiry", _task)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashTaskWithExpiry is a free data retrieval call binding the contract method 0x949528fc.
//
// Solidity: function hashTaskWithExpiry((string,address,address,uint256,bytes,string,uint32,uint256) _task) pure returns(bytes32)
func (_ServiceManager *ServiceManagerSession) HashTaskWithExpiry(_task Task) ([32]byte, error) {
	return _ServiceManager.Contract.HashTaskWithExpiry(&_ServiceManager.CallOpts, _task)
}

// HashTaskWithExpiry is a free data retrieval call binding the contract method 0x949528fc.
//
// Solidity: function hashTaskWithExpiry((string,address,address,uint256,bytes,string,uint32,uint256) _task) pure returns(bytes32)
func (_ServiceManager *ServiceManagerCallerSession) HashTaskWithExpiry(_task Task) ([32]byte, error) {
	return _ServiceManager.Contract.HashTaskWithExpiry(&_ServiceManager.CallOpts, _task)
}

// IdToPolicy is a free data retrieval call binding the contract method 0xd9d4e99f.
//
// Solidity: function idToPolicy(string ) view returns(string)
func (_ServiceManager *ServiceManagerCaller) IdToPolicy(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "idToPolicy", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IdToPolicy is a free data retrieval call binding the contract method 0xd9d4e99f.
//
// Solidity: function idToPolicy(string ) view returns(string)
func (_ServiceManager *ServiceManagerSession) IdToPolicy(arg0 string) (string, error) {
	return _ServiceManager.Contract.IdToPolicy(&_ServiceManager.CallOpts, arg0)
}

// IdToPolicy is a free data retrieval call binding the contract method 0xd9d4e99f.
//
// Solidity: function idToPolicy(string ) view returns(string)
func (_ServiceManager *ServiceManagerCallerSession) IdToPolicy(arg0 string) (string, error) {
	return _ServiceManager.Contract.IdToPolicy(&_ServiceManager.CallOpts, arg0)
}

// IdToSocialGraph is a free data retrieval call binding the contract method 0xddf6a51b.
//
// Solidity: function idToSocialGraph(string ) view returns(string)
func (_ServiceManager *ServiceManagerCaller) IdToSocialGraph(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "idToSocialGraph", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// IdToSocialGraph is a free data retrieval call binding the contract method 0xddf6a51b.
//
// Solidity: function idToSocialGraph(string ) view returns(string)
func (_ServiceManager *ServiceManagerSession) IdToSocialGraph(arg0 string) (string, error) {
	return _ServiceManager.Contract.IdToSocialGraph(&_ServiceManager.CallOpts, arg0)
}

// IdToSocialGraph is a free data retrieval call binding the contract method 0xddf6a51b.
//
// Solidity: function idToSocialGraph(string ) view returns(string)
func (_ServiceManager *ServiceManagerCallerSession) IdToSocialGraph(arg0 string) (string, error) {
	return _ServiceManager.Contract.IdToSocialGraph(&_ServiceManager.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 totalStake, uint8 status)
func (_ServiceManager *ServiceManagerCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalStake *big.Int
	Status     uint8
}, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "operators", arg0)

	outstruct := new(struct {
		TotalStake *big.Int
		Status     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalStake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 totalStake, uint8 status)
func (_ServiceManager *ServiceManagerSession) Operators(arg0 common.Address) (struct {
	TotalStake *big.Int
	Status     uint8
}, error) {
	return _ServiceManager.Contract.Operators(&_ServiceManager.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256 totalStake, uint8 status)
func (_ServiceManager *ServiceManagerCallerSession) Operators(arg0 common.Address) (struct {
	TotalStake *big.Int
	Status     uint8
}, error) {
	return _ServiceManager.Contract.Operators(&_ServiceManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceManager *ServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceManager *ServiceManagerSession) Owner() (common.Address, error) {
	return _ServiceManager.Contract.Owner(&_ServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceManager *ServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ServiceManager.Contract.Owner(&_ServiceManager.CallOpts)
}

// SigningKeyToOperator is a free data retrieval call binding the contract method 0x0ff26fd1.
//
// Solidity: function signingKeyToOperator(address ) view returns(address)
func (_ServiceManager *ServiceManagerCaller) SigningKeyToOperator(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "signingKeyToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigningKeyToOperator is a free data retrieval call binding the contract method 0x0ff26fd1.
//
// Solidity: function signingKeyToOperator(address ) view returns(address)
func (_ServiceManager *ServiceManagerSession) SigningKeyToOperator(arg0 common.Address) (common.Address, error) {
	return _ServiceManager.Contract.SigningKeyToOperator(&_ServiceManager.CallOpts, arg0)
}

// SigningKeyToOperator is a free data retrieval call binding the contract method 0x0ff26fd1.
//
// Solidity: function signingKeyToOperator(address ) view returns(address)
func (_ServiceManager *ServiceManagerCallerSession) SigningKeyToOperator(arg0 common.Address) (common.Address, error) {
	return _ServiceManager.Contract.SigningKeyToOperator(&_ServiceManager.CallOpts, arg0)
}

// SocialGraphIDs is a free data retrieval call binding the contract method 0xdf935065.
//
// Solidity: function socialGraphIDs(uint256 ) view returns(string)
func (_ServiceManager *ServiceManagerCaller) SocialGraphIDs(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "socialGraphIDs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SocialGraphIDs is a free data retrieval call binding the contract method 0xdf935065.
//
// Solidity: function socialGraphIDs(uint256 ) view returns(string)
func (_ServiceManager *ServiceManagerSession) SocialGraphIDs(arg0 *big.Int) (string, error) {
	return _ServiceManager.Contract.SocialGraphIDs(&_ServiceManager.CallOpts, arg0)
}

// SocialGraphIDs is a free data retrieval call binding the contract method 0xdf935065.
//
// Solidity: function socialGraphIDs(uint256 ) view returns(string)
func (_ServiceManager *ServiceManagerCallerSession) SocialGraphIDs(arg0 *big.Int) (string, error) {
	return _ServiceManager.Contract.SocialGraphIDs(&_ServiceManager.CallOpts, arg0)
}

// SpentTaskIds is a free data retrieval call binding the contract method 0x749dccc7.
//
// Solidity: function spentTaskIds(string ) view returns(bool)
func (_ServiceManager *ServiceManagerCaller) SpentTaskIds(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "spentTaskIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpentTaskIds is a free data retrieval call binding the contract method 0x749dccc7.
//
// Solidity: function spentTaskIds(string ) view returns(bool)
func (_ServiceManager *ServiceManagerSession) SpentTaskIds(arg0 string) (bool, error) {
	return _ServiceManager.Contract.SpentTaskIds(&_ServiceManager.CallOpts, arg0)
}

// SpentTaskIds is a free data retrieval call binding the contract method 0x749dccc7.
//
// Solidity: function spentTaskIds(string ) view returns(bool)
func (_ServiceManager *ServiceManagerCallerSession) SpentTaskIds(arg0 string) (bool, error) {
	return _ServiceManager.Contract.SpentTaskIds(&_ServiceManager.CallOpts, arg0)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ServiceManager *ServiceManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ServiceManager *ServiceManagerSession) StakeRegistry() (common.Address, error) {
	return _ServiceManager.Contract.StakeRegistry(&_ServiceManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ServiceManager *ServiceManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ServiceManager.Contract.StakeRegistry(&_ServiceManager.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0xd574ea3d.
//
// Solidity: function strategies(uint256 ) view returns(address)
func (_ServiceManager *ServiceManagerCaller) Strategies(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "strategies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0xd574ea3d.
//
// Solidity: function strategies(uint256 ) view returns(address)
func (_ServiceManager *ServiceManagerSession) Strategies(arg0 *big.Int) (common.Address, error) {
	return _ServiceManager.Contract.Strategies(&_ServiceManager.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0xd574ea3d.
//
// Solidity: function strategies(uint256 ) view returns(address)
func (_ServiceManager *ServiceManagerCallerSession) Strategies(arg0 *big.Int) (common.Address, error) {
	return _ServiceManager.Contract.Strategies(&_ServiceManager.CallOpts, arg0)
}

// ThresholdStake is a free data retrieval call binding the contract method 0x20af7390.
//
// Solidity: function thresholdStake() view returns(uint256)
func (_ServiceManager *ServiceManagerCaller) ThresholdStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ServiceManager.contract.Call(opts, &out, "thresholdStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThresholdStake is a free data retrieval call binding the contract method 0x20af7390.
//
// Solidity: function thresholdStake() view returns(uint256)
func (_ServiceManager *ServiceManagerSession) ThresholdStake() (*big.Int, error) {
	return _ServiceManager.Contract.ThresholdStake(&_ServiceManager.CallOpts)
}

// ThresholdStake is a free data retrieval call binding the contract method 0x20af7390.
//
// Solidity: function thresholdStake() view returns(uint256)
func (_ServiceManager *ServiceManagerCallerSession) ThresholdStake() (*big.Int, error) {
	return _ServiceManager.Contract.ThresholdStake(&_ServiceManager.CallOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x66f17e73.
//
// Solidity: function addStrategy(address _strategy, uint8 quorumNumber, uint256 index) returns()
func (_ServiceManager *ServiceManagerTransactor) AddStrategy(opts *bind.TransactOpts, _strategy common.Address, quorumNumber uint8, index *big.Int) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "addStrategy", _strategy, quorumNumber, index)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x66f17e73.
//
// Solidity: function addStrategy(address _strategy, uint8 quorumNumber, uint256 index) returns()
func (_ServiceManager *ServiceManagerSession) AddStrategy(_strategy common.Address, quorumNumber uint8, index *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.AddStrategy(&_ServiceManager.TransactOpts, _strategy, quorumNumber, index)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x66f17e73.
//
// Solidity: function addStrategy(address _strategy, uint8 quorumNumber, uint256 index) returns()
func (_ServiceManager *ServiceManagerTransactorSession) AddStrategy(_strategy common.Address, quorumNumber uint8, index *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.AddStrategy(&_ServiceManager.TransactOpts, _strategy, quorumNumber, index)
}

// DeployPolicy is a paid mutator transaction binding the contract method 0xc0443c5f.
//
// Solidity: function deployPolicy(string _policyID, string _policy, uint256 _quorumThreshold) returns()
func (_ServiceManager *ServiceManagerTransactor) DeployPolicy(opts *bind.TransactOpts, _policyID string, _policy string, _quorumThreshold *big.Int) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "deployPolicy", _policyID, _policy, _quorumThreshold)
}

// DeployPolicy is a paid mutator transaction binding the contract method 0xc0443c5f.
//
// Solidity: function deployPolicy(string _policyID, string _policy, uint256 _quorumThreshold) returns()
func (_ServiceManager *ServiceManagerSession) DeployPolicy(_policyID string, _policy string, _quorumThreshold *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.DeployPolicy(&_ServiceManager.TransactOpts, _policyID, _policy, _quorumThreshold)
}

// DeployPolicy is a paid mutator transaction binding the contract method 0xc0443c5f.
//
// Solidity: function deployPolicy(string _policyID, string _policy, uint256 _quorumThreshold) returns()
func (_ServiceManager *ServiceManagerTransactorSession) DeployPolicy(_policyID string, _policy string, _quorumThreshold *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.DeployPolicy(&_ServiceManager.TransactOpts, _policyID, _policy, _quorumThreshold)
}

// DeploySocialGraph is a paid mutator transaction binding the contract method 0xd20e78c9.
//
// Solidity: function deploySocialGraph(string _socialGraphID, string _socialGraphConfig) returns()
func (_ServiceManager *ServiceManagerTransactor) DeploySocialGraph(opts *bind.TransactOpts, _socialGraphID string, _socialGraphConfig string) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "deploySocialGraph", _socialGraphID, _socialGraphConfig)
}

// DeploySocialGraph is a paid mutator transaction binding the contract method 0xd20e78c9.
//
// Solidity: function deploySocialGraph(string _socialGraphID, string _socialGraphConfig) returns()
func (_ServiceManager *ServiceManagerSession) DeploySocialGraph(_socialGraphID string, _socialGraphConfig string) (*types.Transaction, error) {
	return _ServiceManager.Contract.DeploySocialGraph(&_ServiceManager.TransactOpts, _socialGraphID, _socialGraphConfig)
}

// DeploySocialGraph is a paid mutator transaction binding the contract method 0xd20e78c9.
//
// Solidity: function deploySocialGraph(string _socialGraphID, string _socialGraphConfig) returns()
func (_ServiceManager *ServiceManagerTransactorSession) DeploySocialGraph(_socialGraphID string, _socialGraphConfig string) (*types.Transaction, error) {
	return _ServiceManager.Contract.DeploySocialGraph(&_ServiceManager.TransactOpts, _socialGraphID, _socialGraphConfig)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address _operator) returns()
func (_ServiceManager *ServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", _operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address _operator) returns()
func (_ServiceManager *ServiceManagerSession) DeregisterOperatorFromAVS(_operator common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.DeregisterOperatorFromAVS(&_ServiceManager.TransactOpts, _operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address _operator) returns()
func (_ServiceManager *ServiceManagerTransactorSession) DeregisterOperatorFromAVS(_operator common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.DeregisterOperatorFromAVS(&_ServiceManager.TransactOpts, _operator)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address _owner, address _aggregator, address _delegationManager, address _stakeRegistry, address _avsDirectory, uint256 _thresholdStake) returns()
func (_ServiceManager *ServiceManagerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _aggregator common.Address, _delegationManager common.Address, _stakeRegistry common.Address, _avsDirectory common.Address, _thresholdStake *big.Int) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "initialize", _owner, _aggregator, _delegationManager, _stakeRegistry, _avsDirectory, _thresholdStake)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address _owner, address _aggregator, address _delegationManager, address _stakeRegistry, address _avsDirectory, uint256 _thresholdStake) returns()
func (_ServiceManager *ServiceManagerSession) Initialize(_owner common.Address, _aggregator common.Address, _delegationManager common.Address, _stakeRegistry common.Address, _avsDirectory common.Address, _thresholdStake *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.Initialize(&_ServiceManager.TransactOpts, _owner, _aggregator, _delegationManager, _stakeRegistry, _avsDirectory, _thresholdStake)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address _owner, address _aggregator, address _delegationManager, address _stakeRegistry, address _avsDirectory, uint256 _thresholdStake) returns()
func (_ServiceManager *ServiceManagerTransactorSession) Initialize(_owner common.Address, _aggregator common.Address, _delegationManager common.Address, _stakeRegistry common.Address, _avsDirectory common.Address, _thresholdStake *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.Initialize(&_ServiceManager.TransactOpts, _owner, _aggregator, _delegationManager, _stakeRegistry, _avsDirectory, _thresholdStake)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address _operatorSigningKey, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_ServiceManager *ServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, _operatorSigningKey common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "registerOperatorToAVS", _operatorSigningKey, _operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address _operatorSigningKey, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_ServiceManager *ServiceManagerSession) RegisterOperatorToAVS(_operatorSigningKey common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ServiceManager.Contract.RegisterOperatorToAVS(&_ServiceManager.TransactOpts, _operatorSigningKey, _operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address _operatorSigningKey, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_ServiceManager *ServiceManagerTransactorSession) RegisterOperatorToAVS(_operatorSigningKey common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ServiceManager.Contract.RegisterOperatorToAVS(&_ServiceManager.TransactOpts, _operatorSigningKey, _operatorSignature)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x34099ba1.
//
// Solidity: function removePolicy(string _policyID) returns()
func (_ServiceManager *ServiceManagerTransactor) RemovePolicy(opts *bind.TransactOpts, _policyID string) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "removePolicy", _policyID)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x34099ba1.
//
// Solidity: function removePolicy(string _policyID) returns()
func (_ServiceManager *ServiceManagerSession) RemovePolicy(_policyID string) (*types.Transaction, error) {
	return _ServiceManager.Contract.RemovePolicy(&_ServiceManager.TransactOpts, _policyID)
}

// RemovePolicy is a paid mutator transaction binding the contract method 0x34099ba1.
//
// Solidity: function removePolicy(string _policyID) returns()
func (_ServiceManager *ServiceManagerTransactorSession) RemovePolicy(_policyID string) (*types.Transaction, error) {
	return _ServiceManager.Contract.RemovePolicy(&_ServiceManager.TransactOpts, _policyID)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_ServiceManager *ServiceManagerTransactor) RemoveStrategy(opts *bind.TransactOpts, _strategy common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "removeStrategy", _strategy)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_ServiceManager *ServiceManagerSession) RemoveStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.RemoveStrategy(&_ServiceManager.TransactOpts, _strategy)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0x175188e8.
//
// Solidity: function removeStrategy(address _strategy) returns()
func (_ServiceManager *ServiceManagerTransactorSession) RemoveStrategy(_strategy common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.RemoveStrategy(&_ServiceManager.TransactOpts, _strategy)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceManager *ServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceManager *ServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceManager.Contract.RenounceOwnership(&_ServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceManager *ServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceManager.Contract.RenounceOwnership(&_ServiceManager.TransactOpts)
}

// RotatePredicateSigningKey is a paid mutator transaction binding the contract method 0x001ba1eb.
//
// Solidity: function rotatePredicateSigningKey(address _oldSigningKey, address _newSigningKey) returns()
func (_ServiceManager *ServiceManagerTransactor) RotatePredicateSigningKey(opts *bind.TransactOpts, _oldSigningKey common.Address, _newSigningKey common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "rotatePredicateSigningKey", _oldSigningKey, _newSigningKey)
}

// RotatePredicateSigningKey is a paid mutator transaction binding the contract method 0x001ba1eb.
//
// Solidity: function rotatePredicateSigningKey(address _oldSigningKey, address _newSigningKey) returns()
func (_ServiceManager *ServiceManagerSession) RotatePredicateSigningKey(_oldSigningKey common.Address, _newSigningKey common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.RotatePredicateSigningKey(&_ServiceManager.TransactOpts, _oldSigningKey, _newSigningKey)
}

// RotatePredicateSigningKey is a paid mutator transaction binding the contract method 0x001ba1eb.
//
// Solidity: function rotatePredicateSigningKey(address _oldSigningKey, address _newSigningKey) returns()
func (_ServiceManager *ServiceManagerTransactorSession) RotatePredicateSigningKey(_oldSigningKey common.Address, _newSigningKey common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.RotatePredicateSigningKey(&_ServiceManager.TransactOpts, _oldSigningKey, _newSigningKey)
}

// SetAVSDirectory is a paid mutator transaction binding the contract method 0x862621ef.
//
// Solidity: function setAVSDirectory(address _avsDirectory) returns()
func (_ServiceManager *ServiceManagerTransactor) SetAVSDirectory(opts *bind.TransactOpts, _avsDirectory common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "setAVSDirectory", _avsDirectory)
}

// SetAVSDirectory is a paid mutator transaction binding the contract method 0x862621ef.
//
// Solidity: function setAVSDirectory(address _avsDirectory) returns()
func (_ServiceManager *ServiceManagerSession) SetAVSDirectory(_avsDirectory common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetAVSDirectory(&_ServiceManager.TransactOpts, _avsDirectory)
}

// SetAVSDirectory is a paid mutator transaction binding the contract method 0x862621ef.
//
// Solidity: function setAVSDirectory(address _avsDirectory) returns()
func (_ServiceManager *ServiceManagerTransactorSession) SetAVSDirectory(_avsDirectory common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetAVSDirectory(&_ServiceManager.TransactOpts, _avsDirectory)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_ServiceManager *ServiceManagerTransactor) SetAggregator(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "setAggregator", _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_ServiceManager *ServiceManagerSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetAggregator(&_ServiceManager.TransactOpts, _aggregator)
}

// SetAggregator is a paid mutator transaction binding the contract method 0xf9120af6.
//
// Solidity: function setAggregator(address _aggregator) returns()
func (_ServiceManager *ServiceManagerTransactorSession) SetAggregator(_aggregator common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetAggregator(&_ServiceManager.TransactOpts, _aggregator)
}

// SetDelegationManager is a paid mutator transaction binding the contract method 0x1a8d0de2.
//
// Solidity: function setDelegationManager(address _delegationManager) returns()
func (_ServiceManager *ServiceManagerTransactor) SetDelegationManager(opts *bind.TransactOpts, _delegationManager common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "setDelegationManager", _delegationManager)
}

// SetDelegationManager is a paid mutator transaction binding the contract method 0x1a8d0de2.
//
// Solidity: function setDelegationManager(address _delegationManager) returns()
func (_ServiceManager *ServiceManagerSession) SetDelegationManager(_delegationManager common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetDelegationManager(&_ServiceManager.TransactOpts, _delegationManager)
}

// SetDelegationManager is a paid mutator transaction binding the contract method 0x1a8d0de2.
//
// Solidity: function setDelegationManager(address _delegationManager) returns()
func (_ServiceManager *ServiceManagerTransactorSession) SetDelegationManager(_delegationManager common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetDelegationManager(&_ServiceManager.TransactOpts, _delegationManager)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ServiceManager *ServiceManagerTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ServiceManager *ServiceManagerSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetMetadataURI(&_ServiceManager.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ServiceManager *ServiceManagerTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetMetadataURI(&_ServiceManager.TransactOpts, _metadataURI)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x6b4c991b.
//
// Solidity: function setPolicy(string _policyID) returns()
func (_ServiceManager *ServiceManagerTransactor) SetPolicy(opts *bind.TransactOpts, _policyID string) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "setPolicy", _policyID)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x6b4c991b.
//
// Solidity: function setPolicy(string _policyID) returns()
func (_ServiceManager *ServiceManagerSession) SetPolicy(_policyID string) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetPolicy(&_ServiceManager.TransactOpts, _policyID)
}

// SetPolicy is a paid mutator transaction binding the contract method 0x6b4c991b.
//
// Solidity: function setPolicy(string _policyID) returns()
func (_ServiceManager *ServiceManagerTransactorSession) SetPolicy(_policyID string) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetPolicy(&_ServiceManager.TransactOpts, _policyID)
}

// SetStakeRegistry is a paid mutator transaction binding the contract method 0xe3b05f2f.
//
// Solidity: function setStakeRegistry(address _stakeRegistry) returns()
func (_ServiceManager *ServiceManagerTransactor) SetStakeRegistry(opts *bind.TransactOpts, _stakeRegistry common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "setStakeRegistry", _stakeRegistry)
}

// SetStakeRegistry is a paid mutator transaction binding the contract method 0xe3b05f2f.
//
// Solidity: function setStakeRegistry(address _stakeRegistry) returns()
func (_ServiceManager *ServiceManagerSession) SetStakeRegistry(_stakeRegistry common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetStakeRegistry(&_ServiceManager.TransactOpts, _stakeRegistry)
}

// SetStakeRegistry is a paid mutator transaction binding the contract method 0xe3b05f2f.
//
// Solidity: function setStakeRegistry(address _stakeRegistry) returns()
func (_ServiceManager *ServiceManagerTransactorSession) SetStakeRegistry(_stakeRegistry common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetStakeRegistry(&_ServiceManager.TransactOpts, _stakeRegistry)
}

// SetThresholdStake is a paid mutator transaction binding the contract method 0x8ad96602.
//
// Solidity: function setThresholdStake(uint256 _thresholdStake) returns()
func (_ServiceManager *ServiceManagerTransactor) SetThresholdStake(opts *bind.TransactOpts, _thresholdStake *big.Int) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "setThresholdStake", _thresholdStake)
}

// SetThresholdStake is a paid mutator transaction binding the contract method 0x8ad96602.
//
// Solidity: function setThresholdStake(uint256 _thresholdStake) returns()
func (_ServiceManager *ServiceManagerSession) SetThresholdStake(_thresholdStake *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetThresholdStake(&_ServiceManager.TransactOpts, _thresholdStake)
}

// SetThresholdStake is a paid mutator transaction binding the contract method 0x8ad96602.
//
// Solidity: function setThresholdStake(uint256 _thresholdStake) returns()
func (_ServiceManager *ServiceManagerTransactorSession) SetThresholdStake(_thresholdStake *big.Int) (*types.Transaction, error) {
	return _ServiceManager.Contract.SetThresholdStake(&_ServiceManager.TransactOpts, _thresholdStake)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceManager *ServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceManager *ServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.TransferOwnership(&_ServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceManager *ServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceManager.Contract.TransferOwnership(&_ServiceManager.TransactOpts, newOwner)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ServiceManager *ServiceManagerTransactor) UpdateOperatorsForQuorum(opts *bind.TransactOpts, operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "updateOperatorsForQuorum", operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ServiceManager *ServiceManagerSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ServiceManager.Contract.UpdateOperatorsForQuorum(&_ServiceManager.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// UpdateOperatorsForQuorum is a paid mutator transaction binding the contract method 0x5140a548.
//
// Solidity: function updateOperatorsForQuorum(address[][] operatorsPerQuorum, bytes quorumNumbers) returns()
func (_ServiceManager *ServiceManagerTransactorSession) UpdateOperatorsForQuorum(operatorsPerQuorum [][]common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ServiceManager.Contract.UpdateOperatorsForQuorum(&_ServiceManager.TransactOpts, operatorsPerQuorum, quorumNumbers)
}

// ValidateSignatures is a paid mutator transaction binding the contract method 0x18cea58d.
//
// Solidity: function validateSignatures((string,address,address,uint256,bytes,string,uint32,uint256) _task, address[] signerAddresses, bytes[] signatures) returns(bool isVerified)
func (_ServiceManager *ServiceManagerTransactor) ValidateSignatures(opts *bind.TransactOpts, _task Task, signerAddresses []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _ServiceManager.contract.Transact(opts, "validateSignatures", _task, signerAddresses, signatures)
}

// ValidateSignatures is a paid mutator transaction binding the contract method 0x18cea58d.
//
// Solidity: function validateSignatures((string,address,address,uint256,bytes,string,uint32,uint256) _task, address[] signerAddresses, bytes[] signatures) returns(bool isVerified)
func (_ServiceManager *ServiceManagerSession) ValidateSignatures(_task Task, signerAddresses []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _ServiceManager.Contract.ValidateSignatures(&_ServiceManager.TransactOpts, _task, signerAddresses, signatures)
}

// ValidateSignatures is a paid mutator transaction binding the contract method 0x18cea58d.
//
// Solidity: function validateSignatures((string,address,address,uint256,bytes,string,uint32,uint256) _task, address[] signerAddresses, bytes[] signatures) returns(bool isVerified)
func (_ServiceManager *ServiceManagerTransactorSession) ValidateSignatures(_task Task, signerAddresses []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _ServiceManager.Contract.ValidateSignatures(&_ServiceManager.TransactOpts, _task, signerAddresses, signatures)
}

// ServiceManagerAVSDirectoryUpdatedIterator is returned from FilterAVSDirectoryUpdated and is used to iterate over the raw logs and unpacked data for AVSDirectoryUpdated events raised by the ServiceManager contract.
type ServiceManagerAVSDirectoryUpdatedIterator struct {
	Event *ServiceManagerAVSDirectoryUpdated // Event containing the contract specifics and raw log

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
func (it *ServiceManagerAVSDirectoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerAVSDirectoryUpdated)
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
		it.Event = new(ServiceManagerAVSDirectoryUpdated)
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
func (it *ServiceManagerAVSDirectoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerAVSDirectoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerAVSDirectoryUpdated represents a AVSDirectoryUpdated event raised by the ServiceManager contract.
type ServiceManagerAVSDirectoryUpdated struct {
	AvsDirectory common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAVSDirectoryUpdated is a free log retrieval operation binding the contract event 0x2a623245c0f1c5741f2a4c247a58842872f1fdf8a31e66d031dd1cd1532e89a7.
//
// Solidity: event AVSDirectoryUpdated(address indexed avsDirectory)
func (_ServiceManager *ServiceManagerFilterer) FilterAVSDirectoryUpdated(opts *bind.FilterOpts, avsDirectory []common.Address) (*ServiceManagerAVSDirectoryUpdatedIterator, error) {

	var avsDirectoryRule []interface{}
	for _, avsDirectoryItem := range avsDirectory {
		avsDirectoryRule = append(avsDirectoryRule, avsDirectoryItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "AVSDirectoryUpdated", avsDirectoryRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerAVSDirectoryUpdatedIterator{contract: _ServiceManager.contract, event: "AVSDirectoryUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSDirectoryUpdated is a free log subscription operation binding the contract event 0x2a623245c0f1c5741f2a4c247a58842872f1fdf8a31e66d031dd1cd1532e89a7.
//
// Solidity: event AVSDirectoryUpdated(address indexed avsDirectory)
func (_ServiceManager *ServiceManagerFilterer) WatchAVSDirectoryUpdated(opts *bind.WatchOpts, sink chan<- *ServiceManagerAVSDirectoryUpdated, avsDirectory []common.Address) (event.Subscription, error) {

	var avsDirectoryRule []interface{}
	for _, avsDirectoryItem := range avsDirectory {
		avsDirectoryRule = append(avsDirectoryRule, avsDirectoryItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "AVSDirectoryUpdated", avsDirectoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerAVSDirectoryUpdated)
				if err := _ServiceManager.contract.UnpackLog(event, "AVSDirectoryUpdated", log); err != nil {
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

// ParseAVSDirectoryUpdated is a log parse operation binding the contract event 0x2a623245c0f1c5741f2a4c247a58842872f1fdf8a31e66d031dd1cd1532e89a7.
//
// Solidity: event AVSDirectoryUpdated(address indexed avsDirectory)
func (_ServiceManager *ServiceManagerFilterer) ParseAVSDirectoryUpdated(log types.Log) (*ServiceManagerAVSDirectoryUpdated, error) {
	event := new(ServiceManagerAVSDirectoryUpdated)
	if err := _ServiceManager.contract.UnpackLog(event, "AVSDirectoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the ServiceManager contract.
type ServiceManagerAggregatorUpdatedIterator struct {
	Event *ServiceManagerAggregatorUpdated // Event containing the contract specifics and raw log

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
func (it *ServiceManagerAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerAggregatorUpdated)
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
		it.Event = new(ServiceManagerAggregatorUpdated)
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
func (it *ServiceManagerAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerAggregatorUpdated represents a AggregatorUpdated event raised by the ServiceManager contract.
type ServiceManagerAggregatorUpdated struct {
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address indexed aggregator)
func (_ServiceManager *ServiceManagerFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts, aggregator []common.Address) (*ServiceManagerAggregatorUpdatedIterator, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "AggregatorUpdated", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerAggregatorUpdatedIterator{contract: _ServiceManager.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address indexed aggregator)
func (_ServiceManager *ServiceManagerFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *ServiceManagerAggregatorUpdated, aggregator []common.Address) (event.Subscription, error) {

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "AggregatorUpdated", aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerAggregatorUpdated)
				if err := _ServiceManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
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

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x602cec4b1583b07d071161da5eb9589444d2459201e2fab7753dc941e9351c21.
//
// Solidity: event AggregatorUpdated(address indexed aggregator)
func (_ServiceManager *ServiceManagerFilterer) ParseAggregatorUpdated(log types.Log) (*ServiceManagerAggregatorUpdated, error) {
	event := new(ServiceManagerAggregatorUpdated)
	if err := _ServiceManager.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerDelegationManagerUpdatedIterator is returned from FilterDelegationManagerUpdated and is used to iterate over the raw logs and unpacked data for DelegationManagerUpdated events raised by the ServiceManager contract.
type ServiceManagerDelegationManagerUpdatedIterator struct {
	Event *ServiceManagerDelegationManagerUpdated // Event containing the contract specifics and raw log

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
func (it *ServiceManagerDelegationManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerDelegationManagerUpdated)
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
		it.Event = new(ServiceManagerDelegationManagerUpdated)
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
func (it *ServiceManagerDelegationManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerDelegationManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerDelegationManagerUpdated represents a DelegationManagerUpdated event raised by the ServiceManager contract.
type ServiceManagerDelegationManagerUpdated struct {
	DelegationManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelegationManagerUpdated is a free log retrieval operation binding the contract event 0xa1b8fcd417a2bb56a91d1fc6708faf8283b5730e55821393e70303e544aeec92.
//
// Solidity: event DelegationManagerUpdated(address indexed delegationManager)
func (_ServiceManager *ServiceManagerFilterer) FilterDelegationManagerUpdated(opts *bind.FilterOpts, delegationManager []common.Address) (*ServiceManagerDelegationManagerUpdatedIterator, error) {

	var delegationManagerRule []interface{}
	for _, delegationManagerItem := range delegationManager {
		delegationManagerRule = append(delegationManagerRule, delegationManagerItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "DelegationManagerUpdated", delegationManagerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerDelegationManagerUpdatedIterator{contract: _ServiceManager.contract, event: "DelegationManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationManagerUpdated is a free log subscription operation binding the contract event 0xa1b8fcd417a2bb56a91d1fc6708faf8283b5730e55821393e70303e544aeec92.
//
// Solidity: event DelegationManagerUpdated(address indexed delegationManager)
func (_ServiceManager *ServiceManagerFilterer) WatchDelegationManagerUpdated(opts *bind.WatchOpts, sink chan<- *ServiceManagerDelegationManagerUpdated, delegationManager []common.Address) (event.Subscription, error) {

	var delegationManagerRule []interface{}
	for _, delegationManagerItem := range delegationManager {
		delegationManagerRule = append(delegationManagerRule, delegationManagerItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "DelegationManagerUpdated", delegationManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerDelegationManagerUpdated)
				if err := _ServiceManager.contract.UnpackLog(event, "DelegationManagerUpdated", log); err != nil {
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

// ParseDelegationManagerUpdated is a log parse operation binding the contract event 0xa1b8fcd417a2bb56a91d1fc6708faf8283b5730e55821393e70303e544aeec92.
//
// Solidity: event DelegationManagerUpdated(address indexed delegationManager)
func (_ServiceManager *ServiceManagerFilterer) ParseDelegationManagerUpdated(log types.Log) (*ServiceManagerDelegationManagerUpdated, error) {
	event := new(ServiceManagerDelegationManagerUpdated)
	if err := _ServiceManager.contract.UnpackLog(event, "DelegationManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerDeployedPolicyIterator is returned from FilterDeployedPolicy and is used to iterate over the raw logs and unpacked data for DeployedPolicy events raised by the ServiceManager contract.
type ServiceManagerDeployedPolicyIterator struct {
	Event *ServiceManagerDeployedPolicy // Event containing the contract specifics and raw log

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
func (it *ServiceManagerDeployedPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerDeployedPolicy)
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
		it.Event = new(ServiceManagerDeployedPolicy)
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
func (it *ServiceManagerDeployedPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerDeployedPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerDeployedPolicy represents a DeployedPolicy event raised by the ServiceManager contract.
type ServiceManagerDeployedPolicy struct {
	PolicyID common.Hash
	Policy   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeployedPolicy is a free log retrieval operation binding the contract event 0xb6487025b06543ad686bdaa829e2b07863fd163cacd2e0f031340308ec09584e.
//
// Solidity: event DeployedPolicy(string indexed policyID, string policy)
func (_ServiceManager *ServiceManagerFilterer) FilterDeployedPolicy(opts *bind.FilterOpts, policyID []string) (*ServiceManagerDeployedPolicyIterator, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "DeployedPolicy", policyIDRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerDeployedPolicyIterator{contract: _ServiceManager.contract, event: "DeployedPolicy", logs: logs, sub: sub}, nil
}

// WatchDeployedPolicy is a free log subscription operation binding the contract event 0xb6487025b06543ad686bdaa829e2b07863fd163cacd2e0f031340308ec09584e.
//
// Solidity: event DeployedPolicy(string indexed policyID, string policy)
func (_ServiceManager *ServiceManagerFilterer) WatchDeployedPolicy(opts *bind.WatchOpts, sink chan<- *ServiceManagerDeployedPolicy, policyID []string) (event.Subscription, error) {

	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "DeployedPolicy", policyIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerDeployedPolicy)
				if err := _ServiceManager.contract.UnpackLog(event, "DeployedPolicy", log); err != nil {
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

// ParseDeployedPolicy is a log parse operation binding the contract event 0xb6487025b06543ad686bdaa829e2b07863fd163cacd2e0f031340308ec09584e.
//
// Solidity: event DeployedPolicy(string indexed policyID, string policy)
func (_ServiceManager *ServiceManagerFilterer) ParseDeployedPolicy(log types.Log) (*ServiceManagerDeployedPolicy, error) {
	event := new(ServiceManagerDeployedPolicy)
	if err := _ServiceManager.contract.UnpackLog(event, "DeployedPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ServiceManager contract.
type ServiceManagerInitializedIterator struct {
	Event *ServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerInitialized)
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
		it.Event = new(ServiceManagerInitialized)
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
func (it *ServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerInitialized represents a Initialized event raised by the ServiceManager contract.
type ServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ServiceManager *ServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ServiceManagerInitializedIterator, error) {

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ServiceManagerInitializedIterator{contract: _ServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ServiceManager *ServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerInitialized)
				if err := _ServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ServiceManager *ServiceManagerFilterer) ParseInitialized(log types.Log) (*ServiceManagerInitialized, error) {
	event := new(ServiceManagerInitialized)
	if err := _ServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerNonCompliantTaskIterator is returned from FilterNonCompliantTask and is used to iterate over the raw logs and unpacked data for NonCompliantTask events raised by the ServiceManager contract.
type ServiceManagerNonCompliantTaskIterator struct {
	Event *ServiceManagerNonCompliantTask // Event containing the contract specifics and raw log

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
func (it *ServiceManagerNonCompliantTaskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerNonCompliantTask)
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
		it.Event = new(ServiceManagerNonCompliantTask)
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
func (it *ServiceManagerNonCompliantTaskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerNonCompliantTaskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerNonCompliantTask represents a NonCompliantTask event raised by the ServiceManager contract.
type ServiceManagerNonCompliantTask struct {
	TaskID *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNonCompliantTask is a free log retrieval operation binding the contract event 0x2a8b097d3dd38ed5fe131a385c16133baae1d9bd3d75f3576068b9211dfd7b5c.
//
// Solidity: event NonCompliantTask(uint256 indexed taskID)
func (_ServiceManager *ServiceManagerFilterer) FilterNonCompliantTask(opts *bind.FilterOpts, taskID []*big.Int) (*ServiceManagerNonCompliantTaskIterator, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "NonCompliantTask", taskIDRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerNonCompliantTaskIterator{contract: _ServiceManager.contract, event: "NonCompliantTask", logs: logs, sub: sub}, nil
}

// WatchNonCompliantTask is a free log subscription operation binding the contract event 0x2a8b097d3dd38ed5fe131a385c16133baae1d9bd3d75f3576068b9211dfd7b5c.
//
// Solidity: event NonCompliantTask(uint256 indexed taskID)
func (_ServiceManager *ServiceManagerFilterer) WatchNonCompliantTask(opts *bind.WatchOpts, sink chan<- *ServiceManagerNonCompliantTask, taskID []*big.Int) (event.Subscription, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "NonCompliantTask", taskIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerNonCompliantTask)
				if err := _ServiceManager.contract.UnpackLog(event, "NonCompliantTask", log); err != nil {
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

// ParseNonCompliantTask is a log parse operation binding the contract event 0x2a8b097d3dd38ed5fe131a385c16133baae1d9bd3d75f3576068b9211dfd7b5c.
//
// Solidity: event NonCompliantTask(uint256 indexed taskID)
func (_ServiceManager *ServiceManagerFilterer) ParseNonCompliantTask(log types.Log) (*ServiceManagerNonCompliantTask, error) {
	event := new(ServiceManagerNonCompliantTask)
	if err := _ServiceManager.contract.UnpackLog(event, "NonCompliantTask", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ServiceManager contract.
type ServiceManagerOperatorRegisteredIterator struct {
	Event *ServiceManagerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *ServiceManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerOperatorRegistered)
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
		it.Event = new(ServiceManagerOperatorRegistered)
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
func (it *ServiceManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerOperatorRegistered represents a OperatorRegistered event raised by the ServiceManager contract.
type ServiceManagerOperatorRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ServiceManager *ServiceManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ServiceManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerOperatorRegisteredIterator{contract: _ServiceManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ServiceManager *ServiceManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ServiceManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerOperatorRegistered)
				if err := _ServiceManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_ServiceManager *ServiceManagerFilterer) ParseOperatorRegistered(log types.Log) (*ServiceManagerOperatorRegistered, error) {
	event := new(ServiceManagerOperatorRegistered)
	if err := _ServiceManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the ServiceManager contract.
type ServiceManagerOperatorRemovedIterator struct {
	Event *ServiceManagerOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *ServiceManagerOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerOperatorRemoved)
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
		it.Event = new(ServiceManagerOperatorRemoved)
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
func (it *ServiceManagerOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerOperatorRemoved represents a OperatorRemoved event raised by the ServiceManager contract.
type ServiceManagerOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_ServiceManager *ServiceManagerFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operator []common.Address) (*ServiceManagerOperatorRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerOperatorRemovedIterator{contract: _ServiceManager.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_ServiceManager *ServiceManagerFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ServiceManagerOperatorRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "OperatorRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerOperatorRemoved)
				if err := _ServiceManager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed operator)
func (_ServiceManager *ServiceManagerFilterer) ParseOperatorRemoved(log types.Log) (*ServiceManagerOperatorRemoved, error) {
	event := new(ServiceManagerOperatorRemoved)
	if err := _ServiceManager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerOperatorsStakesUpdatedIterator is returned from FilterOperatorsStakesUpdated and is used to iterate over the raw logs and unpacked data for OperatorsStakesUpdated events raised by the ServiceManager contract.
type ServiceManagerOperatorsStakesUpdatedIterator struct {
	Event *ServiceManagerOperatorsStakesUpdated // Event containing the contract specifics and raw log

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
func (it *ServiceManagerOperatorsStakesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerOperatorsStakesUpdated)
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
		it.Event = new(ServiceManagerOperatorsStakesUpdated)
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
func (it *ServiceManagerOperatorsStakesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerOperatorsStakesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerOperatorsStakesUpdated represents a OperatorsStakesUpdated event raised by the ServiceManager contract.
type ServiceManagerOperatorsStakesUpdated struct {
	OperatorsPerQuorum [][]common.Address
	QuorumNumbers      common.Hash
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorsStakesUpdated is a free log retrieval operation binding the contract event 0xf3fcd2dadc5f6ebf9c1e9310a9e986a14079afc7ecf26784853ea9a3ac90721b.
//
// Solidity: event OperatorsStakesUpdated(address[][] indexed operatorsPerQuorum, bytes indexed quorumNumbers)
func (_ServiceManager *ServiceManagerFilterer) FilterOperatorsStakesUpdated(opts *bind.FilterOpts, operatorsPerQuorum [][][]common.Address, quorumNumbers [][]byte) (*ServiceManagerOperatorsStakesUpdatedIterator, error) {

	var operatorsPerQuorumRule []interface{}
	for _, operatorsPerQuorumItem := range operatorsPerQuorum {
		operatorsPerQuorumRule = append(operatorsPerQuorumRule, operatorsPerQuorumItem)
	}
	var quorumNumbersRule []interface{}
	for _, quorumNumbersItem := range quorumNumbers {
		quorumNumbersRule = append(quorumNumbersRule, quorumNumbersItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "OperatorsStakesUpdated", operatorsPerQuorumRule, quorumNumbersRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerOperatorsStakesUpdatedIterator{contract: _ServiceManager.contract, event: "OperatorsStakesUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorsStakesUpdated is a free log subscription operation binding the contract event 0xf3fcd2dadc5f6ebf9c1e9310a9e986a14079afc7ecf26784853ea9a3ac90721b.
//
// Solidity: event OperatorsStakesUpdated(address[][] indexed operatorsPerQuorum, bytes indexed quorumNumbers)
func (_ServiceManager *ServiceManagerFilterer) WatchOperatorsStakesUpdated(opts *bind.WatchOpts, sink chan<- *ServiceManagerOperatorsStakesUpdated, operatorsPerQuorum [][][]common.Address, quorumNumbers [][]byte) (event.Subscription, error) {

	var operatorsPerQuorumRule []interface{}
	for _, operatorsPerQuorumItem := range operatorsPerQuorum {
		operatorsPerQuorumRule = append(operatorsPerQuorumRule, operatorsPerQuorumItem)
	}
	var quorumNumbersRule []interface{}
	for _, quorumNumbersItem := range quorumNumbers {
		quorumNumbersRule = append(quorumNumbersRule, quorumNumbersItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "OperatorsStakesUpdated", operatorsPerQuorumRule, quorumNumbersRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerOperatorsStakesUpdated)
				if err := _ServiceManager.contract.UnpackLog(event, "OperatorsStakesUpdated", log); err != nil {
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

// ParseOperatorsStakesUpdated is a log parse operation binding the contract event 0xf3fcd2dadc5f6ebf9c1e9310a9e986a14079afc7ecf26784853ea9a3ac90721b.
//
// Solidity: event OperatorsStakesUpdated(address[][] indexed operatorsPerQuorum, bytes indexed quorumNumbers)
func (_ServiceManager *ServiceManagerFilterer) ParseOperatorsStakesUpdated(log types.Log) (*ServiceManagerOperatorsStakesUpdated, error) {
	event := new(ServiceManagerOperatorsStakesUpdated)
	if err := _ServiceManager.contract.UnpackLog(event, "OperatorsStakesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ServiceManager contract.
type ServiceManagerOwnershipTransferredIterator struct {
	Event *ServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerOwnershipTransferred)
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
		it.Event = new(ServiceManagerOwnershipTransferred)
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
func (it *ServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ServiceManager contract.
type ServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceManager *ServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerOwnershipTransferredIterator{contract: _ServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceManager *ServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerOwnershipTransferred)
				if err := _ServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ServiceManager *ServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ServiceManagerOwnershipTransferred, error) {
	event := new(ServiceManagerOwnershipTransferred)
	if err := _ServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerRemovedPolicyIterator is returned from FilterRemovedPolicy and is used to iterate over the raw logs and unpacked data for RemovedPolicy events raised by the ServiceManager contract.
type ServiceManagerRemovedPolicyIterator struct {
	Event *ServiceManagerRemovedPolicy // Event containing the contract specifics and raw log

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
func (it *ServiceManagerRemovedPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerRemovedPolicy)
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
		it.Event = new(ServiceManagerRemovedPolicy)
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
func (it *ServiceManagerRemovedPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerRemovedPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerRemovedPolicy represents a RemovedPolicy event raised by the ServiceManager contract.
type ServiceManagerRemovedPolicy struct {
	Client   common.Address
	PolicyID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemovedPolicy is a free log retrieval operation binding the contract event 0x64296597d208d70296fadc8eb749ff22ab651c8e6e455f35ee8ed66b74d1f775.
//
// Solidity: event RemovedPolicy(address indexed client, string indexed policyID)
func (_ServiceManager *ServiceManagerFilterer) FilterRemovedPolicy(opts *bind.FilterOpts, client []common.Address, policyID []string) (*ServiceManagerRemovedPolicyIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "RemovedPolicy", clientRule, policyIDRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerRemovedPolicyIterator{contract: _ServiceManager.contract, event: "RemovedPolicy", logs: logs, sub: sub}, nil
}

// WatchRemovedPolicy is a free log subscription operation binding the contract event 0x64296597d208d70296fadc8eb749ff22ab651c8e6e455f35ee8ed66b74d1f775.
//
// Solidity: event RemovedPolicy(address indexed client, string indexed policyID)
func (_ServiceManager *ServiceManagerFilterer) WatchRemovedPolicy(opts *bind.WatchOpts, sink chan<- *ServiceManagerRemovedPolicy, client []common.Address, policyID []string) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "RemovedPolicy", clientRule, policyIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerRemovedPolicy)
				if err := _ServiceManager.contract.UnpackLog(event, "RemovedPolicy", log); err != nil {
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

// ParseRemovedPolicy is a log parse operation binding the contract event 0x64296597d208d70296fadc8eb749ff22ab651c8e6e455f35ee8ed66b74d1f775.
//
// Solidity: event RemovedPolicy(address indexed client, string indexed policyID)
func (_ServiceManager *ServiceManagerFilterer) ParseRemovedPolicy(log types.Log) (*ServiceManagerRemovedPolicy, error) {
	event := new(ServiceManagerRemovedPolicy)
	if err := _ServiceManager.contract.UnpackLog(event, "RemovedPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerSetPolicyIterator is returned from FilterSetPolicy and is used to iterate over the raw logs and unpacked data for SetPolicy events raised by the ServiceManager contract.
type ServiceManagerSetPolicyIterator struct {
	Event *ServiceManagerSetPolicy // Event containing the contract specifics and raw log

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
func (it *ServiceManagerSetPolicyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerSetPolicy)
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
		it.Event = new(ServiceManagerSetPolicy)
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
func (it *ServiceManagerSetPolicyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerSetPolicyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerSetPolicy represents a SetPolicy event raised by the ServiceManager contract.
type ServiceManagerSetPolicy struct {
	Client   common.Address
	PolicyID common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPolicy is a free log retrieval operation binding the contract event 0xfbc30d1514eac402cb2045f1dd80ec75fbc997db6f719421b6d7490f4bfb779d.
//
// Solidity: event SetPolicy(address indexed client, string indexed policyID)
func (_ServiceManager *ServiceManagerFilterer) FilterSetPolicy(opts *bind.FilterOpts, client []common.Address, policyID []string) (*ServiceManagerSetPolicyIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "SetPolicy", clientRule, policyIDRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerSetPolicyIterator{contract: _ServiceManager.contract, event: "SetPolicy", logs: logs, sub: sub}, nil
}

// WatchSetPolicy is a free log subscription operation binding the contract event 0xfbc30d1514eac402cb2045f1dd80ec75fbc997db6f719421b6d7490f4bfb779d.
//
// Solidity: event SetPolicy(address indexed client, string indexed policyID)
func (_ServiceManager *ServiceManagerFilterer) WatchSetPolicy(opts *bind.WatchOpts, sink chan<- *ServiceManagerSetPolicy, client []common.Address, policyID []string) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var policyIDRule []interface{}
	for _, policyIDItem := range policyID {
		policyIDRule = append(policyIDRule, policyIDItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "SetPolicy", clientRule, policyIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerSetPolicy)
				if err := _ServiceManager.contract.UnpackLog(event, "SetPolicy", log); err != nil {
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

// ParseSetPolicy is a log parse operation binding the contract event 0xfbc30d1514eac402cb2045f1dd80ec75fbc997db6f719421b6d7490f4bfb779d.
//
// Solidity: event SetPolicy(address indexed client, string indexed policyID)
func (_ServiceManager *ServiceManagerFilterer) ParseSetPolicy(log types.Log) (*ServiceManagerSetPolicy, error) {
	event := new(ServiceManagerSetPolicy)
	if err := _ServiceManager.contract.UnpackLog(event, "SetPolicy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerSocialGraphDeployedIterator is returned from FilterSocialGraphDeployed and is used to iterate over the raw logs and unpacked data for SocialGraphDeployed events raised by the ServiceManager contract.
type ServiceManagerSocialGraphDeployedIterator struct {
	Event *ServiceManagerSocialGraphDeployed // Event containing the contract specifics and raw log

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
func (it *ServiceManagerSocialGraphDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerSocialGraphDeployed)
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
		it.Event = new(ServiceManagerSocialGraphDeployed)
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
func (it *ServiceManagerSocialGraphDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerSocialGraphDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerSocialGraphDeployed represents a SocialGraphDeployed event raised by the ServiceManager contract.
type ServiceManagerSocialGraphDeployed struct {
	SocialGraphID     common.Hash
	SocialGraphConfig string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSocialGraphDeployed is a free log retrieval operation binding the contract event 0x0547683ac2bbf227b34fe50fb3ef0f2b8a67dfad502dc794e7a69b0a8dd3e680.
//
// Solidity: event SocialGraphDeployed(string indexed socialGraphID, string socialGraphConfig)
func (_ServiceManager *ServiceManagerFilterer) FilterSocialGraphDeployed(opts *bind.FilterOpts, socialGraphID []string) (*ServiceManagerSocialGraphDeployedIterator, error) {

	var socialGraphIDRule []interface{}
	for _, socialGraphIDItem := range socialGraphID {
		socialGraphIDRule = append(socialGraphIDRule, socialGraphIDItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "SocialGraphDeployed", socialGraphIDRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerSocialGraphDeployedIterator{contract: _ServiceManager.contract, event: "SocialGraphDeployed", logs: logs, sub: sub}, nil
}

// WatchSocialGraphDeployed is a free log subscription operation binding the contract event 0x0547683ac2bbf227b34fe50fb3ef0f2b8a67dfad502dc794e7a69b0a8dd3e680.
//
// Solidity: event SocialGraphDeployed(string indexed socialGraphID, string socialGraphConfig)
func (_ServiceManager *ServiceManagerFilterer) WatchSocialGraphDeployed(opts *bind.WatchOpts, sink chan<- *ServiceManagerSocialGraphDeployed, socialGraphID []string) (event.Subscription, error) {

	var socialGraphIDRule []interface{}
	for _, socialGraphIDItem := range socialGraphID {
		socialGraphIDRule = append(socialGraphIDRule, socialGraphIDItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "SocialGraphDeployed", socialGraphIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerSocialGraphDeployed)
				if err := _ServiceManager.contract.UnpackLog(event, "SocialGraphDeployed", log); err != nil {
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

// ParseSocialGraphDeployed is a log parse operation binding the contract event 0x0547683ac2bbf227b34fe50fb3ef0f2b8a67dfad502dc794e7a69b0a8dd3e680.
//
// Solidity: event SocialGraphDeployed(string indexed socialGraphID, string socialGraphConfig)
func (_ServiceManager *ServiceManagerFilterer) ParseSocialGraphDeployed(log types.Log) (*ServiceManagerSocialGraphDeployed, error) {
	event := new(ServiceManagerSocialGraphDeployed)
	if err := _ServiceManager.contract.UnpackLog(event, "SocialGraphDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerStakeRegistryUpdatedIterator is returned from FilterStakeRegistryUpdated and is used to iterate over the raw logs and unpacked data for StakeRegistryUpdated events raised by the ServiceManager contract.
type ServiceManagerStakeRegistryUpdatedIterator struct {
	Event *ServiceManagerStakeRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *ServiceManagerStakeRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerStakeRegistryUpdated)
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
		it.Event = new(ServiceManagerStakeRegistryUpdated)
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
func (it *ServiceManagerStakeRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerStakeRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerStakeRegistryUpdated represents a StakeRegistryUpdated event raised by the ServiceManager contract.
type ServiceManagerStakeRegistryUpdated struct {
	StakeRegistry common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeRegistryUpdated is a free log retrieval operation binding the contract event 0xcb08c26360210256eb3fd98640509ba95a8a716bc3aa4aadc738e333102de737.
//
// Solidity: event StakeRegistryUpdated(address indexed stakeRegistry)
func (_ServiceManager *ServiceManagerFilterer) FilterStakeRegistryUpdated(opts *bind.FilterOpts, stakeRegistry []common.Address) (*ServiceManagerStakeRegistryUpdatedIterator, error) {

	var stakeRegistryRule []interface{}
	for _, stakeRegistryItem := range stakeRegistry {
		stakeRegistryRule = append(stakeRegistryRule, stakeRegistryItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "StakeRegistryUpdated", stakeRegistryRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerStakeRegistryUpdatedIterator{contract: _ServiceManager.contract, event: "StakeRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchStakeRegistryUpdated is a free log subscription operation binding the contract event 0xcb08c26360210256eb3fd98640509ba95a8a716bc3aa4aadc738e333102de737.
//
// Solidity: event StakeRegistryUpdated(address indexed stakeRegistry)
func (_ServiceManager *ServiceManagerFilterer) WatchStakeRegistryUpdated(opts *bind.WatchOpts, sink chan<- *ServiceManagerStakeRegistryUpdated, stakeRegistry []common.Address) (event.Subscription, error) {

	var stakeRegistryRule []interface{}
	for _, stakeRegistryItem := range stakeRegistry {
		stakeRegistryRule = append(stakeRegistryRule, stakeRegistryItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "StakeRegistryUpdated", stakeRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerStakeRegistryUpdated)
				if err := _ServiceManager.contract.UnpackLog(event, "StakeRegistryUpdated", log); err != nil {
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

// ParseStakeRegistryUpdated is a log parse operation binding the contract event 0xcb08c26360210256eb3fd98640509ba95a8a716bc3aa4aadc738e333102de737.
//
// Solidity: event StakeRegistryUpdated(address indexed stakeRegistry)
func (_ServiceManager *ServiceManagerFilterer) ParseStakeRegistryUpdated(log types.Log) (*ServiceManagerStakeRegistryUpdated, error) {
	event := new(ServiceManagerStakeRegistryUpdated)
	if err := _ServiceManager.contract.UnpackLog(event, "StakeRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerStrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the ServiceManager contract.
type ServiceManagerStrategyAddedIterator struct {
	Event *ServiceManagerStrategyAdded // Event containing the contract specifics and raw log

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
func (it *ServiceManagerStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerStrategyAdded)
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
		it.Event = new(ServiceManagerStrategyAdded)
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
func (it *ServiceManagerStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerStrategyAdded represents a StrategyAdded event raised by the ServiceManager contract.
type ServiceManagerStrategyAdded struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAdded is a free log retrieval operation binding the contract event 0x3f008fd510eae7a9e7bee13513d7b83bef8003d488b5a3d0b0da4de71d6846f1.
//
// Solidity: event StrategyAdded(address indexed strategy)
func (_ServiceManager *ServiceManagerFilterer) FilterStrategyAdded(opts *bind.FilterOpts, strategy []common.Address) (*ServiceManagerStrategyAddedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerStrategyAddedIterator{contract: _ServiceManager.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x3f008fd510eae7a9e7bee13513d7b83bef8003d488b5a3d0b0da4de71d6846f1.
//
// Solidity: event StrategyAdded(address indexed strategy)
func (_ServiceManager *ServiceManagerFilterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *ServiceManagerStrategyAdded, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerStrategyAdded)
				if err := _ServiceManager.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
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

// ParseStrategyAdded is a log parse operation binding the contract event 0x3f008fd510eae7a9e7bee13513d7b83bef8003d488b5a3d0b0da4de71d6846f1.
//
// Solidity: event StrategyAdded(address indexed strategy)
func (_ServiceManager *ServiceManagerFilterer) ParseStrategyAdded(log types.Log) (*ServiceManagerStrategyAdded, error) {
	event := new(ServiceManagerStrategyAdded)
	if err := _ServiceManager.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerStrategyRemovedIterator is returned from FilterStrategyRemoved and is used to iterate over the raw logs and unpacked data for StrategyRemoved events raised by the ServiceManager contract.
type ServiceManagerStrategyRemovedIterator struct {
	Event *ServiceManagerStrategyRemoved // Event containing the contract specifics and raw log

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
func (it *ServiceManagerStrategyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerStrategyRemoved)
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
		it.Event = new(ServiceManagerStrategyRemoved)
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
func (it *ServiceManagerStrategyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerStrategyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerStrategyRemoved represents a StrategyRemoved event raised by the ServiceManager contract.
type ServiceManagerStrategyRemoved struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemoved is a free log retrieval operation binding the contract event 0x09a1db4b80c32706328728508c941a6b954f31eb5affd32f236c1fd405f8fea4.
//
// Solidity: event StrategyRemoved(address indexed strategy)
func (_ServiceManager *ServiceManagerFilterer) FilterStrategyRemoved(opts *bind.FilterOpts, strategy []common.Address) (*ServiceManagerStrategyRemovedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "StrategyRemoved", strategyRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerStrategyRemovedIterator{contract: _ServiceManager.contract, event: "StrategyRemoved", logs: logs, sub: sub}, nil
}

// WatchStrategyRemoved is a free log subscription operation binding the contract event 0x09a1db4b80c32706328728508c941a6b954f31eb5affd32f236c1fd405f8fea4.
//
// Solidity: event StrategyRemoved(address indexed strategy)
func (_ServiceManager *ServiceManagerFilterer) WatchStrategyRemoved(opts *bind.WatchOpts, sink chan<- *ServiceManagerStrategyRemoved, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "StrategyRemoved", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerStrategyRemoved)
				if err := _ServiceManager.contract.UnpackLog(event, "StrategyRemoved", log); err != nil {
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

// ParseStrategyRemoved is a log parse operation binding the contract event 0x09a1db4b80c32706328728508c941a6b954f31eb5affd32f236c1fd405f8fea4.
//
// Solidity: event StrategyRemoved(address indexed strategy)
func (_ServiceManager *ServiceManagerFilterer) ParseStrategyRemoved(log types.Log) (*ServiceManagerStrategyRemoved, error) {
	event := new(ServiceManagerStrategyRemoved)
	if err := _ServiceManager.contract.UnpackLog(event, "StrategyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerTaskExecutedIterator is returned from FilterTaskExecuted and is used to iterate over the raw logs and unpacked data for TaskExecuted events raised by the ServiceManager contract.
type ServiceManagerTaskExecutedIterator struct {
	Event *ServiceManagerTaskExecuted // Event containing the contract specifics and raw log

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
func (it *ServiceManagerTaskExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerTaskExecuted)
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
		it.Event = new(ServiceManagerTaskExecuted)
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
func (it *ServiceManagerTaskExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerTaskExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerTaskExecuted represents a TaskExecuted event raised by the ServiceManager contract.
type ServiceManagerTaskExecuted struct {
	TaskHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskExecuted is a free log retrieval operation binding the contract event 0x3a04c0621dfd116e61ea63184b2b9c8cf514b0075675645493095b7be2e492d5.
//
// Solidity: event TaskExecuted(bytes32 indexed taskHash)
func (_ServiceManager *ServiceManagerFilterer) FilterTaskExecuted(opts *bind.FilterOpts, taskHash [][32]byte) (*ServiceManagerTaskExecutedIterator, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "TaskExecuted", taskHashRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerTaskExecutedIterator{contract: _ServiceManager.contract, event: "TaskExecuted", logs: logs, sub: sub}, nil
}

// WatchTaskExecuted is a free log subscription operation binding the contract event 0x3a04c0621dfd116e61ea63184b2b9c8cf514b0075675645493095b7be2e492d5.
//
// Solidity: event TaskExecuted(bytes32 indexed taskHash)
func (_ServiceManager *ServiceManagerFilterer) WatchTaskExecuted(opts *bind.WatchOpts, sink chan<- *ServiceManagerTaskExecuted, taskHash [][32]byte) (event.Subscription, error) {

	var taskHashRule []interface{}
	for _, taskHashItem := range taskHash {
		taskHashRule = append(taskHashRule, taskHashItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "TaskExecuted", taskHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerTaskExecuted)
				if err := _ServiceManager.contract.UnpackLog(event, "TaskExecuted", log); err != nil {
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

// ParseTaskExecuted is a log parse operation binding the contract event 0x3a04c0621dfd116e61ea63184b2b9c8cf514b0075675645493095b7be2e492d5.
//
// Solidity: event TaskExecuted(bytes32 indexed taskHash)
func (_ServiceManager *ServiceManagerFilterer) ParseTaskExecuted(log types.Log) (*ServiceManagerTaskExecuted, error) {
	event := new(ServiceManagerTaskExecuted)
	if err := _ServiceManager.contract.UnpackLog(event, "TaskExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerTaskValidatedIterator is returned from FilterTaskValidated and is used to iterate over the raw logs and unpacked data for TaskValidated events raised by the ServiceManager contract.
type ServiceManagerTaskValidatedIterator struct {
	Event *ServiceManagerTaskValidated // Event containing the contract specifics and raw log

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
func (it *ServiceManagerTaskValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerTaskValidated)
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
		it.Event = new(ServiceManagerTaskValidated)
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
func (it *ServiceManagerTaskValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerTaskValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerTaskValidated represents a TaskValidated event raised by the ServiceManager contract.
type ServiceManagerTaskValidated struct {
	MsgSender            common.Address
	Target               common.Address
	Value                *big.Int
	PolicyID             string
	TaskId               string
	QuorumThresholdCount uint32
	ExpireByBlockNumber  *big.Int
	SignerAddresses      []common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskValidated is a free log retrieval operation binding the contract event 0x9f851f9b4b4a1023918bda08f974ba8a757d5e8c1d8d4131e9963fdd64bd1a4d.
//
// Solidity: event TaskValidated(address indexed msgSender, address indexed target, uint256 indexed value, string policyID, string taskId, uint32 quorumThresholdCount, uint256 expireByBlockNumber, address[] signerAddresses)
func (_ServiceManager *ServiceManagerFilterer) FilterTaskValidated(opts *bind.FilterOpts, msgSender []common.Address, target []common.Address, value []*big.Int) (*ServiceManagerTaskValidatedIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "TaskValidated", msgSenderRule, targetRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerTaskValidatedIterator{contract: _ServiceManager.contract, event: "TaskValidated", logs: logs, sub: sub}, nil
}

// WatchTaskValidated is a free log subscription operation binding the contract event 0x9f851f9b4b4a1023918bda08f974ba8a757d5e8c1d8d4131e9963fdd64bd1a4d.
//
// Solidity: event TaskValidated(address indexed msgSender, address indexed target, uint256 indexed value, string policyID, string taskId, uint32 quorumThresholdCount, uint256 expireByBlockNumber, address[] signerAddresses)
func (_ServiceManager *ServiceManagerFilterer) WatchTaskValidated(opts *bind.WatchOpts, sink chan<- *ServiceManagerTaskValidated, msgSender []common.Address, target []common.Address, value []*big.Int) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "TaskValidated", msgSenderRule, targetRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerTaskValidated)
				if err := _ServiceManager.contract.UnpackLog(event, "TaskValidated", log); err != nil {
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

// ParseTaskValidated is a log parse operation binding the contract event 0x9f851f9b4b4a1023918bda08f974ba8a757d5e8c1d8d4131e9963fdd64bd1a4d.
//
// Solidity: event TaskValidated(address indexed msgSender, address indexed target, uint256 indexed value, string policyID, string taskId, uint32 quorumThresholdCount, uint256 expireByBlockNumber, address[] signerAddresses)
func (_ServiceManager *ServiceManagerFilterer) ParseTaskValidated(log types.Log) (*ServiceManagerTaskValidated, error) {
	event := new(ServiceManagerTaskValidated)
	if err := _ServiceManager.contract.UnpackLog(event, "TaskValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceManagerThresholdStakeUpdatedIterator is returned from FilterThresholdStakeUpdated and is used to iterate over the raw logs and unpacked data for ThresholdStakeUpdated events raised by the ServiceManager contract.
type ServiceManagerThresholdStakeUpdatedIterator struct {
	Event *ServiceManagerThresholdStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ServiceManagerThresholdStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceManagerThresholdStakeUpdated)
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
		it.Event = new(ServiceManagerThresholdStakeUpdated)
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
func (it *ServiceManagerThresholdStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceManagerThresholdStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceManagerThresholdStakeUpdated represents a ThresholdStakeUpdated event raised by the ServiceManager contract.
type ServiceManagerThresholdStakeUpdated struct {
	ThresholdStake *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterThresholdStakeUpdated is a free log retrieval operation binding the contract event 0x76f50103a7b1e136f23c29045235143c62e36705e8253019cb1253d9062bc6d4.
//
// Solidity: event ThresholdStakeUpdated(uint256 indexed thresholdStake)
func (_ServiceManager *ServiceManagerFilterer) FilterThresholdStakeUpdated(opts *bind.FilterOpts, thresholdStake []*big.Int) (*ServiceManagerThresholdStakeUpdatedIterator, error) {

	var thresholdStakeRule []interface{}
	for _, thresholdStakeItem := range thresholdStake {
		thresholdStakeRule = append(thresholdStakeRule, thresholdStakeItem)
	}

	logs, sub, err := _ServiceManager.contract.FilterLogs(opts, "ThresholdStakeUpdated", thresholdStakeRule)
	if err != nil {
		return nil, err
	}
	return &ServiceManagerThresholdStakeUpdatedIterator{contract: _ServiceManager.contract, event: "ThresholdStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdStakeUpdated is a free log subscription operation binding the contract event 0x76f50103a7b1e136f23c29045235143c62e36705e8253019cb1253d9062bc6d4.
//
// Solidity: event ThresholdStakeUpdated(uint256 indexed thresholdStake)
func (_ServiceManager *ServiceManagerFilterer) WatchThresholdStakeUpdated(opts *bind.WatchOpts, sink chan<- *ServiceManagerThresholdStakeUpdated, thresholdStake []*big.Int) (event.Subscription, error) {

	var thresholdStakeRule []interface{}
	for _, thresholdStakeItem := range thresholdStake {
		thresholdStakeRule = append(thresholdStakeRule, thresholdStakeItem)
	}

	logs, sub, err := _ServiceManager.contract.WatchLogs(opts, "ThresholdStakeUpdated", thresholdStakeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceManagerThresholdStakeUpdated)
				if err := _ServiceManager.contract.UnpackLog(event, "ThresholdStakeUpdated", log); err != nil {
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

// ParseThresholdStakeUpdated is a log parse operation binding the contract event 0x76f50103a7b1e136f23c29045235143c62e36705e8253019cb1253d9062bc6d4.
//
// Solidity: event ThresholdStakeUpdated(uint256 indexed thresholdStake)
func (_ServiceManager *ServiceManagerFilterer) ParseThresholdStakeUpdated(log types.Log) (*ServiceManagerThresholdStakeUpdated, error) {
	event := new(ServiceManagerThresholdStakeUpdated)
	if err := _ServiceManager.contract.UnpackLog(event, "ThresholdStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
