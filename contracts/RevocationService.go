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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue1\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue2\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue3\",\"type\":\"bytes1\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"_mtValue4\",\"type\":\"bytes1\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcLeaf\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"VerificationPhase2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"checkRevocationStatusInBloomFilter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isExistInMTAccumulator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"issueVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleTree\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfHashFunctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printMerkleTree\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerIssuers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[]\",\"name\":\"_mtIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_mtValues\",\"type\":\"string[]\"}],\"name\":\"revokeVC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_indexes\",\"type\":\"uint256[4]\"}],\"name\":\"updateBloomFilter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indexes\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_values\",\"type\":\"string[]\"}],\"name\":\"updateMerkleTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"updateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"_bfIndexes\",\"type\":\"uint256[4]\"}],\"name\":\"verificationPhase1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verificationPhase2Test\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b503360055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506118898061005d5f395ff3fe608060405234801561000f575f80fd5b50600436106100fe575f3560e01c8063376a659011610095578063bbb7e2ef11610064578063bbb7e2ef14610282578063cc70dd15146102a0578063d3e53042146102d0578063e5440a9b146102da576100fe565b8063376a6590146101fc57806348db53361461021a57806353491af51461024a57806392ed770e14610266576100fe565b80632eb4a7ab116100d15780632eb4a7ab146101765780632ef76b1814610194578063309ddb30146101b057806334d6c77f146101e0576100fe565b80630de54b85146101025780630df0ff901461010c57806317a304921461013c5780632337db3514610158575b5f80fd5b61010a61030a565b005b61012660048036038101906101219190610d3e565b61030c565b6040516101339190610df3565b60405180910390f35b61015660048036038101906101519190610f41565b6103a7565b005b610160610460565b60405161016d9190610df3565b60405180910390f35b61017e6104ff565b60405161018b9190610df3565b60405180910390f35b6101ae60048036038101906101a991906111ba565b61058b565b005b6101ca60048036038101906101c59190610f41565b61076b565b6040516101d7919061124a565b60405180910390f35b6101fa60048036038101906101f591906111ba565b61077c565b005b6102046107e2565b6040516102119190611272565b60405180910390f35b610234600480360381019061022f9190610d3e565b6107e7565b604051610241919061124a565b60405180910390f35b610264600480360381019061025f919061128b565b610803565b005b610280600480360381019061027b91906112e5565b61087e565b005b61028a6108ee565b6040516102979190610df3565b60405180910390f35b6102ba60048036038101906102b59190610f41565b61098d565b6040516102c7919061124a565b60405180910390f35b6102d8610a01565b005b6102f460048036038101906102ef9190610d3e565b610b57565b604051610301919061124a565b60405180910390f35b565b6001602052805f5260405f205f9150905080546103289061139a565b80601f01602080910402602001604051908101604052809291908181526020018280546103549061139a565b801561039f5780601f106103765761010080835404028352916020019161039f565b820191905f5260205f20905b81548152906001019060200180831161038257829003601f168201915b505050505081565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ff575f80fd5b5f5b600481101561045c5760015f80848460048110610421576104206113ca565b5b602002015181526020019081526020015f205f6101000a81548160ff021916908315150217905550808061045490611424565b915050610401565b5050565b606060015f8081526020019081526020015f20805461047e9061139a565b80601f01602080910402602001604051908101604052809291908181526020018280546104aa9061139a565b80156104f55780601f106104cc576101008083540402835291602001916104f5565b820191905f5260205f20905b8154815290600101906020018083116104d857829003601f168201915b5050505050905090565b6002805461050c9061139a565b80601f01602080910402602001604051908101604052809291908181526020018280546105389061139a565b80156105835780601f1061055a57610100808354040283529160200191610583565b820191905f5260205f20905b81548152906001019060200180831161056657829003601f168201915b505050505081565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105e3575f80fd5b80518251146105f0575f80fd5b5f5b82518110156106615781818151811061060e5761060d6113ca565b5b602002602001015160015f85848151811061062c5761062b6113ca565b5b602002602001015181526020019081526020015f20908161064d9190611608565b50808061065990611424565b9150506105f2565b505f5b8251811015610746575f151560045f858481518110610686576106856113ca565b5b602002602001015181526020019081526020015f205f9054906101000a900460ff1615150361073357600160045f8584815181106106c7576106c66113ca565b5b602002602001015181526020019081526020015f205f6101000a81548160ff0219169083151502179055506003838281518110610707576107066113ca565b5b6020026020010151908060018154018082558091505060019003905f5260205f20015f90919091909150555b808061073e90611424565b915050610664565b5060015f8081526020019081526020015f206002908161076691906116fe565b505050565b5f6107758261098d565b9050919050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107d4575f80fd5b6107de828261058b565b5050565b600481565b5f602052805f5260405f205f915054906101000a900460ff1681565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461085b575f80fd5b8060015f8481526020019081526020015f2090816108799190611608565b505050565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108d6575f80fd5b6108df836103a7565b6108e9828261058b565b505050565b606060015f8081526020019081526020015f20805461090c9061139a565b80601f01602080910402602001604051908101604052809291908181526020018280546109389061139a565b80156109835780601f1061095a57610100808354040283529160200191610983565b820191905f5260205f20905b81548152906001019060200180831161096657829003601f168201915b5050505050905090565b5f805f90505f5b60048110156109f7575f15155f808684600481106109b5576109b46113ca565b5b602002015181526020019081526020015f205f9054906101000a900460ff161515036109e457600191506109f7565b80806109ef90611424565b915050610994565b5080915050919050565b600115156001151503610a4d57610a4c6040518060400160405280601381526020017f70726974696e67206d65726b6c65207472656500000000000000000000000000815250610b74565b5b5f5b600380549050811015610b5457600115156001151503610b4157610b406040518060400160405280601781526020017f696e646578203a20256420092076616c7565203a2025730000000000000000008152508260015f8581526020019081526020015f208054610abf9061139a565b80601f0160208091040260200160405190810160405280929190818152602001828054610aeb9061139a565b8015610b365780601f10610b0d57610100808354040283529160200191610b36565b820191905f5260205f20905b815481529060010190602001808311610b1957829003601f168201915b5050505050610c0d565b5b8080610b4c90611424565b915050610a4f565b50565b6004602052805f5260405f205f915054906101000a900460ff1681565b610c0a81604051602401610b889190610df3565b6040516020818303038152906040527f41304fac000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610cac565b50565b610ca7838383604051602401610c25939291906117e3565b6040516020818303038152906040527f5970e089000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610cac565b505050565b610cc381610cbb610cc6610ce5565b63ffffffff16565b50565b5f6a636f6e736f6c652e6c6f6790505f80835160208501845afa505050565b610cf0819050919050565b610cf8611826565b565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610d1d81610d0b565b8114610d27575f80fd5b50565b5f81359050610d3881610d14565b92915050565b5f60208284031215610d5357610d52610d03565b5b5f610d6084828501610d2a565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610da0578082015181840152602081019050610d85565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610dc582610d69565b610dcf8185610d73565b9350610ddf818560208601610d83565b610de881610dab565b840191505092915050565b5f6020820190508181035f830152610e0b8184610dbb565b905092915050565b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e4d82610dab565b810181811067ffffffffffffffff82111715610e6c57610e6b610e17565b5b80604052505050565b5f610e7e610cfa565b9050610e8a8282610e44565b919050565b5f67ffffffffffffffff821115610ea957610ea8610e17565b5b602082029050919050565b5f80fd5b5f610eca610ec584610e8f565b610e75565b90508060208402830185811115610ee457610ee3610eb4565b5b835b81811015610f0d5780610ef98882610d2a565b845260208401935050602081019050610ee6565b5050509392505050565b5f82601f830112610f2b57610f2a610e13565b5b6004610f38848285610eb8565b91505092915050565b5f60808284031215610f5657610f55610d03565b5b5f610f6384828501610f17565b91505092915050565b5f67ffffffffffffffff821115610f8657610f85610e17565b5b602082029050602081019050919050565b5f610fa9610fa484610f6c565b610e75565b90508083825260208201905060208402830185811115610fcc57610fcb610eb4565b5b835b81811015610ff55780610fe18882610d2a565b845260208401935050602081019050610fce565b5050509392505050565b5f82601f83011261101357611012610e13565b5b8135611023848260208601610f97565b91505092915050565b5f67ffffffffffffffff82111561104657611045610e17565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff82111561107557611074610e17565b5b61107e82610dab565b9050602081019050919050565b828183375f83830152505050565b5f6110ab6110a68461105b565b610e75565b9050828152602081018484840111156110c7576110c6611057565b5b6110d284828561108b565b509392505050565b5f82601f8301126110ee576110ed610e13565b5b81356110fe848260208601611099565b91505092915050565b5f6111196111148461102c565b610e75565b9050808382526020820190506020840283018581111561113c5761113b610eb4565b5b835b8181101561118357803567ffffffffffffffff81111561116157611160610e13565b5b80860161116e89826110da565b8552602085019450505060208101905061113e565b5050509392505050565b5f82601f8301126111a1576111a0610e13565b5b81356111b1848260208601611107565b91505092915050565b5f80604083850312156111d0576111cf610d03565b5b5f83013567ffffffffffffffff8111156111ed576111ec610d07565b5b6111f985828601610fff565b925050602083013567ffffffffffffffff81111561121a57611219610d07565b5b6112268582860161118d565b9150509250929050565b5f8115159050919050565b61124481611230565b82525050565b5f60208201905061125d5f83018461123b565b92915050565b61126c81610d0b565b82525050565b5f6020820190506112855f830184611263565b92915050565b5f80604083850312156112a1576112a0610d03565b5b5f6112ae85828601610d2a565b925050602083013567ffffffffffffffff8111156112cf576112ce610d07565b5b6112db858286016110da565b9150509250929050565b5f805f60c084860312156112fc576112fb610d03565b5b5f61130986828701610f17565b935050608084013567ffffffffffffffff81111561132a57611329610d07565b5b61133686828701610fff565b92505060a084013567ffffffffffffffff81111561135757611356610d07565b5b6113638682870161118d565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806113b157607f821691505b6020821081036113c4576113c361136d565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61142e82610d0b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114605761145f6113f7565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026114c77fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261148c565b6114d1868361148c565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61150c61150761150284610d0b565b6114e9565b610d0b565b9050919050565b5f819050919050565b611525836114f2565b61153961153182611513565b848454611498565b825550505050565b5f90565b61154d611541565b61155881848461151c565b505050565b5b8181101561157b576115705f82611545565b60018101905061155e565b5050565b601f8211156115c0576115918161146b565b61159a8461147d565b810160208510156115a9578190505b6115bd6115b58561147d565b83018261155d565b50505b505050565b5f82821c905092915050565b5f6115e05f19846008026115c5565b1980831691505092915050565b5f6115f883836115d1565b9150826002028217905092915050565b61161182610d69565b67ffffffffffffffff81111561162a57611629610e17565b5b611634825461139a565b61163f82828561157f565b5f60209050601f831160018114611670575f841561165e578287015190505b61166885826115ed565b8655506116cf565b601f19841661167e8661146b565b5f5b828110156116a557848901518255600182019150602085019450602081019050611680565b868310156116c257848901516116be601f8916826115d1565b8355505b6001600288020188555050505b505050505050565b5f815490506116e58161139a565b9050919050565b5f819050815f5260205f209050919050565b81810361170c5750506117e1565b611715826116d7565b67ffffffffffffffff81111561172e5761172d610e17565b5b611738825461139a565b61174382828561157f565b5f601f831160018114611770575f841561175e578287015490505b61176885826115ed565b8655506117da565b601f19841661177e876116ec565b96506117898661146b565b5f5b828110156117b05784890154825560018201915060018501945060208101905061178b565b868310156117cd57848901546117c9601f8916826115d1565b8355505b6001600288020188555050505b5050505050505b565b5f6060820190508181035f8301526117fb8186610dbb565b905061180a6020830185611263565b818103604083015261181c8184610dbb565b9050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52605160045260245ffdfea26469706673582212202d86eefe4040a32352f7ea51beda42dd4b6ac15ce14a34c4b73d61e2a927973d64736f6c63430008150033",
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
