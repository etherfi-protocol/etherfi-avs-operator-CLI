// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arpa

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

// INodeRegistryNode is an auto generated low-level Go binding around an user-defined struct.
type INodeRegistryNode struct {
	IdAddress         common.Address
	DkgPublicKey      []byte
	IsEigenlayerNode  bool
	State             bool
	PendingUntilBlock *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// NodeRegistryMetaData contains all meta data concerning the NodeRegistry contract.
var NodeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EIP1271SignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EIP1271SignatureNotFromSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EIP1271SignatureSaltAlreadySpent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EIP1271SignatureVerificationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingUntilBlock\",\"type\":\"uint256\"}],\"name\":\"NodeStillPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorUnderStaking\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotController\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetAccountAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"AssetAccountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dkgPublicKey\",\"type\":\"bytes\"}],\"name\":\"DkgPublicKeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupIndex\",\"type\":\"uint256\"}],\"name\":\"NodeActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"NodeQuit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dkgPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"groupIndex\",\"type\":\"uint256\"}],\"name\":\"NodeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arpaAmount\",\"type\":\"uint256\"}],\"name\":\"NodeRewarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeIdAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingRewardPenalty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingBlock\",\"type\":\"uint256\"}],\"name\":\"NodeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_NODE_REGISTRATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nodes\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"ethAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arpaAmount\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAccountAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"assetAccountSaltIsSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAccountAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateNativeNodeRegistrationDigestHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dkgPublicKey\",\"type\":\"bytes\"}],\"name\":\"changeDkgPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeIdAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pendingBlock\",\"type\":\"uint256\"}],\"name\":\"dismissNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getAssetAccountAddressByNodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getDKGPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNode\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"idAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dkgPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isEigenlayerNode\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pendingUntilBlock\",\"type\":\"uint256\"}],\"internalType\":\"structINodeRegistry.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetAccountAddress\",\"type\":\"address\"}],\"name\":\"getNodeAddressByAssetAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeRegistryConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"controllerContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"serviceManagerContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeNodeStakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eigenlayerNodeStakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingBlockAfterQuit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeWithdrawableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arpa\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"assetAccountSignature\",\"type\":\"tuple\"}],\"name\":\"nodeActivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeLogOff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeQuit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"dkgPublicKey\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isEigenlayerNode\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"assetAccountAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"assetAccountSignature\",\"type\":\"tuple\"}],\"name\":\"nodeRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"nodeWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assetAccountAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"nodeAddresses\",\"type\":\"address[]\"}],\"name\":\"setAssetAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controllerContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"serviceManagerContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nativeNodeStakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eigenlayerNodeStakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingBlockAfterQuit\",\"type\":\"uint256\"}],\"name\":\"setNodeRegistryConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeIdAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakingRewardPenalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingBlock\",\"type\":\"uint256\"}],\"name\":\"slashNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// NodeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use NodeRegistryMetaData.ABI instead.
var NodeRegistryABI = NodeRegistryMetaData.ABI

// NodeRegistry is an auto generated Go binding around an Ethereum contract.
type NodeRegistry struct {
	NodeRegistryCaller     // Read-only binding to the contract
	NodeRegistryTransactor // Write-only binding to the contract
	NodeRegistryFilterer   // Log filterer for contract events
}

// NodeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeRegistrySession struct {
	Contract     *NodeRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeRegistryCallerSession struct {
	Contract *NodeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NodeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeRegistryTransactorSession struct {
	Contract     *NodeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NodeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRegistryRaw struct {
	Contract *NodeRegistry // Generic contract binding to access the raw methods on
}

// NodeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeRegistryCallerRaw struct {
	Contract *NodeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// NodeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeRegistryTransactorRaw struct {
	Contract *NodeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeRegistry creates a new instance of NodeRegistry, bound to a specific deployed contract.
func NewNodeRegistry(address common.Address, backend bind.ContractBackend) (*NodeRegistry, error) {
	contract, err := bindNodeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeRegistry{NodeRegistryCaller: NodeRegistryCaller{contract: contract}, NodeRegistryTransactor: NodeRegistryTransactor{contract: contract}, NodeRegistryFilterer: NodeRegistryFilterer{contract: contract}}, nil
}

// NewNodeRegistryCaller creates a new read-only instance of NodeRegistry, bound to a specific deployed contract.
func NewNodeRegistryCaller(address common.Address, caller bind.ContractCaller) (*NodeRegistryCaller, error) {
	contract, err := bindNodeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryCaller{contract: contract}, nil
}

// NewNodeRegistryTransactor creates a new write-only instance of NodeRegistry, bound to a specific deployed contract.
func NewNodeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeRegistryTransactor, error) {
	contract, err := bindNodeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryTransactor{contract: contract}, nil
}

// NewNodeRegistryFilterer creates a new log filterer instance of NodeRegistry, bound to a specific deployed contract.
func NewNodeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeRegistryFilterer, error) {
	contract, err := bindNodeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryFilterer{contract: contract}, nil
}

// bindNodeRegistry binds a generic wrapper to an already deployed contract.
func bindNodeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRegistry *NodeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRegistry.Contract.NodeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRegistry *NodeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRegistry *NodeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeRegistry *NodeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeRegistry *NodeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeRegistry *NodeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeRegistry.Contract.contract.Transact(opts, method, params...)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NodeRegistry *NodeRegistrySession) DOMAINTYPEHASH() ([32]byte, error) {
	return _NodeRegistry.Contract.DOMAINTYPEHASH(&_NodeRegistry.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _NodeRegistry.Contract.DOMAINTYPEHASH(&_NodeRegistry.CallOpts)
}

// NATIVENODEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xb447f451.
//
// Solidity: function NATIVE_NODE_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCaller) NATIVENODEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "NATIVE_NODE_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NATIVENODEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xb447f451.
//
// Solidity: function NATIVE_NODE_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_NodeRegistry *NodeRegistrySession) NATIVENODEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _NodeRegistry.Contract.NATIVENODEREGISTRATIONTYPEHASH(&_NodeRegistry.CallOpts)
}

// NATIVENODEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xb447f451.
//
// Solidity: function NATIVE_NODE_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCallerSession) NATIVENODEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _NodeRegistry.Contract.NATIVENODEREGISTRATIONTYPEHASH(&_NodeRegistry.CallOpts)
}

