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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkRevocationStatusInMerkleTreeAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInMTAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfHashFunctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testRevocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_values\",\"type\":\"bytes32[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verificationPhase2Old\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verificationPhase2TestOld\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061271c8061005d5f395ff3fe608060405234801561000f575f80fd5b506004361061012a575f3560e01c806351600698116100ab578063bbb7e2ef1161006f578063bbb7e2ef1461032c578063cc70dd151461034a578063ce4b3f341461037a578063d3e5304214610396578063e5440a9b146103a05761012a565b8063516006981461028a5780638421b2d0146102a65780639ddf3a63146102c2578063ad108b75146102f2578063b0d2b15e146103225761012a565b80632eb4a7ab116100f25780632eb4a7ab146101be578063309ddb30146101dc578063376a65901461020c57806348db53361461022a5780634993101b1461025a5761012a565b80630de54b851461012e5780630df0ff901461013857806317a304921461016857806319f26717146101845780632337db35146101a0575b5f80fd5b6101366103d0565b005b610152600480360381019061014d91906119dc565b6103d2565b60405161015f9190611a1f565b60405180910390f35b610182600480360381019061017d9190611b76565b6103e7565b005b61019e60048036038101906101999190611bcb565b6104a0565b005b6101a8610512565b6040516101b59190611a1f565b60405180910390f35b6101c661052a565b6040516101d39190611a1f565b60405180910390f35b6101f660048036038101906101f19190611b76565b610530565b6040516102039190611c23565b60405180910390f35b610214610541565b6040516102219190611c4b565b60405180910390f35b610244600480360381019061023f91906119dc565b610546565b6040516102519190611c23565b60405180910390f35b610274600480360381019061026f9190611d24565b610562565b6040516102819190611c23565b60405180910390f35b6102a4600480360381019061029f9190611e3e565b61058c565b005b6102c060048036038101906102bb9190611eb4565b61085e565b005b6102dc60048036038101906102d79190611d24565b6108ce565b6040516102e99190611c23565b60405180910390f35b61030c60048036038101906103079190611d24565b61092d565b6040516103199190611c23565b60405180910390f35b61032a610940565b005b610334611410565b6040516103419190611a1f565b60405180910390f35b610364600480360381019061035f9190611b76565b611428565b6040516103719190611c23565b60405180910390f35b610394600480360381019061038f9190611e3e565b61149c565b005b61039e611502565b005b6103ba60048036038101906103b591906119dc565b6115f9565b6040516103c79190611c23565b60405180910390f35b565b6001602052805f5260405f205f915090505481565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461043f575f80fd5b5f5b600481101561049c5760015f8084846004811061046157610460611f3c565b5b602002015181526020019081526020015f205f6101000a81548160ff021916908315150217905550808061049490611f96565b915050610441565b5050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104f8575f80fd5b8060015f8481526020019081526020015f20819055505050565b5f60015f8081526020019081526020015f2054905090565b60025481565b5f61053a82611428565b9050919050565b600481565b5f602052805f5260405f205f915054906101000a900460ff1681565b5f806105808360015f8081526020019081526020015f205486611616565b90508091505092915050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105e4575f80fd5b80518251146105f1575f80fd5b5f5b82518110156106595781818151811061060f5761060e611f3c565b5b602002602001015160015f85848151811061062d5761062c611f3c565b5b602002602001015181526020019081526020015f2081905550808061065190611f96565b9150506105f3565b505f5b825181101561073e575f151560045f85848151811061067e5761067d611f3c565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff1615150361072b57600160045f8584815181106106bf576106be611f3c565b5b602002602001015181526020019081526020015f205f6101000a81548160ff02191690831515021790555060038382815181106106ff576106fe611f3c565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b808061073690611f96565b91505061065c565b5060015f8081526020019081526020015f20546002819055507f31f3d2d784ee13ea1252b844fa60f0be609e591b43ab4ace2cade6d6d4525f5b6003825f8151811061078d5761078c611f3c565b5b60200260200101515f602081106107a7576107a6611f3c565b5b1a60f81b835f815181106107be576107bd611f3c565b5b60200260200101516001602081106107d9576107d8611f3c565b5b1a60f81b845f815181106107f0576107ef611f3c565b5b602002602001015160026020811061080b5761080a611f3c565b5b1a60f81b855f8151811061082257610821611f3c565b5b602002602001015160036020811061083d5761083c611f3c565b5b1a60f81b604051610852959493929190612117565b60405180910390a15050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108b6575f80fd5b6108bf836103e7565b6108c9828261058c565b505050565b5f7f518b0f6d8d08ac9bfd6000c702f772cb73ca7c04cbb7dc2421a005c821a2dcc960015f8081526020019081526020015f2054848460405161091393929190612226565b60405180910390a16109258383610562565b905092915050565b5f6109388383610562565b905092915050565b61094861196c565b5f600190505b6004811015610987578082826004811061096b5761096a611f3c565b5b602002018181525050808061097f90611f96565b91505061094e565b505f604051602001610998906122bc565b6040516020818303038152906040528051906020012090505f6040516020016109c090612324565b6040516020818303038152906040528051906020012090505f6040516020016109e89061238c565b6040516020818303038152906040528051906020012090505f604051602001610a10906123f4565b6040516020818303038152906040528051906020012090505f8484604051602001610a3c929190612432565b6040516020818303038152906040528051906020012090505f8383604051602001610a68929190612432565b6040516020818303038152906040528051906020012090505f8282604051602001610a94929190612432565b6040516020818303038152906040528051906020012090505f600767ffffffffffffffff811115610ac857610ac7611a4c565b5b604051908082528060200260200182016040528015610af65781602001602082028036833780820191505090505b5090505f815f81518110610b0d57610b0c611f3c565b5b602002602001018181525050600181600181518110610b2f57610b2e611f3c565b5b602002602001018181525050600281600281518110610b5157610b50611f3c565b5b602002602001018181525050600381600381518110610b7357610b72611f3c565b5b602002602001018181525050600481600481518110610b9557610b94611f3c565b5b602002602001018181525050600581600581518110610bb757610bb6611f3c565b5b602002602001018181525050600681600681518110610bd957610bd8611f3c565b5b6020026020010181815250505f600767ffffffffffffffff811115610c0157610c00611a4c565b5b604051908082528060200260200182016040528015610c2f5781602001602082028036833780820191505090505b50905082815f81518110610c4657610c45611f3c565b5b6020026020010181815250508481600181518110610c6757610c66611f3c565b5b6020026020010181815250508381600281518110610c8857610c87611f3c565b5b6020026020010181815250508881600381518110610ca957610ca8611f3c565b5b6020026020010181815250508781600481518110610cca57610cc9611f3c565b5b6020026020010181815250508681600581518110610ceb57610cea611f3c565b5b6020026020010181815250508581600681518110610d0c57610d0b611f3c565b5b602002602001018181525050600115156001151503610d8557610d636040518060400160405280602081526020017f69737375696e67205643733a205643312c205643322c205643332c205643333481525061162c565b610d8460405180606001604052806035815260200161268f6035913961162c565b5b610d8f828261149c565b50505f600267ffffffffffffffff811115610dad57610dac611a4c565b5b604051908082528060200260200182016040528015610ddb5781602001602082028036833780820191505090505b50905086815f81518110610df257610df1611f3c565b5b6020026020010181815250508281600181518110610e1357610e12611f3c565b5b602002602001018181525050600115156001151503610e6b57610e6a6040518060400160405280601381526020017f6e6f772076657269666979696e6720766331200000000000000000000000000081525061162c565b5b600115156001151503610fe6575f610e828a610530565b90505f81610ec5576040518060400160405280601081526020017f70726f6261626c79207265766f6b656400000000000000000000000000000000815250610efc565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b9050600115156001151503610f2e57610f2d60405180606001604052806023815260200161266c60239139826116c5565b5b610f388a8461092d565b915081610f7a576040518060400160405280600881526020017f207265766f6b6564000000000000000000000000000000000000000000000000815250610fb1565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b9050600115156001151503610fe357610fe26040518060600160405280602381526020016126c460239139826116c5565b5b50505b600115156001151503611032576110316040518060400160405280601681526020017f6e6f77207265766f6b696e67205643312e2e2e2e2e2e0000000000000000000081525061162c565b5b5f604051602001611042906124a7565b604051602081830303815290604052805190602001209050808860405160200161106d929190612432565b6040516020818303038152906040528051906020012094508484604051602001611098929190612432565b6040516020818303038152906040528051906020012092505f600367ffffffffffffffff8111156110cc576110cb611a4c565b5b6040519080825280602002602001820160405280156110fa5781602001602082028036833780820191505090505b5090505f600367ffffffffffffffff81111561111957611118611a4c565b5b6040519080825280602002602001820160405280156111475781602001602082028036833780820191505090505b5090505f825f8151811061115e5761115d611f3c565b5b60200260200101818152505084815f8151811061117e5761117d611f3c565b5b6020026020010181815250506001826001815181106111a05761119f611f3c565b5b60200260200101818152505086816001815181106111c1576111c0611f3c565b5b6020026020010181815250506003826002815181106111e3576111e2611f3c565b5b602002602001018181525050828160028151811061120457611203611f3c565b5b60200260200101818152505061121b8c838361085e565b5050505f6112288a610530565b90505f8161126b576040518060400160405280601081526020017f70726f6261626c79207265766f6b6564000000000000000000000000000000008152506112a2565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b90506001151560011515036112d4576112d360405180606001604052806023815260200161266c60239139826116c5565b5b5f6112df8b8561092d565b905080611321576040518060400160405280600881526020017f207265766f6b6564000000000000000000000000000000000000000000000000815250611358565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b915060011515600115150361138a576113896040518060600160405280602381526020016126c460239139836116c5565b5b5f611393611410565b6040516020016113a391906124c5565b6040516020818303038152906040529050600115156001151503611401576114006040518060400160405280600d81526020017f6d65726b6c6520726f6f743a2000000000000000000000000000000000000000815250826116c5565b5b50505050505050505050505050565b5f60015f8081526020019081526020015f2054905090565b5f805f90505f5b6004811015611492575f15155f808684600481106114505761144f611f3c565b5b602002015181526020019081526020015f205f9054906101000a900460ff1615150361147f5760019150611492565b808061148a90611f96565b91505061142f565b5080915050919050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114f4575f80fd5b6114fe828261058c565b5050565b60011515600115150361154e5761154d6040518060400160405280601381526020017f70726974696e67206d65726b6c6520747265650000000000000000000000000081525061162c565b5b5f5b6003805490508110156115f6576001151560011515036115e3576115a96040518060400160405280601481526020017f696e646578203a20256420092076616c7565203a00000000000000000000000081525082611761565b6115e260015f8381526020019081526020015f20546040516020016115ce91906124c5565b6040516020818303038152906040526117fd565b5b80806115ee90611f96565b915050611550565b50565b6004602052805f5260405f205f915054906101000a900460ff1681565b5f826116228584611896565b1490509392505050565b6116c2816040516024016116409190612549565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118ea565b50565b61175d82826040516024016116db929190612569565b6040516020818303038152906040527f4b5c4277000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118ea565b5050565b6117f9828260405160240161177792919061259e565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118ea565b5050565b61189381604051602401611811919061261e565b6040516020818303038152906040527f0be77f56000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118ea565b50565b5f808290505f5b84518110156118df576118ca828683815181106118bd576118bc611f3c565b5b6020026020010151611904565b915080806118d790611f96565b91505061189d565b508091505092915050565b611901816118f961192e61194d565b63ffffffff16565b50565b5f81831061191b576119168284611958565b611926565b6119258383611958565b5b905092915050565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b61198e819050919050565b5f825f528160205260405f20905092915050565b6040518060800160405280600490602082028036833780820191505090505090565b61199661263e565b565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6119bb816119a9565b81146119c5575f80fd5b50565b5f813590506119d6816119b2565b92915050565b5f602082840312156119f1576119f06119a1565b5b5f6119fe848285016119c8565b91505092915050565b5f819050919050565b611a1981611a07565b82525050565b5f602082019050611a325f830184611a10565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611a8282611a3c565b810181811067ffffffffffffffff82111715611aa157611aa0611a4c565b5b80604052505050565b5f611ab3611998565b9050611abf8282611a79565b919050565b5f67ffffffffffffffff821115611ade57611add611a4c565b5b602082029050919050565b5f80fd5b5f611aff611afa84611ac4565b611aaa565b90508060208402830185811115611b1957611b18611ae9565b5b835b81811015611b425780611b2e88826119c8565b845260208401935050602081019050611b1b565b5050509392505050565b5f82601f830112611b6057611b5f611a38565b5b6004611b6d848285611aed565b91505092915050565b5f60808284031215611b8b57611b8a6119a1565b5b5f611b9884828501611b4c565b91505092915050565b611baa81611a07565b8114611bb4575f80fd5b50565b5f81359050611bc581611ba1565b92915050565b5f8060408385031215611be157611be06119a1565b5b5f611bee858286016119c8565b9250506020611bff85828601611bb7565b9150509250929050565b5f8115159050919050565b611c1d81611c09565b82525050565b5f602082019050611c365f830184611c14565b92915050565b611c45816119a9565b82525050565b5f602082019050611c5e5f830184611c3c565b92915050565b5f67ffffffffffffffff821115611c7e57611c7d611a4c565b5b602082029050602081019050919050565b5f611ca1611c9c84611c64565b611aaa565b90508083825260208201905060208402830185811115611cc457611cc3611ae9565b5b835b81811015611ced5780611cd98882611bb7565b845260208401935050602081019050611cc6565b5050509392505050565b5f82601f830112611d0b57611d0a611a38565b5b8135611d1b848260208601611c8f565b91505092915050565b5f8060408385031215611d3a57611d396119a1565b5b5f611d4785828601611bb7565b925050602083013567ffffffffffffffff811115611d6857611d676119a5565b5b611d7485828601611cf7565b9150509250929050565b5f67ffffffffffffffff821115611d9857611d97611a4c565b5b602082029050602081019050919050565b5f611dbb611db684611d7e565b611aaa565b90508083825260208201905060208402830185811115611dde57611ddd611ae9565b5b835b81811015611e075780611df388826119c8565b845260208401935050602081019050611de0565b5050509392505050565b5f82601f830112611e2557611e24611a38565b5b8135611e35848260208601611da9565b91505092915050565b5f8060408385031215611e5457611e536119a1565b5b5f83013567ffffffffffffffff811115611e7157611e706119a5565b5b611e7d85828601611e11565b925050602083013567ffffffffffffffff811115611e9e57611e9d6119a5565b5b611eaa85828601611cf7565b9150509250929050565b5f805f60c08486031215611ecb57611eca6119a1565b5b5f611ed886828701611b4c565b935050608084013567ffffffffffffffff811115611ef957611ef86119a5565b5b611f0586828701611e11565b92505060a084013567ffffffffffffffff811115611f2657611f256119a5565b5b611f3286828701611cf7565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fa0826119a9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611fd257611fd1611f69565b5b600182019050919050565b5f81549050919050565b5f82825260208201905092915050565b5f819050815f5260205f209050919050565b612012816119a9565b82525050565b5f6120238383612009565b60208301905092915050565b5f815f1c9050919050565b5f819050919050565b5f6120556120508361202f565b61203a565b9050919050565b5f6120678254612043565b9050919050565b5f600182019050919050565b5f61208482611fdd565b61208e8185611fe7565b935061209983611ff7565b805f5b838110156120d0576120ad8261205c565b6120b78882612018565b97506120c28361206e565b92505060018101905061209c565b5085935050505092915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b612111816120dd565b82525050565b5f60a0820190508181035f83015261212f818861207a565b905061213e6020830187612108565b61214b6040830186612108565b6121586060830185612108565b6121656080830184612108565b9695505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6121a181611a07565b82525050565b5f6121b28383612198565b60208301905092915050565b5f602082019050919050565b5f6121d48261216f565b6121de8185612179565b93506121e983612189565b805f5b8381101561221957815161220088826121a7565b975061220b836121be565b9250506001810190506121ec565b5085935050505092915050565b5f6060820190506122395f830186611a10565b6122466020830185611a10565b818103604083015261225881846121ca565b9050949350505050565b5f82825260208201905092915050565b7f76633100000000000000000000000000000000000000000000000000000000005f82015250565b5f6122a6600383612262565b91506122b182612272565b602082019050919050565b5f6020820190508181035f8301526122d38161229a565b9050919050565b7f76633200000000000000000000000000000000000000000000000000000000005f82015250565b5f61230e600383612262565b9150612319826122da565b602082019050919050565b5f6020820190508181035f83015261233b81612302565b9050919050565b7f76633300000000000000000000000000000000000000000000000000000000005f82015250565b5f612376600383612262565b915061238182612342565b602082019050919050565b5f6020820190508181035f8301526123a38161236a565b9050919050565b7f76633400000000000000000000000000000000000000000000000000000000005f82015250565b5f6123de600383612262565b91506123e9826123aa565b602082019050919050565b5f6020820190508181035f83015261240b816123d2565b9050919050565b5f819050919050565b61242c61242782611a07565b612412565b82525050565b5f61243d828561241b565b60208201915061244d828461241b565b6020820191508190509392505050565b7f76632031207265766f6b656400000000000000000000000000000000000000005f82015250565b5f612491600c83612262565b915061249c8261245d565b602082019050919050565b5f6020820190508181035f8301526124be81612485565b9050919050565b5f6124d0828461241b565b60208201915081905092915050565b5f81519050919050565b5f5b838110156125065780820151818401526020810190506124eb565b5f8484015250505050565b5f61251b826124df565b6125258185612262565b93506125358185602086016124e9565b61253e81611a3c565b840191505092915050565b5f6020820190508181035f8301526125618184612511565b905092915050565b5f6040820190508181035f8301526125818185612511565b905081810360208301526125958184612511565b90509392505050565b5f6040820190508181035f8301526125b68185612511565b90506125c56020830184611c3c565b9392505050565b5f81519050919050565b5f82825260208201905092915050565b5f6125f0826125cc565b6125fa81856125d6565b935061260a8185602086016124e9565b61261381611a3c565b840191505092915050565b5f6020820190508181035f83015261263681846125e6565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfe706861736520313a207265766f636174696f6e20737461747573206f66205643313a206d65726b6c65207472656520616363756d756c61746f7220697320696e697469616c697a656420776974682076616c696420766373706861736520323a207265766f636174696f6e20737461747573206f66205643313a20a264697066735822122088a0da8d9faf7b04ecb2db6ff901ef7244bfdd0dcfabb56a2611212d2db18fd064736f6c63430008150033",
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

