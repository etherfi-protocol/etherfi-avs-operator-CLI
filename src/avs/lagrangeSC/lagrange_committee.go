// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lagrangesc

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

// IBLSKeyCheckerBLSKeyWithProof is an auto generated low-level Go binding around an user-defined struct.
type IBLSKeyCheckerBLSKeyWithProof struct {
	BlsG1PublicKeys [][2]*big.Int
	AggG2PublicKey  [2][2]*big.Int
	Signature       [2]*big.Int
	Salt            [32]byte
	Expiry          *big.Int
}

// ILagrangeCommitteeCommitteeData is an auto generated low-level Go binding around an user-defined struct.
type ILagrangeCommitteeCommitteeData struct {
	Root         [32]byte
	UpdatedBlock *big.Int
	LeafCount    uint32
}

// LagrangeCommitteeMetaData contains all meta data concerning the LagrangeCommittee contract.
var LagrangeCommitteeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILagrangeService\",\"name\":\"_service\",\"type\":\"address\"},{\"internalType\":\"contractIVoteWeigher\",\"name\":\"_voteWeigher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orgLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"added\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"}],\"name\":\"BlsKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"genesisBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"minWeight\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"maxWeight\",\"type\":\"uint96\"}],\"name\":\"InitCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signAddress\",\"type\":\"address\"}],\"name\":\"SignAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"current\",\"type\":\"bytes32\"}],\"name\":\"UpdateCommittee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"l1Bias\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"genesisBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"minWeight\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"maxWeight\",\"type\":\"uint96\"}],\"name\":\"UpdateCommitteeParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLS_KEY_WITH_PROOF_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INNER_NODE_PREFIX\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEAF_NODE_PREFIX\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2][]\",\"name\":\"blsG1PublicKeys\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"aggG2PublicKey\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structIBLSKeyChecker.BLSKeyWithProof\",\"name\":\"blsKeyWithProof\",\"type\":\"tuple\"}],\"name\":\"addBlsPubKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256[2][]\",\"name\":\"blsG1PublicKeys\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"aggG2PublicKey\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structIBLSKeyChecker.BLSKeyWithProof\",\"name\":\"blsKeyWithProof\",\"type\":\"tuple\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"calculateKeyWithProofHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainIDs\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committeeAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"committeeParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"l1Bias\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"genesisBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"minWeight\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"maxWeight\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint224\",\"name\":\"updatedBlock\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"getBlsPubKeyVotingPowers\",\"outputs\":[{\"internalType\":\"uint96[]\",\"name\":\"individualVotingPowers\",\"type\":\"uint96[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getBlsPubKeys\",\"outputs\":[{\"internalType\":\"uint256[2][]\",\"name\":\"\",\"type\":\"uint256[2][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getCommittee\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint224\",\"name\":\"updatedBlock\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"leafCount\",\"type\":\"uint32\"}],\"internalType\":\"structILagrangeCommittee.CommitteeData\",\"name\":\"currentCommittee\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"opAddr\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"getOperatorVotingPower\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getTokenListForOperator\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"isSaltSpent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isUnregisterable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"isUpdatable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorsStatus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"subscribedChainCount\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"genesisBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"minWeight\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"maxWeight\",\"type\":\"uint96\"}],\"name\":\"registerChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32[]\",\"name\":\"indices\",\"type\":\"uint32[]\"}],\"name\":\"removeBlsPubKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"revertEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"service\",\"outputs\":[{\"internalType\":\"contractILagrangeService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"subscribeChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"subscribedChains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"unsubscribeByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"}],\"name\":\"unsubscribeChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint256[2][]\",\"name\":\"blsG1PublicKeys\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"aggG2PublicKey\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structIBLSKeyChecker.BLSKeyWithProof\",\"name\":\"blsKeyWithProof\",\"type\":\"tuple\"}],\"name\":\"updateBlsPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"chainID\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"l1Bias\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"genesisBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freezeDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"minWeight\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"maxWeight\",\"type\":\"uint96\"}],\"name\":\"updateChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newSignAddress\",\"type\":\"address\"}],\"name\":\"updateSignAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"updatedEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteWeigher\",\"outputs\":[{\"internalType\":\"contractIVoteWeigher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LagrangeCommitteeABI is the input ABI used to generate the binding from.
// Deprecated: Use LagrangeCommitteeMetaData.ABI instead.
var LagrangeCommitteeABI = LagrangeCommitteeMetaData.ABI

// LagrangeCommittee is an auto generated Go binding around an Ethereum contract.
type LagrangeCommittee struct {
	LagrangeCommitteeCaller     // Read-only binding to the contract
	LagrangeCommitteeTransactor // Write-only binding to the contract
	LagrangeCommitteeFilterer   // Log filterer for contract events
}

// LagrangeCommitteeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LagrangeCommitteeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeCommitteeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LagrangeCommitteeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeCommitteeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LagrangeCommitteeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LagrangeCommitteeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LagrangeCommitteeSession struct {
	Contract     *LagrangeCommittee // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LagrangeCommitteeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LagrangeCommitteeCallerSession struct {
	Contract *LagrangeCommitteeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LagrangeCommitteeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LagrangeCommitteeTransactorSession struct {
	Contract     *LagrangeCommitteeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LagrangeCommitteeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LagrangeCommitteeRaw struct {
	Contract *LagrangeCommittee // Generic contract binding to access the raw methods on
}

// LagrangeCommitteeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LagrangeCommitteeCallerRaw struct {
	Contract *LagrangeCommitteeCaller // Generic read-only contract binding to access the raw methods on
}

// LagrangeCommitteeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LagrangeCommitteeTransactorRaw struct {
	Contract *LagrangeCommitteeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLagrangeCommittee creates a new instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommittee(address common.Address, backend bind.ContractBackend) (*LagrangeCommittee, error) {
	contract, err := bindLagrangeCommittee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommittee{LagrangeCommitteeCaller: LagrangeCommitteeCaller{contract: contract}, LagrangeCommitteeTransactor: LagrangeCommitteeTransactor{contract: contract}, LagrangeCommitteeFilterer: LagrangeCommitteeFilterer{contract: contract}}, nil
}

// NewLagrangeCommitteeCaller creates a new read-only instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommitteeCaller(address common.Address, caller bind.ContractCaller) (*LagrangeCommitteeCaller, error) {
	contract, err := bindLagrangeCommittee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeCaller{contract: contract}, nil
}

// NewLagrangeCommitteeTransactor creates a new write-only instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommitteeTransactor(address common.Address, transactor bind.ContractTransactor) (*LagrangeCommitteeTransactor, error) {
	contract, err := bindLagrangeCommittee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeTransactor{contract: contract}, nil
}

// NewLagrangeCommitteeFilterer creates a new log filterer instance of LagrangeCommittee, bound to a specific deployed contract.
func NewLagrangeCommitteeFilterer(address common.Address, filterer bind.ContractFilterer) (*LagrangeCommitteeFilterer, error) {
	contract, err := bindLagrangeCommittee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeFilterer{contract: contract}, nil
}

// bindLagrangeCommittee binds a generic wrapper to an already deployed contract.
func bindLagrangeCommittee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LagrangeCommitteeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LagrangeCommittee *LagrangeCommitteeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LagrangeCommittee.Contract.LagrangeCommitteeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LagrangeCommittee *LagrangeCommitteeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.LagrangeCommitteeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LagrangeCommittee *LagrangeCommitteeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.LagrangeCommitteeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LagrangeCommittee *LagrangeCommitteeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LagrangeCommittee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LagrangeCommittee *LagrangeCommitteeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LagrangeCommittee *LagrangeCommitteeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.contract.Transact(opts, method, params...)
}

