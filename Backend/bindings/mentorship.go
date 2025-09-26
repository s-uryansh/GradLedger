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

// MentorshipMetaData contains all meta data concerning the Mentorship contract.
var MentorshipMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reputation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"}],\"name\":\"Accepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"}],\"name\":\"Completed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"student\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"upvote\",\"type\":\"bool\"}],\"name\":\"Feedback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"student\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"}],\"name\":\"Requested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"accept\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"complete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getSession\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"student\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"},{\"internalType\":\"enumMentorship.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"feedbackGiven\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"upvote\",\"type\":\"bool\"}],\"name\":\"giveFeedback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reputation\",\"outputs\":[{\"internalType\":\"contractIReputation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"}],\"name\":\"request\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"student\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mentor\",\"type\":\"address\"},{\"internalType\":\"enumMentorship.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"feedbackGiven\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSessions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userRegistry\",\"outputs\":[{\"internalType\":\"contractIUserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MentorshipABI is the input ABI used to generate the binding from.
// Deprecated: Use MentorshipMetaData.ABI instead.
var MentorshipABI = MentorshipMetaData.ABI

// Mentorship is an auto generated Go binding around an Ethereum contract.
type Mentorship struct {
	MentorshipCaller     // Read-only binding to the contract
	MentorshipTransactor // Write-only binding to the contract
	MentorshipFilterer   // Log filterer for contract events
}

