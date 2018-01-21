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

// CommentsABI is the input ABI used to generate the binding from.
const CommentsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_entryAuthor\",\"type\":\"address\"},{\"name\":\"_commentId\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"}],\"name\":\"edit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"}],\"name\":\"totalComments\",\"outputs\":[{\"name\":\"_total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_commentId\",\"type\":\"bytes32\"}],\"name\":\"isDeleted\",\"outputs\":[{\"name\":\"_deleted\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"discount_every\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_votes\",\"type\":\"address\"}],\"name\":\"setVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_entryAuthor\",\"type\":\"address\"},{\"name\":\"_commentId\",\"type\":\"bytes32\"}],\"name\":\"deleteComment\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_commentId\",\"type\":\"bytes32\"}],\"name\":\"getComment\",\"outputs\":[{\"name\":\"parent\",\"type\":\"bytes32\"},{\"name\":\"author\",\"type\":\"address\"},{\"name\":\"deleted\",\"type\":\"bool\"},{\"name\":\"publishDate\",\"type\":\"uint256\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_entryAuthor\",\"type\":\"address\"},{\"name\":\"_parent\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"}],\"name\":\"publish\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_commentId\",\"type\":\"bytes32\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_essence\",\"type\":\"address\"}],\"name\":\"setEssence\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\"}],\"name\":\"totalCommentsOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_commentId\",\"type\":\"bytes32\"}],\"name\":\"commentAuthor\",\"outputs\":[{\"name\":\"_author\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_entries\",\"type\":\"address\"}],\"name\":\"setEntries\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voting_period\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required_essence\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setRequiredEssence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"entryId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"parent\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"entryId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Comments is an auto generated Go binding around an Ethereum contract.
type Comments struct {
	CommentsCaller     // Read-only binding to the contract
	CommentsTransactor // Write-only binding to the contract
	CommentsFilterer   // Log filterer for contract events
}

// CommentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommentsSession struct {
	Contract     *Comments         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommentsCallerSession struct {
	Contract *CommentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CommentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommentsTransactorSession struct {
	Contract     *CommentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CommentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommentsRaw struct {
	Contract *Comments // Generic contract binding to access the raw methods on
}

// CommentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommentsCallerRaw struct {
	Contract *CommentsCaller // Generic read-only contract binding to access the raw methods on
}

// CommentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommentsTransactorRaw struct {
	Contract *CommentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComments creates a new instance of Comments, bound to a specific deployed contract.
func NewComments(address common.Address, backend bind.ContractBackend) (*Comments, error) {
	contract, err := bindComments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comments{CommentsCaller: CommentsCaller{contract: contract}, CommentsTransactor: CommentsTransactor{contract: contract}, CommentsFilterer: CommentsFilterer{contract: contract}}, nil
}

// NewCommentsCaller creates a new read-only instance of Comments, bound to a specific deployed contract.
func NewCommentsCaller(address common.Address, caller bind.ContractCaller) (*CommentsCaller, error) {
	contract, err := bindComments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommentsCaller{contract: contract}, nil
}

// NewCommentsTransactor creates a new write-only instance of Comments, bound to a specific deployed contract.
func NewCommentsTransactor(address common.Address, transactor bind.ContractTransactor) (*CommentsTransactor, error) {
	contract, err := bindComments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommentsTransactor{contract: contract}, nil
}

// NewCommentsFilterer creates a new log filterer instance of Comments, bound to a specific deployed contract.
func NewCommentsFilterer(address common.Address, filterer bind.ContractFilterer) (*CommentsFilterer, error) {
	contract, err := bindComments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommentsFilterer{contract: contract}, nil
}

// bindComments binds a generic wrapper to an already deployed contract.
func bindComments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommentsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comments *CommentsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Comments.Contract.CommentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comments *CommentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comments.Contract.CommentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comments *CommentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comments.Contract.CommentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comments *CommentsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Comments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comments *CommentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comments *CommentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comments.Contract.contract.Transact(opts, method, params...)
}

