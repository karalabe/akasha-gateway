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

// ProfileRegistrarABI is the input ABI used to generate the binding from.
const ProfileRegistrarABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_subNode\",\"type\":\"bytes32\"},{\"name\":\"_enableDonations\",\"type\":\"bool\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_subNode\",\"type\":\"bytes32\"}],\"name\":\"check_format\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"}],\"name\":\"setEns\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"moduleName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeRootOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"setRootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_subNode\",\"type\":\"bytes32\"},{\"name\":\"_enableDonations\",\"type\":\"bool\"},{\"name\":\"_alphaUser\",\"type\":\"address\"}],\"name\":\"adminRegisterFor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_subNode\",\"type\":\"bytes32\"}],\"name\":\"hash\",\"outputs\":[{\"name\":\"nameHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_subNode\",\"type\":\"bytes32\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"adminSetSubNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rootNode\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_rootNode\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"Register\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ProfileRegistrar is an auto generated Go binding around an Ethereum contract.
type ProfileRegistrar struct {
	ProfileRegistrarCaller     // Read-only binding to the contract
	ProfileRegistrarTransactor // Write-only binding to the contract
	ProfileRegistrarFilterer   // Log filterer for contract events
}

// ProfileRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileRegistrarSession struct {
	Contract     *ProfileRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileRegistrarCallerSession struct {
	Contract *ProfileRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProfileRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileRegistrarTransactorSession struct {
	Contract     *ProfileRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProfileRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileRegistrarRaw struct {
	Contract *ProfileRegistrar // Generic contract binding to access the raw methods on
}

// ProfileRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileRegistrarCallerRaw struct {
	Contract *ProfileRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileRegistrarTransactorRaw struct {
	Contract *ProfileRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfileRegistrar creates a new instance of ProfileRegistrar, bound to a specific deployed contract.
func NewProfileRegistrar(address common.Address, backend bind.ContractBackend) (*ProfileRegistrar, error) {
	contract, err := bindProfileRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistrar{ProfileRegistrarCaller: ProfileRegistrarCaller{contract: contract}, ProfileRegistrarTransactor: ProfileRegistrarTransactor{contract: contract}, ProfileRegistrarFilterer: ProfileRegistrarFilterer{contract: contract}}, nil
}

// NewProfileRegistrarCaller creates a new read-only instance of ProfileRegistrar, bound to a specific deployed contract.
func NewProfileRegistrarCaller(address common.Address, caller bind.ContractCaller) (*ProfileRegistrarCaller, error) {
	contract, err := bindProfileRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistrarCaller{contract: contract}, nil
}

// NewProfileRegistrarTransactor creates a new write-only instance of ProfileRegistrar, bound to a specific deployed contract.
func NewProfileRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileRegistrarTransactor, error) {
	contract, err := bindProfileRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistrarTransactor{contract: contract}, nil
}

// NewProfileRegistrarFilterer creates a new log filterer instance of ProfileRegistrar, bound to a specific deployed contract.
func NewProfileRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileRegistrarFilterer, error) {
	contract, err := bindProfileRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistrarFilterer{contract: contract}, nil
}

// bindProfileRegistrar binds a generic wrapper to an already deployed contract.
func bindProfileRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProfileRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfileRegistrar *ProfileRegistrarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProfileRegistrar.Contract.ProfileRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfileRegistrar *ProfileRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ProfileRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfileRegistrar *ProfileRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ProfileRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfileRegistrar *ProfileRegistrarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProfileRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfileRegistrar *ProfileRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfileRegistrar *ProfileRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.contract.Transact(opts, method, params...)
}

