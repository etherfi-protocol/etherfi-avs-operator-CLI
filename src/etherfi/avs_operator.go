// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package etherfi

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

// AvsOperatorAvsInfo is an auto generated low-level Go binding around an user-defined struct.
type AvsOperatorAvsInfo struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
}

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// IDelegationManagerOperatorDetails is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerOperatorDetails struct {
	EarningsReceiver         common.Address
	DelegationApprover       common.Address
	StakerOptOutWindowBlocks uint32
}

// AvsOperatorMetaData contains all meta data concerning the AvsOperator contract.
var AvsOperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"avsInfos\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsNodeRunner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsOperatorsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecdsaSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"forwardCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"}],\"name\":\"getAvsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structAvsOperator.AvsInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avsOperatorsManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_digestHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"modifyOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_detail\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_metaDataURI\",\"type\":\"string\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avsNodeRunner\",\"type\":\"address\"}],\"name\":\"updateAvsNodeRunner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ecdsaSigner\",\"type\":\"address\"}],\"name\":\"updateEcdsaSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateOperatorMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"verifyBlsKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"verifyBlsKeyAgainstHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AvsOperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AvsOperatorMetaData.ABI instead.
var AvsOperatorABI = AvsOperatorMetaData.ABI

// AvsOperator is an auto generated Go binding around an Ethereum contract.
type AvsOperator struct {
	AvsOperatorCaller     // Read-only binding to the contract
	AvsOperatorTransactor // Write-only binding to the contract
	AvsOperatorFilterer   // Log filterer for contract events
}

// AvsOperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvsOperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsOperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvsOperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsOperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvsOperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsOperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvsOperatorSession struct {
	Contract     *AvsOperator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvsOperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvsOperatorCallerSession struct {
	Contract *AvsOperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AvsOperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvsOperatorTransactorSession struct {
	Contract     *AvsOperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AvsOperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvsOperatorRaw struct {
	Contract *AvsOperator // Generic contract binding to access the raw methods on
}

// AvsOperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvsOperatorCallerRaw struct {
	Contract *AvsOperatorCaller // Generic read-only contract binding to access the raw methods on
}

// AvsOperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvsOperatorTransactorRaw struct {
	Contract *AvsOperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvsOperator creates a new instance of AvsOperator, bound to a specific deployed contract.
func NewAvsOperator(address common.Address, backend bind.ContractBackend) (*AvsOperator, error) {
	contract, err := bindAvsOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvsOperator{AvsOperatorCaller: AvsOperatorCaller{contract: contract}, AvsOperatorTransactor: AvsOperatorTransactor{contract: contract}, AvsOperatorFilterer: AvsOperatorFilterer{contract: contract}}, nil
}

// NewAvsOperatorCaller creates a new read-only instance of AvsOperator, bound to a specific deployed contract.
func NewAvsOperatorCaller(address common.Address, caller bind.ContractCaller) (*AvsOperatorCaller, error) {
	contract, err := bindAvsOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorCaller{contract: contract}, nil
}

// NewAvsOperatorTransactor creates a new write-only instance of AvsOperator, bound to a specific deployed contract.
func NewAvsOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AvsOperatorTransactor, error) {
	contract, err := bindAvsOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorTransactor{contract: contract}, nil
}

// NewAvsOperatorFilterer creates a new log filterer instance of AvsOperator, bound to a specific deployed contract.
func NewAvsOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AvsOperatorFilterer, error) {
	contract, err := bindAvsOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorFilterer{contract: contract}, nil
}

// bindAvsOperator binds a generic wrapper to an already deployed contract.
func bindAvsOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvsOperatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvsOperator *AvsOperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvsOperator.Contract.AvsOperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvsOperator *AvsOperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsOperator.Contract.AvsOperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvsOperator *AvsOperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvsOperator.Contract.AvsOperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvsOperator *AvsOperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvsOperator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvsOperator *AvsOperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsOperator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvsOperator *AvsOperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvsOperator.Contract.contract.Transact(opts, method, params...)
}

