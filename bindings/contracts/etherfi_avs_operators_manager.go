// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package etherfi_avs_operators_manager

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// EtherfiAvsOperatorsManagerMetaData contains all meta data concerning the EtherfiAvsOperatorsManager contract.
var EtherfiAvsOperatorsManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"etherFiAvsOperator\",\"type\":\"address\"}],\"name\":\"CreatedEtherFiAvsOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"DeregisteredOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"ModifiedOperatorDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"detail\",\"type\":\"tuple\"}],\"name\":\"RegisteredAsOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"RegisteredBlsKeyAsDelegatedNodeOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"RegisteredOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsNodeRunner\",\"type\":\"address\"}],\"name\":\"UpdatedAvsNodeRunner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"name\":\"UpdatedAvsWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ecdsaSigner\",\"type\":\"address\"}],\"name\":\"UpdatedEcdsaSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"UpdatedOperatorMetadataURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"UpdatedSocket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"avsOperators\",\"outputs\":[{\"internalType\":\"contractEtherFiAvsOperator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_etherFiAvsOperatorImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nums\",\"type\":\"uint256\"}],\"name\":\"instantiateEtherFiAvsOperator\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"modifyOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvsOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_detail\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_metaDataURI\",\"type\":\"string\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"registerBlsKeyAsDelegatedNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"name\":\"_operatorKickParams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_churnApproverSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorWithChurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsNodeRunner\",\"type\":\"address\"}],\"name\":\"updateAvsNodeRunner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isWhitelisted\",\"type\":\"bool\"}],\"name\":\"updateAvsWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_ecdsaSigner\",\"type\":\"address\"}],\"name\":\"updateEcdsaSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateOperatorMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"}],\"name\":\"updateSocket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradableBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeEtherFiAvsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// EtherfiAvsOperatorsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherfiAvsOperatorsManagerMetaData.ABI instead.
var EtherfiAvsOperatorsManagerABI = EtherfiAvsOperatorsManagerMetaData.ABI

// EtherfiAvsOperatorsManager is an auto generated Go binding around an Ethereum contract.
type EtherfiAvsOperatorsManager struct {
	EtherfiAvsOperatorsManagerCaller     // Read-only binding to the contract
	EtherfiAvsOperatorsManagerTransactor // Write-only binding to the contract
	EtherfiAvsOperatorsManagerFilterer   // Log filterer for contract events
}

// EtherfiAvsOperatorsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherfiAvsOperatorsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAvsOperatorsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherfiAvsOperatorsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAvsOperatorsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherfiAvsOperatorsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAvsOperatorsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherfiAvsOperatorsManagerSession struct {
	Contract     *EtherfiAvsOperatorsManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EtherfiAvsOperatorsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherfiAvsOperatorsManagerCallerSession struct {
	Contract *EtherfiAvsOperatorsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// EtherfiAvsOperatorsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherfiAvsOperatorsManagerTransactorSession struct {
	Contract     *EtherfiAvsOperatorsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// EtherfiAvsOperatorsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherfiAvsOperatorsManagerRaw struct {
	Contract *EtherfiAvsOperatorsManager // Generic contract binding to access the raw methods on
}

// EtherfiAvsOperatorsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherfiAvsOperatorsManagerCallerRaw struct {
	Contract *EtherfiAvsOperatorsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EtherfiAvsOperatorsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherfiAvsOperatorsManagerTransactorRaw struct {
	Contract *EtherfiAvsOperatorsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherfiAvsOperatorsManager creates a new instance of EtherfiAvsOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAvsOperatorsManager(address common.Address, backend bind.ContractBackend) (*EtherfiAvsOperatorsManager, error) {
	contract, err := bindEtherfiAvsOperatorsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManager{EtherfiAvsOperatorsManagerCaller: EtherfiAvsOperatorsManagerCaller{contract: contract}, EtherfiAvsOperatorsManagerTransactor: EtherfiAvsOperatorsManagerTransactor{contract: contract}, EtherfiAvsOperatorsManagerFilterer: EtherfiAvsOperatorsManagerFilterer{contract: contract}}, nil
}

// NewEtherfiAvsOperatorsManagerCaller creates a new read-only instance of EtherfiAvsOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAvsOperatorsManagerCaller(address common.Address, caller bind.ContractCaller) (*EtherfiAvsOperatorsManagerCaller, error) {
	contract, err := bindEtherfiAvsOperatorsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerCaller{contract: contract}, nil
}