// CheckRevocationStatusInMerkleTreeAccumulator is a free data retrieval call binding the contract method 0x4993101b.
//
// Solidity: function checkRevocationStatusInMerkleTreeAccumulator(bytes32 leaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceCaller) CheckRevocationStatusInMerkleTreeAccumulator(opts *bind.CallOpts, leaf [32]byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "checkRevocationStatusInMerkleTreeAccumulator", leaf, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRevocationStatusInMerkleTreeAccumulator is a free data retrieval call binding the contract method 0x4993101b.
//
// Solidity: function checkRevocationStatusInMerkleTreeAccumulator(bytes32 leaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceSession) CheckRevocationStatusInMerkleTreeAccumulator(leaf [32]byte, proof [][32]byte) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInMerkleTreeAccumulator(&_RevocationService.CallOpts, leaf, proof)
}

// CheckRevocationStatusInMerkleTreeAccumulator is a free data retrieval call binding the contract method 0x4993101b.
//
// Solidity: function checkRevocationStatusInMerkleTreeAccumulator(bytes32 leaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) CheckRevocationStatusInMerkleTreeAccumulator(leaf [32]byte, proof [][32]byte) (bool, error) {
	return _RevocationService.Contract.CheckRevocationStatusInMerkleTreeAccumulator(&_RevocationService.CallOpts, leaf, proof)
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

