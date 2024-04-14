package entities

import (
	"encoding/gob"
	"fyne.io/fyne/v2"
	"github.com/Revocation-Service/common"
	"github.com/Revocation-Service/config"
	"github.com/Revocation-Service/models"
	"github.com/Revocation-Service/techniques"
	"github.com/Revocation-Service/vc"
	"go.uber.org/zap"
	"net"
	"strconv"
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
	if holder.Debug==true {
		zap.S().Infoln(" HOLDER - received vc from issuer via: ", conn.RemoteAddr().String())
		zap.S().Infoln("HOLDER - new vc: ", cred.GetId())
	}
	holder.Lock()
	holder.StoreVC(*cred)
	defer holder.Unlock()
}

func(holder *Holder) sendVP(vcID string, vp models.VerifiablePresentation, address string, results *common.Results) (bool, bool, bool){
	conn, err := net.Dial("tcp", address)
	falsePositive := false
	fromDLT := false
	if err != nil {
		zap.S().Infoln("HOLDER - verifier is unavailabe")
		conn.Close()
		return false, false, false
	}

	// step 1 - Holder sends a connection request to send a VP to a verifier

	encoder := gob.NewEncoder(conn)
	//encoder.Encode(s.GetType())
	req := common.NewRequest()
	req.SetId(holder.name)
	req.SetType(common.VerifyVC)
	reqJson, _ := req.Json()
	//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
	encoder.Encode(reqJson)


	// step 2 - Holder receives  a reply from the verifier asking to send VP

	dec := gob.NewDecoder(conn)
	var replyJson []byte
	dec.Decode(&replyJson)
	reply := common.JsonToRequest(replyJson)
	if reply.GetType() == common.SendVP {

		// step 3 - Holder sends a VP to the verifier. Initiates the phase 1 verification
		encoder := gob.NewEncoder(conn)
		vpJson := vp.Json()
		//zap.S().Infoln("HOLDER - sending vp: ")
		encoder.Encode(vpJson)
	}

	// step 4 - Holder receives the result of phase 1 verification from the verifier


	phase1ReplyDecoder := gob.NewDecoder(conn)
	var phase1ReplyJson []byte
	phase1ReplyDecoder.Decode(&phase1ReplyJson)
	phase1Reply := common.JsonToRequest(phase1ReplyJson)
	//zap.S().Infoln("HOLDER - phase 1 reply: ",phase1Reply)

	// step 5 - return true if the verification resulted in success
	if phase1Reply.GetType()== common.SuccessfulVerification {
		conn.Close()
		return true, falsePositive, false
	}

	// step 6 - phase 1 resulted in failure, verifier asks the holder to witness
	if phase1Reply.GetType()== common.SendWitness {


		// step 7 - Holder checks the merkle tree from the smart contract and identifies whether the holder
		// can update the directly from the smart contract

		TimeTofetchWitnessFromSCStart := time.Now()
		mTree := holder.RevocationService.FetchMerkleTree()
		TimeTofetchWitnessFromSC := time.Since(TimeTofetchWitnessFromSCStart).Seconds()
		results.AddAvgTimeToFetchWitnessFromSmartContract(TimeTofetchWitnessFromSC)
		//zap.S().Infoln("HOLDER - merkle tree from smart contract: ", mTree)
		localMerkleProof := holder.merkleProofStore[vcID]
		ancesstorIndex := localMerkleProof.AncesstorIndex
		//zap.S().Infoln("HOLDER - local merkle proof: ", localMerkleProof.Witnesses, "\t ancesstor index: ",ancesstorIndex)
		//zap.S().Infoln("HOLDER - MT Height: ", holder.MTHeight, "\t MT Level in DLT:", holder.MTLevelInDLT)

		//currentHash := localMerkleProof.LeafHash
		//hashValue := currentHash
		//j:=0

		//for i:=holder.MTHeight;i>holder.MTLevelInDLT;i--{
		//	witness:=localMerkleProof.OrderedWitnesses[j]
		//	j++
		//	if witness.Position==techniques.LEFT{
		//		hashValue = techniques.GetHash(witness.HashValue + currentHash)
		//	}
		//	if witness.Position==techniques.RIGHT{
		//		hashValue = techniques.GetHash(currentHash + witness.HashValue)
		//	}
		//}
		if localMerkleProof.AncesstorValue == mTree[ancesstorIndex]{

			TimeToComputeWitnessFromSCStart := time.Now()
			if holder.Debug==true {
				zap.S().Infoln("HOLDER - fetches witness from the smart contract: vc id: ", vcID)
			}
			for i:=0;i< len(mTree);i++{
				_, ok := localMerkleProof.Witnesses[i]
				if ok{
					localMerkleProof.Witnesses[i].HashValue = mTree[i]
				}
			}
			localMerkleProof.OrderedWitnesses = techniques.OrderWitnesses(localMerkleProof)
			TimeToComputeWitness := time.Since(TimeToComputeWitnessFromSCStart).Seconds()
			results.AddAvgTimeToComputeCorrectWitnessAtHolder(TimeToComputeWitness)
			results.IncrementNumberofVCsRetrievedWitnessesFromDLT()

			fromDLT = true
		} else{
			// step 8 - Holder can't update the witness using smart contract. Holder contacts the issuer to retreive the
			// updated witness
			conn1, err := net.Dial("tcp", holder.issuerAddress)
			if err != nil {
				zap.S().Infoln("HOLDER - issuer is unavailabe")
				conn.Close()
				return false, false, false
			}
			if holder.Debug==true {
				zap.S().Infoln("HOLDER - requests merkle proof from issuer: vc id: ", vcID)
			}

			TimeTofetchWitnessFromIssuerStart := time.Now()
			witReqEncoder := gob.NewEncoder(conn1)
			//encoder.Encode(s.GetType())
			witReq := common.NewRequest()
			witReq.SetId(vcID)
			witReq.SetType(common.SendWitness)
			witReqJson, _ := witReq.Json()
			//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
			witReqEncoder.Encode(witReqJson)


			// step 9 - Holder receives the updated witness from the issuer and stores it
			proofDecoder := gob.NewDecoder(conn1)
			var merkleProofJson []byte
			proofDecoder.Decode(&merkleProofJson)
			merkleProof, err := techniques.JsonToMerkleProof(merkleProofJson)
			if err!=nil{
				reply := common.JsonToRequest(merkleProofJson)
				zap.S().Infoln("Holder - vc id: ", vcID, "\t reply from issuer: ",reply.GetType())
			}
			if err==nil {
				holder.Lock()
				holder.StoreMerkleProof(vcID, *merkleProof)
				holder.Unlock()
				conn1.Close()
			}
			TimeTofetchWitnessFromIssuer :=time.Since(TimeTofetchWitnessFromIssuerStart).Seconds()
			results.AddAvgTimeToFetchWitnessFromIssuer(TimeTofetchWitnessFromIssuer)
			//Todo: Revoked VCs are also counted
			results.IncrementNumberofVCsRetrievedWitnessesFromIssuer()

		}


		// step 10 - Holder sends the merkle witness to the verifier
		proof := holder.merkleProofStore[vcID]
		proof.OrderedWitnesses = techniques.OrderWitnesses(proof)
		proofEncoder := gob.NewEncoder(conn)
		proofJson := proof.Json()
		//zap.S().Infoln("HOLDER - sending vp: ")
		proofEncoder.Encode(proofJson)


		// step 11 - Holder receives the phase 2 result
		phase2ReplyDecoder := gob.NewDecoder(conn)
		var phase2ReplyJson []byte
		phase2ReplyDecoder.Decode(&phase2ReplyJson)
		phase2Reply := common.JsonToRequest(phase2ReplyJson)

		if phase2Reply.GetType()== common.SuccessfulVerification {
			results.NumberOfFalsePositives = results.NumberOfFalsePositives+1
			conn.Close()
			return true, true, fromDLT
		}
		if phase2Reply.GetType()== common.FailedVerification {
			conn.Close()
			return false, false, fromDLT
		}
	}


	conn.Close()
	return false, false, false
}

