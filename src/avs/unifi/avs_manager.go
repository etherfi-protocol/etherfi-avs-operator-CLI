// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unifi

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

// OperatorCommitment is an auto generated low-level Go binding around an user-defined struct.
type OperatorCommitment struct {
	DelegateKey   []byte
	ChainIDBitMap *big.Int
}

// OperatorDataExtended is an auto generated low-level Go binding around an user-defined struct.
type OperatorDataExtended struct {
	Commitment                   OperatorCommitment
	PendingCommitment            OperatorCommitment
	ValidatorCount               *big.Int
	StartDeregisterOperatorBlock *big.Int
	CommitmentValidAfter         *big.Int
	IsRegistered                 bool
}

// ValidatorDataExtended is an auto generated low-level Go binding around an user-defined struct.
type ValidatorDataExtended struct {
	Operator       common.Address
	EigenPod       common.Address
	ValidatorIndex uint64
	Status         uint8
	DelegateKey    []byte
	ChainIDBitMap  *big.Int
	BackedByStake  bool
	Registered     bool
}

// AvsManagerMetaData contains all meta data concerning the AvsManager contract.
var AvsManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"eigenPodManagerAddress\",\"type\":\"address\"},{\"internalType\":\"contractIDelegationManager\",\"name\":\"eigenDelegationManagerAddress\",\"type\":\"address\"},{\"internalType\":\"contractIAVSDirectory\",\"name\":\"avsDirectoryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"AccessManagedInvalidAuthority\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"AccessManagedRequiredDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AccessManagedUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentChangeNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelegateKeyNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeregistrationAlreadyStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeregistrationDelayNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeregistrationNotStarted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoEigenPod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDelegatedToOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidatorOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorHasValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RestakingStrategyAllowlistUpdateFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyDeregistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"ChainIDSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldDelay\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newDelay\",\"type\":\"uint64\"}],\"name\":\"DeregistrationDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOperatorCommitment\",\"name\":\"oldCommitment\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOperatorCommitment\",\"name\":\"newCommitment\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"validAfter\",\"type\":\"uint128\"}],\"name\":\"OperatorCommitmentChangeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOperatorCommitment\",\"name\":\"oldCommitment\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOperatorCommitment\",\"name\":\"newCommitment\",\"type\":\"tuple\"}],\"name\":\"OperatorCommitmentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorDeregisterStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOperatorCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"OperatorRegisteredWithCommitment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"RestakingStrategyAllowlistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blsPubKeyHash\",\"type\":\"bytes32\"}],\"name\":\"ValidatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blsPubKeyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorIndex\",\"type\":\"uint256\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AVS_DIRECTORY\",\"outputs\":[{\"internalType\":\"contractIAVSDirectoryExtended\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BEACON_CHAIN_STRATEGY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIGEN_DELEGATION_MANAGER\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIGEN_POD_MANAGER\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bitmap\",\"type\":\"uint256\"}],\"name\":\"bitmapToChainIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"blsPubKeyHashes\",\"type\":\"bytes32[]\"}],\"name\":\"deregisterValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finishDeregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getBitmapIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeregistrationDelay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"internalType\":\"structOperatorCommitment\",\"name\":\"commitment\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"internalType\":\"structOperatorCommitment\",\"name\":\"pendingCommitment\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"validatorCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"startDeregisterOperatorBlock\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"commitmentValidAfter\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structOperatorDataExtended\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"restakedStrategies\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorIndex\",\"type\":\"uint256\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eigenPod\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validatorIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"backedByStake\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"internalType\":\"structValidatorDataExtended\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blsPubKeyHash\",\"type\":\"bytes32\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eigenPod\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validatorIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"backedByStake\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"internalType\":\"structValidatorDataExtended\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"blsPubKeyHashes\",\"type\":\"bytes32[]\"}],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eigenPod\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validatorIndex\",\"type\":\"uint64\"},{\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"backedByStake\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"internalType\":\"structValidatorDataExtended[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManager\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"initialDeregistrationDelay\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isConsumingScheduledOp\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blsPubKeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"isValidatorInChainId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"internalType\":\"structOperatorCommitment\",\"name\":\"initialCommitment\",\"type\":\"tuple\"}],\"name\":\"registerOperatorWithCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"blsPubKeyHashes\",\"type\":\"bytes32[]\"}],\"name\":\"registerValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setAllowlistRestakingStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"setChainID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newDelay\",\"type\":\"uint64\"}],\"name\":\"setDeregistrationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"delegateKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"chainIDBitMap\",\"type\":\"uint256\"}],\"internalType\":\"structOperatorCommitment\",\"name\":\"newCommitment\",\"type\":\"tuple\"}],\"name\":\"setOperatorCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startDeregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateOperatorCommitment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AvsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AvsManagerMetaData.ABI instead.
var AvsManagerABI = AvsManagerMetaData.ABI

// AvsManager is an auto generated Go binding around an Ethereum contract.
type AvsManager struct {
	AvsManagerCaller     // Read-only binding to the contract
	AvsManagerTransactor // Write-only binding to the contract
	AvsManagerFilterer   // Log filterer for contract events
}

// AvsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AvsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AvsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AvsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AvsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AvsManagerSession struct {
	Contract     *AvsManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AvsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AvsManagerCallerSession struct {
	Contract *AvsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AvsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AvsManagerTransactorSession struct {
	Contract     *AvsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AvsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AvsManagerRaw struct {
	Contract *AvsManager // Generic contract binding to access the raw methods on
}

// AvsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AvsManagerCallerRaw struct {
	Contract *AvsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AvsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AvsManagerTransactorRaw struct {
	Contract *AvsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAvsManager creates a new instance of AvsManager, bound to a specific deployed contract.
func NewAvsManager(address common.Address, backend bind.ContractBackend) (*AvsManager, error) {
	contract, err := bindAvsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AvsManager{AvsManagerCaller: AvsManagerCaller{contract: contract}, AvsManagerTransactor: AvsManagerTransactor{contract: contract}, AvsManagerFilterer: AvsManagerFilterer{contract: contract}}, nil
}

// NewAvsManagerCaller creates a new read-only instance of AvsManager, bound to a specific deployed contract.
func NewAvsManagerCaller(address common.Address, caller bind.ContractCaller) (*AvsManagerCaller, error) {
	contract, err := bindAvsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AvsManagerCaller{contract: contract}, nil
}

// NewAvsManagerTransactor creates a new write-only instance of AvsManager, bound to a specific deployed contract.
func NewAvsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AvsManagerTransactor, error) {
	contract, err := bindAvsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AvsManagerTransactor{contract: contract}, nil
}

// NewAvsManagerFilterer creates a new log filterer instance of AvsManager, bound to a specific deployed contract.
func NewAvsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AvsManagerFilterer, error) {
	contract, err := bindAvsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AvsManagerFilterer{contract: contract}, nil
}

// bindAvsManager binds a generic wrapper to an already deployed contract.
func bindAvsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AvsManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvsManager *AvsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvsManager.Contract.AvsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvsManager *AvsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsManager.Contract.AvsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvsManager *AvsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvsManager.Contract.AvsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AvsManager *AvsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AvsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AvsManager *AvsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AvsManager *AvsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AvsManager.Contract.contract.Transact(opts, method, params...)
}

// AVSDIRECTORY is a free data retrieval call binding the contract method 0xad656c56.
//
// Solidity: function AVS_DIRECTORY() view returns(address)
func (_AvsManager *AvsManagerCaller) AVSDIRECTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "AVS_DIRECTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AVSDIRECTORY is a free data retrieval call binding the contract method 0xad656c56.
//
// Solidity: function AVS_DIRECTORY() view returns(address)
func (_AvsManager *AvsManagerSession) AVSDIRECTORY() (common.Address, error) {
	return _AvsManager.Contract.AVSDIRECTORY(&_AvsManager.CallOpts)
}

// AVSDIRECTORY is a free data retrieval call binding the contract method 0xad656c56.
//
// Solidity: function AVS_DIRECTORY() view returns(address)
func (_AvsManager *AvsManagerCallerSession) AVSDIRECTORY() (common.Address, error) {
	return _AvsManager.Contract.AVSDIRECTORY(&_AvsManager.CallOpts)
}

