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

// UserRegistryMetaData contains all meta data concerning the UserRegistry contract.
var UserRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumUserRegistry.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollNumber\",\"type\":\"string\"}],\"name\":\"UserVerified\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"internalType\":\"enumUserRegistry.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"rollNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"program\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"major\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pictureHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isVerified\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"internalType\":\"enumUserRegistry.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"rollNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"program\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"major\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pictureHash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"enumUserRegistry.Role\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_rollNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_program\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_major\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_pictureHash\",\"type\":\"string\"}],\"name\":\"verifyUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UserRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use UserRegistryMetaData.ABI instead.
var UserRegistryABI = UserRegistryMetaData.ABI

// UserRegistry is an auto generated Go binding around an Ethereum contract.
type UserRegistry struct {
	UserRegistryCaller     // Read-only binding to the contract
	UserRegistryTransactor // Write-only binding to the contract
	UserRegistryFilterer   // Log filterer for contract events
}

// UserRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserRegistrySession struct {
	Contract     *UserRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserRegistryCallerSession struct {
	Contract *UserRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UserRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserRegistryTransactorSession struct {
	Contract     *UserRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UserRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRegistryRaw struct {
	Contract *UserRegistry // Generic contract binding to access the raw methods on
}

// UserRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserRegistryCallerRaw struct {
	Contract *UserRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// UserRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserRegistryTransactorRaw struct {
	Contract *UserRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserRegistry creates a new instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistry(address common.Address, backend bind.ContractBackend) (*UserRegistry, error) {
	contract, err := bindUserRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserRegistry{UserRegistryCaller: UserRegistryCaller{contract: contract}, UserRegistryTransactor: UserRegistryTransactor{contract: contract}, UserRegistryFilterer: UserRegistryFilterer{contract: contract}}, nil
}

// NewUserRegistryCaller creates a new read-only instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistryCaller(address common.Address, caller bind.ContractCaller) (*UserRegistryCaller, error) {
	contract, err := bindUserRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserRegistryCaller{contract: contract}, nil
}

// NewUserRegistryTransactor creates a new write-only instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*UserRegistryTransactor, error) {
	contract, err := bindUserRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserRegistryTransactor{contract: contract}, nil
}

// NewUserRegistryFilterer creates a new log filterer instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*UserRegistryFilterer, error) {
	contract, err := bindUserRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserRegistryFilterer{contract: contract}, nil
}

