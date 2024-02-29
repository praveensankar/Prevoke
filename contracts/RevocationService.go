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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetMerkleTreeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveBloomFilter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveMerkleTree\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrievePublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"addPublicKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bfIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInMTAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"isExistInPublicKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_values\",\"type\":\"string[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506126928061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610156575f3560e01c80635a34497a116100c1578063bbb7e2ef1161007a578063bbb7e2ef146103c2578063c2f0b9f1146103e0578063c680f41014610410578063d3e5304214610440578063e5440a9b1461044a578063f174053b1461047a57610156565b80635a34497a146102ee5780635d2c4a381461030a5780636a6f20631461033a5780636e0a59041461036a57806399223a5a14610388578063a5eb0de8146103a657610156565b80632eb4a7ab116101135780632eb4a7ab1461021c5780632ef76b181461023a57806334d6c77f14610256578063392886641461027257806348db5336146102a257806353491af5146102d257610156565b8063070f010e1461015a5780630de54b85146101785780630df0ff9014610182578063203c8ab0146101b257806323268912146101e25780632337db35146101fe575b5f80fd5b610162610498565b60405161016f91906115cb565b60405180910390f35b6101806105d0565b005b61019c6004803603810190610197919061162f565b6105d2565b6040516101a991906116a2565b60405180910390f35b6101cc60048036038101906101c79190611802565b61066d565b6040516101d99190611863565b60405180910390f35b6101fc60048036038101906101f79190611a0a565b6106e4565b005b610206610754565b60405161021391906116a2565b60405180910390f35b6102246107f3565b60405161023191906116a2565b60405180910390f35b610254600480360381019061024f9190611aae565b61087f565b005b610270600480360381019061026b9190611aae565b610a5f565b005b61028c60048036038101906102879190611b24565b610ac5565b6040516102999190611863565b60405180910390f35b6102bc60048036038101906102b7919061162f565b610afa565b6040516102c99190611863565b60405180910390f35b6102ec60048036038101906102e79190611b6b565b610b16565b005b61030860048036038101906103039190611d41565b610b91565b005b610324600480360381019061031f919061162f565b610c57565b6040516103319190611d97565b60405180910390f35b610354600480360381019061034f9190611802565b610c77565b6040516103619190611863565b60405180910390f35b610372610c88565b60405161037f9190611d97565b60405180910390f35b610390610d60565b60405161039d9190611ebd565b60405180910390f35b6103c060048036038101906103bb9190611802565b610e34565b005b6103ca610fd5565b6040516103d791906116a2565b60405180910390f35b6103fa60048036038101906103f5919061162f565b611074565b6040516104079190611863565b60405180910390f35b61042a6004803603810190610425919061162f565b611091565b6040516104379190611f25565b60405180910390f35b610448611137565b005b610464600480360381019061045f919061162f565b61128d565b6040516104719190611863565b60405180910390f35b6104826112aa565b60405161048f9190611ffc565b60405180910390f35b60605f60058054905067ffffffffffffffff8111156104ba576104b96116c6565b5b6040519080825280602002602001820160405280156104ed57816020015b60608152602001906001900390816104d85790505b5090505f5b6005805490508110156105c85760035f8281526020019081526020015f20805461051b90612049565b80601f016020809104026020016040519081016040528092919081815260200182805461054790612049565b80156105925780601f1061056957610100808354040283529160200191610592565b820191905f5260205f20905b81548152906001019060200180831161057557829003601f168201915b50505050508282815181106105aa576105a9612079565b5b602002602001018190525080806105c0906120d3565b9150506104f2565b508091505090565b565b6003602052805f5260405f205f9150905080546105ee90612049565b80601f016020809104026020016040519081016040528092919081815260200182805461061a90612049565b80156106655780601f1061063c57610100808354040283529160200191610665565b820191905f5260205f20905b81548152906001019060200180831161064857829003601f168201915b505050505081565b5f805f90505f5b83518110156106da575f15155f8086848151811061069557610694612079565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff161515036106c757600191506106da565b80806106d2906120d3565b915050610674565b5080915050919050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461073c575f80fd5b61074583610e34565b61074f828261087f565b505050565b606060035f8081526020019081526020015f20805461077290612049565b80601f016020809104026020016040519081016040528092919081815260200182805461079e90612049565b80156107e95780601f106107c0576101008083540402835291602001916107e9565b820191905f5260205f20905b8154815290600101906020018083116107cc57829003601f168201915b5050505050905090565b6004805461080090612049565b80601f016020809104026020016040519081016040528092919081815260200182805461082c90612049565b80156108775780601f1061084e57610100808354040283529160200191610877565b820191905f5260205f20905b81548152906001019060200180831161085a57829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108d7575f80fd5b80518251146108e4575f80fd5b5f5b82518110156109555781818151811061090257610901612079565b5b602002602001015160035f8584815181106109205761091f612079565b5b602002602001015181526020019081526020015f20908161094191906122b7565b50808061094d906120d3565b9150506108e6565b505f5b8251811015610a3a575f151560065f85848151811061097a57610979612079565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff16151503610a2757600160065f8584815181106109bb576109ba612079565b5b602002602001015181526020019081526020015f205f6101000a81548160ff02191690831515021790555060058382815181106109fb576109fa612079565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b8080610a32906120d3565b915050610958565b5060035f8081526020019081526020015f2060049081610a5a91906123ad565b505050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ab7575f80fd5b610ac1828261087f565b5050565b6009818051602081018201805184825260208301602085012081835280955050505050505f915054906101000a900460ff1681565b5f602052805f5260405f205f915054906101000a900460ff1681565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b6e575f80fd5b8060035f8481526020019081526020015f209081610b8c91906122b7565b505050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610be9575f80fd5b5f5b8151811015610c53576008828281518110610c0957610c08612079565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f909190919091509081610c3f91906124ea565b508080610c4b906120d3565b915050610beb565b5050565b60018181548110610c66575f80fd5b905f5260205f20015f915090505481565b5f610c818261066d565b9050919050565b5f60605f805b600580549050811015610d575760035f8281526020019081526020015f208054610cb790612049565b80601f0160208091040260200160405190810160405280929190818152602001828054610ce390612049565b8015610d2e5780601f10610d0557610100808354040283529160200191610d2e565b820191905f5260205f20905b815481529060010190602001808311610d1157829003601f168201915b50505050509250825182610d4291906125b9565b91508080610d4f906120d3565b915050610c8e565b50809250505090565b60606008805480602002602001604051908101604052809291908181526020015f905b82821015610e2b578382905f5260205f20018054610da090612049565b80601f0160208091040260200160405190810160405280929190818152602001828054610dcc90612049565b8015610e175780601f10610dee57610100808354040283529160200191610e17565b820191905f5260205f20905b815481529060010190602001808311610dfa57829003601f168201915b505050505081526020019060010190610d83565b50505050905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e8c575f80fd5b5f5b8151811015610eec5760015f80848481518110610eae57610ead612079565b5b602002602001015181526020019081526020015f205f6101000a81548160ff0219169083151502179055508080610ee4906120d3565b915050610e8e565b505f5b8151811015610fd1575f151560025f848481518110610f1157610f10612079565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff16151503610fbe57600160025f848481518110610f5257610f51612079565b5b602002602001015181526020019081526020015f205f6101000a81548160ff0219169083151502179055506001828281518110610f9257610f91612079565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b8080610fc9906120d3565b915050610eef565b5050565b606060035f8081526020019081526020015f208054610ff390612049565b80601f016020809104026020016040519081016040528092919081815260200182805461101f90612049565b801561106a5780601f106110415761010080835404028352916020019161106a565b820191905f5260205f20905b81548152906001019060200180831161104d57829003601f168201915b5050505050905090565b6002602052805f5260405f205f915054906101000a900460ff1681565b600881815481106110a0575f80fd5b905f5260205f20015f9150905080546110b890612049565b80601f01602080910402602001604051908101604052809291908181526020018280546110e490612049565b801561112f5780601f106111065761010080835404028352916020019161112f565b820191905f5260205f20905b81548152906001019060200180831161111257829003601f168201915b505050505081565b600115156001151503611183576111826040518060400160405280601381526020017f70726974696e67206d65726b6c65207472656500000000000000000000000000815250611300565b5b5f5b60058054905081101561128a57600115156001151503611277576112766040518060400160405280601781526020017f696e646578203a20256420092076616c7565203a2025730000000000000000008152508260035f8581526020019081526020015f2080546111f590612049565b80601f016020809104026020016040519081016040528092919081815260200182805461122190612049565b801561126c5780601f106112435761010080835404028352916020019161126c565b820191905f5260205f20905b81548152906001019060200180831161124f57829003601f168201915b5050505050611399565b5b8080611282906120d3565b915050611185565b50565b6006602052805f5260405f205f915054906101000a900460ff1681565b606060018054806020026020016040519081016040528092919081815260200182805480156112f657602002820191905f5260205f20905b8154815260200190600101908083116112e2575b5050505050905090565b6113968160405160240161131491906116a2565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611438565b50565b6114338383836040516024016113b1939291906125ec565b6040516020818303038152906040527f5970e089000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611438565b505050565b61144f81611447611452611471565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b61147c819050919050565b61148461262f565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156114e65780820151818401526020810190506114cb565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61150b826114af565b61151581856114b9565b93506115258185602086016114c9565b61152e816114f1565b840191505092915050565b5f6115448383611501565b905092915050565b5f602082019050919050565b5f61156282611486565b61156c8185611490565b93508360208202850161157e856114a0565b805f5b858110156115b9578484038952815161159a8582611539565b94506115a58361154c565b925060208a01995050600181019050611581565b50829750879550505050505092915050565b5f6020820190508181035f8301526115e38184611558565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61160e816115fc565b8114611618575f80fd5b50565b5f8135905061162981611605565b92915050565b5f60208284031215611644576116436115f4565b5b5f6116518482850161161b565b91505092915050565b5f82825260208201905092915050565b5f611674826114af565b61167e818561165a565b935061168e8185602086016114c9565b611697816114f1565b840191505092915050565b5f6020820190508181035f8301526116ba818461166a565b905092915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6116fc826114f1565b810181811067ffffffffffffffff8211171561171b5761171a6116c6565b5b80604052505050565b5f61172d6115eb565b905061173982826116f3565b919050565b5f67ffffffffffffffff821115611758576117576116c6565b5b602082029050602081019050919050565b5f80fd5b5f61177f61177a8461173e565b611724565b905080838252602082019050602084028301858111156117a2576117a1611769565b5b835b818110156117cb57806117b7888261161b565b8452602084019350506020810190506117a4565b5050509392505050565b5f82601f8301126117e9576117e86116c2565b5b81356117f984826020860161176d565b91505092915050565b5f60208284031215611817576118166115f4565b5b5f82013567ffffffffffffffff811115611834576118336115f8565b5b611840848285016117d5565b91505092915050565b5f8115159050919050565b61185d81611849565b82525050565b5f6020820190506118765f830184611854565b92915050565b5f67ffffffffffffffff821115611896576118956116c6565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff8211156118c5576118c46116c6565b5b6118ce826114f1565b9050602081019050919050565b828183375f83830152505050565b5f6118fb6118f6846118ab565b611724565b905082815260208101848484011115611917576119166118a7565b5b6119228482856118db565b509392505050565b5f82601f83011261193e5761193d6116c2565b5b813561194e8482602086016118e9565b91505092915050565b5f6119696119648461187c565b611724565b9050808382526020820190506020840283018581111561198c5761198b611769565b5b835b818110156119d357803567ffffffffffffffff8111156119b1576119b06116c2565b5b8086016119be898261192a565b8552602085019450505060208101905061198e565b5050509392505050565b5f82601f8301126119f1576119f06116c2565b5b8135611a01848260208601611957565b91505092915050565b5f805f60608486031215611a2157611a206115f4565b5b5f84013567ffffffffffffffff811115611a3e57611a3d6115f8565b5b611a4a868287016117d5565b935050602084013567ffffffffffffffff811115611a6b57611a6a6115f8565b5b611a77868287016117d5565b925050604084013567ffffffffffffffff811115611a9857611a976115f8565b5b611aa4868287016119dd565b9150509250925092565b5f8060408385031215611ac457611ac36115f4565b5b5f83013567ffffffffffffffff811115611ae157611ae06115f8565b5b611aed858286016117d5565b925050602083013567ffffffffffffffff811115611b0e57611b0d6115f8565b5b611b1a858286016119dd565b9150509250929050565b5f60208284031215611b3957611b386115f4565b5b5f82013567ffffffffffffffff811115611b5657611b556115f8565b5b611b628482850161192a565b91505092915050565b5f8060408385031215611b8157611b806115f4565b5b5f611b8e8582860161161b565b925050602083013567ffffffffffffffff811115611baf57611bae6115f8565b5b611bbb8582860161192a565b9150509250929050565b5f67ffffffffffffffff821115611bdf57611bde6116c6565b5b602082029050602081019050919050565b5f67ffffffffffffffff821115611c0a57611c096116c6565b5b611c13826114f1565b9050602081019050919050565b5f611c32611c2d84611bf0565b611724565b905082815260208101848484011115611c4e57611c4d6118a7565b5b611c598482856118db565b509392505050565b5f82601f830112611c7557611c746116c2565b5b8135611c85848260208601611c20565b91505092915050565b5f611ca0611c9b84611bc5565b611724565b90508083825260208201905060208402830185811115611cc357611cc2611769565b5b835b81811015611d0a57803567ffffffffffffffff811115611ce857611ce76116c2565b5b808601611cf58982611c61565b85526020850194505050602081019050611cc5565b5050509392505050565b5f82601f830112611d2857611d276116c2565b5b8135611d38848260208601611c8e565b91505092915050565b5f60208284031215611d5657611d556115f4565b5b5f82013567ffffffffffffffff811115611d7357611d726115f8565b5b611d7f84828501611d14565b91505092915050565b611d91816115fc565b82525050565b5f602082019050611daa5f830184611d88565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f611dfd82611dd9565b611e078185611de3565b9350611e178185602086016114c9565b611e20816114f1565b840191505092915050565b5f611e368383611df3565b905092915050565b5f602082019050919050565b5f611e5482611db0565b611e5e8185611dba565b935083602082028501611e7085611dca565b805f5b85811015611eab5784840389528151611e8c8582611e2b565b9450611e9783611e3e565b925060208a01995050600181019050611e73565b50829750879550505050505092915050565b5f6020820190508181035f830152611ed58184611e4a565b905092915050565b5f82825260208201905092915050565b5f611ef782611dd9565b611f018185611edd565b9350611f118185602086016114c9565b611f1a816114f1565b840191505092915050565b5f6020820190508181035f830152611f3d8184611eed565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611f77816115fc565b82525050565b5f611f888383611f6e565b60208301905092915050565b5f602082019050919050565b5f611faa82611f45565b611fb48185611f4f565b9350611fbf83611f5f565b805f5b83811015611fef578151611fd68882611f7d565b9750611fe183611f94565b925050600181019050611fc2565b5085935050505092915050565b5f6020820190508181035f8301526120148184611fa0565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061206057607f821691505b6020821081036120735761207261201c565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6120dd826115fc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361210f5761210e6120a6565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026121767fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261213b565b612180868361213b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6121bb6121b66121b1846115fc565b612198565b6115fc565b9050919050565b5f819050919050565b6121d4836121a1565b6121e86121e0826121c2565b848454612147565b825550505050565b5f90565b6121fc6121f0565b6122078184846121cb565b505050565b5b8181101561222a5761221f5f826121f4565b60018101905061220d565b5050565b601f82111561226f576122408161211a565b6122498461212c565b81016020851015612258578190505b61226c6122648561212c565b83018261220c565b50505b505050565b5f82821c905092915050565b5f61228f5f1984600802612274565b1980831691505092915050565b5f6122a78383612280565b9150826002028217905092915050565b6122c0826114af565b67ffffffffffffffff8111156122d9576122d86116c6565b5b6122e38254612049565b6122ee82828561222e565b5f60209050601f83116001811461231f575f841561230d578287015190505b612317858261229c565b86555061237e565b601f19841661232d8661211a565b5f5b828110156123545784890151825560018201915060208501945060208101905061232f565b86831015612371578489015161236d601f891682612280565b8355505b6001600288020188555050505b505050505050565b5f8154905061239481612049565b9050919050565b5f819050815f5260205f209050919050565b8181036123bb575050612490565b6123c482612386565b67ffffffffffffffff8111156123dd576123dc6116c6565b5b6123e78254612049565b6123f282828561222e565b5f601f83116001811461241f575f841561240d578287015490505b612417858261229c565b865550612489565b601f19841661242d8761239b565b96506124388661211a565b5f5b8281101561245f5784890154825560018201915060018501945060208101905061243a565b8683101561247c5784890154612478601f891682612280565b8355505b6001600288020188555050505b5050505050505b565b5f819050815f5260205f209050919050565b601f8211156124e5576124b681612492565b6124bf8461212c565b810160208510156124ce578190505b6124e26124da8561212c565b83018261220c565b50505b505050565b6124f382611dd9565b67ffffffffffffffff81111561250c5761250b6116c6565b5b6125168254612049565b6125218282856124a4565b5f60209050601f831160018114612552575f8415612540578287015190505b61254a858261229c565b8655506125b1565b601f19841661256086612492565b5f5b8281101561258757848901518255600182019150602085019450602081019050612562565b868310156125a457848901516125a0601f891682612280565b8355505b6001600288020188555050505b505050505050565b5f6125c3826115fc565b91506125ce836115fc565b92508282019050808211156125e6576125e56120a6565b5b92915050565b5f6060820190508181035f830152612604818661166a565b90506126136020830185611d88565b8181036040830152612625818461166a565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220c5fa103410cdd246f0b2460483c62dde6435d5167492b12b5a3842f1895a89b464736f6c63430008150033",
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

