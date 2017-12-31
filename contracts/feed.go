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

// FeedABI is the input ABI used to generate the binding from.
const FeedABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"unFollow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"follow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_follower\",\"type\":\"address\"},{\"name\":\"_following\",\"type\":\"address\"}],\"name\":\"follows\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"totalFollowing\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"totalFollowers\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"followsCount\",\"outputs\":[{\"name\":\"_followersCount\",\"type\":\"uint256\"},{\"name\":\"_followingCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"followed\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"follower\",\"type\":\"address\"}],\"name\":\"Follow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"followed\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"follower\",\"type\":\"address\"}],\"name\":\"UnFollow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Feed is an auto generated Go binding around an Ethereum contract.
type Feed struct {
	FeedCaller     // Read-only binding to the contract
	FeedTransactor // Write-only binding to the contract
}

// FeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeedSession struct {
	Contract     *Feed             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeedCallerSession struct {
	Contract *FeedCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeedTransactorSession struct {
	Contract     *FeedTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeedRaw struct {
	Contract *Feed // Generic contract binding to access the raw methods on
}

// FeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeedCallerRaw struct {
	Contract *FeedCaller // Generic read-only contract binding to access the raw methods on
}

// FeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeedTransactorRaw struct {
	Contract *FeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeed creates a new instance of Feed, bound to a specific deployed contract.
func NewFeed(address common.Address, backend bind.ContractBackend) (*Feed, error) {
	contract, err := bindFeed(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Feed{FeedCaller: FeedCaller{contract: contract}, FeedTransactor: FeedTransactor{contract: contract}}, nil
}

// NewFeedCaller creates a new read-only instance of Feed, bound to a specific deployed contract.
func NewFeedCaller(address common.Address, caller bind.ContractCaller) (*FeedCaller, error) {
	contract, err := bindFeed(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &FeedCaller{contract: contract}, nil
}

// NewFeedTransactor creates a new write-only instance of Feed, bound to a specific deployed contract.
func NewFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*FeedTransactor, error) {
	contract, err := bindFeed(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &FeedTransactor{contract: contract}, nil
}

// bindFeed binds a generic wrapper to an already deployed contract.
func bindFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feed *FeedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Feed.Contract.FeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feed *FeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feed.Contract.FeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feed *FeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feed.Contract.FeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Feed *FeedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Feed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Feed *FeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Feed *FeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Feed.Contract.contract.Transact(opts, method, params...)
}

// Follows is a free data retrieval call binding the contract method 0x5d1ca476.
//
// Solidity: function follows(_follower address, _following address) constant returns(bool)
func (_Feed *FeedCaller) Follows(opts *bind.CallOpts, _follower common.Address, _following common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Feed.contract.Call(opts, out, "follows", _follower, _following)
	return *ret0, err
}

// Follows is a free data retrieval call binding the contract method 0x5d1ca476.
//
// Solidity: function follows(_follower address, _following address) constant returns(bool)
func (_Feed *FeedSession) Follows(_follower common.Address, _following common.Address) (bool, error) {
	return _Feed.Contract.Follows(&_Feed.CallOpts, _follower, _following)
}

// Follows is a free data retrieval call binding the contract method 0x5d1ca476.
//
// Solidity: function follows(_follower address, _following address) constant returns(bool)
func (_Feed *FeedCallerSession) Follows(_follower common.Address, _following common.Address) (bool, error) {
	return _Feed.Contract.Follows(&_Feed.CallOpts, _follower, _following)
}

// FollowsCount is a free data retrieval call binding the contract method 0xbb5f4d34.
//
// Solidity: function followsCount(_profile address) constant returns(_followersCount uint256, _followingCount uint256)
func (_Feed *FeedCaller) FollowsCount(opts *bind.CallOpts, _profile common.Address) (struct {
	FollowersCount *big.Int
	FollowingCount *big.Int
}, error) {
	ret := new(struct {
		FollowersCount *big.Int
		FollowingCount *big.Int
	})
	out := ret
	err := _Feed.contract.Call(opts, out, "followsCount", _profile)
	return *ret, err
}

// FollowsCount is a free data retrieval call binding the contract method 0xbb5f4d34.
//
// Solidity: function followsCount(_profile address) constant returns(_followersCount uint256, _followingCount uint256)
func (_Feed *FeedSession) FollowsCount(_profile common.Address) (struct {
	FollowersCount *big.Int
	FollowingCount *big.Int
}, error) {
	return _Feed.Contract.FollowsCount(&_Feed.CallOpts, _profile)
}

// FollowsCount is a free data retrieval call binding the contract method 0xbb5f4d34.
//
// Solidity: function followsCount(_profile address) constant returns(_followersCount uint256, _followingCount uint256)
func (_Feed *FeedCallerSession) FollowsCount(_profile common.Address) (struct {
	FollowersCount *big.Int
	FollowingCount *big.Int
}, error) {
	return _Feed.Contract.FollowsCount(&_Feed.CallOpts, _profile)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feed *FeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Feed.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feed *FeedSession) Owner() (common.Address, error) {
	return _Feed.Contract.Owner(&_Feed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Feed *FeedCallerSession) Owner() (common.Address, error) {
	return _Feed.Contract.Owner(&_Feed.CallOpts)
}

// TotalFollowers is a free data retrieval call binding the contract method 0xa0aba3a1.
//
// Solidity: function totalFollowers(_profile address) constant returns(_total uint256)
func (_Feed *FeedCaller) TotalFollowers(opts *bind.CallOpts, _profile common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feed.contract.Call(opts, out, "totalFollowers", _profile)
	return *ret0, err
}

// TotalFollowers is a free data retrieval call binding the contract method 0xa0aba3a1.
//
// Solidity: function totalFollowers(_profile address) constant returns(_total uint256)
func (_Feed *FeedSession) TotalFollowers(_profile common.Address) (*big.Int, error) {
	return _Feed.Contract.TotalFollowers(&_Feed.CallOpts, _profile)
}

// TotalFollowers is a free data retrieval call binding the contract method 0xa0aba3a1.
//
// Solidity: function totalFollowers(_profile address) constant returns(_total uint256)
func (_Feed *FeedCallerSession) TotalFollowers(_profile common.Address) (*big.Int, error) {
	return _Feed.Contract.TotalFollowers(&_Feed.CallOpts, _profile)
}

// TotalFollowing is a free data retrieval call binding the contract method 0x8dea1cb8.
//
// Solidity: function totalFollowing(_profile address) constant returns(_total uint256)
func (_Feed *FeedCaller) TotalFollowing(opts *bind.CallOpts, _profile common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Feed.contract.Call(opts, out, "totalFollowing", _profile)
	return *ret0, err
}

// TotalFollowing is a free data retrieval call binding the contract method 0x8dea1cb8.
//
// Solidity: function totalFollowing(_profile address) constant returns(_total uint256)
func (_Feed *FeedSession) TotalFollowing(_profile common.Address) (*big.Int, error) {
	return _Feed.Contract.TotalFollowing(&_Feed.CallOpts, _profile)
}

// TotalFollowing is a free data retrieval call binding the contract method 0x8dea1cb8.
//
// Solidity: function totalFollowing(_profile address) constant returns(_total uint256)
func (_Feed *FeedCallerSession) TotalFollowing(_profile common.Address) (*big.Int, error) {
	return _Feed.Contract.TotalFollowing(&_Feed.CallOpts, _profile)
}

// Follow is a paid mutator transaction binding the contract method 0x4dbf27cc.
//
// Solidity: function follow(_profile address) returns(bool)
func (_Feed *FeedTransactor) Follow(opts *bind.TransactOpts, _profile common.Address) (*types.Transaction, error) {
	return _Feed.contract.Transact(opts, "follow", _profile)
}

// Follow is a paid mutator transaction binding the contract method 0x4dbf27cc.
//
// Solidity: function follow(_profile address) returns(bool)
func (_Feed *FeedSession) Follow(_profile common.Address) (*types.Transaction, error) {
	return _Feed.Contract.Follow(&_Feed.TransactOpts, _profile)
}

// Follow is a paid mutator transaction binding the contract method 0x4dbf27cc.
//
// Solidity: function follow(_profile address) returns(bool)
func (_Feed *FeedTransactorSession) Follow(_profile common.Address) (*types.Transaction, error) {
	return _Feed.Contract.Follow(&_Feed.TransactOpts, _profile)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Feed *FeedTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Feed.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Feed *FeedSession) ReclaimEther() (*types.Transaction, error) {
	return _Feed.Contract.ReclaimEther(&_Feed.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Feed *FeedTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _Feed.Contract.ReclaimEther(&_Feed.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Feed *FeedTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Feed.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Feed *FeedSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Feed.Contract.ReclaimToken(&_Feed.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Feed *FeedTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Feed.Contract.ReclaimToken(&_Feed.TransactOpts, token)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Feed *FeedTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Feed.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Feed *FeedSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Feed.Contract.TokenFallback(&_Feed.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Feed *FeedTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Feed.Contract.TokenFallback(&_Feed.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Feed *FeedTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Feed.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Feed *FeedSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Feed.Contract.TransferOwnership(&_Feed.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Feed *FeedTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Feed.Contract.TransferOwnership(&_Feed.TransactOpts, newOwner)
}

// UnFollow is a paid mutator transaction binding the contract method 0x034889ff.
//
// Solidity: function unFollow(_profile address) returns(bool)
func (_Feed *FeedTransactor) UnFollow(opts *bind.TransactOpts, _profile common.Address) (*types.Transaction, error) {
	return _Feed.contract.Transact(opts, "unFollow", _profile)
}

// UnFollow is a paid mutator transaction binding the contract method 0x034889ff.
//
// Solidity: function unFollow(_profile address) returns(bool)
func (_Feed *FeedSession) UnFollow(_profile common.Address) (*types.Transaction, error) {
	return _Feed.Contract.UnFollow(&_Feed.TransactOpts, _profile)
}

// UnFollow is a paid mutator transaction binding the contract method 0x034889ff.
//
// Solidity: function unFollow(_profile address) returns(bool)
func (_Feed *FeedTransactorSession) UnFollow(_profile common.Address) (*types.Transaction, error) {
	return _Feed.Contract.UnFollow(&_Feed.TransactOpts, _profile)
}
