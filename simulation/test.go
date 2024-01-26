package simulation

import (
	"github.com/deckarep/golang-set"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/issuer"
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
	}
}

func PerformExperimentTest(config config.Config){
	issuer1 := issuer.CreateTestIssuer(config)
	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
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
	affectedIndexes := mapset.NewSet()
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
				indexes, amount := issuer1.Revoke(config, *vcs[i])
				affectedIndexes = affectedIndexes.Union(indexes)
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
