// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RevocationServiceMetaData contains all meta data concerning the RevocationService contract.
var RevocationServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_indexes\",\"type\":\"uint256[3]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkRevocationStatusInMerkleTreeAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfHashFunctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testRevocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_indexes\",\"type\":\"uint256[3]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_values\",\"type\":\"bytes32[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[3]\"},{\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyVC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RevocationServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use RevocationServiceMetaData.ABI instead.
var RevocationServiceABI = RevocationServiceMetaData.ABI

// RevocationService is an auto generated Go binding around an Ethereum contract.
type RevocationService struct {
	RevocationServiceCaller     // Read-only binding to the contract
	RevocationServiceTransactor // Write-only binding to the contract
	RevocationServiceFilterer   // Log filterer for contract events
}

// RevocationServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevocationServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevocationServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevocationServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevocationServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevocationServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevocationServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevocationServiceSession struct {
	Contract     *RevocationService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RevocationServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevocationServiceCallerSession struct {
	Contract *RevocationServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RevocationServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevocationServiceTransactorSession struct {
	Contract     *RevocationServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RevocationServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevocationServiceRaw struct {
	Contract *RevocationService // Generic contract binding to access the raw methods on
}

// RevocationServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevocationServiceCallerRaw struct {
	Contract *RevocationServiceCaller // Generic read-only contract binding to access the raw methods on
}

// RevocationServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevocationServiceTransactorRaw struct {
	Contract *RevocationServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevocationService creates a new instance of RevocationService, bound to a specific deployed contract.
func NewRevocationService(address common.Address, backend bind.ContractBackend) (*RevocationService, error) {
	contract, err := bindRevocationService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevocationService{RevocationServiceCaller: RevocationServiceCaller{contract: contract}, RevocationServiceTransactor: RevocationServiceTransactor{contract: contract}, RevocationServiceFilterer: RevocationServiceFilterer{contract: contract}}, nil
}

// NewRevocationServiceCaller creates a new read-only instance of RevocationService, bound to a specific deployed contract.
func NewRevocationServiceCaller(address common.Address, caller bind.ContractCaller) (*RevocationServiceCaller, error) {
	contract, err := bindRevocationService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevocationServiceCaller{contract: contract}, nil
}

// NewRevocationServiceTransactor creates a new write-only instance of RevocationService, bound to a specific deployed contract.
func NewRevocationServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*RevocationServiceTransactor, error) {
	contract, err := bindRevocationService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevocationServiceTransactor{contract: contract}, nil
}

// NewRevocationServiceFilterer creates a new log filterer instance of RevocationService, bound to a specific deployed contract.
func NewRevocationServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*RevocationServiceFilterer, error) {
	contract, err := bindRevocationService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevocationServiceFilterer{contract: contract}, nil
}

// bindRevocationService binds a generic wrapper to an already deployed contract.
func bindRevocationService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevocationServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevocationService *RevocationServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevocationService.Contract.RevocationServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevocationService *RevocationServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.Contract.RevocationServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevocationService *RevocationServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevocationService.Contract.RevocationServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevocationService *RevocationServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevocationService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevocationService *RevocationServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevocationService *RevocationServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevocationService.Contract.contract.Transact(opts, method, params...)
}

// BloomFilter is a free data retrieval call binding the contract method 0x48db5336.
//
// Solidity: function bloomFilter(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceCaller) BloomFilter(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "bloomFilter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BloomFilter is a free data retrieval call binding the contract method 0x48db5336.
//
// Solidity: function bloomFilter(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceSession) BloomFilter(arg0 *big.Int) (bool, error) {
	return _RevocationService.Contract.BloomFilter(&_RevocationService.CallOpts, arg0)
}

