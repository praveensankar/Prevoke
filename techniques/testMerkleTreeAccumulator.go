package techniques

import (
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
)

func TestMerkleTreeAccumulator(conf config.Config){

	conf.ExpectedNumberOfTotalVCs=4
	conf.MTHeight=2
	elements :=  make([]string, 0)
	elements = append(elements,"1","2", "3", "4")


	accumulator := CreateMerkleTreeAccumulator(conf)

	TestInsertion2(elements, accumulator)


	proofs := make(map[string]*MerkleProof)

	for i := 0; i < len(elements); i++ {
		proof := accumulator.GetProof(elements[i])
		proofs[elements[i]]=proof
	}
	TestProofs2(proofs, accumulator)

	newElements :=  make([]string, 0)
	newElements = append(newElements,"10", "20")

	replacementIndexes := []int{0,2}


	for i:=0; i<len(replacementIndexes);i++{
		TestUpdate2(elements[replacementIndexes[i]], newElements[i], accumulator)

	}

	proof := accumulator.GetProof(elements[1])
	proofs[elements[1]]=proof
	proof = accumulator.GetProof(elements[3])
	proofs[elements[3]]=proof
	proof = accumulator.GetProof(newElements[0])
	proofs[newElements[0]]=proof
	proof = accumulator.GetProof(newElements[0])
	proofs[newElements[1]]=proof
	TestProofs2(proofs, accumulator)
}


func TestInsertion2(elements []string, accumulator *MerkleTreeAccumulator2) {

	for i:=0; i< len(elements); i++{
		hash := accumulator.AddLeaf(elements[i])

		//zap.S().Infoln("TEST MERKLE TREE- \t new leaf added with hash in hex: ",newTree.GetHashValueOfLeafInHex(elements[i]))
		zap.S().Infoln("TEST MERKLE TREE- \t new leaf added: \t int value: ", elements[i], "\t hash string: ",accumulator.PrintShortString(hash))
		accumulator.PrintTree()
	}


}


func TestProofs2(proofs map[string]*MerkleProof, accumulator *MerkleTreeAccumulator2) {
	for _, proof := range proofs{
		verificationStatus := accumulator.VerifyProof(proof.LeafHash, proof.OrderedWitnesses, accumulator.RootHash)
		zap.S().Infoln("TEST MERKLE TREE- \t  leaf: ", proof.LeafValue, "\t  proof: ",
			accumulator.ProofToString(proof), "  \t verification status: ", verificationStatus)
	}
}



func TestUpdate2(oldLeaf string, newLeaf string, accumulator *MerkleTreeAccumulator2) *MerkleProof{
	//update tree with new elements

	oldHash := accumulator.GetHashValueOfLeaf(oldLeaf)
	newHash := accumulator.UpdateLeaf(oldLeaf, newLeaf)
	zap.S().Infoln("TEST MERKLE TREE- \t updated leaf: \t old leaf value: ",oldLeaf, "\t hash value: ", accumulator.PrintShortString(oldHash),
		"\t new leaf value: ", newLeaf, "\t new hash: ",accumulator.PrintShortString(newHash))

	accumulator.PrintTree()

	proof := accumulator.GetProof(newLeaf)
	return proof
}