// AssetAccountSaltIsSpent is a free data retrieval call binding the contract method 0xc71e3c78.
//
// Solidity: function assetAccountSaltIsSpent(address assetAccountAddress, bytes32 salt) view returns(bool)
func (_NodeRegistry *NodeRegistryCaller) AssetAccountSaltIsSpent(opts *bind.CallOpts, assetAccountAddress common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "assetAccountSaltIsSpent", assetAccountAddress, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetAccountSaltIsSpent is a free data retrieval call binding the contract method 0xc71e3c78.
//
// Solidity: function assetAccountSaltIsSpent(address assetAccountAddress, bytes32 salt) view returns(bool)
func (_NodeRegistry *NodeRegistrySession) AssetAccountSaltIsSpent(assetAccountAddress common.Address, salt [32]byte) (bool, error) {
	return _NodeRegistry.Contract.AssetAccountSaltIsSpent(&_NodeRegistry.CallOpts, assetAccountAddress, salt)
}

// AssetAccountSaltIsSpent is a free data retrieval call binding the contract method 0xc71e3c78.
//
// Solidity: function assetAccountSaltIsSpent(address assetAccountAddress, bytes32 salt) view returns(bool)
func (_NodeRegistry *NodeRegistryCallerSession) AssetAccountSaltIsSpent(assetAccountAddress common.Address, salt [32]byte) (bool, error) {
	return _NodeRegistry.Contract.AssetAccountSaltIsSpent(&_NodeRegistry.CallOpts, assetAccountAddress, salt)
}

// CalculateNativeNodeRegistrationDigestHash is a free data retrieval call binding the contract method 0xa2cd23e6.
//
// Solidity: function calculateNativeNodeRegistrationDigestHash(address assetAccountAddress, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_NodeRegistry *NodeRegistryCaller) CalculateNativeNodeRegistrationDigestHash(opts *bind.CallOpts, assetAccountAddress common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "calculateNativeNodeRegistrationDigestHash", assetAccountAddress, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateNativeNodeRegistrationDigestHash is a free data retrieval call binding the contract method 0xa2cd23e6.
//
// Solidity: function calculateNativeNodeRegistrationDigestHash(address assetAccountAddress, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_NodeRegistry *NodeRegistrySession) CalculateNativeNodeRegistrationDigestHash(assetAccountAddress common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _NodeRegistry.Contract.CalculateNativeNodeRegistrationDigestHash(&_NodeRegistry.CallOpts, assetAccountAddress, salt, expiry)
}

// CalculateNativeNodeRegistrationDigestHash is a free data retrieval call binding the contract method 0xa2cd23e6.
//
// Solidity: function calculateNativeNodeRegistrationDigestHash(address assetAccountAddress, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_NodeRegistry *NodeRegistryCallerSession) CalculateNativeNodeRegistrationDigestHash(assetAccountAddress common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _NodeRegistry.Contract.CalculateNativeNodeRegistrationDigestHash(&_NodeRegistry.CallOpts, assetAccountAddress, salt, expiry)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NodeRegistry *NodeRegistrySession) DomainSeparator() ([32]byte, error) {
	return _NodeRegistry.Contract.DomainSeparator(&_NodeRegistry.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCallerSession) DomainSeparator() ([32]byte, error) {
	return _NodeRegistry.Contract.DomainSeparator(&_NodeRegistry.CallOpts)
}

