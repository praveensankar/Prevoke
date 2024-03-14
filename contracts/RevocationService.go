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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"GetMerkleTreeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"RetrieveMerkleTree\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrievePublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"addPublicKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_values\",\"type\":\"bytes32[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611dcf8061005d5f395ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c806351600698116100ab578063bbb7e2ef1161006f578063bbb7e2ef14610301578063c0a9290b1461031f578063c680f4101461034f578063ce4b3f341461037f578063f8a8fd6d1461039b5761011f565b8063516006981461025f5780635a34497a1461027b5780636a6f20631461029757806399223a5a146102c7578063a5eb0de8146102e55761011f565b80632337db35116100f25780632337db35146101a9578063266344c2146101c757806348db5336146101e357806350308ca214610213578063515586751461022f5761011f565b80630de54b85146101235780630df0ff901461012d57806319f267171461015d578063203c8ab014610179575b5f80fd5b61012b6103a5565b005b61014760048036038101906101429190610ffc565b6103a7565b604051610154919061103f565b60405180910390f35b61017760048036038101906101729190611082565b6103bc565b005b610193600480360381019061018e9190611210565b61042e565b6040516101a09190611271565b60405180910390f35b6101b16104d5565b6040516101be919061103f565b60405180910390f35b6101e160048036038101906101dc919061134a565b6104ed565b005b6101fd60048036038101906101f89190610ffc565b61055d565b60405161020a91906113fd565b60405180910390f35b61022d60048036038101906102289190610ffc565b610571565b005b61024960048036038101906102449190610ffc565b610664565b60405161025691906114cd565b60405180910390f35b610279600480360381019061027491906114ed565b61070a565b005b610295600480360381019061029091906116f1565b6107dc565b005b6102b160048036038101906102ac9190611210565b6108a2565b6040516102be9190611271565b60405180910390f35b6102cf6108b3565b6040516102dc919061186d565b60405180910390f35b6102ff60048036038101906102fa9190611210565b610987565b005b610309610a67565b604051610316919061103f565b60405180910390f35b61033960048036038101906103349190610ffc565b610a7f565b60405161034691906113fd565b60405180910390f35b61036960048036038101906103649190610ffc565b610acf565b60405161037691906118d5565b60405180910390f35b610399600480360381019061039491906114ed565b610b75565b005b6103a3610bdb565b005b565b6001602052805f5260405f205f915090505481565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610414575f80fd5b8060015f8481526020019081526020015f20819055505050565b5f805f90505f5b83518110156104cb575f6008858381518110610454576104536118f5565b5b6020026020010151901c90505f60ff868481518110610476576104756118f5565b5b6020026020010151166001901b90505f80825f808681526020019081526020015f205416141590505f1515811515036104b557600194505050506104cb565b50505080806104c39061194f565b915050610435565b5080915050919050565b5f60015f8081526020019081526020015f2054905090565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610545575f80fd5b61054e83610987565b610558828261070a565b505050565b5f602052805f5260405f205f915090505481565b6001151560011515036105bd576105bc6040518060400160405280601381526020017f70726974696e67206d65726b6c65207472656500000000000000000000000000815250610d9c565b5b5f5b818110156106605760011515600115150361064d576106136040518060400160405280601481526020017f696e646578203a20256420092076616c7565203a00000000000000000000000081525082610e35565b61064c60015f8381526020019081526020015f205460405160200161063891906119b6565b604051602081830303815290604052610ed1565b5b80806106589061194f565b9150506105bf565b5050565b60605f8267ffffffffffffffff811115610681576106806110d4565b5b6040519080825280602002602001820160405280156106af5781602001602082028036833780820191505090505b5090505f5b838110156107005760015f8281526020019081526020015f20548282815181106106e1576106e06118f5565b5b60200260200101818152505080806106f89061194f565b9150506106b4565b5080915050919050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610762575f80fd5b805182511461076f575f80fd5b5f5b82518110156107d75781818151811061078d5761078c6118f5565b5b602002602001015160015f8584815181106107ab576107aa6118f5565b5b602002602001015181526020019081526020015f208190555080806107cf9061194f565b915050610771565b505050565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610834575f80fd5b5f5b815181101561089e576003828281518110610854576108536118f5565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150908161088a9190611bca565b5080806108969061194f565b915050610836565b5050565b5f6108ac8261042e565b9050919050565b60606003805480602002602001604051908101604052809291908181526020015f905b8282101561097e578382905f5260205f200180546108f3906119fd565b80601f016020809104026020016040519081016040528092919081815260200182805461091f906119fd565b801561096a5780601f106109415761010080835404028352916020019161096a565b820191905f5260205f20905b81548152906001019060200180831161094d57829003601f168201915b5050505050815260200190600101906108d6565b50505050905090565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109df575f80fd5b5f5b8151811015610a63575f6008838381518110610a00576109ff6118f5565b5b6020026020010151901c90505f60ff848481518110610a2257610a216118f5565b5b6020026020010151166001901b9050805f808481526020019081526020015f205f828254179250508190555050508080610a5b9061194f565b9150506109e1565b5050565b5f60015f8081526020019081526020015f2054905090565b5f805f90505f5b83811015610ac55760015f8281526020019081526020015f2050602060ff1682610ab09190611c99565b91508080610abd9061194f565b915050610a86565b5080915050919050565b60038181548110610ade575f80fd5b905f5260205f20015f915090508054610af6906119fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610b22906119fd565b8015610b6d5780601f10610b4457610100808354040283529160200191610b6d565b820191905f5260205f20905b815481529060010190602001808311610b5057829003601f168201915b505050505081565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610bcd575f80fd5b610bd7828261070a565b5050565b5f600267ffffffffffffffff811115610bf757610bf66110d4565b5b604051908082528060200260200182016040528015610c255781602001602082028036833780820191505090505b5090505f815f81518110610c3c57610c3b6118f5565b5b602002602001018181525050600181600181518110610c5e57610c5d6118f5565b5b6020026020010181815250505f600267ffffffffffffffff811115610c8657610c856110d4565b5b604051908082528060200260200182016040528015610cb45781602001602082028036833780820191505090505b5090507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b815f81518110610ced57610cec6118f5565b5b6020026020010181815250507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b81600181518110610d3057610d2f6118f5565b5b602002602001018181525050610d468282610b75565b610d506002610571565b610d986040518060400160405280601581526020017f6d65726b6c6520747265652073697a653a202564200000000000000000000000815250610d936002610a7f565b610e35565b5050565b610e3281604051602401610db09190611d1e565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610f6a565b50565b610ecd8282604051602401610e4b929190611d3e565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610f6a565b5050565b610f6781604051602401610ee591906118d5565b6040516020818303038152906040527f0be77f56000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610f6a565b50565b610f8181610f79610f84610fa3565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b610fae819050919050565b610fb6611d6c565b565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610fdb81610fc9565b8114610fe5575f80fd5b50565b5f81359050610ff681610fd2565b92915050565b5f6020828403121561101157611010610fc1565b5b5f61101e84828501610fe8565b91505092915050565b5f819050919050565b61103981611027565b82525050565b5f6020820190506110525f830184611030565b92915050565b61106181611027565b811461106b575f80fd5b50565b5f8135905061107c81611058565b92915050565b5f806040838503121561109857611097610fc1565b5b5f6110a585828601610fe8565b92505060206110b68582860161106e565b9150509250929050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61110a826110c4565b810181811067ffffffffffffffff82111715611129576111286110d4565b5b80604052505050565b5f61113b610fb8565b90506111478282611101565b919050565b5f67ffffffffffffffff821115611166576111656110d4565b5b602082029050602081019050919050565b5f80fd5b5f61118d6111888461114c565b611132565b905080838252602082019050602084028301858111156111b0576111af611177565b5b835b818110156111d957806111c58882610fe8565b8452602084019350506020810190506111b2565b5050509392505050565b5f82601f8301126111f7576111f66110c0565b5b813561120784826020860161117b565b91505092915050565b5f6020828403121561122557611224610fc1565b5b5f82013567ffffffffffffffff81111561124257611241610fc5565b5b61124e848285016111e3565b91505092915050565b5f8115159050919050565b61126b81611257565b82525050565b5f6020820190506112845f830184611262565b92915050565b5f67ffffffffffffffff8211156112a4576112a36110d4565b5b602082029050602081019050919050565b5f6112c76112c28461128a565b611132565b905080838252602082019050602084028301858111156112ea576112e9611177565b5b835b8181101561131357806112ff888261106e565b8452602084019350506020810190506112ec565b5050509392505050565b5f82601f830112611331576113306110c0565b5b81356113418482602086016112b5565b91505092915050565b5f805f6060848603121561136157611360610fc1565b5b5f84013567ffffffffffffffff81111561137e5761137d610fc5565b5b61138a868287016111e3565b935050602084013567ffffffffffffffff8111156113ab576113aa610fc5565b5b6113b7868287016111e3565b925050604084013567ffffffffffffffff8111156113d8576113d7610fc5565b5b6113e48682870161131d565b9150509250925092565b6113f781610fc9565b82525050565b5f6020820190506114105f8301846113ee565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61144881611027565b82525050565b5f611459838361143f565b60208301905092915050565b5f602082019050919050565b5f61147b82611416565b6114858185611420565b935061149083611430565b805f5b838110156114c05781516114a7888261144e565b97506114b283611465565b925050600181019050611493565b5085935050505092915050565b5f6020820190508181035f8301526114e58184611471565b905092915050565b5f806040838503121561150357611502610fc1565b5b5f83013567ffffffffffffffff8111156115205761151f610fc5565b5b61152c858286016111e3565b925050602083013567ffffffffffffffff81111561154d5761154c610fc5565b5b6115598582860161131d565b9150509250929050565b5f67ffffffffffffffff82111561157d5761157c6110d4565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff8211156115ac576115ab6110d4565b5b6115b5826110c4565b9050602081019050919050565b828183375f83830152505050565b5f6115e26115dd84611592565b611132565b9050828152602081018484840111156115fe576115fd61158e565b5b6116098482856115c2565b509392505050565b5f82601f830112611625576116246110c0565b5b81356116358482602086016115d0565b91505092915050565b5f61165061164b84611563565b611132565b9050808382526020820190506020840283018581111561167357611672611177565b5b835b818110156116ba57803567ffffffffffffffff811115611698576116976110c0565b5b8086016116a58982611611565b85526020850194505050602081019050611675565b5050509392505050565b5f82601f8301126116d8576116d76110c0565b5b81356116e884826020860161163e565b91505092915050565b5f6020828403121561170657611705610fc1565b5b5f82013567ffffffffffffffff81111561172357611722610fc5565b5b61172f848285016116c4565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561179857808201518184015260208101905061177d565b5f8484015250505050565b5f6117ad82611761565b6117b7818561176b565b93506117c781856020860161177b565b6117d0816110c4565b840191505092915050565b5f6117e683836117a3565b905092915050565b5f602082019050919050565b5f61180482611738565b61180e8185611742565b93508360208202850161182085611752565b805f5b8581101561185b578484038952815161183c85826117db565b9450611847836117ee565b925060208a01995050600181019050611823565b50829750879550505050505092915050565b5f6020820190508181035f83015261188581846117fa565b905092915050565b5f82825260208201905092915050565b5f6118a782611761565b6118b1818561188d565b93506118c181856020860161177b565b6118ca816110c4565b840191505092915050565b5f6020820190508181035f8301526118ed818461189d565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61195982610fc9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361198b5761198a611922565b5b600182019050919050565b5f819050919050565b6119b06119ab82611027565b611996565b82525050565b5f6119c1828461199f565b60208201915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a1457607f821691505b602082108103611a2757611a266119d0565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611a897fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611a4e565b611a938683611a4e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611ace611ac9611ac484610fc9565b611aab565b610fc9565b9050919050565b5f819050919050565b611ae783611ab4565b611afb611af382611ad5565b848454611a5a565b825550505050565b5f90565b611b0f611b03565b611b1a818484611ade565b505050565b5b81811015611b3d57611b325f82611b07565b600181019050611b20565b5050565b601f821115611b8257611b5381611a2d565b611b5c84611a3f565b81016020851015611b6b578190505b611b7f611b7785611a3f565b830182611b1f565b50505b505050565b5f82821c905092915050565b5f611ba25f1984600802611b87565b1980831691505092915050565b5f611bba8383611b93565b9150826002028217905092915050565b611bd382611761565b67ffffffffffffffff811115611bec57611beb6110d4565b5b611bf682546119fd565b611c01828285611b41565b5f60209050601f831160018114611c32575f8415611c20578287015190505b611c2a8582611baf565b865550611c91565b601f198416611c4086611a2d565b5f5b82811015611c6757848901518255600182019150602085019450602081019050611c42565b86831015611c845784890151611c80601f891682611b93565b8355505b6001600288020188555050505b505050505050565b5f611ca382610fc9565b9150611cae83610fc9565b9250828201905080821115611cc657611cc5611922565b5b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f611cf082611ccc565b611cfa8185611cd6565b9350611d0a81856020860161177b565b611d13816110c4565b840191505092915050565b5f6020820190508181035f830152611d368184611ce6565b905092915050565b5f6040820190508181035f830152611d568185611ce6565b9050611d6560208301846113ee565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212208b29177e07fab7a1413df3c5751bad20e6b932c14a7d6fe1795c2c09e3b123b064736f6c63430008150033",
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

