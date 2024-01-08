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

func (accumulator *MerkleTreeAccumulator2)  AddLeaf(value string) string {

	if accumulator.CurrentIndex > accumulator.TotalNodes{
		zap.S().Errorln("merkle tree accumulator is full")
		return ""
	}

	index := accumulator.CurrentIndex
	hashValue := accumulator.GetHash(value)
	accumulator.Tree[index].Value = hashValue
	accumulator.leafsToIndexes[value] = index

	// calculate hash from leaf to root
	accumulator.UpdateMiddleandRootNodes(index)

	accumulator.CurrentIndex++
	accumulator.RootHash=accumulator.Tree[0].Value
	return hashValue

}

func (accumulator *MerkleTreeAccumulator2)  UpdateMiddleandRootNodes(index int) {

	for i := accumulator.Height -1 ; i > 0 ; i--{
		parentIndex := int(math.Floor(float64((index - 1) / 2)))
		leftChildIndex := int(math.Pow(2, float64(parentIndex)))+1
		rightChildIndex := int(math.Pow(2, float64(parentIndex)))+2

		lefChildValue := accumulator.Tree[leftChildIndex].Value
		rightChildValue := accumulator.Tree[rightChildIndex].Value

		hashValue := accumulator.GetHash(lefChildValue + rightChildValue)
		accumulator.Tree[parentIndex].Value = hashValue

		index = parentIndex
	}

	accumulator.Tree[0].Value = accumulator.GetHash(accumulator.Tree[1].Value+accumulator.Tree[2].Value)
}


func (accumulator *MerkleTreeAccumulator2)  UpdateLeaf(oldLeaf string, newLeaf string) string {
	index := accumulator.leafsToIndexes[oldLeaf]
	accumulator.leafsToIndexes[oldLeaf] = -1
	hashValue := accumulator.GetHash(newLeaf)
	accumulator.Tree[index].Value = hashValue
	accumulator.leafsToIndexes[newLeaf] = index

	// calculate hash from leaf to root
	accumulator.UpdateMiddleandRootNodes(index)

	accumulator.RootHash=accumulator.Tree[0].Value
	return hashValue
}

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
		temp := int(math.Pow(2, float64(parentIndex)))
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




