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

// AvsOperatorManagerMetaData contains all meta data concerning the AvsOperatorManager contract.
var AvsOperatorManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adminForwardCall\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowedOperatorCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsNodeRunner\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_avsServiceManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAvsOperator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_avsServiceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ecdsaSigner\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwardOperatorCall\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardOperatorCall\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_args\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAvsInfo\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structAvsOperator.AvsInfo\",\"components\":[{\"name\":\"isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"isRegistered\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_etherFiAvsOperatorImpl\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeAvsDirectory\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"instantiateEtherFiAvsOperator\",\"inputs\":[{\"name\":\"_nums\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_ids\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isValidOperatorCall\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_newOperatorDetails\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextAvsOperatorId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorDetails\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pausers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_detail\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"_metaDataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAdmin\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isAdmin\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedOperatorCalls\",\"inputs\":[{\"name\":\"_operatorId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAvsNodeRunner\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_avsNodeRunner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateEcdsaSigner\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_ecdsaSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradableBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractUpgradeableBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeEtherFiAvsOperator\",\"inputs\":[{\"name\":\"_newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminUpdated\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isAdmin\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowedOperatorCallsUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreatedEtherFiAvsOperator\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"etherFiAvsOperator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwardedOperatorCall\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModifiedOperatorDetails\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newOperatorDetails\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegisteredAsOperator\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"detail\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"components\":[{\"name\":\"earningsReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedAvsNodeRunner\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"avsNodeRunner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedEcdsaSigner\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"ecdsaSigner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedOperatorMetadataURI\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidOperatorCall\",\"inputs\":[]}]",
}

// AvsOperatorManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AvsOperatorManagerMetaData.ABI instead.
var AvsOperatorManagerABI = AvsOperatorManagerMetaData.ABI

// AvsOperatorManager is an auto generated Go binding around an Ethereum contract.
type AvsOperatorManager struct {
	AvsOperatorManagerCaller     // Read-only binding to the contract
	AvsOperatorManagerTransactor // Write-only binding to the contract
	AvsOperatorManagerFilterer   // Log filterer for contract events
}

// AvsOperatorManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvsOperatorManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsOperatorManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvsOperatorManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsOperatorManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvsOperatorManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsOperatorManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvsOperatorManagerSession struct {
	Contract     *AvsOperatorManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AvsOperatorManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvsOperatorManagerCallerSession struct {
	Contract *AvsOperatorManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AvsOperatorManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvsOperatorManagerTransactorSession struct {
	Contract     *AvsOperatorManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AvsOperatorManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvsOperatorManagerRaw struct {
	Contract *AvsOperatorManager // Generic contract binding to access the raw methods on
}

// AvsOperatorManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvsOperatorManagerCallerRaw struct {
	Contract *AvsOperatorManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AvsOperatorManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvsOperatorManagerTransactorRaw struct {
	Contract *AvsOperatorManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvsOperatorManager creates a new instance of AvsOperatorManager, bound to a specific deployed contract.
func NewAvsOperatorManager(address common.Address, backend bind.ContractBackend) (*AvsOperatorManager, error) {
	contract, err := bindAvsOperatorManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManager{AvsOperatorManagerCaller: AvsOperatorManagerCaller{contract: contract}, AvsOperatorManagerTransactor: AvsOperatorManagerTransactor{contract: contract}, AvsOperatorManagerFilterer: AvsOperatorManagerFilterer{contract: contract}}, nil
}

// NewAvsOperatorManagerCaller creates a new read-only instance of AvsOperatorManager, bound to a specific deployed contract.
func NewAvsOperatorManagerCaller(address common.Address, caller bind.ContractCaller) (*AvsOperatorManagerCaller, error) {
	contract, err := bindAvsOperatorManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerCaller{contract: contract}, nil
}

// NewAvsOperatorManagerTransactor creates a new write-only instance of AvsOperatorManager, bound to a specific deployed contract.
func NewAvsOperatorManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AvsOperatorManagerTransactor, error) {
	contract, err := bindAvsOperatorManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerTransactor{contract: contract}, nil
}

// NewAvsOperatorManagerFilterer creates a new log filterer instance of AvsOperatorManager, bound to a specific deployed contract.
func NewAvsOperatorManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AvsOperatorManagerFilterer, error) {
	contract, err := bindAvsOperatorManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerFilterer{contract: contract}, nil
}

// bindAvsOperatorManager binds a generic wrapper to an already deployed contract.
func bindAvsOperatorManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvsOperatorManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvsOperatorManager *AvsOperatorManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvsOperatorManager.Contract.AvsOperatorManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvsOperatorManager *AvsOperatorManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.AvsOperatorManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvsOperatorManager *AvsOperatorManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.AvsOperatorManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvsOperatorManager *AvsOperatorManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvsOperatorManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvsOperatorManager *AvsOperatorManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvsOperatorManager *AvsOperatorManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerSession) Admins(arg0 common.Address) (bool, error) {
	return _AvsOperatorManager.Contract.Admins(&_AvsOperatorManager.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _AvsOperatorManager.Contract.Admins(&_AvsOperatorManager.CallOpts, arg0)
}

// AllowedOperatorCalls is a free data retrieval call binding the contract method 0x369b37f1.
//
// Solidity: function allowedOperatorCalls(uint256 , address , bytes4 ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCaller) AllowedOperatorCalls(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 [4]byte) (bool, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "allowedOperatorCalls", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOperatorCalls is a free data retrieval call binding the contract method 0x369b37f1.
//
// Solidity: function allowedOperatorCalls(uint256 , address , bytes4 ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerSession) AllowedOperatorCalls(arg0 *big.Int, arg1 common.Address, arg2 [4]byte) (bool, error) {
	return _AvsOperatorManager.Contract.AllowedOperatorCalls(&_AvsOperatorManager.CallOpts, arg0, arg1, arg2)
}

