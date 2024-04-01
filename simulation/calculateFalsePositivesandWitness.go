package simulation

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/praveensankar/Revocation-Service/common"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"
)


type RevocationMode string

// Random - revokes vcs randomly
// Oldest - revokes the oldest vcs
const (
	Random    RevocationMode = "random"
	Oldest = "oldest"
)

type FPandWitnessResult struct {
	NumberOfFalsePositives int
	NumberOfVCsRetrievingVCsFromDLT int
	VcIDsFromDLT []string
	FpVCIDs []string
	AffectedVCIDs []string
}

type Container struct {
	mu       sync.Mutex
	Results []common.FalsePositiveAndWitnessUpdateResults
	RawResults []common.FalsePositiveAndWitnessUpdateResults
}

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


func GenerateRevokedVCIDs(conf config.Config, vcIDs []string, revocationMode RevocationMode) ([]string) {
	revocationBatchSize := int(conf.RevocationBatchSize)

	totalRevokedVCs := int(conf.ExpectedNumberofRevokedVCs)
	var revokedVCs []string


	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for batch := 0; batch < int(int64(math.Ceil(float64(totalRevokedVCs/revocationBatchSize)))); batch++ {
		revokedVCsInBatch := make([]string, 0)
		for i, counter := 0, 0; counter < revocationBatchSize; {


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

				if revocationMode==Random{
					rand.Seed(time.Now().UnixNano())
					i = rand.Intn(int(conf.ExpectedNumberOfTotalVCs))
				}
				if revocationMode==Oldest{
					i++
				}

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

	expStart := time.Now()
	totalVCs:=[]int{10000, 50000, 100000, 1000000}
	//revocationPercentages = 1 to 100 (increment by 1%)

	revocationModes := []RevocationMode{Random, Oldest}
	falsePositiveRates:= []float64{0.1,0.01,0.001,0.0001}

	// mtHeight: log_2(<totalVCs>)
	// mtLevelInDLT: 0 to mtHeight


	//totalVCs=[]int{100}
	//falsePositiveRates= []float64{0.1}


	rawFilename := fmt.Sprintf("results/results_computed_raw.json")
	resultFileName := fmt.Sprintf("results/results_computed.json")
	var results []common.FalsePositiveAndWitnessUpdateResults
	var rawResults []common.FalsePositiveAndWitnessUpdateResults


	container := Container{}
	var wg sync.WaitGroup


	for i:=0;i< len(totalVCs);i++ {

		totalVC := totalVCs[i]
		mtHeight := int(math.Ceil(math.Log2(float64(totalVC))))

		//SetUpExpParamters(&config, *exp)
		//exp.MtHeight=1

		conf.ExpectedNumberOfTotalVCs = uint(totalVC)
		conf.MTHeight = uint(mtHeight)
		vcIDs := GenerateVCIDs(conf)

		for j := 0; j < len(falsePositiveRates); j++ {

			for mtLevelInDLT := 1; mtLevelInDLT <= mtHeight; mtLevelInDLT++ {

				for revocationPercentage := 1; revocationPercentage <= 100; revocationPercentage = revocationPercentage+5 {

					for k := 0; k < len(revocationModes); k++ {

						revokedVCCount := int(math.Ceil(float64(totalVC * revocationPercentage / 100)))
						falsePositiveRate := falsePositiveRates[j]
						revocationMode := revocationModes[k]
						conf.ExpectedNumberofRevokedVCs = uint(revokedVCCount)

						conf.FalsePositiveRate = falsePositiveRate
						conf.MTHeight = uint(mtHeight)
						conf.MtLevelInDLT = uint(mtLevelInDLT)
						conf.RevocationBatchSize = 1
						//zap.S().Infoln("ISSUER - updated config with experiment config: ", exp.String())

						wg.Add(1)

						go func(conf config.Config, mode RevocationMode) {
							defer wg.Done()
							container.PerformCalculation(conf, vcIDs, mode)
						}(conf, revocationMode)





					}
				}
			}
		}
	}
	wg.Wait()
	common.WriteFalsePositiveAndWitnessUpdateResultsToFile(rawFilename, rawResults)
	common.WriteFalsePositiveAndWitnessUpdateResultsToFile(resultFileName, results)
	expEnd := time.Since(expStart)
	zap.S().Infoln("Total time to run the experiments: ", expEnd.Minutes(), "  minutes")
}

func (c *Container) PerformCalculation(conf config.Config, vcIDs []string, revocationMode RevocationMode) {
	numberOfVCsRetrievingVCsFromDLT := 0
	numberOfFalsePositives := 0

	//zap.S().Infoln("vc ids: ", vcIDs)
	var vcIDsFromDLT []string
	var fpVCIDs []string


	mtAcc := techniques.CreateMerkleTreeAccumulator(conf)
	bf := techniques.CreateBloomFilter(conf.ExpectedNumberofRevokedVCs, conf.FalsePositiveRate)

	mtIndexStore := InsertIntoMT(conf, vcIDs, mtAcc)
	//zap.S().Infoln("mt index store: ", mtIndexStore)

	//mtAcc.PrintTree()

	revokedVcIDs := GenerateRevokedVCIDs(conf, vcIDs, revocationMode)
	//zap.S().Infoln("revoked vc ids: ", revokedVcIDs)

	revokedVCIDMaps, affectedIndexes := RevokeVCs(conf, bf, mtAcc, revokedVcIDs)
	//zap.S().Infoln("affected indexes: ", affectedIndexes)

	affectedVCs := make(map[int]string)
	var affectedVCIDs []string
	for x := 0; x < len(vcIDs); x++ {
		if affectedIndexes.Contains(mtIndexStore[vcIDs[x]]) == true {
			affectedVCs[mtIndexStore[vcIDs[x]]] = vcIDs[x]
			affectedVCIDs = append(affectedVCIDs, vcIDs[x])
		}
	}
	//zap.S().Infoln("affected vc ids: ", affectedVCs)

	for y := 0; y < int(conf.ExpectedNumberOfTotalVCs); y++ {
		vcId := vcIDs[y]
		if bf.CheckStatusInBloomFilter(vcId) == false {
			if revokedVCIDMaps[vcId] == false {
				fpVCIDs = append(fpVCIDs, vcId)
				numberOfFalsePositives++
				mtIndex := mtIndexStore[vcId]
				if affectedIndexes.Contains(mtIndex) == false {
					numberOfVCsRetrievingVCsFromDLT++
					vcIDsFromDLT = append(vcIDsFromDLT, vcId)
				}
			}
		}
	}
	//zap.S().Infoln("false positive vc ids: ",fpVCIDs)
	//zap.S().Infoln("VCs that would retrieve witness from DLTs: ", vcIDsFromDLT)
	//zap.S().Infoln("number of vcs affected by z levels: ",affectedIndexes.Cardinality())
	zap.S().Infoln("total vc: ", conf.ExpectedNumberOfTotalVCs, " revoked vcs: ", conf.ExpectedNumberofRevokedVCs,
		" false positive rate: ", conf.FalsePositiveRate, " mt level in dlt: ", conf.MtLevelInDLT,
		" revocation mode: ", revocationMode, " number of false positives: ", numberOfFalsePositives, " number of vcs retrieved witness"+
			"from dlts: ", numberOfVCsRetrievingVCsFromDLT)

	result := common.CreateFalsePositiveAndWitnessUpdateResults()
	result.TotalVCs = int(conf.ExpectedNumberOfTotalVCs)
	result.RevokedVCs = int(conf.ExpectedNumberofRevokedVCs)
	result.FalsePositiveRate =  conf.FalsePositiveRate
	result.MTHeight = int(conf.MTHeight)
	result.MtLevelInDLT = int(conf.MtLevelInDLT)
	result.NumberOfFalsePositives = numberOfFalsePositives
	result.NumberOfVCsRetrievedWitnessFromDLT = numberOfVCsRetrievingVCsFromDLT
	result.RevocationMode = string(revocationMode)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.Results = append(c.Results, *result)
	result.AffectedVCIDs = affectedVCIDs
	result.FalsePositiveResults = fpVCIDs
	result.FetchedWitnessesFromDLT = vcIDsFromDLT
	c.RawResults = append(c.RawResults, *result)


}