// GetAssetAccountAddressByNodeAddress is a free data retrieval call binding the contract method 0x6cc7fb5d.
//
// Solidity: function getAssetAccountAddressByNodeAddress(address nodeAddress) view returns(address)
func (_NodeRegistry *NodeRegistryCaller) GetAssetAccountAddressByNodeAddress(opts *bind.CallOpts, nodeAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "getAssetAccountAddressByNodeAddress", nodeAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAssetAccountAddressByNodeAddress is a free data retrieval call binding the contract method 0x6cc7fb5d.
//
// Solidity: function getAssetAccountAddressByNodeAddress(address nodeAddress) view returns(address)
func (_NodeRegistry *NodeRegistrySession) GetAssetAccountAddressByNodeAddress(nodeAddress common.Address) (common.Address, error) {
	return _NodeRegistry.Contract.GetAssetAccountAddressByNodeAddress(&_NodeRegistry.CallOpts, nodeAddress)
}

// GetAssetAccountAddressByNodeAddress is a free data retrieval call binding the contract method 0x6cc7fb5d.
//
// Solidity: function getAssetAccountAddressByNodeAddress(address nodeAddress) view returns(address)
func (_NodeRegistry *NodeRegistryCallerSession) GetAssetAccountAddressByNodeAddress(nodeAddress common.Address) (common.Address, error) {
	return _NodeRegistry.Contract.GetAssetAccountAddressByNodeAddress(&_NodeRegistry.CallOpts, nodeAddress)
}

// GetDKGPublicKey is a free data retrieval call binding the contract method 0x15dac9b2.
//
// Solidity: function getDKGPublicKey(address nodeAddress) view returns(bytes)
func (_NodeRegistry *NodeRegistryCaller) GetDKGPublicKey(opts *bind.CallOpts, nodeAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "getDKGPublicKey", nodeAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDKGPublicKey is a free data retrieval call binding the contract method 0x15dac9b2.
//
// Solidity: function getDKGPublicKey(address nodeAddress) view returns(bytes)
func (_NodeRegistry *NodeRegistrySession) GetDKGPublicKey(nodeAddress common.Address) ([]byte, error) {
	return _NodeRegistry.Contract.GetDKGPublicKey(&_NodeRegistry.CallOpts, nodeAddress)
}

// GetDKGPublicKey is a free data retrieval call binding the contract method 0x15dac9b2.
//
// Solidity: function getDKGPublicKey(address nodeAddress) view returns(bytes)
func (_NodeRegistry *NodeRegistryCallerSession) GetDKGPublicKey(nodeAddress common.Address) ([]byte, error) {
	return _NodeRegistry.Contract.GetDKGPublicKey(&_NodeRegistry.CallOpts, nodeAddress)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address nodeAddress) view returns((address,bytes,bool,bool,uint256))
func (_NodeRegistry *NodeRegistryCaller) GetNode(opts *bind.CallOpts, nodeAddress common.Address) (INodeRegistryNode, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "getNode", nodeAddress)

	if err != nil {
		return *new(INodeRegistryNode), err
	}

	out0 := *abi.ConvertType(out[0], new(INodeRegistryNode)).(*INodeRegistryNode)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address nodeAddress) view returns((address,bytes,bool,bool,uint256))
func (_NodeRegistry *NodeRegistrySession) GetNode(nodeAddress common.Address) (INodeRegistryNode, error) {
	return _NodeRegistry.Contract.GetNode(&_NodeRegistry.CallOpts, nodeAddress)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address nodeAddress) view returns((address,bytes,bool,bool,uint256))
func (_NodeRegistry *NodeRegistryCallerSession) GetNode(nodeAddress common.Address) (INodeRegistryNode, error) {
	return _NodeRegistry.Contract.GetNode(&_NodeRegistry.CallOpts, nodeAddress)
}

// GetNodeAddressByAssetAccountAddress is a free data retrieval call binding the contract method 0xd20cc152.
//
// Solidity: function getNodeAddressByAssetAccountAddress(address assetAccountAddress) view returns(address)
func (_NodeRegistry *NodeRegistryCaller) GetNodeAddressByAssetAccountAddress(opts *bind.CallOpts, assetAccountAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "getNodeAddressByAssetAccountAddress", assetAccountAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeAddressByAssetAccountAddress is a free data retrieval call binding the contract method 0xd20cc152.
//
// Solidity: function getNodeAddressByAssetAccountAddress(address assetAccountAddress) view returns(address)
func (_NodeRegistry *NodeRegistrySession) GetNodeAddressByAssetAccountAddress(assetAccountAddress common.Address) (common.Address, error) {
	return _NodeRegistry.Contract.GetNodeAddressByAssetAccountAddress(&_NodeRegistry.CallOpts, assetAccountAddress)
}

// GetNodeAddressByAssetAccountAddress is a free data retrieval call binding the contract method 0xd20cc152.
//
// Solidity: function getNodeAddressByAssetAccountAddress(address assetAccountAddress) view returns(address)
func (_NodeRegistry *NodeRegistryCallerSession) GetNodeAddressByAssetAccountAddress(assetAccountAddress common.Address) (common.Address, error) {
	return _NodeRegistry.Contract.GetNodeAddressByAssetAccountAddress(&_NodeRegistry.CallOpts, assetAccountAddress)
}

// GetNodeRegistryConfig is a free data retrieval call binding the contract method 0xe40e744b.
//
// Solidity: function getNodeRegistryConfig() view returns(address controllerContractAddress, address stakingContractAddress, address serviceManagerContractAddress, uint256 nativeNodeStakingAmount, uint256 eigenlayerNodeStakingAmount, uint256 pendingBlockAfterQuit)
func (_NodeRegistry *NodeRegistryCaller) GetNodeRegistryConfig(opts *bind.CallOpts) (struct {
	ControllerContractAddress     common.Address
	StakingContractAddress        common.Address
	ServiceManagerContractAddress common.Address
	NativeNodeStakingAmount       *big.Int
	EigenlayerNodeStakingAmount   *big.Int
	PendingBlockAfterQuit         *big.Int
}, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "getNodeRegistryConfig")

	outstruct := new(struct {
		ControllerContractAddress     common.Address
		StakingContractAddress        common.Address
		ServiceManagerContractAddress common.Address
		NativeNodeStakingAmount       *big.Int
		EigenlayerNodeStakingAmount   *big.Int
		PendingBlockAfterQuit         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ControllerContractAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakingContractAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ServiceManagerContractAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NativeNodeStakingAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EigenlayerNodeStakingAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PendingBlockAfterQuit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeRegistryConfig is a free data retrieval call binding the contract method 0xe40e744b.
//
// Solidity: function getNodeRegistryConfig() view returns(address controllerContractAddress, address stakingContractAddress, address serviceManagerContractAddress, uint256 nativeNodeStakingAmount, uint256 eigenlayerNodeStakingAmount, uint256 pendingBlockAfterQuit)
func (_NodeRegistry *NodeRegistrySession) GetNodeRegistryConfig() (struct {
	ControllerContractAddress     common.Address
	StakingContractAddress        common.Address
	ServiceManagerContractAddress common.Address
	NativeNodeStakingAmount       *big.Int
	EigenlayerNodeStakingAmount   *big.Int
	PendingBlockAfterQuit         *big.Int
}, error) {
	return _NodeRegistry.Contract.GetNodeRegistryConfig(&_NodeRegistry.CallOpts)
}

// GetNodeRegistryConfig is a free data retrieval call binding the contract method 0xe40e744b.
//
// Solidity: function getNodeRegistryConfig() view returns(address controllerContractAddress, address stakingContractAddress, address serviceManagerContractAddress, uint256 nativeNodeStakingAmount, uint256 eigenlayerNodeStakingAmount, uint256 pendingBlockAfterQuit)
func (_NodeRegistry *NodeRegistryCallerSession) GetNodeRegistryConfig() (struct {
	ControllerContractAddress     common.Address
	StakingContractAddress        common.Address
	ServiceManagerContractAddress common.Address
	NativeNodeStakingAmount       *big.Int
	EigenlayerNodeStakingAmount   *big.Int
	PendingBlockAfterQuit         *big.Int
}, error) {
	return _NodeRegistry.Contract.GetNodeRegistryConfig(&_NodeRegistry.CallOpts)
}

// GetNodeWithdrawableTokens is a free data retrieval call binding the contract method 0x227d0f46.
//
// Solidity: function getNodeWithdrawableTokens(address nodeAddress) view returns(uint256, uint256)
func (_NodeRegistry *NodeRegistryCaller) GetNodeWithdrawableTokens(opts *bind.CallOpts, nodeAddress common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "getNodeWithdrawableTokens", nodeAddress)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetNodeWithdrawableTokens is a free data retrieval call binding the contract method 0x227d0f46.
//
// Solidity: function getNodeWithdrawableTokens(address nodeAddress) view returns(uint256, uint256)
func (_NodeRegistry *NodeRegistrySession) GetNodeWithdrawableTokens(nodeAddress common.Address) (*big.Int, *big.Int, error) {
	return _NodeRegistry.Contract.GetNodeWithdrawableTokens(&_NodeRegistry.CallOpts, nodeAddress)
}

// GetNodeWithdrawableTokens is a free data retrieval call binding the contract method 0x227d0f46.
//
// Solidity: function getNodeWithdrawableTokens(address nodeAddress) view returns(uint256, uint256)
func (_NodeRegistry *NodeRegistryCallerSession) GetNodeWithdrawableTokens(nodeAddress common.Address) (*big.Int, *big.Int, error) {
	return _NodeRegistry.Contract.GetNodeWithdrawableTokens(&_NodeRegistry.CallOpts, nodeAddress)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeRegistry *NodeRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeRegistry *NodeRegistrySession) Owner() (common.Address, error) {
	return _NodeRegistry.Contract.Owner(&_NodeRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NodeRegistry *NodeRegistryCallerSession) Owner() (common.Address, error) {
	return _NodeRegistry.Contract.Owner(&_NodeRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeRegistry *NodeRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _NodeRegistry.Contract.ProxiableUUID(&_NodeRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NodeRegistry *NodeRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NodeRegistry.Contract.ProxiableUUID(&_NodeRegistry.CallOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x914eb34d.
//
// Solidity: function addReward(address[] nodes, uint256 ethAmount, uint256 arpaAmount) returns()
func (_NodeRegistry *NodeRegistryTransactor) AddReward(opts *bind.TransactOpts, nodes []common.Address, ethAmount *big.Int, arpaAmount *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "addReward", nodes, ethAmount, arpaAmount)
}

// AddReward is a paid mutator transaction binding the contract method 0x914eb34d.
//
// Solidity: function addReward(address[] nodes, uint256 ethAmount, uint256 arpaAmount) returns()
func (_NodeRegistry *NodeRegistrySession) AddReward(nodes []common.Address, ethAmount *big.Int, arpaAmount *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.AddReward(&_NodeRegistry.TransactOpts, nodes, ethAmount, arpaAmount)
}

// AddReward is a paid mutator transaction binding the contract method 0x914eb34d.
//
// Solidity: function addReward(address[] nodes, uint256 ethAmount, uint256 arpaAmount) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) AddReward(nodes []common.Address, ethAmount *big.Int, arpaAmount *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.AddReward(&_NodeRegistry.TransactOpts, nodes, ethAmount, arpaAmount)
}

// ChangeDkgPublicKey is a paid mutator transaction binding the contract method 0xe275cde6.
//
// Solidity: function changeDkgPublicKey(bytes dkgPublicKey) returns()
func (_NodeRegistry *NodeRegistryTransactor) ChangeDkgPublicKey(opts *bind.TransactOpts, dkgPublicKey []byte) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "changeDkgPublicKey", dkgPublicKey)
}

// ChangeDkgPublicKey is a paid mutator transaction binding the contract method 0xe275cde6.
//
// Solidity: function changeDkgPublicKey(bytes dkgPublicKey) returns()
func (_NodeRegistry *NodeRegistrySession) ChangeDkgPublicKey(dkgPublicKey []byte) (*types.Transaction, error) {
	return _NodeRegistry.Contract.ChangeDkgPublicKey(&_NodeRegistry.TransactOpts, dkgPublicKey)
}

// ChangeDkgPublicKey is a paid mutator transaction binding the contract method 0xe275cde6.
//
// Solidity: function changeDkgPublicKey(bytes dkgPublicKey) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) ChangeDkgPublicKey(dkgPublicKey []byte) (*types.Transaction, error) {
	return _NodeRegistry.Contract.ChangeDkgPublicKey(&_NodeRegistry.TransactOpts, dkgPublicKey)
}

// DismissNode is a paid mutator transaction binding the contract method 0xadd60b4c.
//
// Solidity: function dismissNode(address nodeIdAddress, uint256 pendingBlock) returns()
func (_NodeRegistry *NodeRegistryTransactor) DismissNode(opts *bind.TransactOpts, nodeIdAddress common.Address, pendingBlock *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "dismissNode", nodeIdAddress, pendingBlock)
}

// DismissNode is a paid mutator transaction binding the contract method 0xadd60b4c.
//
// Solidity: function dismissNode(address nodeIdAddress, uint256 pendingBlock) returns()
func (_NodeRegistry *NodeRegistrySession) DismissNode(nodeIdAddress common.Address, pendingBlock *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.DismissNode(&_NodeRegistry.TransactOpts, nodeIdAddress, pendingBlock)
}

// DismissNode is a paid mutator transaction binding the contract method 0xadd60b4c.
//
// Solidity: function dismissNode(address nodeIdAddress, uint256 pendingBlock) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) DismissNode(nodeIdAddress common.Address, pendingBlock *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.DismissNode(&_NodeRegistry.TransactOpts, nodeIdAddress, pendingBlock)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address arpa) returns()
func (_NodeRegistry *NodeRegistryTransactor) Initialize(opts *bind.TransactOpts, arpa common.Address) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "initialize", arpa)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address arpa) returns()
func (_NodeRegistry *NodeRegistrySession) Initialize(arpa common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.Initialize(&_NodeRegistry.TransactOpts, arpa)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address arpa) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) Initialize(arpa common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.Initialize(&_NodeRegistry.TransactOpts, arpa)
}