// BloomFilter is a free data retrieval call binding the contract method 0x48db5336.
//
// Solidity: function bloomFilter(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) BloomFilter(arg0 *big.Int) (bool, error) {
	return _RevocationService.Contract.BloomFilter(&_RevocationService.CallOpts, arg0)
}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xc3eed8e7.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[3] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceCaller) CheckRevocationStatusInBloomFilter(opts *bind.CallOpts, _indexes [3]*big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "checkRevocationStatusInBloomFilter", _indexes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xc3eed8e7.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[3] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceSession) CheckRevocationStatusInBloomFilter(_indexes [3]*big.Int) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInBloomFilter(&_RevocationService.CallOpts, _indexes)
}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xc3eed8e7.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[3] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) CheckRevocationStatusInBloomFilter(_indexes [3]*big.Int) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInBloomFilter(&_RevocationService.CallOpts, _indexes)
}

// CheckRevocationStatusInMerkleTreeAccumulator is a free data retrieval call binding the contract method 0x4993101b.
//
// Solidity: function checkRevocationStatusInMerkleTreeAccumulator(bytes32 leaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceCaller) CheckRevocationStatusInMerkleTreeAccumulator(opts *bind.CallOpts, leaf [32]byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "checkRevocationStatusInMerkleTreeAccumulator", leaf, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRevocationStatusInMerkleTreeAccumulator is a free data retrieval call binding the contract method 0x4993101b.
//
// Solidity: function checkRevocationStatusInMerkleTreeAccumulator(bytes32 leaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceSession) CheckRevocationStatusInMerkleTreeAccumulator(leaf [32]byte, proof [][32]byte) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInMerkleTreeAccumulator(&_RevocationService.CallOpts, leaf, proof)
}

// CheckRevocationStatusInMerkleTreeAccumulator is a free data retrieval call binding the contract method 0x4993101b.
//
// Solidity: function checkRevocationStatusInMerkleTreeAccumulator(bytes32 leaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) CheckRevocationStatusInMerkleTreeAccumulator(leaf [32]byte, proof [][32]byte) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInMerkleTreeAccumulator(&_RevocationService.CallOpts, leaf, proof)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_RevocationService *RevocationServiceCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "issuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_RevocationService *RevocationServiceSession) Issuer() (common.Address, error) {
	return _RevocationService.Contract.Issuer(&_RevocationService.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_RevocationService *RevocationServiceCallerSession) Issuer() (common.Address, error) {
	return _RevocationService.Contract.Issuer(&_RevocationService.CallOpts)
}

// MerkleTree is a free data retrieval call binding the contract method 0x0df0ff90.
//
// Solidity: function merkleTree(uint256 ) view returns(bytes32)
func (_RevocationService *RevocationServiceCaller) MerkleTree(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "merkleTree", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleTree is a free data retrieval call binding the contract method 0x0df0ff90.
//
// Solidity: function merkleTree(uint256 ) view returns(bytes32)
func (_RevocationService *RevocationServiceSession) MerkleTree(arg0 *big.Int) ([32]byte, error) {
	return _RevocationService.Contract.MerkleTree(&_RevocationService.CallOpts, arg0)
}

// MerkleTree is a free data retrieval call binding the contract method 0x0df0ff90.
//
// Solidity: function merkleTree(uint256 ) view returns(bytes32)
func (_RevocationService *RevocationServiceCallerSession) MerkleTree(arg0 *big.Int) ([32]byte, error) {
	return _RevocationService.Contract.MerkleTree(&_RevocationService.CallOpts, arg0)
}

// NumberOfHashFunctions is a free data retrieval call binding the contract method 0x376a6590.
//
// Solidity: function numberOfHashFunctions() view returns(uint256)
func (_RevocationService *RevocationServiceCaller) NumberOfHashFunctions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "numberOfHashFunctions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfHashFunctions is a free data retrieval call binding the contract method 0x376a6590.
//
// Solidity: function numberOfHashFunctions() view returns(uint256)
func (_RevocationService *RevocationServiceSession) NumberOfHashFunctions() (*big.Int, error) {
	return _RevocationService.Contract.NumberOfHashFunctions(&_RevocationService.CallOpts)
}

// NumberOfHashFunctions is a free data retrieval call binding the contract method 0x376a6590.
//
// Solidity: function numberOfHashFunctions() view returns(uint256)
func (_RevocationService *RevocationServiceCallerSession) NumberOfHashFunctions() (*big.Int, error) {
	return _RevocationService.Contract.NumberOfHashFunctions(&_RevocationService.CallOpts)
}

// PrintMerkleTree is a free data retrieval call binding the contract method 0xd3e53042.
//
// Solidity: function printMerkleTree() view returns()
func (_RevocationService *RevocationServiceCaller) PrintMerkleTree(opts *bind.CallOpts) error {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "printMerkleTree")

	if err != nil {
		return err
	}

	return err

}

