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
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061243a806100606000396000f3fe60806040526004361061007b5760003560e01c8063b4a99a4e1161004e578063b4a99a4e146100fa578063bb9c53b914610125578063e63f341f1461014e578063ed21248c1461018b5761007b565b80630ef678871461008057806357ea89b6146100ab57806382156760146100b557806393f28237146100de575b600080fd5b34801561008c57600080fd5b50610095610195565b6040516100a29190611000565b60405180910390f35b6100b36101dc565b005b3480156100c157600080fd5b506100dc60048036038101906100d79190611336565b610371565b005b6100f860048036038101906100f39190611410565b61060c565b005b34801561010657600080fd5b5061010f610718565b60405161011c919061144c565b60405180910390f35b34801561013157600080fd5b5061014c600480360381019061014791906116b0565b61073c565b005b34801561015a57600080fd5b5061017560048036038101906101709190611410565b610a03565b6040516101829190611000565b60405180910390f35b610193610aa5565b005b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60003390506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008103610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025f9061184f565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156102ae573d6000803e3d6000fd5b5080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546102fe919061189e565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61032f33610ba4565b61033883610d67565b6040516020016103499291906119b5565b6040516020818303038152906040526040516103659190611a3f565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103c957600080fd5b60006103d6888887610eef565b90506000600182868686604051600081526020016040526040516103fd9493929190611a7f565b6020604051602081039080840390855afa15801561041f573d6000803e3d6000fd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff1614610499576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049090611b10565b60405180910390fd5b6000885160028b6040516104ad9190611b30565b9081526020016040518091039020546104c69190611b76565b905060005b89518110156105db5781600160008c84815181106104ec576104eb611ba7565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461053d9190611bd6565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a8b6105898c848151811061057c5761057b611ba7565b5b6020026020010151610ba4565b61059285610d67565b6040516020016105a493929190611c7c565b6040516020818303038152906040526040516105c09190611a3f565b60405180910390a180806105d390611ce9565b9150506104cb565b50600060028b6040516105ee9190611b30565b90815260200160405180910390208190555050505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066457600080fd5b600047905060008290508073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156106b4573d6000803e3d6000fd5b507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106df83610d67565b6040516020016106ef9190611d57565b60405160208183030381529060405260405161070b9190611a3f565b60405180910390a1505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461079457600080fd5b60005b84518110156109f85760006107c98a8a8a8a8a87815181106107bc576107bb611ba7565b5b6020026020010151610f2d565b905060006001828785815181106107e3576107e2611ba7565b5b60200260200101518786815181106107fe576107fd611ba7565b5b602002602001015187878151811061081957610818611ba7565b5b60200260200101516040516000815260200160405260405161083e9493929190611a7f565b6020604051602081039080840390855afa158015610860573d6000803e3d6000fd5b50505060206040510351905088600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156108ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e590611dd8565b60405180910390fd5b8860028c6040516108ff9190611b30565b9081526020016040518091039020600082825461091c9190611bd6565b9250508190555088600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610972919061189e565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a8b6109a483610ba4565b6109ad8c610d67565b6040516020016109bf93929190611e1e565b6040516020818303038152906040526040516109db9190611a3f565b60405180910390a1505080806109f090611ce9565b915050610797565b505050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a5e57600080fd5b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b34600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610af49190611bd6565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610b2533610ba4565b610b6d600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610d67565b604051602001610b7e929190611ed7565b604051602081830303815290604052604051610b9a9190611a3f565b60405180910390a1565b60606000602867ffffffffffffffff811115610bc357610bc261104a565b5b6040519080825280601f01601f191660200182016040528015610bf55781602001600182028036833780820191505090505b50905060005b6014811015610d5d576000816013610c13919061189e565b6008610c1f9190611f28565b6002610c2b919061209d565b8573ffffffffffffffffffffffffffffffffffffffff16610c4c9190611b76565b60f81b9050600060108260f81c610c6391906120e8565b60f81b905060008160f81c6010610c7a9190612119565b8360f81c610c889190612156565b60f81b9050610c9682610f71565b85856002610ca49190611f28565b81518110610cb557610cb4611ba7565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610ced81610f71565b856001866002610cfd9190611f28565b610d079190611bd6565b81518110610d1857610d17611ba7565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505050508080610d5590611ce9565b915050610bfb565b5080915050919050565b606060008203610dae576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610eea565b600082905060005b60008214610de0578080610dc990611ce9565b915050600a82610dd99190611b76565b9150610db6565b60008167ffffffffffffffff811115610dfc57610dfb61104a565b5b6040519080825280601f01601f191660200182016040528015610e2e5781602001600182028036833780820191505090505b50905060008290505b60008614610ee257600181610e4c919061189e565b90506000600a8088610e5e9190611b76565b610e689190611f28565b87610e73919061189e565b6030610e7f919061218b565b905060008160f81b905080848481518110610e9d57610e9c611ba7565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a88610ed99190611b76565b97505050610e37565b819450505050505b919050565b6000610f24848484604051602001610f0993929190612299565b60405160208183030381529060405280519060200120610fb7565b90509392505050565b6000610f668686868686604051602001610f4b959493929190612316565b60405160208183030381529060405280519060200120610fb7565b905095945050505050565b6000600a8260f81c60ff161015610f9c5760308260f81c610f92919061218b565b60f81b9050610fb2565b60578260f81c610fac919061218b565b60f81b90505b919050565b600081604051602001610fca91906123de565b604051602081830303815290604052805190602001209050919050565b6000819050919050565b610ffa81610fe7565b82525050565b60006020820190506110156000830184610ff1565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61108282611039565b810181811067ffffffffffffffff821117156110a1576110a061104a565b5b80604052505050565b60006110b461101b565b90506110c08282611079565b919050565b600067ffffffffffffffff8211156110e0576110df61104a565b5b6110e982611039565b9050602081019050919050565b82818337600083830152505050565b6000611118611113846110c5565b6110aa565b90508281526020810184848401111561113457611133611034565b5b61113f8482856110f6565b509392505050565b600082601f83011261115c5761115b61102f565b5b813561116c848260208601611105565b91505092915050565b600067ffffffffffffffff8211156111905761118f61104a565b5b602082029050602081019050919050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006111d1826111a6565b9050919050565b6111e1816111c6565b81146111ec57600080fd5b50565b6000813590506111fe816111d8565b92915050565b600061121761121284611175565b6110aa565b9050808382526020820190506020840283018581111561123a576112396111a1565b5b835b81811015611263578061124f88826111ef565b84526020840193505060208101905061123c565b5050509392505050565b600082601f8301126112825761128161102f565b5b8135611292848260208601611204565b91505092915050565b6112a481610fe7565b81146112af57600080fd5b50565b6000813590506112c18161129b565b92915050565b600060ff82169050919050565b6112dd816112c7565b81146112e857600080fd5b50565b6000813590506112fa816112d4565b92915050565b6000819050919050565b61131381611300565b811461131e57600080fd5b50565b6000813590506113308161130a565b92915050565b600080600080600080600060e0888a03121561135557611354611025565b5b600088013567ffffffffffffffff8111156113735761137261102a565b5b61137f8a828b01611147565b975050602088013567ffffffffffffffff8111156113a05761139f61102a565b5b6113ac8a828b0161126d565b96505060406113bd8a828b016111ef565b95505060606113ce8a828b016112b2565b94505060806113df8a828b016112eb565b93505060a06113f08a828b01611321565b92505060c06114018a828b01611321565b91505092959891949750929550565b60006020828403121561142657611425611025565b5b6000611434848285016111ef565b91505092915050565b611446816111c6565b82525050565b6000602082019050611461600083018461143d565b92915050565b600067ffffffffffffffff8211156114825761148161104a565b5b602082029050602081019050919050565b60006114a66114a184611467565b6110aa565b905080838252602082019050602084028301858111156114c9576114c86111a1565b5b835b818110156114f257806114de88826112b2565b8452602084019350506020810190506114cb565b5050509392505050565b600082601f8301126115115761151061102f565b5b8135611521848260208601611493565b91505092915050565b600067ffffffffffffffff8211156115455761154461104a565b5b602082029050602081019050919050565b60006115696115648461152a565b6110aa565b9050808382526020820190506020840283018581111561158c5761158b6111a1565b5b835b818110156115b557806115a188826112eb565b84526020840193505060208101905061158e565b5050509392505050565b600082601f8301126115d4576115d361102f565b5b81356115e4848260208601611556565b91505092915050565b600067ffffffffffffffff8211156116085761160761104a565b5b602082029050602081019050919050565b600061162c611627846115ed565b6110aa565b9050808382526020820190506020840283018581111561164f5761164e6111a1565b5b835b8181101561167857806116648882611321565b845260208401935050602081019050611651565b5050509392505050565b600082601f8301126116975761169661102f565b5b81356116a7848260208601611619565b91505092915050565b600080600080600080600080610100898b0312156116d1576116d0611025565b5b600089013567ffffffffffffffff8111156116ef576116ee61102a565b5b6116fb8b828c01611147565b985050602061170c8b828c016111ef565b975050604061171d8b828c016112b2565b965050606061172e8b828c016112b2565b955050608089013567ffffffffffffffff81111561174f5761174e61102a565b5b61175b8b828c016114fc565b94505060a089013567ffffffffffffffff81111561177c5761177b61102a565b5b6117888b828c016115bf565b93505060c089013567ffffffffffffffff8111156117a9576117a861102a565b5b6117b58b828c01611682565b92505060e089013567ffffffffffffffff8111156117d6576117d561102a565b5b6117e28b828c01611682565b9150509295985092959890939650565b600082825260208201905092915050565b7f6e6f7420656e6f7567682062616c616e63650000000000000000000000000000600082015250565b60006118396012836117f2565b915061184482611803565b602082019050919050565b600060208201905081810360008301526118688161182c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118a982610fe7565b91506118b483610fe7565b92508282039050818111156118cc576118cb61186f565b5b92915050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b600081519050919050565b600081905092915050565b60005b8381101561192c578082015181840152602081019050611911565b60008484015250505050565b6000611943826118f8565b61194d8185611903565b935061195d81856020860161190e565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b60006119c0826118d2565b6009820191506119d08285611938565b91506119db82611969565b6009820191506119eb8284611938565b91506119f68261198f565b6001820191508190509392505050565b6000611a11826118f8565b611a1b81856117f2565b9350611a2b81856020860161190e565b611a3481611039565b840191505092915050565b60006020820190508181036000830152611a598184611a06565b905092915050565b611a6a81611300565b82525050565b611a79816112c7565b82525050565b6000608082019050611a946000830187611a61565b611aa16020830186611a70565b611aae6040830185611a61565b611abb6060830184611a61565b95945050505050565b7f696e76616c6964206d6f64657261746f72207369676e61747572650000000000600082015250565b6000611afa601b836117f2565b9150611b0582611ac4565b602082019050919050565b60006020820190508181036000830152611b2981611aed565b9050919050565b6000611b3c8284611938565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611b8182610fe7565b9150611b8c83610fe7565b925082611b9c57611b9b611b47565b5b828204905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611be182610fe7565b9150611bec83610fe7565b9250828201905080821115611c0457611c0361186f565b5b92915050565b7f62657449645b0000000000000000000000000000000000000000000000000000815250565b7f5d20626574746f725b0000000000000000000000000000000000000000000000815250565b7f5d2077696e6e696e67735b000000000000000000000000000000000000000000815250565b6000611c8782611c0a565b600682019150611c978286611938565b9150611ca282611c30565b600982019150611cb28285611938565b9150611cbd82611c56565b600b82019150611ccd8284611938565b9150611cd88261198f565b600182019150819050949350505050565b6000611cf482610fe7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d2657611d2561186f565b5b600182019050919050565b7f7472616e736665725b0000000000000000000000000000000000000000000000815250565b6000611d6282611d31565b600982019150611d728284611938565b9150611d7d8261198f565b60018201915081905092915050565b7f696e73756666696369656e742066756e64730000000000000000000000000000600082015250565b6000611dc26012836117f2565b9150611dcd82611d8c565b602082019050919050565b60006020820190508181036000830152611df181611db5565b9050919050565b7f5d206265745b0000000000000000000000000000000000000000000000000000815250565b6000611e2982611c0a565b600682019150611e398286611938565b9150611e4482611c30565b600982019150611e548285611938565b9150611e5f82611df8565b600682019150611e6f8284611938565b9150611e7a8261198f565b600182019150819050949350505050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b6000611ee282611e8b565b600882019150611ef28285611938565b9150611efd82611eb1565b600a82019150611f0d8284611938565b9150611f188261198f565b6001820191508190509392505050565b6000611f3382610fe7565b9150611f3e83610fe7565b9250828202611f4c81610fe7565b91508282048414831517611f6357611f6261186f565b5b5092915050565b60008160011c9050919050565b6000808291508390505b6001851115611fc157808604811115611f9d57611f9c61186f565b5b6001851615611fac5780820291505b8081029050611fba85611f6a565b9450611f81565b94509492505050565b600082611fda5760019050612096565b81611fe85760009050612096565b8160018114611ffe576002811461200857612037565b6001915050612096565b60ff84111561201a5761201961186f565b5b8360020a9150848211156120315761203061186f565b5b50612096565b5060208310610133831016604e8410600b841016171561206c5782820a9050838111156120675761206661186f565b5b612096565b6120798484846001611f77565b925090508184048111156120905761208f61186f565b5b81810290505b9392505050565b60006120a882610fe7565b91506120b383610fe7565b92506120e07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611fca565b905092915050565b60006120f3826112c7565b91506120fe836112c7565b92508261210e5761210d611b47565b5b828204905092915050565b6000612124826112c7565b915061212f836112c7565b925082820261213d816112c7565b915080821461214f5761214e61186f565b5b5092915050565b6000612161826112c7565b915061216c836112c7565b9250828203905060ff8111156121855761218461186f565b5b92915050565b6000612196826112c7565b91506121a1836112c7565b9250828201905060ff8111156121ba576121b961186f565b5b92915050565b600081519050919050565b600081905092915050565b6000819050602082019050919050565b6121ef816111c6565b82525050565b600061220183836121e6565b60208301905092915050565b6000602082019050919050565b6000612225826121c0565b61222f81856121cb565b935061223a836121d6565b8060005b8381101561226b57815161225288826121f5565b975061225d8361220d565b92505060018101905061223e565b5085935050505092915050565b6000819050919050565b61229361228e82610fe7565b612278565b82525050565b60006122a58286611938565b91506122b1828561221a565b91506122bd8284612282565b602082019150819050949350505050565b60008160601b9050919050565b60006122e6826122ce565b9050919050565b60006122f8826122db565b9050919050565b61231061230b826111c6565b6122ed565b82525050565b60006123228288611938565b915061232e82876122ff565b60148201915061233e8286612282565b60208201915061234e8285612282565b60208201915061235e8284612282565b6020820191508190509695505050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b60006123a7601c83611903565b91506123b282612371565b601c82019050919050565b6000819050919050565b6123d86123d382611300565b6123bd565b82525050565b60006123e98261239a565b91506123f582846123c7565b6020820191508190509291505056fea26469706673582212209dbebea950437754b4b9817120ddd8171a7af6556b660f71c21b1850596a1efc64736f6c63430008110033",
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
