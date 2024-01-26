package techniques

import (
	"encoding/hex"
	"fmt"
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	"math"
)

const DIGEST_SIZE = 32
const SHORT_STRING_SIZE = 4

type Position int64
const (
	LEFT Position = iota
	RIGHT
)

func (p Position) String() string {
	switch p {
	case LEFT:
		return "left"
	case RIGHT:
		return "right"
	}
	return "unknown"
}

type IMerkleTreeAccumulator interface {
	GetHash(input string) string
	AddLeaf(value string) string
	GetHashValueOfLeaf(leaf string) string
	GetProof(leaf string) *MerkleProof
	IsEmpty(index int) bool
  	CreateMerkleTreeAccumulator(conf config.Config) *MerkleTreeAccumulator2
}

type Node struct{
	Value string
}

type Witness struct{
	HashValue string
	Position Position
}

type MerkleProof struct {
	LeafHash string
	LeafValue string
	Witnesses map[int]*Witness
	OrderedWitnesses []*Witness
	Order []int
}

// leafsToIndexes - stores (leaf, index) pairs. The leaf value is stored as it is, not the hash digest of leaf.
type MerkleTreeAccumulator2 struct {
	RootHash string
	CurrentIndex int
	Height int
	TotalLeafs int
	TotalNodes int
	Tree []*Node
	leafsToIndexes map[string]int
	DEBUG bool
}

/*
This function creates a new merkle tree accumulator.

If total number of vcs are not in power of 2, then the leaf counts will be adjusted to the nearest power of 2.

Inputs:
	TotalNumberofVCs - number of VCs issuer expects to issue in its lifetime
	height - Height of Merkle Tree

Output:
	MerkleTreeAccumulator object

*/
func  CreateMerkleTreeAccumulator(conf config.Config) *MerkleTreeAccumulator2 {

	totalLeafs := conf.ExpectedNumberOfTotalVCs
	height := conf.MTHeight

	// round the number of leaves to power of 2
	remainingSpace := int(math.Pow(2, float64(height)))-int(totalLeafs)
	totalLeafs = totalLeafs + uint(remainingSpace)
	totalNodes := int(math.Pow(2, float64(height+1))-1)

	accumulator := MerkleTreeAccumulator2{}
	accumulator.TotalLeafs = int(totalLeafs)
	accumulator.TotalNodes = totalNodes

	zap.S().Errorln("\"MERKLE TREE Accumulator- \t total nodes: ", accumulator.TotalNodes)
	accumulator.Tree=make([]*Node, totalNodes)
	for i:=0; i<totalNodes; i++{
		accumulator.Tree[i] = &Node{}
		accumulator.Tree[i].Value = ""
	}

	accumulator.leafsToIndexes = make(map[string]int)
	accumulator.CurrentIndex = int(math.Pow(2, float64(height)))-1

	accumulator.Height = int(height)
	return &accumulator
}

func (accumulator *MerkleTreeAccumulator2) ByteArrayToString(input []byte) string {
	digestInSting := hex.EncodeToString(input)
	return digestInSting
}

func (accumulator *MerkleTreeAccumulator2) StringToByteArray(input string) [DIGEST_SIZE]byte {
	digestInByteArray, _ := hex.DecodeString(input)
	res := [DIGEST_SIZE]byte{}
	copy(res[:], digestInByteArray)
	return res
}


func (accumulator *MerkleTreeAccumulator2) GetHash(input string) string {
	h :=  sha3.New256()
	h.Write([]byte(input))
	digest := h.Sum(nil)
	return accumulator.ByteArrayToString(digest)

}