// PrintMerkleTree is a free data retrieval call binding the contract method 0xd3e53042.
//
// Solidity: function printMerkleTree() view returns()
func (_RevocationService *RevocationServiceSession) PrintMerkleTree() error {
	return _RevocationService.Contract.PrintMerkleTree(&_RevocationService.CallOpts)
}

// PrintMerkleTree is a free data retrieval call binding the contract method 0xd3e53042.
//
// Solidity: function printMerkleTree() view returns()
func (_RevocationService *RevocationServiceCallerSession) PrintMerkleTree() error {
	return _RevocationService.Contract.PrintMerkleTree(&_RevocationService.CallOpts)
}

// IssueVC is a paid mutator transaction binding the contract method 0xce4b3f34.
//
// Solidity: function issueVC(uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactor) IssueVC(opts *bind.TransactOpts, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "issueVC", _mtIndexes, _mtValues)
}

// IssueVC is a paid mutator transaction binding the contract method 0xce4b3f34.
//
// Solidity: function issueVC(uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceSession) IssueVC(_mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.IssueVC(&_RevocationService.TransactOpts, _mtIndexes, _mtValues)
}

// IssueVC is a paid mutator transaction binding the contract method 0xce4b3f34.
//
// Solidity: function issueVC(uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactorSession) IssueVC(_mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.IssueVC(&_RevocationService.TransactOpts, _mtIndexes, _mtValues)
}

// RegisterIssuers is a paid mutator transaction binding the contract method 0x0de54b85.
//
// Solidity: function registerIssuers() returns()
func (_RevocationService *RevocationServiceTransactor) RegisterIssuers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "registerIssuers")
}

// RegisterIssuers is a paid mutator transaction binding the contract method 0x0de54b85.
//
// Solidity: function registerIssuers() returns()
func (_RevocationService *RevocationServiceSession) RegisterIssuers() (*types.Transaction, error) {
	return _RevocationService.Contract.RegisterIssuers(&_RevocationService.TransactOpts)
}

// RegisterIssuers is a paid mutator transaction binding the contract method 0x0de54b85.
//
// Solidity: function registerIssuers() returns()
func (_RevocationService *RevocationServiceTransactorSession) RegisterIssuers() (*types.Transaction, error) {
	return _RevocationService.Contract.RegisterIssuers(&_RevocationService.TransactOpts)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x5c3702fb.
//
// Solidity: function revokeVC(uint256[3] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactor) RevokeVC(opts *bind.TransactOpts, _bfIndexes [3]*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "revokeVC", _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x5c3702fb.
//
// Solidity: function revokeVC(uint256[3] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceSession) RevokeVC(_bfIndexes [3]*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x5c3702fb.
//
// Solidity: function revokeVC(uint256[3] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactorSession) RevokeVC(_bfIndexes [3]*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// TestRevocation is a paid mutator transaction binding the contract method 0xb0d2b15e.
//
// Solidity: function testRevocation() returns()
func (_RevocationService *RevocationServiceTransactor) TestRevocation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "testRevocation")
}

// TestRevocation is a paid mutator transaction binding the contract method 0xb0d2b15e.
//
// Solidity: function testRevocation() returns()
func (_RevocationService *RevocationServiceSession) TestRevocation() (*types.Transaction, error) {
	return _RevocationService.Contract.TestRevocation(&_RevocationService.TransactOpts)
}