// AvsInfos is a free data retrieval call binding the contract method 0x9ffd8970.
//
// Solidity: function avsInfos(address ) view returns(bool isWhitelisted, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, bool isRegistered)
func (_AvsOperator *AvsOperatorCaller) AvsInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
}, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "avsInfos", arg0)

	outstruct := new(struct {
		IsWhitelisted bool
		QuorumNumbers []byte
		Socket        string
		Params        IBLSApkRegistryPubkeyRegistrationParams
		IsRegistered  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsWhitelisted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.QuorumNumbers = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Socket = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Params = *abi.ConvertType(out[3], new(IBLSApkRegistryPubkeyRegistrationParams)).(*IBLSApkRegistryPubkeyRegistrationParams)
	outstruct.IsRegistered = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// AvsInfos is a free data retrieval call binding the contract method 0x9ffd8970.
//
// Solidity: function avsInfos(address ) view returns(bool isWhitelisted, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, bool isRegistered)
func (_AvsOperator *AvsOperatorSession) AvsInfos(arg0 common.Address) (struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
}, error) {
	return _AvsOperator.Contract.AvsInfos(&_AvsOperator.CallOpts, arg0)
}

// AvsInfos is a free data retrieval call binding the contract method 0x9ffd8970.
//
// Solidity: function avsInfos(address ) view returns(bool isWhitelisted, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, bool isRegistered)
func (_AvsOperator *AvsOperatorCallerSession) AvsInfos(arg0 common.Address) (struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
}, error) {
	return _AvsOperator.Contract.AvsInfos(&_AvsOperator.CallOpts, arg0)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0xd68db80a.
//
// Solidity: function avsNodeRunner() view returns(address)
func (_AvsOperator *AvsOperatorCaller) AvsNodeRunner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "avsNodeRunner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsNodeRunner is a free data retrieval call binding the contract method 0xd68db80a.
//
// Solidity: function avsNodeRunner() view returns(address)
func (_AvsOperator *AvsOperatorSession) AvsNodeRunner() (common.Address, error) {
	return _AvsOperator.Contract.AvsNodeRunner(&_AvsOperator.CallOpts)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0xd68db80a.
//
// Solidity: function avsNodeRunner() view returns(address)
func (_AvsOperator *AvsOperatorCallerSession) AvsNodeRunner() (common.Address, error) {
	return _AvsOperator.Contract.AvsNodeRunner(&_AvsOperator.CallOpts)
}

// AvsOperatorsManager is a free data retrieval call binding the contract method 0x1065a488.
//
// Solidity: function avsOperatorsManager() view returns(address)
func (_AvsOperator *AvsOperatorCaller) AvsOperatorsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "avsOperatorsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsOperatorsManager is a free data retrieval call binding the contract method 0x1065a488.
//
// Solidity: function avsOperatorsManager() view returns(address)
func (_AvsOperator *AvsOperatorSession) AvsOperatorsManager() (common.Address, error) {
	return _AvsOperator.Contract.AvsOperatorsManager(&_AvsOperator.CallOpts)
}

// AvsOperatorsManager is a free data retrieval call binding the contract method 0x1065a488.
//
// Solidity: function avsOperatorsManager() view returns(address)
func (_AvsOperator *AvsOperatorCallerSession) AvsOperatorsManager() (common.Address, error) {
	return _AvsOperator.Contract.AvsOperatorsManager(&_AvsOperator.CallOpts)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x7854c9f7.
//
// Solidity: function ecdsaSigner() view returns(address)
func (_AvsOperator *AvsOperatorCaller) EcdsaSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "ecdsaSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcdsaSigner is a free data retrieval call binding the contract method 0x7854c9f7.
//
// Solidity: function ecdsaSigner() view returns(address)
func (_AvsOperator *AvsOperatorSession) EcdsaSigner() (common.Address, error) {
	return _AvsOperator.Contract.EcdsaSigner(&_AvsOperator.CallOpts)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x7854c9f7.
//
// Solidity: function ecdsaSigner() view returns(address)
func (_AvsOperator *AvsOperatorCallerSession) EcdsaSigner() (common.Address, error) {
	return _AvsOperator.Contract.EcdsaSigner(&_AvsOperator.CallOpts)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x5638c88e.
//
// Solidity: function getAvsInfo(address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_AvsOperator *AvsOperatorCaller) GetAvsInfo(opts *bind.CallOpts, _avsRegistryCoordinator common.Address) (AvsOperatorAvsInfo, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "getAvsInfo", _avsRegistryCoordinator)

	if err != nil {
		return *new(AvsOperatorAvsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AvsOperatorAvsInfo)).(*AvsOperatorAvsInfo)

	return out0, err

}

// GetAvsInfo is a free data retrieval call binding the contract method 0x5638c88e.
//
// Solidity: function getAvsInfo(address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_AvsOperator *AvsOperatorSession) GetAvsInfo(_avsRegistryCoordinator common.Address) (AvsOperatorAvsInfo, error) {
	return _AvsOperator.Contract.GetAvsInfo(&_AvsOperator.CallOpts, _avsRegistryCoordinator)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x5638c88e.
//
// Solidity: function getAvsInfo(address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_AvsOperator *AvsOperatorCallerSession) GetAvsInfo(_avsRegistryCoordinator common.Address) (AvsOperatorAvsInfo, error) {
	return _AvsOperator.Contract.GetAvsInfo(&_AvsOperator.CallOpts, _avsRegistryCoordinator)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AvsOperator *AvsOperatorCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AvsOperator *AvsOperatorSession) Implementation() (common.Address, error) {
	return _AvsOperator.Contract.Implementation(&_AvsOperator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AvsOperator *AvsOperatorCallerSession) Implementation() (common.Address, error) {
	return _AvsOperator.Contract.Implementation(&_AvsOperator.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _digestHash, bytes _signature) view returns(bytes4 magicValue)
func (_AvsOperator *AvsOperatorCaller) IsValidSignature(opts *bind.CallOpts, _digestHash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "isValidSignature", _digestHash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _digestHash, bytes _signature) view returns(bytes4 magicValue)
func (_AvsOperator *AvsOperatorSession) IsValidSignature(_digestHash [32]byte, _signature []byte) ([4]byte, error) {
	return _AvsOperator.Contract.IsValidSignature(&_AvsOperator.CallOpts, _digestHash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _digestHash, bytes _signature) view returns(bytes4 magicValue)
func (_AvsOperator *AvsOperatorCallerSession) IsValidSignature(_digestHash [32]byte, _signature []byte) ([4]byte, error) {
	return _AvsOperator.Contract.IsValidSignature(&_AvsOperator.CallOpts, _digestHash, _signature)
}

// VerifyBlsKey is a free data retrieval call binding the contract method 0xd7a31a42.
//
// Solidity: function verifyBlsKey(address registryCoordinator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params) view returns(bool)
func (_AvsOperator *AvsOperatorCaller) VerifyBlsKey(opts *bind.CallOpts, registryCoordinator common.Address, params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "verifyBlsKey", registryCoordinator, params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlsKey is a free data retrieval call binding the contract method 0xd7a31a42.
//
// Solidity: function verifyBlsKey(address registryCoordinator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params) view returns(bool)
func (_AvsOperator *AvsOperatorSession) VerifyBlsKey(registryCoordinator common.Address, params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _AvsOperator.Contract.VerifyBlsKey(&_AvsOperator.CallOpts, registryCoordinator, params)
}

// VerifyBlsKey is a free data retrieval call binding the contract method 0xd7a31a42.
//
// Solidity: function verifyBlsKey(address registryCoordinator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params) view returns(bool)
func (_AvsOperator *AvsOperatorCallerSession) VerifyBlsKey(registryCoordinator common.Address, params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _AvsOperator.Contract.VerifyBlsKey(&_AvsOperator.CallOpts, registryCoordinator, params)
}

// VerifyBlsKeyAgainstHash is a free data retrieval call binding the contract method 0x8fdf3cbe.
//
// Solidity: function verifyBlsKeyAgainstHash((uint256,uint256) pubkeyRegistrationMessageHash, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params) view returns(bool)
func (_AvsOperator *AvsOperatorCaller) VerifyBlsKeyAgainstHash(opts *bind.CallOpts, pubkeyRegistrationMessageHash BN254G1Point, params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	var out []interface{}
	err := _AvsOperator.contract.Call(opts, &out, "verifyBlsKeyAgainstHash", pubkeyRegistrationMessageHash, params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlsKeyAgainstHash is a free data retrieval call binding the contract method 0x8fdf3cbe.
//
// Solidity: function verifyBlsKeyAgainstHash((uint256,uint256) pubkeyRegistrationMessageHash, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params) view returns(bool)
func (_AvsOperator *AvsOperatorSession) VerifyBlsKeyAgainstHash(pubkeyRegistrationMessageHash BN254G1Point, params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _AvsOperator.Contract.VerifyBlsKeyAgainstHash(&_AvsOperator.CallOpts, pubkeyRegistrationMessageHash, params)
}

// VerifyBlsKeyAgainstHash is a free data retrieval call binding the contract method 0x8fdf3cbe.
//
// Solidity: function verifyBlsKeyAgainstHash((uint256,uint256) pubkeyRegistrationMessageHash, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params) view returns(bool)
func (_AvsOperator *AvsOperatorCallerSession) VerifyBlsKeyAgainstHash(pubkeyRegistrationMessageHash BN254G1Point, params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _AvsOperator.Contract.VerifyBlsKeyAgainstHash(&_AvsOperator.CallOpts, pubkeyRegistrationMessageHash, params)
}

// ForwardCall is a paid mutator transaction binding the contract method 0x22bee494.
//
// Solidity: function forwardCall(address to, bytes data) returns(bytes)
func (_AvsOperator *AvsOperatorTransactor) ForwardCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _AvsOperator.contract.Transact(opts, "forwardCall", to, data)
}

// ForwardCall is a paid mutator transaction binding the contract method 0x22bee494.
//
// Solidity: function forwardCall(address to, bytes data) returns(bytes)
func (_AvsOperator *AvsOperatorSession) ForwardCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _AvsOperator.Contract.ForwardCall(&_AvsOperator.TransactOpts, to, data)
}

// ForwardCall is a paid mutator transaction binding the contract method 0x22bee494.
//
// Solidity: function forwardCall(address to, bytes data) returns(bytes)
func (_AvsOperator *AvsOperatorTransactorSession) ForwardCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _AvsOperator.Contract.ForwardCall(&_AvsOperator.TransactOpts, to, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _avsOperatorsManager) returns()
func (_AvsOperator *AvsOperatorTransactor) Initialize(opts *bind.TransactOpts, _avsOperatorsManager common.Address) (*types.Transaction, error) {
	return _AvsOperator.contract.Transact(opts, "initialize", _avsOperatorsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _avsOperatorsManager) returns()
func (_AvsOperator *AvsOperatorSession) Initialize(_avsOperatorsManager common.Address) (*types.Transaction, error) {
	return _AvsOperator.Contract.Initialize(&_AvsOperator.TransactOpts, _avsOperatorsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _avsOperatorsManager) returns()
func (_AvsOperator *AvsOperatorTransactorSession) Initialize(_avsOperatorsManager common.Address) (*types.Transaction, error) {
	return _AvsOperator.Contract.Initialize(&_AvsOperator.TransactOpts, _avsOperatorsManager)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xc69de2b1.
//
// Solidity: function modifyOperatorDetails(address _delegationManager, (address,address,uint32) _newOperatorDetails) returns()
func (_AvsOperator *AvsOperatorTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, _delegationManager common.Address, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AvsOperator.contract.Transact(opts, "modifyOperatorDetails", _delegationManager, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xc69de2b1.
//
// Solidity: function modifyOperatorDetails(address _delegationManager, (address,address,uint32) _newOperatorDetails) returns()
func (_AvsOperator *AvsOperatorSession) ModifyOperatorDetails(_delegationManager common.Address, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AvsOperator.Contract.ModifyOperatorDetails(&_AvsOperator.TransactOpts, _delegationManager, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xc69de2b1.
//
// Solidity: function modifyOperatorDetails(address _delegationManager, (address,address,uint32) _newOperatorDetails) returns()
func (_AvsOperator *AvsOperatorTransactorSession) ModifyOperatorDetails(_delegationManager common.Address, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AvsOperator.Contract.ModifyOperatorDetails(&_AvsOperator.TransactOpts, _delegationManager, _newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xd905faf1.
//
// Solidity: function registerAsOperator(address _delegationManager, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AvsOperator *AvsOperatorTransactor) RegisterAsOperator(opts *bind.TransactOpts, _delegationManager common.Address, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AvsOperator.contract.Transact(opts, "registerAsOperator", _delegationManager, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xd905faf1.
//
// Solidity: function registerAsOperator(address _delegationManager, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AvsOperator *AvsOperatorSession) RegisterAsOperator(_delegationManager common.Address, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AvsOperator.Contract.RegisterAsOperator(&_AvsOperator.TransactOpts, _delegationManager, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xd905faf1.
//
// Solidity: function registerAsOperator(address _delegationManager, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AvsOperator *AvsOperatorTransactorSession) RegisterAsOperator(_delegationManager common.Address, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AvsOperator.Contract.RegisterAsOperator(&_AvsOperator.TransactOpts, _delegationManager, _detail, _metaDataURI)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x77b98c66.
//
// Solidity: function updateAvsNodeRunner(address _avsNodeRunner) returns()
func (_AvsOperator *AvsOperatorTransactor) UpdateAvsNodeRunner(opts *bind.TransactOpts, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AvsOperator.contract.Transact(opts, "updateAvsNodeRunner", _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x77b98c66.
//
// Solidity: function updateAvsNodeRunner(address _avsNodeRunner) returns()
func (_AvsOperator *AvsOperatorSession) UpdateAvsNodeRunner(_avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AvsOperator.Contract.UpdateAvsNodeRunner(&_AvsOperator.TransactOpts, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x77b98c66.
//
// Solidity: function updateAvsNodeRunner(address _avsNodeRunner) returns()
func (_AvsOperator *AvsOperatorTransactorSession) UpdateAvsNodeRunner(_avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AvsOperator.Contract.UpdateAvsNodeRunner(&_AvsOperator.TransactOpts, _avsNodeRunner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0x215221f5.
//
// Solidity: function updateEcdsaSigner(address _ecdsaSigner) returns()
func (_AvsOperator *AvsOperatorTransactor) UpdateEcdsaSigner(opts *bind.TransactOpts, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AvsOperator.contract.Transact(opts, "updateEcdsaSigner", _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0x215221f5.
//
// Solidity: function updateEcdsaSigner(address _ecdsaSigner) returns()
func (_AvsOperator *AvsOperatorSession) UpdateEcdsaSigner(_ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AvsOperator.Contract.UpdateEcdsaSigner(&_AvsOperator.TransactOpts, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0x215221f5.
//
// Solidity: function updateEcdsaSigner(address _ecdsaSigner) returns()
func (_AvsOperator *AvsOperatorTransactorSession) UpdateEcdsaSigner(_ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AvsOperator.Contract.UpdateEcdsaSigner(&_AvsOperator.TransactOpts, _ecdsaSigner)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address _delegationManager, string _metadataURI) returns()
func (_AvsOperator *AvsOperatorTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, _delegationManager common.Address, _metadataURI string) (*types.Transaction, error) {
	return _AvsOperator.contract.Transact(opts, "updateOperatorMetadataURI", _delegationManager, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address _delegationManager, string _metadataURI) returns()
func (_AvsOperator *AvsOperatorSession) UpdateOperatorMetadataURI(_delegationManager common.Address, _metadataURI string) (*types.Transaction, error) {
	return _AvsOperator.Contract.UpdateOperatorMetadataURI(&_AvsOperator.TransactOpts, _delegationManager, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address _delegationManager, string _metadataURI) returns()
func (_AvsOperator *AvsOperatorTransactorSession) UpdateOperatorMetadataURI(_delegationManager common.Address, _metadataURI string) (*types.Transaction, error) {
	return _AvsOperator.Contract.UpdateOperatorMetadataURI(&_AvsOperator.TransactOpts, _delegationManager, _metadataURI)
}
