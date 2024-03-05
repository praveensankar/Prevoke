package entities

import (
	"encoding/json"
	"fmt"
	"github.com/praveensankar/Revocation-Service/Results"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/revocation_service"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"net"
	"sync"
	"time"
)

type IHolder interface {
	StoreVC(vc models.VerifiableCredential)
	StoreMerkleProof(vcID string, proof techniques.MerkleProof)
	RetrieveVC(id interface{}) models.VerifiableCredential
	ConstructVP(vc models.VerifiableCredential) (vp *models.VerifiablePresentation)
	ShareVP(vp models.VerifiablePresentation)
}

type Holder struct {
	sync.RWMutex
	name string
	Type Entity

	Conn net.Conn

	issuerAddress string
	verifierAddress string
	verfiableCredentials []models.VerifiableCredential
	merkleProofStore map[string]techniques.MerkleProof
	totalVCs int
	MTHeight int
	MTLevelInDLT int
	RevocationService revocation_service.IRevocationService

	Results []Results.Results
}


func NewHolder(config config.Config) *Holder{
	holder := Holder{ Type: HOLDER}
	holder.merkleProofStore = make(map[string]techniques.MerkleProof)
	holder.setName(config.HolderName)
	holder.MTHeight= int(config.MTHeight)
	holder.MTLevelInDLT = int(config.MtLevelInDLT)
	return &holder
}


func (h *Holder) RequestVCFromIssuer(){
	zap.S().Infoln("HOLDER - requesting ", h.totalVCs, " vcs from issuer")
	h.receiveVCs(h.issuerAddress)
}

func (h *Holder) RetrieveandResetResultsAtIssuers(result  *Results.Results ){
	zap.S().Infoln("HOLDER - requesting results from the issuer")
	res := h.retrieveandResetResultsAtIssuers(h.issuerAddress)

	result.RevocationTimeperBatch = res.RevocationTimeperBatch
	result.RevocationTimeTotal = res.RevocationTimeTotal
	result.AmountPaid = res.AmountPaid
	result.RevocationBatchSize = res.RevocationBatchSize
	result.NumberOfWitnessUpdatesForMT = res.NumberOfWitnessUpdatesForMT
	result.MerkleTreeSizeInDLT = res.MerkleTreeSizeInDLT
	result.MerkleTreeNodesCountTotal = res.MerkleTreeNodesCountTotal
	result.BloomFilterSize = 	res.BloomFilterSize
}

func (h *Holder) RetrieveandResetResultsAtVerifiers(result  *Results.Results ){
	zap.S().Infoln("HOLDER - requesting results from the verifier")
	res := h.retrieveandResetResultsAtVerifiers(h.verifierAddress)

	result.BBSVerificationTimePerVP = res.BBSVerificationTimePerVP
	result.VerificationTimePerValidVC = res.VerificationTimePerValidVC
	result.VerificationTimeTotalValidVCs = res.VerificationTimeTotalValidVCs

	result.VerificationTimePerRevokedorFalsePositiveVC = res.VerificationTimePerRevokedorFalsePositiveVC
	result.VerificationTimeTotalRevokedorFalsePositiveVCs = res.VerificationTimeTotalRevokedorFalsePositiveVCs

	result.VerificationTimeTotal = res.VerificationTimeTotal


}

func (h *Holder) StoreVC(vc models.VerifiableCredential) {
	h.verfiableCredentials = append(h.verfiableCredentials, vc)
}

func (h *Holder) StoreResults(result Results.Results) {
	h.Results = append(	h.Results , result)
}

func (h *Holder) StoreMerkleProof(vcID string, proof techniques.MerkleProof) {
	h.merkleProofStore[vcID]=proof
}

func (h *Holder) RetrieveVC(id interface{}) models.VerifiableCredential {
	for _, vc := range h.verfiableCredentials{
		if vc.GetId()==id{
			return vc
		}
	}
	return models.VerifiableCredential{}
}