// CommentAuthor is a free data retrieval call binding the contract method 0xbbd34ac0.
//
// Solidity: function commentAuthor(_entryId bytes32, _commentId bytes32) constant returns(_author address)
func (_Comments *CommentsCaller) CommentAuthor(opts *bind.CallOpts, _entryId [32]byte, _commentId [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "commentAuthor", _entryId, _commentId)
	return *ret0, err
}

// CommentAuthor is a free data retrieval call binding the contract method 0xbbd34ac0.
//
// Solidity: function commentAuthor(_entryId bytes32, _commentId bytes32) constant returns(_author address)
func (_Comments *CommentsSession) CommentAuthor(_entryId [32]byte, _commentId [32]byte) (common.Address, error) {
	return _Comments.Contract.CommentAuthor(&_Comments.CallOpts, _entryId, _commentId)
}

// CommentAuthor is a free data retrieval call binding the contract method 0xbbd34ac0.
//
// Solidity: function commentAuthor(_entryId bytes32, _commentId bytes32) constant returns(_author address)
func (_Comments *CommentsCallerSession) CommentAuthor(_entryId [32]byte, _commentId [32]byte) (common.Address, error) {
	return _Comments.Contract.CommentAuthor(&_Comments.CallOpts, _entryId, _commentId)
}

// Discount_every is a free data retrieval call binding the contract method 0x152c85aa.
//
// Solidity: function discount_every() constant returns(uint256)
func (_Comments *CommentsCaller) Discount_every(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "discount_every")
	return *ret0, err
}

// Discount_every is a free data retrieval call binding the contract method 0x152c85aa.
//
// Solidity: function discount_every() constant returns(uint256)
func (_Comments *CommentsSession) Discount_every() (*big.Int, error) {
	return _Comments.Contract.Discount_every(&_Comments.CallOpts)
}

