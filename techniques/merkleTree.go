package techniques

import (
	"context"
	"errors"
	"github.com/iden3/go-iden3-crypto/poseidon"
	cryptoUtils "github.com/iden3/go-iden3-crypto/utils"
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/iden3/go-merkletree-sql/v2/db/memory"
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
	"math/big"
)

type MerkleTreeAccumulator struct {
	RootHash []byte
	currentIndex int
	Tree *merkletree.MerkleTree
	leafsToIndexes map[string]int64
	levelOrder map[uint]*merkletree.Hash
	DEBUG bool
}

// the library uses sparse merkle tree. In sparse merkle tree, each leaf is indexed.
// sparse merkle tree also allows generation of proof of non-membership
func  CreateMerkleTree(conf config.Config) *MerkleTreeAccumulator {
	ctx := context.Background()
	var treeStorage merkletree.Storage
	treeStorage = memory.NewMemoryStorage()
	mtDepth := conf.MtDepth
	tree, err := merkletree.NewMerkleTree(ctx, treeStorage, int(mtDepth))

	if err!=nil{
		zap.S().Errorln("error creating merkle tree: ", err)
	}
	accumulator := MerkleTreeAccumulator{Tree: tree, currentIndex: 0, DEBUG: true}
	accumulator.leafsToIndexes = make(map[string]int64)
	accumulator.levelOrder = make(map[uint]*merkletree.Hash)
	return &accumulator
}

func (accumulator *MerkleTreeAccumulator) GetRoot() string{
	return accumulator.Tree.Root().String()
}

// returns index used in merkle tree
func (accumulator *MerkleTreeAccumulator) AddLeaf(leaf *big.Int) *big.Int{
	ctx := context.Background()
	accumulator.currentIndex++
	index := big.NewInt(int64(accumulator.currentIndex))
	accumulator.leafsToIndexes[leaf.String()] = int64(accumulator.currentIndex)
	err := accumulator.Tree.Add(ctx, index, leaf)
	//zap.S().Infoln("MERKLE TREE- \t new leaf added: \t index: ", index, "\t value: ",leaf)
	if err != nil {
		zap.S().Errorf("error adding %s to merkle tree: ", err)
	}
	return index
}

func (accumulator *MerkleTreeAccumulator) UpdateLeaf(oldLeaf *big.Int, newLeaf *big.Int){
	ctx := context.Background()
	index := big.NewInt(accumulator.leafsToIndexes[oldLeaf.String()])
	_, err := accumulator.Tree.Update(ctx, index, newLeaf)
	accumulator.leafsToIndexes[oldLeaf.String()] = -1
	accumulator.leafsToIndexes[newLeaf.String()] = index.Int64()
	if err != nil {
		zap.S().Errorf("error updating leaf to merkle tree: ", err)
	}
}

func (accumulator *MerkleTreeAccumulator) GetProof(leaf *big.Int) *merkletree.Proof{
	ctx := context.Background()
	index := big.NewInt(accumulator.leafsToIndexes[leaf.String()])
	proof, _, err := accumulator.Tree.GenerateProof(ctx, index, nil)
	if err != nil {
		zap.S().Errorf("error generating proof for %s in merkle tree: ", err)
	}
	//zap.S().Infoln("proof",proof.AllSiblings(), "\tlength of each hash in the proof: ", len(proof.AllSiblings()[0]),"bytes")
	return proof
}

func (accumulator *MerkleTreeAccumulator) GetProofHashes(leaf *big.Int)  []*merkletree.Hash{
	ctx := context.Background()
	index := big.NewInt(accumulator.leafsToIndexes[leaf.String()])
	proof, _, err := accumulator.Tree.GenerateProof(ctx, index, nil)
	if err != nil {
		zap.S().Errorf("error generating proof for %s in merkle tree: ", err)
	}
	return proof.AllSiblings()
}


func (accumulator *MerkleTreeAccumulator) GetHashValueOfLeaf(leaf *big.Int) *merkletree.Hash {
	leafHash, _ := merkletree.NewHashFromBigInt(leaf)
	index := big.NewInt(accumulator.leafsToIndexes[leaf.String()])
	keyHash,_ := merkletree.NewHashFromBigInt(index)
	hash, err := merkletree.LeafKey(keyHash, leafHash)
	//zap.S().Infoln("MERKLE TREE- \t leaf: \t index: ", index, "\t hash: ", keyHash, "\t value: ",leaf, "\t hash: ",leafHash, "\t leaf key: ", hash.Hex())
	if err!=nil {
		zap.S().Infoln("MERKLE TREE- \thash value of the leaf node: ", hash, "\tlength: ",len(hash))
	}
	return hash
}

