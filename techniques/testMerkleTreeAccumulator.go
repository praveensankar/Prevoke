package techniques

import (
	"encoding/hex"
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
	"math"
	"strconv"
)

func TestMerkleTreeAccumulator(conf config.Config){


	remainingSpace := int(math.Pow(2, float64(conf.MTHeight)))-int(conf.ExpectedNumberOfTotalVCs)
	totalElements := int(conf.ExpectedNumberOfTotalVCs)+remainingSpace
	elements :=  make([]string, 0)
	for i:=0; i<totalElements-2; i++{
		elements = append(elements,strconv.Itoa(i))
	}

	accumulator := CreateMerkleTreeAccumulator(conf)

	TestInsertion2(elements, accumulator)


	proofs := make(map[string]*MerkleProof)

	for i := 0; i < len(elements); i++ {
		proof := accumulator.GetProof(elements[i])
		proofs[elements[i]]=proof
	}
	zap.S().Infoln("TEST MERKLE TREE- \t  proof witnesses: ", proofs[elements[0]].Witnesses)
	TestProofs2(proofs, accumulator)

	newElements :=  make([]string, 0)
	newElements = append(newElements,"10", "20")

	replacementIndexes := []int{0,2}


	for i:=0; i<len(replacementIndexes);i++{
		TestUpdate2(elements[replacementIndexes[i]], newElements[i], accumulator)

	}

	updatedProofs := make(map[string]*MerkleProof)
	proof := accumulator.GetProof(elements[1])
	updatedProofs[elements[1]]=proof
	proof = accumulator.GetProof(elements[3])
	updatedProofs[elements[3]]=proof

	updatedProofs[elements[0]]=proofs[elements[0]]
	proof = accumulator.GetProof(newElements[1])
	updatedProofs[newElements[1]]=proof
	TestProofs2(updatedProofs, accumulator)
	roothash,_ := hex.DecodeString(accumulator.RootHash)
	zap.S().Infoln("accumulator node length: ", len(accumulator.RootHash), "\t byte representation length: ",len(roothash))
}


func TestInsertion2(elements []string, accumulator *MerkleTreeAccumulator2) {

	for i:=0; i< len(elements); i++{
		_, hash := accumulator.AddLeaf(elements[i])

		//zap.S().Infoln("TEST MERKLE TREE- \t new leaf added with hash in hex: ",newTree.GetHashValueOfLeafInHex(elements[i]))
		zap.S().Infoln("TEST MERKLE TREE- \t element: ", elements[i], "\t leaf hash: ",accumulator.PrintShortString(hash))
		//accumulator.PrintTree()
	}


}


func TestProofs2(proofs map[string]*MerkleProof, accumulator *MerkleTreeAccumulator2) {
	for _, proof := range proofs{
		if proof == nil{
			zap.S().Infoln("TEST MERKLE TREE - empty proof")
			continue
		}
		verificationStatus := accumulator.VerifyProof(proof.LeafHash, proof.OrderedWitnesses, accumulator.RootHash)
		zap.S().Infoln("TEST MERKLE TREE- \t  leaf: ", proof.LeafValue, "\t  proof: ",
			accumulator.ProofToString(proof), "  \t verification status: ", verificationStatus)
	}
}



func TestUpdate2(oldLeaf string, newLeaf string, accumulator *MerkleTreeAccumulator2) *MerkleProof{
	//update tree with new elements

	oldHash := accumulator.GetHashValueOfLeaf(oldLeaf)
	_, newHash := accumulator.UpdateLeaf(oldLeaf, newLeaf)
	zap.S().Infoln("TEST MERKLE TREE- \t updated leaf: \t old leaf value: ",oldLeaf, "\t hash value: ", accumulator.PrintShortString(oldHash),
		"\t new leaf value: ", newLeaf, "\t new hash: ",accumulator.PrintShortString(newHash))

	//accumulator.PrintTree()

	proof := accumulator.GetProof(newLeaf)
	return proof
}