// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relay

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BSCRelayContractABI is the input ABI used to generate the binding from.
const BSCRelayContractABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validatorSet\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_currentHeight\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getCurrentHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"isBlockFinal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"unsignedHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signedHeader\",\"type\":\"bytes\"}],\"name\":\"submitBlockHeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"unsignedHeader\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signedHeader\",\"type\":\"bytes[]\"}],\"name\":\"submitBlockHeaderBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeInWei\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signedHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedValue\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeInWei\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signedHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedValue\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyTransaction\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// BSCRelayContract is an auto generated Go binding around an Ethereum contract.
type BSCRelayContract struct {
	BSCRelayContractCaller     // Read-only binding to the contract
	BSCRelayContractTransactor // Write-only binding to the contract
	BSCRelayContractFilterer   // Log filterer for contract events
}

// BSCRelayContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BSCRelayContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCRelayContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCRelayContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCRelayContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCRelayContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCRelayContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCRelayContractSession struct {
	Contract     *BSCRelayContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCRelayContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCRelayContractCallerSession struct {
	Contract *BSCRelayContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BSCRelayContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCRelayContractTransactorSession struct {
	Contract     *BSCRelayContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BSCRelayContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BSCRelayContractRaw struct {
	Contract *BSCRelayContract // Generic contract binding to access the raw methods on
}

// BSCRelayContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCRelayContractCallerRaw struct {
	Contract *BSCRelayContractCaller // Generic read-only contract binding to access the raw methods on
}

// BSCRelayContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCRelayContractTransactorRaw struct {
	Contract *BSCRelayContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCRelayContract creates a new instance of BSCRelayContract, bound to a specific deployed contract.
func NewBSCRelayContract(address common.Address, backend bind.ContractBackend) (*BSCRelayContract, error) {
	contract, err := bindBSCRelayContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCRelayContract{BSCRelayContractCaller: BSCRelayContractCaller{contract: contract}, BSCRelayContractTransactor: BSCRelayContractTransactor{contract: contract}, BSCRelayContractFilterer: BSCRelayContractFilterer{contract: contract}}, nil
}

// NewBSCRelayContractCaller creates a new read-only instance of BSCRelayContract, bound to a specific deployed contract.
func NewBSCRelayContractCaller(address common.Address, caller bind.ContractCaller) (*BSCRelayContractCaller, error) {
	contract, err := bindBSCRelayContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCRelayContractCaller{contract: contract}, nil
}

// NewBSCRelayContractTransactor creates a new write-only instance of BSCRelayContract, bound to a specific deployed contract.
func NewBSCRelayContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BSCRelayContractTransactor, error) {
	contract, err := bindBSCRelayContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCRelayContractTransactor{contract: contract}, nil
}

// NewBSCRelayContractFilterer creates a new log filterer instance of BSCRelayContract, bound to a specific deployed contract.
func NewBSCRelayContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BSCRelayContractFilterer, error) {
	contract, err := bindBSCRelayContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCRelayContractFilterer{contract: contract}, nil
}

// bindBSCRelayContract binds a generic wrapper to an already deployed contract.
func bindBSCRelayContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCRelayContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCRelayContract *BSCRelayContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCRelayContract.Contract.BSCRelayContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCRelayContract *BSCRelayContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.BSCRelayContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCRelayContract *BSCRelayContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.BSCRelayContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCRelayContract *BSCRelayContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCRelayContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCRelayContract *BSCRelayContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCRelayContract *BSCRelayContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.contract.Transact(opts, method, params...)
}

