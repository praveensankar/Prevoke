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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RetrieveBloomFilter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveMerkleTree\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrievePublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"addPublicKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bfIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInMTAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"isExistInPublicKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfHashFunctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_values\",\"type\":\"string[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506125f38061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610156575f3560e01c806353491af5116100c1578063c2f0b9f11161007a578063c2f0b9f1146103b0578063c680f410146103e0578063cc70dd1514610410578063d3e5304214610440578063e5440a9b1461044a578063f174053b1461047a57610156565b806353491af5146102f05780635a34497a1461030c5780635d2c4a381461032857806392ed770e1461035857806399223a5a14610374578063bbb7e2ef1461039257610156565b80632ef76b18116101135780632ef76b181461020a578063309ddb301461022657806334d6c77f14610256578063376a659014610272578063392886641461029057806348db5336146102c057610156565b8063070f010e1461015a5780630de54b85146101785780630df0ff901461018257806317a30492146101b25780632337db35146101ce5780632eb4a7ab146101ec575b5f80fd5b610162610498565b60405161016f91906114e9565b60405180910390f35b6101806105d0565b005b61019c6004803603810190610197919061154d565b6105d2565b6040516101a991906115c0565b60405180910390f35b6101cc60048036038101906101c7919061170e565b61066d565b005b6101d6610802565b6040516101e391906115c0565b60405180910390f35b6101f46108a1565b60405161020191906115c0565b60405180910390f35b610224600480360381019061021f9190611987565b61092d565b005b610240600480360381019061023b919061170e565b610b0d565b60405161024d9190611a17565b60405180910390f35b610270600480360381019061026b9190611987565b610b1e565b005b61027a610b84565b6040516102879190611a3f565b60405180910390f35b6102aa60048036038101906102a59190611a58565b610b89565b6040516102b79190611a17565b60405180910390f35b6102da60048036038101906102d5919061154d565b610bbe565b6040516102e79190611a17565b60405180910390f35b61030a60048036038101906103059190611a9f565b610bda565b005b61032660048036038101906103219190611c75565b610c55565b005b610342600480360381019061033d919061154d565b610d1b565b60405161034f9190611a3f565b60405180910390f35b610372600480360381019061036d9190611cbc565b610d3b565b005b61037c610dab565b6040516103899190611e51565b60405180910390f35b61039a610e7f565b6040516103a791906115c0565b60405180910390f35b6103ca60048036038101906103c5919061154d565b610f1e565b6040516103d79190611a17565b60405180910390f35b6103fa60048036038101906103f5919061154d565b610f3b565b6040516104079190611eb9565b60405180910390f35b61042a6004803603810190610425919061170e565b610fe1565b6040516104379190611a17565b60405180910390f35b610448611055565b005b610464600480360381019061045f919061154d565b6111ab565b6040516104719190611a17565b60405180910390f35b6104826111c8565b60405161048f9190611f90565b60405180910390f35b60605f60058054905067ffffffffffffffff8111156104ba576104b96115e4565b5b6040519080825280602002602001820160405280156104ed57816020015b60608152602001906001900390816104d85790505b5090505f5b6005805490508110156105c85760035f8281526020019081526020015f20805461051b90611fdd565b80601f016020809104026020016040519081016040528092919081815260200182805461054790611fdd565b80156105925780601f1061056957610100808354040283529160200191610592565b820191905f5260205f20905b81548152906001019060200180831161057557829003601f168201915b50505050508282815181106105aa576105a961200d565b5b602002602001018190525080806105c090612067565b9150506104f2565b508091505090565b565b6003602052805f5260405f205f9150905080546105ee90611fdd565b80601f016020809104026020016040519081016040528092919081815260200182805461061a90611fdd565b80156106655780601f1061063c57610100808354040283529160200191610665565b820191905f5260205f20905b81548152906001019060200180831161064857829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106c5575f80fd5b5f5b60048110156107225760015f808484600481106106e7576106e661200d565b5b602002015181526020019081526020015f205f6101000a81548160ff021916908315150217905550808061071a90612067565b9150506106c7565b505f5b60048110156107fe575f151560025f8484600481106107475761074661200d565b5b602002015181526020019081526020015f205f9054906101000a900460ff161515036107eb57600160025f8484600481106107855761078461200d565b5b602002015181526020019081526020015f205f6101000a81548160ff02191690831515021790555060018282600481106107c2576107c161200d565b5b6020020151908060018154018082558091505060019003905f5260205f20015f90919091909150555b80806107f690612067565b915050610725565b5050565b606060035f8081526020019081526020015f20805461082090611fdd565b80601f016020809104026020016040519081016040528092919081815260200182805461084c90611fdd565b80156108975780601f1061086e57610100808354040283529160200191610897565b820191905f5260205f20905b81548152906001019060200180831161087a57829003601f168201915b5050505050905090565b600480546108ae90611fdd565b80601f01602080910402602001604051908101604052809291908181526020018280546108da90611fdd565b80156109255780601f106108fc57610100808354040283529160200191610925565b820191905f5260205f20905b81548152906001019060200180831161090857829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610985575f80fd5b8051825114610992575f80fd5b5f5b8251811015610a03578181815181106109b0576109af61200d565b5b602002602001015160035f8584815181106109ce576109cd61200d565b5b602002602001015181526020019081526020015f2090816109ef919061224b565b5080806109fb90612067565b915050610994565b505f5b8251811015610ae8575f151560065f858481518110610a2857610a2761200d565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff16151503610ad557600160065f858481518110610a6957610a6861200d565b5b602002602001015181526020019081526020015f205f6101000a81548160ff0219169083151502179055506005838281518110610aa957610aa861200d565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b8080610ae090612067565b915050610a06565b5060035f8081526020019081526020015f2060049081610b089190612341565b505050565b5f610b1782610fe1565b9050919050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b76575f80fd5b610b80828261092d565b5050565b600481565b6009818051602081018201805184825260208301602085012081835280955050505050505f915054906101000a900460ff1681565b5f602052805f5260405f205f915054906101000a900460ff1681565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c32575f80fd5b8060035f8481526020019081526020015f209081610c50919061224b565b505050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cad575f80fd5b5f5b8151811015610d17576008828281518110610ccd57610ccc61200d565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f909190919091509081610d03919061247e565b508080610d0f90612067565b915050610caf565b5050565b60018181548110610d2a575f80fd5b905f5260205f20015f915090505481565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d93575f80fd5b610d9c8361066d565b610da6828261092d565b505050565b60606008805480602002602001604051908101604052809291908181526020015f905b82821015610e76578382905f5260205f20018054610deb90611fdd565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1790611fdd565b8015610e625780601f10610e3957610100808354040283529160200191610e62565b820191905f5260205f20905b815481529060010190602001808311610e4557829003601f168201915b505050505081526020019060010190610dce565b50505050905090565b606060035f8081526020019081526020015f208054610e9d90611fdd565b80601f0160208091040260200160405190810160405280929190818152602001828054610ec990611fdd565b8015610f145780601f10610eeb57610100808354040283529160200191610f14565b820191905f5260205f20905b815481529060010190602001808311610ef757829003601f168201915b5050505050905090565b6002602052805f5260405f205f915054906101000a900460ff1681565b60088181548110610f4a575f80fd5b905f5260205f20015f915090508054610f6290611fdd565b80601f0160208091040260200160405190810160405280929190818152602001828054610f8e90611fdd565b8015610fd95780601f10610fb057610100808354040283529160200191610fd9565b820191905f5260205f20905b815481529060010190602001808311610fbc57829003601f168201915b505050505081565b5f805f90505f5b600481101561104b575f15155f808684600481106110095761100861200d565b5b602002015181526020019081526020015f205f9054906101000a900460ff16151503611038576001915061104b565b808061104390612067565b915050610fe8565b5080915050919050565b6001151560011515036110a1576110a06040518060400160405280601381526020017f70726974696e67206d65726b6c6520747265650000000000000000000000000081525061121e565b5b5f5b6005805490508110156111a857600115156001151503611195576111946040518060400160405280601781526020017f696e646578203a20256420092076616c7565203a2025730000000000000000008152508260035f8581526020019081526020015f20805461111390611fdd565b80601f016020809104026020016040519081016040528092919081815260200182805461113f90611fdd565b801561118a5780601f106111615761010080835404028352916020019161118a565b820191905f5260205f20905b81548152906001019060200180831161116d57829003601f168201915b50505050506112b7565b5b80806111a090612067565b9150506110a3565b50565b6006602052805f5260405f205f915054906101000a900460ff1681565b6060600180548060200260200160405190810160405280929190818152602001828054801561121457602002820191905f5260205f20905b815481526020019060010190808311611200575b5050505050905090565b6112b48160405160240161123291906115c0565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611356565b50565b6113518383836040516024016112cf9392919061254d565b6040516020818303038152906040527f5970e089000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611356565b505050565b61136d8161136561137061138f565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b61139a819050919050565b6113a2612590565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156114045780820151818401526020810190506113e9565b5f8484015250505050565b5f601f19601f8301169050919050565b5f611429826113cd565b61143381856113d7565b93506114438185602086016113e7565b61144c8161140f565b840191505092915050565b5f611462838361141f565b905092915050565b5f602082019050919050565b5f611480826113a4565b61148a81856113ae565b93508360208202850161149c856113be565b805f5b858110156114d757848403895281516114b88582611457565b94506114c38361146a565b925060208a0199505060018101905061149f565b50829750879550505050505092915050565b5f6020820190508181035f8301526115018184611476565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61152c8161151a565b8114611536575f80fd5b50565b5f8135905061154781611523565b92915050565b5f6020828403121561156257611561611512565b5b5f61156f84828501611539565b91505092915050565b5f82825260208201905092915050565b5f611592826113cd565b61159c8185611578565b93506115ac8185602086016113e7565b6115b58161140f565b840191505092915050565b5f6020820190508181035f8301526115d88184611588565b905092915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61161a8261140f565b810181811067ffffffffffffffff82111715611639576116386115e4565b5b80604052505050565b5f61164b611509565b90506116578282611611565b919050565b5f67ffffffffffffffff821115611676576116756115e4565b5b602082029050919050565b5f80fd5b5f6116976116928461165c565b611642565b905080602084028301858111156116b1576116b0611681565b5b835b818110156116da57806116c68882611539565b8452602084019350506020810190506116b3565b5050509392505050565b5f82601f8301126116f8576116f76115e0565b5b6004611705848285611685565b91505092915050565b5f6080828403121561172357611722611512565b5b5f611730848285016116e4565b91505092915050565b5f67ffffffffffffffff821115611753576117526115e4565b5b602082029050602081019050919050565b5f61177661177184611739565b611642565b9050808382526020820190506020840283018581111561179957611798611681565b5b835b818110156117c257806117ae8882611539565b84526020840193505060208101905061179b565b5050509392505050565b5f82601f8301126117e0576117df6115e0565b5b81356117f0848260208601611764565b91505092915050565b5f67ffffffffffffffff821115611813576118126115e4565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff821115611842576118416115e4565b5b61184b8261140f565b9050602081019050919050565b828183375f83830152505050565b5f61187861187384611828565b611642565b90508281526020810184848401111561189457611893611824565b5b61189f848285611858565b509392505050565b5f82601f8301126118bb576118ba6115e0565b5b81356118cb848260208601611866565b91505092915050565b5f6118e66118e1846117f9565b611642565b9050808382526020820190506020840283018581111561190957611908611681565b5b835b8181101561195057803567ffffffffffffffff81111561192e5761192d6115e0565b5b80860161193b89826118a7565b8552602085019450505060208101905061190b565b5050509392505050565b5f82601f83011261196e5761196d6115e0565b5b813561197e8482602086016118d4565b91505092915050565b5f806040838503121561199d5761199c611512565b5b5f83013567ffffffffffffffff8111156119ba576119b9611516565b5b6119c6858286016117cc565b925050602083013567ffffffffffffffff8111156119e7576119e6611516565b5b6119f38582860161195a565b9150509250929050565b5f8115159050919050565b611a11816119fd565b82525050565b5f602082019050611a2a5f830184611a08565b92915050565b611a398161151a565b82525050565b5f602082019050611a525f830184611a30565b92915050565b5f60208284031215611a6d57611a6c611512565b5b5f82013567ffffffffffffffff811115611a8a57611a89611516565b5b611a96848285016118a7565b91505092915050565b5f8060408385031215611ab557611ab4611512565b5b5f611ac285828601611539565b925050602083013567ffffffffffffffff811115611ae357611ae2611516565b5b611aef858286016118a7565b9150509250929050565b5f67ffffffffffffffff821115611b1357611b126115e4565b5b602082029050602081019050919050565b5f67ffffffffffffffff821115611b3e57611b3d6115e4565b5b611b478261140f565b9050602081019050919050565b5f611b66611b6184611b24565b611642565b905082815260208101848484011115611b8257611b81611824565b5b611b8d848285611858565b509392505050565b5f82601f830112611ba957611ba86115e0565b5b8135611bb9848260208601611b54565b91505092915050565b5f611bd4611bcf84611af9565b611642565b90508083825260208201905060208402830185811115611bf757611bf6611681565b5b835b81811015611c3e57803567ffffffffffffffff811115611c1c57611c1b6115e0565b5b808601611c298982611b95565b85526020850194505050602081019050611bf9565b5050509392505050565b5f82601f830112611c5c57611c5b6115e0565b5b8135611c6c848260208601611bc2565b91505092915050565b5f60208284031215611c8a57611c89611512565b5b5f82013567ffffffffffffffff811115611ca757611ca6611516565b5b611cb384828501611c48565b91505092915050565b5f805f60c08486031215611cd357611cd2611512565b5b5f611ce0868287016116e4565b935050608084013567ffffffffffffffff811115611d0157611d00611516565b5b611d0d868287016117cc565b92505060a084013567ffffffffffffffff811115611d2e57611d2d611516565b5b611d3a8682870161195a565b9150509250925092565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f611d9182611d6d565b611d9b8185611d77565b9350611dab8185602086016113e7565b611db48161140f565b840191505092915050565b5f611dca8383611d87565b905092915050565b5f602082019050919050565b5f611de882611d44565b611df28185611d4e565b935083602082028501611e0485611d5e565b805f5b85811015611e3f5784840389528151611e208582611dbf565b9450611e2b83611dd2565b925060208a01995050600181019050611e07565b50829750879550505050505092915050565b5f6020820190508181035f830152611e698184611dde565b905092915050565b5f82825260208201905092915050565b5f611e8b82611d6d565b611e958185611e71565b9350611ea58185602086016113e7565b611eae8161140f565b840191505092915050565b5f6020820190508181035f830152611ed18184611e81565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611f0b8161151a565b82525050565b5f611f1c8383611f02565b60208301905092915050565b5f602082019050919050565b5f611f3e82611ed9565b611f488185611ee3565b9350611f5383611ef3565b805f5b83811015611f83578151611f6a8882611f11565b9750611f7583611f28565b925050600181019050611f56565b5085935050505092915050565b5f6020820190508181035f830152611fa88184611f34565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611ff457607f821691505b60208210810361200757612006611fb0565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6120718261151a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120a3576120a261203a565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261210a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826120cf565b61211486836120cf565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61214f61214a6121458461151a565b61212c565b61151a565b9050919050565b5f819050919050565b61216883612135565b61217c61217482612156565b8484546120db565b825550505050565b5f90565b612190612184565b61219b81848461215f565b505050565b5b818110156121be576121b35f82612188565b6001810190506121a1565b5050565b601f821115612203576121d4816120ae565b6121dd846120c0565b810160208510156121ec578190505b6122006121f8856120c0565b8301826121a0565b50505b505050565b5f82821c905092915050565b5f6122235f1984600802612208565b1980831691505092915050565b5f61223b8383612214565b9150826002028217905092915050565b612254826113cd565b67ffffffffffffffff81111561226d5761226c6115e4565b5b6122778254611fdd565b6122828282856121c2565b5f60209050601f8311600181146122b3575f84156122a1578287015190505b6122ab8582612230565b865550612312565b601f1984166122c1866120ae565b5f5b828110156122e8578489015182556001820191506020850194506020810190506122c3565b868310156123055784890151612301601f891682612214565b8355505b6001600288020188555050505b505050505050565b5f8154905061232881611fdd565b9050919050565b5f819050815f5260205f209050919050565b81810361234f575050612424565b6123588261231a565b67ffffffffffffffff811115612371576123706115e4565b5b61237b8254611fdd565b6123868282856121c2565b5f601f8311600181146123b3575f84156123a1578287015490505b6123ab8582612230565b86555061241d565b601f1984166123c18761232f565b96506123cc866120ae565b5f5b828110156123f3578489015482556001820191506001850194506020810190506123ce565b86831015612410578489015461240c601f891682612214565b8355505b6001600288020188555050505b5050505050505b565b5f819050815f5260205f209050919050565b601f8211156124795761244a81612426565b612453846120c0565b81016020851015612462578190505b61247661246e856120c0565b8301826121a0565b50505b505050565b61248782611d6d565b67ffffffffffffffff8111156124a05761249f6115e4565b5b6124aa8254611fdd565b6124b5828285612438565b5f60209050601f8311600181146124e6575f84156124d4578287015190505b6124de8582612230565b865550612545565b601f1984166124f486612426565b5f5b8281101561251b578489015182556001820191506020850194506020810190506124f6565b868310156125385784890151612534601f891682612214565b8355505b6001600288020188555050505b505050505050565b5f6060820190508181035f8301526125658186611588565b90506125746020830185611a30565b81810360408301526125868184611588565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220daa297f04a071f7148a1fb693973955badbc8ca72c56bfc533d319d80a79b22d64736f6c63430008150033",
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