// NodeActivate is a paid mutator transaction binding the contract method 0x8d2f3e6b.
//
// Solidity: function nodeActivate((bytes,bytes32,uint256) assetAccountSignature) returns()
func (_NodeRegistry *NodeRegistryTransactor) NodeActivate(opts *bind.TransactOpts, assetAccountSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "nodeActivate", assetAccountSignature)
}

// NodeActivate is a paid mutator transaction binding the contract method 0x8d2f3e6b.
//
// Solidity: function nodeActivate((bytes,bytes32,uint256) assetAccountSignature) returns()
func (_NodeRegistry *NodeRegistrySession) NodeActivate(assetAccountSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeActivate(&_NodeRegistry.TransactOpts, assetAccountSignature)
}

// NodeActivate is a paid mutator transaction binding the contract method 0x8d2f3e6b.
//
// Solidity: function nodeActivate((bytes,bytes32,uint256) assetAccountSignature) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) NodeActivate(assetAccountSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeActivate(&_NodeRegistry.TransactOpts, assetAccountSignature)
}

// NodeLogOff is a paid mutator transaction binding the contract method 0x15eb6661.
//
// Solidity: function nodeLogOff() returns()
func (_NodeRegistry *NodeRegistryTransactor) NodeLogOff(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "nodeLogOff")
}

// NodeLogOff is a paid mutator transaction binding the contract method 0x15eb6661.
//
// Solidity: function nodeLogOff() returns()
func (_NodeRegistry *NodeRegistrySession) NodeLogOff() (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeLogOff(&_NodeRegistry.TransactOpts)
}

// NodeLogOff is a paid mutator transaction binding the contract method 0x15eb6661.
//
// Solidity: function nodeLogOff() returns()
func (_NodeRegistry *NodeRegistryTransactorSession) NodeLogOff() (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeLogOff(&_NodeRegistry.TransactOpts)
}

// NodeQuit is a paid mutator transaction binding the contract method 0x7a2af56e.
//
// Solidity: function nodeQuit() returns()
func (_NodeRegistry *NodeRegistryTransactor) NodeQuit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "nodeQuit")
}

// NodeQuit is a paid mutator transaction binding the contract method 0x7a2af56e.
//
// Solidity: function nodeQuit() returns()
func (_NodeRegistry *NodeRegistrySession) NodeQuit() (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeQuit(&_NodeRegistry.TransactOpts)
}

// NodeQuit is a paid mutator transaction binding the contract method 0x7a2af56e.
//
// Solidity: function nodeQuit() returns()
func (_NodeRegistry *NodeRegistryTransactorSession) NodeQuit() (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeQuit(&_NodeRegistry.TransactOpts)
}

// NodeRegister is a paid mutator transaction binding the contract method 0x30d640b2.
//
// Solidity: function nodeRegister(bytes dkgPublicKey, bool isEigenlayerNode, address assetAccountAddress, (bytes,bytes32,uint256) assetAccountSignature) returns()
func (_NodeRegistry *NodeRegistryTransactor) NodeRegister(opts *bind.TransactOpts, dkgPublicKey []byte, isEigenlayerNode bool, assetAccountAddress common.Address, assetAccountSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "nodeRegister", dkgPublicKey, isEigenlayerNode, assetAccountAddress, assetAccountSignature)
}

// NodeRegister is a paid mutator transaction binding the contract method 0x30d640b2.
//
// Solidity: function nodeRegister(bytes dkgPublicKey, bool isEigenlayerNode, address assetAccountAddress, (bytes,bytes32,uint256) assetAccountSignature) returns()
func (_NodeRegistry *NodeRegistrySession) NodeRegister(dkgPublicKey []byte, isEigenlayerNode bool, assetAccountAddress common.Address, assetAccountSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeRegister(&_NodeRegistry.TransactOpts, dkgPublicKey, isEigenlayerNode, assetAccountAddress, assetAccountSignature)
}

// NodeRegister is a paid mutator transaction binding the contract method 0x30d640b2.
//
// Solidity: function nodeRegister(bytes dkgPublicKey, bool isEigenlayerNode, address assetAccountAddress, (bytes,bytes32,uint256) assetAccountSignature) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) NodeRegister(dkgPublicKey []byte, isEigenlayerNode bool, assetAccountAddress common.Address, assetAccountSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeRegister(&_NodeRegistry.TransactOpts, dkgPublicKey, isEigenlayerNode, assetAccountAddress, assetAccountSignature)
}

// NodeWithdraw is a paid mutator transaction binding the contract method 0x4ecea80d.
//
// Solidity: function nodeWithdraw(address recipient) returns()
func (_NodeRegistry *NodeRegistryTransactor) NodeWithdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "nodeWithdraw", recipient)
}

// NodeWithdraw is a paid mutator transaction binding the contract method 0x4ecea80d.
//
// Solidity: function nodeWithdraw(address recipient) returns()
func (_NodeRegistry *NodeRegistrySession) NodeWithdraw(recipient common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeWithdraw(&_NodeRegistry.TransactOpts, recipient)
}

