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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkRevocationStatusInMerkleTreeAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInMTAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfHashFunctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_mtValues\",\"type\":\"bytes32[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testRevocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_values\",\"type\":\"bytes32[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verificationPhase2Old\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"verificationPhase2TestOld\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506127988061005d5f395ff3fe608060405234801561000f575f80fd5b506004361061011f575f3560e01c806351600698116100ab578063bbb7e2ef1161006f578063bbb7e2ef14610303578063cc70dd1514610321578063ce4b3f3414610351578063d3e530421461036d578063e5440a9b146103775761011f565b806351600698146102615780638421b2d01461027d5780639ddf3a6314610299578063ad108b75146102c9578063b0d2b15e146102f95761011f565b80632eb4a7ab116100f25780632eb4a7ab14610195578063309ddb30146101b3578063376a6590146101e357806348db5336146102015780634993101b146102315761011f565b80630de54b85146101235780630df0ff901461012d57806317a304921461015d57806319f2671714610179575b5f80fd5b61012b6103a7565b005b61014760048036038101906101429190611a58565b6103a9565b6040516101549190611a9b565b60405180910390f35b61017760048036038101906101729190611bf2565b6103be565b005b610193600480360381019061018e9190611c47565b610477565b005b61019d6104e9565b6040516101aa9190611a9b565b60405180910390f35b6101cd60048036038101906101c89190611bf2565b6104ef565b6040516101da9190611c9f565b60405180910390f35b6101eb610500565b6040516101f89190611cc7565b60405180910390f35b61021b60048036038101906102169190611a58565b610505565b6040516102289190611c9f565b60405180910390f35b61024b60048036038101906102469190611da0565b610521565b6040516102589190611c9f565b60405180910390f35b61027b60048036038101906102769190611eba565b61054b565b005b61029760048036038101906102929190611f30565b61081d565b005b6102b360048036038101906102ae9190611da0565b61088d565b6040516102c09190611c9f565b60405180910390f35b6102e360048036038101906102de9190611da0565b6108ec565b6040516102f09190611c9f565b60405180910390f35b6103016108ff565b005b61030b6113b8565b6040516103189190611a9b565b60405180910390f35b61033b60048036038101906103369190611bf2565b6113d0565b6040516103489190611c9f565b60405180910390f35b61036b60048036038101906103669190611eba565b611445565b005b6103756114ab565b005b610391600480360381019061038c9190611a58565b6115a2565b60405161039e9190611c9f565b60405180910390f35b565b6001602052805f5260405f205f915090505481565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610416575f80fd5b5f5b60048110156104735760015f8084846004811061043857610437611fb8565b5b602002015181526020019081526020015f205f6101000a81548160ff021916908315150217905550808061046b90612012565b915050610418565b5050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104cf575f80fd5b8060015f8481526020019081526020015f20819055505050565b60025481565b5f6104f9826113d0565b9050919050565b600481565b5f602052805f5260405f205f915054906101000a900460ff1681565b5f8061053f8360015f8081526020019081526020015f2054866115bf565b90508091505092915050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105a3575f80fd5b80518251146105b0575f80fd5b5f5b8251811015610618578181815181106105ce576105cd611fb8565b5b602002602001015160015f8584815181106105ec576105eb611fb8565b5b602002602001015181526020019081526020015f2081905550808061061090612012565b9150506105b2565b505f5b82518110156106fd575f151560045f85848151811061063d5761063c611fb8565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff161515036106ea57600160045f85848151811061067e5761067d611fb8565b5b602002602001015181526020019081526020015f205f6101000a81548160ff02191690831515021790555060038382815181106106be576106bd611fb8565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b80806106f590612012565b91505061061b565b5060015f8081526020019081526020015f20546002819055507f31f3d2d784ee13ea1252b844fa60f0be609e591b43ab4ace2cade6d6d4525f5b6003825f8151811061074c5761074b611fb8565b5b60200260200101515f6020811061076657610765611fb8565b5b1a60f81b835f8151811061077d5761077c611fb8565b5b602002602001015160016020811061079857610797611fb8565b5b1a60f81b845f815181106107af576107ae611fb8565b5b60200260200101516002602081106107ca576107c9611fb8565b5b1a60f81b855f815181106107e1576107e0611fb8565b5b60200260200101516003602081106107fc576107fb611fb8565b5b1a60f81b604051610811959493929190612193565b60405180910390a15050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610875575f80fd5b61087e836103be565b610888828261054b565b505050565b5f7f518b0f6d8d08ac9bfd6000c702f772cb73ca7c04cbb7dc2421a005c821a2dcc960015f8081526020019081526020015f205484846040516108d2939291906122a2565b60405180910390a16108e48383610521565b905092915050565b5f6108f78383610521565b905092915050565b6109076119e8565b5f600190505b6004811015610946578082826004811061092a57610929611fb8565b5b602002018181525050808061093e90612012565b91505061090d565b505f60405160200161095790612338565b6040516020818303038152906040528051906020012090505f60405160200161097f906123a0565b6040516020818303038152906040528051906020012090505f6040516020016109a790612408565b6040516020818303038152906040528051906020012090505f6040516020016109cf90612470565b6040516020818303038152906040528051906020012090505f84846040516020016109fb9291906124ae565b6040516020818303038152906040528051906020012090505f8383604051602001610a279291906124ae565b6040516020818303038152906040528051906020012090505f8282604051602001610a539291906124ae565b6040516020818303038152906040528051906020012090505f600767ffffffffffffffff811115610a8757610a86611ac8565b5b604051908082528060200260200182016040528015610ab55781602001602082028036833780820191505090505b5090505f815f81518110610acc57610acb611fb8565b5b602002602001018181525050600181600181518110610aee57610aed611fb8565b5b602002602001018181525050600281600281518110610b1057610b0f611fb8565b5b602002602001018181525050600381600381518110610b3257610b31611fb8565b5b602002602001018181525050600481600481518110610b5457610b53611fb8565b5b602002602001018181525050600581600581518110610b7657610b75611fb8565b5b602002602001018181525050600681600681518110610b9857610b97611fb8565b5b6020026020010181815250505f600767ffffffffffffffff811115610bc057610bbf611ac8565b5b604051908082528060200260200182016040528015610bee5781602001602082028036833780820191505090505b50905082815f81518110610c0557610c04611fb8565b5b6020026020010181815250508481600181518110610c2657610c25611fb8565b5b6020026020010181815250508381600281518110610c4757610c46611fb8565b5b6020026020010181815250508881600381518110610c6857610c67611fb8565b5b6020026020010181815250508781600481518110610c8957610c88611fb8565b5b6020026020010181815250508681600581518110610caa57610ca9611fb8565b5b6020026020010181815250508581600681518110610ccb57610cca611fb8565b5b602002602001018181525050600115156001151503610d4457610d226040518060400160405280602081526020017f69737375696e67205643733a205643312c205643322c205643332c20564333348152506115d5565b610d4360405180606001604052806035815260200161270b603591396115d5565b5b610d4e8282611445565b50505f600267ffffffffffffffff811115610d6c57610d6b611ac8565b5b604051908082528060200260200182016040528015610d9a5781602001602082028036833780820191505090505b50905086815f81518110610db157610db0611fb8565b5b6020026020010181815250508281600181518110610dd257610dd1611fb8565b5b602002602001018181525050600115156001151503610e2a57610e296040518060400160405280601381526020017f6e6f772076657269666979696e672076633120000000000000000000000000008152506115d5565b5b600115156001151503610fa5575f610e418a6104ef565b90505f81610e84576040518060400160405280601081526020017f70726f6261626c79207265766f6b656400000000000000000000000000000000815250610ebb565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b9050600115156001151503610eed57610eec6040518060600160405280602381526020016126e8602391398261166e565b5b610ef78a846108ec565b915081610f39576040518060400160405280600881526020017f207265766f6b6564000000000000000000000000000000000000000000000000815250610f70565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b9050600115156001151503610fa257610fa1604051806060016040528060238152602001612740602391398261166e565b5b50505b600115156001151503610ff157610ff06040518060400160405280601681526020017f6e6f77207265766f6b696e67205643312e2e2e2e2e2e000000000000000000008152506115d5565b5b5f60405160200161100190612523565b604051602081830303815290604052805190602001209050808860405160200161102c9291906124ae565b60405160208183030381529060405280519060200120945084846040516020016110579291906124ae565b6040516020818303038152906040528051906020012092505f600367ffffffffffffffff81111561108b5761108a611ac8565b5b6040519080825280602002602001820160405280156110b95781602001602082028036833780820191505090505b5090505f600367ffffffffffffffff8111156110d8576110d7611ac8565b5b6040519080825280602002602001820160405280156111065781602001602082028036833780820191505090505b5090505f825f8151811061111d5761111c611fb8565b5b60200260200101818152505084815f8151811061113d5761113c611fb8565b5b60200260200101818152505060018260018151811061115f5761115e611fb8565b5b60200260200101818152505086816001815181106111805761117f611fb8565b5b6020026020010181815250506003826002815181106111a2576111a1611fb8565b5b60200260200101818152505082816002815181106111c3576111c2611fb8565b5b6020026020010181815250506111da8c838361081d565b5050505f6111e78a6104ef565b90505f8161122a576040518060400160405280601081526020017f70726f6261626c79207265766f6b656400000000000000000000000000000000815250611261565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b9050600115156001151503611293576112926040518060600160405280602381526020016126e8602391398261166e565b5b5f61129e8b856108ec565b9050806112e0576040518060400160405280600881526020017f207265766f6b6564000000000000000000000000000000000000000000000000815250611317565b6040518060400160405280600b81526020017f6e6f74207265766f6b65640000000000000000000000000000000000000000008152505b915060011515600115150361134957611348604051806060016040528060238152602001612740602391398361166e565b5b5f6113526113b8565b90506001151560011515036113a9576113a86040518060400160405280600d81526020017f6d65726b6c6520726f6f743a20000000000000000000000000000000000000008152506113a38361170a565b61166e565b5b50505050505050505050505050565b5f60015f8081526020019081526020015f2054905090565b5f80600190505f5b600481101561143b57600115155f808684600481106113fa576113f9611fb8565b5b602002015181526020019081526020015f205f9054906101000a900460ff16151503611428575f915061143b565b808061143390612012565b9150506113d8565b5080915050919050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461149d575f80fd5b6114a7828261054b565b5050565b6001151560011515036114f7576114f66040518060400160405280601381526020017f70726974696e67206d65726b6c652074726565000000000000000000000000008152506115d5565b5b5f5b60038054905081101561159f5760011515600115150361158c576115526040518060400160405280601481526020017f696e646578203a20256420092076616c7565203a000000000000000000000000815250826117dd565b61158b60015f8381526020019081526020015f20546040516020016115779190612541565b604051602081830303815290604052611879565b5b808061159790612012565b9150506114f9565b50565b6004602052805f5260405f205f915054906101000a900460ff1681565b5f826115cb8584611912565b1490509392505050565b61166b816040516024016115e991906125c5565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611966565b50565b61170682826040516024016116849291906125e5565b6040516020818303038152906040527f4b5c4277000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611966565b5050565b60605f602067ffffffffffffffff81111561172857611727611ac8565b5b6040519080825280601f01601f19166020018201604052801561175a5781602001600182028036833780820191505090505b5090505f5b60208110156117d35783816020811061177b5761177a611fb8565b5b1a60f81b82828151811061179257611791611fb8565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535080806117cb90612012565b91505061175f565b5080915050919050565b61187582826040516024016117f392919061261a565b6040516020818303038152906040527fb60e72cc000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611966565b5050565b61190f8160405160240161188d919061269a565b6040516020818303038152906040527f0be77f56000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050611966565b50565b5f808290505f5b845181101561195b576119468286838151811061193957611938611fb8565b5b6020026020010151611980565b9150808061195390612012565b915050611919565b508091505092915050565b61197d816119756119aa6119c9565b63ffffffff16565b50565b5f8183106119975761199282846119d4565b6119a2565b6119a183836119d4565b5b905092915050565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b611a0a819050919050565b5f825f528160205260405f20905092915050565b6040518060800160405280600490602082028036833780820191505090505090565b611a126126ba565b565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b611a3781611a25565b8114611a41575f80fd5b50565b5f81359050611a5281611a2e565b92915050565b5f60208284031215611a6d57611a6c611a1d565b5b5f611a7a84828501611a44565b91505092915050565b5f819050919050565b611a9581611a83565b82525050565b5f602082019050611aae5f830184611a8c565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611afe82611ab8565b810181811067ffffffffffffffff82111715611b1d57611b1c611ac8565b5b80604052505050565b5f611b2f611a14565b9050611b3b8282611af5565b919050565b5f67ffffffffffffffff821115611b5a57611b59611ac8565b5b602082029050919050565b5f80fd5b5f611b7b611b7684611b40565b611b26565b90508060208402830185811115611b9557611b94611b65565b5b835b81811015611bbe5780611baa8882611a44565b845260208401935050602081019050611b97565b5050509392505050565b5f82601f830112611bdc57611bdb611ab4565b5b6004611be9848285611b69565b91505092915050565b5f60808284031215611c0757611c06611a1d565b5b5f611c1484828501611bc8565b91505092915050565b611c2681611a83565b8114611c30575f80fd5b50565b5f81359050611c4181611c1d565b92915050565b5f8060408385031215611c5d57611c5c611a1d565b5b5f611c6a85828601611a44565b9250506020611c7b85828601611c33565b9150509250929050565b5f8115159050919050565b611c9981611c85565b82525050565b5f602082019050611cb25f830184611c90565b92915050565b611cc181611a25565b82525050565b5f602082019050611cda5f830184611cb8565b92915050565b5f67ffffffffffffffff821115611cfa57611cf9611ac8565b5b602082029050602081019050919050565b5f611d1d611d1884611ce0565b611b26565b90508083825260208201905060208402830185811115611d4057611d3f611b65565b5b835b81811015611d695780611d558882611c33565b845260208401935050602081019050611d42565b5050509392505050565b5f82601f830112611d8757611d86611ab4565b5b8135611d97848260208601611d0b565b91505092915050565b5f8060408385031215611db657611db5611a1d565b5b5f611dc385828601611c33565b925050602083013567ffffffffffffffff811115611de457611de3611a21565b5b611df085828601611d73565b9150509250929050565b5f67ffffffffffffffff821115611e1457611e13611ac8565b5b602082029050602081019050919050565b5f611e37611e3284611dfa565b611b26565b90508083825260208201905060208402830185811115611e5a57611e59611b65565b5b835b81811015611e835780611e6f8882611a44565b845260208401935050602081019050611e5c565b5050509392505050565b5f82601f830112611ea157611ea0611ab4565b5b8135611eb1848260208601611e25565b91505092915050565b5f8060408385031215611ed057611ecf611a1d565b5b5f83013567ffffffffffffffff811115611eed57611eec611a21565b5b611ef985828601611e8d565b925050602083013567ffffffffffffffff811115611f1a57611f19611a21565b5b611f2685828601611d73565b9150509250929050565b5f805f60c08486031215611f4757611f46611a1d565b5b5f611f5486828701611bc8565b935050608084013567ffffffffffffffff811115611f7557611f74611a21565b5b611f8186828701611e8d565b92505060a084013567ffffffffffffffff811115611fa257611fa1611a21565b5b611fae86828701611d73565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61201c82611a25565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361204e5761204d611fe5565b5b600182019050919050565b5f81549050919050565b5f82825260208201905092915050565b5f819050815f5260205f209050919050565b61208e81611a25565b82525050565b5f61209f8383612085565b60208301905092915050565b5f815f1c9050919050565b5f819050919050565b5f6120d16120cc836120ab565b6120b6565b9050919050565b5f6120e382546120bf565b9050919050565b5f600182019050919050565b5f61210082612059565b61210a8185612063565b935061211583612073565b805f5b8381101561214c57612129826120d8565b6121338882612094565b975061213e836120ea565b925050600181019050612118565b5085935050505092915050565b5f7fff0000000000000000000000000000000000000000000000000000000000000082169050919050565b61218d81612159565b82525050565b5f60a0820190508181035f8301526121ab81886120f6565b90506121ba6020830187612184565b6121c76040830186612184565b6121d46060830185612184565b6121e16080830184612184565b9695505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61221d81611a83565b82525050565b5f61222e8383612214565b60208301905092915050565b5f602082019050919050565b5f612250826121eb565b61225a81856121f5565b935061226583612205565b805f5b8381101561229557815161227c8882612223565b97506122878361223a565b925050600181019050612268565b5085935050505092915050565b5f6060820190506122b55f830186611a8c565b6122c26020830185611a8c565b81810360408301526122d48184612246565b9050949350505050565b5f82825260208201905092915050565b7f76633100000000000000000000000000000000000000000000000000000000005f82015250565b5f6123226003836122de565b915061232d826122ee565b602082019050919050565b5f6020820190508181035f83015261234f81612316565b9050919050565b7f76633200000000000000000000000000000000000000000000000000000000005f82015250565b5f61238a6003836122de565b915061239582612356565b602082019050919050565b5f6020820190508181035f8301526123b78161237e565b9050919050565b7f76633300000000000000000000000000000000000000000000000000000000005f82015250565b5f6123f26003836122de565b91506123fd826123be565b602082019050919050565b5f6020820190508181035f83015261241f816123e6565b9050919050565b7f76633400000000000000000000000000000000000000000000000000000000005f82015250565b5f61245a6003836122de565b915061246582612426565b602082019050919050565b5f6020820190508181035f8301526124878161244e565b9050919050565b5f819050919050565b6124a86124a382611a83565b61248e565b82525050565b5f6124b98285612497565b6020820191506124c98284612497565b6020820191508190509392505050565b7f76632031207265766f6b656400000000000000000000000000000000000000005f82015250565b5f61250d600c836122de565b9150612518826124d9565b602082019050919050565b5f6020820190508181035f83015261253a81612501565b9050919050565b5f61254c8284612497565b60208201915081905092915050565b5f81519050919050565b5f5b83811015612582578082015181840152602081019050612567565b5f8484015250505050565b5f6125978261255b565b6125a181856122de565b93506125b1818560208601612565565b6125ba81611ab8565b840191505092915050565b5f6020820190508181035f8301526125dd818461258d565b905092915050565b5f6040820190508181035f8301526125fd818561258d565b90508181036020830152612611818461258d565b90509392505050565b5f6040820190508181035f830152612632818561258d565b90506126416020830184611cb8565b9392505050565b5f81519050919050565b5f82825260208201905092915050565b5f61266c82612648565b6126768185612652565b9350612686818560208601612565565b61268f81611ab8565b840191505092915050565b5f6020820190508181035f8301526126b28184612662565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfe706861736520313a207265766f636174696f6e20737461747573206f66205643313a206d65726b6c65207472656520616363756d756c61746f7220697320696e697469616c697a656420776974682076616c696420766373706861736520323a207265766f636174696f6e20737461747573206f66205643313a20a26469706673582212209b83bd207d7c40e92e83093e565d7c02b873fbf4a492a8221d65e466ad0e1e8364736f6c63430008150033",
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
