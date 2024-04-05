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


	totalRevokedVCs := int(conf.ExpectedNumberofRevokedVCs)
	var revokedVCs []string

	//stores revoked vc IDs
	revokedVCIDs := make(map[string]bool)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	i:=0
	if revocationMode==Random{
		for counter := 0; counter < totalRevokedVCs; counter++ {
			for
			{
				rand.Seed(time.Now().UnixNano())
				i = rand.Intn(int(conf.ExpectedNumberOfTotalVCs))

				vcID := fmt.Sprintf("%v", vcIDs[i])

				if revokedVCIDs[vcID]==true{
					continue
				}
				if revokedVCIDs[vcID]==false{
					revokedVCs = append(revokedVCs, vcID)
					revokedVCIDs[vcID]=true
					break
				}
			}
		}
	}
	if revocationMode==Oldest{
		for counter := 0; counter < totalRevokedVCs; counter++ {
			vcID := fmt.Sprintf("%v", vcIDs[i])
			revokedVCs = append(revokedVCs, vcID)
			i++
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
	//totalVCs:=[]int{10000}
	//totalVCs:=[]int{50000}
	//totalVCs:=[]int{100000}
	totalVCs:=[]int{1000000}
	processRawData := false
	revocationPercentages := []int{1,2,3,4,5,6,7,8,9,10,20,30,40}

	revocationModes := []RevocationMode{Random, Oldest}
	falsePositiveRates:= []float64{0.1,0.01,0.001,0.0001}

	// mtHeight: log_2(<totalVCs>)
	// mtLevelInDLT: 0 to mtHeight


	//totalVCs=[]int{100}
	//falsePositiveRates= []float64{0.1}


	rawFilename := fmt.Sprintf("results/results_computed_raw_50000.json")
	resultFileName := fmt.Sprintf("results/results_computed_1M.json")


	container := Container{}
	var wg sync.WaitGroup

	var exps []config.Experiment

	/*
	Construct the parameters for the experiments
	 */

	for i:=0;i< len(totalVCs);i++ {

		totalVC := totalVCs[i]
		mtHeight := int(math.Ceil(math.Log2(float64(totalVC))))
		//SetUpExpParamters(&config, *exp)
		//exp.MtHeight=1
		conf.ExpectedNumberOfTotalVCs = uint(totalVC)
		conf.MTHeight = uint(mtHeight)
		vcIDs := GenerateVCIDs(conf)

		if totalVC>100000{
			revocationPercentages = []int{1,2,3,4,5,10,20,30}
		}

		for j := 0; j < len(falsePositiveRates); j++ {

			for mtLevelInDLT := 1; mtLevelInDLT <= mtHeight; mtLevelInDLT++ {

				for revocationPercentageCounter := 0; revocationPercentageCounter<len(revocationPercentages); revocationPercentageCounter++{

					for k := 0; k < len(revocationModes); k++ {

						revokedVCCount := int(math.Ceil(float64(totalVC * revocationPercentages[revocationPercentageCounter] / 100)))
						falsePositiveRate := falsePositiveRates[j]
						revocationMode := revocationModes[k]
						exp:= config.Experiment{
							TotalVCs:            totalVC,
							RevokedVCs:          revokedVCCount,
							FalsePositiveRate:   falsePositiveRate,
							MtLevelInDLT:        mtLevelInDLT,
							MtHeight:            mtHeight,
							RevocationBatchSize: 1,
							RevocationMode: string(revocationMode),
							VCIDs: vcIDs,
						}
						exps = append(exps, exp)

					}
				}

				if totalVC>100000{
					mtLevelInDLT++
				}
			}
		}
	}

	goRountineCounter:=0
	for i:=0;i< len(exps);i++ {

		exp := exps[i]
		conf.ExpectedNumberOfTotalVCs = uint(exp.TotalVCs)
		conf.MTHeight = uint(exp.MtHeight)
		conf.ExpectedNumberofRevokedVCs = uint(exp.RevokedVCs)

		conf.FalsePositiveRate = exp.FalsePositiveRate
		conf.MtLevelInDLT = uint(exp.MtLevelInDLT)
		conf.RevocationBatchSize = 1

		vcIDs := exp.VCIDs
		var revocationMode RevocationMode
		if exp.RevocationMode==string(Oldest){
			revocationMode = Oldest
		}
		if exp.RevocationMode==string(Random){
			revocationMode=Random
		}
		//zap.S().Infoln("ISSUER - updated config with experiment config: ", exp.String())

		wg.Add(1)

		go func(conf config.Config, vcIDs []string, processRawData bool, mode RevocationMode, expCounter int, totalExps int) {
			defer wg.Done()
			container.PerformCalculation(conf, vcIDs, processRawData, mode, expCounter, totalExps)
		}(conf, vcIDs, processRawData, revocationMode,i+1, len(exps))


		goRountineCounter++
		if goRountineCounter==150{
			wg.Wait()
			goRountineCounter=0
		}
	}

	wg.Wait()
	if processRawData==true {
		common.WriteFalsePositiveAndWitnessUpdateRawResultsToFile(rawFilename, container.RawResults)
	}
	common.WriteFalsePositiveAndWitnessUpdateResultsToFile(resultFileName, container.Results)
	expEnd := time.Since(expStart)
	zap.S().Infoln("Total time to run the experiments: ", expEnd.Minutes(), "  minutes")
}

func (c *Container) PerformCalculation(conf config.Config, vcIDs []string, processRawData bool, revocationMode RevocationMode,  expCounter int, totalExps int) {
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

	if processRawData==true {
		for x := 0; x < len(vcIDs); x++ {
			if affectedIndexes.Contains(mtIndexStore[vcIDs[x]]) == true {
				affectedVCs[mtIndexStore[vcIDs[x]]] = vcIDs[x]
				affectedVCIDs = append(affectedVCIDs, vcIDs[x])
			}
		}
	}
	//zap.S().Infoln("affected vc ids: ", affectedVCs)

	for y := 0; y < int(conf.ExpectedNumberOfTotalVCs); y++ {
		vcId := vcIDs[y]
		if bf.CheckStatusInBloomFilter(vcId) == false {
			if revokedVCIDMaps[vcId] == false {
				if processRawData==true {
					fpVCIDs = append(fpVCIDs, vcId)
				}
				numberOfFalsePositives++
				mtIndex := mtIndexStore[vcId]
				if affectedIndexes.Contains(mtIndex) == false {
					numberOfVCsRetrievingVCsFromDLT++
					if processRawData==true {
						vcIDsFromDLT = append(vcIDsFromDLT, vcId)
					}
				}
			}
		}
	}
	//zap.S().Infoln("false positive vc ids: ",fpVCIDs)
	//zap.S().Infoln("VCs that would retrieve witness from DLTs: ", vcIDsFromDLT)
	//zap.S().Infoln("number of vcs affected by z levels: ",affectedIndexes.Cardinality())
	zap.S().Infof("exp: %d/%d  time: %d:%d total vcs:%d revoked vcs:%d  false positive rate:%f  mt level in dlt:%d  revocation mode:%s  number of false positives:%d  number of vcs retrieved witness from dlt:%d",
		expCounter, totalExps, time.Now().Hour(), time.Now().Second(), conf.ExpectedNumberOfTotalVCs, conf.ExpectedNumberofRevokedVCs,
		conf.FalsePositiveRate, conf.MtLevelInDLT, revocationMode, numberOfFalsePositives, numberOfVCsRetrievingVCsFromDLT)
	//zap.S().Infoln("exp: ",expCounter,"/",totalExps," timestamp: ", time.Now().Hour(),":",time.Now().Minute(), "total vc: ", conf.ExpectedNumberOfTotalVCs, " revoked vcs: ", conf.ExpectedNumberofRevokedVCs,
	//	" false positive rate: ", conf.FalsePositiveRate, " mt level in dlt: ", conf.MtLevelInDLT,
	//	" revocation mode: ", revocationMode, " number of false positives: ", numberOfFalsePositives, " number of vcs retrieved witness"+
	//		"from dlts: ", numberOfVCsRetrievingVCsFromDLT)

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

	if processRawData==true {
		result.AffectedVCIDs = affectedVCIDs
		result.FalsePositiveResults = fpVCIDs
		result.FetchedWitnessesFromDLT = vcIDsFromDLT
		c.RawResults = append(c.RawResults, *result)
	}

}


