package revocation_service

import (
	"github.com/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math/big"
)

type RevocationData struct {
	VcId    string
	MtIndex int
	BloomFilterIndexes []*big.Int
	MerkleTreeLeafValue string
	MerkleProof  *techniques.MerkleProof
}

func CreateRevocationData(vcId string, mtIndex int, bfIndexes []*big.Int, mtLeaf string, mtProof *techniques.MerkleProof)  *RevocationData {
	rd := RevocationData{}
	rd.VcId = vcId;
	rd.MtIndex=mtIndex
	rd.BloomFilterIndexes = bfIndexes;
	rd.MerkleTreeLeafValue = mtLeaf;
	rd.MerkleProof = mtProof;
	return &rd
}

func (rd *RevocationData) PrintRevocationData(){
	zap.S().Infoln("REVOCATION DATA- ","vc ID: ", rd.VcId,
		"\tmerkle tree index: ", rd.MtIndex,
	"\tbloom filter indexes: ", rd.BloomFilterIndexes,
	"\t merkle tree leaf: ", rd.MerkleTreeLeafValue,
	"\t merkle proof: ",rd.MerkleProof)
}