// bindUserRegistry binds a generic wrapper to an already deployed contract.
func bindUserRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserRegistry *UserRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserRegistry.Contract.UserRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserRegistry *UserRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserRegistry.Contract.UserRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserRegistry *UserRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserRegistry.Contract.UserRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserRegistry *UserRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserRegistry *UserRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserRegistry *UserRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns(bool verified, uint8 role, string name, string rollNumber, string program, string major, string pictureHash)
func (_UserRegistry *UserRegistryCaller) GetUser(opts *bind.CallOpts, _user common.Address) (struct {
	Verified    bool
	Role        uint8
	Name        string
	RollNumber  string
	Program     string
	Major       string
	PictureHash string
}, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getUser", _user)

	outstruct := new(struct {
		Verified    bool
		Role        uint8
		Name        string
		RollNumber  string
		Program     string
		Major       string
		PictureHash string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Verified = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Role = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.RollNumber = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Program = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Major = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.PictureHash = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns(bool verified, uint8 role, string name, string rollNumber, string program, string major, string pictureHash)
func (_UserRegistry *UserRegistrySession) GetUser(_user common.Address) (struct {
	Verified    bool
	Role        uint8
	Name        string
	RollNumber  string
	Program     string
	Major       string
	PictureHash string
}, error) {
	return _UserRegistry.Contract.GetUser(&_UserRegistry.CallOpts, _user)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address _user) view returns(bool verified, uint8 role, string name, string rollNumber, string program, string major, string pictureHash)
func (_UserRegistry *UserRegistryCallerSession) GetUser(_user common.Address) (struct {
	Verified    bool
	Role        uint8
	Name        string
	RollNumber  string
	Program     string
	Major       string
	PictureHash string
}, error) {
	return _UserRegistry.Contract.GetUser(&_UserRegistry.CallOpts, _user)
}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address _user) view returns(bool)
func (_UserRegistry *UserRegistryCaller) IsVerified(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "isVerified", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address _user) view returns(bool)
func (_UserRegistry *UserRegistrySession) IsVerified(_user common.Address) (bool, error) {
	return _UserRegistry.Contract.IsVerified(&_UserRegistry.CallOpts, _user)
}

// IsVerified is a free data retrieval call binding the contract method 0xb9209e33.
//
// Solidity: function isVerified(address _user) view returns(bool)
func (_UserRegistry *UserRegistryCallerSession) IsVerified(_user common.Address) (bool, error) {
	return _UserRegistry.Contract.IsVerified(&_UserRegistry.CallOpts, _user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserRegistry *UserRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserRegistry *UserRegistrySession) Owner() (common.Address, error) {
	return _UserRegistry.Contract.Owner(&_UserRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UserRegistry *UserRegistryCallerSession) Owner() (common.Address, error) {
	return _UserRegistry.Contract.Owner(&_UserRegistry.CallOpts)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(bool verified, uint8 role, string name, string rollNumber, string program, string major, string pictureHash)
func (_UserRegistry *UserRegistryCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	Verified    bool
	Role        uint8
	Name        string
	RollNumber  string
	Program     string
	Major       string
	PictureHash string
}, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		Verified    bool
		Role        uint8
		Name        string
		RollNumber  string
		Program     string
		Major       string
		PictureHash string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Verified = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Role = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.RollNumber = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Program = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Major = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.PictureHash = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(bool verified, uint8 role, string name, string rollNumber, string program, string major, string pictureHash)
func (_UserRegistry *UserRegistrySession) Users(arg0 common.Address) (struct {
	Verified    bool
	Role        uint8
	Name        string
	RollNumber  string
	Program     string
	Major       string
	PictureHash string
}, error) {
	return _UserRegistry.Contract.Users(&_UserRegistry.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(bool verified, uint8 role, string name, string rollNumber, string program, string major, string pictureHash)
func (_UserRegistry *UserRegistryCallerSession) Users(arg0 common.Address) (struct {
	Verified    bool
	Role        uint8
	Name        string
	RollNumber  string
	Program     string
	Major       string
	PictureHash string
}, error) {
	return _UserRegistry.Contract.Users(&_UserRegistry.CallOpts, arg0)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x7937a348.
//
// Solidity: function verifyUser(address _user, uint8 _role, string _name, string _rollNumber, string _program, string _major, string _pictureHash) returns()
func (_UserRegistry *UserRegistryTransactor) VerifyUser(opts *bind.TransactOpts, _user common.Address, _role uint8, _name string, _rollNumber string, _program string, _major string, _pictureHash string) (*types.Transaction, error) {
	return _UserRegistry.contract.Transact(opts, "verifyUser", _user, _role, _name, _rollNumber, _program, _major, _pictureHash)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x7937a348.
//
// Solidity: function verifyUser(address _user, uint8 _role, string _name, string _rollNumber, string _program, string _major, string _pictureHash) returns()
func (_UserRegistry *UserRegistrySession) VerifyUser(_user common.Address, _role uint8, _name string, _rollNumber string, _program string, _major string, _pictureHash string) (*types.Transaction, error) {
	return _UserRegistry.Contract.VerifyUser(&_UserRegistry.TransactOpts, _user, _role, _name, _rollNumber, _program, _major, _pictureHash)
}

// VerifyUser is a paid mutator transaction binding the contract method 0x7937a348.
//
// Solidity: function verifyUser(address _user, uint8 _role, string _name, string _rollNumber, string _program, string _major, string _pictureHash) returns()
func (_UserRegistry *UserRegistryTransactorSession) VerifyUser(_user common.Address, _role uint8, _name string, _rollNumber string, _program string, _major string, _pictureHash string) (*types.Transaction, error) {
	return _UserRegistry.Contract.VerifyUser(&_UserRegistry.TransactOpts, _user, _role, _name, _rollNumber, _program, _major, _pictureHash)
}

// UserRegistryUserVerifiedIterator is returned from FilterUserVerified and is used to iterate over the raw logs and unpacked data for UserVerified events raised by the UserRegistry contract.
type UserRegistryUserVerifiedIterator struct {
	Event *UserRegistryUserVerified // Event containing the contract specifics and raw log

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
func (it *UserRegistryUserVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserRegistryUserVerified)
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
		it.Event = new(UserRegistryUserVerified)
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
func (it *UserRegistryUserVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserRegistryUserVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserRegistryUserVerified represents a UserVerified event raised by the UserRegistry contract.
type UserRegistryUserVerified struct {
	User       common.Address
	Role       uint8
	Name       string
	RollNumber string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserVerified is a free log retrieval operation binding the contract event 0xbf336d49dae0a6defb13e95f85c9df24dbd2b429dae44068e3c79bbc3ce6c566.
//
// Solidity: event UserVerified(address indexed user, uint8 role, string name, string rollNumber)
func (_UserRegistry *UserRegistryFilterer) FilterUserVerified(opts *bind.FilterOpts, user []common.Address) (*UserRegistryUserVerifiedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserRegistry.contract.FilterLogs(opts, "UserVerified", userRule)
	if err != nil {
		return nil, err
	}
	return &UserRegistryUserVerifiedIterator{contract: _UserRegistry.contract, event: "UserVerified", logs: logs, sub: sub}, nil
}

// WatchUserVerified is a free log subscription operation binding the contract event 0xbf336d49dae0a6defb13e95f85c9df24dbd2b429dae44068e3c79bbc3ce6c566.
//
// Solidity: event UserVerified(address indexed user, uint8 role, string name, string rollNumber)
func (_UserRegistry *UserRegistryFilterer) WatchUserVerified(opts *bind.WatchOpts, sink chan<- *UserRegistryUserVerified, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _UserRegistry.contract.WatchLogs(opts, "UserVerified", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserRegistryUserVerified)
				if err := _UserRegistry.contract.UnpackLog(event, "UserVerified", log); err != nil {
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

// ParseUserVerified is a log parse operation binding the contract event 0xbf336d49dae0a6defb13e95f85c9df24dbd2b429dae44068e3c79bbc3ce6c566.
//
// Solidity: event UserVerified(address indexed user, uint8 role, string name, string rollNumber)
func (_UserRegistry *UserRegistryFilterer) ParseUserVerified(log types.Log) (*UserRegistryUserVerified, error) {
	event := new(UserRegistryUserVerified)
	if err := _UserRegistry.contract.UnpackLog(event, "UserVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
