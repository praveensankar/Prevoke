package simulation

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"strconv"
	"time"
)
func GenerateVCIDs(conf config.Config) ([]string) {
	remainingSpace := int(math.Pow(2, float64(conf.MTHeight)))-int(conf.ExpectedNumberOfTotalVCs)
	totalVCs := int(conf.ExpectedNumberOfTotalVCs)+remainingSpace
	rand.Seed(time.Now().UnixNano())
	vcCounter := rand.Intn(100000)
	var vcIDs []string
	for i:=0 ; i<totalVCs ; i++ {
		vcCounter = vcCounter + 1
		vcId := strconv.Itoa(vcCounter)
		vcIDs = append(vcIDs, vcId)
	}
	return vcIDs
}

func InsertIntoMT(conf config.Config, vcIDs []string, mtAcc *techniques.MerkleTreeAccumulator2) (map[string]int){
	mtIndexStore := make(map[string]int)
	for i:=0;i< len(vcIDs);i++{
		mtIndex, _ := mtAcc.AddLeaf(vcIDs[i])
		mtIndexStore[vcIDs[i]] = mtIndex
	}
	return mtIndexStore
}


func GenerateRevokedVCIDs(conf config.Config, vcIDs []string) ([]string) {
	revocationBatchSize := int(conf.RevocationBatchSize)

	totalRevokedVCs := int(conf.ExpectedNumberofRevokedVCs)
	var revokedVCs []string
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for batch := 0; batch < int(int64(math.Ceil(float64(totalRevokedVCs/revocationBatchSize)))); batch++ {
		revokedVCsInBatch := make([]string, 0)
		for i, counter := 0, 0; counter < revocationBatchSize; {

			i = 2
			for {
				vcID := fmt.Sprintf("%v", vcIDs[i])
				isalreadyRevoked := false
				for _, revokedId := range revokedVCs {
					if vcID == revokedId {
						isalreadyRevoked = true
						break
					}
				}
				if isalreadyRevoked == false {
					revokedVCsInBatch = append(revokedVCsInBatch, vcID)
					revokedVCs = append(revokedVCs, vcID)
					counter++
					break
				}
				rand.Seed(time.Now().UnixNano())
				i = rand.Intn(int(conf.ExpectedNumberOfTotalVCs))
			}
		}
	}
	return revokedVCs
}


func RevokeVCs(conf config.Config, bf *techniques.BloomFilter, acc *techniques.MerkleTreeAccumulator2, revokedVcIDs []string) (map[string]bool, mapset.Set) {
	revokedVCIDMaps := make(map[string]bool)
	affectedIndexes := mapset.NewSet()

	for i := 0; i < len(revokedVcIDs); i++ {
		revokedVCIDMaps[revokedVcIDs[i]] = true
		bf.RevokeInBloomFilter(revokedVcIDs[i])
		mtIndex, _ :=	acc.UpdateLeaf(revokedVcIDs[i], "-1")
		indexes := UpdateAffectedVCs(conf, mtIndex)
		affectedIndexes = affectedIndexes.Union(indexes)
	}
	return revokedVCIDMaps, affectedIndexes

}


// returns number of vcs that are affected
func UpdateAffectedVCs(conf config.Config, mtIndex int) (mapset.Set ) {


	height := int(conf.MTHeight)
	levelStoredInDLT := int(conf.MtLevelInDLT)

	var numberOfEstimatedAffectedVCs int
	affectedIndexes := mapset.NewSet()


	if height==levelStoredInDLT{
		return affectedIndexes
	} else {
		numberOfEstimatedAffectedVCs = int(math.Pow(2, float64(height-levelStoredInDLT)))
	}


	foundBlock := false
	firstLeafsIndex := int(math.Pow(2, float64(height)))-1
	lastLeafsIndex := int(math.Pow(2, float64(height+1)))-1

	for i:=firstLeafsIndex; i<= (lastLeafsIndex-numberOfEstimatedAffectedVCs+1);  i = i + numberOfEstimatedAffectedVCs {
		if foundBlock==true{
			break
		}
		end := i + numberOfEstimatedAffectedVCs
		if mtIndex < end{
			foundBlock=true
			for j:=i; j < int(i + numberOfEstimatedAffectedVCs); j++{
				if mtIndex == j {
					continue
				}
				affectedIndexes.Add(j)


			}
			//zap.S().Infoln("ISSUER: WITNESS UPDATE - \t mt index: ",mtIndex, "\t block starting index: ",i, "\t end index: ", int64(i + numberOfEstimatedAffectedVCs)-1,
			//	"\t affected vcs: ", affectedIndexes)
		}
	}
	return affectedIndexes
}

