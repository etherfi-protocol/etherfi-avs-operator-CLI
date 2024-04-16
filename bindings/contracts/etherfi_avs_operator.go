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

// EtherFiAvsOperatorAvsInfo is an auto generated low-level Go binding around an user-defined struct.
type EtherFiAvsOperatorAvsInfo struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
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

// IRegistryCoordinatorOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperatorKickParam struct {
	QuorumNumber uint8
	Operator     common.Address
}

// EtherfiAVSOperatorMetaData contains all meta data concerning the EtherfiAVSOperator contract.
var EtherfiAVSOperatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"avsInfos\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsNodeRunner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperatorsManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ecdsaSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwardCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAvsInfo\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structEtherFiAvsOperator.AvsInfo\",\"components\":[{\"name\":\"isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_avsOperatorsManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAvsRegistered\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAvsWhitelisted\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRegisteredBlsKey\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidOperatorCall\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_remainingCalldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidSignature\",\"inputs\":[{\"name\":\"_digestHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"magicValue\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_detail\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_metaDataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerBlsKeyAsDelegatedNodeOperator\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorWithChurn\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operatorKickParams\",\"type\":\"tuple[]\",\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"components\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"_churnApproverSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"runnerForwardCall\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_remainingCalldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAvsNodeRunner\",\"inputs\":[{\"name\":\"_avsNodeRunner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAvsWhitelist\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateEcdsaSigner\",\"inputs\":[{\"name\":\"_ecdsaSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// EtherfiAVSOperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherfiAVSOperatorMetaData.ABI instead.
var EtherfiAVSOperatorABI = EtherfiAVSOperatorMetaData.ABI

// EtherfiAVSOperator is an auto generated Go binding around an Ethereum contract.
type EtherfiAVSOperator struct {
	EtherfiAVSOperatorCaller     // Read-only binding to the contract
	EtherfiAVSOperatorTransactor // Write-only binding to the contract
	EtherfiAVSOperatorFilterer   // Log filterer for contract events
}

// EtherfiAVSOperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAVSOperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAVSOperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherfiAVSOperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAVSOperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherfiAVSOperatorSession struct {
	Contract     *EtherfiAVSOperator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EtherfiAVSOperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherfiAVSOperatorCallerSession struct {
	Contract *EtherfiAVSOperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// EtherfiAVSOperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherfiAVSOperatorTransactorSession struct {
	Contract     *EtherfiAVSOperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// EtherfiAVSOperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherfiAVSOperatorRaw struct {
	Contract *EtherfiAVSOperator // Generic contract binding to access the raw methods on
}

// EtherfiAVSOperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorCallerRaw struct {
	Contract *EtherfiAVSOperatorCaller // Generic read-only contract binding to access the raw methods on
}

// EtherfiAVSOperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorTransactorRaw struct {
	Contract *EtherfiAVSOperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherfiAVSOperator creates a new instance of EtherfiAVSOperator, bound to a specific deployed contract.
func NewEtherfiAVSOperator(address common.Address, backend bind.ContractBackend) (*EtherfiAVSOperator, error) {
	contract, err := bindEtherfiAVSOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperator{EtherfiAVSOperatorCaller: EtherfiAVSOperatorCaller{contract: contract}, EtherfiAVSOperatorTransactor: EtherfiAVSOperatorTransactor{contract: contract}, EtherfiAVSOperatorFilterer: EtherfiAVSOperatorFilterer{contract: contract}}, nil
}

// NewEtherfiAVSOperatorCaller creates a new read-only instance of EtherfiAVSOperator, bound to a specific deployed contract.
func NewEtherfiAVSOperatorCaller(address common.Address, caller bind.ContractCaller) (*EtherfiAVSOperatorCaller, error) {
	contract, err := bindEtherfiAVSOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorCaller{contract: contract}, nil
}

// NewEtherfiAVSOperatorTransactor creates a new write-only instance of EtherfiAVSOperator, bound to a specific deployed contract.
func NewEtherfiAVSOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherfiAVSOperatorTransactor, error) {
	contract, err := bindEtherfiAVSOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorTransactor{contract: contract}, nil
}

// NewEtherfiAVSOperatorFilterer creates a new log filterer instance of EtherfiAVSOperator, bound to a specific deployed contract.
func NewEtherfiAVSOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherfiAVSOperatorFilterer, error) {
	contract, err := bindEtherfiAVSOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorFilterer{contract: contract}, nil
}

