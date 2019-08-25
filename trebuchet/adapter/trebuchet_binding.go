// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC725ABI is the input ABI used to generate the binding from.
const ERC725ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getData\",\"outputs\":[{\"name\":\"_value\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"bytes\"}],\"name\":\"setData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"DataChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"ownerAddress\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractCreated\",\"type\":\"event\"}]"

// ERC725Bin is the compiled bytecode used for deploying new contracts.
const ERC725Bin = `0x`

// DeployERC725 deploys a new Ethereum contract, binding an instance of ERC725 to it.
func DeployERC725(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC725, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC725ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC725Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC725{ERC725Caller: ERC725Caller{contract: contract}, ERC725Transactor: ERC725Transactor{contract: contract}, ERC725Filterer: ERC725Filterer{contract: contract}}, nil
}

// ERC725 is an auto generated Go binding around an Ethereum contract.
type ERC725 struct {
	ERC725Caller     // Read-only binding to the contract
	ERC725Transactor // Write-only binding to the contract
	ERC725Filterer   // Log filterer for contract events
}

// ERC725Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC725Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC725Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC725Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC725Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC725Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC725Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC725Session struct {
	Contract     *ERC725           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC725CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC725CallerSession struct {
	Contract *ERC725Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC725TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC725TransactorSession struct {
	Contract     *ERC725Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC725Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC725Raw struct {
	Contract *ERC725 // Generic contract binding to access the raw methods on
}

// ERC725CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC725CallerRaw struct {
	Contract *ERC725Caller // Generic read-only contract binding to access the raw methods on
}

// ERC725TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC725TransactorRaw struct {
	Contract *ERC725Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC725 creates a new instance of ERC725, bound to a specific deployed contract.
func NewERC725(address common.Address, backend bind.ContractBackend) (*ERC725, error) {
	contract, err := bindERC725(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC725{ERC725Caller: ERC725Caller{contract: contract}, ERC725Transactor: ERC725Transactor{contract: contract}, ERC725Filterer: ERC725Filterer{contract: contract}}, nil
}

// NewERC725Caller creates a new read-only instance of ERC725, bound to a specific deployed contract.
func NewERC725Caller(address common.Address, caller bind.ContractCaller) (*ERC725Caller, error) {
	contract, err := bindERC725(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC725Caller{contract: contract}, nil
}

// NewERC725Transactor creates a new write-only instance of ERC725, bound to a specific deployed contract.
func NewERC725Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC725Transactor, error) {
	contract, err := bindERC725(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC725Transactor{contract: contract}, nil
}

// NewERC725Filterer creates a new log filterer instance of ERC725, bound to a specific deployed contract.
func NewERC725Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC725Filterer, error) {
	contract, err := bindERC725(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC725Filterer{contract: contract}, nil
}

// bindERC725 binds a generic wrapper to an already deployed contract.
func bindERC725(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC725ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC725 *ERC725Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC725.Contract.ERC725Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC725 *ERC725Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC725.Contract.ERC725Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC725 *ERC725Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC725.Contract.ERC725Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC725 *ERC725CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC725.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC725 *ERC725TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC725.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC725 *ERC725TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC725.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes _value)
func (_ERC725 *ERC725Caller) GetData(opts *bind.CallOpts, _key [32]byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ERC725.contract.Call(opts, out, "getData", _key)
	return *ret0, err
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes _value)
func (_ERC725 *ERC725Session) GetData(_key [32]byte) ([]byte, error) {
	return _ERC725.Contract.GetData(&_ERC725.CallOpts, _key)
}

// GetData is a free data retrieval call binding the contract method 0x54f6127f.
//
// Solidity: function getData(bytes32 _key) constant returns(bytes _value)
func (_ERC725 *ERC725CallerSession) GetData(_key [32]byte) ([]byte, error) {
	return _ERC725.Contract.GetData(&_ERC725.CallOpts, _key)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_ERC725 *ERC725Transactor) ChangeOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ERC725.contract.Transact(opts, "changeOwner", _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_ERC725 *ERC725Session) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _ERC725.Contract.ChangeOwner(&_ERC725.TransactOpts, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address _owner) returns()
func (_ERC725 *ERC725TransactorSession) ChangeOwner(_owner common.Address) (*types.Transaction, error) {
	return _ERC725.Contract.ChangeOwner(&_ERC725.TransactOpts, _owner)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _data) returns()
func (_ERC725 *ERC725Transactor) Execute(opts *bind.TransactOpts, _data []byte) (*types.Transaction, error) {
	return _ERC725.contract.Transact(opts, "execute", _data)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _data) returns()
func (_ERC725 *ERC725Session) Execute(_data []byte) (*types.Transaction, error) {
	return _ERC725.Contract.Execute(&_ERC725.TransactOpts, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _data) returns()
func (_ERC725 *ERC725TransactorSession) Execute(_data []byte) (*types.Transaction, error) {
	return _ERC725.Contract.Execute(&_ERC725.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 _key, bytes _value) returns()
func (_ERC725 *ERC725Transactor) SetData(opts *bind.TransactOpts, _key [32]byte, _value []byte) (*types.Transaction, error) {
	return _ERC725.contract.Transact(opts, "setData", _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 _key, bytes _value) returns()
func (_ERC725 *ERC725Session) SetData(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _ERC725.Contract.SetData(&_ERC725.TransactOpts, _key, _value)
}

// SetData is a paid mutator transaction binding the contract method 0x7f23690c.
//
// Solidity: function setData(bytes32 _key, bytes _value) returns()
func (_ERC725 *ERC725TransactorSession) SetData(_key [32]byte, _value []byte) (*types.Transaction, error) {
	return _ERC725.Contract.SetData(&_ERC725.TransactOpts, _key, _value)
}

// ERC725ContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the ERC725 contract.
type ERC725ContractCreatedIterator struct {
	Event *ERC725ContractCreated // Event containing the contract specifics and raw log

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
func (it *ERC725ContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC725ContractCreated)
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
		it.Event = new(ERC725ContractCreated)
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
func (it *ERC725ContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC725ContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC725ContractCreated represents a ContractCreated event raised by the ERC725 contract.
type ERC725ContractCreated struct {
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_ERC725 *ERC725Filterer) FilterContractCreated(opts *bind.FilterOpts, contractAddress []common.Address) (*ERC725ContractCreatedIterator, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ERC725.contract.FilterLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC725ContractCreatedIterator{contract: _ERC725.contract, event: "ContractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xcf78cf0d6f3d8371e1075c69c492ab4ec5d8cf23a1a239b6a51a1d00be7ca312.
//
// Solidity: event ContractCreated(address indexed contractAddress)
func (_ERC725 *ERC725Filterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *ERC725ContractCreated, contractAddress []common.Address) (event.Subscription, error) {

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ERC725.contract.WatchLogs(opts, "ContractCreated", contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC725ContractCreated)
				if err := _ERC725.contract.UnpackLog(event, "ContractCreated", log); err != nil {
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

// ERC725DataChangedIterator is returned from FilterDataChanged and is used to iterate over the raw logs and unpacked data for DataChanged events raised by the ERC725 contract.
type ERC725DataChangedIterator struct {
	Event *ERC725DataChanged // Event containing the contract specifics and raw log

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
func (it *ERC725DataChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC725DataChanged)
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
		it.Event = new(ERC725DataChanged)
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
func (it *ERC725DataChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC725DataChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC725DataChanged represents a DataChanged event raised by the ERC725 contract.
type ERC725DataChanged struct {
	Key   [32]byte
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDataChanged is a free log retrieval operation binding the contract event 0xece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes value)
func (_ERC725 *ERC725Filterer) FilterDataChanged(opts *bind.FilterOpts, key [][32]byte) (*ERC725DataChangedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC725.contract.FilterLogs(opts, "DataChanged", keyRule)
	if err != nil {
		return nil, err
	}
	return &ERC725DataChangedIterator{contract: _ERC725.contract, event: "DataChanged", logs: logs, sub: sub}, nil
}

// WatchDataChanged is a free log subscription operation binding the contract event 0xece574603820d07bc9b91f2a932baadf4628aabcb8afba49776529c14a6104b2.
//
// Solidity: event DataChanged(bytes32 indexed key, bytes value)
func (_ERC725 *ERC725Filterer) WatchDataChanged(opts *bind.WatchOpts, sink chan<- *ERC725DataChanged, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _ERC725.contract.WatchLogs(opts, "DataChanged", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC725DataChanged)
				if err := _ERC725.contract.UnpackLog(event, "DataChanged", log); err != nil {
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

// ERC725OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the ERC725 contract.
type ERC725OwnerChangedIterator struct {
	Event *ERC725OwnerChanged // Event containing the contract specifics and raw log

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
func (it *ERC725OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC725OwnerChanged)
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
		it.Event = new(ERC725OwnerChanged)
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
func (it *ERC725OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC725OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC725OwnerChanged represents a OwnerChanged event raised by the ERC725 contract.
type ERC725OwnerChanged struct {
	OwnerAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed ownerAddress)
func (_ERC725 *ERC725Filterer) FilterOwnerChanged(opts *bind.FilterOpts, ownerAddress []common.Address) (*ERC725OwnerChangedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ERC725.contract.FilterLogs(opts, "OwnerChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC725OwnerChangedIterator{contract: _ERC725.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address indexed ownerAddress)
func (_ERC725 *ERC725Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ERC725OwnerChanged, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _ERC725.contract.WatchLogs(opts, "OwnerChanged", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC725OwnerChanged)
				if err := _ERC725.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// TrebuchetABI is the input ABI used to generate the binding from.
const TrebuchetABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// TrebuchetBin is the compiled bytecode used for deploying new contracts.
const TrebuchetBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556102d0806100326000396000f3fe6080604052600436106100295760003560e01c8063444714151461002e5780638da5cb5b14610043575b600080fd5b61004161003c3660046101a0565b61006e565b005b34801561004f57600080fd5b506100586100d6565b60405161006591906101ec565b60405180910390f35b6000546001600160a01b0316331461008557600080fd5b60005b81518110156100d257606082828151811061009f57fe5b602002602001015190506000806017835103603784016000602386015160601c602087015160e81cf15050600101610088565b5050565b6000546001600160a01b031681565b600082601f8301126100f657600080fd5b813561010961010482610227565b610200565b81815260209384019390925082018360005b8381101561014757813586016101318882610151565b845250602092830192919091019060010161011b565b5050505092915050565b600082601f83011261016257600080fd5b813561017061010482610248565b9150808252602083016020830185838301111561018c57600080fd5b610197838284610281565b50505092915050565b6000602082840312156101b257600080fd5b813567ffffffffffffffff8111156101c957600080fd5b6101d5848285016100e5565b949350505050565b6101e681610270565b82525050565b602081016101fa82846101dd565b92915050565b60405181810167ffffffffffffffff8111828210171561021f57600080fd5b604052919050565b600067ffffffffffffffff82111561023e57600080fd5b5060209081020190565b600067ffffffffffffffff82111561025f57600080fd5b506020601f91909101601f19160190565b60006001600160a01b0382166101fa565b8281833750600091015256fea365627a7a72305820f879f0ade5b4e08956a2727bb555a03951abed0d159775ad98425e167604528f6c6578706572696d656e74616cf564736f6c63430005090040`

// DeployTrebuchet deploys a new Ethereum contract, binding an instance of Trebuchet to it.
func DeployTrebuchet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trebuchet, error) {
	parsed, err := abi.JSON(strings.NewReader(TrebuchetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TrebuchetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trebuchet{TrebuchetCaller: TrebuchetCaller{contract: contract}, TrebuchetTransactor: TrebuchetTransactor{contract: contract}, TrebuchetFilterer: TrebuchetFilterer{contract: contract}}, nil
}

// Trebuchet is an auto generated Go binding around an Ethereum contract.
type Trebuchet struct {
	TrebuchetCaller     // Read-only binding to the contract
	TrebuchetTransactor // Write-only binding to the contract
	TrebuchetFilterer   // Log filterer for contract events
}

// TrebuchetCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrebuchetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrebuchetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrebuchetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrebuchetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrebuchetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrebuchetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrebuchetSession struct {
	Contract     *Trebuchet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrebuchetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrebuchetCallerSession struct {
	Contract *TrebuchetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TrebuchetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrebuchetTransactorSession struct {
	Contract     *TrebuchetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TrebuchetRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrebuchetRaw struct {
	Contract *Trebuchet // Generic contract binding to access the raw methods on
}

// TrebuchetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrebuchetCallerRaw struct {
	Contract *TrebuchetCaller // Generic read-only contract binding to access the raw methods on
}

// TrebuchetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrebuchetTransactorRaw struct {
	Contract *TrebuchetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrebuchet creates a new instance of Trebuchet, bound to a specific deployed contract.
func NewTrebuchet(address common.Address, backend bind.ContractBackend) (*Trebuchet, error) {
	contract, err := bindTrebuchet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trebuchet{TrebuchetCaller: TrebuchetCaller{contract: contract}, TrebuchetTransactor: TrebuchetTransactor{contract: contract}, TrebuchetFilterer: TrebuchetFilterer{contract: contract}}, nil
}

// NewTrebuchetCaller creates a new read-only instance of Trebuchet, bound to a specific deployed contract.
func NewTrebuchetCaller(address common.Address, caller bind.ContractCaller) (*TrebuchetCaller, error) {
	contract, err := bindTrebuchet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrebuchetCaller{contract: contract}, nil
}

// NewTrebuchetTransactor creates a new write-only instance of Trebuchet, bound to a specific deployed contract.
func NewTrebuchetTransactor(address common.Address, transactor bind.ContractTransactor) (*TrebuchetTransactor, error) {
	contract, err := bindTrebuchet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrebuchetTransactor{contract: contract}, nil
}

// NewTrebuchetFilterer creates a new log filterer instance of Trebuchet, bound to a specific deployed contract.
func NewTrebuchetFilterer(address common.Address, filterer bind.ContractFilterer) (*TrebuchetFilterer, error) {
	contract, err := bindTrebuchet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrebuchetFilterer{contract: contract}, nil
}

// bindTrebuchet binds a generic wrapper to an already deployed contract.
func bindTrebuchet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrebuchetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trebuchet *TrebuchetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trebuchet.Contract.TrebuchetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trebuchet *TrebuchetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trebuchet.Contract.TrebuchetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trebuchet *TrebuchetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trebuchet.Contract.TrebuchetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trebuchet *TrebuchetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trebuchet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trebuchet *TrebuchetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trebuchet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trebuchet *TrebuchetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trebuchet.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trebuchet *TrebuchetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trebuchet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trebuchet *TrebuchetSession) Owner() (common.Address, error) {
	return _Trebuchet.Contract.Owner(&_Trebuchet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trebuchet *TrebuchetCallerSession) Owner() (common.Address, error) {
	return _Trebuchet.Contract.Owner(&_Trebuchet.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x44471415.
//
// Solidity: function execute(bytes[] _data) returns()
func (_Trebuchet *TrebuchetTransactor) Execute(opts *bind.TransactOpts, _data [][]byte) (*types.Transaction, error) {
	return _Trebuchet.contract.Transact(opts, "execute", _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44471415.
//
// Solidity: function execute(bytes[] _data) returns()
func (_Trebuchet *TrebuchetSession) Execute(_data [][]byte) (*types.Transaction, error) {
	return _Trebuchet.Contract.Execute(&_Trebuchet.TransactOpts, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x44471415.
//
// Solidity: function execute(bytes[] _data) returns()
func (_Trebuchet *TrebuchetTransactorSession) Execute(_data [][]byte) (*types.Transaction, error) {
	return _Trebuchet.Contract.Execute(&_Trebuchet.TransactOpts, _data)
}
