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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"Drain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"PlaceBet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Reconcile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506123d8806100606000396000f3fe60806040526004361061007b5760003560e01c8063b4a99a4e1161004e578063b4a99a4e146100fa578063d612429e14610125578063e63f341f1461014e578063ed21248c1461018b5761007b565b80630ef6788714610080578063474ad2f1146100ab57806357ea89b6146100d457806393f28237146100de575b600080fd5b34801561008c57600080fd5b50610095610195565b6040516100a29190611700565b60405180910390f35b3480156100b757600080fd5b506100d260048036038101906100cd91906117aa565b610226565b005b6100dc610605565b005b6100f860048036038101906100f391906117fd565b6108a5565b005b34801561010657600080fd5b5061010f6109b1565b60405161011c9190611839565b60405180910390f35b34801561013157600080fd5b5061014c60048036038101906101479190611854565b6109d5565b005b34801561015a57600080fd5b50610175600480360381019061017091906117fd565b61113b565b6040516101829190611966565b60405180910390f35b610193611235565b005b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461022191906119b0565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461027e57600080fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156103c557600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461037491906119e4565b925050819055506000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610493565b80600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461043591906119e4565b9250508190555080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461048b91906119b0565b925050819055505b81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610515576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050c90611a75565b60405180910390fd5b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61058383611334565b6105cb600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611334565b6040516020016105dc929190611b78565b6040516020818303038152906040526040516105f89190611c13565b60405180910390a1505050565b6000339050600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411156106cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c290611c81565b60405180910390fd5b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461075791906119b0565b90506000810361079c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079390611c81565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156107e2573d6000803e3d6000fd5b5080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461083291906119b0565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610863336114bc565b61086c83611334565b60405160200161087d929190611ced565b6040516020818303038152906040526040516108999190611c13565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108fd57600080fd5b600047905060008290508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505015801561094d573d6000803e3d6000fd5b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61097883611334565b6040516020016109889190611d64565b6040516020818303038152906040526040516109a49190611c13565b60405180910390a1505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a2d57600080fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610ba9577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a604051610aa190611e0b565b60405180910390a1600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b5891906119e4565b925050819055506000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610c77565b80600160008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c1991906119e4565b9250508190555080600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610c6f91906119b0565b925050819055505b81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610dd2577fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a604051610ceb90611e9d565b60405180910390a1600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d8191906119e4565b925050819055506000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610e7f565b81600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e2191906119e4565b9250508190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e7791906119b0565b925050819055505b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610f10576000600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610f67565b81600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f5f91906119b0565b925050819055505b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610ff8576000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061104f565b81600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461104791906119b0565b925050819055505b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6110b8600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611334565b611100600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611334565b604051602001611111929190611f09565b60405160208183030381529060405260405161112d9190611c13565b60405180910390a150505050565b6111436116c5565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461119b57600080fd5b6040518060400160405280600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152602001600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548152509050919050565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461128491906119e4565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6112b5336114bc565b6112fd600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611334565b60405160200161130e929190611fa6565b60405160208183030381529060405260405161132a9190611c13565b60405180910390a1565b60606000820361137b576040518060400160405280600181526020017f300000000000000000000000000000000000000000000000000000000000000081525090506114b7565b600082905060005b600082146113ad57808061139690611ff7565b915050600a826113a6919061206e565b9150611383565b60008167ffffffffffffffff8111156113c9576113c861209f565b5b6040519080825280601f01601f1916602001820160405280156113fb5781602001600182028036833780820191505090505b50905060008290505b600086146114af5760018161141991906119b0565b90506000600a808861142b919061206e565b61143591906120ce565b8761144091906119b0565b603061144c919061211d565b905060008160f81b90508084848151811061146a57611469612152565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a886114a6919061206e565b97505050611404565b819450505050505b919050565b60606000602867ffffffffffffffff8111156114db576114da61209f565b5b6040519080825280601f01601f19166020018201604052801561150d5781602001600182028036833780820191505090505b50905060005b601481101561167557600081601361152b91906119b0565b600861153791906120ce565b600261154391906122b4565b8573ffffffffffffffffffffffffffffffffffffffff16611564919061206e565b60f81b9050600060108260f81c61157b91906122ff565b60f81b905060008160f81c60106115929190612330565b8360f81c6115a0919061236d565b60f81b90506115ae8261167f565b858560026115bc91906120ce565b815181106115cd576115cc612152565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506116058161167f565b85600186600261161591906120ce565b61161f91906119e4565b815181106116305761162f612152565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350505050808061166d90611ff7565b915050611513565b5080915050919050565b6000600a8260f81c60ff1610156116aa5760308260f81c6116a0919061211d565b60f81b90506116c0565b60578260f81c6116ba919061211d565b60f81b90505b919050565b6040518060400160405280600290602082028036833780820191505090505090565b6000819050919050565b6116fa816116e7565b82525050565b600060208201905061171560008301846116f1565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061174b82611720565b9050919050565b61175b81611740565b811461176657600080fd5b50565b60008135905061177881611752565b92915050565b611787816116e7565b811461179257600080fd5b50565b6000813590506117a48161177e565b92915050565b6000806000606084860312156117c3576117c261171b565b5b60006117d186828701611769565b93505060206117e286828701611795565b92505060406117f386828701611795565b9150509250925092565b6000602082840312156118135761181261171b565b5b600061182184828501611769565b91505092915050565b61183381611740565b82525050565b600060208201905061184e600083018461182a565b92915050565b6000806000806080858703121561186e5761186d61171b565b5b600061187c87828801611769565b945050602061188d87828801611769565b935050604061189e87828801611795565b92505060606118af87828801611795565b91505092959194509250565b600060029050919050565b600081905092915050565b6000819050919050565b6118e4816116e7565b82525050565b60006118f683836118db565b60208301905092915050565b6000602082019050919050565b611918816118bb565b61192281846118c6565b925061192d826118d1565b8060005b8381101561195e57815161194587826118ea565b965061195083611902565b925050600181019050611931565b505050505050565b600060408201905061197b600083018461190f565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006119bb826116e7565b91506119c6836116e7565b92508282039050818111156119de576119dd611981565b5b92915050565b60006119ef826116e7565b91506119fa836116e7565b9250828201905080821115611a1257611a11611981565b5b92915050565b600082825260208201905092915050565b7f6163636f756e742062616c616e636520746f6f206c6f77000000000000000000600082015250565b6000611a5f601783611a18565b9150611a6a82611a29565b602082019050919050565b60006020820190508181036000830152611a8e81611a52565b9050919050565b7f6265745b00000000000000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015611aef578082015181840152602081019050611ad4565b60008484015250505050565b6000611b0682611abb565b611b108185611ac6565b9350611b20818560208601611ad1565b80840191505092915050565b7f5d20746f74616c5b000000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b6000611b8382611a95565b600482019150611b938285611afb565b9150611b9e82611b2c565b600882019150611bae8284611afb565b9150611bb982611b52565b6001820191508190509392505050565b6000601f19601f8301169050919050565b6000611be582611abb565b611bef8185611a18565b9350611bff818560208601611ad1565b611c0881611bc9565b840191505092915050565b60006020820190508181036000830152611c2d8184611bda565b905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b6000611c6b601283611a18565b9150611c7682611c35565b602082019050919050565b60006020820190508181036000830152611c9a81611c5e565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b6000611cf882611ca1565b600982019150611d088285611afb565b9150611d1382611cc7565b600982019150611d238284611afb565b9150611d2e82611b52565b6001820191508190509392505050565b7f7472616e736665725b0000000000000000000000000000000000000000000000815250565b6000611d6f82611d3e565b600982019150611d7f8284611afb565b9150611d8a82611b52565b60018201915081905092915050565b7f6c6f7365722062616c616e636520746f6f206c6f772c2074616b696e6720726560008201527f6d61696e64657220617320666565000000000000000000000000000000000000602082015250565b6000611df5602e83611a18565b9150611e0082611d99565b604082019050919050565b60006020820190508181036000830152611e2481611de8565b9050919050565b7f6c6f7365722062616c616e636520746f6f206c6f772c206d6f76696e6720726560008201527f6d61696e64657220746f2077696e6e6572206163636f756e7400000000000000602082015250565b6000611e87603983611a18565b9150611e9282611e2b565b604082019050919050565b60006020820190508181036000830152611eb681611e7a565b9050919050565b7f77696e6e65722062616c616e63655b0000000000000000000000000000000000815250565b7f5d206c6f7365722062616c616e63655b00000000000000000000000000000000815250565b6000611f1482611ebd565b600f82019150611f248285611afb565b9150611f2f82611ee3565b601082019150611f3f8284611afb565b9150611f4a82611b52565b6001820191508190509392505050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b6000611fb182611f5a565b600882019150611fc18285611afb565b9150611fcc82611f80565b600a82019150611fdc8284611afb565b9150611fe782611b52565b6001820191508190509392505050565b6000612002826116e7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361203457612033611981565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000612079826116e7565b9150612084836116e7565b9250826120945761209361203f565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006120d9826116e7565b91506120e4836116e7565b92508282026120f2816116e7565b9150828204841483151761210957612108611981565b5b5092915050565b600060ff82169050919050565b600061212882612110565b915061213383612110565b9250828201905060ff81111561214c5761214b611981565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156121d8578086048111156121b4576121b3611981565b5b60018516156121c35780820291505b80810290506121d185612181565b9450612198565b94509492505050565b6000826121f157600190506122ad565b816121ff57600090506122ad565b8160018114612215576002811461221f5761224e565b60019150506122ad565b60ff84111561223157612230611981565b5b8360020a91508482111561224857612247611981565b5b506122ad565b5060208310610133831016604e8410600b84101617156122835782820a90508381111561227e5761227d611981565b5b6122ad565b612290848484600161218e565b925090508184048111156122a7576122a6611981565b5b81810290505b9392505050565b60006122bf826116e7565b91506122ca836116e7565b92506122f77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846121e1565b905092915050565b600061230a82612110565b915061231583612110565b9250826123255761232461203f565b5b828204905092915050565b600061233b82612110565b915061234683612110565b925082820261235481612110565b915080821461236657612365611981565b5b5092915050565b600061237882612110565b915061238383612110565b9250828203905060ff81111561239c5761239b611981565b5b9291505056fea26469706673582212205c259889dd171e430f8b37c4defc795c27047e427d95da30dbc5ad0db0aa9c5764736f6c63430008110033",
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

// Drain is a paid mutator transaction binding the contract method 0x93f28237.
//
// Solidity: function Drain(address target) payable returns()
func (_Bank *BankTransactor) Drain(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Drain", target)
}

// Drain is a paid mutator transaction binding the contract method 0x93f28237.
//
// Solidity: function Drain(address target) payable returns()
func (_Bank *BankSession) Drain(target common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Drain(&_Bank.TransactOpts, target)
}

// Drain is a paid mutator transaction binding the contract method 0x93f28237.
//
// Solidity: function Drain(address target) payable returns()
func (_Bank *BankTransactorSession) Drain(target common.Address) (*types.Transaction, error) {
	return _Bank.Contract.Drain(&_Bank.TransactOpts, target)
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
