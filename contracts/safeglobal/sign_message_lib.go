// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safeglobal

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

// SignMessageLibMetaData contains all meta data concerning the SignMessageLib contract.
var SignMessageLibMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"signMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506105f28061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80630a1028c41461003857806385a5affe14610068575b5f80fd5b610052600480360381019061004d9190610368565b610084565b60405161005f91906103c7565b60405180910390f35b610082600480360381019061007d919061043d565b610183565b005b5f807f60b3cbf8b4a223d68d641b3b6ddf9a298e7f33710cf3d3a9d1146b5a6150fbca5f1b83805190602001206040516020016100c2929190610488565b604051602081830303815290604052805190602001209050601960f81b600160f81b3073ffffffffffffffffffffffffffffffffffffffff1663f698da256040518163ffffffff1660e01b8152600401602060405180830381865afa15801561012d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061015191906104d9565b83604051602001610165949392919061056f565b60405160208183030381529060405280519060200120915050919050565b5f6101d083838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050610084565b9050600160075f8381526020019081526020015f2081905550807fe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e460405160405180910390a2505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61027a82610234565b810181811067ffffffffffffffff8211171561029957610298610244565b5b80604052505050565b5f6102ab61021b565b90506102b78282610271565b919050565b5f67ffffffffffffffff8211156102d6576102d5610244565b5b6102df82610234565b9050602081019050919050565b828183375f83830152505050565b5f61030c610307846102bc565b6102a2565b90508281526020810184848401111561032857610327610230565b5b6103338482856102ec565b509392505050565b5f82601f83011261034f5761034e61022c565b5b813561035f8482602086016102fa565b91505092915050565b5f6020828403121561037d5761037c610224565b5b5f82013567ffffffffffffffff81111561039a57610399610228565b5b6103a68482850161033b565b91505092915050565b5f819050919050565b6103c1816103af565b82525050565b5f6020820190506103da5f8301846103b8565b92915050565b5f80fd5b5f80fd5b5f8083601f8401126103fd576103fc61022c565b5b8235905067ffffffffffffffff81111561041a576104196103e0565b5b602083019150836001820283011115610436576104356103e4565b5b9250929050565b5f806020838503121561045357610452610224565b5b5f83013567ffffffffffffffff8111156104705761046f610228565b5b61047c858286016103e8565b92509250509250929050565b5f60408201905061049b5f8301856103b8565b6104a860208301846103b8565b9392505050565b6104b8816103af565b81146104c2575f80fd5b50565b5f815190506104d3816104af565b92915050565b5f602082840312156104ee576104ed610224565b5b5f6104fb848285016104c5565b91505092915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b5f819050919050565b61054961054482610504565b61052f565b82525050565b5f819050919050565b610569610564826103af565b61054f565b82525050565b5f61057a8287610538565b60018201915061058a8286610538565b60018201915061059a8285610558565b6020820191506105aa8284610558565b6020820191508190509594505050505056fea26469706673582212201e48fe83f8a9b3ed0a91b0fc6905eedc1b2bb6c7eb8059758b77b7eee826369d64736f6c63430008190033",
}

// SignMessageLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SignMessageLibMetaData.ABI instead.
var SignMessageLibABI = SignMessageLibMetaData.ABI

// SignMessageLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignMessageLibMetaData.Bin instead.
var SignMessageLibBin = SignMessageLibMetaData.Bin

// DeploySignMessageLib deploys a new Ethereum contract, binding an instance of SignMessageLib to it.
func DeploySignMessageLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignMessageLib, error) {
	parsed, err := SignMessageLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignMessageLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignMessageLib{SignMessageLibCaller: SignMessageLibCaller{contract: contract}, SignMessageLibTransactor: SignMessageLibTransactor{contract: contract}, SignMessageLibFilterer: SignMessageLibFilterer{contract: contract}}, nil
}

// SignMessageLib is an auto generated Go binding around an Ethereum contract.
type SignMessageLib struct {
	SignMessageLibCaller     // Read-only binding to the contract
	SignMessageLibTransactor // Write-only binding to the contract
	SignMessageLibFilterer   // Log filterer for contract events
}

// SignMessageLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignMessageLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignMessageLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignMessageLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignMessageLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignMessageLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignMessageLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignMessageLibSession struct {
	Contract     *SignMessageLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignMessageLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignMessageLibCallerSession struct {
	Contract *SignMessageLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SignMessageLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignMessageLibTransactorSession struct {
	Contract     *SignMessageLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SignMessageLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignMessageLibRaw struct {
	Contract *SignMessageLib // Generic contract binding to access the raw methods on
}

// SignMessageLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignMessageLibCallerRaw struct {
	Contract *SignMessageLibCaller // Generic read-only contract binding to access the raw methods on
}

// SignMessageLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignMessageLibTransactorRaw struct {
	Contract *SignMessageLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignMessageLib creates a new instance of SignMessageLib, bound to a specific deployed contract.
func NewSignMessageLib(address common.Address, backend bind.ContractBackend) (*SignMessageLib, error) {
	contract, err := bindSignMessageLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignMessageLib{SignMessageLibCaller: SignMessageLibCaller{contract: contract}, SignMessageLibTransactor: SignMessageLibTransactor{contract: contract}, SignMessageLibFilterer: SignMessageLibFilterer{contract: contract}}, nil
}

// NewSignMessageLibCaller creates a new read-only instance of SignMessageLib, bound to a specific deployed contract.
func NewSignMessageLibCaller(address common.Address, caller bind.ContractCaller) (*SignMessageLibCaller, error) {
	contract, err := bindSignMessageLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignMessageLibCaller{contract: contract}, nil
}

// NewSignMessageLibTransactor creates a new write-only instance of SignMessageLib, bound to a specific deployed contract.
func NewSignMessageLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SignMessageLibTransactor, error) {
	contract, err := bindSignMessageLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignMessageLibTransactor{contract: contract}, nil
}

// NewSignMessageLibFilterer creates a new log filterer instance of SignMessageLib, bound to a specific deployed contract.
func NewSignMessageLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SignMessageLibFilterer, error) {
	contract, err := bindSignMessageLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignMessageLibFilterer{contract: contract}, nil
}

// bindSignMessageLib binds a generic wrapper to an already deployed contract.
func bindSignMessageLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignMessageLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignMessageLib *SignMessageLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignMessageLib.Contract.SignMessageLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignMessageLib *SignMessageLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignMessageLib.Contract.SignMessageLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignMessageLib *SignMessageLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignMessageLib.Contract.SignMessageLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignMessageLib *SignMessageLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignMessageLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignMessageLib *SignMessageLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignMessageLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignMessageLib *SignMessageLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignMessageLib.Contract.contract.Transact(opts, method, params...)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_SignMessageLib *SignMessageLibCaller) GetMessageHash(opts *bind.CallOpts, message []byte) ([32]byte, error) {
	var out []interface{}
	err := _SignMessageLib.contract.Call(opts, &out, "getMessageHash", message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_SignMessageLib *SignMessageLibSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _SignMessageLib.Contract.GetMessageHash(&_SignMessageLib.CallOpts, message)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x0a1028c4.
//
// Solidity: function getMessageHash(bytes message) view returns(bytes32)
func (_SignMessageLib *SignMessageLibCallerSession) GetMessageHash(message []byte) ([32]byte, error) {
	return _SignMessageLib.Contract.GetMessageHash(&_SignMessageLib.CallOpts, message)
}

// SignMessage is a paid mutator transaction binding the contract method 0x85a5affe.
//
// Solidity: function signMessage(bytes _data) returns()
func (_SignMessageLib *SignMessageLibTransactor) SignMessage(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _SignMessageLib.contract.Transact(opts, "signMessage", _data)
}

// SignMessage is a paid mutator transaction binding the contract method 0x85a5affe.
//
// Solidity: function signMessage(bytes _data) returns()
func (_SignMessageLib *SignMessageLibSession) SignMessage(_data []byte) (*types.Transaction, error) {
	return _SignMessageLib.Contract.SignMessage(&_SignMessageLib.TransactOpts, _data)
}

// SignMessage is a paid mutator transaction binding the contract method 0x85a5affe.
//
// Solidity: function signMessage(bytes _data) returns()
func (_SignMessageLib *SignMessageLibTransactorSession) SignMessage(_data []byte) (*types.Transaction, error) {
	return _SignMessageLib.Contract.SignMessage(&_SignMessageLib.TransactOpts, _data)
}

// SignMessageLibSignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the SignMessageLib contract.
type SignMessageLibSignMsgIterator struct {
	Event *SignMessageLibSignMsg // Event containing the contract specifics and raw log

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
func (it *SignMessageLibSignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignMessageLibSignMsg)
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
		it.Event = new(SignMessageLibSignMsg)
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
func (it *SignMessageLibSignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignMessageLibSignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignMessageLibSignMsg represents a SignMsg event raised by the SignMessageLib contract.
type SignMessageLibSignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SignMessageLib *SignMessageLibFilterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*SignMessageLibSignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _SignMessageLib.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &SignMessageLibSignMsgIterator{contract: _SignMessageLib.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SignMessageLib *SignMessageLibFilterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *SignMessageLibSignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _SignMessageLib.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignMessageLibSignMsg)
				if err := _SignMessageLib.contract.UnpackLog(event, "SignMsg", log); err != nil {
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

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SignMessageLib *SignMessageLibFilterer) ParseSignMsg(log types.Log) (*SignMessageLibSignMsg, error) {
	event := new(SignMessageLibSignMsg)
	if err := _SignMessageLib.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
