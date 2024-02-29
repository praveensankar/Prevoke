package signature

import (
	"fmt"
	"github.com/suutaku/go-bbs/pkg/bbs"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	"math/rand"
	"time"
)



type BBS struct {
	PublicKey *bbs.PublicKey
	PrivateKey *bbs.PrivateKey
}

func GenerateKeyPair() *BBS {
	rand.Seed(time.Now().UnixNano())
	seed := make([]byte, 32)
	_, err := rand.Read(seed)
	if err != nil {
		zap.S().Infoln("BBS - error while generating random string: %s", err)
	}
	//zap.S().Infoln("BBS - key pair generation: seed", seed)
	publicKey, privateKey, err  := bbs.GenerateKeyPair(sha3.New512, seed)
	if err != nil {
		zap.S().Infoln("BBS - error creating new key pair: ", err)
	}
	bbs1 := BBS{
		PublicKey:    publicKey,
		PrivateKey: privateKey,
	}
	return &bbs1
}

func Sign(privateKey *bbs.PrivateKey, messages [][]byte) []byte{
	bbsInstance := bbs.NewBbs()
	signature, err := bbsInstance.SignWithKey(messages, privateKey)
	if err != nil {
		zap.S().Infoln("BBS - error signing: ", err)
	}
	return signature
}

func Verify(publicKey []byte, signature []byte, messages [][]byte) bool{

	bbsInstance := &bbs.Bbs{}
	err := bbsInstance.Verify(messages, signature, publicKey)
	if err != nil {
		zap.S().Infoln("BBS - verification failed: ",err)
		return false
	}

	//zap.S().Infoln("BBS - digital signature verification successful")
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
func SelectiveDisclosure(publicKey []byte, signature []byte, messages [][]byte, revealedIndexes []int) ([]byte, []byte ){

	rand.Seed(time.Now().UnixNano())
	nonce := make([]byte, 32)
	_, err := rand.Read(nonce)
	if err != nil {
		zap.S().Infoln("BBS - error while generating random string: %s", err)
	}
	bbsInstance := bbs.NewBbs()
	//pk , _ := bbs.UnmarshalPublicKey(publicKey)
	//zap.S().Infoln("BBS - Selective disclosure - public key: ", pk)
	proof, err := bbsInstance.DeriveProof(messages, signature, nonce, publicKey, revealedIndexes)
	if err!=nil{
		zap.S().Infoln("BBS - error creating proof for selective disclosure: ",err)
	}
	return proof, nonce
}

func VerifySelectiveDisclosureProof( publicKey []byte,  proof []byte, selectiveMessages [][]byte, nonce []byte) bool{

	bbsInstance := bbs.NewBbs()

	err := bbsInstance.VerifyProof(selectiveMessages, proof, nonce, publicKey)
	if err != nil {
		zap.S().Infoln("BBS - selective disclosure verification failed: ",err)
		return false
	}

	zap.S().Infoln("BBS - selective disclosure verification successful")
	return true
}

/*
PublicKeyToString function converts public key to string format

Input:
	public key - *keyset.Handle

Output:
	public key - string
*/
func PublicKeyToString(publicKey *bbs.PublicKey) string{
	res, err := publicKey.Marshal()
	if err!=nil{
		zap.S().Infoln("BBS - error marshing public key")
	}

	str := fmt.Sprintf("%s", res)
	//zap.S().Infoln("BBS - public key byte: ",res)
	return str
}



