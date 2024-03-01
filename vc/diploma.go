package vc

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	_ "github.com/google/tink/go/keyset"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/signature"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/suutaku/go-bbs/pkg/bbs"
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
	for i:=0;i<techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER;i++ {
		messages = append(messages, []byte(bfIndexes[i]))
		//zap.S().Infoln("BBS - bf index: ", []byte(bfIndexes[i]))
	}
	// 2) append the mt leaf hash to the messages
	messages = append(messages, []byte(mtLeafHash))
	//zap.S().Infoln("BBS - mt leaf hash: ", []byte(mtLeafHash))

	// 2) append vc id to the messages
	messages = append(messages, []byte(vcId))
	//zap.S().Infoln("BBS - vc  id: ", []byte(vcId))

	// 4) append the claims. All the claims are appended together. However, it is trivial to implement
	messages = append(messages, []byte(fmt.Sprintf("%s", claims.Grade)))
	//zap.S().Infoln("BBS - grade: ",[]byte(fmt.Sprintf("%s", claims.Grade)))
	messages = append(messages, []byte(claims.Degree))
	messages = append(messages, []byte(fmt.Sprintf("%v", claims.Id)))
	messages = append(messages, []byte(fmt.Sprintf("%v", claims.StudentId)))

	messages = append(messages, []byte(claims.University))
	messages = append(messages, []byte(claims.StudentName))
	messages = append(messages, []byte(fmt.Sprintf("%d", claims.GraduationYear)))
	return messages
}


func generateProofForDiploma(privateKey *bbs.PrivateKey, vcId string, claims DiplomaClaim, bfIndexes []string,
	mtLeafHash string) models.Proof{
	messages:= generateMessages(vcId, claims, bfIndexes, mtLeafHash)
	signature  := signature.Sign(privateKey, messages)
	proof := models.Proof{Type: "bbs+", ProofValue: signature}
	return proof
}

