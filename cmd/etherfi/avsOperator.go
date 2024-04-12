// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// AVSOperatorMetaData contains all meta data concerning the AVSOperator contract.
var AVSOperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"etherFiAvsOperator\",\"type\":\"address\"}],\"name\":\"CreatedEtherFiAvsOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"DeregisteredOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"ModifiedOperatorDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"detail\",\"type\":\"tuple\"}],\"name\":\"RegisteredAsOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"RegisteredBlsKeyAsDelegatedNodeOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"RegisteredOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsNodeRunner\",\"type\":\"address\"}],\"name\":\"UpdatedAvsNodeRunner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"name\":\"UpdatedAvsWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ecdsaSigner\",\"type\":\"address\"}],\"name\":\"UpdatedEcdsaSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"UpdatedOperatorMetadataURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsRegistryCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"UpdatedSocket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"avsOperators\",\"outputs\":[{\"internalType\":\"contractEtherFiAvsOperator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_etherFiAvsOperatorImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nums\",\"type\":\"uint256\"}],\"name\":\"instantiateEtherFiAvsOperator\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"modifyOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvsOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_detail\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_metaDataURI\",\"type\":\"string\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"registerBlsKeyAsDelegatedNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"name\":\"_operatorKickParams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_churnApproverSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorWithChurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsNodeRunner\",\"type\":\"address\"}],\"name\":\"updateAvsNodeRunner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isWhitelisted\",\"type\":\"bool\"}],\"name\":\"updateAvsWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_ecdsaSigner\",\"type\":\"address\"}],\"name\":\"updateEcdsaSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateOperatorMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"}],\"name\":\"updateSocket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradableBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeEtherFiAvsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AVSOperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSOperatorMetaData.ABI instead.
var AVSOperatorABI = AVSOperatorMetaData.ABI

// AVSOperator is an auto generated Go binding around an Ethereum contract.
type AVSOperator struct {
	AVSOperatorCaller     // Read-only binding to the contract
	AVSOperatorTransactor // Write-only binding to the contract
	AVSOperatorFilterer   // Log filterer for contract events
}

// AVSOperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSOperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSOperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSOperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSOperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSOperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSOperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSOperatorSession struct {
	Contract     *AVSOperator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AVSOperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSOperatorCallerSession struct {
	Contract *AVSOperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AVSOperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSOperatorTransactorSession struct {
	Contract     *AVSOperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AVSOperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSOperatorRaw struct {
	Contract *AVSOperator // Generic contract binding to access the raw methods on
}

// AVSOperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSOperatorCallerRaw struct {
	Contract *AVSOperatorCaller // Generic read-only contract binding to access the raw methods on
}

// AVSOperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSOperatorTransactorRaw struct {
	Contract *AVSOperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSOperator creates a new instance of AVSOperator, bound to a specific deployed contract.
func NewAVSOperator(address common.Address, backend bind.ContractBackend) (*AVSOperator, error) {
	contract, err := bindAVSOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSOperator{AVSOperatorCaller: AVSOperatorCaller{contract: contract}, AVSOperatorTransactor: AVSOperatorTransactor{contract: contract}, AVSOperatorFilterer: AVSOperatorFilterer{contract: contract}}, nil
}

// NewAVSOperatorCaller creates a new read-only instance of AVSOperator, bound to a specific deployed contract.
func NewAVSOperatorCaller(address common.Address, caller bind.ContractCaller) (*AVSOperatorCaller, error) {
	contract, err := bindAVSOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorCaller{contract: contract}, nil
}

// NewAVSOperatorTransactor creates a new write-only instance of AVSOperator, bound to a specific deployed contract.
func NewAVSOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSOperatorTransactor, error) {
	contract, err := bindAVSOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorTransactor{contract: contract}, nil
}

// NewAVSOperatorFilterer creates a new log filterer instance of AVSOperator, bound to a specific deployed contract.
func NewAVSOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSOperatorFilterer, error) {
	contract, err := bindAVSOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorFilterer{contract: contract}, nil
}

