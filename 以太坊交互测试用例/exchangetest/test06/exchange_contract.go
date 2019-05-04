// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820f04fa94e6629532bcb99324dd6a785a39f61a5e8d213bbd2dc539c36e4cd2bc60029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"orderid\",\"type\":\"uint64\"},{\"name\":\"useraddress\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"tradetype\",\"type\":\"bool\"},{\"name\":\"sellamount\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"buyaddress\",\"type\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"depositsOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restartContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"getMoney\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txid\",\"type\":\"uint64\"}],\"name\":\"withdrawal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderid\",\"type\":\"uint64\"},{\"name\":\"useraddress\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"tradetype\",\"type\":\"bool\"},{\"name\":\"sellamount\",\"type\":\"uint256\"},{\"name\":\"selladdress\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"depositsSell\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destoryContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractTypeAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"orderidToUsers\",\"outputs\":[{\"name\":\"order_id\",\"type\":\"uint64\"},{\"name\":\"user_address\",\"type\":\"address\"},{\"name\":\"user_receiver\",\"type\":\"string\"},{\"name\":\"trade_type\",\"type\":\"bool\"},{\"name\":\"sell_amount\",\"type\":\"uint256\"},{\"name\":\"sell_address\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"buy_amount\",\"type\":\"uint256\"},{\"name\":\"buy_address\",\"type\":\"address\"},{\"name\":\"fees\",\"type\":\"uint256\"},{\"name\":\"withdrawalable\",\"type\":\"bool\"},{\"name\":\"ex_success\",\"type\":\"bool\"},{\"name\":\"create_at\",\"type\":\"uint256\"},{\"name\":\"update_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txid\",\"type\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"sell\",\"type\":\"uint256\"},{\"name\":\"buy\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"tokenid\",\"type\":\"uint64\"},{\"name\":\"types\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint64\"},{\"name\":\"chnageType\",\"type\":\"bool\"}],\"name\":\"changeContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txid\",\"type\":\"uint64\"}],\"name\":\"upstate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenid\",\"type\":\"uint64\"}],\"name\":\"getTokenInfor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderid\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"tradetype\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"sellamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"selladdress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"DepositsSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderid\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"tradetype\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"sellamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"Deposits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txid\",\"type\":\"uint64\"}],\"name\":\"Withdrawl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txid\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sell\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Submit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"GetMoney\",\"type\":\"event\"}]"

