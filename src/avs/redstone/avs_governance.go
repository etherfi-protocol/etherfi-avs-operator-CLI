// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package redstone

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

// BLSAuthLibrarySignature is an auto generated low-level Go binding around an user-defined struct.
type BLSAuthLibrarySignature struct {
	Signature [2]*big.Int
}

// IAvsGovernanceInitializationParams is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceInitializationParams struct {
	AvsGovernanceMultisigOwner common.Address
	OperationsMultisig         common.Address
	CommunityMultisig          common.Address
	OthenticRegistry           common.Address
	MessageHandler             common.Address
	Vault                      common.Address
	AvsDirectoryContract       common.Address
	AllowlistSigner            common.Address
	AvsName                    string
	BlsAuthSingleton           common.Address
}

// IAvsGovernancePaymentRequestMessage is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernancePaymentRequestMessage struct {
	Operator   common.Address
	FeeToClaim *big.Int
}

// IAvsGovernanceStrategyMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IAvsGovernanceStrategyMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// AVSGovernanceMetaData contains all meta data concerning the AVSGovernance contract.
var AVSGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccessControlInvalidMultiplierSyncer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowlistEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlowIsCurrentlyUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAllowlistAuthToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlsRegistrationSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiplierNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardsReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlashingRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStrategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModificationDelayNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughVotingPower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfOperatorsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numOfActiveOperators\",\"type\":\"uint256\"}],\"name\":\"NumOfActiveOperatorsIsGreaterThanNumOfOperatorLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfOperatorsLimit\",\"type\":\"uint256\"}],\"name\":\"NumOfOperatorsLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseFlowIsAlreadyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnpausingFlowIsAlreadyUnpaused\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"blsAuthSingleton\",\"type\":\"address\"}],\"name\":\"BLSAuthSingletonSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"}],\"name\":\"FlowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"_pausableFlowFlag\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_unpauser\",\"type\":\"address\"}],\"name\":\"FlowUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxEffectiveBalance\",\"type\":\"uint256\"}],\"name\":\"MaxEffectiveBalanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minShares\",\"type\":\"uint256\"}],\"name\":\"MinSharesPerStrategySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minVotingPower\",\"type\":\"uint256\"}],\"name\":\"MinVotingPowerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[4]\",\"name\":\"blsKey\",\"type\":\"uint256[4]\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"QueuedRewardsReceiverModification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"allowlistSigner\",\"type\":\"address\"}],\"name\":\"SetAllowlistSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsGovernanceLogic\",\"type\":\"address\"}],\"name\":\"SetAvsGovernanceLogic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsGovernanceMultiplierSyncer\",\"type\":\"address\"}],\"name\":\"SetAvsGovernanceMultiplierSyncer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAvsGovernanceMultisig\",\"type\":\"address\"}],\"name\":\"SetAvsGovernanceMultisig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"}],\"name\":\"SetAvsName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowlisted\",\"type\":\"bool\"}],\"name\":\"SetIsAllowlisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newMessageHandler\",\"type\":\"address\"}],\"name\":\"SetMessageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimitOfNumOfOperators\",\"type\":\"uint256\"}],\"name\":\"SetNumOfOperatorsLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"othenticRegistry\",\"type\":\"address\"}],\"name\":\"SetOthenticRegistry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"SetRewardsReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modificationDelay\",\"type\":\"uint256\"}],\"name\":\"SetRewardsReceiverModificationDelay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"SetStrategyMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"SetSupportedStrategies\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SetToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"completeRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsAllowlisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumOfOperatorsLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfOperatorsLimitView\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getRewardsReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"avsGovernanceMultisigOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operationsMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityMultisig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"othenticRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"avsDirectoryContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowlistSigner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"avsName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"blsAuthSingleton\",\"type\":\"address\"}],\"internalType\":\"structIAvsGovernance.InitializationParams\",\"name\":\"_initializationParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"isFlowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEffectiveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"minSharesForStrategy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minVotingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfActiveOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numOfOperators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"numOfShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRewardsReceiver\",\"type\":\"address\"}],\"name\":\"queueRewardsReceiverModification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"_authToken\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_rewardsReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBLSAuthLibrary.Signature\",\"name\":\"_blsRegistrationSignature\",\"type\":\"tuple\"}],\"name\":\"registerAsAllowedOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_blsKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"_rewardsReceiver\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBLSAuthLibrary.Signature\",\"name\":\"_blsRegistrationSignature\",\"type\":\"tuple\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allowlistSigner\",\"type\":\"address\"}],\"name\":\"setAllowlistSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAvsGovernanceLogic\",\"name\":\"_avsGovernanceLogic\",\"type\":\"address\"}],\"name\":\"setAvsGovernanceLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAvsGovernanceMultiplierSyncer\",\"type\":\"address\"}],\"name\":\"setAvsGovernanceMultiplierSyncer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_avsName\",\"type\":\"string\"}],\"name\":\"setAvsName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blsAuthSingleton\",\"type\":\"address\"}],\"name\":\"setBLSAuthSingleton\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowlisted\",\"type\":\"bool\"}],\"name\":\"setIsAllowlisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxBalance\",\"type\":\"uint256\"}],\"name\":\"setMaxEffectiveBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minShares\",\"type\":\"uint256\"}],\"name\":\"setMinSharesForStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minVotingPower\",\"type\":\"uint256\"}],\"name\":\"setMinVotingPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLimitOfNumOfOperators\",\"type\":\"uint256\"}],\"name\":\"setNumOfOperatorsLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOthenticRegistry\",\"name\":\"_othenticRegistry\",\"type\":\"address\"}],\"name\":\"setOthenticRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsReceiverModificationDelay\",\"type\":\"uint256\"}],\"name\":\"setRewardsReceiverModificationDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsGovernance.StrategyMultiplier\",\"name\":\"_strategyMultiplier\",\"type\":\"tuple\"}],\"name\":\"setStrategyMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsGovernance.StrategyMultiplier[]\",\"name\":\"_strategyMultipliers\",\"type\":\"tuple[]\"}],\"name\":\"setStrategyMultiplierBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_strategies\",\"type\":\"address[]\"}],\"name\":\"setSupportedStrategies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategy\",\"type\":\"address\"}],\"name\":\"strategyMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAvsGovernanceMultisig\",\"type\":\"address\"}],\"name\":\"transferAvsGovernanceMultisig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMessageHandler\",\"type\":\"address\"}],\"name\":\"transferMessageHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_pausableFlow\",\"type\":\"bytes4\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"updateAVSMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"votingPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeToClaim\",\"type\":\"uint256\"}],\"internalType\":\"structIAvsGovernance.PaymentRequestMessage[]\",\"name\":\"_operators\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_lastPayedTask\",\"type\":\"uint256\"}],\"name\":\"withdrawBatchRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lastPayedTask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeToClaim\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AVSGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSGovernanceMetaData.ABI instead.
var AVSGovernanceABI = AVSGovernanceMetaData.ABI

// AVSGovernance is an auto generated Go binding around an Ethereum contract.
type AVSGovernance struct {
	AVSGovernanceCaller     // Read-only binding to the contract
	AVSGovernanceTransactor // Write-only binding to the contract
	AVSGovernanceFilterer   // Log filterer for contract events
}

// AVSGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSGovernanceSession struct {
	Contract     *AVSGovernance    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AVSGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSGovernanceCallerSession struct {
	Contract *AVSGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AVSGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSGovernanceTransactorSession struct {
	Contract     *AVSGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AVSGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSGovernanceRaw struct {
	Contract *AVSGovernance // Generic contract binding to access the raw methods on
}

// AVSGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSGovernanceCallerRaw struct {
	Contract *AVSGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// AVSGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSGovernanceTransactorRaw struct {
	Contract *AVSGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSGovernance creates a new instance of AVSGovernance, bound to a specific deployed contract.
func NewAVSGovernance(address common.Address, backend bind.ContractBackend) (*AVSGovernance, error) {
	contract, err := bindAVSGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSGovernance{AVSGovernanceCaller: AVSGovernanceCaller{contract: contract}, AVSGovernanceTransactor: AVSGovernanceTransactor{contract: contract}, AVSGovernanceFilterer: AVSGovernanceFilterer{contract: contract}}, nil
}

// NewAVSGovernanceCaller creates a new read-only instance of AVSGovernance, bound to a specific deployed contract.
func NewAVSGovernanceCaller(address common.Address, caller bind.ContractCaller) (*AVSGovernanceCaller, error) {
	contract, err := bindAVSGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceCaller{contract: contract}, nil
}

// NewAVSGovernanceTransactor creates a new write-only instance of AVSGovernance, bound to a specific deployed contract.
func NewAVSGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSGovernanceTransactor, error) {
	contract, err := bindAVSGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceTransactor{contract: contract}, nil
}

// NewAVSGovernanceFilterer creates a new log filterer instance of AVSGovernance, bound to a specific deployed contract.
func NewAVSGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSGovernanceFilterer, error) {
	contract, err := bindAVSGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceFilterer{contract: contract}, nil
}

// bindAVSGovernance binds a generic wrapper to an already deployed contract.
func bindAVSGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSGovernance *AVSGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSGovernance.Contract.AVSGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSGovernance *AVSGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSGovernance.Contract.AVSGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSGovernance *AVSGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSGovernance.Contract.AVSGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSGovernance *AVSGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSGovernance *AVSGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSGovernance *AVSGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSGovernance.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AVSGovernance *AVSGovernanceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AVSGovernance *AVSGovernanceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AVSGovernance.Contract.DEFAULTADMINROLE(&_AVSGovernance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AVSGovernance *AVSGovernanceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AVSGovernance.Contract.DEFAULTADMINROLE(&_AVSGovernance.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AVSGovernance *AVSGovernanceCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AVSGovernance *AVSGovernanceSession) AvsDirectory() (common.Address, error) {
	return _AVSGovernance.Contract.AvsDirectory(&_AVSGovernance.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_AVSGovernance *AVSGovernanceCallerSession) AvsDirectory() (common.Address, error) {
	return _AVSGovernance.Contract.AvsDirectory(&_AVSGovernance.CallOpts)
}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_AVSGovernance *AVSGovernanceCaller) AvsName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "avsName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_AVSGovernance *AVSGovernanceSession) AvsName() (string, error) {
	return _AVSGovernance.Contract.AvsName(&_AVSGovernance.CallOpts)
}

// AvsName is a free data retrieval call binding the contract method 0x41b92a29.
//
// Solidity: function avsName() view returns(string)
func (_AVSGovernance *AVSGovernanceCallerSession) AvsName() (string, error) {
	return _AVSGovernance.Contract.AvsName(&_AVSGovernance.CallOpts)
}

// GetDefaultStrategies is a free data retrieval call binding the contract method 0xe86685d9.
//
// Solidity: function getDefaultStrategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceCaller) GetDefaultStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "getDefaultStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDefaultStrategies is a free data retrieval call binding the contract method 0xe86685d9.
//
// Solidity: function getDefaultStrategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceSession) GetDefaultStrategies() ([]common.Address, error) {
	return _AVSGovernance.Contract.GetDefaultStrategies(&_AVSGovernance.CallOpts)
}

// GetDefaultStrategies is a free data retrieval call binding the contract method 0xe86685d9.
//
// Solidity: function getDefaultStrategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceCallerSession) GetDefaultStrategies() ([]common.Address, error) {
	return _AVSGovernance.Contract.GetDefaultStrategies(&_AVSGovernance.CallOpts)
}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_AVSGovernance *AVSGovernanceCaller) GetIsAllowlisted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "getIsAllowlisted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_AVSGovernance *AVSGovernanceSession) GetIsAllowlisted() (bool, error) {
	return _AVSGovernance.Contract.GetIsAllowlisted(&_AVSGovernance.CallOpts)
}

// GetIsAllowlisted is a free data retrieval call binding the contract method 0xb525fa88.
//
// Solidity: function getIsAllowlisted() view returns(bool)
func (_AVSGovernance *AVSGovernanceCallerSession) GetIsAllowlisted() (bool, error) {
	return _AVSGovernance.Contract.GetIsAllowlisted(&_AVSGovernance.CallOpts)
}

// GetNumOfOperatorsLimit is a free data retrieval call binding the contract method 0xf251c9a6.
//
// Solidity: function getNumOfOperatorsLimit() view returns(uint256 numOfOperatorsLimitView)
func (_AVSGovernance *AVSGovernanceCaller) GetNumOfOperatorsLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "getNumOfOperatorsLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumOfOperatorsLimit is a free data retrieval call binding the contract method 0xf251c9a6.
//
// Solidity: function getNumOfOperatorsLimit() view returns(uint256 numOfOperatorsLimitView)
func (_AVSGovernance *AVSGovernanceSession) GetNumOfOperatorsLimit() (*big.Int, error) {
	return _AVSGovernance.Contract.GetNumOfOperatorsLimit(&_AVSGovernance.CallOpts)
}

