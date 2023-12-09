package issuer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

type IIsser interface {
	generateDummyVC() verifiable.Credential
	generateDummyVCs(count int) []*verifiable.Credential
	issue(config config.Config, credential verifiable.Credential)
	revoke(config config.Config, credential verifiable.Credential)
	setRevocationService(rs IRevocationService)
	UpdateMerkleProofsInStorage()
}


type Issuer struct{
	name string
	credentialStore []verifiable.Credential
	revokedVcIDs []string
	revocationProofs map[string]*RevocationData
	vcCounter int
	blockchainEndPoint *ethclient.Client
	RevocationService IRevocationService
}

/*
creates new issuer instance.
sets up revocation service and blockchain endpoint
 */
func  CreateIssuer(config config.Config) *Issuer{

	issuer := Issuer{}
	// register public keys at the revocation service
	// ideally, this step should be performed before the starting of the issuance process

	// connect to the blockchain network
	var err error
	issuer.blockchainEndPoint, err = ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	issuer.name = config.IssuerName
	issuer.credentialStore = []verifiable.Credential{}
	issuer.revokedVcIDs = []string{}
	issuer.revocationProofs = make(map[string]*RevocationData)
	rand.Seed(time.Now().UnixNano())
	issuer.vcCounter = rand.Intn(100000)
	rs := CreateRevocationService(config)
	issuer.setRevocationService(rs)
	zap.S().Infoln("ISSUER-","new issuer created: issuer name - ",issuer.name)
	zap.S().Infoln("\n\n********************************************************************************************************************************")

	return &issuer
}


func (issuer *Issuer) setRevocationService(rs IRevocationService) {
	issuer.RevocationService = rs
}

func (issuer *Issuer) generateDummyVC() *verifiable.Credential {
	// step 1 - issuer generates new VC
	issuer.vcCounter = issuer.vcCounter+1
	vcId := strconv.Itoa(issuer.vcCounter)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	zap.S().Infoln("ISSUER- ", "generating dummy vc- \t id: ",vcId)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")

	vc := vc.CreateEmployementProofCredential(vcId)
	return vc
}

func (issuer *Issuer) generateDummyVCs(count int) []*verifiable.Credential {
	// step 1 - issuer generates new VC

	var vcId string
	vcs := []*verifiable.Credential{}
	for i:=0 ; i<count ; i++ {
		issuer.vcCounter = issuer.vcCounter+1
		vcId = strconv.Itoa(issuer.vcCounter)
		//zap.S().Infoln("\n\n********************************************************************************************************************************")
		zap.S().Infoln("ISSUER- ", "generating dummy vc- \t id: ",vcId)
		//zap.S().Infoln("\n\n********************************************************************************************************************************")

		vc := vc.CreateEmployementProofCredential(vcId)
		vcs = append(vcs, vc)
	}
	zap.S().Infoln("\n\n********************************************************************************************************************************")


	return vcs
}

func (issuer *Issuer) AddCretentialToStore(vc verifiable.Credential) {
	issuer.credentialStore = append(issuer.credentialStore, vc)
}

func (issuer *Issuer) AddRevocationProofForNewVC(data *RevocationData){
	issuer.revocationProofs[data.vcId] = data;
}

func (issuer *Issuer) UpdateMerkleProofInRevocationData(vcId string, proof *merkletree.Proof){
	issuer.revocationProofs[vcId].MerkleProof = proof;
}

func (issuer *Issuer) UpdateMerkleProofsInStorage() {
	for _,vc := range issuer.credentialStore{
		merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vc)
		zap.S().Infoln("ISSUER-  UPDATING MERKLE PROOF", issuer.name, "\t vc:", vc.ID, "\t updated merkle proof: ",merkleProof)
		issuer.UpdateMerkleProofInRevocationData(vc.ID, merkleProof)
	}

}

func (issuer *Issuer) issue(vc verifiable.Credential) {
	// when issuer issue new credentials, the credential is created
	issuer.AddCretentialToStore(vc)
	zap.S().Infoln("ISSUER- ",issuer.name, "***ISSUED*** vc:", vc.ID)
	revocationData := issuer.RevocationService.IssueVC(vc)
	//issuer.RevocationService.PrintMerkleTree()



	issuer.AddRevocationProofForNewVC(revocationData)
	zap.S().Infoln("ISSUER- \t sending revocation data to holder: ", revocationData.PrintRevocationData)
	zap.S().Infoln("\n\n********************************************************************************************************************************")

	//Todo: send revocationData and vc to Holder

}