// AllowedOperatorCalls is a free data retrieval call binding the contract method 0x369b37f1.
//
// Solidity: function allowedOperatorCalls(uint256 , address , bytes4 ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) AllowedOperatorCalls(arg0 *big.Int, arg1 common.Address, arg2 [4]byte) (bool, error) {
	return _AvsOperatorManager.Contract.AllowedOperatorCalls(&_AvsOperatorManager.CallOpts, arg0, arg1, arg2)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerSession) AvsDirectory() (common.Address, error) {
	return _AvsOperatorManager.Contract.AvsDirectory(&_AvsOperatorManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _AvsOperatorManager.Contract.AvsDirectory(&_AvsOperatorManager.CallOpts)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0x05dba646.
//
// Solidity: function avsNodeRunner(uint256 _id) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCaller) AvsNodeRunner(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "avsNodeRunner", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsNodeRunner is a free data retrieval call binding the contract method 0x05dba646.
//
// Solidity: function avsNodeRunner(uint256 _id) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerSession) AvsNodeRunner(_id *big.Int) (common.Address, error) {
	return _AvsOperatorManager.Contract.AvsNodeRunner(&_AvsOperatorManager.CallOpts, _id)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0x05dba646.
//
// Solidity: function avsNodeRunner(uint256 _id) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) AvsNodeRunner(_id *big.Int) (common.Address, error) {
	return _AvsOperatorManager.Contract.AvsNodeRunner(&_AvsOperatorManager.CallOpts, _id)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0xd7eb63c7.
//
// Solidity: function avsOperatorStatus(uint256 _id, address _avsServiceManager) view returns(uint8)
func (_AvsOperatorManager *AvsOperatorManagerCaller) AvsOperatorStatus(opts *bind.CallOpts, _id *big.Int, _avsServiceManager common.Address) (uint8, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "avsOperatorStatus", _id, _avsServiceManager)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0xd7eb63c7.
//
// Solidity: function avsOperatorStatus(uint256 _id, address _avsServiceManager) view returns(uint8)
func (_AvsOperatorManager *AvsOperatorManagerSession) AvsOperatorStatus(_id *big.Int, _avsServiceManager common.Address) (uint8, error) {
	return _AvsOperatorManager.Contract.AvsOperatorStatus(&_AvsOperatorManager.CallOpts, _id, _avsServiceManager)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0xd7eb63c7.
//
// Solidity: function avsOperatorStatus(uint256 _id, address _avsServiceManager) view returns(uint8)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) AvsOperatorStatus(_id *big.Int, _avsServiceManager common.Address) (uint8, error) {
	return _AvsOperatorManager.Contract.AvsOperatorStatus(&_AvsOperatorManager.CallOpts, _id, _avsServiceManager)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCaller) AvsOperators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "avsOperators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _AvsOperatorManager.Contract.AvsOperators(&_AvsOperatorManager.CallOpts, arg0)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _AvsOperatorManager.Contract.AvsOperators(&_AvsOperatorManager.CallOpts, arg0)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xba393bb8.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(uint256 _id, address _avsServiceManager, bytes32 _salt, uint256 _expiry) view returns(bytes32)
func (_AvsOperatorManager *AvsOperatorManagerCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, _id *big.Int, _avsServiceManager common.Address, _salt [32]byte, _expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", _id, _avsServiceManager, _salt, _expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xba393bb8.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(uint256 _id, address _avsServiceManager, bytes32 _salt, uint256 _expiry) view returns(bytes32)
func (_AvsOperatorManager *AvsOperatorManagerSession) CalculateOperatorAVSRegistrationDigestHash(_id *big.Int, _avsServiceManager common.Address, _salt [32]byte, _expiry *big.Int) ([32]byte, error) {
	return _AvsOperatorManager.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AvsOperatorManager.CallOpts, _id, _avsServiceManager, _salt, _expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xba393bb8.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(uint256 _id, address _avsServiceManager, bytes32 _salt, uint256 _expiry) view returns(bytes32)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) CalculateOperatorAVSRegistrationDigestHash(_id *big.Int, _avsServiceManager common.Address, _salt [32]byte, _expiry *big.Int) ([32]byte, error) {
	return _AvsOperatorManager.Contract.CalculateOperatorAVSRegistrationDigestHash(&_AvsOperatorManager.CallOpts, _id, _avsServiceManager, _salt, _expiry)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerSession) DelegationManager() (common.Address, error) {
	return _AvsOperatorManager.Contract.DelegationManager(&_AvsOperatorManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) DelegationManager() (common.Address, error) {
	return _AvsOperatorManager.Contract.DelegationManager(&_AvsOperatorManager.CallOpts)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x5408cff6.
//
// Solidity: function ecdsaSigner(uint256 _id) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCaller) EcdsaSigner(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "ecdsaSigner", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcdsaSigner is a free data retrieval call binding the contract method 0x5408cff6.
//
// Solidity: function ecdsaSigner(uint256 _id) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerSession) EcdsaSigner(_id *big.Int) (common.Address, error) {
	return _AvsOperatorManager.Contract.EcdsaSigner(&_AvsOperatorManager.CallOpts, _id)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x5408cff6.
//
// Solidity: function ecdsaSigner(uint256 _id) view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) EcdsaSigner(_id *big.Int) (common.Address, error) {
	return _AvsOperatorManager.Contract.EcdsaSigner(&_AvsOperatorManager.CallOpts, _id)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x10eed82a.
//
// Solidity: function getAvsInfo(uint256 _id, address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_AvsOperatorManager *AvsOperatorManagerCaller) GetAvsInfo(opts *bind.CallOpts, _id *big.Int, _avsRegistryCoordinator common.Address) (AvsOperatorAvsInfo, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "getAvsInfo", _id, _avsRegistryCoordinator)

	if err != nil {
		return *new(AvsOperatorAvsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AvsOperatorAvsInfo)).(*AvsOperatorAvsInfo)

	return out0, err

}

// GetAvsInfo is a free data retrieval call binding the contract method 0x10eed82a.
//
// Solidity: function getAvsInfo(uint256 _id, address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_AvsOperatorManager *AvsOperatorManagerSession) GetAvsInfo(_id *big.Int, _avsRegistryCoordinator common.Address) (AvsOperatorAvsInfo, error) {
	return _AvsOperatorManager.Contract.GetAvsInfo(&_AvsOperatorManager.CallOpts, _id, _avsRegistryCoordinator)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x10eed82a.
//
// Solidity: function getAvsInfo(uint256 _id, address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) GetAvsInfo(_id *big.Int, _avsRegistryCoordinator common.Address) (AvsOperatorAvsInfo, error) {
	return _AvsOperatorManager.Contract.GetAvsInfo(&_AvsOperatorManager.CallOpts, _id, _avsRegistryCoordinator)
}

// IsValidOperatorCall is a free data retrieval call binding the contract method 0x4a44a53e.
//
// Solidity: function isValidOperatorCall(uint256 _id, address _target, bytes4 _selector, bytes ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCaller) IsValidOperatorCall(opts *bind.CallOpts, _id *big.Int, _target common.Address, _selector [4]byte, arg3 []byte) (bool, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "isValidOperatorCall", _id, _target, _selector, arg3)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidOperatorCall is a free data retrieval call binding the contract method 0x4a44a53e.
//
// Solidity: function isValidOperatorCall(uint256 _id, address _target, bytes4 _selector, bytes ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerSession) IsValidOperatorCall(_id *big.Int, _target common.Address, _selector [4]byte, arg3 []byte) (bool, error) {
	return _AvsOperatorManager.Contract.IsValidOperatorCall(&_AvsOperatorManager.CallOpts, _id, _target, _selector, arg3)
}

// IsValidOperatorCall is a free data retrieval call binding the contract method 0x4a44a53e.
//
// Solidity: function isValidOperatorCall(uint256 _id, address _target, bytes4 _selector, bytes ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) IsValidOperatorCall(_id *big.Int, _target common.Address, _selector [4]byte, arg3 []byte) (bool, error) {
	return _AvsOperatorManager.Contract.IsValidOperatorCall(&_AvsOperatorManager.CallOpts, _id, _target, _selector, arg3)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_AvsOperatorManager *AvsOperatorManagerCaller) NextAvsOperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "nextAvsOperatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_AvsOperatorManager *AvsOperatorManagerSession) NextAvsOperatorId() (*big.Int, error) {
	return _AvsOperatorManager.Contract.NextAvsOperatorId(&_AvsOperatorManager.CallOpts)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) NextAvsOperatorId() (*big.Int, error) {
	return _AvsOperatorManager.Contract.NextAvsOperatorId(&_AvsOperatorManager.CallOpts)
}

