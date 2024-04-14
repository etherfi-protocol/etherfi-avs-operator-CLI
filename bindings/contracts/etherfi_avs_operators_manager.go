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

// EtherfiAVSOperatorsManagerMetaData contains all meta data concerning the EtherfiAVSOperatorsManager contract.
var EtherfiAVSOperatorsManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"etherFiAvsOperator\",\"type\":\"address\"}],\"name\":\"CreatedEtherFiAvsOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsServiceManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"DeregisteredOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ForwardedRunnerCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"ModifiedOperatorDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"detail\",\"type\":\"tuple\"}],\"name\":\"RegisteredAsOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsServiceManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"RegisteredBlsKeyAsDelegatedNodeOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsServiceManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"RegisteredOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsNodeRunner\",\"type\":\"address\"}],\"name\":\"UpdatedAvsNodeRunner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsServiceManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"name\":\"UpdatedAvsWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ecdsaSigner\",\"type\":\"address\"}],\"name\":\"UpdatedEcdsaSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"UpdatedOperatorMetadataURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"avsServiceManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"UpdatedSocket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avsDirectory\",\"outputs\":[{\"internalType\":\"contractIAVSDirectory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"avsNodeRunner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsServiceManager\",\"type\":\"address\"}],\"name\":\"avsOperatorStatus\",\"outputs\":[{\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"avsOperators\",\"outputs\":[{\"internalType\":\"contractEtherFiAvsOperator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsServiceManager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"}],\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationManager\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"deregisterOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"ecdsaSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"}],\"name\":\"getAvsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"internalType\":\"structEtherFiAvsOperator.AvsInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_avsDirectory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_etherFiAvsOperatorImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_avsDirectory\",\"type\":\"address\"}],\"name\":\"initializeAvsDirectory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nums\",\"type\":\"uint256\"}],\"name\":\"instantiateEtherFiAvsOperator\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"}],\"name\":\"isAvsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"}],\"name\":\"isAvsWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"isRegisteredBlsKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_newOperatorDetails\",\"type\":\"tuple\"}],\"name\":\"modifyOperatorDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvsOperatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"operatorDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"earningsReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegationApprover\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakerOptOutWindowBlocks\",\"type\":\"uint32\"}],\"internalType\":\"structIDelegationManager.OperatorDetails\",\"name\":\"_detail\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_metaDataURI\",\"type\":\"string\"}],\"name\":\"registerAsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"registerBlsKeyAsDelegatedNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"internalType\":\"structIBLSApkRegistry.PubkeyRegistrationParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIRegistryCoordinator.OperatorKickParam[]\",\"name\":\"_operatorKickParams\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_churnApproverSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"_operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorWithChurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"_remainingCalldata\",\"type\":\"bytes\"}],\"name\":\"runnerForwardCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isAdmin\",\"type\":\"bool\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsNodeRunner\",\"type\":\"address\"}],\"name\":\"updateAvsNodeRunner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isWhitelisted\",\"type\":\"bool\"}],\"name\":\"updateAvsWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_ecdsaSigner\",\"type\":\"address\"}],\"name\":\"updateEcdsaSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"updateOperatorMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_avsRegistryCoordinator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_socket\",\"type\":\"string\"}],\"name\":\"updateSocket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradableBeacon\",\"outputs\":[{\"internalType\":\"contractUpgradeableBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeEtherFiAvsOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// EtherfiAVSOperatorsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherfiAVSOperatorsManagerMetaData.ABI instead.
var EtherfiAVSOperatorsManagerABI = EtherfiAVSOperatorsManagerMetaData.ABI

// EtherfiAVSOperatorsManager is an auto generated Go binding around an Ethereum contract.
type EtherfiAVSOperatorsManager struct {
	EtherfiAVSOperatorsManagerCaller     // Read-only binding to the contract
	EtherfiAVSOperatorsManagerTransactor // Write-only binding to the contract
	EtherfiAVSOperatorsManagerFilterer   // Log filterer for contract events
}

// EtherfiAVSOperatorsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAVSOperatorsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAVSOperatorsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherfiAVSOperatorsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAVSOperatorsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherfiAVSOperatorsManagerSession struct {
	Contract     *EtherfiAVSOperatorsManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EtherfiAVSOperatorsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherfiAVSOperatorsManagerCallerSession struct {
	Contract *EtherfiAVSOperatorsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// EtherfiAVSOperatorsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherfiAVSOperatorsManagerTransactorSession struct {
	Contract     *EtherfiAVSOperatorsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// EtherfiAVSOperatorsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherfiAVSOperatorsManagerRaw struct {
	Contract *EtherfiAVSOperatorsManager // Generic contract binding to access the raw methods on
}

// EtherfiAVSOperatorsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorsManagerCallerRaw struct {
	Contract *EtherfiAVSOperatorsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EtherfiAVSOperatorsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherfiAVSOperatorsManagerTransactorRaw struct {
	Contract *EtherfiAVSOperatorsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherfiAVSOperatorsManager creates a new instance of EtherfiAVSOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAVSOperatorsManager(address common.Address, backend bind.ContractBackend) (*EtherfiAVSOperatorsManager, error) {
	contract, err := bindEtherfiAVSOperatorsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManager{EtherfiAVSOperatorsManagerCaller: EtherfiAVSOperatorsManagerCaller{contract: contract}, EtherfiAVSOperatorsManagerTransactor: EtherfiAVSOperatorsManagerTransactor{contract: contract}, EtherfiAVSOperatorsManagerFilterer: EtherfiAVSOperatorsManagerFilterer{contract: contract}}, nil
}

// NewEtherfiAVSOperatorsManagerCaller creates a new read-only instance of EtherfiAVSOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAVSOperatorsManagerCaller(address common.Address, caller bind.ContractCaller) (*EtherfiAVSOperatorsManagerCaller, error) {
	contract, err := bindEtherfiAVSOperatorsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerCaller{contract: contract}, nil
}

// NewEtherfiAVSOperatorsManagerTransactor creates a new write-only instance of EtherfiAVSOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAVSOperatorsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherfiAVSOperatorsManagerTransactor, error) {
	contract, err := bindEtherfiAVSOperatorsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerTransactor{contract: contract}, nil
}

// NewEtherfiAVSOperatorsManagerFilterer creates a new log filterer instance of EtherfiAVSOperatorsManager, bound to a specific deployed contract.
func NewEtherfiAVSOperatorsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherfiAVSOperatorsManagerFilterer, error) {
	contract, err := bindEtherfiAVSOperatorsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerFilterer{contract: contract}, nil
}