// VerificationPhase2Old is a free data retrieval call binding the contract method 0xad108b75.
//
// Solidity: function verificationPhase2Old(bytes32 vcLeaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceCaller) VerificationPhase2Old(opts *bind.CallOpts, vcLeaf [32]byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _RevocationService.contract.Call(opts, &out, "verificationPhase2Old", vcLeaf, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificationPhase2Old is a free data retrieval call binding the contract method 0xad108b75.
//
// Solidity: function verificationPhase2Old(bytes32 vcLeaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceSession) VerificationPhase2Old(vcLeaf [32]byte, proof [][32]byte) (bool, error) {
	return _RevocationService.Contract.VerificationPhase2Old(&_RevocationService.CallOpts, vcLeaf, proof)
}

// VerificationPhase2Old is a free data retrieval call binding the contract method 0xad108b75.
//
// Solidity: function verificationPhase2Old(bytes32 vcLeaf, bytes32[] proof) view returns(bool)
func (_RevocationService *RevocationServiceCallerSession) VerificationPhase2Old(vcLeaf [32]byte, proof [][32]byte) (bool, error) {
	return _RevocationService.Contract.VerificationPhase2Old(&_RevocationService.CallOpts, vcLeaf, proof)
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

// RevokeVC is a paid mutator transaction binding the contract method 0x8421b2d0.
//
// Solidity: function revokeVC(uint256[4] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactor) RevokeVC(opts *bind.TransactOpts, _bfIndexes [4]*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "revokeVC", _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x8421b2d0.
//
// Solidity: function revokeVC(uint256[4] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceSession) RevokeVC(_bfIndexes [4]*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// RevokeVC is a paid mutator transaction binding the contract method 0x8421b2d0.
//
// Solidity: function revokeVC(uint256[4] _bfIndexes, uint256[] _mtIndexes, bytes32[] _mtValues) returns()
func (_RevocationService *RevocationServiceTransactorSession) RevokeVC(_bfIndexes [4]*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.RevokeVC(&_RevocationService.TransactOpts, _bfIndexes, _mtIndexes, _mtValues)
}

// TestRevocation is a paid mutator transaction binding the contract method 0xb0d2b15e.
//
// Solidity: function testRevocation() returns()
func (_RevocationService *RevocationServiceTransactor) TestRevocation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "testRevocation")
}

