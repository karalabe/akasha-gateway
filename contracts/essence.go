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

// EssenceABI is the input ABI used to generate the binding from.
const EssenceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_action\",\"type\":\"bytes32\"},{\"name\":\"_source\",\"type\":\"bytes32\"}],\"name\":\"collectFor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_collected\",\"type\":\"uint256\"}],\"name\":\"aethValueFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_aeth\",\"type\":\"address\"}],\"name\":\"setAeth\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addToWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_total\",\"type\":\"uint256\"}],\"name\":\"newHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transformEssence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newFactor\",\"type\":\"uint256\"}],\"name\":\"setFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_min\",\"type\":\"uint256\"}],\"name\":\"setMinAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"mana\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint256\"},{\"name\":\"_spent\",\"type\":\"uint256\"},{\"name\":\"_remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_initiator\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_scope\",\"type\":\"bytes32\"}],\"name\":\"spendMana\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transformFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_collector\",\"type\":\"address\"}],\"name\":\"getCollected\",\"outputs\":[{\"name\":\"_karma\",\"type\":\"uint256\"},{\"name\":\"_essence\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"totalToMint\",\"type\":\"uint256\"}],\"name\":\"RefreshMana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"scope\",\"type\":\"bytes32\"}],\"name\":\"SpendMana\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"total\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"source\",\"type\":\"bytes32\"}],\"name\":\"CollectEssence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ConvertEssence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Essence is an auto generated Go binding around an Ethereum contract.
type Essence struct {
	EssenceCaller     // Read-only binding to the contract
	EssenceTransactor // Write-only binding to the contract
}

// EssenceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EssenceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EssenceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EssenceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EssenceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EssenceSession struct {
	Contract     *Essence          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EssenceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EssenceCallerSession struct {
	Contract *EssenceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EssenceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EssenceTransactorSession struct {
	Contract     *EssenceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EssenceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EssenceRaw struct {
	Contract *Essence // Generic contract binding to access the raw methods on
}

// EssenceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EssenceCallerRaw struct {
	Contract *EssenceCaller // Generic read-only contract binding to access the raw methods on
}

// EssenceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EssenceTransactorRaw struct {
	Contract *EssenceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEssence creates a new instance of Essence, bound to a specific deployed contract.
func NewEssence(address common.Address, backend bind.ContractBackend) (*Essence, error) {
	contract, err := bindEssence(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Essence{EssenceCaller: EssenceCaller{contract: contract}, EssenceTransactor: EssenceTransactor{contract: contract}}, nil
}

// NewEssenceCaller creates a new read-only instance of Essence, bound to a specific deployed contract.
func NewEssenceCaller(address common.Address, caller bind.ContractCaller) (*EssenceCaller, error) {
	contract, err := bindEssence(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &EssenceCaller{contract: contract}, nil
}

// NewEssenceTransactor creates a new write-only instance of Essence, bound to a specific deployed contract.
func NewEssenceTransactor(address common.Address, transactor bind.ContractTransactor) (*EssenceTransactor, error) {
	contract, err := bindEssence(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &EssenceTransactor{contract: contract}, nil
}

// bindEssence binds a generic wrapper to an already deployed contract.
func bindEssence(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EssenceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Essence *EssenceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Essence.Contract.EssenceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Essence *EssenceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Essence.Contract.EssenceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Essence *EssenceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Essence.Contract.EssenceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Essence *EssenceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Essence.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Essence *EssenceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Essence.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Essence *EssenceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Essence.Contract.contract.Transact(opts, method, params...)
}

// AethValueFrom is a free data retrieval call binding the contract method 0x2b062c91.
//
// Solidity: function aethValueFrom(_collected uint256) constant returns(uint256)
func (_Essence *EssenceCaller) AethValueFrom(opts *bind.CallOpts, _collected *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Essence.contract.Call(opts, out, "aethValueFrom", _collected)
	return *ret0, err
}

// AethValueFrom is a free data retrieval call binding the contract method 0x2b062c91.
//
// Solidity: function aethValueFrom(_collected uint256) constant returns(uint256)
func (_Essence *EssenceSession) AethValueFrom(_collected *big.Int) (*big.Int, error) {
	return _Essence.Contract.AethValueFrom(&_Essence.CallOpts, _collected)
}

// AethValueFrom is a free data retrieval call binding the contract method 0x2b062c91.
//
// Solidity: function aethValueFrom(_collected uint256) constant returns(uint256)
func (_Essence *EssenceCallerSession) AethValueFrom(_collected *big.Int) (*big.Int, error) {
	return _Essence.Contract.AethValueFrom(&_Essence.CallOpts, _collected)
}

// GetCollected is a free data retrieval call binding the contract method 0xfb166bb0.
//
// Solidity: function getCollected(_collector address) constant returns(_karma uint256, _essence uint256)
func (_Essence *EssenceCaller) GetCollected(opts *bind.CallOpts, _collector common.Address) (struct {
	Karma   *big.Int
	Essence *big.Int
}, error) {
	ret := new(struct {
		Karma   *big.Int
		Essence *big.Int
	})
	out := ret
	err := _Essence.contract.Call(opts, out, "getCollected", _collector)
	return *ret, err
}

// GetCollected is a free data retrieval call binding the contract method 0xfb166bb0.
//
// Solidity: function getCollected(_collector address) constant returns(_karma uint256, _essence uint256)
func (_Essence *EssenceSession) GetCollected(_collector common.Address) (struct {
	Karma   *big.Int
	Essence *big.Int
}, error) {
	return _Essence.Contract.GetCollected(&_Essence.CallOpts, _collector)
}

// GetCollected is a free data retrieval call binding the contract method 0xfb166bb0.
//
// Solidity: function getCollected(_collector address) constant returns(_karma uint256, _essence uint256)
func (_Essence *EssenceCallerSession) GetCollected(_collector common.Address) (struct {
	Karma   *big.Int
	Essence *big.Int
}, error) {
	return _Essence.Contract.GetCollected(&_Essence.CallOpts, _collector)
}

// Mana is a free data retrieval call binding the contract method 0x94dd7a4e.
//
// Solidity: function mana(_spender address) constant returns(_total uint256, _spent uint256, _remaining uint256)
func (_Essence *EssenceCaller) Mana(opts *bind.CallOpts, _spender common.Address) (struct {
	Total     *big.Int
	Spent     *big.Int
	Remaining *big.Int
}, error) {
	ret := new(struct {
		Total     *big.Int
		Spent     *big.Int
		Remaining *big.Int
	})
	out := ret
	err := _Essence.contract.Call(opts, out, "mana", _spender)
	return *ret, err
}

// Mana is a free data retrieval call binding the contract method 0x94dd7a4e.
//
// Solidity: function mana(_spender address) constant returns(_total uint256, _spent uint256, _remaining uint256)
func (_Essence *EssenceSession) Mana(_spender common.Address) (struct {
	Total     *big.Int
	Spent     *big.Int
	Remaining *big.Int
}, error) {
	return _Essence.Contract.Mana(&_Essence.CallOpts, _spender)
}

// Mana is a free data retrieval call binding the contract method 0x94dd7a4e.
//
// Solidity: function mana(_spender address) constant returns(_total uint256, _spent uint256, _remaining uint256)
func (_Essence *EssenceCallerSession) Mana(_spender common.Address) (struct {
	Total     *big.Int
	Spent     *big.Int
	Remaining *big.Int
}, error) {
	return _Essence.Contract.Mana(&_Essence.CallOpts, _spender)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() constant returns(uint256)
func (_Essence *EssenceCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Essence.contract.Call(opts, out, "minAmount")
	return *ret0, err
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() constant returns(uint256)
func (_Essence *EssenceSession) MinAmount() (*big.Int, error) {
	return _Essence.Contract.MinAmount(&_Essence.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() constant returns(uint256)
func (_Essence *EssenceCallerSession) MinAmount() (*big.Int, error) {
	return _Essence.Contract.MinAmount(&_Essence.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Essence *EssenceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Essence.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Essence *EssenceSession) Owner() (common.Address, error) {
	return _Essence.Contract.Owner(&_Essence.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Essence *EssenceCallerSession) Owner() (common.Address, error) {
	return _Essence.Contract.Owner(&_Essence.CallOpts)
}

// TransformFactor is a free data retrieval call binding the contract method 0xe9fbb684.
//
// Solidity: function transformFactor() constant returns(uint256)
func (_Essence *EssenceCaller) TransformFactor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Essence.contract.Call(opts, out, "transformFactor")
	return *ret0, err
}

// TransformFactor is a free data retrieval call binding the contract method 0xe9fbb684.
//
// Solidity: function transformFactor() constant returns(uint256)
func (_Essence *EssenceSession) TransformFactor() (*big.Int, error) {
	return _Essence.Contract.TransformFactor(&_Essence.CallOpts)
}

// TransformFactor is a free data retrieval call binding the contract method 0xe9fbb684.
//
// Solidity: function transformFactor() constant returns(uint256)
func (_Essence *EssenceCallerSession) TransformFactor() (*big.Int, error) {
	return _Essence.Contract.TransformFactor(&_Essence.CallOpts)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(_contract address) returns(bool)
func (_Essence *EssenceTransactor) AddToWhiteList(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "addToWhiteList", _contract)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(_contract address) returns(bool)
func (_Essence *EssenceSession) AddToWhiteList(_contract common.Address) (*types.Transaction, error) {
	return _Essence.Contract.AddToWhiteList(&_Essence.TransactOpts, _contract)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x47ee0394.
//
// Solidity: function addToWhiteList(_contract address) returns(bool)
func (_Essence *EssenceTransactorSession) AddToWhiteList(_contract common.Address) (*types.Transaction, error) {
	return _Essence.Contract.AddToWhiteList(&_Essence.TransactOpts, _contract)
}

// CollectFor is a paid mutator transaction binding the contract method 0x10e32b04.
//
// Solidity: function collectFor(_receiver address, _amount uint256, _action bytes32, _source bytes32) returns(bool)
func (_Essence *EssenceTransactor) CollectFor(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int, _action [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "collectFor", _receiver, _amount, _action, _source)
}

// CollectFor is a paid mutator transaction binding the contract method 0x10e32b04.
//
// Solidity: function collectFor(_receiver address, _amount uint256, _action bytes32, _source bytes32) returns(bool)
func (_Essence *EssenceSession) CollectFor(_receiver common.Address, _amount *big.Int, _action [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _Essence.Contract.CollectFor(&_Essence.TransactOpts, _receiver, _amount, _action, _source)
}

// CollectFor is a paid mutator transaction binding the contract method 0x10e32b04.
//
// Solidity: function collectFor(_receiver address, _amount uint256, _action bytes32, _source bytes32) returns(bool)
func (_Essence *EssenceTransactorSession) CollectFor(_receiver common.Address, _amount *big.Int, _action [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _Essence.Contract.CollectFor(&_Essence.TransactOpts, _receiver, _amount, _action, _source)
}

// NewHash is a paid mutator transaction binding the contract method 0x51ded49f.
//
// Solidity: function newHash(_hash bytes32, _total uint256) returns(bool)
func (_Essence *EssenceTransactor) NewHash(opts *bind.TransactOpts, _hash [32]byte, _total *big.Int) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "newHash", _hash, _total)
}

// NewHash is a paid mutator transaction binding the contract method 0x51ded49f.
//
// Solidity: function newHash(_hash bytes32, _total uint256) returns(bool)
func (_Essence *EssenceSession) NewHash(_hash [32]byte, _total *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.NewHash(&_Essence.TransactOpts, _hash, _total)
}

// NewHash is a paid mutator transaction binding the contract method 0x51ded49f.
//
// Solidity: function newHash(_hash bytes32, _total uint256) returns(bool)
func (_Essence *EssenceTransactorSession) NewHash(_hash [32]byte, _total *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.NewHash(&_Essence.TransactOpts, _hash, _total)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Essence *EssenceTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Essence *EssenceSession) ReclaimEther() (*types.Transaction, error) {
	return _Essence.Contract.ReclaimEther(&_Essence.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Essence *EssenceTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _Essence.Contract.ReclaimEther(&_Essence.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Essence *EssenceTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Essence *EssenceSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Essence.Contract.ReclaimToken(&_Essence.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Essence *EssenceTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Essence.Contract.ReclaimToken(&_Essence.TransactOpts, token)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(_contract address) returns(bool)
func (_Essence *EssenceTransactor) RemoveWhitelisted(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "removeWhitelisted", _contract)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(_contract address) returns(bool)
func (_Essence *EssenceSession) RemoveWhitelisted(_contract common.Address) (*types.Transaction, error) {
	return _Essence.Contract.RemoveWhitelisted(&_Essence.TransactOpts, _contract)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(_contract address) returns(bool)
func (_Essence *EssenceTransactorSession) RemoveWhitelisted(_contract common.Address) (*types.Transaction, error) {
	return _Essence.Contract.RemoveWhitelisted(&_Essence.TransactOpts, _contract)
}

// SetAeth is a paid mutator transaction binding the contract method 0x46c76ff8.
//
// Solidity: function setAeth(_aeth address) returns(bool)
func (_Essence *EssenceTransactor) SetAeth(opts *bind.TransactOpts, _aeth common.Address) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "setAeth", _aeth)
}

// SetAeth is a paid mutator transaction binding the contract method 0x46c76ff8.
//
// Solidity: function setAeth(_aeth address) returns(bool)
func (_Essence *EssenceSession) SetAeth(_aeth common.Address) (*types.Transaction, error) {
	return _Essence.Contract.SetAeth(&_Essence.TransactOpts, _aeth)
}

// SetAeth is a paid mutator transaction binding the contract method 0x46c76ff8.
//
// Solidity: function setAeth(_aeth address) returns(bool)
func (_Essence *EssenceTransactorSession) SetAeth(_aeth common.Address) (*types.Transaction, error) {
	return _Essence.Contract.SetAeth(&_Essence.TransactOpts, _aeth)
}

// SetFactor is a paid mutator transaction binding the contract method 0x817e9d31.
//
// Solidity: function setFactor(_newFactor uint256) returns(bool)
func (_Essence *EssenceTransactor) SetFactor(opts *bind.TransactOpts, _newFactor *big.Int) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "setFactor", _newFactor)
}

// SetFactor is a paid mutator transaction binding the contract method 0x817e9d31.
//
// Solidity: function setFactor(_newFactor uint256) returns(bool)
func (_Essence *EssenceSession) SetFactor(_newFactor *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.SetFactor(&_Essence.TransactOpts, _newFactor)
}

// SetFactor is a paid mutator transaction binding the contract method 0x817e9d31.
//
// Solidity: function setFactor(_newFactor uint256) returns(bool)
func (_Essence *EssenceTransactorSession) SetFactor(_newFactor *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.SetFactor(&_Essence.TransactOpts, _newFactor)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(_min uint256) returns(bool)
func (_Essence *EssenceTransactor) SetMinAmount(opts *bind.TransactOpts, _min *big.Int) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "setMinAmount", _min)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(_min uint256) returns(bool)
func (_Essence *EssenceSession) SetMinAmount(_min *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.SetMinAmount(&_Essence.TransactOpts, _min)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(_min uint256) returns(bool)
func (_Essence *EssenceTransactorSession) SetMinAmount(_min *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.SetMinAmount(&_Essence.TransactOpts, _min)
}

// SpendMana is a paid mutator transaction binding the contract method 0xda5ecde0.
//
// Solidity: function spendMana(_initiator address, _amount uint256, _scope bytes32) returns(bool)
func (_Essence *EssenceTransactor) SpendMana(opts *bind.TransactOpts, _initiator common.Address, _amount *big.Int, _scope [32]byte) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "spendMana", _initiator, _amount, _scope)
}

// SpendMana is a paid mutator transaction binding the contract method 0xda5ecde0.
//
// Solidity: function spendMana(_initiator address, _amount uint256, _scope bytes32) returns(bool)
func (_Essence *EssenceSession) SpendMana(_initiator common.Address, _amount *big.Int, _scope [32]byte) (*types.Transaction, error) {
	return _Essence.Contract.SpendMana(&_Essence.TransactOpts, _initiator, _amount, _scope)
}

// SpendMana is a paid mutator transaction binding the contract method 0xda5ecde0.
//
// Solidity: function spendMana(_initiator address, _amount uint256, _scope bytes32) returns(bool)
func (_Essence *EssenceTransactorSession) SpendMana(_initiator common.Address, _amount *big.Int, _scope [32]byte) (*types.Transaction, error) {
	return _Essence.Contract.SpendMana(&_Essence.TransactOpts, _initiator, _amount, _scope)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Essence *EssenceTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Essence *EssenceSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Essence.Contract.TokenFallback(&_Essence.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Essence *EssenceTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Essence.Contract.TokenFallback(&_Essence.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Essence *EssenceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Essence *EssenceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Essence.Contract.TransferOwnership(&_Essence.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Essence *EssenceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Essence.Contract.TransferOwnership(&_Essence.TransactOpts, newOwner)
}

// TransformEssence is a paid mutator transaction binding the contract method 0x53021332.
//
// Solidity: function transformEssence(_amount uint256) returns(bool)
func (_Essence *EssenceTransactor) TransformEssence(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Essence.contract.Transact(opts, "transformEssence", _amount)
}

// TransformEssence is a paid mutator transaction binding the contract method 0x53021332.
//
// Solidity: function transformEssence(_amount uint256) returns(bool)
func (_Essence *EssenceSession) TransformEssence(_amount *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.TransformEssence(&_Essence.TransactOpts, _amount)
}

// TransformEssence is a paid mutator transaction binding the contract method 0x53021332.
//
// Solidity: function transformEssence(_amount uint256) returns(bool)
func (_Essence *EssenceTransactorSession) TransformEssence(_amount *big.Int) (*types.Transaction, error) {
	return _Essence.Contract.TransformEssence(&_Essence.TransactOpts, _amount)
}