// bindEtherfiAVSOperatorsManager binds a generic wrapper to an already deployed contract.
func bindEtherfiAVSOperatorsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherfiAVSOperatorsManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAVSOperatorsManager.Contract.EtherfiAVSOperatorsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.EtherfiAVSOperatorsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.EtherfiAVSOperatorsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAVSOperatorsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.Admins(&_EtherfiAVSOperatorsManager.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.Admins(&_EtherfiAVSOperatorsManager.CallOpts, arg0)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) AvsDirectory() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsDirectory(&_EtherfiAVSOperatorsManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsDirectory(&_EtherfiAVSOperatorsManager.CallOpts)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0x05dba646.
//
// Solidity: function avsNodeRunner(uint256 _id) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) AvsNodeRunner(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "avsNodeRunner", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsNodeRunner is a free data retrieval call binding the contract method 0x05dba646.
//
// Solidity: function avsNodeRunner(uint256 _id) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) AvsNodeRunner(_id *big.Int) (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsNodeRunner(&_EtherfiAVSOperatorsManager.CallOpts, _id)
}

// AvsNodeRunner is a free data retrieval call binding the contract method 0x05dba646.
//
// Solidity: function avsNodeRunner(uint256 _id) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) AvsNodeRunner(_id *big.Int) (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsNodeRunner(&_EtherfiAVSOperatorsManager.CallOpts, _id)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0xd7eb63c7.
//
// Solidity: function avsOperatorStatus(uint256 _id, address _avsServiceManager) view returns(uint8)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) AvsOperatorStatus(opts *bind.CallOpts, _id *big.Int, _avsServiceManager common.Address) (uint8, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "avsOperatorStatus", _id, _avsServiceManager)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0xd7eb63c7.
//
// Solidity: function avsOperatorStatus(uint256 _id, address _avsServiceManager) view returns(uint8)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) AvsOperatorStatus(_id *big.Int, _avsServiceManager common.Address) (uint8, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsOperatorStatus(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsServiceManager)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0xd7eb63c7.
//
// Solidity: function avsOperatorStatus(uint256 _id, address _avsServiceManager) view returns(uint8)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) AvsOperatorStatus(_id *big.Int, _avsServiceManager common.Address) (uint8, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsOperatorStatus(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsServiceManager)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) AvsOperators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "avsOperators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsOperators(&_EtherfiAVSOperatorsManager.CallOpts, arg0)
}

// AvsOperators is a free data retrieval call binding the contract method 0x3d9da502.
//
// Solidity: function avsOperators(uint256 ) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) AvsOperators(arg0 *big.Int) (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.AvsOperators(&_EtherfiAVSOperatorsManager.CallOpts, arg0)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xba393bb8.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(uint256 _id, address _avsServiceManager, bytes32 _salt, uint256 _expiry) view returns(bytes32)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, _id *big.Int, _avsServiceManager common.Address, _salt [32]byte, _expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", _id, _avsServiceManager, _salt, _expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xba393bb8.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(uint256 _id, address _avsServiceManager, bytes32 _salt, uint256 _expiry) view returns(bytes32)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) CalculateOperatorAVSRegistrationDigestHash(_id *big.Int, _avsServiceManager common.Address, _salt [32]byte, _expiry *big.Int) ([32]byte, error) {
	return _EtherfiAVSOperatorsManager.Contract.CalculateOperatorAVSRegistrationDigestHash(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsServiceManager, _salt, _expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xba393bb8.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(uint256 _id, address _avsServiceManager, bytes32 _salt, uint256 _expiry) view returns(bytes32)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) CalculateOperatorAVSRegistrationDigestHash(_id *big.Int, _avsServiceManager common.Address, _salt [32]byte, _expiry *big.Int) ([32]byte, error) {
	return _EtherfiAVSOperatorsManager.Contract.CalculateOperatorAVSRegistrationDigestHash(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsServiceManager, _salt, _expiry)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) DelegationManager() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.DelegationManager(&_EtherfiAVSOperatorsManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) DelegationManager() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.DelegationManager(&_EtherfiAVSOperatorsManager.CallOpts)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x5408cff6.
//
// Solidity: function ecdsaSigner(uint256 _id) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) EcdsaSigner(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "ecdsaSigner", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcdsaSigner is a free data retrieval call binding the contract method 0x5408cff6.
//
// Solidity: function ecdsaSigner(uint256 _id) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) EcdsaSigner(_id *big.Int) (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.EcdsaSigner(&_EtherfiAVSOperatorsManager.CallOpts, _id)
}

// EcdsaSigner is a free data retrieval call binding the contract method 0x5408cff6.
//
// Solidity: function ecdsaSigner(uint256 _id) view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) EcdsaSigner(_id *big.Int) (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.EcdsaSigner(&_EtherfiAVSOperatorsManager.CallOpts, _id)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x10eed82a.
//
// Solidity: function getAvsInfo(uint256 _id, address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) GetAvsInfo(opts *bind.CallOpts, _id *big.Int, _avsRegistryCoordinator common.Address) (EtherFiAvsOperatorAvsInfo, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "getAvsInfo", _id, _avsRegistryCoordinator)

	if err != nil {
		return *new(EtherFiAvsOperatorAvsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EtherFiAvsOperatorAvsInfo)).(*EtherFiAvsOperatorAvsInfo)

	return out0, err

}

// GetAvsInfo is a free data retrieval call binding the contract method 0x10eed82a.
//
// Solidity: function getAvsInfo(uint256 _id, address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) GetAvsInfo(_id *big.Int, _avsRegistryCoordinator common.Address) (EtherFiAvsOperatorAvsInfo, error) {
	return _EtherfiAVSOperatorsManager.Contract.GetAvsInfo(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator)
}

// GetAvsInfo is a free data retrieval call binding the contract method 0x10eed82a.
//
// Solidity: function getAvsInfo(uint256 _id, address _avsRegistryCoordinator) view returns((bool,bytes,string,((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])),bool))
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) GetAvsInfo(_id *big.Int, _avsRegistryCoordinator common.Address) (EtherFiAvsOperatorAvsInfo, error) {
	return _EtherfiAVSOperatorsManager.Contract.GetAvsInfo(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator)
}

// IsAvsRegistered is a free data retrieval call binding the contract method 0xd8f0efe3.
//
// Solidity: function isAvsRegistered(uint256 _id, address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) IsAvsRegistered(opts *bind.CallOpts, _id *big.Int, _avsRegistryCoordinator common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "isAvsRegistered", _id, _avsRegistryCoordinator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvsRegistered is a free data retrieval call binding the contract method 0xd8f0efe3.
//
// Solidity: function isAvsRegistered(uint256 _id, address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) IsAvsRegistered(_id *big.Int, _avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.IsAvsRegistered(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator)
}

// IsAvsRegistered is a free data retrieval call binding the contract method 0xd8f0efe3.
//
// Solidity: function isAvsRegistered(uint256 _id, address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) IsAvsRegistered(_id *big.Int, _avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.IsAvsRegistered(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator)
}

