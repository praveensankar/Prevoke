package common

import (
	"encoding/json"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/techniques"
	"net"
)
type RequestType string
const (
	GetVC    RequestType = "get vc"
	GetVCs    RequestType = "get vcs"
	VerifyVC             = "verify vc"
	GetMerkleProof        = "get merkle proof"
	GetMerkleProofs        = "get merkle proofs"
	SendWitness        = "send witness"
	CalculateVCsRetreivingWitnessFromDLT ="calculate witness"
	SendVP			  = "send vp"
	SuccessfulVerification = "successful verification"
	FailedVerification = "failed verification"
	StoreResults			  = "store results"
	RevokedVC = "revoked vc"
	GetContractAddress = "get contract address"
	UpdateContractAddress = "update contract address"
	ContractAddress = "contract address"
	GetandResetResult = "send results"
	Result = "results"
	SetExpConfigs = "set experiment configurations"
	SendExpConfigs = "send experiment configurations"
)

type Request struct {
	Id string
	Conn net.Conn
	Type RequestType
	VcID string
	TotalVCs uint
}

func (r *Request) SetId(id string){
	r.Id = id
}

func (r *Request) GetId() string{
	return r.Id
}


func (r *Request) SetType(requestType RequestType){
	r.Type = requestType
}

func (r *Request) GetType() RequestType {
	return r.Type
}

func (r *Request) SetTotalVCs(count uint){
	r.TotalVCs =  count
}

func (r *Request) GetTotalVCs() uint{
	return r.TotalVCs
}

func (r *Request) SetConn(conn net.Conn){
	r.Conn = conn
}

func (r *Request) SetVcID(vcID string){
	r.VcID = vcID
}


func (r *Request) Json() ([]byte, error){
	return json.MarshalIndent(r, "","    ")
}



func JsonToRequest(jsonObj []byte) *Request {
	request := Request{}
	json.Unmarshal(jsonObj, &request)
	return &request
}

func NewRequest() Request {
	r := Request{}
	return r
}



type VCOffer struct {
	VC *models.VerifiableCredential
	MerkleProof *techniques.MerkleProof
}

func VCoffersToJson(vcs []*VCOffer) []byte{
	jsonObj, _ := json.Marshal(vcs)
	return jsonObj
}

func JsonToVCOffers(jsonObj []byte) []VCOffer{
	var vcs []VCOffer
	json.Unmarshal(jsonObj, &vcs)
	return vcs
}