// NewEtherfiAvsOperatorsManagerTransactor creates a new write-only instance of EtherfiAvsOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAvsOperatorsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherfiAvsOperatorsManagerTransactor, error) {
	contract, err := bindEtherfiAvsOperatorsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerTransactor{contract: contract}, nil
}

// NewEtherfiAvsOperatorsManagerFilterer creates a new log filterer instance of EtherfiAvsOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAvsOperatorsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherfiAvsOperatorsManagerFilterer, error) {
	contract, err := bindEtherfiAvsOperatorsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerFilterer{contract: contract}, nil
}

// bindEtherfiAvsOperatorsManager binds a generic wrapper to an already deployed contract.
func bindEtherfiAvsOperatorsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherfiAvsOperatorsManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAvsOperatorsManager.Contract.EtherfiAvsOperatorsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.EtherfiAvsOperatorsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.EtherfiAvsOperatorsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAvsOperatorsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherfiAvsOperatorsManager.Contract.Admins(&_EtherfiAvsOperatorsManager.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherfiAvsOperatorsManager.Contract.Admins(&_EtherfiAvsOperatorsManager.CallOpts, arg0)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) AvsOperators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "avsOperators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.AvsOperators(&_EtherfiAvsOperatorsManager.CallOpts, arg0)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.AvsOperators(&_EtherfiAvsOperatorsManager.CallOpts, arg0)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) DelegationManager() (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.DelegationManager(&_EtherfiAvsOperatorsManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) DelegationManager() (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.DelegationManager(&_EtherfiAvsOperatorsManager.CallOpts)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) NextAvsOperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "nextAvsOperatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) NextAvsOperatorId() (*big.Int, error) {
	return _EtherfiAvsOperatorsManager.Contract.NextAvsOperatorId(&_EtherfiAvsOperatorsManager.CallOpts)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) NextAvsOperatorId() (*big.Int, error) {
	return _EtherfiAvsOperatorsManager.Contract.NextAvsOperatorId(&_EtherfiAvsOperatorsManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) Owner() (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.Owner(&_EtherfiAvsOperatorsManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) Owner() (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.Owner(&_EtherfiAvsOperatorsManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) Paused() (bool, error) {
	return _EtherfiAvsOperatorsManager.Contract.Paused(&_EtherfiAvsOperatorsManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) Paused() (bool, error) {
	return _EtherfiAvsOperatorsManager.Contract.Paused(&_EtherfiAvsOperatorsManager.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) Pausers(arg0 common.Address) (bool, error) {
	return _EtherfiAvsOperatorsManager.Contract.Pausers(&_EtherfiAvsOperatorsManager.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _EtherfiAvsOperatorsManager.Contract.Pausers(&_EtherfiAvsOperatorsManager.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherfiAvsOperatorsManager.Contract.ProxiableUUID(&_EtherfiAvsOperatorsManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherfiAvsOperatorsManager.Contract.ProxiableUUID(&_EtherfiAvsOperatorsManager.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCaller) UpgradableBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAvsOperatorsManager.contract.Call(opts, &out, "upgradableBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpgradableBeacon() (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradableBeacon(&_EtherfiAvsOperatorsManager.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerCallerSession) UpgradableBeacon() (common.Address, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradableBeacon(&_EtherfiAvsOperatorsManager.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) DeregisterOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "deregisterOperator", _id, _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) DeregisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.DeregisterOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) DeregisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.DeregisterOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManager, address _etherFiAvsOperatorImpl) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) Initialize(opts *bind.TransactOpts, _delegationManager common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "initialize", _delegationManager, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManager, address _etherFiAvsOperatorImpl) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) Initialize(_delegationManager common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.Initialize(&_EtherfiAvsOperatorsManager.TransactOpts, _delegationManager, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManager, address _etherFiAvsOperatorImpl) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) Initialize(_delegationManager common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.Initialize(&_EtherfiAvsOperatorsManager.TransactOpts, _delegationManager, _etherFiAvsOperatorImpl)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) InstantiateEtherFiAvsOperator(opts *bind.TransactOpts, _nums *big.Int) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "instantiateEtherFiAvsOperator", _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.InstantiateEtherFiAvsOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.InstantiateEtherFiAvsOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _nums)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, _id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "modifyOperatorDetails", _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.ModifyOperatorDetails(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.ModifyOperatorDetails(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, _id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "registerAsOperator", _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterAsOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterAsOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _detail, _metaDataURI)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) RegisterBlsKeyAsDelegatedNodeOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "registerBlsKeyAsDelegatedNodeOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) RegisterBlsKeyAsDelegatedNodeOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) RegisterBlsKeyAsDelegatedNodeOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) RegisterOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "registerOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) RegisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) RegisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "registerOperatorWithChurn", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) RegisterOperatorWithChurn(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterOperatorWithChurn(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) RegisterOperatorWithChurn(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RegisterOperatorWithChurn(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RenounceOwnership(&_EtherfiAvsOperatorsManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.RenounceOwnership(&_EtherfiAvsOperatorsManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.TransferOwnership(&_EtherfiAvsOperatorsManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.TransferOwnership(&_EtherfiAvsOperatorsManager.TransactOpts, newOwner)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateAdmin(&_EtherfiAvsOperatorsManager.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateAdmin(&_EtherfiAvsOperatorsManager.TransactOpts, _address, _isAdmin)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpdateAvsNodeRunner(opts *bind.TransactOpts, _id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "updateAvsNodeRunner", _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateAvsNodeRunner(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateAvsNodeRunner(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsNodeRunner)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpdateAvsWhitelist(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "updateAvsWhitelist", _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpdateAvsWhitelist(_id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateAvsWhitelist(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpdateAvsWhitelist(_id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateAvsWhitelist(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpdateEcdsaSigner(opts *bind.TransactOpts, _id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "updateEcdsaSigner", _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateEcdsaSigner(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateEcdsaSigner(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _ecdsaSigner)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, _id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "updateOperatorMetadataURI", _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateOperatorMetadataURI(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateOperatorMetadataURI(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _metadataURI)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpdateSocket(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "updateSocket", _id, _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpdateSocket(_id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateSocket(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpdateSocket(_id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpdateSocket(&_EtherfiAvsOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _socket)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpgradeEtherFiAvsOperator(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "upgradeEtherFiAvsOperator", _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradeEtherFiAvsOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradeEtherFiAvsOperator(&_EtherfiAvsOperatorsManager.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradeTo(&_EtherfiAvsOperatorsManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradeTo(&_EtherfiAvsOperatorsManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradeToAndCall(&_EtherfiAvsOperatorsManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAvsOperatorsManager.Contract.UpgradeToAndCall(&_EtherfiAvsOperatorsManager.TransactOpts, newImplementation, data)
}

// EtherfiAvsOperatorsManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerAdminChangedIterator struct {
	Event *EtherfiAvsOperatorsManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerAdminChanged)
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
		it.Event = new(EtherfiAvsOperatorsManagerAdminChanged)
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
func (it *EtherfiAvsOperatorsManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerAdminChanged represents a AdminChanged event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EtherfiAvsOperatorsManagerAdminChangedIterator, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerAdminChangedIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerAdminChanged)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseAdminChanged(log types.Log) (*EtherfiAvsOperatorsManagerAdminChanged, error) {
	event := new(EtherfiAvsOperatorsManagerAdminChanged)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerBeaconUpgradedIterator struct {
	Event *EtherfiAvsOperatorsManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerBeaconUpgraded)
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
		it.Event = new(EtherfiAvsOperatorsManagerBeaconUpgraded)
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
func (it *EtherfiAvsOperatorsManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerBeaconUpgraded represents a BeaconUpgraded event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EtherfiAvsOperatorsManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerBeaconUpgradedIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerBeaconUpgraded)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseBeaconUpgraded(log types.Log) (*EtherfiAvsOperatorsManagerBeaconUpgraded, error) {
	event := new(EtherfiAvsOperatorsManagerBeaconUpgraded)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperatorIterator is returned from FilterCreatedEtherFiAvsOperator and is used to iterate over the raw logs and unpacked data for CreatedEtherFiAvsOperator events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperatorIterator struct {
	Event *EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator)
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
		it.Event = new(EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator)
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
func (it *EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator represents a CreatedEtherFiAvsOperator event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator struct {
	Id                 *big.Int
	EtherFiAvsOperator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCreatedEtherFiAvsOperator is a free log retrieval operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterCreatedEtherFiAvsOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperatorIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "CreatedEtherFiAvsOperator", logs: logs, sub: sub}, nil
}

// WatchCreatedEtherFiAvsOperator is a free log subscription operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchCreatedEtherFiAvsOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseCreatedEtherFiAvsOperator(log types.Log) (*EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator, error) {
	event := new(EtherfiAvsOperatorsManagerCreatedEtherFiAvsOperator)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerDeregisteredOperatorIterator is returned from FilterDeregisteredOperator and is used to iterate over the raw logs and unpacked data for DeregisteredOperator events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerDeregisteredOperatorIterator struct {
	Event *EtherfiAvsOperatorsManagerDeregisteredOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerDeregisteredOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerDeregisteredOperator)
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
		it.Event = new(EtherfiAvsOperatorsManagerDeregisteredOperator)
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
func (it *EtherfiAvsOperatorsManagerDeregisteredOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerDeregisteredOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerDeregisteredOperator represents a DeregisteredOperator event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerDeregisteredOperator struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	QuorumNumbers          []byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDeregisteredOperator is a free log retrieval operation binding the contract event 0x1376d0cba684c4d21991c950915a2bf8a9181c389d343993ff185ba5944b5a2d.
//
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterDeregisteredOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerDeregisteredOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "DeregisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerDeregisteredOperatorIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "DeregisteredOperator", logs: logs, sub: sub}, nil
}

// WatchDeregisteredOperator is a free log subscription operation binding the contract event 0x1376d0cba684c4d21991c950915a2bf8a9181c389d343993ff185ba5944b5a2d.
//
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchDeregisteredOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerDeregisteredOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "DeregisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerDeregisteredOperator)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "DeregisteredOperator", log); err != nil {
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

// ParseDeregisteredOperator is a log parse operation binding the contract event 0x1376d0cba684c4d21991c950915a2bf8a9181c389d343993ff185ba5944b5a2d.
//
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseDeregisteredOperator(log types.Log) (*EtherfiAvsOperatorsManagerDeregisteredOperator, error) {
	event := new(EtherfiAvsOperatorsManagerDeregisteredOperator)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "DeregisteredOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerInitializedIterator struct {
	Event *EtherfiAvsOperatorsManagerInitialized // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerInitialized)
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
		it.Event = new(EtherfiAvsOperatorsManagerInitialized)
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
func (it *EtherfiAvsOperatorsManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerInitialized represents a Initialized event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EtherfiAvsOperatorsManagerInitializedIterator, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerInitializedIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerInitialized)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseInitialized(log types.Log) (*EtherfiAvsOperatorsManagerInitialized, error) {
	event := new(EtherfiAvsOperatorsManagerInitialized)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerModifiedOperatorDetailsIterator is returned from FilterModifiedOperatorDetails and is used to iterate over the raw logs and unpacked data for ModifiedOperatorDetails events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerModifiedOperatorDetailsIterator struct {
	Event *EtherfiAvsOperatorsManagerModifiedOperatorDetails // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerModifiedOperatorDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerModifiedOperatorDetails)
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
		it.Event = new(EtherfiAvsOperatorsManagerModifiedOperatorDetails)
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
func (it *EtherfiAvsOperatorsManagerModifiedOperatorDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerModifiedOperatorDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerModifiedOperatorDetails represents a ModifiedOperatorDetails event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerModifiedOperatorDetails struct {
	Id                 *big.Int
	NewOperatorDetails IDelegationManagerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterModifiedOperatorDetails is a free log retrieval operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterModifiedOperatorDetails(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerModifiedOperatorDetailsIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerModifiedOperatorDetailsIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "ModifiedOperatorDetails", logs: logs, sub: sub}, nil
}

// WatchModifiedOperatorDetails is a free log subscription operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchModifiedOperatorDetails(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerModifiedOperatorDetails, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerModifiedOperatorDetails)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseModifiedOperatorDetails(log types.Log) (*EtherfiAvsOperatorsManagerModifiedOperatorDetails, error) {
	event := new(EtherfiAvsOperatorsManagerModifiedOperatorDetails)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerOwnershipTransferredIterator struct {
	Event *EtherfiAvsOperatorsManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerOwnershipTransferred)
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
		it.Event = new(EtherfiAvsOperatorsManagerOwnershipTransferred)
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
func (it *EtherfiAvsOperatorsManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EtherfiAvsOperatorsManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerOwnershipTransferredIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerOwnershipTransferred)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EtherfiAvsOperatorsManagerOwnershipTransferred, error) {
	event := new(EtherfiAvsOperatorsManagerOwnershipTransferred)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerPausedIterator struct {
	Event *EtherfiAvsOperatorsManagerPaused // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerPaused)
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
		it.Event = new(EtherfiAvsOperatorsManagerPaused)
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
func (it *EtherfiAvsOperatorsManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerPaused represents a Paused event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EtherfiAvsOperatorsManagerPausedIterator, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerPausedIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerPaused)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParsePaused(log types.Log) (*EtherfiAvsOperatorsManagerPaused, error) {
	event := new(EtherfiAvsOperatorsManagerPaused)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerRegisteredAsOperatorIterator is returned from FilterRegisteredAsOperator and is used to iterate over the raw logs and unpacked data for RegisteredAsOperator events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerRegisteredAsOperatorIterator struct {
	Event *EtherfiAvsOperatorsManagerRegisteredAsOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerRegisteredAsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerRegisteredAsOperator)
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
		it.Event = new(EtherfiAvsOperatorsManagerRegisteredAsOperator)
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
func (it *EtherfiAvsOperatorsManagerRegisteredAsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerRegisteredAsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerRegisteredAsOperator represents a RegisteredAsOperator event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerRegisteredAsOperator struct {
	Id     *big.Int
	Detail IDelegationManagerOperatorDetails
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisteredAsOperator is a free log retrieval operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterRegisteredAsOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerRegisteredAsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerRegisteredAsOperatorIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "RegisteredAsOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredAsOperator is a free log subscription operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchRegisteredAsOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerRegisteredAsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerRegisteredAsOperator)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseRegisteredAsOperator(log types.Log) (*EtherfiAvsOperatorsManagerRegisteredAsOperator, error) {
	event := new(EtherfiAvsOperatorsManagerRegisteredAsOperator)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator is returned from FilterRegisteredBlsKeyAsDelegatedNodeOperator and is used to iterate over the raw logs and unpacked data for RegisteredBlsKeyAsDelegatedNodeOperator events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator struct {
	Event *EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
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
		it.Event = new(EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
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
func (it *EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator represents a RegisteredBlsKeyAsDelegatedNodeOperator event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	QuorumNumbers          []byte
	Socket                 string
	Params                 IBLSApkRegistryPubkeyRegistrationParams
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRegisteredBlsKeyAsDelegatedNodeOperator is a free log retrieval operation binding the contract event 0x3bf95f161a4d8780e658d5f1476e87c49b56ff67c2ecec2baeadb4a474f64d51.
//
// Solidity: event RegisteredBlsKeyAsDelegatedNodeOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterRegisteredBlsKeyAsDelegatedNodeOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "RegisteredBlsKeyAsDelegatedNodeOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "RegisteredBlsKeyAsDelegatedNodeOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredBlsKeyAsDelegatedNodeOperator is a free log subscription operation binding the contract event 0x3bf95f161a4d8780e658d5f1476e87c49b56ff67c2ecec2baeadb4a474f64d51.
//
// Solidity: event RegisteredBlsKeyAsDelegatedNodeOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchRegisteredBlsKeyAsDelegatedNodeOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "RegisteredBlsKeyAsDelegatedNodeOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "RegisteredBlsKeyAsDelegatedNodeOperator", log); err != nil {
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

// ParseRegisteredBlsKeyAsDelegatedNodeOperator is a log parse operation binding the contract event 0x3bf95f161a4d8780e658d5f1476e87c49b56ff67c2ecec2baeadb4a474f64d51.
//
// Solidity: event RegisteredBlsKeyAsDelegatedNodeOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseRegisteredBlsKeyAsDelegatedNodeOperator(log types.Log) (*EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator, error) {
	event := new(EtherfiAvsOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "RegisteredBlsKeyAsDelegatedNodeOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerRegisteredOperatorIterator is returned from FilterRegisteredOperator and is used to iterate over the raw logs and unpacked data for RegisteredOperator events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerRegisteredOperatorIterator struct {
	Event *EtherfiAvsOperatorsManagerRegisteredOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerRegisteredOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerRegisteredOperator)
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
		it.Event = new(EtherfiAvsOperatorsManagerRegisteredOperator)
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
func (it *EtherfiAvsOperatorsManagerRegisteredOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerRegisteredOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerRegisteredOperator represents a RegisteredOperator event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerRegisteredOperator struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	QuorumNumbers          []byte
	Socket                 string
	Params                 IBLSApkRegistryPubkeyRegistrationParams
	OperatorSignature      ISignatureUtilsSignatureWithSaltAndExpiry
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRegisteredOperator is a free log retrieval operation binding the contract event 0x7fcf4f067d4a0674927732b3bfbbccaa47ef090c869bd6189e8dc842b46cc62c.
//
// Solidity: event RegisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterRegisteredOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerRegisteredOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "RegisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerRegisteredOperatorIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "RegisteredOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredOperator is a free log subscription operation binding the contract event 0x7fcf4f067d4a0674927732b3bfbbccaa47ef090c869bd6189e8dc842b46cc62c.
//
// Solidity: event RegisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchRegisteredOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerRegisteredOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "RegisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerRegisteredOperator)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "RegisteredOperator", log); err != nil {
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

// ParseRegisteredOperator is a log parse operation binding the contract event 0x7fcf4f067d4a0674927732b3bfbbccaa47ef090c869bd6189e8dc842b46cc62c.
//
// Solidity: event RegisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseRegisteredOperator(log types.Log) (*EtherfiAvsOperatorsManagerRegisteredOperator, error) {
	event := new(EtherfiAvsOperatorsManagerRegisteredOperator)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "RegisteredOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUnpausedIterator struct {
	Event *EtherfiAvsOperatorsManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerUnpaused)
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
		it.Event = new(EtherfiAvsOperatorsManagerUnpaused)
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
func (it *EtherfiAvsOperatorsManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerUnpaused represents a Unpaused event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EtherfiAvsOperatorsManagerUnpausedIterator, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerUnpausedIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerUnpaused)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseUnpaused(log types.Log) (*EtherfiAvsOperatorsManagerUnpaused, error) {
	event := new(EtherfiAvsOperatorsManagerUnpaused)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerUpdatedAvsNodeRunnerIterator is returned from FilterUpdatedAvsNodeRunner and is used to iterate over the raw logs and unpacked data for UpdatedAvsNodeRunner events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedAvsNodeRunnerIterator struct {
	Event *EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerUpdatedAvsNodeRunnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner)
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
		it.Event = new(EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner)
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
func (it *EtherfiAvsOperatorsManagerUpdatedAvsNodeRunnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerUpdatedAvsNodeRunnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner represents a UpdatedAvsNodeRunner event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner struct {
	Id            *big.Int
	AvsNodeRunner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvsNodeRunner is a free log retrieval operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterUpdatedAvsNodeRunner(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerUpdatedAvsNodeRunnerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerUpdatedAvsNodeRunnerIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "UpdatedAvsNodeRunner", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvsNodeRunner is a free log subscription operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchUpdatedAvsNodeRunner(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseUpdatedAvsNodeRunner(log types.Log) (*EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner, error) {
	event := new(EtherfiAvsOperatorsManagerUpdatedAvsNodeRunner)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerUpdatedAvsWhitelistIterator is returned from FilterUpdatedAvsWhitelist and is used to iterate over the raw logs and unpacked data for UpdatedAvsWhitelist events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedAvsWhitelistIterator struct {
	Event *EtherfiAvsOperatorsManagerUpdatedAvsWhitelist // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerUpdatedAvsWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerUpdatedAvsWhitelist)
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
		it.Event = new(EtherfiAvsOperatorsManagerUpdatedAvsWhitelist)
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
func (it *EtherfiAvsOperatorsManagerUpdatedAvsWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerUpdatedAvsWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerUpdatedAvsWhitelist represents a UpdatedAvsWhitelist event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedAvsWhitelist struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	IsWhitelisted          bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvsWhitelist is a free log retrieval operation binding the contract event 0x0cf8cb3148c96f3b418d80fc77f713f454ab17bfa4692f340ddd3057e65a78ac.
//
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsRegistryCoordinator, bool isWhitelisted)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterUpdatedAvsWhitelist(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerUpdatedAvsWhitelistIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "UpdatedAvsWhitelist", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerUpdatedAvsWhitelistIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "UpdatedAvsWhitelist", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvsWhitelist is a free log subscription operation binding the contract event 0x0cf8cb3148c96f3b418d80fc77f713f454ab17bfa4692f340ddd3057e65a78ac.
//
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsRegistryCoordinator, bool isWhitelisted)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchUpdatedAvsWhitelist(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerUpdatedAvsWhitelist, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "UpdatedAvsWhitelist", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerUpdatedAvsWhitelist)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedAvsWhitelist", log); err != nil {
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

// ParseUpdatedAvsWhitelist is a log parse operation binding the contract event 0x0cf8cb3148c96f3b418d80fc77f713f454ab17bfa4692f340ddd3057e65a78ac.
//
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsRegistryCoordinator, bool isWhitelisted)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseUpdatedAvsWhitelist(log types.Log) (*EtherfiAvsOperatorsManagerUpdatedAvsWhitelist, error) {
	event := new(EtherfiAvsOperatorsManagerUpdatedAvsWhitelist)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedAvsWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerUpdatedEcdsaSignerIterator is returned from FilterUpdatedEcdsaSigner and is used to iterate over the raw logs and unpacked data for UpdatedEcdsaSigner events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedEcdsaSignerIterator struct {
	Event *EtherfiAvsOperatorsManagerUpdatedEcdsaSigner // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerUpdatedEcdsaSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerUpdatedEcdsaSigner)
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
		it.Event = new(EtherfiAvsOperatorsManagerUpdatedEcdsaSigner)
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
func (it *EtherfiAvsOperatorsManagerUpdatedEcdsaSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerUpdatedEcdsaSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerUpdatedEcdsaSigner represents a UpdatedEcdsaSigner event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedEcdsaSigner struct {
	Id          *big.Int
	EcdsaSigner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEcdsaSigner is a free log retrieval operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterUpdatedEcdsaSigner(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerUpdatedEcdsaSignerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerUpdatedEcdsaSignerIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "UpdatedEcdsaSigner", logs: logs, sub: sub}, nil
}

// WatchUpdatedEcdsaSigner is a free log subscription operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchUpdatedEcdsaSigner(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerUpdatedEcdsaSigner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerUpdatedEcdsaSigner)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseUpdatedEcdsaSigner(log types.Log) (*EtherfiAvsOperatorsManagerUpdatedEcdsaSigner, error) {
	event := new(EtherfiAvsOperatorsManagerUpdatedEcdsaSigner)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURIIterator is returned from FilterUpdatedOperatorMetadataURI and is used to iterate over the raw logs and unpacked data for UpdatedOperatorMetadataURI events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURIIterator struct {
	Event *EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI)
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
		it.Event = new(EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI)
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
func (it *EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI represents a UpdatedOperatorMetadataURI event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI struct {
	Id          *big.Int
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOperatorMetadataURI is a free log retrieval operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterUpdatedOperatorMetadataURI(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURIIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "UpdatedOperatorMetadataURI", logs: logs, sub: sub}, nil
}

// WatchUpdatedOperatorMetadataURI is a free log subscription operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchUpdatedOperatorMetadataURI(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseUpdatedOperatorMetadataURI(log types.Log) (*EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI, error) {
	event := new(EtherfiAvsOperatorsManagerUpdatedOperatorMetadataURI)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerUpdatedSocketIterator is returned from FilterUpdatedSocket and is used to iterate over the raw logs and unpacked data for UpdatedSocket events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedSocketIterator struct {
	Event *EtherfiAvsOperatorsManagerUpdatedSocket // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerUpdatedSocketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerUpdatedSocket)
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
		it.Event = new(EtherfiAvsOperatorsManagerUpdatedSocket)
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
func (it *EtherfiAvsOperatorsManagerUpdatedSocketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerUpdatedSocketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerUpdatedSocket represents a UpdatedSocket event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpdatedSocket struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	Socket                 string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSocket is a free log retrieval operation binding the contract event 0x6e72bae16f3bc0863210747946d52680ce537fca1cfb695c91bdb9bf872966b9.
//
// Solidity: event UpdatedSocket(uint256 indexed id, address avsRegistryCoordinator, string socket)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterUpdatedSocket(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAvsOperatorsManagerUpdatedSocketIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "UpdatedSocket", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerUpdatedSocketIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "UpdatedSocket", logs: logs, sub: sub}, nil
}

// WatchUpdatedSocket is a free log subscription operation binding the contract event 0x6e72bae16f3bc0863210747946d52680ce537fca1cfb695c91bdb9bf872966b9.
//
// Solidity: event UpdatedSocket(uint256 indexed id, address avsRegistryCoordinator, string socket)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchUpdatedSocket(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerUpdatedSocket, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "UpdatedSocket", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerUpdatedSocket)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedSocket", log); err != nil {
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

// ParseUpdatedSocket is a log parse operation binding the contract event 0x6e72bae16f3bc0863210747946d52680ce537fca1cfb695c91bdb9bf872966b9.
//
// Solidity: event UpdatedSocket(uint256 indexed id, address avsRegistryCoordinator, string socket)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseUpdatedSocket(log types.Log) (*EtherfiAvsOperatorsManagerUpdatedSocket, error) {
	event := new(EtherfiAvsOperatorsManagerUpdatedSocket)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "UpdatedSocket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAvsOperatorsManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpgradedIterator struct {
	Event *EtherfiAvsOperatorsManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherfiAvsOperatorsManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAvsOperatorsManagerUpgraded)
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
		it.Event = new(EtherfiAvsOperatorsManagerUpgraded)
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
func (it *EtherfiAvsOperatorsManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAvsOperatorsManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAvsOperatorsManagerUpgraded represents a Upgraded event raised by the EtherfiAvsOperatorsManager contract.
type EtherfiAvsOperatorsManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EtherfiAvsOperatorsManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAvsOperatorsManagerUpgradedIterator{contract: _EtherfiAvsOperatorsManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EtherfiAvsOperatorsManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherfiAvsOperatorsManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAvsOperatorsManagerUpgraded)
				if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EtherfiAvsOperatorsManager *EtherfiAvsOperatorsManagerFilterer) ParseUpgraded(log types.Log) (*EtherfiAvsOperatorsManagerUpgraded, error) {
	event := new(EtherfiAvsOperatorsManagerUpgraded)
	if err := _EtherfiAvsOperatorsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