// IsAvsWhitelisted is a free data retrieval call binding the contract method 0x78b4c573.
//
// Solidity: function isAvsWhitelisted(uint256 _id, address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) IsAvsWhitelisted(opts *bind.CallOpts, _id *big.Int, _avsRegistryCoordinator common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "isAvsWhitelisted", _id, _avsRegistryCoordinator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAvsWhitelisted is a free data retrieval call binding the contract method 0x78b4c573.
//
// Solidity: function isAvsWhitelisted(uint256 _id, address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) IsAvsWhitelisted(_id *big.Int, _avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.IsAvsWhitelisted(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator)
}

// IsAvsWhitelisted is a free data retrieval call binding the contract method 0x78b4c573.
//
// Solidity: function isAvsWhitelisted(uint256 _id, address _avsRegistryCoordinator) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) IsAvsWhitelisted(_id *big.Int, _avsRegistryCoordinator common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.IsAvsWhitelisted(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator)
}

// IsRegisteredBlsKey is a free data retrieval call binding the contract method 0x5d54a27a.
//
// Solidity: function isRegisteredBlsKey(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) IsRegisteredBlsKey(opts *bind.CallOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "isRegisteredBlsKey", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredBlsKey is a free data retrieval call binding the contract method 0x5d54a27a.
//
// Solidity: function isRegisteredBlsKey(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) IsRegisteredBlsKey(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.IsRegisteredBlsKey(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// IsRegisteredBlsKey is a free data retrieval call binding the contract method 0x5d54a27a.
//
// Solidity: function isRegisteredBlsKey(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) IsRegisteredBlsKey(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.IsRegisteredBlsKey(&_EtherfiAVSOperatorsManager.CallOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) NextAvsOperatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "nextAvsOperatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) NextAvsOperatorId() (*big.Int, error) {
	return _EtherfiAVSOperatorsManager.Contract.NextAvsOperatorId(&_EtherfiAVSOperatorsManager.CallOpts)
}

// NextAvsOperatorId is a free data retrieval call binding the contract method 0xae6f25c6.
//
// Solidity: function nextAvsOperatorId() view returns(uint256)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) NextAvsOperatorId() (*big.Int, error) {
	return _EtherfiAVSOperatorsManager.Contract.NextAvsOperatorId(&_EtherfiAVSOperatorsManager.CallOpts)
}

// OperatorDetails is a free data retrieval call binding the contract method 0x38a14d6d.
//
// Solidity: function operatorDetails(uint256 _id) view returns((address,address,uint32))
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) OperatorDetails(opts *bind.CallOpts, _id *big.Int) (IDelegationManagerOperatorDetails, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "operatorDetails", _id)

	if err != nil {
		return *new(IDelegationManagerOperatorDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationManagerOperatorDetails)).(*IDelegationManagerOperatorDetails)

	return out0, err

}

// OperatorDetails is a free data retrieval call binding the contract method 0x38a14d6d.
//
// Solidity: function operatorDetails(uint256 _id) view returns((address,address,uint32))
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) OperatorDetails(_id *big.Int) (IDelegationManagerOperatorDetails, error) {
	return _EtherfiAVSOperatorsManager.Contract.OperatorDetails(&_EtherfiAVSOperatorsManager.CallOpts, _id)
}