func (issuer *Issuer) getUpdatedMerkleProof(vc verifiable.Credential) *merkletree.Proof {

	merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vc)
	zap.S().Infoln("ISSUER- ", issuer.name, "\t vc:", vc.ID, "\t merkle tree accumulator witness updated..... ")
	issuer.UpdateMerkleProofInRevocationData(vc.ID, merkleProof)
	return merkleProof
}

func (issuer *Issuer) revoke(vc verifiable.Credential) {

	zap.S().Infoln("\n\n********************************************************************************************************************************")


	zap.S().Infoln("ISSUER-", issuer.name, "***REVOKED*** vc:", vc.ID)

	issuer.RevocationService.RevokeVC(vc)
	issuer.revokedVcIDs = append(issuer.revokedVcIDs, vc.ID)

	zap.S().Infoln("\n\n********************************************************************************************************************************")

}

func (issuer *Issuer) verifyTest(vc verifiable.Credential) {

	zap.S().Infoln("\n********************************************************************************************************************************")
	zap.S().Infoln("***********************\t  Verification test: \t VC id: ", vc.ID, "***********************")
	var bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int


	// ***************************** Phase 1 **************************************************

	rd := issuer.revocationProofs[vc.ID]

	for i,v := range rd.BloomFilterIndexes{
		bfIndexes[i]=v;
	}
	_, err := issuer.RevocationService.VerificationPhase1(bfIndexes)
	if err != nil {
		return
	}

	//if revocationStatus == true{
	//	return
	//}


	// ***************************** update witness only for valid vcs ***********************************
	status := true
	for _, vcID := range issuer.revokedVcIDs {
		if vcID == vc.ID {
			status = false
		}
	}
	//zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)
	if status == true {
		issuer.getUpdatedMerkleProof(vc)
	}

	rd1 := issuer.revocationProofs[vc.ID]
	//rd1.PrintRevocationData()

	// ***************************** Phase 2 **************************************************
	//mtRoot, _ := issuer.RevocationService.GetMerkleRoot()
	//issuer.RevocationService.LocalMTVerification(mtRoot, rd1)
	_, err = issuer.RevocationService.VerificationPhase2(rd1)
	if err != nil {
		return
	}
	// This is to check whether the VC is actuall revoked or not
	//revokedStatus := true
	//for _, vcID := range issuer.revokedVcIDs {
	//	if vcID == vc.ID {
	//		revokedStatus = false
	//	}
	//}
	////zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)
	//if revokedStatus == true {
	//	issuer.getUpdatedMerkleProof(vc)
	//}


	//zap.S().Infoln("ISSUER- \t vc id: ",vc.ID, "\t status: : ", status)


}
func (issuer *Issuer) verifyLocalTest(vc verifiable.Credential) {

	zap.S().Infoln("ISSUER- ",issuer.name, "***VERIFY*** vc:", vc.ID)
	revokedStatus := true
	for  _, vcID:= range issuer.revokedVcIDs{
		if vcID==vc.ID{
			revokedStatus=false
		}
	}
	zap.S().Infoln("ISSUER- \t vc id: ",vc.ID, "\t status: : ", revokedStatus)
	if revokedStatus==true {
		issuer.getUpdatedMerkleProof(vc)
	}
	for _,rd := range issuer.revocationProofs{
		if (rd.vcId==vc.ID){
			var bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int
			for i,v := range rd.BloomFilterIndexes{
				bfIndexes[i]=v;
			}
			status,_ := issuer.RevocationService.VerificationPhase1(bfIndexes)
			zap.S().Infoln("ISSUER- \t Verification of VC: \t id: ",vc.ID, "\t phase 1 (bloomfilter) result : ",status)

			rootHash, _ := issuer.RevocationService.GetMerkleRoot()
			//zap.S().Infoln("ISSUER- Verification test- merkle tree root: ", rootHash.Hex())
			issuer.RevocationService.LocalMTVerification(rootHash, rd)
			zap.S().Infoln("\n\n********************************************************************************************************************************")

			break;
		}
	}
}