// TestRevocation is a paid mutator transaction binding the contract method 0xb0d2b15e.
//
// Solidity: function testRevocation() returns()
func (_RevocationService *RevocationServiceSession) TestRevocation() (*types.Transaction, error) {
	return _RevocationService.Contract.TestRevocation(&_RevocationService.TransactOpts)
}

// TestRevocation is a paid mutator transaction binding the contract method 0xb0d2b15e.
//
// Solidity: function testRevocation() returns()
func (_RevocationService *RevocationServiceTransactorSession) TestRevocation() (*types.Transaction, error) {
	return _RevocationService.Contract.TestRevocation(&_RevocationService.TransactOpts)
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

// VerificationPhase2Test is a paid mutator transaction binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() returns(bytes32)
func (_RevocationService *RevocationServiceTransactor) VerificationPhase2Test(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "verificationPhase2Test")
}

// VerificationPhase2Test is a paid mutator transaction binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() returns(bytes32)
func (_RevocationService *RevocationServiceSession) VerificationPhase2Test() (*types.Transaction, error) {
	return _RevocationService.Contract.VerificationPhase2Test(&_RevocationService.TransactOpts)
}

// VerificationPhase2Test is a paid mutator transaction binding the contract method 0x2337db35.
//
// Solidity: function verificationPhase2Test() returns(bytes32)
func (_RevocationService *RevocationServiceTransactorSession) VerificationPhase2Test() (*types.Transaction, error) {
	return _RevocationService.Contract.VerificationPhase2Test(&_RevocationService.TransactOpts)
}

