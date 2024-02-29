package entities

import (
	"encoding/gob"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"net"
	"time"
)

func StartVerifierServer(config config.Config){

	verifier := CreateVerifier(config)


	//if app!=nil{
	//	go issuer.setupUIForUniversity(app)
	//}
	server, err := net.Listen("tcp", config.VerifierAddress)
	if err!=nil{
		zap.S().Infof("VERIFIER - error creating server")
		return
	}
	defer server.Close()
	verifier.serverListener(server)
	timer1 := time.NewTimer(10000 * time.Second)
	<-timer1.C
}


// This function handles the incomming connections. It puts all the incoming connections into a list
func(verifier *Verifier) serverListener(server net.Listener){

	zap.S().Infoln("VERIFIER - server set up and listening at : ",server.Addr().String())
	for{

		conn, err := server.Accept()
		if err != nil {
			zap.S().Errorln("VERIFIER - error : %v", err)
		} else {
			dec := gob.NewDecoder(conn)
			//dec.Decode(&entity)
			var reqJson []byte
			dec.Decode(&reqJson)
			req := JsonToRequest(reqJson)
			if req.GetType() ==VerifyVC{
				zap.S().Infoln("VERFIER - received new request: ",req)
				encoder := gob.NewEncoder(conn)
				req := NewRequest()
				req.SetId(verifier.name)
				req.SetType(SendVP)
				reqJson, _ := req.Json()
				encoder.Encode(reqJson)
				dec := gob.NewDecoder(conn)
				var vpJson []byte
				dec.Decode(&vpJson)
				diplomaVP := vc.JsonToDiplomaVP(vpJson)
				zap.S().Infoln("VERIFIER - received VP with following claims: \t degree: ", diplomaVP.Messages.(vc.SampleDiplomaPresentation).Degree,
					"\t grade: ", diplomaVP.Messages.(vc.SampleDiplomaPresentation).Grade)
				verifier.VerifyVP(diplomaVP)
			}


		}
	}
}
