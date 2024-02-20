package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/deckarep/golang-set"
	"github.com/praveensankar/Revocation-Service/blockchain"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/issuer"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/vc"
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
	conf.MTHeight=uint(exp.MtHeight)
}

func PerformExperiment(config config.Config){
	issuer1 := issuer.CreateIssuer(config)
	publicKey, _ := issuer1.BbsKeyPair[0].PublicKey.Marshal()
	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
	claimsSet := issuer1.GenerateMultipleDummyVCClaims(int(config.ExpectedNumberOfTotalVCs)+remainingSpace)
	revocationBatchSize := int(config.RevocationBatchSize)

	issuer1.IssueBulk(claimsSet, int(config.ExpectedNumberOfTotalVCs)+remainingSpace)

	credentials := issuer1.CredentialStore
	for _, vc := range credentials{
		issuer1.UpdateMerkleProof(vc)
	}

	vcs:= []models.VerifiableCredential{}

	for i:=0; i<int(config.ExpectedNumberOfTotalVCs);i++{
		vcs = append(vcs, credentials[i])
	}

	//for _, vc := range vcs{
	//	issuer1.VerifyTest(*vc)
	//}



	var amountPaid int64
	amountPaid = 0
	affectedIndexes := mapset.NewSet()
	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	revokedVCs := make([]string, totalRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for batch:=0; batch<int(int64(math.Ceil(float64(totalRevokedVCs/revocationBatchSize)))); batch++ {
		revokedVCsInBatch := make([]string, 0)
		for i, counter := 0, 0; counter < revocationBatchSize; {

			i = 2
			for {
				vcID := fmt.Sprintf("%v", vcs[i].Metadata.Id)
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
				i = rand.Intn(int(config.ExpectedNumberOfTotalVCs))
			}
		}
		indexes, amount := issuer1.RevokeVCInBatches(config, revokedVCsInBatch)
		affectedIndexes = affectedIndexes.Union(indexes)
		amountPaid = amountPaid + amount;
		amountPaid = amountPaid / 2;
	}

	issuer1.RevocationService.CacheRevocationDataStructuresFromSmartContract()

	var falsePositiveStatus bool
	falsePositiveStatus = false
	var isAffectedInMTAcc bool
	isAffectedInMTAcc = false
	numberOfOccuredFalsePositives := 0
	numberOfVCsRetrievedWitnessFromIssuer := 0
	falsePositiveResults := mapset.NewSet()
	fetchedWitnessesFromIssuers := mapset.NewSet()
	for _, credential := range vcs {

		vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
		vcId := fmt.Sprintf("%v",credential.Metadata.Id)

		falsePositiveStatus, isAffectedInMTAcc = issuer1.VerifyTest(vcId, *vp)
		if falsePositiveStatus == true {
			numberOfOccuredFalsePositives++
			falsePositiveResults.Add(vcId)
			if isAffectedInMTAcc == true {
				numberOfVCsRetrievedWitnessFromIssuer++
				fetchedWitnessesFromIssuers.Add(vcId)
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
	zap.S().Infoln("SIMULATOR - \t indexes of VCs that are affected by revocation: ", affectedIndexes)
	zap.S().Infoln("SIMULATOR - \t indexes of VCs that are affected by false positives: ", falsePositiveResults)
	zap.S().Infoln("SIMULATOR - \t indexes of VCs that retrieved witnesses from issuer: ", fetchedWitnessesFromIssuers)
	size, k := BloomFilterConfigurationGenerators(config.ExpectedNumberofRevokedVCs,config.FalsePositiveRate)
	result := &Results{
		TotalVCs:                              int(config.ExpectedNumberOfTotalVCs),
		RevokedVCs:                            int(config.ExpectedNumberofRevokedVCs),
		FalsePositiveRate:                     config.FalsePositiveRate,
		MTHeight:                               int(config.MTHeight),
		MtLevelInDLT:                          int(config.MtLevelInDLT),
		NumberOfFalsePositives:                numberOfOccuredFalsePositives,
		AmountPaid:                            amountPaid,
		NumberOfWitnessUpdatesForMT:           affectedIndexes.Cardinality(),
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
	jsonFile, err := os.Open("results.json")
	if err != nil {
		zap.S().Errorln("ERROR - results.json file open error")
	}
	byte, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byte, &results)
	results = append(results, result)
	jsonRes, _ := json.MarshalIndent(results,"","")
	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	err = ioutil.WriteFile("results.json", jsonRes, 0644)
	if err != nil {
		zap.S().Errorln("unable to write results to file")
	}

}

func BloomFilterConfigurationGenerators(totalNumberOfVCs uint, falsePositiveRate float64) (uint, uint) {
	size, numberOfIndexesPerEntry := bloom.EstimateParameters(totalNumberOfVCs, falsePositiveRate)
	return size, numberOfIndexesPerEntry
}