// GetNumOfOperatorsLimit is a free data retrieval call binding the contract method 0xf251c9a6.
//
// Solidity: function getNumOfOperatorsLimit() view returns(uint256 numOfOperatorsLimitView)
func (_AVSGovernance *AVSGovernanceCallerSession) GetNumOfOperatorsLimit() (*big.Int, error) {
	return _AVSGovernance.Contract.GetNumOfOperatorsLimit(&_AVSGovernance.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_AVSGovernance *AVSGovernanceCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, _operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "getOperatorRestakedStrategies", _operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_AVSGovernance *AVSGovernanceSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _AVSGovernance.Contract.GetOperatorRestakedStrategies(&_AVSGovernance.CallOpts, _operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address _operator) view returns(address[])
func (_AVSGovernance *AVSGovernanceCallerSession) GetOperatorRestakedStrategies(_operator common.Address) ([]common.Address, error) {
	return _AVSGovernance.Contract.GetOperatorRestakedStrategies(&_AVSGovernance.CallOpts, _operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _AVSGovernance.Contract.GetRestakeableStrategies(&_AVSGovernance.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _AVSGovernance.Contract.GetRestakeableStrategies(&_AVSGovernance.CallOpts)
}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_AVSGovernance *AVSGovernanceCaller) GetRewardsReceiver(opts *bind.CallOpts, _operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "getRewardsReceiver", _operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_AVSGovernance *AVSGovernanceSession) GetRewardsReceiver(_operator common.Address) (common.Address, error) {
	return _AVSGovernance.Contract.GetRewardsReceiver(&_AVSGovernance.CallOpts, _operator)
}

// GetRewardsReceiver is a free data retrieval call binding the contract method 0x5e95cee2.
//
// Solidity: function getRewardsReceiver(address _operator) view returns(address)
func (_AVSGovernance *AVSGovernanceCallerSession) GetRewardsReceiver(_operator common.Address) (common.Address, error) {
	return _AVSGovernance.Contract.GetRewardsReceiver(&_AVSGovernance.CallOpts, _operator)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AVSGovernance *AVSGovernanceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AVSGovernance *AVSGovernanceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AVSGovernance.Contract.GetRoleAdmin(&_AVSGovernance.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AVSGovernance *AVSGovernanceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AVSGovernance.Contract.GetRoleAdmin(&_AVSGovernance.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AVSGovernance *AVSGovernanceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AVSGovernance *AVSGovernanceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AVSGovernance.Contract.HasRole(&_AVSGovernance.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AVSGovernance *AVSGovernanceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AVSGovernance.Contract.HasRole(&_AVSGovernance.CallOpts, role, account)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_AVSGovernance *AVSGovernanceCaller) IsFlowPaused(opts *bind.CallOpts, _pausableFlow [4]byte) (bool, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "isFlowPaused", _pausableFlow)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_AVSGovernance *AVSGovernanceSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _AVSGovernance.Contract.IsFlowPaused(&_AVSGovernance.CallOpts, _pausableFlow)
}

// IsFlowPaused is a free data retrieval call binding the contract method 0xefd96978.
//
// Solidity: function isFlowPaused(bytes4 _pausableFlow) view returns(bool _isPaused)
func (_AVSGovernance *AVSGovernanceCallerSession) IsFlowPaused(_pausableFlow [4]byte) (bool, error) {
	return _AVSGovernance.Contract.IsFlowPaused(&_AVSGovernance.CallOpts, _pausableFlow)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_AVSGovernance *AVSGovernanceCaller) IsOperatorRegistered(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "isOperatorRegistered", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_AVSGovernance *AVSGovernanceSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _AVSGovernance.Contract.IsOperatorRegistered(&_AVSGovernance.CallOpts, operator)
}

// IsOperatorRegistered is a free data retrieval call binding the contract method 0x6b1906f8.
//
// Solidity: function isOperatorRegistered(address operator) view returns(bool)
func (_AVSGovernance *AVSGovernanceCallerSession) IsOperatorRegistered(operator common.Address) (bool, error) {
	return _AVSGovernance.Contract.IsOperatorRegistered(&_AVSGovernance.CallOpts, operator)
}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) MaxEffectiveBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "maxEffectiveBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) MaxEffectiveBalance() (*big.Int, error) {
	return _AVSGovernance.Contract.MaxEffectiveBalance(&_AVSGovernance.CallOpts)
}

// MaxEffectiveBalance is a free data retrieval call binding the contract method 0xa88171ee.
//
// Solidity: function maxEffectiveBalance() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) MaxEffectiveBalance() (*big.Int, error) {
	return _AVSGovernance.Contract.MaxEffectiveBalance(&_AVSGovernance.CallOpts)
}

// MinSharesForStrategy is a free data retrieval call binding the contract method 0xc3814e5b.
//
// Solidity: function minSharesForStrategy(address _strategy) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) MinSharesForStrategy(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "minSharesForStrategy", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSharesForStrategy is a free data retrieval call binding the contract method 0xc3814e5b.
//
// Solidity: function minSharesForStrategy(address _strategy) view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) MinSharesForStrategy(_strategy common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.MinSharesForStrategy(&_AVSGovernance.CallOpts, _strategy)
}

// MinSharesForStrategy is a free data retrieval call binding the contract method 0xc3814e5b.
//
// Solidity: function minSharesForStrategy(address _strategy) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) MinSharesForStrategy(_strategy common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.MinSharesForStrategy(&_AVSGovernance.CallOpts, _strategy)
}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) MinVotingPower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "minVotingPower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) MinVotingPower() (*big.Int, error) {
	return _AVSGovernance.Contract.MinVotingPower(&_AVSGovernance.CallOpts)
}

// MinVotingPower is a free data retrieval call binding the contract method 0x36fffde0.
//
// Solidity: function minVotingPower() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) MinVotingPower() (*big.Int, error) {
	return _AVSGovernance.Contract.MinVotingPower(&_AVSGovernance.CallOpts)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) NumOfActiveOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "numOfActiveOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) NumOfActiveOperators() (*big.Int, error) {
	return _AVSGovernance.Contract.NumOfActiveOperators(&_AVSGovernance.CallOpts)
}

// NumOfActiveOperators is a free data retrieval call binding the contract method 0x7897dec3.
//
// Solidity: function numOfActiveOperators() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) NumOfActiveOperators() (*big.Int, error) {
	return _AVSGovernance.Contract.NumOfActiveOperators(&_AVSGovernance.CallOpts)
}

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) NumOfOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "numOfOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) NumOfOperators() (*big.Int, error) {
	return _AVSGovernance.Contract.NumOfOperators(&_AVSGovernance.CallOpts)
}

// NumOfOperators is a free data retrieval call binding the contract method 0x6ade02da.
//
// Solidity: function numOfOperators() view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) NumOfOperators() (*big.Int, error) {
	return _AVSGovernance.Contract.NumOfOperators(&_AVSGovernance.CallOpts)
}

// NumOfShares is a free data retrieval call binding the contract method 0x6a907803.
//
// Solidity: function numOfShares(address _operator) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) NumOfShares(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "numOfShares", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumOfShares is a free data retrieval call binding the contract method 0x6a907803.
//
// Solidity: function numOfShares(address _operator) view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) NumOfShares(_operator common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.NumOfShares(&_AVSGovernance.CallOpts, _operator)
}

// NumOfShares is a free data retrieval call binding the contract method 0x6a907803.
//
// Solidity: function numOfShares(address _operator) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) NumOfShares(_operator common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.NumOfShares(&_AVSGovernance.CallOpts, _operator)
}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceCaller) Strategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "strategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceSession) Strategies() ([]common.Address, error) {
	return _AVSGovernance.Contract.Strategies(&_AVSGovernance.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0xd9f9027f.
//
// Solidity: function strategies() view returns(address[])
func (_AVSGovernance *AVSGovernanceCallerSession) Strategies() ([]common.Address, error) {
	return _AVSGovernance.Contract.Strategies(&_AVSGovernance.CallOpts)
}

// StrategyMultiplier is a free data retrieval call binding the contract method 0x8f53bc50.
//
// Solidity: function strategyMultiplier(address _strategy) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) StrategyMultiplier(opts *bind.CallOpts, _strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "strategyMultiplier", _strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyMultiplier is a free data retrieval call binding the contract method 0x8f53bc50.
//
// Solidity: function strategyMultiplier(address _strategy) view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) StrategyMultiplier(_strategy common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.StrategyMultiplier(&_AVSGovernance.CallOpts, _strategy)
}

// StrategyMultiplier is a free data retrieval call binding the contract method 0x8f53bc50.
//
// Solidity: function strategyMultiplier(address _strategy) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) StrategyMultiplier(_strategy common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.StrategyMultiplier(&_AVSGovernance.CallOpts, _strategy)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AVSGovernance *AVSGovernanceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AVSGovernance *AVSGovernanceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AVSGovernance.Contract.SupportsInterface(&_AVSGovernance.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AVSGovernance *AVSGovernanceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AVSGovernance.Contract.SupportsInterface(&_AVSGovernance.CallOpts, interfaceId)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_AVSGovernance *AVSGovernanceCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_AVSGovernance *AVSGovernanceSession) Vault() (common.Address, error) {
	return _AVSGovernance.Contract.Vault(&_AVSGovernance.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_AVSGovernance *AVSGovernanceCallerSession) Vault() (common.Address, error) {
	return _AVSGovernance.Contract.Vault(&_AVSGovernance.CallOpts)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCaller) VotingPower(opts *bind.CallOpts, _operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AVSGovernance.contract.Call(opts, &out, "votingPower", _operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_AVSGovernance *AVSGovernanceSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.VotingPower(&_AVSGovernance.CallOpts, _operator)
}

// VotingPower is a free data retrieval call binding the contract method 0xc07473f6.
//
// Solidity: function votingPower(address _operator) view returns(uint256)
func (_AVSGovernance *AVSGovernanceCallerSession) VotingPower(_operator common.Address) (*big.Int, error) {
	return _AVSGovernance.Contract.VotingPower(&_AVSGovernance.CallOpts, _operator)
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_AVSGovernance *AVSGovernanceTransactor) CompleteRewardsReceiverModification(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "completeRewardsReceiverModification")
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_AVSGovernance *AVSGovernanceSession) CompleteRewardsReceiverModification() (*types.Transaction, error) {
	return _AVSGovernance.Contract.CompleteRewardsReceiverModification(&_AVSGovernance.TransactOpts)
}

