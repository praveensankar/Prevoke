package techniques

import (
	"go.uber.org/zap"
	"math/big"
)

func TestMerkleTree(){
	newTree := createMerkleTree()
	elements :=  make([]*big.Int, 0)
	elements = append(elements,big.NewInt(10), big.NewInt(20), big.NewInt(30), big.NewInt(40))
	for i:=0; i< len(elements); i++{
		newTree.addLeaf(elements[i])
	}
	newTree.PrintTree()
	proof := newTree.getProof(elements[0])

	status := newTree.verifyProof(elements[0],proof)
	if status==true{
		zap.S().Infoln("used right proof - verification is successful")
	}else{
		zap.S().Infoln("used right proof - verification failed")
	}


	status = newTree.verifyProof(elements[1],proof)
	if status==true{
		zap.S().Infoln("used wrong proof - verification is successful")
	}else{
		zap.S().Infoln("used wrong proof - verification failed")
	}




}