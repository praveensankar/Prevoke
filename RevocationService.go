// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"indexes\",\"type\":\"uint256[3]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"indexes\",\"type\":\"uint256[3]\"}],\"name\":\"revokeInBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testRevocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506107bc8061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80630de54b85146100645780631d1438481461006e57806348db53361461008c57806389cdfa6e146100bc578063b0d2b15e146100d8578063c3eed8e7146100e2575b5f80fd5b61006c610112565b005b610076610114565b60405161008391906104c5565b60405180910390f35b6100a660048036038101906100a1919061051e565b610139565b6040516100b39190610563565b60405180910390f35b6100d660048036038101906100d191906106ba565b610155565b005b6100e061020e565b005b6100fc60048036038101906100f791906106ba565b610411565b6040516101099190610563565b60405180910390f35b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f602052805f5260405f205f915054906101000a900460ff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101ad575f80fd5b5f5b600381101561020a5760015f808484600381106101cf576101ce6106e5565b5b602002015181526020019081526020015f205f6101000a81548160ff02191690831515021790555080806102029061073f565b9150506101af565b5050565b5f60405180606001604052806001815260200160028152602001600381525090505f60405180606001604052806004815260200160058152602001600681525090505f604051806060016040528060078152602001600881526020016009815250905061027a83610155565b61028382610155565b5f61028d84610411565b6102cc576040518060400160405280600b81526020017f6e6f74207265766f6b6564000000000000000000000000000000000000000000815250610303565b6040518060400160405280600781526020017f7265766f6b6564000000000000000000000000000000000000000000000000008152505b90505f61030f84610411565b61034e576040518060400160405280600b81526020017f6e6f74207265766f6b6564000000000000000000000000000000000000000000815250610385565b6040518060400160405280600781526020017f7265766f6b6564000000000000000000000000000000000000000000000000008152505b90505f61039184610411565b6103d0576040518060400160405280600b81526020017f6e6f74207265766f6b6564000000000000000000000000000000000000000000815250610407565b6040518060400160405280600781526020017f7265766f6b6564000000000000000000000000000000000000000000000000008152505b9050505050505050565b5f805f90505f5b600381101561047c57600115155f8086846003811061043a576104396106e5565b5b602002015181526020019081526020015f205f9054906101000a900460ff16151503610469576001915061047c565b80806104749061073f565b915050610418565b5080915050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6104af82610486565b9050919050565b6104bf816104a5565b82525050565b5f6020820190506104d85f8301846104b6565b92915050565b5f604051905090565b5f80fd5b5f819050919050565b6104fd816104eb565b8114610507575f80fd5b50565b5f81359050610518816104f4565b92915050565b5f60208284031215610533576105326104e7565b5b5f6105408482850161050a565b91505092915050565b5f8115159050919050565b61055d81610549565b82525050565b5f6020820190506105765f830184610554565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6105c682610580565b810181811067ffffffffffffffff821117156105e5576105e4610590565b5b80604052505050565b5f6105f76104de565b905061060382826105bd565b919050565b5f67ffffffffffffffff82111561062257610621610590565b5b602082029050919050565b5f80fd5b5f61064361063e84610608565b6105ee565b9050806020840283018581111561065d5761065c61062d565b5b835b818110156106865780610672888261050a565b84526020840193505060208101905061065f565b5050509392505050565b5f82601f8301126106a4576106a361057c565b5b60036106b1848285610631565b91505092915050565b5f606082840312156106cf576106ce6104e7565b5b5f6106dc84828501610690565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610749826104eb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361077b5761077a610712565b5b60018201905091905056fea2646970667358221220a852493a644af0660046ad5da5c5223e6b1fd79d661dad719efa64154ff0d47864736f6c63430008150033",
}

// RevocationServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use RevocationServiceMetaData.ABI instead.
var RevocationServiceABI = RevocationServiceMetaData.ABI

// RevocationServiceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RevocationServiceMetaData.Bin instead.
var RevocationServiceBin = RevocationServiceMetaData.Bin

// DeployRevocationService deploys a new Ethereum contract, binding an instance of RevocationService to it.
func DeployRevocationService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RevocationService, error) {
	parsed, err := RevocationServiceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RevocationServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RevocationService{RevocationServiceCaller: RevocationServiceCaller{contract: contract}, RevocationServiceTransactor: RevocationServiceTransactor{contract: contract}, RevocationServiceFilterer: RevocationServiceFilterer{contract: contract}}, nil
}

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
// Solidity: function checkRevocationStatusInBloomFilter(uint256[3] indexes) view returns(bool)
func (_RevocationService *RevocationServiceCaller) CheckRevocationStatusInBloomFilter(opts *bind.CallOpts, indexes [3]*big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "checkRevocationStatusInBloomFilter", indexes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xc3eed8e7.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[3] indexes) view returns(bool)
func (_RevocationService *RevocationServiceSession) CheckRevocationStatusInBloomFilter(indexes [3]*big.Int) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInBloomFilter(&_RevocationService.CallOpts, indexes)
}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xc3eed8e7.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[3] indexes) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) CheckRevocationStatusInBloomFilter(indexes [3]*big.Int) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInBloomFilter(&_RevocationService.CallOpts, indexes)
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

// RevokeInBloomFilter is a paid mutator transaction binding the contract method 0x89cdfa6e.
//
// Solidity: function revokeInBloomFilter(uint256[3] indexes) returns()
func (_RevocationService *RevocationServiceTransactor) RevokeInBloomFilter(opts *bind.TransactOpts, indexes [3]*big.Int) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "revokeInBloomFilter", indexes)
}

// RevokeInBloomFilter is a paid mutator transaction binding the contract method 0x89cdfa6e.
//
// Solidity: function revokeInBloomFilter(uint256[3] indexes) returns()
func (_RevocationService *RevocationServiceSession) RevokeInBloomFilter(indexes [3]*big.Int) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeInBloomFilter(&_RevocationService.TransactOpts, indexes)
}

// RevokeInBloomFilter is a paid mutator transaction binding the contract method 0x89cdfa6e.
//
// Solidity: function revokeInBloomFilter(uint256[3] indexes) returns()
func (_RevocationService *RevocationServiceTransactorSession) RevokeInBloomFilter(indexes [3]*big.Int) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeInBloomFilter(&_RevocationService.TransactOpts, indexes)
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