// BLSKEYWITHPROOFTYPEHASH is a free data retrieval call binding the contract method 0x1de8d88e.
//
// Solidity: function BLS_KEY_WITH_PROOF_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) BLSKEYWITHPROOFTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "BLS_KEY_WITH_PROOF_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLSKEYWITHPROOFTYPEHASH is a free data retrieval call binding the contract method 0x1de8d88e.
//
// Solidity: function BLS_KEY_WITH_PROOF_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) BLSKEYWITHPROOFTYPEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.BLSKEYWITHPROOFTYPEHASH(&_LagrangeCommittee.CallOpts)
}

// BLSKEYWITHPROOFTYPEHASH is a free data retrieval call binding the contract method 0x1de8d88e.
//
// Solidity: function BLS_KEY_WITH_PROOF_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) BLSKEYWITHPROOFTYPEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.BLSKEYWITHPROOFTYPEHASH(&_LagrangeCommittee.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.DOMAINTYPEHASH(&_LagrangeCommittee.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _LagrangeCommittee.Contract.DOMAINTYPEHASH(&_LagrangeCommittee.CallOpts)
}

// INNERNODEPREFIX is a free data retrieval call binding the contract method 0x727cb30f.
//
// Solidity: function INNER_NODE_PREFIX() view returns(bytes1)
func (_LagrangeCommittee *LagrangeCommitteeCaller) INNERNODEPREFIX(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "INNER_NODE_PREFIX")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INNERNODEPREFIX is a free data retrieval call binding the contract method 0x727cb30f.
//
// Solidity: function INNER_NODE_PREFIX() view returns(bytes1)
func (_LagrangeCommittee *LagrangeCommitteeSession) INNERNODEPREFIX() ([1]byte, error) {
	return _LagrangeCommittee.Contract.INNERNODEPREFIX(&_LagrangeCommittee.CallOpts)
}

// INNERNODEPREFIX is a free data retrieval call binding the contract method 0x727cb30f.
//
// Solidity: function INNER_NODE_PREFIX() view returns(bytes1)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) INNERNODEPREFIX() ([1]byte, error) {
	return _LagrangeCommittee.Contract.INNERNODEPREFIX(&_LagrangeCommittee.CallOpts)
}

// LEAFNODEPREFIX is a free data retrieval call binding the contract method 0x38fa59d1.
//
// Solidity: function LEAF_NODE_PREFIX() view returns(bytes1)
func (_LagrangeCommittee *LagrangeCommitteeCaller) LEAFNODEPREFIX(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "LEAF_NODE_PREFIX")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// LEAFNODEPREFIX is a free data retrieval call binding the contract method 0x38fa59d1.
//
// Solidity: function LEAF_NODE_PREFIX() view returns(bytes1)
func (_LagrangeCommittee *LagrangeCommitteeSession) LEAFNODEPREFIX() ([1]byte, error) {
	return _LagrangeCommittee.Contract.LEAFNODEPREFIX(&_LagrangeCommittee.CallOpts)
}

// LEAFNODEPREFIX is a free data retrieval call binding the contract method 0x38fa59d1.
//
// Solidity: function LEAF_NODE_PREFIX() view returns(bytes1)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) LEAFNODEPREFIX() ([1]byte, error) {
	return _LagrangeCommittee.Contract.LEAFNODEPREFIX(&_LagrangeCommittee.CallOpts)
}

// CalculateKeyWithProofHash is a free data retrieval call binding the contract method 0x9e224428.
//
// Solidity: function calculateKeyWithProofHash(address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CalculateKeyWithProofHash(opts *bind.CallOpts, operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "calculateKeyWithProofHash", operator, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateKeyWithProofHash is a free data retrieval call binding the contract method 0x9e224428.
//
// Solidity: function calculateKeyWithProofHash(address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) CalculateKeyWithProofHash(operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _LagrangeCommittee.Contract.CalculateKeyWithProofHash(&_LagrangeCommittee.CallOpts, operator, salt, expiry)
}

// CalculateKeyWithProofHash is a free data retrieval call binding the contract method 0x9e224428.
//
// Solidity: function calculateKeyWithProofHash(address operator, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CalculateKeyWithProofHash(operator common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _LagrangeCommittee.Contract.CalculateKeyWithProofHash(&_LagrangeCommittee.CallOpts, operator, salt, expiry)
}

// ChainIDs is a free data retrieval call binding the contract method 0x09d23e24.
//
// Solidity: function chainIDs(uint256 ) view returns(uint32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) ChainIDs(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "chainIDs", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChainIDs is a free data retrieval call binding the contract method 0x09d23e24.
//
// Solidity: function chainIDs(uint256 ) view returns(uint32)
func (_LagrangeCommittee *LagrangeCommitteeSession) ChainIDs(arg0 *big.Int) (uint32, error) {
	return _LagrangeCommittee.Contract.ChainIDs(&_LagrangeCommittee.CallOpts, arg0)
}

// ChainIDs is a free data retrieval call binding the contract method 0x09d23e24.
//
// Solidity: function chainIDs(uint256 ) view returns(uint32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) ChainIDs(arg0 *big.Int) (uint32, error) {
	return _LagrangeCommittee.Contract.ChainIDs(&_LagrangeCommittee.CallOpts, arg0)
}

// CommitteeAddrs is a free data retrieval call binding the contract method 0xbf988ab6.
//
// Solidity: function committeeAddrs(uint32 , uint256 ) view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CommitteeAddrs(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "committeeAddrs", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommitteeAddrs is a free data retrieval call binding the contract method 0xbf988ab6.
//
// Solidity: function committeeAddrs(uint32 , uint256 ) view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeAddrs(arg0 uint32, arg1 *big.Int) (common.Address, error) {
	return _LagrangeCommittee.Contract.CommitteeAddrs(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeAddrs is a free data retrieval call binding the contract method 0xbf988ab6.
//
// Solidity: function committeeAddrs(uint32 , uint256 ) view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CommitteeAddrs(arg0 uint32, arg1 *big.Int) (common.Address, error) {
	return _LagrangeCommittee.Contract.CommitteeAddrs(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// CommitteeParams is a free data retrieval call binding the contract method 0x72856455.
//
// Solidity: function committeeParams(uint32 ) view returns(uint256 startBlock, int256 l1Bias, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeCaller) CommitteeParams(opts *bind.CallOpts, arg0 uint32) (struct {
	StartBlock     *big.Int
	L1Bias         *big.Int
	GenesisBlock   *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	QuorumNumber   uint8
	MinWeight      *big.Int
	MaxWeight      *big.Int
}, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "committeeParams", arg0)

	outstruct := new(struct {
		StartBlock     *big.Int
		L1Bias         *big.Int
		GenesisBlock   *big.Int
		Duration       *big.Int
		FreezeDuration *big.Int
		QuorumNumber   uint8
		MinWeight      *big.Int
		MaxWeight      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.L1Bias = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.GenesisBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FreezeDuration = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.QuorumNumber = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.MinWeight = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MaxWeight = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CommitteeParams is a free data retrieval call binding the contract method 0x72856455.
//
// Solidity: function committeeParams(uint32 ) view returns(uint256 startBlock, int256 l1Bias, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeSession) CommitteeParams(arg0 uint32) (struct {
	StartBlock     *big.Int
	L1Bias         *big.Int
	GenesisBlock   *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	QuorumNumber   uint8
	MinWeight      *big.Int
	MaxWeight      *big.Int
}, error) {
	return _LagrangeCommittee.Contract.CommitteeParams(&_LagrangeCommittee.CallOpts, arg0)
}

// CommitteeParams is a free data retrieval call binding the contract method 0x72856455.
//
// Solidity: function committeeParams(uint32 ) view returns(uint256 startBlock, int256 l1Bias, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) CommitteeParams(arg0 uint32) (struct {
	StartBlock     *big.Int
	L1Bias         *big.Int
	GenesisBlock   *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	QuorumNumber   uint8
	MinWeight      *big.Int
	MaxWeight      *big.Int
}, error) {
	return _LagrangeCommittee.Contract.CommitteeParams(&_LagrangeCommittee.CallOpts, arg0)
}