// OperatorDetails is a free data retrieval call binding the contract method 0x38a14d6d.
//
// Solidity: function operatorDetails(uint256 _id) view returns((address,address,uint32))
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) OperatorDetails(_id *big.Int) (IDelegationManagerOperatorDetails, error) {
	return _EtherfiAVSOperatorsManager.Contract.OperatorDetails(&_EtherfiAVSOperatorsManager.CallOpts, _id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) Owner() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.Owner(&_EtherfiAVSOperatorsManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) Owner() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.Owner(&_EtherfiAVSOperatorsManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) Paused() (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.Paused(&_EtherfiAVSOperatorsManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) Paused() (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.Paused(&_EtherfiAVSOperatorsManager.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) Pausers(arg0 common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.Pausers(&_EtherfiAVSOperatorsManager.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _EtherfiAVSOperatorsManager.Contract.Pausers(&_EtherfiAVSOperatorsManager.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherfiAVSOperatorsManager.Contract.ProxiableUUID(&_EtherfiAVSOperatorsManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _EtherfiAVSOperatorsManager.Contract.ProxiableUUID(&_EtherfiAVSOperatorsManager.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCaller) UpgradableBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAVSOperatorsManager.contract.Call(opts, &out, "upgradableBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpgradableBeacon() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradableBeacon(&_EtherfiAVSOperatorsManager.CallOpts)
}

// UpgradableBeacon is a free data retrieval call binding the contract method 0xe266f26a.
//
// Solidity: function upgradableBeacon() view returns(address)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerCallerSession) UpgradableBeacon() (common.Address, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradableBeacon(&_EtherfiAVSOperatorsManager.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) DeregisterOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "deregisterOperator", _id, _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) DeregisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.DeregisterOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xeee100fd.
//
// Solidity: function deregisterOperator(uint256 _id, address _avsRegistryCoordinator, bytes quorumNumbers) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) DeregisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.DeregisterOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delegationManager, address _avsDirectory, address _etherFiAvsOperatorImpl) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) Initialize(opts *bind.TransactOpts, _delegationManager common.Address, _avsDirectory common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "initialize", _delegationManager, _avsDirectory, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delegationManager, address _avsDirectory, address _etherFiAvsOperatorImpl) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) Initialize(_delegationManager common.Address, _avsDirectory common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.Initialize(&_EtherfiAVSOperatorsManager.TransactOpts, _delegationManager, _avsDirectory, _etherFiAvsOperatorImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _delegationManager, address _avsDirectory, address _etherFiAvsOperatorImpl) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) Initialize(_delegationManager common.Address, _avsDirectory common.Address, _etherFiAvsOperatorImpl common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.Initialize(&_EtherfiAVSOperatorsManager.TransactOpts, _delegationManager, _avsDirectory, _etherFiAvsOperatorImpl)
}

// InitializeAvsDirectory is a paid mutator transaction binding the contract method 0xafe833f7.
//
// Solidity: function initializeAvsDirectory(address _avsDirectory) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) InitializeAvsDirectory(opts *bind.TransactOpts, _avsDirectory common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "initializeAvsDirectory", _avsDirectory)
}

// InitializeAvsDirectory is a paid mutator transaction binding the contract method 0xafe833f7.
//
// Solidity: function initializeAvsDirectory(address _avsDirectory) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) InitializeAvsDirectory(_avsDirectory common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.InitializeAvsDirectory(&_EtherfiAVSOperatorsManager.TransactOpts, _avsDirectory)
}

// InitializeAvsDirectory is a paid mutator transaction binding the contract method 0xafe833f7.
//
// Solidity: function initializeAvsDirectory(address _avsDirectory) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) InitializeAvsDirectory(_avsDirectory common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.InitializeAvsDirectory(&_EtherfiAVSOperatorsManager.TransactOpts, _avsDirectory)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) InstantiateEtherFiAvsOperator(opts *bind.TransactOpts, _nums *big.Int) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "instantiateEtherFiAvsOperator", _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.InstantiateEtherFiAvsOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _nums)
}

// InstantiateEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x5a5cebf5.
//
// Solidity: function instantiateEtherFiAvsOperator(uint256 _nums) returns(uint256[] _ids)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) InstantiateEtherFiAvsOperator(_nums *big.Int) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.InstantiateEtherFiAvsOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _nums)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, _id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "modifyOperatorDetails", _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.ModifyOperatorDetails(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _newOperatorDetails)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x8acb29da.
//
// Solidity: function modifyOperatorDetails(uint256 _id, (address,address,uint32) _newOperatorDetails) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) ModifyOperatorDetails(_id *big.Int, _newOperatorDetails IDelegationManagerOperatorDetails) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.ModifyOperatorDetails(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _newOperatorDetails)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, _id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "registerAsOperator", _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterAsOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _detail, _metaDataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x92d59d3f.
//
// Solidity: function registerAsOperator(uint256 _id, (address,address,uint32) _detail, string _metaDataURI) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) RegisterAsOperator(_id *big.Int, _detail IDelegationManagerOperatorDetails, _metaDataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterAsOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _detail, _metaDataURI)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) RegisterBlsKeyAsDelegatedNodeOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "registerBlsKeyAsDelegatedNodeOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) RegisterBlsKeyAsDelegatedNodeOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterBlsKeyAsDelegatedNodeOperator is a paid mutator transaction binding the contract method 0xc74e7ae5.
//
// Solidity: function registerBlsKeyAsDelegatedNodeOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) RegisterBlsKeyAsDelegatedNodeOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterBlsKeyAsDelegatedNodeOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) RegisterOperator(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "registerOperator", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) RegisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x346c6fd0.
//
// Solidity: function registerOperator(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) RegisterOperator(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) RegisterOperatorWithChurn(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "registerOperatorWithChurn", _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) RegisterOperatorWithChurn(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterOperatorWithChurn(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RegisterOperatorWithChurn is a paid mutator transaction binding the contract method 0xe253829d.
//
// Solidity: function registerOperatorWithChurn(uint256 _id, address _avsRegistryCoordinator, bytes _quorumNumbers, string _socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) _params, (uint8,address)[] _operatorKickParams, (bytes,bytes32,uint256) _churnApproverSignature, (bytes,bytes32,uint256) _operatorSignature) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) RegisterOperatorWithChurn(_id *big.Int, _avsRegistryCoordinator common.Address, _quorumNumbers []byte, _socket string, _params IBLSApkRegistryPubkeyRegistrationParams, _operatorKickParams []IRegistryCoordinatorOperatorKickParam, _churnApproverSignature ISignatureUtilsSignatureWithSaltAndExpiry, _operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RegisterOperatorWithChurn(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _quorumNumbers, _socket, _params, _operatorKickParams, _churnApproverSignature, _operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RenounceOwnership(&_EtherfiAVSOperatorsManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RenounceOwnership(&_EtherfiAVSOperatorsManager.TransactOpts)
}

// RunnerForwardCall is a paid mutator transaction binding the contract method 0x71790c1f.
//
// Solidity: function runnerForwardCall(uint256 _id, address _target, bytes4 _selector, bytes _remainingCalldata) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) RunnerForwardCall(opts *bind.TransactOpts, _id *big.Int, _target common.Address, _selector [4]byte, _remainingCalldata []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "runnerForwardCall", _id, _target, _selector, _remainingCalldata)
}

// RunnerForwardCall is a paid mutator transaction binding the contract method 0x71790c1f.
//
// Solidity: function runnerForwardCall(uint256 _id, address _target, bytes4 _selector, bytes _remainingCalldata) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) RunnerForwardCall(_id *big.Int, _target common.Address, _selector [4]byte, _remainingCalldata []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RunnerForwardCall(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _target, _selector, _remainingCalldata)
}

// RunnerForwardCall is a paid mutator transaction binding the contract method 0x71790c1f.
//
// Solidity: function runnerForwardCall(uint256 _id, address _target, bytes4 _selector, bytes _remainingCalldata) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) RunnerForwardCall(_id *big.Int, _target common.Address, _selector [4]byte, _remainingCalldata []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.RunnerForwardCall(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _target, _selector, _remainingCalldata)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.TransferOwnership(&_EtherfiAVSOperatorsManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.TransferOwnership(&_EtherfiAVSOperatorsManager.TransactOpts, newOwner)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpdateAdmin(opts *bind.TransactOpts, _address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "updateAdmin", _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateAdmin(&_EtherfiAVSOperatorsManager.TransactOpts, _address, _isAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0x670a6fd9.
//
// Solidity: function updateAdmin(address _address, bool _isAdmin) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpdateAdmin(_address common.Address, _isAdmin bool) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateAdmin(&_EtherfiAVSOperatorsManager.TransactOpts, _address, _isAdmin)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpdateAvsNodeRunner(opts *bind.TransactOpts, _id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "updateAvsNodeRunner", _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateAvsNodeRunner(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsNodeRunner)
}

// UpdateAvsNodeRunner is a paid mutator transaction binding the contract method 0x626cd8fc.
//
// Solidity: function updateAvsNodeRunner(uint256 _id, address _avsNodeRunner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpdateAvsNodeRunner(_id *big.Int, _avsNodeRunner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateAvsNodeRunner(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsNodeRunner)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpdateAvsWhitelist(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "updateAvsWhitelist", _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpdateAvsWhitelist(_id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateAvsWhitelist(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateAvsWhitelist is a paid mutator transaction binding the contract method 0xccc0693b.
//
// Solidity: function updateAvsWhitelist(uint256 _id, address _avsRegistryCoordinator, bool _isWhitelisted) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpdateAvsWhitelist(_id *big.Int, _avsRegistryCoordinator common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateAvsWhitelist(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _isWhitelisted)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpdateEcdsaSigner(opts *bind.TransactOpts, _id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "updateEcdsaSigner", _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateEcdsaSigner(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _ecdsaSigner)
}