func(holder *Holder) getContractAddressFromIssuer(address string) (string){
	conn, err := net.Dial("tcp",address)
	if err != nil {
		zap.S().Infoln("HOLDER - issuer is unavailabe")
		return ""
	}

	encoder := gob.NewEncoder(conn)
	//encoder.Encode(s.GetType())
	req := common.NewRequest()
	req.SetId(holder.name)
	req.SetType(common.GetContractAddress)
	reqJson, _ := req.Json()

	encoder.Encode(reqJson)

	dec := gob.NewDecoder(conn)
	//dec.Decode(&entity)
	var jsonObj []byte
	dec.Decode(&jsonObj)
	reply := common.JsonToRequest(jsonObj)
	if holder.Debug==true {
		zap.S().Infoln("HOLDER - contract address from issuer: ", reply.GetId())
	}
	conn.Close()
	return reply.GetId()
}

func(holder *Holder) receiveVCsAtOnce(address string){

		conn, err := net.Dial("tcp",address)
		if err != nil {
			zap.S().Infoln("HOLDER - issuer is unavailabe")
			return
		}

		//zap.S().Infoln("HOLDER -  address : ",conn.LocalAddr().String())
		//zap.S().Infoln("connecting with the issuer via ", conn.RemoteAddr().String())

		encoder := gob.NewEncoder(conn)
		//encoder.Encode(s.GetType())
		req := common.NewRequest()
		req.SetId(holder.name)
		req.SetType(common.GetVCs)
		reqJson, _ := req.Json()
		//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
		encoder.Encode(reqJson)
		dec := gob.NewDecoder(conn)
		var vcoffersJson []byte
		//ticker := time.NewTicker(1 * time.Millisecond)
		//for {
		//	select {
		//	case <-ticker.C:
		dec.Decode(&vcoffersJson)
		vcOffers := common.JsonToVCOffers(vcoffersJson)


		for i:=0; i< len(vcOffers);i++{
		cred := vcOffers[i].VC
		merkleProof :=vcOffers[i].MerkleProof
		holder.Lock()
		holder.StoreVC(*cred)
		holder.Unlock()
		holder.StoreMerkleProof(cred.GetId(), *merkleProof)
		holder.Unlock()
		}
		conn.Close()
		//break
		//	}
		//}
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
		req := common.NewRequest()
		req.SetId(holder.name)
		req.SetType(common.GetVC)
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

		holder.Lock()
		holder.StoreVC(*cred)
		holder.Unlock()

		proofEncoder := gob.NewEncoder(conn)
		//encoder.Encode(s.GetType())
		proofReq := common.NewRequest()
		proofReq.SetId(holder.name)
		proofReq.SetType(common.GetMerkleProof)
		proofReqJson, _ := proofReq.Json()
		//zap.S().Infoln("HOLDER - sending new request: ", JsonToRequest(reqJson))
		proofEncoder.Encode(proofReqJson)

		proofDecoder := gob.NewDecoder(conn)
		var merkleProofJson []byte
		proofDecoder.Decode(&merkleProofJson)
		merkleProof, err := techniques.JsonToMerkleProof(merkleProofJson)
		if err!=nil{
			zap.S().Infoln("HOLDER - error retrieving merkle proof from issuer: ",err)
			return
		}

		if cred.GetId()==""{
			zap.S().Infoln("HOLDER - error retrieving VC from issuer: ",err)
			return
		}
		if holder.Debug==true {
			zap.S().Infoln("HOLDER - received new vc: ", cred.GetId(), "\t merkle proof: ", merkleProof.String())
		}
		holder.Lock()
		holder.StoreMerkleProof(cred.GetId(), *merkleProof)
		holder.Unlock()

		conn.Close()
		//break
		//	}
		//}
	}

}

