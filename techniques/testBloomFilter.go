package techniques

import "fmt"

func TestBloomFilter(estimatedVCs ...int){
	testBloomFilter(estimatedVCs...);
}


func testBloomFilter(estimatedVCs ...int){

	numberOfVCs := 1000
	for _, count := range estimatedVCs {
		numberOfVCs=count
	}

	fpr:= 0.000001

	bloomFilter1 := CreateBloomFilter(uint(numberOfVCs), fpr)
	fmt.Println("Total vcs: ", numberOfVCs, "\t false positive rate: ", fpr, "\nbloom filter size in bits: ",bloomFilter1.size, "\t number of hash functions: ",bloomFilter1.numberOfIndexesPerEntry)
	statusBool := bloomFilter1.CheckStatusInBloomFilter("vc1")
	var status string
	if statusBool==true{
		status = "vc is valid"
	} else{
		status = "vc is revoked"
	}
	fmt.Println("revocation status of vc1: ", status)

	fmt.Println("revoking vc1")
	indexes:= bloomFilter1.RevokeInBloomFilter("vc1")
	fmt.Println("vc1 indexes in bloom filter: ", indexes)
	statusBool = bloomFilter1.CheckStatusInBloomFilter("vc1")
	statusLocationsBool := bloomFilter1.bloomFilter.TestLocations(indexes)
	if statusBool==true{
		status = "vc is valid"
	} else{
		status = "vc is revoked"
	}
	fmt.Println("revocation status of vc1: ", status)

	if statusLocationsBool==false{
		status = "vc is valid"
	} else{
		status = "vc is revoked"
	}
	fmt.Println("revocation status of vc1 in by testing indexes: ", status)



	fmt.Println("revoking 10 vcs")
	for i := 1; i <= numberOfVCs; i++ {
		revokedId := fmt.Sprintf("id_%d", i)
		bloomFilter1.RevokeInBloomFilter(revokedId)
	}


	fmt.Println("number of collisions: ",bloomFilter1.getCollisions())


}