// NodeWithdraw is a paid mutator transaction binding the contract method 0x4ecea80d.
//
// Solidity: function nodeWithdraw(address recipient) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) NodeWithdraw(recipient common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.NodeWithdraw(&_NodeRegistry.TransactOpts, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeRegistry *NodeRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeRegistry *NodeRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeRegistry.Contract.RenounceOwnership(&_NodeRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NodeRegistry *NodeRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NodeRegistry.Contract.RenounceOwnership(&_NodeRegistry.TransactOpts)
}

// SetAssetAccount is a paid mutator transaction binding the contract method 0x081342ba.
//
// Solidity: function setAssetAccount(address[] assetAccountAddresses, address[] nodeAddresses) returns()
func (_NodeRegistry *NodeRegistryTransactor) SetAssetAccount(opts *bind.TransactOpts, assetAccountAddresses []common.Address, nodeAddresses []common.Address) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "setAssetAccount", assetAccountAddresses, nodeAddresses)
}

// SetAssetAccount is a paid mutator transaction binding the contract method 0x081342ba.
//
// Solidity: function setAssetAccount(address[] assetAccountAddresses, address[] nodeAddresses) returns()
func (_NodeRegistry *NodeRegistrySession) SetAssetAccount(assetAccountAddresses []common.Address, nodeAddresses []common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.SetAssetAccount(&_NodeRegistry.TransactOpts, assetAccountAddresses, nodeAddresses)
}

// SetAssetAccount is a paid mutator transaction binding the contract method 0x081342ba.
//
// Solidity: function setAssetAccount(address[] assetAccountAddresses, address[] nodeAddresses) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) SetAssetAccount(assetAccountAddresses []common.Address, nodeAddresses []common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.SetAssetAccount(&_NodeRegistry.TransactOpts, assetAccountAddresses, nodeAddresses)
}

// SetNodeRegistryConfig is a paid mutator transaction binding the contract method 0x33f010a1.
//
// Solidity: function setNodeRegistryConfig(address controllerContractAddress, address stakingContractAddress, address serviceManagerContractAddress, uint256 nativeNodeStakingAmount, uint256 eigenlayerNodeStakingAmount, uint256 pendingBlockAfterQuit) returns()
func (_NodeRegistry *NodeRegistryTransactor) SetNodeRegistryConfig(opts *bind.TransactOpts, controllerContractAddress common.Address, stakingContractAddress common.Address, serviceManagerContractAddress common.Address, nativeNodeStakingAmount *big.Int, eigenlayerNodeStakingAmount *big.Int, pendingBlockAfterQuit *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "setNodeRegistryConfig", controllerContractAddress, stakingContractAddress, serviceManagerContractAddress, nativeNodeStakingAmount, eigenlayerNodeStakingAmount, pendingBlockAfterQuit)
}

// SetNodeRegistryConfig is a paid mutator transaction binding the contract method 0x33f010a1.
//
// Solidity: function setNodeRegistryConfig(address controllerContractAddress, address stakingContractAddress, address serviceManagerContractAddress, uint256 nativeNodeStakingAmount, uint256 eigenlayerNodeStakingAmount, uint256 pendingBlockAfterQuit) returns()
func (_NodeRegistry *NodeRegistrySession) SetNodeRegistryConfig(controllerContractAddress common.Address, stakingContractAddress common.Address, serviceManagerContractAddress common.Address, nativeNodeStakingAmount *big.Int, eigenlayerNodeStakingAmount *big.Int, pendingBlockAfterQuit *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.SetNodeRegistryConfig(&_NodeRegistry.TransactOpts, controllerContractAddress, stakingContractAddress, serviceManagerContractAddress, nativeNodeStakingAmount, eigenlayerNodeStakingAmount, pendingBlockAfterQuit)
}

// SetNodeRegistryConfig is a paid mutator transaction binding the contract method 0x33f010a1.
//
// Solidity: function setNodeRegistryConfig(address controllerContractAddress, address stakingContractAddress, address serviceManagerContractAddress, uint256 nativeNodeStakingAmount, uint256 eigenlayerNodeStakingAmount, uint256 pendingBlockAfterQuit) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) SetNodeRegistryConfig(controllerContractAddress common.Address, stakingContractAddress common.Address, serviceManagerContractAddress common.Address, nativeNodeStakingAmount *big.Int, eigenlayerNodeStakingAmount *big.Int, pendingBlockAfterQuit *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.SetNodeRegistryConfig(&_NodeRegistry.TransactOpts, controllerContractAddress, stakingContractAddress, serviceManagerContractAddress, nativeNodeStakingAmount, eigenlayerNodeStakingAmount, pendingBlockAfterQuit)
}

// SlashNode is a paid mutator transaction binding the contract method 0x8ed47008.
//
// Solidity: function slashNode(address nodeIdAddress, uint256 stakingRewardPenalty, uint256 pendingBlock) returns()
func (_NodeRegistry *NodeRegistryTransactor) SlashNode(opts *bind.TransactOpts, nodeIdAddress common.Address, stakingRewardPenalty *big.Int, pendingBlock *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "slashNode", nodeIdAddress, stakingRewardPenalty, pendingBlock)
}

// SlashNode is a paid mutator transaction binding the contract method 0x8ed47008.
//
// Solidity: function slashNode(address nodeIdAddress, uint256 stakingRewardPenalty, uint256 pendingBlock) returns()
func (_NodeRegistry *NodeRegistrySession) SlashNode(nodeIdAddress common.Address, stakingRewardPenalty *big.Int, pendingBlock *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.SlashNode(&_NodeRegistry.TransactOpts, nodeIdAddress, stakingRewardPenalty, pendingBlock)
}