/*
AddLeaf adds a new leaf to the merkle tree accumulator.

The Merkle Tree is represented in Level Order. Therefore, root is stored at index 0, middle nodes from index 1.
The leafs are stored from the index (2^height)-1.

E.g. If Merkle Tree can store 4 leafs, then the tree height is 2. Root is at height 0, middle nodes
are at height 1, and 4 leafs are at height 2. The leafs are stored at indexes from 3 to 6. The root node is stored
at index 0.

When the merkle tree is empty, the index to insert the element points to the left most position. e.g. for a 4 leaf
merkle tree, the leaf index starts 3.

Algorithm:
1) fetch the current index for leaf
2) calculate the hash value of leaf
3) store the hash at the current index in the merkle tree
4) store the (leaf, index) pair in leafsToIndexes
5) update the hash values of middle nodes and root nodes
6) increment the current leaf index


Input:
	1) (string) - leaf value

Output:
	1) (int) - index of leaf in the merkle tree accumulator
	2) (string) - hash value of leaf that is stored in the merkle tree accumulator

*/
func (accumulator *MerkleTreeAccumulator2)  AddLeaf(leaf string) (int, string) {

	if accumulator.CurrentIndex > accumulator.TotalNodes{
		zap.S().Errorln("merkle tree accumulator is full")
		return -1,""
	}

	// 1) fetch the current index for leaf
	index := accumulator.CurrentIndex

	// 2) calculate the hash value of leaf
	hashValue := accumulator.GetHash(leaf)
	//zap.S().Infoln("MERKLE TREE ACCUMULATOR: \t current index : ", index)

	// 3) store the hash at the current index in the merkle tree
	accumulator.Tree[index].Value = hashValue

	// 4) store the (leaf, index) pair in leafsToIndexes
	accumulator.leafsToIndexes[leaf] = index

	// 5) update the hash values of middle nodes and root nodes
	accumulator.UpdateMiddleandRootNodes(index)
	accumulator.RootHash=accumulator.Tree[0].Value

	// 6) increment the current leaf index
	accumulator.CurrentIndex++
	
	return index, hashValue

}


/*
UpdateMiddleandRootNodes updates the middle and root nodes in the path from a leaf node to the root.

The leaf node is marked by its index.

This function is called whenever a new leaf is added or existing leaf is modified. The following functions call this
function:
1) AddLeaf()
2) UpdateLeaf()

Algorithm:
1) Loop through the parent of leaf node till root (not including root): For each iteration,
2) calculate the parent's index
3) calculate left child's index of the parent
4) calculate right child's index of the parent
5) update the parent's value with the hash of left and right value
6) Once the loop is over, update the root value

Input:
	1) (int) - index of leaf in the merkle tree accumulator

Output:
	-
*/
func (accumulator *MerkleTreeAccumulator2)  UpdateMiddleandRootNodes(index int) {
	var parentIndex int
	//zap.S().Infoln("MERKLE TREE ACCUMULATOR: \t index: ", index)

	//1) Loop through the parent of leaf node till root (not including root):
	for i := accumulator.Height -1 ; i > 0 ; i--{

		//zap.S().Infoln("MERKLE TREE ACCUMULATOR: \t i: ", i)

		//2) calculate the parent's index
		parentIndex = int(math.Floor(float64((index - 1) / 2)))
		//leftChildIndex := int(math.Pow(2, float64(parentIndex)))+1
		//rightChildIndex := int(math.Pow(2, float64(parentIndex)))+2

		//3) calculate left child's index of the parent
		leftChildIndex := (2*parentIndex)+1

		//4) calculate right child's index of the parent
		rightChildIndex := (2*parentIndex)+2

		//5) update the parent's value with the hash of left and right value
		lefChildValue := accumulator.Tree[leftChildIndex].Value
		rightChildValue := accumulator.Tree[rightChildIndex].Value
		hashValue := accumulator.GetHash(lefChildValue + rightChildValue)
		accumulator.Tree[parentIndex].Value = hashValue

		index = parentIndex
	}
	//6) Once the loop is over, update the root value
	accumulator.Tree[0].Value = accumulator.GetHash(accumulator.Tree[1].Value+accumulator.Tree[2].Value)
}



