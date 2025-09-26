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

// ContentAccessMetaData contains all meta data concerning the ContentAccess contract.
var ContentAccessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contentRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"Granted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"canAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contentRegistry\",\"outputs\":[{\"internalType\":\"contractIContentRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"grant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"contentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"viewer\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContentAccessABI is the input ABI used to generate the binding from.
// Deprecated: Use ContentAccessMetaData.ABI instead.
var ContentAccessABI = ContentAccessMetaData.ABI

// ContentAccess is an auto generated Go binding around an Ethereum contract.
type ContentAccess struct {
	ContentAccessCaller     // Read-only binding to the contract
	ContentAccessTransactor // Write-only binding to the contract
	ContentAccessFilterer   // Log filterer for contract events
}

// ContentAccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContentAccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentAccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContentAccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentAccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContentAccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentAccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContentAccessSession struct {
	Contract     *ContentAccess    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContentAccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContentAccessCallerSession struct {
	Contract *ContentAccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ContentAccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContentAccessTransactorSession struct {
	Contract     *ContentAccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContentAccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContentAccessRaw struct {
	Contract *ContentAccess // Generic contract binding to access the raw methods on
}

// ContentAccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContentAccessCallerRaw struct {
	Contract *ContentAccessCaller // Generic read-only contract binding to access the raw methods on
}

// ContentAccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContentAccessTransactorRaw struct {
	Contract *ContentAccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContentAccess creates a new instance of ContentAccess, bound to a specific deployed contract.
func NewContentAccess(address common.Address, backend bind.ContractBackend) (*ContentAccess, error) {
	contract, err := bindContentAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContentAccess{ContentAccessCaller: ContentAccessCaller{contract: contract}, ContentAccessTransactor: ContentAccessTransactor{contract: contract}, ContentAccessFilterer: ContentAccessFilterer{contract: contract}}, nil
}

// NewContentAccessCaller creates a new read-only instance of ContentAccess, bound to a specific deployed contract.
func NewContentAccessCaller(address common.Address, caller bind.ContractCaller) (*ContentAccessCaller, error) {
	contract, err := bindContentAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContentAccessCaller{contract: contract}, nil
}

// NewContentAccessTransactor creates a new write-only instance of ContentAccess, bound to a specific deployed contract.
func NewContentAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*ContentAccessTransactor, error) {
	contract, err := bindContentAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContentAccessTransactor{contract: contract}, nil
}

// NewContentAccessFilterer creates a new log filterer instance of ContentAccess, bound to a specific deployed contract.
func NewContentAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*ContentAccessFilterer, error) {
	contract, err := bindContentAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContentAccessFilterer{contract: contract}, nil
}

// bindContentAccess binds a generic wrapper to an already deployed contract.
func bindContentAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContentAccessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentAccess *ContentAccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentAccess.Contract.ContentAccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentAccess *ContentAccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentAccess.Contract.ContentAccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentAccess *ContentAccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentAccess.Contract.ContentAccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentAccess *ContentAccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentAccess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentAccess *ContentAccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentAccess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentAccess *ContentAccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentAccess.Contract.contract.Transact(opts, method, params...)
}