func(holder *Holder) retrieveandResetResultsAtIssuers(address string)*common.Results {

		conn, err := net.Dial("tcp",address)
		if err != nil {
			zap.S().Infoln("HOLDER - issuer is unavailabe")
			return nil
		}

		//zap.S().Infoln("HOLDER -  address : ",conn.LocalAddr().String())
		//zap.S().Infoln("connecting with the issuer via ", conn.RemoteAddr().String())

		encoder := gob.NewEncoder(conn)

		req := common.NewRequest()
		req.SetId(holder.name)
		req.SetType(common.GetandResetResult)
		reqJson, _ := req.Json()
		encoder.Encode(reqJson)


		dec := gob.NewDecoder(conn)
		var resJson []byte
		dec.Decode(&resJson)
		res := common.JsonToResults(resJson)

	if holder.Debug==true {
		zap.S().Infoln("HOLDER - received revocation results from issuer")
	}
		conn.Close()
		return res
		//break
		//	}
		//}
}


func(holder *Holder) retrieveandResetResultsAtVerifiers(address string) *common.Results {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		zap.S().Infoln("HOLDER - verifier is unavailabe")
		conn.Close()
	}

	encoder := gob.NewEncoder(conn)

	req := common.NewRequest()
	req.SetId(holder.name)
	req.SetType(common.GetandResetResult)
	reqJson, _ := req.Json()

	encoder.Encode(reqJson)

	dec := gob.NewDecoder(conn)
	var resJson []byte
	dec.Decode(&resJson)
	res := common.JsonToResults(resJson)

	if holder.Debug==true {
		zap.S().Infoln("HOLDER - received revocation results from verifier")
	}
	return res
}

