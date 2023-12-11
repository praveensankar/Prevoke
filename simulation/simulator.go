package simulation

import (
	"encoding/json"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/issuer"
	"go.uber.org/zap"
	"io/ioutil"
	"math/rand"
	"os"
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
	for _, vc := range vcs{
		issuer1.VerifyTest(*vc)
	}

	numberOfAffectedVCs := 0
	numberOfOccuredFalsePositives := 0
	numberOfVCsRetrievedWitnessFromIssuer := 0
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
				numberOfAffectedVCs += issuer1.Revoke(config, *vcs[i])
				revokedVCs = append(revokedVCs, vcID)
				break
			}
			rand.Seed(time.Now().UnixNano())
			i = rand.Intn(int(config.ExpectedNumberOfTotalVCs))
		}
	}

	var falsePositiveStatus bool
	falsePositiveStatus = false
	var isAffectedInMTAcc bool
	isAffectedInMTAcc = false

	for _, vc := range vcs{
		falsePositiveStatus, isAffectedInMTAcc = issuer1.VerifyTest(*vc)
		// it means false positive
		if falsePositiveStatus==true{
			numberOfOccuredFalsePositives++
			if isAffectedInMTAcc==true{
				numberOfVCsRetrievedWitnessFromIssuer++
			}
		}
	}
	size, k := BloomFilterConfigurationGenerators(config.ExpectedNumberofRevokedVCs,config.FalsePositiveRate)
	result := &Results{
		TotalVCs:                    int(config.ExpectedNumberOfTotalVCs),
		RevokedVCs:                  int(config.ExpectedNumberofRevokedVCs),
		FalsePositiveRate:           config.FalsePositiveRate,
		MtDepth:                     int(config.MtDepth),
		MtLevelInDLT:                int(config.MtLevelInDLT),
		NumberOfFalsePositives:      numberOfOccuredFalsePositives,
		NumberOfAffectedVCs:         numberOfAffectedVCs,
		NumberOfVCsRetrievedWitnessFromIssuer: numberOfVCsRetrievedWitnessFromIssuer,
		NumberOfWitnessUpdatesSaved: numberOfAffectedVCs-numberOfOccuredFalsePositives,
		BloomFilterSize:             int(size),
		BloomFilterIndexesPerEntry: int(k),
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