// ExchangeBin is the compiled bytecode used for deploying new contracts.
const ExchangeBin = `0x60806040526000805460a060020a60ff02197fffffffffffffffffffff00ff000000000000000000000000000000000000000090911633171674010000000000000000000000000000000000000000179055611f4c806100606000396000f3006080604052600436106100cc5763ffffffff60e060020a60003504166327ab173281146100d157806330a52d5b1461016d5780633262fd9a14610184578063439766ce1461019c5780635fb02f4d146101b157806365f1d0ad146101c65780636f9fb98a146101db5780637a5216c414610202578063b568cfa01461028a578063c79f28951461029f578063cda33c6e146102c0578063ce606ee01461040b578063e78909c11461043c578063ee8e2ae414610466578063f20781d9146104ea578063f6891a2f1461050c575b600080fd5b604080516020600460443581810135601f810184900484028501840190955284845261015994823567ffffffffffffffff169460248035600160a060020a0316953695946064949201919081908401838280828437509497505050508235151593505050602081013590604081013590600160a060020a0360608201351690608001356105db565b604080519115158252519081900360200190f35b34801561017957600080fd5b50610182610918565b005b34801561019057600080fd5b5061018260043561096f565b3480156101a857600080fd5b506101826109f0565b3480156101bd57600080fd5b50610159610a45565b61015967ffffffffffffffff60043516610a55565b3480156101e757600080fd5b506101f0610edf565b60408051918252519081900360200190f35b604080516020600460443581810135601f810184900484028501840190955284845261015994823567ffffffffffffffff169460248035600160a060020a0316953695946064949201919081908401838280828437509497505050508235151593505050602081013590600160a060020a036040820135169060608101359060800135610f02565b34801561029657600080fd5b50610182611429565b3480156102ab57600080fd5b50610159600160a060020a0360043516611450565b3480156102cc57600080fd5b506102e267ffffffffffffffff60043516611465565b604051808f67ffffffffffffffff1667ffffffffffffffff1681526020018e600160a060020a0316600160a060020a03168152602001806020018d1515151581526020018c81526020018b600160a060020a0316600160a060020a031681526020018a815260200189815260200188600160a060020a0316600160a060020a03168152602001878152602001861515151581526020018515151515815260200184815260200183815260200182810382528e818151815260200191508051906020019080838360005b838110156103c35781810151838201526020016103ab565b50505050905090810190601f1680156103f05780820380516001836020036101000a031916815260200191505b509f5050505050505050505050505050505060405180910390f35b34801561041757600080fd5b5061042061158d565b60408051600160a060020a039092168252519081900360200190f35b61015967ffffffffffffffff60043516600160a060020a036024351660443560643560843561159c565b34801561047257600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610182948235600160a060020a0316946024803567ffffffffffffffff16953695946064949201919081908401838280828437509497505067ffffffffffffffff8535169550505050506020013515156119a5565b3480156104f657600080fd5b5061018267ffffffffffffffff60043516611b10565b34801561051857600080fd5b5061052e67ffffffffffffffff60043516611be4565b6040805167ffffffffffffffff8088168252600160a060020a038616928201929092529083166060820152811515608082015260a06020808301828152875192840192909252865160c084019188019080838360005b8381101561059c578181015183820152602001610584565b50505050905090810190601f1680156105c95780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b60006105e5611d78565b600160a060020a03891633146105fa57600080fd5b8751151561060757600080fd5b60005460a060020a900460ff16151560011461062257600080fd5b34861461062e57600080fd5b86151561067457600160a060020a038416151561064a57600080fd5b600160a060020a03841660009081526002602052604090205460ff16151560011461067457600080fd5b50604080516101c08101825267ffffffffffffffff8b8116808352600160a060020a038c811660208086019182528587018e81528d15156060880152608087018d9052600060a0880181905260c088018d905260e088018190528b8516610100890152610120880181905260016101408901819052610160890182905261018089018c90526101a089018c90529581528583529790972086518154935167ffffffffffffffff199094169616959095177fffffffff0000000000000000000000000000000000000000ffffffffffffffff1668010000000000000000929093169190910291909117835593518051939485946107769385019290910190611e13565b5060608201518160020160006101000a81548160ff0219169083151502179055506080820151816003015560a08201518160040160006101000a815481600160a060020a030219169083600160a060020a0316021790555060c0820151816005015560e082015181600601556101008201518160070160006101000a815481600160a060020a030219169083600160a060020a0316021790555061012082015181600801556101408201518160090160006101000a81548160ff0219169083151502179055506101608201518160090160016101000a81548160ff02191690831515021790555061018082015181600a01556101a082015181600b01559050507fb619533d6184a37cebded416c74cda54f801f2a4bfd06c175fa8afaf85740b258a88888860008989604051808867ffffffffffffffff1667ffffffffffffffff1681526020018715151515815260200186815260200185815260200184815260200183600160a060020a0316600160a060020a0316815260200182815260200197505050505050505060405180910390a15060019998505050505050505050565b6000543390600160a060020a0316811461093157600080fd5b60005460a060020a900460ff161561094857600080fd5b506000805474ff0000000000000000000000000000000000000000191660a060020a179055565b6000543390600160a060020a0316811461098857600080fd5b600082118015610999575030318211155b15156109a457600080fd5b6109ad82611d13565b6040805133815260208101849052428183015290517f34f457f9b40515dc7505fa1ca125d98bad483e73f905d02ab62a027f27cb292b9181900360600190a15050565b6000543390600160a060020a03168114610a0957600080fd5b60005460a060020a900460ff161515600114610a2457600080fd5b506000805474ff000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b60008054819060a060020a900460ff161515600114610a7357600080fd5b67ffffffffffffffff8316600090815260016020526040902054680100000000000000009004600160a060020a03163314610aad57600080fd5b5067ffffffffffffffff821660009081526001602081905260409091206003810154600990910154909160ff9091161515148015610b13575067ffffffffffffffff8316600090815260016020819052604090912060090154610100900460ff16151514155b8015610b1f5750600081115b1515610b2a57600080fd5b6000547501000000000000000000000000000000000000000000900460ff1615610b5357600080fd5b67ffffffffffffffff8316600090815260016020526040902060040154600160a060020a031615610deb576000805475ff0000000000000000000000000000000000000000001916750100000000000000000000000000000000000000000017815567ffffffffffffffff8416815260016020908152604080832060049081015482517fa9059cbb0000000000000000000000000000000000000000000000000000000081523392810192909252602482018690529151600160a060020a039092169363a9059cbb9360448084019491939192918390030190829087803b158015610c3d57600080fd5b505af1158015610c51573d6000803e3d6000fd5b505050506040513d6020811015610c6757600080fd5b50506000805475ff0000000000000000000000000000000000000000001916815567ffffffffffffffff84168152600160205260409020600601541515610d6b5767ffffffffffffffff83166000908152600160208190526040822080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681559190610cf490830182611e8d565b5060028101805460ff1916905560006003820181905560048201805473ffffffffffffffffffffffffffffffffffffffff19908116909155600583018290556006830182905560078301805490911690556008820181905560098201805461ffff19169055600a8201819055600b90910155610da5565b67ffffffffffffffff83166000908152600160205260408120600381019190915560098101805461ffff191661010017905542600b909101555b6040805167ffffffffffffffff8516815290517f0269114ad2350df1094ec1642497bb8df36a692522a7f0ddef9f9cd2142d91429181900360200190a160019150610ed9565b6000805475ff00000000000000000000000000000000000000000019167501000000000000000000000000000000000000000000178155604051339183156108fc02918491818181858888f19350505050158015610e4d573d6000803e3d6000fd5b506000805475ff0000000000000000000000000000000000000000001916815567ffffffffffffffff84168152600160205260409020600601541515610d6b5767ffffffffffffffff83166000908152600160208190526040822080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681559190610cf490830182611e8d565b50919050565b600080543390600160a060020a03168114610ef957600080fd5b303191505b5090565b6000610f0c611d78565b6000543390600160a060020a03168114610f2557600080fd5b60005460a060020a900460ff161515600114610f4057600080fd5b88511515610f4d57600080fd5b600188151514610f5c57600080fd5b600160a060020a0386161515610f7157600080fd5b600160a060020a03861660009081526002602052604090205460ff161515600114610f9b57600080fd5b604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a038c811660048301523060248301529151899289169163dd62ed3e9160448083019260209291908290030181600087803b15801561100657600080fd5b505af115801561101a573d6000803e3d6000fd5b505050506040513d602081101561103057600080fd5b5051101561103d57600080fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a038c81166004830152306024830152604482018a90529151918816916323b872dd916064808201926020929091908290030181600087803b1580156110af57600080fd5b505af11580156110c3573d6000803e3d6000fd5b505050506040513d60208110156110d957600080fd5b8101908080519060200190929190505050506101c0604051908101604052808c67ffffffffffffffff1681526020018b600160a060020a031681526020018a8152602001891515815260200188815260200187600160a060020a03168152602001868152602001600081526020016000600160a060020a031681526020016000815260200160011515815260200160001515815260200185815260200185815250915081600160008d67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a815481600160a060020a030219169083600160a060020a03160217905550604082015181600101908051906020019061121c929190611e13565b5060608201518160020160006101000a81548160ff0219169083151502179055506080820151816003015560a08201518160040160006101000a815481600160a060020a030219169083600160a060020a0316021790555060c0820151816005015560e082015181600601556101008201518160070160006101000a815481600160a060020a030219169083600160a060020a0316021790555061012082015181600801556101408201518160090160006101000a81548160ff0219169083151502179055506101608201518160090160016101000a81548160ff02191690831515021790555061018082015181600a01556101a082015181600b01559050507fbfec2afe29691471c0ef7c2539f3f323340ab54e7e34c6e6c6c1de48277c36928b8a8a8a8a8a60008b604051808967ffffffffffffffff1667ffffffffffffffff168152602001806020018815151515815260200187815260200186600160a060020a0316600160a060020a03168152602001858152602001848152602001838152602001828103825289818151815260200191508051906020019080838360005b838110156113d75781810151838201526020016113bf565b50505050905090810190601f1680156114045780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a15060019a9950505050505050505050565b6000543390600160a060020a0316811461144257600080fd5b600054600160a060020a0316ff5b60026020526000908152604090205460ff1681565b600160208181526000928352604092839020805481840180548651600261010097831615979097026000190190911695909504601f810185900485028601850190965285855267ffffffffffffffff82169568010000000000000000909204600160a060020a03169492939091908301828280156115245780601f106114f957610100808354040283529160200191611524565b820191906000526020600020905b81548152906001019060200180831161150757829003601f168201915b50505050600283015460038401546004850154600586015460068701546007880154600889015460098a0154600a8b0154600b909b0154999a60ff9889169a979950600160a060020a03968716989597949690931694919382821693610100909204909216918e565b600054600160a060020a031681565b6000805481903390600160a060020a031681146115b857600080fd5b60005460a060020a900460ff1615156001146115d357600080fd5b6115e3868563ffffffff611d5016565b67ffffffffffffffff89166000908152600160205260409020600301541080159061160e5750600086115b801561161a5750600084115b80156116265750600085115b151561163157600080fd5b67ffffffffffffffff881660009081526001602052604090206003810154600490910154909250600160a060020a0316156117dd5767ffffffffffffffff8816600090815260016020908152604080832060049081015482517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038d811693820193909352602481018c9052925191169363a9059cbb93604480850194919392918390030190829087803b1580156116f157600080fd5b505af1158015611705573d6000803e3d6000fd5b505050506040513d602081101561171b57600080fd5b505067ffffffffffffffff88166000908152600160209081526040808320600490810154845483517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a0391821693810193909352602483018a9052925192169363a9059cbb9360448084019491939192918390030190829087803b1580156117ab57600080fd5b505af11580156117bf573d6000803e3d6000fd5b505050506040513d60208110156117d557600080fd5b506118159050565b604051600160a060020a0388169087156108fc029088906000818181858888f19350505050158015611813573d6000803e3d6000fd5b505b611835611828878663ffffffff611d5016565b839063ffffffff611d6616565b67ffffffffffffffff8916600090815260016020526040902060038101919091556008015461186a908563ffffffff611d5016565b67ffffffffffffffff8916600090815260016020526040902060088101919091556006015461189f908663ffffffff611d5016565b67ffffffffffffffff89166000908152600160205260409020600681019190915542600b909101556118d7868563ffffffff611d5016565b82141561190c5767ffffffffffffffff88166000908152600160205260409020600901805461ffff1916610100179055611937565b67ffffffffffffffff88166000908152600160208190526040909120600901805460ff191690911790555b6040805167ffffffffffffffff8a168152600160a060020a0389166020820152808201889052606081018790526080810186905290517f4cb073b1936b49cc564cd29a6510c77130225593415dbd914defbceb27d417da9181900360a00190a1506001979650505050505050565b6119ad611ed4565b6000543390600160a060020a031681146119c657600080fd5b600160a060020a0387166000818152600260209081526040808320805460ff1916881515908117909155815160a08101835267ffffffffffffffff8c81168083528286018d8152838601989098528b82166060840152608083019390935291855260038452919093208151815467ffffffffffffffff19169416939093178355925180519395508593611a5f9260018501920190611e13565b506040820151600290910180546060840151608090940151151560e060020a027fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff67ffffffffffffffff90951660a060020a027fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff600160a060020a0390951673ffffffffffffffffffffffffffffffffffffffff199093169290921793909316179290921617905550505050505050565b600080543390600160a060020a03168114611b2a57600080fd5b60005460a060020a900460ff161515600114611b4557600080fd5b67ffffffffffffffff831660009081526001602081905260409091206009015460ff80821694506101009091041615151415611b8057600080fd5b60018215151415611bb45767ffffffffffffffff83166000908152600160205260409020600901805460ff19169055611bdf565b67ffffffffffffffff83166000908152600160208190526040909120600901805460ff191690911790555b505050565b60008054606090829081908190819081908190819081903390600160a060020a03168114611c1157600080fd5b67ffffffffffffffff8c811660009081526003602090815260409182902080546002808301546001938401805487516101009682161596909602600019011692909204601f81018690048602850186019096528584529186169b509950600160a060020a038116985060a060020a8104909416965060e060020a90930460ff169450889288928892889288929091869190830182828015611cf35780601f10611cc857610100808354040283529160200191611cf3565b820191906000526020600020905b815481529060010190602001808311611cd657829003601f168201915b505050505093509a509a509a509a509a5050505050505091939590929450565b60008054604051600160a060020a039091169183156108fc02918491818181858888f19350505050158015611d4c573d6000803e3d6000fd5b5050565b600082820183811015611d5f57fe5b9392505050565b600082821115611d7257fe5b50900390565b6101c060405190810160405280600067ffffffffffffffff1681526020016000600160a060020a0316815260200160608152602001600015158152602001600081526020016000600160a060020a0316815260200160008152602001600081526020016000600160a060020a031681526020016000815260200160001515815260200160001515815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611e5457805160ff1916838001178555611e81565b82800160010185558215611e81579182015b82811115611e81578251825591602001919060010190611e66565b50610efe929150611f03565b50805460018160011615610100020316600290046000825580601f10611eb35750611ed1565b601f016020900490600052602060002090810190611ed19190611f03565b50565b6040805160a0810182526000808252606060208301819052928201819052918101829052608081019190915290565b611f1d91905b80821115610efe5760008155600101611f09565b905600a165627a7a723058209695d4b362e624a56394070cf5c7c128ea5106abe3f1044dd13b5558fd4c0f9e0029`

