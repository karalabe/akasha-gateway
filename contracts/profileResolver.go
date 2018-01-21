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

// ProfileResolverABI is the input ABI used to generate the binding from.
const ProfileResolverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_nodeHash\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"}],\"name\":\"setHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_akashaId\",\"type\":\"bytes32\"},{\"name\":\"_nodeHash\",\"type\":\"bytes32\"},{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"bool\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"}],\"name\":\"registerHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"moduleAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_moduleAddress\",\"type\":\"address\"}],\"name\":\"setModule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[{\"name\":\"_akashaId\",\"type\":\"bytes32\"},{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_donationsEnabled\",\"type\":\"bool\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"removeProfile\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"},{\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"toggleDonations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgradeController\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_moduleAddress\",\"type\":\"address\"}],\"name\":\"updateModule\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"donationsEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalProfiles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\"},{\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"bytes32\"}],\"name\":\"hash\",\"outputs\":[{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"reverse\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"disabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ens\",\"type\":\"address\"},{\"name\":\"_totalProfiles\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ProfileResolver is an auto generated Go binding around an Ethereum contract.
type ProfileResolver struct {
	ProfileResolverCaller     // Read-only binding to the contract
	ProfileResolverTransactor // Write-only binding to the contract
	ProfileResolverFilterer   // Log filterer for contract events
}

// ProfileResolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileResolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileResolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileResolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileResolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileResolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileResolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileResolverSession struct {
	Contract     *ProfileResolver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileResolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileResolverCallerSession struct {
	Contract *ProfileResolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProfileResolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileResolverTransactorSession struct {
	Contract     *ProfileResolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProfileResolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileResolverRaw struct {
	Contract *ProfileResolver // Generic contract binding to access the raw methods on
}

// ProfileResolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileResolverCallerRaw struct {
	Contract *ProfileResolverCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileResolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileResolverTransactorRaw struct {
	Contract *ProfileResolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfileResolver creates a new instance of ProfileResolver, bound to a specific deployed contract.
func NewProfileResolver(address common.Address, backend bind.ContractBackend) (*ProfileResolver, error) {
	contract, err := bindProfileResolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProfileResolver{ProfileResolverCaller: ProfileResolverCaller{contract: contract}, ProfileResolverTransactor: ProfileResolverTransactor{contract: contract}, ProfileResolverFilterer: ProfileResolverFilterer{contract: contract}}, nil
}

// NewProfileResolverCaller creates a new read-only instance of ProfileResolver, bound to a specific deployed contract.
func NewProfileResolverCaller(address common.Address, caller bind.ContractCaller) (*ProfileResolverCaller, error) {
	contract, err := bindProfileResolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileResolverCaller{contract: contract}, nil
}

// NewProfileResolverTransactor creates a new write-only instance of ProfileResolver, bound to a specific deployed contract.
func NewProfileResolverTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileResolverTransactor, error) {
	contract, err := bindProfileResolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileResolverTransactor{contract: contract}, nil
}

// NewProfileResolverFilterer creates a new log filterer instance of ProfileResolver, bound to a specific deployed contract.
func NewProfileResolverFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileResolverFilterer, error) {
	contract, err := bindProfileResolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileResolverFilterer{contract: contract}, nil
}

// bindProfileResolver binds a generic wrapper to an already deployed contract.
func bindProfileResolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProfileResolverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfileResolver *ProfileResolverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProfileResolver.Contract.ProfileResolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfileResolver *ProfileResolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileResolver.Contract.ProfileResolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfileResolver *ProfileResolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfileResolver.Contract.ProfileResolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfileResolver *ProfileResolverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ProfileResolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfileResolver *ProfileResolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileResolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfileResolver *ProfileResolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfileResolver.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(address)
func (_ProfileResolver *ProfileResolverCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfileResolver.contract.Call(opts, out, "addr", node)
	return *ret0, err
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(address)
func (_ProfileResolver *ProfileResolverSession) Addr(node [32]byte) (common.Address, error) {
	return _ProfileResolver.Contract.Addr(&_ProfileResolver.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(node bytes32) constant returns(address)
func (_ProfileResolver *ProfileResolverCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _ProfileResolver.Contract.Addr(&_ProfileResolver.CallOpts, node)
}

// Disabled is a free data retrieval call binding the contract method 0xee070805.
//
// Solidity: function disabled() constant returns(bool)
func (_ProfileResolver *ProfileResolverCaller) Disabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProfileResolver.contract.Call(opts, out, "disabled")
	return *ret0, err
}

// Disabled is a free data retrieval call binding the contract method 0xee070805.
//
// Solidity: function disabled() constant returns(bool)
func (_ProfileResolver *ProfileResolverSession) Disabled() (bool, error) {
	return _ProfileResolver.Contract.Disabled(&_ProfileResolver.CallOpts)
}

// Disabled is a free data retrieval call binding the contract method 0xee070805.
//
// Solidity: function disabled() constant returns(bool)
func (_ProfileResolver *ProfileResolverCallerSession) Disabled() (bool, error) {
	return _ProfileResolver.Contract.Disabled(&_ProfileResolver.CallOpts)
}

// DonationsEnabled is a free data retrieval call binding the contract method 0xc59315bf.
//
// Solidity: function donationsEnabled(_node bytes32) constant returns(bool)
func (_ProfileResolver *ProfileResolverCaller) DonationsEnabled(opts *bind.CallOpts, _node [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ProfileResolver.contract.Call(opts, out, "donationsEnabled", _node)
	return *ret0, err
}

// DonationsEnabled is a free data retrieval call binding the contract method 0xc59315bf.
//
// Solidity: function donationsEnabled(_node bytes32) constant returns(bool)
func (_ProfileResolver *ProfileResolverSession) DonationsEnabled(_node [32]byte) (bool, error) {
	return _ProfileResolver.Contract.DonationsEnabled(&_ProfileResolver.CallOpts, _node)
}

// DonationsEnabled is a free data retrieval call binding the contract method 0xc59315bf.
//
// Solidity: function donationsEnabled(_node bytes32) constant returns(bool)
func (_ProfileResolver *ProfileResolverCallerSession) DonationsEnabled(_node [32]byte) (bool, error) {
	return _ProfileResolver.Contract.DonationsEnabled(&_ProfileResolver.CallOpts, _node)
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(_node bytes32) constant returns(_fn uint8, _digestSize uint8, _hash bytes32)
func (_ProfileResolver *ProfileResolverCaller) Hash(opts *bind.CallOpts, _node [32]byte) (struct {
	Fn         uint8
	DigestSize uint8
	Hash       [32]byte
}, error) {
	ret := new(struct {
		Fn         uint8
		DigestSize uint8
		Hash       [32]byte
	})
	out := ret
	err := _ProfileResolver.contract.Call(opts, out, "hash", _node)
	return *ret, err
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(_node bytes32) constant returns(_fn uint8, _digestSize uint8, _hash bytes32)
func (_ProfileResolver *ProfileResolverSession) Hash(_node [32]byte) (struct {
	Fn         uint8
	DigestSize uint8
	Hash       [32]byte
}, error) {
	return _ProfileResolver.Contract.Hash(&_ProfileResolver.CallOpts, _node)
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(_node bytes32) constant returns(_fn uint8, _digestSize uint8, _hash bytes32)
func (_ProfileResolver *ProfileResolverCallerSession) Hash(_node [32]byte) (struct {
	Fn         uint8
	DigestSize uint8
	Hash       [32]byte
}, error) {
	return _ProfileResolver.Contract.Hash(&_ProfileResolver.CallOpts, _node)
}

// ModuleAddress is a free data retrieval call binding the contract method 0x35b2bd2d.
//
// Solidity: function moduleAddress() constant returns(address)
func (_ProfileResolver *ProfileResolverCaller) ModuleAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfileResolver.contract.Call(opts, out, "moduleAddress")
	return *ret0, err
}

// ModuleAddress is a free data retrieval call binding the contract method 0x35b2bd2d.
//
// Solidity: function moduleAddress() constant returns(address)
func (_ProfileResolver *ProfileResolverSession) ModuleAddress() (common.Address, error) {
	return _ProfileResolver.Contract.ModuleAddress(&_ProfileResolver.CallOpts)
}

// ModuleAddress is a free data retrieval call binding the contract method 0x35b2bd2d.
//
// Solidity: function moduleAddress() constant returns(address)
func (_ProfileResolver *ProfileResolverCallerSession) ModuleAddress() (common.Address, error) {
	return _ProfileResolver.Contract.ModuleAddress(&_ProfileResolver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProfileResolver *ProfileResolverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ProfileResolver.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProfileResolver *ProfileResolverSession) Owner() (common.Address, error) {
	return _ProfileResolver.Contract.Owner(&_ProfileResolver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ProfileResolver *ProfileResolverCallerSession) Owner() (common.Address, error) {
	return _ProfileResolver.Contract.Owner(&_ProfileResolver.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(_node bytes32) constant returns(_akashaId bytes32, _addr address, _donationsEnabled bool, _fn uint8, _digestSize uint8, _hash bytes32)
func (_ProfileResolver *ProfileResolverCaller) Resolve(opts *bind.CallOpts, _node [32]byte) (struct {
	AkashaId         [32]byte
	Addr             common.Address
	DonationsEnabled bool
	Fn               uint8
	DigestSize       uint8
	Hash             [32]byte
}, error) {
	ret := new(struct {
		AkashaId         [32]byte
		Addr             common.Address
		DonationsEnabled bool
		Fn               uint8
		DigestSize       uint8
		Hash             [32]byte
	})
	out := ret
	err := _ProfileResolver.contract.Call(opts, out, "resolve", _node)
	return *ret, err
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(_node bytes32) constant returns(_akashaId bytes32, _addr address, _donationsEnabled bool, _fn uint8, _digestSize uint8, _hash bytes32)
func (_ProfileResolver *ProfileResolverSession) Resolve(_node [32]byte) (struct {
	AkashaId         [32]byte
	Addr             common.Address
	DonationsEnabled bool
	Fn               uint8
	DigestSize       uint8
	Hash             [32]byte
}, error) {
	return _ProfileResolver.Contract.Resolve(&_ProfileResolver.CallOpts, _node)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(_node bytes32) constant returns(_akashaId bytes32, _addr address, _donationsEnabled bool, _fn uint8, _digestSize uint8, _hash bytes32)
func (_ProfileResolver *ProfileResolverCallerSession) Resolve(_node [32]byte) (struct {
	AkashaId         [32]byte
	Addr             common.Address
	DonationsEnabled bool
	Fn               uint8
	DigestSize       uint8
	Hash             [32]byte
}, error) {
	return _ProfileResolver.Contract.Resolve(&_ProfileResolver.CallOpts, _node)
}

// Reverse is a free data retrieval call binding the contract method 0xe30bd740.
//
// Solidity: function reverse(owner address) constant returns(bytes32)
func (_ProfileResolver *ProfileResolverCaller) Reverse(opts *bind.CallOpts, owner common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ProfileResolver.contract.Call(opts, out, "reverse", owner)
	return *ret0, err
}

// Reverse is a free data retrieval call binding the contract method 0xe30bd740.
//
// Solidity: function reverse(owner address) constant returns(bytes32)
func (_ProfileResolver *ProfileResolverSession) Reverse(owner common.Address) ([32]byte, error) {
	return _ProfileResolver.Contract.Reverse(&_ProfileResolver.CallOpts, owner)
}

// Reverse is a free data retrieval call binding the contract method 0xe30bd740.
//
// Solidity: function reverse(owner address) constant returns(bytes32)
func (_ProfileResolver *ProfileResolverCallerSession) Reverse(owner common.Address) ([32]byte, error) {
	return _ProfileResolver.Contract.Reverse(&_ProfileResolver.CallOpts, owner)
}

// TotalProfiles is a free data retrieval call binding the contract method 0xd2adbe79.
//
// Solidity: function totalProfiles() constant returns(uint256)
func (_ProfileResolver *ProfileResolverCaller) TotalProfiles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ProfileResolver.contract.Call(opts, out, "totalProfiles")
	return *ret0, err
}

// TotalProfiles is a free data retrieval call binding the contract method 0xd2adbe79.
//
// Solidity: function totalProfiles() constant returns(uint256)
func (_ProfileResolver *ProfileResolverSession) TotalProfiles() (*big.Int, error) {
	return _ProfileResolver.Contract.TotalProfiles(&_ProfileResolver.CallOpts)
}

// TotalProfiles is a free data retrieval call binding the contract method 0xd2adbe79.
//
// Solidity: function totalProfiles() constant returns(uint256)
func (_ProfileResolver *ProfileResolverCallerSession) TotalProfiles() (*big.Int, error) {
	return _ProfileResolver.Contract.TotalProfiles(&_ProfileResolver.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ProfileResolver *ProfileResolverTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ProfileResolver *ProfileResolverSession) Destroy() (*types.Transaction, error) {
	return _ProfileResolver.Contract.Destroy(&_ProfileResolver.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_ProfileResolver *ProfileResolverTransactorSession) Destroy() (*types.Transaction, error) {
	return _ProfileResolver.Contract.Destroy(&_ProfileResolver.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_ProfileResolver *ProfileResolverTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_ProfileResolver *ProfileResolverSession) ReclaimEther() (*types.Transaction, error) {
	return _ProfileResolver.Contract.ReclaimEther(&_ProfileResolver.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_ProfileResolver *ProfileResolverTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _ProfileResolver.Contract.ReclaimEther(&_ProfileResolver.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_ProfileResolver *ProfileResolverTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_ProfileResolver *ProfileResolverSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.ReclaimToken(&_ProfileResolver.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_ProfileResolver *ProfileResolverTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.ReclaimToken(&_ProfileResolver.TransactOpts, token)
}

// RegisterHash is a paid mutator transaction binding the contract method 0x17a65f4c.
//
// Solidity: function registerHash(_akashaId bytes32, _nodeHash bytes32, _owner address, _status bool, _hash bytes32, _fn uint8, _digestSize uint8) returns(uint256)
func (_ProfileResolver *ProfileResolverTransactor) RegisterHash(opts *bind.TransactOpts, _akashaId [32]byte, _nodeHash [32]byte, _owner common.Address, _status bool, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "registerHash", _akashaId, _nodeHash, _owner, _status, _hash, _fn, _digestSize)
}

// RegisterHash is a paid mutator transaction binding the contract method 0x17a65f4c.
//
// Solidity: function registerHash(_akashaId bytes32, _nodeHash bytes32, _owner address, _status bool, _hash bytes32, _fn uint8, _digestSize uint8) returns(uint256)
func (_ProfileResolver *ProfileResolverSession) RegisterHash(_akashaId [32]byte, _nodeHash [32]byte, _owner common.Address, _status bool, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileResolver.Contract.RegisterHash(&_ProfileResolver.TransactOpts, _akashaId, _nodeHash, _owner, _status, _hash, _fn, _digestSize)
}

// RegisterHash is a paid mutator transaction binding the contract method 0x17a65f4c.
//
// Solidity: function registerHash(_akashaId bytes32, _nodeHash bytes32, _owner address, _status bool, _hash bytes32, _fn uint8, _digestSize uint8) returns(uint256)
func (_ProfileResolver *ProfileResolverTransactorSession) RegisterHash(_akashaId [32]byte, _nodeHash [32]byte, _owner common.Address, _status bool, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileResolver.Contract.RegisterHash(&_ProfileResolver.TransactOpts, _akashaId, _nodeHash, _owner, _status, _hash, _fn, _digestSize)
}

// RemoveProfile is a paid mutator transaction binding the contract method 0x7340c30a.
//
// Solidity: function removeProfile(_node bytes32) returns(bool)
func (_ProfileResolver *ProfileResolverTransactor) RemoveProfile(opts *bind.TransactOpts, _node [32]byte) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "removeProfile", _node)
}

// RemoveProfile is a paid mutator transaction binding the contract method 0x7340c30a.
//
// Solidity: function removeProfile(_node bytes32) returns(bool)
func (_ProfileResolver *ProfileResolverSession) RemoveProfile(_node [32]byte) (*types.Transaction, error) {
	return _ProfileResolver.Contract.RemoveProfile(&_ProfileResolver.TransactOpts, _node)
}

// RemoveProfile is a paid mutator transaction binding the contract method 0x7340c30a.
//
// Solidity: function removeProfile(_node bytes32) returns(bool)
func (_ProfileResolver *ProfileResolverTransactorSession) RemoveProfile(_node [32]byte) (*types.Transaction, error) {
	return _ProfileResolver.Contract.RemoveProfile(&_ProfileResolver.TransactOpts, _node)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, newAddress address) returns()
func (_ProfileResolver *ProfileResolverTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "setAddr", node, newAddress)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, newAddress address) returns()
func (_ProfileResolver *ProfileResolverSession) SetAddr(node [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.SetAddr(&_ProfileResolver.TransactOpts, node, newAddress)
}

// SetAddr is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(node bytes32, newAddress address) returns()
func (_ProfileResolver *ProfileResolverTransactorSession) SetAddr(node [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.SetAddr(&_ProfileResolver.TransactOpts, node, newAddress)
}

// SetHash is a paid mutator transaction binding the contract method 0x02a77059.
//
// Solidity: function setHash(_nodeHash bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_ProfileResolver *ProfileResolverTransactor) SetHash(opts *bind.TransactOpts, _nodeHash [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "setHash", _nodeHash, _hash, _fn, _digestSize)
}

// SetHash is a paid mutator transaction binding the contract method 0x02a77059.
//
// Solidity: function setHash(_nodeHash bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_ProfileResolver *ProfileResolverSession) SetHash(_nodeHash [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileResolver.Contract.SetHash(&_ProfileResolver.TransactOpts, _nodeHash, _hash, _fn, _digestSize)
}

// SetHash is a paid mutator transaction binding the contract method 0x02a77059.
//
// Solidity: function setHash(_nodeHash bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_ProfileResolver *ProfileResolverTransactorSession) SetHash(_nodeHash [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _ProfileResolver.Contract.SetHash(&_ProfileResolver.TransactOpts, _nodeHash, _hash, _fn, _digestSize)
}

// SetModule is a paid mutator transaction binding the contract method 0x47f543bc.
//
// Solidity: function setModule(_moduleAddress address) returns(bool)
func (_ProfileResolver *ProfileResolverTransactor) SetModule(opts *bind.TransactOpts, _moduleAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "setModule", _moduleAddress)
}

// SetModule is a paid mutator transaction binding the contract method 0x47f543bc.
//
// Solidity: function setModule(_moduleAddress address) returns(bool)
func (_ProfileResolver *ProfileResolverSession) SetModule(_moduleAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.SetModule(&_ProfileResolver.TransactOpts, _moduleAddress)
}

// SetModule is a paid mutator transaction binding the contract method 0x47f543bc.
//
// Solidity: function setModule(_moduleAddress address) returns(bool)
func (_ProfileResolver *ProfileResolverTransactorSession) SetModule(_moduleAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.SetModule(&_ProfileResolver.TransactOpts, _moduleAddress)
}

// ToggleDonations is a paid mutator transaction binding the contract method 0x76b74a78.
//
// Solidity: function toggleDonations(_node bytes32, _status bool) returns(bool)
func (_ProfileResolver *ProfileResolverTransactor) ToggleDonations(opts *bind.TransactOpts, _node [32]byte, _status bool) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "toggleDonations", _node, _status)
}

// ToggleDonations is a paid mutator transaction binding the contract method 0x76b74a78.
//
// Solidity: function toggleDonations(_node bytes32, _status bool) returns(bool)
func (_ProfileResolver *ProfileResolverSession) ToggleDonations(_node [32]byte, _status bool) (*types.Transaction, error) {
	return _ProfileResolver.Contract.ToggleDonations(&_ProfileResolver.TransactOpts, _node, _status)
}

// ToggleDonations is a paid mutator transaction binding the contract method 0x76b74a78.
//
// Solidity: function toggleDonations(_node bytes32, _status bool) returns(bool)
func (_ProfileResolver *ProfileResolverTransactorSession) ToggleDonations(_node [32]byte, _status bool) (*types.Transaction, error) {
	return _ProfileResolver.Contract.ToggleDonations(&_ProfileResolver.TransactOpts, _node, _status)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_ProfileResolver *ProfileResolverTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_ProfileResolver *ProfileResolverSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _ProfileResolver.Contract.TokenFallback(&_ProfileResolver.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_ProfileResolver *ProfileResolverTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _ProfileResolver.Contract.TokenFallback(&_ProfileResolver.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ProfileResolver *ProfileResolverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ProfileResolver *ProfileResolverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.TransferOwnership(&_ProfileResolver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_ProfileResolver *ProfileResolverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.TransferOwnership(&_ProfileResolver.TransactOpts, newOwner)
}

// UpdateModule is a paid mutator transaction binding the contract method 0x8c5f89dc.
//
// Solidity: function updateModule(_moduleAddress address) returns(bool)
func (_ProfileResolver *ProfileResolverTransactor) UpdateModule(opts *bind.TransactOpts, _moduleAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "updateModule", _moduleAddress)
}

// UpdateModule is a paid mutator transaction binding the contract method 0x8c5f89dc.
//
// Solidity: function updateModule(_moduleAddress address) returns(bool)
func (_ProfileResolver *ProfileResolverSession) UpdateModule(_moduleAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.UpdateModule(&_ProfileResolver.TransactOpts, _moduleAddress)
}

// UpdateModule is a paid mutator transaction binding the contract method 0x8c5f89dc.
//
// Solidity: function updateModule(_moduleAddress address) returns(bool)
func (_ProfileResolver *ProfileResolverTransactorSession) UpdateModule(_moduleAddress common.Address) (*types.Transaction, error) {
	return _ProfileResolver.Contract.UpdateModule(&_ProfileResolver.TransactOpts, _moduleAddress)
}

// UpgradeController is a paid mutator transaction binding the contract method 0x87543ef6.
//
// Solidity: function upgradeController() returns(bool)
func (_ProfileResolver *ProfileResolverTransactor) UpgradeController(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileResolver.contract.Transact(opts, "upgradeController")
}

// UpgradeController is a paid mutator transaction binding the contract method 0x87543ef6.
//
// Solidity: function upgradeController() returns(bool)
func (_ProfileResolver *ProfileResolverSession) UpgradeController() (*types.Transaction, error) {
	return _ProfileResolver.Contract.UpgradeController(&_ProfileResolver.TransactOpts)
}

// UpgradeController is a paid mutator transaction binding the contract method 0x87543ef6.
//
// Solidity: function upgradeController() returns(bool)
func (_ProfileResolver *ProfileResolverTransactorSession) UpgradeController() (*types.Transaction, error) {
	return _ProfileResolver.Contract.UpgradeController(&_ProfileResolver.TransactOpts)
}

// ProfileResolverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProfileResolver contract.
type ProfileResolverOwnershipTransferredIterator struct {
	Event *ProfileResolverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProfileResolverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileResolverOwnershipTransferred)
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
		it.Event = new(ProfileResolverOwnershipTransferred)
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

// Error retruned any retrieval or parsing error occurred during filtering.
func (it *ProfileResolverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileResolverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileResolverOwnershipTransferred represents a OwnershipTransferred event raised by the ProfileResolver contract.
type ProfileResolverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProfileResolver *ProfileResolverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProfileResolverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProfileResolver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProfileResolverOwnershipTransferredIterator{contract: _ProfileResolver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ProfileResolver *ProfileResolverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProfileResolverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProfileResolver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileResolverOwnershipTransferred)
				if err := _ProfileResolver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