// CompleteRewardsReceiverModification is a paid mutator transaction binding the contract method 0xe6474b0f.
//
// Solidity: function completeRewardsReceiverModification() returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) CompleteRewardsReceiverModification() (*types.Transaction, error) {
	return _AVSGovernance.Contract.CompleteRewardsReceiverModification(&_AVSGovernance.TransactOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xb79092fd.
//
// Solidity: function depositERC20(uint256 _amount) returns()
func (_AVSGovernance *AVSGovernanceTransactor) DepositERC20(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "depositERC20", _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xb79092fd.
//
// Solidity: function depositERC20(uint256 _amount) returns()
func (_AVSGovernance *AVSGovernanceSession) DepositERC20(_amount *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.DepositERC20(&_AVSGovernance.TransactOpts, _amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xb79092fd.
//
// Solidity: function depositERC20(uint256 _amount) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) DepositERC20(_amount *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.DepositERC20(&_AVSGovernance.TransactOpts, _amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AVSGovernance *AVSGovernanceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AVSGovernance *AVSGovernanceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.GrantRole(&_AVSGovernance.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.GrantRole(&_AVSGovernance.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab57b8f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address) _initializationParams) returns()
func (_AVSGovernance *AVSGovernanceTransactor) Initialize(opts *bind.TransactOpts, _initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "initialize", _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab57b8f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address) _initializationParams) returns()
func (_AVSGovernance *AVSGovernanceSession) Initialize(_initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _AVSGovernance.Contract.Initialize(&_AVSGovernance.TransactOpts, _initializationParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab57b8f.
//
// Solidity: function initialize((address,address,address,address,address,address,address,address,string,address) _initializationParams) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) Initialize(_initializationParams IAvsGovernanceInitializationParams) (*types.Transaction, error) {
	return _AVSGovernance.Contract.Initialize(&_AVSGovernance.TransactOpts, _initializationParams)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_AVSGovernance *AVSGovernanceTransactor) Pause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "pause", _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_AVSGovernance *AVSGovernanceSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _AVSGovernance.Contract.Pause(&_AVSGovernance.TransactOpts, _pausableFlow)
}

// Pause is a paid mutator transaction binding the contract method 0x3aa83ec7.
//
// Solidity: function pause(bytes4 _pausableFlow) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) Pause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _AVSGovernance.Contract.Pause(&_AVSGovernance.TransactOpts, _pausableFlow)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_AVSGovernance *AVSGovernanceTransactor) QueueRewardsReceiverModification(opts *bind.TransactOpts, _newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "queueRewardsReceiverModification", _newRewardsReceiver)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_AVSGovernance *AVSGovernanceSession) QueueRewardsReceiverModification(_newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.QueueRewardsReceiverModification(&_AVSGovernance.TransactOpts, _newRewardsReceiver)
}

// QueueRewardsReceiverModification is a paid mutator transaction binding the contract method 0x1b21ba72.
//
// Solidity: function queueRewardsReceiverModification(address _newRewardsReceiver) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) QueueRewardsReceiverModification(_newRewardsReceiver common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.QueueRewardsReceiverModification(&_AVSGovernance.TransactOpts, _newRewardsReceiver)
}

// RegisterAsAllowedOperator is a paid mutator transaction binding the contract method 0x93304a9d.
//
// Solidity: function registerAsAllowedOperator(uint256[4] _blsKey, bytes _authToken, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_AVSGovernance *AVSGovernanceTransactor) RegisterAsAllowedOperator(opts *bind.TransactOpts, _blsKey [4]*big.Int, _authToken []byte, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "registerAsAllowedOperator", _blsKey, _authToken, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsAllowedOperator is a paid mutator transaction binding the contract method 0x93304a9d.
//
// Solidity: function registerAsAllowedOperator(uint256[4] _blsKey, bytes _authToken, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_AVSGovernance *AVSGovernanceSession) RegisterAsAllowedOperator(_blsKey [4]*big.Int, _authToken []byte, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RegisterAsAllowedOperator(&_AVSGovernance.TransactOpts, _blsKey, _authToken, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsAllowedOperator is a paid mutator transaction binding the contract method 0x93304a9d.
//
// Solidity: function registerAsAllowedOperator(uint256[4] _blsKey, bytes _authToken, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) RegisterAsAllowedOperator(_blsKey [4]*big.Int, _authToken []byte, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RegisterAsAllowedOperator(&_AVSGovernance.TransactOpts, _blsKey, _authToken, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x22609a4d.
//
// Solidity: function registerAsOperator(uint256[4] _blsKey, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_AVSGovernance *AVSGovernanceTransactor) RegisterAsOperator(opts *bind.TransactOpts, _blsKey [4]*big.Int, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "registerAsOperator", _blsKey, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x22609a4d.
//
// Solidity: function registerAsOperator(uint256[4] _blsKey, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_AVSGovernance *AVSGovernanceSession) RegisterAsOperator(_blsKey [4]*big.Int, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RegisterAsOperator(&_AVSGovernance.TransactOpts, _blsKey, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x22609a4d.
//
// Solidity: function registerAsOperator(uint256[4] _blsKey, address _rewardsReceiver, (bytes,bytes32,uint256) _operatorSignature, (uint256[2]) _blsRegistrationSignature) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) RegisterAsOperator(_blsKey [4]*big.Int, _rewardsReceiver common.Address, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry, _blsRegistrationSignature BLSAuthLibrarySignature) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RegisterAsOperator(&_AVSGovernance.TransactOpts, _blsKey, _rewardsReceiver, _operatorSignature, _blsRegistrationSignature)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AVSGovernance *AVSGovernanceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AVSGovernance *AVSGovernanceSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RenounceRole(&_AVSGovernance.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RenounceRole(&_AVSGovernance.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AVSGovernance *AVSGovernanceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AVSGovernance *AVSGovernanceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RevokeRole(&_AVSGovernance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.RevokeRole(&_AVSGovernance.TransactOpts, role, account)
}

// SetAllowlistSigner is a paid mutator transaction binding the contract method 0xe474def4.
//
// Solidity: function setAllowlistSigner(address _allowlistSigner) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetAllowlistSigner(opts *bind.TransactOpts, _allowlistSigner common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setAllowlistSigner", _allowlistSigner)
}

// SetAllowlistSigner is a paid mutator transaction binding the contract method 0xe474def4.
//
// Solidity: function setAllowlistSigner(address _allowlistSigner) returns()
func (_AVSGovernance *AVSGovernanceSession) SetAllowlistSigner(_allowlistSigner common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAllowlistSigner(&_AVSGovernance.TransactOpts, _allowlistSigner)
}

// SetAllowlistSigner is a paid mutator transaction binding the contract method 0xe474def4.
//
// Solidity: function setAllowlistSigner(address _allowlistSigner) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetAllowlistSigner(_allowlistSigner common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAllowlistSigner(&_AVSGovernance.TransactOpts, _allowlistSigner)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetAvsGovernanceLogic(opts *bind.TransactOpts, _avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setAvsGovernanceLogic", _avsGovernanceLogic)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_AVSGovernance *AVSGovernanceSession) SetAvsGovernanceLogic(_avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAvsGovernanceLogic(&_AVSGovernance.TransactOpts, _avsGovernanceLogic)
}

// SetAvsGovernanceLogic is a paid mutator transaction binding the contract method 0x8987c767.
//
// Solidity: function setAvsGovernanceLogic(address _avsGovernanceLogic) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetAvsGovernanceLogic(_avsGovernanceLogic common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAvsGovernanceLogic(&_AVSGovernance.TransactOpts, _avsGovernanceLogic)
}

// SetAvsGovernanceMultiplierSyncer is a paid mutator transaction binding the contract method 0x3425e8d8.
//
// Solidity: function setAvsGovernanceMultiplierSyncer(address _newAvsGovernanceMultiplierSyncer) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetAvsGovernanceMultiplierSyncer(opts *bind.TransactOpts, _newAvsGovernanceMultiplierSyncer common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setAvsGovernanceMultiplierSyncer", _newAvsGovernanceMultiplierSyncer)
}

// SetAvsGovernanceMultiplierSyncer is a paid mutator transaction binding the contract method 0x3425e8d8.
//
// Solidity: function setAvsGovernanceMultiplierSyncer(address _newAvsGovernanceMultiplierSyncer) returns()
func (_AVSGovernance *AVSGovernanceSession) SetAvsGovernanceMultiplierSyncer(_newAvsGovernanceMultiplierSyncer common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAvsGovernanceMultiplierSyncer(&_AVSGovernance.TransactOpts, _newAvsGovernanceMultiplierSyncer)
}

// SetAvsGovernanceMultiplierSyncer is a paid mutator transaction binding the contract method 0x3425e8d8.
//
// Solidity: function setAvsGovernanceMultiplierSyncer(address _newAvsGovernanceMultiplierSyncer) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetAvsGovernanceMultiplierSyncer(_newAvsGovernanceMultiplierSyncer common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAvsGovernanceMultiplierSyncer(&_AVSGovernance.TransactOpts, _newAvsGovernanceMultiplierSyncer)
}

// SetAvsName is a paid mutator transaction binding the contract method 0x7d38e926.
//
// Solidity: function setAvsName(string _avsName) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetAvsName(opts *bind.TransactOpts, _avsName string) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setAvsName", _avsName)
}

// SetAvsName is a paid mutator transaction binding the contract method 0x7d38e926.
//
// Solidity: function setAvsName(string _avsName) returns()
func (_AVSGovernance *AVSGovernanceSession) SetAvsName(_avsName string) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAvsName(&_AVSGovernance.TransactOpts, _avsName)
}

// SetAvsName is a paid mutator transaction binding the contract method 0x7d38e926.
//
// Solidity: function setAvsName(string _avsName) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetAvsName(_avsName string) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetAvsName(&_AVSGovernance.TransactOpts, _avsName)
}

// SetBLSAuthSingleton is a paid mutator transaction binding the contract method 0x4ef1476e.
//
// Solidity: function setBLSAuthSingleton(address _blsAuthSingleton) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetBLSAuthSingleton(opts *bind.TransactOpts, _blsAuthSingleton common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setBLSAuthSingleton", _blsAuthSingleton)
}

// SetBLSAuthSingleton is a paid mutator transaction binding the contract method 0x4ef1476e.
//
// Solidity: function setBLSAuthSingleton(address _blsAuthSingleton) returns()
func (_AVSGovernance *AVSGovernanceSession) SetBLSAuthSingleton(_blsAuthSingleton common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetBLSAuthSingleton(&_AVSGovernance.TransactOpts, _blsAuthSingleton)
}

// SetBLSAuthSingleton is a paid mutator transaction binding the contract method 0x4ef1476e.
//
// Solidity: function setBLSAuthSingleton(address _blsAuthSingleton) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetBLSAuthSingleton(_blsAuthSingleton common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetBLSAuthSingleton(&_AVSGovernance.TransactOpts, _blsAuthSingleton)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetIsAllowlisted(opts *bind.TransactOpts, _isAllowlisted bool) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setIsAllowlisted", _isAllowlisted)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_AVSGovernance *AVSGovernanceSession) SetIsAllowlisted(_isAllowlisted bool) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetIsAllowlisted(&_AVSGovernance.TransactOpts, _isAllowlisted)
}

// SetIsAllowlisted is a paid mutator transaction binding the contract method 0x9e965cc1.
//
// Solidity: function setIsAllowlisted(bool _isAllowlisted) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetIsAllowlisted(_isAllowlisted bool) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetIsAllowlisted(&_AVSGovernance.TransactOpts, _isAllowlisted)
}

// SetMaxEffectiveBalance is a paid mutator transaction binding the contract method 0x76086c70.
//
// Solidity: function setMaxEffectiveBalance(uint256 _maxBalance) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetMaxEffectiveBalance(opts *bind.TransactOpts, _maxBalance *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setMaxEffectiveBalance", _maxBalance)
}

// SetMaxEffectiveBalance is a paid mutator transaction binding the contract method 0x76086c70.
//
// Solidity: function setMaxEffectiveBalance(uint256 _maxBalance) returns()
func (_AVSGovernance *AVSGovernanceSession) SetMaxEffectiveBalance(_maxBalance *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetMaxEffectiveBalance(&_AVSGovernance.TransactOpts, _maxBalance)
}

// SetMaxEffectiveBalance is a paid mutator transaction binding the contract method 0x76086c70.
//
// Solidity: function setMaxEffectiveBalance(uint256 _maxBalance) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetMaxEffectiveBalance(_maxBalance *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetMaxEffectiveBalance(&_AVSGovernance.TransactOpts, _maxBalance)
}

// SetMinSharesForStrategy is a paid mutator transaction binding the contract method 0x305df58a.
//
// Solidity: function setMinSharesForStrategy(address _strategy, uint256 _minShares) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetMinSharesForStrategy(opts *bind.TransactOpts, _strategy common.Address, _minShares *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setMinSharesForStrategy", _strategy, _minShares)
}

// SetMinSharesForStrategy is a paid mutator transaction binding the contract method 0x305df58a.
//
// Solidity: function setMinSharesForStrategy(address _strategy, uint256 _minShares) returns()
func (_AVSGovernance *AVSGovernanceSession) SetMinSharesForStrategy(_strategy common.Address, _minShares *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetMinSharesForStrategy(&_AVSGovernance.TransactOpts, _strategy, _minShares)
}

// SetMinSharesForStrategy is a paid mutator transaction binding the contract method 0x305df58a.
//
// Solidity: function setMinSharesForStrategy(address _strategy, uint256 _minShares) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetMinSharesForStrategy(_strategy common.Address, _minShares *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetMinSharesForStrategy(&_AVSGovernance.TransactOpts, _strategy, _minShares)
}

// SetMinVotingPower is a paid mutator transaction binding the contract method 0x55e48918.
//
// Solidity: function setMinVotingPower(uint256 _minVotingPower) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetMinVotingPower(opts *bind.TransactOpts, _minVotingPower *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setMinVotingPower", _minVotingPower)
}

// SetMinVotingPower is a paid mutator transaction binding the contract method 0x55e48918.
//
// Solidity: function setMinVotingPower(uint256 _minVotingPower) returns()
func (_AVSGovernance *AVSGovernanceSession) SetMinVotingPower(_minVotingPower *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetMinVotingPower(&_AVSGovernance.TransactOpts, _minVotingPower)
}

// SetMinVotingPower is a paid mutator transaction binding the contract method 0x55e48918.
//
// Solidity: function setMinVotingPower(uint256 _minVotingPower) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetMinVotingPower(_minVotingPower *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetMinVotingPower(&_AVSGovernance.TransactOpts, _minVotingPower)
}

// SetNumOfOperatorsLimit is a paid mutator transaction binding the contract method 0x9d79e4a7.
//
// Solidity: function setNumOfOperatorsLimit(uint256 _newLimitOfNumOfOperators) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetNumOfOperatorsLimit(opts *bind.TransactOpts, _newLimitOfNumOfOperators *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setNumOfOperatorsLimit", _newLimitOfNumOfOperators)
}

// SetNumOfOperatorsLimit is a paid mutator transaction binding the contract method 0x9d79e4a7.
//
// Solidity: function setNumOfOperatorsLimit(uint256 _newLimitOfNumOfOperators) returns()
func (_AVSGovernance *AVSGovernanceSession) SetNumOfOperatorsLimit(_newLimitOfNumOfOperators *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetNumOfOperatorsLimit(&_AVSGovernance.TransactOpts, _newLimitOfNumOfOperators)
}

