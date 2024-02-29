package entities

import (
	"fmt"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/revocation_service"
	"github.com/praveensankar/Revocation-Service/signature"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"net"
	"time"
)


/*
creates new entities instance for testing purpose.
This instance connects to local stub for revoation service instead of connecting to blockchain
*/
func  CreateTestIssuer(config config.Config) *Issuer {
	issuer := Issuer{}
	issuer.name = config.IssuerName
	issuer.CredentialStore = []models.VerifiableCredential{}
	issuer.revokedVcIDs = []string{}
	issuer.revocationProofs = make(map[string]*revocation_service.RevocationData)
	issuer.AffectedVCIndexes = make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	issuer.vcCounter = rand.Intn(100000)
	issuer.credentialType="diploma"
	rs := revocation_service.CreateRevocationServiceStub(config)
	issuer.setRevocationService(rs)
	keyPair1 := signature.GenerateKeyPair()
	issuer.BbsKeyPair = make([]*signature.BBS, 2)
	issuer.BbsKeyPair[0] = keyPair1
	publicKey1, _ := keyPair1.PublicKey.Marshal()
	keys := make([][]byte, 1)
	keys[0]=publicKey1
	rs.AddPublicKeys(keys)
	issuer.activeConnections = []net.Conn{}
	issuer.processedConnections = []net.Conn{}
	zap.S().Infoln("ISSUER-","entities test instance created")
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	return &issuer
}


func TestIssuer(config config.Config){

	issuer := CreateTestIssuer(config)
	publicKey, _ := issuer.BbsKeyPair[0].PublicKey.Marshal()
	claimsSet := issuer.GenerateMultipleDummyVCClaims(int(config.ExpectedNumberOfTotalVCs))
	revocationBatchSize :=5
	for _, claims := range claimsSet{
		issuer.Issue(claims)
	}


	credentials := issuer.CredentialStore
	var vcs []models.VerifiableCredential

	for _, credential := range credentials{
		vcs = append(vcs, credential)
		issuer.UpdateMerkleProof(credential)
	}
	issuer.RevocationService.PrintMerkleTree()
	for _, credential := range vcs{
			vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
			vcId := fmt.Sprintf("%v",credential.Metadata.Id)
			issuer.VerifyTest(vcId, vp)
	}

	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)

	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	revokedVCs := make([]string, 0)
	for batch:=0; batch<revocationBatchSize; batch++ {
		revokedVCsInBatch := make([]string, 0)
		for i, counter := 0, 0; counter < int(int64(math.Ceil(float64(totalRevokedVCs/revocationBatchSize)))); {
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
					revokedVCs = append(revokedVCs, vcID)
					revokedVCsInBatch = append(revokedVCsInBatch, vcID)
					counter++
					break
				}
				rand.Seed(time.Now().UnixNano())
				i = rand.Intn(int(config.ExpectedNumberOfTotalVCs))
			}
		}
		//zap.S().Infoln("ISSUER TEST - revoked vcIDs: ", revokedVCsInBatch)
		issuer.RevokeVCInBatches(config, revokedVCsInBatch)
	}
	for _, credential := range credentials{
		issuer.UpdateMerkleProof(credential)
	}



	for _, credential := range credentials{
		vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
		vcId := fmt.Sprintf("%v",credential.Metadata.Id)
		issuer.VerifyTest(vcId, vp)
	}
	//vc1 := entities.generateDummyVC()
	//entities.issue(*vc1)
	////entities.verifyTest(*vc1)
	//vc2 := entities.generateDummyVC()
	//entities.issue(*vc2)
	//entities.verifyTest(*vc2)
	//vc3 := entities.generateDummyVC()
	//entities.issue(*vc3)
	////entities.verifyTest(*vc3)
	//vc4 := entities.generateDummyVC()
	//entities.issue(*vc4)
	//entities.verifyTest(*vc4)
	////for index := range entities.credentialStore{
	////	fmt.Println(entities.credentialStore[index])
	////}
	//
	//entities.getUpdatedMerkleProof(*vc1)
	//
	//
	//entities.verifyTest(*vcs[2])
	//entities.revoke(*vcs[2])

}