// Check_format is a free data retrieval call binding the contract method 0x6b4316db.
//
// Solidity: function check_format(_subNode bytes32) constant returns(bool)
func (_ProfileRegistrar *ProfileRegistrarCaller) Check_format(opts *bind.CallOpts, _subNode [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProfileRegistrar.contract.Call(opts, out, "check_format", _subNode)
	return *ret0, err
}

// Check_format is a free data retrieval call binding the contract method 0x6b4316db.
//
// Solidity: function check_format(_subNode bytes32) constant returns(bool)
func (_ProfileRegistrar *ProfileRegistrarSession) Check_format(_subNode [32]byte) (bool, error) {
	return _ProfileRegistrar.Contract.Check_format(&_ProfileRegistrar.CallOpts, _subNode)
}

// Check_format is a free data retrieval call binding the contract method 0x6b4316db.
//
// Solidity: function check_format(_subNode bytes32) constant returns(bool)
func (_ProfileRegistrar *ProfileRegistrarCallerSession) Check_format(_subNode [32]byte) (bool, error) {
	return _ProfileRegistrar.Contract.Check_format(&_ProfileRegistrar.CallOpts, _subNode)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ProfileRegistrar *ProfileRegistrarCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfileRegistrar.contract.Call(opts, out, "ens")
	return *ret0, err
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ProfileRegistrar *ProfileRegistrarSession) Ens() (common.Address, error) {
	return _ProfileRegistrar.Contract.Ens(&_ProfileRegistrar.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() constant returns(address)
func (_ProfileRegistrar *ProfileRegistrarCallerSession) Ens() (common.Address, error) {
	return _ProfileRegistrar.Contract.Ens(&_ProfileRegistrar.CallOpts)
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(_subNode bytes32) constant returns(nameHash bytes32)
func (_ProfileRegistrar *ProfileRegistrarCaller) Hash(opts *bind.CallOpts, _subNode [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ProfileRegistrar.contract.Call(opts, out, "hash", _subNode)
	return *ret0, err
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(_subNode bytes32) constant returns(nameHash bytes32)
func (_ProfileRegistrar *ProfileRegistrarSession) Hash(_subNode [32]byte) ([32]byte, error) {
	return _ProfileRegistrar.Contract.Hash(&_ProfileRegistrar.CallOpts, _subNode)
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(_subNode bytes32) constant returns(nameHash bytes32)
func (_ProfileRegistrar *ProfileRegistrarCallerSession) Hash(_subNode [32]byte) ([32]byte, error) {
	return _ProfileRegistrar.Contract.Hash(&_ProfileRegistrar.CallOpts, _subNode)
}

// ModuleName is a free data retrieval call binding the contract method 0x93f0899a.
//
// Solidity: function moduleName() constant returns(string)
func (_ProfileRegistrar *ProfileRegistrarCaller) ModuleName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ProfileRegistrar.contract.Call(opts, out, "moduleName")
	return *ret0, err
}

// ModuleName is a free data retrieval call binding the contract method 0x93f0899a.
//
// Solidity: function moduleName() constant returns(string)
func (_ProfileRegistrar *ProfileRegistrarSession) ModuleName() (string, error) {
	return _ProfileRegistrar.Contract.ModuleName(&_ProfileRegistrar.CallOpts)
}

// ModuleName is a free data retrieval call binding the contract method 0x93f0899a.
//
// Solidity: function moduleName() constant returns(string)
func (_ProfileRegistrar *ProfileRegistrarCallerSession) ModuleName() (string, error) {
	return _ProfileRegistrar.Contract.ModuleName(&_ProfileRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProfileRegistrar *ProfileRegistrarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfileRegistrar.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProfileRegistrar *ProfileRegistrarSession) Owner() (common.Address, error) {
	return _ProfileRegistrar.Contract.Owner(&_ProfileRegistrar.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProfileRegistrar *ProfileRegistrarCallerSession) Owner() (common.Address, error) {
	return _ProfileRegistrar.Contract.Owner(&_ProfileRegistrar.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_ProfileRegistrar *ProfileRegistrarCaller) RootNode(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ProfileRegistrar.contract.Call(opts, out, "rootNode")
	return *ret0, err
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_ProfileRegistrar *ProfileRegistrarSession) RootNode() ([32]byte, error) {
	return _ProfileRegistrar.Contract.RootNode(&_ProfileRegistrar.CallOpts)
}

// RootNode is a free data retrieval call binding the contract method 0xfaff50a8.
//
// Solidity: function rootNode() constant returns(bytes32)
func (_ProfileRegistrar *ProfileRegistrarCallerSession) RootNode() ([32]byte, error) {
	return _ProfileRegistrar.Contract.RootNode(&_ProfileRegistrar.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(uint256)
func (_ProfileRegistrar *ProfileRegistrarCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProfileRegistrar.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(uint256)
func (_ProfileRegistrar *ProfileRegistrarSession) Version() (*big.Int, error) {
	return _ProfileRegistrar.Contract.Version(&_ProfileRegistrar.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(uint256)
func (_ProfileRegistrar *ProfileRegistrarCallerSession) Version() (*big.Int, error) {
	return _ProfileRegistrar.Contract.Version(&_ProfileRegistrar.CallOpts)
}

// AdminRegisterFor is a paid mutator transaction binding the contract method 0xd20fdeec.
//
// Solidity: function adminRegisterFor(_subNode bytes32, _enableDonations bool, _alphaUser address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactor) AdminRegisterFor(opts *bind.TransactOpts, _subNode [32]byte, _enableDonations bool, _alphaUser common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "adminRegisterFor", _subNode, _enableDonations, _alphaUser)
}

// AdminRegisterFor is a paid mutator transaction binding the contract method 0xd20fdeec.
//
// Solidity: function adminRegisterFor(_subNode bytes32, _enableDonations bool, _alphaUser address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarSession) AdminRegisterFor(_subNode [32]byte, _enableDonations bool, _alphaUser common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.AdminRegisterFor(&_ProfileRegistrar.TransactOpts, _subNode, _enableDonations, _alphaUser)
}

// AdminRegisterFor is a paid mutator transaction binding the contract method 0xd20fdeec.
//
// Solidity: function adminRegisterFor(_subNode bytes32, _enableDonations bool, _alphaUser address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) AdminRegisterFor(_subNode [32]byte, _enableDonations bool, _alphaUser common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.AdminRegisterFor(&_ProfileRegistrar.TransactOpts, _subNode, _enableDonations, _alphaUser)
}

// AdminSetSubNode is a paid mutator transaction binding the contract method 0xe3650b30.
//
// Solidity: function adminSetSubNode(_subNode bytes32, _newOwner address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactor) AdminSetSubNode(opts *bind.TransactOpts, _subNode [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "adminSetSubNode", _subNode, _newOwner)
}

// AdminSetSubNode is a paid mutator transaction binding the contract method 0xe3650b30.
//
// Solidity: function adminSetSubNode(_subNode bytes32, _newOwner address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarSession) AdminSetSubNode(_subNode [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.AdminSetSubNode(&_ProfileRegistrar.TransactOpts, _subNode, _newOwner)
}

// AdminSetSubNode is a paid mutator transaction binding the contract method 0xe3650b30.
//
// Solidity: function adminSetSubNode(_subNode bytes32, _newOwner address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) AdminSetSubNode(_subNode [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.AdminSetSubNode(&_ProfileRegistrar.TransactOpts, _subNode, _newOwner)
}

// ChangeRootOwner is a paid mutator transaction binding the contract method 0xb5f5aad2.
//
// Solidity: function changeRootOwner(_newOwner address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactor) ChangeRootOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "changeRootOwner", _newOwner)
}

// ChangeRootOwner is a paid mutator transaction binding the contract method 0xb5f5aad2.
//
// Solidity: function changeRootOwner(_newOwner address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarSession) ChangeRootOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ChangeRootOwner(&_ProfileRegistrar.TransactOpts, _newOwner)
}

// ChangeRootOwner is a paid mutator transaction binding the contract method 0xb5f5aad2.
//
// Solidity: function changeRootOwner(_newOwner address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) ChangeRootOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ChangeRootOwner(&_ProfileRegistrar.TransactOpts, _newOwner)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ProfileRegistrar *ProfileRegistrarTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ProfileRegistrar *ProfileRegistrarSession) Destroy() (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.Destroy(&_ProfileRegistrar.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) Destroy() (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.Destroy(&_ProfileRegistrar.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_ProfileRegistrar *ProfileRegistrarTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_ProfileRegistrar *ProfileRegistrarSession) ReclaimEther() (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ReclaimEther(&_ProfileRegistrar.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ReclaimEther(&_ProfileRegistrar.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_ProfileRegistrar *ProfileRegistrarSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ReclaimToken(&_ProfileRegistrar.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.ReclaimToken(&_ProfileRegistrar.TransactOpts, token)
}

// Register is a paid mutator transaction binding the contract method 0x3a4217ba.
//
// Solidity: function register(_subNode bytes32, _enableDonations bool, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactor) Register(opts *bind.TransactOpts, _subNode [32]byte, _enableDonations bool, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "register", _subNode, _enableDonations, _hash, _fn, _digestSize)
}

// Register is a paid mutator transaction binding the contract method 0x3a4217ba.
//
// Solidity: function register(_subNode bytes32, _enableDonations bool, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarSession) Register(_subNode [32]byte, _enableDonations bool, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.Register(&_ProfileRegistrar.TransactOpts, _subNode, _enableDonations, _hash, _fn, _digestSize)
}

// Register is a paid mutator transaction binding the contract method 0x3a4217ba.
//
// Solidity: function register(_subNode bytes32, _enableDonations bool, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) Register(_subNode [32]byte, _enableDonations bool, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.Register(&_ProfileRegistrar.TransactOpts, _subNode, _enableDonations, _hash, _fn, _digestSize)
}

// SetEns is a paid mutator transaction binding the contract method 0x6e8f2be0.
//
// Solidity: function setEns(_ens address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactor) SetEns(opts *bind.TransactOpts, _ens common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "setEns", _ens)
}

// SetEns is a paid mutator transaction binding the contract method 0x6e8f2be0.
//
// Solidity: function setEns(_ens address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarSession) SetEns(_ens common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.SetEns(&_ProfileRegistrar.TransactOpts, _ens)
}