// Discount_every is a free data retrieval call binding the contract method 0x152c85aa.
//
// Solidity: function discount_every() constant returns(uint256)
func (_Comments *CommentsCallerSession) Discount_every() (*big.Int, error) {
	return _Comments.Contract.Discount_every(&_Comments.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x8683612b.
//
// Solidity: function exists(_entryId bytes32, _commentId bytes32) constant returns(bool)
func (_Comments *CommentsCaller) Exists(opts *bind.CallOpts, _entryId [32]byte, _commentId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "exists", _entryId, _commentId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x8683612b.
//
// Solidity: function exists(_entryId bytes32, _commentId bytes32) constant returns(bool)
func (_Comments *CommentsSession) Exists(_entryId [32]byte, _commentId [32]byte) (bool, error) {
	return _Comments.Contract.Exists(&_Comments.CallOpts, _entryId, _commentId)
}

// Exists is a free data retrieval call binding the contract method 0x8683612b.
//
// Solidity: function exists(_entryId bytes32, _commentId bytes32) constant returns(bool)
func (_Comments *CommentsCallerSession) Exists(_entryId [32]byte, _commentId [32]byte) (bool, error) {
	return _Comments.Contract.Exists(&_Comments.CallOpts, _entryId, _commentId)
}

// GetComment is a free data retrieval call binding the contract method 0x6920ef70.
//
// Solidity: function getComment(_entryId bytes32, _commentId bytes32) constant returns(parent bytes32, author address, deleted bool, publishDate uint256, _fn uint8, _digestSize uint8, _hash bytes32)
func (_Comments *CommentsCaller) GetComment(opts *bind.CallOpts, _entryId [32]byte, _commentId [32]byte) (struct {
	Parent      [32]byte
	Author      common.Address
	Deleted     bool
	PublishDate *big.Int
	Fn          uint8
	DigestSize  uint8
	Hash        [32]byte
}, error) {
	ret := new(struct {
		Parent      [32]byte
		Author      common.Address
		Deleted     bool
		PublishDate *big.Int
		Fn          uint8
		DigestSize  uint8
		Hash        [32]byte
	})
	out := ret
	err := _Comments.contract.Call(opts, out, "getComment", _entryId, _commentId)
	return *ret, err
}

// GetComment is a free data retrieval call binding the contract method 0x6920ef70.
//
// Solidity: function getComment(_entryId bytes32, _commentId bytes32) constant returns(parent bytes32, author address, deleted bool, publishDate uint256, _fn uint8, _digestSize uint8, _hash bytes32)
func (_Comments *CommentsSession) GetComment(_entryId [32]byte, _commentId [32]byte) (struct {
	Parent      [32]byte
	Author      common.Address
	Deleted     bool
	PublishDate *big.Int
	Fn          uint8
	DigestSize  uint8
	Hash        [32]byte
}, error) {
	return _Comments.Contract.GetComment(&_Comments.CallOpts, _entryId, _commentId)
}

// GetComment is a free data retrieval call binding the contract method 0x6920ef70.
//
// Solidity: function getComment(_entryId bytes32, _commentId bytes32) constant returns(parent bytes32, author address, deleted bool, publishDate uint256, _fn uint8, _digestSize uint8, _hash bytes32)
func (_Comments *CommentsCallerSession) GetComment(_entryId [32]byte, _commentId [32]byte) (struct {
	Parent      [32]byte
	Author      common.Address
	Deleted     bool
	PublishDate *big.Int
	Fn          uint8
	DigestSize  uint8
	Hash        [32]byte
}, error) {
	return _Comments.Contract.GetComment(&_Comments.CallOpts, _entryId, _commentId)
}

// IsDeleted is a free data retrieval call binding the contract method 0x1026418a.
//
// Solidity: function isDeleted(_entryId bytes32, _commentId bytes32) constant returns(_deleted bool)
func (_Comments *CommentsCaller) IsDeleted(opts *bind.CallOpts, _entryId [32]byte, _commentId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "isDeleted", _entryId, _commentId)
	return *ret0, err
}

// IsDeleted is a free data retrieval call binding the contract method 0x1026418a.
//
// Solidity: function isDeleted(_entryId bytes32, _commentId bytes32) constant returns(_deleted bool)
func (_Comments *CommentsSession) IsDeleted(_entryId [32]byte, _commentId [32]byte) (bool, error) {
	return _Comments.Contract.IsDeleted(&_Comments.CallOpts, _entryId, _commentId)
}

// IsDeleted is a free data retrieval call binding the contract method 0x1026418a.
//
// Solidity: function isDeleted(_entryId bytes32, _commentId bytes32) constant returns(_deleted bool)
func (_Comments *CommentsCallerSession) IsDeleted(_entryId [32]byte, _commentId [32]byte) (bool, error) {
	return _Comments.Contract.IsDeleted(&_Comments.CallOpts, _entryId, _commentId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Comments *CommentsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Comments *CommentsSession) Owner() (common.Address, error) {
	return _Comments.Contract.Owner(&_Comments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Comments *CommentsCallerSession) Owner() (common.Address, error) {
	return _Comments.Contract.Owner(&_Comments.CallOpts)
}

// Required_essence is a free data retrieval call binding the contract method 0xea754469.
//
// Solidity: function required_essence() constant returns(uint256)
func (_Comments *CommentsCaller) Required_essence(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "required_essence")
	return *ret0, err
}

// Required_essence is a free data retrieval call binding the contract method 0xea754469.
//
// Solidity: function required_essence() constant returns(uint256)
func (_Comments *CommentsSession) Required_essence() (*big.Int, error) {
	return _Comments.Contract.Required_essence(&_Comments.CallOpts)
}

// Required_essence is a free data retrieval call binding the contract method 0xea754469.
//
// Solidity: function required_essence() constant returns(uint256)
func (_Comments *CommentsCallerSession) Required_essence() (*big.Int, error) {
	return _Comments.Contract.Required_essence(&_Comments.CallOpts)
}

// TotalComments is a free data retrieval call binding the contract method 0x09227663.
//
// Solidity: function totalComments(_entryId bytes32) constant returns(_total uint256)
func (_Comments *CommentsCaller) TotalComments(opts *bind.CallOpts, _entryId [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "totalComments", _entryId)
	return *ret0, err
}

// TotalComments is a free data retrieval call binding the contract method 0x09227663.
//
// Solidity: function totalComments(_entryId bytes32) constant returns(_total uint256)
func (_Comments *CommentsSession) TotalComments(_entryId [32]byte) (*big.Int, error) {
	return _Comments.Contract.TotalComments(&_Comments.CallOpts, _entryId)
}

// TotalComments is a free data retrieval call binding the contract method 0x09227663.
//
// Solidity: function totalComments(_entryId bytes32) constant returns(_total uint256)
func (_Comments *CommentsCallerSession) TotalComments(_entryId [32]byte) (*big.Int, error) {
	return _Comments.Contract.TotalComments(&_Comments.CallOpts, _entryId)
}

// TotalCommentsOf is a free data retrieval call binding the contract method 0xa3180d33.
//
// Solidity: function totalCommentsOf(_publisher address) constant returns(uint256)
func (_Comments *CommentsCaller) TotalCommentsOf(opts *bind.CallOpts, _publisher common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "totalCommentsOf", _publisher)
	return *ret0, err
}

// TotalCommentsOf is a free data retrieval call binding the contract method 0xa3180d33.
//
// Solidity: function totalCommentsOf(_publisher address) constant returns(uint256)
func (_Comments *CommentsSession) TotalCommentsOf(_publisher common.Address) (*big.Int, error) {
	return _Comments.Contract.TotalCommentsOf(&_Comments.CallOpts, _publisher)
}

// TotalCommentsOf is a free data retrieval call binding the contract method 0xa3180d33.
//
// Solidity: function totalCommentsOf(_publisher address) constant returns(uint256)
func (_Comments *CommentsCallerSession) TotalCommentsOf(_publisher common.Address) (*big.Int, error) {
	return _Comments.Contract.TotalCommentsOf(&_Comments.CallOpts, _publisher)
}

// Voting_period is a free data retrieval call binding the contract method 0xe100bfe6.
//
// Solidity: function voting_period() constant returns(uint256)
func (_Comments *CommentsCaller) Voting_period(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Comments.contract.Call(opts, out, "voting_period")
	return *ret0, err
}

// Voting_period is a free data retrieval call binding the contract method 0xe100bfe6.
//
// Solidity: function voting_period() constant returns(uint256)
func (_Comments *CommentsSession) Voting_period() (*big.Int, error) {
	return _Comments.Contract.Voting_period(&_Comments.CallOpts)
}

// Voting_period is a free data retrieval call binding the contract method 0xe100bfe6.
//
// Solidity: function voting_period() constant returns(uint256)
func (_Comments *CommentsCallerSession) Voting_period() (*big.Int, error) {
	return _Comments.Contract.Voting_period(&_Comments.CallOpts)
}

// DeleteComment is a paid mutator transaction binding the contract method 0x3c56ec36.
//
// Solidity: function deleteComment(_entryId bytes32, _entryAuthor address, _commentId bytes32) returns(bool)
func (_Comments *CommentsTransactor) DeleteComment(opts *bind.TransactOpts, _entryId [32]byte, _entryAuthor common.Address, _commentId [32]byte) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "deleteComment", _entryId, _entryAuthor, _commentId)
}

