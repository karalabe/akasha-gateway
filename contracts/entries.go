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

// EntriesABI is the input ABI used to generate the binding from.
const EntriesABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_essence\",\"type\":\"address\"}],\"name\":\"setEssenceAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"discount_every\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setDiscountEvery\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\"},{\"name\":\"_entryId\",\"type\":\"bytes32\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"}],\"name\":\"edit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_tags\",\"type\":\"bytes32[]\"}],\"name\":\"publishArticle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_votes\",\"type\":\"address\"}],\"name\":\"setVotesAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_tags\",\"type\":\"bytes32[]\"}],\"name\":\"publishOther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tags\",\"type\":\"address\"}],\"name\":\"setTagsAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_entryId\",\"type\":\"bytes32\"}],\"name\":\"claim\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\"}],\"name\":\"getEntryCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_publisher\",\"type\":\"address\"},{\"name\":\"_entryId\",\"type\":\"bytes32\"}],\"name\":\"getEntry\",\"outputs\":[{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_tags\",\"type\":\"bytes32[]\"}],\"name\":\"publishLink\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_fn\",\"type\":\"uint8\"},{\"name\":\"_digestSize\",\"type\":\"uint8\"},{\"name\":\"_tags\",\"type\":\"bytes32[]\"}],\"name\":\"publishMedia\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voting_period\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setVotingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required_essence\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setRequiredEssence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_essence\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"entryId\",\"type\":\"bytes32\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"tagName\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"entryId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"entryType\",\"type\":\"uint8\"}],\"name\":\"TagIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"author\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"entryId\",\"type\":\"bytes32\"}],\"name\":\"Update\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Entries is an auto generated Go binding around an Ethereum contract.
type Entries struct {
	EntriesCaller     // Read-only binding to the contract
	EntriesTransactor // Write-only binding to the contract
	EntriesFilterer   // Log filterer for contract events
}

// EntriesCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntriesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntriesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntriesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntriesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntriesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntriesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntriesSession struct {
	Contract     *Entries          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntriesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntriesCallerSession struct {
	Contract *EntriesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EntriesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntriesTransactorSession struct {
	Contract     *EntriesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EntriesRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntriesRaw struct {
	Contract *Entries // Generic contract binding to access the raw methods on
}

// EntriesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntriesCallerRaw struct {
	Contract *EntriesCaller // Generic read-only contract binding to access the raw methods on
}

// EntriesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntriesTransactorRaw struct {
	Contract *EntriesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntries creates a new instance of Entries, bound to a specific deployed contract.
func NewEntries(address common.Address, backend bind.ContractBackend) (*Entries, error) {
	contract, err := bindEntries(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Entries{EntriesCaller: EntriesCaller{contract: contract}, EntriesTransactor: EntriesTransactor{contract: contract}, EntriesFilterer: EntriesFilterer{contract: contract}}, nil
}

// NewEntriesCaller creates a new read-only instance of Entries, bound to a specific deployed contract.
func NewEntriesCaller(address common.Address, caller bind.ContractCaller) (*EntriesCaller, error) {
	contract, err := bindEntries(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntriesCaller{contract: contract}, nil
}

// NewEntriesTransactor creates a new write-only instance of Entries, bound to a specific deployed contract.
func NewEntriesTransactor(address common.Address, transactor bind.ContractTransactor) (*EntriesTransactor, error) {
	contract, err := bindEntries(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntriesTransactor{contract: contract}, nil
}

// NewEntriesFilterer creates a new log filterer instance of Entries, bound to a specific deployed contract.
func NewEntriesFilterer(address common.Address, filterer bind.ContractFilterer) (*EntriesFilterer, error) {
	contract, err := bindEntries(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntriesFilterer{contract: contract}, nil
}

// bindEntries binds a generic wrapper to an already deployed contract.
func bindEntries(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EntriesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entries *EntriesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entries.Contract.EntriesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entries *EntriesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entries.Contract.EntriesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entries *EntriesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entries.Contract.EntriesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entries *EntriesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entries.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entries *EntriesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entries.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entries *EntriesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entries.Contract.contract.Transact(opts, method, params...)
}

// Discount_every is a free data retrieval call binding the contract method 0x152c85aa.
//
// Solidity: function discount_every() constant returns(uint256)
func (_Entries *EntriesCaller) Discount_every(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "discount_every")
	return *ret0, err
}

// Discount_every is a free data retrieval call binding the contract method 0x152c85aa.
//
// Solidity: function discount_every() constant returns(uint256)
func (_Entries *EntriesSession) Discount_every() (*big.Int, error) {
	return _Entries.Contract.Discount_every(&_Entries.CallOpts)
}

// Discount_every is a free data retrieval call binding the contract method 0x152c85aa.
//
// Solidity: function discount_every() constant returns(uint256)
func (_Entries *EntriesCallerSession) Discount_every() (*big.Int, error) {
	return _Entries.Contract.Discount_every(&_Entries.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x45120511.
//
// Solidity: function exists(_publisher address, _entryId bytes32) constant returns(bool)
func (_Entries *EntriesCaller) Exists(opts *bind.CallOpts, _publisher common.Address, _entryId [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "exists", _publisher, _entryId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x45120511.
//
// Solidity: function exists(_publisher address, _entryId bytes32) constant returns(bool)
func (_Entries *EntriesSession) Exists(_publisher common.Address, _entryId [32]byte) (bool, error) {
	return _Entries.Contract.Exists(&_Entries.CallOpts, _publisher, _entryId)
}

// Exists is a free data retrieval call binding the contract method 0x45120511.
//
// Solidity: function exists(_publisher address, _entryId bytes32) constant returns(bool)
func (_Entries *EntriesCallerSession) Exists(_publisher common.Address, _entryId [32]byte) (bool, error) {
	return _Entries.Contract.Exists(&_Entries.CallOpts, _publisher, _entryId)
}

// GetEntry is a free data retrieval call binding the contract method 0xcad2ab23.
//
// Solidity: function getEntry(_publisher address, _entryId bytes32) constant returns(_fn uint8, _digestSize uint8, _hash bytes32)
func (_Entries *EntriesCaller) GetEntry(opts *bind.CallOpts, _publisher common.Address, _entryId [32]byte) (struct {
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
	err := _Entries.contract.Call(opts, out, "getEntry", _publisher, _entryId)
	return *ret, err
}

// GetEntry is a free data retrieval call binding the contract method 0xcad2ab23.
//
// Solidity: function getEntry(_publisher address, _entryId bytes32) constant returns(_fn uint8, _digestSize uint8, _hash bytes32)
func (_Entries *EntriesSession) GetEntry(_publisher common.Address, _entryId [32]byte) (struct {
	Fn         uint8
	DigestSize uint8
	Hash       [32]byte
}, error) {
	return _Entries.Contract.GetEntry(&_Entries.CallOpts, _publisher, _entryId)
}

// GetEntry is a free data retrieval call binding the contract method 0xcad2ab23.
//
// Solidity: function getEntry(_publisher address, _entryId bytes32) constant returns(_fn uint8, _digestSize uint8, _hash bytes32)
func (_Entries *EntriesCallerSession) GetEntry(_publisher common.Address, _entryId [32]byte) (struct {
	Fn         uint8
	DigestSize uint8
	Hash       [32]byte
}, error) {
	return _Entries.Contract.GetEntry(&_Entries.CallOpts, _publisher, _entryId)
}

// GetEntryCount is a free data retrieval call binding the contract method 0xc5c9bd25.
//
// Solidity: function getEntryCount(_publisher address) constant returns(uint256)
func (_Entries *EntriesCaller) GetEntryCount(opts *bind.CallOpts, _publisher common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "getEntryCount", _publisher)
	return *ret0, err
}

// GetEntryCount is a free data retrieval call binding the contract method 0xc5c9bd25.
//
// Solidity: function getEntryCount(_publisher address) constant returns(uint256)
func (_Entries *EntriesSession) GetEntryCount(_publisher common.Address) (*big.Int, error) {
	return _Entries.Contract.GetEntryCount(&_Entries.CallOpts, _publisher)
}

// GetEntryCount is a free data retrieval call binding the contract method 0xc5c9bd25.
//
// Solidity: function getEntryCount(_publisher address) constant returns(uint256)
func (_Entries *EntriesCallerSession) GetEntryCount(_publisher common.Address) (*big.Int, error) {
	return _Entries.Contract.GetEntryCount(&_Entries.CallOpts, _publisher)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Entries *EntriesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Entries *EntriesSession) Owner() (common.Address, error) {
	return _Entries.Contract.Owner(&_Entries.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Entries *EntriesCallerSession) Owner() (common.Address, error) {
	return _Entries.Contract.Owner(&_Entries.CallOpts)
}

// Required_essence is a free data retrieval call binding the contract method 0xea754469.
//
// Solidity: function required_essence() constant returns(uint256)
func (_Entries *EntriesCaller) Required_essence(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "required_essence")
	return *ret0, err
}

// Required_essence is a free data retrieval call binding the contract method 0xea754469.
//
// Solidity: function required_essence() constant returns(uint256)
func (_Entries *EntriesSession) Required_essence() (*big.Int, error) {
	return _Entries.Contract.Required_essence(&_Entries.CallOpts)
}

// Required_essence is a free data retrieval call binding the contract method 0xea754469.
//
// Solidity: function required_essence() constant returns(uint256)
func (_Entries *EntriesCallerSession) Required_essence() (*big.Int, error) {
	return _Entries.Contract.Required_essence(&_Entries.CallOpts)
}

// Voting_period is a free data retrieval call binding the contract method 0xe100bfe6.
//
// Solidity: function voting_period() constant returns(uint256)
func (_Entries *EntriesCaller) Voting_period(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "voting_period")
	return *ret0, err
}

// Voting_period is a free data retrieval call binding the contract method 0xe100bfe6.
//
// Solidity: function voting_period() constant returns(uint256)
func (_Entries *EntriesSession) Voting_period() (*big.Int, error) {
	return _Entries.Contract.Voting_period(&_Entries.CallOpts)
}

// Voting_period is a free data retrieval call binding the contract method 0xe100bfe6.
//
// Solidity: function voting_period() constant returns(uint256)
func (_Entries *EntriesCallerSession) Voting_period() (*big.Int, error) {
	return _Entries.Contract.Voting_period(&_Entries.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0xbd66528a.
//
// Solidity: function claim(_entryId bytes32) returns(bool)
func (_Entries *EntriesTransactor) Claim(opts *bind.TransactOpts, _entryId [32]byte) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "claim", _entryId)
}

// Claim is a paid mutator transaction binding the contract method 0xbd66528a.
//
// Solidity: function claim(_entryId bytes32) returns(bool)
func (_Entries *EntriesSession) Claim(_entryId [32]byte) (*types.Transaction, error) {
	return _Entries.Contract.Claim(&_Entries.TransactOpts, _entryId)
}

// Claim is a paid mutator transaction binding the contract method 0xbd66528a.
//
// Solidity: function claim(_entryId bytes32) returns(bool)
func (_Entries *EntriesTransactorSession) Claim(_entryId [32]byte) (*types.Transaction, error) {
	return _Entries.Contract.Claim(&_Entries.TransactOpts, _entryId)
}

// Edit is a paid mutator transaction binding the contract method 0x758a592d.
//
// Solidity: function edit(_entryId bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns()
func (_Entries *EntriesTransactor) Edit(opts *bind.TransactOpts, _entryId [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "edit", _entryId, _hash, _fn, _digestSize)
}

// Edit is a paid mutator transaction binding the contract method 0x758a592d.
//
// Solidity: function edit(_entryId bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns()
func (_Entries *EntriesSession) Edit(_entryId [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Entries.Contract.Edit(&_Entries.TransactOpts, _entryId, _hash, _fn, _digestSize)
}

// Edit is a paid mutator transaction binding the contract method 0x758a592d.
//
// Solidity: function edit(_entryId bytes32, _hash bytes32, _fn uint8, _digestSize uint8) returns()
func (_Entries *EntriesTransactorSession) Edit(_entryId [32]byte, _hash [32]byte, _fn uint8, _digestSize uint8) (*types.Transaction, error) {
	return _Entries.Contract.Edit(&_Entries.TransactOpts, _entryId, _hash, _fn, _digestSize)
}

// PublishArticle is a paid mutator transaction binding the contract method 0x8ab62406.
//
// Solidity: function publishArticle(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactor) PublishArticle(opts *bind.TransactOpts, _hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "publishArticle", _hash, _fn, _digestSize, _tags)
}

// PublishArticle is a paid mutator transaction binding the contract method 0x8ab62406.
//
// Solidity: function publishArticle(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesSession) PublishArticle(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishArticle(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// PublishArticle is a paid mutator transaction binding the contract method 0x8ab62406.
//
// Solidity: function publishArticle(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactorSession) PublishArticle(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishArticle(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// PublishLink is a paid mutator transaction binding the contract method 0xd1b33ca4.
//
// Solidity: function publishLink(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactor) PublishLink(opts *bind.TransactOpts, _hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "publishLink", _hash, _fn, _digestSize, _tags)
}

// PublishLink is a paid mutator transaction binding the contract method 0xd1b33ca4.
//
// Solidity: function publishLink(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesSession) PublishLink(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishLink(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// PublishLink is a paid mutator transaction binding the contract method 0xd1b33ca4.
//
// Solidity: function publishLink(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactorSession) PublishLink(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishLink(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// PublishMedia is a paid mutator transaction binding the contract method 0xd2aa00f4.
//
// Solidity: function publishMedia(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactor) PublishMedia(opts *bind.TransactOpts, _hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "publishMedia", _hash, _fn, _digestSize, _tags)
}

// PublishMedia is a paid mutator transaction binding the contract method 0xd2aa00f4.
//
// Solidity: function publishMedia(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesSession) PublishMedia(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishMedia(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// PublishMedia is a paid mutator transaction binding the contract method 0xd2aa00f4.
//
// Solidity: function publishMedia(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactorSession) PublishMedia(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishMedia(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// PublishOther is a paid mutator transaction binding the contract method 0xa92ac252.
//
// Solidity: function publishOther(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactor) PublishOther(opts *bind.TransactOpts, _hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "publishOther", _hash, _fn, _digestSize, _tags)
}

// PublishOther is a paid mutator transaction binding the contract method 0xa92ac252.
//
// Solidity: function publishOther(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesSession) PublishOther(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishOther(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// PublishOther is a paid mutator transaction binding the contract method 0xa92ac252.
//
// Solidity: function publishOther(_hash bytes32, _fn uint8, _digestSize uint8, _tags bytes32[]) returns()
func (_Entries *EntriesTransactorSession) PublishOther(_hash [32]byte, _fn uint8, _digestSize uint8, _tags [][32]byte) (*types.Transaction, error) {
	return _Entries.Contract.PublishOther(&_Entries.TransactOpts, _hash, _fn, _digestSize, _tags)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Entries *EntriesTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Entries *EntriesSession) ReclaimEther() (*types.Transaction, error) {
	return _Entries.Contract.ReclaimEther(&_Entries.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_Entries *EntriesTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _Entries.Contract.ReclaimEther(&_Entries.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Entries *EntriesTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Entries *EntriesSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Entries.Contract.ReclaimToken(&_Entries.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_Entries *EntriesTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _Entries.Contract.ReclaimToken(&_Entries.TransactOpts, token)
}

// SetDiscountEvery is a paid mutator transaction binding the contract method 0x1eec8289.
//
// Solidity: function setDiscountEvery(_amount uint256) returns(bool)
func (_Entries *EntriesTransactor) SetDiscountEvery(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "setDiscountEvery", _amount)
}

// SetDiscountEvery is a paid mutator transaction binding the contract method 0x1eec8289.
//
// Solidity: function setDiscountEvery(_amount uint256) returns(bool)
func (_Entries *EntriesSession) SetDiscountEvery(_amount *big.Int) (*types.Transaction, error) {
	return _Entries.Contract.SetDiscountEvery(&_Entries.TransactOpts, _amount)
}

// SetDiscountEvery is a paid mutator transaction binding the contract method 0x1eec8289.
//
// Solidity: function setDiscountEvery(_amount uint256) returns(bool)
func (_Entries *EntriesTransactorSession) SetDiscountEvery(_amount *big.Int) (*types.Transaction, error) {
	return _Entries.Contract.SetDiscountEvery(&_Entries.TransactOpts, _amount)
}

// SetEssenceAddress is a paid mutator transaction binding the contract method 0x0455d89d.
//
// Solidity: function setEssenceAddress(_essence address) returns(bool)
func (_Entries *EntriesTransactor) SetEssenceAddress(opts *bind.TransactOpts, _essence common.Address) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "setEssenceAddress", _essence)
}

// SetEssenceAddress is a paid mutator transaction binding the contract method 0x0455d89d.
//
// Solidity: function setEssenceAddress(_essence address) returns(bool)
func (_Entries *EntriesSession) SetEssenceAddress(_essence common.Address) (*types.Transaction, error) {
	return _Entries.Contract.SetEssenceAddress(&_Entries.TransactOpts, _essence)
}

// SetEssenceAddress is a paid mutator transaction binding the contract method 0x0455d89d.
//
// Solidity: function setEssenceAddress(_essence address) returns(bool)
func (_Entries *EntriesTransactorSession) SetEssenceAddress(_essence common.Address) (*types.Transaction, error) {
	return _Entries.Contract.SetEssenceAddress(&_Entries.TransactOpts, _essence)
}

// SetRequiredEssence is a paid mutator transaction binding the contract method 0xfe4bc0f0.
//
// Solidity: function setRequiredEssence(_amount uint256) returns(bool)
func (_Entries *EntriesTransactor) SetRequiredEssence(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "setRequiredEssence", _amount)
}

// SetRequiredEssence is a paid mutator transaction binding the contract method 0xfe4bc0f0.
//
// Solidity: function setRequiredEssence(_amount uint256) returns(bool)
func (_Entries *EntriesSession) SetRequiredEssence(_amount *big.Int) (*types.Transaction, error) {
	return _Entries.Contract.SetRequiredEssence(&_Entries.TransactOpts, _amount)
}

// SetRequiredEssence is a paid mutator transaction binding the contract method 0xfe4bc0f0.
//
// Solidity: function setRequiredEssence(_amount uint256) returns(bool)
func (_Entries *EntriesTransactorSession) SetRequiredEssence(_amount *big.Int) (*types.Transaction, error) {
	return _Entries.Contract.SetRequiredEssence(&_Entries.TransactOpts, _amount)
}

// SetTagsAddress is a paid mutator transaction binding the contract method 0xab6e5eda.
//
// Solidity: function setTagsAddress(_tags address) returns(bool)
func (_Entries *EntriesTransactor) SetTagsAddress(opts *bind.TransactOpts, _tags common.Address) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "setTagsAddress", _tags)
}

// SetTagsAddress is a paid mutator transaction binding the contract method 0xab6e5eda.
//
// Solidity: function setTagsAddress(_tags address) returns(bool)
func (_Entries *EntriesSession) SetTagsAddress(_tags common.Address) (*types.Transaction, error) {
	return _Entries.Contract.SetTagsAddress(&_Entries.TransactOpts, _tags)
}

// SetTagsAddress is a paid mutator transaction binding the contract method 0xab6e5eda.
//
// Solidity: function setTagsAddress(_tags address) returns(bool)
func (_Entries *EntriesTransactorSession) SetTagsAddress(_tags common.Address) (*types.Transaction, error) {
	return _Entries.Contract.SetTagsAddress(&_Entries.TransactOpts, _tags)
}

// SetVotesAddress is a paid mutator transaction binding the contract method 0x8d4e5963.
//
// Solidity: function setVotesAddress(_votes address) returns(bool)
func (_Entries *EntriesTransactor) SetVotesAddress(opts *bind.TransactOpts, _votes common.Address) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "setVotesAddress", _votes)
}

