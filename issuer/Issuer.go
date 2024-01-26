package issuer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/praveensankar/Revocation-Service/vc"
	"go.uber.org/zap"
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

type IIsser interface {
	GenerateDummyVC() verifiable.Credential
	GenerateDummyVCs(count int) []*verifiable.Credential
	Issue(config config.Config, credential verifiable.Credential)
	IssueBulk(config config.Config, credential []*verifiable.Credential, total int)
	Revoke(config config.Config, credential verifiable.Credential) (int, int64)
	setRevocationService(rs IRevocationService)
	UpdateMerkleProofsInStorage()
	UpdateMerkleProof(vc verifiable.Credential)
	UpdateAffectedVCs(conf config.Config, mtIndex *big.Int) int
	GetAffectedVCsCount() int
	// returns whether it resulted in false positive in phase 1
	VerifyTest(vc verifiable.Credential) (bool, bool)
}


type Issuer struct{
	name string
	credentialStore []verifiable.Credential
	revokedVcIDs []string
	AffectedVCIndexes map[int]bool
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
	issuer.AffectedVCIndexes = make(map[int]bool)
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
	//zap.S().Infoln("ISSUER- ", "generating dummy vc- \t id: ",vcId)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")

	vc := vc.CreateEmployementProofCredential(vcId)
	return vc
}

func (issuer *Issuer) GenerateDummyVCs(count int) []*verifiable.Credential {
	// step 1 - issuer generates new VC

	var vcId string
	vcs := []*verifiable.Credential{}
	for i:=0 ; i<count ; i++ {
		issuer.vcCounter = issuer.vcCounter+1
		vcId = strconv.Itoa(issuer.vcCounter)
		//zap.S().Infoln("\n\n********************************************************************************************************************************")
		//zap.S().Infoln("ISSUER- ", "generating dummy vc- \t id: ",vcId)
		//zap.S().Infoln("\n\n********************************************************************************************************************************")

		vc := vc.CreateEmployementProofCredential(vcId)
		vcs = append(vcs, vc)
	}
	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	return vcs
}

func (issuer *Issuer) AddCretentialToStore(vc verifiable.Credential) {
	issuer.credentialStore = append(issuer.credentialStore, vc)
}

func (issuer *Issuer) AddRevocationProofForNewVC(data *RevocationData){
	issuer.revocationProofs[data.vcId] = data;
}

func (issuer *Issuer) UpdateMerkleProofInRevocationData(vcId string, proof *techniques.MerkleProof){
	issuer.revocationProofs[vcId].MerkleProof = proof;
}

func (issuer *Issuer) UpdateMerkleProofsInStorage() {
	for _,vc := range issuer.credentialStore{
		merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vc)
		zap.S().Infoln("ISSUER-  UPDATING MERKLE PROOF", issuer.name, "\t vc:", vc.ID, "\t updated merkle proof: ",merkleProof)
		issuer.UpdateMerkleProofInRevocationData(vc.ID, merkleProof)
	}

}

func (issuer *Issuer) Issue(vc verifiable.Credential) {
	// when issuer issue new credentials, the credential is created
	issuer.AddCretentialToStore(vc)
	revocationData := issuer.RevocationService.IssueVC(vc)
	zap.S().Infoln("ISSUER- ",issuer.name, "***ISSUED*** vc:", vc.ID, "\t mt index: ", revocationData.MtIndex,
		"\t mt leaf: ", revocationData.MerkleProof.LeafHash[:techniques.SHORT_STRING_SIZE] + "..",
		"\t bf indexes: ",revocationData.BloomFilterIndexes)
	//issuer.RevocationService.PrintMerkleTree()

	issuer.AddRevocationProofForNewVC(revocationData)
	//zap.S().Infoln("ISSUER- \t sending revocation data to holder: ", revocationData.PrintRevocationData)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")



}

func (issuer *Issuer) IssueBulk(config config.Config, vcs []*verifiable.Credential, total int){



	revocationData := issuer.RevocationService.IssueVCsInBulk(vcs)

	for i:=0; i<total; i++{
		issuer.AddCretentialToStore(*vcs[i])
	}

	for _, rd := range revocationData{
		issuer.AddRevocationProofForNewVC(rd)
		zap.S().Infoln("ISSUER- ",issuer.name, "***ISSUED*** vc:", rd.vcId, "\t leaf: ", rd.MerkleProof.LeafHash)
	}
}

func (issuer *Issuer) UpdateMerkleProof(vc verifiable.Credential)  {

	status := true
	for _, vcID := range issuer.revokedVcIDs {
		if vcID == vc.ID {
			status = false
		}
	}

	if status==false{
		return
	}

	merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vc)
	//zap.S().Infoln("ISSUER- ", issuer.name, "\t vc:", vc.ID, "\t merkle tree accumulator witness updated..... ")
	 issuer.UpdateMerkleProofInRevocationData(vc.ID, merkleProof)

}

func (issuer *Issuer) getUpdatedMerkleProof(vc verifiable.Credential) *techniques.MerkleProof {

	merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vc)
	//zap.S().Infoln("ISSUER- ", issuer.name, "\t vc:", vc.ID, "\t merkle tree accumulator witness updated..... ")
	issuer.UpdateMerkleProofInRevocationData(vc.ID, merkleProof)
	return merkleProof
}