// OperatorDetails is a free data retrieval call binding the contract method 0x38a14d6d.
//
// Solidity: function operatorDetails(uint256 _id) view returns((address,address,uint32))
func (_AvsOperatorManager *AvsOperatorManagerCaller) OperatorDetails(opts *bind.CallOpts, _id *big.Int) (IDelegationManagerOperatorDetails, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "operatorDetails", _id)

	if err != nil {
		return *new(IDelegationManagerOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationManagerOperatorDetails)).(*IDelegationManagerOperatorDetails)

	return out0, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0x38a14d6d.
//
// Solidity: function operatorDetails(uint256 _id) view returns((address,address,uint32))
func (_AvsOperatorManager *AvsOperatorManagerSession) OperatorDetails(_id *big.Int) (IDelegationManagerOperatorDetails, error) {
	return _AvsOperatorManager.Contract.OperatorDetails(&_AvsOperatorManager.CallOpts, _id)
}

// OperatorDetails is a free data retrieval call binding the contract method 0x38a14d6d.
//
// Solidity: function operatorDetails(uint256 _id) view returns((address,address,uint32))
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) OperatorDetails(_id *big.Int) (IDelegationManagerOperatorDetails, error) {
	return _AvsOperatorManager.Contract.OperatorDetails(&_AvsOperatorManager.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerSession) Owner() (common.Address, error) {
	return _AvsOperatorManager.Contract.Owner(&_AvsOperatorManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) Owner() (common.Address, error) {
	return _AvsOperatorManager.Contract.Owner(&_AvsOperatorManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerSession) Paused() (bool, error) {
	return _AvsOperatorManager.Contract.Paused(&_AvsOperatorManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) Paused() (bool, error) {
	return _AvsOperatorManager.Contract.Paused(&_AvsOperatorManager.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerSession) Pausers(arg0 common.Address) (bool, error) {
	return _AvsOperatorManager.Contract.Pausers(&_AvsOperatorManager.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _AvsOperatorManager.Contract.Pausers(&_AvsOperatorManager.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AvsOperatorManager *AvsOperatorManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AvsOperatorManager *AvsOperatorManagerSession) ProxiableUUID() ([32]byte, error) {
	return _AvsOperatorManager.Contract.ProxiableUUID(&_AvsOperatorManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AvsOperatorManager.Contract.ProxiableUUID(&_AvsOperatorManager.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCaller) UpgradableBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsOperatorManager.contract.Call(opts, &out, "upgradableBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerSession) UpgradableBeacon() (common.Address, error) {
	return _AvsOperatorManager.Contract.UpgradableBeacon(&_AvsOperatorManager.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_AvsOperatorManager *AvsOperatorManagerCallerSession) UpgradableBeacon() (common.Address, error) {
	return _AvsOperatorManager.Contract.UpgradableBeacon(&_AvsOperatorManager.CallOpts)
}

// AdminForwardCall is a paid mutator transaction binding the contract method 0x907382ac.
//
// Solidity: function adminForwardCall(uint256 _id, address _target, bytes4 _selector, bytes _args) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) AdminForwardCall(opts *bind.TransactOpts, _id *big.Int, _target common.Address, _selector [4]byte, _args []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "adminForwardCall", _id, _target, _selector, _args)
}

// AdminForwardCall is a paid mutator transaction binding the contract method 0x907382ac.
//
// Solidity: function adminForwardCall(uint256 _id, address _target, bytes4 _selector, bytes _args) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) AdminForwardCall(_id *big.Int, _target common.Address, _selector [4]byte, _args []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.AdminForwardCall(&_AvsOperatorManager.TransactOpts, _id, _target, _selector, _args)
}

// AdminForwardCall is a paid mutator transaction binding the contract method 0x907382ac.
//
// Solidity: function adminForwardCall(uint256 _id, address _target, bytes4 _selector, bytes _args) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) AdminForwardCall(_id *big.Int, _target common.Address, _selector [4]byte, _args []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.AdminForwardCall(&_AvsOperatorManager.TransactOpts, _id, _target, _selector, _args)
}

// ForwardOperatorCall is a paid mutator transaction binding the contract method 0x26099a40.
//
// Solidity: function forwardOperatorCall(uint256 _id, address _target, bytes _input) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) ForwardOperatorCall(opts *bind.TransactOpts, _id *big.Int, _target common.Address, _input []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "forwardOperatorCall", _id, _target, _input)
}

// ForwardOperatorCall is a paid mutator transaction binding the contract method 0x26099a40.
//
// Solidity: function forwardOperatorCall(uint256 _id, address _target, bytes _input) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) ForwardOperatorCall(_id *big.Int, _target common.Address, _input []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.ForwardOperatorCall(&_AvsOperatorManager.TransactOpts, _id, _target, _input)
}

// ForwardOperatorCall is a paid mutator transaction binding the contract method 0x26099a40.
//
// Solidity: function forwardOperatorCall(uint256 _id, address _target, bytes _input) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) ForwardOperatorCall(_id *big.Int, _target common.Address, _input []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.ForwardOperatorCall(&_AvsOperatorManager.TransactOpts, _id, _target, _input)
}

// ForwardOperatorCall0 is a paid mutator transaction binding the contract method 0xc2935fa6.
//
// Solidity: function forwardOperatorCall(uint256 _id, address _target, bytes4 _selector, bytes _args) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) ForwardOperatorCall0(opts *bind.TransactOpts, _id *big.Int, _target common.Address, _selector [4]byte, _args []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "forwardOperatorCall0", _id, _target, _selector, _args)
}

// ForwardOperatorCall0 is a paid mutator transaction binding the contract method 0xc2935fa6.
//
// Solidity: function forwardOperatorCall(uint256 _id, address _target, bytes4 _selector, bytes _args) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) ForwardOperatorCall0(_id *big.Int, _target common.Address, _selector [4]byte, _args []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.ForwardOperatorCall0(&_AvsOperatorManager.TransactOpts, _id, _target, _selector, _args)
}

// ForwardOperatorCall0 is a paid mutator transaction binding the contract method 0xc2935fa6.
//
// Solidity: function forwardOperatorCall(uint256 _id, address _target, bytes4 _selector, bytes _args) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) ForwardOperatorCall0(_id *big.Int, _target common.Address, _selector [4]byte, _args []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.ForwardOperatorCall0(&_AvsOperatorManager.TransactOpts, _id, _target, _selector, _args)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delegationManager, address _avsDirectory, address _etherFiAvsOperatorImpl) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) Initialize(opts *bind.TransactOpts, _delegationManager common.Address, _avsDirectory common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "initialize", _delegationManager, _avsDirectory, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delegationManager, address _avsDirectory, address _etherFiAvsOperatorImpl) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) Initialize(_delegationManager common.Address, _avsDirectory common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.Initialize(&_AvsOperatorManager.TransactOpts, _delegationManager, _avsDirectory, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delegationManager, address _avsDirectory, address _etherFiAvsOperatorImpl) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) Initialize(_delegationManager common.Address, _avsDirectory common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.Initialize(&_AvsOperatorManager.TransactOpts, _delegationManager, _avsDirectory, _etherFiAvsOperatorImpl)
}

// InitializeAvsDirectory is a paid mutator transaction binding the contract method 0xafe833f7.
//
// Solidity: function initializeAvsDirectory(address _avsDirectory) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) InitializeAvsDirectory(opts *bind.TransactOpts, _avsDirectory common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "initializeAvsDirectory", _avsDirectory)
}