// bindAVSOperator binds a generic wrapper to an already deployed contract.
func bindAVSOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AVSOperatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSOperator *AVSOperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSOperator.Contract.AVSOperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSOperator *AVSOperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSOperator.Contract.AVSOperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSOperator *AVSOperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSOperator.Contract.AVSOperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSOperator *AVSOperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSOperator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSOperator *AVSOperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSOperator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSOperator *AVSOperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSOperator.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AVSOperator *AVSOperatorCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AVSOperator *AVSOperatorSession) Admins(arg0 common.Address) (bool, error) {
	return _AVSOperator.Contract.Admins(&_AVSOperator.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_AVSOperator *AVSOperatorCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _AVSOperator.Contract.Admins(&_AVSOperator.CallOpts, arg0)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_AVSOperator *AVSOperatorCaller) AvsOperators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "avsOperators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_AVSOperator *AVSOperatorSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _AVSOperator.Contract.AvsOperators(&_AVSOperator.CallOpts, arg0)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_AVSOperator *AVSOperatorCallerSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _AVSOperator.Contract.AvsOperators(&_AVSOperator.CallOpts, arg0)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_AVSOperator *AVSOperatorCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_AVSOperator *AVSOperatorSession) DelegationManager() (common.Address, error) {
	return _AVSOperator.Contract.DelegationManager(&_AVSOperator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_AVSOperator *AVSOperatorCallerSession) DelegationManager() (common.Address, error) {
	return _AVSOperator.Contract.DelegationManager(&_AVSOperator.CallOpts)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_AVSOperator *AVSOperatorCaller) NextAvsOperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "nextAvsOperatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_AVSOperator *AVSOperatorSession) NextAvsOperatorId() (*big.Int, error) {
	return _AVSOperator.Contract.NextAvsOperatorId(&_AVSOperator.CallOpts)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_AVSOperator *AVSOperatorCallerSession) NextAvsOperatorId() (*big.Int, error) {
	return _AVSOperator.Contract.NextAvsOperatorId(&_AVSOperator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSOperator *AVSOperatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSOperator *AVSOperatorSession) Owner() (common.Address, error) {
	return _AVSOperator.Contract.Owner(&_AVSOperator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AVSOperator *AVSOperatorCallerSession) Owner() (common.Address, error) {
	return _AVSOperator.Contract.Owner(&_AVSOperator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AVSOperator *AVSOperatorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AVSOperator *AVSOperatorSession) Paused() (bool, error) {
	return _AVSOperator.Contract.Paused(&_AVSOperator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AVSOperator *AVSOperatorCallerSession) Paused() (bool, error) {
	return _AVSOperator.Contract.Paused(&_AVSOperator.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_AVSOperator *AVSOperatorCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_AVSOperator *AVSOperatorSession) Pausers(arg0 common.Address) (bool, error) {
	return _AVSOperator.Contract.Pausers(&_AVSOperator.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_AVSOperator *AVSOperatorCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _AVSOperator.Contract.Pausers(&_AVSOperator.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AVSOperator *AVSOperatorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AVSOperator *AVSOperatorSession) ProxiableUUID() ([32]byte, error) {
	return _AVSOperator.Contract.ProxiableUUID(&_AVSOperator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AVSOperator *AVSOperatorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AVSOperator.Contract.ProxiableUUID(&_AVSOperator.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_AVSOperator *AVSOperatorCaller) UpgradableBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSOperator.contract.Call(opts, &out, "upgradableBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_AVSOperator *AVSOperatorSession) UpgradableBeacon() (common.Address, error) {
	return _AVSOperator.Contract.UpgradableBeacon(&_AVSOperator.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_AVSOperator *AVSOperatorCallerSession) UpgradableBeacon() (common.Address, error) {
	return _AVSOperator.Contract.UpgradableBeacon(&_AVSOperator.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_AVSOperator *AVSOperatorTransactor) DeregisterOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "deregisterOperator", _id, _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_AVSOperator *AVSOperatorSession) DeregisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _AVSOperator.Contract.DeregisterOperator(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_AVSOperator *AVSOperatorTransactorSession) DeregisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _AVSOperator.Contract.DeregisterOperator(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManager, address _etherFiAvsOperatorImpl) returns()
func (_AVSOperator *AVSOperatorTransactor) Initialize(opts *bind.TransactOpts, _delegationManager common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "initialize", _delegationManager, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManager, address _etherFiAvsOperatorImpl) returns()
func (_AVSOperator *AVSOperatorSession) Initialize(_delegationManager common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.Initialize(&_AVSOperator.TransactOpts, _delegationManager, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _delegationManager, address _etherFiAvsOperatorImpl) returns()
func (_AVSOperator *AVSOperatorTransactorSession) Initialize(_delegationManager common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.Initialize(&_AVSOperator.TransactOpts, _delegationManager, _etherFiAvsOperatorImpl)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_AVSOperator *AVSOperatorTransactor) InstantiateEtherFiAvsOperator(opts *bind.TransactOpts, _nums *big.Int) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "instantiateEtherFiAvsOperator", _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_AVSOperator *AVSOperatorSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _AVSOperator.Contract.InstantiateEtherFiAvsOperator(&_AVSOperator.TransactOpts, _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_AVSOperator *AVSOperatorTransactorSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _AVSOperator.Contract.InstantiateEtherFiAvsOperator(&_AVSOperator.TransactOpts, _nums)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_AVSOperator *AVSOperatorTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, _id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "modifyOperatorDetails", _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_AVSOperator *AVSOperatorSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AVSOperator.Contract.ModifyOperatorDetails(&_AVSOperator.TransactOpts, _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_AVSOperator *AVSOperatorTransactorSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _AVSOperator.Contract.ModifyOperatorDetails(&_AVSOperator.TransactOpts, _id, _newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AVSOperator *AVSOperatorTransactor) RegisterAsOperator(opts *bind.TransactOpts, _id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "registerAsOperator", _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AVSOperator *AVSOperatorSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterAsOperator(&_AVSOperator.TransactOpts, _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_AVSOperator *AVSOperatorTransactorSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterAsOperator(&_AVSOperator.TransactOpts, _id, _detail, _metaDataURI)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_AVSOperator *AVSOperatorTransactor) RegisterBlsKeyAsDelegatedNodeOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "registerBlsKeyAsDelegatedNodeOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_AVSOperator *AVSOperatorSession) RegisterBlsKeyAsDelegatedNodeOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_AVSOperator *AVSOperatorTransactorSession) RegisterBlsKeyAsDelegatedNodeOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_AVSOperator *AVSOperatorTransactor) RegisterOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "registerOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_AVSOperator *AVSOperatorSession) RegisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterOperator(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_AVSOperator *AVSOperatorTransactorSession) RegisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterOperator(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_AVSOperator *AVSOperatorTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "registerOperatorWithChurn", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_AVSOperator *AVSOperatorSession) RegisterOperatorWithChurn(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterOperatorWithChurn(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_AVSOperator *AVSOperatorTransactorSession) RegisterOperatorWithChurn(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _AVSOperator.Contract.RegisterOperatorWithChurn(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSOperator *AVSOperatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSOperator *AVSOperatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AVSOperator.Contract.RenounceOwnership(&_AVSOperator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AVSOperator *AVSOperatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AVSOperator.Contract.RenounceOwnership(&_AVSOperator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSOperator *AVSOperatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSOperator *AVSOperatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.TransferOwnership(&_AVSOperator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AVSOperator *AVSOperatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.TransferOwnership(&_AVSOperator.TransactOpts, newOwner)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AVSOperator *AVSOperatorTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AVSOperator *AVSOperatorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateAdmin(&_AVSOperator.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateAdmin(&_AVSOperator.TransactOpts, _address, _isAdmin)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_AVSOperator *AVSOperatorTransactor) UpdateAvsNodeRunner(opts *bind.TransactOpts, _id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "updateAvsNodeRunner", _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_AVSOperator *AVSOperatorSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateAvsNodeRunner(&_AVSOperator.TransactOpts, _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateAvsNodeRunner(&_AVSOperator.TransactOpts, _id, _avsNodeRunner)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_AVSOperator *AVSOperatorTransactor) UpdateAvsWhitelist(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "updateAvsWhitelist", _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_AVSOperator *AVSOperatorSession) UpdateAvsWhitelist(_id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateAvsWhitelist(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpdateAvsWhitelist(_id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateAvsWhitelist(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_AVSOperator *AVSOperatorTransactor) UpdateEcdsaSigner(opts *bind.TransactOpts, _id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "updateEcdsaSigner", _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_AVSOperator *AVSOperatorSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateEcdsaSigner(&_AVSOperator.TransactOpts, _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateEcdsaSigner(&_AVSOperator.TransactOpts, _id, _ecdsaSigner)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_AVSOperator *AVSOperatorTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, _id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "updateOperatorMetadataURI", _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_AVSOperator *AVSOperatorSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateOperatorMetadataURI(&_AVSOperator.TransactOpts, _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateOperatorMetadataURI(&_AVSOperator.TransactOpts, _id, _metadataURI)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_AVSOperator *AVSOperatorTransactor) UpdateSocket(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "updateSocket", _id, _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_AVSOperator *AVSOperatorSession) UpdateSocket(_id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateSocket(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpdateSocket(_id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpdateSocket(&_AVSOperator.TransactOpts, _id, _avsRegistryCoordinator, _socket)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_AVSOperator *AVSOperatorTransactor) UpgradeEtherFiAvsOperator(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "upgradeEtherFiAvsOperator", _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_AVSOperator *AVSOperatorSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpgradeEtherFiAvsOperator(&_AVSOperator.TransactOpts, _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpgradeEtherFiAvsOperator(&_AVSOperator.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AVSOperator *AVSOperatorTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AVSOperator *AVSOperatorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpgradeTo(&_AVSOperator.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpgradeTo(&_AVSOperator.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AVSOperator *AVSOperatorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AVSOperator.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AVSOperator *AVSOperatorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpgradeToAndCall(&_AVSOperator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AVSOperator *AVSOperatorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AVSOperator.Contract.UpgradeToAndCall(&_AVSOperator.TransactOpts, newImplementation, data)
}

// AVSOperatorAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AVSOperator contract.
type AVSOperatorAdminChangedIterator struct {
	Event *AVSOperatorAdminChanged // Event containing the contract specifics and raw log

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
func (it *AVSOperatorAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorAdminChanged)
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
		it.Event = new(AVSOperatorAdminChanged)
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
func (it *AVSOperatorAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorAdminChanged represents a AdminChanged event raised by the AVSOperator contract.
type AVSOperatorAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AVSOperator *AVSOperatorFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AVSOperatorAdminChangedIterator, error) {

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AVSOperatorAdminChangedIterator{contract: _AVSOperator.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AVSOperator *AVSOperatorFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AVSOperatorAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorAdminChanged)
				if err := _AVSOperator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseAdminChanged(log types.Log) (*AVSOperatorAdminChanged, error) {
	event := new(AVSOperatorAdminChanged)
	if err := _AVSOperator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AVSOperator contract.
type AVSOperatorBeaconUpgradedIterator struct {
	Event *AVSOperatorBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AVSOperatorBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorBeaconUpgraded)
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
		it.Event = new(AVSOperatorBeaconUpgraded)
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
func (it *AVSOperatorBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorBeaconUpgraded represents a BeaconUpgraded event raised by the AVSOperator contract.
type AVSOperatorBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AVSOperator *AVSOperatorFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AVSOperatorBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorBeaconUpgradedIterator{contract: _AVSOperator.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AVSOperator *AVSOperatorFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AVSOperatorBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorBeaconUpgraded)
				if err := _AVSOperator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseBeaconUpgraded(log types.Log) (*AVSOperatorBeaconUpgraded, error) {
	event := new(AVSOperatorBeaconUpgraded)
	if err := _AVSOperator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorCreatedEtherFiAvsOperatorIterator is returned from FilterCreatedEtherFiAvsOperator and is used to iterate over the raw logs and unpacked data for CreatedEtherFiAvsOperator events raised by the AVSOperator contract.
type AVSOperatorCreatedEtherFiAvsOperatorIterator struct {
	Event *AVSOperatorCreatedEtherFiAvsOperator // Event containing the contract specifics and raw log

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
func (it *AVSOperatorCreatedEtherFiAvsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorCreatedEtherFiAvsOperator)
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
		it.Event = new(AVSOperatorCreatedEtherFiAvsOperator)
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
func (it *AVSOperatorCreatedEtherFiAvsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorCreatedEtherFiAvsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorCreatedEtherFiAvsOperator represents a CreatedEtherFiAvsOperator event raised by the AVSOperator contract.
type AVSOperatorCreatedEtherFiAvsOperator struct {
	Id                 *big.Int
	EtherFiAvsOperator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCreatedEtherFiAvsOperator is a free log retrieval operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_AVSOperator *AVSOperatorFilterer) FilterCreatedEtherFiAvsOperator(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorCreatedEtherFiAvsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorCreatedEtherFiAvsOperatorIterator{contract: _AVSOperator.contract, event: "CreatedEtherFiAvsOperator", logs: logs, sub: sub}, nil
}

// WatchCreatedEtherFiAvsOperator is a free log subscription operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_AVSOperator *AVSOperatorFilterer) WatchCreatedEtherFiAvsOperator(opts *bind.WatchOpts, sink chan<- *AVSOperatorCreatedEtherFiAvsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorCreatedEtherFiAvsOperator)
				if err := _AVSOperator.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseCreatedEtherFiAvsOperator(log types.Log) (*AVSOperatorCreatedEtherFiAvsOperator, error) {
	event := new(AVSOperatorCreatedEtherFiAvsOperator)
	if err := _AVSOperator.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorDeregisteredOperatorIterator is returned from FilterDeregisteredOperator and is used to iterate over the raw logs and unpacked data for DeregisteredOperator events raised by the AVSOperator contract.
type AVSOperatorDeregisteredOperatorIterator struct {
	Event *AVSOperatorDeregisteredOperator // Event containing the contract specifics and raw log

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
func (it *AVSOperatorDeregisteredOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorDeregisteredOperator)
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
		it.Event = new(AVSOperatorDeregisteredOperator)
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
func (it *AVSOperatorDeregisteredOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorDeregisteredOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorDeregisteredOperator represents a DeregisteredOperator event raised by the AVSOperator contract.
type AVSOperatorDeregisteredOperator struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	QuorumNumbers          []byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDeregisteredOperator is a free log retrieval operation binding the contract event 0x1376d0cba684c4d21991c950915a2bf8a9181c389d343993ff185ba5944b5a2d.
//
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers)
func (_AVSOperator *AVSOperatorFilterer) FilterDeregisteredOperator(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorDeregisteredOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "DeregisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorDeregisteredOperatorIterator{contract: _AVSOperator.contract, event: "DeregisteredOperator", logs: logs, sub: sub}, nil
}

// WatchDeregisteredOperator is a free log subscription operation binding the contract event 0x1376d0cba684c4d21991c950915a2bf8a9181c389d343993ff185ba5944b5a2d.
//
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers)
func (_AVSOperator *AVSOperatorFilterer) WatchDeregisteredOperator(opts *bind.WatchOpts, sink chan<- *AVSOperatorDeregisteredOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "DeregisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorDeregisteredOperator)
				if err := _AVSOperator.contract.UnpackLog(event, "DeregisteredOperator", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseDeregisteredOperator(log types.Log) (*AVSOperatorDeregisteredOperator, error) {
	event := new(AVSOperatorDeregisteredOperator)
	if err := _AVSOperator.contract.UnpackLog(event, "DeregisteredOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AVSOperator contract.
type AVSOperatorInitializedIterator struct {
	Event *AVSOperatorInitialized // Event containing the contract specifics and raw log

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
func (it *AVSOperatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorInitialized)
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
		it.Event = new(AVSOperatorInitialized)
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
func (it *AVSOperatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorInitialized represents a Initialized event raised by the AVSOperator contract.
type AVSOperatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSOperator *AVSOperatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*AVSOperatorInitializedIterator, error) {

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AVSOperatorInitializedIterator{contract: _AVSOperator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSOperator *AVSOperatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AVSOperatorInitialized) (event.Subscription, error) {

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorInitialized)
				if err := _AVSOperator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseInitialized(log types.Log) (*AVSOperatorInitialized, error) {
	event := new(AVSOperatorInitialized)
	if err := _AVSOperator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorModifiedOperatorDetailsIterator is returned from FilterModifiedOperatorDetails and is used to iterate over the raw logs and unpacked data for ModifiedOperatorDetails events raised by the AVSOperator contract.
type AVSOperatorModifiedOperatorDetailsIterator struct {
	Event *AVSOperatorModifiedOperatorDetails // Event containing the contract specifics and raw log

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
func (it *AVSOperatorModifiedOperatorDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorModifiedOperatorDetails)
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
		it.Event = new(AVSOperatorModifiedOperatorDetails)
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
func (it *AVSOperatorModifiedOperatorDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorModifiedOperatorDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorModifiedOperatorDetails represents a ModifiedOperatorDetails event raised by the AVSOperator contract.
type AVSOperatorModifiedOperatorDetails struct {
	Id                 *big.Int
	NewOperatorDetails IDelegationManagerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterModifiedOperatorDetails is a free log retrieval operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_AVSOperator *AVSOperatorFilterer) FilterModifiedOperatorDetails(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorModifiedOperatorDetailsIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorModifiedOperatorDetailsIterator{contract: _AVSOperator.contract, event: "ModifiedOperatorDetails", logs: logs, sub: sub}, nil
}

// WatchModifiedOperatorDetails is a free log subscription operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_AVSOperator *AVSOperatorFilterer) WatchModifiedOperatorDetails(opts *bind.WatchOpts, sink chan<- *AVSOperatorModifiedOperatorDetails, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorModifiedOperatorDetails)
				if err := _AVSOperator.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseModifiedOperatorDetails(log types.Log) (*AVSOperatorModifiedOperatorDetails, error) {
	event := new(AVSOperatorModifiedOperatorDetails)
	if err := _AVSOperator.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AVSOperator contract.
type AVSOperatorOwnershipTransferredIterator struct {
	Event *AVSOperatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AVSOperatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorOwnershipTransferred)
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
		it.Event = new(AVSOperatorOwnershipTransferred)
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
func (it *AVSOperatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorOwnershipTransferred represents a OwnershipTransferred event raised by the AVSOperator contract.
type AVSOperatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AVSOperator *AVSOperatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AVSOperatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorOwnershipTransferredIterator{contract: _AVSOperator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AVSOperator *AVSOperatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AVSOperatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorOwnershipTransferred)
				if err := _AVSOperator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseOwnershipTransferred(log types.Log) (*AVSOperatorOwnershipTransferred, error) {
	event := new(AVSOperatorOwnershipTransferred)
	if err := _AVSOperator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AVSOperator contract.
type AVSOperatorPausedIterator struct {
	Event *AVSOperatorPaused // Event containing the contract specifics and raw log

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
func (it *AVSOperatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorPaused)
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
		it.Event = new(AVSOperatorPaused)
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
func (it *AVSOperatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorPaused represents a Paused event raised by the AVSOperator contract.
type AVSOperatorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AVSOperator *AVSOperatorFilterer) FilterPaused(opts *bind.FilterOpts) (*AVSOperatorPausedIterator, error) {

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AVSOperatorPausedIterator{contract: _AVSOperator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AVSOperator *AVSOperatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AVSOperatorPaused) (event.Subscription, error) {

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorPaused)
				if err := _AVSOperator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParsePaused(log types.Log) (*AVSOperatorPaused, error) {
	event := new(AVSOperatorPaused)
	if err := _AVSOperator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorRegisteredAsOperatorIterator is returned from FilterRegisteredAsOperator and is used to iterate over the raw logs and unpacked data for RegisteredAsOperator events raised by the AVSOperator contract.
type AVSOperatorRegisteredAsOperatorIterator struct {
	Event *AVSOperatorRegisteredAsOperator // Event containing the contract specifics and raw log

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
func (it *AVSOperatorRegisteredAsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorRegisteredAsOperator)
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
		it.Event = new(AVSOperatorRegisteredAsOperator)
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
func (it *AVSOperatorRegisteredAsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorRegisteredAsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorRegisteredAsOperator represents a RegisteredAsOperator event raised by the AVSOperator contract.
type AVSOperatorRegisteredAsOperator struct {
	Id     *big.Int
	Detail IDelegationManagerOperatorDetails
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisteredAsOperator is a free log retrieval operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_AVSOperator *AVSOperatorFilterer) FilterRegisteredAsOperator(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorRegisteredAsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorRegisteredAsOperatorIterator{contract: _AVSOperator.contract, event: "RegisteredAsOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredAsOperator is a free log subscription operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_AVSOperator *AVSOperatorFilterer) WatchRegisteredAsOperator(opts *bind.WatchOpts, sink chan<- *AVSOperatorRegisteredAsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorRegisteredAsOperator)
				if err := _AVSOperator.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseRegisteredAsOperator(log types.Log) (*AVSOperatorRegisteredAsOperator, error) {
	event := new(AVSOperatorRegisteredAsOperator)
	if err := _AVSOperator.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperatorIterator is returned from FilterRegisteredBlsKeyAsDelegatedNodeOperator and is used to iterate over the raw logs and unpacked data for RegisteredBlsKeyAsDelegatedNodeOperator events raised by the AVSOperator contract.
type AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperatorIterator struct {
	Event *AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator // Event containing the contract specifics and raw log

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
func (it *AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator)
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
		it.Event = new(AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator)
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
func (it *AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator represents a RegisteredBlsKeyAsDelegatedNodeOperator event raised by the AVSOperator contract.
type AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator struct {
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
func (_AVSOperator *AVSOperatorFilterer) FilterRegisteredBlsKeyAsDelegatedNodeOperator(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "RegisteredBlsKeyAsDelegatedNodeOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperatorIterator{contract: _AVSOperator.contract, event: "RegisteredBlsKeyAsDelegatedNodeOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredBlsKeyAsDelegatedNodeOperator is a free log subscription operation binding the contract event 0x3bf95f161a4d8780e658d5f1476e87c49b56ff67c2ecec2baeadb4a474f64d51.
//
// Solidity: event RegisteredBlsKeyAsDelegatedNodeOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params)
func (_AVSOperator *AVSOperatorFilterer) WatchRegisteredBlsKeyAsDelegatedNodeOperator(opts *bind.WatchOpts, sink chan<- *AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "RegisteredBlsKeyAsDelegatedNodeOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator)
				if err := _AVSOperator.contract.UnpackLog(event, "RegisteredBlsKeyAsDelegatedNodeOperator", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseRegisteredBlsKeyAsDelegatedNodeOperator(log types.Log) (*AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator, error) {
	event := new(AVSOperatorRegisteredBlsKeyAsDelegatedNodeOperator)
	if err := _AVSOperator.contract.UnpackLog(event, "RegisteredBlsKeyAsDelegatedNodeOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorRegisteredOperatorIterator is returned from FilterRegisteredOperator and is used to iterate over the raw logs and unpacked data for RegisteredOperator events raised by the AVSOperator contract.
type AVSOperatorRegisteredOperatorIterator struct {
	Event *AVSOperatorRegisteredOperator // Event containing the contract specifics and raw log

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
func (it *AVSOperatorRegisteredOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorRegisteredOperator)
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
		it.Event = new(AVSOperatorRegisteredOperator)
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
func (it *AVSOperatorRegisteredOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorRegisteredOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorRegisteredOperator represents a RegisteredOperator event raised by the AVSOperator contract.
type AVSOperatorRegisteredOperator struct {
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
func (_AVSOperator *AVSOperatorFilterer) FilterRegisteredOperator(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorRegisteredOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "RegisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorRegisteredOperatorIterator{contract: _AVSOperator.contract, event: "RegisteredOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredOperator is a free log subscription operation binding the contract event 0x7fcf4f067d4a0674927732b3bfbbccaa47ef090c869bd6189e8dc842b46cc62c.
//
// Solidity: event RegisteredOperator(uint256 indexed id, address avsRegistryCoordinator, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature)
func (_AVSOperator *AVSOperatorFilterer) WatchRegisteredOperator(opts *bind.WatchOpts, sink chan<- *AVSOperatorRegisteredOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "RegisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorRegisteredOperator)
				if err := _AVSOperator.contract.UnpackLog(event, "RegisteredOperator", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseRegisteredOperator(log types.Log) (*AVSOperatorRegisteredOperator, error) {
	event := new(AVSOperatorRegisteredOperator)
	if err := _AVSOperator.contract.UnpackLog(event, "RegisteredOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AVSOperator contract.
type AVSOperatorUnpausedIterator struct {
	Event *AVSOperatorUnpaused // Event containing the contract specifics and raw log

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
func (it *AVSOperatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorUnpaused)
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
		it.Event = new(AVSOperatorUnpaused)
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
func (it *AVSOperatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorUnpaused represents a Unpaused event raised by the AVSOperator contract.
type AVSOperatorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AVSOperator *AVSOperatorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AVSOperatorUnpausedIterator, error) {

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AVSOperatorUnpausedIterator{contract: _AVSOperator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AVSOperator *AVSOperatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AVSOperatorUnpaused) (event.Subscription, error) {

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorUnpaused)
				if err := _AVSOperator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseUnpaused(log types.Log) (*AVSOperatorUnpaused, error) {
	event := new(AVSOperatorUnpaused)
	if err := _AVSOperator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorUpdatedAvsNodeRunnerIterator is returned from FilterUpdatedAvsNodeRunner and is used to iterate over the raw logs and unpacked data for UpdatedAvsNodeRunner events raised by the AVSOperator contract.
type AVSOperatorUpdatedAvsNodeRunnerIterator struct {
	Event *AVSOperatorUpdatedAvsNodeRunner // Event containing the contract specifics and raw log

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
func (it *AVSOperatorUpdatedAvsNodeRunnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorUpdatedAvsNodeRunner)
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
		it.Event = new(AVSOperatorUpdatedAvsNodeRunner)
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
func (it *AVSOperatorUpdatedAvsNodeRunnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorUpdatedAvsNodeRunnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorUpdatedAvsNodeRunner represents a UpdatedAvsNodeRunner event raised by the AVSOperator contract.
type AVSOperatorUpdatedAvsNodeRunner struct {
	Id            *big.Int
	AvsNodeRunner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvsNodeRunner is a free log retrieval operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_AVSOperator *AVSOperatorFilterer) FilterUpdatedAvsNodeRunner(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorUpdatedAvsNodeRunnerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorUpdatedAvsNodeRunnerIterator{contract: _AVSOperator.contract, event: "UpdatedAvsNodeRunner", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvsNodeRunner is a free log subscription operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_AVSOperator *AVSOperatorFilterer) WatchUpdatedAvsNodeRunner(opts *bind.WatchOpts, sink chan<- *AVSOperatorUpdatedAvsNodeRunner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorUpdatedAvsNodeRunner)
				if err := _AVSOperator.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseUpdatedAvsNodeRunner(log types.Log) (*AVSOperatorUpdatedAvsNodeRunner, error) {
	event := new(AVSOperatorUpdatedAvsNodeRunner)
	if err := _AVSOperator.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorUpdatedAvsWhitelistIterator is returned from FilterUpdatedAvsWhitelist and is used to iterate over the raw logs and unpacked data for UpdatedAvsWhitelist events raised by the AVSOperator contract.
type AVSOperatorUpdatedAvsWhitelistIterator struct {
	Event *AVSOperatorUpdatedAvsWhitelist // Event containing the contract specifics and raw log

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
func (it *AVSOperatorUpdatedAvsWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorUpdatedAvsWhitelist)
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
		it.Event = new(AVSOperatorUpdatedAvsWhitelist)
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
func (it *AVSOperatorUpdatedAvsWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorUpdatedAvsWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorUpdatedAvsWhitelist represents a UpdatedAvsWhitelist event raised by the AVSOperator contract.
type AVSOperatorUpdatedAvsWhitelist struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	IsWhitelisted          bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvsWhitelist is a free log retrieval operation binding the contract event 0x0cf8cb3148c96f3b418d80fc77f713f454ab17bfa4692f340ddd3057e65a78ac.
//
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsRegistryCoordinator, bool isWhitelisted)
func (_AVSOperator *AVSOperatorFilterer) FilterUpdatedAvsWhitelist(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorUpdatedAvsWhitelistIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "UpdatedAvsWhitelist", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorUpdatedAvsWhitelistIterator{contract: _AVSOperator.contract, event: "UpdatedAvsWhitelist", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvsWhitelist is a free log subscription operation binding the contract event 0x0cf8cb3148c96f3b418d80fc77f713f454ab17bfa4692f340ddd3057e65a78ac.
//
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsRegistryCoordinator, bool isWhitelisted)
func (_AVSOperator *AVSOperatorFilterer) WatchUpdatedAvsWhitelist(opts *bind.WatchOpts, sink chan<- *AVSOperatorUpdatedAvsWhitelist, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "UpdatedAvsWhitelist", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorUpdatedAvsWhitelist)
				if err := _AVSOperator.contract.UnpackLog(event, "UpdatedAvsWhitelist", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseUpdatedAvsWhitelist(log types.Log) (*AVSOperatorUpdatedAvsWhitelist, error) {
	event := new(AVSOperatorUpdatedAvsWhitelist)
	if err := _AVSOperator.contract.UnpackLog(event, "UpdatedAvsWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorUpdatedEcdsaSignerIterator is returned from FilterUpdatedEcdsaSigner and is used to iterate over the raw logs and unpacked data for UpdatedEcdsaSigner events raised by the AVSOperator contract.
type AVSOperatorUpdatedEcdsaSignerIterator struct {
	Event *AVSOperatorUpdatedEcdsaSigner // Event containing the contract specifics and raw log

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
func (it *AVSOperatorUpdatedEcdsaSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorUpdatedEcdsaSigner)
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
		it.Event = new(AVSOperatorUpdatedEcdsaSigner)
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
func (it *AVSOperatorUpdatedEcdsaSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorUpdatedEcdsaSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorUpdatedEcdsaSigner represents a UpdatedEcdsaSigner event raised by the AVSOperator contract.
type AVSOperatorUpdatedEcdsaSigner struct {
	Id          *big.Int
	EcdsaSigner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEcdsaSigner is a free log retrieval operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_AVSOperator *AVSOperatorFilterer) FilterUpdatedEcdsaSigner(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorUpdatedEcdsaSignerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorUpdatedEcdsaSignerIterator{contract: _AVSOperator.contract, event: "UpdatedEcdsaSigner", logs: logs, sub: sub}, nil
}

// WatchUpdatedEcdsaSigner is a free log subscription operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_AVSOperator *AVSOperatorFilterer) WatchUpdatedEcdsaSigner(opts *bind.WatchOpts, sink chan<- *AVSOperatorUpdatedEcdsaSigner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorUpdatedEcdsaSigner)
				if err := _AVSOperator.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseUpdatedEcdsaSigner(log types.Log) (*AVSOperatorUpdatedEcdsaSigner, error) {
	event := new(AVSOperatorUpdatedEcdsaSigner)
	if err := _AVSOperator.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorUpdatedOperatorMetadataURIIterator is returned from FilterUpdatedOperatorMetadataURI and is used to iterate over the raw logs and unpacked data for UpdatedOperatorMetadataURI events raised by the AVSOperator contract.
type AVSOperatorUpdatedOperatorMetadataURIIterator struct {
	Event *AVSOperatorUpdatedOperatorMetadataURI // Event containing the contract specifics and raw log

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
func (it *AVSOperatorUpdatedOperatorMetadataURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorUpdatedOperatorMetadataURI)
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
		it.Event = new(AVSOperatorUpdatedOperatorMetadataURI)
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
func (it *AVSOperatorUpdatedOperatorMetadataURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorUpdatedOperatorMetadataURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorUpdatedOperatorMetadataURI represents a UpdatedOperatorMetadataURI event raised by the AVSOperator contract.
type AVSOperatorUpdatedOperatorMetadataURI struct {
	Id          *big.Int
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOperatorMetadataURI is a free log retrieval operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_AVSOperator *AVSOperatorFilterer) FilterUpdatedOperatorMetadataURI(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorUpdatedOperatorMetadataURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorUpdatedOperatorMetadataURIIterator{contract: _AVSOperator.contract, event: "UpdatedOperatorMetadataURI", logs: logs, sub: sub}, nil
}

// WatchUpdatedOperatorMetadataURI is a free log subscription operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_AVSOperator *AVSOperatorFilterer) WatchUpdatedOperatorMetadataURI(opts *bind.WatchOpts, sink chan<- *AVSOperatorUpdatedOperatorMetadataURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorUpdatedOperatorMetadataURI)
				if err := _AVSOperator.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseUpdatedOperatorMetadataURI(log types.Log) (*AVSOperatorUpdatedOperatorMetadataURI, error) {
	event := new(AVSOperatorUpdatedOperatorMetadataURI)
	if err := _AVSOperator.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorUpdatedSocketIterator is returned from FilterUpdatedSocket and is used to iterate over the raw logs and unpacked data for UpdatedSocket events raised by the AVSOperator contract.
type AVSOperatorUpdatedSocketIterator struct {
	Event *AVSOperatorUpdatedSocket // Event containing the contract specifics and raw log

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
func (it *AVSOperatorUpdatedSocketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorUpdatedSocket)
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
		it.Event = new(AVSOperatorUpdatedSocket)
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
func (it *AVSOperatorUpdatedSocketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorUpdatedSocketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorUpdatedSocket represents a UpdatedSocket event raised by the AVSOperator contract.
type AVSOperatorUpdatedSocket struct {
	Id                     *big.Int
	AvsRegistryCoordinator common.Address
	Socket                 string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSocket is a free log retrieval operation binding the contract event 0x6e72bae16f3bc0863210747946d52680ce537fca1cfb695c91bdb9bf872966b9.
//
// Solidity: event UpdatedSocket(uint256 indexed id, address avsRegistryCoordinator, string socket)
func (_AVSOperator *AVSOperatorFilterer) FilterUpdatedSocket(opts *bind.FilterOpts, id []*big.Int) (*AVSOperatorUpdatedSocketIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "UpdatedSocket", idRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorUpdatedSocketIterator{contract: _AVSOperator.contract, event: "UpdatedSocket", logs: logs, sub: sub}, nil
}

// WatchUpdatedSocket is a free log subscription operation binding the contract event 0x6e72bae16f3bc0863210747946d52680ce537fca1cfb695c91bdb9bf872966b9.
//
// Solidity: event UpdatedSocket(uint256 indexed id, address avsRegistryCoordinator, string socket)
func (_AVSOperator *AVSOperatorFilterer) WatchUpdatedSocket(opts *bind.WatchOpts, sink chan<- *AVSOperatorUpdatedSocket, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "UpdatedSocket", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorUpdatedSocket)
				if err := _AVSOperator.contract.UnpackLog(event, "UpdatedSocket", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseUpdatedSocket(log types.Log) (*AVSOperatorUpdatedSocket, error) {
	event := new(AVSOperatorUpdatedSocket)
	if err := _AVSOperator.contract.UnpackLog(event, "UpdatedSocket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSOperatorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AVSOperator contract.
type AVSOperatorUpgradedIterator struct {
	Event *AVSOperatorUpgraded // Event containing the contract specifics and raw log

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
func (it *AVSOperatorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSOperatorUpgraded)
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
		it.Event = new(AVSOperatorUpgraded)
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
func (it *AVSOperatorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSOperatorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSOperatorUpgraded represents a Upgraded event raised by the AVSOperator contract.
type AVSOperatorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AVSOperator *AVSOperatorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AVSOperatorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AVSOperator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AVSOperatorUpgradedIterator{contract: _AVSOperator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AVSOperator *AVSOperatorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AVSOperatorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AVSOperator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSOperatorUpgraded)
				if err := _AVSOperator.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AVSOperator *AVSOperatorFilterer) ParseUpgraded(log types.Log) (*AVSOperatorUpgraded, error) {
	event := new(AVSOperatorUpgraded)
	if err := _AVSOperator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
