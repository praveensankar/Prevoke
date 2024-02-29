package entities

import (
	"encoding/json"
	"net"
)
type RequestType string
const (
	GetVC RequestType = "get vc"
	VerifyVC		  = "verify vc"
	GetWitness        = "get witness"
	SendVP			  = "send vp"
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

func (r *Request) GetType() RequestType{
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



func JsonToRequest(jsonObj []byte) *Request{
	request := Request{}
	json.Unmarshal(jsonObj, &request)
	return &request
}

func NewRequest() Request{
	r := Request{}
	return r
}


