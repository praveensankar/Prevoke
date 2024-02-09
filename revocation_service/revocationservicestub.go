package revocation_service

import (
	_ "github.com/ethereum/go-ethereum/core/types"
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
	PublicKeys [][]byte
}



func CreateRevocationServiceStub(config config.Config) *RevocationServiceStub {
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
	rs.PublicKeys = make([][]byte, 0)
	return &rs
}



/*
Issues VC to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	_mtIndexes: merkle tree indexes
	_mtValues: merkle tree values
*/
func (r *RevocationServiceStub) IssueVC(vcID string) (*RevocationData) {


	mtIndex, leafHash := r.merkleTreeAcc.AddLeaf(vcID)

	//mtIndexes, mtValues := r.merkleTreeAcc.GetEntriesInLevelOrder(r.NumberOfEntriesForMTInDLT)
	//zap.S().Infoln("REVOCATION SERVICE - merkle tree indexes and values: ", mtIndexes, mtValues)
	//zap.S().Infoln("REVOCATION SERVICE- \t number of non-leaf nodes of MT accumulator stored in smart contract ",levelCounter)

	//generate bloom filter indexes for the vc and give it to the holders
	bfIndexes := r.bloomFilter.GetIndexes(vcID)
	//mtLeaf := r.merkleTreeAcc.GetHashValueOfLeaf(r.VCToBigInts[vc.ID])§
	merkleProof := r.merkleTreeAcc.GetProof(vcID)
	//zap.S().Infoln("REVOCATION SERVICE- \t vc to big int: ",r.VCToBigInts[vc.ID])
	//zap.S().Infoln("REVOCATION SERVICE- \t hash to hex: ",mtLeaf.Hex())
	//zap.S().Infoln("REVOCATION SERVICE- \t mt leaf hash value: ", mtLeaf[:])
	//zap.S().Infoln("REVOCATION SERVICE- \t checks proof: ", r.merkleTreeAcc.VerifyProof(r.VCToBigInts[vc.ID], merkleProof))

	revocationData := CreateRevocationData(vcID, mtIndex, bfIndexes, leafHash, merkleProof)
	//revocationData.PrintRevocationData()

	return revocationData
}
/*
Issues VC to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	_mtIndexes: merkle tree indexes
	_mtValues: merkle tree values
*/
func (r *RevocationServiceStub) IssueVCsInBulk(vcIDs []string) ([]*RevocationData) {

	var revocationDataALl []*RevocationData
	for _, vcID := range vcIDs {
		//zap.S().Infoln("REVOCATION SERVICE STUB - VC ID: ",vcID)
		mtIndex, leafHash := r.merkleTreeAcc.AddLeaf(vcID)
		bfIndexes := r.bloomFilter.GetIndexes(vcID)
		merkleProof := r.merkleTreeAcc.GetProof(vcID)
		revocationData := CreateRevocationData(vcID, mtIndex, bfIndexes, leafHash, merkleProof)
		revocationDataALl = append(revocationDataALl, revocationData)
	}
	//mtIndexes, mtValues := r.merkleTreeAcc.GetEntriesInLevelOrder(r.NumberOfEntriesForMTInDLT)
	//zap.S().Infoln("REVOCATION SERVICE - merkle tree indexes and values: ", mtIndexes, mtValues)
	return revocationDataALl
}

func (r RevocationServiceStub) RetreiveUpdatedProof(vcID string)  *techniques.MerkleProof{
	//merkleProof := r.merkleTreeAcc.GetProofHashes(r.VCToBigInts[vc.ID])
	merkleProof := r.merkleTreeAcc.GetProof(vcID)
	return merkleProof
}

// returns old mt index and amount of gwei paid
func (r *RevocationServiceStub) RevokeVC(vcID string) (int, int64, error) {
	r.bloomFilter.RevokeInBloomFilter(vcID)
	//oldMTIndex := r.VCToBigInts[vc.ID]
	oldMTIndex, _ := r.merkleTreeAcc.UpdateLeaf(vcID, "-1")
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

func (r RevocationServiceStub) FetchMerkleTree() ([]string){


	var mtValues []string

	_, mtValues = r.merkleTreeAcc.GetEntriesInLevelOrder(r.NumberOfEntriesForMTInDLT)

	//zap.S().Errorln("REVOCATION SERVICE- merkle values: ",GetShortString(mtValues))
	return mtValues;
}

func (r RevocationServiceStub) VerifyVC( bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int, data *RevocationData) (bool, error) {


	//Todo: this function should be moved to the verifiers. The parameters should be shared to the holders.
	status, err := r.VerificationPhase1(bfIndexes)
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	zap.S().Infoln("REVOCATION SERVICE- ", "***VERIFY*** vc:",data.VcId)
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

	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification - \t vc id: ",data.VcId, "\t root: ",mtRoot,
		"\t leaf value: ", data.MerkleTreeLeafValue, "\t proof: ",data.MerkleProof.OrderedWitnesses)
	r.merkleTreeAcc.PrintTree()
	status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)
	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification : ", status)
	//statusLocal := r.merkleTreeAcc.VerifyProof(data.merkleTreeIndex, data.MerkleProof)
	//zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification local : ", statusLocal)
}

func (r* RevocationServiceStub) AddPublicKeys(publicKeys [][]byte) {
	r.PublicKeys = append(r.PublicKeys, publicKeys...)
}

/*
FetchPublicKeys retrieves the issuer's public keys from the smart contract

Output:
	public Keys - []string
*/
func (r RevocationServiceStub) FetchPublicKeys()([][]byte) {
	//zap.S().Infoln("ReVOCATION SERVICE STUB - public keys: ",r.PublicKeys)
	return r.PublicKeys
}