// GetMerkleTreeSize is a free data retrieval call binding the contract method 0xc0a9290b.
//
// Solidity: function GetMerkleTreeSize(uint256 _length) view returns(uint256)
func (_RevocationService *RevocationServiceCaller) GetMerkleTreeSize(opts *bind.CallOpts, _length *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "GetMerkleTreeSize", _length)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerkleTreeSize is a free data retrieval call binding the contract method 0xc0a9290b.
//
// Solidity: function GetMerkleTreeSize(uint256 _length) view returns(uint256)
func (_RevocationService *RevocationServiceSession) GetMerkleTreeSize(_length *big.Int) (*big.Int, error) {
	return _RevocationService.Contract.GetMerkleTreeSize(&_RevocationService.CallOpts, _length)
}

// GetMerkleTreeSize is a free data retrieval call binding the contract method 0xc0a9290b.
//
// Solidity: function GetMerkleTreeSize(uint256 _length) view returns(uint256)
func (_RevocationService *RevocationServiceCallerSession) GetMerkleTreeSize(_length *big.Int) (*big.Int, error) {
	return _RevocationService.Contract.GetMerkleTreeSize(&_RevocationService.CallOpts, _length)
}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x51558675.
//
// Solidity: function RetrieveMerkleTree(uint256 _length) view returns(bytes32[])
func (_RevocationService *RevocationServiceCaller) RetrieveMerkleTree(opts *bind.CallOpts, _length *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "RetrieveMerkleTree", _length)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x51558675.
//
// Solidity: function RetrieveMerkleTree(uint256 _length) view returns(bytes32[])
func (_RevocationService *RevocationServiceSession) RetrieveMerkleTree(_length *big.Int) ([][32]byte, error) {
	return _RevocationService.Contract.RetrieveMerkleTree(&_RevocationService.CallOpts, _length)
}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x51558675.
//
// Solidity: function RetrieveMerkleTree(uint256 _length) view returns(bytes32[])
func (_RevocationService *RevocationServiceCallerSession) RetrieveMerkleTree(_length *big.Int) ([][32]byte, error) {
	return _RevocationService.Contract.RetrieveMerkleTree(&_RevocationService.CallOpts, _length)
}