/*
UpdateLeaf updates an existing leaf with new leaf in the merkle tree accumulator.

The Merkle Tree is represented in Level Order. Therefore, root is stored at index 0, middle nodes from index 1.
The leafs are stored from the index (2^height)-1.

E.g. If Merkle Tree can store 4 leafs, then the tree height is 2. Root is at height 0, middle nodes
are at height 1, and 4 leafs are at height 2. The leafs are stored at indexes from 3 to 6. The root node is stored
at index 0.



Algorithm:
1) fetch the index for oldLeaf
2) calculate the hash value of newLeaf
3) store the hash at the index in the merkle tree
4) update leafsToIndexes map:  insert -1 at oldLeaf, insert index at newLeaf
5) update the hash values of middle nodes and root nodes based on the index


Input:
	1) (string) - oldLeaf value
	2) (string) - newLeaf value

Output:
	1) (int) - index of newLeaf in the merkle tree accumulator
	2) (string) - hash value of newLeaf that is stored in the merkle tree accumulator
*/

func (accumulator *MerkleTreeAccumulator2)  UpdateLeaf(oldLeaf string, newLeaf string) (int,string) {

	//1) fetch the index for oldLeaf
	index := accumulator.leafsToIndexes[oldLeaf]

	//2) calculate the hash value of newLeaf
	hashValue := accumulator.GetHash(newLeaf)

	//3) store the hash at the index in the merkle tree
	accumulator.Tree[index].Value = hashValue


	//4) update leafsToIndexes map:  insert -1 at oldLeaf, insert index at newLeaf
	accumulator.leafsToIndexes[oldLeaf] = -1
	accumulator.leafsToIndexes[newLeaf] = index

	// calculate hash from leaf to root
	accumulator.UpdateMiddleandRootNodes(index)

	//5) update the hash values of middle nodes and root nodes based on the index
	accumulator.RootHash=accumulator.Tree[0].Value


	return index, hashValue
}

/*
GetProof generates witness for a leaf.

The witness consists of
1) hash values of the leaf's sibling
2) hash values of nodes's siblings that are in the path from the leaf to the root.

E.g. let's consider the following merkle tree.
----- H(H(1,2), H(3,4)) ------------
----- H(1,2) ------ H(3,4) ---------
--- H(1) - H(2) -- H(3) - H(4) -----
---- "1" --- "2" --- "3" --- "4" ---

The merkle tree accumulator consists:
[ 0 -> H(H(1,2), H(3,4)), 1 ->  H(1,2), 2-> H(3,4), 3 -> "1", 4 -> "2", 5 -> "3",  6 -> "4" ]

The witness of "4" ---> ( H(3), H(1,2) )

Algorithm:
1)
First, the index value of leaf is fetched using the leafsToIndexes map.
 */
func (accumulator *MerkleTreeAccumulator2)  GetProof(leaf string) *MerkleProof{

	merkleProof := &MerkleProof{}
	merkleProof.Witnesses = make(map[int]*Witness)
	merkleProof.Order=make([]int, 0, accumulator.Height)

	index := accumulator.leafsToIndexes[leaf]

	if index==-1{
		zap.S().Infoln("MERKLE TREE ACCUMULATOR- \t the element is not present in the merkle tree accumulator")
		return nil
	}

	merkleProof.LeafHash=accumulator.Tree[index].Value
	merkleProof.LeafValue = leaf

	for i := accumulator.Height; i > 0 ; i-- {

		parentIndex := int(math.Floor(float64((index - 1) / 2)))
		//temp := int(math.Pow(2, float64(parentIndex)))
		temp := 2 * parentIndex
		if parentIndex==0{
			temp = 0
		}
		leftChildIndex := temp+1
		rightChildIndex := temp+2

		//zap.S().Infoln("MERKLE TREE ACCUMULATOR- \t index: ", index, "\t parent index: ", parentIndex, "\t temp: ", temp,
		//	"\t left child index: ", leftChildIndex, "\t right child index: ", rightChildIndex)

		witnessIndex := leftChildIndex
		pos := LEFT
		if index==leftChildIndex{
			witnessIndex=rightChildIndex
			pos = RIGHT
		}

		witnessValue := accumulator.Tree[witnessIndex].Value


		merkleProof.Order = append(merkleProof.Order, witnessIndex)
		wit := Witness{
			HashValue: witnessValue,
			Position:  pos,
		}
		merkleProof.Witnesses[witnessIndex] = &wit


		index = parentIndex
	}

	merkleProof.OrderedWitnesses = accumulator.OrderWitnesses(*merkleProof)

	return merkleProof
}