// InitializeAvsDirectory is a paid mutator transaction binding the contract method 0xafe833f7.
//
// Solidity: function initializeAvsDirectory(address _avsDirectory) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) InitializeAvsDirectory(_avsDirectory common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.InitializeAvsDirectory(&_AvsOperatorManager.TransactOpts, _avsDirectory)
}

// InitializeAvsDirectory is a paid mutator transaction binding the contract method 0xafe833f7.
//
// Solidity: function initializeAvsDirectory(address _avsDirectory) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) InitializeAvsDirectory(_avsDirectory common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.InitializeAvsDirectory(&_AvsOperatorManager.TransactOpts, _avsDirectory)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_AvsOperatorManager *AvsOperatorManagerTransactor) InstantiateEtherFiAvsOperator(opts *bind.TransactOpts, _nums *big.Int) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "instantiateEtherFiAvsOperator", _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_AvsOperatorManager *AvsOperatorManagerSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.InstantiateEtherFiAvsOperator(&_AvsOperatorManager.TransactOpts, _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.InstantiateEtherFiAvsOperator(&_AvsOperatorManager.TransactOpts, _nums)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, _id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "modifyOperatorDetails", _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.ModifyOperatorDetails(&_AvsOperatorManager.TransactOpts, _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.ModifyOperatorDetails(&_AvsOperatorManager.TransactOpts, _id, _newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, _id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "registerAsOperator", _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.RegisterAsOperator(&_AvsOperatorManager.TransactOpts, _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.RegisterAsOperator(&_AvsOperatorManager.TransactOpts, _id, _detail, _metaDataURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.RenounceOwnership(&_AvsOperatorManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.RenounceOwnership(&_AvsOperatorManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.TransferOwnership(&_AvsOperatorManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.TransferOwnership(&_AvsOperatorManager.TransactOpts, newOwner)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateAdmin(&_AvsOperatorManager.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateAdmin(&_AvsOperatorManager.TransactOpts, _address, _isAdmin)
}

// UpdateAllowedOperatorCalls is a paid mutator transaction binding the contract method 0x8ef7f1ac.
//
// Solidity: function updateAllowedOperatorCalls(uint256 _operatorId, address _target, bytes4 _selector, bool _allowed) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpdateAllowedOperatorCalls(opts *bind.TransactOpts, _operatorId *big.Int, _target common.Address, _selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "updateAllowedOperatorCalls", _operatorId, _target, _selector, _allowed)
}