// RetrievePublicKeys is a free data retrieval call binding the contract method 0x99223a5a.
//
// Solidity: function RetrievePublicKeys() view returns(bytes[])
func (_RevocationService *RevocationServiceCaller) RetrievePublicKeys(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "RetrievePublicKeys")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// RetrievePublicKeys is a free data retrieval call binding the contract method 0x99223a5a.
//
// Solidity: function RetrievePublicKeys() view returns(bytes[])
func (_RevocationService *RevocationServiceSession) RetrievePublicKeys() ([][]byte, error) {
	return _RevocationService.Contract.RetrievePublicKeys(&_RevocationService.CallOpts)
}

// RetrievePublicKeys is a free data retrieval call binding the contract method 0x99223a5a.
//
// Solidity: function RetrievePublicKeys() view returns(bytes[])
func (_RevocationService *RevocationServiceCallerSession) RetrievePublicKeys() ([][]byte, error) {
	return _RevocationService.Contract.RetrievePublicKeys(&_RevocationService.CallOpts)
}

// BloomFilter is a free data retrieval call binding the contract method 0x48db5336.
//
// Solidity: function bloomFilter(uint256 ) view returns(uint256)
func (_RevocationService *RevocationServiceCaller) BloomFilter(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "bloomFilter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BloomFilter is a free data retrieval call binding the contract method 0x48db5336.
//
// Solidity: function bloomFilter(uint256 ) view returns(uint256)
func (_RevocationService *RevocationServiceSession) BloomFilter(arg0 *big.Int) (*big.Int, error) {
	return _RevocationService.Contract.BloomFilter(&_RevocationService.CallOpts, arg0)
}

// BloomFilter is a free data retrieval call binding the contract method 0x48db5336.
//
// Solidity: function bloomFilter(uint256 ) view returns(uint256)
func (_RevocationService *RevocationServiceCallerSession) BloomFilter(arg0 *big.Int) (*big.Int, error) {
	return _RevocationService.Contract.BloomFilter(&_RevocationService.CallOpts, arg0)
}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0x203c8ab0.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceCaller) CheckRevocationStatusInBloomFilter(opts *bind.CallOpts, _indexes []*big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "checkRevocationStatusInBloomFilter", _indexes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0x203c8ab0.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceSession) CheckRevocationStatusInBloomFilter(_indexes []*big.Int) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInBloomFilter(&_RevocationService.CallOpts, _indexes)
}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0x203c8ab0.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) CheckRevocationStatusInBloomFilter(_indexes []*big.Int) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInBloomFilter(&_RevocationService.CallOpts, _indexes)
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