// returns number of vcs that are affected
func (issuer *Issuer) UpdateAffectedVCs(conf config.Config, mtIndex *big.Int) ([]int, int) {

	n := int(conf.ExpectedNumberOfTotalVCs)
	actualAffectedVCs := 0
	height := int(math.Ceil(math.Log2(float64(n))))
	levelStoredInDLT := int(conf.MtLevelInDLT)
	var numberOfEstimatedAffectedVCs int
	if height==levelStoredInDLT{
		numberOfEstimatedAffectedVCs = 0
		return nil, actualAffectedVCs
	} else {
		numberOfEstimatedAffectedVCs = int(math.Pow(2, float64(height-levelStoredInDLT)))
	}
	affectedIndexes := make([]int,numberOfEstimatedAffectedVCs)

	foundBlock := false
	for i:=1; i<= (n-numberOfEstimatedAffectedVCs+1);  i = i + numberOfEstimatedAffectedVCs {
		if foundBlock==true{
			break
		}
		end := big.NewInt(int64(i + numberOfEstimatedAffectedVCs))
		if mtIndex.Cmp(end) == -1{
			foundBlock=true
			for j:=i; j < int(i + numberOfEstimatedAffectedVCs); j++{
				J := big.NewInt(int64(j))
				if mtIndex.Cmp(J)==0{
					continue
				}
				//if issuer.AffectedVCIndexes[j]!=true{
				//	issuer.AffectedVCIndexes[j] = true
				//	affectedIndexes = append(affectedIndexes, j)
				//	actualAffectedVCs++
				//}
				issuer.AffectedVCIndexes[j] = true
					affectedIndexes = append(affectedIndexes, j)
					actualAffectedVCs++

			}
			//zap.S().Infoln("ISSUER: WITNESS UPDATE - \t mt index: ",mtIndex, "\t block starting index: ",i, "\t end index: ", int64(i + numberOfEstimatedAffectedVCs)-1,
			//	"\t affected vcs: ", affectedIndexes)
		}
	}
	return affectedIndexes, actualAffectedVCs
}

func (issuer *Issuer) GetAffectedVCsCount() (int) {
	return len(issuer.AffectedVCIndexes)
}

// returns number of affected vcs, amount of gwei paid
func (issuer *Issuer) Revoke(conf config.Config, vc verifiable.Credential) (int, int64) {

	//zap.S().Infoln("\n\n********************************************************************************************************************************")




	mtIndex, amountPaid, _ := issuer.RevocationService.RevokeVC(vc)
	issuer.revokedVcIDs = append(issuer.revokedVcIDs, vc.ID)
	affectedIndexes, numberOfAffectedVCs := issuer.UpdateAffectedVCs(conf, mtIndex)
	zap.S().Infoln("ISSUER-", issuer.name, "***REVOKED*** vc:", vc.ID,"\t mt index: ",mtIndex,
		"\t affected VCs Indexes: ",affectedIndexes, "\t number of affected VCs: ", numberOfAffectedVCs)

	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	return numberOfAffectedVCs, amountPaid
}

// returns
//1st argument - whether it resulted in false positive in phase 1
// 2nd argument - whether the merkle tree accumulator witness is updated only from the dlt
func (issuer *Issuer) VerifyTest(vc verifiable.Credential) (bool, bool) {

	//zap.S().Infoln("\n********************************************************************************************************************************")
	//zap.S().Infoln("***********************\t  Verification test: \t VC id: ", vc.ID, "***********************")
	var bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int

	falsePositive := false
	isAffected := false
	actualStatus := true
	for _, vcID := range issuer.revokedVcIDs {
		if vcID == vc.ID {
			actualStatus = false
		}
	}
	// ***************************** Phase 1 **************************************************

	rd := issuer.revocationProofs[vc.ID]

	for i,v := range rd.BloomFilterIndexes{
		bfIndexes[i]=v;
	}
	phase1Result, err := issuer.RevocationService.VerificationPhase1(bfIndexes)
	if err != nil {
		return false, false
	}

	//if revocationStatus == true{
	//	return
	//}


	// ***************************** update witness only for valid vcs ***********************************

	//zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)

	if actualStatus == true && phase1Result==false{
		falsePositive = true
		//zap.S().Infoln("ISSUER- \t affected vc: vc id: ", vc.ID)
		issuer.getUpdatedMerkleProof(vc)
		index := rd.MtIndex
		if issuer.AffectedVCIndexes[int(index)]==true{
			isAffected = true
		}
	}



	// this step is performed to simulate retreiving proofs from DLT
	// Todo: Create a function in revocation service to allow holders to fetch remaining parts of their witness
	if actualStatus == true {
		issuer.getUpdatedMerkleProof(vc)
	}

	rd1 := issuer.revocationProofs[vc.ID]
	//rd1.PrintRevocationData()

	// ***************************** Phase 2 **************************************************
	//mtRoot, _ := issuer.RevocationService.GetMerkleRoot()
	//issuer.RevocationService.LocalMTVerification(mtRoot, rd1)
	phase2Result, err := issuer.RevocationService.VerificationPhase2(rd1)
	if err != nil {
		return false, false
	}

	var vcStatus string
	if actualStatus==true{
		vcStatus="valid"
	} else{
		vcStatus="revoked"
	}
	zap.S().Infoln("VERIFICAION TEST- \t ***VERIFICATION*** vc:", vc.ID, "\t actual status: ",vcStatus,
		"\t phase1 result: ", phase1Result, "\t phase2 result: ", phase2Result)
	// This is to check whether the VC is actuall revoked or not
	//revokedStatus := true
	//for _, vcID := range issuer.revokedVcIDs {
	//	if vcID == vc.ID {
	//		revokedStatus = false
	//	}
	//}
	////zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)



	//zap.S().Infoln("ISSUER- \t vc id: ",vc.ID, "\t status: : ", status)

	return falsePositive, isAffected

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


