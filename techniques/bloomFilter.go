package techniques

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/deckarep/golang-set"
	"math/big"
)


// Todo: change the value of this constant based on value in smart contract
//const NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER = 4;

type BloomFilter struct{
	bloomFilter *bloom.BloomFilter
	numberOfIndexesPerEntry uint
	size uint
	assignedIndexes mapset.Set
	collisions uint
}
/*
This function creates a new bloomFilter objects.
At first, the size of bloom filter and number of hash functions required  are estimatesd based on two input parameters.
Then a new BloomFilter object is created. and retured.

Inputs:
	TotalNumberofVCs - number of VCs entities expects to issue in its lifetime
	falsePositiveRate - false positive rate of bloomfilter

Output:
	BloomFilter object
	error
*/
func CreateBloomFilter(expectedNumberOfRevokedVCs uint, falsePositiveRate float64) (*BloomFilter){
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(expectedNumberOfRevokedVCs, falsePositiveRate)
	bf := bloom.NewWithEstimates(expectedNumberOfRevokedVCs,falsePositiveRate)
	newBloomFilter := BloomFilter{
		bloomFilter: bf,
		size: size,
		numberOfIndexesPerEntry: numberOfIndexesPerEntry,
	}


	//zap.S().Infoln("BlOOM FILTER: size : ", size, "\t number of indexes per entry : ", numberOfIndexesPerEntry)

	//if numberOfIndexesPerEntry !=NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER{
	//	zap.S().Errorln("ERROR:    \t bloom filter index mismatch.", numberOfIndexesPerEntry, "is given by go code")
	//}

	newBloomFilter.assignedIndexes = mapset.NewSet()


	return &newBloomFilter
}



/*
Output:
	true - the VC is not revoked.
	false - the VC is probably revoked.
*/
func (bf BloomFilter) CheckStatusInBloomFilter(vc string) bool{
	input := []byte(vc)
	isPresentInBloomFilter := bf.bloomFilter.Test(input)

	// if present in bloom filter, then the vc is revoked. so retuns false to indicate the status that vc is not valid
	// otherwise it returns true, to indicate that vc is valid
	if isPresentInBloomFilter==true{
		return false
	}else{
		return true
	}
}

/*
CheckBitInBloomFilter checks an index in the bloomfilter.
Returns
	true - all indexes are set in the BloomFilter
	false - if one or more indexes are not set
 */
func (bf BloomFilter) CheckIndexesInBloomFilter(indexes []uint64) bool {
	isPresentInBloomFilter := bf.bloomFilter.TestLocations(indexes)
	if isPresentInBloomFilter==true{
		return false
	}else{
		return true
	}
}

/*
This function revokes a VC. The bloom filter is set with the revoked VC.

Input:
	vc: unique string representing the vc that is going to be revoked

Output:
	the indexes of the revoked VC in the bloomfilter.
 */
func (bf *BloomFilter) RevokeInBloomFilter(vc string) []uint64 {
	input := []byte(vc)
	bf.bloomFilter.Add(input)
	indexes := bloom.Locations(input, bf.numberOfIndexesPerEntry)
	for i:=0; i< len(indexes);i++{
		indexes[i] = indexes[i]%uint64(bf.size)
}

	// keep track of assigned indexes
	for _, value := range indexes{
		if bf.assignedIndexes.Contains(value)==true{
			bf.collisions = bf.collisions+1
		}
		bf.assignedIndexes.Add(value)
	}

	return indexes
}

/*
This function retuns the indexes for a vc.

Input:
	vc: unique string representing the vc

Output:
	the indexes of the VC in the bloomfilter.
*/
func (bf *BloomFilter) GetIndexes(vc string) []*big.Int {
	input := []byte(vc)
	indexes := bloom.Locations(input, bf.numberOfIndexesPerEntry)
	for i:=0; i< len(indexes);i++{
		indexes[i] = indexes[i]%uint64(bf.size)
	}
	results := []*big.Int{}
	for i:=0; i< len(indexes);i++ {
		results = append(results, big.NewInt(int64(indexes[i])))
	}
	return results
}


func (bf BloomFilter) getCollisions() uint{
return bf.collisions
}

