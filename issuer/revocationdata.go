package issuer

import (
	"github.com/iden3/go-merkletree-sql/v2"
	"go.uber.org/zap"
	"math/big"
)

type RevocationData struct {
	vcId string
	merkleTreeIndex *big.Int
	BloomFilterIndexes []*big.Int
	MerkleTreeLeafValue *big.Int
	MerkleProof  *merkletree.Proof
}

func CreateRevocationData(vcId string, mtIndex *big.Int, bfIndexes []*big.Int, mtLeaf *big.Int, mtProof *merkletree.Proof)  *RevocationData{
	rd := RevocationData{}
	rd.vcId = vcId;
	rd.merkleTreeIndex = mtIndex;
	rd.BloomFilterIndexes = bfIndexes;
	rd.MerkleTreeLeafValue = mtLeaf;
	rd.MerkleProof = mtProof;
	return &rd
}

func (rd *RevocationData) PrintRevocationData(){
	var merkleProofInHex []string
	for _, hash := range rd.MerkleProof.AllSiblings(){
		merkleProofInHex = append(merkleProofInHex, hash.Hex())
	}
	zap.S().Infoln("REVOCATION DATA- ","vc ID: ", rd.vcId,
		"\tmerkle tree index: ", rd.merkleTreeIndex,
	"\tbloom filter indexes: ", rd.BloomFilterIndexes,
	"\t merkle tree leaf: ", rd.MerkleTreeLeafValue,
	"\t merkle proof: ",merkleProofInHex)
}