// bindEtherfiAVSOperator binds a generic wrapper to an already deployed contract.
func bindEtherfiAVSOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherfiAVSOperatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAVSOperator *EtherfiAVSOperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAVSOperator.Contract.EtherfiAVSOperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAVSOperator *EtherfiAVSOperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.EtherfiAVSOperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAVSOperator *EtherfiAVSOperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.EtherfiAVSOperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAVSOperator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.contract.Transact(opts, method, params...)
}

// AvsInfos is a free data retrieval call binding the contract method 0x9ffd8970.
//
// Solidity: function avsInfos(address ) view returns(bool isWhitelisted, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, bool isRegistered)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) AvsInfos(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
}, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "avsInfos", arg0)

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
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) AvsInfos(arg0 common.Address) (struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
}, error) {
	return _EtherfiAVSOperator.Contract.AvsInfos(&_EtherfiAVSOperator.CallOpts, arg0)
}

// AvsInfos is a free data retrieval call binding the contract method 0x9ffd8970.
//
// Solidity: function avsInfos(address ) view returns(bool isWhitelisted, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, bool isRegistered)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) AvsInfos(arg0 common.Address) (struct {
	IsWhitelisted bool
	QuorumNumbers []byte
	Socket        string
	Params        IBLSApkRegistryPubkeyRegistrationParams
	IsRegistered  bool
}, error) {
	return _EtherfiAVSOperator.Contract.AvsInfos(&_EtherfiAVSOperator.CallOpts, arg0)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0xd68db80a.
//
// Solidity: function avsNodeRunner() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) AvsNodeRunner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "avsNodeRunner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsNodeRunner is a free data retrieval call binding the contract method 0xd68db80a.
//
// Solidity: function avsNodeRunner() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) AvsNodeRunner() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.AvsNodeRunner(&_EtherfiAVSOperator.CallOpts)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0xd68db80a.
//
// Solidity: function avsNodeRunner() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) AvsNodeRunner() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.AvsNodeRunner(&_EtherfiAVSOperator.CallOpts)
}

// AvsOperatorsManager is a free data retrieval call binding the contract method 0x1065a488.
//
// Solidity: function avsOperatorsManager() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) AvsOperatorsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "avsOperatorsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsOperatorsManager is a free data retrieval call binding the contract method 0x1065a488.
//
// Solidity: function avsOperatorsManager() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) AvsOperatorsManager() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.AvsOperatorsManager(&_EtherfiAVSOperator.CallOpts)
}

// AvsOperatorsManager is a free data retrieval call binding the contract method 0x1065a488.
//
// Solidity: function avsOperatorsManager() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) AvsOperatorsManager() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.AvsOperatorsManager(&_EtherfiAVSOperator.CallOpts)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x7854c9f7.
//
// Solidity: function ecdsaSigner() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) EcdsaSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "ecdsaSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcdsaSigner is a free data retrieval call binding the contract method 0x7854c9f7.
//
// Solidity: function ecdsaSigner() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) EcdsaSigner() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.EcdsaSigner(&_EtherfiAVSOperator.CallOpts)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x7854c9f7.
//
// Solidity: function ecdsaSigner() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) EcdsaSigner() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.EcdsaSigner(&_EtherfiAVSOperator.CallOpts)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x5638c88e.
//
// Solidity: function getAvsInfo(address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) GetAvsInfo(opts *bind.CallOpts, _avsRegistryCoordinator common.Address) (EtherFiAvsOperatorAvsInfo, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "getAvsInfo", _avsRegistryCoordinator)

	if err != nil {
		return *new(EtherFiAvsOperatorAvsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EtherFiAvsOperatorAvsInfo)).(*EtherFiAvsOperatorAvsInfo)

	return out0, err

}

// GetAvsInfo is a free data retrieval call binding the contract method 0x5638c88e.
//
// Solidity: function getAvsInfo(address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) GetAvsInfo(_avsRegistryCoordinator common.Address) (EtherFiAvsOperatorAvsInfo, error) {
	return _EtherfiAVSOperator.Contract.GetAvsInfo(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x5638c88e.
//
// Solidity: function getAvsInfo(address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) GetAvsInfo(_avsRegistryCoordinator common.Address) (EtherFiAvsOperatorAvsInfo, error) {
	return _EtherfiAVSOperator.Contract.GetAvsInfo(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) Implementation() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.Implementation(&_EtherfiAVSOperator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) Implementation() (common.Address, error) {
	return _EtherfiAVSOperator.Contract.Implementation(&_EtherfiAVSOperator.CallOpts)
}

// IsAvsRegistered is a free data retrieval call binding the contract method 0xe3d276ab.
//
// Solidity: function isAvsRegistered(address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) IsAvsRegistered(opts *bind.CallOpts, _avsRegistryCoordinator common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "isAvsRegistered", _avsRegistryCoordinator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvsRegistered is a free data retrieval call binding the contract method 0xe3d276ab.
//
// Solidity: function isAvsRegistered(address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) IsAvsRegistered(_avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsAvsRegistered(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator)
}