// GetMerkleTreeSize is a free data retrieval call binding the contract method 0x6e0a5904.
//
// Solidity: function GetMerkleTreeSize() view returns(uint256)
func (_RevocationService *RevocationServiceCaller) GetMerkleTreeSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "GetMerkleTreeSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerkleTreeSize is a free data retrieval call binding the contract method 0x6e0a5904.
//
// Solidity: function GetMerkleTreeSize() view returns(uint256)
func (_RevocationService *RevocationServiceSession) GetMerkleTreeSize() (*big.Int, error) {
	return _RevocationService.Contract.GetMerkleTreeSize(&_RevocationService.CallOpts)
}

// GetMerkleTreeSize is a free data retrieval call binding the contract method 0x6e0a5904.
//
// Solidity: function GetMerkleTreeSize() view returns(uint256)
func (_RevocationService *RevocationServiceCallerSession) GetMerkleTreeSize() (*big.Int, error) {
	return _RevocationService.Contract.GetMerkleTreeSize(&_RevocationService.CallOpts)
}

// RetrieveBloomFilter is a free data retrieval call binding the contract method 0xf174053b.
//
// Solidity: function RetrieveBloomFilter() view returns(uint256[])
func (_RevocationService *RevocationServiceCaller) RetrieveBloomFilter(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "RetrieveBloomFilter")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RetrieveBloomFilter is a free data retrieval call binding the contract method 0xf174053b.
//
// Solidity: function RetrieveBloomFilter() view returns(uint256[])
func (_RevocationService *RevocationServiceSession) RetrieveBloomFilter() ([]*big.Int, error) {
	return _RevocationService.Contract.RetrieveBloomFilter(&_RevocationService.CallOpts)
}