// DeployExchange deploys a new Ethereum contract, binding an instance of Exchange to it.
func DeployExchange(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exchange, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() constant returns(address)
func (_Exchange *ExchangeCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "contractOwner")
	return *ret0, err
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() constant returns(address)
func (_Exchange *ExchangeSession) ContractOwner() (common.Address, error) {
	return _Exchange.Contract.ContractOwner(&_Exchange.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() constant returns(address)
func (_Exchange *ExchangeCallerSession) ContractOwner() (common.Address, error) {
	return _Exchange.Contract.ContractOwner(&_Exchange.CallOpts)
}

// ContractTypeAddress is a free data retrieval call binding the contract method 0xc79f2895.
//
// Solidity: function contractTypeAddress(address ) constant returns(bool)
func (_Exchange *ExchangeCaller) ContractTypeAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "contractTypeAddress", arg0)
	return *ret0, err
}

// ContractTypeAddress is a free data retrieval call binding the contract method 0xc79f2895.
//
// Solidity: function contractTypeAddress(address ) constant returns(bool)
func (_Exchange *ExchangeSession) ContractTypeAddress(arg0 common.Address) (bool, error) {
	return _Exchange.Contract.ContractTypeAddress(&_Exchange.CallOpts, arg0)
}

// ContractTypeAddress is a free data retrieval call binding the contract method 0xc79f2895.
//
// Solidity: function contractTypeAddress(address ) constant returns(bool)
func (_Exchange *ExchangeCallerSession) ContractTypeAddress(arg0 common.Address) (bool, error) {
	return _Exchange.Contract.ContractTypeAddress(&_Exchange.CallOpts, arg0)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() constant returns(uint256)
func (_Exchange *ExchangeCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getContractBalance")
	return *ret0, err
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() constant returns(uint256)
func (_Exchange *ExchangeSession) GetContractBalance() (*big.Int, error) {
	return _Exchange.Contract.GetContractBalance(&_Exchange.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() constant returns(uint256)
func (_Exchange *ExchangeCallerSession) GetContractBalance() (*big.Int, error) {
	return _Exchange.Contract.GetContractBalance(&_Exchange.CallOpts)
}

// GetTokenInfor is a free data retrieval call binding the contract method 0xf6891a2f.
//
// Solidity: function getTokenInfor(uint64 tokenid) constant returns(uint64, string, address, uint64, bool)
func (_Exchange *ExchangeCaller) GetTokenInfor(opts *bind.CallOpts, tokenid uint64) (uint64, string, common.Address, uint64, bool, error) {
	var (
		ret0 = new(uint64)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new(uint64)
		ret4 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Exchange.contract.Call(opts, out, "getTokenInfor", tokenid)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetTokenInfor is a free data retrieval call binding the contract method 0xf6891a2f.
//
// Solidity: function getTokenInfor(uint64 tokenid) constant returns(uint64, string, address, uint64, bool)
func (_Exchange *ExchangeSession) GetTokenInfor(tokenid uint64) (uint64, string, common.Address, uint64, bool, error) {
	return _Exchange.Contract.GetTokenInfor(&_Exchange.CallOpts, tokenid)
}

// GetTokenInfor is a free data retrieval call binding the contract method 0xf6891a2f.
//
// Solidity: function getTokenInfor(uint64 tokenid) constant returns(uint64, string, address, uint64, bool)
func (_Exchange *ExchangeCallerSession) GetTokenInfor(tokenid uint64) (uint64, string, common.Address, uint64, bool, error) {
	return _Exchange.Contract.GetTokenInfor(&_Exchange.CallOpts, tokenid)
}

// OrderidToUsers is a free data retrieval call binding the contract method 0xcda33c6e.
//
// Solidity: function orderidToUsers(uint64 ) constant returns(uint64 order_id, address user_address, string user_receiver, bool trade_type, uint256 sell_amount, address sell_address, uint256 price, uint256 buy_amount, address buy_address, uint256 fees, bool withdrawalable, bool ex_success, uint256 create_at, uint256 update_at)
func (_Exchange *ExchangeCaller) OrderidToUsers(opts *bind.CallOpts, arg0 uint64) (struct {
	OrderId        uint64
	UserAddress    common.Address
	UserReceiver   string
	TradeType      bool
	SellAmount     *big.Int
	SellAddress    common.Address
	Price          *big.Int
	BuyAmount      *big.Int
	BuyAddress     common.Address
	Fees           *big.Int
	Withdrawalable bool
	ExSuccess      bool
	CreateAt       *big.Int
	UpdateAt       *big.Int
}, error) {
	ret := new(struct {
		OrderId        uint64
		UserAddress    common.Address
		UserReceiver   string
		TradeType      bool
		SellAmount     *big.Int
		SellAddress    common.Address
		Price          *big.Int
		BuyAmount      *big.Int
		BuyAddress     common.Address
		Fees           *big.Int
		Withdrawalable bool
		ExSuccess      bool
		CreateAt       *big.Int
		UpdateAt       *big.Int
	})
	out := ret
	err := _Exchange.contract.Call(opts, out, "orderidToUsers", arg0)
	return *ret, err
}

// OrderidToUsers is a free data retrieval call binding the contract method 0xcda33c6e.
//
// Solidity: function orderidToUsers(uint64 ) constant returns(uint64 order_id, address user_address, string user_receiver, bool trade_type, uint256 sell_amount, address sell_address, uint256 price, uint256 buy_amount, address buy_address, uint256 fees, bool withdrawalable, bool ex_success, uint256 create_at, uint256 update_at)
func (_Exchange *ExchangeSession) OrderidToUsers(arg0 uint64) (struct {
	OrderId        uint64
	UserAddress    common.Address
	UserReceiver   string
	TradeType      bool
	SellAmount     *big.Int
	SellAddress    common.Address
	Price          *big.Int
	BuyAmount      *big.Int
	BuyAddress     common.Address
	Fees           *big.Int
	Withdrawalable bool
	ExSuccess      bool
	CreateAt       *big.Int
	UpdateAt       *big.Int
}, error) {
	return _Exchange.Contract.OrderidToUsers(&_Exchange.CallOpts, arg0)
}

// OrderidToUsers is a free data retrieval call binding the contract method 0xcda33c6e.
//
// Solidity: function orderidToUsers(uint64 ) constant returns(uint64 order_id, address user_address, string user_receiver, bool trade_type, uint256 sell_amount, address sell_address, uint256 price, uint256 buy_amount, address buy_address, uint256 fees, bool withdrawalable, bool ex_success, uint256 create_at, uint256 update_at)
func (_Exchange *ExchangeCallerSession) OrderidToUsers(arg0 uint64) (struct {
	OrderId        uint64
	UserAddress    common.Address
	UserReceiver   string
	TradeType      bool
	SellAmount     *big.Int
	SellAddress    common.Address
	Price          *big.Int
	BuyAmount      *big.Int
	BuyAddress     common.Address
	Fees           *big.Int
	Withdrawalable bool
	ExSuccess      bool
	CreateAt       *big.Int
	UpdateAt       *big.Int
}, error) {
	return _Exchange.Contract.OrderidToUsers(&_Exchange.CallOpts, arg0)
}

// StartContract is a free data retrieval call binding the contract method 0x5fb02f4d.
//
// Solidity: function startContract() constant returns(bool)
func (_Exchange *ExchangeCaller) StartContract(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "startContract")
	return *ret0, err
}

// StartContract is a free data retrieval call binding the contract method 0x5fb02f4d.
//
// Solidity: function startContract() constant returns(bool)
func (_Exchange *ExchangeSession) StartContract() (bool, error) {
	return _Exchange.Contract.StartContract(&_Exchange.CallOpts)
}

// StartContract is a free data retrieval call binding the contract method 0x5fb02f4d.
//
// Solidity: function startContract() constant returns(bool)
func (_Exchange *ExchangeCallerSession) StartContract() (bool, error) {
	return _Exchange.Contract.StartContract(&_Exchange.CallOpts)
}

// ChangeContractAddress is a paid mutator transaction binding the contract method 0xee8e2ae4.
//
// Solidity: function changeContractAddress(address addr, uint64 tokenid, string types, uint64 decimals, bool chnageType) returns()
func (_Exchange *ExchangeTransactor) ChangeContractAddress(opts *bind.TransactOpts, addr common.Address, tokenid uint64, types string, decimals uint64, chnageType bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "changeContractAddress", addr, tokenid, types, decimals, chnageType)
}

// ChangeContractAddress is a paid mutator transaction binding the contract method 0xee8e2ae4.
//
// Solidity: function changeContractAddress(address addr, uint64 tokenid, string types, uint64 decimals, bool chnageType) returns()
func (_Exchange *ExchangeSession) ChangeContractAddress(addr common.Address, tokenid uint64, types string, decimals uint64, chnageType bool) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeContractAddress(&_Exchange.TransactOpts, addr, tokenid, types, decimals, chnageType)
}

// ChangeContractAddress is a paid mutator transaction binding the contract method 0xee8e2ae4.
//
// Solidity: function changeContractAddress(address addr, uint64 tokenid, string types, uint64 decimals, bool chnageType) returns()
func (_Exchange *ExchangeTransactorSession) ChangeContractAddress(addr common.Address, tokenid uint64, types string, decimals uint64, chnageType bool) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeContractAddress(&_Exchange.TransactOpts, addr, tokenid, types, decimals, chnageType)
}

// DepositsOrder is a paid mutator transaction binding the contract method 0x27ab1732.
//
// Solidity: function depositsOrder(uint64 orderid, address useraddress, string receiver, bool tradetype, uint256 sellamount, uint256 price, address buyaddress, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactor) DepositsOrder(opts *bind.TransactOpts, orderid uint64, useraddress common.Address, receiver string, tradetype bool, sellamount *big.Int, price *big.Int, buyaddress common.Address, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositsOrder", orderid, useraddress, receiver, tradetype, sellamount, price, buyaddress, createTime)
}

// DepositsOrder is a paid mutator transaction binding the contract method 0x27ab1732.
//
// Solidity: function depositsOrder(uint64 orderid, address useraddress, string receiver, bool tradetype, uint256 sellamount, uint256 price, address buyaddress, uint256 createTime) returns(bool)
func (_Exchange *ExchangeSession) DepositsOrder(orderid uint64, useraddress common.Address, receiver string, tradetype bool, sellamount *big.Int, price *big.Int, buyaddress common.Address, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsOrder(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, price, buyaddress, createTime)
}

// DepositsOrder is a paid mutator transaction binding the contract method 0x27ab1732.
//
// Solidity: function depositsOrder(uint64 orderid, address useraddress, string receiver, bool tradetype, uint256 sellamount, uint256 price, address buyaddress, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactorSession) DepositsOrder(orderid uint64, useraddress common.Address, receiver string, tradetype bool, sellamount *big.Int, price *big.Int, buyaddress common.Address, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsOrder(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, price, buyaddress, createTime)
}

// DepositsSell is a paid mutator transaction binding the contract method 0x7a5216c4.
//
// Solidity: function depositsSell(uint64 orderid, address useraddress, string receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactor) DepositsSell(opts *bind.TransactOpts, orderid uint64, useraddress common.Address, receiver string, tradetype bool, sellamount *big.Int, selladdress common.Address, price *big.Int, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositsSell", orderid, useraddress, receiver, tradetype, sellamount, selladdress, price, createTime)
}