// Committees is a free data retrieval call binding the contract method 0xa63490a2.
//
// Solidity: function committees(uint32 , uint256 ) view returns(bytes32 root, uint224 updatedBlock, uint32 leafCount)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Committees(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int) (struct {
	Root         [32]byte
	UpdatedBlock *big.Int
	LeafCount    uint32
}, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "committees", arg0, arg1)

	outstruct := new(struct {
		Root         [32]byte
		UpdatedBlock *big.Int
		LeafCount    uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.UpdatedBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LeafCount = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// Committees is a free data retrieval call binding the contract method 0xa63490a2.
//
// Solidity: function committees(uint32 , uint256 ) view returns(bytes32 root, uint224 updatedBlock, uint32 leafCount)
func (_LagrangeCommittee *LagrangeCommitteeSession) Committees(arg0 uint32, arg1 *big.Int) (struct {
	Root         [32]byte
	UpdatedBlock *big.Int
	LeafCount    uint32
}, error) {
	return _LagrangeCommittee.Contract.Committees(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// Committees is a free data retrieval call binding the contract method 0xa63490a2.
//
// Solidity: function committees(uint32 , uint256 ) view returns(bytes32 root, uint224 updatedBlock, uint32 leafCount)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Committees(arg0 uint32, arg1 *big.Int) (struct {
	Root         [32]byte
	UpdatedBlock *big.Int
	LeafCount    uint32
}, error) {
	return _LagrangeCommittee.Contract.Committees(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeSession) DomainSeparator() ([32]byte, error) {
	return _LagrangeCommittee.Contract.DomainSeparator(&_LagrangeCommittee.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) DomainSeparator() ([32]byte, error) {
	return _LagrangeCommittee.Contract.DomainSeparator(&_LagrangeCommittee.CallOpts)
}

// GetBlsPubKeyVotingPowers is a free data retrieval call binding the contract method 0x24ca7a48.
//
// Solidity: function getBlsPubKeyVotingPowers(address opAddr, uint32 chainID) view returns(uint96[] individualVotingPowers)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetBlsPubKeyVotingPowers(opts *bind.CallOpts, opAddr common.Address, chainID uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getBlsPubKeyVotingPowers", opAddr, chainID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBlsPubKeyVotingPowers is a free data retrieval call binding the contract method 0x24ca7a48.
//
// Solidity: function getBlsPubKeyVotingPowers(address opAddr, uint32 chainID) view returns(uint96[] individualVotingPowers)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetBlsPubKeyVotingPowers(opAddr common.Address, chainID uint32) ([]*big.Int, error) {
	return _LagrangeCommittee.Contract.GetBlsPubKeyVotingPowers(&_LagrangeCommittee.CallOpts, opAddr, chainID)
}

// GetBlsPubKeyVotingPowers is a free data retrieval call binding the contract method 0x24ca7a48.
//
// Solidity: function getBlsPubKeyVotingPowers(address opAddr, uint32 chainID) view returns(uint96[] individualVotingPowers)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetBlsPubKeyVotingPowers(opAddr common.Address, chainID uint32) ([]*big.Int, error) {
	return _LagrangeCommittee.Contract.GetBlsPubKeyVotingPowers(&_LagrangeCommittee.CallOpts, opAddr, chainID)
}

// GetBlsPubKeys is a free data retrieval call binding the contract method 0xa43dd8dd.
//
// Solidity: function getBlsPubKeys(address operator) view returns(uint256[2][])
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetBlsPubKeys(opts *bind.CallOpts, operator common.Address) ([][2]*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getBlsPubKeys", operator)

	if err != nil {
		return *new([][2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]*big.Int)).(*[][2]*big.Int)

	return out0, err

}

// GetBlsPubKeys is a free data retrieval call binding the contract method 0xa43dd8dd.
//
// Solidity: function getBlsPubKeys(address operator) view returns(uint256[2][])
func (_LagrangeCommittee *LagrangeCommitteeSession) GetBlsPubKeys(operator common.Address) ([][2]*big.Int, error) {
	return _LagrangeCommittee.Contract.GetBlsPubKeys(&_LagrangeCommittee.CallOpts, operator)
}

// GetBlsPubKeys is a free data retrieval call binding the contract method 0xa43dd8dd.
//
// Solidity: function getBlsPubKeys(address operator) view returns(uint256[2][])
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetBlsPubKeys(operator common.Address) ([][2]*big.Int, error) {
	return _LagrangeCommittee.Contract.GetBlsPubKeys(&_LagrangeCommittee.CallOpts, operator)
}

// GetCommittee is a free data retrieval call binding the contract method 0xdef9e7d5.
//
// Solidity: function getCommittee(uint32 chainID, uint256 blockNumber) view returns((bytes32,uint224,uint32) currentCommittee)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetCommittee(opts *bind.CallOpts, chainID uint32, blockNumber *big.Int) (ILagrangeCommitteeCommitteeData, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getCommittee", chainID, blockNumber)

	if err != nil {
		return *new(ILagrangeCommitteeCommitteeData), err
	}

	out0 := *abi.ConvertType(out[0], new(ILagrangeCommitteeCommitteeData)).(*ILagrangeCommitteeCommitteeData)

	return out0, err

}

// GetCommittee is a free data retrieval call binding the contract method 0xdef9e7d5.
//
// Solidity: function getCommittee(uint32 chainID, uint256 blockNumber) view returns((bytes32,uint224,uint32) currentCommittee)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetCommittee(chainID uint32, blockNumber *big.Int) (ILagrangeCommitteeCommitteeData, error) {
	return _LagrangeCommittee.Contract.GetCommittee(&_LagrangeCommittee.CallOpts, chainID, blockNumber)
}

// GetCommittee is a free data retrieval call binding the contract method 0xdef9e7d5.
//
// Solidity: function getCommittee(uint32 chainID, uint256 blockNumber) view returns((bytes32,uint224,uint32) currentCommittee)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetCommittee(chainID uint32, blockNumber *big.Int) (ILagrangeCommitteeCommitteeData, error) {
	return _LagrangeCommittee.Contract.GetCommittee(&_LagrangeCommittee.CallOpts, chainID, blockNumber)
}

// GetEpochInterval is a free data retrieval call binding the contract method 0x0f225b84.
//
// Solidity: function getEpochInterval(uint32 chainID, uint256 epochNumber) view returns(uint256 startBlock, uint256 freezeBlock, uint256 endBlock)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetEpochInterval(opts *bind.CallOpts, chainID uint32, epochNumber *big.Int) (struct {
	StartBlock  *big.Int
	FreezeBlock *big.Int
	EndBlock    *big.Int
}, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getEpochInterval", chainID, epochNumber)

	outstruct := new(struct {
		StartBlock  *big.Int
		FreezeBlock *big.Int
		EndBlock    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FreezeBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEpochInterval is a free data retrieval call binding the contract method 0x0f225b84.
//
// Solidity: function getEpochInterval(uint32 chainID, uint256 epochNumber) view returns(uint256 startBlock, uint256 freezeBlock, uint256 endBlock)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetEpochInterval(chainID uint32, epochNumber *big.Int) (struct {
	StartBlock  *big.Int
	FreezeBlock *big.Int
	EndBlock    *big.Int
}, error) {
	return _LagrangeCommittee.Contract.GetEpochInterval(&_LagrangeCommittee.CallOpts, chainID, epochNumber)
}

