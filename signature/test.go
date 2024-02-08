package signature

import "github.com/praveensankar/Revocation-Service/config"

func TestBBC(conf config.Config){
	bbs := GenerateKeyPair()
	messages := [][]byte{[]byte("message 1"), []byte("message 2"), []byte("message 3"), []byte("message 4")}
	signature := Sign(bbs.PrivateKey, messages)
	Verify(bbs.PublicKey, signature, messages)
	selectiveMessageIndexes := []int{0,2}
	proof, nonce := SelectiveDisclosure(bbs.PublicKey, signature, messages, selectiveMessageIndexes)
	selectiveMessages := [][]byte{messages[0], messages[2]}
	VerifySelectiveDisclosureProof(bbs.PublicKey, proof, selectiveMessages, nonce)
	wrongSelectiveMessages := [][]byte{messages[0], messages[1]}
	VerifySelectiveDisclosureProof(bbs.PublicKey, proof, wrongSelectiveMessages, nonce)
}