// PrintMerkleTree is a free data retrieval call binding the contract method 0x50308ca2.
//
// Solidity: function printMerkleTree(uint256 _length) view returns()
func (_RevocationService *RevocationServiceCaller) PrintMerkleTree(opts *bind.CallOpts, _length *big.Int) error {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "printMerkleTree", _length)

	if err != nil {
		return err
	}

	return err

}

// PrintMerkleTree is a free data retrieval call binding the contract method 0x50308ca2.
//
// Solidity: function printMerkleTree(uint256 _length) view returns()
func (_RevocationService *RevocationServiceSession) PrintMerkleTree(_length *big.Int) error {
	return _RevocationService.Contract.PrintMerkleTree(&_RevocationService.CallOpts, _length)
}

// PrintMerkleTree is a free data retrieval call binding the contract method 0x50308ca2.
//
// Solidity: function printMerkleTree(uint256 _length) view returns()
func (_RevocationService *RevocationServiceCallerSession) PrintMerkleTree(_length *big.Int) error {
	return _RevocationService.Contract.PrintMerkleTree(&_RevocationService.CallOpts, _length)
}

// PublicKeys is a free data retrieval call binding the contract method 0xc680f410.
//
// Solidity: function publicKeys(uint256 ) view returns(bytes)
func (_RevocationService *RevocationServiceCaller) PublicKeys(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "publicKeys", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PublicKeys is a free data retrieval call binding the contract method 0xc680f410.
//
// Solidity: function publicKeys(uint256 ) view returns(bytes)
func (_RevocationService *RevocationServiceSession) PublicKeys(arg0 *big.Int) ([]byte, error) {
	return _RevocationService.Contract.PublicKeys(&_RevocationService.CallOpts, arg0)
}

// PublicKeys is a free data retrieval call binding the contract method 0xc680f410.
//
// Solidity: function publicKeys(uint256 ) view returns(bytes)
func (_RevocationService *RevocationServiceCallerSession) PublicKeys(arg0 *big.Int) ([]byte, error) {
	return _RevocationService.Contract.PublicKeys(&_RevocationService.CallOpts, arg0)
}

// VerificationPhase1 is a free data retrieval call binding the contract method 0x6a6f2063.
//
// Solidity: function verificationPhase1(uint256[] _bfIndexes) view returns(bool)
func (_RevocationService *RevocationServiceCaller) VerificationPhase1(opts *bind.CallOpts, _bfIndexes []*big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "verificationPhase1", _bfIndexes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificationPhase1 is a free data retrieval call binding the contract method 0x6a6f2063.
//
// Solidity: function verificationPhase1(uint256[] _bfIndexes) view returns(bool)
func (_RevocationService *RevocationServiceSession) VerificationPhase1(_bfIndexes []*big.Int) (bool, error) {
	return _RevocationService.Contract.VerificationPhase1(&_RevocationService.CallOpts, _bfIndexes)
}

// VerificationPhase1 is a free data retrieval call binding the contract method 0x6a6f2063.
//
// Solidity: function verificationPhase1(uint256[] _bfIndexes) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) VerificationPhase1(_bfIndexes []*big.Int) (bool, error) {
	return _RevocationService.Contract.VerificationPhase1(&_RevocationService.CallOpts, _bfIndexes)
}

// VerificationPhase2 is a free data retrieval call binding the contract method 0xbbb7e2ef.
//
// Solidity: function verificationPhase2() view returns(bytes32)
func (_RevocationService *RevocationServiceCaller) VerificationPhase2(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "verificationPhase2")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerificationPhase2 is a free data retrieval call binding the contract method 0xbbb7e2ef.
//
// Solidity: function verificationPhase2() view returns(bytes32)
func (_RevocationService *RevocationServiceSession) VerificationPhase2() ([32]byte, error) {
	return _RevocationService.Contract.VerificationPhase2(&_RevocationService.CallOpts)
}

// VerificationPhase2 is a free data retrieval call binding the contract method 0xbbb7e2ef.
//
// Solidity: function verificationPhase2() view returns(bytes32)
func (_RevocationService *RevocationServiceCallerSession) VerificationPhase2() ([32]byte, error) {
	return _RevocationService.Contract.VerificationPhase2(&_RevocationService.CallOpts)
}

// VerificationPhase2Test is a free data retrieval call binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() view returns(bytes32)
func (_RevocationService *RevocationServiceCaller) VerificationPhase2Test(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "verificationPhase2Test")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerificationPhase2Test is a free data retrieval call binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() view returns(bytes32)
func (_RevocationService *RevocationServiceSession) VerificationPhase2Test() ([32]byte, error) {
	return _RevocationService.Contract.VerificationPhase2Test(&_RevocationService.CallOpts)
}

// VerificationPhase2Test is a free data retrieval call binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() view returns(bytes32)
func (_RevocationService *RevocationServiceCallerSession) VerificationPhase2Test() ([32]byte, error) {
	return _RevocationService.Contract.VerificationPhase2Test(&_RevocationService.CallOpts)
}

// AddPublicKeys is a paid mutator transaction binding the contract method 0x5a34497a.
//
// Solidity: function addPublicKeys(bytes[] _publicKeys) returns()
func (_RevocationService *RevocationServiceTransactor) AddPublicKeys(opts *bind.TransactOpts, _publicKeys [][]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "addPublicKeys", _publicKeys)
}