// UpdateAllowedOperatorCalls is a paid mutator transaction binding the contract method 0x8ef7f1ac.
//
// Solidity: function updateAllowedOperatorCalls(uint256 _operatorId, address _target, bytes4 _selector, bool _allowed) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpdateAllowedOperatorCalls(_operatorId *big.Int, _target common.Address, _selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateAllowedOperatorCalls(&_AvsOperatorManager.TransactOpts, _operatorId, _target, _selector, _allowed)
}

// UpdateAllowedOperatorCalls is a paid mutator transaction binding the contract method 0x8ef7f1ac.
//
// Solidity: function updateAllowedOperatorCalls(uint256 _operatorId, address _target, bytes4 _selector, bool _allowed) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpdateAllowedOperatorCalls(_operatorId *big.Int, _target common.Address, _selector [4]byte, _allowed bool) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateAllowedOperatorCalls(&_AvsOperatorManager.TransactOpts, _operatorId, _target, _selector, _allowed)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpdateAvsNodeRunner(opts *bind.TransactOpts, _id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "updateAvsNodeRunner", _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateAvsNodeRunner(&_AvsOperatorManager.TransactOpts, _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateAvsNodeRunner(&_AvsOperatorManager.TransactOpts, _id, _avsNodeRunner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpdateEcdsaSigner(opts *bind.TransactOpts, _id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "updateEcdsaSigner", _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateEcdsaSigner(&_AvsOperatorManager.TransactOpts, _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateEcdsaSigner(&_AvsOperatorManager.TransactOpts, _id, _ecdsaSigner)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, _id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "updateOperatorMetadataURI", _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateOperatorMetadataURI(&_AvsOperatorManager.TransactOpts, _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpdateOperatorMetadataURI(&_AvsOperatorManager.TransactOpts, _id, _metadataURI)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpgradeEtherFiAvsOperator(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "upgradeEtherFiAvsOperator", _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpgradeEtherFiAvsOperator(&_AvsOperatorManager.TransactOpts, _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpgradeEtherFiAvsOperator(&_AvsOperatorManager.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpgradeTo(&_AvsOperatorManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpgradeTo(&_AvsOperatorManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AvsOperatorManager *AvsOperatorManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpgradeToAndCall(&_AvsOperatorManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AvsOperatorManager *AvsOperatorManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AvsOperatorManager.Contract.UpgradeToAndCall(&_AvsOperatorManager.TransactOpts, newImplementation, data)
}

// AvsOperatorManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AvsOperatorManager contract.
type AvsOperatorManagerAdminChangedIterator struct {
	Event *AvsOperatorManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerAdminChanged)
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
		it.Event = new(AvsOperatorManagerAdminChanged)
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
func (it *AvsOperatorManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerAdminChanged represents a AdminChanged event raised by the AvsOperatorManager contract.
type AvsOperatorManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AvsOperatorManagerAdminChangedIterator, error) {

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerAdminChangedIterator{contract: _AvsOperatorManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerAdminChanged)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseAdminChanged(log types.Log) (*AvsOperatorManagerAdminChanged, error) {
	event := new(AvsOperatorManagerAdminChanged)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerAdminUpdatedIterator is returned from FilterAdminUpdated and is used to iterate over the raw logs and unpacked data for AdminUpdated events raised by the AvsOperatorManager contract.
type AvsOperatorManagerAdminUpdatedIterator struct {
	Event *AvsOperatorManagerAdminUpdated // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerAdminUpdated)
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
		it.Event = new(AvsOperatorManagerAdminUpdated)
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
func (it *AvsOperatorManagerAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerAdminUpdated represents a AdminUpdated event raised by the AvsOperatorManager contract.
type AvsOperatorManagerAdminUpdated struct {
	Admin   common.Address
	IsAdmin bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminUpdated is a free log retrieval operation binding the contract event 0x235bc17e7930760029e9f4d860a2a8089976de5b381cf8380fc11c1d88a11133.
//
// Solidity: event AdminUpdated(address indexed admin, bool isAdmin)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterAdminUpdated(opts *bind.FilterOpts, admin []common.Address) (*AvsOperatorManagerAdminUpdatedIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "AdminUpdated", adminRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerAdminUpdatedIterator{contract: _AvsOperatorManager.contract, event: "AdminUpdated", logs: logs, sub: sub}, nil
}

// WatchAdminUpdated is a free log subscription operation binding the contract event 0x235bc17e7930760029e9f4d860a2a8089976de5b381cf8380fc11c1d88a11133.
//
// Solidity: event AdminUpdated(address indexed admin, bool isAdmin)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchAdminUpdated(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerAdminUpdated, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "AdminUpdated", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerAdminUpdated)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "AdminUpdated", log); err != nil {
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

// ParseAdminUpdated is a log parse operation binding the contract event 0x235bc17e7930760029e9f4d860a2a8089976de5b381cf8380fc11c1d88a11133.
//
// Solidity: event AdminUpdated(address indexed admin, bool isAdmin)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseAdminUpdated(log types.Log) (*AvsOperatorManagerAdminUpdated, error) {
	event := new(AvsOperatorManagerAdminUpdated)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "AdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerAllowedOperatorCallsUpdatedIterator is returned from FilterAllowedOperatorCallsUpdated and is used to iterate over the raw logs and unpacked data for AllowedOperatorCallsUpdated events raised by the AvsOperatorManager contract.
type AvsOperatorManagerAllowedOperatorCallsUpdatedIterator struct {
	Event *AvsOperatorManagerAllowedOperatorCallsUpdated // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerAllowedOperatorCallsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerAllowedOperatorCallsUpdated)
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
		it.Event = new(AvsOperatorManagerAllowedOperatorCallsUpdated)
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
func (it *AvsOperatorManagerAllowedOperatorCallsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerAllowedOperatorCallsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerAllowedOperatorCallsUpdated represents a AllowedOperatorCallsUpdated event raised by the AvsOperatorManager contract.
type AvsOperatorManagerAllowedOperatorCallsUpdated struct {
	Id       *big.Int
	Target   common.Address
	Selector [4]byte
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedOperatorCallsUpdated is a free log retrieval operation binding the contract event 0xe13ff0fb7503ce05251b2bb396e6213850f08d85e6dfa625a4d3a203a0059e4d.
//
// Solidity: event AllowedOperatorCallsUpdated(uint256 indexed id, address indexed target, bytes4 indexed selector, bool allowed)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterAllowedOperatorCallsUpdated(opts *bind.FilterOpts, id []*big.Int, target []common.Address, selector [][4]byte) (*AvsOperatorManagerAllowedOperatorCallsUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "AllowedOperatorCallsUpdated", idRule, targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerAllowedOperatorCallsUpdatedIterator{contract: _AvsOperatorManager.contract, event: "AllowedOperatorCallsUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedOperatorCallsUpdated is a free log subscription operation binding the contract event 0xe13ff0fb7503ce05251b2bb396e6213850f08d85e6dfa625a4d3a203a0059e4d.
//
// Solidity: event AllowedOperatorCallsUpdated(uint256 indexed id, address indexed target, bytes4 indexed selector, bool allowed)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchAllowedOperatorCallsUpdated(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerAllowedOperatorCallsUpdated, id []*big.Int, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "AllowedOperatorCallsUpdated", idRule, targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerAllowedOperatorCallsUpdated)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "AllowedOperatorCallsUpdated", log); err != nil {
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

// ParseAllowedOperatorCallsUpdated is a log parse operation binding the contract event 0xe13ff0fb7503ce05251b2bb396e6213850f08d85e6dfa625a4d3a203a0059e4d.
//
// Solidity: event AllowedOperatorCallsUpdated(uint256 indexed id, address indexed target, bytes4 indexed selector, bool allowed)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseAllowedOperatorCallsUpdated(log types.Log) (*AvsOperatorManagerAllowedOperatorCallsUpdated, error) {
	event := new(AvsOperatorManagerAllowedOperatorCallsUpdated)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "AllowedOperatorCallsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AvsOperatorManager contract.
type AvsOperatorManagerBeaconUpgradedIterator struct {
	Event *AvsOperatorManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerBeaconUpgraded)
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
		it.Event = new(AvsOperatorManagerBeaconUpgraded)
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
func (it *AvsOperatorManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerBeaconUpgraded represents a BeaconUpgraded event raised by the AvsOperatorManager contract.
type AvsOperatorManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AvsOperatorManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerBeaconUpgradedIterator{contract: _AvsOperatorManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerBeaconUpgraded)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseBeaconUpgraded(log types.Log) (*AvsOperatorManagerBeaconUpgraded, error) {
	event := new(AvsOperatorManagerBeaconUpgraded)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerCreatedEtherFiAvsOperatorIterator is returned from FilterCreatedEtherFiAvsOperator and is used to iterate over the raw logs and unpacked data for CreatedEtherFiAvsOperator events raised by the AvsOperatorManager contract.
type AvsOperatorManagerCreatedEtherFiAvsOperatorIterator struct {
	Event *AvsOperatorManagerCreatedEtherFiAvsOperator // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerCreatedEtherFiAvsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerCreatedEtherFiAvsOperator)
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
		it.Event = new(AvsOperatorManagerCreatedEtherFiAvsOperator)
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
func (it *AvsOperatorManagerCreatedEtherFiAvsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerCreatedEtherFiAvsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerCreatedEtherFiAvsOperator represents a CreatedEtherFiAvsOperator event raised by the AvsOperatorManager contract.
type AvsOperatorManagerCreatedEtherFiAvsOperator struct {
	Id                 *big.Int
	EtherFiAvsOperator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCreatedEtherFiAvsOperator is a free log retrieval operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterCreatedEtherFiAvsOperator(opts *bind.FilterOpts, id []*big.Int) (*AvsOperatorManagerCreatedEtherFiAvsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerCreatedEtherFiAvsOperatorIterator{contract: _AvsOperatorManager.contract, event: "CreatedEtherFiAvsOperator", logs: logs, sub: sub}, nil
}

// WatchCreatedEtherFiAvsOperator is a free log subscription operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchCreatedEtherFiAvsOperator(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerCreatedEtherFiAvsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerCreatedEtherFiAvsOperator)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
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

// ParseCreatedEtherFiAvsOperator is a log parse operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseCreatedEtherFiAvsOperator(log types.Log) (*AvsOperatorManagerCreatedEtherFiAvsOperator, error) {
	event := new(AvsOperatorManagerCreatedEtherFiAvsOperator)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerForwardedOperatorCallIterator is returned from FilterForwardedOperatorCall and is used to iterate over the raw logs and unpacked data for ForwardedOperatorCall events raised by the AvsOperatorManager contract.
type AvsOperatorManagerForwardedOperatorCallIterator struct {
	Event *AvsOperatorManagerForwardedOperatorCall // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerForwardedOperatorCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerForwardedOperatorCall)
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
		it.Event = new(AvsOperatorManagerForwardedOperatorCall)
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
func (it *AvsOperatorManagerForwardedOperatorCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerForwardedOperatorCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerForwardedOperatorCall represents a ForwardedOperatorCall event raised by the AvsOperatorManager contract.
type AvsOperatorManagerForwardedOperatorCall struct {
	Id       *big.Int
	Target   common.Address
	Selector [4]byte
	Data     []byte
	Sender   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterForwardedOperatorCall is a free log retrieval operation binding the contract event 0x3a58cd79827a08a20cd91c8a256b4769891ee1a36774089671fd5f0db5a3a6f2.
//
// Solidity: event ForwardedOperatorCall(uint256 indexed id, address indexed target, bytes4 indexed selector, bytes data, address sender)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterForwardedOperatorCall(opts *bind.FilterOpts, id []*big.Int, target []common.Address, selector [][4]byte) (*AvsOperatorManagerForwardedOperatorCallIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "ForwardedOperatorCall", idRule, targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerForwardedOperatorCallIterator{contract: _AvsOperatorManager.contract, event: "ForwardedOperatorCall", logs: logs, sub: sub}, nil
}

// WatchForwardedOperatorCall is a free log subscription operation binding the contract event 0x3a58cd79827a08a20cd91c8a256b4769891ee1a36774089671fd5f0db5a3a6f2.
//
// Solidity: event ForwardedOperatorCall(uint256 indexed id, address indexed target, bytes4 indexed selector, bytes data, address sender)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchForwardedOperatorCall(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerForwardedOperatorCall, id []*big.Int, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "ForwardedOperatorCall", idRule, targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerForwardedOperatorCall)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "ForwardedOperatorCall", log); err != nil {
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

// ParseForwardedOperatorCall is a log parse operation binding the contract event 0x3a58cd79827a08a20cd91c8a256b4769891ee1a36774089671fd5f0db5a3a6f2.
//
// Solidity: event ForwardedOperatorCall(uint256 indexed id, address indexed target, bytes4 indexed selector, bytes data, address sender)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseForwardedOperatorCall(log types.Log) (*AvsOperatorManagerForwardedOperatorCall, error) {
	event := new(AvsOperatorManagerForwardedOperatorCall)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "ForwardedOperatorCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AvsOperatorManager contract.
type AvsOperatorManagerInitializedIterator struct {
	Event *AvsOperatorManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerInitialized)
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
		it.Event = new(AvsOperatorManagerInitialized)
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
func (it *AvsOperatorManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerInitialized represents a Initialized event raised by the AvsOperatorManager contract.
type AvsOperatorManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AvsOperatorManagerInitializedIterator, error) {

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerInitializedIterator{contract: _AvsOperatorManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerInitialized)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseInitialized(log types.Log) (*AvsOperatorManagerInitialized, error) {
	event := new(AvsOperatorManagerInitialized)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerModifiedOperatorDetailsIterator is returned from FilterModifiedOperatorDetails and is used to iterate over the raw logs and unpacked data for ModifiedOperatorDetails events raised by the AvsOperatorManager contract.
type AvsOperatorManagerModifiedOperatorDetailsIterator struct {
	Event *AvsOperatorManagerModifiedOperatorDetails // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerModifiedOperatorDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerModifiedOperatorDetails)
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
		it.Event = new(AvsOperatorManagerModifiedOperatorDetails)
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
func (it *AvsOperatorManagerModifiedOperatorDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerModifiedOperatorDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerModifiedOperatorDetails represents a ModifiedOperatorDetails event raised by the AvsOperatorManager contract.
type AvsOperatorManagerModifiedOperatorDetails struct {
	Id                 *big.Int
	NewOperatorDetails IDelegationManagerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterModifiedOperatorDetails is a free log retrieval operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterModifiedOperatorDetails(opts *bind.FilterOpts, id []*big.Int) (*AvsOperatorManagerModifiedOperatorDetailsIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerModifiedOperatorDetailsIterator{contract: _AvsOperatorManager.contract, event: "ModifiedOperatorDetails", logs: logs, sub: sub}, nil
}

// WatchModifiedOperatorDetails is a free log subscription operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchModifiedOperatorDetails(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerModifiedOperatorDetails, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerModifiedOperatorDetails)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
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

// ParseModifiedOperatorDetails is a log parse operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseModifiedOperatorDetails(log types.Log) (*AvsOperatorManagerModifiedOperatorDetails, error) {
	event := new(AvsOperatorManagerModifiedOperatorDetails)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AvsOperatorManager contract.
type AvsOperatorManagerOwnershipTransferredIterator struct {
	Event *AvsOperatorManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerOwnershipTransferred)
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
		it.Event = new(AvsOperatorManagerOwnershipTransferred)
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
func (it *AvsOperatorManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AvsOperatorManager contract.
type AvsOperatorManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AvsOperatorManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerOwnershipTransferredIterator{contract: _AvsOperatorManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerOwnershipTransferred)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AvsOperatorManagerOwnershipTransferred, error) {
	event := new(AvsOperatorManagerOwnershipTransferred)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AvsOperatorManager contract.
type AvsOperatorManagerPausedIterator struct {
	Event *AvsOperatorManagerPaused // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerPaused)
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
		it.Event = new(AvsOperatorManagerPaused)
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
func (it *AvsOperatorManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerPaused represents a Paused event raised by the AvsOperatorManager contract.
type AvsOperatorManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*AvsOperatorManagerPausedIterator, error) {

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerPausedIterator{contract: _AvsOperatorManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerPaused) (event.Subscription, error) {

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerPaused)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParsePaused(log types.Log) (*AvsOperatorManagerPaused, error) {
	event := new(AvsOperatorManagerPaused)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerRegisteredAsOperatorIterator is returned from FilterRegisteredAsOperator and is used to iterate over the raw logs and unpacked data for RegisteredAsOperator events raised by the AvsOperatorManager contract.
type AvsOperatorManagerRegisteredAsOperatorIterator struct {
	Event *AvsOperatorManagerRegisteredAsOperator // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerRegisteredAsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerRegisteredAsOperator)
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
		it.Event = new(AvsOperatorManagerRegisteredAsOperator)
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
func (it *AvsOperatorManagerRegisteredAsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerRegisteredAsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerRegisteredAsOperator represents a RegisteredAsOperator event raised by the AvsOperatorManager contract.
type AvsOperatorManagerRegisteredAsOperator struct {
	Id     *big.Int
	Detail IDelegationManagerOperatorDetails
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisteredAsOperator is a free log retrieval operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterRegisteredAsOperator(opts *bind.FilterOpts, id []*big.Int) (*AvsOperatorManagerRegisteredAsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerRegisteredAsOperatorIterator{contract: _AvsOperatorManager.contract, event: "RegisteredAsOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredAsOperator is a free log subscription operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchRegisteredAsOperator(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerRegisteredAsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerRegisteredAsOperator)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
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

// ParseRegisteredAsOperator is a log parse operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseRegisteredAsOperator(log types.Log) (*AvsOperatorManagerRegisteredAsOperator, error) {
	event := new(AvsOperatorManagerRegisteredAsOperator)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AvsOperatorManager contract.
type AvsOperatorManagerUnpausedIterator struct {
	Event *AvsOperatorManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerUnpaused)
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
		it.Event = new(AvsOperatorManagerUnpaused)
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
func (it *AvsOperatorManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerUnpaused represents a Unpaused event raised by the AvsOperatorManager contract.
type AvsOperatorManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AvsOperatorManagerUnpausedIterator, error) {

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerUnpausedIterator{contract: _AvsOperatorManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerUnpaused)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseUnpaused(log types.Log) (*AvsOperatorManagerUnpaused, error) {
	event := new(AvsOperatorManagerUnpaused)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerUpdatedAvsNodeRunnerIterator is returned from FilterUpdatedAvsNodeRunner and is used to iterate over the raw logs and unpacked data for UpdatedAvsNodeRunner events raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpdatedAvsNodeRunnerIterator struct {
	Event *AvsOperatorManagerUpdatedAvsNodeRunner // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerUpdatedAvsNodeRunnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerUpdatedAvsNodeRunner)
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
		it.Event = new(AvsOperatorManagerUpdatedAvsNodeRunner)
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
func (it *AvsOperatorManagerUpdatedAvsNodeRunnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerUpdatedAvsNodeRunnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerUpdatedAvsNodeRunner represents a UpdatedAvsNodeRunner event raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpdatedAvsNodeRunner struct {
	Id            *big.Int
	AvsNodeRunner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvsNodeRunner is a free log retrieval operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterUpdatedAvsNodeRunner(opts *bind.FilterOpts, id []*big.Int) (*AvsOperatorManagerUpdatedAvsNodeRunnerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerUpdatedAvsNodeRunnerIterator{contract: _AvsOperatorManager.contract, event: "UpdatedAvsNodeRunner", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvsNodeRunner is a free log subscription operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchUpdatedAvsNodeRunner(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerUpdatedAvsNodeRunner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerUpdatedAvsNodeRunner)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
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

// ParseUpdatedAvsNodeRunner is a log parse operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseUpdatedAvsNodeRunner(log types.Log) (*AvsOperatorManagerUpdatedAvsNodeRunner, error) {
	event := new(AvsOperatorManagerUpdatedAvsNodeRunner)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerUpdatedEcdsaSignerIterator is returned from FilterUpdatedEcdsaSigner and is used to iterate over the raw logs and unpacked data for UpdatedEcdsaSigner events raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpdatedEcdsaSignerIterator struct {
	Event *AvsOperatorManagerUpdatedEcdsaSigner // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerUpdatedEcdsaSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerUpdatedEcdsaSigner)
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
		it.Event = new(AvsOperatorManagerUpdatedEcdsaSigner)
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
func (it *AvsOperatorManagerUpdatedEcdsaSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerUpdatedEcdsaSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerUpdatedEcdsaSigner represents a UpdatedEcdsaSigner event raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpdatedEcdsaSigner struct {
	Id          *big.Int
	EcdsaSigner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEcdsaSigner is a free log retrieval operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterUpdatedEcdsaSigner(opts *bind.FilterOpts, id []*big.Int) (*AvsOperatorManagerUpdatedEcdsaSignerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerUpdatedEcdsaSignerIterator{contract: _AvsOperatorManager.contract, event: "UpdatedEcdsaSigner", logs: logs, sub: sub}, nil
}

// WatchUpdatedEcdsaSigner is a free log subscription operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchUpdatedEcdsaSigner(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerUpdatedEcdsaSigner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerUpdatedEcdsaSigner)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
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

// ParseUpdatedEcdsaSigner is a log parse operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseUpdatedEcdsaSigner(log types.Log) (*AvsOperatorManagerUpdatedEcdsaSigner, error) {
	event := new(AvsOperatorManagerUpdatedEcdsaSigner)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerUpdatedOperatorMetadataURIIterator is returned from FilterUpdatedOperatorMetadataURI and is used to iterate over the raw logs and unpacked data for UpdatedOperatorMetadataURI events raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpdatedOperatorMetadataURIIterator struct {
	Event *AvsOperatorManagerUpdatedOperatorMetadataURI // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerUpdatedOperatorMetadataURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerUpdatedOperatorMetadataURI)
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
		it.Event = new(AvsOperatorManagerUpdatedOperatorMetadataURI)
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
func (it *AvsOperatorManagerUpdatedOperatorMetadataURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerUpdatedOperatorMetadataURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerUpdatedOperatorMetadataURI represents a UpdatedOperatorMetadataURI event raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpdatedOperatorMetadataURI struct {
	Id          *big.Int
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOperatorMetadataURI is a free log retrieval operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterUpdatedOperatorMetadataURI(opts *bind.FilterOpts, id []*big.Int) (*AvsOperatorManagerUpdatedOperatorMetadataURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerUpdatedOperatorMetadataURIIterator{contract: _AvsOperatorManager.contract, event: "UpdatedOperatorMetadataURI", logs: logs, sub: sub}, nil
}

// WatchUpdatedOperatorMetadataURI is a free log subscription operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchUpdatedOperatorMetadataURI(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerUpdatedOperatorMetadataURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerUpdatedOperatorMetadataURI)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
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

// ParseUpdatedOperatorMetadataURI is a log parse operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseUpdatedOperatorMetadataURI(log types.Log) (*AvsOperatorManagerUpdatedOperatorMetadataURI, error) {
	event := new(AvsOperatorManagerUpdatedOperatorMetadataURI)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsOperatorManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpgradedIterator struct {
	Event *AvsOperatorManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *AvsOperatorManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsOperatorManagerUpgraded)
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
		it.Event = new(AvsOperatorManagerUpgraded)
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
func (it *AvsOperatorManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsOperatorManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsOperatorManagerUpgraded represents a Upgraded event raised by the AvsOperatorManager contract.
type AvsOperatorManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AvsOperatorManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AvsOperatorManagerUpgradedIterator{contract: _AvsOperatorManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AvsOperatorManager *AvsOperatorManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AvsOperatorManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AvsOperatorManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsOperatorManagerUpgraded)
				if err := _AvsOperatorManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AvsOperatorManager *AvsOperatorManagerFilterer) ParseUpgraded(log types.Log) (*AvsOperatorManagerUpgraded, error) {
	event := new(AvsOperatorManagerUpgraded)
	if err := _AvsOperatorManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