// VerificationPhase2TestOld is a paid mutator transaction binding the contract method 0x9ddf3a63.
//
// Solidity: function verificationPhase2TestOld(bytes32 vcLeaf, bytes32[] proof) returns(bool)
func (_RevocationService *RevocationServiceTransactor) VerificationPhase2TestOld(opts *bind.TransactOpts, vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _RevocationService.contract.Transact(opts, "verificationPhase2TestOld", vcLeaf, proof)
}

// VerificationPhase2TestOld is a paid mutator transaction binding the contract method 0x9ddf3a63.
//
// Solidity: function verificationPhase2TestOld(bytes32 vcLeaf, bytes32[] proof) returns(bool)
func (_RevocationService *RevocationServiceSession) VerificationPhase2TestOld(vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.VerificationPhase2TestOld(&_RevocationService.TransactOpts, vcLeaf, proof)
}

// VerificationPhase2TestOld is a paid mutator transaction binding the contract method 0x9ddf3a63.
//
// Solidity: function verificationPhase2TestOld(bytes32 vcLeaf, bytes32[] proof) returns(bool)
func (_RevocationService *RevocationServiceTransactorSession) VerificationPhase2TestOld(vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _RevocationService.Contract.VerificationPhase2TestOld(&_RevocationService.TransactOpts, vcLeaf, proof)
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