func verifyProofForDiploma(publicKey []byte, sign []byte, vcId string, claims DiplomaClaim, bfIndexes []string,
	mtLeafHash string) bool{
	messages:= generateMessages(vcId, claims, bfIndexes, mtLeafHash)
	status  := signature.Verify(publicKey, sign, messages)
	zap.S().Infoln("DIPLOMA - digital signature verification: ",status)
	return status
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

func CreateDiploma(privateKey *bbs.PrivateKey, vcId string, claims interface{},
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
func GenerateProofForSelectiveDisclosure(publicKey []byte, diploma models.VerifiableCredential) (models.VerifiablePresentation, error){
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
	//zap.S().Infoln("DIPLOMA - revealed indexes: ", revealedIndexes)
	messages := generateMessages(fmt.Sprintf("%v", diploma.Metadata.Id), claims, bfIndexes , mtLeafHash)
	//zap.S().Infoln("DIPLOMA - messages: \t ", messages)

	signStatus := signature.Verify(publicKey, sign, messages)
	if signStatus==false{
		zap.S().Infoln("BBS - digital signature verification failed")
	}
	SDproof, nonce := signature.SelectiveDisclosure(publicKey, sign, messages, revealedIndexes)


	diplomaPresentation := SampleDiplomaPresentation{}
	diplomaPresentation.Degree=degree
	diplomaPresentation.Grade=grade
	diplomaPresentation.BfIndexes=bfIndexes
	diplomaPresentation.MtLeafHash = mtLeafHash
	diplomaPresentation.Nonce = make([]byte, len(nonce))
	copy(diplomaPresentation.Nonce[:], nonce[:])
	diplomaPresentation.Proof = make([]byte, len(SDproof))
	copy(diplomaPresentation.Proof[:], SDproof[:])
	vp := models.VerifiablePresentation{}
	vp.Messages=diplomaPresentation
	vp.Proof = make([]byte, len(SDproof))

	copy(vp.Proof[:], SDproof[:])

	//zap.S().Infoln("DIPLOMA - DIPLOMA - proof: ", vp.Proof[0:5])

	//var revealedMessages [][]byte
	//presentation := vp.Messages.(SampleDiplomaPresentation)
	//for i:=0; i<techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER; i++{
	//	revealedMessages = append(revealedMessages, []byte(presentation.BfIndexes[i]))
	//	//zap.S().Infoln("DIPLOMA - verification of selective disclosure: bf index: ", []byte(bfIndex))
	//}
	//// 2) append the mt leaf hash to the messages
	//revealedMessages = append(revealedMessages, []byte(presentation.MtLeafHash))
	//
	//
	//// 4) append the claims. All the claims are appended together. However, it is trivial to implement
	//revealedMessages = append(revealedMessages, []byte(presentation.Grade))
	//revealedMessages = append(revealedMessages, []byte(presentation.Degree))
	//
	//
	////zap.S().Infoln("DIPLOMA - DIPLOMA - proof: ", presentation.Proof[0:5])
	//
	//
	//status:= signature.VerifySelectiveDisclosureProof(publicKey, presentation.Proof, revealedMessages, presentation.Nonce)
	//zap.S().Infoln("DIPLOMA - selective disclosure:\t verification status: ", status)
	//zap.S().Infoln("BBS - selective disclosure: proof: ", SDproof, "\t nonce: ", nonce, "\t verification status: ", status)

	return vp, nil

}

func VerifySelectiveDisclosureDiploma( publicKey []byte, vp SampleDiplomaPresentation) bool{


	var messages [][]byte

	// 1) append the bf indexes to the messages
	for i:=0; i<techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER; i++{
		messages = append(messages, []byte(vp.BfIndexes[i]))
		//zap.S().Infoln("DIPLOMA - verification of selective disclosure: bf index: ", []byte(bfIndex))
	}
	// 2) append the mt leaf hash to the messages
	messages = append(messages, []byte(vp.MtLeafHash))


	// 4) append the claims. All the claims are appended together. However, it is trivial to implement
	messages = append(messages, []byte(vp.Grade))
	messages = append(messages, []byte(vp.Degree))


	//zap.S().Infoln("DIPLOMA - DIPLOMA - proof: ", vp.Proof[0:5])
	//Todo: I hardcoded bytes 2,3,4. Find a fix.
	//vp.Proof[2]=0
	//vp.Proof[3]=6
	//vp.Proof[4]=255

	status := signature.VerifySelectiveDisclosureProof( publicKey, vp.Proof, messages, vp.Nonce)


	if status == true {
		//zap.S().Infoln("DIPLOMA PRESENTATION - verification successful: ")
		return true
	}
	zap.S().Infoln("DIPLOMA PRESENTATION - verification failed")
	return false
}


func JsonToDiplomaVC(jsonObj []byte) (*models.VerifiableCredential){
	credential := models.VerifiableCredential{}
	var claims DiplomaClaim = DiplomaClaim{}
	json.Unmarshal(jsonObj, &credential)
	claimsJson,_ := json.Marshal(credential.Claims)
	json.Unmarshal(claimsJson, &claims)
	credential.Claims = claims
	return &credential
}


func JsonToDiplomaVP(jsonObj []byte) (*models.VerifiablePresentation){
	vp := models.VerifiablePresentation{}
	var diplomaVP SampleDiplomaPresentation = SampleDiplomaPresentation{}
	json.Unmarshal(jsonObj, &vp)
	messagesJson,_ := json.Marshal(vp.Messages)
	json.Unmarshal(messagesJson, &diplomaVP)
	vp.Messages = diplomaVP
	return &vp
}

func TestDiploma(){
	bbsKeys := signature.GenerateKeyPair()

	privateKey := bbsKeys.PrivateKey


	publicKey, _ := bbsKeys.PublicKey.Marshal()

	var bfIndexes []string
	for i := 0; i < techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER; i++ {
		bfIndexes = append(bfIndexes, strconv.Itoa(rand.Int()))
	}

	zap.S().Infoln("bf indexes: ", bfIndexes)
	mtLeafHash := "leaf1Hash"

	vcId := "vc#1"
	claims, _ := CreateDiplomaClaims(vcId)

	myDiploma,_ := CreateDiploma(privateKey, vcId, claims, bfIndexes, mtLeafHash )

	verifyProofForDiploma(publicKey, myDiploma.Proofs[0].ProofValue, vcId, claims, bfIndexes, mtLeafHash )
	//zap.S().Infoln("DIPLOMA - json \t: ", string(myDiploma.Json()))

	vp, _ := GenerateProofForSelectiveDisclosure(publicKey, *myDiploma)


	diplomaPresentation := vp.Messages.(SampleDiplomaPresentation)
	zap.S().Infoln("DIPLOMA - Presentation: ",diplomaPresentation.Proof[0:5])



	VerifySelectiveDisclosureDiploma(publicKey, diplomaPresentation)
	//
	//diplomaPresentation.Grade="C"
	//
	//VerifySelectiveDisclosureDiploma( publicKey, diplomaPresentation)

	jsonObj := myDiploma.Json()
	JsonToDiplomaVC(jsonObj)

	//vpJson := vp.Json()
	//vp1 := JsonToDiplomaVP(vpJson)
	//
	//VerifySelectiveDisclosureDiploma( publicKey, vp1.Messages.(SampleDiplomaPresentation))


}


