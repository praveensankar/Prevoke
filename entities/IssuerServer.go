package entities

import (
	"encoding/gob"
	"encoding/json"
	"github.com/praveensankar/Revocation-Service/blockchain"
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"net"
	"time"
)


func  StartIssueVCsInBulkToHolders(config config.Config) {

	issuer := CreateIssuer(config)

	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
	totalVCs := int(config.ExpectedNumberOfTotalVCs)+remainingSpace

	claimsSet := issuer.GenerateMultipleDummyVCClaims(int(totalVCs))
	issuer.IssueBulk(claimsSet, int(totalVCs))

	credentials := issuer.CredentialStore
	for _, cred := range credentials{
		issuer.UpdateMerkleProof(cred)
	}

	for _, cred := range credentials{
		conn, err := net.Dial("tcp", config.HolderAddress)
		if err!=nil{
			zap.S().Infoln("ISSUER - holder is unavailable")
			return
		}
		encoder := gob.NewEncoder(conn)
		encoder.Encode(cred.Json())
		conn.Close()
		zap.S().Infoln("ISSUER - issued vc : ",cred.GetId(), " \t to: ",conn.RemoteAddr().String())
	}




}

func StartIssuerServer(conf config.Config){
	DeployContract(&conf, 0)
	issuer := CreateIssuer(conf)

	issuer.BulkIssuance(conf)

	//if app!=nil{
	//	go issuer.setupUIForUniversity(app)
	//}

	server, err := net.Listen("tcp", conf.IssuerAddress)
	if err!=nil{
		zap.S().Infof("ISSUER - error creating server")
		return
	}
	defer server.Close()
	issuer.serverListener(server, &conf)
	timer1 := time.NewTimer(100 * time.Second)
	<-timer1.C
}


func (issuer *Issuer) BulkIssuance(config config.Config) {
	remainingSpace := int(math.Pow(2, float64(config.MTHeight)))-int(config.ExpectedNumberOfTotalVCs)
	totalVCs := int(config.ExpectedNumberOfTotalVCs)+remainingSpace
	claimsSet := issuer.GenerateMultipleDummyVCClaims(int(totalVCs))
	issuer.IssueBulk(claimsSet, totalVCs)

	credentials := issuer.CredentialStore
	for _, vc := range credentials{
		issuer.UpdateMerkleProof(vc)
	}
}
// This function handles the incomming connections. It puts all the incoming connections into a list
func(issuer *Issuer) serverListener(server net.Listener, config *config.Config){

	count :=0
	revoked := false
	zap.S().Infoln("ISSUER - server set up and listening at : ",server.Addr().String())
	for{

		conn, err := server.Accept()
		if err != nil {
			zap.S().Errorln("ISSUER - error : %v", err)
		} else {
			dec := gob.NewDecoder(conn)
			//dec.Decode(&entity)
			var reqJson []byte
			dec.Decode(&reqJson)
			req := JsonToRequest(reqJson)
			if req.GetType() == GetContractAddress{
				contractAddressEncoder := gob.NewEncoder(conn)
				//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
				contractAddressReply := NewRequest()
				contractAddressReply.SetId(issuer.ContractAddress)
				contractAddressReply.SetType(RevokedVC)
				contractAddressReplyJson, _ := contractAddressReply.Json()
				contractAddressEncoder.Encode(contractAddressReplyJson)
				conn.Close()
			}
			if req.GetType() == GetandResetResult{
				resultEncoder := gob.NewEncoder(conn)
				zap.S().Infoln("ISSUER - sending results to holder: \t", issuer.Result.String())
				resJson, _ := issuer.Result.Json()
				resultEncoder.Encode(resJson)
				contractAddress := DeployContract(config, 0)
				issuer.Reset(*config)
				issuer.ContractAddress=contractAddress
				issuer.BulkIssuance(*config)

				count = 0
				revoked=false
				conn.Close()
			}
			if req.GetType() ==SendWitness {
				isRevoked := false
				vcID := req.GetId()
				for i := 0; i < len(issuer.revokedVcIDs); i++ {
					if vcID == issuer.revokedVcIDs[i] {
						revokedVCEncoder := gob.NewEncoder(conn)
						revokedVCReply := NewRequest()
						revokedVCReply.SetId(issuer.name)
						revokedVCReply.SetType(RevokedVC)
						revokedVCReplyJson, _ := revokedVCReply.Json()
						//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
						revokedVCEncoder.Encode(revokedVCReplyJson)
						isRevoked = true
						zap.S().Infoln("ISSUER - vc id: ", vcID, "\t revoked. Did not send merkle proof ")
						break
					}
				}
				if isRevoked == false {
					proofEncoder := gob.NewEncoder(conn)
					merkleProof := issuer.getUpdatedMerkleProof(vcID)
					proofEncoder.Encode(merkleProof.Json())
					zap.S().Infoln("ISSUER - vc id: ", vcID, "\t send merkle proof: ", merkleProof)
				}
				conn.Close()
			}
			if req.GetType() ==GetVC {
				zap.S().Infoln("ISSUER - received new request: ", req)
				encoder := gob.NewEncoder(conn)
				encoder.Encode(issuer.CredentialStore[count].Json())

				dec := gob.NewDecoder(conn)
				//dec.Decode(&entity)
				var reqJson []byte
				dec.Decode(&reqJson)
				req := JsonToRequest(reqJson)
				if req.GetType() == GetMerkleProof {
					proofEncoder := gob.NewEncoder(conn)
					merkleProof := issuer.getUpdatedMerkleProof(issuer.CredentialStore[count].GetId())
					zap.S().Infoln("ISSUER - issued vc : ", issuer.CredentialStore[count].GetId(), "  \t to: ", req.GetId())
					count = count + 1

					if count==int(config.ExpectedNumberOfTotalVCs){
						if revoked==false {
							credentials := issuer.CredentialStore
							for _, vc := range credentials {
								issuer.UpdateMerkleProof(vc)
							}

							issuer.SimulateRevocation(*config)
							revoked=true
						}
					}

					proofEncoder.Encode(merkleProof.Json())

				}
				conn.Close()
			}


		}
	}
}

func DeployContract(conf *config.Config,counter int) string{
	address, err := blockchain.DeployContract(*conf, counter)

	if err != nil {
		zap.S().Errorln("error deploying contract")
	}
	conf.SmartContractAddress = address

	if err != nil {
		zap.S().Errorln("ERROR - config.json file open error")
	}
	contractAddressMap := make(map[string]string)
	contractAddressMap["contractAddress"] = address
	jsonRes, _ := json.MarshalIndent(contractAddressMap,"","")
	//filename := fmt.Sprintf("Simulation/results/result_%v_%v_%v.json",numberOfVcs, numberOfRevokedVcs, mtLevelInDLT)
	err = ioutil.WriteFile("contractAddress.json", jsonRes, 0644)
	if err != nil {
		zap.S().Errorln("unable to write results to file")
	}
	return address

}




