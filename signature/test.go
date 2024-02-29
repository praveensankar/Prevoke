package signature

import (
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
)

func TestBBS(conf config.Config){
	bbsKeys := GenerateKeyPair()
	//bytes, _ := bbsKeys.PublicKey.Marshal()
	publicKey,  _ := bbsKeys.PublicKey.Marshal()
	zap.S().Infoln("BBS TEST - public key: ", publicKey)
	messages := [][]byte{[]byte("message 1"), []byte("message 2"), []byte("message 3"), []byte("message 4")}
	signature := Sign(bbsKeys.PrivateKey, messages)

	Verify(publicKey, signature, messages)
	selectiveMessageIndexes := []int{0,2}
	proof, nonce := SelectiveDisclosure(publicKey, signature, messages, selectiveMessageIndexes)

	selectiveMessages := [][]byte{messages[0], messages[2]}
	VerifySelectiveDisclosureProof(publicKey, proof, selectiveMessages, nonce)

	//wrongSelectiveMessages := [][]byte{messages[0], messages[1]}
	//VerifySelectiveDisclosureProof(publicKey, proof, wrongSelectiveMessages, nonce)

}


func TestBBS1(conf config.Config){
	bbsKeys := GenerateKeyPair1()
	bytes, _ := bbsKeys.PublicKey.Marshal()
	zap.S().Infoln("BBS TEST - public key: ", len(bytes))
	//publicKey := []byte(PublicKeyToString1(bbsKeys.PublicKey))
	messages := [][]byte{[]byte("message 1"), []byte("message 2"), []byte("message 3"), []byte("message 4")}
	signature := Sign1(bbsKeys.PrivateKey, messages)

	Verify1(bbsKeys.PublicKey, signature, messages)


	selectiveMessageIndexes := []int{0,2}
	proof, nonce := SelectiveDisclosure1(bbsKeys.PublicKey, signature, messages, selectiveMessageIndexes)
	selectiveMessages := [][]byte{messages[0], messages[2]}
	wrongSelectiveMessages := [][]byte{messages[0], messages[1]}
	VerifySelectiveDisclosureProof1(bbsKeys.PublicKey, proof, selectiveMessages, nonce)
	VerifySelectiveDisclosureProof1(bbsKeys.PublicKey, proof, wrongSelectiveMessages, nonce)

}


