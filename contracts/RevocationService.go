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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetBloomFilterSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMerkleTreeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveBloomFilter\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveMerkleTree\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrievePublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"addPublicKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_values\",\"type\":\"bytes32[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506121058061005d5f395ff3fe608060405234801561000f575f80fd5b5060043610610140575f3560e01c80635a34497a116100b6578063bbb7e2ef1161007a578063bbb7e2ef1461034e578063c680f4101461036c578063ce4b3f341461039c578063d3e53042146103b8578063f174053b146103c2578063f8a8fd6d146103e057610140565b80635a34497a146102aa5780636a6f2063146102c65780636e0a5904146102f657806399223a5a14610314578063a5eb0de81461033257610140565b80632337db35116101085780632337db35146101e8578063266344c2146102065780632eb4a7ab14610222578063336a76b41461024057806348db53361461025e578063516006981461028e57610140565b8063070f010e146101445780630de54b85146101625780630df0ff901461016c57806319f267171461019c578063203c8ab0146101b8575b5f80fd5b61014c6103ea565b60405161015991906112d7565b60405180910390f35b61016a610498565b005b6101866004803603810190610181919061133b565b61049a565b6040516101939190611375565b60405180910390f35b6101b660048036038101906101b191906113b8565b6104af565b005b6101d260048036038101906101cd9190611546565b610521565b6040516101df91906115a7565b60405180910390f35b6101f06105c8565b6040516101fd9190611375565b60405180910390f35b610220600480360381019061021b9190611680565b6105e0565b005b61022a610650565b6040516102379190611375565b60405180910390f35b610248610656565b6040516102559190611733565b60405180910390f35b6102786004803603810190610273919061133b565b610694565b6040516102859190611733565b60405180910390f35b6102a860048036038101906102a3919061174c565b6106a8565b005b6102c460048036038101906102bf9190611950565b610877565b005b6102e060048036038101906102db9190611546565b61093d565b6040516102ed91906115a7565b60405180910390f35b6102fe61094e565b60405161030b9190611733565b60405180910390f35b61031c6109a1565b6040516103299190611acc565b60405180910390f35b61034c60048036038101906103479190611546565b610a75565b005b610356610bcd565b6040516103639190611375565b60405180910390f35b6103866004803603810190610381919061133b565b610be5565b6040516103939190611b34565b60405180910390f35b6103b660048036038101906103b1919061174c565b610c8b565b005b6103c0610cf1565b005b6103ca610de8565b6040516103d79190611c0b565b60405180910390f35b6103e8610e3e565b005b60605f60058054905067ffffffffffffffff81111561040c5761040b61140a565b5b60405190808252806020026020018201604052801561043a5781602001602082028036833780820191505090505b5090505f5b6005805490508110156104905760035f8281526020019081526020015f205482828151811061047157610470611c2b565b5b602002602001018181525050808061048890611c85565b91505061043f565b508091505090565b565b6003602052805f5260405f205f915090505481565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610507575f80fd5b8060035f8481526020019081526020015f20819055505050565b5f805f90505f5b83518110156105be575f600885838151811061054757610546611c2b565b5b6020026020010151901c90505f60ff86848151811061056957610568611c2b565b5b6020026020010151166001901b90505f80825f808681526020019081526020015f205416141590505f1515811515036105a857600194505050506105be565b50505080806105b690611c85565b915050610528565b5080915050919050565b5f60035f8081526020019081526020015f2054905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610638575f80fd5b61064183610a75565b61064b82826106a8565b505050565b60045481565b5f805f90505f5b60018054905081101561068c576020826106779190611ccc565b9150808061068490611c85565b91505061065d565b508091505090565b5f602052805f5260405f205f915090505481565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610700575f80fd5b805182511461070d575f80fd5b5f5b82518110156107755781818151811061072b5761072a611c2b565b5b602002602001015160035f85848151811061074957610748611c2b565b5b602002602001015181526020019081526020015f2081905550808061076d90611c85565b91505061070f565b505f5b825181101561085a575f151560065f85848151811061079a57610799611c2b565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff1615150361084757600160065f8584815181106107db576107da611c2b565b5b602002602001015181526020019081526020015f205f6101000a81548160ff021916908315150217905550600583828151811061081b5761081a611c2b565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b808061085290611c85565b915050610778565b5060035f8081526020019081526020015f20546004819055505050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108cf575f80fd5b5f5b81518110156109395760088282815181106108ef576108ee611c2b565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f9091909190915090816109259190611ef9565b50808061093190611c85565b9150506108d1565b5050565b5f61094782610521565b9050919050565b5f805f90505f5b6005805490508110156109995760035f8281526020019081526020015f2050602060ff16826109849190611ccc565b9150808061099190611c85565b915050610955565b508091505090565b60606008805480602002602001604051908101604052809291908181526020015f905b82821015610a6c578382905f5260205f200180546109e190611d2c565b80601f0160208091040260200160405190810160405280929190818152602001828054610a0d90611d2c565b8015610a585780601f10610a2f57610100808354040283529160200191610a58565b820191905f5260205f20905b815481529060010190602001808311610a3b57829003601f168201915b5050505050815260200190600101906109c4565b50505050905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610acd575f80fd5b5f5b8151811015610bc9575f6008838381518110610aee57610aed611c2b565b5b6020026020010151901c90505f60ff848481518110610b1057610b0f611c2b565b5b6020026020010151166001901b9050805f808481526020019081526020015f205f82825417925050819055505f151560025f8481526020019081526020015f205f9054906101000a900460ff16151503610bb457600160025f8481526020019081526020015f205f6101000a81548160ff021916908315150217905550600182908060018154018082558091505060019003905f5260205f20015f90919091909150555b50508080610bc190611c85565b915050610acf565b5050565b5f60035f8081526020019081526020015f2054905090565b60088181548110610bf4575f80fd5b905f5260205f20015f915090508054610c0c90611d2c565b80601f0160208091040260200160405190810160405280929190818152602001828054610c3890611d2c565b8015610c835780601f10610c5a57610100808354040283529160200191610c83565b820191905f5260205f20905b815481529060010190602001808311610c6657829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ce3575f80fd5b610ced82826106a8565b5050565b600115156001151503610d3d57610d3c6040518060400160405280601381526020017f70726974696e67206d65726b6c65207472656500000000000000000000000000815250610ffb565b5b5f5b600580549050811015610de557600115156001151503610dd257610d986040518060400160405280601481526020017f696e646578203a20256420092076616c7565203a00000000000000000000000081525082611094565b610dd160035f8381526020019081526020015f2054604051602001610dbd9190611fe8565b604051602081830303815290604052611130565b5b8080610ddd90611c85565b915050610d3f565b50565b60606001805480602002602001604051908101604052809291908181526020018280548015610e3457602002820191905f5260205f20905b815481526020019060010190808311610e20575b5050505050905090565b5f600267ffffffffffffffff811115610e5a57610e5961140a565b5b604051908082528060200260200182016040528015610e885781602001602082028036833780820191505090505b5090505f815f81518110610e9f57610e9e611c2b565b5b602002602001018181525050600181600181518110610ec157610ec0611c2b565b5b6020026020010181815250505f600267ffffffffffffffff811115610ee957610ee861140a565b5b604051908082528060200260200182016040528015610f175781602001602082028036833780820191505090505b5090507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b815f81518110610f5057610f4f611c2b565b5b6020026020010181815250507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b81600181518110610f9357610f92611c2b565b5b602002602001018181525050610fa98282610c8b565b610fb1610cf1565b610ff76040518060400160405280601581526020017f6d65726b6c6520747265652073697a653a202564200000000000000000000000815250610ff261094e565b611094565b5050565b6110918160405160240161100f9190612054565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111c9565b50565b61112c82826040516024016110aa929190612074565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111c9565b5050565b6111c6816040516024016111449190611b34565b6040516020818303038152906040527f0be77f56000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506111c9565b50565b6111e0816111d86111e3611202565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b61120d819050919050565b6112156120a2565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61125281611240565b82525050565b5f6112638383611249565b60208301905092915050565b5f602082019050919050565b5f61128582611217565b61128f8185611221565b935061129a83611231565b805f5b838110156112ca5781516112b18882611258565b97506112bc8361126f565b92505060018101905061129d565b5085935050505092915050565b5f6020820190508181035f8301526112ef818461127b565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61131a81611308565b8114611324575f80fd5b50565b5f8135905061133581611311565b92915050565b5f602082840312156113505761134f611300565b5b5f61135d84828501611327565b91505092915050565b61136f81611240565b82525050565b5f6020820190506113885f830184611366565b92915050565b61139781611240565b81146113a1575f80fd5b50565b5f813590506113b28161138e565b92915050565b5f80604083850312156113ce576113cd611300565b5b5f6113db85828601611327565b92505060206113ec858286016113a4565b9150509250929050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611440826113fa565b810181811067ffffffffffffffff8211171561145f5761145e61140a565b5b80604052505050565b5f6114716112f7565b905061147d8282611437565b919050565b5f67ffffffffffffffff82111561149c5761149b61140a565b5b602082029050602081019050919050565b5f80fd5b5f6114c36114be84611482565b611468565b905080838252602082019050602084028301858111156114e6576114e56114ad565b5b835b8181101561150f57806114fb8882611327565b8452602084019350506020810190506114e8565b5050509392505050565b5f82601f83011261152d5761152c6113f6565b5b813561153d8482602086016114b1565b91505092915050565b5f6020828403121561155b5761155a611300565b5b5f82013567ffffffffffffffff81111561157857611577611304565b5b61158484828501611519565b91505092915050565b5f8115159050919050565b6115a18161158d565b82525050565b5f6020820190506115ba5f830184611598565b92915050565b5f67ffffffffffffffff8211156115da576115d961140a565b5b602082029050602081019050919050565b5f6115fd6115f8846115c0565b611468565b905080838252602082019050602084028301858111156116205761161f6114ad565b5b835b81811015611649578061163588826113a4565b845260208401935050602081019050611622565b5050509392505050565b5f82601f830112611667576116666113f6565b5b81356116778482602086016115eb565b91505092915050565b5f805f6060848603121561169757611696611300565b5b5f84013567ffffffffffffffff8111156116b4576116b3611304565b5b6116c086828701611519565b935050602084013567ffffffffffffffff8111156116e1576116e0611304565b5b6116ed86828701611519565b925050604084013567ffffffffffffffff81111561170e5761170d611304565b5b61171a86828701611653565b9150509250925092565b61172d81611308565b82525050565b5f6020820190506117465f830184611724565b92915050565b5f806040838503121561176257611761611300565b5b5f83013567ffffffffffffffff81111561177f5761177e611304565b5b61178b85828601611519565b925050602083013567ffffffffffffffff8111156117ac576117ab611304565b5b6117b885828601611653565b9150509250929050565b5f67ffffffffffffffff8211156117dc576117db61140a565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff82111561180b5761180a61140a565b5b611814826113fa565b9050602081019050919050565b828183375f83830152505050565b5f61184161183c846117f1565b611468565b90508281526020810184848401111561185d5761185c6117ed565b5b611868848285611821565b509392505050565b5f82601f830112611884576118836113f6565b5b813561189484826020860161182f565b91505092915050565b5f6118af6118aa846117c2565b611468565b905080838252602082019050602084028301858111156118d2576118d16114ad565b5b835b8181101561191957803567ffffffffffffffff8111156118f7576118f66113f6565b5b8086016119048982611870565b855260208501945050506020810190506118d4565b5050509392505050565b5f82601f830112611937576119366113f6565b5b813561194784826020860161189d565b91505092915050565b5f6020828403121561196557611964611300565b5b5f82013567ffffffffffffffff81111561198257611981611304565b5b61198e84828501611923565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156119f75780820151818401526020810190506119dc565b5f8484015250505050565b5f611a0c826119c0565b611a1681856119ca565b9350611a268185602086016119da565b611a2f816113fa565b840191505092915050565b5f611a458383611a02565b905092915050565b5f602082019050919050565b5f611a6382611997565b611a6d81856119a1565b935083602082028501611a7f856119b1565b805f5b85811015611aba5784840389528151611a9b8582611a3a565b9450611aa683611a4d565b925060208a01995050600181019050611a82565b50829750879550505050505092915050565b5f6020820190508181035f830152611ae48184611a59565b905092915050565b5f82825260208201905092915050565b5f611b06826119c0565b611b108185611aec565b9350611b208185602086016119da565b611b29816113fa565b840191505092915050565b5f6020820190508181035f830152611b4c8184611afc565b905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611b8681611308565b82525050565b5f611b978383611b7d565b60208301905092915050565b5f602082019050919050565b5f611bb982611b54565b611bc38185611b5e565b9350611bce83611b6e565b805f5b83811015611bfe578151611be58882611b8c565b9750611bf083611ba3565b925050600181019050611bd1565b5085935050505092915050565b5f6020820190508181035f830152611c238184611baf565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611c8f82611308565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611cc157611cc0611c58565b5b600182019050919050565b5f611cd682611308565b9150611ce183611308565b9250828201905080821115611cf957611cf8611c58565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611d4357607f821691505b602082108103611d5657611d55611cff565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611db87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611d7d565b611dc28683611d7d565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611dfd611df8611df384611308565b611dda565b611308565b9050919050565b5f819050919050565b611e1683611de3565b611e2a611e2282611e04565b848454611d89565b825550505050565b5f90565b611e3e611e32565b611e49818484611e0d565b505050565b5b81811015611e6c57611e615f82611e36565b600181019050611e4f565b5050565b601f821115611eb157611e8281611d5c565b611e8b84611d6e565b81016020851015611e9a578190505b611eae611ea685611d6e565b830182611e4e565b50505b505050565b5f82821c905092915050565b5f611ed15f1984600802611eb6565b1980831691505092915050565b5f611ee98383611ec2565b9150826002028217905092915050565b611f02826119c0565b67ffffffffffffffff811115611f1b57611f1a61140a565b5b611f258254611d2c565b611f30828285611e70565b5f60209050601f831160018114611f61575f8415611f4f578287015190505b611f598582611ede565b865550611fc0565b601f198416611f6f86611d5c565b5f5b82811015611f9657848901518255600182019150602085019450602081019050611f71565b86831015611fb35784890151611faf601f891682611ec2565b8355505b6001600288020188555050505b505050505050565b5f819050919050565b611fe2611fdd82611240565b611fc8565b82525050565b5f611ff38284611fd1565b60208201915081905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f61202682612002565b612030818561200c565b93506120408185602086016119da565b612049816113fa565b840191505092915050565b5f6020820190508181035f83015261206c818461201c565b905092915050565b5f6040820190508181035f83015261208c818561201c565b905061209b6020830184611724565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220b069f3fa5f462e9e8ddbe9d3c12d6f9d4b06baad7eeb0a2ba5f1b3762cf1e8fc64736f6c63430008150033",
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

// GetBloomFilterSize is a free data retrieval call binding the contract method 0x336a76b4.
//
// Solidity: function GetBloomFilterSize() view returns(uint256)
func (_RevocationService *RevocationServiceCaller) GetBloomFilterSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "GetBloomFilterSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBloomFilterSize is a free data retrieval call binding the contract method 0x336a76b4.
//
// Solidity: function GetBloomFilterSize() view returns(uint256)
func (_RevocationService *RevocationServiceSession) GetBloomFilterSize() (*big.Int, error) {
	return _RevocationService.Contract.GetBloomFilterSize(&_RevocationService.CallOpts)
}

// GetBloomFilterSize is a free data retrieval call binding the contract method 0x336a76b4.
//
// Solidity: function GetBloomFilterSize() view returns(uint256)
func (_RevocationService *RevocationServiceCallerSession) GetBloomFilterSize() (*big.Int, error) {
	return _RevocationService.Contract.GetBloomFilterSize(&_RevocationService.CallOpts)
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