// RetrieveBloomFilter is a free data retrieval call binding the contract method 0xf174053b.
//
// Solidity: function RetrieveBloomFilter() view returns(uint256[])
func (_RevocationService *RevocationServiceCallerSession) RetrieveBloomFilter() ([]*big.Int, error) {
	return _RevocationService.Contract.RetrieveBloomFilter(&_RevocationService.CallOpts)
}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x070f010e.
//
// Solidity: function RetrieveMerkleTree() view returns(string[])
func (_RevocationService *RevocationServiceCaller) RetrieveMerkleTree(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "RetrieveMerkleTree")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x070f010e.
//
// Solidity: function RetrieveMerkleTree() view returns(string[])
func (_RevocationService *RevocationServiceSession) RetrieveMerkleTree() ([]string, error) {
	return _RevocationService.Contract.RetrieveMerkleTree(&_RevocationService.CallOpts)
}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x070f010e.
//
// Solidity: function RetrieveMerkleTree() view returns(string[])
func (_RevocationService *RevocationServiceCallerSession) RetrieveMerkleTree() ([]string, error) {
	return _RevocationService.Contract.RetrieveMerkleTree(&_RevocationService.CallOpts)
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

// BfIndexes is a free data retrieval call binding the contract method 0x5d2c4a38.
//
// Solidity: function bfIndexes(uint256 ) view returns(uint256)
func (_RevocationService *RevocationServiceCaller) BfIndexes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "bfIndexes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BfIndexes is a free data retrieval call binding the contract method 0x5d2c4a38.
//
// Solidity: function bfIndexes(uint256 ) view returns(uint256)
func (_RevocationService *RevocationServiceSession) BfIndexes(arg0 *big.Int) (*big.Int, error) {
	return _RevocationService.Contract.BfIndexes(&_RevocationService.CallOpts, arg0)
}

// BfIndexes is a free data retrieval call binding the contract method 0x5d2c4a38.
//
// Solidity: function bfIndexes(uint256 ) view returns(uint256)
func (_RevocationService *RevocationServiceCallerSession) BfIndexes(arg0 *big.Int) (*big.Int, error) {
	return _RevocationService.Contract.BfIndexes(&_RevocationService.CallOpts, arg0)
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

// IsExistInBloomFilter is a free data retrieval call binding the contract method 0xc2f0b9f1.
//
// Solidity: function isExistInBloomFilter(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceCaller) IsExistInBloomFilter(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "isExistInBloomFilter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistInBloomFilter is a free data retrieval call binding the contract method 0xc2f0b9f1.
//
// Solidity: function isExistInBloomFilter(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceSession) IsExistInBloomFilter(arg0 *big.Int) (bool, error) {
	return _RevocationService.Contract.IsExistInBloomFilter(&_RevocationService.CallOpts, arg0)
}

// IsExistInBloomFilter is a free data retrieval call binding the contract method 0xc2f0b9f1.
//
// Solidity: function isExistInBloomFilter(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) IsExistInBloomFilter(arg0 *big.Int) (bool, error) {
	return _RevocationService.Contract.IsExistInBloomFilter(&_RevocationService.CallOpts, arg0)
}

