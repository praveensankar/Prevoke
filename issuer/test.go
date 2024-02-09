package issuer

import (
	"fmt"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/revocation_service"
	"github.com/praveensankar/Revocation-Service/signature"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"math/rand"
	"time"
)


/*
creates new issuer instance for testing purpose.
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
	zap.S().Infoln("ISSUER-","issuer test instance created")
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	return &issuer
}


func TestIssuer(config config.Config){

	issuer := CreateTestIssuer(config)
	publicKey, _ := issuer.BbsKeyPair[0].PublicKey.Marshal()
	claimsSet := issuer.GenerateMultipleDummyVCClaims(int(config.ExpectedNumberOfTotalVCs))

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
			issuer.VerifyTest(vcId, *vp)
	}

	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	revokedVCs := make([]string, totalRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for i, counter:=0, 0; counter< totalRevokedVCs; {
		for {
			vcID := fmt.Sprintf("%v", vcs[i].Metadata.Id)
			isalreadyRevoked := false
			for _, revokedId := range revokedVCs {
				if vcID == revokedId {
					isalreadyRevoked = true
					break
				}
			}
			if isalreadyRevoked==false{
				issuer.Revoke(config, credentials[i])
				revokedVCs = append(revokedVCs, vcID)
				counter++
				break
			}
			rand.Seed(time.Now().UnixNano())
			i = rand.Intn(int(config.ExpectedNumberOfTotalVCs))
		}
	}

	for _, credential := range credentials{
		issuer.UpdateMerkleProof(credential)
	}



	for _, credential := range credentials{
		vp, _ := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
		vcId := fmt.Sprintf("%v",credential.Metadata.Id)
		issuer.VerifyTest(vcId, *vp)
	}
	//vc1 := issuer.generateDummyVC()
	//issuer.issue(*vc1)
	////issuer.verifyTest(*vc1)
	//vc2 := issuer.generateDummyVC()
	//issuer.issue(*vc2)
	//issuer.verifyTest(*vc2)
	//vc3 := issuer.generateDummyVC()
	//issuer.issue(*vc3)
	////issuer.verifyTest(*vc3)
	//vc4 := issuer.generateDummyVC()
	//issuer.issue(*vc4)
	//issuer.verifyTest(*vc4)
	////for index := range issuer.credentialStore{
	////	fmt.Println(issuer.credentialStore[index])
	////}
	//
	//issuer.getUpdatedMerkleProof(*vc1)
	//
	//
	//issuer.verifyTest(*vcs[2])
	//issuer.revoke(*vcs[2])

}
