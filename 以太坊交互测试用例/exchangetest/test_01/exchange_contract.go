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
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582048d505b8a6b89eb6e3e65e566831f606190cb1fa42f2ea64c98a2f5f2a57f5f30029`

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
const ExchangeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"orderid\",\"type\":\"uint64\"},{\"name\":\"useraddress\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"tradetype\",\"type\":\"bool\"},{\"name\":\"sellamount\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"buyaddress\",\"type\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"depositsSell\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restartContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"money\",\"type\":\"uint256\"}],\"name\":\"getMoney\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txid\",\"type\":\"uint64\"}],\"name\":\"withdrawal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractAddressToType\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destoryContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"contractTypeAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"types\",\"type\":\"string\"},{\"name\":\"chnageType\",\"type\":\"bool\"}],\"name\":\"changeContractAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"orderidToUsers\",\"outputs\":[{\"name\":\"order_id\",\"type\":\"uint64\"},{\"name\":\"user_address\",\"type\":\"address\"},{\"name\":\"user_receiver\",\"type\":\"address\"},{\"name\":\"trade_type\",\"type\":\"bool\"},{\"name\":\"sell_amount\",\"type\":\"uint256\"},{\"name\":\"sell_address\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"buy_amount\",\"type\":\"uint256\"},{\"name\":\"buy_address\",\"type\":\"address\"},{\"name\":\"fees\",\"type\":\"uint256\"},{\"name\":\"withdrawalable\",\"type\":\"bool\"},{\"name\":\"ex_success\",\"type\":\"bool\"},{\"name\":\"create_at\",\"type\":\"uint256\"},{\"name\":\"update_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txid\",\"type\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"sell\",\"type\":\"uint256\"},{\"name\":\"buy\",\"type\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txid\",\"type\":\"uint64\"}],\"name\":\"upstate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderid\",\"type\":\"uint64\"},{\"name\":\"useraddress\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"tradetype\",\"type\":\"bool\"},{\"name\":\"sellamount\",\"type\":\"uint256\"},{\"name\":\"selladdress\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"depositsBuy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderid\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tradetype\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"sellamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"selladdress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"DepositsBuy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"orderid\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"tradetype\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"sellamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyamount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buyaddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"createTime\",\"type\":\"uint256\"}],\"name\":\"DepositsSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txid\",\"type\":\"uint64\"}],\"name\":\"Withdrawl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"txid\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sell\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"buy\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Submit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"money\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"GetMoney\",\"type\":\"event\"}]"

// ExchangeBin is the compiled bytecode used for deploying new contracts.
const ExchangeBin = `0x60806040526000805460a060020a60ff02197fffffffffffffffffffff00ff000000000000000000000000000000000000000090911633171674010000000000000000000000000000000000000000179055611c4f806100606000396000f3006080604052600436106100e55763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632cda39d981146100ea57806330a52d5b1461013c5780633262fd9a14610153578063439766ce1461016b5780635fb02f4d1461018057806365f1d0ad146101955780636f9fb98a146101aa578063958e2bcb146101d1578063b568cfa014610267578063c79f28951461027c578063cb037adc1461029d578063cda33c6e14610308578063ce606ee014610403578063e78909c114610434578063f20781d91461045e578063f4b06dda14610480575b600080fd5b61012867ffffffffffffffff60043516600160a060020a036024358116906044358116906064351515906084359060a4359060c4351660e4356104bd565b604080519115158252519081900360200190f35b34801561014857600080fd5b50610151610837565b005b34801561015f57600080fd5b5061015160043561088e565b34801561017757600080fd5b5061015161090f565b34801561018c57600080fd5b50610128610964565b61012867ffffffffffffffff60043516610974565b3480156101b657600080fd5b506101bf610e77565b60408051918252519081900360200190f35b3480156101dd57600080fd5b506101f2600160a060020a0360043516610e9a565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561022c578181015183820152602001610214565b50505050905090810190601f1680156102595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561027357600080fd5b50610151610f35565b34801561028857600080fd5b50610128600160a060020a0360043516610f5c565b3480156102a957600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610151958335600160a060020a0316953695604494919390910191908190840183828082843750949750505050913515159250610f71915050565b34801561031457600080fd5b5061032a67ffffffffffffffff6004351661100f565b604051808f67ffffffffffffffff1667ffffffffffffffff1681526020018e600160a060020a0316600160a060020a031681526020018d600160a060020a0316600160a060020a031681526020018c1515151581526020018b81526020018a600160a060020a0316600160a060020a0316815260200189815260200188815260200187600160a060020a0316600160a060020a0316815260200186815260200185151515158152602001841515151581526020018381526020018281526020019e50505050505050505050505050505060405180910390f35b34801561040f57600080fd5b506104186110a5565b60408051600160a060020a039092168252519081900360200190f35b61012867ffffffffffffffff60043516600160a060020a03602435166044356064356084356110b4565b34801561046a57600080fd5b5061015167ffffffffffffffff600435166114b9565b61012867ffffffffffffffff60043516600160a060020a036024358116906044358116906064351515906084359060a4351660c43560e43561158d565b60006104c7611ad1565b600160a060020a03891633146104dc57600080fd5b60005460a060020a900460ff1615156001146104f757600080fd5b60018715151461050657600080fd5b34861461051257600080fd5b600160a060020a0384161561054b57600160a060020a03841660009081526002602052604090205460ff16151560011461054b57600080fd5b6101c0604051908101604052808b67ffffffffffffffff1681526020018a600160a060020a0316815260200189600160a060020a0316815260200188151581526020018781526020016000600160a060020a031681526020018681526020016000815260200185600160a060020a031681526020016000815260200160011515815260200160001515815260200184815260200184815250905080600160008c67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a815481600160a060020a030219169083600160a060020a0316021790555060408201518160010160006101000a815481600160a060020a030219169083600160a060020a0316021790555060608201518160010160146101000a81548160ff0219169083151502179055506080820151816002015560a08201518160030160006101000a815481600160a060020a030219169083600160a060020a0316021790555060c0820151816004015560e082015181600501556101008201518160060160006101000a815481600160a060020a030219169083600160a060020a0316021790555061012082015181600701556101408201518160080160006101000a81548160ff0219169083151502179055506101608201518160080160016101000a81548160ff02191690831515021790555061018082015181600901556101a082015181600a01559050507fbad96d0d267222ee6012f37e5f48e919d6fb5db9b0601bc9e75f95413ed35d8d8a88888860008989604051808867ffffffffffffffff1667ffffffffffffffff1681526020018715151515815260200186815260200185815260200184815260200183600160a060020a0316600160a060020a0316815260200182815260200197505050505050505060405180910390a15060019998505050505050505050565b6000543390600160a060020a0316811461085057600080fd5b60005460a060020a900460ff161561086757600080fd5b506000805474ff0000000000000000000000000000000000000000191660a060020a179055565b6000543390600160a060020a031681146108a757600080fd5b6000821180156108b8575030318211155b15156108c357600080fd5b6108cc82611a6c565b6040805133815260208101849052428183015290517f34f457f9b40515dc7505fa1ca125d98bad483e73f905d02ab62a027f27cb292b9181900360600190a15050565b6000543390600160a060020a0316811461092857600080fd5b60005460a060020a900460ff16151560011461094357600080fd5b506000805474ff000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b60008054819060a060020a900460ff16151560011461099257600080fd5b67ffffffffffffffff8316600090815260016020526040902054680100000000000000009004600160a060020a031633146109cc57600080fd5b5067ffffffffffffffff821660009081526001602081905260409091206002810154600890910154909160ff9091161515148015610a32575067ffffffffffffffff8316600090815260016020819052604090912060080154610100900460ff16151514155b8015610a3e5750600081115b1515610a4957600080fd5b6000547501000000000000000000000000000000000000000000900460ff1615610a7257600080fd5b67ffffffffffffffff8316600090815260016020526040902060030154600160a060020a031615610d0c576000805475ff0000000000000000000000000000000000000000001916750100000000000000000000000000000000000000000017815567ffffffffffffffff841681526001602090815260408083206003015481517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482018190526024820152604481018690529151600160a060020a03909116936323b872dd93606480850194919392918390030190829087803b158015610b5e57600080fd5b505af1158015610b72573d6000803e3d6000fd5b505050506040513d6020811015610b8857600080fd5b50506000805475ff0000000000000000000000000000000000000000001916815567ffffffffffffffff84168152600160205260409020600501541515610c8c5767ffffffffffffffff83166000908152600160208190526040822080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168155908101805474ffffffffffffffffffffffffffffffffffffffffff191690556002810182905560038101805473ffffffffffffffffffffffffffffffffffffffff19908116909155600482018390556005820183905560068201805490911690556007810182905560088101805461ffff1916905560098101829055600a0155610cc6565b67ffffffffffffffff83166000908152600160205260408120600281019190915560088101805461ffff191661010017905542600a909101555b6040805167ffffffffffffffff8516815290517f0269114ad2350df1094ec1642497bb8df36a692522a7f0ddef9f9cd2142d91429181900360200190a160019150610e71565b6000805475ff00000000000000000000000000000000000000000019167501000000000000000000000000000000000000000000178155604051339183156108fc02918491818181858888f19350505050158015610d6e573d6000803e3d6000fd5b506000805475ff0000000000000000000000000000000000000000001916815567ffffffffffffffff84168152600160205260409020600501541515610c8c5767ffffffffffffffff83166000908152600160208190526040822080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168155908101805474ffffffffffffffffffffffffffffffffffffffffff191690556002810182905560038101805473ffffffffffffffffffffffffffffffffffffffff19908116909155600482018390556005820183905560068201805490911690556007810182905560088101805461ffff1916905560098101829055600a0155610cc6565b50919050565b600080543390600160a060020a03168114610e9157600080fd5b303191505b5090565b60036020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610f2d5780601f10610f0257610100808354040283529160200191610f2d565b820191906000526020600020905b815481529060010190602001808311610f1057829003601f168201915b505050505081565b6000543390600160a060020a03168114610f4e57600080fd5b600054600160a060020a0316ff5b60026020526000908152604090205460ff1681565b6000543390600160a060020a03168114610f8a57600080fd5b60018215151415610fd657600160a060020a0384166000908152600260209081526040808320805460ff19166001179055600382529091208451610fd092860190611b45565b50611009565b600160a060020a0384166000908152600260209081526040808320805460ff191690556003909152812061100991611bbf565b50505050565b600160208190526000918252604090912080549181015460028201546003830154600484015460058501546006860154600787015460088801546009890154600a9099015467ffffffffffffffff8b169a600160a060020a036801000000000000000090910481169a8a82169a60ff60a060020a90910481169a998316989796909216949380831693610100909104909216918e565b600054600160a060020a031681565b6000805481903390600160a060020a031681146110d057600080fd5b60005460a060020a900460ff1615156001146110eb57600080fd5b6110fb868563ffffffff611aa916565b67ffffffffffffffff8916600090815260016020526040902060020154108015906111265750600086115b80156111325750600084115b801561113e5750600085115b151561114957600080fd5b67ffffffffffffffff88166000908152600160208190526040909120600281015491015490925060a060020a900460ff1615156112f15767ffffffffffffffff881660009081526001602090815260408083206003015481517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a038c81166004830152602482018c90529251929091169363a9059cbb9360448084019491939192918390030190829087803b15801561120a57600080fd5b505af115801561121e573d6000803e3d6000fd5b505050506040513d602081101561123457600080fd5b505067ffffffffffffffff8816600090815260016020908152604080832060030154835482517fa9059cbb000000000000000000000000000000000000000000000000000000008152600160a060020a039182166004820152602481018a9052925191169363a9059cbb93604480850194919392918390030190829087803b1580156112bf57600080fd5b505af11580156112d3573d6000803e3d6000fd5b505050506040513d60208110156112e957600080fd5b506113299050565b604051600160a060020a0388169087156108fc029088906000818181858888f19350505050158015611327573d6000803e3d6000fd5b505b61134961133c878663ffffffff611aa916565b839063ffffffff611abf16565b67ffffffffffffffff8916600090815260016020526040902060028101919091556007015461137e908563ffffffff611aa916565b67ffffffffffffffff891660009081526001602052604090206007810191909155600501546113b3908663ffffffff611aa916565b67ffffffffffffffff89166000908152600160205260409020600581019190915542600a909101556113eb868563ffffffff611aa916565b8214156114205767ffffffffffffffff88166000908152600160205260409020600801805461ffff191661010017905561144b565b67ffffffffffffffff88166000908152600160208190526040909120600801805460ff191690911790555b6040805167ffffffffffffffff8a168152600160a060020a0389166020820152808201889052606081018790526080810186905290517f4cb073b1936b49cc564cd29a6510c77130225593415dbd914defbceb27d417da9181900360a00190a1506001979650505050505050565b600080543390600160a060020a031681146114d357600080fd5b60005460a060020a900460ff1615156001146114ee57600080fd5b67ffffffffffffffff831660009081526001602081905260409091206008015460ff8082169450610100909104161515141561152957600080fd5b6001821515141561155d5767ffffffffffffffff83166000908152600160205260409020600801805460ff19169055611588565b67ffffffffffffffff83166000908152600160208190526040909120600801805460ff191690911790555b505050565b6000611597611ad1565b6000543390600160a060020a031681146115b057600080fd5b60005460a060020a900460ff1615156001146115cb57600080fd5b87156115d657600080fd5b600160a060020a03861615156115eb57600080fd5b600160a060020a03861660009081526002602052604090205460ff16151560011461161557600080fd5b604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a038c811660048301523060248301529151899289169163dd62ed3e9160448083019260209291908290030181600087803b15801561168057600080fd5b505af1158015611694573d6000803e3d6000fd5b505050506040513d60208110156116aa57600080fd5b505110156116b757600080fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a038c81166004830152306024830152604482018a90529151918816916323b872dd916064808201926020929091908290030181600087803b15801561172957600080fd5b505af115801561173d573d6000803e3d6000fd5b505050506040513d602081101561175357600080fd5b8101908080519060200190929190505050506101c0604051908101604052808c67ffffffffffffffff1681526020018b600160a060020a031681526020018a600160a060020a03168152602001891515815260200188815260200187600160a060020a03168152602001868152602001600081526020016000600160a060020a031681526020016000815260200160011515815260200160001515815260200185815260200185815250915081600160008d67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060208201518160000160086101000a815481600160a060020a030219169083600160a060020a0316021790555060408201518160010160006101000a815481600160a060020a030219169083600160a060020a0316021790555060608201518160010160146101000a81548160ff0219169083151502179055506080820151816002015560a08201518160030160006101000a815481600160a060020a030219169083600160a060020a0316021790555060c0820151816004015560e082015181600501556101008201518160060160006101000a815481600160a060020a030219169083600160a060020a0316021790555061012082015181600701556101408201518160080160006101000a81548160ff0219169083151502179055506101608201518160080160016101000a81548160ff02191690831515021790555061018082015181600901556101a082015181600a01559050507f32880d39225648d4e2ab5afba3b87d917f0346fe35d66e48202117b5cf8249578b8a8a8a8a8a60008b604051808967ffffffffffffffff1667ffffffffffffffff16815260200188600160a060020a0316600160a060020a031681526020018715151515815260200186815260200185600160a060020a0316600160a060020a031681526020018481526020018381526020018281526020019850505050505050505060405180910390a15060019a9950505050505050505050565b60008054604051600160a060020a039091169183156108fc02918491818181858888f19350505050158015611aa5573d6000803e3d6000fd5b5050565b600082820183811015611ab857fe5b9392505050565b600082821115611acb57fe5b50900390565b604080516101c081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611b8657805160ff1916838001178555611bb3565b82800160010185558215611bb3579182015b82811115611bb3578251825591602001919060010190611b98565b50610e96929150611c06565b50805460018160011615610100020316600290046000825580601f10611be55750611c03565b601f016020900490600052602060002090810190611c039190611c06565b50565b611c2091905b80821115610e965760008155600101611c0c565b905600a165627a7a723058204f644eb44142f857d4ac9f8b97f19570c7e1d1824d8def1a4df430c6c4db60660029`

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

// ContractAddressToType is a free data retrieval call binding the contract method 0x958e2bcb.
//
// Solidity: function contractAddressToType(address ) constant returns(string)
func (_Exchange *ExchangeCaller) ContractAddressToType(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "contractAddressToType", arg0)
	return *ret0, err
}

// ContractAddressToType is a free data retrieval call binding the contract method 0x958e2bcb.
//
// Solidity: function contractAddressToType(address ) constant returns(string)
func (_Exchange *ExchangeSession) ContractAddressToType(arg0 common.Address) (string, error) {
	return _Exchange.Contract.ContractAddressToType(&_Exchange.CallOpts, arg0)
}

// ContractAddressToType is a free data retrieval call binding the contract method 0x958e2bcb.
//
// Solidity: function contractAddressToType(address ) constant returns(string)
func (_Exchange *ExchangeCallerSession) ContractAddressToType(arg0 common.Address) (string, error) {
	return _Exchange.Contract.ContractAddressToType(&_Exchange.CallOpts, arg0)
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

// OrderidToUsers is a free data retrieval call binding the contract method 0xcda33c6e.
//
// Solidity: function orderidToUsers(uint64 ) constant returns(uint64 order_id, address user_address, address user_receiver, bool trade_type, uint256 sell_amount, address sell_address, uint256 price, uint256 buy_amount, address buy_address, uint256 fees, bool withdrawalable, bool ex_success, uint256 create_at, uint256 update_at)
func (_Exchange *ExchangeCaller) OrderidToUsers(opts *bind.CallOpts, arg0 uint64) (struct {
	OrderId        uint64
	UserAddress    common.Address
	UserReceiver   common.Address
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
		UserReceiver   common.Address
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
// Solidity: function orderidToUsers(uint64 ) constant returns(uint64 order_id, address user_address, address user_receiver, bool trade_type, uint256 sell_amount, address sell_address, uint256 price, uint256 buy_amount, address buy_address, uint256 fees, bool withdrawalable, bool ex_success, uint256 create_at, uint256 update_at)
func (_Exchange *ExchangeSession) OrderidToUsers(arg0 uint64) (struct {
	OrderId        uint64
	UserAddress    common.Address
	UserReceiver   common.Address
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
// Solidity: function orderidToUsers(uint64 ) constant returns(uint64 order_id, address user_address, address user_receiver, bool trade_type, uint256 sell_amount, address sell_address, uint256 price, uint256 buy_amount, address buy_address, uint256 fees, bool withdrawalable, bool ex_success, uint256 create_at, uint256 update_at)
func (_Exchange *ExchangeCallerSession) OrderidToUsers(arg0 uint64) (struct {
	OrderId        uint64
	UserAddress    common.Address
	UserReceiver   common.Address
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

// ChangeContractAddress is a paid mutator transaction binding the contract method 0xcb037adc.
//
// Solidity: function changeContractAddress(address addr, string types, bool chnageType) returns()
func (_Exchange *ExchangeTransactor) ChangeContractAddress(opts *bind.TransactOpts, addr common.Address, types string, chnageType bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "changeContractAddress", addr, types, chnageType)
}

// ChangeContractAddress is a paid mutator transaction binding the contract method 0xcb037adc.
//
// Solidity: function changeContractAddress(address addr, string types, bool chnageType) returns()
func (_Exchange *ExchangeSession) ChangeContractAddress(addr common.Address, types string, chnageType bool) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeContractAddress(&_Exchange.TransactOpts, addr, types, chnageType)
}

// ChangeContractAddress is a paid mutator transaction binding the contract method 0xcb037adc.
//
// Solidity: function changeContractAddress(address addr, string types, bool chnageType) returns()
func (_Exchange *ExchangeTransactorSession) ChangeContractAddress(addr common.Address, types string, chnageType bool) (*types.Transaction, error) {
	return _Exchange.Contract.ChangeContractAddress(&_Exchange.TransactOpts, addr, types, chnageType)
}

// DepositsBuy is a paid mutator transaction binding the contract method 0xf4b06dda.
//
// Solidity: function depositsBuy(uint64 orderid, address useraddress, address receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactor) DepositsBuy(opts *bind.TransactOpts, orderid uint64, useraddress common.Address, receiver common.Address, tradetype bool, sellamount *big.Int, selladdress common.Address, price *big.Int, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositsBuy", orderid, useraddress, receiver, tradetype, sellamount, selladdress, price, createTime)
}

// DepositsBuy is a paid mutator transaction binding the contract method 0xf4b06dda.
//
// Solidity: function depositsBuy(uint64 orderid, address useraddress, address receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 createTime) returns(bool)
func (_Exchange *ExchangeSession) DepositsBuy(orderid uint64, useraddress common.Address, receiver common.Address, tradetype bool, sellamount *big.Int, selladdress common.Address, price *big.Int, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsBuy(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, selladdress, price, createTime)
}

// DepositsBuy is a paid mutator transaction binding the contract method 0xf4b06dda.
//
// Solidity: function depositsBuy(uint64 orderid, address useraddress, address receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactorSession) DepositsBuy(orderid uint64, useraddress common.Address, receiver common.Address, tradetype bool, sellamount *big.Int, selladdress common.Address, price *big.Int, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsBuy(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, selladdress, price, createTime)
}

// DepositsSell is a paid mutator transaction binding the contract method 0x2cda39d9.
//
// Solidity: function depositsSell(uint64 orderid, address useraddress, address receiver, bool tradetype, uint256 sellamount, uint256 price, address buyaddress, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactor) DepositsSell(opts *bind.TransactOpts, orderid uint64, useraddress common.Address, receiver common.Address, tradetype bool, sellamount *big.Int, price *big.Int, buyaddress common.Address, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "depositsSell", orderid, useraddress, receiver, tradetype, sellamount, price, buyaddress, createTime)
}

// DepositsSell is a paid mutator transaction binding the contract method 0x2cda39d9.
//
// Solidity: function depositsSell(uint64 orderid, address useraddress, address receiver, bool tradetype, uint256 sellamount, uint256 price, address buyaddress, uint256 createTime) returns(bool)
func (_Exchange *ExchangeSession) DepositsSell(orderid uint64, useraddress common.Address, receiver common.Address, tradetype bool, sellamount *big.Int, price *big.Int, buyaddress common.Address, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsSell(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, price, buyaddress, createTime)
}

// DepositsSell is a paid mutator transaction binding the contract method 0x2cda39d9.
//
// Solidity: function depositsSell(uint64 orderid, address useraddress, address receiver, bool tradetype, uint256 sellamount, uint256 price, address buyaddress, uint256 createTime) returns(bool)
func (_Exchange *ExchangeTransactorSession) DepositsSell(orderid uint64, useraddress common.Address, receiver common.Address, tradetype bool, sellamount *big.Int, price *big.Int, buyaddress common.Address, createTime *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.DepositsSell(&_Exchange.TransactOpts, orderid, useraddress, receiver, tradetype, sellamount, price, buyaddress, createTime)
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

// ExchangeDepositsBuyIterator is returned from FilterDepositsBuy and is used to iterate over the raw logs and unpacked data for DepositsBuy events raised by the Exchange contract.
type ExchangeDepositsBuyIterator struct {
	Event *ExchangeDepositsBuy // Event containing the contract specifics and raw log

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
func (it *ExchangeDepositsBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeDepositsBuy)
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
		it.Event = new(ExchangeDepositsBuy)
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
func (it *ExchangeDepositsBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeDepositsBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeDepositsBuy represents a DepositsBuy event raised by the Exchange contract.
type ExchangeDepositsBuy struct {
	Orderid     uint64
	Receiver    common.Address
	Tradetype   bool
	Sellamount  *big.Int
	Selladdress common.Address
	Price       *big.Int
	Buyamount   *big.Int
	CreateTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositsBuy is a free log retrieval operation binding the contract event 0x32880d39225648d4e2ab5afba3b87d917f0346fe35d66e48202117b5cf824957.
//
// Solidity: event DepositsBuy(uint64 orderid, address receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 buyamount, uint256 createTime)
func (_Exchange *ExchangeFilterer) FilterDepositsBuy(opts *bind.FilterOpts) (*ExchangeDepositsBuyIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "DepositsBuy")
	if err != nil {
		return nil, err
	}
	return &ExchangeDepositsBuyIterator{contract: _Exchange.contract, event: "DepositsBuy", logs: logs, sub: sub}, nil
}

// WatchDepositsBuy is a free log subscription operation binding the contract event 0x32880d39225648d4e2ab5afba3b87d917f0346fe35d66e48202117b5cf824957.
//
// Solidity: event DepositsBuy(uint64 orderid, address receiver, bool tradetype, uint256 sellamount, address selladdress, uint256 price, uint256 buyamount, uint256 createTime)
func (_Exchange *ExchangeFilterer) WatchDepositsBuy(opts *bind.WatchOpts, sink chan<- *ExchangeDepositsBuy) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "DepositsBuy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeDepositsBuy)
				if err := _Exchange.contract.UnpackLog(event, "DepositsBuy", log); err != nil {
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
	Orderid    uint64
	Tradetype  bool
	Sellamount *big.Int
	Price      *big.Int
	Buyamount  *big.Int
	Buyaddress common.Address
	CreateTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDepositsSell is a free log retrieval operation binding the contract event 0xbad96d0d267222ee6012f37e5f48e919d6fb5db9b0601bc9e75f95413ed35d8d.
//
// Solidity: event DepositsSell(uint64 orderid, bool tradetype, uint256 sellamount, uint256 price, uint256 buyamount, address buyaddress, uint256 createTime)
func (_Exchange *ExchangeFilterer) FilterDepositsSell(opts *bind.FilterOpts) (*ExchangeDepositsSellIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "DepositsSell")
	if err != nil {
		return nil, err
	}
	return &ExchangeDepositsSellIterator{contract: _Exchange.contract, event: "DepositsSell", logs: logs, sub: sub}, nil
}

// WatchDepositsSell is a free log subscription operation binding the contract event 0xbad96d0d267222ee6012f37e5f48e919d6fb5db9b0601bc9e75f95413ed35d8d.
//
// Solidity: event DepositsSell(uint64 orderid, bool tradetype, uint256 sellamount, uint256 price, uint256 buyamount, address buyaddress, uint256 createTime)
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