// AddPublicKeys is a paid mutator transaction binding the contract method 0x5a34497a.
//
// Solidity: function addPublicKeys(bytes[] _publicKeys) returns()
func (_RevocationService *RevocationServiceSession) AddPublicKeys(_publicKeys [][]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.AddPublicKeys(&_RevocationService.TransactOpts, _publicKeys)
}

// AddPublicKeys is a paid mutator transaction binding the contract method 0x5a34497a.
//
// Solidity: function addPublicKeys(bytes[] _publicKeys) returns()
func (_RevocationService *RevocationServiceTransactorSession) AddPublicKeys(_publicKeys [][]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.AddPublicKeys(&_RevocationService.TransactOpts, _publicKeys)
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

// RevokeVC is a paid mutator transaction binding the contract method 0x266344c2.
//
// Solidity: function revokeVC(uint256[] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactor) RevokeVC(opts *bind.TransactOpts, _bfIndexes []*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "revokeVC", _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x266344c2.
//
// Solidity: function revokeVC(uint256[] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceSession) RevokeVC(_bfIndexes []*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x266344c2.
//
// Solidity: function revokeVC(uint256[] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactorSession) RevokeVC(_bfIndexes []*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_RevocationService *RevocationServiceTransactor) Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "test")
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_RevocationService *RevocationServiceSession) Test() (*types.Transaction, error) {
	return _RevocationService.Contract.Test(&_RevocationService.TransactOpts)
}

// Test is a paid mutator transaction binding the contract method 0xf8a8fd6d.
//
// Solidity: function test() returns()
func (_RevocationService *RevocationServiceTransactorSession) Test() (*types.Transaction, error) {
	return _RevocationService.Contract.Test(&_RevocationService.TransactOpts)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0xa5eb0de8.
//
// Solidity: function updateBloomFilter(uint256[] _indexes) returns()
func (_RevocationService *RevocationServiceTransactor) UpdateBloomFilter(opts *bind.TransactOpts, _indexes []*big.Int) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "updateBloomFilter", _indexes)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0xa5eb0de8.
//
// Solidity: function updateBloomFilter(uint256[] _indexes) returns()
func (_RevocationService *RevocationServiceSession) UpdateBloomFilter(_indexes []*big.Int) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateBloomFilter(&_RevocationService.TransactOpts, _indexes)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0xa5eb0de8.
//
// Solidity: function updateBloomFilter(uint256[] _indexes) returns()
func (_RevocationService *RevocationServiceTransactorSession) UpdateBloomFilter(_indexes []*big.Int) (*types.Transaction, error) {
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