// IsAvsRegistered is a free data retrieval call binding the contract method 0xe3d276ab.
//
// Solidity: function isAvsRegistered(address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) IsAvsRegistered(_avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsAvsRegistered(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator)
}

// IsAvsWhitelisted is a free data retrieval call binding the contract method 0xaf5ea0f4.
//
// Solidity: function isAvsWhitelisted(address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) IsAvsWhitelisted(opts *bind.CallOpts, _avsRegistryCoordinator common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "isAvsWhitelisted", _avsRegistryCoordinator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvsWhitelisted is a free data retrieval call binding the contract method 0xaf5ea0f4.
//
// Solidity: function isAvsWhitelisted(address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) IsAvsWhitelisted(_avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsAvsWhitelisted(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator)
}

// IsAvsWhitelisted is a free data retrieval call binding the contract method 0xaf5ea0f4.
//
// Solidity: function isAvsWhitelisted(address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) IsAvsWhitelisted(_avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsAvsWhitelisted(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator)
}

// IsRegisteredBlsKey is a free data retrieval call binding the contract method 0x26b64e1e.
//
// Solidity: function isRegisteredBlsKey(address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) IsRegisteredBlsKey(opts *bind.CallOpts, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "isRegisteredBlsKey", _avsRegistryCoordinator, _quorumNumbers, _socket, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredBlsKey is a free data retrieval call binding the contract method 0x26b64e1e.
//
// Solidity: function isRegisteredBlsKey(address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) IsRegisteredBlsKey(_avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsRegisteredBlsKey(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// IsRegisteredBlsKey is a free data retrieval call binding the contract method 0x26b64e1e.
//
// Solidity: function isRegisteredBlsKey(address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) IsRegisteredBlsKey(_avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsRegisteredBlsKey(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// IsValidOperatorCall is a free data retrieval call binding the contract method 0x498a34ad.
//
// Solidity: function isValidOperatorCall(address _avsRegistryCoordinator, bytes4 _selector, bytes _remainingCalldata) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) IsValidOperatorCall(opts *bind.CallOpts, _avsRegistryCoordinator common.Address, _selector [4]byte, _remainingCalldata []byte) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "isValidOperatorCall", _avsRegistryCoordinator, _selector, _remainingCalldata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOperatorCall is a free data retrieval call binding the contract method 0x498a34ad.
//
// Solidity: function isValidOperatorCall(address _avsRegistryCoordinator, bytes4 _selector, bytes _remainingCalldata) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) IsValidOperatorCall(_avsRegistryCoordinator common.Address, _selector [4]byte, _remainingCalldata []byte) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsValidOperatorCall(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator, _selector, _remainingCalldata)
}

// IsValidOperatorCall is a free data retrieval call binding the contract method 0x498a34ad.
//
// Solidity: function isValidOperatorCall(address _avsRegistryCoordinator, bytes4 _selector, bytes _remainingCalldata) view returns(bool)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) IsValidOperatorCall(_avsRegistryCoordinator common.Address, _selector [4]byte, _remainingCalldata []byte) (bool, error) {
	return _EtherfiAVSOperator.Contract.IsValidOperatorCall(&_EtherfiAVSOperator.CallOpts, _avsRegistryCoordinator, _selector, _remainingCalldata)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _digestHash, bytes _signature) view returns(bytes4 magicValue)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCaller) IsValidSignature(opts *bind.CallOpts, _digestHash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _EtherfiAVSOperator.contract.Call(opts, &out, "isValidSignature", _digestHash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _digestHash, bytes _signature) view returns(bytes4 magicValue)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) IsValidSignature(_digestHash [32]byte, _signature []byte) ([4]byte, error) {
	return _EtherfiAVSOperator.Contract.IsValidSignature(&_EtherfiAVSOperator.CallOpts, _digestHash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _digestHash, bytes _signature) view returns(bytes4 magicValue)
func (_EtherfiAVSOperator *EtherfiAVSOperatorCallerSession) IsValidSignature(_digestHash [32]byte, _signature []byte) ([4]byte, error) {
	return _EtherfiAVSOperator.Contract.IsValidSignature(&_EtherfiAVSOperator.CallOpts, _digestHash, _signature)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) DeregisterOperator(opts *bind.TransactOpts, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "deregisterOperator", _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) DeregisterOperator(_avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.DeregisterOperator(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) DeregisterOperator(_avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.DeregisterOperator(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, quorumNumbers)
}

// ForwardCall is a paid mutator transaction binding the contract method 0x22bee494.
//
// Solidity: function forwardCall(address to, bytes data) returns(bytes)
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) ForwardCall(opts *bind.TransactOpts, to common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "forwardCall", to, data)
}

