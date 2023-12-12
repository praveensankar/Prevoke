package issuer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/iden3/go-merkletree-sql/v2"
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
	RetreiveUpdatedProof(vc verifiable.Credential) *merkletree.Proof
	VerificationPhase1(bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int) (bool, error)
	VerificationPhase2(data *RevocationData) (bool, error)
	VerifyVC( _bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int, data *RevocationData) (bool, error)
	GetMerkleRoot()(*merkletree.Hash, error)
	PrintMerkleTree()
	LocalMTVerification(mtRoot *merkletree.Hash, data *RevocationData)
}


type RevocationService struct{
	merkleTreeAcc *techniques.MerkleTreeAccumulator
	VCToBigInts map[string]*big.Int
	vcCounter int64
	bloomFilter *techniques.BloomFilter
	MtLevelInDLT int
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
	rs.merkleTreeAcc = techniques.CreateMerkleTree(config)
	rs.bloomFilter = techniques.CreateBloomFilter(config.ExpectedNumberofRevokedVCs, config.FalsePositiveRate)
	rs.smartContractAddress= common.HexToAddress(config.SmartContractAddress)
	rs.privateKey = config.PrivateKey
	rs.gasLimit = config.GasLimit
	rs.gasPrice = config.GasPrice
	rs.VCToBigInts = make(map[string]*big.Int)
	rs.MtLevelInDLT = int(config.MtLevelInDLT)
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
	mtIndex := r.merkleTreeAcc.AddLeaf(r.VCToBigInts[vc.ID])
	//root := r.merkleTreeAcc.GetRoot()
	//zap.S().Errorln("REVOCATION SERVICE- merkle root in string from local: ",root)

	levelOrderRepr := r.merkleTreeAcc.GetLevelOrderRepresentation()
	var mtIndexes []*big.Int
	var mtValues [][32]byte
	levelCounter := 0
	for i:=0; i<len(levelOrderRepr); i++{
		mtIndexes = append(mtIndexes, big.NewInt(int64(i)))
		h := levelOrderRepr[uint(i)]
		byteRepr := [32]byte{}
		copy(byteRepr[:], h[:])
		mtValues = append(mtValues, byteRepr)
		levelCounter += 1
		if levelCounter==r.NumberOfEntriesForMTInDLT{
			break
		}
	}
	//for index, value := range levelOrderRepr{
	//	mtIndexes = append(mtIndexes, big.NewInt(int64(index)))
	//	h := value
	//	byteRepr := [32]byte{}
	//	copy(byteRepr[:], h[:])
	//	mtValues = append(mtValues, byteRepr)
	//	levelCounter += 1
	//	if levelCounter==r.NumberOfEntriesForMTInDLT{
	//		break
	//	}
	//}
	//zap.S().Infoln("REVOCATION SERVICE- \t number of non-leaf nodes of MT accumulator stored in smart contract ",levelCounter)
	_, err =revocationService.IssueVC(auth, mtIndexes, mtValues)
	if err != nil {
		zap.S().Fatalln("failed to issue vc", err)
	} else{
		//zap.S().Infoln("REVOCATION SERVICE - \t issued vc: \t id: ", vc.ID)
	}
	//
	//byteRepr := [32]byte{}
	//copy(byteRepr[:], mtValues[0][:])
	//hexString := hex.EncodeToString(mtValues[0][:])
	//rootHash,_ := merkletree.NewHashFromHex(hexString)
	//zap.S().Errorln("REVOCATION SERVICE- merkle root in big int: ",rootHash.BigInt())


	//generate bloom filter indexes for the vc and give it to the holders
	bfIndexes := r.bloomFilter.GetIndexes(vc.ID)
	//mtLeaf := r.merkleTreeAcc.GetHashValueOfLeaf(r.VCToBigInts[vc.ID])§
	merkleProof := r.merkleTreeAcc.GetProof(r.VCToBigInts[vc.ID])
	//zap.S().Infoln("REVOCATION SERVICE- \t vc to big int: ",r.VCToBigInts[vc.ID])
	//zap.S().Infoln("REVOCATION SERVICE- \t hash to hex: ",mtLeaf.Hex())
	//zap.S().Infoln("REVOCATION SERVICE- \t mt leaf hash value: ", mtLeaf[:])
	//zap.S().Infoln("REVOCATION SERVICE- \t checks proof: ", r.merkleTreeAcc.VerifyProof(r.VCToBigInts[vc.ID], merkleProof))

	revocationData := CreateRevocationData(vc.ID, mtIndex,bfIndexes, r.VCToBigInts[vc.ID], merkleProof)
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
		mtIndex := r.merkleTreeAcc.AddLeaf(r.VCToBigInts[vc.ID])
		bfIndexes := r.bloomFilter.GetIndexes(vc.ID)
		merkleProof := r.merkleTreeAcc.GetProof(r.VCToBigInts[vc.ID])
		revocationData := CreateRevocationData(vc.ID, mtIndex,bfIndexes, r.VCToBigInts[vc.ID], merkleProof)
		revocationDataALl = append(revocationDataALl, revocationData)
	}
		//root := r.merkleTreeAcc.GetRoot()
		//zap.S().Errorln("REVOCATION SERVICE- merkle root in string from local: ",root)

		levelOrderRepr := r.merkleTreeAcc.GetLevelOrderRepresentation()
		var mtIndexes []*big.Int
		var mtValues [][32]byte
		levelCounter := 0
		for i := 0; i < len(levelOrderRepr); i++ {
			mtIndexes = append(mtIndexes, big.NewInt(int64(i)))
			h := levelOrderRepr[uint(i)]
			byteRepr := [32]byte{}
			copy(byteRepr[:], h[:])
			mtValues = append(mtValues, byteRepr)
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

func (r RevocationService) RetreiveUpdatedProof(vc verifiable.Credential)  *merkletree.Proof{
	//merkleProof := r.merkleTreeAcc.GetProofHashes(r.VCToBigInts[vc.ID])
	merkleProof := r.merkleTreeAcc.GetProof(r.VCToBigInts[vc.ID])
	return merkleProof
}

// returns old mt index and amount of gwei paid
func (r RevocationService) RevokeVC(vc verifiable.Credential) (*big.Int, int64, error) {
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
	revokedMTIndex := big.NewInt(-1)
	r.merkleTreeAcc.UpdateLeaf(oldMTIndex, revokedMTIndex)
	levelOrderRepr := r.merkleTreeAcc.GetLevelOrderRepresentation()
	var mtIndexes []*big.Int
	var mtValues [][32]byte
	levelCounter := 0
	for i:=0; i<len(levelOrderRepr); i++{
		mtIndexes = append(mtIndexes, big.NewInt(int64(i)))
		h := levelOrderRepr[uint(i)]
		byteRepr := [32]byte{}
		copy(byteRepr[:], h[:])
		mtValues = append(mtValues, byteRepr)
		levelCounter += 1
		if levelCounter==r.NumberOfEntriesForMTInDLT{
			break
		}
	}
	//for index, value := range levelOrderRepr{
	//	mtIndexes = append(mtIndexes, big.NewInt(int64(index)))
	//	h := value
	//	byteRepr := [32]byte{}
	//	copy(byteRepr[:], h[:])
	//	mtValues = append(mtValues, byteRepr)
	//	levelCounter += 1
	//	if levelCounter==r.NumberOfEntriesForMTInDLT{
	//		break
	//	}
	//}
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
	byteRepr := [32]byte{}
	copy(byteRepr[:], mtRoot[:])
	hexString := hex.EncodeToString(mtRoot[:])
	rootHash,_ := merkletree.NewHashFromHex(hexString)


	status := merkletree.VerifyProof(rootHash, data.MerkleProof, data.merkleTreeIndex, data.MerkleTreeLeafValue)
	//zap.S().Errorln("REVOCATION SERVICE-  verification phase 2: ",status)
	return status, nil
}

func (r RevocationService) GetMerkleRoot()(*merkletree.Hash, error) {
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
	byteRepr := [32]byte{}
	copy(byteRepr[:], mtRoot[:])
	hexString := hex.EncodeToString(mtRoot[:])
	rootHash,_ := merkletree.NewHashFromHex(hexString)
	zap.S().Errorln("REVOCATION SERVICE- merkle root in big int: ",rootHash.BigInt())
	return rootHash, nil
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
		byteRepr := [32]byte{}
		copy(byteRepr[:], mtRoot[:])
		hexString := hex.EncodeToString(mtRoot[:])
		rootHash,_ := merkletree.NewHashFromHex(hexString)


		status = merkletree.VerifyProof(rootHash, data.MerkleProof, data.merkleTreeIndex, data.MerkleTreeLeafValue)
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


func (r RevocationService) LocalMTVerification(mtRoot *merkletree.Hash, data *RevocationData) {
	var proofInHex []string
	for _, hash := range data.MerkleProof.AllSiblings(){
		proofInHex = append(proofInHex, hash.Hex())
	}
	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification - \t vc id: ",data.vcId, "\t root: ",mtRoot.BigInt(), "\t index: ",data.merkleTreeIndex,
		"\t leaf value: ", data.MerkleTreeLeafValue, "\t proof: ",proofInHex)
	r.merkleTreeAcc.PrintTree()
	status := merkletree.VerifyProof(mtRoot, data.MerkleProof, data.MerkleTreeLeafValue, data.merkleTreeIndex)
	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification : ", status)
	//statusLocal := r.merkleTreeAcc.VerifyProof(data.merkleTreeIndex, data.MerkleProof)
	//zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification local : ", statusLocal)
}