// DeleteComment is a paid mutator transaction binding the contract method 0x3c56ec36.
//
// Solidity: function deleteComment(_entryId bytes32, _entryAuthor address, _commentId bytes32) returns(bool)
func (_Comments *CommentsSession) DeleteComment(_entryId [32]byte, _entryAuthor common.Address, _commentId [32]byte) (*types.Transaction, error) {
	return _Comments.Contract.DeleteComment(&_Comments.TransactOpts, _entryId, _entryAuthor, _commentId)
}

// DeleteComment is a paid mutator transaction binding the contract method 0x3c56ec36.
//
// Solidity: function deleteComment(_entryId bytes32, _entryAuthor address, _commentId bytes32) returns(bool)
func (_Comments *CommentsTransactorSession) DeleteComment(_entryId [32]byte, _entryAuthor common.Address, _commentId [32]byte) (*types.Transaction, error) {
	return _Comments.Contract.DeleteComment(&_Comments.TransactOpts, _entryId, _entryAuthor, _commentId)
}

// Edit is a paid mutator transaction binding the contract method 0x054bcf9a.
//
// Solidity: function edit(_entryId bytes32, _entryAuthor address, _commentId bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_Comments *CommentsTransactor) Edit(opts *bind.TransactOpts, _entryId [32]byte, _entryAuthor common.Address, _commentId [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "edit", _entryId, _entryAuthor, _commentId, _hash, _fn, _digestSize)
}

