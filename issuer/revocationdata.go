package issuer

import (
	"github.com/iden3/go-merkletree-sql/v2"
	"go.uber.org/zap"
)

type RevocationData struct {
	BloomFilterIndexes []uint64
	MerkleTreeLeafValue *merkletree.Hash
	MerkleProof  []*merkletree.Hash
}

func CreateRevocationData(bfIndexes []uint64, mtLeaf *merkletree.Hash, mtProof []*merkletree.Hash)  *RevocationData{
	rd := RevocationData{}
	rd.BloomFilterIndexes = bfIndexes;
	rd.MerkleTreeLeafValue = mtLeaf;
	rd.MerkleProof = mtProof;
	return &rd
}

func (rd RevocationData) PrintRevocationData(){
	zap.S().Infoln("bloom filter indexes: ", rd.BloomFilterIndexes)
	zap.S().Infoln("merkle tree leaf: ", rd.MerkleTreeLeafValue)
	zap.S().Infoln("merkle proof: ",rd.MerkleProof)
}
