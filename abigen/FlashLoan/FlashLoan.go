// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flashLoan

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

// FlashLoanABI is the input ABI used to generate the binding from.
const FlashLoanABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LENDING_POOL\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"premiums\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FlashLoanBin is the compiled bytecode used for deploying new contracts.
var FlashLoanBin = "0x60c060405260008060146101000a81548160ff0219169083151502179055503480156200002b57600080fd5b506040516200160a3803806200160a833981810160405281019062000051919062000236565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff16630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b158015620000d057600080fd5b505afa158015620000e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010b919062000236565b73ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b8152505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36001600060146101000a81548160ff02191690831515021790555050620002b0565b600081519050620002308162000296565b92915050565b6000602082840312156200024957600080fd5b600062000259848285016200021f565b91505092915050565b60006200026f8262000276565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620002a18162000262565b8114620002ad57600080fd5b50565b60805160601c60a05160601c611320620002ea60003960008181610469015281816107ac015261084d0152600061019001526113206000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063919840ad11610066578063919840ad146100fc578063920f5c841461011a5780639ad5981e1461014a578063b4dcfc7714610154578063f2fde38b1461017257610093565b80630542975c14610098578063715018a6146100b65780638da5cb5b146100c05780638f32d59b146100de575b600080fd5b6100a061018e565b6040516100ad9190610f7f565b60405180910390f35b6100be6101b2565b005b6100c86102b7565b6040516100d59190610e95565b60405180910390f35b6100e66102e0565b6040516100f39190610f64565b60405180910390f35b610104610337565b6040516101119190610f64565b60405180910390f35b610134600480360381019061012f9190610b7e565b61034a565b6040516101419190610f64565b60405180910390f35b61015261050b565b005b61015c61084b565b6040516101699190610f9a565b60405180910390f35b61018c60048036038101906101879190610b55565b61086f565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b6101ba6102e0565b6101f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f090610ff5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600060149054906101000a900460ff1681565b600080600060146101000a81548160ff02191690831515021790555060006103fb878760008181106103a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200201358a8a60008181106103e6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200201356108c290919063ffffffff16565b90508a8a6000818110610437577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b905060200201602081019061044c9190610b55565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000836040518363ffffffff1660e01b81526004016104a6929190610f3b565b602060405180830381600087803b1580156104c057600080fd5b505af11580156104d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f89190610c66565b5060019150509998505050505050505050565b60003090506000662386f26fc1000090506000600167ffffffffffffffff81111561055f577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561058d5781602001602082028036833780820191505090505b50905073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2816000815181106105df577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506000600167ffffffffffffffff81111561065c577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561068a5781602001602082028036833780820191505090505b50905082816000815181106106c8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181815250506000600167ffffffffffffffff811115610717577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156107455781602001602082028036833780820191505090505b509050600081600081518110610784577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101818152505060006040518060200160405280600081525090506000803090507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663ab9c4b5d898888888689896040518863ffffffff1660e01b815260040161080f9796959493929190610eb0565b600060405180830381600087803b15801561082957600080fd5b505af115801561083d573d6000803e3d6000fd5b505050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6108776102e0565b6108b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ad90610ff5565b60405180910390fd5b6108bf81610920565b50565b60008082846108d191906110b4565b905083811015610916576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090d90610fd5565b60405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610990576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098790610fb5565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600081359050610a5c816112bc565b92915050565b60008083601f840112610a7457600080fd5b8235905067ffffffffffffffff811115610a8d57600080fd5b602083019150836020820283011115610aa557600080fd5b9250929050565b60008083601f840112610abe57600080fd5b8235905067ffffffffffffffff811115610ad757600080fd5b602083019150836020820283011115610aef57600080fd5b9250929050565b600081519050610b05816112d3565b92915050565b60008083601f840112610b1d57600080fd5b8235905067ffffffffffffffff811115610b3657600080fd5b602083019150836001820283011115610b4e57600080fd5b9250929050565b600060208284031215610b6757600080fd5b6000610b7584828501610a4d565b91505092915050565b600080600080600080600080600060a08a8c031215610b9c57600080fd5b60008a013567ffffffffffffffff811115610bb657600080fd5b610bc28c828d01610a62565b995099505060208a013567ffffffffffffffff811115610be157600080fd5b610bed8c828d01610aac565b975097505060408a013567ffffffffffffffff811115610c0c57600080fd5b610c188c828d01610aac565b95509550506060610c2b8c828d01610a4d565b93505060808a013567ffffffffffffffff811115610c4857600080fd5b610c548c828d01610b0b565b92509250509295985092959850929598565b600060208284031215610c7857600080fd5b6000610c8684828501610af6565b91505092915050565b6000610c9b8383610cbf565b60208301905092915050565b6000610cb38383610e77565b60208301905092915050565b610cc88161110a565b82525050565b610cd78161110a565b82525050565b6000610ce882611035565b610cf28185611070565b9350610cfd83611015565b8060005b83811015610d2e578151610d158882610c8f565b9750610d2083611056565b925050600181019050610d01565b5085935050505092915050565b6000610d4682611040565b610d508185611081565b9350610d5b83611025565b8060005b83811015610d8c578151610d738882610ca7565b9750610d7e83611063565b925050600181019050610d5f565b5085935050505092915050565b610da28161111c565b82525050565b6000610db38261104b565b610dbd8185611092565b9350610dcd8185602086016111a8565b610dd68161120a565b840191505092915050565b610dea81611160565b82525050565b610df981611184565b82525050565b6000610e0c6026836110a3565b9150610e178261121b565b604082019050919050565b6000610e2f601b836110a3565b9150610e3a8261126a565b602082019050919050565b6000610e526020836110a3565b9150610e5d82611293565b602082019050919050565b610e7181611128565b82525050565b610e8081611156565b82525050565b610e8f81611156565b82525050565b6000602082019050610eaa6000830184610cce565b92915050565b600060e082019050610ec5600083018a610cce565b8181036020830152610ed78189610cdd565b90508181036040830152610eeb8188610d3b565b90508181036060830152610eff8187610d3b565b9050610f0e6080830186610cce565b81810360a0830152610f208185610da8565b9050610f2f60c0830184610e68565b98975050505050505050565b6000604082019050610f506000830185610cce565b610f5d6020830184610e86565b9392505050565b6000602082019050610f796000830184610d99565b92915050565b6000602082019050610f946000830184610de1565b92915050565b6000602082019050610faf6000830184610df0565b92915050565b60006020820190508181036000830152610fce81610dff565b9050919050565b60006020820190508181036000830152610fee81610e22565b9050919050565b6000602082019050818103600083015261100e81610e45565b9050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b60006110bf82611156565b91506110ca83611156565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156110ff576110fe6111db565b5b828201905092915050565b600061111582611136565b9050919050565b60008115159050919050565b600061ffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061116b82611172565b9050919050565b600061117d82611136565b9050919050565b600061118f82611196565b9050919050565b60006111a182611136565b9050919050565b60005b838110156111c65780820151818401526020810190506111ab565b838111156111d5576000848401525b50505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6112c58161110a565b81146112d057600080fd5b50565b6112dc8161111c565b81146112e757600080fd5b5056fea264697066735822122079447b0bc5a0d909e830a2f726a140debea410ccbe10c376a519370a087fe46c64736f6c63430008040033"