// MentorshipCaller is an auto generated read-only Go binding around an Ethereum contract.
type MentorshipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MentorshipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MentorshipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MentorshipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MentorshipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MentorshipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MentorshipSession struct {
	Contract     *Mentorship       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MentorshipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MentorshipCallerSession struct {
	Contract *MentorshipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MentorshipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MentorshipTransactorSession struct {
	Contract     *MentorshipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MentorshipRaw is an auto generated low-level Go binding around an Ethereum contract.
type MentorshipRaw struct {
	Contract *Mentorship // Generic contract binding to access the raw methods on
}

// MentorshipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MentorshipCallerRaw struct {
	Contract *MentorshipCaller // Generic read-only contract binding to access the raw methods on
}

// MentorshipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MentorshipTransactorRaw struct {
	Contract *MentorshipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMentorship creates a new instance of Mentorship, bound to a specific deployed contract.
func NewMentorship(address common.Address, backend bind.ContractBackend) (*Mentorship, error) {
	contract, err := bindMentorship(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mentorship{MentorshipCaller: MentorshipCaller{contract: contract}, MentorshipTransactor: MentorshipTransactor{contract: contract}, MentorshipFilterer: MentorshipFilterer{contract: contract}}, nil
}

// NewMentorshipCaller creates a new read-only instance of Mentorship, bound to a specific deployed contract.
func NewMentorshipCaller(address common.Address, caller bind.ContractCaller) (*MentorshipCaller, error) {
	contract, err := bindMentorship(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MentorshipCaller{contract: contract}, nil
}

// NewMentorshipTransactor creates a new write-only instance of Mentorship, bound to a specific deployed contract.
func NewMentorshipTransactor(address common.Address, transactor bind.ContractTransactor) (*MentorshipTransactor, error) {
	contract, err := bindMentorship(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MentorshipTransactor{contract: contract}, nil
}

// NewMentorshipFilterer creates a new log filterer instance of Mentorship, bound to a specific deployed contract.
func NewMentorshipFilterer(address common.Address, filterer bind.ContractFilterer) (*MentorshipFilterer, error) {
	contract, err := bindMentorship(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MentorshipFilterer{contract: contract}, nil
}

// bindMentorship binds a generic wrapper to an already deployed contract.
func bindMentorship(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MentorshipMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mentorship *MentorshipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mentorship.Contract.MentorshipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mentorship *MentorshipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mentorship.Contract.MentorshipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mentorship *MentorshipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mentorship.Contract.MentorshipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mentorship *MentorshipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mentorship.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mentorship *MentorshipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mentorship.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mentorship *MentorshipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mentorship.Contract.contract.Transact(opts, method, params...)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 id) view returns(address student, address mentor, uint8 status, bool feedbackGiven)
func (_Mentorship *MentorshipCaller) GetSession(opts *bind.CallOpts, id *big.Int) (struct {
	Student       common.Address
	Mentor        common.Address
	Status        uint8
	FeedbackGiven bool
}, error) {
	var out []interface{}
	err := _Mentorship.contract.Call(opts, &out, "getSession", id)

	outstruct := new(struct {
		Student       common.Address
		Mentor        common.Address
		Status        uint8
		FeedbackGiven bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Student = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Mentor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.FeedbackGiven = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 id) view returns(address student, address mentor, uint8 status, bool feedbackGiven)
func (_Mentorship *MentorshipSession) GetSession(id *big.Int) (struct {
	Student       common.Address
	Mentor        common.Address
	Status        uint8
	FeedbackGiven bool
}, error) {
	return _Mentorship.Contract.GetSession(&_Mentorship.CallOpts, id)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 id) view returns(address student, address mentor, uint8 status, bool feedbackGiven)
func (_Mentorship *MentorshipCallerSession) GetSession(id *big.Int) (struct {
	Student       common.Address
	Mentor        common.Address
	Status        uint8
	FeedbackGiven bool
}, error) {
	return _Mentorship.Contract.GetSession(&_Mentorship.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mentorship *MentorshipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mentorship.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mentorship *MentorshipSession) Owner() (common.Address, error) {
	return _Mentorship.Contract.Owner(&_Mentorship.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mentorship *MentorshipCallerSession) Owner() (common.Address, error) {
	return _Mentorship.Contract.Owner(&_Mentorship.CallOpts)
}

// Reputation is a free data retrieval call binding the contract method 0xc52164c6.
//
// Solidity: function reputation() view returns(address)
func (_Mentorship *MentorshipCaller) Reputation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mentorship.contract.Call(opts, &out, "reputation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Reputation is a free data retrieval call binding the contract method 0xc52164c6.
//
// Solidity: function reputation() view returns(address)
func (_Mentorship *MentorshipSession) Reputation() (common.Address, error) {
	return _Mentorship.Contract.Reputation(&_Mentorship.CallOpts)
}

// Reputation is a free data retrieval call binding the contract method 0xc52164c6.
//
// Solidity: function reputation() view returns(address)
func (_Mentorship *MentorshipCallerSession) Reputation() (common.Address, error) {
	return _Mentorship.Contract.Reputation(&_Mentorship.CallOpts)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(address student, address mentor, uint8 status, bool feedbackGiven)
func (_Mentorship *MentorshipCaller) Sessions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Student       common.Address
	Mentor        common.Address
	Status        uint8
	FeedbackGiven bool
}, error) {
	var out []interface{}
	err := _Mentorship.contract.Call(opts, &out, "sessions", arg0)

	outstruct := new(struct {
		Student       common.Address
		Mentor        common.Address
		Status        uint8
		FeedbackGiven bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Student = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Mentor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.FeedbackGiven = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(address student, address mentor, uint8 status, bool feedbackGiven)
func (_Mentorship *MentorshipSession) Sessions(arg0 *big.Int) (struct {
	Student       common.Address
	Mentor        common.Address
	Status        uint8
	FeedbackGiven bool
}, error) {
	return _Mentorship.Contract.Sessions(&_Mentorship.CallOpts, arg0)
}

// Sessions is a free data retrieval call binding the contract method 0x83c4b7a3.
//
// Solidity: function sessions(uint256 ) view returns(address student, address mentor, uint8 status, bool feedbackGiven)
func (_Mentorship *MentorshipCallerSession) Sessions(arg0 *big.Int) (struct {
	Student       common.Address
	Mentor        common.Address
	Status        uint8
	FeedbackGiven bool
}, error) {
	return _Mentorship.Contract.Sessions(&_Mentorship.CallOpts, arg0)
}

// TotalSessions is a free data retrieval call binding the contract method 0x6f5db0b5.
//
// Solidity: function totalSessions() view returns(uint256)
func (_Mentorship *MentorshipCaller) TotalSessions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mentorship.contract.Call(opts, &out, "totalSessions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSessions is a free data retrieval call binding the contract method 0x6f5db0b5.
//
// Solidity: function totalSessions() view returns(uint256)
func (_Mentorship *MentorshipSession) TotalSessions() (*big.Int, error) {
	return _Mentorship.Contract.TotalSessions(&_Mentorship.CallOpts)
}

// TotalSessions is a free data retrieval call binding the contract method 0x6f5db0b5.
//
// Solidity: function totalSessions() view returns(uint256)
func (_Mentorship *MentorshipCallerSession) TotalSessions() (*big.Int, error) {
	return _Mentorship.Contract.TotalSessions(&_Mentorship.CallOpts)
}

// UserRegistry is a free data retrieval call binding the contract method 0x5c7460d6.
//
// Solidity: function userRegistry() view returns(address)
func (_Mentorship *MentorshipCaller) UserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mentorship.contract.Call(opts, &out, "userRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserRegistry is a free data retrieval call binding the contract method 0x5c7460d6.
//
// Solidity: function userRegistry() view returns(address)
func (_Mentorship *MentorshipSession) UserRegistry() (common.Address, error) {
	return _Mentorship.Contract.UserRegistry(&_Mentorship.CallOpts)
}

// UserRegistry is a free data retrieval call binding the contract method 0x5c7460d6.
//
// Solidity: function userRegistry() view returns(address)
func (_Mentorship *MentorshipCallerSession) UserRegistry() (common.Address, error) {
	return _Mentorship.Contract.UserRegistry(&_Mentorship.CallOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 id) returns()
func (_Mentorship *MentorshipTransactor) Accept(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mentorship.contract.Transact(opts, "accept", id)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 id) returns()
func (_Mentorship *MentorshipSession) Accept(id *big.Int) (*types.Transaction, error) {
	return _Mentorship.Contract.Accept(&_Mentorship.TransactOpts, id)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 id) returns()
func (_Mentorship *MentorshipTransactorSession) Accept(id *big.Int) (*types.Transaction, error) {
	return _Mentorship.Contract.Accept(&_Mentorship.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 id) returns()
func (_Mentorship *MentorshipTransactor) Cancel(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mentorship.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 id) returns()
func (_Mentorship *MentorshipSession) Cancel(id *big.Int) (*types.Transaction, error) {
	return _Mentorship.Contract.Cancel(&_Mentorship.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 id) returns()
func (_Mentorship *MentorshipTransactorSession) Cancel(id *big.Int) (*types.Transaction, error) {
	return _Mentorship.Contract.Cancel(&_Mentorship.TransactOpts, id)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 id) returns()
func (_Mentorship *MentorshipTransactor) Complete(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Mentorship.contract.Transact(opts, "complete", id)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 id) returns()
func (_Mentorship *MentorshipSession) Complete(id *big.Int) (*types.Transaction, error) {
	return _Mentorship.Contract.Complete(&_Mentorship.TransactOpts, id)
}

// Complete is a paid mutator transaction binding the contract method 0x971d852f.
//
// Solidity: function complete(uint256 id) returns()
func (_Mentorship *MentorshipTransactorSession) Complete(id *big.Int) (*types.Transaction, error) {
	return _Mentorship.Contract.Complete(&_Mentorship.TransactOpts, id)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0xa8efd3bb.
//
// Solidity: function giveFeedback(uint256 id, bool upvote) returns()
func (_Mentorship *MentorshipTransactor) GiveFeedback(opts *bind.TransactOpts, id *big.Int, upvote bool) (*types.Transaction, error) {
	return _Mentorship.contract.Transact(opts, "giveFeedback", id, upvote)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0xa8efd3bb.
//
// Solidity: function giveFeedback(uint256 id, bool upvote) returns()
func (_Mentorship *MentorshipSession) GiveFeedback(id *big.Int, upvote bool) (*types.Transaction, error) {
	return _Mentorship.Contract.GiveFeedback(&_Mentorship.TransactOpts, id, upvote)
}

// GiveFeedback is a paid mutator transaction binding the contract method 0xa8efd3bb.
//
// Solidity: function giveFeedback(uint256 id, bool upvote) returns()
func (_Mentorship *MentorshipTransactorSession) GiveFeedback(id *big.Int, upvote bool) (*types.Transaction, error) {
	return _Mentorship.Contract.GiveFeedback(&_Mentorship.TransactOpts, id, upvote)
}

// Request is a paid mutator transaction binding the contract method 0x27c78c42.
//
// Solidity: function request(address mentor) returns(uint256 id)
func (_Mentorship *MentorshipTransactor) Request(opts *bind.TransactOpts, mentor common.Address) (*types.Transaction, error) {
	return _Mentorship.contract.Transact(opts, "request", mentor)
}

// Request is a paid mutator transaction binding the contract method 0x27c78c42.
//
// Solidity: function request(address mentor) returns(uint256 id)
func (_Mentorship *MentorshipSession) Request(mentor common.Address) (*types.Transaction, error) {
	return _Mentorship.Contract.Request(&_Mentorship.TransactOpts, mentor)
}

// Request is a paid mutator transaction binding the contract method 0x27c78c42.
//
// Solidity: function request(address mentor) returns(uint256 id)
func (_Mentorship *MentorshipTransactorSession) Request(mentor common.Address) (*types.Transaction, error) {
	return _Mentorship.Contract.Request(&_Mentorship.TransactOpts, mentor)
}

// MentorshipAcceptedIterator is returned from FilterAccepted and is used to iterate over the raw logs and unpacked data for Accepted events raised by the Mentorship contract.
type MentorshipAcceptedIterator struct {
	Event *MentorshipAccepted // Event containing the contract specifics and raw log

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
func (it *MentorshipAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MentorshipAccepted)
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
		it.Event = new(MentorshipAccepted)
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
func (it *MentorshipAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MentorshipAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MentorshipAccepted represents a Accepted event raised by the Mentorship contract.
type MentorshipAccepted struct {
	Id     *big.Int
	Mentor common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAccepted is a free log retrieval operation binding the contract event 0xec4eee35e0123e111ca1e78b349d631065c4e0d60bd85b917217cdf2fc749dd1.
//
// Solidity: event Accepted(uint256 indexed id, address indexed mentor)
func (_Mentorship *MentorshipFilterer) FilterAccepted(opts *bind.FilterOpts, id []*big.Int, mentor []common.Address) (*MentorshipAcceptedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.FilterLogs(opts, "Accepted", idRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return &MentorshipAcceptedIterator{contract: _Mentorship.contract, event: "Accepted", logs: logs, sub: sub}, nil
}

// WatchAccepted is a free log subscription operation binding the contract event 0xec4eee35e0123e111ca1e78b349d631065c4e0d60bd85b917217cdf2fc749dd1.
//
// Solidity: event Accepted(uint256 indexed id, address indexed mentor)
func (_Mentorship *MentorshipFilterer) WatchAccepted(opts *bind.WatchOpts, sink chan<- *MentorshipAccepted, id []*big.Int, mentor []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.WatchLogs(opts, "Accepted", idRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MentorshipAccepted)
				if err := _Mentorship.contract.UnpackLog(event, "Accepted", log); err != nil {
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

// ParseAccepted is a log parse operation binding the contract event 0xec4eee35e0123e111ca1e78b349d631065c4e0d60bd85b917217cdf2fc749dd1.
//
// Solidity: event Accepted(uint256 indexed id, address indexed mentor)
func (_Mentorship *MentorshipFilterer) ParseAccepted(log types.Log) (*MentorshipAccepted, error) {
	event := new(MentorshipAccepted)
	if err := _Mentorship.contract.UnpackLog(event, "Accepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MentorshipCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the Mentorship contract.
type MentorshipCanceledIterator struct {
	Event *MentorshipCanceled // Event containing the contract specifics and raw log

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
func (it *MentorshipCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MentorshipCanceled)
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
		it.Event = new(MentorshipCanceled)
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
func (it *MentorshipCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MentorshipCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MentorshipCanceled represents a Canceled event raised by the Mentorship contract.
type MentorshipCanceled struct {
	Id  *big.Int
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0xba05e869a56a216a7344a605ccad13464848815ad8bdf8d93ad788963a3cfdac.
//
// Solidity: event Canceled(uint256 indexed id, address indexed by)
func (_Mentorship *MentorshipFilterer) FilterCanceled(opts *bind.FilterOpts, id []*big.Int, by []common.Address) (*MentorshipCanceledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Mentorship.contract.FilterLogs(opts, "Canceled", idRule, byRule)
	if err != nil {
		return nil, err
	}
	return &MentorshipCanceledIterator{contract: _Mentorship.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0xba05e869a56a216a7344a605ccad13464848815ad8bdf8d93ad788963a3cfdac.
//
// Solidity: event Canceled(uint256 indexed id, address indexed by)
func (_Mentorship *MentorshipFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *MentorshipCanceled, id []*big.Int, by []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Mentorship.contract.WatchLogs(opts, "Canceled", idRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MentorshipCanceled)
				if err := _Mentorship.contract.UnpackLog(event, "Canceled", log); err != nil {
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

// ParseCanceled is a log parse operation binding the contract event 0xba05e869a56a216a7344a605ccad13464848815ad8bdf8d93ad788963a3cfdac.
//
// Solidity: event Canceled(uint256 indexed id, address indexed by)
func (_Mentorship *MentorshipFilterer) ParseCanceled(log types.Log) (*MentorshipCanceled, error) {
	event := new(MentorshipCanceled)
	if err := _Mentorship.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MentorshipCompletedIterator is returned from FilterCompleted and is used to iterate over the raw logs and unpacked data for Completed events raised by the Mentorship contract.
type MentorshipCompletedIterator struct {
	Event *MentorshipCompleted // Event containing the contract specifics and raw log

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
func (it *MentorshipCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MentorshipCompleted)
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
		it.Event = new(MentorshipCompleted)
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
func (it *MentorshipCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MentorshipCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MentorshipCompleted represents a Completed event raised by the Mentorship contract.
type MentorshipCompleted struct {
	Id     *big.Int
	Mentor common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCompleted is a free log retrieval operation binding the contract event 0xa4c87919f856d464ae137e2cb78e129ea90b599d0661f1755dbd7d733af7925a.
//
// Solidity: event Completed(uint256 indexed id, address indexed mentor)
func (_Mentorship *MentorshipFilterer) FilterCompleted(opts *bind.FilterOpts, id []*big.Int, mentor []common.Address) (*MentorshipCompletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.FilterLogs(opts, "Completed", idRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return &MentorshipCompletedIterator{contract: _Mentorship.contract, event: "Completed", logs: logs, sub: sub}, nil
}

// WatchCompleted is a free log subscription operation binding the contract event 0xa4c87919f856d464ae137e2cb78e129ea90b599d0661f1755dbd7d733af7925a.
//
// Solidity: event Completed(uint256 indexed id, address indexed mentor)
func (_Mentorship *MentorshipFilterer) WatchCompleted(opts *bind.WatchOpts, sink chan<- *MentorshipCompleted, id []*big.Int, mentor []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.WatchLogs(opts, "Completed", idRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MentorshipCompleted)
				if err := _Mentorship.contract.UnpackLog(event, "Completed", log); err != nil {
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

// ParseCompleted is a log parse operation binding the contract event 0xa4c87919f856d464ae137e2cb78e129ea90b599d0661f1755dbd7d733af7925a.
//
// Solidity: event Completed(uint256 indexed id, address indexed mentor)
func (_Mentorship *MentorshipFilterer) ParseCompleted(log types.Log) (*MentorshipCompleted, error) {
	event := new(MentorshipCompleted)
	if err := _Mentorship.contract.UnpackLog(event, "Completed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MentorshipFeedbackIterator is returned from FilterFeedback and is used to iterate over the raw logs and unpacked data for Feedback events raised by the Mentorship contract.
type MentorshipFeedbackIterator struct {
	Event *MentorshipFeedback // Event containing the contract specifics and raw log

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
func (it *MentorshipFeedbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MentorshipFeedback)
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
		it.Event = new(MentorshipFeedback)
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
func (it *MentorshipFeedbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MentorshipFeedbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MentorshipFeedback represents a Feedback event raised by the Mentorship contract.
type MentorshipFeedback struct {
	Id      *big.Int
	Student common.Address
	Mentor  common.Address
	Upvote  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeedback is a free log retrieval operation binding the contract event 0xdce7dd30820551e74fdb161d847acfdc8e2cec55fc7ab43707ea8f2a9ba4a273.
//
// Solidity: event Feedback(uint256 indexed id, address indexed student, address indexed mentor, bool upvote)
func (_Mentorship *MentorshipFilterer) FilterFeedback(opts *bind.FilterOpts, id []*big.Int, student []common.Address, mentor []common.Address) (*MentorshipFeedbackIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.FilterLogs(opts, "Feedback", idRule, studentRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return &MentorshipFeedbackIterator{contract: _Mentorship.contract, event: "Feedback", logs: logs, sub: sub}, nil
}

// WatchFeedback is a free log subscription operation binding the contract event 0xdce7dd30820551e74fdb161d847acfdc8e2cec55fc7ab43707ea8f2a9ba4a273.
//
// Solidity: event Feedback(uint256 indexed id, address indexed student, address indexed mentor, bool upvote)
func (_Mentorship *MentorshipFilterer) WatchFeedback(opts *bind.WatchOpts, sink chan<- *MentorshipFeedback, id []*big.Int, student []common.Address, mentor []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.WatchLogs(opts, "Feedback", idRule, studentRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MentorshipFeedback)
				if err := _Mentorship.contract.UnpackLog(event, "Feedback", log); err != nil {
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

// ParseFeedback is a log parse operation binding the contract event 0xdce7dd30820551e74fdb161d847acfdc8e2cec55fc7ab43707ea8f2a9ba4a273.
//
// Solidity: event Feedback(uint256 indexed id, address indexed student, address indexed mentor, bool upvote)
func (_Mentorship *MentorshipFilterer) ParseFeedback(log types.Log) (*MentorshipFeedback, error) {
	event := new(MentorshipFeedback)
	if err := _Mentorship.contract.UnpackLog(event, "Feedback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MentorshipRequestedIterator is returned from FilterRequested and is used to iterate over the raw logs and unpacked data for Requested events raised by the Mentorship contract.
type MentorshipRequestedIterator struct {
	Event *MentorshipRequested // Event containing the contract specifics and raw log

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
func (it *MentorshipRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MentorshipRequested)
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
		it.Event = new(MentorshipRequested)
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
func (it *MentorshipRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MentorshipRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MentorshipRequested represents a Requested event raised by the Mentorship contract.
type MentorshipRequested struct {
	Id      *big.Int
	Student common.Address
	Mentor  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRequested is a free log retrieval operation binding the contract event 0xe7e8c1f4c6d5a8fa417e2037c16b06ae7b175a22362d2c26d824f54bf91b0bf4.
//
// Solidity: event Requested(uint256 indexed id, address indexed student, address indexed mentor)
func (_Mentorship *MentorshipFilterer) FilterRequested(opts *bind.FilterOpts, id []*big.Int, student []common.Address, mentor []common.Address) (*MentorshipRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.FilterLogs(opts, "Requested", idRule, studentRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return &MentorshipRequestedIterator{contract: _Mentorship.contract, event: "Requested", logs: logs, sub: sub}, nil
}

// WatchRequested is a free log subscription operation binding the contract event 0xe7e8c1f4c6d5a8fa417e2037c16b06ae7b175a22362d2c26d824f54bf91b0bf4.
//
// Solidity: event Requested(uint256 indexed id, address indexed student, address indexed mentor)
func (_Mentorship *MentorshipFilterer) WatchRequested(opts *bind.WatchOpts, sink chan<- *MentorshipRequested, id []*big.Int, student []common.Address, mentor []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var studentRule []interface{}
	for _, studentItem := range student {
		studentRule = append(studentRule, studentItem)
	}
	var mentorRule []interface{}
	for _, mentorItem := range mentor {
		mentorRule = append(mentorRule, mentorItem)
	}

	logs, sub, err := _Mentorship.contract.WatchLogs(opts, "Requested", idRule, studentRule, mentorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MentorshipRequested)
				if err := _Mentorship.contract.UnpackLog(event, "Requested", log); err != nil {
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

// ParseRequested is a log parse operation binding the contract event 0xe7e8c1f4c6d5a8fa417e2037c16b06ae7b175a22362d2c26d824f54bf91b0bf4.
//
// Solidity: event Requested(uint256 indexed id, address indexed student, address indexed mentor)
func (_Mentorship *MentorshipFilterer) ParseRequested(log types.Log) (*MentorshipRequested, error) {
	event := new(MentorshipRequested)
	if err := _Mentorship.contract.UnpackLog(event, "Requested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