func (accumulator *MerkleTreeAccumulator) GetHashValueOfLeafInHex(leaf *big.Int) string {
	hash := accumulator.GetHashValueOfLeaf(leaf)
	return hash.Hex()
	//if err==nil {
	//	zap.S().Infoln("hash value of the leaf node: ", hash, "\tlength: ",len(hash))
	//}

}

func (accumulator *MerkleTreeAccumulator) GetLevelOrderRepresentation() map[uint]*merkletree.Hash{
	ctx := context.Background()
	var counter uint
	counter = 0
	accumulator.levelOrder[counter] = accumulator.Tree.Root()
	counter++
	//levelOrderRepr = append(levelOrderRepr, accumulator.Tree.Root())
	_ = accumulator.Tree.Walk(ctx, nil, func(n *merkletree.Node) {
		switch n.Type {
		case merkletree.NodeTypeMiddle:
			accumulator.levelOrder[counter] = n.ChildL
			counter++
			accumulator.levelOrder[counter] = n.ChildR
			counter++
			//levelOrderRepr = append(levelOrderRepr, n.ChildL, n.ChildR)
		default:
		}
	})
	//zap.S().Infoln("level order representation: ", accumulator.levelOrder)
	return accumulator.levelOrder
}

// returns true if proof is valid. otherwise returns false
func (accumulator *MerkleTreeAccumulator) VerifyProof(leaf *big.Int, proof *merkletree.Proof) bool {
	index := big.NewInt(accumulator.leafsToIndexes[leaf.String()])
	//zap.S().Infoln("MERKLE TREE- Verification: \tindex: ", index, "\t leaf: ",leaf)
	valid := merkletree.VerifyProof(accumulator.Tree.Root(), proof, index, leaf)
	return valid
}


func (accumulator *MerkleTreeAccumulator) LocalMTVerification(leaf *big.Int, proof *merkletree.Proof) bool{
	var merkleProofInHex []*merkletree.Hash
	var root *merkletree.Node
	var err error

	//index := big.NewInt(accumulator.leafsToIndexes[leaf.String()])
	//keyHash,_ := merkletree.NewHashFromBigInt(index)
	//path := getPath(int(accumulator.Tree.MaxLevels()), keyHash[:])
	//siblings := proof.AllSiblings()
	//var siblingKey *merkletree.Hash
	//var midKey *merkletree.Hash
	//sibIdx := len(siblings) - 1
	//for lvl := int(accumulator.Tree.MaxLevels()) - 1; lvl >= 0; lvl-- {
	//	siblingKey = siblings[sibIdx]
	//	if path[lvl] {
	//		midKey, err = NewNodeMiddle(siblingKey, midKey).Key()
	//		if err != nil {
	//			return false
	//		}
	//	} else {
	//		midKey, err = NewNodeMiddle(midKey, siblingKey).Key()
	//		if err != nil {
	//			return false
	//		}
	//	}
	//}
	//zap.S().Infoln("TEST MERKLE TREE- \t root value in mid key: ",midKey)

	//zap.S().Infoln("TEST MERKLE TREE- \t root value in mid key: ",midKey)


	for _, hash := range proof.AllSiblings(){
		merkleProofInHex = append(merkleProofInHex, hash)
		zap.S().Infoln("TEST MERKLE TREE- \t leaf in big int: ",leaf,"\t proof: ", hash.BigInt())
		leafHash := accumulator.GetHashValueOfLeaf(leaf)
		root = NewNodeMiddle(hash, leafHash)
		if err!=nil{
			zap.S().Errorln("\"MERKLE TREE- \t unable to calculate hash: ", err)
		}
	}
	leftChild := root.ChildL.BigInt()
	rightChild := root.ChildR.BigInt()
	rootHash, _ := root.Key()
	if root!=nil {
		zap.S().Infoln("MERKLE TREE- ", "\t local MT verification : \t root", rootHash.BigInt(), "\t left child: ",leftChild, "\t right child: ",rightChild)
	}

	mtRoot := accumulator.Tree.Root().BigInt()
	zap.S().Infoln("MERKLE TREE- ", "\t root", mtRoot)

	if mtRoot.Cmp(rootHash.BigInt())==0{
		return true
	} else{
		return false
	}
}