// DepositsSell is a paid mutator transaction binding the contract method 0x7a5216c4.
//
// Solidity: function depositsSell(uint64 orderid, address useraddress, string receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 createTime) returns(bool)
func (_Exchange *ExchangeSession) DepositsSell(orderid uint64, useraddress common.Address, receiver string, tradetype bool, sellamount *big.Int, selladdress common.Address, price *big.Int, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsSell(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, selladdress, price, createTime)
}

// DepositsSell is a paid mutator transaction binding the contract method 0x7a5216c4.
//
// Solidity: function depositsSell(uint64 orderid, address useraddress, string receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactorSession) DepositsSell(orderid uint64, useraddress common.Address, receiver string, tradetype bool, sellamount *big.Int, selladdress common.Address, price *big.Int, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsSell(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, selladdress, price, createTime)
}

// DestoryContract is a paid mutator transaction binding the contract method 0xb568cfa0.
//
// Solidity: function destoryContract() returns()
func (_Exchange *ExchangeTransactor) DestoryContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "destoryContract")
}

// DestoryContract is a paid mutator transaction binding the contract method 0xb568cfa0.
//
// Solidity: function destoryContract() returns()
func (_Exchange *ExchangeSession) DestoryContract() (*types.Transaction, error) {
	return _Exchange.Contract.DestoryContract(&_Exchange.TransactOpts)
}

