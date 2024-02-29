package entities

import (
	"encoding/gob"
	"fyne.io/fyne/v2"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"net"
	"time"
)

func StartHolderServer(app fyne.App, config config.Config){

	holder := NewHolder(config)
	holder.issuerAddress = config.IssuerAddress


	//if app!=nil{
	//	go holder.setupUIForHolder(app)
	//}
	server, err := net.Listen("tcp", config.HolderAddress)
	if err!=nil{
		zap.S().Infof("HOLDER - error creating server")
		return
	}
	defer server.Close()
	go holder.serverListener(server)
	timer1 := time.NewTimer(100 * time.Second)
	<-timer1.C
}

// This function handles the incomming connections. It puts all the incoming connections into a list
func(holder *Holder) serverListener(server net.Listener){

	zap.S().Infoln("HOLDER - server set up and listening at : ",server.Addr().String())
	for{

		conn, err := server.Accept()
		if err != nil {
			zap.S().Errorln("HOLDER - error : %v", err)
		} else {
			go holder.processConnection(conn)
		}
	}
}

func(holder *Holder) processConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	var vcJson []byte
	//dec.Decode(&entity)
	dec.Decode(&vcJson)
	//os.WriteFile("vc.json", vcJson, 0644)
	cred := vc.JsonToDiplomaVC(vcJson)
	//len, _ := conn.Read(buffer)
	zap.S().Infoln(" HOLDER - received vc from issuer via: ",conn.RemoteAddr().String() )
	zap.S().Infoln("HOLDER - new vc: ", cred.GetId())
	holder.Lock()
	holder.StoreVC(*cred)
	defer holder.Unlock()
}

func(holder *Holder) sendVP(vp models.VerifiablePresentation, address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		zap.S().Infoln("HOLDER - verifier is unavailabe")
		return
	}

	encoder := gob.NewEncoder(conn)
	//encoder.Encode(s.GetType())
	req := NewRequest()
	req.SetId(holder.name)
	req.SetType(VerifyVC)
	reqJson, _ := req.Json()
	//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
	encoder.Encode(reqJson)
	dec := gob.NewDecoder(conn)
	var replyJson []byte
	dec.Decode(&replyJson)
	reply := JsonToRequest(replyJson)
	if reply.GetType() ==SendVP{
		encoder := gob.NewEncoder(conn)
		vpJson := vp.Json()
		//zap.S().Infoln("HOLDER - sending vp: ")
		encoder.Encode(vpJson)
	}
	conn.Close()
}

func(holder *Holder) receiveVCs(address string){
	for i:=0;i<holder.totalVCs;i++ {
		conn, err := net.Dial("tcp",address)
		if err != nil {
			zap.S().Infoln("HOLDER - issuer is unavailabe")
			return
		}

		//zap.S().Infoln("HOLDER -  address : ",conn.LocalAddr().String())
		//zap.S().Infoln("connecting with the issuer via ", conn.RemoteAddr().String())

		encoder := gob.NewEncoder(conn)
		//encoder.Encode(s.GetType())
		req := NewRequest()
		req.SetId(holder.name)
		req.SetType(GetVC)
		reqJson, _ := req.Json()
		//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
		encoder.Encode(reqJson)
		dec := gob.NewDecoder(conn)
		var vcJson []byte
		//ticker := time.NewTicker(1 * time.Millisecond)
		//for {
		//	select {
		//	case <-ticker.C:
		dec.Decode(&vcJson)
		cred := vc.JsonToDiplomaVC(vcJson)
		zap.S().Infoln("HOLDER - received new vc: ", cred.GetId())
		holder.Lock()
		holder.StoreVC(*cred)
		holder.Unlock()
		conn.Close()
		//break
		//	}
		//}
	}
}
