// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gweiPump

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

// GweiPumpMetaData contains all meta data concerning the GweiPump contract.
var GweiPumpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"msgValueTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"notOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"oraclePriceFeedZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"pumpNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"upKeepNotNeeded\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"oilBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"updateWti\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Wti40Milliliters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WtiPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyOil40Milliliters\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainlinkNodeRequestWtiPrice\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20LINK\",\"outputs\":[{\"internalType\":\"contractERC20TokenContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"memoryWtiPriceOracle\",\"type\":\"uint256\"}],\"name\":\"fulfill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestMaticUsd\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestWtiMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPumpFilled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastWtiPriceCheckUnixTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manualUpKeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"name\":\"ownerPumpFilledStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"performUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526001600455600160065534801561001a57600080fd5b503373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505073d0d5e3db44de05e9f294bb0a3beeaf030de24ada600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555073326c977e6efc84e512bb9c30f76e30c160ed06fb600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061011c73326c977e6efc84e512bb9c30f76e30c160ed06fb61012160201b60201c565b610165565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60805161233b61018e600039600081816107650152818161094d0152610b37015261233b6000f3fe6080604052600436106100e85760003560e01c8063861b31ce1161008a578063d42506b811610059578063d42506b8146102c1578063d6666612146102ea578063d698a37e14610315578063f5da229e14610340576100e8565b8063861b31ce14610215578063b4a99a4e14610240578063bcc9f5911461026b578063c62d34ab14610296576100e8565b80635fce9d34116100c65780635fce9d341461016a5780636e04ff0d14610195578063784eb75b146101d35780637d5ded7c146101fe576100e8565b80632514f7e3146100ed5780634357855e146101185780634585e33b14610141575b600080fd5b3480156100f957600080fd5b5061010261034a565b60405161010f9190611632565b60405180910390f35b34801561012457600080fd5b5061013f600480360381019061013a91906116b9565b610374565b005b34801561014d57600080fd5b506101686004803603810190610163919061175e565b6104c2565b005b34801561017657600080fd5b5061017f6104cf565b60405161018c9190611632565b60405180910390f35b3480156101a157600080fd5b506101bc60048036038101906101b7919061175e565b6104d5565b6040516101ca929190611856565b60405180910390f35b3480156101df57600080fd5b506101e86105a4565b6040516101f5919061189f565b60405180910390f35b34801561020a57600080fd5b50610213610653565b005b34801561022157600080fd5b5061022a61075d565b6040516102379190611632565b60405180910390f35b34801561024c57600080fd5b50610255610763565b60405161026291906118fb565b60405180910390f35b34801561027757600080fd5b50610280610787565b60405161028d9190611925565b60405180910390f35b3480156102a257600080fd5b506102ab610902565b6040516102b89190611632565b60405180910390f35b3480156102cd57600080fd5b506102e860048036038101906102e39190611940565b61094b565b005b3480156102f657600080fd5b506102ff6109da565b60405161030c91906119cc565b60405180910390f35b34801561032157600080fd5b5061032a610a00565b6040516103379190611632565b60405180910390f35b610348610a06565b005b600062f29869610358610902565b610fa06103659190611a16565b61036f9190611a87565b905090565b816005600082815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610416576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040d90611b3b565b60405180910390fd5b6005600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055807f7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a60405160405180910390a2600082111561048a57816008819055505b426007819055507f90c6eb475b1dc27758b0b1c75bd44d59a523bab636839816edfad763432d1e5a60405160405180910390a1505050565b6104ca610787565b505050565b60065481565b60006060620151806007546104ea9190611b5b565b421015801561059b5750662386f26fc10000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161055791906118fb565b602060405180830381865afa158015610574573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105989190611ba4565b10155b91509250929050565b600080600080600080600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa15801561061a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063e9190611c3f565b94509450945094509450839550505050505090565b620151806007546106649190611b5b565b42101580156107155750662386f26fc10000600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106d191906118fb565b602060405180830381865afa1580156106ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107129190611ba4565b10155b15156000151503610752576040517ff33f738f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61075a610787565b50565b60075481565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000806107bc7f626266306261646164323964343964633838373530346261636662623930356230634357855e60e01b610bca565b90506108206040518060400160405280600381526020017f67657400000000000000000000000000000000000000000000000000000000008152506040518060a00160405280607f8152602001612287607f913983610bfb9092919063ffffffff16565b6108826040518060400160405280600481526020017f70617468000000000000000000000000000000000000000000000000000000008152506040518060600160405280603481526020016122536034913983610bfb9092919063ffffffff16565b60006305f5e10090506108d56040518060400160405280600581526020017f74696d65730000000000000000000000000000000000000000000000000000008152508284610c2e9092919063ffffffff16565b6108fb73c8d925525ca8759812d0c299b90247917d4d4b7c83662386f26fc10000610c61565b9250505090565b600061090c6105a4565b6103e8670de0b6b3a76400006103eb6008546109289190611a16565b6109329190611cba565b61093c9190611d32565b6109469190611d32565b905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109d0576040517f251c9d6300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060068190555050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60085481565b600060065403610a42576040517fee94d61d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610a4c61034a565b03610a83576040517fd5da0f4300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a8b61034a565b341015610ac4576040517f4837092100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600681905550610ad461034a565b341115610b35573373ffffffffffffffffffffffffffffffffffffffff166108fc610afd61034a565b34610b089190611d9c565b9081150290604051600060405180830381858888f19350505050158015610b33573d6000803e3d6000fd5b505b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610b9b573d6000803e3d6000fd5b507f260abb0ce3f225ef1530c6f1a606b8331910fd3f66650c133f35b5817624ffe260405160405180910390a1565b610bd2611592565b610bda611592565b610bf185858584610d2d909392919063ffffffff16565b9150509392505050565b610c12828460800151610ddd90919063ffffffff16565b610c29818460800151610ddd90919063ffffffff16565b505050565b610c45828460800151610ddd90919063ffffffff16565b610c5c818460800151610e0290919063ffffffff16565b505050565b6000806004549050600181610c769190611b5b565b6004819055506000634042994660e01b60008087600001513089604001518760018c6080015160000151604051602401610cb7989796959493929190611e0b565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050610d2286838684610eaf565b925050509392505050565b610d35611592565b610d458560800151610100611044565b508385600001818152505082856020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508185604001907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050849050949350505050565b610dea82600383516110ae565b610dfd818361123390919063ffffffff16565b505050565b7fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000811215610e3957610e348282611255565b610eab565b67ffffffffffffffff811315610e5857610e5382826112cc565b610eaa565b60008112610e7157610e6c826000836110ae565b610ea9565b610ea8826001837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610ea39190611e90565b6110ae565b5b5b5b5050565b60003084604051602001610ec4929190611f4e565b604051602081830303815290604052805190602001209050846005600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550807fb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af960405160405180910390a2600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634000aea08685856040518463ffffffff1660e01b8152600401610fba93929190611f7a565b6020604051808303816000875af1158015610fd9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ffd9190611fe4565b61103c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103390612083565b60405180910390fd5b949350505050565b61104c6115ff565b600060208361105b91906120a3565b146110875760208261106d91906120a3565b60206110799190611d9c565b826110849190611b5b565b91505b81836020018181525050604051808452600081528281016020016040525082905092915050565b60178167ffffffffffffffff16116110e5576110df8160058460ff16901b60ff16178461131890919063ffffffff16565b5061122e565b60ff8167ffffffffffffffff161161113b57611114601860058460ff16901b178461131890919063ffffffff16565b506111358167ffffffffffffffff166001856113389092919063ffffffff16565b5061122d565b61ffff8167ffffffffffffffff16116111925761116b601960058460ff16901b178461131890919063ffffffff16565b5061118c8167ffffffffffffffff166002856113389092919063ffffffff16565b5061122c565b63ffffffff8167ffffffffffffffff16116111eb576111c4601a60058460ff16901b178461131890919063ffffffff16565b506111e58167ffffffffffffffff166004856113389092919063ffffffff16565b5061122b565b611208601b60058460ff16901b178461131890919063ffffffff16565b506112298167ffffffffffffffff166008856113389092919063ffffffff16565b505b5b5b5b505050565b61123b6115ff565b61124d8384600001515184855161135a565b905092915050565b61127360036005600660ff16901b178361131890919063ffffffff16565b506112c882827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6112a49190611e90565b6040516020016112b49190611632565b604051602081830303815290604052611449565b5050565b6112ea60026005600660ff16901b178361131890919063ffffffff16565b5061131482826040516020016113009190611632565b604051602081830303815290604052611449565b5050565b6113206115ff565b611330838460000151518461146e565b905092915050565b6113406115ff565b6113518485600001515185856114c4565b90509392505050565b6113626115ff565b825182111561137057600080fd5b846020015182856113819190611b5b565b11156113b6576113b58560026113a6886020015188876113a19190611b5b565b611552565b6113b09190611a16565b61156e565b5b6000808651805187602083010193508088870111156113d55787860182525b60208701925050505b6020841061141c57805182526020826113f79190611b5b565b91506020816114069190611b5b565b90506020846114159190611d9c565b93506113de565b60006001856020036101000a03905080198251168184511681811785525050508692505050949350505050565b61145682600283516110ae565b611469818361123390919063ffffffff16565b505050565b6114766115ff565b8360200151831061149c5761149b84600286602001516114969190611a16565b61156e565b5b835180516020858301018481538186036114b7576001820183525b5050508390509392505050565b6114cc6115ff565b846020015184836114dd9190611b5b565b11156115055761150485600286856114f59190611b5b565b6114ff9190611a16565b61156e565b5b60006001836101006115179190612207565b6115219190611d9c565b905085518386820101858319825116178152815185880111156115445784870182525b505085915050949350505050565b60008183111561156457829050611568565b8190505b92915050565b6000826000015190506115818383611044565b5061158c8382611233565b50505050565b6040518060a0016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001600081526020016115f96115ff565b81525090565b604051806040016040528060608152602001600081525090565b6000819050919050565b61162c81611619565b82525050565b60006020820190506116476000830184611623565b92915050565b600080fd5b600080fd5b6000819050919050565b61166a81611657565b811461167557600080fd5b50565b60008135905061168781611661565b92915050565b61169681611619565b81146116a157600080fd5b50565b6000813590506116b38161168d565b92915050565b600080604083850312156116d0576116cf61164d565b5b60006116de85828601611678565b92505060206116ef858286016116a4565b9150509250929050565b600080fd5b600080fd5b600080fd5b60008083601f84011261171e5761171d6116f9565b5b8235905067ffffffffffffffff81111561173b5761173a6116fe565b5b60208301915083600182028301111561175757611756611703565b5b9250929050565b600080602083850312156117755761177461164d565b5b600083013567ffffffffffffffff81111561179357611792611652565b5b61179f85828601611708565b92509250509250929050565b60008115159050919050565b6117c0816117ab565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156118005780820151818401526020810190506117e5565b60008484015250505050565b6000601f19601f8301169050919050565b6000611828826117c6565b61183281856117d1565b93506118428185602086016117e2565b61184b8161180c565b840191505092915050565b600060408201905061186b60008301856117b7565b818103602083015261187d818461181d565b90509392505050565b6000819050919050565b61189981611886565b82525050565b60006020820190506118b46000830184611890565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006118e5826118ba565b9050919050565b6118f5816118da565b82525050565b600060208201905061191060008301846118ec565b92915050565b61191f81611657565b82525050565b600060208201905061193a6000830184611916565b92915050565b6000602082840312156119565761195561164d565b5b6000611964848285016116a4565b91505092915050565b6000819050919050565b600061199261198d611988846118ba565b61196d565b6118ba565b9050919050565b60006119a482611977565b9050919050565b60006119b682611999565b9050919050565b6119c6816119ab565b82525050565b60006020820190506119e160008301846119bd565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611a2182611619565b9150611a2c83611619565b9250828202611a3a81611619565b91508282048414831517611a5157611a506119e7565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611a9282611619565b9150611a9d83611619565b925082611aad57611aac611a58565b5b828204905092915050565b600082825260208201905092915050565b7f536f75726365206d75737420626520746865206f7261636c65206f662074686560008201527f2072657175657374000000000000000000000000000000000000000000000000602082015250565b6000611b25602883611ab8565b9150611b3082611ac9565b604082019050919050565b60006020820190508181036000830152611b5481611b18565b9050919050565b6000611b6682611619565b9150611b7183611619565b9250828201905080821115611b8957611b886119e7565b5b92915050565b600081519050611b9e8161168d565b92915050565b600060208284031215611bba57611bb961164d565b5b6000611bc884828501611b8f565b91505092915050565b600069ffffffffffffffffffff82169050919050565b611bf081611bd1565b8114611bfb57600080fd5b50565b600081519050611c0d81611be7565b92915050565b611c1c81611886565b8114611c2757600080fd5b50565b600081519050611c3981611c13565b92915050565b600080600080600060a08688031215611c5b57611c5a61164d565b5b6000611c6988828901611bfe565b9550506020611c7a88828901611c2a565b9450506040611c8b88828901611b8f565b9350506060611c9c88828901611b8f565b9250506080611cad88828901611bfe565b9150509295509295909350565b6000611cc582611886565b9150611cd083611886565b9250828202611cde81611886565b91507f80000000000000000000000000000000000000000000000000000000000000008414600084121615611d1657611d156119e7565b5b8282058414831517611d2b57611d2a6119e7565b5b5092915050565b6000611d3d82611886565b9150611d4883611886565b925082611d5857611d57611a58565b5b600160000383147f800000000000000000000000000000000000000000000000000000000000000083141615611d9157611d906119e7565b5b828205905092915050565b6000611da782611619565b9150611db283611619565b9250828203905081811115611dca57611dc96119e7565b5b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611e0581611dd0565b82525050565b600061010082019050611e21600083018b6118ec565b611e2e602083018a611623565b611e3b6040830189611916565b611e4860608301886118ec565b611e556080830187611dfc565b611e6260a0830186611623565b611e6f60c0830185611623565b81810360e0830152611e81818461181d565b90509998505050505050505050565b6000611e9b82611886565b9150611ea683611886565b9250828203905081811260008412168282136000851215161715611ecd57611ecc6119e7565b5b92915050565b6000611ede82611999565b9050919050565b60008160601b9050919050565b6000611efd82611ee5565b9050919050565b6000611f0f82611ef2565b9050919050565b611f27611f2282611ed3565b611f04565b82525050565b6000819050919050565b611f48611f4382611619565b611f2d565b82525050565b6000611f5a8285611f16565b601482019150611f6a8284611f37565b6020820191508190509392505050565b6000606082019050611f8f60008301866118ec565b611f9c6020830185611623565b8181036040830152611fae818461181d565b9050949350505050565b611fc1816117ab565b8114611fcc57600080fd5b50565b600081519050611fde81611fb8565b92915050565b600060208284031215611ffa57611ff961164d565b5b600061200884828501611fcf565b91505092915050565b7f756e61626c6520746f207472616e73666572416e6443616c6c20746f206f726160008201527f636c650000000000000000000000000000000000000000000000000000000000602082015250565b600061206d602383611ab8565b915061207882612011565b604082019050919050565b6000602082019050818103600083015261209c81612060565b9050919050565b60006120ae82611619565b91506120b983611619565b9250826120c9576120c8611a58565b5b828206905092915050565b60008160011c9050919050565b6000808291508390505b600185111561212b57808604811115612107576121066119e7565b5b60018516156121165780820291505b8081029050612124856120d4565b94506120eb565b94509492505050565b6000826121445760019050612200565b816121525760009050612200565b81600181146121685760028114612172576121a1565b6001915050612200565b60ff841115612184576121836119e7565b5b8360020a91508482111561219b5761219a6119e7565b5b50612200565b5060208310610133831016604e8410600b84101617156121d65782820a9050838111156121d1576121d06119e7565b5b612200565b6121e384848460016120e1565b925090508184048111156121fa576121f96119e7565b5b81810290505b9392505050565b600061221282611619565b915061221d83611619565b925061224a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612134565b90509291505056fe7265636f7264732e302e6669656c64732e63757368696e675f6f6b5f7774695f73706f745f70726963655f666f625f6461696c7968747470733a2f2f64617461736f757263652e6b6170736172632e6f72672f6170692f7265636f7264732f312e302f7365617263682f3f646174617365743d73706f742d7072696365732d666f722d63727564652d6f696c2d616e642d706574726f6c65756d2d70726f647563747326713d2666616365743d706572696f64a26469706673582212206068a8e288c8d1e5bbd34072e19b464f2fb4c23e23ae0e72ac9256aff203ce1a64736f6c63430008110033",
}

