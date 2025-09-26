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

// ContentRegistryMetaData contains all meta data concerning the ContentRegistry contract.
var ContentRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"ContentUploaded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getContent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"uploader\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalContents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserUploads\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isPublic\",\"type\":\"bool\"}],\"name\":\"uploadContent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userUploads\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContentRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContentRegistryMetaData.ABI instead.
var ContentRegistryABI = ContentRegistryMetaData.ABI

// ContentRegistry is an auto generated Go binding around an Ethereum contract.
type ContentRegistry struct {
	ContentRegistryCaller     // Read-only binding to the contract
	ContentRegistryTransactor // Write-only binding to the contract
	ContentRegistryFilterer   // Log filterer for contract events
}

// ContentRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContentRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContentRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContentRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContentRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContentRegistrySession struct {
	Contract     *ContentRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContentRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContentRegistryCallerSession struct {
	Contract *ContentRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContentRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContentRegistryTransactorSession struct {
	Contract     *ContentRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContentRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContentRegistryRaw struct {
	Contract *ContentRegistry // Generic contract binding to access the raw methods on
}

// ContentRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContentRegistryCallerRaw struct {
	Contract *ContentRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContentRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContentRegistryTransactorRaw struct {
	Contract *ContentRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContentRegistry creates a new instance of ContentRegistry, bound to a specific deployed contract.
func NewContentRegistry(address common.Address, backend bind.ContractBackend) (*ContentRegistry, error) {
	contract, err := bindContentRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContentRegistry{ContentRegistryCaller: ContentRegistryCaller{contract: contract}, ContentRegistryTransactor: ContentRegistryTransactor{contract: contract}, ContentRegistryFilterer: ContentRegistryFilterer{contract: contract}}, nil
}

// NewContentRegistryCaller creates a new read-only instance of ContentRegistry, bound to a specific deployed contract.
func NewContentRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContentRegistryCaller, error) {
	contract, err := bindContentRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContentRegistryCaller{contract: contract}, nil
}

// NewContentRegistryTransactor creates a new write-only instance of ContentRegistry, bound to a specific deployed contract.
func NewContentRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContentRegistryTransactor, error) {
	contract, err := bindContentRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContentRegistryTransactor{contract: contract}, nil
}

// NewContentRegistryFilterer creates a new log filterer instance of ContentRegistry, bound to a specific deployed contract.
func NewContentRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContentRegistryFilterer, error) {
	contract, err := bindContentRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContentRegistryFilterer{contract: contract}, nil
}

// bindContentRegistry binds a generic wrapper to an already deployed contract.
func bindContentRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContentRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentRegistry *ContentRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentRegistry.Contract.ContentRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentRegistry *ContentRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentRegistry.Contract.ContentRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentRegistry *ContentRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentRegistry.Contract.ContentRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContentRegistry *ContentRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContentRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContentRegistry *ContentRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContentRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContentRegistry *ContentRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContentRegistry.Contract.contract.Transact(opts, method, params...)
}

