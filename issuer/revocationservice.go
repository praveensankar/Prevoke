package issuer

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/contracts"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math"
	"math/big"
)

type IRevocationService interface {
	IssueVC(vc verifiable.Credential) *RevocationData
	IssueVCsInBulk(vcs []*verifiable.Credential) ([]*RevocationData)
	RevokeVC(vc verifiable.Credential) (*big.Int, int64, error)
	RetreiveUpdatedProof(vc verifiable.Credential) *techniques.MerkleProof
	VerificationPhase1(bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int) (bool, error)
	VerificationPhase2(data *RevocationData) (bool, error)
	VerifyVC( _bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int, data *RevocationData) (bool, error)
	GetMerkleRoot()(string, error)
	PrintMerkleTree()
	LocalMTVerification(mtRoot string, data *RevocationData)
}


type RevocationService struct{
	merkleTreeAcc *techniques.MerkleTreeAccumulator2
	VCToBigInts map[string]*big.Int
	vcCounter int64
	bloomFilter *techniques.BloomFilter
	MtLevelInDLT int
	mtHeight int
	NumberOfEntriesForMTInDLT int
	blockchainRPCEndpoint string
	account common.Address
	smartContractAddress common.Address
	privateKey string
	gasLimit uint64
	gasPrice *big.Int
}



func CreateRevocationService(config config.Config) *RevocationService{
	rs := RevocationService{}
	rs.blockchainRPCEndpoint = config.BlockchainRpcEndpoint
	rs.merkleTreeAcc = techniques.CreateMerkleTreeAccumulator(config)
	rs.bloomFilter = techniques.CreateBloomFilter(config.ExpectedNumberofRevokedVCs, config.FalsePositiveRate)
	rs.smartContractAddress= common.HexToAddress(config.SmartContractAddress)
	rs.privateKey = config.PrivateKey
	rs.gasLimit = config.GasLimit
	rs.gasPrice = config.GasPrice
	rs.VCToBigInts = make(map[string]*big.Int)
	rs.MtLevelInDLT = int(config.MtLevelInDLT)
	rs.mtHeight = int(config.MtDepth-1)
	rs.account = common.HexToAddress(config.SenderAddress)
	rs.NumberOfEntriesForMTInDLT = 0
	for i := 0; i <= rs.MtLevelInDLT; i++ {
		rs.NumberOfEntriesForMTInDLT += int(math.Pow(2, float64(i)))
	}
	rs.vcCounter = 0
	return &rs
}

func (r RevocationService) getAuth()  *bind.TransactOpts{
	// step 1: connect to a blockchain node using RPC endpoint
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(r.privateKey)
	if err != nil {
		zap.S().Fatalln("REVOCATION SERVICE: auth error: ",err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		zap.S().Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		zap.S().Fatalln(err)
	}


	gasLimit := uint64(r.gasLimit)                // in units
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	zap.S().Fatalln(err)
	//}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = r.gasPrice

	return auth
}

func SwapEndianness(b []byte) []byte {
	o := make([]byte, len(b))
	for i := range b {
		o[len(b)-1-i] = b[i]
	}
	return o
}

/*
Issues VC to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	_mtIndexes: merkle tree indexes
	_mtValues: merkle tree values
*/
func (r *RevocationService) IssueVC(vc verifiable.Credential) (*RevocationData) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()

	r.vcCounter++
	r.VCToBigInts[vc.ID] = big.NewInt(r.vcCounter)
	mtIndex, leafHash := r.merkleTreeAcc.AddLeaf(vc.ID)

	levelOrderRepr := r.merkleTreeAcc.Tree
	levelCounter := 0
	var mtIndexes []*big.Int
	var mtValues []string
	for i:=0; i<len(levelOrderRepr); i++ {
		mtIndexes = append(mtIndexes, big.NewInt(int64(i)))
		h := levelOrderRepr[uint(i)].Value
		mtValues = append(mtValues, h)
		levelCounter += 1
		if levelCounter == r.NumberOfEntriesForMTInDLT {
			break
		}
	}
	//zap.S().Infoln("REVOCATION SERVICE- \t number of non-leaf nodes of MT accumulator stored in smart contract ",levelCounter)
	_, err =revocationService.IssueVC(auth, mtIndexes, mtValues)
	if err != nil {
		zap.S().Fatalln("failed to issue vc", err)
	} else{
		//zap.S().Infoln("REVOCATION SERVICE - \t issued vc: \t id: ", vc.ID)
	}

	//generate bloom filter indexes for the vc and give it to the holders
	bfIndexes := r.bloomFilter.GetIndexes(vc.ID)
	//mtLeaf := r.merkleTreeAcc.GetHashValueOfLeaf(r.VCToBigInts[vc.ID])§
	merkleProof := r.merkleTreeAcc.GetProof(vc.ID)
	//zap.S().Infoln("REVOCATION SERVICE- \t vc to big int: ",r.VCToBigInts[vc.ID])
	//zap.S().Infoln("REVOCATION SERVICE- \t hash to hex: ",mtLeaf.Hex())
	//zap.S().Infoln("REVOCATION SERVICE- \t mt leaf hash value: ", mtLeaf[:])
	//zap.S().Infoln("REVOCATION SERVICE- \t checks proof: ", r.merkleTreeAcc.VerifyProof(r.VCToBigInts[vc.ID], merkleProof))

	revocationData := CreateRevocationData(vc.ID, mtIndex, bfIndexes, leafHash, merkleProof)
	//revocationData.PrintRevocationData()

	return revocationData
}
/*
Issues VC to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	_mtIndexes: merkle tree indexes
	_mtValues: merkle tree values
 */
