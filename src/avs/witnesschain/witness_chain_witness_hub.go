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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// IWitnessHubStrategyParam is an auto generated low-level Go binding around an user-defined struct.
type IWitnessHubStrategyParam struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// TypesBountyRewards is an auto generated low-level Go binding around an user-defined struct.
type TypesBountyRewards struct {
	InclusionProofBounties *big.Int
	DiligenceProofBounties *big.Int
}

// WitnessChainWitnessHubMetaData contains all meta data concerning the WitnessChainWitnessHub contract.
var WitnessChainWitnessHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIAVSDirectory\",\"name\":\"__avsDirectory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAggregator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAggregator\",\"type\":\"address\"}],\"name\":\"AggregatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"InvalidOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldL2ChainMapping\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newL2ChainMapping\",\"type\":\"address\"}],\"name\":\"L2ChainMappingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2BlockNumberEnd\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rewardHash\",\"type\":\"bytes32\"}],\"name\":\"NewRewardsUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRegistry\",\"type\":\"address\"}],\"name\":\"RegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"indexed\":false,\"internalType\":\"structIWitnessHub.StrategyParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"SetStrategyParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_proofCommitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumberBegin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumberEnd\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rewardHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"submissionBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"}],\"name\":\"getNextBlockByChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"}],\"name\":\"getOperatorRewardsByChainID\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inclusionProofBounties\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diligenceProofBounties\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BountyRewards\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2ChainMapping\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_agg\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ChainMapping\",\"outputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUpdateBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"}],\"name\":\"setAggregatorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIL2ChainMapping\",\"name\":\"_l2chainmapping\",\"type\":\"address\"}],\"name\":\"setL2ChainMapping\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOperatorRegistry\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"multiplier\",\"type\":\"uint96\"}],\"internalType\":\"structIWitnessHub.StrategyParam[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"setStrategyParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumBegin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumEnd\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_operatorsList\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"inclusionProofBounties\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"diligenceProofBounties\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.BountyRewards[]\",\"name\":\"_proofRewards\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardHash\",\"type\":\"bytes32\"}],\"name\":\"updateReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// WitnessChainWitnessHubABI is the input ABI used to generate the binding from.
// Deprecated: Use WitnessChainWitnessHubMetaData.ABI instead.
var WitnessChainWitnessHubABI = WitnessChainWitnessHubMetaData.ABI

// WitnessChainWitnessHub is an auto generated Go binding around an Ethereum contract.
type WitnessChainWitnessHub struct {
	WitnessChainWitnessHubCaller     // Read-only binding to the contract
	WitnessChainWitnessHubTransactor // Write-only binding to the contract
	WitnessChainWitnessHubFilterer   // Log filterer for contract events
}

// WitnessChainWitnessHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type WitnessChainWitnessHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessChainWitnessHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WitnessChainWitnessHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessChainWitnessHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WitnessChainWitnessHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WitnessChainWitnessHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WitnessChainWitnessHubSession struct {
	Contract     *WitnessChainWitnessHub // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WitnessChainWitnessHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WitnessChainWitnessHubCallerSession struct {
	Contract *WitnessChainWitnessHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// WitnessChainWitnessHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WitnessChainWitnessHubTransactorSession struct {
	Contract     *WitnessChainWitnessHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// WitnessChainWitnessHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type WitnessChainWitnessHubRaw struct {
	Contract *WitnessChainWitnessHub // Generic contract binding to access the raw methods on
}

// WitnessChainWitnessHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WitnessChainWitnessHubCallerRaw struct {
	Contract *WitnessChainWitnessHubCaller // Generic read-only contract binding to access the raw methods on
}

// WitnessChainWitnessHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WitnessChainWitnessHubTransactorRaw struct {
	Contract *WitnessChainWitnessHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWitnessChainWitnessHub creates a new instance of WitnessChainWitnessHub, bound to a specific deployed contract.
func NewWitnessChainWitnessHub(address common.Address, backend bind.ContractBackend) (*WitnessChainWitnessHub, error) {
	contract, err := bindWitnessChainWitnessHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHub{WitnessChainWitnessHubCaller: WitnessChainWitnessHubCaller{contract: contract}, WitnessChainWitnessHubTransactor: WitnessChainWitnessHubTransactor{contract: contract}, WitnessChainWitnessHubFilterer: WitnessChainWitnessHubFilterer{contract: contract}}, nil
}