// TestRevocation is a paid mutator transaction binding the contract method 0xb0d2b15e.
//
// Solidity: function testRevocation() returns()
func (_RevocationService *RevocationServiceTransactorSession) TestRevocation() (*types.Transaction, error) {
	return _RevocationService.Contract.TestRevocation(&_RevocationService.TransactOpts)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0x33379d9a.
//
// Solidity: function updateBloomFilter(uint256[3] _indexes) returns()
func (_RevocationService *RevocationServiceTransactor) UpdateBloomFilter(opts *bind.TransactOpts, _indexes [3]*big.Int) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "updateBloomFilter", _indexes)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0x33379d9a.
//
// Solidity: function updateBloomFilter(uint256[3] _indexes) returns()
func (_RevocationService *RevocationServiceSession) UpdateBloomFilter(_indexes [3]*big.Int) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateBloomFilter(&_RevocationService.TransactOpts, _indexes)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0x33379d9a.
//
// Solidity: function updateBloomFilter(uint256[3] _indexes) returns()
func (_RevocationService *RevocationServiceTransactorSession) UpdateBloomFilter(_indexes [3]*big.Int) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateBloomFilter(&_RevocationService.TransactOpts, _indexes)
}

// UpdateMerkleTree is a paid mutator transaction binding the contract method 0x51600698.
//
// Solidity: function updateMerkleTree(uint256[] _indexes, bytes32[] _values) returns()
func (_RevocationService *RevocationServiceTransactor) UpdateMerkleTree(opts *bind.TransactOpts, _indexes []*big.Int, _values [][32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "updateMerkleTree", _indexes, _values)
}

// UpdateMerkleTree is a paid mutator transaction binding the contract method 0x51600698.
//
// Solidity: function updateMerkleTree(uint256[] _indexes, bytes32[] _values) returns()
func (_RevocationService *RevocationServiceSession) UpdateMerkleTree(_indexes []*big.Int, _values [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateMerkleTree(&_RevocationService.TransactOpts, _indexes, _values)
}

// UpdateMerkleTree is a paid mutator transaction binding the contract method 0x51600698.
//
// Solidity: function updateMerkleTree(uint256[] _indexes, bytes32[] _values) returns()
func (_RevocationService *RevocationServiceTransactorSession) UpdateMerkleTree(_indexes []*big.Int, _values [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateMerkleTree(&_RevocationService.TransactOpts, _indexes, _values)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x19f26717.
//
// Solidity: function updateNode(uint256 index, bytes32 value) returns()
func (_RevocationService *RevocationServiceTransactor) UpdateNode(opts *bind.TransactOpts, index *big.Int, value [32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "updateNode", index, value)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x19f26717.
//
// Solidity: function updateNode(uint256 index, bytes32 value) returns()
func (_RevocationService *RevocationServiceSession) UpdateNode(index *big.Int, value [32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateNode(&_RevocationService.TransactOpts, index, value)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x19f26717.
//
// Solidity: function updateNode(uint256 index, bytes32 value) returns()
func (_RevocationService *RevocationServiceTransactorSession) UpdateNode(index *big.Int, value [32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateNode(&_RevocationService.TransactOpts, index, value)
}

// VerifyVC is a paid mutator transaction binding the contract method 0x4b72f59a.
//
// Solidity: function verifyVC(uint256[3] _bfIndexes, bytes32 vcLeaf, bytes32[] proof) returns(bool)
func (_RevocationService *RevocationServiceTransactor) VerifyVC(opts *bind.TransactOpts, _bfIndexes [3]*big.Int, vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "verifyVC", _bfIndexes, vcLeaf, proof)
}

// VerifyVC is a paid mutator transaction binding the contract method 0x4b72f59a.
//
// Solidity: function verifyVC(uint256[3] _bfIndexes, bytes32 vcLeaf, bytes32[] proof) returns(bool)
func (_RevocationService *RevocationServiceSession) VerifyVC(_bfIndexes [3]*big.Int, vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.VerifyVC(&_RevocationService.TransactOpts, _bfIndexes, vcLeaf, proof)
}

// VerifyVC is a paid mutator transaction binding the contract method 0x4b72f59a.
//
// Solidity: function verifyVC(uint256[3] _bfIndexes, bytes32 vcLeaf, bytes32[] proof) returns(bool)
func (_RevocationService *RevocationServiceTransactorSession) VerifyVC(_bfIndexes [3]*big.Int, vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.VerifyVC(&_RevocationService.TransactOpts, _bfIndexes, vcLeaf, proof)
}