func (accumulator *MerkleTreeAccumulator) PrintTree(){
	//zap.S().Info("**************************************** printing merkle tree ***********************************************************")
	ctx := context.Background()
	tree := make(map[uint]string)
	treeInBigInt := make(map[uint]*big.Int)
	treeLeaves := make(map[uint]string)
	var counter uint
	counter = 0
	tree[counter] = accumulator.Tree.Root().Hex()
	treeInBigInt[counter] = accumulator.Tree.Root().BigInt()
	treeLeaves[counter] = accumulator.Tree.Root().String()
	counter++
	//levelOrderRepr = append(levelOrderRepr, accumulator.Tree.Root())
	_ = accumulator.Tree.Walk(ctx, nil, func(n *merkletree.Node) {
		switch n.Type {
		case merkletree.NodeTypeMiddle:
			tree[counter] = n.ChildL.Hex()
			treeInBigInt[counter] = n.ChildL.BigInt()
			treeLeaves[counter] = n.ChildL.String()
			counter++
			tree[counter] = n.ChildR.Hex()
			treeInBigInt[counter] = n.ChildR.BigInt()
			treeLeaves[counter] = n.ChildR.String()
			counter++
			//levelOrderRepr = append(levelOrderRepr, n.ChildL, n.ChildR)
		default:
		}
	})
	if accumulator.DEBUG==true {
		//zap.S().Infoln("MERKLE TREE- hex values: ", tree)
		//zap.S().Infoln("MERKLE TREE- big int: ", treeInBigInt)
		zap.S().Infoln("MERKLE TREE- string: ", treeLeaves)
	}

	//err := accumulator.Tree.PrintGraphViz(ctx,accumulator.Tree.Root())
	//if err!=nil{
	//	zap.S().Errorln("error while visualizing the merkle tree: ",err)
	//}
	//file, _ := os.Create("merkleTree.gv")
	//accumulator.Tree.GraphViz(ctx, file, accumulator.Tree.Root())
	//leaves, _ := accumulator.Tree.DumpLeafs(ctx, nil)
	//zap.S().Infoln("leaves: ",leaves)
	//zap.L().Info("********************************************************************************************************************************\n")
}

// HashElems performs a poseidon hash over the array of ElemBytes, currently we
// are using 2 elements.  Uses poseidon.Hash to be compatible with the circom
// circuits implementations.
func HashElems(elems ...*big.Int) (*big.Int, error) {
	poseidonHash, err := poseidon.Hash(elems)
	if err != nil {
		return nil, err
	}
	return poseidonHash, err
}


// NewHashFromBigInt returns a *Hash representation of the given *big.Int
func NewHashFromBigInt(b *big.Int) (*merkletree.Hash, error) {
	if !cryptoUtils.CheckBigIntInField(b) {
		return nil, errors.New(
			"NewHashFromBigInt: Value not inside the Finite Field")
	}
	r := &merkletree.Hash{}
	copy(r[:], SwapEndianness(b.Bytes()))
	return r, nil
}

// SwapEndianness swaps the order of the bytes in the slice.
func SwapEndianness(b []byte) []byte {
	o := make([]byte, len(b))
	for i := range b {
		o[len(b)-1-i] = b[i]
	}
	return o
}


// LeafKey computes the key of a leaf node given the hIndex and hValue of the
// entry of the leaf.
func LeafKey(k, v *merkletree.Hash) (*big.Int, error) {
	return HashElems(big.NewInt(1), k.BigInt(), v.BigInt())
}

// NewNodeMiddle creates a new middle node.
func NewNodeMiddle(childL *merkletree.Hash, childR *merkletree.Hash) *merkletree.Node {
	return &merkletree.Node{Type: 0, ChildL: childL, ChildR: childR}
}

// getPath returns the binary path, from the root to the leaf.
func getPath(numLevels int, k []byte) []bool {
	path := make([]bool, numLevels)
	for n := 0; n < numLevels; n++ {
		path[n] = TestBit(k[:], uint(n))
	}
	return path
}

// TestBit tests whether the bit n in bitmap is 1.
func TestBit(bitmap []byte, n uint) bool {
	return bitmap[n/8]&(1<<(n%8)) != 0
}
