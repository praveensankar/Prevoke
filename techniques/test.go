package techniques

import (
	"github.com/iden3/go-merkletree-sql/v2"
	"go.uber.org/zap"
	"math/big"
)

func TestMerkleTree(){
	newTree := CreateMerkleTree()
	elements :=  make([]*big.Int, 0)
	//elements = append(elements,big.NewInt(10), big.NewInt(20), big.NewInt(30), big.NewInt(40))
	elements = append(elements,big.NewInt(10), big.NewInt(20))
	var leaf []*merkletree.Hash
	for i:=0; i< len(elements); i++{
		newTree.AddLeaf(elements[i])
		leaf = append(leaf, newTree.GetHashValueOfLeaf(elements[i]))
		zap.S().Infoln("TEST MERKLE TREE- \t new leaf added with hash in hex: ",newTree.GetHashValueOfLeafInHex(elements[i]))
	}
	newTree.PrintTree()
	proofForele0 := newTree.GetProof(elements[0])
	var proofForele0InHex []string
	for _, hash := range proofForele0.AllSiblings(){
		proofForele0InHex = append(proofForele0InHex, hash.Hex())
	}
	proofForele1 := newTree.GetProof(elements[1])
	var proofForele1InHex []string
	for _, hash := range proofForele1.AllSiblings(){
		proofForele1InHex = append(proofForele1InHex, hash.Hex())
	}
	zap.S().Infoln("TEST MERKLE TREE- \t leaf in big int: ",leaf[0].BigInt(),"\t proof: ", proofForele0InHex, "\t node aux: ", proofForele0.NodeAux)
	zap.S().Infoln("TEST MERKLE TREE- \t leaf in big int: ",leaf[1].BigInt(),"\t proof: ", proofForele1InHex, "\t node aux: ", proofForele1.NodeAux)
	status := newTree.VerifyProof(elements[0],proofForele0)
	if status==true{
		zap.S().Infoln("TEST MERKLE TREE- \t leaf in hex: ",newTree.GetHashValueOfLeafInHex(elements[0])," \t used right proof - verification is successful")
	}else{
		zap.S().Infoln("used right proof - verification failed")
	}


	status = newTree.VerifyProof(elements[0],proofForele0)
	if status==true{
		zap.S().Infoln("TEST MERKLE TREE- \t leaf in hex: ",newTree.GetHashValueOfLeafInHex(elements[1])," \t used right proof - verification is successful")
	}else{
		zap.S().Infoln("TEST MERKLE TREE- \t leaf in hex: ",newTree.GetHashValueOfLeafInHex(elements[1])," \t used right proof - verification failed")
	}
	//
	//bytes := newTree.GetHashValueOfLeaf(elements[0])
	//zap.S().Infoln("hash of ", elements[0]," is: ",bytes)
	//_ = newTree.GetLevelOrderRepresentation()
}


func TestBloomFilter(estimatedVCs ...int){
	testBloomFilter(estimatedVCs...);
}
