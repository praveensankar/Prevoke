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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RetrieveBloomFilter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveMerkleTree\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrievePublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"addPublicKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bfIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInMTAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"isExistInPublicKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfHashFunctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_values\",\"type\":\"string[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506126628061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610156575f3560e01c806353491af5116100c1578063c2f0b9f11161007a578063c2f0b9f1146103b0578063c680f410146103e0578063cc70dd1514610410578063d3e5304214610440578063e5440a9b1461044a578063f174053b1461047a57610156565b806353491af5146102f05780635a34497a1461030c5780635d2c4a381461032857806399223a5a14610358578063a5eb0de814610376578063bbb7e2ef1461039257610156565b80632ef76b18116101135780632ef76b181461020a578063309ddb301461022657806334d6c77f14610256578063376a659014610272578063392886641461029057806348db5336146102c057610156565b8063070f010e1461015a5780630de54b85146101785780630df0ff901461018257806323268912146101b25780632337db35146101ce5780632eb4a7ab146101ec575b5f80fd5b610162610498565b60405161016f91906114f5565b60405180910390f35b6101806105d0565b005b61019c60048036038101906101979190611559565b6105d2565b6040516101a991906115cc565b60405180910390f35b6101cc60048036038101906101c791906118ba565b61066d565b005b6101d66106dd565b6040516101e391906115cc565b60405180910390f35b6101f461077c565b60405161020191906115cc565b60405180910390f35b610224600480360381019061021f919061195e565b610808565b005b610240600480360381019061023b9190611a82565b6109e8565b60405161024d9190611ac7565b60405180910390f35b610270600480360381019061026b919061195e565b6109f9565b005b61027a610a5f565b6040516102879190611aef565b60405180910390f35b6102aa60048036038101906102a59190611b08565b610a64565b6040516102b79190611ac7565b60405180910390f35b6102da60048036038101906102d59190611559565b610a99565b6040516102e79190611ac7565b60405180910390f35b61030a60048036038101906103059190611b4f565b610ab5565b005b61032660048036038101906103219190611d25565b610b30565b005b610342600480360381019061033d9190611559565b610bf6565b60405161034f9190611aef565b60405180910390f35b610360610c16565b60405161036d9190611e79565b60405180910390f35b610390600480360381019061038b9190611e99565b610cea565b005b61039a610e8b565b6040516103a791906115cc565b60405180910390f35b6103ca60048036038101906103c59190611559565b610f2a565b6040516103d79190611ac7565b60405180910390f35b6103fa60048036038101906103f59190611559565b610f47565b6040516104079190611f28565b60405180910390f35b61042a60048036038101906104259190611a82565b610fed565b6040516104379190611ac7565b60405180910390f35b610448611061565b005b610464600480360381019061045f9190611559565b6111b7565b6040516104719190611ac7565b60405180910390f35b6104826111d4565b60405161048f9190611fff565b60405180910390f35b60605f60058054905067ffffffffffffffff8111156104ba576104b96115f0565b5b6040519080825280602002602001820160405280156104ed57816020015b60608152602001906001900390816104d85790505b5090505f5b6005805490508110156105c85760035f8281526020019081526020015f20805461051b9061204c565b80601f01602080910402602001604051908101604052809291908181526020018280546105479061204c565b80156105925780601f1061056957610100808354040283529160200191610592565b820191905f5260205f20905b81548152906001019060200180831161057557829003601f168201915b50505050508282815181106105aa576105a961207c565b5b602002602001018190525080806105c0906120d6565b9150506104f2565b508091505090565b565b6003602052805f5260405f205f9150905080546105ee9061204c565b80601f016020809104026020016040519081016040528092919081815260200182805461061a9061204c565b80156106655780601f1061063c57610100808354040283529160200191610665565b820191905f5260205f20905b81548152906001019060200180831161064857829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c5575f80fd5b6106ce83610cea565b6106d88282610808565b505050565b606060035f8081526020019081526020015f2080546106fb9061204c565b80601f01602080910402602001604051908101604052809291908181526020018280546107279061204c565b80156107725780601f1061074957610100808354040283529160200191610772565b820191905f5260205f20905b81548152906001019060200180831161075557829003601f168201915b5050505050905090565b600480546107899061204c565b80601f01602080910402602001604051908101604052809291908181526020018280546107b59061204c565b80156108005780601f106107d757610100808354040283529160200191610800565b820191905f5260205f20905b8154815290600101906020018083116107e357829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610860575f80fd5b805182511461086d575f80fd5b5f5b82518110156108de5781818151811061088b5761088a61207c565b5b602002602001015160035f8584815181106108a9576108a861207c565b5b602002602001015181526020019081526020015f2090816108ca91906122ba565b5080806108d6906120d6565b91505061086f565b505f5b82518110156109c3575f151560065f8584815181106109035761090261207c565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff161515036109b057600160065f8584815181106109445761094361207c565b5b602002602001015181526020019081526020015f205f6101000a81548160ff02191690831515021790555060058382815181106109845761098361207c565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b80806109bb906120d6565b9150506108e1565b5060035f8081526020019081526020015f20600490816109e391906123b0565b505050565b5f6109f282610fed565b9050919050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a51575f80fd5b610a5b8282610808565b5050565b600481565b6009818051602081018201805184825260208301602085012081835280955050505050505f915054906101000a900460ff1681565b5f602052805f5260405f205f915054906101000a900460ff1681565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b0d575f80fd5b8060035f8481526020019081526020015f209081610b2b91906122ba565b505050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b88575f80fd5b5f5b8151811015610bf2576008828281518110610ba857610ba761207c565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f909190919091509081610bde91906124ed565b508080610bea906120d6565b915050610b8a565b5050565b60018181548110610c05575f80fd5b905f5260205f20015f915090505481565b60606008805480602002602001604051908101604052809291908181526020015f905b82821015610ce1578382905f5260205f20018054610c569061204c565b80601f0160208091040260200160405190810160405280929190818152602001828054610c829061204c565b8015610ccd5780601f10610ca457610100808354040283529160200191610ccd565b820191905f5260205f20905b815481529060010190602001808311610cb057829003601f168201915b505050505081526020019060010190610c39565b50505050905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d42575f80fd5b5f5b8151811015610da25760015f80848481518110610d6457610d6361207c565b5b602002602001015181526020019081526020015f205f6101000a81548160ff0219169083151502179055508080610d9a906120d6565b915050610d44565b505f5b8151811015610e87575f151560025f848481518110610dc757610dc661207c565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff16151503610e7457600160025f848481518110610e0857610e0761207c565b5b602002602001015181526020019081526020015f205f6101000a81548160ff0219169083151502179055506001828281518110610e4857610e4761207c565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b8080610e7f906120d6565b915050610da5565b5050565b606060035f8081526020019081526020015f208054610ea99061204c565b80601f0160208091040260200160405190810160405280929190818152602001828054610ed59061204c565b8015610f205780601f10610ef757610100808354040283529160200191610f20565b820191905f5260205f20905b815481529060010190602001808311610f0357829003601f168201915b5050505050905090565b6002602052805f5260405f205f915054906101000a900460ff1681565b60088181548110610f56575f80fd5b905f5260205f20015f915090508054610f6e9061204c565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9a9061204c565b8015610fe55780601f10610fbc57610100808354040283529160200191610fe5565b820191905f5260205f20905b815481529060010190602001808311610fc857829003601f168201915b505050505081565b5f805f90505f5b6004811015611057575f15155f808684600481106110155761101461207c565b5b602002015181526020019081526020015f205f9054906101000a900460ff161515036110445760019150611057565b808061104f906120d6565b915050610ff4565b5080915050919050565b6001151560011515036110ad576110ac6040518060400160405280601381526020017f70726974696e67206d65726b6c6520747265650000000000000000000000000081525061122a565b5b5f5b6005805490508110156111b4576001151560011515036111a1576111a06040518060400160405280601781526020017f696e646578203a20256420092076616c7565203a2025730000000000000000008152508260035f8581526020019081526020015f20805461111f9061204c565b80601f016020809104026020016040519081016040528092919081815260200182805461114b9061204c565b80156111965780601f1061116d57610100808354040283529160200191611196565b820191905f5260205f20905b81548152906001019060200180831161117957829003601f168201915b50505050506112c3565b5b80806111ac906120d6565b9150506110af565b50565b6006602052805f5260405f205f915054906101000a900460ff1681565b6060600180548060200260200160405190810160405280929190818152602001828054801561122057602002820191905f5260205f20905b81548152602001906001019080831161120c575b5050505050905090565b6112c08160405160240161123e91906115cc565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611362565b50565b61135d8383836040516024016112db939291906125bc565b6040516020818303038152906040527f5970e089000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611362565b505050565b6113798161137161137c61139b565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b6113a6819050919050565b6113ae6125ff565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156114105780820151818401526020810190506113f5565b5f8484015250505050565b5f601f19601f8301169050919050565b5f611435826113d9565b61143f81856113e3565b935061144f8185602086016113f3565b6114588161141b565b840191505092915050565b5f61146e838361142b565b905092915050565b5f602082019050919050565b5f61148c826113b0565b61149681856113ba565b9350836020820285016114a8856113ca565b805f5b858110156114e357848403895281516114c48582611463565b94506114cf83611476565b925060208a019950506001810190506114ab565b50829750879550505050505092915050565b5f6020820190508181035f83015261150d8184611482565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61153881611526565b8114611542575f80fd5b50565b5f813590506115538161152f565b92915050565b5f6020828403121561156e5761156d61151e565b5b5f61157b84828501611545565b91505092915050565b5f82825260208201905092915050565b5f61159e826113d9565b6115a88185611584565b93506115b88185602086016113f3565b6115c18161141b565b840191505092915050565b5f6020820190508181035f8301526115e48184611594565b905092915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6116268261141b565b810181811067ffffffffffffffff82111715611645576116446115f0565b5b80604052505050565b5f611657611515565b9050611663828261161d565b919050565b5f67ffffffffffffffff821115611682576116816115f0565b5b602082029050602081019050919050565b5f80fd5b5f6116a96116a484611668565b61164e565b905080838252602082019050602084028301858111156116cc576116cb611693565b5b835b818110156116f557806116e18882611545565b8452602084019350506020810190506116ce565b5050509392505050565b5f82601f830112611713576117126115ec565b5b8135611723848260208601611697565b91505092915050565b5f67ffffffffffffffff821115611746576117456115f0565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff821115611775576117746115f0565b5b61177e8261141b565b9050602081019050919050565b828183375f83830152505050565b5f6117ab6117a68461175b565b61164e565b9050828152602081018484840111156117c7576117c6611757565b5b6117d284828561178b565b509392505050565b5f82601f8301126117ee576117ed6115ec565b5b81356117fe848260208601611799565b91505092915050565b5f6118196118148461172c565b61164e565b9050808382526020820190506020840283018581111561183c5761183b611693565b5b835b8181101561188357803567ffffffffffffffff811115611861576118606115ec565b5b80860161186e89826117da565b8552602085019450505060208101905061183e565b5050509392505050565b5f82601f8301126118a1576118a06115ec565b5b81356118b1848260208601611807565b91505092915050565b5f805f606084860312156118d1576118d061151e565b5b5f84013567ffffffffffffffff8111156118ee576118ed611522565b5b6118fa868287016116ff565b935050602084013567ffffffffffffffff81111561191b5761191a611522565b5b611927868287016116ff565b925050604084013567ffffffffffffffff81111561194857611947611522565b5b6119548682870161188d565b9150509250925092565b5f80604083850312156119745761197361151e565b5b5f83013567ffffffffffffffff81111561199157611990611522565b5b61199d858286016116ff565b925050602083013567ffffffffffffffff8111156119be576119bd611522565b5b6119ca8582860161188d565b9150509250929050565b5f67ffffffffffffffff8211156119ee576119ed6115f0565b5b602082029050919050565b5f611a0b611a06846119d4565b61164e565b90508060208402830185811115611a2557611a24611693565b5b835b81811015611a4e5780611a3a8882611545565b845260208401935050602081019050611a27565b5050509392505050565b5f82601f830112611a6c57611a6b6115ec565b5b6004611a798482856119f9565b91505092915050565b5f60808284031215611a9757611a9661151e565b5b5f611aa484828501611a58565b91505092915050565b5f8115159050919050565b611ac181611aad565b82525050565b5f602082019050611ada5f830184611ab8565b92915050565b611ae981611526565b82525050565b5f602082019050611b025f830184611ae0565b92915050565b5f60208284031215611b1d57611b1c61151e565b5b5f82013567ffffffffffffffff811115611b3a57611b39611522565b5b611b46848285016117da565b91505092915050565b5f8060408385031215611b6557611b6461151e565b5b5f611b7285828601611545565b925050602083013567ffffffffffffffff811115611b9357611b92611522565b5b611b9f858286016117da565b9150509250929050565b5f67ffffffffffffffff821115611bc357611bc26115f0565b5b602082029050602081019050919050565b5f67ffffffffffffffff821115611bee57611bed6115f0565b5b611bf78261141b565b9050602081019050919050565b5f611c16611c1184611bd4565b61164e565b905082815260208101848484011115611c3257611c31611757565b5b611c3d84828561178b565b509392505050565b5f82601f830112611c5957611c586115ec565b5b8135611c69848260208601611c04565b91505092915050565b5f611c84611c7f84611ba9565b61164e565b90508083825260208201905060208402830185811115611ca757611ca6611693565b5b835b81811015611cee57803567ffffffffffffffff811115611ccc57611ccb6115ec565b5b808601611cd98982611c45565b85526020850194505050602081019050611ca9565b5050509392505050565b5f82601f830112611d0c57611d0b6115ec565b5b8135611d1c848260208601611c72565b91505092915050565b5f60208284031215611d3a57611d3961151e565b5b5f82013567ffffffffffffffff811115611d5757611d56611522565b5b611d6384828501611cf8565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f611db982611d95565b611dc38185611d9f565b9350611dd38185602086016113f3565b611ddc8161141b565b840191505092915050565b5f611df28383611daf565b905092915050565b5f602082019050919050565b5f611e1082611d6c565b611e1a8185611d76565b935083602082028501611e2c85611d86565b805f5b85811015611e675784840389528151611e488582611de7565b9450611e5383611dfa565b925060208a01995050600181019050611e2f565b50829750879550505050505092915050565b5f6020820190508181035f830152611e918184611e06565b905092915050565b5f60208284031215611eae57611ead61151e565b5b5f82013567ffffffffffffffff811115611ecb57611eca611522565b5b611ed7848285016116ff565b91505092915050565b5f82825260208201905092915050565b5f611efa82611d95565b611f048185611ee0565b9350611f148185602086016113f3565b611f1d8161141b565b840191505092915050565b5f6020820190508181035f830152611f408184611ef0565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611f7a81611526565b82525050565b5f611f8b8383611f71565b60208301905092915050565b5f602082019050919050565b5f611fad82611f48565b611fb78185611f52565b9350611fc283611f62565b805f5b83811015611ff2578151611fd98882611f80565b9750611fe483611f97565b925050600181019050611fc5565b5085935050505092915050565b5f6020820190508181035f8301526120178184611fa3565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061206357607f821691505b6020821081036120765761207561201f565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6120e082611526565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612112576121116120a9565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026121797fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261213e565b612183868361213e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6121be6121b96121b484611526565b61219b565b611526565b9050919050565b5f819050919050565b6121d7836121a4565b6121eb6121e3826121c5565b84845461214a565b825550505050565b5f90565b6121ff6121f3565b61220a8184846121ce565b505050565b5b8181101561222d576122225f826121f7565b600181019050612210565b5050565b601f821115612272576122438161211d565b61224c8461212f565b8101602085101561225b578190505b61226f6122678561212f565b83018261220f565b50505b505050565b5f82821c905092915050565b5f6122925f1984600802612277565b1980831691505092915050565b5f6122aa8383612283565b9150826002028217905092915050565b6122c3826113d9565b67ffffffffffffffff8111156122dc576122db6115f0565b5b6122e6825461204c565b6122f1828285612231565b5f60209050601f831160018114612322575f8415612310578287015190505b61231a858261229f565b865550612381565b601f1984166123308661211d565b5f5b8281101561235757848901518255600182019150602085019450602081019050612332565b868310156123745784890151612370601f891682612283565b8355505b6001600288020188555050505b505050505050565b5f815490506123978161204c565b9050919050565b5f819050815f5260205f209050919050565b8181036123be575050612493565b6123c782612389565b67ffffffffffffffff8111156123e0576123df6115f0565b5b6123ea825461204c565b6123f5828285612231565b5f601f831160018114612422575f8415612410578287015490505b61241a858261229f565b86555061248c565b601f1984166124308761239e565b965061243b8661211d565b5f5b828110156124625784890154825560018201915060018501945060208101905061243d565b8683101561247f578489015461247b601f891682612283565b8355505b6001600288020188555050505b5050505050505b565b5f819050815f5260205f209050919050565b601f8211156124e8576124b981612495565b6124c28461212f565b810160208510156124d1578190505b6124e56124dd8561212f565b83018261220f565b50505b505050565b6124f682611d95565b67ffffffffffffffff81111561250f5761250e6115f0565b5b612519825461204c565b6125248282856124a7565b5f60209050601f831160018114612555575f8415612543578287015190505b61254d858261229f565b8655506125b4565b601f19841661256386612495565b5f5b8281101561258a57848901518255600182019150602085019450602081019050612565565b868310156125a757848901516125a3601f891682612283565b8355505b6001600288020188555050505b505050505050565b5f6060820190508181035f8301526125d48186611594565b90506125e36020830185611ae0565b81810360408301526125f58184611594565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220b753e8101a3b7dd14ddd2cdf54887f8f15d28dbc2706ac693b6b193ed24ab51e64736f6c63430008150033",
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

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xcc70dd15.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[4] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceCaller) CheckRevocationStatusInBloomFilter(opts *bind.CallOpts, _indexes [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "checkRevocationStatusInBloomFilter", _indexes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xcc70dd15.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[4] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceSession) CheckRevocationStatusInBloomFilter(_indexes [4]*big.Int) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInBloomFilter(&_RevocationService.CallOpts, _indexes)
}

