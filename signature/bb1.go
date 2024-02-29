


//The package url: https://pkg.go.dev/github.com/hyperledger/aries-framework-go@v0.3.2/pkg/crypto/tinkcrypto/primitive/bbs

package signature

import (
	"crypto/sha256"
	_ "github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/tinkcrypto"
	"github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/tinkcrypto/primitive/bbs/subtle"
	_ "github.com/hyperledger/aries-framework-go/component/models/signature/verifier"
	"github.com/hyperledger/aries-framework-go/pkg/crypto/primitive/bbs12381g2pub"
	_ "github.com/hyperledger/aries-framework-go/pkg/crypto/tinkcrypto/primitive/bbs"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

//type BBS struct {
//	PublicKey *keyset.Handle
//	PrivateKey *keyset.Handle
//}

type BBS1 struct {
	PublicKey *bbs12381g2pub.PublicKey
	PrivateKey *bbs12381g2pub.PrivateKey
}

func GenerateKeyPair1() *BBS1 {
	// create signer keyset handle
	//kh, err := keyset.NewHandle(bbs.BLS12381G2KeyTemplate())
	//
	//if err != nil {
	//	zap.S().Infoln("BBS - error creating new key pair: ", err)
	//}
	//
	//// extract signer public keyset handle and key for signature verification and proof derivation/verification
	//publicKey, err := kh.Public()
	//if err != nil {
	//	zap.S().Infoln("BBS - error creating public key: ", err)
	//}
	//
	//verifier, err := bbs.NewVerifier(publicKey)
	//zap.S().Infoln("BBS - verifier: ", verifier)
	//bbs := BBS{
	//	PublicKey:    publicKey,
	//	PrivateKey: kh,
	//}
	//zap.S().Infoln("BBS - \t private key: ", kh.String(), "\n public key: ", bbs.PublicKey.KeysetInfo().KeyInfo)


	publicKey, privateKey, err := bbs12381g2pub.GenerateKeyPair(sha256.New, []byte(""))
	if err != nil {
		zap.S().Infoln("BBS - error creating new key pair: ", err)
	}
	bbs1 := BBS1{
		PublicKey:    publicKey,
		PrivateKey: privateKey,
	}

	publicKeyInByte, _ := publicKey.Marshal()

	verifier := subtle.NewBLS12381G2Verifier(publicKeyInByte)
	zap.S().Infoln("BBS - subtle verifier: ", verifier)
	return &bbs1
}

func Sign1(privateKey *bbs12381g2pub.PrivateKey, messages [][]byte) []byte{
	// finally get the BBS+ signing primitive from the private key handle created above
	//signerHandle, _ := bbs.NewSigner(privateKey)
	privateKeyInByte, _ := privateKey.Marshal()
	signerHandle := subtle.NewBLS12381G2Signer(privateKeyInByte)
	signature, err := signerHandle.Sign(messages)
	if err != nil{
		zap.S().Infoln("BBS - error signing: ",err)
	}
	//zap.S().Infoln("BBS - signature: ",signature)
	return signature
}

func Verify1(publicKey *bbs12381g2pub.PublicKey, signature []byte, messages [][]byte) bool{

	publicKeyInByte, _ := publicKey.Marshal()

	verifier := subtle.NewBLS12381G2Verifier(publicKeyInByte)
	//if err != nil{
	//	zap.S().Infoln("BBS - error creating verifier: ",err)
	//}
	err := verifier.Verify(messages, signature)
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
func SelectiveDisclosure1(publicKey *bbs12381g2pub.PublicKey, signature []byte, messages [][]byte, revealedIndexes []int) ([]byte, []byte ){
	//verifier, err := bbs.NewVerifier(publicKey)
	//if err != nil{
	//	zap.S().Infoln("BBS - error creating verifier: ",err)
	//}
	publicKeyInByte, _ := publicKey.Marshal()

	verifier := subtle.NewBLS12381G2Verifier(publicKeyInByte)
	rand.Seed(time.Now().UnixNano())

	nonce := make([]byte, 10)
	proof, err := verifier.DeriveProof(messages, signature, nonce, revealedIndexes)
	if err!=nil{
		zap.S().Infoln("BBS - error creating proof for selective disclosure: ",err)
	}
	return proof, nonce
}

func VerifySelectiveDisclosureProof1(publicKey *bbs12381g2pub.PublicKey,  proof []byte, selectiveMessages [][]byte, nonce []byte) bool{
	//verifier, err := bbs.NewVerifier(publicKey)
	//if err != nil{
	//	zap.S().Infoln("BBS - error creating verifier: ",err)
	//}
	publicKeyInByte, _ := publicKey.Marshal()

	verifier := subtle.NewBLS12381G2Verifier(publicKeyInByte)
	err := verifier.VerifyProof(selectiveMessages, proof, nonce)
	if err != nil {
		zap.S().Infoln("BBS - verification failed: ",err)
		return false
	}
	//zap.S().Infoln("BBS - verification successful")
	return true
}

/*
PublicKeyToString function converts public key to string format

Input:
	public key - *keyset.Handle

Output:
	public key - string
*/
func PublicKeyToString1(publicKey *bbs12381g2pub.PublicKey) string{
	res, _ := publicKey.Marshal()
	return string(res)
}

/*
PublicKeyToVerifier function converts public key to string format

Input:
	public key - string

Output:
	verifier handld - *subtle.BLS12381G2Verifier
*/
func PublicKeyToVerifier(publicKey string) *subtle.BLS12381G2Verifier{
	verifier := subtle.NewBLS12381G2Verifier([]byte(publicKey))
	zap.S().Infoln("BBS - verifier: ", verifier)
	return verifier
}

