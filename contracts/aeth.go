// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// AETHABI is the input ABI used to generate the binding from.
const AETHABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_fromDate\",\"type\":\"uint256\"}],\"name\":\"getCyclingStatesNr\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint8\"},{\"name\":\"_available\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_fromIndex\",\"type\":\"uint8\"}],\"name\":\"getCyclingState\",\"outputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_unlockDate\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bondAeth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"freeAeth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cycleAeth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"cycling\",\"outputs\":[{\"name\":\"_cycling\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_aethAmount\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"string\"}],\"name\":\"donate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_essence\",\"type\":\"address\"}],\"name\":\"setEssence\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setLockTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"bonded\",\"outputs\":[{\"name\":\"_bonded\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"getTokenRecords\",\"outputs\":[{\"name\":\"_free\",\"type\":\"uint256\"},{\"name\":\"_bonded\",\"type\":\"uint256\"},{\"name\":\"_cycling\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"consumeEssence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"aeth\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eth\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"extraData\",\"type\":\"string\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// AETH is an auto generated Go binding around an Ethereum contract.
type AETH struct {
	AETHCaller     // Read-only binding to the contract
	AETHTransactor // Write-only binding to the contract
	AETHFilterer   // Log filterer for contract events
}

// AETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type AETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AETHSession struct {
	Contract     *AETH             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AETHCallerSession struct {
	Contract *AETHCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AETHTransactorSession struct {
	Contract     *AETHTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type AETHRaw struct {
	Contract *AETH // Generic contract binding to access the raw methods on
}

// AETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AETHCallerRaw struct {
	Contract *AETHCaller // Generic read-only contract binding to access the raw methods on
}

// AETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AETHTransactorRaw struct {
	Contract *AETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAETH creates a new instance of AETH, bound to a specific deployed contract.
func NewAETH(address common.Address, backend bind.ContractBackend) (*AETH, error) {
	contract, err := bindAETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AETH{AETHCaller: AETHCaller{contract: contract}, AETHTransactor: AETHTransactor{contract: contract}, AETHFilterer: AETHFilterer{contract: contract}}, nil
}

// NewAETHCaller creates a new read-only instance of AETH, bound to a specific deployed contract.
func NewAETHCaller(address common.Address, caller bind.ContractCaller) (*AETHCaller, error) {
	contract, err := bindAETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AETHCaller{contract: contract}, nil
}

// NewAETHTransactor creates a new write-only instance of AETH, bound to a specific deployed contract.
func NewAETHTransactor(address common.Address, transactor bind.ContractTransactor) (*AETHTransactor, error) {
	contract, err := bindAETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AETHTransactor{contract: contract}, nil
}

// NewAETHFilterer creates a new log filterer instance of AETH, bound to a specific deployed contract.
func NewAETHFilterer(address common.Address, filterer bind.ContractFilterer) (*AETHFilterer, error) {
	contract, err := bindAETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AETHFilterer{contract: contract}, nil
}

// bindAETH binds a generic wrapper to an already deployed contract.
func bindAETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AETH *AETHRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AETH.Contract.AETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AETH *AETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETH.Contract.AETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AETH *AETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AETH.Contract.AETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AETH *AETHCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AETH *AETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AETH *AETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_AETH *AETHCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_AETH *AETHSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _AETH.Contract.Allowance(&_AETH.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_AETH *AETHCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _AETH.Contract.Allowance(&_AETH.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_AETH *AETHCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_AETH *AETHSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AETH.Contract.BalanceOf(&_AETH.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_AETH *AETHCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AETH.Contract.BalanceOf(&_AETH.CallOpts, _owner)
}

// Bonded is a free data retrieval call binding the contract method 0xf19e2a21.
//
// Solidity: function bonded(_holder address) constant returns(_bonded uint256)
func (_AETH *AETHCaller) Bonded(opts *bind.CallOpts, _holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "bonded", _holder)
	return *ret0, err
}

// Bonded is a free data retrieval call binding the contract method 0xf19e2a21.
//
// Solidity: function bonded(_holder address) constant returns(_bonded uint256)
func (_AETH *AETHSession) Bonded(_holder common.Address) (*big.Int, error) {
	return _AETH.Contract.Bonded(&_AETH.CallOpts, _holder)
}

// Bonded is a free data retrieval call binding the contract method 0xf19e2a21.
//
// Solidity: function bonded(_holder address) constant returns(_bonded uint256)
func (_AETH *AETHCallerSession) Bonded(_holder common.Address) (*big.Int, error) {
	return _AETH.Contract.Bonded(&_AETH.CallOpts, _holder)
}

// Cycling is a free data retrieval call binding the contract method 0x6d2b9de0.
//
// Solidity: function cycling(_holder address) constant returns(_cycling uint256)
func (_AETH *AETHCaller) Cycling(opts *bind.CallOpts, _holder common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "cycling", _holder)
	return *ret0, err
}

// Cycling is a free data retrieval call binding the contract method 0x6d2b9de0.
//
// Solidity: function cycling(_holder address) constant returns(_cycling uint256)
func (_AETH *AETHSession) Cycling(_holder common.Address) (*big.Int, error) {
	return _AETH.Contract.Cycling(&_AETH.CallOpts, _holder)
}

// Cycling is a free data retrieval call binding the contract method 0x6d2b9de0.
//
// Solidity: function cycling(_holder address) constant returns(_cycling uint256)
func (_AETH *AETHCallerSession) Cycling(_holder common.Address) (*big.Int, error) {
	return _AETH.Contract.Cycling(&_AETH.CallOpts, _holder)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AETH *AETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AETH *AETHSession) Decimals() (uint8, error) {
	return _AETH.Contract.Decimals(&_AETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_AETH *AETHCallerSession) Decimals() (uint8, error) {
	return _AETH.Contract.Decimals(&_AETH.CallOpts)
}

// GetCyclingState is a free data retrieval call binding the contract method 0x2a3f5080.
//
// Solidity: function getCyclingState(_holder address, _fromIndex uint8) constant returns(_amount uint256, _unlockDate uint256, _index uint8)
func (_AETH *AETHCaller) GetCyclingState(opts *bind.CallOpts, _holder common.Address, _fromIndex uint8) (struct {
	Amount     *big.Int
	UnlockDate *big.Int
	Index      uint8
}, error) {
	ret := new(struct {
		Amount     *big.Int
		UnlockDate *big.Int
		Index      uint8
	})
	out := ret
	err := _AETH.contract.Call(opts, out, "getCyclingState", _holder, _fromIndex)
	return *ret, err
}

// GetCyclingState is a free data retrieval call binding the contract method 0x2a3f5080.
//
// Solidity: function getCyclingState(_holder address, _fromIndex uint8) constant returns(_amount uint256, _unlockDate uint256, _index uint8)
func (_AETH *AETHSession) GetCyclingState(_holder common.Address, _fromIndex uint8) (struct {
	Amount     *big.Int
	UnlockDate *big.Int
	Index      uint8
}, error) {
	return _AETH.Contract.GetCyclingState(&_AETH.CallOpts, _holder, _fromIndex)
}

// GetCyclingState is a free data retrieval call binding the contract method 0x2a3f5080.
//
// Solidity: function getCyclingState(_holder address, _fromIndex uint8) constant returns(_amount uint256, _unlockDate uint256, _index uint8)
func (_AETH *AETHCallerSession) GetCyclingState(_holder common.Address, _fromIndex uint8) (struct {
	Amount     *big.Int
	UnlockDate *big.Int
	Index      uint8
}, error) {
	return _AETH.Contract.GetCyclingState(&_AETH.CallOpts, _holder, _fromIndex)
}

// GetCyclingStatesNr is a free data retrieval call binding the contract method 0x2108b354.
//
// Solidity: function getCyclingStatesNr(_holder address, _fromDate uint256) constant returns(_total uint8, _available uint8)
func (_AETH *AETHCaller) GetCyclingStatesNr(opts *bind.CallOpts, _holder common.Address, _fromDate *big.Int) (struct {
	Total     uint8
	Available uint8
}, error) {
	ret := new(struct {
		Total     uint8
		Available uint8
	})
	out := ret
	err := _AETH.contract.Call(opts, out, "getCyclingStatesNr", _holder, _fromDate)
	return *ret, err
}

// GetCyclingStatesNr is a free data retrieval call binding the contract method 0x2108b354.
//
// Solidity: function getCyclingStatesNr(_holder address, _fromDate uint256) constant returns(_total uint8, _available uint8)
func (_AETH *AETHSession) GetCyclingStatesNr(_holder common.Address, _fromDate *big.Int) (struct {
	Total     uint8
	Available uint8
}, error) {
	return _AETH.Contract.GetCyclingStatesNr(&_AETH.CallOpts, _holder, _fromDate)
}

// GetCyclingStatesNr is a free data retrieval call binding the contract method 0x2108b354.
//
// Solidity: function getCyclingStatesNr(_holder address, _fromDate uint256) constant returns(_total uint8, _available uint8)
func (_AETH *AETHCallerSession) GetCyclingStatesNr(_holder common.Address, _fromDate *big.Int) (struct {
	Total     uint8
	Available uint8
}, error) {
	return _AETH.Contract.GetCyclingStatesNr(&_AETH.CallOpts, _holder, _fromDate)
}

// GetTokenRecords is a free data retrieval call binding the contract method 0xf2a36684.
//
// Solidity: function getTokenRecords(_holder address) constant returns(_free uint256, _bonded uint256, _cycling uint256)
func (_AETH *AETHCaller) GetTokenRecords(opts *bind.CallOpts, _holder common.Address) (struct {
	Free    *big.Int
	Bonded  *big.Int
	Cycling *big.Int
}, error) {
	ret := new(struct {
		Free    *big.Int
		Bonded  *big.Int
		Cycling *big.Int
	})
	out := ret
	err := _AETH.contract.Call(opts, out, "getTokenRecords", _holder)
	return *ret, err
}

// GetTokenRecords is a free data retrieval call binding the contract method 0xf2a36684.
//
// Solidity: function getTokenRecords(_holder address) constant returns(_free uint256, _bonded uint256, _cycling uint256)
func (_AETH *AETHSession) GetTokenRecords(_holder common.Address) (struct {
	Free    *big.Int
	Bonded  *big.Int
	Cycling *big.Int
}, error) {
	return _AETH.Contract.GetTokenRecords(&_AETH.CallOpts, _holder)
}

// GetTokenRecords is a free data retrieval call binding the contract method 0xf2a36684.
//
// Solidity: function getTokenRecords(_holder address) constant returns(_free uint256, _bonded uint256, _cycling uint256)
func (_AETH *AETHCallerSession) GetTokenRecords(_holder common.Address) (struct {
	Free    *big.Int
	Bonded  *big.Int
	Cycling *big.Int
}, error) {
	return _AETH.Contract.GetTokenRecords(&_AETH.CallOpts, _holder)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() constant returns(uint256)
func (_AETH *AETHCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "lockTime")
	return *ret0, err
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() constant returns(uint256)
func (_AETH *AETHSession) LockTime() (*big.Int, error) {
	return _AETH.Contract.LockTime(&_AETH.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() constant returns(uint256)
func (_AETH *AETHCallerSession) LockTime() (*big.Int, error) {
	return _AETH.Contract.LockTime(&_AETH.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_AETH *AETHCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_AETH *AETHSession) MintingFinished() (bool, error) {
	return _AETH.Contract.MintingFinished(&_AETH.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_AETH *AETHCallerSession) MintingFinished() (bool, error) {
	return _AETH.Contract.MintingFinished(&_AETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AETH *AETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AETH *AETHSession) Name() (string, error) {
	return _AETH.Contract.Name(&_AETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AETH *AETHCallerSession) Name() (string, error) {
	return _AETH.Contract.Name(&_AETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AETH *AETHCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AETH *AETHSession) Owner() (common.Address, error) {
	return _AETH.Contract.Owner(&_AETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AETH *AETHCallerSession) Owner() (common.Address, error) {
	return _AETH.Contract.Owner(&_AETH.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AETH *AETHCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AETH *AETHSession) Paused() (bool, error) {
	return _AETH.Contract.Paused(&_AETH.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_AETH *AETHCallerSession) Paused() (bool, error) {
	return _AETH.Contract.Paused(&_AETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AETH *AETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AETH *AETHSession) Symbol() (string, error) {
	return _AETH.Contract.Symbol(&_AETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_AETH *AETHCallerSession) Symbol() (string, error) {
	return _AETH.Contract.Symbol(&_AETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AETH *AETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AETH.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AETH *AETHSession) TotalSupply() (*big.Int, error) {
	return _AETH.Contract.TotalSupply(&_AETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_AETH *AETHCallerSession) TotalSupply() (*big.Int, error) {
	return _AETH.Contract.TotalSupply(&_AETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_AETH *AETHTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_AETH *AETHSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.Approve(&_AETH.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_AETH *AETHTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.Approve(&_AETH.TransactOpts, _spender, _value)
}

// BondAeth is a paid mutator transaction binding the contract method 0x501951ed.
//
// Solidity: function bondAeth(_amount uint256) returns(bool)
func (_AETH *AETHTransactor) BondAeth(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "bondAeth", _amount)
}

// BondAeth is a paid mutator transaction binding the contract method 0x501951ed.
//
// Solidity: function bondAeth(_amount uint256) returns(bool)
func (_AETH *AETHSession) BondAeth(_amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.BondAeth(&_AETH.TransactOpts, _amount)
}

// BondAeth is a paid mutator transaction binding the contract method 0x501951ed.
//
// Solidity: function bondAeth(_amount uint256) returns(bool)
func (_AETH *AETHTransactorSession) BondAeth(_amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.BondAeth(&_AETH.TransactOpts, _amount)
}

// ConsumeEssence is a paid mutator transaction binding the contract method 0xf51724df.
//
// Solidity: function consumeEssence(_to address, _amount uint256) returns(bool)
func (_AETH *AETHTransactor) ConsumeEssence(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "consumeEssence", _to, _amount)
}

// ConsumeEssence is a paid mutator transaction binding the contract method 0xf51724df.
//
// Solidity: function consumeEssence(_to address, _amount uint256) returns(bool)
func (_AETH *AETHSession) ConsumeEssence(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.ConsumeEssence(&_AETH.TransactOpts, _to, _amount)
}

// ConsumeEssence is a paid mutator transaction binding the contract method 0xf51724df.
//
// Solidity: function consumeEssence(_to address, _amount uint256) returns(bool)
func (_AETH *AETHTransactorSession) ConsumeEssence(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.ConsumeEssence(&_AETH.TransactOpts, _to, _amount)
}

// CycleAeth is a paid mutator transaction binding the contract method 0x6ab465bd.
//
// Solidity: function cycleAeth(_amount uint256) returns(bool)
func (_AETH *AETHTransactor) CycleAeth(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "cycleAeth", _amount)
}

// CycleAeth is a paid mutator transaction binding the contract method 0x6ab465bd.
//
// Solidity: function cycleAeth(_amount uint256) returns(bool)
func (_AETH *AETHSession) CycleAeth(_amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.CycleAeth(&_AETH.TransactOpts, _amount)
}

// CycleAeth is a paid mutator transaction binding the contract method 0x6ab465bd.
//
// Solidity: function cycleAeth(_amount uint256) returns(bool)
func (_AETH *AETHTransactorSession) CycleAeth(_amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.CycleAeth(&_AETH.TransactOpts, _amount)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_AETH *AETHTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_AETH *AETHSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.DecreaseApproval(&_AETH.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_AETH *AETHTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.DecreaseApproval(&_AETH.TransactOpts, _spender, _subtractedValue)
}

// Donate is a paid mutator transaction binding the contract method 0x8488893a.
//
// Solidity: function donate(_to address, _aethAmount uint256, _extraData string) returns(bool)
func (_AETH *AETHTransactor) Donate(opts *bind.TransactOpts, _to common.Address, _aethAmount *big.Int, _extraData string) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "donate", _to, _aethAmount, _extraData)
}

// Donate is a paid mutator transaction binding the contract method 0x8488893a.
//
// Solidity: function donate(_to address, _aethAmount uint256, _extraData string) returns(bool)
func (_AETH *AETHSession) Donate(_to common.Address, _aethAmount *big.Int, _extraData string) (*types.Transaction, error) {
	return _AETH.Contract.Donate(&_AETH.TransactOpts, _to, _aethAmount, _extraData)
}

// Donate is a paid mutator transaction binding the contract method 0x8488893a.
//
// Solidity: function donate(_to address, _aethAmount uint256, _extraData string) returns(bool)
func (_AETH *AETHTransactorSession) Donate(_to common.Address, _aethAmount *big.Int, _extraData string) (*types.Transaction, error) {
	return _AETH.Contract.Donate(&_AETH.TransactOpts, _to, _aethAmount, _extraData)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_AETH *AETHTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_AETH *AETHSession) FinishMinting() (*types.Transaction, error) {
	return _AETH.Contract.FinishMinting(&_AETH.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_AETH *AETHTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _AETH.Contract.FinishMinting(&_AETH.TransactOpts)
}

// FreeAeth is a paid mutator transaction binding the contract method 0x5f896246.
//
// Solidity: function freeAeth() returns(bool)
func (_AETH *AETHTransactor) FreeAeth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "freeAeth")
}

// FreeAeth is a paid mutator transaction binding the contract method 0x5f896246.
//
// Solidity: function freeAeth() returns(bool)
func (_AETH *AETHSession) FreeAeth() (*types.Transaction, error) {
	return _AETH.Contract.FreeAeth(&_AETH.TransactOpts)
}

// FreeAeth is a paid mutator transaction binding the contract method 0x5f896246.
//
// Solidity: function freeAeth() returns(bool)
func (_AETH *AETHTransactorSession) FreeAeth() (*types.Transaction, error) {
	return _AETH.Contract.FreeAeth(&_AETH.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_AETH *AETHTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_AETH *AETHSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.IncreaseApproval(&_AETH.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_AETH *AETHTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.IncreaseApproval(&_AETH.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_AETH *AETHTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_AETH *AETHSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.Mint(&_AETH.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_AETH *AETHTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.Mint(&_AETH.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AETH *AETHTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AETH *AETHSession) Pause() (*types.Transaction, error) {
	return _AETH.Contract.Pause(&_AETH.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AETH *AETHTransactorSession) Pause() (*types.Transaction, error) {
	return _AETH.Contract.Pause(&_AETH.TransactOpts)
}

// SetEssence is a paid mutator transaction binding the contract method 0x8f0ecae2.
//
// Solidity: function setEssence(_essence address) returns()
func (_AETH *AETHTransactor) SetEssence(opts *bind.TransactOpts, _essence common.Address) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "setEssence", _essence)
}

// SetEssence is a paid mutator transaction binding the contract method 0x8f0ecae2.
//
// Solidity: function setEssence(_essence address) returns()
func (_AETH *AETHSession) SetEssence(_essence common.Address) (*types.Transaction, error) {
	return _AETH.Contract.SetEssence(&_AETH.TransactOpts, _essence)
}

// SetEssence is a paid mutator transaction binding the contract method 0x8f0ecae2.
//
// Solidity: function setEssence(_essence address) returns()
func (_AETH *AETHTransactorSession) SetEssence(_essence common.Address) (*types.Transaction, error) {
	return _AETH.Contract.SetEssence(&_AETH.TransactOpts, _essence)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(_time uint256) returns()
func (_AETH *AETHTransactor) SetLockTime(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "setLockTime", _time)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(_time uint256) returns()
func (_AETH *AETHSession) SetLockTime(_time *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.SetLockTime(&_AETH.TransactOpts, _time)
}

// SetLockTime is a paid mutator transaction binding the contract method 0xae04d45d.
//
// Solidity: function setLockTime(_time uint256) returns()
func (_AETH *AETHTransactorSession) SetLockTime(_time *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.SetLockTime(&_AETH.TransactOpts, _time)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(_resolver address) returns()
func (_AETH *AETHTransactor) SetResolver(opts *bind.TransactOpts, _resolver common.Address) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "setResolver", _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(_resolver address) returns()
func (_AETH *AETHSession) SetResolver(_resolver common.Address) (*types.Transaction, error) {
	return _AETH.Contract.SetResolver(&_AETH.TransactOpts, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(_resolver address) returns()
func (_AETH *AETHTransactorSession) SetResolver(_resolver common.Address) (*types.Transaction, error) {
	return _AETH.Contract.SetResolver(&_AETH.TransactOpts, _resolver)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_AETH *AETHTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_AETH *AETHSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.Transfer(&_AETH.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_AETH *AETHTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.Transfer(&_AETH.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_AETH *AETHTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_AETH *AETHSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.TransferFrom(&_AETH.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_AETH *AETHTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AETH.Contract.TransferFrom(&_AETH.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AETH *AETHTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AETH *AETHSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AETH.Contract.TransferOwnership(&_AETH.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AETH *AETHTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AETH.Contract.TransferOwnership(&_AETH.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AETH *AETHTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AETH.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AETH *AETHSession) Unpause() (*types.Transaction, error) {
	return _AETH.Contract.Unpause(&_AETH.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AETH *AETHTransactorSession) Unpause() (*types.Transaction, error) {
	return _AETH.Contract.Unpause(&_AETH.TransactOpts)
}

// AETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AETH contract.
type AETHApprovalIterator struct {
	Event *AETHApproval // Event containing the contract specifics and raw log

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
func (it *AETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHApproval)
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
		it.Event = new(AETHApproval)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHApproval represents a Approval event raised by the AETH contract.
type AETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_AETH *AETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AETHApprovalIterator{contract: _AETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_AETH *AETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHApproval)
				if err := _AETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
		return nil
	}), nil
}

// AETHDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the AETH contract.
type AETHDonateIterator struct {
	Event *AETHDonate // Event containing the contract specifics and raw log

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
func (it *AETHDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHDonate)
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
		it.Event = new(AETHDonate)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHDonate represents a Donate event raised by the AETH contract.
type AETHDonate struct {
	From      common.Address
	To        common.Address
	Aeth      *big.Int
	Eth       *big.Int
	ExtraData string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0xf0867503c5147c27b10a8e12f7c0ca20f597dc2174d3a6b51a25cb7c5191460a.
//
// Solidity: event Donate(from indexed address, to indexed address, aeth uint256, eth uint256, extraData string)
func (_AETH *AETHFilterer) FilterDonate(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AETHDonateIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AETH.contract.FilterLogs(opts, "Donate", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AETHDonateIterator{contract: _AETH.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0xf0867503c5147c27b10a8e12f7c0ca20f597dc2174d3a6b51a25cb7c5191460a.
//
// Solidity: event Donate(from indexed address, to indexed address, aeth uint256, eth uint256, extraData string)
func (_AETH *AETHFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *AETHDonate, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AETH.contract.WatchLogs(opts, "Donate", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHDonate)
				if err := _AETH.contract.UnpackLog(event, "Donate", log); err != nil {
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
		return nil
	}), nil
}

// AETHMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AETH contract.
type AETHMintIterator struct {
	Event *AETHMint // Event containing the contract specifics and raw log

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
func (it *AETHMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHMint)
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
		it.Event = new(AETHMint)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHMint represents a Mint event raised by the AETH contract.
type AETHMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_AETH *AETHFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*AETHMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AETH.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &AETHMintIterator{contract: _AETH.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_AETH *AETHFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AETHMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AETH.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHMint)
				if err := _AETH.contract.UnpackLog(event, "Mint", log); err != nil {
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
		return nil
	}), nil
}

// AETHMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the AETH contract.
type AETHMintFinishedIterator struct {
	Event *AETHMintFinished // Event containing the contract specifics and raw log

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
func (it *AETHMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHMintFinished)
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
		it.Event = new(AETHMintFinished)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHMintFinished represents a MintFinished event raised by the AETH contract.
type AETHMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_AETH *AETHFilterer) FilterMintFinished(opts *bind.FilterOpts) (*AETHMintFinishedIterator, error) {

	logs, sub, err := _AETH.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &AETHMintFinishedIterator{contract: _AETH.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_AETH *AETHFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *AETHMintFinished) (event.Subscription, error) {

	logs, sub, err := _AETH.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHMintFinished)
				if err := _AETH.contract.UnpackLog(event, "MintFinished", log); err != nil {
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
		return nil
	}), nil
}

// AETHOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AETH contract.
type AETHOwnershipTransferredIterator struct {
	Event *AETHOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AETHOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHOwnershipTransferred)
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
		it.Event = new(AETHOwnershipTransferred)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHOwnershipTransferred represents a OwnershipTransferred event raised by the AETH contract.
type AETHOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AETH *AETHFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AETHOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AETH.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AETHOwnershipTransferredIterator{contract: _AETH.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AETH *AETHFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AETHOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AETH.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHOwnershipTransferred)
				if err := _AETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
		return nil
	}), nil
}

// AETHPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the AETH contract.
type AETHPauseIterator struct {
	Event *AETHPause // Event containing the contract specifics and raw log

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
func (it *AETHPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHPause)
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
		it.Event = new(AETHPause)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHPause represents a Pause event raised by the AETH contract.
type AETHPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_AETH *AETHFilterer) FilterPause(opts *bind.FilterOpts) (*AETHPauseIterator, error) {

	logs, sub, err := _AETH.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &AETHPauseIterator{contract: _AETH.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_AETH *AETHFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *AETHPause) (event.Subscription, error) {

	logs, sub, err := _AETH.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHPause)
				if err := _AETH.contract.UnpackLog(event, "Pause", log); err != nil {
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
		return nil
	}), nil
}

// AETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AETH contract.
type AETHTransferIterator struct {
	Event *AETHTransfer // Event containing the contract specifics and raw log

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
func (it *AETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHTransfer)
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
		it.Event = new(AETHTransfer)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHTransfer represents a Transfer event raised by the AETH contract.
type AETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_AETH *AETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AETHTransferIterator{contract: _AETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_AETH *AETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHTransfer)
				if err := _AETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		return nil
	}), nil
}

// AETHTransitionIterator is returned from FilterTransition and is used to iterate over the raw logs and unpacked data for Transition events raised by the AETH contract.
type AETHTransitionIterator struct {
	Event *AETHTransition // Event containing the contract specifics and raw log

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
func (it *AETHTransitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHTransition)
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
		it.Event = new(AETHTransition)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHTransitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHTransitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHTransition represents a Transition event raised by the AETH contract.
type AETHTransition struct {
	Owner common.Address
	State uint8
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransition is a free log retrieval operation binding the contract event 0x0f30aa5a186db94018a20857f8cbe6f1156cb35de19b2f988e78733e44f6590f.
//
// Solidity: event Transition(owner address, state uint8, value uint256)
func (_AETH *AETHFilterer) FilterTransition(opts *bind.FilterOpts) (*AETHTransitionIterator, error) {

	logs, sub, err := _AETH.contract.FilterLogs(opts, "Transition")
	if err != nil {
		return nil, err
	}
	return &AETHTransitionIterator{contract: _AETH.contract, event: "Transition", logs: logs, sub: sub}, nil
}

// WatchTransition is a free log subscription operation binding the contract event 0x0f30aa5a186db94018a20857f8cbe6f1156cb35de19b2f988e78733e44f6590f.
//
// Solidity: event Transition(owner address, state uint8, value uint256)
func (_AETH *AETHFilterer) WatchTransition(opts *bind.WatchOpts, sink chan<- *AETHTransition) (event.Subscription, error) {

	logs, sub, err := _AETH.contract.WatchLogs(opts, "Transition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHTransition)
				if err := _AETH.contract.UnpackLog(event, "Transition", log); err != nil {
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
		return nil
	}), nil
}

// AETHUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the AETH contract.
type AETHUnpauseIterator struct {
	Event *AETHUnpause // Event containing the contract specifics and raw log

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
func (it *AETHUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AETHUnpause)
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
		it.Event = new(AETHUnpause)
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

// Error retruned any retrieval or parsing error occured during filtering.
func (it *AETHUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AETHUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AETHUnpause represents a Unpause event raised by the AETH contract.
type AETHUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_AETH *AETHFilterer) FilterUnpause(opts *bind.FilterOpts) (*AETHUnpauseIterator, error) {

	logs, sub, err := _AETH.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &AETHUnpauseIterator{contract: _AETH.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_AETH *AETHFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *AETHUnpause) (event.Subscription, error) {

	logs, sub, err := _AETH.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AETHUnpause)
				if err := _AETH.contract.UnpackLog(event, "Unpause", log); err != nil {
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
		return nil
	}), nil
}
