package techniques

import (
	"context"
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/iden3/go-merkletree-sql/v2/db/memory"
	"go.uber.org/zap"
	"math/big"
	"os"
)

type MerkleTreeAccumulator struct {
	RootHash []byte
	currentIndex int
	Tree *merkletree.MerkleTree
	leafsToIndexes map[string]big.Int
}

func  CreateMerkleTree() *MerkleTreeAccumulator {
	ctx := context.Background()
	var treeStorage merkletree.Storage
	treeStorage = memory.NewMemoryStorage()
	mtDepth := 40
	tree, err := merkletree.NewMerkleTree(ctx, treeStorage, mtDepth)
	if err!=nil{
		zap.S().Errorln("error creating merkle tree: ", err)
	}
	accumulator := MerkleTreeAccumulator{Tree: tree, currentIndex: 0}
	accumulator.leafsToIndexes = make(map[string]big.Int)
	return &accumulator
}

func (accumulator *MerkleTreeAccumulator) getRoot() string{
	return accumulator.Tree.Root().String()
}

func (accumulator *MerkleTreeAccumulator) addLeaf(leaf *big.Int){
	ctx := context.Background()
	accumulator.currentIndex++
	index := big.NewInt(int64(accumulator.currentIndex))
	accumulator.leafsToIndexes[leaf.String()]=*index
	err := accumulator.Tree.Add(ctx, index, leaf)
	if err != nil {
		zap.S().Errorf("error adding %s to merkle tree: ", err)
	}
}

func (accumulator *MerkleTreeAccumulator) getProof(leaf *big.Int) *merkletree.Proof{
	ctx := context.Background()
	index := accumulator.leafsToIndexes[leaf.String()]
	proof, _, err := accumulator.Tree.GenerateProof(ctx, &index, nil)
	if err != nil {
		zap.S().Errorf("error generating proof for %s in merkle tree: ", err)
	}
	zap.S().Infoln(
		"proof",proof.Bytes())
	return proof
}

// returns true if proof is valid. otherwise returns false
func (accumulator *MerkleTreeAccumulator) verifyProof(leaf *big.Int, proof *merkletree.Proof) bool {
	index := accumulator.leafsToIndexes[leaf.String()]
	valid := merkletree.VerifyProof(accumulator.Tree.Root(), proof, &index, leaf)
	return valid
}

func (accumulator *MerkleTreeAccumulator) PrintTree(){
	zap.S().Info("**************************************** printing merkle tree ***********************************************************")
	ctx := context.Background()
	err := accumulator.Tree.PrintGraphViz(ctx,accumulator.Tree.Root())
	if err!=nil{
		zap.S().Errorln("error while visualizing the merkle tree: ",err)
	}
	file, _ := os.Create("merkleTree.gv")
	accumulator.Tree.GraphViz(ctx, file, accumulator.Tree.Root())
	leaves, _ := accumulator.Tree.DumpLeafs(ctx, nil)
	zap.S().Infoln("leaves: ",leaves)
	zap.L().Info("********************************************************************************************************************************\n")
}
