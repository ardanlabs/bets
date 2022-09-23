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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"Drain\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"betId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"nonce\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"PlaceBetsSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"betId\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"winners\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"moderator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ReconcileSigned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061244f806100606000396000f3fe60806040526004361061007b5760003560e01c8063b4a99a4e1161004e578063b4a99a4e146100fa578063bb9c53b914610125578063e63f341f1461014e578063ed21248c1461018b5761007b565b80630ef678871461008057806357ea89b6146100ab57806382156760146100b557806393f28237146100de575b600080fd5b34801561008c57600080fd5b50610095610195565b6040516100a29190611004565b60405180910390f35b6100b36101dc565b005b3480156100c157600080fd5b506100dc60048036038101906100d7919061133a565b610371565b005b6100f860048036038101906100f39190611414565b61060d565b005b34801561010657600080fd5b5061010f610719565b60405161011c9190611450565b60405180910390f35b34801561013157600080fd5b5061014c600480360381019061014791906116b4565b61073d565b005b34801561015a57600080fd5b5061017560048036038101906101709190611414565b610a04565b6040516101829190611004565b60405180910390f35b610193610aa6565b005b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008103610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025f90611853565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102ae573d6000803e3d6000fd5b5080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102fe91906118a2565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61032f33610ba5565b61033883610d68565b6040516020016103499291906119b9565b6040516020818303038152906040526040516103659190611a43565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103c957600080fd5b60006103d788888888610ef0565b90506000600182868686604051600081526020016040526040516103fe9493929190611a83565b6020604051602081039080840390855afa158015610420573d6000803e3d6000fd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff161461049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049190611b14565b60405180910390fd5b6000885160028b6040516104ae9190611b34565b9081526020016040518091039020546104c79190611b7a565b905060005b89518110156105dc5781600160008c84815181106104ed576104ec611bab565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461053e9190611bda565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a8b61058a8c848151811061057d5761057c611bab565b5b6020026020010151610ba5565b61059385610d68565b6040516020016105a593929190611c80565b6040516020818303038152906040526040516105c19190611a43565b60405180910390a180806105d490611ced565b9150506104cc565b50600060028b6040516105ef9190611b34565b90815260200160405180910390208190555050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066557600080fd5b600047905060008290508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156106b5573d6000803e3d6000fd5b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106e083610d68565b6040516020016106f09190611d5b565b60405160208183030381529060405260405161070c9190611a43565b60405180910390a1505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461079557600080fd5b60005b84518110156109f95760006107ca8a8a8a8a8a87815181106107bd576107bc611bab565b5b6020026020010151610f31565b905060006001828785815181106107e4576107e3611bab565b5b60200260200101518786815181106107ff576107fe611bab565b5b602002602001015187878151811061081a57610819611bab565b5b60200260200101516040516000815260200160405260405161083f9493929190611a83565b6020604051602081039080840390855afa158015610861573d6000803e3d6000fd5b50505060206040510351905088600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156108ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e690611ddc565b60405180910390fd5b8860028c6040516109009190611b34565b9081526020016040518091039020600082825461091d9190611bda565b9250508190555088600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461097391906118a2565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a8b6109a583610ba5565b6109ae8c610d68565b6040516020016109c093929190611e22565b6040516020818303038152906040526040516109dc9190611a43565b60405180910390a1505080806109f190611ced565b915050610798565b505050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a5f57600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610af59190611bda565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610b2633610ba5565b610b6e600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d68565b604051602001610b7f929190611edb565b604051602081830303815290604052604051610b9b9190611a43565b60405180910390a1565b60606000602867ffffffffffffffff811115610bc457610bc361104e565b5b6040519080825280601f01601f191660200182016040528015610bf65781602001600182028036833780820191505090505b50905060005b6014811015610d5e576000816013610c1491906118a2565b6008610c209190611f2c565b6002610c2c91906120a1565b8573ffffffffffffffffffffffffffffffffffffffff16610c4d9190611b7a565b60f81b9050600060108260f81c610c6491906120ec565b60f81b905060008160f81c6010610c7b919061211d565b8360f81c610c89919061215a565b60f81b9050610c9782610f75565b85856002610ca59190611f2c565b81518110610cb657610cb5611bab565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610cee81610f75565b856001866002610cfe9190611f2c565b610d089190611bda565b81518110610d1957610d18611bab565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610d5690611ced565b915050610bfc565b5080915050919050565b606060008203610daf576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610eeb565b600082905060005b60008214610de1578080610dca90611ced565b915050600a82610dda9190611b7a565b9150610db7565b60008167ffffffffffffffff811115610dfd57610dfc61104e565b5b6040519080825280601f01601f191660200182016040528015610e2f5781602001600182028036833780820191505090505b50905060008290505b60008614610ee357600181610e4d91906118a2565b90506000600a8088610e5f9190611b7a565b610e699190611f2c565b87610e7491906118a2565b6030610e80919061218f565b905060008160f81b905080848481518110610e9e57610e9d611bab565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610eda9190611b7a565b97505050610e38565b819450505050505b919050565b6000610f2785858585604051602001610f0c94939291906122e5565b60405160208183030381529060405280519060200120610fbb565b9050949350505050565b6000610f6a8686868686604051602001610f4f95949392919061232b565b60405160208183030381529060405280519060200120610fbb565b905095945050505050565b6000600a8260f81c60ff161015610fa05760308260f81c610f96919061218f565b60f81b9050610fb6565b60578260f81c610fb0919061218f565b60f81b90505b919050565b600081604051602001610fce91906123f3565b604051602081830303815290604052805190602001209050919050565b6000819050919050565b610ffe81610feb565b82525050565b60006020820190506110196000830184610ff5565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110868261103d565b810181811067ffffffffffffffff821117156110a5576110a461104e565b5b80604052505050565b60006110b861101f565b90506110c4828261107d565b919050565b600067ffffffffffffffff8211156110e4576110e361104e565b5b6110ed8261103d565b9050602081019050919050565b82818337600083830152505050565b600061111c611117846110c9565b6110ae565b90508281526020810184848401111561113857611137611038565b5b6111438482856110fa565b509392505050565b600082601f8301126111605761115f611033565b5b8135611170848260208601611109565b91505092915050565b600067ffffffffffffffff8211156111945761119361104e565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006111d5826111aa565b9050919050565b6111e5816111ca565b81146111f057600080fd5b50565b600081359050611202816111dc565b92915050565b600061121b61121684611179565b6110ae565b9050808382526020820190506020840283018581111561123e5761123d6111a5565b5b835b81811015611267578061125388826111f3565b845260208401935050602081019050611240565b5050509392505050565b600082601f83011261128657611285611033565b5b8135611296848260208601611208565b91505092915050565b6112a881610feb565b81146112b357600080fd5b50565b6000813590506112c58161129f565b92915050565b600060ff82169050919050565b6112e1816112cb565b81146112ec57600080fd5b50565b6000813590506112fe816112d8565b92915050565b6000819050919050565b61131781611304565b811461132257600080fd5b50565b6000813590506113348161130e565b92915050565b600080600080600080600060e0888a03121561135957611358611029565b5b600088013567ffffffffffffffff8111156113775761137661102e565b5b6113838a828b0161114b565b975050602088013567ffffffffffffffff8111156113a4576113a361102e565b5b6113b08a828b01611271565b96505060406113c18a828b016111f3565b95505060606113d28a828b016112b6565b94505060806113e38a828b016112ef565b93505060a06113f48a828b01611325565b92505060c06114058a828b01611325565b91505092959891949750929550565b60006020828403121561142a57611429611029565b5b6000611438848285016111f3565b91505092915050565b61144a816111ca565b82525050565b60006020820190506114656000830184611441565b92915050565b600067ffffffffffffffff8211156114865761148561104e565b5b602082029050602081019050919050565b60006114aa6114a58461146b565b6110ae565b905080838252602082019050602084028301858111156114cd576114cc6111a5565b5b835b818110156114f657806114e288826112b6565b8452602084019350506020810190506114cf565b5050509392505050565b600082601f83011261151557611514611033565b5b8135611525848260208601611497565b91505092915050565b600067ffffffffffffffff8211156115495761154861104e565b5b602082029050602081019050919050565b600061156d6115688461152e565b6110ae565b905080838252602082019050602084028301858111156115905761158f6111a5565b5b835b818110156115b957806115a588826112ef565b845260208401935050602081019050611592565b5050509392505050565b600082601f8301126115d8576115d7611033565b5b81356115e884826020860161155a565b91505092915050565b600067ffffffffffffffff82111561160c5761160b61104e565b5b602082029050602081019050919050565b600061163061162b846115f1565b6110ae565b90508083825260208201905060208402830185811115611653576116526111a5565b5b835b8181101561167c57806116688882611325565b845260208401935050602081019050611655565b5050509392505050565b600082601f83011261169b5761169a611033565b5b81356116ab84826020860161161d565b91505092915050565b600080600080600080600080610100898b0312156116d5576116d4611029565b5b600089013567ffffffffffffffff8111156116f3576116f261102e565b5b6116ff8b828c0161114b565b98505060206117108b828c016111f3565b97505060406117218b828c016112b6565b96505060606117328b828c016112b6565b955050608089013567ffffffffffffffff8111156117535761175261102e565b5b61175f8b828c01611500565b94505060a089013567ffffffffffffffff8111156117805761177f61102e565b5b61178c8b828c016115c3565b93505060c089013567ffffffffffffffff8111156117ad576117ac61102e565b5b6117b98b828c01611686565b92505060e089013567ffffffffffffffff8111156117da576117d961102e565b5b6117e68b828c01611686565b9150509295985092959890939650565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b600061183d6012836117f6565b915061184882611807565b602082019050919050565b6000602082019050818103600083015261186c81611830565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118ad82610feb565b91506118b883610feb565b92508282039050818111156118d0576118cf611873565b5b92915050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b83811015611930578082015181840152602081019050611915565b60008484015250505050565b6000611947826118fc565b6119518185611907565b9350611961818560208601611912565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006119c4826118d6565b6009820191506119d4828561193c565b91506119df8261196d565b6009820191506119ef828461193c565b91506119fa82611993565b6001820191508190509392505050565b6000611a15826118fc565b611a1f81856117f6565b9350611a2f818560208601611912565b611a388161103d565b840191505092915050565b60006020820190508181036000830152611a5d8184611a0a565b905092915050565b611a6e81611304565b82525050565b611a7d816112cb565b82525050565b6000608082019050611a986000830187611a65565b611aa56020830186611a74565b611ab26040830185611a65565b611abf6060830184611a65565b95945050505050565b7f696e76616c6964206d6f64657261746f72207369676e61747572650000000000600082015250565b6000611afe601b836117f6565b9150611b0982611ac8565b602082019050919050565b60006020820190508181036000830152611b2d81611af1565b9050919050565b6000611b40828461193c565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611b8582610feb565b9150611b9083610feb565b925082611ba057611b9f611b4b565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611be582610feb565b9150611bf083610feb565b9250828201905080821115611c0857611c07611873565b5b92915050565b7f62657449645b0000000000000000000000000000000000000000000000000000815250565b7f5d20626574746f725b0000000000000000000000000000000000000000000000815250565b7f5d2077696e6e696e67735b000000000000000000000000000000000000000000815250565b6000611c8b82611c0e565b600682019150611c9b828661193c565b9150611ca682611c34565b600982019150611cb6828561193c565b9150611cc182611c5a565b600b82019150611cd1828461193c565b9150611cdc82611993565b600182019150819050949350505050565b6000611cf882610feb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d2a57611d29611873565b5b600182019050919050565b7f7472616e736665725b0000000000000000000000000000000000000000000000815250565b6000611d6682611d35565b600982019150611d76828461193c565b9150611d8182611993565b60018201915081905092915050565b7f696e73756666696369656e742066756e64730000000000000000000000000000600082015250565b6000611dc66012836117f6565b9150611dd182611d90565b602082019050919050565b60006020820190508181036000830152611df581611db9565b9050919050565b7f5d206265745b0000000000000000000000000000000000000000000000000000815250565b6000611e2d82611c0e565b600682019150611e3d828661193c565b9150611e4882611c34565b600982019150611e58828561193c565b9150611e6382611dfc565b600682019150611e73828461193c565b9150611e7e82611993565b600182019150819050949350505050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b6000611ee682611e8f565b600882019150611ef6828561193c565b9150611f0182611eb5565b600a82019150611f11828461193c565b9150611f1c82611993565b6001820191508190509392505050565b6000611f3782610feb565b9150611f4283610feb565b9250828202611f5081610feb565b91508282048414831517611f6757611f66611873565b5b5092915050565b60008160011c9050919050565b6000808291508390505b6001851115611fc557808604811115611fa157611fa0611873565b5b6001851615611fb05780820291505b8081029050611fbe85611f6e565b9450611f85565b94509492505050565b600082611fde576001905061209a565b81611fec576000905061209a565b8160018114612002576002811461200c5761203b565b600191505061209a565b60ff84111561201e5761201d611873565b5b8360020a91508482111561203557612034611873565b5b5061209a565b5060208310610133831016604e8410600b84101617156120705782820a90508381111561206b5761206a611873565b5b61209a565b61207d8484846001611f7b565b9250905081840481111561209457612093611873565b5b81810290505b9392505050565b60006120ac82610feb565b91506120b783610feb565b92506120e47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611fce565b905092915050565b60006120f7826112cb565b9150612102836112cb565b92508261211257612111611b4b565b5b828204905092915050565b6000612128826112cb565b9150612133836112cb565b9250828202612141816112cb565b915080821461215357612152611873565b5b5092915050565b6000612165826112cb565b9150612170836112cb565b9250828203905060ff81111561218957612188611873565b5b92915050565b600061219a826112cb565b91506121a5836112cb565b9250828201905060ff8111156121be576121bd611873565b5b92915050565b600081519050919050565b600081905092915050565b6000819050602082019050919050565b6121f3816111ca565b82525050565b600061220583836121ea565b60208301905092915050565b6000602082019050919050565b6000612229826121c4565b61223381856121cf565b935061223e836121da565b8060005b8381101561226f57815161225688826121f9565b975061226183612211565b925050600181019050612242565b5085935050505092915050565b60008160601b9050919050565b60006122948261227c565b9050919050565b60006122a682612289565b9050919050565b6122be6122b9826111ca565b61229b565b82525050565b6000819050919050565b6122df6122da82610feb565b6122c4565b82525050565b60006122f1828761193c565b91506122fd828661221e565b915061230982856122ad565b60148201915061231982846122ce565b60208201915081905095945050505050565b6000612337828861193c565b915061234382876122ad565b60148201915061235382866122ce565b60208201915061236382856122ce565b60208201915061237382846122ce565b6020820191508190509695505050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b60006123bc601c83611907565b91506123c782612386565b601c82019050919050565b6000819050919050565b6123ed6123e882611304565b6123d2565b82525050565b60006123fe826123af565b915061240a82846123dc565b6020820191508190509291505056fea26469706673582212204df0828e1495a89c6625be55301a97099fe7b1f94ad82541fd1d7f0ee905fcf864736f6c63430008110033",
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
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
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