// GetCurrentHeight is a free data retrieval call binding the contract method 0x3caf44dc.
//
// Solidity: function getCurrentHeight() view returns(uint256)
func (_BSCRelayContract *BSCRelayContractCaller) GetCurrentHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCRelayContract.contract.Call(opts, &out, "getCurrentHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentHeight is a free data retrieval call binding the contract method 0x3caf44dc.
//
// Solidity: function getCurrentHeight() view returns(uint256)
func (_BSCRelayContract *BSCRelayContractSession) GetCurrentHeight() (*big.Int, error) {
	return _BSCRelayContract.Contract.GetCurrentHeight(&_BSCRelayContract.CallOpts)
}

// GetCurrentHeight is a free data retrieval call binding the contract method 0x3caf44dc.
//
// Solidity: function getCurrentHeight() view returns(uint256)
func (_BSCRelayContract *BSCRelayContractCallerSession) GetCurrentHeight() (*big.Int, error) {
	return _BSCRelayContract.Contract.GetCurrentHeight(&_BSCRelayContract.CallOpts)
}

// IsBlockFinal is a free data retrieval call binding the contract method 0x4a6aeb50.
//
// Solidity: function isBlockFinal(bytes32 blockHash) view returns(bool)
func (_BSCRelayContract *BSCRelayContractCaller) IsBlockFinal(opts *bind.CallOpts, blockHash [32]byte) (bool, error) {
	var out []interface{}
	err := _BSCRelayContract.contract.Call(opts, &out, "isBlockFinal", blockHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockFinal is a free data retrieval call binding the contract method 0x4a6aeb50.
//
// Solidity: function isBlockFinal(bytes32 blockHash) view returns(bool)
func (_BSCRelayContract *BSCRelayContractSession) IsBlockFinal(blockHash [32]byte) (bool, error) {
	return _BSCRelayContract.Contract.IsBlockFinal(&_BSCRelayContract.CallOpts, blockHash)
}

// IsBlockFinal is a free data retrieval call binding the contract method 0x4a6aeb50.
//
// Solidity: function isBlockFinal(bytes32 blockHash) view returns(bool)
func (_BSCRelayContract *BSCRelayContractCallerSession) IsBlockFinal(blockHash [32]byte) (bool, error) {
	return _BSCRelayContract.Contract.IsBlockFinal(&_BSCRelayContract.CallOpts, blockHash)
}

// SubmitBlockHeader is a paid mutator transaction binding the contract method 0x738c38a9.
//
// Solidity: function submitBlockHeader(bytes unsignedHeader, bytes signedHeader) returns()
func (_BSCRelayContract *BSCRelayContractTransactor) SubmitBlockHeader(opts *bind.TransactOpts, unsignedHeader []byte, signedHeader []byte) (*types.Transaction, error) {
	return _BSCRelayContract.contract.Transact(opts, "submitBlockHeader", unsignedHeader, signedHeader)
}

// SubmitBlockHeader is a paid mutator transaction binding the contract method 0x738c38a9.
//
// Solidity: function submitBlockHeader(bytes unsignedHeader, bytes signedHeader) returns()
func (_BSCRelayContract *BSCRelayContractSession) SubmitBlockHeader(unsignedHeader []byte, signedHeader []byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.SubmitBlockHeader(&_BSCRelayContract.TransactOpts, unsignedHeader, signedHeader)
}

// SubmitBlockHeader is a paid mutator transaction binding the contract method 0x738c38a9.
//
// Solidity: function submitBlockHeader(bytes unsignedHeader, bytes signedHeader) returns()
func (_BSCRelayContract *BSCRelayContractTransactorSession) SubmitBlockHeader(unsignedHeader []byte, signedHeader []byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.SubmitBlockHeader(&_BSCRelayContract.TransactOpts, unsignedHeader, signedHeader)
}

// SubmitBlockHeaderBatch is a paid mutator transaction binding the contract method 0x54407e47.
//
// Solidity: function submitBlockHeaderBatch(bytes[] unsignedHeader, bytes[] signedHeader) returns()
func (_BSCRelayContract *BSCRelayContractTransactor) SubmitBlockHeaderBatch(opts *bind.TransactOpts, unsignedHeader [][]byte, signedHeader [][]byte) (*types.Transaction, error) {
	return _BSCRelayContract.contract.Transact(opts, "submitBlockHeaderBatch", unsignedHeader, signedHeader)
}

// SubmitBlockHeaderBatch is a paid mutator transaction binding the contract method 0x54407e47.
//
// Solidity: function submitBlockHeaderBatch(bytes[] unsignedHeader, bytes[] signedHeader) returns()
func (_BSCRelayContract *BSCRelayContractSession) SubmitBlockHeaderBatch(unsignedHeader [][]byte, signedHeader [][]byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.SubmitBlockHeaderBatch(&_BSCRelayContract.TransactOpts, unsignedHeader, signedHeader)
}

// SubmitBlockHeaderBatch is a paid mutator transaction binding the contract method 0x54407e47.
//
// Solidity: function submitBlockHeaderBatch(bytes[] unsignedHeader, bytes[] signedHeader) returns()
func (_BSCRelayContract *BSCRelayContractTransactorSession) SubmitBlockHeaderBatch(unsignedHeader [][]byte, signedHeader [][]byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.SubmitBlockHeaderBatch(&_BSCRelayContract.TransactOpts, unsignedHeader, signedHeader)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xed315dfa.
//
// Solidity: function verifyReceipt(uint256 feeInWei, bytes signedHeader, uint8 noOfConfirmations, bytes rlpEncodedValue, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_BSCRelayContract *BSCRelayContractTransactor) VerifyReceipt(opts *bind.TransactOpts, feeInWei *big.Int, signedHeader []byte, noOfConfirmations uint8, rlpEncodedValue []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _BSCRelayContract.contract.Transact(opts, "verifyReceipt", feeInWei, signedHeader, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xed315dfa.
//
// Solidity: function verifyReceipt(uint256 feeInWei, bytes signedHeader, uint8 noOfConfirmations, bytes rlpEncodedValue, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_BSCRelayContract *BSCRelayContractSession) VerifyReceipt(feeInWei *big.Int, signedHeader []byte, noOfConfirmations uint8, rlpEncodedValue []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.VerifyReceipt(&_BSCRelayContract.TransactOpts, feeInWei, signedHeader, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xed315dfa.
//
// Solidity: function verifyReceipt(uint256 feeInWei, bytes signedHeader, uint8 noOfConfirmations, bytes rlpEncodedValue, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_BSCRelayContract *BSCRelayContractTransactorSession) VerifyReceipt(feeInWei *big.Int, signedHeader []byte, noOfConfirmations uint8, rlpEncodedValue []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.VerifyReceipt(&_BSCRelayContract.TransactOpts, feeInWei, signedHeader, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0x5e29b7da.
//
// Solidity: function verifyTransaction(uint256 feeInWei, bytes signedHeader, uint8 noOfConfirmations, bytes rlpEncodedValue, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_BSCRelayContract *BSCRelayContractTransactor) VerifyTransaction(opts *bind.TransactOpts, feeInWei *big.Int, signedHeader []byte, noOfConfirmations uint8, rlpEncodedValue []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _BSCRelayContract.contract.Transact(opts, "verifyTransaction", feeInWei, signedHeader, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0x5e29b7da.
//
// Solidity: function verifyTransaction(uint256 feeInWei, bytes signedHeader, uint8 noOfConfirmations, bytes rlpEncodedValue, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_BSCRelayContract *BSCRelayContractSession) VerifyTransaction(feeInWei *big.Int, signedHeader []byte, noOfConfirmations uint8, rlpEncodedValue []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.VerifyTransaction(&_BSCRelayContract.TransactOpts, feeInWei, signedHeader, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes)
}

// VerifyTransaction is a paid mutator transaction binding the contract method 0x5e29b7da.
//
// Solidity: function verifyTransaction(uint256 feeInWei, bytes signedHeader, uint8 noOfConfirmations, bytes rlpEncodedValue, bytes path, bytes rlpEncodedNodes) payable returns(uint8)
func (_BSCRelayContract *BSCRelayContractTransactorSession) VerifyTransaction(feeInWei *big.Int, signedHeader []byte, noOfConfirmations uint8, rlpEncodedValue []byte, path []byte, rlpEncodedNodes []byte) (*types.Transaction, error) {
	return _BSCRelayContract.Contract.VerifyTransaction(&_BSCRelayContract.TransactOpts, feeInWei, signedHeader, noOfConfirmations, rlpEncodedValue, path, rlpEncodedNodes)
}
