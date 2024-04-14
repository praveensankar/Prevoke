package techniques

import (
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/Revocation-Service/config"
	"go.uber.org/zap"
	"math/big"
)

func TestMerkleTree(conf config.Config){

	elements :=  make([]*big.Int, 0)
	elements = append(elements,big.NewInt(10), big.NewInt(20), big.NewInt(30), big.NewInt(40))
	//elements = append(elements,big.NewInt(100), big.NewInt(200), big.NewInt(300), big.NewInt(400))
	//elements = append(elements,big.NewInt(10), big.NewInt(20))


	newTree := CreateMerkleTree(conf)


	TestInsertion(elements, newTree)
	//TestLevelOrder(newTree)
	//TestUpdate(elements, newTree)
	TestProofs(elements, newTree)


}


func TestInsertion(elements []*big.Int, newTree *MerkleTreeAccumulator) {
	var leaf []*merkletree.Hash
	for i:=0; i< len(elements); i++{
		newTree.AddLeaf(elements[i])
		leaf = append(leaf, newTree.GetHashValueOfLeaf(elements[i]))
		//zap.S().Infoln("TEST MERKLE TREE- \t new leaf added with hash in hex: ",newTree.GetHashValueOfLeafInHex(elements[i]))
		zap.S().Infoln("TEST MERKLE TREE- \t new leaf added: \t int value: ", elements[i], "\t string: ",newTree.GetHashValueOfLeaf(elements[i]).String())

	}
	newTree.PrintTree()

}

func TestUpdate(elements []*big.Int, newTree *MerkleTreeAccumulator){
	//update tree with new elements
	replaceWith := big.NewInt(500)
	replaceWith2 := big.NewInt(600)

	oldHash := newTree.GetHashValueOfLeaf(elements[0]).String()
	affectedIndexes, affectedNodes := newTree.UpdateLeaf(elements[0], replaceWith)
	zap.S().Infoln("TEST MERKLE TREE- \t updated leaf: \t old leaf value: ", elements[0], "\t hash value: ", oldHash,
		"\t new leaf value: ", replaceWith, "\t new hash: ",newTree.GetHashValueOfLeaf(replaceWith).String())

	newTree.PrintTree()
	zap.S().Infoln("\"TEST MERKLE TREE- \t affected indexes: ", affectedIndexes, "\t affected nodes: ",affectedNodes)


	oldHash = newTree.GetHashValueOfLeaf(elements[2]).String()
	affectedIndexes, affectedNodes = newTree.UpdateLeaf(elements[2], replaceWith2)
	zap.S().Infoln("TEST MERKLE TREE- \t updated leaf: \t old leaf value: ", elements[2], "\t hash value: ", oldHash,
		"\t new leaf value: ", replaceWith, "\t new hash: ",newTree.GetHashValueOfLeaf(replaceWith2).String())

	zap.S().Infoln("\"TEST MERKLE TREE- \t affected indexes: ", affectedIndexes, "\t affected nodes: ",affectedNodes)
	newTree.PrintTree()
}


func TestProofs(elements []*big.Int, newTree *MerkleTreeAccumulator) {
	proofForele0 := newTree.GetProof(elements[0])
	var proofForele0InHex []string
	var proofForele0InString []string
	for _, hash := range proofForele0.AllSiblings() {
		proofForele0InHex = append(proofForele0InHex, hash.Hex())
		proofForele0InString = append(proofForele0InString, hash.String())
	}

	zap.S().Infoln("TEST MERKLE TREE- \t leaf value: ", elements[0], "\t string: ",
		newTree.GetHashValueOfLeaf(elements[0]).String(), "\t proof: ", proofForele0InString)

	status := newTree.VerifyProof(elements[0], proofForele0)
	if status == true {
		//zap.S().Infoln("TEST MERKLE TREE- \t leaf in hex: ",newTree.GetHashValueOfLeafInHex(elements[0])," \t used right proof - verification is successful")
		zap.S().Infoln("TEST MERKLE TREE- \t leaf in string: ", newTree.GetHashValueOfLeaf(elements[0]).String(), " \t used right proof - verification is successful")
	} else {
		zap.S().Infoln("revoked element - verification failed")
	}

	proofForele1 := newTree.GetProof(elements[1])
	var proofForele1InHex []string
	var proofForele1InString []string
	for _, hash := range proofForele1.AllSiblings() {
		proofForele1InHex = append(proofForele1InHex, hash.Hex())
		proofForele1InString = append(proofForele1InString, hash.String())
	}
	//zap.S().Infoln("TEST MERKLE TREE- \t leaf in big int: ",leaf[0].BigInt(),"\t proof: ", proofForele0InHex)
	//zap.S().Infoln("TEST MERKLE TREE- \t leaf in big int: ",leaf[1].BigInt(),"\t proof: ", proofForele1InHex)

	zap.S().Infoln("TEST MERKLE TREE- \t leaf value: ", elements[1], "\t string: ",
		newTree.GetHashValueOfLeaf(elements[1]).String(), "\t proof: ", proofForele1InString)

	status = newTree.VerifyProof(elements[1], proofForele1)
	if status == true {
		//zap.S().Infoln("TEST MERKLE TREE- \t leaf in hex: ",newTree.GetHashValueOfLeafInHex(elements[1])," \t used right proof - verification is successful")
		zap.S().Infoln("TEST MERKLE TREE- \t leaf in string: ", newTree.GetHashValueOfLeaf(elements[1]).String(),
			" \t used right proof - verification is successful")

	} else {
		//zap.S().Infoln("TEST MERKLE TREE- \t leaf in hex: ",newTree.GetHashValueOfLeafInHex(elements[1])," \t used right proof - verification failed")
		zap.S().Infoln("TEST MERKLE TREE- \t leaf in string: ", newTree.GetHashValueOfLeaf(elements[1]).String(),
			" \t used right proof - verification failed")

	}
}

	func TestLevelOrder(newTree *MerkleTreeAccumulator){

		levelOrder := newTree.GetLevelOrderRepresentation()
		zap.S().Infoln("TEST MERKLE TREE- \t level order: ", levelOrder)

	}