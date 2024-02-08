package signature

import (
	"github.com/google/tink/go/keyset"
	"github.com/hyperledger/aries-framework-go/pkg/crypto/tinkcrypto/primitive/bbs"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type BBS struct {

	PublicKey *keyset.Handle
	PrivateKey *keyset.Handle
}

func GenerateKeyPair() *BBS {
	// create signer keyset handle
	kh, err := keyset.NewHandle(bbs.BLS12381G2KeyTemplate())
	if err != nil {
		zap.S().Infoln("BBS - error creating new key pair: ", err)
	}

	// extract signer public keyset handle and key for signature verification and proof derivation/verification
	publicKey, err := kh.Public()
	if err != nil {
		zap.S().Infoln("BBS - error creating public key: ", err)
	}



	bbs := BBS{
		PublicKey:    publicKey,
		PrivateKey: kh,
	}
	zap.S().Infoln("BBS - \t private key: ", kh.String(), "\n public key: ", bbs.PublicKey.KeysetInfo().KeyInfo)
	return &bbs
}

func Sign(privateKey *keyset.Handle, messages [][]byte) []byte{
	// finally get the BBS+ signing primitive from the private key handle created above
	signerHandle, _ := bbs.NewSigner(privateKey)
	signature, err := signerHandle.Sign(messages)
	if err != nil{
		zap.S().Infoln("BBS - error signing: ",err)
	}
	//zap.S().Infoln("BBS - signature: ",signature)
	return signature
}

func Verify(publicKey *keyset.Handle, signature []byte, messages [][]byte) bool{
	verifier, err := bbs.NewVerifier(publicKey)
	if err != nil{
		zap.S().Infoln("BBS - error creating verifier: ",err)
	}
	err = verifier.Verify(messages, signature)
	if err != nil {
		zap.S().Infoln("BBS - verification failed: ",err)
		return false
	}
	//zap.S().Infoln("BBS - verification successful")
	return true
}

/*
SelectiveDisclosure function generates proof for selective disclosure

Input:
	publicKey: public key
	signature : digital signature of the complete messages
	messages: all the messages
	revealedIndexes: list of indexes that need to be revealed

Output:
	(proof) - []byte
	(nonce) - []byte
*/
func SelectiveDisclosure(publicKey *keyset.Handle, signature []byte, messages [][]byte, revealedIndexes []int) ([]byte, []byte ){
	verifier, err := bbs.NewVerifier(publicKey)
	if err != nil{
		zap.S().Infoln("BBS - error creating verifier: ",err)
	}
	rand.Seed(time.Now().UnixNano())

	nonce := make([]byte, 10)
	proof, err := verifier.DeriveProof(messages, signature, nonce, revealedIndexes)
	if err!=nil{
		zap.S().Infoln("BBS - error creating proof for selective disclosure: ",err)
	}
	return proof, nonce
}

func VerifySelectiveDisclosureProof(publicKey *keyset.Handle, proof []byte, selectiveMessages [][]byte, nonce []byte) bool{
	verifier, err := bbs.NewVerifier(publicKey)
	if err != nil{
		zap.S().Infoln("BBS - error creating verifier: ",err)
	}
	err = verifier.VerifyProof(selectiveMessages, proof, nonce)
	if err != nil {
		zap.S().Infoln("BBS - verification failed: ",err)
		return false
	}
	//zap.S().Infoln("BBS - verification successful")
	return true
}