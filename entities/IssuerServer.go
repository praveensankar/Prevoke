package entities

import (
	"encoding/gob"
	"encoding/json"
	"github.com/Revocation-Service/blockchain"
	"github.com/Revocation-Service/common"
	"github.com/Revocation-Service/config"
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
		if issuer.Debug==true {
			zap.S().Infoln("ISSUER - issued vc : ", cred.GetId(), " \t to: ", conn.RemoteAddr().String())
		}
	}




}

func StartIssuerServer(conf config.Config){
	DeployContract(&conf, 0)
	issuer := CreateIssuer(conf)

	//issuer.BulkIssuance(conf)

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
func(issuer *Issuer) serverListener(server net.Listener, conf *config.Config){

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
			req := common.JsonToRequest(reqJson)
			if req.GetType() == common.GetContractAddress {
				contractAddressEncoder := gob.NewEncoder(conn)
				//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
				contractAddressReply := common.NewRequest()
				contractAddressReply.SetId(issuer.ContractAddress)
				contractAddressReply.SetType(common.RevokedVC)
				contractAddressReplyJson, _ := contractAddressReply.Json()
				contractAddressEncoder.Encode(contractAddressReplyJson)
				conn.Close()
			}
			if req.GetType() == common.SetExpConfigs {
				expReqEncoder := gob.NewEncoder(conn)
				expReq := common.NewRequest()
				expReq.SetId(issuer.name)
				expReq.SetType(common.SendExpConfigs)
				expReqJson, _ := expReq.Json()
				expReqEncoder.Encode(expReqJson)


				expDecoder := gob.NewDecoder(conn)
				//dec.Decode(&entity)
				var expJson []byte
				expDecoder.Decode(&expJson)
				exp := config.JsonToExperiment(expJson)
				issuer.SetExperimentConfigs(conf, *exp)
				contractAddress, gasUsed := DeployContract(conf, 0)

				zap.S().Infoln("ISSUER - contract deployment cost (in gas): \t", gasUsed)

				conf.SmartContractAddress=contractAddress
				issuer.Reset(*conf)
				issuer.ContractAddress=contractAddress
				issuer.BulkIssuance(*conf)
				issuer.Result.ContractDeploymentCost = gasUsed
			}

			//if req.GetType() == common.CalculateVCsRetreivingWitnessFromDLT {
			//	expReqEncoder := gob.NewEncoder(conn)
			//	expReq := common.NewRequest()
			//	expReq.SetId(issuer.name)
			//	expReq.SetType(common.SendExpConfigs)
			//	expReqJson, _ := expReq.Json()
			//	expReqEncoder.Encode(expReqJson)
			//
			//
			//	expDecoder := gob.NewDecoder(conn)
			//	//dec.Decode(&entity)
			//	var expJson []byte
			//	expDecoder.Decode(&expJson)
			//	exp := config.JsonToExperiment(expJson)
			//	issuer.SetExperimentConfigs(conf, *exp)
			//	issuer.Reset(*conf)
			//	rs := revocation_service.CreateRevocationServiceStub(*conf)
			//	issuer.setRevocationService(rs)
			//	issuer.BulkIssuance(*conf)
			//	issuer.SimulateRevocation(*conf)
			//
			//
			//	numberOfFalsePositives, numberOfVCsRetrievingVCsFromDLT := issuer.CalculateNumberOfVCsWouldRetrieveWitnessFromDLT(*conf)
			//	calWitReplyEncoder := gob.NewEncoder(conn)
			//	calWitReply := common.NewCalWitnessReply()
			//	calWitReply.SetFalsePositives(strconv.Itoa(numberOfFalsePositives))
			//	calWitReply.SetNumberOfVCsRetrievingVCsFromDLT(strconv.Itoa(numberOfVCsRetrievingVCsFromDLT))
			//	calWitReplyJson, _ := calWitReply.Json()
			//	calwitEncErr := calWitReplyEncoder.Encode(calWitReplyJson)
			//	if calwitEncErr!=nil{
			//		zap.S().Infoln("ISSUER - witness calculation encoding error: ", calwitEncErr)
			//	}
			//	zap.S().Infoln("ISSUER - number of vcs retrieving witness from dlt: \t", calWitReply.String())
			//	conn.Close()
			//}
			if req.GetType() == common.GetandResetResult {
				resultEncoder := gob.NewEncoder(conn)
				zap.S().Infoln("ISSUER - sending results to holder: \t", issuer.Result.String())

				issuer.CalculateResult(*conf)
				resJson, _ := issuer.Result.Json()
				resultEncoder.Encode(resJson)
				//issuer.BulkIssuance(*conf)

				count = 0
				revoked=false
				conn.Close()
			}
			if req.GetType() == common.SendWitness {
				isRevoked := false
				vcID := req.GetId()
				for i := 0; i < len(issuer.revokedVcIDs); i++ {
					if vcID == issuer.revokedVcIDs[i] {
						revokedVCEncoder := gob.NewEncoder(conn)
						revokedVCReply := common.NewRequest()
						revokedVCReply.SetId(issuer.name)
						revokedVCReply.SetType(common.RevokedVC)
						revokedVCReplyJson, _ := revokedVCReply.Json()
						//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
						revokedVCEncoder.Encode(revokedVCReplyJson)
						isRevoked = true
						if issuer.Debug==true {
							zap.S().Infoln("ISSUER - vc id: ", vcID, "\t revoked. Did not send merkle proof ")
						}
						break
					}
				}
				if isRevoked == false {
					proofEncoder := gob.NewEncoder(conn)
					merkleProof := issuer.getUpdatedMerkleProof(vcID)
					proofEncoder.Encode(merkleProof.Json())
					if issuer.Debug==true {
						zap.S().Infoln("ISSUER - vc id: ", vcID, "\t send merkle proof: ", merkleProof.String())
					}
				}
				conn.Close()
			}
			if req.GetType() == common.GetVC {
				if issuer.Debug==true {
					zap.S().Infoln("ISSUER - received new request: ", req)
				}
				encoder := gob.NewEncoder(conn)
				encoder.Encode(issuer.CredentialStore[count].Json())

				dec := gob.NewDecoder(conn)
				//dec.Decode(&entity)
				var reqJson []byte
				dec.Decode(&reqJson)
				req := common.JsonToRequest(reqJson)
				if req.GetType() == common.GetMerkleProof {
					proofEncoder := gob.NewEncoder(conn)
					merkleProof := issuer.getUpdatedMerkleProof(issuer.CredentialStore[count].GetId())

					zap.S().Infoln("ISSUER - issued vc : ", issuer.CredentialStore[count].GetId(), "  \t to: ", req.GetId())

					count = count + 1

					if count==int(conf.ExpectedNumberOfTotalVCs){
						if revoked==false {
							//credentials := issuer.CredentialStore
							//for _, vc := range credentials {
							//	issuer.UpdateMerkleProof(vc)
							//}

							issuer.SimulateRevocation(*conf)
							revoked=true
						}
					}

					proofEncoder.Encode(merkleProof.Json())

				}
				conn.Close()
			}
			if req.GetType() == common.GetVCs {
				if issuer.Debug==true {
					zap.S().Infoln("ISSUER - received new request: ", req)
				}
				encoder := gob.NewEncoder(conn)

				var vcOffers []*common.VCOffer
				for i:=0;i<int(conf.ExpectedNumberOfTotalVCs);i++{
					cred := issuer.CredentialStore[i]
					merkleProof := issuer.getUpdatedMerkleProof(cred.GetId())
					vcOffer := common.VCOffer{VC: &cred, MerkleProof: merkleProof}
					vcOffers = append(vcOffers, &vcOffer)
				}
				jsonObj := common.VCoffersToJson(vcOffers)
				encoder.Encode(jsonObj)

				if revoked==false {
					credentials := issuer.CredentialStore
					for _, vc := range credentials {
						issuer.UpdateMerkleProof(vc)
					}
					issuer.SimulateRevocation(*conf)
					revoked=true
				}
				conn.Close()
				}
			}
		}
	}


func DeployContract(conf *config.Config,counter int) (string, int64){
	address, gasUsed, err := blockchain.DeployContract(*conf, counter)

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
	return address, gasUsed

}




