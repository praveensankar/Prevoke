package simulation

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/issuer"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"time"
)

func TestSimulator(conf config.Config) {
	experiments := conf.ExpParamters

	for _, exp := range experiments{

		//zap.S().Infoln("smart contract: ",config.SmartContractAddress)
		SetUpExpParamters(&conf, *exp)
		PerformExperimentTest(conf)
		break
	}
}

func PerformExperimentTest(config config.Config){
	issuer1 := issuer.CreateTestIssuer(config)
	publicKey, _ := issuer1.BbsKeyPair[0].PublicKey.Marshal()
	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
	claimsSet := issuer1.GenerateMultipleDummyVCClaims(int(config.ExpectedNumberOfTotalVCs)+remainingSpace)
	revocationBatchSize :=5

	issuer1.IssueBulk(claimsSet, int(config.ExpectedNumberOfTotalVCs)+remainingSpace)

	credentials := issuer1.CredentialStore
	for _, vc := range credentials{
		issuer1.UpdateMerkleProof(vc)
	}


	vcs:= []models.VerifiableCredential{}

	for i:=0; i<int(config.ExpectedNumberOfTotalVCs);i++{
		vcs = append(vcs, credentials[i])
	}

	for _, credential := range vcs{

		vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
		vcId := fmt.Sprintf("%v",credential.Metadata.Id)
		issuer1.VerifyTest(vcId, *vp)
	}


	var amountPaid int64
	amountPaid = 0
	affectedIndexes := mapset.NewSet()
	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	revokedVCs := make([]string, totalRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for batch:=0; batch<revocationBatchSize; batch++ {
		revokedVCsInBatch := make([]string, 0)
		for i, counter := 0, 0; counter < int(int64(math.Ceil(float64(totalRevokedVCs/revocationBatchSize)))); {

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

	var falsePositiveStatus bool
	falsePositiveStatus = false
	var isAffectedInMTAcc bool
	isAffectedInMTAcc = false
	numberOfOccuredFalsePositives := 0
	numberOfVCsRetrievedWitnessFromIssuer := 0
	for _, credential := range vcs {
		vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey,credential)
		vcId := fmt.Sprintf("%v",credential.Metadata.Id)
		falsePositiveStatus, isAffectedInMTAcc = issuer1.VerifyTest(vcId, *vp)
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
	zap.S().Infoln("SIMULATOR - \t indexes of VCs that are affected by revocation: ", affectedIndexes)
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