func (accumulator *MerkleTreeAccumulator2)  OrderWitnesses(merkleProof MerkleProof) []*Witness {

	orderedWitnesses := make([]*Witness,0)
	for _, witnessIndex := range merkleProof.Order{
		wit := &Witness{
			HashValue: merkleProof.Witnesses[witnessIndex].HashValue,
			Position:  merkleProof.Witnesses[witnessIndex].Position,
		}
		orderedWitnesses = append(orderedWitnesses, wit)
	}
	return orderedWitnesses
}



func (accumulator *MerkleTreeAccumulator2)  VerifyProof(leafHash string, witnesses []*Witness, rootHash string) bool {

	root := rootHash

	if rootHash==""{
		root = accumulator.Tree[0].Value
	}

	currentHash := leafHash
	hashValue := currentHash


	var intermediateHashes string

	for _,witness := range witnesses{
		if witness.Position==LEFT{
			hashValue = accumulator.GetHash( witness.HashValue + currentHash)
		}
		if witness.Position==RIGHT{
			hashValue = accumulator.GetHash(currentHash + witness.HashValue)
		}

		currentHash = hashValue
		intermediateHashes = intermediateHashes + accumulator.PrintShortString(currentHash)
	}

	zap.S().Infoln("MERKLE TREE ACCUMULATOR: \t root: ",accumulator.PrintShortString(root), "\t leaf hash: ",
		accumulator.PrintShortString(leafHash), "\t hashes: ", intermediateHashes)


	if hashValue==root{
		return true
	} else{
		return false
	}

}

func (accumulator *MerkleTreeAccumulator2)  GetHashValueOfLeaf(leaf string) string {
	index := accumulator.leafsToIndexes[leaf]

	if index==-1{
		zap.S().Infoln("MERKLE TREE ACCUMULATOR- \t the element is not present in the merkle tree accumulator")
		return ""
	}
	return accumulator.Tree[index].Value
}

func (accumulator *MerkleTreeAccumulator2)  PrintTree() {
	var res string
	res = res +"Merkle Tree: "
	for i:=0;i<accumulator.TotalNodes;i++{
		value := accumulator.Tree[i].Value
		res = res + fmt.Sprintf("\t %d: %s",i,accumulator.PrintShortString(value))
	}
	zap.S().Infoln(res)
}

func (accumulator *MerkleTreeAccumulator2)  ProofToString(proof *MerkleProof) string{
var res string
res = res + fmt.Sprintf("leaf hash: %s \t",accumulator.PrintShortString(proof.LeafHash))
res = res + "witness order: "+fmt.Sprint(proof.Order)+"\t witnesses: "

for  index, witness := range proof.Witnesses{
	res = res + fmt.Sprintf("  %d: %s",index,accumulator.PrintShortString(witness.HashValue))
}
return res
}

func (accumulator *MerkleTreeAccumulator2)  WitnessesToString(witnesses []*Witness) string{
	var res string
	for  _, witness := range witnesses {
		res = res + accumulator.PrintShortString(witness.HashValue)
	}
	return res
}

func (accumulator *MerkleTreeAccumulator2) PrintShortString(input string) string{

	if len(input)>0{
		input = input[:SHORT_STRING_SIZE] + ".."
	}
	return fmt.Sprintf("%s",input)
}