// SetNumOfOperatorsLimit is a paid mutator transaction binding the contract method 0x9d79e4a7.
//
// Solidity: function setNumOfOperatorsLimit(uint256 _newLimitOfNumOfOperators) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetNumOfOperatorsLimit(_newLimitOfNumOfOperators *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetNumOfOperatorsLimit(&_AVSGovernance.TransactOpts, _newLimitOfNumOfOperators)
}

// SetOthenticRegistry is a paid mutator transaction binding the contract method 0x45a022fa.
//
// Solidity: function setOthenticRegistry(address _othenticRegistry) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetOthenticRegistry(opts *bind.TransactOpts, _othenticRegistry common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setOthenticRegistry", _othenticRegistry)
}

// SetOthenticRegistry is a paid mutator transaction binding the contract method 0x45a022fa.
//
// Solidity: function setOthenticRegistry(address _othenticRegistry) returns()
func (_AVSGovernance *AVSGovernanceSession) SetOthenticRegistry(_othenticRegistry common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetOthenticRegistry(&_AVSGovernance.TransactOpts, _othenticRegistry)
}

// SetOthenticRegistry is a paid mutator transaction binding the contract method 0x45a022fa.
//
// Solidity: function setOthenticRegistry(address _othenticRegistry) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetOthenticRegistry(_othenticRegistry common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetOthenticRegistry(&_AVSGovernance.TransactOpts, _othenticRegistry)
}

// SetRewardsReceiverModificationDelay is a paid mutator transaction binding the contract method 0x8a70469a.
//
// Solidity: function setRewardsReceiverModificationDelay(uint256 _rewardsReceiverModificationDelay) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetRewardsReceiverModificationDelay(opts *bind.TransactOpts, _rewardsReceiverModificationDelay *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setRewardsReceiverModificationDelay", _rewardsReceiverModificationDelay)
}

// SetRewardsReceiverModificationDelay is a paid mutator transaction binding the contract method 0x8a70469a.
//
// Solidity: function setRewardsReceiverModificationDelay(uint256 _rewardsReceiverModificationDelay) returns()
func (_AVSGovernance *AVSGovernanceSession) SetRewardsReceiverModificationDelay(_rewardsReceiverModificationDelay *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetRewardsReceiverModificationDelay(&_AVSGovernance.TransactOpts, _rewardsReceiverModificationDelay)
}

// SetRewardsReceiverModificationDelay is a paid mutator transaction binding the contract method 0x8a70469a.
//
// Solidity: function setRewardsReceiverModificationDelay(uint256 _rewardsReceiverModificationDelay) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetRewardsReceiverModificationDelay(_rewardsReceiverModificationDelay *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetRewardsReceiverModificationDelay(&_AVSGovernance.TransactOpts, _rewardsReceiverModificationDelay)
}

// SetStrategyMultiplier is a paid mutator transaction binding the contract method 0x076400d5.
//
// Solidity: function setStrategyMultiplier((address,uint256) _strategyMultiplier) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetStrategyMultiplier(opts *bind.TransactOpts, _strategyMultiplier IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setStrategyMultiplier", _strategyMultiplier)
}

// SetStrategyMultiplier is a paid mutator transaction binding the contract method 0x076400d5.
//
// Solidity: function setStrategyMultiplier((address,uint256) _strategyMultiplier) returns()
func (_AVSGovernance *AVSGovernanceSession) SetStrategyMultiplier(_strategyMultiplier IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetStrategyMultiplier(&_AVSGovernance.TransactOpts, _strategyMultiplier)
}

// SetStrategyMultiplier is a paid mutator transaction binding the contract method 0x076400d5.
//
// Solidity: function setStrategyMultiplier((address,uint256) _strategyMultiplier) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetStrategyMultiplier(_strategyMultiplier IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetStrategyMultiplier(&_AVSGovernance.TransactOpts, _strategyMultiplier)
}

// SetStrategyMultiplierBatch is a paid mutator transaction binding the contract method 0xd94a2e1d.
//
// Solidity: function setStrategyMultiplierBatch((address,uint256)[] _strategyMultipliers) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetStrategyMultiplierBatch(opts *bind.TransactOpts, _strategyMultipliers []IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setStrategyMultiplierBatch", _strategyMultipliers)
}

// SetStrategyMultiplierBatch is a paid mutator transaction binding the contract method 0xd94a2e1d.
//
// Solidity: function setStrategyMultiplierBatch((address,uint256)[] _strategyMultipliers) returns()
func (_AVSGovernance *AVSGovernanceSession) SetStrategyMultiplierBatch(_strategyMultipliers []IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetStrategyMultiplierBatch(&_AVSGovernance.TransactOpts, _strategyMultipliers)
}

// SetStrategyMultiplierBatch is a paid mutator transaction binding the contract method 0xd94a2e1d.
//
// Solidity: function setStrategyMultiplierBatch((address,uint256)[] _strategyMultipliers) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetStrategyMultiplierBatch(_strategyMultipliers []IAvsGovernanceStrategyMultiplier) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetStrategyMultiplierBatch(&_AVSGovernance.TransactOpts, _strategyMultipliers)
}

// SetSupportedStrategies is a paid mutator transaction binding the contract method 0x312c150b.
//
// Solidity: function setSupportedStrategies(address[] _strategies) returns()
func (_AVSGovernance *AVSGovernanceTransactor) SetSupportedStrategies(opts *bind.TransactOpts, _strategies []common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "setSupportedStrategies", _strategies)
}

// SetSupportedStrategies is a paid mutator transaction binding the contract method 0x312c150b.
//
// Solidity: function setSupportedStrategies(address[] _strategies) returns()
func (_AVSGovernance *AVSGovernanceSession) SetSupportedStrategies(_strategies []common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetSupportedStrategies(&_AVSGovernance.TransactOpts, _strategies)
}

// SetSupportedStrategies is a paid mutator transaction binding the contract method 0x312c150b.
//
// Solidity: function setSupportedStrategies(address[] _strategies) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) SetSupportedStrategies(_strategies []common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.SetSupportedStrategies(&_AVSGovernance.TransactOpts, _strategies)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_AVSGovernance *AVSGovernanceTransactor) TransferAvsGovernanceMultisig(opts *bind.TransactOpts, _newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "transferAvsGovernanceMultisig", _newAvsGovernanceMultisig)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_AVSGovernance *AVSGovernanceSession) TransferAvsGovernanceMultisig(_newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.TransferAvsGovernanceMultisig(&_AVSGovernance.TransactOpts, _newAvsGovernanceMultisig)
}

// TransferAvsGovernanceMultisig is a paid mutator transaction binding the contract method 0x513c52ba.
//
// Solidity: function transferAvsGovernanceMultisig(address _newAvsGovernanceMultisig) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) TransferAvsGovernanceMultisig(_newAvsGovernanceMultisig common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.TransferAvsGovernanceMultisig(&_AVSGovernance.TransactOpts, _newAvsGovernanceMultisig)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_AVSGovernance *AVSGovernanceTransactor) TransferMessageHandler(opts *bind.TransactOpts, _newMessageHandler common.Address) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "transferMessageHandler", _newMessageHandler)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_AVSGovernance *AVSGovernanceSession) TransferMessageHandler(_newMessageHandler common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.TransferMessageHandler(&_AVSGovernance.TransactOpts, _newMessageHandler)
}

// TransferMessageHandler is a paid mutator transaction binding the contract method 0x4d07f651.
//
// Solidity: function transferMessageHandler(address _newMessageHandler) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) TransferMessageHandler(_newMessageHandler common.Address) (*types.Transaction, error) {
	return _AVSGovernance.Contract.TransferMessageHandler(&_AVSGovernance.TransactOpts, _newMessageHandler)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_AVSGovernance *AVSGovernanceTransactor) Unpause(opts *bind.TransactOpts, _pausableFlow [4]byte) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "unpause", _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_AVSGovernance *AVSGovernanceSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _AVSGovernance.Contract.Unpause(&_AVSGovernance.TransactOpts, _pausableFlow)
}

// Unpause is a paid mutator transaction binding the contract method 0xbac1e94b.
//
// Solidity: function unpause(bytes4 _pausableFlow) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) Unpause(_pausableFlow [4]byte) (*types.Transaction, error) {
	return _AVSGovernance.Contract.Unpause(&_AVSGovernance.TransactOpts, _pausableFlow)
}

// UnregisterAsOperator is a paid mutator transaction binding the contract method 0x09869442.
//
// Solidity: function unregisterAsOperator() returns()
func (_AVSGovernance *AVSGovernanceTransactor) UnregisterAsOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "unregisterAsOperator")
}

// UnregisterAsOperator is a paid mutator transaction binding the contract method 0x09869442.
//
// Solidity: function unregisterAsOperator() returns()
func (_AVSGovernance *AVSGovernanceSession) UnregisterAsOperator() (*types.Transaction, error) {
	return _AVSGovernance.Contract.UnregisterAsOperator(&_AVSGovernance.TransactOpts)
}

// UnregisterAsOperator is a paid mutator transaction binding the contract method 0x09869442.
//
// Solidity: function unregisterAsOperator() returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) UnregisterAsOperator() (*types.Transaction, error) {
	return _AVSGovernance.Contract.UnregisterAsOperator(&_AVSGovernance.TransactOpts)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSGovernance *AVSGovernanceTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSGovernance *AVSGovernanceSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSGovernance.Contract.UpdateAVSMetadataURI(&_AVSGovernance.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _AVSGovernance.Contract.UpdateAVSMetadataURI(&_AVSGovernance.TransactOpts, metadataURI)
}

// WithdrawBatchRewards is a paid mutator transaction binding the contract method 0xbc8be0c8.
//
// Solidity: function withdrawBatchRewards((address,uint256)[] _operators, uint256 _lastPayedTask) returns()
func (_AVSGovernance *AVSGovernanceTransactor) WithdrawBatchRewards(opts *bind.TransactOpts, _operators []IAvsGovernancePaymentRequestMessage, _lastPayedTask *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "withdrawBatchRewards", _operators, _lastPayedTask)
}

// WithdrawBatchRewards is a paid mutator transaction binding the contract method 0xbc8be0c8.
//
// Solidity: function withdrawBatchRewards((address,uint256)[] _operators, uint256 _lastPayedTask) returns()
func (_AVSGovernance *AVSGovernanceSession) WithdrawBatchRewards(_operators []IAvsGovernancePaymentRequestMessage, _lastPayedTask *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.WithdrawBatchRewards(&_AVSGovernance.TransactOpts, _operators, _lastPayedTask)
}