// Edit is a paid mutator transaction binding the contract method 0x054bcf9a.
//
// Solidity: function edit(_entryId bytes32, _entryAuthor address, _commentId bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_Comments *CommentsSession) Edit(_entryId [32]byte, _entryAuthor common.Address, _commentId [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Comments.Contract.Edit(&_Comments.TransactOpts, _entryId, _entryAuthor, _commentId, _hash, _fn, _digestSize)
}

// Edit is a paid mutator transaction binding the contract method 0x054bcf9a.
//
// Solidity: function edit(_entryId bytes32, _entryAuthor address, _commentId bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_Comments *CommentsTransactorSession) Edit(_entryId [32]byte, _entryAuthor common.Address, _commentId [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Comments.Contract.Edit(&_Comments.TransactOpts, _entryId, _entryAuthor, _commentId, _hash, _fn, _digestSize)
}

// Publish is a paid mutator transaction binding the contract method 0x8613456b.
//
// Solidity: function publish(_entryId bytes32, _entryAuthor address, _parent bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_Comments *CommentsTransactor) Publish(opts *bind.TransactOpts, _entryId [32]byte, _entryAuthor common.Address, _parent [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "publish", _entryId, _entryAuthor, _parent, _hash, _fn, _digestSize)
}

// Publish is a paid mutator transaction binding the contract method 0x8613456b.
//
// Solidity: function publish(_entryId bytes32, _entryAuthor address, _parent bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_Comments *CommentsSession) Publish(_entryId [32]byte, _entryAuthor common.Address, _parent [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Comments.Contract.Publish(&_Comments.TransactOpts, _entryId, _entryAuthor, _parent, _hash, _fn, _digestSize)
}

// Publish is a paid mutator transaction binding the contract method 0x8613456b.
//
// Solidity: function publish(_entryId bytes32, _entryAuthor address, _parent bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns(bool)
func (_Comments *CommentsTransactorSession) Publish(_entryId [32]byte, _entryAuthor common.Address, _parent [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Comments.Contract.Publish(&_Comments.TransactOpts, _entryId, _entryAuthor, _parent, _hash, _fn, _digestSize)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Comments *CommentsTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Comments *CommentsSession) ReclaimEther() (*types.Transaction, error) {
	return _Comments.Contract.ReclaimEther(&_Comments.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Comments *CommentsTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _Comments.Contract.ReclaimEther(&_Comments.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Comments *CommentsTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Comments *CommentsSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Comments.Contract.ReclaimToken(&_Comments.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Comments *CommentsTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Comments.Contract.ReclaimToken(&_Comments.TransactOpts, token)
}

// SetEntries is a paid mutator transaction binding the contract method 0xc42a7b52.
//
// Solidity: function setEntries(_entries address) returns()
func (_Comments *CommentsTransactor) SetEntries(opts *bind.TransactOpts, _entries common.Address) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "setEntries", _entries)
}

// SetEntries is a paid mutator transaction binding the contract method 0xc42a7b52.
//
// Solidity: function setEntries(_entries address) returns()
func (_Comments *CommentsSession) SetEntries(_entries common.Address) (*types.Transaction, error) {
	return _Comments.Contract.SetEntries(&_Comments.TransactOpts, _entries)
}

// SetEntries is a paid mutator transaction binding the contract method 0xc42a7b52.
//
// Solidity: function setEntries(_entries address) returns()
func (_Comments *CommentsTransactorSession) SetEntries(_entries common.Address) (*types.Transaction, error) {
	return _Comments.Contract.SetEntries(&_Comments.TransactOpts, _entries)
}

// SetEssence is a paid mutator transaction binding the contract method 0x8f0ecae2.
//
// Solidity: function setEssence(_essence address) returns()
func (_Comments *CommentsTransactor) SetEssence(opts *bind.TransactOpts, _essence common.Address) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "setEssence", _essence)
}