// SetVotesAddress is a paid mutator transaction binding the contract method 0x8d4e5963.
//
// Solidity: function setVotesAddress(_votes address) returns(bool)
func (_Entries *EntriesSession) SetVotesAddress(_votes common.Address) (*types.Transaction, error) {
	return _Entries.Contract.SetVotesAddress(&_Entries.TransactOpts, _votes)
}

// SetVotesAddress is a paid mutator transaction binding the contract method 0x8d4e5963.
//
// Solidity: function setVotesAddress(_votes address) returns(bool)
func (_Entries *EntriesTransactorSession) SetVotesAddress(_votes common.Address) (*types.Transaction, error) {
	return _Entries.Contract.SetVotesAddress(&_Entries.TransactOpts, _votes)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(_period uint256) returns(bool)
func (_Entries *EntriesTransactor) SetVotingPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "setVotingPeriod", _period)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(_period uint256) returns(bool)
func (_Entries *EntriesSession) SetVotingPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Entries.Contract.SetVotingPeriod(&_Entries.TransactOpts, _period)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xea0217cf.
//
// Solidity: function setVotingPeriod(_period uint256) returns(bool)
func (_Entries *EntriesTransactorSession) SetVotingPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Entries.Contract.SetVotingPeriod(&_Entries.TransactOpts, _period)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Entries *EntriesTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Entries *EntriesSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Entries.Contract.TokenFallback(&_Entries.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_Entries *EntriesTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _Entries.Contract.TokenFallback(&_Entries.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Entries *EntriesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Entries.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Entries *EntriesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Entries.Contract.TransferOwnership(&_Entries.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Entries *EntriesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Entries.Contract.TransferOwnership(&_Entries.TransactOpts, newOwner)
}

// EntriesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Entries contract.
type EntriesOwnershipTransferredIterator struct {
	Event *EntriesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EntriesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntriesOwnershipTransferred)
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
		it.Event = new(EntriesOwnershipTransferred)
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
func (it *EntriesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntriesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntriesOwnershipTransferred represents a OwnershipTransferred event raised by the Entries contract.
type EntriesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Entries *EntriesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EntriesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Entries.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EntriesOwnershipTransferredIterator{contract: _Entries.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Entries *EntriesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EntriesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Entries.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntriesOwnershipTransferred)
				if err := _Entries.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// EntriesPublishIterator is returned from FilterPublish and is used to iterate over the raw logs and unpacked data for Publish events raised by the Entries contract.
type EntriesPublishIterator struct {
	Event *EntriesPublish // Event containing the contract specifics and raw log

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
func (it *EntriesPublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntriesPublish)
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
		it.Event = new(EntriesPublish)
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
func (it *EntriesPublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntriesPublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntriesPublish represents a Publish event raised by the Entries contract.
type EntriesPublish struct {
	Author  common.Address
	EntryId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0xac6d015ca4b9b742bdf7acc16fcae54c45052fa461961fa47e75ae6cb182e775.
//
// Solidity: event Publish(author indexed address, entryId indexed bytes32)
func (_Entries *EntriesFilterer) FilterPublish(opts *bind.FilterOpts, author []common.Address, entryId [][32]byte) (*EntriesPublishIterator, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}

	logs, sub, err := _Entries.contract.FilterLogs(opts, "Publish", authorRule, entryIdRule)
	if err != nil {
		return nil, err
	}
	return &EntriesPublishIterator{contract: _Entries.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0xac6d015ca4b9b742bdf7acc16fcae54c45052fa461961fa47e75ae6cb182e775.
//
// Solidity: event Publish(author indexed address, entryId indexed bytes32)
func (_Entries *EntriesFilterer) WatchPublish(opts *bind.WatchOpts, sink chan<- *EntriesPublish, author []common.Address, entryId [][32]byte) (event.Subscription, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}

	logs, sub, err := _Entries.contract.WatchLogs(opts, "Publish", authorRule, entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntriesPublish)
				if err := _Entries.contract.UnpackLog(event, "Publish", log); err != nil {
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

// EntriesTagIndexIterator is returned from FilterTagIndex and is used to iterate over the raw logs and unpacked data for TagIndex events raised by the Entries contract.
type EntriesTagIndexIterator struct {
	Event *EntriesTagIndex // Event containing the contract specifics and raw log

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
func (it *EntriesTagIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntriesTagIndex)
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
		it.Event = new(EntriesTagIndex)
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
func (it *EntriesTagIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntriesTagIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntriesTagIndex represents a TagIndex event raised by the Entries contract.
type EntriesTagIndex struct {
	TagName   [32]byte
	EntryId   [32]byte
	EntryType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTagIndex is a free log retrieval operation binding the contract event 0x7c61ee1955d4d884d26ad3d7b2e7bbccebf672517e391878c03a44a3c3b6526a.
//
// Solidity: event TagIndex(tagName indexed bytes32, entryId indexed bytes32, entryType indexed uint8)
func (_Entries *EntriesFilterer) FilterTagIndex(opts *bind.FilterOpts, tagName [][32]byte, entryId [][32]byte, entryType []uint8) (*EntriesTagIndexIterator, error) {

	var tagNameRule []interface{}
	for _, tagNameItem := range tagName {
		tagNameRule = append(tagNameRule, tagNameItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var entryTypeRule []interface{}
	for _, entryTypeItem := range entryType {
		entryTypeRule = append(entryTypeRule, entryTypeItem)
	}

	logs, sub, err := _Entries.contract.FilterLogs(opts, "TagIndex", tagNameRule, entryIdRule, entryTypeRule)
	if err != nil {
		return nil, err
	}
	return &EntriesTagIndexIterator{contract: _Entries.contract, event: "TagIndex", logs: logs, sub: sub}, nil
}

// WatchTagIndex is a free log subscription operation binding the contract event 0x7c61ee1955d4d884d26ad3d7b2e7bbccebf672517e391878c03a44a3c3b6526a.
//
// Solidity: event TagIndex(tagName indexed bytes32, entryId indexed bytes32, entryType indexed uint8)
func (_Entries *EntriesFilterer) WatchTagIndex(opts *bind.WatchOpts, sink chan<- *EntriesTagIndex, tagName [][32]byte, entryId [][32]byte, entryType []uint8) (event.Subscription, error) {

	var tagNameRule []interface{}
	for _, tagNameItem := range tagName {
		tagNameRule = append(tagNameRule, tagNameItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var entryTypeRule []interface{}
	for _, entryTypeItem := range entryType {
		entryTypeRule = append(entryTypeRule, entryTypeItem)
	}

	logs, sub, err := _Entries.contract.WatchLogs(opts, "TagIndex", tagNameRule, entryIdRule, entryTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntriesTagIndex)
				if err := _Entries.contract.UnpackLog(event, "TagIndex", log); err != nil {
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

// EntriesUpdateIterator is returned from FilterUpdate and is used to iterate over the raw logs and unpacked data for Update events raised by the Entries contract.
type EntriesUpdateIterator struct {
	Event *EntriesUpdate // Event containing the contract specifics and raw log

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
func (it *EntriesUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntriesUpdate)
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
		it.Event = new(EntriesUpdate)
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
func (it *EntriesUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntriesUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntriesUpdate represents a Update event raised by the Entries contract.
type EntriesUpdate struct {
	Author  common.Address
	EntryId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdate is a free log retrieval operation binding the contract event 0xc270eb3f985f34bb5584f38c55f673923cda3c6529ffad6f711e691078ead909.
//
// Solidity: event Update(author indexed address, entryId indexed bytes32)
func (_Entries *EntriesFilterer) FilterUpdate(opts *bind.FilterOpts, author []common.Address, entryId [][32]byte) (*EntriesUpdateIterator, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}

	logs, sub, err := _Entries.contract.FilterLogs(opts, "Update", authorRule, entryIdRule)
	if err != nil {
		return nil, err
	}
	return &EntriesUpdateIterator{contract: _Entries.contract, event: "Update", logs: logs, sub: sub}, nil
}

// WatchUpdate is a free log subscription operation binding the contract event 0xc270eb3f985f34bb5584f38c55f673923cda3c6529ffad6f711e691078ead909.
//
// Solidity: event Update(author indexed address, entryId indexed bytes32)
func (_Entries *EntriesFilterer) WatchUpdate(opts *bind.WatchOpts, sink chan<- *EntriesUpdate, author []common.Address, entryId [][32]byte) (event.Subscription, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}

	logs, sub, err := _Entries.contract.WatchLogs(opts, "Update", authorRule, entryIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntriesUpdate)
				if err := _Entries.contract.UnpackLog(event, "Update", log); err != nil {
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
