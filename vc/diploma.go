package vc

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/google/tink/go/keyset"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/signature"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"log"
	"strconv"

	"math/rand"
	"time"
)



type DiplomaClaim struct{
	Id          models.URI
	StudentName string
	StudentId   models.URI
	University  string
	Degree string
	GraduationYear int
	Grade string
}



func EncodeToBytes(claims DiplomaClaim) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(claims)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func DecodeToClaims(s []byte) DiplomaClaim {

	d := DiplomaClaim{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&d)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func (d DiplomaClaim) String() string  {

	var response string
	response = response + "student name : "+ d.StudentName + "\n"
	response = response + "student id : "+ fmt.Sprintf("%v", d.StudentId) + "\n"
	response = response + "university : "+ d.University +"\n"
	response = response + "degree : "+ d.Degree+"\n"
	response = response + "graduation year : "+ fmt.Sprintf("%v", d.GraduationYear)+"\n"
	response = response + "grade : "+ d.Grade

	return response
}


type SampleDiplomaPresentation struct{
	Degree string
	Grade string
	BfIndexes []string
	MtLeafHash string
	Proof []byte
	Nonce []byte
}

func (vp SampleDiplomaPresentation) Present()  {

}

/*
The order of messages for selective disclosure
bf index1, ..., bf indexn, mtLeafHash, vcId, claims
 */
func generateMessages( vcId string, claims DiplomaClaim,
	bfIndexes []string, mtLeafHash string) [][]byte{
	var messages [][]byte

	// 1) append the bf indexes to the messages
	for _, bfIndex := range  bfIndexes{
		messages = append(messages, []byte(bfIndex))
	}
	// 2) append the mt leaf hash to the messages
	messages = append(messages, []byte(mtLeafHash))

	// 2) append vc id to the messages
	messages = append(messages, []byte(vcId))

	// 4) append the claims. All the claims are appended together. However, it is trivial to implement
	messages = append(messages, []byte(claims.Grade))
	messages = append(messages, []byte(claims.Degree))
	messages = append(messages, []byte(fmt.Sprintf("%v", claims.Id)))
	messages = append(messages, []byte(fmt.Sprintf("%v", claims.StudentId)))

	messages = append(messages, []byte(claims.University))
	messages = append(messages, []byte(claims.StudentName))
	messages = append(messages, []byte(string(claims.GraduationYear)))
	return messages
}


func generateProofForDiploma(privateKey *keyset.Handle, vcId string, claims DiplomaClaim, bfIndexes []string,
	mtLeafHash string) models.Proof{
	messages:= generateMessages(vcId, claims, bfIndexes, mtLeafHash)
	signature  := signature.Sign(privateKey, messages)
	proof := models.Proof{Type: "BBS+", ProofValue: signature}
	return proof
}

func NewDiploma() models.VerifiableCredential {
	diploma := models.VerifiableCredential{}
	diploma.Claims, _ = CreateDiplomaClaims("test")
	return diploma

}

func CreateDiplomaClaims(id string) (DiplomaClaim, error){
	var vcID string
	vcID = "diploma#" + fmt.Sprintf("%s", id)
	myDiplomaClaims := DiplomaClaim{
		Id:             vcID,
		StudentName:    "praveen",
		StudentId:      vcID,
		University:     "University",
		Degree:         "Doctor of Philosophy",
		GraduationYear: 2000,
		Grade:          "A",
	}
	return myDiplomaClaims, nil
}

func CreateDiploma(privateKey *keyset.Handle, vcId string, claims interface{},
	bfIndexes []string, mtLeafHash string) (*models.VerifiableCredential, error){


	diplomaClaims := claims.(DiplomaClaim)
	proof := generateProofForDiploma(privateKey, vcId, diplomaClaims, bfIndexes, mtLeafHash)

	var proofs []models.Proof
	proofs = append(proofs, proof)
	myDiploma := models.VerifiableCredential{
		Metadata: models.Metadata{
			Contexts:       []string{"http:test.com/diploma"},
			Id:             vcId,
			Types:          []models.URI{"Diploma"},
			Issuer:         "university of oslo",
			IssuanceDate:   models.TimeString(time.Now()),
			ExpirationDate: models.TimeString(time.Now().Add(1000*time.Hour)),
			CredentialStatus: models.CredentialStatus{
				Id:     diplomaClaims.Id,
				Method: "2-phase revocation",
				BfIndexes: bfIndexes[:],
				MTLeafValue: mtLeafHash,
			},
		},
		Claims: diplomaClaims,
		Proofs: proofs,
	}

	//zap.S().Infoln("DIPLOMA - new diploma created")

	return &myDiploma, nil
}

/*
GenerateProofToSelectivelyDiscloseBfIndexes function generates proof to selective disclosure bf indexes

Input:
	vc: unique string representing the vc

Output:
	sampleDiplomaPresentation
	error
 */
func GenerateProofForSelectiveDisclosure(publicKey *keyset.Handle, diploma models.VerifiableCredential) (*models.VerifiablePresentation, error){
	var revealedIndexes []int
	claims := diploma.Claims.(DiplomaClaim)

	// add bf indexes
	i:=0
	for ; i < techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER; i++ {
		revealedIndexes = append(revealedIndexes, i)
	}

	// add mt leaf hash
	revealedIndexes = append(revealedIndexes, i)

	// grade index
	i = i + 2
	revealedIndexes = append(revealedIndexes, i)
	grade := claims.Grade

	// degree index
	i = i + 1
	revealedIndexes = append(revealedIndexes, i)
	degree := claims.Degree
	proof := diploma.Proofs[0]
	sign := proof.ProofValue

	bfIndexes := diploma.Metadata.CredentialStatus.BfIndexes
	mtLeafHash := diploma.Metadata.CredentialStatus.MTLeafValue

	messages := generateMessages(fmt.Sprintf("%v", diploma.Metadata.Id), claims, bfIndexes , mtLeafHash)
	SDproof, nonce := signature.SelectiveDisclosure(publicKey, sign, messages, revealedIndexes)

	vp := models.VerifiablePresentation{
		Messages: SampleDiplomaPresentation{
			Degree:     degree,
			Grade:     grade,
			BfIndexes:  bfIndexes,
			MtLeafHash: mtLeafHash,
			Proof:      SDproof,
			Nonce:      nonce,
		},
		Proof:    SDproof,
	}

	return &vp, nil

}

func VerifySelectiveDisclosureDiploma(publicKey *keyset.Handle, vp SampleDiplomaPresentation) bool{


	var messages [][]byte

	// 1) append the bf indexes to the messages
	for _, bfIndex := range  vp.BfIndexes{
		messages = append(messages, []byte(bfIndex))
	}
	// 2) append the mt leaf hash to the messages
	messages = append(messages, []byte(vp.MtLeafHash))


	// 4) append the claims. All the claims are appended together. However, it is trivial to implement
	messages = append(messages, []byte(vp.Grade))
	messages = append(messages, []byte(vp.Degree))
	status := signature.VerifySelectiveDisclosureProof(publicKey, vp.Proof, messages, vp.Nonce)


	if status == true {
		//zap.S().Infoln("DIPLOMA PRESENTATION - verification successful: ")
		return true
	}
	zap.S().Infoln("DIPLOMA PRESENTATION - verification failed")
	return false
}



func TestDiploma(){
	bbs := signature.GenerateKeyPair()
	privateKey := bbs.PrivateKey
	publicKey := bbs.PublicKey

	var bfIndexes []string
	for i := 0; i < techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER; i++ {
		bfIndexes = append(bfIndexes, strconv.Itoa(rand.Int()))
	}

	zap.S().Infoln("bf indexes: ", bfIndexes)
	mtLeafHash := "leaf1Hash"

	vcId := "vc#1"
	claims, _ := CreateDiplomaClaims(vcId)

	myDiploma,_ := CreateDiploma(privateKey, vcId, claims, bfIndexes, mtLeafHash )

	zap.S().Infoln("DIPLOMA - json \t: ", string(myDiploma.Json()))

	vp, _ := GenerateProofForSelectiveDisclosure(publicKey, *myDiploma)

	zap.S().Infoln("DIPLOMA - Presentation: ",vp)

	diplomaPresentation := vp.Messages.(SampleDiplomaPresentation)

	VerifySelectiveDisclosureDiploma(publicKey, diplomaPresentation)

	diplomaPresentation.Grade="C"

	VerifySelectiveDisclosureDiploma(publicKey, diplomaPresentation)

}


