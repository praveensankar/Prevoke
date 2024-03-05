package entities

import (
	"encoding/gob"
	"github.com/praveensankar/Revocation-Service/Results"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
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
	verifier.serverListener(server, config)
	timer1 := time.NewTimer(10000 * time.Second)
	<-timer1.C
}


// This function handles the incomming connections. It puts all the incoming connections into a list
func(verifier *Verifier) serverListener(server net.Listener, conf config.Config){

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
			if req.GetType() ==StoreResults{
				Results.WriteToFile("results_verifier.json",*verifier.Result)
				verifier.Result = Results.CreateResult()
			}
			if req.GetType() ==GetandResetResult{
				resultEncoder := gob.NewEncoder(conn)
				zap.S().Infoln("ISSUER - sending results to holder: \t", verifier.Result.String())
				resJson, _ := verifier.Result.Json()
				resultEncoder.Encode(resJson)
				verifier.Reset(conf)
			}
			if req.GetType() ==VerifyVC{
				zap.S().Infoln("VERFIER - received new request: ",req)
				vpReqEncoder := gob.NewEncoder(conn)
				vpReq := NewRequest()
				vpReq.SetId(verifier.name)
				vpReq.SetType(SendVP)
				reqJson, _ := vpReq.Json()
				vpReqEncoder.Encode(reqJson)
				vpDecoder := gob.NewDecoder(conn)
				var vpJson []byte
				vpDecoder.Decode(&vpJson)
				diplomaVP := vc.JsonToDiplomaVP(vpJson)
				zap.S().Infoln("VERIFIER - received VP with following claims: \t degree: ", diplomaVP.Messages.(vc.SampleDiplomaPresentation).Degree,
					"\t grade: ", diplomaVP.Messages.(vc.SampleDiplomaPresentation).Grade)

				// phase 1 time does not include bbs verification time
				phase1result, bbsVerificationTime, phase1Time := verifier.VerifyVPPhase1(diplomaVP)
				verifier.Result.AddBBSVerificationTimePerVP(bbsVerificationTime)
				verifier.Result.AddVerificationTimeTotal(phase1Time)

				if phase1result==true{
					verifier.Result.AddVerificationTimePerValidVC(phase1Time)
					verifier.Result.AddVerificationTimeTotalValidVCs(phase1Time)
					phase1ResEncoder := gob.NewEncoder(conn)
					phase1ResultReq := NewRequest()
					phase1ResultReq.SetId(verifier.name)
					phase1ResultReq.SetType(SuccessfulVerification)
					phase1ResultReqJson, _ := phase1ResultReq.Json()
					phase1ResEncoder.Encode(phase1ResultReqJson)

				}

				if phase1result==false{
					//fetch witness from holder

					// phase two time calculated as follows: verifier sends witness request to holder, retrieves witness
					// from the holder, verifies the witness
					phase2Start := time.Now()
					witRequestEncoder := gob.NewEncoder(conn)
					witReq := NewRequest()
					witReq.SetId(verifier.name)
					witReq.SetType(SendWitness)
					witreqJson, _ := witReq.Json()
					witRequestEncoder.Encode(witreqJson)
					zap.S().Infoln("VERFIER - sending witness request: ",witReq)
					witReplyDecoder := gob.NewDecoder(conn)
					//dec.Decode(&entity)
					var witJson []byte
					witReplyDecoder.Decode(&witJson)
					merkleProof, _ := techniques.JsonToMerkleProof(witJson)
					zap.S().Infoln("VERIFIER - received merkle proof: ", merkleProof)


					phase2result := verifier.VerifyVPPhase2(diplomaVP, *merkleProof)
					phase2Time := time.Since(phase2Start)

					verifier.Result.AddVerificationTimePerRevokedandFalsePositiveVC(phase2Time.Seconds())
					verifier.Result.AddVerificationTimeTotalRevokedandFalsePositiveVCs(phase2Time.Seconds())
					verifier.Result.AddVerificationTimeTotal(phase2Time.Seconds())

					phase2ResEncoder := gob.NewEncoder(conn)
					phase2ResultReq := NewRequest()
					phase2ResultReq.SetId(verifier.name)
					if phase2result==true{
					phase2ResultReq.SetType(SuccessfulVerification)
					} else{
						phase2ResultReq.SetType(FailedVerification)
					}
					phase2ResultReqJson, _ := phase2ResultReq.Json()
					phase2ResEncoder.Encode(phase2ResultReqJson)
				}
			}


		}
	}
}
func(verifier *Verifier) getContractAddressFromIssuer(address string) (string){
	conn, err := net.Dial("tcp",address)
	if err != nil {
		zap.S().Infoln("HOLDER - issuer is unavailabe")
		return ""
	}

	encoder := gob.NewEncoder(conn)
	//encoder.Encode(s.GetType())
	req := NewRequest()
	req.SetId(verifier.name)
	req.SetType(GetContractAddress)
	reqJson, _ := req.Json()

	encoder.Encode(reqJson)

	dec := gob.NewDecoder(conn)
	//dec.Decode(&entity)
	var jsonObj []byte
	dec.Decode(&jsonObj)
	reply := JsonToRequest(jsonObj)
	zap.S().Infoln("VERIFIER - contract address from issuer: ",reply.GetId())
	conn.Close()
	return reply.GetId()
}
// This function handles the incomming connections. It puts all the incoming connections into a list
func(verifier *Verifier) getWitnessFromHolder(server net.Listener){

}