// SetEssence is a paid mutator transaction binding the contract method 0x8f0ecae2.
//
// Solidity: function setEssence(_essence address) returns()
func (_Comments *CommentsSession) SetEssence(_essence common.Address) (*types.Transaction, error) {
	return _Comments.Contract.SetEssence(&_Comments.TransactOpts, _essence)
}

// SetEssence is a paid mutator transaction binding the contract method 0x8f0ecae2.
//
// Solidity: function setEssence(_essence address) returns()
func (_Comments *CommentsTransactorSession) SetEssence(_essence common.Address) (*types.Transaction, error) {
	return _Comments.Contract.SetEssence(&_Comments.TransactOpts, _essence)
}

// SetRequiredEssence is a paid mutator transaction binding the contract method 0xfe4bc0f0.
//
// Solidity: function setRequiredEssence(_amount uint256) returns(bool)
func (_Comments *CommentsTransactor) SetRequiredEssence(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "setRequiredEssence", _amount)
}

// SetRequiredEssence is a paid mutator transaction binding the contract method 0xfe4bc0f0.
//
// Solidity: function setRequiredEssence(_amount uint256) returns(bool)
func (_Comments *CommentsSession) SetRequiredEssence(_amount *big.Int) (*types.Transaction, error) {
	return _Comments.Contract.SetRequiredEssence(&_Comments.TransactOpts, _amount)
}

// SetRequiredEssence is a paid mutator transaction binding the contract method 0xfe4bc0f0.
//
// Solidity: function setRequiredEssence(_amount uint256) returns(bool)
func (_Comments *CommentsTransactorSession) SetRequiredEssence(_amount *big.Int) (*types.Transaction, error) {
	return _Comments.Contract.SetRequiredEssence(&_Comments.TransactOpts, _amount)
}

// SetVotes is a paid mutator transaction binding the contract method 0x39004b27.
//
// Solidity: function setVotes(_votes address) returns()
func (_Comments *CommentsTransactor) SetVotes(opts *bind.TransactOpts, _votes common.Address) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "setVotes", _votes)
}

// SetVotes is a paid mutator transaction binding the contract method 0x39004b27.
//
// Solidity: function setVotes(_votes address) returns()
func (_Comments *CommentsSession) SetVotes(_votes common.Address) (*types.Transaction, error) {
	return _Comments.Contract.SetVotes(&_Comments.TransactOpts, _votes)
}