// BEACONCHAINSTRATEGY is a free data retrieval call binding the contract method 0x3f2c6d7f.
//
// Solidity: function BEACON_CHAIN_STRATEGY() view returns(address)
func (_AvsManager *AvsManagerCaller) BEACONCHAINSTRATEGY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "BEACON_CHAIN_STRATEGY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BEACONCHAINSTRATEGY is a free data retrieval call binding the contract method 0x3f2c6d7f.
//
// Solidity: function BEACON_CHAIN_STRATEGY() view returns(address)
func (_AvsManager *AvsManagerSession) BEACONCHAINSTRATEGY() (common.Address, error) {
	return _AvsManager.Contract.BEACONCHAINSTRATEGY(&_AvsManager.CallOpts)
}

// BEACONCHAINSTRATEGY is a free data retrieval call binding the contract method 0x3f2c6d7f.
//
// Solidity: function BEACON_CHAIN_STRATEGY() view returns(address)
func (_AvsManager *AvsManagerCallerSession) BEACONCHAINSTRATEGY() (common.Address, error) {
	return _AvsManager.Contract.BEACONCHAINSTRATEGY(&_AvsManager.CallOpts)
}

// EIGENDELEGATIONMANAGER is a free data retrieval call binding the contract method 0xf2251b0f.
//
// Solidity: function EIGEN_DELEGATION_MANAGER() view returns(address)
func (_AvsManager *AvsManagerCaller) EIGENDELEGATIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "EIGEN_DELEGATION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGENDELEGATIONMANAGER is a free data retrieval call binding the contract method 0xf2251b0f.
//
// Solidity: function EIGEN_DELEGATION_MANAGER() view returns(address)
func (_AvsManager *AvsManagerSession) EIGENDELEGATIONMANAGER() (common.Address, error) {
	return _AvsManager.Contract.EIGENDELEGATIONMANAGER(&_AvsManager.CallOpts)
}

// EIGENDELEGATIONMANAGER is a free data retrieval call binding the contract method 0xf2251b0f.
//
// Solidity: function EIGEN_DELEGATION_MANAGER() view returns(address)
func (_AvsManager *AvsManagerCallerSession) EIGENDELEGATIONMANAGER() (common.Address, error) {
	return _AvsManager.Contract.EIGENDELEGATIONMANAGER(&_AvsManager.CallOpts)
}

// EIGENPODMANAGER is a free data retrieval call binding the contract method 0x0a732358.
//
// Solidity: function EIGEN_POD_MANAGER() view returns(address)
func (_AvsManager *AvsManagerCaller) EIGENPODMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "EIGEN_POD_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EIGENPODMANAGER is a free data retrieval call binding the contract method 0x0a732358.
//
// Solidity: function EIGEN_POD_MANAGER() view returns(address)
func (_AvsManager *AvsManagerSession) EIGENPODMANAGER() (common.Address, error) {
	return _AvsManager.Contract.EIGENPODMANAGER(&_AvsManager.CallOpts)
}

