// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"PlaceBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061224a806100606000396000f3fe6080604052600436106100705760003560e01c8063b4a99a4e1161004e578063b4a99a4e146100d3578063d612429e146100fe578063e63f341f14610127578063ed21248c1461016457610070565b80630ef6788714610075578063474ad2f1146100a057806357ea89b6146100c9575b600080fd5b34801561008157600080fd5b5061008a61016e565b60405161009791906115cd565b60405180910390f35b3480156100ac57600080fd5b506100c760048036038101906100c29190611677565b6101ff565b005b6100d16105de565b005b3480156100df57600080fd5b506100e861087e565b6040516100f591906116d9565b60405180910390f35b34801561010a57600080fd5b50610125600480360381019061012091906116f4565b6108a2565b005b34801561013357600080fd5b5061014e6004803603810190610149919061175b565b611008565b60405161015b9190611833565b60405180910390f35b61016c611102565b005b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546101fa919061187d565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461025757600080fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561039e57600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461034d91906118b1565b925050819055506000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061046c565b80600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461040e91906118b1565b9250508190555080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610464919061187d565b925050819055505b81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156104ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e590611942565b60405180910390fd5b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61055c83611201565b6105a4600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611201565b6040516020016105b5929190611a45565b6040516020818303038152906040526040516105d19190611ae0565b60405180910390a1505050565b6000339050600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411156106a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069b90611b4e565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610730919061187d565b905060008103610775576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076c90611b4e565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156107bb573d6000803e3d6000fd5b5080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461080b919061187d565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61083c33611389565b61084583611201565b604051602001610856929190611bba565b6040516020818303038152906040526040516108729190611ae0565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108fa57600080fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610a76577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a60405161096e90611c7d565b60405180910390a1600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610a2591906118b1565b925050819055506000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610b44565b80600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610ae691906118b1565b9250508190555080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b3c919061187d565b925050819055505b81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610c9f577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a604051610bb890611d0f565b60405180910390a1600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c4e91906118b1565b925050819055506000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d4c565b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610cee91906118b1565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d44919061187d565b925050819055505b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610ddd576000600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610e34565b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e2c919061187d565b925050819055505b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610ec5576000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610f1c565b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f14919061187d565b925050819055505b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610f85600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611201565b610fcd600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611201565b604051602001610fde929190611d7b565b604051602081830303815290604052604051610ffa9190611ae0565b60405180910390a150505050565b611010611592565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461106857600080fd5b6040518060400160405280600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152509050919050565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461115191906118b1565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61118233611389565b6111ca600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611201565b6040516020016111db929190611e18565b6040516020818303038152906040526040516111f79190611ae0565b60405180910390a1565b606060008203611248576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611384565b600082905060005b6000821461127a57808061126390611e69565b915050600a826112739190611ee0565b9150611250565b60008167ffffffffffffffff81111561129657611295611f11565b5b6040519080825280601f01601f1916602001820160405280156112c85781602001600182028036833780820191505090505b50905060008290505b6000861461137c576001816112e6919061187d565b90506000600a80886112f89190611ee0565b6113029190611f40565b8761130d919061187d565b60306113199190611f8f565b905060008160f81b90508084848151811061133757611336611fc4565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886113739190611ee0565b975050506112d1565b819450505050505b919050565b60606000602867ffffffffffffffff8111156113a8576113a7611f11565b5b6040519080825280601f01601f1916602001820160405280156113da5781602001600182028036833780820191505090505b50905060005b60148110156115425760008160136113f8919061187d565b60086114049190611f40565b60026114109190612126565b8573ffffffffffffffffffffffffffffffffffffffff166114319190611ee0565b60f81b9050600060108260f81c6114489190612171565b60f81b905060008160f81c601061145f91906121a2565b8360f81c61146d91906121df565b60f81b905061147b8261154c565b858560026114899190611f40565b8151811061149a57611499611fc4565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506114d28161154c565b8560018660026114e29190611f40565b6114ec91906118b1565b815181106114fd576114fc611fc4565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350505050808061153a90611e69565b9150506113e0565b5080915050919050565b6000600a8260f81c60ff1610156115775760308260f81c61156d9190611f8f565b60f81b905061158d565b60578260f81c6115879190611f8f565b60f81b90505b919050565b6040518060400160405280600290602082028036833780820191505090505090565b6000819050919050565b6115c7816115b4565b82525050565b60006020820190506115e260008301846115be565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611618826115ed565b9050919050565b6116288161160d565b811461163357600080fd5b50565b6000813590506116458161161f565b92915050565b611654816115b4565b811461165f57600080fd5b50565b6000813590506116718161164b565b92915050565b6000806000606084860312156116905761168f6115e8565b5b600061169e86828701611636565b93505060206116af86828701611662565b92505060406116c086828701611662565b9150509250925092565b6116d38161160d565b82525050565b60006020820190506116ee60008301846116ca565b92915050565b6000806000806080858703121561170e5761170d6115e8565b5b600061171c87828801611636565b945050602061172d87828801611636565b935050604061173e87828801611662565b925050606061174f87828801611662565b91505092959194509250565b600060208284031215611771576117706115e8565b5b600061177f84828501611636565b91505092915050565b600060029050919050565b600081905092915050565b6000819050919050565b6117b1816115b4565b82525050565b60006117c383836117a8565b60208301905092915050565b6000602082019050919050565b6117e581611788565b6117ef8184611793565b92506117fa8261179e565b8060005b8381101561182b57815161181287826117b7565b965061181d836117cf565b9250506001810190506117fe565b505050505050565b600060408201905061184860008301846117dc565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611888826115b4565b9150611893836115b4565b92508282039050818111156118ab576118aa61184e565b5b92915050565b60006118bc826115b4565b91506118c7836115b4565b92508282019050808211156118df576118de61184e565b5b92915050565b600082825260208201905092915050565b7f6163636f756e742062616c616e636520746f6f206c6f77000000000000000000600082015250565b600061192c6017836118e5565b9150611937826118f6565b602082019050919050565b6000602082019050818103600083015261195b8161191f565b9050919050565b7f6265745b00000000000000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b838110156119bc5780820151818401526020810190506119a1565b60008484015250505050565b60006119d382611988565b6119dd8185611993565b93506119ed81856020860161199e565b80840191505092915050565b7f5d20746f74616c5b000000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000611a5082611962565b600482019150611a6082856119c8565b9150611a6b826119f9565b600882019150611a7b82846119c8565b9150611a8682611a1f565b6001820191508190509392505050565b6000601f19601f8301169050919050565b6000611ab282611988565b611abc81856118e5565b9350611acc81856020860161199e565b611ad581611a96565b840191505092915050565b60006020820190508181036000830152611afa8184611aa7565b905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000611b386012836118e5565b9150611b4382611b02565b602082019050919050565b60006020820190508181036000830152611b6781611b2b565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b6000611bc582611b6e565b600982019150611bd582856119c8565b9150611be082611b94565b600982019150611bf082846119c8565b9150611bfb82611a1f565b6001820191508190509392505050565b7f6c6f7365722062616c616e636520746f6f206c6f772c2074616b696e6720726560008201527f6d61696e64657220617320666565000000000000000000000000000000000000602082015250565b6000611c67602e836118e5565b9150611c7282611c0b565b604082019050919050565b60006020820190508181036000830152611c9681611c5a565b9050919050565b7f6c6f7365722062616c616e636520746f6f206c6f772c206d6f76696e6720726560008201527f6d61696e64657220746f2077696e6e6572206163636f756e7400000000000000602082015250565b6000611cf96039836118e5565b9150611d0482611c9d565b604082019050919050565b60006020820190508181036000830152611d2881611cec565b9050919050565b7f77696e6e65722062616c616e63655b0000000000000000000000000000000000815250565b7f5d206c6f7365722062616c616e63655b00000000000000000000000000000000815250565b6000611d8682611d2f565b600f82019150611d9682856119c8565b9150611da182611d55565b601082019150611db182846119c8565b9150611dbc82611a1f565b6001820191508190509392505050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b6000611e2382611dcc565b600882019150611e3382856119c8565b9150611e3e82611df2565b600a82019150611e4e82846119c8565b9150611e5982611a1f565b6001820191508190509392505050565b6000611e74826115b4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611ea657611ea561184e565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611eeb826115b4565b9150611ef6836115b4565b925082611f0657611f05611eb1565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000611f4b826115b4565b9150611f56836115b4565b9250828202611f64816115b4565b91508282048414831517611f7b57611f7a61184e565b5b5092915050565b600060ff82169050919050565b6000611f9a82611f82565b9150611fa583611f82565b9250828201905060ff811115611fbe57611fbd61184e565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008160011c9050919050565b6000808291508390505b600185111561204a578086048111156120265761202561184e565b5b60018516156120355780820291505b808102905061204385611ff3565b945061200a565b94509492505050565b600082612063576001905061211f565b81612071576000905061211f565b81600181146120875760028114612091576120c0565b600191505061211f565b60ff8411156120a3576120a261184e565b5b8360020a9150848211156120ba576120b961184e565b5b5061211f565b5060208310610133831016604e8410600b84101617156120f55782820a9050838111156120f0576120ef61184e565b5b61211f565b6121028484846001612000565b925090508184048111156121195761211861184e565b5b81810290505b9392505050565b6000612131826115b4565b915061213c836115b4565b92506121697fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484612053565b905092915050565b600061217c82611f82565b915061218783611f82565b92508261219757612196611eb1565b5b828204905092915050565b60006121ad82611f82565b91506121b883611f82565b92508282026121c681611f82565b91508082146121d8576121d761184e565b5b5092915050565b60006121ea82611f82565b91506121f583611f82565b9250828203905060ff81111561220e5761220d61184e565b5b9291505056fea2646970667358221220865e5987f176bd9662864f7e47adbb9f09b9bc18d470c092b202e859dd743c4c64736f6c63430008110033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256[2])
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256[2])
func (_Bank *BankSession) AccountBalance(account common.Address) ([2]*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256[2])
func (_Bank *BankCallerSession) AccountBalance(account common.Address) ([2]*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCallerSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCallerSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x474ad2f1.
//
// Solidity: function PlaceBet(address person, uint256 amount, uint256 fee) returns()
func (_Bank *BankTransactor) PlaceBet(opts *bind.TransactOpts, person common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "PlaceBet", person, amount, fee)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x474ad2f1.
//
// Solidity: function PlaceBet(address person, uint256 amount, uint256 fee) returns()
func (_Bank *BankSession) PlaceBet(person common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.PlaceBet(&_Bank.TransactOpts, person, amount, fee)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x474ad2f1.
//
// Solidity: function PlaceBet(address person, uint256 amount, uint256 fee) returns()
func (_Bank *BankTransactorSession) PlaceBet(person common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.PlaceBet(&_Bank.TransactOpts, person, amount, fee)
}

// Reconcile is a paid mutator transaction binding the contract method 0xd612429e.
//
// Solidity: function Reconcile(address winner, address loser, uint256 amount, uint256 fee) returns()
func (_Bank *BankTransactor) Reconcile(opts *bind.TransactOpts, winner common.Address, loser common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Reconcile", winner, loser, amount, fee)
}

// Reconcile is a paid mutator transaction binding the contract method 0xd612429e.
//
// Solidity: function Reconcile(address winner, address loser, uint256 amount, uint256 fee) returns()
func (_Bank *BankSession) Reconcile(winner common.Address, loser common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, loser, amount, fee)
}

// Reconcile is a paid mutator transaction binding the contract method 0xd612429e.
//
// Solidity: function Reconcile(address winner, address loser, uint256 amount, uint256 fee) returns()
func (_Bank *BankTransactorSession) Reconcile(winner common.Address, loser common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Reconcile(&_Bank.TransactOpts, winner, loser, amount, fee)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// BankEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bank contract.
type BankEventLogIterator struct {
	Event *BankEventLog // Event containing the contract specifics and raw log

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
func (it *BankEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventLog)
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
		it.Event = new(BankEventLog)
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
func (it *BankEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventLog represents a EventLog event raised by the Bank contract.
type BankEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankEventLogIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankEventLogIterator{contract: _Bank.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankEventLog) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventLog)
				if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
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

// ParseEventLog is a log parse operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) ParseEventLog(log types.Log) (*BankEventLog, error) {
	event := new(BankEventLog)
	if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