// GetEpochInterval is a free data retrieval call binding the contract method 0x0f225b84.
//
// Solidity: function getEpochInterval(uint32 chainID, uint256 epochNumber) view returns(uint256 startBlock, uint256 freezeBlock, uint256 endBlock)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetEpochInterval(chainID uint32, epochNumber *big.Int) (struct {
	StartBlock  *big.Int
	FreezeBlock *big.Int
	EndBlock    *big.Int
}, error) {
	return _LagrangeCommittee.Contract.GetEpochInterval(&_LagrangeCommittee.CallOpts, chainID, epochNumber)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x78d81d08.
//
// Solidity: function getEpochNumber(uint32 chainID, uint256 blockNumber) view returns(uint256 epochNumber)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetEpochNumber(opts *bind.CallOpts, chainID uint32, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getEpochNumber", chainID, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x78d81d08.
//
// Solidity: function getEpochNumber(uint32 chainID, uint256 blockNumber) view returns(uint256 epochNumber)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetEpochNumber(chainID uint32, blockNumber *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetEpochNumber(&_LagrangeCommittee.CallOpts, chainID, blockNumber)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x78d81d08.
//
// Solidity: function getEpochNumber(uint32 chainID, uint256 blockNumber) view returns(uint256 epochNumber)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetEpochNumber(chainID uint32, blockNumber *big.Int) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetEpochNumber(&_LagrangeCommittee.CallOpts, chainID, blockNumber)
}

// GetOperatorVotingPower is a free data retrieval call binding the contract method 0x5af1a88f.
//
// Solidity: function getOperatorVotingPower(address opAddr, uint32 chainID) view returns(uint96)
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetOperatorVotingPower(opts *bind.CallOpts, opAddr common.Address, chainID uint32) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getOperatorVotingPower", opAddr, chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorVotingPower is a free data retrieval call binding the contract method 0x5af1a88f.
//
// Solidity: function getOperatorVotingPower(address opAddr, uint32 chainID) view returns(uint96)
func (_LagrangeCommittee *LagrangeCommitteeSession) GetOperatorVotingPower(opAddr common.Address, chainID uint32) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetOperatorVotingPower(&_LagrangeCommittee.CallOpts, opAddr, chainID)
}

// GetOperatorVotingPower is a free data retrieval call binding the contract method 0x5af1a88f.
//
// Solidity: function getOperatorVotingPower(address opAddr, uint32 chainID) view returns(uint96)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetOperatorVotingPower(opAddr common.Address, chainID uint32) (*big.Int, error) {
	return _LagrangeCommittee.Contract.GetOperatorVotingPower(&_LagrangeCommittee.CallOpts, opAddr, chainID)
}

// GetTokenListForOperator is a free data retrieval call binding the contract method 0x56bf7c25.
//
// Solidity: function getTokenListForOperator(address operator) view returns(address[])
func (_LagrangeCommittee *LagrangeCommitteeCaller) GetTokenListForOperator(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "getTokenListForOperator", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenListForOperator is a free data retrieval call binding the contract method 0x56bf7c25.
//
// Solidity: function getTokenListForOperator(address operator) view returns(address[])
func (_LagrangeCommittee *LagrangeCommitteeSession) GetTokenListForOperator(operator common.Address) ([]common.Address, error) {
	return _LagrangeCommittee.Contract.GetTokenListForOperator(&_LagrangeCommittee.CallOpts, operator)
}

// GetTokenListForOperator is a free data retrieval call binding the contract method 0x56bf7c25.
//
// Solidity: function getTokenListForOperator(address operator) view returns(address[])
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) GetTokenListForOperator(operator common.Address) ([]common.Address, error) {
	return _LagrangeCommittee.Contract.GetTokenListForOperator(&_LagrangeCommittee.CallOpts, operator)
}

// IsLocked is a free data retrieval call binding the contract method 0x3d6a2679.
//
// Solidity: function isLocked(uint32 chainID) view returns(bool, uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) IsLocked(opts *bind.CallOpts, chainID uint32) (bool, *big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "isLocked", chainID)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// IsLocked is a free data retrieval call binding the contract method 0x3d6a2679.
//
// Solidity: function isLocked(uint32 chainID) view returns(bool, uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) IsLocked(chainID uint32) (bool, *big.Int, error) {
	return _LagrangeCommittee.Contract.IsLocked(&_LagrangeCommittee.CallOpts, chainID)
}

// IsLocked is a free data retrieval call binding the contract method 0x3d6a2679.
//
// Solidity: function isLocked(uint32 chainID) view returns(bool, uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) IsLocked(chainID uint32) (bool, *big.Int, error) {
	return _LagrangeCommittee.Contract.IsLocked(&_LagrangeCommittee.CallOpts, chainID)
}

// IsSaltSpent is a free data retrieval call binding the contract method 0x3811db0d.
//
// Solidity: function isSaltSpent(address operator, bytes32 salt) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCaller) IsSaltSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "isSaltSpent", operator, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSaltSpent is a free data retrieval call binding the contract method 0x3811db0d.
//
// Solidity: function isSaltSpent(address operator, bytes32 salt) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeSession) IsSaltSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _LagrangeCommittee.Contract.IsSaltSpent(&_LagrangeCommittee.CallOpts, operator, salt)
}

// IsSaltSpent is a free data retrieval call binding the contract method 0x3811db0d.
//
// Solidity: function isSaltSpent(address operator, bytes32 salt) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) IsSaltSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _LagrangeCommittee.Contract.IsSaltSpent(&_LagrangeCommittee.CallOpts, operator, salt)
}

// IsUnregisterable is a free data retrieval call binding the contract method 0x19a74c5f.
//
// Solidity: function isUnregisterable(address operator) view returns(bool, uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) IsUnregisterable(opts *bind.CallOpts, operator common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "isUnregisterable", operator)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// IsUnregisterable is a free data retrieval call binding the contract method 0x19a74c5f.
//
// Solidity: function isUnregisterable(address operator) view returns(bool, uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) IsUnregisterable(operator common.Address) (bool, *big.Int, error) {
	return _LagrangeCommittee.Contract.IsUnregisterable(&_LagrangeCommittee.CallOpts, operator)
}

// IsUnregisterable is a free data retrieval call binding the contract method 0x19a74c5f.
//
// Solidity: function isUnregisterable(address operator) view returns(bool, uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) IsUnregisterable(operator common.Address) (bool, *big.Int, error) {
	return _LagrangeCommittee.Contract.IsUnregisterable(&_LagrangeCommittee.CallOpts, operator)
}

// IsUpdatable is a free data retrieval call binding the contract method 0x85ab9a7a.
//
// Solidity: function isUpdatable(uint32 chainID, uint256 epochNumber) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCaller) IsUpdatable(opts *bind.CallOpts, chainID uint32, epochNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "isUpdatable", chainID, epochNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpdatable is a free data retrieval call binding the contract method 0x85ab9a7a.
//
// Solidity: function isUpdatable(uint32 chainID, uint256 epochNumber) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeSession) IsUpdatable(chainID uint32, epochNumber *big.Int) (bool, error) {
	return _LagrangeCommittee.Contract.IsUpdatable(&_LagrangeCommittee.CallOpts, chainID, epochNumber)
}

// IsUpdatable is a free data retrieval call binding the contract method 0x85ab9a7a.
//
// Solidity: function isUpdatable(uint32 chainID, uint256 epochNumber) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) IsUpdatable(chainID uint32, epochNumber *big.Int) (bool, error) {
	return _LagrangeCommittee.Contract.IsUpdatable(&_LagrangeCommittee.CallOpts, chainID, epochNumber)
}