// CanAccess is a free data retrieval call binding the contract method 0xd5b2d1ed.
//
// Solidity: function canAccess(uint256 contentId, address viewer) view returns(bool)
func (_ContentAccess *ContentAccessCaller) CanAccess(opts *bind.CallOpts, contentId *big.Int, viewer common.Address) (bool, error) {
	var out []interface{}
	err := _ContentAccess.contract.Call(opts, &out, "canAccess", contentId, viewer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanAccess is a free data retrieval call binding the contract method 0xd5b2d1ed.
//
// Solidity: function canAccess(uint256 contentId, address viewer) view returns(bool)
func (_ContentAccess *ContentAccessSession) CanAccess(contentId *big.Int, viewer common.Address) (bool, error) {
	return _ContentAccess.Contract.CanAccess(&_ContentAccess.CallOpts, contentId, viewer)
}

// CanAccess is a free data retrieval call binding the contract method 0xd5b2d1ed.
//
// Solidity: function canAccess(uint256 contentId, address viewer) view returns(bool)
func (_ContentAccess *ContentAccessCallerSession) CanAccess(contentId *big.Int, viewer common.Address) (bool, error) {
	return _ContentAccess.Contract.CanAccess(&_ContentAccess.CallOpts, contentId, viewer)
}

// ContentRegistry is a free data retrieval call binding the contract method 0x27a4f2bb.
//
// Solidity: function contentRegistry() view returns(address)
func (_ContentAccess *ContentAccessCaller) ContentRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContentAccess.contract.Call(opts, &out, "contentRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContentRegistry is a free data retrieval call binding the contract method 0x27a4f2bb.
//
// Solidity: function contentRegistry() view returns(address)
func (_ContentAccess *ContentAccessSession) ContentRegistry() (common.Address, error) {
	return _ContentAccess.Contract.ContentRegistry(&_ContentAccess.CallOpts)
}

// ContentRegistry is a free data retrieval call binding the contract method 0x27a4f2bb.
//
// Solidity: function contentRegistry() view returns(address)
func (_ContentAccess *ContentAccessCallerSession) ContentRegistry() (common.Address, error) {
	return _ContentAccess.Contract.ContentRegistry(&_ContentAccess.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 , address ) view returns(bool)
func (_ContentAccess *ContentAccessCaller) IsWhitelisted(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ContentAccess.contract.Call(opts, &out, "isWhitelisted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 , address ) view returns(bool)
func (_ContentAccess *ContentAccessSession) IsWhitelisted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ContentAccess.Contract.IsWhitelisted(&_ContentAccess.CallOpts, arg0, arg1)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x7d22c35c.
//
// Solidity: function isWhitelisted(uint256 , address ) view returns(bool)
func (_ContentAccess *ContentAccessCallerSession) IsWhitelisted(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ContentAccess.Contract.IsWhitelisted(&_ContentAccess.CallOpts, arg0, arg1)
}

// Grant is a paid mutator transaction binding the contract method 0x887d7c3e.
//
// Solidity: function grant(uint256 contentId, address viewer) returns()
func (_ContentAccess *ContentAccessTransactor) Grant(opts *bind.TransactOpts, contentId *big.Int, viewer common.Address) (*types.Transaction, error) {
	return _ContentAccess.contract.Transact(opts, "grant", contentId, viewer)
}

// Grant is a paid mutator transaction binding the contract method 0x887d7c3e.
//
// Solidity: function grant(uint256 contentId, address viewer) returns()
func (_ContentAccess *ContentAccessSession) Grant(contentId *big.Int, viewer common.Address) (*types.Transaction, error) {
	return _ContentAccess.Contract.Grant(&_ContentAccess.TransactOpts, contentId, viewer)
}

// Grant is a paid mutator transaction binding the contract method 0x887d7c3e.
//
// Solidity: function grant(uint256 contentId, address viewer) returns()
func (_ContentAccess *ContentAccessTransactorSession) Grant(contentId *big.Int, viewer common.Address) (*types.Transaction, error) {
	return _ContentAccess.Contract.Grant(&_ContentAccess.TransactOpts, contentId, viewer)
}

// Revoke is a paid mutator transaction binding the contract method 0x20d154da.
//
// Solidity: function revoke(uint256 contentId, address viewer) returns()
func (_ContentAccess *ContentAccessTransactor) Revoke(opts *bind.TransactOpts, contentId *big.Int, viewer common.Address) (*types.Transaction, error) {
	return _ContentAccess.contract.Transact(opts, "revoke", contentId, viewer)
}

// Revoke is a paid mutator transaction binding the contract method 0x20d154da.
//
// Solidity: function revoke(uint256 contentId, address viewer) returns()
func (_ContentAccess *ContentAccessSession) Revoke(contentId *big.Int, viewer common.Address) (*types.Transaction, error) {
	return _ContentAccess.Contract.Revoke(&_ContentAccess.TransactOpts, contentId, viewer)
}

// Revoke is a paid mutator transaction binding the contract method 0x20d154da.
//
// Solidity: function revoke(uint256 contentId, address viewer) returns()
func (_ContentAccess *ContentAccessTransactorSession) Revoke(contentId *big.Int, viewer common.Address) (*types.Transaction, error) {
	return _ContentAccess.Contract.Revoke(&_ContentAccess.TransactOpts, contentId, viewer)
}

// ContentAccessGrantedIterator is returned from FilterGranted and is used to iterate over the raw logs and unpacked data for Granted events raised by the ContentAccess contract.
type ContentAccessGrantedIterator struct {
	Event *ContentAccessGranted // Event containing the contract specifics and raw log

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
func (it *ContentAccessGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentAccessGranted)
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
		it.Event = new(ContentAccessGranted)
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
func (it *ContentAccessGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentAccessGranted represents a Granted event raised by the ContentAccess contract.
type ContentAccessGranted struct {
	ContentId *big.Int
	Viewer    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGranted is a free log retrieval operation binding the contract event 0x13b0d5f348cbb4d844b91093c270d22f8eb5c4ab9b73ebd7fe352e4d557f423a.
//
// Solidity: event Granted(uint256 indexed contentId, address indexed viewer)
func (_ContentAccess *ContentAccessFilterer) FilterGranted(opts *bind.FilterOpts, contentId []*big.Int, viewer []common.Address) (*ContentAccessGrantedIterator, error) {

	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var viewerRule []interface{}
	for _, viewerItem := range viewer {
		viewerRule = append(viewerRule, viewerItem)
	}

	logs, sub, err := _ContentAccess.contract.FilterLogs(opts, "Granted", contentIdRule, viewerRule)
	if err != nil {
		return nil, err
	}
	return &ContentAccessGrantedIterator{contract: _ContentAccess.contract, event: "Granted", logs: logs, sub: sub}, nil
}

// WatchGranted is a free log subscription operation binding the contract event 0x13b0d5f348cbb4d844b91093c270d22f8eb5c4ab9b73ebd7fe352e4d557f423a.
//
// Solidity: event Granted(uint256 indexed contentId, address indexed viewer)
func (_ContentAccess *ContentAccessFilterer) WatchGranted(opts *bind.WatchOpts, sink chan<- *ContentAccessGranted, contentId []*big.Int, viewer []common.Address) (event.Subscription, error) {

	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var viewerRule []interface{}
	for _, viewerItem := range viewer {
		viewerRule = append(viewerRule, viewerItem)
	}

	logs, sub, err := _ContentAccess.contract.WatchLogs(opts, "Granted", contentIdRule, viewerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentAccessGranted)
				if err := _ContentAccess.contract.UnpackLog(event, "Granted", log); err != nil {
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

// ParseGranted is a log parse operation binding the contract event 0x13b0d5f348cbb4d844b91093c270d22f8eb5c4ab9b73ebd7fe352e4d557f423a.
//
// Solidity: event Granted(uint256 indexed contentId, address indexed viewer)
func (_ContentAccess *ContentAccessFilterer) ParseGranted(log types.Log) (*ContentAccessGranted, error) {
	event := new(ContentAccessGranted)
	if err := _ContentAccess.contract.UnpackLog(event, "Granted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContentAccessRevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the ContentAccess contract.
type ContentAccessRevokedIterator struct {
	Event *ContentAccessRevoked // Event containing the contract specifics and raw log

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
func (it *ContentAccessRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentAccessRevoked)
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
		it.Event = new(ContentAccessRevoked)
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
func (it *ContentAccessRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentAccessRevoked represents a Revoked event raised by the ContentAccess contract.
type ContentAccessRevoked struct {
	ContentId *big.Int
	Viewer    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x0819476e55ffbe23df64353d5f7898429609397cc93c44300b058b51648f7878.
//
// Solidity: event Revoked(uint256 indexed contentId, address indexed viewer)
func (_ContentAccess *ContentAccessFilterer) FilterRevoked(opts *bind.FilterOpts, contentId []*big.Int, viewer []common.Address) (*ContentAccessRevokedIterator, error) {

	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var viewerRule []interface{}
	for _, viewerItem := range viewer {
		viewerRule = append(viewerRule, viewerItem)
	}

	logs, sub, err := _ContentAccess.contract.FilterLogs(opts, "Revoked", contentIdRule, viewerRule)
	if err != nil {
		return nil, err
	}
	return &ContentAccessRevokedIterator{contract: _ContentAccess.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x0819476e55ffbe23df64353d5f7898429609397cc93c44300b058b51648f7878.
//
// Solidity: event Revoked(uint256 indexed contentId, address indexed viewer)
func (_ContentAccess *ContentAccessFilterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *ContentAccessRevoked, contentId []*big.Int, viewer []common.Address) (event.Subscription, error) {

	var contentIdRule []interface{}
	for _, contentIdItem := range contentId {
		contentIdRule = append(contentIdRule, contentIdItem)
	}
	var viewerRule []interface{}
	for _, viewerItem := range viewer {
		viewerRule = append(viewerRule, viewerItem)
	}

	logs, sub, err := _ContentAccess.contract.WatchLogs(opts, "Revoked", contentIdRule, viewerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentAccessRevoked)
				if err := _ContentAccess.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x0819476e55ffbe23df64353d5f7898429609397cc93c44300b058b51648f7878.
//
// Solidity: event Revoked(uint256 indexed contentId, address indexed viewer)
func (_ContentAccess *ContentAccessFilterer) ParseRevoked(log types.Log) (*ContentAccessRevoked, error) {
	event := new(ContentAccessRevoked)
	if err := _ContentAccess.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
