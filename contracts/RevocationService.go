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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetMerkleTreeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveBloomFilter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveMerkleTree\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrievePublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"addPublicKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_values\",\"type\":\"bytes32[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061210b8061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610135575f3560e01c80635a34497a116100b6578063bbb7e2ef1161007a578063bbb7e2ef14610325578063c680f41014610343578063ce4b3f3414610373578063d3e530421461038f578063f174053b14610399578063f8a8fd6d146103b757610135565b80635a34497a146102815780636a6f20631461029d5780636e0a5904146102cd57806399223a5a146102eb578063a5eb0de81461030957610135565b80632337db35116100fd5780632337db35146101dd578063266344c2146101fb5780632eb4a7ab1461021757806348db533614610235578063516006981461026557610135565b8063070f010e146101395780630de54b85146101575780630df0ff901461016157806319f2671714610191578063203c8ab0146101ad575b5f80fd5b6101416103c1565b60405161014e91906112dd565b60405180910390f35b61015f61046f565b005b61017b60048036038101906101769190611341565b610471565b604051610188919061137b565b60405180910390f35b6101ab60048036038101906101a691906113be565b610486565b005b6101c760048036038101906101c2919061154c565b6104f8565b6040516101d491906115ad565b60405180910390f35b6101e561059f565b6040516101f2919061137b565b60405180910390f35b61021560048036038101906102109190611686565b6105b7565b005b61021f610627565b60405161022c919061137b565b60405180910390f35b61024f600480360381019061024a9190611341565b61062d565b60405161025c9190611739565b60405180910390f35b61027f600480360381019061027a9190611752565b610641565b005b61029b60048036038101906102969190611956565b610810565b005b6102b760048036038101906102b2919061154c565b6108d6565b6040516102c491906115ad565b60405180910390f35b6102d56108e7565b6040516102e29190611739565b60405180910390f35b6102f361093a565b6040516103009190611ad2565b60405180910390f35b610323600480360381019061031e919061154c565b610a0e565b005b61032d610bd3565b60405161033a919061137b565b60405180910390f35b61035d60048036038101906103589190611341565b610beb565b60405161036a9190611b3a565b60405180910390f35b61038d60048036038101906103889190611752565b610c91565b005b610397610cf7565b005b6103a1610dee565b6040516103ae9190611c11565b60405180910390f35b6103bf610e44565b005b60605f60058054905067ffffffffffffffff8111156103e3576103e2611410565b5b6040519080825280602002602001820160405280156104115781602001602082028036833780820191505090505b5090505f5b6005805490508110156104675760035f8281526020019081526020015f205482828151811061044857610447611c31565b5b602002602001018181525050808061045f90611c8b565b915050610416565b508091505090565b565b6003602052805f5260405f205f915090505481565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104de575f80fd5b8060035f8481526020019081526020015f20819055505050565b5f805f90505f5b8351811015610595575f600885838151811061051e5761051d611c31565b5b6020026020010151901c90505f60ff8684815181106105405761053f611c31565b5b6020026020010151166001901b90505f80825f808681526020019081526020015f205416141590505f15158115150361057f5760019450505050610595565b505050808061058d90611c8b565b9150506104ff565b5080915050919050565b5f60035f8081526020019081526020015f2054905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461060f575f80fd5b61061883610a0e565b6106228282610641565b505050565b60045481565b5f602052805f5260405f205f915090505481565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610699575f80fd5b80518251146106a6575f80fd5b5f5b825181101561070e578181815181106106c4576106c3611c31565b5b602002602001015160035f8584815181106106e2576106e1611c31565b5b602002602001015181526020019081526020015f2081905550808061070690611c8b565b9150506106a8565b505f5b82518110156107f3575f151560065f85848151811061073357610732611c31565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff161515036107e057600160065f85848151811061077457610773611c31565b5b602002602001015181526020019081526020015f205f6101000a81548160ff02191690831515021790555060058382815181106107b4576107b3611c31565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b80806107eb90611c8b565b915050610711565b5060035f8081526020019081526020015f20546004819055505050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610868575f80fd5b5f5b81518110156108d257600882828151811061088857610887611c31565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f9091909190915090816108be9190611ecc565b5080806108ca90611c8b565b91505061086a565b5050565b5f6108e0826104f8565b9050919050565b5f805f90505f5b6005805490508110156109325760035f8281526020019081526020015f2050602060ff168261091d9190611f9b565b9150808061092a90611c8b565b9150506108ee565b508091505090565b60606008805480602002602001604051908101604052809291908181526020015f905b82821015610a05578382905f5260205f2001805461097a90611cff565b80601f01602080910402602001604051908101604052809291908181526020018280546109a690611cff565b80156109f15780601f106109c8576101008083540402835291602001916109f1565b820191905f5260205f20905b8154815290600101906020018083116109d457829003601f168201915b50505050508152602001906001019061095d565b50505050905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a66575f80fd5b5f5b8151811015610aea575f6008838381518110610a8757610a86611c31565b5b6020026020010151901c90505f60ff848481518110610aa957610aa8611c31565b5b6020026020010151166001901b9050805f808481526020019081526020015f205f828254179250508190555050508080610ae290611c8b565b915050610a68565b505f5b8151811015610bcf575f151560025f848481518110610b0f57610b0e611c31565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff16151503610bbc57600160025f848481518110610b5057610b4f611c31565b5b602002602001015181526020019081526020015f205f6101000a81548160ff0219169083151502179055506001828281518110610b9057610b8f611c31565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b8080610bc790611c8b565b915050610aed565b5050565b5f60035f8081526020019081526020015f2054905090565b60088181548110610bfa575f80fd5b905f5260205f20015f915090508054610c1290611cff565b80601f0160208091040260200160405190810160405280929190818152602001828054610c3e90611cff565b8015610c895780601f10610c6057610100808354040283529160200191610c89565b820191905f5260205f20905b815481529060010190602001808311610c6c57829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ce9575f80fd5b610cf38282610641565b5050565b600115156001151503610d4357610d426040518060400160405280601381526020017f70726974696e67206d65726b6c65207472656500000000000000000000000000815250611001565b5b5f5b600580549050811015610deb57600115156001151503610dd857610d9e6040518060400160405280601481526020017f696e646578203a20256420092076616c7565203a0000000000000000000000008152508261109a565b610dd760035f8381526020019081526020015f2054604051602001610dc39190611fee565b604051602081830303815290604052611136565b5b8080610de390611c8b565b915050610d45565b50565b60606001805480602002602001604051908101604052809291908181526020018280548015610e3a57602002820191905f5260205f20905b815481526020019060010190808311610e26575b5050505050905090565b5f600267ffffffffffffffff811115610e6057610e5f611410565b5b604051908082528060200260200182016040528015610e8e5781602001602082028036833780820191505090505b5090505f815f81518110610ea557610ea4611c31565b5b602002602001018181525050600181600181518110610ec757610ec6611c31565b5b6020026020010181815250505f600267ffffffffffffffff811115610eef57610eee611410565b5b604051908082528060200260200182016040528015610f1d5781602001602082028036833780820191505090505b5090507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b815f81518110610f5657610f55611c31565b5b6020026020010181815250507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b81600181518110610f9957610f98611c31565b5b602002602001018181525050610faf8282610c91565b610fb7610cf7565b610ffd6040518060400160405280601581526020017f6d65726b6c6520747265652073697a653a202564200000000000000000000000815250610ff86108e7565b61109a565b5050565b61109781604051602401611015919061205a565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111cf565b50565b61113282826040516024016110b092919061207a565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111cf565b5050565b6111cc8160405160240161114a9190611b3a565b6040516020818303038152906040527f0be77f56000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111cf565b50565b6111e6816111de6111e9611208565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b611213819050919050565b61121b6120a8565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61125881611246565b82525050565b5f611269838361124f565b60208301905092915050565b5f602082019050919050565b5f61128b8261121d565b6112958185611227565b93506112a083611237565b805f5b838110156112d05781516112b7888261125e565b97506112c283611275565b9250506001810190506112a3565b5085935050505092915050565b5f6020820190508181035f8301526112f58184611281565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6113208161130e565b811461132a575f80fd5b50565b5f8135905061133b81611317565b92915050565b5f6020828403121561135657611355611306565b5b5f6113638482850161132d565b91505092915050565b61137581611246565b82525050565b5f60208201905061138e5f83018461136c565b92915050565b61139d81611246565b81146113a7575f80fd5b50565b5f813590506113b881611394565b92915050565b5f80604083850312156113d4576113d3611306565b5b5f6113e18582860161132d565b92505060206113f2858286016113aa565b9150509250929050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61144682611400565b810181811067ffffffffffffffff8211171561146557611464611410565b5b80604052505050565b5f6114776112fd565b9050611483828261143d565b919050565b5f67ffffffffffffffff8211156114a2576114a1611410565b5b602082029050602081019050919050565b5f80fd5b5f6114c96114c484611488565b61146e565b905080838252602082019050602084028301858111156114ec576114eb6114b3565b5b835b818110156115155780611501888261132d565b8452602084019350506020810190506114ee565b5050509392505050565b5f82601f830112611533576115326113fc565b5b81356115438482602086016114b7565b91505092915050565b5f6020828403121561156157611560611306565b5b5f82013567ffffffffffffffff81111561157e5761157d61130a565b5b61158a8482850161151f565b91505092915050565b5f8115159050919050565b6115a781611593565b82525050565b5f6020820190506115c05f83018461159e565b92915050565b5f67ffffffffffffffff8211156115e0576115df611410565b5b602082029050602081019050919050565b5f6116036115fe846115c6565b61146e565b90508083825260208201905060208402830185811115611626576116256114b3565b5b835b8181101561164f578061163b88826113aa565b845260208401935050602081019050611628565b5050509392505050565b5f82601f83011261166d5761166c6113fc565b5b813561167d8482602086016115f1565b91505092915050565b5f805f6060848603121561169d5761169c611306565b5b5f84013567ffffffffffffffff8111156116ba576116b961130a565b5b6116c68682870161151f565b935050602084013567ffffffffffffffff8111156116e7576116e661130a565b5b6116f38682870161151f565b925050604084013567ffffffffffffffff8111156117145761171361130a565b5b61172086828701611659565b9150509250925092565b6117338161130e565b82525050565b5f60208201905061174c5f83018461172a565b92915050565b5f806040838503121561176857611767611306565b5b5f83013567ffffffffffffffff8111156117855761178461130a565b5b6117918582860161151f565b925050602083013567ffffffffffffffff8111156117b2576117b161130a565b5b6117be85828601611659565b9150509250929050565b5f67ffffffffffffffff8211156117e2576117e1611410565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff82111561181157611810611410565b5b61181a82611400565b9050602081019050919050565b828183375f83830152505050565b5f611847611842846117f7565b61146e565b905082815260208101848484011115611863576118626117f3565b5b61186e848285611827565b509392505050565b5f82601f83011261188a576118896113fc565b5b813561189a848260208601611835565b91505092915050565b5f6118b56118b0846117c8565b61146e565b905080838252602082019050602084028301858111156118d8576118d76114b3565b5b835b8181101561191f57803567ffffffffffffffff8111156118fd576118fc6113fc565b5b80860161190a8982611876565b855260208501945050506020810190506118da565b5050509392505050565b5f82601f83011261193d5761193c6113fc565b5b813561194d8482602086016118a3565b91505092915050565b5f6020828403121561196b5761196a611306565b5b5f82013567ffffffffffffffff8111156119885761198761130a565b5b61199484828501611929565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156119fd5780820151818401526020810190506119e2565b5f8484015250505050565b5f611a12826119c6565b611a1c81856119d0565b9350611a2c8185602086016119e0565b611a3581611400565b840191505092915050565b5f611a4b8383611a08565b905092915050565b5f602082019050919050565b5f611a698261199d565b611a7381856119a7565b935083602082028501611a85856119b7565b805f5b85811015611ac05784840389528151611aa18582611a40565b9450611aac83611a53565b925060208a01995050600181019050611a88565b50829750879550505050505092915050565b5f6020820190508181035f830152611aea8184611a5f565b905092915050565b5f82825260208201905092915050565b5f611b0c826119c6565b611b168185611af2565b9350611b268185602086016119e0565b611b2f81611400565b840191505092915050565b5f6020820190508181035f830152611b528184611b02565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611b8c8161130e565b82525050565b5f611b9d8383611b83565b60208301905092915050565b5f602082019050919050565b5f611bbf82611b5a565b611bc98185611b64565b9350611bd483611b74565b805f5b83811015611c04578151611beb8882611b92565b9750611bf683611ba9565b925050600181019050611bd7565b5085935050505092915050565b5f6020820190508181035f830152611c298184611bb5565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611c958261130e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611cc757611cc6611c5e565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611d1657607f821691505b602082108103611d2957611d28611cd2565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611d8b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611d50565b611d958683611d50565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611dd0611dcb611dc68461130e565b611dad565b61130e565b9050919050565b5f819050919050565b611de983611db6565b611dfd611df582611dd7565b848454611d5c565b825550505050565b5f90565b611e11611e05565b611e1c818484611de0565b505050565b5b81811015611e3f57611e345f82611e09565b600181019050611e22565b5050565b601f821115611e8457611e5581611d2f565b611e5e84611d41565b81016020851015611e6d578190505b611e81611e7985611d41565b830182611e21565b50505b505050565b5f82821c905092915050565b5f611ea45f1984600802611e89565b1980831691505092915050565b5f611ebc8383611e95565b9150826002028217905092915050565b611ed5826119c6565b67ffffffffffffffff811115611eee57611eed611410565b5b611ef88254611cff565b611f03828285611e43565b5f60209050601f831160018114611f34575f8415611f22578287015190505b611f2c8582611eb1565b865550611f93565b601f198416611f4286611d2f565b5f5b82811015611f6957848901518255600182019150602085019450602081019050611f44565b86831015611f865784890151611f82601f891682611e95565b8355505b6001600288020188555050505b505050505050565b5f611fa58261130e565b9150611fb08361130e565b9250828201905080821115611fc857611fc7611c5e565b5b92915050565b5f819050919050565b611fe8611fe382611246565b611fce565b82525050565b5f611ff98284611fd7565b60208201915081905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f61202c82612008565b6120368185612012565b93506120468185602086016119e0565b61204f81611400565b840191505092915050565b5f6020820190508181035f8301526120728184612022565b905092915050565b5f6040820190508181035f8301526120928185612022565b90506120a1602083018461172a565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220e44b51987a69ac54bcc92df74906d6208863c4f6fc6549332164256841904dd164736f6c63430008150033",
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
// Solidity: function RetrieveMerkleTree() view returns(bytes32[])
func (_RevocationService *RevocationServiceCaller) RetrieveMerkleTree(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "RetrieveMerkleTree")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x070f010e.
//
// Solidity: function RetrieveMerkleTree() view returns(bytes32[])
func (_RevocationService *RevocationServiceSession) RetrieveMerkleTree() ([][32]byte, error) {
	return _RevocationService.Contract.RetrieveMerkleTree(&_RevocationService.CallOpts)
}

// RetrieveMerkleTree is a free data retrieval call binding the contract method 0x070f010e.
//
// Solidity: function RetrieveMerkleTree() view returns(bytes32[])
func (_RevocationService *RevocationServiceCallerSession) RetrieveMerkleTree() ([][32]byte, error) {
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

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_RevocationService *RevocationServiceCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_RevocationService *RevocationServiceSession) MerkleRoot() ([32]byte, error) {
	return _RevocationService.Contract.MerkleRoot(&_RevocationService.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_RevocationService *RevocationServiceCallerSession) MerkleRoot() ([32]byte, error) {
	return _RevocationService.Contract.MerkleRoot(&_RevocationService.CallOpts)
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
