package issuer

import "github.com/praveensankar/Revocation-Service/config"

func TestIssuer(config config.Config){

	issuer := CreateIssuer(config)

	vc1 := issuer.generateDummyVC()
	issuer.issue(*vc1)
	//issuer.verifyTest(*vc1)
	vc2 := issuer.generateDummyVC()
	issuer.issue(*vc2)
	//issuer.verifyTest(*vc2)
	vc3 := issuer.generateDummyVC()
	issuer.issue(*vc3)
	//issuer.verifyTest(*vc3)
	vc4 := issuer.generateDummyVC()
	issuer.issue(*vc4)
	//issuer.verifyTest(*vc4)
	////for index := range issuer.credentialStore{
	////	fmt.Println(issuer.credentialStore[index])
	////}
	//
	//issuer.getUpdatedMerkleProof(*vc1)
	//
	//
	issuer.verifyTest(*vc1)
	issuer.revoke(*vc1)
	issuer.verifyTest(*vc1)
}