// OperatorsStatus is a free data retrieval call binding the contract method 0x95703d08.
//
// Solidity: function operatorsStatus(address ) view returns(address signAddress, uint8 subscribedChainCount)
func (_LagrangeCommittee *LagrangeCommitteeCaller) OperatorsStatus(opts *bind.CallOpts, arg0 common.Address) (struct {
	SignAddress          common.Address
	SubscribedChainCount uint8
}, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "operatorsStatus", arg0)

	outstruct := new(struct {
		SignAddress          common.Address
		SubscribedChainCount uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SignAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SubscribedChainCount = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// OperatorsStatus is a free data retrieval call binding the contract method 0x95703d08.
//
// Solidity: function operatorsStatus(address ) view returns(address signAddress, uint8 subscribedChainCount)
func (_LagrangeCommittee *LagrangeCommitteeSession) OperatorsStatus(arg0 common.Address) (struct {
	SignAddress          common.Address
	SubscribedChainCount uint8
}, error) {
	return _LagrangeCommittee.Contract.OperatorsStatus(&_LagrangeCommittee.CallOpts, arg0)
}

// OperatorsStatus is a free data retrieval call binding the contract method 0x95703d08.
//
// Solidity: function operatorsStatus(address ) view returns(address signAddress, uint8 subscribedChainCount)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) OperatorsStatus(arg0 common.Address) (struct {
	SignAddress          common.Address
	SubscribedChainCount uint8
}, error) {
	return _LagrangeCommittee.Contract.OperatorsStatus(&_LagrangeCommittee.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeSession) Owner() (common.Address, error) {
	return _LagrangeCommittee.Contract.Owner(&_LagrangeCommittee.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Owner() (common.Address, error) {
	return _LagrangeCommittee.Contract.Owner(&_LagrangeCommittee.CallOpts)
}

// Service is a free data retrieval call binding the contract method 0xd598d4c9.
//
// Solidity: function service() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCaller) Service(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "service")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Service is a free data retrieval call binding the contract method 0xd598d4c9.
//
// Solidity: function service() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeSession) Service() (common.Address, error) {
	return _LagrangeCommittee.Contract.Service(&_LagrangeCommittee.CallOpts)
}

// Service is a free data retrieval call binding the contract method 0xd598d4c9.
//
// Solidity: function service() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) Service() (common.Address, error) {
	return _LagrangeCommittee.Contract.Service(&_LagrangeCommittee.CallOpts)
}

// SubscribedChains is a free data retrieval call binding the contract method 0x53b7a0ff.
//
// Solidity: function subscribedChains(uint32 , address ) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCaller) SubscribedChains(opts *bind.CallOpts, arg0 uint32, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "subscribedChains", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubscribedChains is a free data retrieval call binding the contract method 0x53b7a0ff.
//
// Solidity: function subscribedChains(uint32 , address ) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeSession) SubscribedChains(arg0 uint32, arg1 common.Address) (bool, error) {
	return _LagrangeCommittee.Contract.SubscribedChains(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// SubscribedChains is a free data retrieval call binding the contract method 0x53b7a0ff.
//
// Solidity: function subscribedChains(uint32 , address ) view returns(bool)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) SubscribedChains(arg0 uint32, arg1 common.Address) (bool, error) {
	return _LagrangeCommittee.Contract.SubscribedChains(&_LagrangeCommittee.CallOpts, arg0, arg1)
}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x4db6f74a.
//
// Solidity: function updatedEpoch(uint32 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCaller) UpdatedEpoch(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "updatedEpoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x4db6f74a.
//
// Solidity: function updatedEpoch(uint32 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeSession) UpdatedEpoch(arg0 uint32) (*big.Int, error) {
	return _LagrangeCommittee.Contract.UpdatedEpoch(&_LagrangeCommittee.CallOpts, arg0)
}

// UpdatedEpoch is a free data retrieval call binding the contract method 0x4db6f74a.
//
// Solidity: function updatedEpoch(uint32 ) view returns(uint256)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) UpdatedEpoch(arg0 uint32) (*big.Int, error) {
	return _LagrangeCommittee.Contract.UpdatedEpoch(&_LagrangeCommittee.CallOpts, arg0)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCaller) VoteWeigher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LagrangeCommittee.contract.Call(opts, &out, "voteWeigher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeSession) VoteWeigher() (common.Address, error) {
	return _LagrangeCommittee.Contract.VoteWeigher(&_LagrangeCommittee.CallOpts)
}

// VoteWeigher is a free data retrieval call binding the contract method 0xef030673.
//
// Solidity: function voteWeigher() view returns(address)
func (_LagrangeCommittee *LagrangeCommitteeCallerSession) VoteWeigher() (common.Address, error) {
	return _LagrangeCommittee.Contract.VoteWeigher(&_LagrangeCommittee.CallOpts)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0xcfedf427.
//
// Solidity: function addBlsPubKeys(address operator, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) AddBlsPubKeys(opts *bind.TransactOpts, operator common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "addBlsPubKeys", operator, blsKeyWithProof)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0xcfedf427.
//
// Solidity: function addBlsPubKeys(address operator, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) AddBlsPubKeys(operator common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.AddBlsPubKeys(&_LagrangeCommittee.TransactOpts, operator, blsKeyWithProof)
}

// AddBlsPubKeys is a paid mutator transaction binding the contract method 0xcfedf427.
//
// Solidity: function addBlsPubKeys(address operator, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) AddBlsPubKeys(operator common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.AddBlsPubKeys(&_LagrangeCommittee.TransactOpts, operator, blsKeyWithProof)
}

// AddOperator is a paid mutator transaction binding the contract method 0x16c72b2f.
//
// Solidity: function addOperator(address operator, address signAddress, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address, signAddress common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "addOperator", operator, signAddress, blsKeyWithProof)
}

// AddOperator is a paid mutator transaction binding the contract method 0x16c72b2f.
//
// Solidity: function addOperator(address operator, address signAddress, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) AddOperator(operator common.Address, signAddress common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.AddOperator(&_LagrangeCommittee.TransactOpts, operator, signAddress, blsKeyWithProof)
}

// AddOperator is a paid mutator transaction binding the contract method 0x16c72b2f.
//
// Solidity: function addOperator(address operator, address signAddress, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) AddOperator(operator common.Address, signAddress common.Address, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.AddOperator(&_LagrangeCommittee.TransactOpts, operator, signAddress, blsKeyWithProof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Initialize(&_LagrangeCommittee.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Initialize(&_LagrangeCommittee.TransactOpts, initialOwner)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xc8491506.
//
// Solidity: function registerChain(uint32 chainID, uint256 genesisBlock, uint256 epochPeriod, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RegisterChain(opts *bind.TransactOpts, chainID uint32, genesisBlock *big.Int, epochPeriod *big.Int, freezeDuration *big.Int, quorumNumber uint8, minWeight *big.Int, maxWeight *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "registerChain", chainID, genesisBlock, epochPeriod, freezeDuration, quorumNumber, minWeight, maxWeight)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xc8491506.
//
// Solidity: function registerChain(uint32 chainID, uint256 genesisBlock, uint256 epochPeriod, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RegisterChain(chainID uint32, genesisBlock *big.Int, epochPeriod *big.Int, freezeDuration *big.Int, quorumNumber uint8, minWeight *big.Int, maxWeight *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RegisterChain(&_LagrangeCommittee.TransactOpts, chainID, genesisBlock, epochPeriod, freezeDuration, quorumNumber, minWeight, maxWeight)
}

// RegisterChain is a paid mutator transaction binding the contract method 0xc8491506.
//
// Solidity: function registerChain(uint32 chainID, uint256 genesisBlock, uint256 epochPeriod, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RegisterChain(chainID uint32, genesisBlock *big.Int, epochPeriod *big.Int, freezeDuration *big.Int, quorumNumber uint8, minWeight *big.Int, maxWeight *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RegisterChain(&_LagrangeCommittee.TransactOpts, chainID, genesisBlock, epochPeriod, freezeDuration, quorumNumber, minWeight, maxWeight)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x0c57b9d4.
//
// Solidity: function removeBlsPubKeys(address operator, uint32[] indices) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RemoveBlsPubKeys(opts *bind.TransactOpts, operator common.Address, indices []uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "removeBlsPubKeys", operator, indices)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x0c57b9d4.
//
// Solidity: function removeBlsPubKeys(address operator, uint32[] indices) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RemoveBlsPubKeys(operator common.Address, indices []uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RemoveBlsPubKeys(&_LagrangeCommittee.TransactOpts, operator, indices)
}

