package issuer

import (
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math"
	"math/big"
)


type RevocationServiceStub struct{
	merkleTreeAcc *techniques.MerkleTreeAccumulator2
	VCToBigInts map[string]*big.Int
	vcCounter int64
	bloomFilter *techniques.BloomFilter
	MtLevelInDLT int
	mtHeight int
	NumberOfEntriesForMTInDLT int
}



func CreateRevocationServiceStub(config config.Config) *RevocationServiceStub{
	rs := RevocationServiceStub{}
	rs.merkleTreeAcc = techniques.CreateMerkleTreeAccumulator(config)
	rs.bloomFilter = techniques.CreateBloomFilter(config.ExpectedNumberofRevokedVCs, config.FalsePositiveRate)
	rs.VCToBigInts = make(map[string]*big.Int)
	rs.MtLevelInDLT = int(config.MtLevelInDLT)
	rs.mtHeight = int(config.MTHeight)
	rs.NumberOfEntriesForMTInDLT = 0
	for i := 0; i <= rs.MtLevelInDLT; i++ {
		rs.NumberOfEntriesForMTInDLT += int(math.Pow(2, float64(i)))
	}
	rs.vcCounter = 0
	return &rs
}



/*
Issues VC to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	_mtIndexes: merkle tree indexes
	_mtValues: merkle tree values
*/
func (r *RevocationServiceStub) IssueVC(vc verifiable.Credential) (*RevocationData) {


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
func (r *RevocationServiceStub) IssueVCsInBulk(vcs []*verifiable.Credential) ([]*RevocationData) {

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


	return revocationDataALl
}

func (r RevocationServiceStub) RetreiveUpdatedProof(vc verifiable.Credential)  *techniques.MerkleProof{
	//merkleProof := r.merkleTreeAcc.GetProofHashes(r.VCToBigInts[vc.ID])
	merkleProof := r.merkleTreeAcc.GetProof(vc.ID)
	return merkleProof
}

// returns old mt index and amount of gwei paid
func (r *RevocationServiceStub) RevokeVC(vc verifiable.Credential) (int, int64, error) {
	r.bloomFilter.RevokeInBloomFilter(vc.ID)
	//oldMTIndex := r.VCToBigInts[vc.ID]
	oldMTIndex, _ := r.merkleTreeAcc.UpdateLeaf(vc.ID, "-1")
	return oldMTIndex, -1, nil
}


func (r RevocationServiceStub) VerificationPhase1(bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int) (bool, error){

	//r.bloomFilter.CheckStatusInBloomFilter()
	var indexes []uint64
	for _, index := range bfIndexes{
		indexes = append(indexes, index.Uint64())
	}
	status := r.bloomFilter.CheckIndexesInBloomFilter(indexes)
	return status, nil
}


func (r RevocationServiceStub) VerificationPhase2( data *RevocationData)(bool, error) {

	mtRoot := r.merkleTreeAcc.RootHash
	status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)

	//zap.S().Errorln("REVOCATION SERVICE-  verification phase 2: ",status)
	return status, nil
}

func (r RevocationServiceStub) GetMerkleRoot()(string, error) {
	mtRoot := r.merkleTreeAcc.RootHash
	return mtRoot, nil
}

func (r RevocationServiceStub) VerifyVC( bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int, data *RevocationData) (bool, error) {


	//Todo: this function should be moved to the verifiers. The parameters should be shared to the holders.
	status, err := r.VerificationPhase1(bfIndexes)
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	zap.S().Infoln("REVOCATION SERVICE- ", "***VERIFY*** vc:",data.vcId)
	zap.S().Errorln("REVOCATION SERVICE-  verification phase 1: ",status)
	if status==true{
		return status, err
	} else{
		mtRoot := r.merkleTreeAcc.RootHash
		status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)
		zap.S().Errorln("REVOCATION SERVICE-  verification phase 2: ",status)
	}
	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}
	zap.S().Infoln("\n\n********************************************************************************************************************************")

	return status, err
}

func (r RevocationServiceStub) PrintMerkleTree(){
	r.merkleTreeAcc.PrintTree()
}



func (r RevocationServiceStub) LocalMTVerification(mtRoot string, data *RevocationData) {

	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification - \t vc id: ",data.vcId, "\t root: ",mtRoot,
		"\t leaf value: ", data.MerkleTreeLeafValue, "\t proof: ",data.MerkleProof.OrderedWitnesses)
	r.merkleTreeAcc.PrintTree()
	status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)
	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification : ", status)
	//statusLocal := r.merkleTreeAcc.VerifyProof(data.merkleTreeIndex, data.MerkleProof)
	//zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification local : ", statusLocal)
}




