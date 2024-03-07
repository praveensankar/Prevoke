package entities

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/praveensankar/Revocation-Service/common"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/revocation_service"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/praveensankar/Revocation-Service/vc"
	"github.com/suutaku/go-bbs/pkg/bbs"
	"go.uber.org/zap"
	"math/big"
	"net"
	"strconv"
	"sync"
	"time"
)

type IVerifier interface {
	VerifyVP(vp models.VerifiablePresentation)  (bool, float64, float64, float64)
}


type Verifier struct {
	sync.RWMutex
	name string
	Type Entity
	Conn net.Conn
	HolderAddress string
	bbs *bbs.Bbs
	blockchainEndPoint *ethclient.Client
	RevocationService revocation_service.IRevocationService
	Result *common.Results
}

/*
CreateVerifier creates a new verifier
 */
func  CreateVerifier(config config.Config) *Verifier{

	verifier := Verifier{}
	verifier.name = config.IssuerName
	verifier.HolderAddress = config.HolderAddress
	contractAddress := verifier.getContractAddressFromIssuer(config.IssuerAddress)
	config.SmartContractAddress=contractAddress
	rs := revocation_service.CreateRevocationService(config)
	verifier.RevocationService = rs
	verifier.bbs = bbs.NewBbs()
	verifier.Result= common.CreateResult()
	zap.S().Infoln("VERIFIER-","new entities created: entities name - ",verifier.name)
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	return &verifier
}

/*
VerifyVP verifies a VP - only phase 1

Input:
vp - models.VerifiablePresentation

Outputs:
1) (bool) - actual status of the VP
2) (float64) - bbs verification time
3) (float64) - phase 1 time (only valid vcs). It doesnot include bbs verification time
 */


func (verifier *Verifier) VerifyVPPhase1(vp *models.VerifiablePresentation) (bool, float64, float64) {

	//zap.S().Infoln("\n********************************************************************************************************************************")
	//zap.S().Infoln("***********************\t  Verification test: \t VC id: ", vc.ID, "***********************")
	var bfIndexes []*big.Int
	var bbsTime time.Duration

	// ***************************** Phase 1 **************************************************



	publicKeys := verifier.RevocationService.FetchPublicKeysCached()
	publicKey := publicKeys[0]
	//pk , _ := bbs.UnmarshalPublicKey(publicKey)
	//zap.S().Infoln("HOLDER - issuer's public keys: ", pk.PointG2)

	//verify selective disclosure

	diplomaPresentation := vp.Messages.(vc.SampleDiplomaPresentation)

	bbsVerificationStart := time.Now()
	vc.VerifySelectiveDisclosureDiploma(publicKey, diplomaPresentation)
	bbsTime = time.Since(bbsVerificationStart)
	for i, v:= range diplomaPresentation.BfIndexes{
		intValue, _ := strconv.Atoi(v)
		bfIndexes[i]=big.NewInt(int64(intValue))
	}

	phase1Start := time.Now()
	phase1Result, err := verifier.RevocationService.VerificationPhase1(bfIndexes[:])
	if err != nil {
		return  false, 0, 0
	}

	phase1Time := time.Since(phase1Start)
	zap.S().Infoln("VERIFIER: \t ***VERIFICATION*** vp: \t phase1 result: ", phase1Result)
	return  phase1Result, bbsTime.Seconds(), phase1Time.Seconds()



}

/*
VerifyVP verifies a VP - only phase 2

Input:
vp - models.VerifiablePresentation

Outputs:
1) (bool) - actual status of the VP (phase 2 result)

*/
func (verifier *Verifier) VerifyVPPhase2(vp *models.VerifiablePresentation, proof techniques.MerkleProof) (bool) {

	// ***************************** update witness only for valid vcs ***********************************

	//zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)
	diplomaPresentation := vp.Messages.(vc.SampleDiplomaPresentation)
	phase2Result, err := verifier.RevocationService.VerificationPhase2(diplomaPresentation.MtLeafHash, proof.OrderedWitnesses)
	if err!=nil{
		zap.S().Infoln("VERIFIER - phase 2 verification failed: ", err)
	}
	zap.S().Infoln("VERIFIER: \t ***VERIFICATION*** vp: \t phase2 result: ", phase2Result)
	return phase2Result
}

func (verifier *Verifier) Reset(config config.Config){
	contractAddress := verifier.getContractAddressFromIssuer(config.IssuerAddress)
	config.SmartContractAddress=contractAddress
	rs := revocation_service.CreateRevocationService(config)
	verifier.RevocationService = rs
	verifier.bbs = bbs.NewBbs()
	verifier.Result= common.CreateResult()
}


func (verifier *Verifier) SetExperimentConfigs(conf *config.Config, exp config.Experiment){
	conf.ExpectedNumberOfTotalVCs = uint(exp.TotalVCs)
	conf.ExpectedNumberofRevokedVCs = uint(exp.RevokedVCs)
	conf.FalsePositiveRate = exp.FalsePositiveRate
	conf.MTHeight = uint(exp.MtHeight)
	conf.MtLevelInDLT = uint(exp.MtLevelInDLT)
	conf.RevocationBatchSize = uint(exp.RevocationBatchSize)
	zap.S().Infoln("VERIFIER - updated config with experiment config: ", exp.String())

}