// RevokeVC is a paid mutator transaction binding the contract method 0x92ed770e.
//
// Solidity: function revokeVC(uint256[4] _bfIndexes, uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactor) RevokeVC(opts *bind.TransactOpts, _bfIndexes [4]*big.Int, _mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "revokeVC", _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x92ed770e.
//
// Solidity: function revokeVC(uint256[4] _bfIndexes, uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceSession) RevokeVC(_bfIndexes [4]*big.Int, _mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x92ed770e.
//
// Solidity: function revokeVC(uint256[4] _bfIndexes, uint256[] _mtIndexes, string[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactorSession) RevokeVC(_bfIndexes [4]*big.Int, _mtIndexes []*big.Int, _mtValues []string) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0x17a30492.
//
// Solidity: function updateBloomFilter(uint256[4] _indexes) returns()
func (_RevocationService *RevocationServiceTransactor) UpdateBloomFilter(opts *bind.TransactOpts, _indexes [4]*big.Int) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "updateBloomFilter", _indexes)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0x17a30492.
//
// Solidity: function updateBloomFilter(uint256[4] _indexes) returns()
func (_RevocationService *RevocationServiceSession) UpdateBloomFilter(_indexes [4]*big.Int) (*types.Transaction, error) {
	return _RevocationService.Contract.UpdateBloomFilter(&_RevocationService.TransactOpts, _indexes)
}

// UpdateBloomFilter is a paid mutator transaction binding the contract method 0x17a30492.
//
// Solidity: function updateBloomFilter(uint256[4] _indexes) returns()
func (_RevocationService *RevocationServiceTransactorSession) UpdateBloomFilter(_indexes [4]*big.Int) (*types.Transaction, error) {
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