// Contents is a free data retrieval call binding the contract method 0xb5ecf912.
//
// Solidity: function contents(uint256 ) view returns(address uploader, string cid, string title, bool isPublic)
func (_ContentRegistry *ContentRegistryCaller) Contents(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Uploader common.Address
	Cid      string
	Title    string
	IsPublic bool
}, error) {
	var out []interface{}
	err := _ContentRegistry.contract.Call(opts, &out, "contents", arg0)

	outstruct := new(struct {
		Uploader common.Address
		Cid      string
		Title    string
		IsPublic bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Uploader = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Cid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Title = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IsPublic = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Contents is a free data retrieval call binding the contract method 0xb5ecf912.
//
// Solidity: function contents(uint256 ) view returns(address uploader, string cid, string title, bool isPublic)
func (_ContentRegistry *ContentRegistrySession) Contents(arg0 *big.Int) (struct {
	Uploader common.Address
	Cid      string
	Title    string
	IsPublic bool
}, error) {
	return _ContentRegistry.Contract.Contents(&_ContentRegistry.CallOpts, arg0)
}

// Contents is a free data retrieval call binding the contract method 0xb5ecf912.
//
// Solidity: function contents(uint256 ) view returns(address uploader, string cid, string title, bool isPublic)
func (_ContentRegistry *ContentRegistryCallerSession) Contents(arg0 *big.Int) (struct {
	Uploader common.Address
	Cid      string
	Title    string
	IsPublic bool
}, error) {
	return _ContentRegistry.Contract.Contents(&_ContentRegistry.CallOpts, arg0)
}

// GetContent is a free data retrieval call binding the contract method 0x0b7ad54c.
//
// Solidity: function getContent(uint256 _id) view returns(address uploader, string cid, string title, bool isPublic)
func (_ContentRegistry *ContentRegistryCaller) GetContent(opts *bind.CallOpts, _id *big.Int) (struct {
	Uploader common.Address
	Cid      string
	Title    string
	IsPublic bool
}, error) {
	var out []interface{}
	err := _ContentRegistry.contract.Call(opts, &out, "getContent", _id)

	outstruct := new(struct {
		Uploader common.Address
		Cid      string
		Title    string
		IsPublic bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Uploader = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Cid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Title = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.IsPublic = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetContent is a free data retrieval call binding the contract method 0x0b7ad54c.
//
// Solidity: function getContent(uint256 _id) view returns(address uploader, string cid, string title, bool isPublic)
func (_ContentRegistry *ContentRegistrySession) GetContent(_id *big.Int) (struct {
	Uploader common.Address
	Cid      string
	Title    string
	IsPublic bool
}, error) {
	return _ContentRegistry.Contract.GetContent(&_ContentRegistry.CallOpts, _id)
}

// GetContent is a free data retrieval call binding the contract method 0x0b7ad54c.
//
// Solidity: function getContent(uint256 _id) view returns(address uploader, string cid, string title, bool isPublic)
func (_ContentRegistry *ContentRegistryCallerSession) GetContent(_id *big.Int) (struct {
	Uploader common.Address
	Cid      string
	Title    string
	IsPublic bool
}, error) {
	return _ContentRegistry.Contract.GetContent(&_ContentRegistry.CallOpts, _id)
}

// GetTotalContents is a free data retrieval call binding the contract method 0xa00dafe0.
//
// Solidity: function getTotalContents() view returns(uint256)
func (_ContentRegistry *ContentRegistryCaller) GetTotalContents(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContentRegistry.contract.Call(opts, &out, "getTotalContents")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalContents is a free data retrieval call binding the contract method 0xa00dafe0.
//
// Solidity: function getTotalContents() view returns(uint256)
func (_ContentRegistry *ContentRegistrySession) GetTotalContents() (*big.Int, error) {
	return _ContentRegistry.Contract.GetTotalContents(&_ContentRegistry.CallOpts)
}

// GetTotalContents is a free data retrieval call binding the contract method 0xa00dafe0.
//
// Solidity: function getTotalContents() view returns(uint256)
func (_ContentRegistry *ContentRegistryCallerSession) GetTotalContents() (*big.Int, error) {
	return _ContentRegistry.Contract.GetTotalContents(&_ContentRegistry.CallOpts)
}

// GetUserUploads is a free data retrieval call binding the contract method 0x412bd348.
//
// Solidity: function getUserUploads(address _user) view returns(uint256[])
func (_ContentRegistry *ContentRegistryCaller) GetUserUploads(opts *bind.CallOpts, _user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ContentRegistry.contract.Call(opts, &out, "getUserUploads", _user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserUploads is a free data retrieval call binding the contract method 0x412bd348.
//
// Solidity: function getUserUploads(address _user) view returns(uint256[])
func (_ContentRegistry *ContentRegistrySession) GetUserUploads(_user common.Address) ([]*big.Int, error) {
	return _ContentRegistry.Contract.GetUserUploads(&_ContentRegistry.CallOpts, _user)
}

// GetUserUploads is a free data retrieval call binding the contract method 0x412bd348.
//
// Solidity: function getUserUploads(address _user) view returns(uint256[])
func (_ContentRegistry *ContentRegistryCallerSession) GetUserUploads(_user common.Address) ([]*big.Int, error) {
	return _ContentRegistry.Contract.GetUserUploads(&_ContentRegistry.CallOpts, _user)
}

// UserUploads is a free data retrieval call binding the contract method 0x030d1efe.
//
// Solidity: function userUploads(address , uint256 ) view returns(uint256)
func (_ContentRegistry *ContentRegistryCaller) UserUploads(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContentRegistry.contract.Call(opts, &out, "userUploads", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUploads is a free data retrieval call binding the contract method 0x030d1efe.
//
// Solidity: function userUploads(address , uint256 ) view returns(uint256)
func (_ContentRegistry *ContentRegistrySession) UserUploads(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ContentRegistry.Contract.UserUploads(&_ContentRegistry.CallOpts, arg0, arg1)
}

// UserUploads is a free data retrieval call binding the contract method 0x030d1efe.
//
// Solidity: function userUploads(address , uint256 ) view returns(uint256)
func (_ContentRegistry *ContentRegistryCallerSession) UserUploads(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _ContentRegistry.Contract.UserUploads(&_ContentRegistry.CallOpts, arg0, arg1)
}

// UploadContent is a paid mutator transaction binding the contract method 0x1e6aab26.
//
// Solidity: function uploadContent(string _cid, string _title, bool _isPublic) returns()
func (_ContentRegistry *ContentRegistryTransactor) UploadContent(opts *bind.TransactOpts, _cid string, _title string, _isPublic bool) (*types.Transaction, error) {
	return _ContentRegistry.contract.Transact(opts, "uploadContent", _cid, _title, _isPublic)
}

// UploadContent is a paid mutator transaction binding the contract method 0x1e6aab26.
//
// Solidity: function uploadContent(string _cid, string _title, bool _isPublic) returns()
func (_ContentRegistry *ContentRegistrySession) UploadContent(_cid string, _title string, _isPublic bool) (*types.Transaction, error) {
	return _ContentRegistry.Contract.UploadContent(&_ContentRegistry.TransactOpts, _cid, _title, _isPublic)
}

// UploadContent is a paid mutator transaction binding the contract method 0x1e6aab26.
//
// Solidity: function uploadContent(string _cid, string _title, bool _isPublic) returns()
func (_ContentRegistry *ContentRegistryTransactorSession) UploadContent(_cid string, _title string, _isPublic bool) (*types.Transaction, error) {
	return _ContentRegistry.Contract.UploadContent(&_ContentRegistry.TransactOpts, _cid, _title, _isPublic)
}

// ContentRegistryContentUploadedIterator is returned from FilterContentUploaded and is used to iterate over the raw logs and unpacked data for ContentUploaded events raised by the ContentRegistry contract.
type ContentRegistryContentUploadedIterator struct {
	Event *ContentRegistryContentUploaded // Event containing the contract specifics and raw log

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
func (it *ContentRegistryContentUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContentRegistryContentUploaded)
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
		it.Event = new(ContentRegistryContentUploaded)
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
func (it *ContentRegistryContentUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContentRegistryContentUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContentRegistryContentUploaded represents a ContentUploaded event raised by the ContentRegistry contract.
type ContentRegistryContentUploaded struct {
	Uploader  common.Address
	ContentId *big.Int
	Title     string
	IsPublic  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContentUploaded is a free log retrieval operation binding the contract event 0xb8a0e9b458b12916c179180157f3e2a048ce050a5a3fbbc7c2d1de9ba1f988e5.
//
// Solidity: event ContentUploaded(address indexed uploader, uint256 contentId, string title, bool isPublic)
func (_ContentRegistry *ContentRegistryFilterer) FilterContentUploaded(opts *bind.FilterOpts, uploader []common.Address) (*ContentRegistryContentUploadedIterator, error) {

	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}

	logs, sub, err := _ContentRegistry.contract.FilterLogs(opts, "ContentUploaded", uploaderRule)
	if err != nil {
		return nil, err
	}
	return &ContentRegistryContentUploadedIterator{contract: _ContentRegistry.contract, event: "ContentUploaded", logs: logs, sub: sub}, nil
}

// WatchContentUploaded is a free log subscription operation binding the contract event 0xb8a0e9b458b12916c179180157f3e2a048ce050a5a3fbbc7c2d1de9ba1f988e5.
//
// Solidity: event ContentUploaded(address indexed uploader, uint256 contentId, string title, bool isPublic)
func (_ContentRegistry *ContentRegistryFilterer) WatchContentUploaded(opts *bind.WatchOpts, sink chan<- *ContentRegistryContentUploaded, uploader []common.Address) (event.Subscription, error) {

	var uploaderRule []interface{}
	for _, uploaderItem := range uploader {
		uploaderRule = append(uploaderRule, uploaderItem)
	}

	logs, sub, err := _ContentRegistry.contract.WatchLogs(opts, "ContentUploaded", uploaderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContentRegistryContentUploaded)
				if err := _ContentRegistry.contract.UnpackLog(event, "ContentUploaded", log); err != nil {
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

// ParseContentUploaded is a log parse operation binding the contract event 0xb8a0e9b458b12916c179180157f3e2a048ce050a5a3fbbc7c2d1de9ba1f988e5.
//
// Solidity: event ContentUploaded(address indexed uploader, uint256 contentId, string title, bool isPublic)
func (_ContentRegistry *ContentRegistryFilterer) ParseContentUploaded(log types.Log) (*ContentRegistryContentUploaded, error) {
	event := new(ContentRegistryContentUploaded)
	if err := _ContentRegistry.contract.UnpackLog(event, "ContentUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
