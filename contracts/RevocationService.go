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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GetMerkleTreeSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrieveMerkleTree\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RetrievePublicKeys\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"name\":\"addPublicKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_values\",\"type\":\"bytes32[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611e898061005d5f395ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c80635a34497a116100ab578063bbb7e2ef1161006f578063bbb7e2ef146102f1578063c680f4101461030f578063ce4b3f341461033f578063d3e530421461035b578063f8a8fd6d146103655761011f565b80635a34497a1461024d5780636a6f2063146102695780636e0a59041461029957806399223a5a146102b7578063a5eb0de8146102d55761011f565b8063203c8ab0116100f2578063203c8ab0146101975780632337db35146101c7578063266344c2146101e557806348db53361461020157806351600698146102315761011f565b8063070f010e146101235780630de54b85146101415780630df0ff901461014b57806319f267171461017b575b5f80fd5b61012b61036f565b6040516101389190611132565b60405180910390f35b61014961041d565b005b61016560048036038101906101609190611196565b61041f565b60405161017291906111d0565b60405180910390f35b61019560048036038101906101909190611213565b610434565b005b6101b160048036038101906101ac91906113a1565b6104a6565b6040516101be9190611402565b60405180910390f35b6101cf61054d565b6040516101dc91906111d0565b60405180910390f35b6101ff60048036038101906101fa91906114db565b610565565b005b61021b60048036038101906102169190611196565b6105d5565b604051610228919061158e565b60405180910390f35b61024b600480360381019061024691906115a7565b6105e9565b005b610267600480360381019061026291906117ab565b6107a0565b005b610283600480360381019061027e91906113a1565b610866565b6040516102909190611402565b60405180910390f35b6102a1610877565b6040516102ae919061158e565b60405180910390f35b6102bf6108ca565b6040516102cc9190611927565b60405180910390f35b6102ef60048036038101906102ea91906113a1565b61099e565b005b6102f9610a7e565b60405161030691906111d0565b60405180910390f35b61032960048036038101906103249190611196565b610a96565b604051610336919061198f565b60405180910390f35b610359600480360381019061035491906115a7565b610b3c565b005b610363610ba2565b005b61036d610c99565b005b60605f60028054905067ffffffffffffffff81111561039157610390611265565b5b6040519080825280602002602001820160405280156103bf5781602001602082028036833780820191505090505b5090505f5b6002805490508110156104155760015f8281526020019081526020015f20548282815181106103f6576103f56119af565b5b602002602001018181525050808061040d90611a09565b9150506103c4565b508091505090565b565b6001602052805f5260405f205f915090505481565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461048c575f80fd5b8060015f8481526020019081526020015f20819055505050565b5f805f90505f5b8351811015610543575f60088583815181106104cc576104cb6119af565b5b6020026020010151901c90505f60ff8684815181106104ee576104ed6119af565b5b6020026020010151166001901b90505f80825f808681526020019081526020015f205416141590505f15158115150361052d5760019450505050610543565b505050808061053b90611a09565b9150506104ad565b5080915050919050565b5f60015f8081526020019081526020015f2054905090565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105bd575f80fd5b6105c68361099e565b6105d082826105e9565b505050565b5f602052805f5260405f205f915090505481565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610641575f80fd5b805182511461064e575f80fd5b5f5b82518110156106b65781818151811061066c5761066b6119af565b5b602002602001015160015f85848151811061068a576106896119af565b5b602002602001015181526020019081526020015f208190555080806106ae90611a09565b915050610650565b505f5b825181101561079b575f151560035f8584815181106106db576106da6119af565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff1615150361078857600160035f85848151811061071c5761071b6119af565b5b602002602001015181526020019081526020015f205f6101000a81548160ff021916908315150217905550600283828151811061075c5761075b6119af565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b808061079390611a09565b9150506106b9565b505050565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f8575f80fd5b5f5b8151811015610862576005828281518110610818576108176119af565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150908161084e9190611c4a565b50808061085a90611a09565b9150506107fa565b5050565b5f610870826104a6565b9050919050565b5f805f90505f5b6002805490508110156108c25760015f8281526020019081526020015f2050602060ff16826108ad9190611d19565b915080806108ba90611a09565b91505061087e565b508091505090565b60606005805480602002602001604051908101604052809291908181526020015f905b82821015610995578382905f5260205f2001805461090a90611a7d565b80601f016020809104026020016040519081016040528092919081815260200182805461093690611a7d565b80156109815780601f1061095857610100808354040283529160200191610981565b820191905f5260205f20905b81548152906001019060200180831161096457829003601f168201915b5050505050815260200190600101906108ed565b50505050905090565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109f6575f80fd5b5f5b8151811015610a7a575f6008838381518110610a1757610a166119af565b5b6020026020010151901c90505f60ff848481518110610a3957610a386119af565b5b6020026020010151166001901b9050805f808481526020019081526020015f205f828254179250508190555050508080610a7290611a09565b9150506109f8565b5050565b5f60015f8081526020019081526020015f2054905090565b60058181548110610aa5575f80fd5b905f5260205f20015f915090508054610abd90611a7d565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae990611a7d565b8015610b345780601f10610b0b57610100808354040283529160200191610b34565b820191905f5260205f20905b815481529060010190602001808311610b1757829003601f168201915b505050505081565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b94575f80fd5b610b9e82826105e9565b5050565b600115156001151503610bee57610bed6040518060400160405280601381526020017f70726974696e67206d65726b6c65207472656500000000000000000000000000815250610e56565b5b5f5b600280549050811015610c9657600115156001151503610c8357610c496040518060400160405280601481526020017f696e646578203a20256420092076616c7565203a00000000000000000000000081525082610eef565b610c8260015f8381526020019081526020015f2054604051602001610c6e9190611d6c565b604051602081830303815290604052610f8b565b5b8080610c8e90611a09565b915050610bf0565b50565b5f600267ffffffffffffffff811115610cb557610cb4611265565b5b604051908082528060200260200182016040528015610ce35781602001602082028036833780820191505090505b5090505f815f81518110610cfa57610cf96119af565b5b602002602001018181525050600181600181518110610d1c57610d1b6119af565b5b6020026020010181815250505f600267ffffffffffffffff811115610d4457610d43611265565b5b604051908082528060200260200182016040528015610d725781602001602082028036833780820191505090505b5090507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b815f81518110610dab57610daa6119af565b5b6020026020010181815250507f68656c6c6f0000000000000000000000000000000000000000000000000000005f1b81600181518110610dee57610ded6119af565b5b602002602001018181525050610e048282610b3c565b610e0c610ba2565b610e526040518060400160405280601581526020017f6d65726b6c6520747265652073697a653a202564200000000000000000000000815250610e4d610877565b610eef565b5050565b610eec81604051602401610e6a9190611dd8565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611024565b50565b610f878282604051602401610f05929190611df8565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611024565b5050565b61102181604051602401610f9f919061198f565b6040516020818303038152906040527f0be77f56000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611024565b50565b61103b8161103361103e61105d565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b611068819050919050565b611070611e26565b565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b6110ad8161109b565b82525050565b5f6110be83836110a4565b60208301905092915050565b5f602082019050919050565b5f6110e082611072565b6110ea818561107c565b93506110f58361108c565b805f5b8381101561112557815161110c88826110b3565b9750611117836110ca565b9250506001810190506110f8565b5085935050505092915050565b5f6020820190508181035f83015261114a81846110d6565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61117581611163565b811461117f575f80fd5b50565b5f813590506111908161116c565b92915050565b5f602082840312156111ab576111aa61115b565b5b5f6111b884828501611182565b91505092915050565b6111ca8161109b565b82525050565b5f6020820190506111e35f8301846111c1565b92915050565b6111f28161109b565b81146111fc575f80fd5b50565b5f8135905061120d816111e9565b92915050565b5f80604083850312156112295761122861115b565b5b5f61123685828601611182565b9250506020611247858286016111ff565b9150509250929050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61129b82611255565b810181811067ffffffffffffffff821117156112ba576112b9611265565b5b80604052505050565b5f6112cc611152565b90506112d88282611292565b919050565b5f67ffffffffffffffff8211156112f7576112f6611265565b5b602082029050602081019050919050565b5f80fd5b5f61131e611319846112dd565b6112c3565b9050808382526020820190506020840283018581111561134157611340611308565b5b835b8181101561136a57806113568882611182565b845260208401935050602081019050611343565b5050509392505050565b5f82601f83011261138857611387611251565b5b813561139884826020860161130c565b91505092915050565b5f602082840312156113b6576113b561115b565b5b5f82013567ffffffffffffffff8111156113d3576113d261115f565b5b6113df84828501611374565b91505092915050565b5f8115159050919050565b6113fc816113e8565b82525050565b5f6020820190506114155f8301846113f3565b92915050565b5f67ffffffffffffffff82111561143557611434611265565b5b602082029050602081019050919050565b5f6114586114538461141b565b6112c3565b9050808382526020820190506020840283018581111561147b5761147a611308565b5b835b818110156114a4578061149088826111ff565b84526020840193505060208101905061147d565b5050509392505050565b5f82601f8301126114c2576114c1611251565b5b81356114d2848260208601611446565b91505092915050565b5f805f606084860312156114f2576114f161115b565b5b5f84013567ffffffffffffffff81111561150f5761150e61115f565b5b61151b86828701611374565b935050602084013567ffffffffffffffff81111561153c5761153b61115f565b5b61154886828701611374565b925050604084013567ffffffffffffffff8111156115695761156861115f565b5b611575868287016114ae565b9150509250925092565b61158881611163565b82525050565b5f6020820190506115a15f83018461157f565b92915050565b5f80604083850312156115bd576115bc61115b565b5b5f83013567ffffffffffffffff8111156115da576115d961115f565b5b6115e685828601611374565b925050602083013567ffffffffffffffff8111156116075761160661115f565b5b611613858286016114ae565b9150509250929050565b5f67ffffffffffffffff82111561163757611636611265565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff82111561166657611665611265565b5b61166f82611255565b9050602081019050919050565b828183375f83830152505050565b5f61169c6116978461164c565b6112c3565b9050828152602081018484840111156116b8576116b7611648565b5b6116c384828561167c565b509392505050565b5f82601f8301126116df576116de611251565b5b81356116ef84826020860161168a565b91505092915050565b5f61170a6117058461161d565b6112c3565b9050808382526020820190506020840283018581111561172d5761172c611308565b5b835b8181101561177457803567ffffffffffffffff81111561175257611751611251565b5b80860161175f89826116cb565b8552602085019450505060208101905061172f565b5050509392505050565b5f82601f83011261179257611791611251565b5b81356117a28482602086016116f8565b91505092915050565b5f602082840312156117c0576117bf61115b565b5b5f82013567ffffffffffffffff8111156117dd576117dc61115f565b5b6117e98482850161177e565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015611852578082015181840152602081019050611837565b5f8484015250505050565b5f6118678261181b565b6118718185611825565b9350611881818560208601611835565b61188a81611255565b840191505092915050565b5f6118a0838361185d565b905092915050565b5f602082019050919050565b5f6118be826117f2565b6118c881856117fc565b9350836020820285016118da8561180c565b805f5b8581101561191557848403895281516118f68582611895565b9450611901836118a8565b925060208a019950506001810190506118dd565b50829750879550505050505092915050565b5f6020820190508181035f83015261193f81846118b4565b905092915050565b5f82825260208201905092915050565b5f6119618261181b565b61196b8185611947565b935061197b818560208601611835565b61198481611255565b840191505092915050565b5f6020820190508181035f8301526119a78184611957565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611a1382611163565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611a4557611a446119dc565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611a9457607f821691505b602082108103611aa757611aa6611a50565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611b097fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611ace565b611b138683611ace565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611b4e611b49611b4484611163565b611b2b565b611163565b9050919050565b5f819050919050565b611b6783611b34565b611b7b611b7382611b55565b848454611ada565b825550505050565b5f90565b611b8f611b83565b611b9a818484611b5e565b505050565b5b81811015611bbd57611bb25f82611b87565b600181019050611ba0565b5050565b601f821115611c0257611bd381611aad565b611bdc84611abf565b81016020851015611beb578190505b611bff611bf785611abf565b830182611b9f565b50505b505050565b5f82821c905092915050565b5f611c225f1984600802611c07565b1980831691505092915050565b5f611c3a8383611c13565b9150826002028217905092915050565b611c538261181b565b67ffffffffffffffff811115611c6c57611c6b611265565b5b611c768254611a7d565b611c81828285611bc1565b5f60209050601f831160018114611cb2575f8415611ca0578287015190505b611caa8582611c2f565b865550611d11565b601f198416611cc086611aad565b5f5b82811015611ce757848901518255600182019150602085019450602081019050611cc2565b86831015611d045784890151611d00601f891682611c13565b8355505b6001600288020188555050505b505050505050565b5f611d2382611163565b9150611d2e83611163565b9250828201905080821115611d4657611d456119dc565b5b92915050565b5f819050919050565b611d66611d618261109b565b611d4c565b82525050565b5f611d778284611d55565b60208201915081905092915050565b5f81519050919050565b5f82825260208201905092915050565b5f611daa82611d86565b611db48185611d90565b9350611dc4818560208601611835565b611dcd81611255565b840191505092915050565b5f6020820190508181035f830152611df08184611da0565b905092915050565b5f6040820190508181035f830152611e108185611da0565b9050611e1f602083018461157f565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea2646970667358221220e459c84bbcd3bcc79e89b8b0af08f30b744e52144fa0aa0a906a3a414dcfadf864736f6c63430008150033",
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