// GweiPumpABI is the input ABI used to generate the binding from.
// Deprecated: Use GweiPumpMetaData.ABI instead.
var GweiPumpABI = GweiPumpMetaData.ABI

// GweiPumpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GweiPumpMetaData.Bin instead.
var GweiPumpBin = GweiPumpMetaData.Bin

// DeployGweiPump deploys a new Ethereum contract, binding an instance of GweiPump to it.
func DeployGweiPump(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GweiPump, error) {
	parsed, err := GweiPumpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GweiPumpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GweiPump{GweiPumpCaller: GweiPumpCaller{contract: contract}, GweiPumpTransactor: GweiPumpTransactor{contract: contract}, GweiPumpFilterer: GweiPumpFilterer{contract: contract}}, nil
}

// GweiPump is an auto generated Go binding around an Ethereum contract.
type GweiPump struct {
	GweiPumpCaller     // Read-only binding to the contract
	GweiPumpTransactor // Write-only binding to the contract
	GweiPumpFilterer   // Log filterer for contract events
}

// GweiPumpCaller is an auto generated read-only Go binding around an Ethereum contract.
type GweiPumpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GweiPumpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GweiPumpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GweiPumpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GweiPumpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GweiPumpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GweiPumpSession struct {
	Contract     *GweiPump         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GweiPumpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GweiPumpCallerSession struct {
	Contract *GweiPumpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GweiPumpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GweiPumpTransactorSession struct {
	Contract     *GweiPumpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GweiPumpRaw is an auto generated low-level Go binding around an Ethereum contract.
type GweiPumpRaw struct {
	Contract *GweiPump // Generic contract binding to access the raw methods on
}

// GweiPumpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GweiPumpCallerRaw struct {
	Contract *GweiPumpCaller // Generic read-only contract binding to access the raw methods on
}

// GweiPumpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GweiPumpTransactorRaw struct {
	Contract *GweiPumpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGweiPump creates a new instance of GweiPump, bound to a specific deployed contract.
func NewGweiPump(address common.Address, backend bind.ContractBackend) (*GweiPump, error) {
	contract, err := bindGweiPump(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GweiPump{GweiPumpCaller: GweiPumpCaller{contract: contract}, GweiPumpTransactor: GweiPumpTransactor{contract: contract}, GweiPumpFilterer: GweiPumpFilterer{contract: contract}}, nil
}

// NewGweiPumpCaller creates a new read-only instance of GweiPump, bound to a specific deployed contract.
func NewGweiPumpCaller(address common.Address, caller bind.ContractCaller) (*GweiPumpCaller, error) {
	contract, err := bindGweiPump(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GweiPumpCaller{contract: contract}, nil
}

// NewGweiPumpTransactor creates a new write-only instance of GweiPump, bound to a specific deployed contract.
func NewGweiPumpTransactor(address common.Address, transactor bind.ContractTransactor) (*GweiPumpTransactor, error) {
	contract, err := bindGweiPump(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GweiPumpTransactor{contract: contract}, nil
}

// NewGweiPumpFilterer creates a new log filterer instance of GweiPump, bound to a specific deployed contract.
func NewGweiPumpFilterer(address common.Address, filterer bind.ContractFilterer) (*GweiPumpFilterer, error) {
	contract, err := bindGweiPump(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GweiPumpFilterer{contract: contract}, nil
}

// bindGweiPump binds a generic wrapper to an already deployed contract.
func bindGweiPump(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GweiPumpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GweiPump *GweiPumpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GweiPump.Contract.GweiPumpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GweiPump *GweiPumpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GweiPump.Contract.GweiPumpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GweiPump *GweiPumpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GweiPump.Contract.GweiPumpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GweiPump *GweiPumpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GweiPump.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GweiPump *GweiPumpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GweiPump.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GweiPump *GweiPumpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GweiPump.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_GweiPump *GweiPumpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_GweiPump *GweiPumpSession) Owner() (common.Address, error) {
	return _GweiPump.Contract.Owner(&_GweiPump.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_GweiPump *GweiPumpCallerSession) Owner() (common.Address, error) {
	return _GweiPump.Contract.Owner(&_GweiPump.CallOpts)
}

// Wti40Milliliters is a free data retrieval call binding the contract method 0x2514f7e3.
//
// Solidity: function Wti40Milliliters() view returns(uint256)
func (_GweiPump *GweiPumpCaller) Wti40Milliliters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "Wti40Milliliters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wti40Milliliters is a free data retrieval call binding the contract method 0x2514f7e3.
//
// Solidity: function Wti40Milliliters() view returns(uint256)
func (_GweiPump *GweiPumpSession) Wti40Milliliters() (*big.Int, error) {
	return _GweiPump.Contract.Wti40Milliliters(&_GweiPump.CallOpts)
}

// Wti40Milliliters is a free data retrieval call binding the contract method 0x2514f7e3.
//
// Solidity: function Wti40Milliliters() view returns(uint256)
func (_GweiPump *GweiPumpCallerSession) Wti40Milliliters() (*big.Int, error) {
	return _GweiPump.Contract.Wti40Milliliters(&_GweiPump.CallOpts)
}

// WtiPriceOracle is a free data retrieval call binding the contract method 0xd698a37e.
//
// Solidity: function WtiPriceOracle() view returns(uint256)
func (_GweiPump *GweiPumpCaller) WtiPriceOracle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "WtiPriceOracle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WtiPriceOracle is a free data retrieval call binding the contract method 0xd698a37e.
//
// Solidity: function WtiPriceOracle() view returns(uint256)
func (_GweiPump *GweiPumpSession) WtiPriceOracle() (*big.Int, error) {
	return _GweiPump.Contract.WtiPriceOracle(&_GweiPump.CallOpts)
}

// WtiPriceOracle is a free data retrieval call binding the contract method 0xd698a37e.
//
// Solidity: function WtiPriceOracle() view returns(uint256)
func (_GweiPump *GweiPumpCallerSession) WtiPriceOracle() (*big.Int, error) {
	return _GweiPump.Contract.WtiPriceOracle(&_GweiPump.CallOpts)
}

// Erc20LINK is a free data retrieval call binding the contract method 0xd6666612.
//
// Solidity: function erc20LINK() view returns(address)
func (_GweiPump *GweiPumpCaller) Erc20LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "erc20LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20LINK is a free data retrieval call binding the contract method 0xd6666612.
//
// Solidity: function erc20LINK() view returns(address)
func (_GweiPump *GweiPumpSession) Erc20LINK() (common.Address, error) {
	return _GweiPump.Contract.Erc20LINK(&_GweiPump.CallOpts)
}

// Erc20LINK is a free data retrieval call binding the contract method 0xd6666612.
//
// Solidity: function erc20LINK() view returns(address)
func (_GweiPump *GweiPumpCallerSession) Erc20LINK() (common.Address, error) {
	return _GweiPump.Contract.Erc20LINK(&_GweiPump.CallOpts)
}

// GetLatestMaticUsd is a free data retrieval call binding the contract method 0x784eb75b.
//
// Solidity: function getLatestMaticUsd() view returns(int256)
func (_GweiPump *GweiPumpCaller) GetLatestMaticUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "getLatestMaticUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestMaticUsd is a free data retrieval call binding the contract method 0x784eb75b.
//
// Solidity: function getLatestMaticUsd() view returns(int256)
func (_GweiPump *GweiPumpSession) GetLatestMaticUsd() (*big.Int, error) {
	return _GweiPump.Contract.GetLatestMaticUsd(&_GweiPump.CallOpts)
}

// GetLatestMaticUsd is a free data retrieval call binding the contract method 0x784eb75b.
//
// Solidity: function getLatestMaticUsd() view returns(int256)
func (_GweiPump *GweiPumpCallerSession) GetLatestMaticUsd() (*big.Int, error) {
	return _GweiPump.Contract.GetLatestMaticUsd(&_GweiPump.CallOpts)
}

// GetLatestWtiMatic is a free data retrieval call binding the contract method 0xc62d34ab.
//
// Solidity: function getLatestWtiMatic() view returns(uint256)
func (_GweiPump *GweiPumpCaller) GetLatestWtiMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "getLatestWtiMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestWtiMatic is a free data retrieval call binding the contract method 0xc62d34ab.
//
// Solidity: function getLatestWtiMatic() view returns(uint256)
func (_GweiPump *GweiPumpSession) GetLatestWtiMatic() (*big.Int, error) {
	return _GweiPump.Contract.GetLatestWtiMatic(&_GweiPump.CallOpts)
}

// GetLatestWtiMatic is a free data retrieval call binding the contract method 0xc62d34ab.
//
// Solidity: function getLatestWtiMatic() view returns(uint256)
func (_GweiPump *GweiPumpCallerSession) GetLatestWtiMatic() (*big.Int, error) {
	return _GweiPump.Contract.GetLatestWtiMatic(&_GweiPump.CallOpts)
}

// IsPumpFilled is a free data retrieval call binding the contract method 0x5fce9d34.
//
// Solidity: function isPumpFilled() view returns(uint256)
func (_GweiPump *GweiPumpCaller) IsPumpFilled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "isPumpFilled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsPumpFilled is a free data retrieval call binding the contract method 0x5fce9d34.
//
// Solidity: function isPumpFilled() view returns(uint256)
func (_GweiPump *GweiPumpSession) IsPumpFilled() (*big.Int, error) {
	return _GweiPump.Contract.IsPumpFilled(&_GweiPump.CallOpts)
}

// IsPumpFilled is a free data retrieval call binding the contract method 0x5fce9d34.
//
// Solidity: function isPumpFilled() view returns(uint256)
func (_GweiPump *GweiPumpCallerSession) IsPumpFilled() (*big.Int, error) {
	return _GweiPump.Contract.IsPumpFilled(&_GweiPump.CallOpts)
}

// LastWtiPriceCheckUnixTime is a free data retrieval call binding the contract method 0x861b31ce.
//
// Solidity: function lastWtiPriceCheckUnixTime() view returns(uint256)
func (_GweiPump *GweiPumpCaller) LastWtiPriceCheckUnixTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GweiPump.contract.Call(opts, &out, "lastWtiPriceCheckUnixTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastWtiPriceCheckUnixTime is a free data retrieval call binding the contract method 0x861b31ce.
//
// Solidity: function lastWtiPriceCheckUnixTime() view returns(uint256)
func (_GweiPump *GweiPumpSession) LastWtiPriceCheckUnixTime() (*big.Int, error) {
	return _GweiPump.Contract.LastWtiPriceCheckUnixTime(&_GweiPump.CallOpts)
}

// LastWtiPriceCheckUnixTime is a free data retrieval call binding the contract method 0x861b31ce.
//
// Solidity: function lastWtiPriceCheckUnixTime() view returns(uint256)
func (_GweiPump *GweiPumpCallerSession) LastWtiPriceCheckUnixTime() (*big.Int, error) {
	return _GweiPump.Contract.LastWtiPriceCheckUnixTime(&_GweiPump.CallOpts)
}

// BuyOil40Milliliters is a paid mutator transaction binding the contract method 0xf5da229e.
//
// Solidity: function buyOil40Milliliters() payable returns()
func (_GweiPump *GweiPumpTransactor) BuyOil40Milliliters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GweiPump.contract.Transact(opts, "buyOil40Milliliters")
}

// BuyOil40Milliliters is a paid mutator transaction binding the contract method 0xf5da229e.
//
// Solidity: function buyOil40Milliliters() payable returns()
func (_GweiPump *GweiPumpSession) BuyOil40Milliliters() (*types.Transaction, error) {
	return _GweiPump.Contract.BuyOil40Milliliters(&_GweiPump.TransactOpts)
}

// BuyOil40Milliliters is a paid mutator transaction binding the contract method 0xf5da229e.
//
// Solidity: function buyOil40Milliliters() payable returns()
func (_GweiPump *GweiPumpTransactorSession) BuyOil40Milliliters() (*types.Transaction, error) {
	return _GweiPump.Contract.BuyOil40Milliliters(&_GweiPump.TransactOpts)
}

// ChainlinkNodeRequestWtiPrice is a paid mutator transaction binding the contract method 0xbcc9f591.
//
// Solidity: function chainlinkNodeRequestWtiPrice() returns(bytes32 requestId)
func (_GweiPump *GweiPumpTransactor) ChainlinkNodeRequestWtiPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GweiPump.contract.Transact(opts, "chainlinkNodeRequestWtiPrice")
}

// ChainlinkNodeRequestWtiPrice is a paid mutator transaction binding the contract method 0xbcc9f591.
//
// Solidity: function chainlinkNodeRequestWtiPrice() returns(bytes32 requestId)
func (_GweiPump *GweiPumpSession) ChainlinkNodeRequestWtiPrice() (*types.Transaction, error) {
	return _GweiPump.Contract.ChainlinkNodeRequestWtiPrice(&_GweiPump.TransactOpts)
}

// ChainlinkNodeRequestWtiPrice is a paid mutator transaction binding the contract method 0xbcc9f591.
//
// Solidity: function chainlinkNodeRequestWtiPrice() returns(bytes32 requestId)
func (_GweiPump *GweiPumpTransactorSession) ChainlinkNodeRequestWtiPrice() (*types.Transaction, error) {
	return _GweiPump.Contract.ChainlinkNodeRequestWtiPrice(&_GweiPump.TransactOpts)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) returns(bool upkeepNeeded, bytes)
func (_GweiPump *GweiPumpTransactor) CheckUpkeep(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _GweiPump.contract.Transact(opts, "checkUpkeep", arg0)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) returns(bool upkeepNeeded, bytes)
func (_GweiPump *GweiPumpSession) CheckUpkeep(arg0 []byte) (*types.Transaction, error) {
	return _GweiPump.Contract.CheckUpkeep(&_GweiPump.TransactOpts, arg0)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) returns(bool upkeepNeeded, bytes)
func (_GweiPump *GweiPumpTransactorSession) CheckUpkeep(arg0 []byte) (*types.Transaction, error) {
	return _GweiPump.Contract.CheckUpkeep(&_GweiPump.TransactOpts, arg0)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 memoryWtiPriceOracle) returns()
func (_GweiPump *GweiPumpTransactor) Fulfill(opts *bind.TransactOpts, _requestId [32]byte, memoryWtiPriceOracle *big.Int) (*types.Transaction, error) {
	return _GweiPump.contract.Transact(opts, "fulfill", _requestId, memoryWtiPriceOracle)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 memoryWtiPriceOracle) returns()
func (_GweiPump *GweiPumpSession) Fulfill(_requestId [32]byte, memoryWtiPriceOracle *big.Int) (*types.Transaction, error) {
	return _GweiPump.Contract.Fulfill(&_GweiPump.TransactOpts, _requestId, memoryWtiPriceOracle)
}

// Fulfill is a paid mutator transaction binding the contract method 0x4357855e.
//
// Solidity: function fulfill(bytes32 _requestId, uint256 memoryWtiPriceOracle) returns()
func (_GweiPump *GweiPumpTransactorSession) Fulfill(_requestId [32]byte, memoryWtiPriceOracle *big.Int) (*types.Transaction, error) {
	return _GweiPump.Contract.Fulfill(&_GweiPump.TransactOpts, _requestId, memoryWtiPriceOracle)
}

// ManualUpKeep is a paid mutator transaction binding the contract method 0x7d5ded7c.
//
// Solidity: function manualUpKeep() returns()
func (_GweiPump *GweiPumpTransactor) ManualUpKeep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GweiPump.contract.Transact(opts, "manualUpKeep")
}

// ManualUpKeep is a paid mutator transaction binding the contract method 0x7d5ded7c.
//
// Solidity: function manualUpKeep() returns()
func (_GweiPump *GweiPumpSession) ManualUpKeep() (*types.Transaction, error) {
	return _GweiPump.Contract.ManualUpKeep(&_GweiPump.TransactOpts)
}

// ManualUpKeep is a paid mutator transaction binding the contract method 0x7d5ded7c.
//
// Solidity: function manualUpKeep() returns()
func (_GweiPump *GweiPumpTransactorSession) ManualUpKeep() (*types.Transaction, error) {
	return _GweiPump.Contract.ManualUpKeep(&_GweiPump.TransactOpts)
}

// OwnerPumpFilledStatus is a paid mutator transaction binding the contract method 0xd42506b8.
//
// Solidity: function ownerPumpFilledStatus(uint256 status) returns()
func (_GweiPump *GweiPumpTransactor) OwnerPumpFilledStatus(opts *bind.TransactOpts, status *big.Int) (*types.Transaction, error) {
	return _GweiPump.contract.Transact(opts, "ownerPumpFilledStatus", status)
}

// OwnerPumpFilledStatus is a paid mutator transaction binding the contract method 0xd42506b8.
//
// Solidity: function ownerPumpFilledStatus(uint256 status) returns()
func (_GweiPump *GweiPumpSession) OwnerPumpFilledStatus(status *big.Int) (*types.Transaction, error) {
	return _GweiPump.Contract.OwnerPumpFilledStatus(&_GweiPump.TransactOpts, status)
}

// OwnerPumpFilledStatus is a paid mutator transaction binding the contract method 0xd42506b8.
//
// Solidity: function ownerPumpFilledStatus(uint256 status) returns()
func (_GweiPump *GweiPumpTransactorSession) OwnerPumpFilledStatus(status *big.Int) (*types.Transaction, error) {
	return _GweiPump.Contract.OwnerPumpFilledStatus(&_GweiPump.TransactOpts, status)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes ) returns()
func (_GweiPump *GweiPumpTransactor) PerformUpkeep(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _GweiPump.contract.Transact(opts, "performUpkeep", arg0)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes ) returns()
func (_GweiPump *GweiPumpSession) PerformUpkeep(arg0 []byte) (*types.Transaction, error) {
	return _GweiPump.Contract.PerformUpkeep(&_GweiPump.TransactOpts, arg0)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes ) returns()
func (_GweiPump *GweiPumpTransactorSession) PerformUpkeep(arg0 []byte) (*types.Transaction, error) {
	return _GweiPump.Contract.PerformUpkeep(&_GweiPump.TransactOpts, arg0)
}

// GweiPumpChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the GweiPump contract.
type GweiPumpChainlinkCancelledIterator struct {
	Event *GweiPumpChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *GweiPumpChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GweiPumpChainlinkCancelled)
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
		it.Event = new(GweiPumpChainlinkCancelled)
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
func (it *GweiPumpChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GweiPumpChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GweiPumpChainlinkCancelled represents a ChainlinkCancelled event raised by the GweiPump contract.
type GweiPumpChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*GweiPumpChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GweiPump.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &GweiPumpChainlinkCancelledIterator{contract: _GweiPump.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *GweiPumpChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GweiPump.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GweiPumpChainlinkCancelled)
				if err := _GweiPump.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) ParseChainlinkCancelled(log types.Log) (*GweiPumpChainlinkCancelled, error) {
	event := new(GweiPumpChainlinkCancelled)
	if err := _GweiPump.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GweiPumpChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the GweiPump contract.
type GweiPumpChainlinkFulfilledIterator struct {
	Event *GweiPumpChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *GweiPumpChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GweiPumpChainlinkFulfilled)
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
		it.Event = new(GweiPumpChainlinkFulfilled)
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
func (it *GweiPumpChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GweiPumpChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GweiPumpChainlinkFulfilled represents a ChainlinkFulfilled event raised by the GweiPump contract.
type GweiPumpChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*GweiPumpChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GweiPump.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &GweiPumpChainlinkFulfilledIterator{contract: _GweiPump.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *GweiPumpChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GweiPump.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GweiPumpChainlinkFulfilled)
				if err := _GweiPump.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) ParseChainlinkFulfilled(log types.Log) (*GweiPumpChainlinkFulfilled, error) {
	event := new(GweiPumpChainlinkFulfilled)
	if err := _GweiPump.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GweiPumpChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the GweiPump contract.
type GweiPumpChainlinkRequestedIterator struct {
	Event *GweiPumpChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *GweiPumpChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GweiPumpChainlinkRequested)
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
		it.Event = new(GweiPumpChainlinkRequested)
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
func (it *GweiPumpChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GweiPumpChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GweiPumpChainlinkRequested represents a ChainlinkRequested event raised by the GweiPump contract.
type GweiPumpChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*GweiPumpChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GweiPump.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &GweiPumpChainlinkRequestedIterator{contract: _GweiPump.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *GweiPumpChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GweiPump.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GweiPumpChainlinkRequested)
				if err := _GweiPump.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_GweiPump *GweiPumpFilterer) ParseChainlinkRequested(log types.Log) (*GweiPumpChainlinkRequested, error) {
	event := new(GweiPumpChainlinkRequested)
	if err := _GweiPump.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GweiPumpOilBoughtIterator is returned from FilterOilBought and is used to iterate over the raw logs and unpacked data for OilBought events raised by the GweiPump contract.
type GweiPumpOilBoughtIterator struct {
	Event *GweiPumpOilBought // Event containing the contract specifics and raw log

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
func (it *GweiPumpOilBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GweiPumpOilBought)
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
		it.Event = new(GweiPumpOilBought)
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
func (it *GweiPumpOilBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GweiPumpOilBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GweiPumpOilBought represents a OilBought event raised by the GweiPump contract.
type GweiPumpOilBought struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOilBought is a free log retrieval operation binding the contract event 0x260abb0ce3f225ef1530c6f1a606b8331910fd3f66650c133f35b5817624ffe2.
//
// Solidity: event oilBought()
func (_GweiPump *GweiPumpFilterer) FilterOilBought(opts *bind.FilterOpts) (*GweiPumpOilBoughtIterator, error) {

	logs, sub, err := _GweiPump.contract.FilterLogs(opts, "oilBought")
	if err != nil {
		return nil, err
	}
	return &GweiPumpOilBoughtIterator{contract: _GweiPump.contract, event: "oilBought", logs: logs, sub: sub}, nil
}

// WatchOilBought is a free log subscription operation binding the contract event 0x260abb0ce3f225ef1530c6f1a606b8331910fd3f66650c133f35b5817624ffe2.
//
// Solidity: event oilBought()
func (_GweiPump *GweiPumpFilterer) WatchOilBought(opts *bind.WatchOpts, sink chan<- *GweiPumpOilBought) (event.Subscription, error) {

	logs, sub, err := _GweiPump.contract.WatchLogs(opts, "oilBought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GweiPumpOilBought)
				if err := _GweiPump.contract.UnpackLog(event, "oilBought", log); err != nil {
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

// ParseOilBought is a log parse operation binding the contract event 0x260abb0ce3f225ef1530c6f1a606b8331910fd3f66650c133f35b5817624ffe2.
//
// Solidity: event oilBought()
func (_GweiPump *GweiPumpFilterer) ParseOilBought(log types.Log) (*GweiPumpOilBought, error) {
	event := new(GweiPumpOilBought)
	if err := _GweiPump.contract.UnpackLog(event, "oilBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GweiPumpUpdateWtiIterator is returned from FilterUpdateWti and is used to iterate over the raw logs and unpacked data for UpdateWti events raised by the GweiPump contract.
type GweiPumpUpdateWtiIterator struct {
	Event *GweiPumpUpdateWti // Event containing the contract specifics and raw log

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
func (it *GweiPumpUpdateWtiIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GweiPumpUpdateWti)
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
		it.Event = new(GweiPumpUpdateWti)
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
func (it *GweiPumpUpdateWtiIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GweiPumpUpdateWtiIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GweiPumpUpdateWti represents a UpdateWti event raised by the GweiPump contract.
type GweiPumpUpdateWti struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateWti is a free log retrieval operation binding the contract event 0x90c6eb475b1dc27758b0b1c75bd44d59a523bab636839816edfad763432d1e5a.
//
// Solidity: event updateWti()
func (_GweiPump *GweiPumpFilterer) FilterUpdateWti(opts *bind.FilterOpts) (*GweiPumpUpdateWtiIterator, error) {

	logs, sub, err := _GweiPump.contract.FilterLogs(opts, "updateWti")
	if err != nil {
		return nil, err
	}
	return &GweiPumpUpdateWtiIterator{contract: _GweiPump.contract, event: "updateWti", logs: logs, sub: sub}, nil
}

// WatchUpdateWti is a free log subscription operation binding the contract event 0x90c6eb475b1dc27758b0b1c75bd44d59a523bab636839816edfad763432d1e5a.
//
// Solidity: event updateWti()
func (_GweiPump *GweiPumpFilterer) WatchUpdateWti(opts *bind.WatchOpts, sink chan<- *GweiPumpUpdateWti) (event.Subscription, error) {

	logs, sub, err := _GweiPump.contract.WatchLogs(opts, "updateWti")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GweiPumpUpdateWti)
				if err := _GweiPump.contract.UnpackLog(event, "updateWti", log); err != nil {
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

// ParseUpdateWti is a log parse operation binding the contract event 0x90c6eb475b1dc27758b0b1c75bd44d59a523bab636839816edfad763432d1e5a.
//
// Solidity: event updateWti()
func (_GweiPump *GweiPumpFilterer) ParseUpdateWti(log types.Log) (*GweiPumpUpdateWti, error) {
	event := new(GweiPumpUpdateWti)
	if err := _GweiPump.contract.UnpackLog(event, "updateWti", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