// PlaceBetsSigned is a paid mutator transaction binding the contract method 0xbb9c53b9.
//
// Solidity: function PlaceBetsSigned(string betId, address moderator, uint256 amount, uint256 expiration, uint256[] nonce, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Bank *BankTransactor) PlaceBetsSigned(opts *bind.TransactOpts, betId string, moderator common.Address, amount *big.Int, expiration *big.Int, nonce []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "PlaceBetsSigned", betId, moderator, amount, expiration, nonce, v, r, s)
}

// PlaceBetsSigned is a paid mutator transaction binding the contract method 0xbb9c53b9.
//
// Solidity: function PlaceBetsSigned(string betId, address moderator, uint256 amount, uint256 expiration, uint256[] nonce, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Bank *BankSession) PlaceBetsSigned(betId string, moderator common.Address, amount *big.Int, expiration *big.Int, nonce []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Bank.Contract.PlaceBetsSigned(&_Bank.TransactOpts, betId, moderator, amount, expiration, nonce, v, r, s)
}

// PlaceBetsSigned is a paid mutator transaction binding the contract method 0xbb9c53b9.
//
// Solidity: function PlaceBetsSigned(string betId, address moderator, uint256 amount, uint256 expiration, uint256[] nonce, uint8[] v, bytes32[] r, bytes32[] s) returns()
func (_Bank *BankTransactorSession) PlaceBetsSigned(betId string, moderator common.Address, amount *big.Int, expiration *big.Int, nonce []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Bank.Contract.PlaceBetsSigned(&_Bank.TransactOpts, betId, moderator, amount, expiration, nonce, v, r, s)
}

// ReconcileSigned is a paid mutator transaction binding the contract method 0x82156760.
//
// Solidity: function ReconcileSigned(string betId, address[] winners, address moderator, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bank *BankTransactor) ReconcileSigned(opts *bind.TransactOpts, betId string, winners []common.Address, moderator common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "ReconcileSigned", betId, winners, moderator, nonce, v, r, s)
}

// ReconcileSigned is a paid mutator transaction binding the contract method 0x82156760.
//
// Solidity: function ReconcileSigned(string betId, address[] winners, address moderator, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bank *BankSession) ReconcileSigned(betId string, winners []common.Address, moderator common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bank.Contract.ReconcileSigned(&_Bank.TransactOpts, betId, winners, moderator, nonce, v, r, s)
}

// ReconcileSigned is a paid mutator transaction binding the contract method 0x82156760.
//
// Solidity: function ReconcileSigned(string betId, address[] winners, address moderator, uint256 nonce, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bank *BankTransactorSession) ReconcileSigned(betId string, winners []common.Address, moderator common.Address, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bank.Contract.ReconcileSigned(&_Bank.TransactOpts, betId, winners, moderator, nonce, v, r, s)
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