// SlashNode is a paid mutator transaction binding the contract method 0x8ed47008.
//
// Solidity: function slashNode(address nodeIdAddress, uint256 stakingRewardPenalty, uint256 pendingBlock) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) SlashNode(nodeIdAddress common.Address, stakingRewardPenalty *big.Int, pendingBlock *big.Int) (*types.Transaction, error) {
	return _NodeRegistry.Contract.SlashNode(&_NodeRegistry.TransactOpts, nodeIdAddress, stakingRewardPenalty, pendingBlock)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeRegistry *NodeRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeRegistry *NodeRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.TransferOwnership(&_NodeRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.TransferOwnership(&_NodeRegistry.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NodeRegistry *NodeRegistryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NodeRegistry *NodeRegistrySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.UpgradeTo(&_NodeRegistry.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_NodeRegistry *NodeRegistryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _NodeRegistry.Contract.UpgradeTo(&_NodeRegistry.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeRegistry *NodeRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeRegistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeRegistry *NodeRegistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeRegistry.Contract.UpgradeToAndCall(&_NodeRegistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NodeRegistry *NodeRegistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NodeRegistry.Contract.UpgradeToAndCall(&_NodeRegistry.TransactOpts, newImplementation, data)
}

// NodeRegistryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the NodeRegistry contract.
type NodeRegistryAdminChangedIterator struct {
	Event *NodeRegistryAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodeRegistryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryAdminChanged)
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
		it.Event = new(NodeRegistryAdminChanged)
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
func (it *NodeRegistryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryAdminChanged represents a AdminChanged event raised by the NodeRegistry contract.
type NodeRegistryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NodeRegistry *NodeRegistryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NodeRegistryAdminChangedIterator, error) {

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NodeRegistryAdminChangedIterator{contract: _NodeRegistry.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_NodeRegistry *NodeRegistryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NodeRegistryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryAdminChanged)
				if err := _NodeRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_NodeRegistry *NodeRegistryFilterer) ParseAdminChanged(log types.Log) (*NodeRegistryAdminChanged, error) {
	event := new(NodeRegistryAdminChanged)
	if err := _NodeRegistry.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryAssetAccountSetIterator is returned from FilterAssetAccountSet and is used to iterate over the raw logs and unpacked data for AssetAccountSet events raised by the NodeRegistry contract.
type NodeRegistryAssetAccountSetIterator struct {
	Event *NodeRegistryAssetAccountSet // Event containing the contract specifics and raw log

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
func (it *NodeRegistryAssetAccountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryAssetAccountSet)
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
		it.Event = new(NodeRegistryAssetAccountSet)
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
func (it *NodeRegistryAssetAccountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryAssetAccountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryAssetAccountSet represents a AssetAccountSet event raised by the NodeRegistry contract.
type NodeRegistryAssetAccountSet struct {
	AssetAccountAddress common.Address
	NodeAddress         common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAssetAccountSet is a free log retrieval operation binding the contract event 0x89ca1a6d1ba2dd7a1222041d072dd6d2e7790f80973456811e834bb9cabedbea.
//
// Solidity: event AssetAccountSet(address indexed assetAccountAddress, address indexed nodeAddress)
func (_NodeRegistry *NodeRegistryFilterer) FilterAssetAccountSet(opts *bind.FilterOpts, assetAccountAddress []common.Address, nodeAddress []common.Address) (*NodeRegistryAssetAccountSetIterator, error) {

	var assetAccountAddressRule []interface{}
	for _, assetAccountAddressItem := range assetAccountAddress {
		assetAccountAddressRule = append(assetAccountAddressRule, assetAccountAddressItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "AssetAccountSet", assetAccountAddressRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryAssetAccountSetIterator{contract: _NodeRegistry.contract, event: "AssetAccountSet", logs: logs, sub: sub}, nil
}

// WatchAssetAccountSet is a free log subscription operation binding the contract event 0x89ca1a6d1ba2dd7a1222041d072dd6d2e7790f80973456811e834bb9cabedbea.
//
// Solidity: event AssetAccountSet(address indexed assetAccountAddress, address indexed nodeAddress)
func (_NodeRegistry *NodeRegistryFilterer) WatchAssetAccountSet(opts *bind.WatchOpts, sink chan<- *NodeRegistryAssetAccountSet, assetAccountAddress []common.Address, nodeAddress []common.Address) (event.Subscription, error) {

	var assetAccountAddressRule []interface{}
	for _, assetAccountAddressItem := range assetAccountAddress {
		assetAccountAddressRule = append(assetAccountAddressRule, assetAccountAddressItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "AssetAccountSet", assetAccountAddressRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryAssetAccountSet)
				if err := _NodeRegistry.contract.UnpackLog(event, "AssetAccountSet", log); err != nil {
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

// ParseAssetAccountSet is a log parse operation binding the contract event 0x89ca1a6d1ba2dd7a1222041d072dd6d2e7790f80973456811e834bb9cabedbea.
//
// Solidity: event AssetAccountSet(address indexed assetAccountAddress, address indexed nodeAddress)
func (_NodeRegistry *NodeRegistryFilterer) ParseAssetAccountSet(log types.Log) (*NodeRegistryAssetAccountSet, error) {
	event := new(NodeRegistryAssetAccountSet)
	if err := _NodeRegistry.contract.UnpackLog(event, "AssetAccountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the NodeRegistry contract.
type NodeRegistryBeaconUpgradedIterator struct {
	Event *NodeRegistryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NodeRegistryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryBeaconUpgraded)
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
		it.Event = new(NodeRegistryBeaconUpgraded)
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
func (it *NodeRegistryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryBeaconUpgraded represents a BeaconUpgraded event raised by the NodeRegistry contract.
type NodeRegistryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NodeRegistry *NodeRegistryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NodeRegistryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryBeaconUpgradedIterator{contract: _NodeRegistry.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_NodeRegistry *NodeRegistryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NodeRegistryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryBeaconUpgraded)
				if err := _NodeRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_NodeRegistry *NodeRegistryFilterer) ParseBeaconUpgraded(log types.Log) (*NodeRegistryBeaconUpgraded, error) {
	event := new(NodeRegistryBeaconUpgraded)
	if err := _NodeRegistry.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryDkgPublicKeyChangedIterator is returned from FilterDkgPublicKeyChanged and is used to iterate over the raw logs and unpacked data for DkgPublicKeyChanged events raised by the NodeRegistry contract.
type NodeRegistryDkgPublicKeyChangedIterator struct {
	Event *NodeRegistryDkgPublicKeyChanged // Event containing the contract specifics and raw log

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
func (it *NodeRegistryDkgPublicKeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryDkgPublicKeyChanged)
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
		it.Event = new(NodeRegistryDkgPublicKeyChanged)
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
func (it *NodeRegistryDkgPublicKeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryDkgPublicKeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryDkgPublicKeyChanged represents a DkgPublicKeyChanged event raised by the NodeRegistry contract.
type NodeRegistryDkgPublicKeyChanged struct {
	NodeAddress  common.Address
	DkgPublicKey []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDkgPublicKeyChanged is a free log retrieval operation binding the contract event 0x4a327ac4843af7a9586b5a2a2c312bd17289ae1d70da32855ed539fe39f86a50.
//
// Solidity: event DkgPublicKeyChanged(address indexed nodeAddress, bytes dkgPublicKey)
func (_NodeRegistry *NodeRegistryFilterer) FilterDkgPublicKeyChanged(opts *bind.FilterOpts, nodeAddress []common.Address) (*NodeRegistryDkgPublicKeyChangedIterator, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "DkgPublicKeyChanged", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryDkgPublicKeyChangedIterator{contract: _NodeRegistry.contract, event: "DkgPublicKeyChanged", logs: logs, sub: sub}, nil
}

// WatchDkgPublicKeyChanged is a free log subscription operation binding the contract event 0x4a327ac4843af7a9586b5a2a2c312bd17289ae1d70da32855ed539fe39f86a50.
//
// Solidity: event DkgPublicKeyChanged(address indexed nodeAddress, bytes dkgPublicKey)
func (_NodeRegistry *NodeRegistryFilterer) WatchDkgPublicKeyChanged(opts *bind.WatchOpts, sink chan<- *NodeRegistryDkgPublicKeyChanged, nodeAddress []common.Address) (event.Subscription, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "DkgPublicKeyChanged", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryDkgPublicKeyChanged)
				if err := _NodeRegistry.contract.UnpackLog(event, "DkgPublicKeyChanged", log); err != nil {
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

// ParseDkgPublicKeyChanged is a log parse operation binding the contract event 0x4a327ac4843af7a9586b5a2a2c312bd17289ae1d70da32855ed539fe39f86a50.
//
// Solidity: event DkgPublicKeyChanged(address indexed nodeAddress, bytes dkgPublicKey)
func (_NodeRegistry *NodeRegistryFilterer) ParseDkgPublicKeyChanged(log types.Log) (*NodeRegistryDkgPublicKeyChanged, error) {
	event := new(NodeRegistryDkgPublicKeyChanged)
	if err := _NodeRegistry.contract.UnpackLog(event, "DkgPublicKeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NodeRegistry contract.
type NodeRegistryInitializedIterator struct {
	Event *NodeRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *NodeRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryInitialized)
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
		it.Event = new(NodeRegistryInitialized)
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
func (it *NodeRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryInitialized represents a Initialized event raised by the NodeRegistry contract.
type NodeRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeRegistry *NodeRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodeRegistryInitializedIterator, error) {

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodeRegistryInitializedIterator{contract: _NodeRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NodeRegistry *NodeRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodeRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryInitialized)
				if err := _NodeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NodeRegistry *NodeRegistryFilterer) ParseInitialized(log types.Log) (*NodeRegistryInitialized, error) {
	event := new(NodeRegistryInitialized)
	if err := _NodeRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryNodeActivatedIterator is returned from FilterNodeActivated and is used to iterate over the raw logs and unpacked data for NodeActivated events raised by the NodeRegistry contract.
type NodeRegistryNodeActivatedIterator struct {
	Event *NodeRegistryNodeActivated // Event containing the contract specifics and raw log

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
func (it *NodeRegistryNodeActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryNodeActivated)
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
		it.Event = new(NodeRegistryNodeActivated)
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
func (it *NodeRegistryNodeActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryNodeActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryNodeActivated represents a NodeActivated event raised by the NodeRegistry contract.
type NodeRegistryNodeActivated struct {
	NodeAddress common.Address
	GroupIndex  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeActivated is a free log retrieval operation binding the contract event 0xfc97cd9154b40031874ef09a9436a4b60052e4dcf40f21b1258be265fac4a397.
//
// Solidity: event NodeActivated(address indexed nodeAddress, uint256 groupIndex)
func (_NodeRegistry *NodeRegistryFilterer) FilterNodeActivated(opts *bind.FilterOpts, nodeAddress []common.Address) (*NodeRegistryNodeActivatedIterator, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "NodeActivated", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryNodeActivatedIterator{contract: _NodeRegistry.contract, event: "NodeActivated", logs: logs, sub: sub}, nil
}

// WatchNodeActivated is a free log subscription operation binding the contract event 0xfc97cd9154b40031874ef09a9436a4b60052e4dcf40f21b1258be265fac4a397.
//
// Solidity: event NodeActivated(address indexed nodeAddress, uint256 groupIndex)
func (_NodeRegistry *NodeRegistryFilterer) WatchNodeActivated(opts *bind.WatchOpts, sink chan<- *NodeRegistryNodeActivated, nodeAddress []common.Address) (event.Subscription, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "NodeActivated", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryNodeActivated)
				if err := _NodeRegistry.contract.UnpackLog(event, "NodeActivated", log); err != nil {
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

// ParseNodeActivated is a log parse operation binding the contract event 0xfc97cd9154b40031874ef09a9436a4b60052e4dcf40f21b1258be265fac4a397.
//
// Solidity: event NodeActivated(address indexed nodeAddress, uint256 groupIndex)
func (_NodeRegistry *NodeRegistryFilterer) ParseNodeActivated(log types.Log) (*NodeRegistryNodeActivated, error) {
	event := new(NodeRegistryNodeActivated)
	if err := _NodeRegistry.contract.UnpackLog(event, "NodeActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryNodeQuitIterator is returned from FilterNodeQuit and is used to iterate over the raw logs and unpacked data for NodeQuit events raised by the NodeRegistry contract.
type NodeRegistryNodeQuitIterator struct {
	Event *NodeRegistryNodeQuit // Event containing the contract specifics and raw log

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
func (it *NodeRegistryNodeQuitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryNodeQuit)
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
		it.Event = new(NodeRegistryNodeQuit)
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
func (it *NodeRegistryNodeQuitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryNodeQuitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryNodeQuit represents a NodeQuit event raised by the NodeRegistry contract.
type NodeRegistryNodeQuit struct {
	NodeAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeQuit is a free log retrieval operation binding the contract event 0x68577adbb6b0647e21353ff032be5797d9fa0879ce7e05fe617e40368441f97d.
//
// Solidity: event NodeQuit(address indexed nodeAddress)
func (_NodeRegistry *NodeRegistryFilterer) FilterNodeQuit(opts *bind.FilterOpts, nodeAddress []common.Address) (*NodeRegistryNodeQuitIterator, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "NodeQuit", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryNodeQuitIterator{contract: _NodeRegistry.contract, event: "NodeQuit", logs: logs, sub: sub}, nil
}

// WatchNodeQuit is a free log subscription operation binding the contract event 0x68577adbb6b0647e21353ff032be5797d9fa0879ce7e05fe617e40368441f97d.
//
// Solidity: event NodeQuit(address indexed nodeAddress)
func (_NodeRegistry *NodeRegistryFilterer) WatchNodeQuit(opts *bind.WatchOpts, sink chan<- *NodeRegistryNodeQuit, nodeAddress []common.Address) (event.Subscription, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "NodeQuit", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryNodeQuit)
				if err := _NodeRegistry.contract.UnpackLog(event, "NodeQuit", log); err != nil {
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

// ParseNodeQuit is a log parse operation binding the contract event 0x68577adbb6b0647e21353ff032be5797d9fa0879ce7e05fe617e40368441f97d.
//
// Solidity: event NodeQuit(address indexed nodeAddress)
func (_NodeRegistry *NodeRegistryFilterer) ParseNodeQuit(log types.Log) (*NodeRegistryNodeQuit, error) {
	event := new(NodeRegistryNodeQuit)
	if err := _NodeRegistry.contract.UnpackLog(event, "NodeQuit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the NodeRegistry contract.
type NodeRegistryNodeRegisteredIterator struct {
	Event *NodeRegistryNodeRegistered // Event containing the contract specifics and raw log

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
func (it *NodeRegistryNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryNodeRegistered)
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
		it.Event = new(NodeRegistryNodeRegistered)
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
func (it *NodeRegistryNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryNodeRegistered represents a NodeRegistered event raised by the NodeRegistry contract.
type NodeRegistryNodeRegistered struct {
	NodeAddress  common.Address
	DkgPublicKey []byte
	GroupIndex   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0xd4ec586f4f9f417f99e20fe821fbaa10a10a4b95f8712a0c57c6d8ed970e98bd.
//
// Solidity: event NodeRegistered(address indexed nodeAddress, bytes dkgPublicKey, uint256 groupIndex)
func (_NodeRegistry *NodeRegistryFilterer) FilterNodeRegistered(opts *bind.FilterOpts, nodeAddress []common.Address) (*NodeRegistryNodeRegisteredIterator, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "NodeRegistered", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryNodeRegisteredIterator{contract: _NodeRegistry.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0xd4ec586f4f9f417f99e20fe821fbaa10a10a4b95f8712a0c57c6d8ed970e98bd.
//
// Solidity: event NodeRegistered(address indexed nodeAddress, bytes dkgPublicKey, uint256 groupIndex)
func (_NodeRegistry *NodeRegistryFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *NodeRegistryNodeRegistered, nodeAddress []common.Address) (event.Subscription, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "NodeRegistered", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryNodeRegistered)
				if err := _NodeRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
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

// ParseNodeRegistered is a log parse operation binding the contract event 0xd4ec586f4f9f417f99e20fe821fbaa10a10a4b95f8712a0c57c6d8ed970e98bd.
//
// Solidity: event NodeRegistered(address indexed nodeAddress, bytes dkgPublicKey, uint256 groupIndex)
func (_NodeRegistry *NodeRegistryFilterer) ParseNodeRegistered(log types.Log) (*NodeRegistryNodeRegistered, error) {
	event := new(NodeRegistryNodeRegistered)
	if err := _NodeRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryNodeRewardedIterator is returned from FilterNodeRewarded and is used to iterate over the raw logs and unpacked data for NodeRewarded events raised by the NodeRegistry contract.
type NodeRegistryNodeRewardedIterator struct {
	Event *NodeRegistryNodeRewarded // Event containing the contract specifics and raw log

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
func (it *NodeRegistryNodeRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryNodeRewarded)
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
		it.Event = new(NodeRegistryNodeRewarded)
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
func (it *NodeRegistryNodeRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryNodeRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryNodeRewarded represents a NodeRewarded event raised by the NodeRegistry contract.
type NodeRegistryNodeRewarded struct {
	NodeAddress common.Address
	EthAmount   *big.Int
	ArpaAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeRewarded is a free log retrieval operation binding the contract event 0x8353a804115421789f3ab2eeb3f5215943906ce12100c91d40fc865caf742b6f.
//
// Solidity: event NodeRewarded(address indexed nodeAddress, uint256 ethAmount, uint256 arpaAmount)
func (_NodeRegistry *NodeRegistryFilterer) FilterNodeRewarded(opts *bind.FilterOpts, nodeAddress []common.Address) (*NodeRegistryNodeRewardedIterator, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "NodeRewarded", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryNodeRewardedIterator{contract: _NodeRegistry.contract, event: "NodeRewarded", logs: logs, sub: sub}, nil
}

// WatchNodeRewarded is a free log subscription operation binding the contract event 0x8353a804115421789f3ab2eeb3f5215943906ce12100c91d40fc865caf742b6f.
//
// Solidity: event NodeRewarded(address indexed nodeAddress, uint256 ethAmount, uint256 arpaAmount)
func (_NodeRegistry *NodeRegistryFilterer) WatchNodeRewarded(opts *bind.WatchOpts, sink chan<- *NodeRegistryNodeRewarded, nodeAddress []common.Address) (event.Subscription, error) {

	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "NodeRewarded", nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryNodeRewarded)
				if err := _NodeRegistry.contract.UnpackLog(event, "NodeRewarded", log); err != nil {
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

// ParseNodeRewarded is a log parse operation binding the contract event 0x8353a804115421789f3ab2eeb3f5215943906ce12100c91d40fc865caf742b6f.
//
// Solidity: event NodeRewarded(address indexed nodeAddress, uint256 ethAmount, uint256 arpaAmount)
func (_NodeRegistry *NodeRegistryFilterer) ParseNodeRewarded(log types.Log) (*NodeRegistryNodeRewarded, error) {
	event := new(NodeRegistryNodeRewarded)
	if err := _NodeRegistry.contract.UnpackLog(event, "NodeRewarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryNodeSlashedIterator is returned from FilterNodeSlashed and is used to iterate over the raw logs and unpacked data for NodeSlashed events raised by the NodeRegistry contract.
type NodeRegistryNodeSlashedIterator struct {
	Event *NodeRegistryNodeSlashed // Event containing the contract specifics and raw log

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
func (it *NodeRegistryNodeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryNodeSlashed)
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
		it.Event = new(NodeRegistryNodeSlashed)
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
func (it *NodeRegistryNodeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryNodeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryNodeSlashed represents a NodeSlashed event raised by the NodeRegistry contract.
type NodeRegistryNodeSlashed struct {
	NodeIdAddress        common.Address
	StakingRewardPenalty *big.Int
	PendingBlock         *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNodeSlashed is a free log retrieval operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeIdAddress, uint256 stakingRewardPenalty, uint256 pendingBlock)
func (_NodeRegistry *NodeRegistryFilterer) FilterNodeSlashed(opts *bind.FilterOpts, nodeIdAddress []common.Address) (*NodeRegistryNodeSlashedIterator, error) {

	var nodeIdAddressRule []interface{}
	for _, nodeIdAddressItem := range nodeIdAddress {
		nodeIdAddressRule = append(nodeIdAddressRule, nodeIdAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "NodeSlashed", nodeIdAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryNodeSlashedIterator{contract: _NodeRegistry.contract, event: "NodeSlashed", logs: logs, sub: sub}, nil
}

// WatchNodeSlashed is a free log subscription operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeIdAddress, uint256 stakingRewardPenalty, uint256 pendingBlock)
func (_NodeRegistry *NodeRegistryFilterer) WatchNodeSlashed(opts *bind.WatchOpts, sink chan<- *NodeRegistryNodeSlashed, nodeIdAddress []common.Address) (event.Subscription, error) {

	var nodeIdAddressRule []interface{}
	for _, nodeIdAddressItem := range nodeIdAddress {
		nodeIdAddressRule = append(nodeIdAddressRule, nodeIdAddressItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "NodeSlashed", nodeIdAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryNodeSlashed)
				if err := _NodeRegistry.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
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

// ParseNodeSlashed is a log parse operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeIdAddress, uint256 stakingRewardPenalty, uint256 pendingBlock)
func (_NodeRegistry *NodeRegistryFilterer) ParseNodeSlashed(log types.Log) (*NodeRegistryNodeSlashed, error) {
	event := new(NodeRegistryNodeSlashed)
	if err := _NodeRegistry.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NodeRegistry contract.
type NodeRegistryOwnershipTransferredIterator struct {
	Event *NodeRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodeRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryOwnershipTransferred)
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
		it.Event = new(NodeRegistryOwnershipTransferred)
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
func (it *NodeRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the NodeRegistry contract.
type NodeRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeRegistry *NodeRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodeRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryOwnershipTransferredIterator{contract: _NodeRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NodeRegistry *NodeRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodeRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryOwnershipTransferred)
				if err := _NodeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NodeRegistry *NodeRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*NodeRegistryOwnershipTransferred, error) {
	event := new(NodeRegistryOwnershipTransferred)
	if err := _NodeRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NodeRegistry contract.
type NodeRegistryUpgradedIterator struct {
	Event *NodeRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *NodeRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeRegistryUpgraded)
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
		it.Event = new(NodeRegistryUpgraded)
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
func (it *NodeRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeRegistryUpgraded represents a Upgraded event raised by the NodeRegistry contract.
type NodeRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NodeRegistry *NodeRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodeRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NodeRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodeRegistryUpgradedIterator{contract: _NodeRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NodeRegistry *NodeRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodeRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NodeRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeRegistryUpgraded)
				if err := _NodeRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_NodeRegistry *NodeRegistryFilterer) ParseUpgraded(log types.Log) (*NodeRegistryUpgraded, error) {
	event := new(NodeRegistryUpgraded)
	if err := _NodeRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