// NewWitnessChainWitnessHubCaller creates a new read-only instance of WitnessChainWitnessHub, bound to a specific deployed contract.
func NewWitnessChainWitnessHubCaller(address common.Address, caller bind.ContractCaller) (*WitnessChainWitnessHubCaller, error) {
	contract, err := bindWitnessChainWitnessHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubCaller{contract: contract}, nil
}

// NewWitnessChainWitnessHubTransactor creates a new write-only instance of WitnessChainWitnessHub, bound to a specific deployed contract.
func NewWitnessChainWitnessHubTransactor(address common.Address, transactor bind.ContractTransactor) (*WitnessChainWitnessHubTransactor, error) {
	contract, err := bindWitnessChainWitnessHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubTransactor{contract: contract}, nil
}

// NewWitnessChainWitnessHubFilterer creates a new log filterer instance of WitnessChainWitnessHub, bound to a specific deployed contract.
func NewWitnessChainWitnessHubFilterer(address common.Address, filterer bind.ContractFilterer) (*WitnessChainWitnessHubFilterer, error) {
	contract, err := bindWitnessChainWitnessHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubFilterer{contract: contract}, nil
}

// bindWitnessChainWitnessHub binds a generic wrapper to an already deployed contract.
func bindWitnessChainWitnessHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WitnessChainWitnessHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessChainWitnessHub *WitnessChainWitnessHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WitnessChainWitnessHub.Contract.WitnessChainWitnessHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessChainWitnessHub *WitnessChainWitnessHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.WitnessChainWitnessHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessChainWitnessHub *WitnessChainWitnessHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.WitnessChainWitnessHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WitnessChainWitnessHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.contract.Transact(opts, method, params...)
}

// ProofCommitments is a free data retrieval call binding the contract method 0x22062c14.
//
// Solidity: function _proofCommitments(uint256 ) view returns(uint256 chainID, uint256 l2BlockNumberBegin, uint256 l2BlockNumberEnd, bytes32 rewardHash, uint256 submissionBlock)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) ProofCommitments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChainID            *big.Int
	L2BlockNumberBegin *big.Int
	L2BlockNumberEnd   *big.Int
	RewardHash         [32]byte
	SubmissionBlock    *big.Int
}, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "_proofCommitments", arg0)

	outstruct := new(struct {
		ChainID            *big.Int
		L2BlockNumberBegin *big.Int
		L2BlockNumberEnd   *big.Int
		RewardHash         [32]byte
		SubmissionBlock    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainID = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.L2BlockNumberBegin = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.L2BlockNumberEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.SubmissionBlock = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProofCommitments is a free data retrieval call binding the contract method 0x22062c14.
//
// Solidity: function _proofCommitments(uint256 ) view returns(uint256 chainID, uint256 l2BlockNumberBegin, uint256 l2BlockNumberEnd, bytes32 rewardHash, uint256 submissionBlock)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) ProofCommitments(arg0 *big.Int) (struct {
	ChainID            *big.Int
	L2BlockNumberBegin *big.Int
	L2BlockNumberEnd   *big.Int
	RewardHash         [32]byte
	SubmissionBlock    *big.Int
}, error) {
	return _WitnessChainWitnessHub.Contract.ProofCommitments(&_WitnessChainWitnessHub.CallOpts, arg0)
}

// ProofCommitments is a free data retrieval call binding the contract method 0x22062c14.
//
// Solidity: function _proofCommitments(uint256 ) view returns(uint256 chainID, uint256 l2BlockNumberBegin, uint256 l2BlockNumberEnd, bytes32 rewardHash, uint256 submissionBlock)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) ProofCommitments(arg0 *big.Int) (struct {
	ChainID            *big.Int
	L2BlockNumberBegin *big.Int
	L2BlockNumberEnd   *big.Int
	RewardHash         [32]byte
	SubmissionBlock    *big.Int
}, error) {
	return _WitnessChainWitnessHub.Contract.ProofCommitments(&_WitnessChainWitnessHub.CallOpts, arg0)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) Aggregator() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.Aggregator(&_WitnessChainWitnessHub.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) Aggregator() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.Aggregator(&_WitnessChainWitnessHub.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) AvsDirectory() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.AvsDirectory(&_WitnessChainWitnessHub.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) AvsDirectory() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.AvsDirectory(&_WitnessChainWitnessHub.CallOpts)
}