func (holder *Holder) CalculateVCsThatWouldRetrieveWitnessFromDLT(address string,  exp *config.Experiment) (int, int){
	conn, err := net.Dial("tcp", address)
	if err != nil {
		zap.S().Infoln("HOLDER - issuer is unavailabe")
		conn.Close()
	}

	encoder := gob.NewEncoder(conn)

	req := common.NewRequest()
	req.SetId(holder.name)
	req.SetType(common.CalculateVCsRetreivingWitnessFromDLT)
	reqJson, _ := req.Json()

	encoder.Encode(reqJson)


	dec := gob.NewDecoder(conn)
	var replyJson []byte
	dec.Decode(&replyJson)
	_ = common.JsonToRequest(replyJson)


		// step 3 - Holder sends a VP to the verifier. Initiates the phase 1 verification
		expEncoder := gob.NewEncoder(conn)
		expJson, _ := exp.Json()
		//zap.S().Infoln("HOLDER - sending vp: ")
		expEncoder.Encode(expJson)

		witReplyDecoder := gob.NewDecoder(conn)
		var resJson []byte
		witReplyDecoder.Decode(&resJson)
		res := common.JsonToCalWitnessReply(resJson)
		conn.Close()
	numberOfFalsePositives, err := strconv.Atoi(res.GetFalsePositives())
	numberOfVCsRetrievingVCsFromDLT, err := strconv.Atoi(res.GetNumberOfVCsRetrievingVCsFromDLT())
	return numberOfFalsePositives, numberOfVCsRetrievingVCsFromDLT
}

func (holder *Holder) sendExpConfig(address string, exp *config.Experiment) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		zap.S().Infoln("HOLDER - issuer is unavailabe")
		conn.Close()
	}

	encoder := gob.NewEncoder(conn)

	req := common.NewRequest()
	req.SetId(holder.name)
	req.SetType(common.SetExpConfigs)
	reqJson, _ := req.Json()
	encoder.Encode(reqJson)

	dec := gob.NewDecoder(conn)
	var replyJson []byte
	dec.Decode(&replyJson)
	reply := common.JsonToRequest(replyJson)

	if reply.GetType() == common.SendExpConfigs {

		// step 3 - Holder sends a VP to the verifier. Initiates the phase 1 verification
		expEncoder := gob.NewEncoder(conn)
		expJson, _ := exp.Json()
		//zap.S().Infoln("HOLDER - sending vp: ")
		expEncoder.Encode(expJson)
	}


}