// SetEns is a paid mutator transaction binding the contract method 0x6e8f2be0.
//
// Solidity: function setEns(_ens address) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) SetEns(_ens common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.SetEns(&_ProfileRegistrar.TransactOpts, _ens)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(_resolver address) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactor) SetResolver(opts *bind.TransactOpts, _resolver common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "setResolver", _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(_resolver address) returns()
func (_ProfileRegistrar *ProfileRegistrarSession) SetResolver(_resolver common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.SetResolver(&_ProfileRegistrar.TransactOpts, _resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(_resolver address) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) SetResolver(_resolver common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.SetResolver(&_ProfileRegistrar.TransactOpts, _resolver)
}

// SetRootNode is a paid mutator transaction binding the contract method 0xd20ab1d7.
//
// Solidity: function setRootNode(_newRoot bytes32) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactor) SetRootNode(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "setRootNode", _newRoot)
}

// SetRootNode is a paid mutator transaction binding the contract method 0xd20ab1d7.
//
// Solidity: function setRootNode(_newRoot bytes32) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarSession) SetRootNode(_newRoot [32]byte) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.SetRootNode(&_ProfileRegistrar.TransactOpts, _newRoot)
}

// SetRootNode is a paid mutator transaction binding the contract method 0xd20ab1d7.
//
// Solidity: function setRootNode(_newRoot bytes32) returns(bool)
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) SetRootNode(_newRoot [32]byte) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.SetRootNode(&_ProfileRegistrar.TransactOpts, _newRoot)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_ProfileRegistrar *ProfileRegistrarSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.TokenFallback(&_ProfileRegistrar.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.TokenFallback(&_ProfileRegistrar.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ProfileRegistrar *ProfileRegistrarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.TransferOwnership(&_ProfileRegistrar.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ProfileRegistrar *ProfileRegistrarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProfileRegistrar.Contract.TransferOwnership(&_ProfileRegistrar.TransactOpts, newOwner)
}

// ProfileRegistrarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProfileRegistrar contract.
type ProfileRegistrarOwnershipTransferredIterator struct {
	Event *ProfileRegistrarOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProfileRegistrarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileRegistrarOwnershipTransferred)
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
		it.Event = new(ProfileRegistrarOwnershipTransferred)
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
func (it *ProfileRegistrarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileRegistrarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileRegistrarOwnershipTransferred represents a OwnershipTransferred event raised by the ProfileRegistrar contract.
type ProfileRegistrarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProfileRegistrar *ProfileRegistrarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProfileRegistrarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProfileRegistrar.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistrarOwnershipTransferredIterator{contract: _ProfileRegistrar.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProfileRegistrar *ProfileRegistrarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProfileRegistrarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProfileRegistrar.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileRegistrarOwnershipTransferred)
				if err := _ProfileRegistrar.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ProfileRegistrarRegisterIterator is returned from FilterRegister and is used to iterate over the raw logs and unpacked data for Register events raised by the ProfileRegistrar contract.
type ProfileRegistrarRegisterIterator struct {
	Event *ProfileRegistrarRegister // Event containing the contract specifics and raw log

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
func (it *ProfileRegistrarRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileRegistrarRegister)
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
		it.Event = new(ProfileRegistrarRegister)
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
func (it *ProfileRegistrarRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileRegistrarRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileRegistrarRegister represents a Register event raised by the ProfileRegistrar contract.
type ProfileRegistrarRegister struct {
	Label   [32]byte
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegister is a free log retrieval operation binding the contract event 0xd5fa0e9a716b3ec4895a48223ad309e2d3fa5e27f04d8dc9b3c33cc738a50eb0.
//
// Solidity: event Register(label indexed bytes32, version indexed uint256)
func (_ProfileRegistrar *ProfileRegistrarFilterer) FilterRegister(opts *bind.FilterOpts, label [][32]byte, version []*big.Int) (*ProfileRegistrarRegisterIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _ProfileRegistrar.contract.FilterLogs(opts, "Register", labelRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistrarRegisterIterator{contract: _ProfileRegistrar.contract, event: "Register", logs: logs, sub: sub}, nil
}

// WatchRegister is a free log subscription operation binding the contract event 0xd5fa0e9a716b3ec4895a48223ad309e2d3fa5e27f04d8dc9b3c33cc738a50eb0.
//
// Solidity: event Register(label indexed bytes32, version indexed uint256)
func (_ProfileRegistrar *ProfileRegistrarFilterer) WatchRegister(opts *bind.WatchOpts, sink chan<- *ProfileRegistrarRegister, label [][32]byte, version []*big.Int) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _ProfileRegistrar.contract.WatchLogs(opts, "Register", labelRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileRegistrarRegister)
				if err := _ProfileRegistrar.contract.UnpackLog(event, "Register", log); err != nil {
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
