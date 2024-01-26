package issuer

import (
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
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
	issuer.credentialStore = []verifiable.Credential{}
	issuer.revokedVcIDs = []string{}
	issuer.revocationProofs = make(map[string]*RevocationData)
	issuer.AffectedVCIndexes = make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	issuer.vcCounter = rand.Intn(100000)
	rs := CreateRevocationServiceStub(config)
	issuer.setRevocationService(rs)
	zap.S().Infoln("ISSUER-","issuer test instance created")
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	return &issuer
}
func TestIssuer(config config.Config){

	issuer := CreateTestIssuer(config)
	vcs := issuer.GenerateDummyVCs(int(config.ExpectedNumberOfTotalVCs))

	for _, vc := range vcs{
		issuer.Issue(*vc)
	}

	issuer.RevocationService.PrintMerkleTree()
	for _, vc := range vcs{
		issuer.VerifyTest(*vc)
	}

	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	revokedVCs := make([]string, totalRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for i, counter:=0, 0; counter< totalRevokedVCs; {


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
				issuer.Revoke(config, *vcs[i])
				revokedVCs = append(revokedVCs, vcID)
				counter++
				break
			}
			rand.Seed(time.Now().UnixNano())
			i = rand.Intn(int(config.ExpectedNumberOfTotalVCs))
		}
	}

	for _, vc := range vcs{
		issuer.UpdateMerkleProof(*vc)
	}



	for _, vc := range vcs{
		issuer.VerifyTest(*vc)
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
	issuer.VerifyTest(*vcs[2])
}