// SetVotes is a paid mutator transaction binding the contract method 0x39004b27.
//
// Solidity: function setVotes(_votes address) returns()
func (_Comments *CommentsTransactorSession) SetVotes(_votes common.Address) (*types.Transaction, error) {
	return _Comments.Contract.SetVotes(&_Comments.TransactOpts, _votes)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(_period uint256) returns(bool)
func (_Comments *CommentsTransactor) SetVotingPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "setVotingPeriod", _period)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(_period uint256) returns(bool)
func (_Comments *CommentsSession) SetVotingPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Comments.Contract.SetVotingPeriod(&_Comments.TransactOpts, _period)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(_period uint256) returns(bool)
func (_Comments *CommentsTransactorSession) SetVotingPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Comments.Contract.SetVotingPeriod(&_Comments.TransactOpts, _period)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Comments *CommentsTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Comments *CommentsSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Comments.Contract.TokenFallback(&_Comments.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Comments *CommentsTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Comments.Contract.TokenFallback(&_Comments.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Comments *CommentsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Comments.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Comments *CommentsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Comments.Contract.TransferOwnership(&_Comments.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Comments *CommentsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Comments.Contract.TransferOwnership(&_Comments.TransactOpts, newOwner)
}

// CommentsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Comments contract.
type CommentsOwnershipTransferredIterator struct {
	Event *CommentsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CommentsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommentsOwnershipTransferred)
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
		it.Event = new(CommentsOwnershipTransferred)
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
func (it *CommentsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommentsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommentsOwnershipTransferred represents a OwnershipTransferred event raised by the Comments contract.
type CommentsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Comments *CommentsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CommentsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Comments.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CommentsOwnershipTransferredIterator{contract: _Comments.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Comments *CommentsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CommentsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Comments.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommentsOwnershipTransferred)
				if err := _Comments.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// CommentsPublishIterator is returned from FilterPublish and is used to iterate over the raw logs and unpacked data for Publish events raised by the Comments contract.
type CommentsPublishIterator struct {
	Event *CommentsPublish // Event containing the contract specifics and raw log

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
func (it *CommentsPublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommentsPublish)
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
		it.Event = new(CommentsPublish)
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
func (it *CommentsPublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommentsPublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommentsPublish represents a Publish event raised by the Comments contract.
type CommentsPublish struct {
	Author  common.Address
	EntryId [32]byte
	Parent  [32]byte
	Id      [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0x7f6b1cd311d41d73f8ed4b87ba42952b5dcdc185b1e84b2f03dfc48d4e9f42b0.
//
// Solidity: event Publish(author indexed address, entryId indexed bytes32, parent indexed bytes32, id bytes32)
func (_Comments *CommentsFilterer) FilterPublish(opts *bind.FilterOpts, author []common.Address, entryId [][32]byte, parent [][32]byte) (*CommentsPublishIterator, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _Comments.contract.FilterLogs(opts, "Publish", authorRule, entryIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return &CommentsPublishIterator{contract: _Comments.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0x7f6b1cd311d41d73f8ed4b87ba42952b5dcdc185b1e84b2f03dfc48d4e9f42b0.
//
// Solidity: event Publish(author indexed address, entryId indexed bytes32, parent indexed bytes32, id bytes32)
func (_Comments *CommentsFilterer) WatchPublish(opts *bind.WatchOpts, sink chan<- *CommentsPublish, author []common.Address, entryId [][32]byte, parent [][32]byte) (event.Subscription, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _Comments.contract.WatchLogs(opts, "Publish", authorRule, entryIdRule, parentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommentsPublish)
				if err := _Comments.contract.UnpackLog(event, "Publish", log); err != nil {
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

// CommentsUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the Comments contract.
type CommentsUpdateIterator struct {
	Event *CommentsUpdate // Event containing the contract specifics and raw log

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
func (it *CommentsUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CommentsUpdate)
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
		it.Event = new(CommentsUpdate)
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
func (it *CommentsUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CommentsUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CommentsUpdate represents a Update event raised by the Comments contract.
type CommentsUpdate struct {
	Author  common.Address
	EntryId [32]byte
	Id      [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0xe7e154f822858c98d161949b394a7608624b34731b51bb82c3f3535d443d7228.
//
// Solidity: event Update(author indexed address, entryId indexed bytes32, id indexed bytes32)
func (_Comments *CommentsFilterer) FilterUpdate(opts *bind.FilterOpts, author []common.Address, entryId [][32]byte, id [][32]byte) (*CommentsUpdateIterator, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Comments.contract.FilterLogs(opts, "Update", authorRule, entryIdRule, idRule)
	if err != nil {
		return nil, err
	}
	return &CommentsUpdateIterator{contract: _Comments.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0xe7e154f822858c98d161949b394a7608624b34731b51bb82c3f3535d443d7228.
//
// Solidity: event Update(author indexed address, entryId indexed bytes32, id indexed bytes32)
func (_Comments *CommentsFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *CommentsUpdate, author []common.Address, entryId [][32]byte, id [][32]byte) (event.Subscription, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Comments.contract.WatchLogs(opts, "Update", authorRule, entryIdRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CommentsUpdate)
				if err := _Comments.contract.UnpackLog(event, "Update", log); err != nil {
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