// RemoveBlsPubKeys is a paid mutator transaction binding the contract method 0x0c57b9d4.
//
// Solidity: function removeBlsPubKeys(address operator, uint32[] indices) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RemoveBlsPubKeys(operator common.Address, indices []uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RemoveBlsPubKeys(&_LagrangeCommittee.TransactOpts, operator, indices)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RemoveOperator(&_LagrangeCommittee.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RemoveOperator(&_LagrangeCommittee.TransactOpts, operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RenounceOwnership() (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RenounceOwnership(&_LagrangeCommittee.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RenounceOwnership(&_LagrangeCommittee.TransactOpts)
}

// RevertEpoch is a paid mutator transaction binding the contract method 0xa6b15068.
//
// Solidity: function revertEpoch(uint32 chainID, uint256 epochNumber) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) RevertEpoch(opts *bind.TransactOpts, chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "revertEpoch", chainID, epochNumber)
}

// RevertEpoch is a paid mutator transaction binding the contract method 0xa6b15068.
//
// Solidity: function revertEpoch(uint32 chainID, uint256 epochNumber) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) RevertEpoch(chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RevertEpoch(&_LagrangeCommittee.TransactOpts, chainID, epochNumber)
}

// RevertEpoch is a paid mutator transaction binding the contract method 0xa6b15068.
//
// Solidity: function revertEpoch(uint32 chainID, uint256 epochNumber) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) RevertEpoch(chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.RevertEpoch(&_LagrangeCommittee.TransactOpts, chainID, epochNumber)
}

// SubscribeChain is a paid mutator transaction binding the contract method 0x6b11c38e.
//
// Solidity: function subscribeChain(address operator, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) SubscribeChain(opts *bind.TransactOpts, operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "subscribeChain", operator, chainID)
}

// SubscribeChain is a paid mutator transaction binding the contract method 0x6b11c38e.
//
// Solidity: function subscribeChain(address operator, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) SubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.SubscribeChain(&_LagrangeCommittee.TransactOpts, operator, chainID)
}

// SubscribeChain is a paid mutator transaction binding the contract method 0x6b11c38e.
//
// Solidity: function subscribeChain(address operator, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) SubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.SubscribeChain(&_LagrangeCommittee.TransactOpts, operator, chainID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.TransferOwnership(&_LagrangeCommittee.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.TransferOwnership(&_LagrangeCommittee.TransactOpts, newOwner)
}

// UnsubscribeByAdmin is a paid mutator transaction binding the contract method 0x935a9b6a.
//
// Solidity: function unsubscribeByAdmin(address[] operators, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) UnsubscribeByAdmin(opts *bind.TransactOpts, operators []common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "unsubscribeByAdmin", operators, chainID)
}

// UnsubscribeByAdmin is a paid mutator transaction binding the contract method 0x935a9b6a.
//
// Solidity: function unsubscribeByAdmin(address[] operators, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) UnsubscribeByAdmin(operators []common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UnsubscribeByAdmin(&_LagrangeCommittee.TransactOpts, operators, chainID)
}

// UnsubscribeByAdmin is a paid mutator transaction binding the contract method 0x935a9b6a.
//
// Solidity: function unsubscribeByAdmin(address[] operators, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) UnsubscribeByAdmin(operators []common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UnsubscribeByAdmin(&_LagrangeCommittee.TransactOpts, operators, chainID)
}

// UnsubscribeChain is a paid mutator transaction binding the contract method 0x0e9f564b.
//
// Solidity: function unsubscribeChain(address operator, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) UnsubscribeChain(opts *bind.TransactOpts, operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "unsubscribeChain", operator, chainID)
}

// UnsubscribeChain is a paid mutator transaction binding the contract method 0x0e9f564b.
//
// Solidity: function unsubscribeChain(address operator, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) UnsubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UnsubscribeChain(&_LagrangeCommittee.TransactOpts, operator, chainID)
}

// UnsubscribeChain is a paid mutator transaction binding the contract method 0x0e9f564b.
//
// Solidity: function unsubscribeChain(address operator, uint32 chainID) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) UnsubscribeChain(operator common.Address, chainID uint32) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UnsubscribeChain(&_LagrangeCommittee.TransactOpts, operator, chainID)
}

// Update is a paid mutator transaction binding the contract method 0x3bc72805.
//
// Solidity: function update(uint32 chainID, uint256 epochNumber) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) Update(opts *bind.TransactOpts, chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "update", chainID, epochNumber)
}

// Update is a paid mutator transaction binding the contract method 0x3bc72805.
//
// Solidity: function update(uint32 chainID, uint256 epochNumber) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) Update(chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Update(&_LagrangeCommittee.TransactOpts, chainID, epochNumber)
}

// Update is a paid mutator transaction binding the contract method 0x3bc72805.
//
// Solidity: function update(uint32 chainID, uint256 epochNumber) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) Update(chainID uint32, epochNumber *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.Update(&_LagrangeCommittee.TransactOpts, chainID, epochNumber)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0x829c7dd7.
//
// Solidity: function updateBlsPubKey(address operator, uint32 index, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) UpdateBlsPubKey(opts *bind.TransactOpts, operator common.Address, index uint32, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "updateBlsPubKey", operator, index, blsKeyWithProof)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0x829c7dd7.
//
// Solidity: function updateBlsPubKey(address operator, uint32 index, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) UpdateBlsPubKey(operator common.Address, index uint32, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UpdateBlsPubKey(&_LagrangeCommittee.TransactOpts, operator, index, blsKeyWithProof)
}

// UpdateBlsPubKey is a paid mutator transaction binding the contract method 0x829c7dd7.
//
// Solidity: function updateBlsPubKey(address operator, uint32 index, (uint256[2][],uint256[2][2],uint256[2],bytes32,uint256) blsKeyWithProof) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) UpdateBlsPubKey(operator common.Address, index uint32, blsKeyWithProof IBLSKeyCheckerBLSKeyWithProof) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UpdateBlsPubKey(&_LagrangeCommittee.TransactOpts, operator, index, blsKeyWithProof)
}

// UpdateChain is a paid mutator transaction binding the contract method 0x13538c7f.
//
// Solidity: function updateChain(uint32 chainID, int256 l1Bias, uint256 genesisBlock, uint256 epochPeriod, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) UpdateChain(opts *bind.TransactOpts, chainID uint32, l1Bias *big.Int, genesisBlock *big.Int, epochPeriod *big.Int, freezeDuration *big.Int, quorumNumber uint8, minWeight *big.Int, maxWeight *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "updateChain", chainID, l1Bias, genesisBlock, epochPeriod, freezeDuration, quorumNumber, minWeight, maxWeight)
}

// UpdateChain is a paid mutator transaction binding the contract method 0x13538c7f.
//
// Solidity: function updateChain(uint32 chainID, int256 l1Bias, uint256 genesisBlock, uint256 epochPeriod, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) UpdateChain(chainID uint32, l1Bias *big.Int, genesisBlock *big.Int, epochPeriod *big.Int, freezeDuration *big.Int, quorumNumber uint8, minWeight *big.Int, maxWeight *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UpdateChain(&_LagrangeCommittee.TransactOpts, chainID, l1Bias, genesisBlock, epochPeriod, freezeDuration, quorumNumber, minWeight, maxWeight)
}