// DestoryContract is a paid mutator transaction binding the contract method 0xb568cfa0.
//
// Solidity: function destoryContract() returns()
func (_Exchange *ExchangeTransactorSession) DestoryContract() (*types.Transaction, error) {
	return _Exchange.Contract.DestoryContract(&_Exchange.TransactOpts)
}

// GetMoney is a paid mutator transaction binding the contract method 0x3262fd9a.
//
// Solidity: function getMoney(uint256 money) returns()
func (_Exchange *ExchangeTransactor) GetMoney(opts *bind.TransactOpts, money *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "getMoney", money)
}

// GetMoney is a paid mutator transaction binding the contract method 0x3262fd9a.
//
// Solidity: function getMoney(uint256 money) returns()
func (_Exchange *ExchangeSession) GetMoney(money *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.GetMoney(&_Exchange.TransactOpts, money)
}

// GetMoney is a paid mutator transaction binding the contract method 0x3262fd9a.
//
// Solidity: function getMoney(uint256 money) returns()
func (_Exchange *ExchangeTransactorSession) GetMoney(money *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.GetMoney(&_Exchange.TransactOpts, money)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Exchange *ExchangeTransactor) PauseContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "pauseContract")
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Exchange *ExchangeSession) PauseContract() (*types.Transaction, error) {
	return _Exchange.Contract.PauseContract(&_Exchange.TransactOpts)
}