// DeployFlashLoan deploys a new Ethereum contract, binding an instance of FlashLoan to it.
func DeployFlashLoan(auth *bind.TransactOpts, backend bind.ContractBackend, _addressProvider common.Address) (common.Address, *types.Transaction, *FlashLoan, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashLoanABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FlashLoanBin), backend, _addressProvider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FlashLoan{FlashLoanCaller: FlashLoanCaller{contract: contract}, FlashLoanTransactor: FlashLoanTransactor{contract: contract}, FlashLoanFilterer: FlashLoanFilterer{contract: contract}}, nil
}

// FlashLoan is an auto generated Go binding around an Ethereum contract.
type FlashLoan struct {
	FlashLoanCaller     // Read-only binding to the contract
	FlashLoanTransactor // Write-only binding to the contract
	FlashLoanFilterer   // Log filterer for contract events
}

// FlashLoanCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashLoanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashLoanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashLoanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashLoanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashLoanSession struct {
	Contract     *FlashLoan        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashLoanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashLoanCallerSession struct {
	Contract *FlashLoanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FlashLoanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashLoanTransactorSession struct {
	Contract     *FlashLoanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FlashLoanRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashLoanRaw struct {
	Contract *FlashLoan // Generic contract binding to access the raw methods on
}

// FlashLoanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashLoanCallerRaw struct {
	Contract *FlashLoanCaller // Generic read-only contract binding to access the raw methods on
}

// FlashLoanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashLoanTransactorRaw struct {
	Contract *FlashLoanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashLoan creates a new instance of FlashLoan, bound to a specific deployed contract.
func NewFlashLoan(address common.Address, backend bind.ContractBackend) (*FlashLoan, error) {
	contract, err := bindFlashLoan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashLoan{FlashLoanCaller: FlashLoanCaller{contract: contract}, FlashLoanTransactor: FlashLoanTransactor{contract: contract}, FlashLoanFilterer: FlashLoanFilterer{contract: contract}}, nil
}

// NewFlashLoanCaller creates a new read-only instance of FlashLoan, bound to a specific deployed contract.
func NewFlashLoanCaller(address common.Address, caller bind.ContractCaller) (*FlashLoanCaller, error) {
	contract, err := bindFlashLoan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanCaller{contract: contract}, nil
}

// NewFlashLoanTransactor creates a new write-only instance of FlashLoan, bound to a specific deployed contract.
func NewFlashLoanTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashLoanTransactor, error) {
	contract, err := bindFlashLoan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashLoanTransactor{contract: contract}, nil
}