// CheckRevocationStatusInBloomFilter is a free data retrieval call binding the contract method 0xcc70dd15.
//
// Solidity: function checkRevocationStatusInBloomFilter(uint256[4] _indexes) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) CheckRevocationStatusInBloomFilter(_indexes [4]*big.Int) (bool, error) {
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

// NumberOfHashFunctions is a free data retrieval call binding the contract method 0x376a6590.
//
// Solidity: function numberOfHashFunctions() view returns(uint256)
func (_RevocationService *RevocationServiceCaller) NumberOfHashFunctions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "numberOfHashFunctions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfHashFunctions is a free data retrieval call binding the contract method 0x376a6590.
//
// Solidity: function numberOfHashFunctions() view returns(uint256)
func (_RevocationService *RevocationServiceSession) NumberOfHashFunctions() (*big.Int, error) {
	return _RevocationService.Contract.NumberOfHashFunctions(&_RevocationService.CallOpts)
}

// NumberOfHashFunctions is a free data retrieval call binding the contract method 0x376a6590.
//
// Solidity: function numberOfHashFunctions() view returns(uint256)
func (_RevocationService *RevocationServiceCallerSession) NumberOfHashFunctions() (*big.Int, error) {
	return _RevocationService.Contract.NumberOfHashFunctions(&_RevocationService.CallOpts)
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

// VerificationPhase1 is a free data retrieval call binding the contract method 0x309ddb30.
//
// Solidity: function verificationPhase1(uint256[4] _bfIndexes) view returns(bool)
func (_RevocationService *RevocationServiceCaller) VerificationPhase1(opts *bind.CallOpts, _bfIndexes [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "verificationPhase1", _bfIndexes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificationPhase1 is a free data retrieval call binding the contract method 0x309ddb30.
//
// Solidity: function verificationPhase1(uint256[4] _bfIndexes) view returns(bool)
func (_RevocationService *RevocationServiceSession) VerificationPhase1(_bfIndexes [4]*big.Int) (bool, error) {
	return _RevocationService.Contract.VerificationPhase1(&_RevocationService.CallOpts, _bfIndexes)
}

// VerificationPhase1 is a free data retrieval call binding the contract method 0x309ddb30.
//
// Solidity: function verificationPhase1(uint256[4] _bfIndexes) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) VerificationPhase1(_bfIndexes [4]*big.Int) (bool, error) {
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