// PauseContract is a paid mutator transaction binding the contract method 0x439766ce.
//
// Solidity: function pauseContract() returns()
func (_Exchange *ExchangeTransactorSession) PauseContract() (*types.Transaction, error) {
	return _Exchange.Contract.PauseContract(&_Exchange.TransactOpts)
}

// RestartContract is a paid mutator transaction binding the contract method 0x30a52d5b.
//
// Solidity: function restartContract() returns()
func (_Exchange *ExchangeTransactor) RestartContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "restartContract")
}

// RestartContract is a paid mutator transaction binding the contract method 0x30a52d5b.
//
// Solidity: function restartContract() returns()
func (_Exchange *ExchangeSession) RestartContract() (*types.Transaction, error) {
	return _Exchange.Contract.RestartContract(&_Exchange.TransactOpts)
}

// RestartContract is a paid mutator transaction binding the contract method 0x30a52d5b.
//
// Solidity: function restartContract() returns()
func (_Exchange *ExchangeTransactorSession) RestartContract() (*types.Transaction, error) {
	return _Exchange.Contract.RestartContract(&_Exchange.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xe78909c1.
//
// Solidity: function submit(uint64 txid, address to, uint256 sell, uint256 buy, uint256 fee) returns(bool)
func (_Exchange *ExchangeTransactor) Submit(opts *bind.TransactOpts, txid uint64, to common.Address, sell *big.Int, buy *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "submit", txid, to, sell, buy, fee)
}

// Submit is a paid mutator transaction binding the contract method 0xe78909c1.
//
// Solidity: function submit(uint64 txid, address to, uint256 sell, uint256 buy, uint256 fee) returns(bool)
func (_Exchange *ExchangeSession) Submit(txid uint64, to common.Address, sell *big.Int, buy *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Submit(&_Exchange.TransactOpts, txid, to, sell, buy, fee)
}

// Submit is a paid mutator transaction binding the contract method 0xe78909c1.
//
// Solidity: function submit(uint64 txid, address to, uint256 sell, uint256 buy, uint256 fee) returns(bool)
func (_Exchange *ExchangeTransactorSession) Submit(txid uint64, to common.Address, sell *big.Int, buy *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Submit(&_Exchange.TransactOpts, txid, to, sell, buy, fee)
}

// Upstate is a paid mutator transaction binding the contract method 0xf20781d9.
//
// Solidity: function upstate(uint64 txid) returns()
func (_Exchange *ExchangeTransactor) Upstate(opts *bind.TransactOpts, txid uint64) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "upstate", txid)
}

