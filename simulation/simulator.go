package simulation

import (
	"encoding/json"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/blockchain"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/issuer"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)



func Start(config config.Config){
		DeployContract(&config)
		//zap.S().Infoln("smart contract: ",config.SmartContractAddress)
		PerformExperiment(config)

}

func StartExperiments(config config.Config){

	experiments := config.ExpParamters

	for _, exp := range experiments{
		DeployContract(&config)
		//zap.S().Infoln("smart contract: ",config.SmartContractAddress)
		SetUpExpParamters(&config, *exp)
		PerformExperiment(config)
	}
}

func DeployContract(conf *config.Config){
	address, err := blockchain.DeployContract(*conf)

	if err != nil {
		zap.S().Errorln("error deploying contract")
	}

	conf.SmartContractAddress = address
}

func SetUpExpParamters(conf *config.Config, exp config.Experiment){
	conf.ExpectedNumberOfTotalVCs= uint(exp.TotalVCs)
	conf.ExpectedNumberofRevokedVCs= uint(exp.RevokedVCs)
	conf.FalsePositiveRate=exp.FalsePositiveRate
	conf.MtLevelInDLT= uint(exp.MtLevelInDLT)
	conf.MtDepth= uint(exp.MtDepth)
}

func PerformExperiment(config config.Config){

	issuer1 := issuer.CreateIssuer(config)
	remainingSpace := int(math.Pow(2, float64(config.MtDepth-1)))-int(config.ExpectedNumberOfTotalVCs)
	vcDummies := issuer1.GenerateDummyVCs(int(config.ExpectedNumberOfTotalVCs)+remainingSpace)

	issuer1.IssueBulk(config, vcDummies, len(vcDummies))

	for _, vc := range vcDummies{
		issuer1.UpdateMerkleProof(*vc)
	}

	vcs:= []*verifiable.Credential{}

	for i:=0; i<int(config.ExpectedNumberOfTotalVCs);i++{
		vcs = append(vcs, vcDummies[i])
	}

	for _, vc := range vcs{
			issuer1.VerifyTest(*vc)
	}



	var amountPaid int64
	amountPaid = 0
	numberOfAffectedVCs := 0
	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	revokedVCs := make([]string, totalRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for i, counter:=0, 0; counter< totalRevokedVCs; {

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
				counter++
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
	numberOfOccuredFalsePositives := 0
	numberOfVCsRetrievedWitnessFromIssuer := 0
	for _, vc := range vcs {
		falsePositiveStatus, isAffectedInMTAcc = issuer1.VerifyTest(*vc)
		if falsePositiveStatus == true {
			numberOfOccuredFalsePositives++
			if isAffectedInMTAcc == true {
				numberOfVCsRetrievedWitnessFromIssuer++
			}
		}
	}

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
		NumberOfFalsePositives:                numberOfOccuredFalsePositives,
		AmountPaid:                            amountPaid,
		NumberOfWitnessUpdatesForMT:                   numberOfAffectedVCs,
		NumberOfVCsRetrievedWitnessFromIssuer: numberOfVCsRetrievedWitnessFromIssuer,
		NumberOfWitnessUpdatesSaved:         numberOfOccuredFalsePositives-numberOfVCsRetrievedWitnessFromIssuer,
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
	jsonRes, _ := json.MarshalIndent(results,"","")
	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	ioutil.WriteFile("Simulation/results.json", jsonRes, 0644)

}

func BloomFilterConfigurationGenerators(totalNumberOfVCs uint, falsePositiveRate float64) (uint, uint) {
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(totalNumberOfVCs, falsePositiveRate)
	return size, numberOfIndexesPerEntry
}