// IsExistInMTAccumulator is a free data retrieval call binding the contract method 0xe5440a9b.
//
// Solidity: function isExistInMTAccumulator(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceCaller) IsExistInMTAccumulator(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "isExistInMTAccumulator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistInMTAccumulator is a free data retrieval call binding the contract method 0xe5440a9b.
//
// Solidity: function isExistInMTAccumulator(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceSession) IsExistInMTAccumulator(arg0 *big.Int) (bool, error) {
	return _RevocationService.Contract.IsExistInMTAccumulator(&_RevocationService.CallOpts, arg0)
}

// IsExistInMTAccumulator is a free data retrieval call binding the contract method 0xe5440a9b.
//
// Solidity: function isExistInMTAccumulator(uint256 ) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) IsExistInMTAccumulator(arg0 *big.Int) (bool, error) {
	return _RevocationService.Contract.IsExistInMTAccumulator(&_RevocationService.CallOpts, arg0)
}

// IsExistInPublicKeys is a free data retrieval call binding the contract method 0x39288664.
//
// Solidity: function isExistInPublicKeys(string ) view returns(bool)
func (_RevocationService *RevocationServiceCaller) IsExistInPublicKeys(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "isExistInPublicKeys", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistInPublicKeys is a free data retrieval call binding the contract method 0x39288664.
//
// Solidity: function isExistInPublicKeys(string ) view returns(bool)
func (_RevocationService *RevocationServiceSession) IsExistInPublicKeys(arg0 string) (bool, error) {
	return _RevocationService.Contract.IsExistInPublicKeys(&_RevocationService.CallOpts, arg0)
}

// IsExistInPublicKeys is a free data retrieval call binding the contract method 0x39288664.
//
// Solidity: function isExistInPublicKeys(string ) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) IsExistInPublicKeys(arg0 string) (bool, error) {
	return _RevocationService.Contract.IsExistInPublicKeys(&_RevocationService.CallOpts, arg0)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(string)
func (_RevocationService *RevocationServiceCaller) MerkleRoot(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(string)
func (_RevocationService *RevocationServiceSession) MerkleRoot() (string, error) {
	return _RevocationService.Contract.MerkleRoot(&_RevocationService.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(string)
func (_RevocationService *RevocationServiceCallerSession) MerkleRoot() (string, error) {
	return _RevocationService.Contract.MerkleRoot(&_RevocationService.CallOpts)
}

// MerkleTree is a free data retrieval call binding the contract method 0x0df0ff90.
//
// Solidity: function merkleTree(uint256 ) view returns(string)
func (_RevocationService *RevocationServiceCaller) MerkleTree(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "merkleTree", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MerkleTree is a free data retrieval call binding the contract method 0x0df0ff90.
//
// Solidity: function merkleTree(uint256 ) view returns(string)
func (_RevocationService *RevocationServiceSession) MerkleTree(arg0 *big.Int) (string, error) {
	return _RevocationService.Contract.MerkleTree(&_RevocationService.CallOpts, arg0)
}

// MerkleTree is a free data retrieval call binding the contract method 0x0df0ff90.
//
// Solidity: function merkleTree(uint256 ) view returns(string)
func (_RevocationService *RevocationServiceCallerSession) MerkleTree(arg0 *big.Int) (string, error) {
	return _RevocationService.Contract.MerkleTree(&_RevocationService.CallOpts, arg0)
}

// PrintMerkleTree is a free data retrieval call binding the contract method 0xd3e53042.
//
// Solidity: function printMerkleTree() view returns()
func (_RevocationService *RevocationServiceCaller) PrintMerkleTree(opts *bind.CallOpts) error {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "printMerkleTree")

	if err != nil {
		return err
	}

	return err

}

// PrintMerkleTree is a free data retrieval call binding the contract method 0xd3e53042.
//
// Solidity: function printMerkleTree() view returns()
func (_RevocationService *RevocationServiceSession) PrintMerkleTree() error {
	return _RevocationService.Contract.PrintMerkleTree(&_RevocationService.CallOpts)
}

// PrintMerkleTree is a free data retrieval call binding the contract method 0xd3e53042.
//
// Solidity: function printMerkleTree() view returns()
func (_RevocationService *RevocationServiceCallerSession) PrintMerkleTree() error {
	return _RevocationService.Contract.PrintMerkleTree(&_RevocationService.CallOpts)
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
// Solidity: function verificationPhase2() view returns(string)
func (_RevocationService *RevocationServiceCaller) VerificationPhase2(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "verificationPhase2")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VerificationPhase2 is a free data retrieval call binding the contract method 0xbbb7e2ef.
//
// Solidity: function verificationPhase2() view returns(string)
func (_RevocationService *RevocationServiceSession) VerificationPhase2() (string, error) {
	return _RevocationService.Contract.VerificationPhase2(&_RevocationService.CallOpts)
}

// VerificationPhase2 is a free data retrieval call binding the contract method 0xbbb7e2ef.
//
// Solidity: function verificationPhase2() view returns(string)
func (_RevocationService *RevocationServiceCallerSession) VerificationPhase2() (string, error) {
	return _RevocationService.Contract.VerificationPhase2(&_RevocationService.CallOpts)
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

// IssueVC is a paid mutator transaction binding the contract method 0x34d6c77f.
//
// Solidity: function issueVC(uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactor) IssueVC(opts *bind.TransactOpts, _mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "issueVC", _mtIndexes, _mtValues)
}

// IssueVC is a paid mutator transaction binding the contract method 0x34d6c77f.
//
// Solidity: function issueVC(uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceSession) IssueVC(_mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.Contract.IssueVC(&_RevocationService.TransactOpts, _mtIndexes, _mtValues)
}

// IssueVC is a paid mutator transaction binding the contract method 0x34d6c77f.
//
// Solidity: function issueVC(uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactorSession) IssueVC(_mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
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

// RevokeVC is a paid mutator transaction binding the contract method 0x23268912.
//
// Solidity: function revokeVC(uint256[] _bfIndexes, uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactor) RevokeVC(opts *bind.TransactOpts, _bfIndexes []*big.Int, _mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "revokeVC", _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x23268912.
//
// Solidity: function revokeVC(uint256[] _bfIndexes, uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceSession) RevokeVC(_bfIndexes []*big.Int, _mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x23268912.
//
// Solidity: function revokeVC(uint256[] _bfIndexes, uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactorSession) RevokeVC(_bfIndexes []*big.Int, _mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
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

// UpdateMerkleTree is a paid mutator transaction binding the contract method 0x2ef76b18.
//
// Solidity: function updateMerkleTree(uint256[] _indexes, string[] _values) returns()
func (_RevocationService *RevocationServiceTransactor) UpdateMerkleTree(opts *bind.TransactOpts, _indexes []*big.Int, _values []string) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "updateMerkleTree", _indexes, _values)
}

// UpdateMerkleTree is a paid mutator transaction binding the contract method 0x2ef76b18.
//
// Solidity: function updateMerkleTree(uint256[] _indexes, string[] _values) returns()
func (_RevocationService *RevocationServiceSession) UpdateMerkleTree(_indexes []*big.Int, _values []string) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateMerkleTree(&_RevocationService.TransactOpts, _indexes, _values)
}

// UpdateMerkleTree is a paid mutator transaction binding the contract method 0x2ef76b18.
//
// Solidity: function updateMerkleTree(uint256[] _indexes, string[] _values) returns()
func (_RevocationService *RevocationServiceTransactorSession) UpdateMerkleTree(_indexes []*big.Int, _values []string) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateMerkleTree(&_RevocationService.TransactOpts, _indexes, _values)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x53491af5.
//
// Solidity: function updateNode(uint256 index, string value) returns()
func (_RevocationService *RevocationServiceTransactor) UpdateNode(opts *bind.TransactOpts, index *big.Int, value string) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "updateNode", index, value)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x53491af5.
//
// Solidity: function updateNode(uint256 index, string value) returns()
func (_RevocationService *RevocationServiceSession) UpdateNode(index *big.Int, value string) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateNode(&_RevocationService.TransactOpts, index, value)
}

