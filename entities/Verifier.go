package entities

import (
	"github.com/ethereum/go-ethereum/ethclient"
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
}

/*
CreateVerifier creates a new verifier
 */
func  CreateVerifier(config config.Config) *Verifier{

	verifier := Verifier{}
	verifier.name = config.IssuerName
	verifier.HolderAddress = config.HolderAddress
	rs := revocation_service.CreateRevocationService(config)
	verifier.RevocationService = rs
	verifier.bbs = bbs.NewBbs()
	zap.S().Infoln("VERIFIER-","new entities created: entities name - ",verifier.name)
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	return &verifier
}

/*
VerifyVP verifies a VP

Input:
vp - models.VerifiablePresentation

Outputs:
1) (bool) - actual status of the VP
2) (float64) - bbs verification time
3) (float64) - phase 1 time (only valid vcs)
4) (float64) - phase 2 time (both revoked and false positive vcs)
 */


func (verifier *Verifier) VerifyVP(vp *models.VerifiablePresentation) (bool, float64, float64, float64) {

	//zap.S().Infoln("\n********************************************************************************************************************************")
	//zap.S().Infoln("***********************\t  Verification test: \t VC id: ", vc.ID, "***********************")
	var bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int



	var bbsTime time.Duration

	// ***************************** Phase 1 **************************************************




	publicKeys := verifier.RevocationService.FetchPublicKeysCached()
	publicKey := publicKeys[0]
	pk , _ := bbs.UnmarshalPublicKey(publicKey)
	zap.S().Infoln("HOLDER - issuer's public keys: ", pk.PointG2)



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
		return  false, 0, 0, 0
	}

	phase1Time := time.Since(phase1Start)
	if phase1Result == true{
		zap.S().Infoln("VERIFIER: \t ***VERIFICATION*** vp: \t phase1 result: ", phase1Result)
		return  phase1Result, bbsTime.Seconds(), phase1Time.Seconds(), 0.0
	}


	// ***************************** update witness only for valid vcs ***********************************

	//zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)

	if  phase1Result==false{

	}
	return false, 0.0, 0.0, 0.0

}