// ForwardCall is a paid mutator transaction binding the contract method 0x22bee494.
//
// Solidity: function forwardCall(address to, bytes data) returns(bytes)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) ForwardCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.ForwardCall(&_EtherfiAVSOperator.TransactOpts, to, data)
}

// ForwardCall is a paid mutator transaction binding the contract method 0x22bee494.
//
// Solidity: function forwardCall(address to, bytes data) returns(bytes)
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) ForwardCall(to common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.ForwardCall(&_EtherfiAVSOperator.TransactOpts, to, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _avsOperatorsManager) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) Initialize(opts *bind.TransactOpts, _avsOperatorsManager common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "initialize", _avsOperatorsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _avsOperatorsManager) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) Initialize(_avsOperatorsManager common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.Initialize(&_EtherfiAVSOperator.TransactOpts, _avsOperatorsManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _avsOperatorsManager) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) Initialize(_avsOperatorsManager common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.Initialize(&_EtherfiAVSOperator.TransactOpts, _avsOperatorsManager)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xc69de2b1.
//
// Solidity: function modifyOperatorDetails(address _delegationManager, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, _delegationManager common.Address, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "modifyOperatorDetails", _delegationManager, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xc69de2b1.
//
// Solidity: function modifyOperatorDetails(address _delegationManager, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) ModifyOperatorDetails(_delegationManager common.Address, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.ModifyOperatorDetails(&_EtherfiAVSOperator.TransactOpts, _delegationManager, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0xc69de2b1.
//
// Solidity: function modifyOperatorDetails(address _delegationManager, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) ModifyOperatorDetails(_delegationManager common.Address, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.ModifyOperatorDetails(&_EtherfiAVSOperator.TransactOpts, _delegationManager, _newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xd905faf1.
//
// Solidity: function registerAsOperator(address _delegationManager, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) RegisterAsOperator(opts *bind.TransactOpts, _delegationManager common.Address, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "registerAsOperator", _delegationManager, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xd905faf1.
//
// Solidity: function registerAsOperator(address _delegationManager, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) RegisterAsOperator(_delegationManager common.Address, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterAsOperator(&_EtherfiAVSOperator.TransactOpts, _delegationManager, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0xd905faf1.
//
// Solidity: function registerAsOperator(address _delegationManager, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) RegisterAsOperator(_delegationManager common.Address, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterAsOperator(&_EtherfiAVSOperator.TransactOpts, _delegationManager, _detail, _metaDataURI)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0x62815a4c.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) RegisterBlsKeyAsDelegatedNodeOperator(opts *bind.TransactOpts, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "registerBlsKeyAsDelegatedNodeOperator", _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0x62815a4c.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) RegisterBlsKeyAsDelegatedNodeOperator(_avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0x62815a4c.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) RegisterBlsKeyAsDelegatedNodeOperator(_avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xaa2bcdbd.
//
// Solidity: function registerOperator(address _avsRegistryCoordinator, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) RegisterOperator(opts *bind.TransactOpts, _avsRegistryCoordinator common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "registerOperator", _avsRegistryCoordinator, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xaa2bcdbd.
//
// Solidity: function registerOperator(address _avsRegistryCoordinator, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) RegisterOperator(_avsRegistryCoordinator common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterOperator(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xaa2bcdbd.
//
// Solidity: function registerOperator(address _avsRegistryCoordinator, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) RegisterOperator(_avsRegistryCoordinator common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterOperator(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x041d2f68.
//
// Solidity: function registerOperatorWithChurn(address _avsRegistryCoordinator, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, _avsRegistryCoordinator common.Address, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "registerOperatorWithChurn", _avsRegistryCoordinator, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x041d2f68.
//
// Solidity: function registerOperatorWithChurn(address _avsRegistryCoordinator, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) RegisterOperatorWithChurn(_avsRegistryCoordinator common.Address, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterOperatorWithChurn(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0x041d2f68.
//
// Solidity: function registerOperatorWithChurn(address _avsRegistryCoordinator, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) RegisterOperatorWithChurn(_avsRegistryCoordinator common.Address, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RegisterOperatorWithChurn(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RunnerForwardCall is a paid mutator transaction binding the contract method 0x4e995e12.
//
// Solidity: function runnerForwardCall(address _target, bytes4 _selector, bytes _remainingCalldata) returns(bytes)
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) RunnerForwardCall(opts *bind.TransactOpts, _target common.Address, _selector [4]byte, _remainingCalldata []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "runnerForwardCall", _target, _selector, _remainingCalldata)
}