// WithdrawBatchRewards is a paid mutator transaction binding the contract method 0xbc8be0c8.
//
// Solidity: function withdrawBatchRewards((address,uint256)[] _operators, uint256 _lastPayedTask) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) WithdrawBatchRewards(_operators []IAvsGovernancePaymentRequestMessage, _lastPayedTask *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.WithdrawBatchRewards(&_AVSGovernance.TransactOpts, _operators, _lastPayedTask)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x3256b4d1.
//
// Solidity: function withdrawRewards(address _operator, uint256 _lastPayedTask, uint256 _feeToClaim) returns()
func (_AVSGovernance *AVSGovernanceTransactor) WithdrawRewards(opts *bind.TransactOpts, _operator common.Address, _lastPayedTask *big.Int, _feeToClaim *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.contract.Transact(opts, "withdrawRewards", _operator, _lastPayedTask, _feeToClaim)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x3256b4d1.
//
// Solidity: function withdrawRewards(address _operator, uint256 _lastPayedTask, uint256 _feeToClaim) returns()
func (_AVSGovernance *AVSGovernanceSession) WithdrawRewards(_operator common.Address, _lastPayedTask *big.Int, _feeToClaim *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.WithdrawRewards(&_AVSGovernance.TransactOpts, _operator, _lastPayedTask, _feeToClaim)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0x3256b4d1.
//
// Solidity: function withdrawRewards(address _operator, uint256 _lastPayedTask, uint256 _feeToClaim) returns()
func (_AVSGovernance *AVSGovernanceTransactorSession) WithdrawRewards(_operator common.Address, _lastPayedTask *big.Int, _feeToClaim *big.Int) (*types.Transaction, error) {
	return _AVSGovernance.Contract.WithdrawRewards(&_AVSGovernance.TransactOpts, _operator, _lastPayedTask, _feeToClaim)
}

// AVSGovernanceBLSAuthSingletonSetIterator is returned from FilterBLSAuthSingletonSet and is used to iterate over the raw logs and unpacked data for BLSAuthSingletonSet events raised by the AVSGovernance contract.
type AVSGovernanceBLSAuthSingletonSetIterator struct {
	Event *AVSGovernanceBLSAuthSingletonSet // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceBLSAuthSingletonSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceBLSAuthSingletonSet)
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
		it.Event = new(AVSGovernanceBLSAuthSingletonSet)
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
func (it *AVSGovernanceBLSAuthSingletonSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceBLSAuthSingletonSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceBLSAuthSingletonSet represents a BLSAuthSingletonSet event raised by the AVSGovernance contract.
type AVSGovernanceBLSAuthSingletonSet struct {
	BlsAuthSingleton common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBLSAuthSingletonSet is a free log retrieval operation binding the contract event 0x4cbffdecf3b5e4b22bfb2bdec99a66f8fcf81e19b060682afd9645c729da1472.
//
// Solidity: event BLSAuthSingletonSet(address blsAuthSingleton)
func (_AVSGovernance *AVSGovernanceFilterer) FilterBLSAuthSingletonSet(opts *bind.FilterOpts) (*AVSGovernanceBLSAuthSingletonSetIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "BLSAuthSingletonSet")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceBLSAuthSingletonSetIterator{contract: _AVSGovernance.contract, event: "BLSAuthSingletonSet", logs: logs, sub: sub}, nil
}

// WatchBLSAuthSingletonSet is a free log subscription operation binding the contract event 0x4cbffdecf3b5e4b22bfb2bdec99a66f8fcf81e19b060682afd9645c729da1472.
//
// Solidity: event BLSAuthSingletonSet(address blsAuthSingleton)
func (_AVSGovernance *AVSGovernanceFilterer) WatchBLSAuthSingletonSet(opts *bind.WatchOpts, sink chan<- *AVSGovernanceBLSAuthSingletonSet) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "BLSAuthSingletonSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceBLSAuthSingletonSet)
				if err := _AVSGovernance.contract.UnpackLog(event, "BLSAuthSingletonSet", log); err != nil {
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

// ParseBLSAuthSingletonSet is a log parse operation binding the contract event 0x4cbffdecf3b5e4b22bfb2bdec99a66f8fcf81e19b060682afd9645c729da1472.
//
// Solidity: event BLSAuthSingletonSet(address blsAuthSingleton)
func (_AVSGovernance *AVSGovernanceFilterer) ParseBLSAuthSingletonSet(log types.Log) (*AVSGovernanceBLSAuthSingletonSet, error) {
	event := new(AVSGovernanceBLSAuthSingletonSet)
	if err := _AVSGovernance.contract.UnpackLog(event, "BLSAuthSingletonSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceFlowPausedIterator is returned from FilterFlowPaused and is used to iterate over the raw logs and unpacked data for FlowPaused events raised by the AVSGovernance contract.
type AVSGovernanceFlowPausedIterator struct {
	Event *AVSGovernanceFlowPaused // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceFlowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceFlowPaused)
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
		it.Event = new(AVSGovernanceFlowPaused)
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
func (it *AVSGovernanceFlowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceFlowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceFlowPaused represents a FlowPaused event raised by the AVSGovernance contract.
type AVSGovernanceFlowPaused struct {
	PausableFlow [4]byte
	Pauser       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlowPaused is a free log retrieval operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_AVSGovernance *AVSGovernanceFilterer) FilterFlowPaused(opts *bind.FilterOpts) (*AVSGovernanceFlowPausedIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceFlowPausedIterator{contract: _AVSGovernance.contract, event: "FlowPaused", logs: logs, sub: sub}, nil
}

// WatchFlowPaused is a free log subscription operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_AVSGovernance *AVSGovernanceFilterer) WatchFlowPaused(opts *bind.WatchOpts, sink chan<- *AVSGovernanceFlowPaused) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "FlowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceFlowPaused)
				if err := _AVSGovernance.contract.UnpackLog(event, "FlowPaused", log); err != nil {
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

// ParseFlowPaused is a log parse operation binding the contract event 0x95c3658c5e0c74e20cf12db371b9b67d26e97a1937f6d2284f88cc44d036b4f6.
//
// Solidity: event FlowPaused(bytes4 _pausableFlow, address _pauser)
func (_AVSGovernance *AVSGovernanceFilterer) ParseFlowPaused(log types.Log) (*AVSGovernanceFlowPaused, error) {
	event := new(AVSGovernanceFlowPaused)
	if err := _AVSGovernance.contract.UnpackLog(event, "FlowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceFlowUnpausedIterator is returned from FilterFlowUnpaused and is used to iterate over the raw logs and unpacked data for FlowUnpaused events raised by the AVSGovernance contract.
type AVSGovernanceFlowUnpausedIterator struct {
	Event *AVSGovernanceFlowUnpaused // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceFlowUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceFlowUnpaused)
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
		it.Event = new(AVSGovernanceFlowUnpaused)
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
func (it *AVSGovernanceFlowUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceFlowUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceFlowUnpaused represents a FlowUnpaused event raised by the AVSGovernance contract.
type AVSGovernanceFlowUnpaused struct {
	PausableFlowFlag [4]byte
	Unpauser         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlowUnpaused is a free log retrieval operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_AVSGovernance *AVSGovernanceFilterer) FilterFlowUnpaused(opts *bind.FilterOpts) (*AVSGovernanceFlowUnpausedIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceFlowUnpausedIterator{contract: _AVSGovernance.contract, event: "FlowUnpaused", logs: logs, sub: sub}, nil
}

// WatchFlowUnpaused is a free log subscription operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_AVSGovernance *AVSGovernanceFilterer) WatchFlowUnpaused(opts *bind.WatchOpts, sink chan<- *AVSGovernanceFlowUnpaused) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "FlowUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceFlowUnpaused)
				if err := _AVSGovernance.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
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

// ParseFlowUnpaused is a log parse operation binding the contract event 0xc7e56e17b0a6c4b467df6495e1eda1baecd7ba20604e80c1058ac06f4578d85e.
//
// Solidity: event FlowUnpaused(bytes4 _pausableFlowFlag, address _unpauser)
func (_AVSGovernance *AVSGovernanceFilterer) ParseFlowUnpaused(log types.Log) (*AVSGovernanceFlowUnpaused, error) {
	event := new(AVSGovernanceFlowUnpaused)
	if err := _AVSGovernance.contract.UnpackLog(event, "FlowUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AVSGovernance contract.
type AVSGovernanceInitializedIterator struct {
	Event *AVSGovernanceInitialized // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceInitialized)
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
		it.Event = new(AVSGovernanceInitialized)
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
func (it *AVSGovernanceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceInitialized represents a Initialized event raised by the AVSGovernance contract.
type AVSGovernanceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AVSGovernance *AVSGovernanceFilterer) FilterInitialized(opts *bind.FilterOpts) (*AVSGovernanceInitializedIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceInitializedIterator{contract: _AVSGovernance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AVSGovernance *AVSGovernanceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AVSGovernanceInitialized) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceInitialized)
				if err := _AVSGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AVSGovernance *AVSGovernanceFilterer) ParseInitialized(log types.Log) (*AVSGovernanceInitialized, error) {
	event := new(AVSGovernanceInitialized)
	if err := _AVSGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceMaxEffectiveBalanceSetIterator is returned from FilterMaxEffectiveBalanceSet and is used to iterate over the raw logs and unpacked data for MaxEffectiveBalanceSet events raised by the AVSGovernance contract.
type AVSGovernanceMaxEffectiveBalanceSetIterator struct {
	Event *AVSGovernanceMaxEffectiveBalanceSet // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceMaxEffectiveBalanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceMaxEffectiveBalanceSet)
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
		it.Event = new(AVSGovernanceMaxEffectiveBalanceSet)
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
func (it *AVSGovernanceMaxEffectiveBalanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceMaxEffectiveBalanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceMaxEffectiveBalanceSet represents a MaxEffectiveBalanceSet event raised by the AVSGovernance contract.
type AVSGovernanceMaxEffectiveBalanceSet struct {
	MaxEffectiveBalance *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMaxEffectiveBalanceSet is a free log retrieval operation binding the contract event 0x00c6fb6db9c52d89a1eaf84e0470a3304db2086d0ac44d64ebf4ea35a905a7d0.
//
// Solidity: event MaxEffectiveBalanceSet(uint256 maxEffectiveBalance)
func (_AVSGovernance *AVSGovernanceFilterer) FilterMaxEffectiveBalanceSet(opts *bind.FilterOpts) (*AVSGovernanceMaxEffectiveBalanceSetIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "MaxEffectiveBalanceSet")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceMaxEffectiveBalanceSetIterator{contract: _AVSGovernance.contract, event: "MaxEffectiveBalanceSet", logs: logs, sub: sub}, nil
}

// WatchMaxEffectiveBalanceSet is a free log subscription operation binding the contract event 0x00c6fb6db9c52d89a1eaf84e0470a3304db2086d0ac44d64ebf4ea35a905a7d0.
//
// Solidity: event MaxEffectiveBalanceSet(uint256 maxEffectiveBalance)
func (_AVSGovernance *AVSGovernanceFilterer) WatchMaxEffectiveBalanceSet(opts *bind.WatchOpts, sink chan<- *AVSGovernanceMaxEffectiveBalanceSet) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "MaxEffectiveBalanceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceMaxEffectiveBalanceSet)
				if err := _AVSGovernance.contract.UnpackLog(event, "MaxEffectiveBalanceSet", log); err != nil {
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

// ParseMaxEffectiveBalanceSet is a log parse operation binding the contract event 0x00c6fb6db9c52d89a1eaf84e0470a3304db2086d0ac44d64ebf4ea35a905a7d0.
//
// Solidity: event MaxEffectiveBalanceSet(uint256 maxEffectiveBalance)
func (_AVSGovernance *AVSGovernanceFilterer) ParseMaxEffectiveBalanceSet(log types.Log) (*AVSGovernanceMaxEffectiveBalanceSet, error) {
	event := new(AVSGovernanceMaxEffectiveBalanceSet)
	if err := _AVSGovernance.contract.UnpackLog(event, "MaxEffectiveBalanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceMinSharesPerStrategySetIterator is returned from FilterMinSharesPerStrategySet and is used to iterate over the raw logs and unpacked data for MinSharesPerStrategySet events raised by the AVSGovernance contract.
type AVSGovernanceMinSharesPerStrategySetIterator struct {
	Event *AVSGovernanceMinSharesPerStrategySet // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceMinSharesPerStrategySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceMinSharesPerStrategySet)
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
		it.Event = new(AVSGovernanceMinSharesPerStrategySet)
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
func (it *AVSGovernanceMinSharesPerStrategySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceMinSharesPerStrategySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceMinSharesPerStrategySet represents a MinSharesPerStrategySet event raised by the AVSGovernance contract.
type AVSGovernanceMinSharesPerStrategySet struct {
	Strategy  common.Address
	MinShares *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinSharesPerStrategySet is a free log retrieval operation binding the contract event 0x3a6c52328a7b3b726d0ec757d68f416b26ec2991ac4d4f95d450c504f5a0e521.
//
// Solidity: event MinSharesPerStrategySet(address strategy, uint256 minShares)
func (_AVSGovernance *AVSGovernanceFilterer) FilterMinSharesPerStrategySet(opts *bind.FilterOpts) (*AVSGovernanceMinSharesPerStrategySetIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "MinSharesPerStrategySet")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceMinSharesPerStrategySetIterator{contract: _AVSGovernance.contract, event: "MinSharesPerStrategySet", logs: logs, sub: sub}, nil
}

// WatchMinSharesPerStrategySet is a free log subscription operation binding the contract event 0x3a6c52328a7b3b726d0ec757d68f416b26ec2991ac4d4f95d450c504f5a0e521.
//
// Solidity: event MinSharesPerStrategySet(address strategy, uint256 minShares)
func (_AVSGovernance *AVSGovernanceFilterer) WatchMinSharesPerStrategySet(opts *bind.WatchOpts, sink chan<- *AVSGovernanceMinSharesPerStrategySet) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "MinSharesPerStrategySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceMinSharesPerStrategySet)
				if err := _AVSGovernance.contract.UnpackLog(event, "MinSharesPerStrategySet", log); err != nil {
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

// ParseMinSharesPerStrategySet is a log parse operation binding the contract event 0x3a6c52328a7b3b726d0ec757d68f416b26ec2991ac4d4f95d450c504f5a0e521.
//
// Solidity: event MinSharesPerStrategySet(address strategy, uint256 minShares)
func (_AVSGovernance *AVSGovernanceFilterer) ParseMinSharesPerStrategySet(log types.Log) (*AVSGovernanceMinSharesPerStrategySet, error) {
	event := new(AVSGovernanceMinSharesPerStrategySet)
	if err := _AVSGovernance.contract.UnpackLog(event, "MinSharesPerStrategySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceMinVotingPowerSetIterator is returned from FilterMinVotingPowerSet and is used to iterate over the raw logs and unpacked data for MinVotingPowerSet events raised by the AVSGovernance contract.
type AVSGovernanceMinVotingPowerSetIterator struct {
	Event *AVSGovernanceMinVotingPowerSet // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceMinVotingPowerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceMinVotingPowerSet)
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
		it.Event = new(AVSGovernanceMinVotingPowerSet)
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
func (it *AVSGovernanceMinVotingPowerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceMinVotingPowerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceMinVotingPowerSet represents a MinVotingPowerSet event raised by the AVSGovernance contract.
type AVSGovernanceMinVotingPowerSet struct {
	MinVotingPower *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinVotingPowerSet is a free log retrieval operation binding the contract event 0x10203ddc048c86cf14172a6ea2565c805ce7320b22d6941b2eb396d0ee077983.
//
// Solidity: event MinVotingPowerSet(uint256 minVotingPower)
func (_AVSGovernance *AVSGovernanceFilterer) FilterMinVotingPowerSet(opts *bind.FilterOpts) (*AVSGovernanceMinVotingPowerSetIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "MinVotingPowerSet")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceMinVotingPowerSetIterator{contract: _AVSGovernance.contract, event: "MinVotingPowerSet", logs: logs, sub: sub}, nil
}

// WatchMinVotingPowerSet is a free log subscription operation binding the contract event 0x10203ddc048c86cf14172a6ea2565c805ce7320b22d6941b2eb396d0ee077983.
//
// Solidity: event MinVotingPowerSet(uint256 minVotingPower)
func (_AVSGovernance *AVSGovernanceFilterer) WatchMinVotingPowerSet(opts *bind.WatchOpts, sink chan<- *AVSGovernanceMinVotingPowerSet) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "MinVotingPowerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceMinVotingPowerSet)
				if err := _AVSGovernance.contract.UnpackLog(event, "MinVotingPowerSet", log); err != nil {
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

// ParseMinVotingPowerSet is a log parse operation binding the contract event 0x10203ddc048c86cf14172a6ea2565c805ce7320b22d6941b2eb396d0ee077983.
//
// Solidity: event MinVotingPowerSet(uint256 minVotingPower)
func (_AVSGovernance *AVSGovernanceFilterer) ParseMinVotingPowerSet(log types.Log) (*AVSGovernanceMinVotingPowerSet, error) {
	event := new(AVSGovernanceMinVotingPowerSet)
	if err := _AVSGovernance.contract.UnpackLog(event, "MinVotingPowerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the AVSGovernance contract.
type AVSGovernanceOperatorRegisteredIterator struct {
	Event *AVSGovernanceOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceOperatorRegistered)
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
		it.Event = new(AVSGovernanceOperatorRegistered)
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
func (it *AVSGovernanceOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceOperatorRegistered represents a OperatorRegistered event raised by the AVSGovernance contract.
type AVSGovernanceOperatorRegistered struct {
	Operator common.Address
	BlsKey   [4]*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac055.
//
// Solidity: event OperatorRegistered(address indexed operator, uint256[4] blsKey)
func (_AVSGovernance *AVSGovernanceFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*AVSGovernanceOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceOperatorRegisteredIterator{contract: _AVSGovernance.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac055.
//
// Solidity: event OperatorRegistered(address indexed operator, uint256[4] blsKey)
func (_AVSGovernance *AVSGovernanceFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *AVSGovernanceOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceOperatorRegistered)
				if err := _AVSGovernance.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x54bc9cf83c2eb0f2ad1abf6e4fab882964404622ba2df6b5a9356a18d3aac055.
//
// Solidity: event OperatorRegistered(address indexed operator, uint256[4] blsKey)
func (_AVSGovernance *AVSGovernanceFilterer) ParseOperatorRegistered(log types.Log) (*AVSGovernanceOperatorRegistered, error) {
	event := new(AVSGovernanceOperatorRegistered)
	if err := _AVSGovernance.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceOperatorUnregisteredIterator is returned from FilterOperatorUnregistered and is used to iterate over the raw logs and unpacked data for OperatorUnregistered events raised by the AVSGovernance contract.
type AVSGovernanceOperatorUnregisteredIterator struct {
	Event *AVSGovernanceOperatorUnregistered // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceOperatorUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceOperatorUnregistered)
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
		it.Event = new(AVSGovernanceOperatorUnregistered)
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
func (it *AVSGovernanceOperatorUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceOperatorUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceOperatorUnregistered represents a OperatorUnregistered event raised by the AVSGovernance contract.
type AVSGovernanceOperatorUnregistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUnregistered is a free log retrieval operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_AVSGovernance *AVSGovernanceFilterer) FilterOperatorUnregistered(opts *bind.FilterOpts) (*AVSGovernanceOperatorUnregisteredIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "OperatorUnregistered")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceOperatorUnregisteredIterator{contract: _AVSGovernance.contract, event: "OperatorUnregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorUnregistered is a free log subscription operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_AVSGovernance *AVSGovernanceFilterer) WatchOperatorUnregistered(opts *bind.WatchOpts, sink chan<- *AVSGovernanceOperatorUnregistered) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "OperatorUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceOperatorUnregistered)
				if err := _AVSGovernance.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
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

// ParseOperatorUnregistered is a log parse operation binding the contract event 0x6f42117a557500c705ddf040a619d86f39101e6b74ac20d7b3e5943ba473fc7f.
//
// Solidity: event OperatorUnregistered(address operator)
func (_AVSGovernance *AVSGovernanceFilterer) ParseOperatorUnregistered(log types.Log) (*AVSGovernanceOperatorUnregistered, error) {
	event := new(AVSGovernanceOperatorUnregistered)
	if err := _AVSGovernance.contract.UnpackLog(event, "OperatorUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceQueuedRewardsReceiverModificationIterator is returned from FilterQueuedRewardsReceiverModification and is used to iterate over the raw logs and unpacked data for QueuedRewardsReceiverModification events raised by the AVSGovernance contract.
type AVSGovernanceQueuedRewardsReceiverModificationIterator struct {
	Event *AVSGovernanceQueuedRewardsReceiverModification // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceQueuedRewardsReceiverModificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceQueuedRewardsReceiverModification)
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
		it.Event = new(AVSGovernanceQueuedRewardsReceiverModification)
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
func (it *AVSGovernanceQueuedRewardsReceiverModificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceQueuedRewardsReceiverModificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceQueuedRewardsReceiverModification represents a QueuedRewardsReceiverModification event raised by the AVSGovernance contract.
type AVSGovernanceQueuedRewardsReceiverModification struct {
	Operator common.Address
	Receiver common.Address
	Delay    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterQueuedRewardsReceiverModification is a free log retrieval operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address operator, address receiver, uint256 delay)
func (_AVSGovernance *AVSGovernanceFilterer) FilterQueuedRewardsReceiverModification(opts *bind.FilterOpts) (*AVSGovernanceQueuedRewardsReceiverModificationIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "QueuedRewardsReceiverModification")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceQueuedRewardsReceiverModificationIterator{contract: _AVSGovernance.contract, event: "QueuedRewardsReceiverModification", logs: logs, sub: sub}, nil
}

// WatchQueuedRewardsReceiverModification is a free log subscription operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address operator, address receiver, uint256 delay)
func (_AVSGovernance *AVSGovernanceFilterer) WatchQueuedRewardsReceiverModification(opts *bind.WatchOpts, sink chan<- *AVSGovernanceQueuedRewardsReceiverModification) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "QueuedRewardsReceiverModification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceQueuedRewardsReceiverModification)
				if err := _AVSGovernance.contract.UnpackLog(event, "QueuedRewardsReceiverModification", log); err != nil {
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

// ParseQueuedRewardsReceiverModification is a log parse operation binding the contract event 0x0d8cfa10a3087b28d3c226ad9a37314860e7c3c0505a25a39e3cdefb3118a98a.
//
// Solidity: event QueuedRewardsReceiverModification(address operator, address receiver, uint256 delay)
func (_AVSGovernance *AVSGovernanceFilterer) ParseQueuedRewardsReceiverModification(log types.Log) (*AVSGovernanceQueuedRewardsReceiverModification, error) {
	event := new(AVSGovernanceQueuedRewardsReceiverModification)
	if err := _AVSGovernance.contract.UnpackLog(event, "QueuedRewardsReceiverModification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AVSGovernance contract.
type AVSGovernanceRoleAdminChangedIterator struct {
	Event *AVSGovernanceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceRoleAdminChanged)
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
		it.Event = new(AVSGovernanceRoleAdminChanged)
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
func (it *AVSGovernanceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceRoleAdminChanged represents a RoleAdminChanged event raised by the AVSGovernance contract.
type AVSGovernanceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AVSGovernance *AVSGovernanceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AVSGovernanceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceRoleAdminChangedIterator{contract: _AVSGovernance.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AVSGovernance *AVSGovernanceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AVSGovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceRoleAdminChanged)
				if err := _AVSGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AVSGovernance *AVSGovernanceFilterer) ParseRoleAdminChanged(log types.Log) (*AVSGovernanceRoleAdminChanged, error) {
	event := new(AVSGovernanceRoleAdminChanged)
	if err := _AVSGovernance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AVSGovernance contract.
type AVSGovernanceRoleGrantedIterator struct {
	Event *AVSGovernanceRoleGranted // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceRoleGranted)
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
		it.Event = new(AVSGovernanceRoleGranted)
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
func (it *AVSGovernanceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceRoleGranted represents a RoleGranted event raised by the AVSGovernance contract.
type AVSGovernanceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AVSGovernance *AVSGovernanceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AVSGovernanceRoleGrantedIterator, error) {

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

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceRoleGrantedIterator{contract: _AVSGovernance.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AVSGovernance *AVSGovernanceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AVSGovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceRoleGranted)
				if err := _AVSGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AVSGovernance *AVSGovernanceFilterer) ParseRoleGranted(log types.Log) (*AVSGovernanceRoleGranted, error) {
	event := new(AVSGovernanceRoleGranted)
	if err := _AVSGovernance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AVSGovernance contract.
type AVSGovernanceRoleRevokedIterator struct {
	Event *AVSGovernanceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceRoleRevoked)
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
		it.Event = new(AVSGovernanceRoleRevoked)
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
func (it *AVSGovernanceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceRoleRevoked represents a RoleRevoked event raised by the AVSGovernance contract.
type AVSGovernanceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AVSGovernance *AVSGovernanceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AVSGovernanceRoleRevokedIterator, error) {

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

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceRoleRevokedIterator{contract: _AVSGovernance.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AVSGovernance *AVSGovernanceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AVSGovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceRoleRevoked)
				if err := _AVSGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AVSGovernance *AVSGovernanceFilterer) ParseRoleRevoked(log types.Log) (*AVSGovernanceRoleRevoked, error) {
	event := new(AVSGovernanceRoleRevoked)
	if err := _AVSGovernance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetAllowlistSignerIterator is returned from FilterSetAllowlistSigner and is used to iterate over the raw logs and unpacked data for SetAllowlistSigner events raised by the AVSGovernance contract.
type AVSGovernanceSetAllowlistSignerIterator struct {
	Event *AVSGovernanceSetAllowlistSigner // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetAllowlistSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetAllowlistSigner)
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
		it.Event = new(AVSGovernanceSetAllowlistSigner)
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
func (it *AVSGovernanceSetAllowlistSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetAllowlistSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetAllowlistSigner represents a SetAllowlistSigner event raised by the AVSGovernance contract.
type AVSGovernanceSetAllowlistSigner struct {
	AllowlistSigner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetAllowlistSigner is a free log retrieval operation binding the contract event 0xfa4acc0aaeb2714e420e9c8339167ddef7bc66c0f94a0c5a7722de21dcb7508c.
//
// Solidity: event SetAllowlistSigner(address allowlistSigner)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetAllowlistSigner(opts *bind.FilterOpts) (*AVSGovernanceSetAllowlistSignerIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetAllowlistSigner")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetAllowlistSignerIterator{contract: _AVSGovernance.contract, event: "SetAllowlistSigner", logs: logs, sub: sub}, nil
}

// WatchSetAllowlistSigner is a free log subscription operation binding the contract event 0xfa4acc0aaeb2714e420e9c8339167ddef7bc66c0f94a0c5a7722de21dcb7508c.
//
// Solidity: event SetAllowlistSigner(address allowlistSigner)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetAllowlistSigner(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetAllowlistSigner) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetAllowlistSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetAllowlistSigner)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetAllowlistSigner", log); err != nil {
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

// ParseSetAllowlistSigner is a log parse operation binding the contract event 0xfa4acc0aaeb2714e420e9c8339167ddef7bc66c0f94a0c5a7722de21dcb7508c.
//
// Solidity: event SetAllowlistSigner(address allowlistSigner)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetAllowlistSigner(log types.Log) (*AVSGovernanceSetAllowlistSigner, error) {
	event := new(AVSGovernanceSetAllowlistSigner)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetAllowlistSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetAvsGovernanceLogicIterator is returned from FilterSetAvsGovernanceLogic and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceLogic events raised by the AVSGovernance contract.
type AVSGovernanceSetAvsGovernanceLogicIterator struct {
	Event *AVSGovernanceSetAvsGovernanceLogic // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetAvsGovernanceLogicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetAvsGovernanceLogic)
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
		it.Event = new(AVSGovernanceSetAvsGovernanceLogic)
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
func (it *AVSGovernanceSetAvsGovernanceLogicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetAvsGovernanceLogicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetAvsGovernanceLogic represents a SetAvsGovernanceLogic event raised by the AVSGovernance contract.
type AVSGovernanceSetAvsGovernanceLogic struct {
	AvsGovernanceLogic common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceLogic is a free log retrieval operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetAvsGovernanceLogic(opts *bind.FilterOpts) (*AVSGovernanceSetAvsGovernanceLogicIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetAvsGovernanceLogic")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetAvsGovernanceLogicIterator{contract: _AVSGovernance.contract, event: "SetAvsGovernanceLogic", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceLogic is a free log subscription operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetAvsGovernanceLogic(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetAvsGovernanceLogic) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetAvsGovernanceLogic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetAvsGovernanceLogic)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsGovernanceLogic", log); err != nil {
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

// ParseSetAvsGovernanceLogic is a log parse operation binding the contract event 0x7c36ee80df183e227956a9f387a48d26bbf4d2f1526410493d11126de5a8942c.
//
// Solidity: event SetAvsGovernanceLogic(address avsGovernanceLogic)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetAvsGovernanceLogic(log types.Log) (*AVSGovernanceSetAvsGovernanceLogic, error) {
	event := new(AVSGovernanceSetAvsGovernanceLogic)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsGovernanceLogic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetAvsGovernanceMultiplierSyncerIterator is returned from FilterSetAvsGovernanceMultiplierSyncer and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceMultiplierSyncer events raised by the AVSGovernance contract.
type AVSGovernanceSetAvsGovernanceMultiplierSyncerIterator struct {
	Event *AVSGovernanceSetAvsGovernanceMultiplierSyncer // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetAvsGovernanceMultiplierSyncerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetAvsGovernanceMultiplierSyncer)
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
		it.Event = new(AVSGovernanceSetAvsGovernanceMultiplierSyncer)
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
func (it *AVSGovernanceSetAvsGovernanceMultiplierSyncerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetAvsGovernanceMultiplierSyncerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetAvsGovernanceMultiplierSyncer represents a SetAvsGovernanceMultiplierSyncer event raised by the AVSGovernance contract.
type AVSGovernanceSetAvsGovernanceMultiplierSyncer struct {
	AvsGovernanceMultiplierSyncer common.Address
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceMultiplierSyncer is a free log retrieval operation binding the contract event 0xb73a70f24733a9265231de5807eae76d1740a9974b31a142ef9e243508987bbe.
//
// Solidity: event SetAvsGovernanceMultiplierSyncer(address avsGovernanceMultiplierSyncer)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetAvsGovernanceMultiplierSyncer(opts *bind.FilterOpts) (*AVSGovernanceSetAvsGovernanceMultiplierSyncerIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetAvsGovernanceMultiplierSyncer")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetAvsGovernanceMultiplierSyncerIterator{contract: _AVSGovernance.contract, event: "SetAvsGovernanceMultiplierSyncer", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceMultiplierSyncer is a free log subscription operation binding the contract event 0xb73a70f24733a9265231de5807eae76d1740a9974b31a142ef9e243508987bbe.
//
// Solidity: event SetAvsGovernanceMultiplierSyncer(address avsGovernanceMultiplierSyncer)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetAvsGovernanceMultiplierSyncer(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetAvsGovernanceMultiplierSyncer) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetAvsGovernanceMultiplierSyncer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetAvsGovernanceMultiplierSyncer)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultiplierSyncer", log); err != nil {
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

// ParseSetAvsGovernanceMultiplierSyncer is a log parse operation binding the contract event 0xb73a70f24733a9265231de5807eae76d1740a9974b31a142ef9e243508987bbe.
//
// Solidity: event SetAvsGovernanceMultiplierSyncer(address avsGovernanceMultiplierSyncer)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetAvsGovernanceMultiplierSyncer(log types.Log) (*AVSGovernanceSetAvsGovernanceMultiplierSyncer, error) {
	event := new(AVSGovernanceSetAvsGovernanceMultiplierSyncer)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultiplierSyncer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetAvsGovernanceMultisigIterator is returned from FilterSetAvsGovernanceMultisig and is used to iterate over the raw logs and unpacked data for SetAvsGovernanceMultisig events raised by the AVSGovernance contract.
type AVSGovernanceSetAvsGovernanceMultisigIterator struct {
	Event *AVSGovernanceSetAvsGovernanceMultisig // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetAvsGovernanceMultisigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetAvsGovernanceMultisig)
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
		it.Event = new(AVSGovernanceSetAvsGovernanceMultisig)
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
func (it *AVSGovernanceSetAvsGovernanceMultisigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetAvsGovernanceMultisigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetAvsGovernanceMultisig represents a SetAvsGovernanceMultisig event raised by the AVSGovernance contract.
type AVSGovernanceSetAvsGovernanceMultisig struct {
	NewAvsGovernanceMultisig common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetAvsGovernanceMultisig is a free log retrieval operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetAvsGovernanceMultisig(opts *bind.FilterOpts) (*AVSGovernanceSetAvsGovernanceMultisigIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetAvsGovernanceMultisig")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetAvsGovernanceMultisigIterator{contract: _AVSGovernance.contract, event: "SetAvsGovernanceMultisig", logs: logs, sub: sub}, nil
}

// WatchSetAvsGovernanceMultisig is a free log subscription operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetAvsGovernanceMultisig(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetAvsGovernanceMultisig) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetAvsGovernanceMultisig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetAvsGovernanceMultisig)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultisig", log); err != nil {
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

// ParseSetAvsGovernanceMultisig is a log parse operation binding the contract event 0x024e98b7d808a3ddb028252dc95dfdcb165a0ca59fcff8984b4fecf9a7222649.
//
// Solidity: event SetAvsGovernanceMultisig(address newAvsGovernanceMultisig)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetAvsGovernanceMultisig(log types.Log) (*AVSGovernanceSetAvsGovernanceMultisig, error) {
	event := new(AVSGovernanceSetAvsGovernanceMultisig)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsGovernanceMultisig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetAvsNameIterator is returned from FilterSetAvsName and is used to iterate over the raw logs and unpacked data for SetAvsName events raised by the AVSGovernance contract.
type AVSGovernanceSetAvsNameIterator struct {
	Event *AVSGovernanceSetAvsName // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetAvsNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetAvsName)
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
		it.Event = new(AVSGovernanceSetAvsName)
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
func (it *AVSGovernanceSetAvsNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetAvsNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetAvsName represents a SetAvsName event raised by the AVSGovernance contract.
type AVSGovernanceSetAvsName struct {
	AvsName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAvsName is a free log retrieval operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetAvsName(opts *bind.FilterOpts) (*AVSGovernanceSetAvsNameIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetAvsName")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetAvsNameIterator{contract: _AVSGovernance.contract, event: "SetAvsName", logs: logs, sub: sub}, nil
}

// WatchSetAvsName is a free log subscription operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetAvsName(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetAvsName) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetAvsName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetAvsName)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsName", log); err != nil {
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

// ParseSetAvsName is a log parse operation binding the contract event 0x7f63aacad63bc1693280450d5c3612ccd4efc53e46d69f3a537db102cd66290c.
//
// Solidity: event SetAvsName(string avsName)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetAvsName(log types.Log) (*AVSGovernanceSetAvsName, error) {
	event := new(AVSGovernanceSetAvsName)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetAvsName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetIsAllowlistedIterator is returned from FilterSetIsAllowlisted and is used to iterate over the raw logs and unpacked data for SetIsAllowlisted events raised by the AVSGovernance contract.
type AVSGovernanceSetIsAllowlistedIterator struct {
	Event *AVSGovernanceSetIsAllowlisted // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetIsAllowlistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetIsAllowlisted)
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
		it.Event = new(AVSGovernanceSetIsAllowlisted)
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
func (it *AVSGovernanceSetIsAllowlistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetIsAllowlistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetIsAllowlisted represents a SetIsAllowlisted event raised by the AVSGovernance contract.
type AVSGovernanceSetIsAllowlisted struct {
	IsAllowlisted bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetIsAllowlisted is a free log retrieval operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetIsAllowlisted(opts *bind.FilterOpts) (*AVSGovernanceSetIsAllowlistedIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetIsAllowlisted")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetIsAllowlistedIterator{contract: _AVSGovernance.contract, event: "SetIsAllowlisted", logs: logs, sub: sub}, nil
}

// WatchSetIsAllowlisted is a free log subscription operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetIsAllowlisted(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetIsAllowlisted) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetIsAllowlisted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetIsAllowlisted)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetIsAllowlisted", log); err != nil {
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

// ParseSetIsAllowlisted is a log parse operation binding the contract event 0x2dcb3282f9b7aa18e1bf7fa254c45f3e270e8f26d9a37ae590d5d8125b58d1b1.
//
// Solidity: event SetIsAllowlisted(bool isAllowlisted)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetIsAllowlisted(log types.Log) (*AVSGovernanceSetIsAllowlisted, error) {
	event := new(AVSGovernanceSetIsAllowlisted)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetIsAllowlisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetMessageHandlerIterator is returned from FilterSetMessageHandler and is used to iterate over the raw logs and unpacked data for SetMessageHandler events raised by the AVSGovernance contract.
type AVSGovernanceSetMessageHandlerIterator struct {
	Event *AVSGovernanceSetMessageHandler // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetMessageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetMessageHandler)
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
		it.Event = new(AVSGovernanceSetMessageHandler)
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
func (it *AVSGovernanceSetMessageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetMessageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetMessageHandler represents a SetMessageHandler event raised by the AVSGovernance contract.
type AVSGovernanceSetMessageHandler struct {
	NewMessageHandler common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMessageHandler is a free log retrieval operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetMessageHandler(opts *bind.FilterOpts) (*AVSGovernanceSetMessageHandlerIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetMessageHandler")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetMessageHandlerIterator{contract: _AVSGovernance.contract, event: "SetMessageHandler", logs: logs, sub: sub}, nil
}

// WatchSetMessageHandler is a free log subscription operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetMessageHandler(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetMessageHandler) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetMessageHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetMessageHandler)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetMessageHandler", log); err != nil {
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

// ParseSetMessageHandler is a log parse operation binding the contract event 0x997f84b541d7b68e210e6f50e3402b51d8411dbbc4d44ed81e508383126e4e94.
//
// Solidity: event SetMessageHandler(address newMessageHandler)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetMessageHandler(log types.Log) (*AVSGovernanceSetMessageHandler, error) {
	event := new(AVSGovernanceSetMessageHandler)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetMessageHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetNumOfOperatorsLimitIterator is returned from FilterSetNumOfOperatorsLimit and is used to iterate over the raw logs and unpacked data for SetNumOfOperatorsLimit events raised by the AVSGovernance contract.
type AVSGovernanceSetNumOfOperatorsLimitIterator struct {
	Event *AVSGovernanceSetNumOfOperatorsLimit // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetNumOfOperatorsLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetNumOfOperatorsLimit)
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
		it.Event = new(AVSGovernanceSetNumOfOperatorsLimit)
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
func (it *AVSGovernanceSetNumOfOperatorsLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetNumOfOperatorsLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetNumOfOperatorsLimit represents a SetNumOfOperatorsLimit event raised by the AVSGovernance contract.
type AVSGovernanceSetNumOfOperatorsLimit struct {
	NewLimitOfNumOfOperators *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetNumOfOperatorsLimit is a free log retrieval operation binding the contract event 0xc0dd1d82df4ae12576f7a7912395305cf73deae26c764dd74a945cd6ba81591b.
//
// Solidity: event SetNumOfOperatorsLimit(uint256 newLimitOfNumOfOperators)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetNumOfOperatorsLimit(opts *bind.FilterOpts) (*AVSGovernanceSetNumOfOperatorsLimitIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetNumOfOperatorsLimit")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetNumOfOperatorsLimitIterator{contract: _AVSGovernance.contract, event: "SetNumOfOperatorsLimit", logs: logs, sub: sub}, nil
}

// WatchSetNumOfOperatorsLimit is a free log subscription operation binding the contract event 0xc0dd1d82df4ae12576f7a7912395305cf73deae26c764dd74a945cd6ba81591b.
//
// Solidity: event SetNumOfOperatorsLimit(uint256 newLimitOfNumOfOperators)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetNumOfOperatorsLimit(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetNumOfOperatorsLimit) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetNumOfOperatorsLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetNumOfOperatorsLimit)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetNumOfOperatorsLimit", log); err != nil {
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

// ParseSetNumOfOperatorsLimit is a log parse operation binding the contract event 0xc0dd1d82df4ae12576f7a7912395305cf73deae26c764dd74a945cd6ba81591b.
//
// Solidity: event SetNumOfOperatorsLimit(uint256 newLimitOfNumOfOperators)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetNumOfOperatorsLimit(log types.Log) (*AVSGovernanceSetNumOfOperatorsLimit, error) {
	event := new(AVSGovernanceSetNumOfOperatorsLimit)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetNumOfOperatorsLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetOthenticRegistryIterator is returned from FilterSetOthenticRegistry and is used to iterate over the raw logs and unpacked data for SetOthenticRegistry events raised by the AVSGovernance contract.
type AVSGovernanceSetOthenticRegistryIterator struct {
	Event *AVSGovernanceSetOthenticRegistry // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetOthenticRegistryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetOthenticRegistry)
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
		it.Event = new(AVSGovernanceSetOthenticRegistry)
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
func (it *AVSGovernanceSetOthenticRegistryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetOthenticRegistryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetOthenticRegistry represents a SetOthenticRegistry event raised by the AVSGovernance contract.
type AVSGovernanceSetOthenticRegistry struct {
	OthenticRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetOthenticRegistry is a free log retrieval operation binding the contract event 0xf9855cc914fefc396bdeb5a4dcb97a2f6c75f4d6f00a8e71d6f9a40e474afe8d.
//
// Solidity: event SetOthenticRegistry(address othenticRegistry)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetOthenticRegistry(opts *bind.FilterOpts) (*AVSGovernanceSetOthenticRegistryIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetOthenticRegistry")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetOthenticRegistryIterator{contract: _AVSGovernance.contract, event: "SetOthenticRegistry", logs: logs, sub: sub}, nil
}

// WatchSetOthenticRegistry is a free log subscription operation binding the contract event 0xf9855cc914fefc396bdeb5a4dcb97a2f6c75f4d6f00a8e71d6f9a40e474afe8d.
//
// Solidity: event SetOthenticRegistry(address othenticRegistry)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetOthenticRegistry(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetOthenticRegistry) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetOthenticRegistry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetOthenticRegistry)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetOthenticRegistry", log); err != nil {
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

// ParseSetOthenticRegistry is a log parse operation binding the contract event 0xf9855cc914fefc396bdeb5a4dcb97a2f6c75f4d6f00a8e71d6f9a40e474afe8d.
//
// Solidity: event SetOthenticRegistry(address othenticRegistry)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetOthenticRegistry(log types.Log) (*AVSGovernanceSetOthenticRegistry, error) {
	event := new(AVSGovernanceSetOthenticRegistry)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetOthenticRegistry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetRewardsReceiverIterator is returned from FilterSetRewardsReceiver and is used to iterate over the raw logs and unpacked data for SetRewardsReceiver events raised by the AVSGovernance contract.
type AVSGovernanceSetRewardsReceiverIterator struct {
	Event *AVSGovernanceSetRewardsReceiver // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetRewardsReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetRewardsReceiver)
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
		it.Event = new(AVSGovernanceSetRewardsReceiver)
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
func (it *AVSGovernanceSetRewardsReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetRewardsReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetRewardsReceiver represents a SetRewardsReceiver event raised by the AVSGovernance contract.
type AVSGovernanceSetRewardsReceiver struct {
	Operator common.Address
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRewardsReceiver is a free log retrieval operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address operator, address receiver)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetRewardsReceiver(opts *bind.FilterOpts) (*AVSGovernanceSetRewardsReceiverIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetRewardsReceiver")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetRewardsReceiverIterator{contract: _AVSGovernance.contract, event: "SetRewardsReceiver", logs: logs, sub: sub}, nil
}

// WatchSetRewardsReceiver is a free log subscription operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address operator, address receiver)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetRewardsReceiver(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetRewardsReceiver) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetRewardsReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetRewardsReceiver)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetRewardsReceiver", log); err != nil {
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

// ParseSetRewardsReceiver is a log parse operation binding the contract event 0xe906feea2ef60b474e22b4169bdd4de6906a84cd448cbcee99593526fe87082d.
//
// Solidity: event SetRewardsReceiver(address operator, address receiver)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetRewardsReceiver(log types.Log) (*AVSGovernanceSetRewardsReceiver, error) {
	event := new(AVSGovernanceSetRewardsReceiver)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetRewardsReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetRewardsReceiverModificationDelayIterator is returned from FilterSetRewardsReceiverModificationDelay and is used to iterate over the raw logs and unpacked data for SetRewardsReceiverModificationDelay events raised by the AVSGovernance contract.
type AVSGovernanceSetRewardsReceiverModificationDelayIterator struct {
	Event *AVSGovernanceSetRewardsReceiverModificationDelay // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetRewardsReceiverModificationDelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetRewardsReceiverModificationDelay)
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
		it.Event = new(AVSGovernanceSetRewardsReceiverModificationDelay)
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
func (it *AVSGovernanceSetRewardsReceiverModificationDelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetRewardsReceiverModificationDelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetRewardsReceiverModificationDelay represents a SetRewardsReceiverModificationDelay event raised by the AVSGovernance contract.
type AVSGovernanceSetRewardsReceiverModificationDelay struct {
	ModificationDelay *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetRewardsReceiverModificationDelay is a free log retrieval operation binding the contract event 0x47c8c3268759fc47868c5e319217a2e85d47bd3935a4108debe246f6025fb88b.
//
// Solidity: event SetRewardsReceiverModificationDelay(uint256 modificationDelay)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetRewardsReceiverModificationDelay(opts *bind.FilterOpts) (*AVSGovernanceSetRewardsReceiverModificationDelayIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetRewardsReceiverModificationDelay")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetRewardsReceiverModificationDelayIterator{contract: _AVSGovernance.contract, event: "SetRewardsReceiverModificationDelay", logs: logs, sub: sub}, nil
}

// WatchSetRewardsReceiverModificationDelay is a free log subscription operation binding the contract event 0x47c8c3268759fc47868c5e319217a2e85d47bd3935a4108debe246f6025fb88b.
//
// Solidity: event SetRewardsReceiverModificationDelay(uint256 modificationDelay)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetRewardsReceiverModificationDelay(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetRewardsReceiverModificationDelay) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetRewardsReceiverModificationDelay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetRewardsReceiverModificationDelay)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetRewardsReceiverModificationDelay", log); err != nil {
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

// ParseSetRewardsReceiverModificationDelay is a log parse operation binding the contract event 0x47c8c3268759fc47868c5e319217a2e85d47bd3935a4108debe246f6025fb88b.
//
// Solidity: event SetRewardsReceiverModificationDelay(uint256 modificationDelay)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetRewardsReceiverModificationDelay(log types.Log) (*AVSGovernanceSetRewardsReceiverModificationDelay, error) {
	event := new(AVSGovernanceSetRewardsReceiverModificationDelay)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetRewardsReceiverModificationDelay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetStrategyMultiplierIterator is returned from FilterSetStrategyMultiplier and is used to iterate over the raw logs and unpacked data for SetStrategyMultiplier events raised by the AVSGovernance contract.
type AVSGovernanceSetStrategyMultiplierIterator struct {
	Event *AVSGovernanceSetStrategyMultiplier // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetStrategyMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetStrategyMultiplier)
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
		it.Event = new(AVSGovernanceSetStrategyMultiplier)
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
func (it *AVSGovernanceSetStrategyMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetStrategyMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetStrategyMultiplier represents a SetStrategyMultiplier event raised by the AVSGovernance contract.
type AVSGovernanceSetStrategyMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetStrategyMultiplier is a free log retrieval operation binding the contract event 0x8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c81.
//
// Solidity: event SetStrategyMultiplier(address strategy, uint256 multiplier)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetStrategyMultiplier(opts *bind.FilterOpts) (*AVSGovernanceSetStrategyMultiplierIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetStrategyMultiplier")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetStrategyMultiplierIterator{contract: _AVSGovernance.contract, event: "SetStrategyMultiplier", logs: logs, sub: sub}, nil
}

// WatchSetStrategyMultiplier is a free log subscription operation binding the contract event 0x8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c81.
//
// Solidity: event SetStrategyMultiplier(address strategy, uint256 multiplier)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetStrategyMultiplier(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetStrategyMultiplier) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetStrategyMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetStrategyMultiplier)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetStrategyMultiplier", log); err != nil {
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

// ParseSetStrategyMultiplier is a log parse operation binding the contract event 0x8ae53ffd0ebc018acb19342fba690554d49ae9a467a9606a38b49cb5ad775c81.
//
// Solidity: event SetStrategyMultiplier(address strategy, uint256 multiplier)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetStrategyMultiplier(log types.Log) (*AVSGovernanceSetStrategyMultiplier, error) {
	event := new(AVSGovernanceSetStrategyMultiplier)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetStrategyMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetSupportedStrategiesIterator is returned from FilterSetSupportedStrategies and is used to iterate over the raw logs and unpacked data for SetSupportedStrategies events raised by the AVSGovernance contract.
type AVSGovernanceSetSupportedStrategiesIterator struct {
	Event *AVSGovernanceSetSupportedStrategies // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetSupportedStrategiesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetSupportedStrategies)
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
		it.Event = new(AVSGovernanceSetSupportedStrategies)
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
func (it *AVSGovernanceSetSupportedStrategiesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetSupportedStrategiesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetSupportedStrategies represents a SetSupportedStrategies event raised by the AVSGovernance contract.
type AVSGovernanceSetSupportedStrategies struct {
	Strategies []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetSupportedStrategies is a free log retrieval operation binding the contract event 0xf009a6ffded424f714e8904d643a1ea4479453188faf08a3996121996b76684f.
//
// Solidity: event SetSupportedStrategies(address[] strategies)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetSupportedStrategies(opts *bind.FilterOpts) (*AVSGovernanceSetSupportedStrategiesIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetSupportedStrategies")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetSupportedStrategiesIterator{contract: _AVSGovernance.contract, event: "SetSupportedStrategies", logs: logs, sub: sub}, nil
}

// WatchSetSupportedStrategies is a free log subscription operation binding the contract event 0xf009a6ffded424f714e8904d643a1ea4479453188faf08a3996121996b76684f.
//
// Solidity: event SetSupportedStrategies(address[] strategies)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetSupportedStrategies(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetSupportedStrategies) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetSupportedStrategies")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetSupportedStrategies)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetSupportedStrategies", log); err != nil {
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

// ParseSetSupportedStrategies is a log parse operation binding the contract event 0xf009a6ffded424f714e8904d643a1ea4479453188faf08a3996121996b76684f.
//
// Solidity: event SetSupportedStrategies(address[] strategies)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetSupportedStrategies(log types.Log) (*AVSGovernanceSetSupportedStrategies, error) {
	event := new(AVSGovernanceSetSupportedStrategies)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetSupportedStrategies", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSGovernanceSetTokenIterator is returned from FilterSetToken and is used to iterate over the raw logs and unpacked data for SetToken events raised by the AVSGovernance contract.
type AVSGovernanceSetTokenIterator struct {
	Event *AVSGovernanceSetToken // Event containing the contract specifics and raw log

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
func (it *AVSGovernanceSetTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSGovernanceSetToken)
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
		it.Event = new(AVSGovernanceSetToken)
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
func (it *AVSGovernanceSetTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSGovernanceSetTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSGovernanceSetToken represents a SetToken event raised by the AVSGovernance contract.
type AVSGovernanceSetToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetToken is a free log retrieval operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_AVSGovernance *AVSGovernanceFilterer) FilterSetToken(opts *bind.FilterOpts) (*AVSGovernanceSetTokenIterator, error) {

	logs, sub, err := _AVSGovernance.contract.FilterLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return &AVSGovernanceSetTokenIterator{contract: _AVSGovernance.contract, event: "SetToken", logs: logs, sub: sub}, nil
}

// WatchSetToken is a free log subscription operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_AVSGovernance *AVSGovernanceFilterer) WatchSetToken(opts *bind.WatchOpts, sink chan<- *AVSGovernanceSetToken) (event.Subscription, error) {

	logs, sub, err := _AVSGovernance.contract.WatchLogs(opts, "SetToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSGovernanceSetToken)
				if err := _AVSGovernance.contract.UnpackLog(event, "SetToken", log); err != nil {
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

// ParseSetToken is a log parse operation binding the contract event 0xefc1fd16ea80a922086ee4e995739d59b025c1bcea6d1f67855747480c83214b.
//
// Solidity: event SetToken(address token)
func (_AVSGovernance *AVSGovernanceFilterer) ParseSetToken(log types.Log) (*AVSGovernanceSetToken, error) {
	event := new(AVSGovernanceSetToken)
	if err := _AVSGovernance.contract.UnpackLog(event, "SetToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
