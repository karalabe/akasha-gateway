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

// FeedABI is the input ABI used to generate the binding from.
const FeedABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"unFollow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"follow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_follower\",\"type\":\"address\"},{\"name\":\"_following\",\"type\":\"address\"}],\"name\":\"follows\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"totalFollowing\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"totalFollowers\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_profile\",\"type\":\"address\"}],\"name\":\"followsCount\",\"outputs\":[{\"name\":\"_followersCount\",\"type\":\"uint256\"},{\"name\":\"_followingCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"followed\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"follower\",\"type\":\"address\"}],\"name\":\"Follow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"followed\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"follower\",\"type\":\"address\"}],\"name\":\"UnFollow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Feed is an auto generated Go binding around an Ethereum contract.
type Feed struct {
	FeedCaller     // Read-only binding to the contract
	FeedTransactor // Write-only binding to the contract
	FeedFilterer   // Log filterer for contract events
}

// FeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeedFilterer struct {
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
	contract, err := bindFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Feed{FeedCaller: FeedCaller{contract: contract}, FeedTransactor: FeedTransactor{contract: contract}, FeedFilterer: FeedFilterer{contract: contract}}, nil
}

// NewFeedCaller creates a new read-only instance of Feed, bound to a specific deployed contract.
func NewFeedCaller(address common.Address, caller bind.ContractCaller) (*FeedCaller, error) {
	contract, err := bindFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeedCaller{contract: contract}, nil
}

// NewFeedTransactor creates a new write-only instance of Feed, bound to a specific deployed contract.
func NewFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*FeedTransactor, error) {
	contract, err := bindFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeedTransactor{contract: contract}, nil
}

// NewFeedFilterer creates a new log filterer instance of Feed, bound to a specific deployed contract.
func NewFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*FeedFilterer, error) {
	contract, err := bindFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeedFilterer{contract: contract}, nil
}

// bindFeed binds a generic wrapper to an already deployed contract.
func bindFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// FeedFollowIterator is returned from FilterFollow and is used to iterate over the raw logs and unpacked data for Follow events raised by the Feed contract.
type FeedFollowIterator struct {
	Event *FeedFollow // Event containing the contract specifics and raw log

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
func (it *FeedFollowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedFollow)
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
		it.Event = new(FeedFollow)
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
func (it *FeedFollowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedFollowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedFollow represents a Follow event raised by the Feed contract.
type FeedFollow struct {
	Followed common.Address
	Follower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFollow is a free log retrieval operation binding the contract event 0xbccc71dc7842b86291138666aa18e133ee6d41aa71e6d7c650debad1a0576635.
//
// Solidity: event Follow(followed indexed address, follower indexed address)
func (_Feed *FeedFilterer) FilterFollow(opts *bind.FilterOpts, followed []common.Address, follower []common.Address) (*FeedFollowIterator, error) {

	var followedRule []interface{}
	for _, followedItem := range followed {
		followedRule = append(followedRule, followedItem)
	}
	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Feed.contract.FilterLogs(opts, "Follow", followedRule, followerRule)
	if err != nil {
		return nil, err
	}
	return &FeedFollowIterator{contract: _Feed.contract, event: "Follow", logs: logs, sub: sub}, nil
}

// WatchFollow is a free log subscription operation binding the contract event 0xbccc71dc7842b86291138666aa18e133ee6d41aa71e6d7c650debad1a0576635.
//
// Solidity: event Follow(followed indexed address, follower indexed address)
func (_Feed *FeedFilterer) WatchFollow(opts *bind.WatchOpts, sink chan<- *FeedFollow, followed []common.Address, follower []common.Address) (event.Subscription, error) {

	var followedRule []interface{}
	for _, followedItem := range followed {
		followedRule = append(followedRule, followedItem)
	}
	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Feed.contract.WatchLogs(opts, "Follow", followedRule, followerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedFollow)
				if err := _Feed.contract.UnpackLog(event, "Follow", log); err != nil {
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

// FeedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Feed contract.
type FeedOwnershipTransferredIterator struct {
	Event *FeedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedOwnershipTransferred)
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
		it.Event = new(FeedOwnershipTransferred)
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
func (it *FeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedOwnershipTransferred represents a OwnershipTransferred event raised by the Feed contract.
type FeedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Feed *FeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Feed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeedOwnershipTransferredIterator{contract: _Feed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Feed *FeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Feed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedOwnershipTransferred)
				if err := _Feed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// FeedUnFollowIterator is returned from FilterUnFollow and is used to iterate over the raw logs and unpacked data for UnFollow events raised by the Feed contract.
type FeedUnFollowIterator struct {
	Event *FeedUnFollow // Event containing the contract specifics and raw log

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
func (it *FeedUnFollowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedUnFollow)
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
		it.Event = new(FeedUnFollow)
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
func (it *FeedUnFollowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedUnFollowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedUnFollow represents a UnFollow event raised by the Feed contract.
type FeedUnFollow struct {
	Followed common.Address
	Follower common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnFollow is a free log retrieval operation binding the contract event 0xe5dcccfe8e7890c6d4aa94d44c3b409fb6f023337f29c3308077c0c4068192c4.
//
// Solidity: event UnFollow(followed indexed address, follower indexed address)
func (_Feed *FeedFilterer) FilterUnFollow(opts *bind.FilterOpts, followed []common.Address, follower []common.Address) (*FeedUnFollowIterator, error) {

	var followedRule []interface{}
	for _, followedItem := range followed {
		followedRule = append(followedRule, followedItem)
	}
	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Feed.contract.FilterLogs(opts, "UnFollow", followedRule, followerRule)
	if err != nil {
		return nil, err
	}
	return &FeedUnFollowIterator{contract: _Feed.contract, event: "UnFollow", logs: logs, sub: sub}, nil
}

// WatchUnFollow is a free log subscription operation binding the contract event 0xe5dcccfe8e7890c6d4aa94d44c3b409fb6f023337f29c3308077c0c4068192c4.
//
// Solidity: event UnFollow(followed indexed address, follower indexed address)
func (_Feed *FeedFilterer) WatchUnFollow(opts *bind.WatchOpts, sink chan<- *FeedUnFollow, followed []common.Address, follower []common.Address) (event.Subscription, error) {

	var followedRule []interface{}
	for _, followedItem := range followed {
		followedRule = append(followedRule, followedItem)
	}
	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _Feed.contract.WatchLogs(opts, "UnFollow", followedRule, followerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedUnFollow)
				if err := _Feed.contract.UnpackLog(event, "UnFollow", log); err != nil {
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