/*
CalculateNumberOfVCsWouldRetrieveWitnessFromDLT calculates how many valid vcs need to retrieve witness from dlt

First, it computes the list of valid vcs that are affected by the bloom filter
*/
func CalculateNumberOfVCsWouldRetrieveWitnessFromDLT(conf config.Config) {

	experiments := conf.ExpParamters

	for _, exp := range experiments {

		//SetUpExpParamters(&config, *exp)
		//exp.MtHeight=1

		conf.ExpectedNumberOfTotalVCs = uint(exp.TotalVCs)
		conf.ExpectedNumberofRevokedVCs = uint(exp.RevokedVCs)
		conf.FalsePositiveRate = exp.FalsePositiveRate
		conf.MTHeight = uint(exp.MtHeight)
		conf.MtLevelInDLT = uint(exp.MtLevelInDLT)
		conf.RevocationBatchSize = uint(exp.RevocationBatchSize)
		zap.S().Infoln("ISSUER - updated config with experiment config: ", exp.String())

		if exp.TotalVCs != 0 {

			numberOfVCsRetrievingVCsFromDLT := 0
			numberOfFalsePositives := 0
			vcIDs := GenerateVCIDs(conf)
			zap.S().Infoln("vc ids: ", vcIDs)
			var vcIDsFromDLT []string
			var fpVCIDs []string

			mtAcc := techniques.CreateMerkleTreeAccumulator(conf)
			bf := techniques.CreateBloomFilter(conf.ExpectedNumberofRevokedVCs, conf.FalsePositiveRate)

			mtIndexStore := InsertIntoMT(conf, vcIDs, mtAcc)
			zap.S().Infoln("mt index store: ", mtIndexStore)

			//mtAcc.PrintTree()

			revokedVcIDs := GenerateRevokedVCIDs(conf, vcIDs)
			zap.S().Infoln("revoked vc ids: ", revokedVcIDs)

			revokedVCIDMaps, affectedIndexes := RevokeVCs(conf, bf, mtAcc, revokedVcIDs)
			zap.S().Infoln("affected indexes: ", affectedIndexes)

			affectedVCs := make(map[int]string)
			for i:=0;i< len(vcIDs);i++{
				if affectedIndexes.Contains(mtIndexStore[vcIDs[i]])==true{
					affectedVCs[mtIndexStore[vcIDs[i]]]=vcIDs[i]
				}
			}
			zap.S().Infoln("affected vc ids: ", affectedVCs)

			for i := 0; i < int(conf.ExpectedNumberOfTotalVCs); i++ {
				vcId := vcIDs[i]
				if bf.CheckStatusInBloomFilter(vcId) == false {
					if revokedVCIDMaps[vcId] == false {
						fpVCIDs = append(fpVCIDs, vcId)
						numberOfFalsePositives++
						mtIndex := mtIndexStore[vcId]
						if affectedIndexes.Contains(mtIndex)==false {
							numberOfVCsRetrievingVCsFromDLT++
							vcIDsFromDLT = append(vcIDsFromDLT, vcId)
						}
					}
				}
			}
			zap.S().Infoln("false positive vc ids: ",fpVCIDs)
			zap.S().Infoln("VCs that would retrieve witness from DLTs: ", vcIDsFromDLT)
			zap.S().Infoln("number of vcs affected by z levels: ",affectedIndexes.Cardinality())
			zap.S().Infoln("number of false positives: ",numberOfFalsePositives, "\t number of vcs retrieved witness" +
				"from dlts: ",numberOfVCsRetrievingVCsFromDLT)
		}

	}
}


