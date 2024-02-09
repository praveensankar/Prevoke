package signature

import (
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
)

func TestBBS(conf config.Config){
	bbs := GenerateKeyPair()
	bytes, _ := bbs.PublicKey.Marshal()
	zap.S().Infoln("BBS TEST - public key: ", len(bytes))
	publicKey := []byte(PublicKeyToString(bbs.PublicKey))
	messages := [][]byte{[]byte("message 1"), []byte("message 2"), []byte("message 3"), []byte("message 4")}
	signature := Sign(bbs.PrivateKey, messages)
	Verify(publicKey, signature, messages)
	selectiveMessageIndexes := []int{0,2}
	proof, nonce := SelectiveDisclosure(publicKey, signature, messages, selectiveMessageIndexes)
	selectiveMessages := [][]byte{messages[0], messages[2]}
	VerifySelectiveDisclosureProof(publicKey, proof, selectiveMessages, nonce)
	wrongSelectiveMessages := [][]byte{messages[0], messages[1]}
	VerifySelectiveDisclosureProof(publicKey, proof, wrongSelectiveMessages, nonce)
}
