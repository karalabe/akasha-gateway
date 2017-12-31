// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AETHABI is the input ABI used to generate the binding from.
const AETHABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_fromDate\",\"type\":\"uint256\"}],\"name\":\"getCyclingStatesNr\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint8\"},{\"name\":\"_available\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_fromIndex\",\"type\":\"uint8\"}],\"name\":\"getCyclingState\",\"outputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_unlockDate\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bondAeth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"freeAeth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cycleAeth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"cycling\",\"outputs\":[{\"name\":\"_cycling\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_aethAmount\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"string\"}],\"name\":\"donate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_essence\",\"type\":\"address\"}],\"name\":\"setEssence\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setLockTime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"bonded\",\"outputs\":[{\"name\":\"_bonded\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"}],\"name\":\"getTokenRecords\",\"outputs\":[{\"name\":\"_free\",\"type\":\"uint256\"},{\"name\":\"_bonded\",\"type\":\"uint256\"},{\"name\":\"_cycling\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"consumeEssence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"state\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"aeth\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eth\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"extraData\",\"type\":\"string\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// AETH is an auto generated Go binding around an Ethereum contract.
type AETH struct {
	AETHCaller     // Read-only binding to the contract
	AETHTransactor // Write-only binding to the contract
}

// AETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type AETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AETHTransactor struct {
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
	contract, err := bindAETH(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AETH{AETHCaller: AETHCaller{contract: contract}, AETHTransactor: AETHTransactor{contract: contract}}, nil
}

// NewAETHCaller creates a new read-only instance of AETH, bound to a specific deployed contract.
func NewAETHCaller(address common.Address, caller bind.ContractCaller) (*AETHCaller, error) {
	contract, err := bindAETH(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AETHCaller{contract: contract}, nil
}

// NewAETHTransactor creates a new write-only instance of AETH, bound to a specific deployed contract.
func NewAETHTransactor(address common.Address, transactor bind.ContractTransactor) (*AETHTransactor, error) {
	contract, err := bindAETH(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AETHTransactor{contract: contract}, nil
}

// bindAETH binds a generic wrapper to an already deployed contract.
func bindAETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
