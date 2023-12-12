package simulation

import (
	"encoding/json"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/issuer"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Experiment struct {
	totalVCs int `json:"TotalVCs"`
	revokedVCs int `json:"TotalRevokedVCs"`
	falsePositiveRate float64 `json:"FalsePositiveRate"`
	mtLevelInDLT int `json:"MtLevelInDLT"`
	mtDepth int `json:"MtDepth"`
}

func Start(config config.Config){


	issuer1 := issuer.CreateIssuer(config)
	vcs := issuer1.GenerateDummyVCs(int(config.ExpectedNumberOfTotalVCs))

	issuer1.IssueBulk(config, vcs, len(vcs))

	for _, vc := range vcs{
		issuer1.UpdateMerkleProof(*vc)
	}

	var wg sync.WaitGroup

	for _, vc := range vcs{
		wg.Add(1)
		go func(vc verifiable.Credential) {
			defer wg.Done()
			issuer1.VerifyTest(vc)
		}(*vc)
	}

	wg.Wait()

	var amountPaid int64
	amountPaid = 0
	numberOfAffectedVCs := 0
	//numberOfOccuredFalsePositives := 0
	//numberOfVCsRetrievedWitnessFromIssuer := 0
	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	revokedVCs := make([]string, totalRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for i, counter:=0, 0; counter< totalRevokedVCs; counter++{

		i = 2
		for {
			vcID := vcs[i].ID
			isalreadyRevoked := false
			for _, revokedId := range revokedVCs {
				if vcID == revokedId {
					isalreadyRevoked = true
					break
				}
			}
			if isalreadyRevoked==false{
				n, amount := issuer1.Revoke(config, *vcs[i])
				numberOfAffectedVCs += n
				amountPaid = amountPaid + amount;
				amountPaid = amountPaid/2;
				revokedVCs = append(revokedVCs, vcID)
				break
			}
			rand.Seed(time.Now().UnixNano())
			i = rand.Intn(int(config.ExpectedNumberOfTotalVCs))
		}
	}

	//var falsePositiveStatus bool
	//falsePositiveStatus = false
	//var isAffectedInMTAcc bool
	//isAffectedInMTAcc = false
	var numberOfOccuredFalsePositives atomic.Uint64
	var numberOfVCsRetrievedWitnessFromIssuer atomic.Uint64
	for _, vc := range vcs{
		wg.Add(1)
		go func(vc verifiable.Credential) {
			defer wg.Done()
			falsePositiveStatus, isAffectedInMTAcc := issuer1.VerifyTest(vc)
			if falsePositiveStatus==true{
				numberOfOccuredFalsePositives.Add(1)
				if isAffectedInMTAcc==true{
					numberOfVCsRetrievedWitnessFromIssuer.Add(1)
				}
			}
		}(*vc)
	}
	wg.Wait()
		//falsePositiveStatus, isAffectedInMTAcc = issuer1.VerifyTest(*vc)
		//// it means false positive
		//if falsePositiveStatus==true{
		//	numberOfOccuredFalsePositives++
		//	if isAffectedInMTAcc==true{
		//		numberOfVCsRetrievedWitnessFromIssuer++
		//	}
		//}

	size, k := BloomFilterConfigurationGenerators(config.ExpectedNumberofRevokedVCs,config.FalsePositiveRate)
	result := &Results{
		TotalVCs:                              int(config.ExpectedNumberOfTotalVCs),
		RevokedVCs:                            int(config.ExpectedNumberofRevokedVCs),
		FalsePositiveRate:                     config.FalsePositiveRate,
		MtDepth:                               int(config.MtDepth),
		MtLevelInDLT:                          int(config.MtLevelInDLT),
		NumberOfFalsePositives:                int(numberOfOccuredFalsePositives.Load()),
		AmountPaid:                            amountPaid,
		NumberOfAffectedVCs:                   numberOfAffectedVCs,
		NumberOfVCsRetrievedWitnessFromIssuer: int(numberOfVCsRetrievedWitnessFromIssuer.Load()),
		NumberOfWitnessUpdatesSaved:          int(numberOfOccuredFalsePositives.Load())-int(numberOfVCsRetrievedWitnessFromIssuer.Load()),
		BloomFilterSize:                       int(size),
		BloomFilterIndexesPerEntry:            int(k),
	}

	//jsonObj, err := json.Marshal(result)
	//if err!=nil{
	//	zap.S().Errorln("marshall json errror: ",err)
	//}
	zap.S().Infoln("SIMULATOR : \t results: ", result.String())

	WriteToFile(*result)



}


func  WriteToFile( result Results) {

	var results []Results
	jsonFile, _ := os.Open("Simulation/results.json")
	byte, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byte, &results)
	results = append(results, result)
	jsonRes, _ := json.Marshal(results)
	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	ioutil.WriteFile("Simulation/results.json", jsonRes, 0644)

}

func BloomFilterConfigurationGenerators(totalNumberOfVCs uint, falsePositiveRate float64) (uint, uint) {
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(totalNumberOfVCs, falsePositiveRate)
	return size, numberOfIndexesPerEntry
}