func (r *RevocationService) IssueVCsInBulk(vcs []*verifiable.Credential) ([]*RevocationData) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()
	var revocationDataALl []*RevocationData
	for _, vc := range vcs {
		r.vcCounter++
		r.VCToBigInts[vc.ID] = big.NewInt(r.vcCounter)
		mtIndex, leafHash := r.merkleTreeAcc.AddLeaf(vc.ID)
		bfIndexes := r.bloomFilter.GetIndexes(vc.ID)
		merkleProof := r.merkleTreeAcc.GetProof(vc.ID)
		revocationData := CreateRevocationData(vc.ID, mtIndex, bfIndexes, leafHash, merkleProof)
		revocationDataALl = append(revocationDataALl, revocationData)
	}

		//root := r.merkleTreeAcc.GetRoot()
		//zap.S().Errorln("REVOCATION SERVICE- merkle root in string from local: ",root)

	levelOrderRepr := r.merkleTreeAcc.Tree
	levelCounter := 0
	var mtIndexes []*big.Int
	var mtValues []string
	for i:=0; i<len(levelOrderRepr); i++ {
		mtIndexes = append(mtIndexes, big.NewInt(int64(i)))
		h := levelOrderRepr[uint(i)].Value
		mtValues = append(mtValues, h)
		levelCounter += 1
		if levelCounter == r.NumberOfEntriesForMTInDLT {
			break
		}
	}
	_, err =revocationService.IssueVC(auth, mtIndexes, mtValues)
	if err != nil {
		zap.S().Fatalln("failed to issue vc", err)
	} else{
		//zap.S().Infoln("REVOCATION SERVICE - \t issued vc: \t id: ", vc.ID)
	}

	return revocationDataALl
}

func (r RevocationService) RetreiveUpdatedProof(vc verifiable.Credential)  *techniques.MerkleProof{
	//merkleProof := r.merkleTreeAcc.GetProofHashes(r.VCToBigInts[vc.ID])
	merkleProof := r.merkleTreeAcc.GetProof(vc.ID)
	return merkleProof
}