// GetNextBlockByChainID is a free data retrieval call binding the contract method 0xac359bd3.
//
// Solidity: function getNextBlockByChainID(uint256 _chainID) view returns(uint256)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) GetNextBlockByChainID(opts *bind.CallOpts, _chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "getNextBlockByChainID", _chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextBlockByChainID is a free data retrieval call binding the contract method 0xac359bd3.
//
// Solidity: function getNextBlockByChainID(uint256 _chainID) view returns(uint256)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) GetNextBlockByChainID(_chainID *big.Int) (*big.Int, error) {
	return _WitnessChainWitnessHub.Contract.GetNextBlockByChainID(&_WitnessChainWitnessHub.CallOpts, _chainID)
}

// GetNextBlockByChainID is a free data retrieval call binding the contract method 0xac359bd3.
//
// Solidity: function getNextBlockByChainID(uint256 _chainID) view returns(uint256)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) GetNextBlockByChainID(_chainID *big.Int) (*big.Int, error) {
	return _WitnessChainWitnessHub.Contract.GetNextBlockByChainID(&_WitnessChainWitnessHub.CallOpts, _chainID)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _WitnessChainWitnessHub.Contract.GetOperatorRestakedStrategies(&_WitnessChainWitnessHub.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _WitnessChainWitnessHub.Contract.GetOperatorRestakedStrategies(&_WitnessChainWitnessHub.CallOpts, operator)
}

// GetOperatorRewardsByChainID is a free data retrieval call binding the contract method 0x93a949fb.
//
// Solidity: function getOperatorRewardsByChainID(address _operator, uint256 _chainID) view returns((uint256,uint256))
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) GetOperatorRewardsByChainID(opts *bind.CallOpts, _operator common.Address, _chainID *big.Int) (TypesBountyRewards, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "getOperatorRewardsByChainID", _operator, _chainID)

	if err != nil {
		return *new(TypesBountyRewards), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesBountyRewards)).(*TypesBountyRewards)

	return out0, err

}

// GetOperatorRewardsByChainID is a free data retrieval call binding the contract method 0x93a949fb.
//
// Solidity: function getOperatorRewardsByChainID(address _operator, uint256 _chainID) view returns((uint256,uint256))
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) GetOperatorRewardsByChainID(_operator common.Address, _chainID *big.Int) (TypesBountyRewards, error) {
	return _WitnessChainWitnessHub.Contract.GetOperatorRewardsByChainID(&_WitnessChainWitnessHub.CallOpts, _operator, _chainID)
}

// GetOperatorRewardsByChainID is a free data retrieval call binding the contract method 0x93a949fb.
//
// Solidity: function getOperatorRewardsByChainID(address _operator, uint256 _chainID) view returns((uint256,uint256))
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) GetOperatorRewardsByChainID(_operator common.Address, _chainID *big.Int) (TypesBountyRewards, error) {
	return _WitnessChainWitnessHub.Contract.GetOperatorRewardsByChainID(&_WitnessChainWitnessHub.CallOpts, _operator, _chainID)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _WitnessChainWitnessHub.Contract.GetRestakeableStrategies(&_WitnessChainWitnessHub.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _WitnessChainWitnessHub.Contract.GetRestakeableStrategies(&_WitnessChainWitnessHub.CallOpts)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) L2ChainMapping(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "l2ChainMapping")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) L2ChainMapping() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.L2ChainMapping(&_WitnessChainWitnessHub.CallOpts)
}

// L2ChainMapping is a free data retrieval call binding the contract method 0x820b6f39.
//
// Solidity: function l2ChainMapping() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) L2ChainMapping() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.L2ChainMapping(&_WitnessChainWitnessHub.CallOpts)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x5472534d.
//
// Solidity: function operatorRewards(uint256 ) view returns(uint256 lastUpdateBlock)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) OperatorRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "operatorRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorRewards is a free data retrieval call binding the contract method 0x5472534d.
//
// Solidity: function operatorRewards(uint256 ) view returns(uint256 lastUpdateBlock)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) OperatorRewards(arg0 *big.Int) (*big.Int, error) {
	return _WitnessChainWitnessHub.Contract.OperatorRewards(&_WitnessChainWitnessHub.CallOpts, arg0)
}