// EIGENPODMANAGER is a free data retrieval call binding the contract method 0x0a732358.
//
// Solidity: function EIGEN_POD_MANAGER() view returns(address)
func (_AvsManager *AvsManagerCallerSession) EIGENPODMANAGER() (common.Address, error) {
	return _AvsManager.Contract.EIGENPODMANAGER(&_AvsManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AvsManager *AvsManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AvsManager *AvsManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AvsManager.Contract.UPGRADEINTERFACEVERSION(&_AvsManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_AvsManager *AvsManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _AvsManager.Contract.UPGRADEINTERFACEVERSION(&_AvsManager.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_AvsManager *AvsManagerCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_AvsManager *AvsManagerSession) Authority() (common.Address, error) {
	return _AvsManager.Contract.Authority(&_AvsManager.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_AvsManager *AvsManagerCallerSession) Authority() (common.Address, error) {
	return _AvsManager.Contract.Authority(&_AvsManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AvsManager *AvsManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AvsManager *AvsManagerSession) AvsDirectory() (common.Address, error) {
	return _AvsManager.Contract.AvsDirectory(&_AvsManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AvsManager *AvsManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _AvsManager.Contract.AvsDirectory(&_AvsManager.CallOpts)
}

// BitmapToChainIDs is a free data retrieval call binding the contract method 0x96c0fa5f.
//
// Solidity: function bitmapToChainIDs(uint256 bitmap) view returns(uint256[])
func (_AvsManager *AvsManagerCaller) BitmapToChainIDs(opts *bind.CallOpts, bitmap *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "bitmapToChainIDs", bitmap)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BitmapToChainIDs is a free data retrieval call binding the contract method 0x96c0fa5f.
//
// Solidity: function bitmapToChainIDs(uint256 bitmap) view returns(uint256[])
func (_AvsManager *AvsManagerSession) BitmapToChainIDs(bitmap *big.Int) ([]*big.Int, error) {
	return _AvsManager.Contract.BitmapToChainIDs(&_AvsManager.CallOpts, bitmap)
}

// BitmapToChainIDs is a free data retrieval call binding the contract method 0x96c0fa5f.
//
// Solidity: function bitmapToChainIDs(uint256 bitmap) view returns(uint256[])
func (_AvsManager *AvsManagerCallerSession) BitmapToChainIDs(bitmap *big.Int) ([]*big.Int, error) {
	return _AvsManager.Contract.BitmapToChainIDs(&_AvsManager.CallOpts, bitmap)
}

// GetBitmapIndex is a free data retrieval call binding the contract method 0xc8fee669.
//
// Solidity: function getBitmapIndex(uint256 chainID) view returns(uint8)
func (_AvsManager *AvsManagerCaller) GetBitmapIndex(opts *bind.CallOpts, chainID *big.Int) (uint8, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getBitmapIndex", chainID)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetBitmapIndex is a free data retrieval call binding the contract method 0xc8fee669.
//
// Solidity: function getBitmapIndex(uint256 chainID) view returns(uint8)
func (_AvsManager *AvsManagerSession) GetBitmapIndex(chainID *big.Int) (uint8, error) {
	return _AvsManager.Contract.GetBitmapIndex(&_AvsManager.CallOpts, chainID)
}

// GetBitmapIndex is a free data retrieval call binding the contract method 0xc8fee669.
//
// Solidity: function getBitmapIndex(uint256 chainID) view returns(uint8)
func (_AvsManager *AvsManagerCallerSession) GetBitmapIndex(chainID *big.Int) (uint8, error) {
	return _AvsManager.Contract.GetBitmapIndex(&_AvsManager.CallOpts, chainID)
}

// GetChainID is a free data retrieval call binding the contract method 0x435ee6b0.
//
// Solidity: function getChainID(uint8 index) view returns(uint256)
func (_AvsManager *AvsManagerCaller) GetChainID(opts *bind.CallOpts, index uint8) (*big.Int, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getChainID", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainID is a free data retrieval call binding the contract method 0x435ee6b0.
//
// Solidity: function getChainID(uint8 index) view returns(uint256)
func (_AvsManager *AvsManagerSession) GetChainID(index uint8) (*big.Int, error) {
	return _AvsManager.Contract.GetChainID(&_AvsManager.CallOpts, index)
}

// GetChainID is a free data retrieval call binding the contract method 0x435ee6b0.
//
// Solidity: function getChainID(uint8 index) view returns(uint256)
func (_AvsManager *AvsManagerCallerSession) GetChainID(index uint8) (*big.Int, error) {
	return _AvsManager.Contract.GetChainID(&_AvsManager.CallOpts, index)
}

// GetDeregistrationDelay is a free data retrieval call binding the contract method 0x8495feab.
//
// Solidity: function getDeregistrationDelay() view returns(uint64)
func (_AvsManager *AvsManagerCaller) GetDeregistrationDelay(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getDeregistrationDelay")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDeregistrationDelay is a free data retrieval call binding the contract method 0x8495feab.
//
// Solidity: function getDeregistrationDelay() view returns(uint64)
func (_AvsManager *AvsManagerSession) GetDeregistrationDelay() (uint64, error) {
	return _AvsManager.Contract.GetDeregistrationDelay(&_AvsManager.CallOpts)
}

// GetDeregistrationDelay is a free data retrieval call binding the contract method 0x8495feab.
//
// Solidity: function getDeregistrationDelay() view returns(uint64)
func (_AvsManager *AvsManagerCallerSession) GetDeregistrationDelay() (uint64, error) {
	return _AvsManager.Contract.GetDeregistrationDelay(&_AvsManager.CallOpts)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns(((bytes,uint256),(bytes,uint256),uint128,uint128,uint128,bool))
func (_AvsManager *AvsManagerCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (OperatorDataExtended, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(OperatorDataExtended), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorDataExtended)).(*OperatorDataExtended)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns(((bytes,uint256),(bytes,uint256),uint128,uint128,uint128,bool))
func (_AvsManager *AvsManagerSession) GetOperator(operator common.Address) (OperatorDataExtended, error) {
	return _AvsManager.Contract.GetOperator(&_AvsManager.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns(((bytes,uint256),(bytes,uint256),uint128,uint128,uint128,bool))
func (_AvsManager *AvsManagerCallerSession) GetOperator(operator common.Address) (OperatorDataExtended, error) {
	return _AvsManager.Contract.GetOperator(&_AvsManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[] restakedStrategies)
func (_AvsManager *AvsManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[] restakedStrategies)
func (_AvsManager *AvsManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _AvsManager.Contract.GetOperatorRestakedStrategies(&_AvsManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[] restakedStrategies)
func (_AvsManager *AvsManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _AvsManager.Contract.GetOperatorRestakedStrategies(&_AvsManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_AvsManager *AvsManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_AvsManager *AvsManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _AvsManager.Contract.GetRestakeableStrategies(&_AvsManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_AvsManager *AvsManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _AvsManager.Contract.GetRestakeableStrategies(&_AvsManager.CallOpts)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 validatorIndex) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool))
func (_AvsManager *AvsManagerCaller) GetValidator(opts *bind.CallOpts, validatorIndex *big.Int) (ValidatorDataExtended, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getValidator", validatorIndex)

	if err != nil {
		return *new(ValidatorDataExtended), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorDataExtended)).(*ValidatorDataExtended)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 validatorIndex) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool))
func (_AvsManager *AvsManagerSession) GetValidator(validatorIndex *big.Int) (ValidatorDataExtended, error) {
	return _AvsManager.Contract.GetValidator(&_AvsManager.CallOpts, validatorIndex)
}

// GetValidator is a free data retrieval call binding the contract method 0xb5d89627.
//
// Solidity: function getValidator(uint256 validatorIndex) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool))
func (_AvsManager *AvsManagerCallerSession) GetValidator(validatorIndex *big.Int) (ValidatorDataExtended, error) {
	return _AvsManager.Contract.GetValidator(&_AvsManager.CallOpts, validatorIndex)
}

// GetValidator0 is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 blsPubKeyHash) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool))
func (_AvsManager *AvsManagerCaller) GetValidator0(opts *bind.CallOpts, blsPubKeyHash [32]byte) (ValidatorDataExtended, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getValidator0", blsPubKeyHash)

	if err != nil {
		return *new(ValidatorDataExtended), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorDataExtended)).(*ValidatorDataExtended)

	return out0, err

}

// GetValidator0 is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 blsPubKeyHash) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool))
func (_AvsManager *AvsManagerSession) GetValidator0(blsPubKeyHash [32]byte) (ValidatorDataExtended, error) {
	return _AvsManager.Contract.GetValidator0(&_AvsManager.CallOpts, blsPubKeyHash)
}

// GetValidator0 is a free data retrieval call binding the contract method 0xd5f20ff6.
//
// Solidity: function getValidator(bytes32 blsPubKeyHash) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool))
func (_AvsManager *AvsManagerCallerSession) GetValidator0(blsPubKeyHash [32]byte) (ValidatorDataExtended, error) {
	return _AvsManager.Contract.GetValidator0(&_AvsManager.CallOpts, blsPubKeyHash)
}

// GetValidators is a free data retrieval call binding the contract method 0x786ffbc2.
//
// Solidity: function getValidators(bytes32[] blsPubKeyHashes) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool)[])
func (_AvsManager *AvsManagerCaller) GetValidators(opts *bind.CallOpts, blsPubKeyHashes [][32]byte) ([]ValidatorDataExtended, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "getValidators", blsPubKeyHashes)

	if err != nil {
		return *new([]ValidatorDataExtended), err
	}

	out0 := *abi.ConvertType(out[0], new([]ValidatorDataExtended)).(*[]ValidatorDataExtended)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0x786ffbc2.
//
// Solidity: function getValidators(bytes32[] blsPubKeyHashes) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool)[])
func (_AvsManager *AvsManagerSession) GetValidators(blsPubKeyHashes [][32]byte) ([]ValidatorDataExtended, error) {
	return _AvsManager.Contract.GetValidators(&_AvsManager.CallOpts, blsPubKeyHashes)
}

// GetValidators is a free data retrieval call binding the contract method 0x786ffbc2.
//
// Solidity: function getValidators(bytes32[] blsPubKeyHashes) view returns((address,address,uint64,uint8,bytes,uint256,bool,bool)[])
func (_AvsManager *AvsManagerCallerSession) GetValidators(blsPubKeyHashes [][32]byte) ([]ValidatorDataExtended, error) {
	return _AvsManager.Contract.GetValidators(&_AvsManager.CallOpts, blsPubKeyHashes)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_AvsManager *AvsManagerCaller) IsConsumingScheduledOp(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "isConsumingScheduledOp")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_AvsManager *AvsManagerSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _AvsManager.Contract.IsConsumingScheduledOp(&_AvsManager.CallOpts)
}

// IsConsumingScheduledOp is a free data retrieval call binding the contract method 0x8fb36037.
//
// Solidity: function isConsumingScheduledOp() view returns(bytes4)
func (_AvsManager *AvsManagerCallerSession) IsConsumingScheduledOp() ([4]byte, error) {
	return _AvsManager.Contract.IsConsumingScheduledOp(&_AvsManager.CallOpts)
}

// IsValidatorInChainId is a free data retrieval call binding the contract method 0x619b9080.
//
// Solidity: function isValidatorInChainId(bytes32 blsPubKeyHash, uint256 chainId) view returns(bool)
func (_AvsManager *AvsManagerCaller) IsValidatorInChainId(opts *bind.CallOpts, blsPubKeyHash [32]byte, chainId *big.Int) (bool, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "isValidatorInChainId", blsPubKeyHash, chainId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorInChainId is a free data retrieval call binding the contract method 0x619b9080.
//
// Solidity: function isValidatorInChainId(bytes32 blsPubKeyHash, uint256 chainId) view returns(bool)
func (_AvsManager *AvsManagerSession) IsValidatorInChainId(blsPubKeyHash [32]byte, chainId *big.Int) (bool, error) {
	return _AvsManager.Contract.IsValidatorInChainId(&_AvsManager.CallOpts, blsPubKeyHash, chainId)
}

// IsValidatorInChainId is a free data retrieval call binding the contract method 0x619b9080.
//
// Solidity: function isValidatorInChainId(bytes32 blsPubKeyHash, uint256 chainId) view returns(bool)
func (_AvsManager *AvsManagerCallerSession) IsValidatorInChainId(blsPubKeyHash [32]byte, chainId *big.Int) (bool, error) {
	return _AvsManager.Contract.IsValidatorInChainId(&_AvsManager.CallOpts, blsPubKeyHash, chainId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AvsManager *AvsManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AvsManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AvsManager *AvsManagerSession) ProxiableUUID() ([32]byte, error) {
	return _AvsManager.Contract.ProxiableUUID(&_AvsManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AvsManager *AvsManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AvsManager.Contract.ProxiableUUID(&_AvsManager.CallOpts)
}

// DeregisterValidators is a paid mutator transaction binding the contract method 0x29d6786d.
//
// Solidity: function deregisterValidators(bytes32[] blsPubKeyHashes) returns()
func (_AvsManager *AvsManagerTransactor) DeregisterValidators(opts *bind.TransactOpts, blsPubKeyHashes [][32]byte) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "deregisterValidators", blsPubKeyHashes)
}

// DeregisterValidators is a paid mutator transaction binding the contract method 0x29d6786d.
//
// Solidity: function deregisterValidators(bytes32[] blsPubKeyHashes) returns()
func (_AvsManager *AvsManagerSession) DeregisterValidators(blsPubKeyHashes [][32]byte) (*types.Transaction, error) {
	return _AvsManager.Contract.DeregisterValidators(&_AvsManager.TransactOpts, blsPubKeyHashes)
}

// DeregisterValidators is a paid mutator transaction binding the contract method 0x29d6786d.
//
// Solidity: function deregisterValidators(bytes32[] blsPubKeyHashes) returns()
func (_AvsManager *AvsManagerTransactorSession) DeregisterValidators(blsPubKeyHashes [][32]byte) (*types.Transaction, error) {
	return _AvsManager.Contract.DeregisterValidators(&_AvsManager.TransactOpts, blsPubKeyHashes)
}

// FinishDeregisterOperator is a paid mutator transaction binding the contract method 0xe3672163.
//
// Solidity: function finishDeregisterOperator() returns()
func (_AvsManager *AvsManagerTransactor) FinishDeregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "finishDeregisterOperator")
}

// FinishDeregisterOperator is a paid mutator transaction binding the contract method 0xe3672163.
//
// Solidity: function finishDeregisterOperator() returns()
func (_AvsManager *AvsManagerSession) FinishDeregisterOperator() (*types.Transaction, error) {
	return _AvsManager.Contract.FinishDeregisterOperator(&_AvsManager.TransactOpts)
}

// FinishDeregisterOperator is a paid mutator transaction binding the contract method 0xe3672163.
//
// Solidity: function finishDeregisterOperator() returns()
func (_AvsManager *AvsManagerTransactorSession) FinishDeregisterOperator() (*types.Transaction, error) {
	return _AvsManager.Contract.FinishDeregisterOperator(&_AvsManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1798de81.
//
// Solidity: function initialize(address accessManager, uint64 initialDeregistrationDelay) returns()
func (_AvsManager *AvsManagerTransactor) Initialize(opts *bind.TransactOpts, accessManager common.Address, initialDeregistrationDelay uint64) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "initialize", accessManager, initialDeregistrationDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x1798de81.
//
// Solidity: function initialize(address accessManager, uint64 initialDeregistrationDelay) returns()
func (_AvsManager *AvsManagerSession) Initialize(accessManager common.Address, initialDeregistrationDelay uint64) (*types.Transaction, error) {
	return _AvsManager.Contract.Initialize(&_AvsManager.TransactOpts, accessManager, initialDeregistrationDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x1798de81.
//
// Solidity: function initialize(address accessManager, uint64 initialDeregistrationDelay) returns()
func (_AvsManager *AvsManagerTransactorSession) Initialize(accessManager common.Address, initialDeregistrationDelay uint64) (*types.Transaction, error) {
	return _AvsManager.Contract.Initialize(&_AvsManager.TransactOpts, accessManager, initialDeregistrationDelay)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8317781d.
//
// Solidity: function registerOperator((bytes,bytes32,uint256) operatorSignature) returns()
func (_AvsManager *AvsManagerTransactor) RegisterOperator(opts *bind.TransactOpts, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "registerOperator", operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8317781d.
//
// Solidity: function registerOperator((bytes,bytes32,uint256) operatorSignature) returns()
func (_AvsManager *AvsManagerSession) RegisterOperator(operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AvsManager.Contract.RegisterOperator(&_AvsManager.TransactOpts, operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x8317781d.
//
// Solidity: function registerOperator((bytes,bytes32,uint256) operatorSignature) returns()
func (_AvsManager *AvsManagerTransactorSession) RegisterOperator(operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AvsManager.Contract.RegisterOperator(&_AvsManager.TransactOpts, operatorSignature)
}

// RegisterOperatorWithCommitment is a paid mutator transaction binding the contract method 0x858b539d.
//
// Solidity: function registerOperatorWithCommitment((bytes,bytes32,uint256) operatorSignature, (bytes,uint256) initialCommitment) returns()
func (_AvsManager *AvsManagerTransactor) RegisterOperatorWithCommitment(opts *bind.TransactOpts, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, initialCommitment OperatorCommitment) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "registerOperatorWithCommitment", operatorSignature, initialCommitment)
}

// RegisterOperatorWithCommitment is a paid mutator transaction binding the contract method 0x858b539d.
//
// Solidity: function registerOperatorWithCommitment((bytes,bytes32,uint256) operatorSignature, (bytes,uint256) initialCommitment) returns()
func (_AvsManager *AvsManagerSession) RegisterOperatorWithCommitment(operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, initialCommitment OperatorCommitment) (*types.Transaction, error) {
	return _AvsManager.Contract.RegisterOperatorWithCommitment(&_AvsManager.TransactOpts, operatorSignature, initialCommitment)
}

// RegisterOperatorWithCommitment is a paid mutator transaction binding the contract method 0x858b539d.
//
// Solidity: function registerOperatorWithCommitment((bytes,bytes32,uint256) operatorSignature, (bytes,uint256) initialCommitment) returns()
func (_AvsManager *AvsManagerTransactorSession) RegisterOperatorWithCommitment(operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, initialCommitment OperatorCommitment) (*types.Transaction, error) {
	return _AvsManager.Contract.RegisterOperatorWithCommitment(&_AvsManager.TransactOpts, operatorSignature, initialCommitment)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xd4d73589.
//
// Solidity: function registerValidators(address podOwner, bytes32[] blsPubKeyHashes) returns()
func (_AvsManager *AvsManagerTransactor) RegisterValidators(opts *bind.TransactOpts, podOwner common.Address, blsPubKeyHashes [][32]byte) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "registerValidators", podOwner, blsPubKeyHashes)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xd4d73589.
//
// Solidity: function registerValidators(address podOwner, bytes32[] blsPubKeyHashes) returns()
func (_AvsManager *AvsManagerSession) RegisterValidators(podOwner common.Address, blsPubKeyHashes [][32]byte) (*types.Transaction, error) {
	return _AvsManager.Contract.RegisterValidators(&_AvsManager.TransactOpts, podOwner, blsPubKeyHashes)
}

// RegisterValidators is a paid mutator transaction binding the contract method 0xd4d73589.
//
// Solidity: function registerValidators(address podOwner, bytes32[] blsPubKeyHashes) returns()
func (_AvsManager *AvsManagerTransactorSession) RegisterValidators(podOwner common.Address, blsPubKeyHashes [][32]byte) (*types.Transaction, error) {
	return _AvsManager.Contract.RegisterValidators(&_AvsManager.TransactOpts, podOwner, blsPubKeyHashes)
}

// SetAllowlistRestakingStrategy is a paid mutator transaction binding the contract method 0x43adc491.
//
// Solidity: function setAllowlistRestakingStrategy(address strategy, bool allowed) returns()
func (_AvsManager *AvsManagerTransactor) SetAllowlistRestakingStrategy(opts *bind.TransactOpts, strategy common.Address, allowed bool) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "setAllowlistRestakingStrategy", strategy, allowed)
}

// SetAllowlistRestakingStrategy is a paid mutator transaction binding the contract method 0x43adc491.
//
// Solidity: function setAllowlistRestakingStrategy(address strategy, bool allowed) returns()
func (_AvsManager *AvsManagerSession) SetAllowlistRestakingStrategy(strategy common.Address, allowed bool) (*types.Transaction, error) {
	return _AvsManager.Contract.SetAllowlistRestakingStrategy(&_AvsManager.TransactOpts, strategy, allowed)
}

// SetAllowlistRestakingStrategy is a paid mutator transaction binding the contract method 0x43adc491.
//
// Solidity: function setAllowlistRestakingStrategy(address strategy, bool allowed) returns()
func (_AvsManager *AvsManagerTransactorSession) SetAllowlistRestakingStrategy(strategy common.Address, allowed bool) (*types.Transaction, error) {
	return _AvsManager.Contract.SetAllowlistRestakingStrategy(&_AvsManager.TransactOpts, strategy, allowed)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_AvsManager *AvsManagerTransactor) SetAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "setAuthority", newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_AvsManager *AvsManagerSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _AvsManager.Contract.SetAuthority(&_AvsManager.TransactOpts, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_AvsManager *AvsManagerTransactorSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _AvsManager.Contract.SetAuthority(&_AvsManager.TransactOpts, newAuthority)
}

// SetChainID is a paid mutator transaction binding the contract method 0xc5a444cd.
//
// Solidity: function setChainID(uint8 index, uint256 chainID) returns()
func (_AvsManager *AvsManagerTransactor) SetChainID(opts *bind.TransactOpts, index uint8, chainID *big.Int) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "setChainID", index, chainID)
}

// SetChainID is a paid mutator transaction binding the contract method 0xc5a444cd.
//
// Solidity: function setChainID(uint8 index, uint256 chainID) returns()
func (_AvsManager *AvsManagerSession) SetChainID(index uint8, chainID *big.Int) (*types.Transaction, error) {
	return _AvsManager.Contract.SetChainID(&_AvsManager.TransactOpts, index, chainID)
}

// SetChainID is a paid mutator transaction binding the contract method 0xc5a444cd.
//
// Solidity: function setChainID(uint8 index, uint256 chainID) returns()
func (_AvsManager *AvsManagerTransactorSession) SetChainID(index uint8, chainID *big.Int) (*types.Transaction, error) {
	return _AvsManager.Contract.SetChainID(&_AvsManager.TransactOpts, index, chainID)
}

// SetDeregistrationDelay is a paid mutator transaction binding the contract method 0x84a3f19b.
//
// Solidity: function setDeregistrationDelay(uint64 newDelay) returns()
func (_AvsManager *AvsManagerTransactor) SetDeregistrationDelay(opts *bind.TransactOpts, newDelay uint64) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "setDeregistrationDelay", newDelay)
}

// SetDeregistrationDelay is a paid mutator transaction binding the contract method 0x84a3f19b.
//
// Solidity: function setDeregistrationDelay(uint64 newDelay) returns()
func (_AvsManager *AvsManagerSession) SetDeregistrationDelay(newDelay uint64) (*types.Transaction, error) {
	return _AvsManager.Contract.SetDeregistrationDelay(&_AvsManager.TransactOpts, newDelay)
}

// SetDeregistrationDelay is a paid mutator transaction binding the contract method 0x84a3f19b.
//
// Solidity: function setDeregistrationDelay(uint64 newDelay) returns()
func (_AvsManager *AvsManagerTransactorSession) SetDeregistrationDelay(newDelay uint64) (*types.Transaction, error) {
	return _AvsManager.Contract.SetDeregistrationDelay(&_AvsManager.TransactOpts, newDelay)
}

// SetOperatorCommitment is a paid mutator transaction binding the contract method 0x6a01b2c9.
//
// Solidity: function setOperatorCommitment((bytes,uint256) newCommitment) returns()
func (_AvsManager *AvsManagerTransactor) SetOperatorCommitment(opts *bind.TransactOpts, newCommitment OperatorCommitment) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "setOperatorCommitment", newCommitment)
}

// SetOperatorCommitment is a paid mutator transaction binding the contract method 0x6a01b2c9.
//
// Solidity: function setOperatorCommitment((bytes,uint256) newCommitment) returns()
func (_AvsManager *AvsManagerSession) SetOperatorCommitment(newCommitment OperatorCommitment) (*types.Transaction, error) {
	return _AvsManager.Contract.SetOperatorCommitment(&_AvsManager.TransactOpts, newCommitment)
}

// SetOperatorCommitment is a paid mutator transaction binding the contract method 0x6a01b2c9.
//
// Solidity: function setOperatorCommitment((bytes,uint256) newCommitment) returns()
func (_AvsManager *AvsManagerTransactorSession) SetOperatorCommitment(newCommitment OperatorCommitment) (*types.Transaction, error) {
	return _AvsManager.Contract.SetOperatorCommitment(&_AvsManager.TransactOpts, newCommitment)
}

// StartDeregisterOperator is a paid mutator transaction binding the contract method 0x389517e4.
//
// Solidity: function startDeregisterOperator() returns()
func (_AvsManager *AvsManagerTransactor) StartDeregisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "startDeregisterOperator")
}

// StartDeregisterOperator is a paid mutator transaction binding the contract method 0x389517e4.
//
// Solidity: function startDeregisterOperator() returns()
func (_AvsManager *AvsManagerSession) StartDeregisterOperator() (*types.Transaction, error) {
	return _AvsManager.Contract.StartDeregisterOperator(&_AvsManager.TransactOpts)
}

// StartDeregisterOperator is a paid mutator transaction binding the contract method 0x389517e4.
//
// Solidity: function startDeregisterOperator() returns()
func (_AvsManager *AvsManagerTransactorSession) StartDeregisterOperator() (*types.Transaction, error) {
	return _AvsManager.Contract.StartDeregisterOperator(&_AvsManager.TransactOpts)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_AvsManager *AvsManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_AvsManager *AvsManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _AvsManager.Contract.UpdateAVSMetadataURI(&_AvsManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_AvsManager *AvsManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _AvsManager.Contract.UpdateAVSMetadataURI(&_AvsManager.TransactOpts, _metadataURI)
}

// UpdateOperatorCommitment is a paid mutator transaction binding the contract method 0x4d9739bf.
//
// Solidity: function updateOperatorCommitment() returns()
func (_AvsManager *AvsManagerTransactor) UpdateOperatorCommitment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "updateOperatorCommitment")
}

// UpdateOperatorCommitment is a paid mutator transaction binding the contract method 0x4d9739bf.
//
// Solidity: function updateOperatorCommitment() returns()
func (_AvsManager *AvsManagerSession) UpdateOperatorCommitment() (*types.Transaction, error) {
	return _AvsManager.Contract.UpdateOperatorCommitment(&_AvsManager.TransactOpts)
}

// UpdateOperatorCommitment is a paid mutator transaction binding the contract method 0x4d9739bf.
//
// Solidity: function updateOperatorCommitment() returns()
func (_AvsManager *AvsManagerTransactorSession) UpdateOperatorCommitment() (*types.Transaction, error) {
	return _AvsManager.Contract.UpdateOperatorCommitment(&_AvsManager.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AvsManager *AvsManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AvsManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AvsManager *AvsManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AvsManager.Contract.UpgradeToAndCall(&_AvsManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AvsManager *AvsManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AvsManager.Contract.UpgradeToAndCall(&_AvsManager.TransactOpts, newImplementation, data)
}

// AvsManagerAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the AvsManager contract.
type AvsManagerAuthorityUpdatedIterator struct {
	Event *AvsManagerAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *AvsManagerAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerAuthorityUpdated)
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
		it.Event = new(AvsManagerAuthorityUpdated)
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
func (it *AvsManagerAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerAuthorityUpdated represents a AuthorityUpdated event raised by the AvsManager contract.
type AvsManagerAuthorityUpdated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_AvsManager *AvsManagerFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts) (*AvsManagerAuthorityUpdatedIterator, error) {

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return &AvsManagerAuthorityUpdatedIterator{contract: _AvsManager.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_AvsManager *AvsManagerFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *AvsManagerAuthorityUpdated) (event.Subscription, error) {

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "AuthorityUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerAuthorityUpdated)
				if err := _AvsManager.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address authority)
func (_AvsManager *AvsManagerFilterer) ParseAuthorityUpdated(log types.Log) (*AvsManagerAuthorityUpdated, error) {
	event := new(AvsManagerAuthorityUpdated)
	if err := _AvsManager.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerChainIDSetIterator is returned from FilterChainIDSet and is used to iterate over the raw logs and unpacked data for ChainIDSet events raised by the AvsManager contract.
type AvsManagerChainIDSetIterator struct {
	Event *AvsManagerChainIDSet // Event containing the contract specifics and raw log

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
func (it *AvsManagerChainIDSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerChainIDSet)
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
		it.Event = new(AvsManagerChainIDSet)
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
func (it *AvsManagerChainIDSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerChainIDSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerChainIDSet represents a ChainIDSet event raised by the AvsManager contract.
type AvsManagerChainIDSet struct {
	Index   uint8
	ChainID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainIDSet is a free log retrieval operation binding the contract event 0x593f668a875f6d1f8fdcff32e09eba9648a341ac0bad6168933dd5ebd545513c.
//
// Solidity: event ChainIDSet(uint8 index, uint256 chainID)
func (_AvsManager *AvsManagerFilterer) FilterChainIDSet(opts *bind.FilterOpts) (*AvsManagerChainIDSetIterator, error) {

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "ChainIDSet")
	if err != nil {
		return nil, err
	}
	return &AvsManagerChainIDSetIterator{contract: _AvsManager.contract, event: "ChainIDSet", logs: logs, sub: sub}, nil
}

// WatchChainIDSet is a free log subscription operation binding the contract event 0x593f668a875f6d1f8fdcff32e09eba9648a341ac0bad6168933dd5ebd545513c.
//
// Solidity: event ChainIDSet(uint8 index, uint256 chainID)
func (_AvsManager *AvsManagerFilterer) WatchChainIDSet(opts *bind.WatchOpts, sink chan<- *AvsManagerChainIDSet) (event.Subscription, error) {

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "ChainIDSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerChainIDSet)
				if err := _AvsManager.contract.UnpackLog(event, "ChainIDSet", log); err != nil {
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

// ParseChainIDSet is a log parse operation binding the contract event 0x593f668a875f6d1f8fdcff32e09eba9648a341ac0bad6168933dd5ebd545513c.
//
// Solidity: event ChainIDSet(uint8 index, uint256 chainID)
func (_AvsManager *AvsManagerFilterer) ParseChainIDSet(log types.Log) (*AvsManagerChainIDSet, error) {
	event := new(AvsManagerChainIDSet)
	if err := _AvsManager.contract.UnpackLog(event, "ChainIDSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerDeregistrationDelaySetIterator is returned from FilterDeregistrationDelaySet and is used to iterate over the raw logs and unpacked data for DeregistrationDelaySet events raised by the AvsManager contract.
type AvsManagerDeregistrationDelaySetIterator struct {
	Event *AvsManagerDeregistrationDelaySet // Event containing the contract specifics and raw log

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
func (it *AvsManagerDeregistrationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerDeregistrationDelaySet)
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
		it.Event = new(AvsManagerDeregistrationDelaySet)
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
func (it *AvsManagerDeregistrationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerDeregistrationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerDeregistrationDelaySet represents a DeregistrationDelaySet event raised by the AvsManager contract.
type AvsManagerDeregistrationDelaySet struct {
	OldDelay uint64
	NewDelay uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeregistrationDelaySet is a free log retrieval operation binding the contract event 0x7d9c9ce14fd1f11f09bf0faca4d408d3bb5815ceb44368d65b408ebc4401165c.
//
// Solidity: event DeregistrationDelaySet(uint64 oldDelay, uint64 newDelay)
func (_AvsManager *AvsManagerFilterer) FilterDeregistrationDelaySet(opts *bind.FilterOpts) (*AvsManagerDeregistrationDelaySetIterator, error) {

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "DeregistrationDelaySet")
	if err != nil {
		return nil, err
	}
	return &AvsManagerDeregistrationDelaySetIterator{contract: _AvsManager.contract, event: "DeregistrationDelaySet", logs: logs, sub: sub}, nil
}

// WatchDeregistrationDelaySet is a free log subscription operation binding the contract event 0x7d9c9ce14fd1f11f09bf0faca4d408d3bb5815ceb44368d65b408ebc4401165c.
//
// Solidity: event DeregistrationDelaySet(uint64 oldDelay, uint64 newDelay)
func (_AvsManager *AvsManagerFilterer) WatchDeregistrationDelaySet(opts *bind.WatchOpts, sink chan<- *AvsManagerDeregistrationDelaySet) (event.Subscription, error) {

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "DeregistrationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerDeregistrationDelaySet)
				if err := _AvsManager.contract.UnpackLog(event, "DeregistrationDelaySet", log); err != nil {
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

// ParseDeregistrationDelaySet is a log parse operation binding the contract event 0x7d9c9ce14fd1f11f09bf0faca4d408d3bb5815ceb44368d65b408ebc4401165c.
//
// Solidity: event DeregistrationDelaySet(uint64 oldDelay, uint64 newDelay)
func (_AvsManager *AvsManagerFilterer) ParseDeregistrationDelaySet(log types.Log) (*AvsManagerDeregistrationDelaySet, error) {
	event := new(AvsManagerDeregistrationDelaySet)
	if err := _AvsManager.contract.UnpackLog(event, "DeregistrationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AvsManager contract.
type AvsManagerInitializedIterator struct {
	Event *AvsManagerInitialized // Event containing the contract specifics and raw log

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
func (it *AvsManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerInitialized)
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
		it.Event = new(AvsManagerInitialized)
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
func (it *AvsManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerInitialized represents a Initialized event raised by the AvsManager contract.
type AvsManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AvsManager *AvsManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AvsManagerInitializedIterator, error) {

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AvsManagerInitializedIterator{contract: _AvsManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AvsManager *AvsManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AvsManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerInitialized)
				if err := _AvsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AvsManager *AvsManagerFilterer) ParseInitialized(log types.Log) (*AvsManagerInitialized, error) {
	event := new(AvsManagerInitialized)
	if err := _AvsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerOperatorCommitmentChangeInitiatedIterator is returned from FilterOperatorCommitmentChangeInitiated and is used to iterate over the raw logs and unpacked data for OperatorCommitmentChangeInitiated events raised by the AvsManager contract.
type AvsManagerOperatorCommitmentChangeInitiatedIterator struct {
	Event *AvsManagerOperatorCommitmentChangeInitiated // Event containing the contract specifics and raw log

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
func (it *AvsManagerOperatorCommitmentChangeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerOperatorCommitmentChangeInitiated)
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
		it.Event = new(AvsManagerOperatorCommitmentChangeInitiated)
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
func (it *AvsManagerOperatorCommitmentChangeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerOperatorCommitmentChangeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerOperatorCommitmentChangeInitiated represents a OperatorCommitmentChangeInitiated event raised by the AvsManager contract.
type AvsManagerOperatorCommitmentChangeInitiated struct {
	Operator      common.Address
	OldCommitment OperatorCommitment
	NewCommitment OperatorCommitment
	ValidAfter    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorCommitmentChangeInitiated is a free log retrieval operation binding the contract event 0xbeb3e5f671ae7d66f24476749e90eea4117fd593dc77a9e8405f0199cd4dafd9.
//
// Solidity: event OperatorCommitmentChangeInitiated(address indexed operator, (bytes,uint256) oldCommitment, (bytes,uint256) newCommitment, uint128 validAfter)
func (_AvsManager *AvsManagerFilterer) FilterOperatorCommitmentChangeInitiated(opts *bind.FilterOpts, operator []common.Address) (*AvsManagerOperatorCommitmentChangeInitiatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "OperatorCommitmentChangeInitiated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerOperatorCommitmentChangeInitiatedIterator{contract: _AvsManager.contract, event: "OperatorCommitmentChangeInitiated", logs: logs, sub: sub}, nil
}

// WatchOperatorCommitmentChangeInitiated is a free log subscription operation binding the contract event 0xbeb3e5f671ae7d66f24476749e90eea4117fd593dc77a9e8405f0199cd4dafd9.
//
// Solidity: event OperatorCommitmentChangeInitiated(address indexed operator, (bytes,uint256) oldCommitment, (bytes,uint256) newCommitment, uint128 validAfter)
func (_AvsManager *AvsManagerFilterer) WatchOperatorCommitmentChangeInitiated(opts *bind.WatchOpts, sink chan<- *AvsManagerOperatorCommitmentChangeInitiated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "OperatorCommitmentChangeInitiated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerOperatorCommitmentChangeInitiated)
				if err := _AvsManager.contract.UnpackLog(event, "OperatorCommitmentChangeInitiated", log); err != nil {
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

// ParseOperatorCommitmentChangeInitiated is a log parse operation binding the contract event 0xbeb3e5f671ae7d66f24476749e90eea4117fd593dc77a9e8405f0199cd4dafd9.
//
// Solidity: event OperatorCommitmentChangeInitiated(address indexed operator, (bytes,uint256) oldCommitment, (bytes,uint256) newCommitment, uint128 validAfter)
func (_AvsManager *AvsManagerFilterer) ParseOperatorCommitmentChangeInitiated(log types.Log) (*AvsManagerOperatorCommitmentChangeInitiated, error) {
	event := new(AvsManagerOperatorCommitmentChangeInitiated)
	if err := _AvsManager.contract.UnpackLog(event, "OperatorCommitmentChangeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerOperatorCommitmentSetIterator is returned from FilterOperatorCommitmentSet and is used to iterate over the raw logs and unpacked data for OperatorCommitmentSet events raised by the AvsManager contract.
type AvsManagerOperatorCommitmentSetIterator struct {
	Event *AvsManagerOperatorCommitmentSet // Event containing the contract specifics and raw log

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
func (it *AvsManagerOperatorCommitmentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerOperatorCommitmentSet)
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
		it.Event = new(AvsManagerOperatorCommitmentSet)
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
func (it *AvsManagerOperatorCommitmentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerOperatorCommitmentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerOperatorCommitmentSet represents a OperatorCommitmentSet event raised by the AvsManager contract.
type AvsManagerOperatorCommitmentSet struct {
	Operator      common.Address
	OldCommitment OperatorCommitment
	NewCommitment OperatorCommitment
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorCommitmentSet is a free log retrieval operation binding the contract event 0x7a8fb671d54b440e1e9cdf2a99020b77bdb1e7b8d217302483cfc00558be6588.
//
// Solidity: event OperatorCommitmentSet(address indexed operator, (bytes,uint256) oldCommitment, (bytes,uint256) newCommitment)
func (_AvsManager *AvsManagerFilterer) FilterOperatorCommitmentSet(opts *bind.FilterOpts, operator []common.Address) (*AvsManagerOperatorCommitmentSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "OperatorCommitmentSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerOperatorCommitmentSetIterator{contract: _AvsManager.contract, event: "OperatorCommitmentSet", logs: logs, sub: sub}, nil
}

// WatchOperatorCommitmentSet is a free log subscription operation binding the contract event 0x7a8fb671d54b440e1e9cdf2a99020b77bdb1e7b8d217302483cfc00558be6588.
//
// Solidity: event OperatorCommitmentSet(address indexed operator, (bytes,uint256) oldCommitment, (bytes,uint256) newCommitment)
func (_AvsManager *AvsManagerFilterer) WatchOperatorCommitmentSet(opts *bind.WatchOpts, sink chan<- *AvsManagerOperatorCommitmentSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "OperatorCommitmentSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerOperatorCommitmentSet)
				if err := _AvsManager.contract.UnpackLog(event, "OperatorCommitmentSet", log); err != nil {
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

// ParseOperatorCommitmentSet is a log parse operation binding the contract event 0x7a8fb671d54b440e1e9cdf2a99020b77bdb1e7b8d217302483cfc00558be6588.
//
// Solidity: event OperatorCommitmentSet(address indexed operator, (bytes,uint256) oldCommitment, (bytes,uint256) newCommitment)
func (_AvsManager *AvsManagerFilterer) ParseOperatorCommitmentSet(log types.Log) (*AvsManagerOperatorCommitmentSet, error) {
	event := new(AvsManagerOperatorCommitmentSet)
	if err := _AvsManager.contract.UnpackLog(event, "OperatorCommitmentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerOperatorDeregisterStartedIterator is returned from FilterOperatorDeregisterStarted and is used to iterate over the raw logs and unpacked data for OperatorDeregisterStarted events raised by the AvsManager contract.
type AvsManagerOperatorDeregisterStartedIterator struct {
	Event *AvsManagerOperatorDeregisterStarted // Event containing the contract specifics and raw log

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
func (it *AvsManagerOperatorDeregisterStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerOperatorDeregisterStarted)
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
		it.Event = new(AvsManagerOperatorDeregisterStarted)
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
func (it *AvsManagerOperatorDeregisterStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerOperatorDeregisterStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerOperatorDeregisterStarted represents a OperatorDeregisterStarted event raised by the AvsManager contract.
type AvsManagerOperatorDeregisterStarted struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregisterStarted is a free log retrieval operation binding the contract event 0x34e20e11f51d19d6e90f5aaea605cb369f34ad23805c51d2e77d05c8208c8585.
//
// Solidity: event OperatorDeregisterStarted(address indexed operator)
func (_AvsManager *AvsManagerFilterer) FilterOperatorDeregisterStarted(opts *bind.FilterOpts, operator []common.Address) (*AvsManagerOperatorDeregisterStartedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "OperatorDeregisterStarted", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerOperatorDeregisterStartedIterator{contract: _AvsManager.contract, event: "OperatorDeregisterStarted", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregisterStarted is a free log subscription operation binding the contract event 0x34e20e11f51d19d6e90f5aaea605cb369f34ad23805c51d2e77d05c8208c8585.
//
// Solidity: event OperatorDeregisterStarted(address indexed operator)
func (_AvsManager *AvsManagerFilterer) WatchOperatorDeregisterStarted(opts *bind.WatchOpts, sink chan<- *AvsManagerOperatorDeregisterStarted, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "OperatorDeregisterStarted", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerOperatorDeregisterStarted)
				if err := _AvsManager.contract.UnpackLog(event, "OperatorDeregisterStarted", log); err != nil {
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

// ParseOperatorDeregisterStarted is a log parse operation binding the contract event 0x34e20e11f51d19d6e90f5aaea605cb369f34ad23805c51d2e77d05c8208c8585.
//
// Solidity: event OperatorDeregisterStarted(address indexed operator)
func (_AvsManager *AvsManagerFilterer) ParseOperatorDeregisterStarted(log types.Log) (*AvsManagerOperatorDeregisterStarted, error) {
	event := new(AvsManagerOperatorDeregisterStarted)
	if err := _AvsManager.contract.UnpackLog(event, "OperatorDeregisterStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the AvsManager contract.
type AvsManagerOperatorDeregisteredIterator struct {
	Event *AvsManagerOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *AvsManagerOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerOperatorDeregistered)
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
		it.Event = new(AvsManagerOperatorDeregistered)
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
func (it *AvsManagerOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerOperatorDeregistered represents a OperatorDeregistered event raised by the AvsManager contract.
type AvsManagerOperatorDeregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_AvsManager *AvsManagerFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*AvsManagerOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerOperatorDeregisteredIterator{contract: _AvsManager.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_AvsManager *AvsManagerFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *AvsManagerOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerOperatorDeregistered)
				if err := _AvsManager.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x6dd4ca66565fb3dee8076c654634c6c4ad949022d809d0394308617d6791218d.
//
// Solidity: event OperatorDeregistered(address indexed operator)
func (_AvsManager *AvsManagerFilterer) ParseOperatorDeregistered(log types.Log) (*AvsManagerOperatorDeregistered, error) {
	event := new(AvsManagerOperatorDeregistered)
	if err := _AvsManager.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the AvsManager contract.
type AvsManagerOperatorRegisteredIterator struct {
	Event *AvsManagerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *AvsManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerOperatorRegistered)
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
		it.Event = new(AvsManagerOperatorRegistered)
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
func (it *AvsManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerOperatorRegistered represents a OperatorRegistered event raised by the AvsManager contract.
type AvsManagerOperatorRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_AvsManager *AvsManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*AvsManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerOperatorRegisteredIterator{contract: _AvsManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x4d0eb1f4bac8744fd2be119845e23b3befc88094b42bcda1204c65694a00f9e5.
//
// Solidity: event OperatorRegistered(address indexed operator)
func (_AvsManager *AvsManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *AvsManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerOperatorRegistered)
				if err := _AvsManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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
func (_AvsManager *AvsManagerFilterer) ParseOperatorRegistered(log types.Log) (*AvsManagerOperatorRegistered, error) {
	event := new(AvsManagerOperatorRegistered)
	if err := _AvsManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerOperatorRegisteredWithCommitmentIterator is returned from FilterOperatorRegisteredWithCommitment and is used to iterate over the raw logs and unpacked data for OperatorRegisteredWithCommitment events raised by the AvsManager contract.
type AvsManagerOperatorRegisteredWithCommitmentIterator struct {
	Event *AvsManagerOperatorRegisteredWithCommitment // Event containing the contract specifics and raw log

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
func (it *AvsManagerOperatorRegisteredWithCommitmentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerOperatorRegisteredWithCommitment)
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
		it.Event = new(AvsManagerOperatorRegisteredWithCommitment)
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
func (it *AvsManagerOperatorRegisteredWithCommitmentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerOperatorRegisteredWithCommitmentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerOperatorRegisteredWithCommitment represents a OperatorRegisteredWithCommitment event raised by the AvsManager contract.
type AvsManagerOperatorRegisteredWithCommitment struct {
	Operator   common.Address
	Commitment OperatorCommitment
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegisteredWithCommitment is a free log retrieval operation binding the contract event 0xc4dd9a1c9f9fdfb9b1ec214ecaa456e160eeb015e9a68f5d241f896158d2cab8.
//
// Solidity: event OperatorRegisteredWithCommitment(address indexed operator, (bytes,uint256) commitment)
func (_AvsManager *AvsManagerFilterer) FilterOperatorRegisteredWithCommitment(opts *bind.FilterOpts, operator []common.Address) (*AvsManagerOperatorRegisteredWithCommitmentIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "OperatorRegisteredWithCommitment", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerOperatorRegisteredWithCommitmentIterator{contract: _AvsManager.contract, event: "OperatorRegisteredWithCommitment", logs: logs, sub: sub}, nil
}

// WatchOperatorRegisteredWithCommitment is a free log subscription operation binding the contract event 0xc4dd9a1c9f9fdfb9b1ec214ecaa456e160eeb015e9a68f5d241f896158d2cab8.
//
// Solidity: event OperatorRegisteredWithCommitment(address indexed operator, (bytes,uint256) commitment)
func (_AvsManager *AvsManagerFilterer) WatchOperatorRegisteredWithCommitment(opts *bind.WatchOpts, sink chan<- *AvsManagerOperatorRegisteredWithCommitment, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "OperatorRegisteredWithCommitment", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerOperatorRegisteredWithCommitment)
				if err := _AvsManager.contract.UnpackLog(event, "OperatorRegisteredWithCommitment", log); err != nil {
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

// ParseOperatorRegisteredWithCommitment is a log parse operation binding the contract event 0xc4dd9a1c9f9fdfb9b1ec214ecaa456e160eeb015e9a68f5d241f896158d2cab8.
//
// Solidity: event OperatorRegisteredWithCommitment(address indexed operator, (bytes,uint256) commitment)
func (_AvsManager *AvsManagerFilterer) ParseOperatorRegisteredWithCommitment(log types.Log) (*AvsManagerOperatorRegisteredWithCommitment, error) {
	event := new(AvsManagerOperatorRegisteredWithCommitment)
	if err := _AvsManager.contract.UnpackLog(event, "OperatorRegisteredWithCommitment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerRestakingStrategyAllowlistUpdatedIterator is returned from FilterRestakingStrategyAllowlistUpdated and is used to iterate over the raw logs and unpacked data for RestakingStrategyAllowlistUpdated events raised by the AvsManager contract.
type AvsManagerRestakingStrategyAllowlistUpdatedIterator struct {
	Event *AvsManagerRestakingStrategyAllowlistUpdated // Event containing the contract specifics and raw log

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
func (it *AvsManagerRestakingStrategyAllowlistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerRestakingStrategyAllowlistUpdated)
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
		it.Event = new(AvsManagerRestakingStrategyAllowlistUpdated)
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
func (it *AvsManagerRestakingStrategyAllowlistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerRestakingStrategyAllowlistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerRestakingStrategyAllowlistUpdated represents a RestakingStrategyAllowlistUpdated event raised by the AvsManager contract.
type AvsManagerRestakingStrategyAllowlistUpdated struct {
	Strategy common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRestakingStrategyAllowlistUpdated is a free log retrieval operation binding the contract event 0x74d2e573146a4c9fbfd23124ce9007a97aebeb715219b83507ec3a05209486bd.
//
// Solidity: event RestakingStrategyAllowlistUpdated(address indexed strategy, bool allowed)
func (_AvsManager *AvsManagerFilterer) FilterRestakingStrategyAllowlistUpdated(opts *bind.FilterOpts, strategy []common.Address) (*AvsManagerRestakingStrategyAllowlistUpdatedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "RestakingStrategyAllowlistUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerRestakingStrategyAllowlistUpdatedIterator{contract: _AvsManager.contract, event: "RestakingStrategyAllowlistUpdated", logs: logs, sub: sub}, nil
}

// WatchRestakingStrategyAllowlistUpdated is a free log subscription operation binding the contract event 0x74d2e573146a4c9fbfd23124ce9007a97aebeb715219b83507ec3a05209486bd.
//
// Solidity: event RestakingStrategyAllowlistUpdated(address indexed strategy, bool allowed)
func (_AvsManager *AvsManagerFilterer) WatchRestakingStrategyAllowlistUpdated(opts *bind.WatchOpts, sink chan<- *AvsManagerRestakingStrategyAllowlistUpdated, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "RestakingStrategyAllowlistUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerRestakingStrategyAllowlistUpdated)
				if err := _AvsManager.contract.UnpackLog(event, "RestakingStrategyAllowlistUpdated", log); err != nil {
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

// ParseRestakingStrategyAllowlistUpdated is a log parse operation binding the contract event 0x74d2e573146a4c9fbfd23124ce9007a97aebeb715219b83507ec3a05209486bd.
//
// Solidity: event RestakingStrategyAllowlistUpdated(address indexed strategy, bool allowed)
func (_AvsManager *AvsManagerFilterer) ParseRestakingStrategyAllowlistUpdated(log types.Log) (*AvsManagerRestakingStrategyAllowlistUpdated, error) {
	event := new(AvsManagerRestakingStrategyAllowlistUpdated)
	if err := _AvsManager.contract.UnpackLog(event, "RestakingStrategyAllowlistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AvsManager contract.
type AvsManagerUpgradedIterator struct {
	Event *AvsManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *AvsManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerUpgraded)
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
		it.Event = new(AvsManagerUpgraded)
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
func (it *AvsManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerUpgraded represents a Upgraded event raised by the AvsManager contract.
type AvsManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AvsManager *AvsManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AvsManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerUpgradedIterator{contract: _AvsManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AvsManager *AvsManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AvsManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerUpgraded)
				if err := _AvsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AvsManager *AvsManagerFilterer) ParseUpgraded(log types.Log) (*AvsManagerUpgraded, error) {
	event := new(AvsManagerUpgraded)
	if err := _AvsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerValidatorDeregisteredIterator is returned from FilterValidatorDeregistered and is used to iterate over the raw logs and unpacked data for ValidatorDeregistered events raised by the AvsManager contract.
type AvsManagerValidatorDeregisteredIterator struct {
	Event *AvsManagerValidatorDeregistered // Event containing the contract specifics and raw log

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
func (it *AvsManagerValidatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerValidatorDeregistered)
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
		it.Event = new(AvsManagerValidatorDeregistered)
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
func (it *AvsManagerValidatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerValidatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerValidatorDeregistered represents a ValidatorDeregistered event raised by the AvsManager contract.
type AvsManagerValidatorDeregistered struct {
	Operator      common.Address
	BlsPubKeyHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeregistered is a free log retrieval operation binding the contract event 0x9100d6f32d745f208c0e70d0679e591d81509157db766dc5eb3d092ac047f067.
//
// Solidity: event ValidatorDeregistered(address indexed operator, bytes32 blsPubKeyHash)
func (_AvsManager *AvsManagerFilterer) FilterValidatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*AvsManagerValidatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "ValidatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerValidatorDeregisteredIterator{contract: _AvsManager.contract, event: "ValidatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchValidatorDeregistered is a free log subscription operation binding the contract event 0x9100d6f32d745f208c0e70d0679e591d81509157db766dc5eb3d092ac047f067.
//
// Solidity: event ValidatorDeregistered(address indexed operator, bytes32 blsPubKeyHash)
func (_AvsManager *AvsManagerFilterer) WatchValidatorDeregistered(opts *bind.WatchOpts, sink chan<- *AvsManagerValidatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "ValidatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerValidatorDeregistered)
				if err := _AvsManager.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
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

// ParseValidatorDeregistered is a log parse operation binding the contract event 0x9100d6f32d745f208c0e70d0679e591d81509157db766dc5eb3d092ac047f067.
//
// Solidity: event ValidatorDeregistered(address indexed operator, bytes32 blsPubKeyHash)
func (_AvsManager *AvsManagerFilterer) ParseValidatorDeregistered(log types.Log) (*AvsManagerValidatorDeregistered, error) {
	event := new(AvsManagerValidatorDeregistered)
	if err := _AvsManager.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AvsManagerValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the AvsManager contract.
type AvsManagerValidatorRegisteredIterator struct {
	Event *AvsManagerValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *AvsManagerValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AvsManagerValidatorRegistered)
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
		it.Event = new(AvsManagerValidatorRegistered)
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
func (it *AvsManagerValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AvsManagerValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AvsManagerValidatorRegistered represents a ValidatorRegistered event raised by the AvsManager contract.
type AvsManagerValidatorRegistered struct {
	PodOwner       common.Address
	Operator       common.Address
	DelegateKey    []byte
	BlsPubKeyHash  [32]byte
	ValidatorIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x28e80e6d7e989497bd036ed5c5484b9b46fe0d438a8ea0918e31d3cc59817f65.
//
// Solidity: event ValidatorRegistered(address indexed podOwner, address indexed operator, bytes delegateKey, bytes32 blsPubKeyHash, uint256 validatorIndex)
func (_AvsManager *AvsManagerFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, podOwner []common.Address, operator []common.Address) (*AvsManagerValidatorRegisteredIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.FilterLogs(opts, "ValidatorRegistered", podOwnerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AvsManagerValidatorRegisteredIterator{contract: _AvsManager.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x28e80e6d7e989497bd036ed5c5484b9b46fe0d438a8ea0918e31d3cc59817f65.
//
// Solidity: event ValidatorRegistered(address indexed podOwner, address indexed operator, bytes delegateKey, bytes32 blsPubKeyHash, uint256 validatorIndex)
func (_AvsManager *AvsManagerFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *AvsManagerValidatorRegistered, podOwner []common.Address, operator []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AvsManager.contract.WatchLogs(opts, "ValidatorRegistered", podOwnerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AvsManagerValidatorRegistered)
				if err := _AvsManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0x28e80e6d7e989497bd036ed5c5484b9b46fe0d438a8ea0918e31d3cc59817f65.
//
// Solidity: event ValidatorRegistered(address indexed podOwner, address indexed operator, bytes delegateKey, bytes32 blsPubKeyHash, uint256 validatorIndex)
func (_AvsManager *AvsManagerFilterer) ParseValidatorRegistered(log types.Log) (*AvsManagerValidatorRegistered, error) {
	event := new(AvsManagerValidatorRegistered)
	if err := _AvsManager.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