func (h *Holder) ConstructVP(credential models.VerifiableCredential) (vp models.VerifiablePresentation, err error){
	publicKeys := h.RevocationService.FetchPublicKeysCached()
	publicKey := publicKeys[0]
	//pk , _ := bbs.UnmarshalPublicKey(publicKey)
	//zap.S().Infoln("HOLDER - issuer's public keys: ", pk.PointG2)


	//zap.S().Infoln("HOLDER - proof: ",credential.Proofs)
	presentation, err := vc.GenerateProofForSelectiveDisclosure(publicKey, credential)
	if err!=nil{
		zap.S().Infoln("HOLDER - error in generating vp")
	}

	//Verification check for the newly generated presentation
	//diplomaVP := presentation.Messages.(vc.SampleDiplomaPresentation)
	//zap.S().Infoln("HOLDER - new vp: \t degree: ", diplomaVP.Degree, "\t grade:", diplomaVP.Grade)
	//zap.S().Infoln("HOLDER - new vp: : ", diplomaVP)
	//status := vc.VerifySelectiveDisclosureDiploma(publicKey, diplomaVP)
	//zap.S().Infoln("HOLDER - new vp: signature check: ",status)
	//if status==false{
	//return models.VerifiablePresentation{}, errors.New("signature check failed")
	//}

	return presentation, nil
}

func (h *Holder) ShareallVPs(results *Results.Results){
	if len(h.verfiableCredentials)==0{
		zap.S().Infoln("HOLDER - haven't recived any vcs yet")
		return
	}
	for i:=0;i<len(h.verfiableCredentials);i++ {
		vc := h.verfiableCredentials[i]
		zap.S().Infoln("HOLDER - sending presentation of vc: ", vc.GetId())
		vp, err := h.ConstructVP(vc)
		if err != nil {
			return
		}
		status := h.ShareVP(vc.GetId(), vp, results)
		zap.S().Infoln("HOLDER - verification result: ", status)
	}
}
func (h *Holder) ShareVP(vcID string, vp models.VerifiablePresentation, results *Results.Results) (bool){
	return h.sendVP(vcID, vp, h.verifierAddress, results)
}

func (h Holder) GetType() Entity{
	return h.Type
}

func (h *Holder) setConnection(conn net.Conn){
	h.Conn = conn
}


func (h *Holder) setName(name string){
	h.name = name
}

func (h *Holder) getName() string{
	return h.name
}


func (h *Holder) Json() ([]byte, error){
	return json.Marshal(h)
}

//func JsonToHolder(jsonObj []byte) *Holder{
//	holder := NewHolder("")
//	json.Unmarshal(jsonObj, holder)
//	return holder
//}

func (h *Holder) String() string  {

	var response string
	response = response + fmt.Sprintf("%v", h.Type)+"\n"
	response = response + fmt.Sprintf("%v", h.Conn)+"\n"
	return response
}





func StartHolder(config config.Config){



	for i:=0;i<2;i++ {
		holder := NewHolder(config)
		holder.issuerAddress = config.IssuerAddress
		holder.verifierAddress = config.VerifierAddress
		holder.totalVCs = int(config.ExpectedNumberOfTotalVCs)
		contractAddress := holder.getContractAddressFromIssuer(holder.issuerAddress)
		config.SmartContractAddress = contractAddress
		holder.RevocationService = revocation_service.CreateRevocationService(config)
		result := Results.CreateResult()

		start := time.Now()
		holder.RequestVCFromIssuer()
		holder.ShareallVPs(result)
		holder.RetrieveandResetResultsAtIssuers(result)
		holder.RetrieveandResetResultsAtVerifiers(result)

		Results.ConstructResults(config, start, result)
		Results.WriteToFile("results.json", *result)
	}



	timer1 := time.NewTimer(30 * time.Second)
	<-timer1.C

}







