package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/praveensankar/Revocation-Service/Results"
	"github.com/praveensankar/Revocation-Service/blockchain"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/entities"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"math/rand"
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

	if err != nil {
		zap.S().Errorln("ERROR - config.json file open error")
	}
	jsonRes, _ := json.MarshalIndent(address,"","")
	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	err = ioutil.WriteFile("contractAddress.json", jsonRes, 0644)
	if err != nil {
		zap.S().Errorln("unable to write results to file")
	}

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

	issuer1 := entities.CreateIssuer(config)
	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
	totalVCs := int(config.ExpectedNumberOfTotalVCs)+remainingSpace

	claimsSet := issuer1.GenerateMultipleDummyVCClaims(totalVCs)
	results := Results.CreateResult()
	vcs := SimulateIssuance(config, issuer1, claimsSet,totalVCs )
	SimulateRevocation(config, issuer1, vcs, results)
	SimulateVerification( issuer1, vcs, results)
	Results.ConstructResults(config, start, results)
	Results.WriteToFile("results.json", *results)
}


func SimulateIssuance(config config.Config, issuer1 *entities.Issuer, claimsSet []interface{},totalVCs int) []models.VerifiableCredential{
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

func SimulateRevocation(config config.Config, issuer1 *entities.Issuer, vcs []models.VerifiableCredential, result *Results.Results){
	revocationBatchSize := int(config.RevocationBatchSize)
	var amountPaid int64
	amountPaid = 0
	revocationTimePerBatch := 0.0
	revocationTimeTotal := 0.0
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
		indexes, amount, revocationTime := issuer1.RevokeVCInBatches(config, revokedVCsInBatch)
		result.AffectedIndexes = result.AffectedIndexes.Union(indexes)
		revocationTimeTotal = revocationTimeTotal + revocationTime.Seconds()
		if revocationTimePerBatch == 0.0{
			revocationTimePerBatch = revocationTimePerBatch + revocationTime.Seconds();
		} else{
			revocationTimePerBatch = (revocationTimePerBatch + revocationTime.Seconds())/2;
		}

		if amountPaid==0{
			amountPaid = amountPaid + amount;
		} else{
			amountPaid = amountPaid + amount;
			amountPaid = amountPaid / 2;
		}


	}


	result.AmountPaid = amountPaid
	result.RevocationBatchSize = revocationBatchSize
	result.RevocationTimeperBatch = revocationTimePerBatch
	result.RevocationTimeTotal = revocationTimeTotal
}

func SimulateVerification( issuer1 *entities.Issuer, vcs []models.VerifiableCredential, result *Results.Results){
	publicKey, err := issuer1.BbsKeyPair[0].PublicKey.Marshal()

	result.MerkleTreeSizeInDLT = int(issuer1.FetchMerkleTreeSizeInDLT())*8
	result.MerkleTreeSizeTotal = int(issuer1.FetchMerkleTreeSizeLocal())*8

	if err!=nil{
		zap.S().Infoln("SIMULATION - error parsing public key")
	}
	var falsePositiveStatus bool
	falsePositiveStatus = false
	var isAffectedInMTAcc bool
	isAffectedInMTAcc = false
	var vcResult bool
	 phase1time := 0.0
	 phase2time := 0.0
	 bbstime := 0.0
	 avgbbstime := 0.0
	 verificationTimeTotal := 0.0
	 validVCsTime := 0.0
	avgValidVCsTime := 0.0
	 notValidVCsTime := 0.0
	avgNotValidVCsTime := 0.0
	numberOfOccuredFalsePositives := 0
	numberOfVCsRetrievedWitnessFromIssuer := 0

	var fp atomic.Uint64
	var witFromIssuers atomic.Uint64
	//mux := &sync.RWMutex{}
	vps := make(map[string]models.VerifiablePresentation)
	for _, credential := range vcs {
		vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
		vcId := fmt.Sprintf("%v",credential.Metadata.Id)
		vps[vcId] = vp
	}
	for vcId, vp := range vps {

			falsePositiveStatus, isAffectedInMTAcc, vcResult, bbstime, phase1time, phase2time = issuer1.VerifyTest(vcId, vp)
			if falsePositiveStatus == true {
				fp.Add(1)
				result.FalsePositiveResults.Add(vcId)
				if isAffectedInMTAcc == true {
					witFromIssuers.Add(1)
					result.FetchedWitnessesFromIssuers.Add(vcId)
				}
			}

			verificationTimeTotal = verificationTimeTotal + phase1time + phase2time + bbstime

			// valid vcs
			if phase2time ==0 && vcResult==true{
				validVCsTime = validVCsTime + phase1time + bbstime
				if avgValidVCsTime == 0.0{
					avgValidVCsTime = avgValidVCsTime + phase1time
				} else{
					avgValidVCsTime = avgValidVCsTime + phase1time
					avgValidVCsTime = avgValidVCsTime/2
				}
			} else {
				notValidVCsTime = notValidVCsTime + phase1time + phase2time + bbstime
				if avgNotValidVCsTime == 0.0 {
					avgNotValidVCsTime = avgNotValidVCsTime + phase1time + phase2time
				} else {
					avgNotValidVCsTime = avgNotValidVCsTime + phase1time + phase2time
					avgNotValidVCsTime = avgNotValidVCsTime / 2
				}
			}
			if avgbbstime==0.0{
				avgbbstime = avgbbstime + bbstime
			} else {
				avgbbstime = avgbbstime + bbstime
				avgbbstime = avgbbstime / 2
			}



	}


	numberOfOccuredFalsePositives = int(fp.Load())
	numberOfVCsRetrievedWitnessFromIssuer = int(witFromIssuers.Load())
	result.NumberOfFalsePositives = numberOfOccuredFalsePositives
	result.NumberOfVCsRetrievedWitnessFromIssuer = numberOfVCsRetrievedWitnessFromIssuer
	result.VerificationTimeTotalValidVCs = validVCsTime
	result.VerificationTimeTotalRevokedorFalsePositiveVCs = notValidVCsTime
	result.VerificationTimePerValidVC = avgValidVCsTime
	result.VerificationTimePerRevokedorFalsePositiveVC = avgNotValidVCsTime
	result.VerificationTimeTotal = verificationTimeTotal
	result.BBSVerificationTimePerVP = avgbbstime
}