// OperatorRewards is a free data retrieval call binding the contract method 0x5472534d.
//
// Solidity: function operatorRewards(uint256 ) view returns(uint256 lastUpdateBlock)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) OperatorRewards(arg0 *big.Int) (*big.Int, error) {
	return _WitnessChainWitnessHub.Contract.OperatorRewards(&_WitnessChainWitnessHub.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) Owner() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.Owner(&_WitnessChainWitnessHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) Owner() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.Owner(&_WitnessChainWitnessHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) Paused() (bool, error) {
	return _WitnessChainWitnessHub.Contract.Paused(&_WitnessChainWitnessHub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) Paused() (bool, error) {
	return _WitnessChainWitnessHub.Contract.Paused(&_WitnessChainWitnessHub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) ProxiableUUID() ([32]byte, error) {
	return _WitnessChainWitnessHub.Contract.ProxiableUUID(&_WitnessChainWitnessHub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WitnessChainWitnessHub.Contract.ProxiableUUID(&_WitnessChainWitnessHub.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WitnessChainWitnessHub.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) Registry() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.Registry(&_WitnessChainWitnessHub.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubCallerSession) Registry() (common.Address, error) {
	return _WitnessChainWitnessHub.Contract.Registry(&_WitnessChainWitnessHub.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.DeregisterOperatorFromAVS(&_WitnessChainWitnessHub.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.DeregisterOperatorFromAVS(&_WitnessChainWitnessHub.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "initialize", _registry, _l2ChainMapping, _agg)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) Initialize(_registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.Initialize(&_WitnessChainWitnessHub.TransactOpts, _registry, _l2ChainMapping, _agg)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _registry, address _l2ChainMapping, address _agg) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) Initialize(_registry common.Address, _l2ChainMapping common.Address, _agg common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.Initialize(&_WitnessChainWitnessHub.TransactOpts, _registry, _l2ChainMapping, _agg)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) Pause() (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.Pause(&_WitnessChainWitnessHub.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) Pause() (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.Pause(&_WitnessChainWitnessHub.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.RegisterOperatorToAVS(&_WitnessChainWitnessHub.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.RegisterOperatorToAVS(&_WitnessChainWitnessHub.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.RenounceOwnership(&_WitnessChainWitnessHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.RenounceOwnership(&_WitnessChainWitnessHub.TransactOpts)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) SetAggregatorAddress(opts *bind.TransactOpts, _aggregator common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "setAggregatorAddress", _aggregator)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) SetAggregatorAddress(_aggregator common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetAggregatorAddress(&_WitnessChainWitnessHub.TransactOpts, _aggregator)
}

// SetAggregatorAddress is a paid mutator transaction binding the contract method 0x47b32448.
//
// Solidity: function setAggregatorAddress(address _aggregator) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) SetAggregatorAddress(_aggregator common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetAggregatorAddress(&_WitnessChainWitnessHub.TransactOpts, _aggregator)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) SetL2ChainMapping(opts *bind.TransactOpts, _l2chainmapping common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "setL2ChainMapping", _l2chainmapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) SetL2ChainMapping(_l2chainmapping common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetL2ChainMapping(&_WitnessChainWitnessHub.TransactOpts, _l2chainmapping)
}

// SetL2ChainMapping is a paid mutator transaction binding the contract method 0x8d134f8d.
//
// Solidity: function setL2ChainMapping(address _l2chainmapping) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) SetL2ChainMapping(_l2chainmapping common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetL2ChainMapping(&_WitnessChainWitnessHub.TransactOpts, _l2chainmapping)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetRegistry(&_WitnessChainWitnessHub.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetRegistry(&_WitnessChainWitnessHub.TransactOpts, _registry)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) SetStrategyParams(opts *bind.TransactOpts, params []IWitnessHubStrategyParam) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "setStrategyParams", params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) SetStrategyParams(params []IWitnessHubStrategyParam) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetStrategyParams(&_WitnessChainWitnessHub.TransactOpts, params)
}

// SetStrategyParams is a paid mutator transaction binding the contract method 0xae30f16d.
//
// Solidity: function setStrategyParams((address,uint96)[] params) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) SetStrategyParams(params []IWitnessHubStrategyParam) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.SetStrategyParams(&_WitnessChainWitnessHub.TransactOpts, params)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.TransferOwnership(&_WitnessChainWitnessHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.TransferOwnership(&_WitnessChainWitnessHub.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) Unpause() (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.Unpause(&_WitnessChainWitnessHub.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) Unpause() (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.Unpause(&_WitnessChainWitnessHub.TransactOpts)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpdateAVSMetadataURI(&_WitnessChainWitnessHub.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpdateAVSMetadataURI(&_WitnessChainWitnessHub.TransactOpts, _metadataURI)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) UpdateReward(opts *bind.TransactOpts, _chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesBountyRewards, _rewardHash [32]byte) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "updateReward", _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) UpdateReward(_chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesBountyRewards, _rewardHash [32]byte) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpdateReward(&_WitnessChainWitnessHub.TransactOpts, _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpdateReward is a paid mutator transaction binding the contract method 0x395373b0.
//
// Solidity: function updateReward(uint256 _chainID, uint256 _blockNumBegin, uint256 _blockNumEnd, address[] _operatorsList, (uint256,uint256)[] _proofRewards, bytes32 _rewardHash) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) UpdateReward(_chainID *big.Int, _blockNumBegin *big.Int, _blockNumEnd *big.Int, _operatorsList []common.Address, _proofRewards []TypesBountyRewards, _rewardHash [32]byte) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpdateReward(&_WitnessChainWitnessHub.TransactOpts, _chainID, _blockNumBegin, _blockNumEnd, _operatorsList, _proofRewards, _rewardHash)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpgradeTo(&_WitnessChainWitnessHub.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpgradeTo(&_WitnessChainWitnessHub.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpgradeToAndCall(&_WitnessChainWitnessHub.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WitnessChainWitnessHub *WitnessChainWitnessHubTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WitnessChainWitnessHub.Contract.UpgradeToAndCall(&_WitnessChainWitnessHub.TransactOpts, newImplementation, data)
}

// WitnessChainWitnessHubAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubAdminChangedIterator struct {
	Event *WitnessChainWitnessHubAdminChanged // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubAdminChanged)
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
		it.Event = new(WitnessChainWitnessHubAdminChanged)
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
func (it *WitnessChainWitnessHubAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubAdminChanged represents a AdminChanged event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WitnessChainWitnessHubAdminChangedIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubAdminChangedIterator{contract: _WitnessChainWitnessHub.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubAdminChanged) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubAdminChanged)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseAdminChanged(log types.Log) (*WitnessChainWitnessHubAdminChanged, error) {
	event := new(WitnessChainWitnessHubAdminChanged)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubAggregatorUpdatedIterator is returned from FilterAggregatorUpdated and is used to iterate over the raw logs and unpacked data for AggregatorUpdated events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubAggregatorUpdatedIterator struct {
	Event *WitnessChainWitnessHubAggregatorUpdated // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubAggregatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubAggregatorUpdated)
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
		it.Event = new(WitnessChainWitnessHubAggregatorUpdated)
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
func (it *WitnessChainWitnessHubAggregatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubAggregatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubAggregatorUpdated represents a AggregatorUpdated event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubAggregatorUpdated struct {
	OldAggregator common.Address
	NewAggregator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatorUpdated is a free log retrieval operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterAggregatorUpdated(opts *bind.FilterOpts) (*WitnessChainWitnessHubAggregatorUpdatedIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubAggregatorUpdatedIterator{contract: _WitnessChainWitnessHub.contract, event: "AggregatorUpdated", logs: logs, sub: sub}, nil
}

// WatchAggregatorUpdated is a free log subscription operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchAggregatorUpdated(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubAggregatorUpdated) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "AggregatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubAggregatorUpdated)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
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

// ParseAggregatorUpdated is a log parse operation binding the contract event 0x89baabef7dfd0683c0ac16fd2a8431c51b49fbe654c3f7b5ef19763e2ccd88f2.
//
// Solidity: event AggregatorUpdated(address oldAggregator, address newAggregator)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseAggregatorUpdated(log types.Log) (*WitnessChainWitnessHubAggregatorUpdated, error) {
	event := new(WitnessChainWitnessHubAggregatorUpdated)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "AggregatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubBeaconUpgradedIterator struct {
	Event *WitnessChainWitnessHubBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubBeaconUpgraded)
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
		it.Event = new(WitnessChainWitnessHubBeaconUpgraded)
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
func (it *WitnessChainWitnessHubBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubBeaconUpgraded represents a BeaconUpgraded event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WitnessChainWitnessHubBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubBeaconUpgradedIterator{contract: _WitnessChainWitnessHub.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubBeaconUpgraded)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseBeaconUpgraded(log types.Log) (*WitnessChainWitnessHubBeaconUpgraded, error) {
	event := new(WitnessChainWitnessHubBeaconUpgraded)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubInitializedIterator struct {
	Event *WitnessChainWitnessHubInitialized // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubInitialized)
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
		it.Event = new(WitnessChainWitnessHubInitialized)
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
func (it *WitnessChainWitnessHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubInitialized represents a Initialized event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*WitnessChainWitnessHubInitializedIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubInitializedIterator{contract: _WitnessChainWitnessHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubInitialized) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubInitialized)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseInitialized(log types.Log) (*WitnessChainWitnessHubInitialized, error) {
	event := new(WitnessChainWitnessHubInitialized)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubInvalidOperatorIterator is returned from FilterInvalidOperator and is used to iterate over the raw logs and unpacked data for InvalidOperator events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubInvalidOperatorIterator struct {
	Event *WitnessChainWitnessHubInvalidOperator // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubInvalidOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubInvalidOperator)
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
		it.Event = new(WitnessChainWitnessHubInvalidOperator)
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
func (it *WitnessChainWitnessHubInvalidOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubInvalidOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubInvalidOperator represents a InvalidOperator event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubInvalidOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInvalidOperator is a free log retrieval operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterInvalidOperator(opts *bind.FilterOpts) (*WitnessChainWitnessHubInvalidOperatorIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "InvalidOperator")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubInvalidOperatorIterator{contract: _WitnessChainWitnessHub.contract, event: "InvalidOperator", logs: logs, sub: sub}, nil
}

// WatchInvalidOperator is a free log subscription operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchInvalidOperator(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubInvalidOperator) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "InvalidOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubInvalidOperator)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "InvalidOperator", log); err != nil {
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

// ParseInvalidOperator is a log parse operation binding the contract event 0x3eaa03e38745f9ef282a3d3ae730e428a187a9075ccc2a3b7a1a318fe83d25be.
//
// Solidity: event InvalidOperator(address operator)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseInvalidOperator(log types.Log) (*WitnessChainWitnessHubInvalidOperator, error) {
	event := new(WitnessChainWitnessHubInvalidOperator)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "InvalidOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubL2ChainMappingUpdatedIterator is returned from FilterL2ChainMappingUpdated and is used to iterate over the raw logs and unpacked data for L2ChainMappingUpdated events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubL2ChainMappingUpdatedIterator struct {
	Event *WitnessChainWitnessHubL2ChainMappingUpdated // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubL2ChainMappingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubL2ChainMappingUpdated)
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
		it.Event = new(WitnessChainWitnessHubL2ChainMappingUpdated)
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
func (it *WitnessChainWitnessHubL2ChainMappingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubL2ChainMappingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubL2ChainMappingUpdated represents a L2ChainMappingUpdated event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubL2ChainMappingUpdated struct {
	OldL2ChainMapping common.Address
	NewL2ChainMapping common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterL2ChainMappingUpdated is a free log retrieval operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterL2ChainMappingUpdated(opts *bind.FilterOpts) (*WitnessChainWitnessHubL2ChainMappingUpdatedIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "L2ChainMappingUpdated")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubL2ChainMappingUpdatedIterator{contract: _WitnessChainWitnessHub.contract, event: "L2ChainMappingUpdated", logs: logs, sub: sub}, nil
}

// WatchL2ChainMappingUpdated is a free log subscription operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchL2ChainMappingUpdated(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubL2ChainMappingUpdated) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "L2ChainMappingUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubL2ChainMappingUpdated)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "L2ChainMappingUpdated", log); err != nil {
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

// ParseL2ChainMappingUpdated is a log parse operation binding the contract event 0xafcfe8816011d6bd68dd7f59cbd0865f721cc58085f72778b69065c75a19b5ca.
//
// Solidity: event L2ChainMappingUpdated(address oldL2ChainMapping, address newL2ChainMapping)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseL2ChainMappingUpdated(log types.Log) (*WitnessChainWitnessHubL2ChainMappingUpdated, error) {
	event := new(WitnessChainWitnessHubL2ChainMappingUpdated)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "L2ChainMappingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubNewRewardsUpdateIterator is returned from FilterNewRewardsUpdate and is used to iterate over the raw logs and unpacked data for NewRewardsUpdate events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubNewRewardsUpdateIterator struct {
	Event *WitnessChainWitnessHubNewRewardsUpdate // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubNewRewardsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubNewRewardsUpdate)
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
		it.Event = new(WitnessChainWitnessHubNewRewardsUpdate)
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
func (it *WitnessChainWitnessHubNewRewardsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubNewRewardsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubNewRewardsUpdate represents a NewRewardsUpdate event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubNewRewardsUpdate struct {
	ChainId          *big.Int
	L2BlockNumberEnd *big.Int
	RewardHash       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewRewardsUpdate is a free log retrieval operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterNewRewardsUpdate(opts *bind.FilterOpts, chainId []*big.Int, l2BlockNumberEnd []*big.Int, rewardHash [][32]byte) (*WitnessChainWitnessHubNewRewardsUpdateIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2BlockNumberEndRule []interface{}
	for _, l2BlockNumberEndItem := range l2BlockNumberEnd {
		l2BlockNumberEndRule = append(l2BlockNumberEndRule, l2BlockNumberEndItem)
	}
	var rewardHashRule []interface{}
	for _, rewardHashItem := range rewardHash {
		rewardHashRule = append(rewardHashRule, rewardHashItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "NewRewardsUpdate", chainIdRule, l2BlockNumberEndRule, rewardHashRule)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubNewRewardsUpdateIterator{contract: _WitnessChainWitnessHub.contract, event: "NewRewardsUpdate", logs: logs, sub: sub}, nil
}

// WatchNewRewardsUpdate is a free log subscription operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchNewRewardsUpdate(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubNewRewardsUpdate, chainId []*big.Int, l2BlockNumberEnd []*big.Int, rewardHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2BlockNumberEndRule []interface{}
	for _, l2BlockNumberEndItem := range l2BlockNumberEnd {
		l2BlockNumberEndRule = append(l2BlockNumberEndRule, l2BlockNumberEndItem)
	}
	var rewardHashRule []interface{}
	for _, rewardHashItem := range rewardHash {
		rewardHashRule = append(rewardHashRule, rewardHashItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "NewRewardsUpdate", chainIdRule, l2BlockNumberEndRule, rewardHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubNewRewardsUpdate)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "NewRewardsUpdate", log); err != nil {
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

// ParseNewRewardsUpdate is a log parse operation binding the contract event 0x36b8234ba3e9f8897b5448ed687509b776a4ad94a88ffd28d6aef8bee37ee610.
//
// Solidity: event NewRewardsUpdate(uint256 indexed chainId, uint256 indexed l2BlockNumberEnd, bytes32 indexed rewardHash)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseNewRewardsUpdate(log types.Log) (*WitnessChainWitnessHubNewRewardsUpdate, error) {
	event := new(WitnessChainWitnessHubNewRewardsUpdate)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "NewRewardsUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubOwnershipTransferredIterator struct {
	Event *WitnessChainWitnessHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubOwnershipTransferred)
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
		it.Event = new(WitnessChainWitnessHubOwnershipTransferred)
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
func (it *WitnessChainWitnessHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubOwnershipTransferred represents a OwnershipTransferred event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WitnessChainWitnessHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubOwnershipTransferredIterator{contract: _WitnessChainWitnessHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubOwnershipTransferred)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseOwnershipTransferred(log types.Log) (*WitnessChainWitnessHubOwnershipTransferred, error) {
	event := new(WitnessChainWitnessHubOwnershipTransferred)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubPausedIterator struct {
	Event *WitnessChainWitnessHubPaused // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubPaused)
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
		it.Event = new(WitnessChainWitnessHubPaused)
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
func (it *WitnessChainWitnessHubPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubPaused represents a Paused event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterPaused(opts *bind.FilterOpts) (*WitnessChainWitnessHubPausedIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubPausedIterator{contract: _WitnessChainWitnessHub.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubPaused) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubPaused)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParsePaused(log types.Log) (*WitnessChainWitnessHubPaused, error) {
	event := new(WitnessChainWitnessHubPaused)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubRegistryUpdatedIterator is returned from FilterRegistryUpdated and is used to iterate over the raw logs and unpacked data for RegistryUpdated events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubRegistryUpdatedIterator struct {
	Event *WitnessChainWitnessHubRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubRegistryUpdated)
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
		it.Event = new(WitnessChainWitnessHubRegistryUpdated)
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
func (it *WitnessChainWitnessHubRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubRegistryUpdated represents a RegistryUpdated event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubRegistryUpdated struct {
	OldRegistry common.Address
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistryUpdated is a free log retrieval operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterRegistryUpdated(opts *bind.FilterOpts) (*WitnessChainWitnessHubRegistryUpdatedIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubRegistryUpdatedIterator{contract: _WitnessChainWitnessHub.contract, event: "RegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchRegistryUpdated is a free log subscription operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchRegistryUpdated(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "RegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubRegistryUpdated)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
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

// ParseRegistryUpdated is a log parse operation binding the contract event 0x482b97c53e48ffa324a976e2738053e9aff6eee04d8aac63b10e19411d869b82.
//
// Solidity: event RegistryUpdated(address oldRegistry, address newRegistry)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseRegistryUpdated(log types.Log) (*WitnessChainWitnessHubRegistryUpdated, error) {
	event := new(WitnessChainWitnessHubRegistryUpdated)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "RegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubSetStrategyParamsIterator is returned from FilterSetStrategyParams and is used to iterate over the raw logs and unpacked data for SetStrategyParams events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubSetStrategyParamsIterator struct {
	Event *WitnessChainWitnessHubSetStrategyParams // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubSetStrategyParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubSetStrategyParams)
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
		it.Event = new(WitnessChainWitnessHubSetStrategyParams)
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
func (it *WitnessChainWitnessHubSetStrategyParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubSetStrategyParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubSetStrategyParams represents a SetStrategyParams event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubSetStrategyParams struct {
	Params []IWitnessHubStrategyParam
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetStrategyParams is a free log retrieval operation binding the contract event 0x78f9b85d45475399beaeb2ca0a5bb0eb3cdc5e60dd4be21f77e1257e641f33fb.
//
// Solidity: event SetStrategyParams((address,uint96)[] params)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterSetStrategyParams(opts *bind.FilterOpts) (*WitnessChainWitnessHubSetStrategyParamsIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "SetStrategyParams")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubSetStrategyParamsIterator{contract: _WitnessChainWitnessHub.contract, event: "SetStrategyParams", logs: logs, sub: sub}, nil
}

// WatchSetStrategyParams is a free log subscription operation binding the contract event 0x78f9b85d45475399beaeb2ca0a5bb0eb3cdc5e60dd4be21f77e1257e641f33fb.
//
// Solidity: event SetStrategyParams((address,uint96)[] params)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchSetStrategyParams(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubSetStrategyParams) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "SetStrategyParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubSetStrategyParams)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "SetStrategyParams", log); err != nil {
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

// ParseSetStrategyParams is a log parse operation binding the contract event 0x78f9b85d45475399beaeb2ca0a5bb0eb3cdc5e60dd4be21f77e1257e641f33fb.
//
// Solidity: event SetStrategyParams((address,uint96)[] params)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseSetStrategyParams(log types.Log) (*WitnessChainWitnessHubSetStrategyParams, error) {
	event := new(WitnessChainWitnessHubSetStrategyParams)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "SetStrategyParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubUnpausedIterator struct {
	Event *WitnessChainWitnessHubUnpaused // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubUnpaused)
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
		it.Event = new(WitnessChainWitnessHubUnpaused)
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
func (it *WitnessChainWitnessHubUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubUnpaused represents a Unpaused event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WitnessChainWitnessHubUnpausedIterator, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubUnpausedIterator{contract: _WitnessChainWitnessHub.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubUnpaused) (event.Subscription, error) {

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubUnpaused)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseUnpaused(log types.Log) (*WitnessChainWitnessHubUnpaused, error) {
	event := new(WitnessChainWitnessHubUnpaused)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WitnessChainWitnessHubUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubUpgradedIterator struct {
	Event *WitnessChainWitnessHubUpgraded // Event containing the contract specifics and raw log

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
func (it *WitnessChainWitnessHubUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WitnessChainWitnessHubUpgraded)
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
		it.Event = new(WitnessChainWitnessHubUpgraded)
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
func (it *WitnessChainWitnessHubUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WitnessChainWitnessHubUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WitnessChainWitnessHubUpgraded represents a Upgraded event raised by the WitnessChainWitnessHub contract.
type WitnessChainWitnessHubUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WitnessChainWitnessHubUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WitnessChainWitnessHubUpgradedIterator{contract: _WitnessChainWitnessHub.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WitnessChainWitnessHubUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WitnessChainWitnessHub.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WitnessChainWitnessHubUpgraded)
				if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_WitnessChainWitnessHub *WitnessChainWitnessHubFilterer) ParseUpgraded(log types.Log) (*WitnessChainWitnessHubUpgraded, error) {
	event := new(WitnessChainWitnessHubUpgraded)
	if err := _WitnessChainWitnessHub.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