// UpdateChain is a paid mutator transaction binding the contract method 0x13538c7f.
//
// Solidity: function updateChain(uint32 chainID, int256 l1Bias, uint256 genesisBlock, uint256 epochPeriod, uint256 freezeDuration, uint8 quorumNumber, uint96 minWeight, uint96 maxWeight) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) UpdateChain(chainID uint32, l1Bias *big.Int, genesisBlock *big.Int, epochPeriod *big.Int, freezeDuration *big.Int, quorumNumber uint8, minWeight *big.Int, maxWeight *big.Int) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UpdateChain(&_LagrangeCommittee.TransactOpts, chainID, l1Bias, genesisBlock, epochPeriod, freezeDuration, quorumNumber, minWeight, maxWeight)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x2c36c718.
//
// Solidity: function updateSignAddress(address operator, address newSignAddress) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactor) UpdateSignAddress(opts *bind.TransactOpts, operator common.Address, newSignAddress common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.contract.Transact(opts, "updateSignAddress", operator, newSignAddress)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x2c36c718.
//
// Solidity: function updateSignAddress(address operator, address newSignAddress) returns()
func (_LagrangeCommittee *LagrangeCommitteeSession) UpdateSignAddress(operator common.Address, newSignAddress common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UpdateSignAddress(&_LagrangeCommittee.TransactOpts, operator, newSignAddress)
}

// UpdateSignAddress is a paid mutator transaction binding the contract method 0x2c36c718.
//
// Solidity: function updateSignAddress(address operator, address newSignAddress) returns()
func (_LagrangeCommittee *LagrangeCommitteeTransactorSession) UpdateSignAddress(operator common.Address, newSignAddress common.Address) (*types.Transaction, error) {
	return _LagrangeCommittee.Contract.UpdateSignAddress(&_LagrangeCommittee.TransactOpts, operator, newSignAddress)
}

