// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vockTails

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
)

// VockTailsMetaData contains all meta data concerning the VockTails contract.
var VockTailsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"drinkVRF\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COORDINATOR\",\"outputs\":[{\"internalType\":\"contractVRFCoordinatorV2Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINKTOKEN\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomDrinkValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50737a1bac17ccc5b313516c5e16fb24f7659aa5ebed8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050737a1bac17ccc5b313516c5e16fb24f7659aa5ebed6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073326c977e6efc84e512bb9c30f76e30c160ed06fb600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080516109cd6101266000396000818161010c015261016001526109cd6000f3fe608060405234801561001057600080fd5b50600436106100615760003560e01c80626d6cae146100665780631fe543e3146100845780633b2bcbf1146100a057806355380dfb146100be578063ad2f0001146100dc578063e0c86289146100fa575b600080fd5b61006e610104565b60405161007b9190610370565b60405180910390f35b61009e60048036038101906100999190610524565b61010a565b005b6100a86101ca565b6040516100b591906105ff565b60405180910390f35b6100c66101ee565b6040516100d3919061063b565b60405180910390f35b6100e4610214565b6040516100f19190610370565b60405180910390f35b61010261021a565b005b60025481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101bc57337f00000000000000000000000000000000000000000000000000000000000000006040517f1cf993f40000000000000000000000000000000000000000000000000000000081526004016101b3929190610677565b60405180910390fd5b6101c682826102ed565b5050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d3b1d307f4b09e658ed251bcafeebbc69400383d49f344ace09b9576fe248bb02c003fe9f6108a96003622625a060016040518663ffffffff1660e01b81526004016102a2959493929190610810565b6020604051808303816000875af11580156102c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e59190610878565b600281905550565b6001600382600081518110610305576103046108a5565b5b60200260200101516103179190610903565b6103219190610963565b6003819055507f43a8f2dfef938718c1c549d3e9e20a7e142b565db7b34b2e1a68335cc5c5fad760405160405180910390a15050565b6000819050919050565b61036a81610357565b82525050565b60006020820190506103856000830184610361565b92915050565b6000604051905090565b600080fd5b600080fd5b6103a881610357565b81146103b357600080fd5b50565b6000813590506103c58161039f565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610419826103d0565b810181811067ffffffffffffffff82111715610438576104376103e1565b5b80604052505050565b600061044b61038b565b90506104578282610410565b919050565b600067ffffffffffffffff821115610477576104766103e1565b5b602082029050602081019050919050565b600080fd5b60006104a061049b8461045c565b610441565b905080838252602082019050602084028301858111156104c3576104c2610488565b5b835b818110156104ec57806104d888826103b6565b8452602084019350506020810190506104c5565b5050509392505050565b600082601f83011261050b5761050a6103cb565b5b813561051b84826020860161048d565b91505092915050565b6000806040838503121561053b5761053a610395565b5b6000610549858286016103b6565b925050602083013567ffffffffffffffff81111561056a5761056961039a565b5b610576858286016104f6565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006105c56105c06105bb84610580565b6105a0565b610580565b9050919050565b60006105d7826105aa565b9050919050565b60006105e9826105cc565b9050919050565b6105f9816105de565b82525050565b600060208201905061061460008301846105f0565b92915050565b6000610625826105cc565b9050919050565b6106358161061a565b82525050565b6000602082019050610650600083018461062c565b92915050565b600061066182610580565b9050919050565b61067181610656565b82525050565b600060408201905061068c6000830185610668565b6106996020830184610668565b9392505050565b6000819050919050565b6000819050919050565b60008160001b9050919050565b60006106dc6106d76106d2846106a0565b6106b4565b6106aa565b9050919050565b6106ec816106c1565b82525050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600061072b610726610721846106f2565b6105a0565b6106fc565b9050919050565b61073b81610710565b82525050565b6000819050919050565b600061ffff82169050919050565b600061077461076f61076a84610741565b6105a0565b61074b565b9050919050565b61078481610759565b82525050565b6000819050919050565b600063ffffffff82169050919050565b60006107bf6107ba6107b58461078a565b6105a0565b610794565b9050919050565b6107cf816107a4565b82525050565b6000819050919050565b60006107fa6107f56107f0846107d5565b6105a0565b610794565b9050919050565b61080a816107df565b82525050565b600060a08201905061082560008301886106e3565b6108326020830187610732565b61083f604083018661077b565b61084c60608301856107c6565b6108596080830184610801565b9695505050505050565b6000815190506108728161039f565b92915050565b60006020828403121561088e5761088d610395565b5b600061089c84828501610863565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061090e82610357565b915061091983610357565b925082610929576109286108d4565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061096e82610357565b915061097983610357565b925082820190508082111561099157610990610934565b5b9291505056fea26469706673582212203c69c2b26d964377ad3b257dedd1523aeaca226c702be0a773169b7f7b8920e364736f6c63430008110033",
}

// VockTailsABI is the input ABI used to generate the binding from.
// Deprecated: Use VockTailsMetaData.ABI instead.
var VockTailsABI = VockTailsMetaData.ABI

// VockTailsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VockTailsMetaData.Bin instead.
var VockTailsBin = VockTailsMetaData.Bin

// DeployVockTails deploys a new Ethereum contract, binding an instance of VockTails to it.
func DeployVockTails(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VockTails, error) {
	parsed, err := VockTailsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VockTailsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VockTails{VockTailsCaller: VockTailsCaller{contract: contract}, VockTailsTransactor: VockTailsTransactor{contract: contract}, VockTailsFilterer: VockTailsFilterer{contract: contract}}, nil
}

// VockTails is an auto generated Go binding around an Ethereum contract.
type VockTails struct {
	VockTailsCaller     // Read-only binding to the contract
	VockTailsTransactor // Write-only binding to the contract
	VockTailsFilterer   // Log filterer for contract events
}

// VockTailsCaller is an auto generated read-only Go binding around an Ethereum contract.
type VockTailsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VockTailsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VockTailsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VockTailsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VockTailsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VockTailsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VockTailsSession struct {
	Contract     *VockTails        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VockTailsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VockTailsCallerSession struct {
	Contract *VockTailsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VockTailsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VockTailsTransactorSession struct {
	Contract     *VockTailsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VockTailsRaw is an auto generated low-level Go binding around an Ethereum contract.
type VockTailsRaw struct {
	Contract *VockTails // Generic contract binding to access the raw methods on
}

// VockTailsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VockTailsCallerRaw struct {
	Contract *VockTailsCaller // Generic read-only contract binding to access the raw methods on
}

// VockTailsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VockTailsTransactorRaw struct {
	Contract *VockTailsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVockTails creates a new instance of VockTails, bound to a specific deployed contract.
func NewVockTails(address common.Address, backend bind.ContractBackend) (*VockTails, error) {
	contract, err := bindVockTails(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VockTails{VockTailsCaller: VockTailsCaller{contract: contract}, VockTailsTransactor: VockTailsTransactor{contract: contract}, VockTailsFilterer: VockTailsFilterer{contract: contract}}, nil
}

// NewVockTailsCaller creates a new read-only instance of VockTails, bound to a specific deployed contract.
func NewVockTailsCaller(address common.Address, caller bind.ContractCaller) (*VockTailsCaller, error) {
	contract, err := bindVockTails(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VockTailsCaller{contract: contract}, nil
}

// NewVockTailsTransactor creates a new write-only instance of VockTails, bound to a specific deployed contract.
func NewVockTailsTransactor(address common.Address, transactor bind.ContractTransactor) (*VockTailsTransactor, error) {
	contract, err := bindVockTails(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VockTailsTransactor{contract: contract}, nil
}

// NewVockTailsFilterer creates a new log filterer instance of VockTails, bound to a specific deployed contract.
func NewVockTailsFilterer(address common.Address, filterer bind.ContractFilterer) (*VockTailsFilterer, error) {
	contract, err := bindVockTails(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VockTailsFilterer{contract: contract}, nil
}

// bindVockTails binds a generic wrapper to an already deployed contract.
func bindVockTails(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VockTailsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VockTails *VockTailsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VockTails.Contract.VockTailsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VockTails *VockTailsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VockTails.Contract.VockTailsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VockTails *VockTailsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VockTails.Contract.VockTailsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VockTails *VockTailsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VockTails.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VockTails *VockTailsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VockTails.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VockTails *VockTailsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VockTails.Contract.contract.Transact(opts, method, params...)
}

// COORDINATOR is a free data retrieval call binding the contract method 0x3b2bcbf1.
//
// Solidity: function COORDINATOR() view returns(address)
func (_VockTails *VockTailsCaller) COORDINATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VockTails.contract.Call(opts, &out, "COORDINATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COORDINATOR is a free data retrieval call binding the contract method 0x3b2bcbf1.
//
// Solidity: function COORDINATOR() view returns(address)
func (_VockTails *VockTailsSession) COORDINATOR() (common.Address, error) {
	return _VockTails.Contract.COORDINATOR(&_VockTails.CallOpts)
}

// COORDINATOR is a free data retrieval call binding the contract method 0x3b2bcbf1.
//
// Solidity: function COORDINATOR() view returns(address)
func (_VockTails *VockTailsCallerSession) COORDINATOR() (common.Address, error) {
	return _VockTails.Contract.COORDINATOR(&_VockTails.CallOpts)
}

// LINKTOKEN is a free data retrieval call binding the contract method 0x55380dfb.
//
// Solidity: function LINKTOKEN() view returns(address)
func (_VockTails *VockTailsCaller) LINKTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VockTails.contract.Call(opts, &out, "LINKTOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINKTOKEN is a free data retrieval call binding the contract method 0x55380dfb.
//
// Solidity: function LINKTOKEN() view returns(address)
func (_VockTails *VockTailsSession) LINKTOKEN() (common.Address, error) {
	return _VockTails.Contract.LINKTOKEN(&_VockTails.CallOpts)
}

// LINKTOKEN is a free data retrieval call binding the contract method 0x55380dfb.
//
// Solidity: function LINKTOKEN() view returns(address)
func (_VockTails *VockTailsCallerSession) LINKTOKEN() (common.Address, error) {
	return _VockTails.Contract.LINKTOKEN(&_VockTails.CallOpts)
}

// RandomDrinkValue is a free data retrieval call binding the contract method 0xad2f0001.
//
// Solidity: function randomDrinkValue() view returns(uint256)
func (_VockTails *VockTailsCaller) RandomDrinkValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VockTails.contract.Call(opts, &out, "randomDrinkValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomDrinkValue is a free data retrieval call binding the contract method 0xad2f0001.
//
// Solidity: function randomDrinkValue() view returns(uint256)
func (_VockTails *VockTailsSession) RandomDrinkValue() (*big.Int, error) {
	return _VockTails.Contract.RandomDrinkValue(&_VockTails.CallOpts)
}

// RandomDrinkValue is a free data retrieval call binding the contract method 0xad2f0001.
//
// Solidity: function randomDrinkValue() view returns(uint256)
func (_VockTails *VockTailsCallerSession) RandomDrinkValue() (*big.Int, error) {
	return _VockTails.Contract.RandomDrinkValue(&_VockTails.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(uint256)
func (_VockTails *VockTailsCaller) RequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VockTails.contract.Call(opts, &out, "requestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(uint256)
func (_VockTails *VockTailsSession) RequestId() (*big.Int, error) {
	return _VockTails.Contract.RequestId(&_VockTails.CallOpts)
}

// RequestId is a free data retrieval call binding the contract method 0x006d6cae.
//
// Solidity: function requestId() view returns(uint256)
func (_VockTails *VockTailsCallerSession) RequestId() (*big.Int, error) {
	return _VockTails.Contract.RequestId(&_VockTails.CallOpts)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VockTails *VockTailsTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VockTails.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VockTails *VockTailsSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VockTails.Contract.RawFulfillRandomWords(&_VockTails.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VockTails *VockTailsTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VockTails.Contract.RawFulfillRandomWords(&_VockTails.TransactOpts, requestId, randomWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0xe0c86289.
//
// Solidity: function requestRandomWords() returns()
func (_VockTails *VockTailsTransactor) RequestRandomWords(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VockTails.contract.Transact(opts, "requestRandomWords")
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0xe0c86289.
//
// Solidity: function requestRandomWords() returns()
func (_VockTails *VockTailsSession) RequestRandomWords() (*types.Transaction, error) {
	return _VockTails.Contract.RequestRandomWords(&_VockTails.TransactOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0xe0c86289.
//
// Solidity: function requestRandomWords() returns()
func (_VockTails *VockTailsTransactorSession) RequestRandomWords() (*types.Transaction, error) {
	return _VockTails.Contract.RequestRandomWords(&_VockTails.TransactOpts)
}

// VockTailsDrinkVRFIterator is returned from FilterDrinkVRF and is used to iterate over the raw logs and unpacked data for DrinkVRF events raised by the VockTails contract.
type VockTailsDrinkVRFIterator struct {
	Event *VockTailsDrinkVRF // Event containing the contract specifics and raw log

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
func (it *VockTailsDrinkVRFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VockTailsDrinkVRF)
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
		it.Event = new(VockTailsDrinkVRF)
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
func (it *VockTailsDrinkVRFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VockTailsDrinkVRFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VockTailsDrinkVRF represents a DrinkVRF event raised by the VockTails contract.
type VockTailsDrinkVRF struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDrinkVRF is a free log retrieval operation binding the contract event 0x43a8f2dfef938718c1c549d3e9e20a7e142b565db7b34b2e1a68335cc5c5fad7.
//
// Solidity: event drinkVRF()
func (_VockTails *VockTailsFilterer) FilterDrinkVRF(opts *bind.FilterOpts) (*VockTailsDrinkVRFIterator, error) {

	logs, sub, err := _VockTails.contract.FilterLogs(opts, "drinkVRF")
	if err != nil {
		return nil, err
	}
	return &VockTailsDrinkVRFIterator{contract: _VockTails.contract, event: "drinkVRF", logs: logs, sub: sub}, nil
}

// WatchDrinkVRF is a free log subscription operation binding the contract event 0x43a8f2dfef938718c1c549d3e9e20a7e142b565db7b34b2e1a68335cc5c5fad7.
//
// Solidity: event drinkVRF()
func (_VockTails *VockTailsFilterer) WatchDrinkVRF(opts *bind.WatchOpts, sink chan<- *VockTailsDrinkVRF) (event.Subscription, error) {

	logs, sub, err := _VockTails.contract.WatchLogs(opts, "drinkVRF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VockTailsDrinkVRF)
				if err := _VockTails.contract.UnpackLog(event, "drinkVRF", log); err != nil {
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

// ParseDrinkVRF is a log parse operation binding the contract event 0x43a8f2dfef938718c1c549d3e9e20a7e142b565db7b34b2e1a68335cc5c5fad7.
//
// Solidity: event drinkVRF()
func (_VockTails *VockTailsFilterer) ParseDrinkVRF(log types.Log) (*VockTailsDrinkVRF, error) {
	event := new(VockTailsDrinkVRF)
	if err := _VockTails.contract.UnpackLog(event, "drinkVRF", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
