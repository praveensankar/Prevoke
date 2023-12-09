package issuer

import "github.com/praveensankar/Revocation-Service/config"

func TestIssuer(config config.Config){

	issuer := CreateIssuer(config)
	vcs := issuer.generateDummyVCs(int(config.ExpectedNumberOfTotalVCs))

	for _, vc := range vcs{
		issuer.issue(*vc)
	}
	for _, vc := range vcs{
		issuer.verifyTest(*vc)
	}
	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	for i:=0; i< totalRevokedVCs; i++{
		issuer.revoke(*vcs[i])
	}
	for _, vc := range vcs{
		issuer.verifyTest(*vc)
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
	issuer.verifyTest(*vcs[2])
}