// LagrangeCommitteeBlsKeyUpdatedIterator is returned from FilterBlsKeyUpdated and is used to iterate over the raw logs and unpacked data for BlsKeyUpdated events raised by the LagrangeCommittee contract.
type LagrangeCommitteeBlsKeyUpdatedIterator struct {
	Event *LagrangeCommitteeBlsKeyUpdated // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeBlsKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeBlsKeyUpdated)
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
		it.Event = new(LagrangeCommitteeBlsKeyUpdated)
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
func (it *LagrangeCommitteeBlsKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeBlsKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeBlsKeyUpdated represents a BlsKeyUpdated event raised by the LagrangeCommittee contract.
type LagrangeCommitteeBlsKeyUpdated struct {
	Operator  common.Address
	OrgLength *big.Int
	Added     *big.Int
	Removed   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlsKeyUpdated is a free log retrieval operation binding the contract event 0xded05fe36328e698032818f06b33ec59b066b9f59547700fa2bc365517f7f9a1.
//
// Solidity: event BlsKeyUpdated(address indexed operator, uint256 orgLength, uint256 added, uint256 removed)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterBlsKeyUpdated(opts *bind.FilterOpts, operator []common.Address) (*LagrangeCommitteeBlsKeyUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "BlsKeyUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeBlsKeyUpdatedIterator{contract: _LagrangeCommittee.contract, event: "BlsKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchBlsKeyUpdated is a free log subscription operation binding the contract event 0xded05fe36328e698032818f06b33ec59b066b9f59547700fa2bc365517f7f9a1.
//
// Solidity: event BlsKeyUpdated(address indexed operator, uint256 orgLength, uint256 added, uint256 removed)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchBlsKeyUpdated(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeBlsKeyUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "BlsKeyUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeBlsKeyUpdated)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "BlsKeyUpdated", log); err != nil {
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

// ParseBlsKeyUpdated is a log parse operation binding the contract event 0xded05fe36328e698032818f06b33ec59b066b9f59547700fa2bc365517f7f9a1.
//
// Solidity: event BlsKeyUpdated(address indexed operator, uint256 orgLength, uint256 added, uint256 removed)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseBlsKeyUpdated(log types.Log) (*LagrangeCommitteeBlsKeyUpdated, error) {
	event := new(LagrangeCommitteeBlsKeyUpdated)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "BlsKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeInitCommitteeIterator is returned from FilterInitCommittee and is used to iterate over the raw logs and unpacked data for InitCommittee events raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitCommitteeIterator struct {
	Event *LagrangeCommitteeInitCommittee // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeInitCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeInitCommittee)
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
		it.Event = new(LagrangeCommitteeInitCommittee)
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
func (it *LagrangeCommitteeInitCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeInitCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeInitCommittee represents a InitCommittee event raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitCommittee struct {
	ChainID        *big.Int
	QuorumNumber   uint8
	GenesisBlock   *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	MinWeight      *big.Int
	MaxWeight      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitCommittee is a free log retrieval operation binding the contract event 0x443f7a0c0eb557351b38f3fed112889e529704fb0ccaac63b1f0914816d880f3.
//
// Solidity: event InitCommittee(uint256 indexed chainID, uint8 indexed quorumNumber, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterInitCommittee(opts *bind.FilterOpts, chainID []*big.Int, quorumNumber []uint8) (*LagrangeCommitteeInitCommitteeIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "InitCommittee", chainIDRule, quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeInitCommitteeIterator{contract: _LagrangeCommittee.contract, event: "InitCommittee", logs: logs, sub: sub}, nil
}

// WatchInitCommittee is a free log subscription operation binding the contract event 0x443f7a0c0eb557351b38f3fed112889e529704fb0ccaac63b1f0914816d880f3.
//
// Solidity: event InitCommittee(uint256 indexed chainID, uint8 indexed quorumNumber, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchInitCommittee(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeInitCommittee, chainID []*big.Int, quorumNumber []uint8) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "InitCommittee", chainIDRule, quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeInitCommittee)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
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

// ParseInitCommittee is a log parse operation binding the contract event 0x443f7a0c0eb557351b38f3fed112889e529704fb0ccaac63b1f0914816d880f3.
//
// Solidity: event InitCommittee(uint256 indexed chainID, uint8 indexed quorumNumber, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseInitCommittee(log types.Log) (*LagrangeCommitteeInitCommittee, error) {
	event := new(LagrangeCommitteeInitCommittee)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "InitCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitializedIterator struct {
	Event *LagrangeCommitteeInitialized // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeInitialized)
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
		it.Event = new(LagrangeCommitteeInitialized)
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
func (it *LagrangeCommitteeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeInitialized represents a Initialized event raised by the LagrangeCommittee contract.
type LagrangeCommitteeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterInitialized(opts *bind.FilterOpts) (*LagrangeCommitteeInitializedIterator, error) {

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeInitializedIterator{contract: _LagrangeCommittee.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeInitialized) (event.Subscription, error) {

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeInitialized)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseInitialized(log types.Log) (*LagrangeCommitteeInitialized, error) {
	event := new(LagrangeCommitteeInitialized)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LagrangeCommittee contract.
type LagrangeCommitteeOwnershipTransferredIterator struct {
	Event *LagrangeCommitteeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeOwnershipTransferred)
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
		it.Event = new(LagrangeCommitteeOwnershipTransferred)
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
func (it *LagrangeCommitteeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeOwnershipTransferred represents a OwnershipTransferred event raised by the LagrangeCommittee contract.
type LagrangeCommitteeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LagrangeCommitteeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeOwnershipTransferredIterator{contract: _LagrangeCommittee.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeOwnershipTransferred)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseOwnershipTransferred(log types.Log) (*LagrangeCommitteeOwnershipTransferred, error) {
	event := new(LagrangeCommitteeOwnershipTransferred)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeSignAddressUpdatedIterator is returned from FilterSignAddressUpdated and is used to iterate over the raw logs and unpacked data for SignAddressUpdated events raised by the LagrangeCommittee contract.
type LagrangeCommitteeSignAddressUpdatedIterator struct {
	Event *LagrangeCommitteeSignAddressUpdated // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeSignAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeSignAddressUpdated)
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
		it.Event = new(LagrangeCommitteeSignAddressUpdated)
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
func (it *LagrangeCommitteeSignAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeSignAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeSignAddressUpdated represents a SignAddressUpdated event raised by the LagrangeCommittee contract.
type LagrangeCommitteeSignAddressUpdated struct {
	Operator    common.Address
	SignAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSignAddressUpdated is a free log retrieval operation binding the contract event 0xf29351472086a2d8b1af521f96d65142d14518260a62923a07aa7c9e3c752d84.
//
// Solidity: event SignAddressUpdated(address indexed operator, address indexed signAddress)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterSignAddressUpdated(opts *bind.FilterOpts, operator []common.Address, signAddress []common.Address) (*LagrangeCommitteeSignAddressUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var signAddressRule []interface{}
	for _, signAddressItem := range signAddress {
		signAddressRule = append(signAddressRule, signAddressItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "SignAddressUpdated", operatorRule, signAddressRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeSignAddressUpdatedIterator{contract: _LagrangeCommittee.contract, event: "SignAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchSignAddressUpdated is a free log subscription operation binding the contract event 0xf29351472086a2d8b1af521f96d65142d14518260a62923a07aa7c9e3c752d84.
//
// Solidity: event SignAddressUpdated(address indexed operator, address indexed signAddress)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchSignAddressUpdated(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeSignAddressUpdated, operator []common.Address, signAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var signAddressRule []interface{}
	for _, signAddressItem := range signAddress {
		signAddressRule = append(signAddressRule, signAddressItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "SignAddressUpdated", operatorRule, signAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeSignAddressUpdated)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "SignAddressUpdated", log); err != nil {
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

// ParseSignAddressUpdated is a log parse operation binding the contract event 0xf29351472086a2d8b1af521f96d65142d14518260a62923a07aa7c9e3c752d84.
//
// Solidity: event SignAddressUpdated(address indexed operator, address indexed signAddress)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseSignAddressUpdated(log types.Log) (*LagrangeCommitteeSignAddressUpdated, error) {
	event := new(LagrangeCommitteeSignAddressUpdated)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "SignAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeUpdateCommitteeIterator is returned from FilterUpdateCommittee and is used to iterate over the raw logs and unpacked data for UpdateCommittee events raised by the LagrangeCommittee contract.
type LagrangeCommitteeUpdateCommitteeIterator struct {
	Event *LagrangeCommitteeUpdateCommittee // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeUpdateCommitteeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeUpdateCommittee)
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
		it.Event = new(LagrangeCommitteeUpdateCommittee)
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
func (it *LagrangeCommitteeUpdateCommitteeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeUpdateCommitteeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeUpdateCommittee represents a UpdateCommittee event raised by the LagrangeCommittee contract.
type LagrangeCommitteeUpdateCommittee struct {
	ChainID     *big.Int
	EpochNumber *big.Int
	Current     [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateCommittee is a free log retrieval operation binding the contract event 0x0762b282466177911d330a0c27e7de4e70f0c930671f2080b7cf2dcb3b100f95.
//
// Solidity: event UpdateCommittee(uint256 indexed chainID, uint256 indexed epochNumber, bytes32 current)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterUpdateCommittee(opts *bind.FilterOpts, chainID []*big.Int, epochNumber []*big.Int) (*LagrangeCommitteeUpdateCommitteeIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "UpdateCommittee", chainIDRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeUpdateCommitteeIterator{contract: _LagrangeCommittee.contract, event: "UpdateCommittee", logs: logs, sub: sub}, nil
}

// WatchUpdateCommittee is a free log subscription operation binding the contract event 0x0762b282466177911d330a0c27e7de4e70f0c930671f2080b7cf2dcb3b100f95.
//
// Solidity: event UpdateCommittee(uint256 indexed chainID, uint256 indexed epochNumber, bytes32 current)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchUpdateCommittee(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeUpdateCommittee, chainID []*big.Int, epochNumber []*big.Int) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "UpdateCommittee", chainIDRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeUpdateCommittee)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "UpdateCommittee", log); err != nil {
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

// ParseUpdateCommittee is a log parse operation binding the contract event 0x0762b282466177911d330a0c27e7de4e70f0c930671f2080b7cf2dcb3b100f95.
//
// Solidity: event UpdateCommittee(uint256 indexed chainID, uint256 indexed epochNumber, bytes32 current)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseUpdateCommittee(log types.Log) (*LagrangeCommitteeUpdateCommittee, error) {
	event := new(LagrangeCommitteeUpdateCommittee)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "UpdateCommittee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LagrangeCommitteeUpdateCommitteeParamsIterator is returned from FilterUpdateCommitteeParams and is used to iterate over the raw logs and unpacked data for UpdateCommitteeParams events raised by the LagrangeCommittee contract.
type LagrangeCommitteeUpdateCommitteeParamsIterator struct {
	Event *LagrangeCommitteeUpdateCommitteeParams // Event containing the contract specifics and raw log

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
func (it *LagrangeCommitteeUpdateCommitteeParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LagrangeCommitteeUpdateCommitteeParams)
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
		it.Event = new(LagrangeCommitteeUpdateCommitteeParams)
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
func (it *LagrangeCommitteeUpdateCommitteeParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LagrangeCommitteeUpdateCommitteeParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LagrangeCommitteeUpdateCommitteeParams represents a UpdateCommitteeParams event raised by the LagrangeCommittee contract.
type LagrangeCommitteeUpdateCommitteeParams struct {
	ChainID        *big.Int
	QuorumNumber   uint8
	L1Bias         *big.Int
	GenesisBlock   *big.Int
	Duration       *big.Int
	FreezeDuration *big.Int
	MinWeight      *big.Int
	MaxWeight      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateCommitteeParams is a free log retrieval operation binding the contract event 0xc0f5fe53277733159fe7a739f75b98e959a28a0859db2c46e04a5f4820b96400.
//
// Solidity: event UpdateCommitteeParams(uint256 indexed chainID, uint8 indexed quorumNumber, int256 l1Bias, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) FilterUpdateCommitteeParams(opts *bind.FilterOpts, chainID []*big.Int, quorumNumber []uint8) (*LagrangeCommitteeUpdateCommitteeParamsIterator, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.FilterLogs(opts, "UpdateCommitteeParams", chainIDRule, quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &LagrangeCommitteeUpdateCommitteeParamsIterator{contract: _LagrangeCommittee.contract, event: "UpdateCommitteeParams", logs: logs, sub: sub}, nil
}

// WatchUpdateCommitteeParams is a free log subscription operation binding the contract event 0xc0f5fe53277733159fe7a739f75b98e959a28a0859db2c46e04a5f4820b96400.
//
// Solidity: event UpdateCommitteeParams(uint256 indexed chainID, uint8 indexed quorumNumber, int256 l1Bias, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) WatchUpdateCommitteeParams(opts *bind.WatchOpts, sink chan<- *LagrangeCommitteeUpdateCommitteeParams, chainID []*big.Int, quorumNumber []uint8) (event.Subscription, error) {

	var chainIDRule []interface{}
	for _, chainIDItem := range chainID {
		chainIDRule = append(chainIDRule, chainIDItem)
	}
	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _LagrangeCommittee.contract.WatchLogs(opts, "UpdateCommitteeParams", chainIDRule, quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LagrangeCommitteeUpdateCommitteeParams)
				if err := _LagrangeCommittee.contract.UnpackLog(event, "UpdateCommitteeParams", log); err != nil {
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

// ParseUpdateCommitteeParams is a log parse operation binding the contract event 0xc0f5fe53277733159fe7a739f75b98e959a28a0859db2c46e04a5f4820b96400.
//
// Solidity: event UpdateCommitteeParams(uint256 indexed chainID, uint8 indexed quorumNumber, int256 l1Bias, uint256 genesisBlock, uint256 duration, uint256 freezeDuration, uint96 minWeight, uint96 maxWeight)
func (_LagrangeCommittee *LagrangeCommitteeFilterer) ParseUpdateCommitteeParams(log types.Log) (*LagrangeCommitteeUpdateCommitteeParams, error) {
	event := new(LagrangeCommitteeUpdateCommitteeParams)
	if err := _LagrangeCommittee.contract.UnpackLog(event, "UpdateCommitteeParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
