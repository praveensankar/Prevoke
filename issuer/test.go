package issuer

import "github.com/praveensankar/Revocation-Service/config"

func TestIssuer(config config.Config){

	issuer := CreateIssuer(config)
	rs := CreateRevocationServiceStub(config)
	issuer.setRevocationService(rs)
	vcs := issuer.GenerateDummyVCs(int(config.ExpectedNumberOfTotalVCs))

	for _, vc := range vcs{
		issuer.Issue(*vc)
	}

	issuer.RevocationService.PrintMerkleTree()
	for _, vc := range vcs{
		issuer.VerifyTest(*vc)
	}

	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	for i, counter:=0, 0; counter< totalRevokedVCs; counter++{
		issuer.Revoke(config, *vcs[i])
		i = i + 2
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