// UpdateEcdsaSigner is a paid mutator transaction binding the contract method 0xc94ed928.
//
// Solidity: function updateEcdsaSigner(uint256 _id, address _ecdsaSigner) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpdateEcdsaSigner(_id *big.Int, _ecdsaSigner common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateEcdsaSigner(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _ecdsaSigner)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, _id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "updateOperatorMetadataURI", _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateOperatorMetadataURI(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0xaccfd6f7.
//
// Solidity: function updateOperatorMetadataURI(uint256 _id, string _metadataURI) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpdateOperatorMetadataURI(_id *big.Int, _metadataURI string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateOperatorMetadataURI(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _metadataURI)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpdateSocket(opts *bind.TransactOpts, _id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "updateSocket", _id, _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpdateSocket(_id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateSocket(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0fae2a7d.
//
// Solidity: function updateSocket(uint256 _id, address _avsRegistryCoordinator, string _socket) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpdateSocket(_id *big.Int, _avsRegistryCoordinator common.Address, _socket string) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpdateSocket(&_EtherfiAVSOperatorsManager.TransactOpts, _id, _avsRegistryCoordinator, _socket)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpgradeEtherFiAvsOperator(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "upgradeEtherFiAvsOperator", _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradeEtherFiAvsOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _newImplementation)
}

// UpgradeEtherFiAvsOperator is a paid mutator transaction binding the contract method 0x8eefe329.
//
// Solidity: function upgradeEtherFiAvsOperator(address _newImplementation) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpgradeEtherFiAvsOperator(_newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradeEtherFiAvsOperator(&_EtherfiAVSOperatorsManager.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradeTo(&_EtherfiAVSOperatorsManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradeTo(&_EtherfiAVSOperatorsManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradeToAndCall(&_EtherfiAVSOperatorsManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _EtherfiAVSOperatorsManager.Contract.UpgradeToAndCall(&_EtherfiAVSOperatorsManager.TransactOpts, newImplementation, data)
}

// EtherfiAVSOperatorsManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerAdminChangedIterator struct {
	Event *EtherfiAVSOperatorsManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerAdminChanged)
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
		it.Event = new(EtherfiAVSOperatorsManagerAdminChanged)
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
func (it *EtherfiAVSOperatorsManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerAdminChanged represents a AdminChanged event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*EtherfiAVSOperatorsManagerAdminChangedIterator, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerAdminChangedIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerAdminChanged)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseAdminChanged(log types.Log) (*EtherfiAVSOperatorsManagerAdminChanged, error) {
	event := new(EtherfiAVSOperatorsManagerAdminChanged)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerBeaconUpgradedIterator struct {
	Event *EtherfiAVSOperatorsManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerBeaconUpgraded)
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
		it.Event = new(EtherfiAVSOperatorsManagerBeaconUpgraded)
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
func (it *EtherfiAVSOperatorsManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerBeaconUpgraded represents a BeaconUpgraded event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*EtherfiAVSOperatorsManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerBeaconUpgradedIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerBeaconUpgraded)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseBeaconUpgraded(log types.Log) (*EtherfiAVSOperatorsManagerBeaconUpgraded, error) {
	event := new(EtherfiAVSOperatorsManagerBeaconUpgraded)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperatorIterator is returned from FilterCreatedEtherFiAvsOperator and is used to iterate over the raw logs and unpacked data for CreatedEtherFiAvsOperator events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperatorIterator struct {
	Event *EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator)
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
		it.Event = new(EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator)
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
func (it *EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator represents a CreatedEtherFiAvsOperator event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator struct {
	Id                 *big.Int
	EtherFiAvsOperator common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCreatedEtherFiAvsOperator is a free log retrieval operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterCreatedEtherFiAvsOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperatorIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "CreatedEtherFiAvsOperator", logs: logs, sub: sub}, nil
}

// WatchCreatedEtherFiAvsOperator is a free log subscription operation binding the contract event 0x4d7b7ceda679acb24213ae2669a7fda45d700eb72145e9c9f7dcc5206529472e.
//
// Solidity: event CreatedEtherFiAvsOperator(uint256 indexed id, address etherFiAvsOperator)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchCreatedEtherFiAvsOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "CreatedEtherFiAvsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseCreatedEtherFiAvsOperator(log types.Log) (*EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator, error) {
	event := new(EtherfiAVSOperatorsManagerCreatedEtherFiAvsOperator)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "CreatedEtherFiAvsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerDeregisteredOperatorIterator is returned from FilterDeregisteredOperator and is used to iterate over the raw logs and unpacked data for DeregisteredOperator events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerDeregisteredOperatorIterator struct {
	Event *EtherfiAVSOperatorsManagerDeregisteredOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerDeregisteredOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerDeregisteredOperator)
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
		it.Event = new(EtherfiAVSOperatorsManagerDeregisteredOperator)
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
func (it *EtherfiAVSOperatorsManagerDeregisteredOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerDeregisteredOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerDeregisteredOperator represents a DeregisteredOperator event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerDeregisteredOperator struct {
	Id                *big.Int
	AvsServiceManager common.Address
	QuorumNumbers     []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDeregisteredOperator is a free log retrieval operation binding the contract event 0x1376d0cba684c4d21991c950915a2bf8a9181c389d343993ff185ba5944b5a2d.
//
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterDeregisteredOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerDeregisteredOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "DeregisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerDeregisteredOperatorIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "DeregisteredOperator", logs: logs, sub: sub}, nil
}

// WatchDeregisteredOperator is a free log subscription operation binding the contract event 0x1376d0cba684c4d21991c950915a2bf8a9181c389d343993ff185ba5944b5a2d.
//
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchDeregisteredOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerDeregisteredOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "DeregisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerDeregisteredOperator)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "DeregisteredOperator", log); err != nil {
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
// Solidity: event DeregisteredOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseDeregisteredOperator(log types.Log) (*EtherfiAVSOperatorsManagerDeregisteredOperator, error) {
	event := new(EtherfiAVSOperatorsManagerDeregisteredOperator)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "DeregisteredOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerForwardedRunnerCallIterator is returned from FilterForwardedRunnerCall and is used to iterate over the raw logs and unpacked data for ForwardedRunnerCall events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerForwardedRunnerCallIterator struct {
	Event *EtherfiAVSOperatorsManagerForwardedRunnerCall // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerForwardedRunnerCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerForwardedRunnerCall)
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
		it.Event = new(EtherfiAVSOperatorsManagerForwardedRunnerCall)
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
func (it *EtherfiAVSOperatorsManagerForwardedRunnerCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerForwardedRunnerCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerForwardedRunnerCall represents a ForwardedRunnerCall event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerForwardedRunnerCall struct {
	Id       *big.Int
	Target   common.Address
	Selector [4]byte
	Data     []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterForwardedRunnerCall is a free log retrieval operation binding the contract event 0x46412d627440b4b2939ce573b54f71d886e9916beabfde64b2572e459a6374a2.
//
// Solidity: event ForwardedRunnerCall(uint256 indexed id, address target, bytes4 selector, bytes data)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterForwardedRunnerCall(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerForwardedRunnerCallIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "ForwardedRunnerCall", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerForwardedRunnerCallIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "ForwardedRunnerCall", logs: logs, sub: sub}, nil
}

// WatchForwardedRunnerCall is a free log subscription operation binding the contract event 0x46412d627440b4b2939ce573b54f71d886e9916beabfde64b2572e459a6374a2.
//
// Solidity: event ForwardedRunnerCall(uint256 indexed id, address target, bytes4 selector, bytes data)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchForwardedRunnerCall(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerForwardedRunnerCall, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "ForwardedRunnerCall", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerForwardedRunnerCall)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "ForwardedRunnerCall", log); err != nil {
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

// ParseForwardedRunnerCall is a log parse operation binding the contract event 0x46412d627440b4b2939ce573b54f71d886e9916beabfde64b2572e459a6374a2.
//
// Solidity: event ForwardedRunnerCall(uint256 indexed id, address target, bytes4 selector, bytes data)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseForwardedRunnerCall(log types.Log) (*EtherfiAVSOperatorsManagerForwardedRunnerCall, error) {
	event := new(EtherfiAVSOperatorsManagerForwardedRunnerCall)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "ForwardedRunnerCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerInitializedIterator struct {
	Event *EtherfiAVSOperatorsManagerInitialized // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerInitialized)
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
		it.Event = new(EtherfiAVSOperatorsManagerInitialized)
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
func (it *EtherfiAVSOperatorsManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerInitialized represents a Initialized event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*EtherfiAVSOperatorsManagerInitializedIterator, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerInitializedIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerInitialized)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseInitialized(log types.Log) (*EtherfiAVSOperatorsManagerInitialized, error) {
	event := new(EtherfiAVSOperatorsManagerInitialized)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerModifiedOperatorDetailsIterator is returned from FilterModifiedOperatorDetails and is used to iterate over the raw logs and unpacked data for ModifiedOperatorDetails events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerModifiedOperatorDetailsIterator struct {
	Event *EtherfiAVSOperatorsManagerModifiedOperatorDetails // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerModifiedOperatorDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerModifiedOperatorDetails)
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
		it.Event = new(EtherfiAVSOperatorsManagerModifiedOperatorDetails)
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
func (it *EtherfiAVSOperatorsManagerModifiedOperatorDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerModifiedOperatorDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerModifiedOperatorDetails represents a ModifiedOperatorDetails event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerModifiedOperatorDetails struct {
	Id                 *big.Int
	NewOperatorDetails IDelegationManagerOperatorDetails
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterModifiedOperatorDetails is a free log retrieval operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterModifiedOperatorDetails(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerModifiedOperatorDetailsIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerModifiedOperatorDetailsIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "ModifiedOperatorDetails", logs: logs, sub: sub}, nil
}

// WatchModifiedOperatorDetails is a free log subscription operation binding the contract event 0xbc072201e8c4896adac46a042af73400cf0b7d3b4ff58cbf14877736a1d551cb.
//
// Solidity: event ModifiedOperatorDetails(uint256 indexed id, (address,address,uint32) newOperatorDetails)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchModifiedOperatorDetails(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerModifiedOperatorDetails, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "ModifiedOperatorDetails", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerModifiedOperatorDetails)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseModifiedOperatorDetails(log types.Log) (*EtherfiAVSOperatorsManagerModifiedOperatorDetails, error) {
	event := new(EtherfiAVSOperatorsManagerModifiedOperatorDetails)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "ModifiedOperatorDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerOwnershipTransferredIterator struct {
	Event *EtherfiAVSOperatorsManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerOwnershipTransferred)
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
		it.Event = new(EtherfiAVSOperatorsManagerOwnershipTransferred)
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
func (it *EtherfiAVSOperatorsManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerOwnershipTransferred represents a OwnershipTransferred event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EtherfiAVSOperatorsManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerOwnershipTransferredIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerOwnershipTransferred)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseOwnershipTransferred(log types.Log) (*EtherfiAVSOperatorsManagerOwnershipTransferred, error) {
	event := new(EtherfiAVSOperatorsManagerOwnershipTransferred)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerPausedIterator struct {
	Event *EtherfiAVSOperatorsManagerPaused // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerPaused)
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
		it.Event = new(EtherfiAVSOperatorsManagerPaused)
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
func (it *EtherfiAVSOperatorsManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerPaused represents a Paused event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*EtherfiAVSOperatorsManagerPausedIterator, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerPausedIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerPaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerPaused)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParsePaused(log types.Log) (*EtherfiAVSOperatorsManagerPaused, error) {
	event := new(EtherfiAVSOperatorsManagerPaused)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerRegisteredAsOperatorIterator is returned from FilterRegisteredAsOperator and is used to iterate over the raw logs and unpacked data for RegisteredAsOperator events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerRegisteredAsOperatorIterator struct {
	Event *EtherfiAVSOperatorsManagerRegisteredAsOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerRegisteredAsOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerRegisteredAsOperator)
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
		it.Event = new(EtherfiAVSOperatorsManagerRegisteredAsOperator)
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
func (it *EtherfiAVSOperatorsManagerRegisteredAsOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerRegisteredAsOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerRegisteredAsOperator represents a RegisteredAsOperator event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerRegisteredAsOperator struct {
	Id     *big.Int
	Detail IDelegationManagerOperatorDetails
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisteredAsOperator is a free log retrieval operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterRegisteredAsOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerRegisteredAsOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerRegisteredAsOperatorIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "RegisteredAsOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredAsOperator is a free log subscription operation binding the contract event 0x702cca91e590bd527801145cb928b821c3bd36196de46ffff13b380d3caa86b4.
//
// Solidity: event RegisteredAsOperator(uint256 indexed id, (address,address,uint32) detail)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchRegisteredAsOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerRegisteredAsOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "RegisteredAsOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerRegisteredAsOperator)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseRegisteredAsOperator(log types.Log) (*EtherfiAVSOperatorsManagerRegisteredAsOperator, error) {
	event := new(EtherfiAVSOperatorsManagerRegisteredAsOperator)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "RegisteredAsOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator is returned from FilterRegisteredBlsKeyAsDelegatedNodeOperator and is used to iterate over the raw logs and unpacked data for RegisteredBlsKeyAsDelegatedNodeOperator events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator struct {
	Event *EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
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
		it.Event = new(EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
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
func (it *EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator represents a RegisteredBlsKeyAsDelegatedNodeOperator event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator struct {
	Id                *big.Int
	AvsServiceManager common.Address
	QuorumNumbers     []byte
	Socket            string
	Params            IBLSApkRegistryPubkeyRegistrationParams
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRegisteredBlsKeyAsDelegatedNodeOperator is a free log retrieval operation binding the contract event 0x3bf95f161a4d8780e658d5f1476e87c49b56ff67c2ecec2baeadb4a474f64d51.
//
// Solidity: event RegisteredBlsKeyAsDelegatedNodeOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterRegisteredBlsKeyAsDelegatedNodeOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "RegisteredBlsKeyAsDelegatedNodeOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperatorIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "RegisteredBlsKeyAsDelegatedNodeOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredBlsKeyAsDelegatedNodeOperator is a free log subscription operation binding the contract event 0x3bf95f161a4d8780e658d5f1476e87c49b56ff67c2ecec2baeadb4a474f64d51.
//
// Solidity: event RegisteredBlsKeyAsDelegatedNodeOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchRegisteredBlsKeyAsDelegatedNodeOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "RegisteredBlsKeyAsDelegatedNodeOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "RegisteredBlsKeyAsDelegatedNodeOperator", log); err != nil {
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
// Solidity: event RegisteredBlsKeyAsDelegatedNodeOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseRegisteredBlsKeyAsDelegatedNodeOperator(log types.Log) (*EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator, error) {
	event := new(EtherfiAVSOperatorsManagerRegisteredBlsKeyAsDelegatedNodeOperator)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "RegisteredBlsKeyAsDelegatedNodeOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerRegisteredOperatorIterator is returned from FilterRegisteredOperator and is used to iterate over the raw logs and unpacked data for RegisteredOperator events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerRegisteredOperatorIterator struct {
	Event *EtherfiAVSOperatorsManagerRegisteredOperator // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerRegisteredOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerRegisteredOperator)
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
		it.Event = new(EtherfiAVSOperatorsManagerRegisteredOperator)
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
func (it *EtherfiAVSOperatorsManagerRegisteredOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerRegisteredOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerRegisteredOperator represents a RegisteredOperator event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerRegisteredOperator struct {
	Id                *big.Int
	AvsServiceManager common.Address
	QuorumNumbers     []byte
	Socket            string
	Params            IBLSApkRegistryPubkeyRegistrationParams
	OperatorSignature ISignatureUtilsSignatureWithSaltAndExpiry
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRegisteredOperator is a free log retrieval operation binding the contract event 0x7fcf4f067d4a0674927732b3bfbbccaa47ef090c869bd6189e8dc842b46cc62c.
//
// Solidity: event RegisteredOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterRegisteredOperator(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerRegisteredOperatorIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "RegisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerRegisteredOperatorIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "RegisteredOperator", logs: logs, sub: sub}, nil
}

// WatchRegisteredOperator is a free log subscription operation binding the contract event 0x7fcf4f067d4a0674927732b3bfbbccaa47ef090c869bd6189e8dc842b46cc62c.
//
// Solidity: event RegisteredOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchRegisteredOperator(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerRegisteredOperator, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "RegisteredOperator", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerRegisteredOperator)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "RegisteredOperator", log); err != nil {
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
// Solidity: event RegisteredOperator(uint256 indexed id, address avsServiceManager, bytes quorumNumbers, string socket, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (bytes,bytes32,uint256) operatorSignature)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseRegisteredOperator(log types.Log) (*EtherfiAVSOperatorsManagerRegisteredOperator, error) {
	event := new(EtherfiAVSOperatorsManagerRegisteredOperator)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "RegisteredOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUnpausedIterator struct {
	Event *EtherfiAVSOperatorsManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerUnpaused)
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
		it.Event = new(EtherfiAVSOperatorsManagerUnpaused)
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
func (it *EtherfiAVSOperatorsManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerUnpaused represents a Unpaused event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EtherfiAVSOperatorsManagerUnpausedIterator, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerUnpausedIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerUnpaused)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseUnpaused(log types.Log) (*EtherfiAVSOperatorsManagerUnpaused, error) {
	event := new(EtherfiAVSOperatorsManagerUnpaused)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerUpdatedAvsNodeRunnerIterator is returned from FilterUpdatedAvsNodeRunner and is used to iterate over the raw logs and unpacked data for UpdatedAvsNodeRunner events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedAvsNodeRunnerIterator struct {
	Event *EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerUpdatedAvsNodeRunnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner)
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
		it.Event = new(EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner)
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
func (it *EtherfiAVSOperatorsManagerUpdatedAvsNodeRunnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerUpdatedAvsNodeRunnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner represents a UpdatedAvsNodeRunner event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner struct {
	Id            *big.Int
	AvsNodeRunner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvsNodeRunner is a free log retrieval operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterUpdatedAvsNodeRunner(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerUpdatedAvsNodeRunnerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerUpdatedAvsNodeRunnerIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "UpdatedAvsNodeRunner", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvsNodeRunner is a free log subscription operation binding the contract event 0x50f9c5b78cf9d87091c218d43d96c1759b0ea6bab4816360425eb08fae896c0b.
//
// Solidity: event UpdatedAvsNodeRunner(uint256 indexed id, address avsNodeRunner)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchUpdatedAvsNodeRunner(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "UpdatedAvsNodeRunner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseUpdatedAvsNodeRunner(log types.Log) (*EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner, error) {
	event := new(EtherfiAVSOperatorsManagerUpdatedAvsNodeRunner)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedAvsNodeRunner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerUpdatedAvsWhitelistIterator is returned from FilterUpdatedAvsWhitelist and is used to iterate over the raw logs and unpacked data for UpdatedAvsWhitelist events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedAvsWhitelistIterator struct {
	Event *EtherfiAVSOperatorsManagerUpdatedAvsWhitelist // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerUpdatedAvsWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerUpdatedAvsWhitelist)
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
		it.Event = new(EtherfiAVSOperatorsManagerUpdatedAvsWhitelist)
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
func (it *EtherfiAVSOperatorsManagerUpdatedAvsWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerUpdatedAvsWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerUpdatedAvsWhitelist represents a UpdatedAvsWhitelist event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedAvsWhitelist struct {
	Id                *big.Int
	AvsServiceManager common.Address
	IsWhitelisted     bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdatedAvsWhitelist is a free log retrieval operation binding the contract event 0x0cf8cb3148c96f3b418d80fc77f713f454ab17bfa4692f340ddd3057e65a78ac.
//
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsServiceManager, bool isWhitelisted)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterUpdatedAvsWhitelist(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerUpdatedAvsWhitelistIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "UpdatedAvsWhitelist", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerUpdatedAvsWhitelistIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "UpdatedAvsWhitelist", logs: logs, sub: sub}, nil
}

// WatchUpdatedAvsWhitelist is a free log subscription operation binding the contract event 0x0cf8cb3148c96f3b418d80fc77f713f454ab17bfa4692f340ddd3057e65a78ac.
//
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsServiceManager, bool isWhitelisted)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchUpdatedAvsWhitelist(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerUpdatedAvsWhitelist, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "UpdatedAvsWhitelist", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerUpdatedAvsWhitelist)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedAvsWhitelist", log); err != nil {
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
// Solidity: event UpdatedAvsWhitelist(uint256 indexed id, address avsServiceManager, bool isWhitelisted)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseUpdatedAvsWhitelist(log types.Log) (*EtherfiAVSOperatorsManagerUpdatedAvsWhitelist, error) {
	event := new(EtherfiAVSOperatorsManagerUpdatedAvsWhitelist)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedAvsWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerUpdatedEcdsaSignerIterator is returned from FilterUpdatedEcdsaSigner and is used to iterate over the raw logs and unpacked data for UpdatedEcdsaSigner events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedEcdsaSignerIterator struct {
	Event *EtherfiAVSOperatorsManagerUpdatedEcdsaSigner // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerUpdatedEcdsaSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerUpdatedEcdsaSigner)
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
		it.Event = new(EtherfiAVSOperatorsManagerUpdatedEcdsaSigner)
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
func (it *EtherfiAVSOperatorsManagerUpdatedEcdsaSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerUpdatedEcdsaSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerUpdatedEcdsaSigner represents a UpdatedEcdsaSigner event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedEcdsaSigner struct {
	Id          *big.Int
	EcdsaSigner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedEcdsaSigner is a free log retrieval operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterUpdatedEcdsaSigner(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerUpdatedEcdsaSignerIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerUpdatedEcdsaSignerIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "UpdatedEcdsaSigner", logs: logs, sub: sub}, nil
}

// WatchUpdatedEcdsaSigner is a free log subscription operation binding the contract event 0x79559f144b1d369bc6e3c73bf52efae8f5af26c59bb9f33a97e6e9221ad28c57.
//
// Solidity: event UpdatedEcdsaSigner(uint256 indexed id, address ecdsaSigner)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchUpdatedEcdsaSigner(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerUpdatedEcdsaSigner, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "UpdatedEcdsaSigner", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerUpdatedEcdsaSigner)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseUpdatedEcdsaSigner(log types.Log) (*EtherfiAVSOperatorsManagerUpdatedEcdsaSigner, error) {
	event := new(EtherfiAVSOperatorsManagerUpdatedEcdsaSigner)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedEcdsaSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURIIterator is returned from FilterUpdatedOperatorMetadataURI and is used to iterate over the raw logs and unpacked data for UpdatedOperatorMetadataURI events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURIIterator struct {
	Event *EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI)
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
		it.Event = new(EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI)
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
func (it *EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI represents a UpdatedOperatorMetadataURI event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI struct {
	Id          *big.Int
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedOperatorMetadataURI is a free log retrieval operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterUpdatedOperatorMetadataURI(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURIIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "UpdatedOperatorMetadataURI", logs: logs, sub: sub}, nil
}

// WatchUpdatedOperatorMetadataURI is a free log subscription operation binding the contract event 0x722e352197cc934495ca287effc2332acd796933ec3302bbfb782edd3f733f62.
//
// Solidity: event UpdatedOperatorMetadataURI(uint256 indexed id, string metadataURI)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchUpdatedOperatorMetadataURI(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "UpdatedOperatorMetadataURI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseUpdatedOperatorMetadataURI(log types.Log) (*EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI, error) {
	event := new(EtherfiAVSOperatorsManagerUpdatedOperatorMetadataURI)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedOperatorMetadataURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerUpdatedSocketIterator is returned from FilterUpdatedSocket and is used to iterate over the raw logs and unpacked data for UpdatedSocket events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedSocketIterator struct {
	Event *EtherfiAVSOperatorsManagerUpdatedSocket // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerUpdatedSocketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerUpdatedSocket)
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
		it.Event = new(EtherfiAVSOperatorsManagerUpdatedSocket)
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
func (it *EtherfiAVSOperatorsManagerUpdatedSocketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerUpdatedSocketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerUpdatedSocket represents a UpdatedSocket event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpdatedSocket struct {
	Id                *big.Int
	AvsServiceManager common.Address
	Socket            string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSocket is a free log retrieval operation binding the contract event 0x6e72bae16f3bc0863210747946d52680ce537fca1cfb695c91bdb9bf872966b9.
//
// Solidity: event UpdatedSocket(uint256 indexed id, address avsServiceManager, string socket)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterUpdatedSocket(opts *bind.FilterOpts, id []*big.Int) (*EtherfiAVSOperatorsManagerUpdatedSocketIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "UpdatedSocket", idRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerUpdatedSocketIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "UpdatedSocket", logs: logs, sub: sub}, nil
}

// WatchUpdatedSocket is a free log subscription operation binding the contract event 0x6e72bae16f3bc0863210747946d52680ce537fca1cfb695c91bdb9bf872966b9.
//
// Solidity: event UpdatedSocket(uint256 indexed id, address avsServiceManager, string socket)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchUpdatedSocket(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerUpdatedSocket, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "UpdatedSocket", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerUpdatedSocket)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedSocket", log); err != nil {
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
// Solidity: event UpdatedSocket(uint256 indexed id, address avsServiceManager, string socket)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseUpdatedSocket(log types.Log) (*EtherfiAVSOperatorsManagerUpdatedSocket, error) {
	event := new(EtherfiAVSOperatorsManagerUpdatedSocket)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "UpdatedSocket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAVSOperatorsManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpgradedIterator struct {
	Event *EtherfiAVSOperatorsManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *EtherfiAVSOperatorsManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAVSOperatorsManagerUpgraded)
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
		it.Event = new(EtherfiAVSOperatorsManagerUpgraded)
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
func (it *EtherfiAVSOperatorsManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAVSOperatorsManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAVSOperatorsManagerUpgraded represents a Upgraded event raised by the EtherfiAVSOperatorsManager contract.
type EtherfiAVSOperatorsManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EtherfiAVSOperatorsManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAVSOperatorsManagerUpgradedIterator{contract: _EtherfiAVSOperatorsManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EtherfiAVSOperatorsManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _EtherfiAVSOperatorsManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAVSOperatorsManagerUpgraded)
				if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_EtherfiAVSOperatorsManager *EtherfiAVSOperatorsManagerFilterer) ParseUpgraded(log types.Log) (*EtherfiAVSOperatorsManagerUpgraded, error) {
	event := new(EtherfiAVSOperatorsManagerUpgraded)
	if err := _EtherfiAVSOperatorsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