// NewFlashLoanFilterer creates a new log filterer instance of FlashLoan, bound to a specific deployed contract.
func NewFlashLoanFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashLoanFilterer, error) {
	contract, err := bindFlashLoan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashLoanFilterer{contract: contract}, nil
}

// bindFlashLoan binds a generic wrapper to an already deployed contract.
func bindFlashLoan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashLoanABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoan *FlashLoanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoan.Contract.FlashLoanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoan *FlashLoanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoan.Contract.FlashLoanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoan *FlashLoanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoan.Contract.FlashLoanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashLoan *FlashLoanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashLoan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashLoan *FlashLoanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashLoan *FlashLoanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashLoan.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_FlashLoan *FlashLoanCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLoan.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_FlashLoan *FlashLoanSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _FlashLoan.Contract.ADDRESSESPROVIDER(&_FlashLoan.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_FlashLoan *FlashLoanCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _FlashLoan.Contract.ADDRESSESPROVIDER(&_FlashLoan.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_FlashLoan *FlashLoanCaller) LENDINGPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLoan.contract.Call(opts, &out, "LENDING_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_FlashLoan *FlashLoanSession) LENDINGPOOL() (common.Address, error) {
	return _FlashLoan.Contract.LENDINGPOOL(&_FlashLoan.CallOpts)
}

// LENDINGPOOL is a free data retrieval call binding the contract method 0xb4dcfc77.
//
// Solidity: function LENDING_POOL() view returns(address)
func (_FlashLoan *FlashLoanCallerSession) LENDINGPOOL() (common.Address, error) {
	return _FlashLoan.Contract.LENDINGPOOL(&_FlashLoan.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x919840ad.
//
// Solidity: function check() view returns(bool)
func (_FlashLoan *FlashLoanCaller) Check(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlashLoan.contract.Call(opts, &out, "check")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0x919840ad.
//
// Solidity: function check() view returns(bool)
func (_FlashLoan *FlashLoanSession) Check() (bool, error) {
	return _FlashLoan.Contract.Check(&_FlashLoan.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x919840ad.
//
// Solidity: function check() view returns(bool)
func (_FlashLoan *FlashLoanCallerSession) Check() (bool, error) {
	return _FlashLoan.Contract.Check(&_FlashLoan.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FlashLoan *FlashLoanCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FlashLoan.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FlashLoan *FlashLoanSession) IsOwner() (bool, error) {
	return _FlashLoan.Contract.IsOwner(&_FlashLoan.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FlashLoan *FlashLoanCallerSession) IsOwner() (bool, error) {
	return _FlashLoan.Contract.IsOwner(&_FlashLoan.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashLoan *FlashLoanCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlashLoan.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashLoan *FlashLoanSession) Owner() (common.Address, error) {
	return _FlashLoan.Contract.Owner(&_FlashLoan.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FlashLoan *FlashLoanCallerSession) Owner() (common.Address, error) {
	return _FlashLoan.Contract.Owner(&_FlashLoan.CallOpts)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_FlashLoan *FlashLoanTransactor) ExecuteOperation(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _FlashLoan.contract.Transact(opts, "executeOperation", assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_FlashLoan *FlashLoanSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _FlashLoan.Contract.ExecuteOperation(&_FlashLoan.TransactOpts, assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_FlashLoan *FlashLoanTransactorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _FlashLoan.Contract.ExecuteOperation(&_FlashLoan.TransactOpts, assets, amounts, premiums, initiator, params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x9ad5981e.
//
// Solidity: function flashLoan() returns()
func (_FlashLoan *FlashLoanTransactor) FlashLoan(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoan.contract.Transact(opts, "flashLoan")
}

// FlashLoan is a paid mutator transaction binding the contract method 0x9ad5981e.
//
// Solidity: function flashLoan() returns()
func (_FlashLoan *FlashLoanSession) FlashLoan() (*types.Transaction, error) {
	return _FlashLoan.Contract.FlashLoan(&_FlashLoan.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x9ad5981e.
//
// Solidity: function flashLoan() returns()
func (_FlashLoan *FlashLoanTransactorSession) FlashLoan() (*types.Transaction, error) {
	return _FlashLoan.Contract.FlashLoan(&_FlashLoan.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FlashLoan *FlashLoanTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashLoan.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FlashLoan *FlashLoanSession) RenounceOwnership() (*types.Transaction, error) {
	return _FlashLoan.Contract.RenounceOwnership(&_FlashLoan.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FlashLoan *FlashLoanTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FlashLoan.Contract.RenounceOwnership(&_FlashLoan.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FlashLoan *FlashLoanTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FlashLoan.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FlashLoan *FlashLoanSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FlashLoan.Contract.TransferOwnership(&_FlashLoan.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FlashLoan *FlashLoanTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FlashLoan.Contract.TransferOwnership(&_FlashLoan.TransactOpts, newOwner)
}

// FlashLoanOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FlashLoan contract.
type FlashLoanOwnershipTransferredIterator struct {
	Event *FlashLoanOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FlashLoanOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashLoanOwnershipTransferred)
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
		it.Event = new(FlashLoanOwnershipTransferred)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlashLoanOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashLoanOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashLoanOwnershipTransferred represents a OwnershipTransferred event raised by the FlashLoan contract.
type FlashLoanOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FlashLoan *FlashLoanFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FlashLoanOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FlashLoan.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlashLoanOwnershipTransferredIterator{contract: _FlashLoan.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FlashLoan *FlashLoanFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FlashLoanOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FlashLoan.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashLoanOwnershipTransferred)
				if err := _FlashLoan.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FlashLoan *FlashLoanFilterer) ParseOwnershipTransferred(log types.Log) (*FlashLoanOwnershipTransferred, error) {
	event := new(FlashLoanOwnershipTransferred)
	if err := _FlashLoan.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