// RunnerForwardCall is a paid mutator transaction binding the contract method 0x4e995e12.
//
// Solidity: function runnerForwardCall(address _target, bytes4 _selector, bytes _remainingCalldata) returns(bytes)
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) RunnerForwardCall(_target common.Address, _selector [4]byte, _remainingCalldata []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RunnerForwardCall(&_EtherfiAVSOperator.TransactOpts, _target, _selector, _remainingCalldata)
}

// RunnerForwardCall is a paid mutator transaction binding the contract method 0x4e995e12.
//
// Solidity: function runnerForwardCall(address _target, bytes4 _selector, bytes _remainingCalldata) returns(bytes)
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) RunnerForwardCall(_target common.Address, _selector [4]byte, _remainingCalldata []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.RunnerForwardCall(&_EtherfiAVSOperator.TransactOpts, _target, _selector, _remainingCalldata)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x77b98c66.
//
// Solidity: function updateAvsNodeRunner(address _avsNodeRunner) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) UpdateAvsNodeRunner(opts *bind.TransactOpts, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "updateAvsNodeRunner", _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x77b98c66.
//
// Solidity: function updateAvsNodeRunner(address _avsNodeRunner) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) UpdateAvsNodeRunner(_avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateAvsNodeRunner(&_EtherfiAVSOperator.TransactOpts, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x77b98c66.
//
// Solidity: function updateAvsNodeRunner(address _avsNodeRunner) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) UpdateAvsNodeRunner(_avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateAvsNodeRunner(&_EtherfiAVSOperator.TransactOpts, _avsNodeRunner)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0x13739972.
//
// Solidity: function updateAvsWhitelist(address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) UpdateAvsWhitelist(opts *bind.TransactOpts, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "updateAvsWhitelist", _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0x13739972.
//
// Solidity: function updateAvsWhitelist(address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) UpdateAvsWhitelist(_avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateAvsWhitelist(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0x13739972.
//
// Solidity: function updateAvsWhitelist(address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) UpdateAvsWhitelist(_avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateAvsWhitelist(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0x215221f5.
//
// Solidity: function updateEcdsaSigner(address _ecdsaSigner) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) UpdateEcdsaSigner(opts *bind.TransactOpts, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "updateEcdsaSigner", _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0x215221f5.
//
// Solidity: function updateEcdsaSigner(address _ecdsaSigner) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) UpdateEcdsaSigner(_ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateEcdsaSigner(&_EtherfiAVSOperator.TransactOpts, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0x215221f5.
//
// Solidity: function updateEcdsaSigner(address _ecdsaSigner) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) UpdateEcdsaSigner(_ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateEcdsaSigner(&_EtherfiAVSOperator.TransactOpts, _ecdsaSigner)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address _delegationManager, string _metadataURI) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, _delegationManager common.Address, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "updateOperatorMetadataURI", _delegationManager, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address _delegationManager, string _metadataURI) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) UpdateOperatorMetadataURI(_delegationManager common.Address, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateOperatorMetadataURI(&_EtherfiAVSOperator.TransactOpts, _delegationManager, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address _delegationManager, string _metadataURI) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) UpdateOperatorMetadataURI(_delegationManager common.Address, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateOperatorMetadataURI(&_EtherfiAVSOperator.TransactOpts, _delegationManager, _metadataURI)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactor) UpdateSocket(opts *bind.TransactOpts, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.contract.Transact(opts, "updateSocket", _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorSession) UpdateSocket(_avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateSocket(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAVSOperator *EtherfiAVSOperatorTransactorSession) UpdateSocket(_avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAVSOperator.Contract.UpdateSocket(&_EtherfiAVSOperator.TransactOpts, _avsRegistryCoordinator, _socket)
}