// Upstate is a paid mutator transaction binding the contract method 0xf20781d9.
//
// Solidity: function upstate(uint64 txid) returns()
func (_Exchange *ExchangeSession) Upstate(txid uint64) (*types.Transaction, error) {
	return _Exchange.Contract.Upstate(&_Exchange.TransactOpts, txid)
}

// Upstate is a paid mutator transaction binding the contract method 0xf20781d9.
//
// Solidity: function upstate(uint64 txid) returns()
func (_Exchange *ExchangeTransactorSession) Upstate(txid uint64) (*types.Transaction, error) {
	return _Exchange.Contract.Upstate(&_Exchange.TransactOpts, txid)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x65f1d0ad.
//
// Solidity: function withdrawal(uint64 txid) returns(bool)
func (_Exchange *ExchangeTransactor) Withdrawal(opts *bind.TransactOpts, txid uint64) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdrawal", txid)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x65f1d0ad.
//
// Solidity: function withdrawal(uint64 txid) returns(bool)
func (_Exchange *ExchangeSession) Withdrawal(txid uint64) (*types.Transaction, error) {
	return _Exchange.Contract.Withdrawal(&_Exchange.TransactOpts, txid)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x65f1d0ad.
//
// Solidity: function withdrawal(uint64 txid) returns(bool)
func (_Exchange *ExchangeTransactorSession) Withdrawal(txid uint64) (*types.Transaction, error) {
	return _Exchange.Contract.Withdrawal(&_Exchange.TransactOpts, txid)
}

// ExchangeDepositsIterator is returned from FilterDeposits and is used to iterate over the raw logs and unpacked data for Deposits events raised by the Exchange contract.
type ExchangeDepositsIterator struct {
	Event *ExchangeDeposits // Event containing the contract specifics and raw log

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
func (it *ExchangeDepositsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeDeposits)
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
		it.Event = new(ExchangeDeposits)
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
func (it *ExchangeDepositsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeDepositsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeDeposits represents a Deposits event raised by the Exchange contract.
type ExchangeDeposits struct {
	Orderid    uint64
	Tradetype  bool
	Sellamount *big.Int
	Price      *big.Int
	Buyamount  *big.Int
	Buyaddress common.Address
	CreateTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposits is a free log retrieval operation binding the contract event 0xb619533d6184a37cebded416c74cda54f801f2a4bfd06c175fa8afaf85740b25.
//
// Solidity: event Deposits(uint64 orderid, bool tradetype, uint256 sellamount, uint256 price, uint256 buyamount, address buyaddress, uint256 createTime)
func (_Exchange *ExchangeFilterer) FilterDeposits(opts *bind.FilterOpts) (*ExchangeDepositsIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Deposits")
	if err != nil {
		return nil, err
	}
	return &ExchangeDepositsIterator{contract: _Exchange.contract, event: "Deposits", logs: logs, sub: sub}, nil
}

// WatchDeposits is a free log subscription operation binding the contract event 0xb619533d6184a37cebded416c74cda54f801f2a4bfd06c175fa8afaf85740b25.
//
// Solidity: event Deposits(uint64 orderid, bool tradetype, uint256 sellamount, uint256 price, uint256 buyamount, address buyaddress, uint256 createTime)
func (_Exchange *ExchangeFilterer) WatchDeposits(opts *bind.WatchOpts, sink chan<- *ExchangeDeposits) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Deposits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeDeposits)
				if err := _Exchange.contract.UnpackLog(event, "Deposits", log); err != nil {
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

// ExchangeDepositsSellIterator is returned from FilterDepositsSell and is used to iterate over the raw logs and unpacked data for DepositsSell events raised by the Exchange contract.
type ExchangeDepositsSellIterator struct {
	Event *ExchangeDepositsSell // Event containing the contract specifics and raw log

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
func (it *ExchangeDepositsSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeDepositsSell)
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
		it.Event = new(ExchangeDepositsSell)
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
func (it *ExchangeDepositsSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeDepositsSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeDepositsSell represents a DepositsSell event raised by the Exchange contract.
type ExchangeDepositsSell struct {
	Orderid     uint64
	Receiver    string
	Tradetype   bool
	Sellamount  *big.Int
	Selladdress common.Address
	Price       *big.Int
	Buyamount   *big.Int
	CreateTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositsSell is a free log retrieval operation binding the contract event 0xbfec2afe29691471c0ef7c2539f3f323340ab54e7e34c6e6c6c1de48277c3692.
//
// Solidity: event DepositsSell(uint64 orderid, string receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 buyamount, uint256 createTime)
func (_Exchange *ExchangeFilterer) FilterDepositsSell(opts *bind.FilterOpts) (*ExchangeDepositsSellIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "DepositsSell")
	if err != nil {
		return nil, err
	}
	return &ExchangeDepositsSellIterator{contract: _Exchange.contract, event: "DepositsSell", logs: logs, sub: sub}, nil
}

// WatchDepositsSell is a free log subscription operation binding the contract event 0xbfec2afe29691471c0ef7c2539f3f323340ab54e7e34c6e6c6c1de48277c3692.
//
// Solidity: event DepositsSell(uint64 orderid, string receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 buyamount, uint256 createTime)
func (_Exchange *ExchangeFilterer) WatchDepositsSell(opts *bind.WatchOpts, sink chan<- *ExchangeDepositsSell) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "DepositsSell")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeDepositsSell)
				if err := _Exchange.contract.UnpackLog(event, "DepositsSell", log); err != nil {
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

// ExchangeGetMoneyIterator is returned from FilterGetMoney and is used to iterate over the raw logs and unpacked data for GetMoney events raised by the Exchange contract.
type ExchangeGetMoneyIterator struct {
	Event *ExchangeGetMoney // Event containing the contract specifics and raw log

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
func (it *ExchangeGetMoneyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeGetMoney)
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
		it.Event = new(ExchangeGetMoney)
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
func (it *ExchangeGetMoneyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeGetMoneyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeGetMoney represents a GetMoney event raised by the Exchange contract.
type ExchangeGetMoney struct {
	Owner common.Address
	Money *big.Int
	Time  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGetMoney is a free log retrieval operation binding the contract event 0x34f457f9b40515dc7505fa1ca125d98bad483e73f905d02ab62a027f27cb292b.
//
// Solidity: event GetMoney(address owner, uint256 money, uint256 time)
func (_Exchange *ExchangeFilterer) FilterGetMoney(opts *bind.FilterOpts) (*ExchangeGetMoneyIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "GetMoney")
	if err != nil {
		return nil, err
	}
	return &ExchangeGetMoneyIterator{contract: _Exchange.contract, event: "GetMoney", logs: logs, sub: sub}, nil
}

// WatchGetMoney is a free log subscription operation binding the contract event 0x34f457f9b40515dc7505fa1ca125d98bad483e73f905d02ab62a027f27cb292b.
//
// Solidity: event GetMoney(address owner, uint256 money, uint256 time)
func (_Exchange *ExchangeFilterer) WatchGetMoney(opts *bind.WatchOpts, sink chan<- *ExchangeGetMoney) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "GetMoney")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeGetMoney)
				if err := _Exchange.contract.UnpackLog(event, "GetMoney", log); err != nil {
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

// ExchangeSubmitIterator is returned from FilterSubmit and is used to iterate over the raw logs and unpacked data for Submit events raised by the Exchange contract.
type ExchangeSubmitIterator struct {
	Event *ExchangeSubmit // Event containing the contract specifics and raw log

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
func (it *ExchangeSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeSubmit)
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
		it.Event = new(ExchangeSubmit)
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
func (it *ExchangeSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeSubmit represents a Submit event raised by the Exchange contract.
type ExchangeSubmit struct {
	Txid uint64
	To   common.Address
	Sell *big.Int
	Buy  *big.Int
	Fee  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSubmit is a free log retrieval operation binding the contract event 0x4cb073b1936b49cc564cd29a6510c77130225593415dbd914defbceb27d417da.
//
// Solidity: event Submit(uint64 txid, address to, uint256 sell, uint256 buy, uint256 fee)
func (_Exchange *ExchangeFilterer) FilterSubmit(opts *bind.FilterOpts) (*ExchangeSubmitIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Submit")
	if err != nil {
		return nil, err
	}
	return &ExchangeSubmitIterator{contract: _Exchange.contract, event: "Submit", logs: logs, sub: sub}, nil
}

// WatchSubmit is a free log subscription operation binding the contract event 0x4cb073b1936b49cc564cd29a6510c77130225593415dbd914defbceb27d417da.
//
// Solidity: event Submit(uint64 txid, address to, uint256 sell, uint256 buy, uint256 fee)
func (_Exchange *ExchangeFilterer) WatchSubmit(opts *bind.WatchOpts, sink chan<- *ExchangeSubmit) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Submit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeSubmit)
				if err := _Exchange.contract.UnpackLog(event, "Submit", log); err != nil {
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

// ExchangeWithdrawlIterator is returned from FilterWithdrawl and is used to iterate over the raw logs and unpacked data for Withdrawl events raised by the Exchange contract.
type ExchangeWithdrawlIterator struct {
	Event *ExchangeWithdrawl // Event containing the contract specifics and raw log

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
func (it *ExchangeWithdrawlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeWithdrawl)
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
		it.Event = new(ExchangeWithdrawl)
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
func (it *ExchangeWithdrawlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeWithdrawlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeWithdrawl represents a Withdrawl event raised by the Exchange contract.
type ExchangeWithdrawl struct {
	Txid uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawl is a free log retrieval operation binding the contract event 0x0269114ad2350df1094ec1642497bb8df36a692522a7f0ddef9f9cd2142d9142.
//
// Solidity: event Withdrawl(uint64 txid)
func (_Exchange *ExchangeFilterer) FilterWithdrawl(opts *bind.FilterOpts) (*ExchangeWithdrawlIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Withdrawl")
	if err != nil {
		return nil, err
	}
	return &ExchangeWithdrawlIterator{contract: _Exchange.contract, event: "Withdrawl", logs: logs, sub: sub}, nil
}

// WatchWithdrawl is a free log subscription operation binding the contract event 0x0269114ad2350df1094ec1642497bb8df36a692522a7f0ddef9f9cd2142d9142.
//
// Solidity: event Withdrawl(uint64 txid)
func (_Exchange *ExchangeFilterer) WatchWithdrawl(opts *bind.WatchOpts, sink chan<- *ExchangeWithdrawl) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Withdrawl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeWithdrawl)
				if err := _Exchange.contract.UnpackLog(event, "Withdrawl", log); err != nil {
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