// UpdateNode is a paid mutator transaction binding the contract method 0x53491af5.
//
// Solidity: function updateNode(uint256 index, string value) returns()
func (_RevocationService *RevocationServiceTransactorSession) UpdateNode(index *big.Int, value string) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateNode(&_RevocationService.TransactOpts, index, value)
}

// VerificationPhase2Test is a paid mutator transaction binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() returns(string)
func (_RevocationService *RevocationServiceTransactor) VerificationPhase2Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "verificationPhase2Test")
}

// VerificationPhase2Test is a paid mutator transaction binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() returns(string)
func (_RevocationService *RevocationServiceSession) VerificationPhase2Test() (*types.Transaction, error) {
	return _RevocationService.Contract.VerificationPhase2Test(&_RevocationService.TransactOpts)
}

// VerificationPhase2Test is a paid mutator transaction binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() returns(string)
func (_RevocationService *RevocationServiceTransactorSession) VerificationPhase2Test() (*types.Transaction, error) {
	return _RevocationService.Contract.VerificationPhase2Test(&_RevocationService.TransactOpts)
}

// RevocationServiceIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the RevocationService contract.
type RevocationServiceIssueIterator struct {
	Event *RevocationServiceIssue // Event containing the contract specifics and raw log

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
func (it *RevocationServiceIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevocationServiceIssue)
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
		it.Event = new(RevocationServiceIssue)
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
func (it *RevocationServiceIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevocationServiceIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevocationServiceIssue represents a Issue event raised by the RevocationService contract.
type RevocationServiceIssue struct {
	MtIndexes []*big.Int
	MtValue1  [1]byte
	MtValue2  [1]byte
	MtValue3  [1]byte
	MtValue4  [1]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0x31f3d2d784ee13ea1252b844fa60f0be609e591b43ab4ace2cade6d6d4525f5b.
//
// Solidity: event Issue(uint256[] _mtIndexes, bytes1 _mtValue1, bytes1 _mtValue2, bytes1 _mtValue3, bytes1 _mtValue4)
func (_RevocationService *RevocationServiceFilterer) FilterIssue(opts *bind.FilterOpts) (*RevocationServiceIssueIterator, error) {

	logs, sub, err := _RevocationService.contract.FilterLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return &RevocationServiceIssueIterator{contract: _RevocationService.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0x31f3d2d784ee13ea1252b844fa60f0be609e591b43ab4ace2cade6d6d4525f5b.
//
// Solidity: event Issue(uint256[] _mtIndexes, bytes1 _mtValue1, bytes1 _mtValue2, bytes1 _mtValue3, bytes1 _mtValue4)
func (_RevocationService *RevocationServiceFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *RevocationServiceIssue) (event.Subscription, error) {

	logs, sub, err := _RevocationService.contract.WatchLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevocationServiceIssue)
				if err := _RevocationService.contract.UnpackLog(event, "Issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0x31f3d2d784ee13ea1252b844fa60f0be609e591b43ab4ace2cade6d6d4525f5b.
//
// Solidity: event Issue(uint256[] _mtIndexes, bytes1 _mtValue1, bytes1 _mtValue2, bytes1 _mtValue3, bytes1 _mtValue4)
func (_RevocationService *RevocationServiceFilterer) ParseIssue(log types.Log) (*RevocationServiceIssue, error) {
	event := new(RevocationServiceIssue)
	if err := _RevocationService.contract.UnpackLog(event, "Issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RevocationServiceVerificationPhase2Iterator is returned from FilterVerificationPhase2 and is used to iterate over the raw logs and unpacked data for VerificationPhase2 events raised by the RevocationService contract.
type RevocationServiceVerificationPhase2Iterator struct {
	Event *RevocationServiceVerificationPhase2 // Event containing the contract specifics and raw log

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
func (it *RevocationServiceVerificationPhase2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RevocationServiceVerificationPhase2)
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
		it.Event = new(RevocationServiceVerificationPhase2)
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
func (it *RevocationServiceVerificationPhase2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RevocationServiceVerificationPhase2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RevocationServiceVerificationPhase2 represents a VerificationPhase2 event raised by the RevocationService contract.
type RevocationServiceVerificationPhase2 struct {
	MerkleRoot [32]byte
	VcLeaf     [32]byte
	Proof      [][32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerificationPhase2 is a free log retrieval operation binding the contract event 0x518b0f6d8d08ac9bfd6000c702f772cb73ca7c04cbb7dc2421a005c821a2dcc9.
//
// Solidity: event VerificationPhase2(bytes32 merkleRoot, bytes32 vcLeaf, bytes32[] proof)
func (_RevocationService *RevocationServiceFilterer) FilterVerificationPhase2(opts *bind.FilterOpts) (*RevocationServiceVerificationPhase2Iterator, error) {

	logs, sub, err := _RevocationService.contract.FilterLogs(opts, "VerificationPhase2")
	if err != nil {
		return nil, err
	}
	return &RevocationServiceVerificationPhase2Iterator{contract: _RevocationService.contract, event: "VerificationPhase2", logs: logs, sub: sub}, nil
}

// WatchVerificationPhase2 is a free log subscription operation binding the contract event 0x518b0f6d8d08ac9bfd6000c702f772cb73ca7c04cbb7dc2421a005c821a2dcc9.
//
// Solidity: event VerificationPhase2(bytes32 merkleRoot, bytes32 vcLeaf, bytes32[] proof)
func (_RevocationService *RevocationServiceFilterer) WatchVerificationPhase2(opts *bind.WatchOpts, sink chan<- *RevocationServiceVerificationPhase2) (event.Subscription, error) {

	logs, sub, err := _RevocationService.contract.WatchLogs(opts, "VerificationPhase2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RevocationServiceVerificationPhase2)
				if err := _RevocationService.contract.UnpackLog(event, "VerificationPhase2", log); err != nil {
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

// ParseVerificationPhase2 is a log parse operation binding the contract event 0x518b0f6d8d08ac9bfd6000c702f772cb73ca7c04cbb7dc2421a005c821a2dcc9.
//
// Solidity: event VerificationPhase2(bytes32 merkleRoot, bytes32 vcLeaf, bytes32[] proof)
func (_RevocationService *RevocationServiceFilterer) ParseVerificationPhase2(log types.Log) (*RevocationServiceVerificationPhase2, error) {
	event := new(RevocationServiceVerificationPhase2)
	if err := _RevocationService.contract.UnpackLog(event, "VerificationPhase2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