// returns old mt index and amount of gwei paid
func (r RevocationService) RevokeVC(vc verifiable.Credential) ( *big.Int, int64, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()

	//Todo: retrieve the bloom filter indexes, merkle tree indexes and merkle tree values

	var bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int
	for i, value := range r.bloomFilter.GetIndexes(vc.ID){
		bfIndexes[i]=value;
	}
	oldMTIndex := r.VCToBigInts[vc.ID]
	index, _ := r.merkleTreeAcc.UpdateLeaf(vc.ID, "-1")
	var mtIndexes []*big.Int
	var mtValues []string
	var parentIndex int
	for i := r.merkleTreeAcc.Height -1 ; i > (r.mtHeight-r.MtLevelInDLT)+1 ; i--{
		//zap.S().Infoln("MERKLE TREE ACCUMULATOR: \t i: ", i)
		parentIndex = int(math.Floor(float64((index - 1) / 2)))
		index = parentIndex
	}
	for i := r.MtLevelInDLT ; i > 0 ; i--{
		//zap.S().Infoln("MERKLE TREE ACCUMULATOR: \t i: ", i)
		parentIndex = int(math.Floor(float64((index - 1) / 2)))
		hashValue := r.merkleTreeAcc.Tree[parentIndex]
		mtIndexes = append(mtIndexes, big.NewInt(int64(parentIndex)))
		mtValues = append(mtValues, hashValue.Value)
		index = parentIndex
	}


	zap.S().Infoln("REVOCATION SERVICE- \t mt indexes: ", mtIndexes, "\t mt values: ",mtValues)
	//zap.S().Infoln("REVOCATION SERVICE- \t number of non-leaf nodes of MT accumulator stored in smart contract ",levelCounter)
	startBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	_, err = revocationService.RevokeVC(auth, bfIndexes, mtIndexes, mtValues)
	endBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	gasUsed := (startBalance.Int64()-endBalance.Int64())/int64(math.Pow(10,9))
	//zap.S().Infoln("REVOCATION SERVICE- \t MT Accumulator levels in DLT: ",r.NumberOfEntriesForMTInDLT, "GAS USAGE in gwei: ", gasUsed)



	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}

	return oldMTIndex, gasUsed, nil
}


func (r RevocationService) VerificationPhase1(bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int) (bool, error){
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	vcStatus, err := revocationService.VerificationPhase1(nil, bfIndexes)
	//zap.S().Errorln("REVOCATION SERVICE-  vc.IDverification phase 1: ",vcStatus)

	return vcStatus, err
}


func (r RevocationService) VerificationPhase2( data *RevocationData)(bool, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)


	mtRoot, err := revocationService.VerificationPhase2(nil)
	if err!=nil{
		zap.S().Errorln("REVOCATION SERVICE- error verification phase 2: ",err)
	}



	status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)

	//zap.S().Errorln("REVOCATION SERVICE-  verification phase 2: ",status)
	return status, nil
}

func (r RevocationService) GetMerkleRoot()(string, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	mtRoot, err := revocationService.VerificationPhase2(nil)
	if err!=nil{
		zap.S().Errorln("REVOCATION SERVICE- error verification phase 2: ",err)
	}

	zap.S().Errorln("REVOCATION SERVICE- merkle root: ",mtRoot)
	return mtRoot, nil
}

func (r RevocationService) VerifyVC( _bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int, data *RevocationData) (bool, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}


	//Todo: this function should be moved to the verifiers. The parameters should be shared to the holders.
	status, err := revocationService.VerificationPhase1(nil, _bfIndexes)
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	zap.S().Infoln("REVOCATION SERVICE- ", "***VERIFY*** vc:",data.vcId)
	zap.S().Errorln("REVOCATION SERVICE-  verification phase 1: ",status)
	if status==true{
		return status, err
	} else{
		mtRoot, err := revocationService.VerificationPhase2(nil)
		if err!=nil{
			zap.S().Errorln("REVOCATION SERVICE- error verification phase 2: ",err)
		}

		status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)
		zap.S().Errorln("REVOCATION SERVICE-  verification phase 2: ",status)
	}
	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}
	zap.S().Infoln("\n\n********************************************************************************************************************************")

	return status, err
}

func (r RevocationService) PrintMerkleTree(){
	r.merkleTreeAcc.PrintTree()
}


func (r RevocationService) LocalMTVerification(mtRoot string, data *RevocationData) {

	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification - \t vc id: ",data.vcId, "\t root: ",mtRoot,
		"\t leaf value: ", data.MerkleTreeLeafValue, "\t proof: ",data.MerkleProof.OrderedWitnesses)
	r.merkleTreeAcc.PrintTree()
	status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)
	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification : ", status)
	//statusLocal := r.merkleTreeAcc.VerifyProof(data.merkleTreeIndex, data.MerkleProof)
	//zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification local : ", statusLocal)
}


