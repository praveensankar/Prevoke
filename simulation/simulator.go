package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
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
	"sync/atomic"
	"time"
)



func Start(config config.Config){
		DeployContract(&config, 0)
		//zap.S().Infoln("smart contract: ",config.SmartContractAddress)
		PerformExperiment(config)

}

func StartExperiments(config config.Config){

	experiments := config.ExpParamters
	counter := 0
	for _, exp := range experiments{
		DeployContract(&config, counter)
		counter++
		if counter==len(config.PrivateKeys){
			counter=0
		}
		//zap.S().Infoln("smart contract: ",config.SmartContractAddress)
		SetUpExpParamters(&config, *exp)
		PerformExperiment(config)
	}
}

func DeployContract(conf *config.Config,counter int){
	address, err := blockchain.DeployContract(*conf, counter)

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
	start := time.Now()

	issuer1 := issuer.CreateIssuer(config)
	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
	totalVCs := int(config.ExpectedNumberOfTotalVCs)+remainingSpace

	claimsSet := issuer1.GenerateMultipleDummyVCClaims(totalVCs)
	results := CreateResult()
	vcs := SimulateIssuance(config, issuer1, claimsSet,totalVCs )
	SimulateRevocation(config, issuer1, vcs, results)
	SimulateVerification( issuer1, vcs, results)
	ConstructResults(config, start, results)
	WriteToFile(*results)
}


func SimulateIssuance(config config.Config, issuer1 *issuer.Issuer, claimsSet []interface{},totalVCs int) []models.VerifiableCredential{
	issuer1.IssueBulk(claimsSet, totalVCs)

	credentials := issuer1.CredentialStore
	for _, vc := range credentials{
		issuer1.UpdateMerkleProof(vc)
	}

	vcs:= []models.VerifiableCredential{}

	for i:=0; i<int(config.ExpectedNumberOfTotalVCs);i++{
		vcs = append(vcs, credentials[i])
	}

	return vcs
}

func SimulateRevocation(config config.Config, issuer1 *issuer.Issuer, vcs []models.VerifiableCredential, result *Results){
	revocationBatchSize := int(config.RevocationBatchSize)
	var amountPaid int64
	amountPaid = 0

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
		result.AffectedIndexes = result.AffectedIndexes.Union(indexes)
		amountPaid = amountPaid + amount;
		amountPaid = amountPaid / 2;
	}

	issuer1.RevocationService.CacheRevocationDataStructuresFromSmartContract()

	result.AmountPaid = amountPaid
	result.NumberOfWitnessUpdatesForMT = result.AffectedIndexes.Cardinality()
}

func SimulateVerification( issuer1 *issuer.Issuer, vcs []models.VerifiableCredential, result *Results){
	publicKey, err := issuer1.BbsKeyPair[0].PublicKey.Marshal()

	if err!=nil{
		zap.S().Infoln("SIMULATION - error parsing public key")
	}
	var falsePositiveStatus bool
	falsePositiveStatus = false
	var isAffectedInMTAcc bool
	isAffectedInMTAcc = false
	numberOfOccuredFalsePositives := 0
	numberOfVCsRetrievedWitnessFromIssuer := 0

	var fp atomic.Uint64
	var witFromIssuers atomic.Uint64
	//mux := &sync.RWMutex{}
	vps := make(map[string]*models.VerifiablePresentation)
	for _, credential := range vcs {
		vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
		vcId := fmt.Sprintf("%v",credential.Metadata.Id)
		vps[vcId] = vp
	}
	for vcId, vp := range vps {

			falsePositiveStatus, isAffectedInMTAcc = issuer1.VerifyTest(vcId, *vp)
			if falsePositiveStatus == true {
				fp.Add(1)
				result.FalsePositiveResults.Add(vcId)
				if isAffectedInMTAcc == true {
					witFromIssuers.Add(1)
					result.FetchedWitnessesFromIssuers.Add(vcId)
				}
			}



	}


	numberOfOccuredFalsePositives = int(fp.Load())
	numberOfVCsRetrievedWitnessFromIssuer = int(witFromIssuers.Load())
	result.NumberOfFalsePositives = numberOfOccuredFalsePositives
	result.NumberOfVCsRetrievedWitnessFromIssuer = numberOfVCsRetrievedWitnessFromIssuer
	result.NumberOfWitnessUpdatesSaved = numberOfOccuredFalsePositives-numberOfVCsRetrievedWitnessFromIssuer
}

func ConstructResults(config config.Config, start  time.Time, result *Results){
	zap.S().Infoln("SIMULATOR - \t indexes of VCs that are affected by revocation: ", result.AffectedIndexes)
	zap.S().Infoln("SIMULATOR - \t indexes of VCs that are affected by false positives: ", result.FalsePositiveResults)
	zap.S().Infoln("SIMULATOR - \t indexes of VCs that retrieved witnesses from issuer: ", result.FetchedWitnessesFromIssuers)
	size, k := BloomFilterConfigurationGenerators(config.ExpectedNumberofRevokedVCs,config.FalsePositiveRate)
	// Code to measure
	end := time.Since(start)
	zap.S().Infof("SIMULATOR : \t total time to run the experiment: %f", end.Seconds())

	result.SimulationTime = end.Seconds()
	result.TotalVCs = int(config.ExpectedNumberOfTotalVCs)
	result.RevokedVCs =  int(config.ExpectedNumberofRevokedVCs)
	result.FalsePositiveRate = config.FalsePositiveRate
	result.MTHeight = int(config.MTHeight)
	result.MtLevelInDLT = int(config.MtLevelInDLT)
	result.BloomFilterSize = int(size)
	result.BloomFilterIndexesPerEntry = int(k)

	zap.S().Infoln("SIMULATOR : \t results: ", result.String())
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




