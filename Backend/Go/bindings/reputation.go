// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ReputationMetaData contains all meta data concerning the Reputation contract.
var ReputationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"delta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newScore\",\"type\":\"uint256\"}],\"name\":\"ReputationChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scoreOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_op\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ReputationABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationMetaData.ABI instead.
var ReputationABI = ReputationMetaData.ABI

// Reputation is an auto generated Go binding around an Ethereum contract.
type Reputation struct {
	ReputationCaller     // Read-only binding to the contract
	ReputationTransactor // Write-only binding to the contract
	ReputationFilterer   // Log filterer for contract events
}

// ReputationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationSession struct {
	Contract     *Reputation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReputationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationCallerSession struct {
	Contract *ReputationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReputationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationTransactorSession struct {
	Contract     *ReputationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReputationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationRaw struct {
	Contract *Reputation // Generic contract binding to access the raw methods on
}

// ReputationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationCallerRaw struct {
	Contract *ReputationCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationTransactorRaw struct {
	Contract *ReputationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputation creates a new instance of Reputation, bound to a specific deployed contract.
func NewReputation(address common.Address, backend bind.ContractBackend) (*Reputation, error) {
	contract, err := bindReputation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reputation{ReputationCaller: ReputationCaller{contract: contract}, ReputationTransactor: ReputationTransactor{contract: contract}, ReputationFilterer: ReputationFilterer{contract: contract}}, nil
}

// NewReputationCaller creates a new read-only instance of Reputation, bound to a specific deployed contract.
func NewReputationCaller(address common.Address, caller bind.ContractCaller) (*ReputationCaller, error) {
	contract, err := bindReputation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationCaller{contract: contract}, nil
}

// NewReputationTransactor creates a new write-only instance of Reputation, bound to a specific deployed contract.
func NewReputationTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationTransactor, error) {
	contract, err := bindReputation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTransactor{contract: contract}, nil
}

// NewReputationFilterer creates a new log filterer instance of Reputation, bound to a specific deployed contract.
func NewReputationFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationFilterer, error) {
	contract, err := bindReputation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationFilterer{contract: contract}, nil
}

// bindReputation binds a generic wrapper to an already deployed contract.
func bindReputation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReputationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.ReputationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transact(opts, method, params...)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Reputation *ReputationCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Reputation *ReputationSession) Operator() (common.Address, error) {
	return _Reputation.Contract.Operator(&_Reputation.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Reputation *ReputationCallerSession) Operator() (common.Address, error) {
	return _Reputation.Contract.Operator(&_Reputation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationSession) Owner() (common.Address, error) {
	return _Reputation.Contract.Owner(&_Reputation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationCallerSession) Owner() (common.Address, error) {
	return _Reputation.Contract.Owner(&_Reputation.CallOpts)
}

// ScoreOf is a free data retrieval call binding the contract method 0x133af456.
//
// Solidity: function scoreOf(address user) view returns(uint256)
func (_Reputation *ReputationCaller) ScoreOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "scoreOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScoreOf is a free data retrieval call binding the contract method 0x133af456.
//
// Solidity: function scoreOf(address user) view returns(uint256)
func (_Reputation *ReputationSession) ScoreOf(user common.Address) (*big.Int, error) {
	return _Reputation.Contract.ScoreOf(&_Reputation.CallOpts, user)
}

// ScoreOf is a free data retrieval call binding the contract method 0x133af456.
//
// Solidity: function scoreOf(address user) view returns(uint256)
func (_Reputation *ReputationCallerSession) ScoreOf(user common.Address) (*big.Int, error) {
	return _Reputation.Contract.ScoreOf(&_Reputation.CallOpts, user)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address mentor, uint256 amount) returns()
func (_Reputation *ReputationTransactor) Add(opts *bind.TransactOpts, mentor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "add", mentor, amount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address mentor, uint256 amount) returns()
func (_Reputation *ReputationSession) Add(mentor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Reputation.Contract.Add(&_Reputation.TransactOpts, mentor, amount)
}

// Add is a paid mutator transaction binding the contract method 0xf5d82b6b.
//
// Solidity: function add(address mentor, uint256 amount) returns()
func (_Reputation *ReputationTransactorSession) Add(mentor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Reputation.Contract.Add(&_Reputation.TransactOpts, mentor, amount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _op) returns()
func (_Reputation *ReputationTransactor) SetOperator(opts *bind.TransactOpts, _op common.Address) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "setOperator", _op)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _op) returns()
func (_Reputation *ReputationSession) SetOperator(_op common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.SetOperator(&_Reputation.TransactOpts, _op)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _op) returns()
func (_Reputation *ReputationTransactorSession) SetOperator(_op common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.SetOperator(&_Reputation.TransactOpts, _op)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address mentor, uint256 amount) returns()
func (_Reputation *ReputationTransactor) Sub(opts *bind.TransactOpts, mentor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "sub", mentor, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address mentor, uint256 amount) returns()
func (_Reputation *ReputationSession) Sub(mentor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Reputation.Contract.Sub(&_Reputation.TransactOpts, mentor, amount)
}

// Sub is a paid mutator transaction binding the contract method 0x26ffee08.
//
// Solidity: function sub(address mentor, uint256 amount) returns()
func (_Reputation *ReputationTransactorSession) Sub(mentor common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Reputation.Contract.Sub(&_Reputation.TransactOpts, mentor, amount)
}

// ReputationOperatorUpdatedIterator is returned from FilterOperatorUpdated and is used to iterate over the raw logs and unpacked data for OperatorUpdated events raised by the Reputation contract.
type ReputationOperatorUpdatedIterator struct {
	Event *ReputationOperatorUpdated // Event containing the contract specifics and raw log

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
func (it *ReputationOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationOperatorUpdated)
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
		it.Event = new(ReputationOperatorUpdated)
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
func (it *ReputationOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationOperatorUpdated represents a OperatorUpdated event raised by the Reputation contract.
type ReputationOperatorUpdated struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorUpdated is a free log retrieval operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address indexed operator)
func (_Reputation *ReputationFilterer) FilterOperatorUpdated(opts *bind.FilterOpts, operator []common.Address) (*ReputationOperatorUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "OperatorUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ReputationOperatorUpdatedIterator{contract: _Reputation.contract, event: "OperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorUpdated is a free log subscription operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address indexed operator)
func (_Reputation *ReputationFilterer) WatchOperatorUpdated(opts *bind.WatchOpts, sink chan<- *ReputationOperatorUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "OperatorUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationOperatorUpdated)
				if err := _Reputation.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
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

// ParseOperatorUpdated is a log parse operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address indexed operator)
func (_Reputation *ReputationFilterer) ParseOperatorUpdated(log types.Log) (*ReputationOperatorUpdated, error) {
	event := new(ReputationOperatorUpdated)
	if err := _Reputation.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationReputationChangedIterator is returned from FilterReputationChanged and is used to iterate over the raw logs and unpacked data for ReputationChanged events raised by the Reputation contract.
type ReputationReputationChangedIterator struct {
	Event *ReputationReputationChanged // Event containing the contract specifics and raw log

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
func (it *ReputationReputationChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationReputationChanged)
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
		it.Event = new(ReputationReputationChanged)
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
func (it *ReputationReputationChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationReputationChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationReputationChanged represents a ReputationChanged event raised by the Reputation contract.
type ReputationReputationChanged struct {
	Mentor   common.Address
	Delta    *big.Int
	NewScore *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReputationChanged is a free log retrieval operation binding the contract event 0xdc3faeadeebcfa827a4c4cc5a430dc9501e614bb5e272ffc223c35f8a48d92ec.
//
// Solidity: event ReputationChanged(address indexed mentor, int256 delta, uint256 newScore)
func (_Reputation *ReputationFilterer) FilterReputationChanged(opts *bind.FilterOpts, mentor []common.Address) (*ReputationReputationChangedIterator, error) {

	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "ReputationChanged", mentorRule)
	if err != nil {
		return nil, err
	}
	return &ReputationReputationChangedIterator{contract: _Reputation.contract, event: "ReputationChanged", logs: logs, sub: sub}, nil
}

// WatchReputationChanged is a free log subscription operation binding the contract event 0xdc3faeadeebcfa827a4c4cc5a430dc9501e614bb5e272ffc223c35f8a48d92ec.
//
// Solidity: event ReputationChanged(address indexed mentor, int256 delta, uint256 newScore)
func (_Reputation *ReputationFilterer) WatchReputationChanged(opts *bind.WatchOpts, sink chan<- *ReputationReputationChanged, mentor []common.Address) (event.Subscription, error) {

	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "ReputationChanged", mentorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationReputationChanged)
				if err := _Reputation.contract.UnpackLog(event, "ReputationChanged", log); err != nil {
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

// ParseReputationChanged is a log parse operation binding the contract event 0xdc3faeadeebcfa827a4c4cc5a430dc9501e614bb5e272ffc223c35f8a48d92ec.
//
// Solidity: event ReputationChanged(address indexed mentor, int256 delta, uint256 newScore)
func (_Reputation *ReputationFilterer) ParseReputationChanged(log types.Log) (*ReputationReputationChanged, error) {
	event := new(ReputationReputationChanged)
	if err := _Reputation.contract.UnpackLog(event, "ReputationChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
