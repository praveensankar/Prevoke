package simulation

import (
	"encoding/json"
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




	experiment1:= Experiment{
		totalVCs:          128,
		revokedVCs:  50,
		falsePositiveRate: 0.01,
		mtLevelInDLT:      3,
		mtDepth:          8,
	}
	stats := make(map[int]int)
	config.ExpectedNumberOfTotalVCs = uint(experiment1.totalVCs)
	config.ExpectedNumberofRevokedVCs = uint(experiment1.revokedVCs)
	config.MtDepth = uint(experiment1.mtDepth)
	config.MtLevelInDLT = uint(experiment1.mtLevelInDLT)
	issuer1 := issuer.CreateIssuer(config)
	vcs := issuer1.GenerateDummyVCs(int(config.ExpectedNumberOfTotalVCs))

	for _, vc := range vcs{
		issuer1.Issue(*vc)
	}
	for _, vc := range vcs{
		issuer1.UpdateMerkleProof(*vc)
	}
	for _, vc := range vcs{
		issuer1.VerifyTest(*vc)
	}

	numberOfAffectedVCs := 0
	numberOfOccuredFalsePositives := 0
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
	stats[totalRevokedVCs]=numberOfAffectedVCs

	for _, vc := range vcs{
		status := issuer1.VerifyTest(*vc)
		// it means false positive
		if status==true{
			numberOfOccuredFalsePositives++
		}
	}

	result := &Results{
		TotalVCs: experiment1.totalVCs,
		RevokedVCs: experiment1.revokedVCs,
		FalsePositiveRate: experiment1.falsePositiveRate,
		MtDepth: experiment1.mtDepth,
		MtLevelInDLT: experiment1.mtLevelInDLT,
		NumberOfFalsePositives: numberOfOccuredFalsePositives,
		NumberOfAffectedVCs:    numberOfAffectedVCs,
		NumberOfWitnessUpdatesSaved: numberOfAffectedVCs-numberOfOccuredFalsePositives,
	}

	//jsonObj, err := json.Marshal(result)
	//if err!=nil{
	//	zap.S().Errorln("marshall json errror: ",err)
	//}
	zap.S().Infoln("SIMULATOR : \t results: ", result.String())

	WriteToFile(experiment1.totalVCs,experiment1.revokedVCs,experiment1.mtLevelInDLT, *result)



}


func  WriteToFile(numberOfVcs int, numberOfRevokedVcs int, mtLevelInDLT int, result Results) {

	var results []Results
	jsonFile, _ := os.Open("Simulation/results.json")
	byte, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byte, &results)
	results = append(results, result)
	jsonRes, _ := json.Marshal(results)
	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	ioutil.WriteFile("Simulation/results.json", jsonRes, 0644)

}




