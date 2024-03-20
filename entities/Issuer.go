package entities

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/praveensankar/Revocation-Service/common"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/models"
	"github.com/praveensankar/Revocation-Service/revocation_service"
	"github.com/praveensankar/Revocation-Service/signature"
	"github.com/praveensankar/Revocation-Service/techniques"
	"github.com/praveensankar/Revocation-Service/vc"
	"github.com/suutaku/go-bbs/pkg/bbs"
	"go.uber.org/zap"
	"math"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"
)

type IIsser interface {
	GenerateMultipleDummyVCClaims(count int) []interface{}
	Issue( claims interface{})
	IssueBulk(claims interface{}, total int)
	Revoke(config config.Config, credential models.VerifiableCredential) (mapset.Set , int64, time.Duration)
	RevokeVCInBatches(conf config.Config, vcIDs []string) (mapset.Set , int64, time.Duration)
	setRevocationService(rs revocation_service.IRevocationService)
	UpdateMerkleProofsInStorage()
	UpdateMerkleProof(vc models.VerifiableCredential)
	UpdateAffectedVCs(conf config.Config, mtIndex int) (mapset.Set , int)
	GetAffectedVCsCount() int
	// returns whether it resulted in false positive in phase 1
	VerifyTest(vp models.VerifiablePresentation) (bool, bool, bool, float64, float64, float64)
}


type Issuer struct{
	sync.RWMutex
	name            string
	CredentialStore []models.VerifiableCredential
	credentialType  string
	revokedVcIDs []string
	AffectedVCIndexes map[int]bool
	revocationProofs map[string]*revocation_service.RevocationData
	vcCounter int
	lock sync.Mutex
	bbs *bbs.Bbs
	blockchainEndPoint *ethclient.Client
	RevocationService revocation_service.IRevocationService
	BbsKeyPair        []*signature.BBS
	activeConnections []net.Conn
	processedConnections []net.Conn
	Result *common.Results
	ContractAddress string
	Debug bool
}

/*
creates new entities instance.
sets up revocation service and blockchain endpoint
 */
func  CreateIssuer(config config.Config) *Issuer{

	issuer := Issuer{}
	// register public keys at the revocation service
	// ideally, this step should be performed before the starting of the issuance process

	// connect to the blockchain network
	//var err error
	//issuer.blockchainEndPoint, err = ethclient.Dial(config.BlockchainRpcEndpoint)
	//if err != nil {
	//	zap.S().Errorln("ISSUER - error connecting to blockchain")
	//}

	issuer.name = config.IssuerName
	issuer.CredentialStore = []models.VerifiableCredential{}
	issuer.revokedVcIDs = []string{}
	issuer.revocationProofs = make(map[string]*revocation_service.RevocationData)
	issuer.AffectedVCIndexes = make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	issuer.credentialType="diploma"
	issuer.vcCounter = rand.Intn(100000)
	rs := revocation_service.CreateRevocationService(config)
	issuer.setRevocationService(rs)
	keyPair1 := signature.GenerateKeyPair()
	//keyPair2 := signature.GenerateKeyPair()
	issuer.BbsKeyPair = make([]*signature.BBS, 1)
	issuer.BbsKeyPair[0] = keyPair1
	//issuer.BbsKeyPair[1] = keyPair2
	issuer.bbs = bbs.NewBbs()
	publicKey1, _ := keyPair1.PublicKey.Marshal()
	//publicKey2, _ := keyPair2.PublicKey.Marshal()
	keys := make([][]byte, 1)
	keys[0]=publicKey1
	//keys[1]=publicKey2
	pk , _ := bbs.UnmarshalPublicKey(publicKey1)
	zap.S().Infoln("ISSUER - BBS public key: ", pk.PointG2)
	rs.AddPublicKeys(keys)

	issuer.activeConnections = []net.Conn{}
	issuer.processedConnections = []net.Conn{}
	issuer.Result = common.CreateResult()

	issuer.ContractAddress = config.SmartContractAddress
	issuer.Debug = config.DEBUG

	zap.S().Infoln("ISSUER-","new entities created: entities name - ",issuer.name)
	zap.S().Infoln("\n\n********************************************************************************************************************************")


	return &issuer
}


func (issuer *Issuer) setRevocationService(rs revocation_service.IRevocationService) {
	issuer.RevocationService = rs
}

func (issuer *Issuer) generateDummyVCClaims() interface{} {
	// step 1 - entities generates new VC
	issuer.vcCounter = issuer.vcCounter+1
	vcId := strconv.Itoa(issuer.vcCounter)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	//zap.S().Infoln("ISSUER- ", "generating dummy vc- \t id: ",vcId)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")

	//vc := vc.CreateEmployementProofCredential(vcId)
	vc, _ := vc.CreateDiplomaClaims(vcId)

	return vc
}

func (issuer *Issuer) GenerateMultipleDummyVCClaims(count int) []interface{} {
	// step 1 - entities generates new VC

	var vcId string
	var vcs []interface{}
	for i:=0 ; i<count ; i++ {
		issuer.vcCounter = issuer.vcCounter+1
		vcId = strconv.Itoa(issuer.vcCounter)
		//zap.S().Infoln("\n\n********************************************************************************************************************************")
		//zap.S().Infoln("ISSUER- ", "generating dummy vc- \t id: ",vcId)
		//zap.S().Infoln("\n\n********************************************************************************************************************************")

		//vc := vc.CreateEmployementProofCredential(vcId)
		vc, _ := vc.CreateDiplomaClaims(vcId)

		vcs = append(vcs, vc)
	}
	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	return vcs
}

func (issuer *Issuer) AddCretentialToStore(vc models.VerifiableCredential) {
	issuer.CredentialStore = append(issuer.CredentialStore, vc)
}

func (issuer *Issuer) AddRevocationProofForNewVC(data *revocation_service.RevocationData){
	issuer.revocationProofs[data.VcId] = data;
}

func (issuer *Issuer) UpdateMerkleProofInRevocationData(vcId string, proof *techniques.MerkleProof){
	issuer.revocationProofs[vcId].MerkleProof = proof;
}

func (issuer *Issuer) UpdateMerkleProofsInStorage() {
	for _,vc := range issuer.CredentialStore {
		vcID := fmt.Sprintf("%v", vc.Metadata.Id)
		merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vcID)
		if issuer.Debug==true {
			zap.S().Infoln("ISSUER-  UPDATING MERKLE PROOF", issuer.name, "\t vc:", vcID, "\t updated merkle proof: ", merkleProof)
		}
		issuer.UpdateMerkleProofInRevocationData(vcID, merkleProof)
	}

}

func (issuer *Issuer) Issue(claims interface{})  {
	// when entities issue new credentials, the credential is created
	var vcID string
	if issuer.credentialType=="diploma"{
		diplomaClaims := claims.(vc.DiplomaClaim)
		vcID = fmt.Sprintf("%v", diplomaClaims.Id)
	}
	revocationData := issuer.RevocationService.IssueVC(vcID)

	var bfIndexes []string
	for _, index := range revocationData.BloomFilterIndexes {
		bfIndexes = append(bfIndexes, string(index.String()))
	}

	mtLeafHash := revocationData.MerkleProof.LeafHash
	var credential *models.VerifiableCredential
	if issuer.credentialType=="diploma"{
		credential, _ = vc.CreateDiploma(issuer.BbsKeyPair[0].PrivateKey, vcID, claims, bfIndexes, mtLeafHash)
	}
	if issuer.Debug==true {
		zap.S().Infoln("ISSUER- ", issuer.name, "***GENERATED*** vc:", vcID, "\t mt index: ", revocationData.MtIndex,
			"\t mt leaf: ", revocationData.MerkleProof.LeafHash[:techniques.SHORT_STRING_SIZE]+"..",
			"\t bf indexes: ", revocationData.BloomFilterIndexes)
	}
	//entities.RevocationService.PrintMerkleTree()
	issuer.AddCretentialToStore(*credential)
	issuer.AddRevocationProofForNewVC(revocationData)
	//zap.S().Infoln("ISSUER- \t sending revocation data to holder: ", revocationData.PrintRevocationData)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")

}

func (issuer *Issuer) IssueBulk(claimsForMutipleVCs []interface{}, total int){


	var vcIDs []string

	for _, claims := range claimsForMutipleVCs{
		var vcID string
		if issuer.credentialType=="diploma"{
			diplomaClaims := claims.(vc.DiplomaClaim)
			vcID = fmt.Sprintf("%v", diplomaClaims.Id)
		}
		vcIDs=append(vcIDs, vcID)
	}

	revocationData, cost := issuer.RevocationService.IssueVCsInBulk(vcIDs)
	issuer.Result.BulkIssuanceCost = cost
	if issuer.Debug==true{
		zap.S().Infoln("ISSUER - bulk issuance cost (in gas):  ",issuer.Result.BulkIssuanceCost)
	}


	for i:=0; i<total; i++{

		var bfIndexes []string
		for _, index := range revocationData[i].BloomFilterIndexes {
			bfIndexes = append(bfIndexes, string(index.String()))
		}

		mtLeafHash := revocationData[i].MerkleProof.LeafHash
		var credential *models.VerifiableCredential
		if issuer.credentialType=="diploma"{
			diplomaClaims := claimsForMutipleVCs[i].(vc.DiplomaClaim)
			vcID := fmt.Sprintf("%v", diplomaClaims.Id)
			if vcID!=revocationData[i].VcId{
				zap.S().Errorln("ISSUER - vc id mismatch in issuing")
			}
			credential, _ = vc.CreateDiploma(issuer.BbsKeyPair[0].PrivateKey, revocationData[i].VcId, diplomaClaims, bfIndexes, mtLeafHash)
		}
		issuer.AddCretentialToStore(*credential)
	}

	for _, rd := range revocationData{
		issuer.AddRevocationProofForNewVC(rd)
		if issuer.Debug==true {
			zap.S().Infoln("ISSUER- ", issuer.name, "***ISSUED*** vc:", rd.VcId, "\t leaf: ", rd.MerkleProof.LeafHash)
		}
	}
}

func (issuer *Issuer) UpdateMerkleProof(credential models.VerifiableCredential)  {


	status := true
	for _, vcID := range issuer.revokedVcIDs {
		if vcID == fmt.Sprintf("%v", credential.Metadata.Id) {
			status = false
		}
	}

	if status==false{
		return
	}

	merkleProof := issuer.RevocationService.RetreiveUpdatedProof(fmt.Sprintf("%v", credential.Metadata.Id))
	//zap.S().Infoln("ISSUER- ", entities.name, "\t vc:", vc.ID, "\t merkle tree accumulator witness updated..... ")
	 issuer.UpdateMerkleProofInRevocationData(fmt.Sprintf("%v", credential.Metadata.Id), merkleProof)

}

func (issuer *Issuer) getUpdatedMerkleProof(vcID string) *techniques.MerkleProof {

	merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vcID)

	ancesstorIndex := issuer.RevocationService.FindAncesstorInMerkleTree(merkleProof.MTIndex)
	merkleProof.AncesstorIndex=ancesstorIndex
	//zap.S().Infoln("ISSUER- ", entities.name, "\t vc:", vc.ID, "\t merkle tree accumulator witness updated..... ")
	issuer.UpdateMerkleProofInRevocationData(vcID, merkleProof)
	return merkleProof
}

// returns number of vcs that are affected
func (issuer *Issuer) UpdateAffectedVCs(conf config.Config, mtIndex int) (mapset.Set , int) {


	height := int(conf.MTHeight)
	levelStoredInDLT := int(conf.MtLevelInDLT)

	var numberOfEstimatedAffectedVCs int
	affectedIndexes := mapset.NewSet()
	actualAffectedVCs := 0

	if height==levelStoredInDLT{
		return affectedIndexes, actualAffectedVCs
	} else {
		numberOfEstimatedAffectedVCs = int(math.Pow(2, float64(height-levelStoredInDLT)))
	}


	foundBlock := false
	firstLeafsIndex := int(math.Pow(2, float64(height)))-1
	lastLeafsIndex := int(math.Pow(2, float64(height+1)))-1

	for i:=firstLeafsIndex; i<= (lastLeafsIndex-numberOfEstimatedAffectedVCs+1);  i = i + numberOfEstimatedAffectedVCs {
		if foundBlock==true{
			break
		}
		end := i + numberOfEstimatedAffectedVCs
		if mtIndex < end{
			foundBlock=true
			for j:=i; j < int(i + numberOfEstimatedAffectedVCs); j++{
				if mtIndex == j {
					continue
				}
				issuer.AffectedVCIndexes[j] = true
				affectedIndexes.Add(j)
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

// returns indexes of affected vcs, amount of gwei paid, revocation time per vc
func (issuer *Issuer) Revoke(conf config.Config, vc models.VerifiableCredential) (mapset.Set , int64, time.Duration) {

	vcID := fmt.Sprintf("%v", vc.Metadata.Id)
	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	start := time.Now()
	mtIndex, amountPaid, _ := issuer.RevocationService.RevokeVC(vcID)
	end := time.Since(start)
	issuer.revokedVcIDs = append(issuer.revokedVcIDs, vcID)
	affectedIndexes, numberOfAffectedVCs := issuer.UpdateAffectedVCs(conf, mtIndex)
	zap.S().Infoln("ISSUER-", issuer.name, "***REVOKED*** vc:", vcID, "\t mt index: ", mtIndex)
	if issuer.Debug==true {
		zap.S().Infoln("ISSUER-", issuer.name, "***REVOKED*** vc:", vcID, "\t mt index: ", mtIndex,
			"\t affected VCs Indexes: ", affectedIndexes, "\t number of affected VCs: ", numberOfAffectedVCs)
	}
	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	return affectedIndexes, amountPaid, end
}

/* returns
1) indexes of affected vcs
2) amount of gwei paid
3) revocation time per batch
*/
func (issuer *Issuer) RevokeVCInBatches(conf config.Config, vcIDs []string) (mapset.Set , int64, time.Duration)  {



	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	start := time.Now()
	mtIndexes, amountPaid, _ := issuer.RevocationService.RevokeVCInBatches(vcIDs)
	end := time.Since(start)
	issuer.revokedVcIDs = append(issuer.revokedVcIDs, vcIDs...)
	 affectedIndexesAll := mapset.NewSet()
	numberOfAffectedVCsTotal:=0
	for _, mtIndex := range mtIndexes{
		affectedIndexes, numberOfAffectedVCs := issuer.UpdateAffectedVCs(conf, mtIndex)
		affectedIndexesAll= affectedIndexesAll.Union(affectedIndexes)
		numberOfAffectedVCsTotal+=numberOfAffectedVCs
	}
	zap.S().Infoln("ISSUER-", issuer.name, "***REVOKED*** vcs:", vcIDs)
	if issuer.Debug==true {
		zap.S().Infoln("ISSUER-", issuer.name, "***REVOKED*** vcs:", vcIDs, "\t mt indexes: ", mtIndexes,
			"\t affected VCs Indexes: ", affectedIndexesAll, "\t number of affected VCs: ", numberOfAffectedVCsTotal)
	}
	//zap.S().Infoln("\n\n********************************************************************************************************************************")
	return affectedIndexesAll, amountPaid, end
}

func (issuer *Issuer) IsRevoked(vcID string) (bool){
	actualStatus := true
	for _, id := range issuer.revokedVcIDs {
		if id == vcID {
			actualStatus = false
		}
	}
	return actualStatus
}

// returns
//1st argument - whether it resulted in false positive in phase 1
// 2nd argument - whether the merkle tree accumulator witness is updated only from the dlt
//3rd argument - the actual status of the vc
//4th argument - bbs verification time
// 5rd argument - phase 1 - valid vcs verification time
// 6th argument - phase 2 - revoked or false positive vcs verification time
func (issuer *Issuer) VerifyTest(vcID string, vp models.VerifiablePresentation) (bool, bool, bool, float64, float64, float64) {

	//zap.S().Infoln("\n********************************************************************************************************************************")
	//zap.S().Infoln("***********************\t  Verification test: \t VC id: ", vc.ID, "***********************")
	var bfIndexes []*big.Int

	falsePositive := false
	isAffected := false
	actualStatus := issuer.IsRevoked(vcID)

	rd := issuer.revocationProofs[vcID]
	index := rd.MtIndex
	if issuer.AffectedVCIndexes[int(index)]==true{
		isAffected = true
	}


	var bbsTime time.Duration

	// ***************************** Phase 1 **************************************************





	for i,v := range rd.BloomFilterIndexes{
		bfIndexes[i]=v;
	}

	publicKeys := issuer.RevocationService.FetchPublicKeysCached()
	publicKey := publicKeys[0]



	//verify selective disclosure
	if issuer.credentialType=="diploma"{
		diplomaPresentation := vp.Messages.(vc.SampleDiplomaPresentation)

		bbsVerificationStart := time.Now()
		vc.VerifySelectiveDisclosureDiploma(publicKey, diplomaPresentation)
		bbsTime = time.Since(bbsVerificationStart)
		for i, v:= range diplomaPresentation.BfIndexes{
			intValue, _ := strconv.Atoi(v)
			bfIndexes[i]=big.NewInt(int64(intValue))
		}
	}

	phase1Start := time.Now()
	phase1Result, err := issuer.RevocationService.VerificationPhase1(bfIndexes[:])
	if err != nil {
		return false, false, false, 0, 0, 0
	}

	phase1Time := time.Since(phase1Start)
	if phase1Result == true{
		if issuer.Debug==true {
			zap.S().Infoln("VERIFICAION TEST- \t ***VERIFICATION*** vc:", vcID, "\t actual status: ", "valid",
				"\t phase1 result: ", phase1Result)
		}
		return false, false, phase1Result, bbsTime.Seconds(), phase1Time.Seconds(), 0.0
	}


	// ***************************** update witness only for valid vcs ***********************************

	//zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)

	if actualStatus == true && phase1Result==false{
		falsePositive = true
		//zap.S().Infoln("ISSUER- \t affected vc: vc id: ", vc.ID)
		issuer.getUpdatedMerkleProof(vcID)

	}



	// this step is performed to simulate retreiving proofs from DLT
	// Todo: Create a function in revocation service to allow holders to fetch remaining parts of their witness
	var rd1 *revocation_service.RevocationData
	if actualStatus == true {
		if issuer.AffectedVCIndexes[int(rd.MtIndex)]==true {
			//Todo: simulate the delay correctly. right now I am using smart contract read time as a delay
			_ = issuer.RevocationService.FetchMerkleTree()
			issuer.getUpdatedMerkleProof(vcID)
		}
		rd1 = issuer.revocationProofs[vcID]
		if issuer.AffectedVCIndexes[int(rd.MtIndex)]==false {
			mtValues := issuer.RevocationService.FetchMerkleTree()
			for i:=0; i< len(mtValues);i++{
				wit, ok :=rd1.MerkleProof.Witnesses[i]
				if ok{
					if wit.HashValue!=mtValues[i]{
						rd1.MerkleProof.Witnesses[i].HashValue=mtValues[i]
					}
				}
			}

			rd1.MerkleProof.OrderedWitnesses = techniques.OrderWitnesses(*rd1.MerkleProof)
			//rd1.PrintRevocationData()
		}
	} else{
		rd1 = issuer.revocationProofs[vcID]
	}


	//rd1.PrintRevocationData()

	// ***************************** Phase 2 **************************************************
	//mtRoot, _ := entities.RevocationService.GetMerkleRoot()
	//entities.RevocationService.LocalMTVerification(mtRoot, rd1)

	phase2Start := time.Now()

	phase2Result, err := issuer.RevocationService.VerificationPhase2(rd1.MerkleProof.LeafHash, rd1.MerkleProof.OrderedWitnesses)

	phase2Time := time.Since(phase2Start)


	if err != nil {
		return false, false, false, 0, 0, 0
	}

	var vcStatus string
	if actualStatus==true{
		vcStatus="valid"
	} else{
		vcStatus="revoked"
	}
	if issuer.Debug==true {
		zap.S().Infoln("VERIFICAION TEST- \t ***VERIFICATION*** vc:", vcID, "\t mt index: ", rd1.MtIndex, "\t actual status: ", vcStatus,
			"\t phase1 result: ", phase1Result, "\t phase2 result: ", phase2Result)
	}
	// This is to check whether the VC is actuall revoked or not
	//revokedStatus := true
	//for _, vcID := range entities.revokedVcIDs {
	//	if vcID == vc.ID {
	//		revokedStatus = false
	//	}
	//}
	////zap.S().Infoln("ISSUER- \t vc id: ", vc.ID, "\t status: : ", revokedStatus)



	//zap.S().Infoln("ISSUER- \t vc id: ",vc.ID, "\t status: : ", status)

	return falsePositive, isAffected,  phase2Result,  bbsTime.Seconds(), phase1Time.Seconds(), phase2Time.Seconds()

}

func (issuer *Issuer) FetchMerkleTreeSizeInDLT()(uint) {
	mtSize := issuer.RevocationService.FetchMerkleTreeSizeInDLT()
	if issuer.Debug==true {
		zap.S().Infoln("ISSUER- \t merkle tree size in smart contract: ", mtSize)
	}
	return mtSize
}

func (issuer *Issuer) FetchMerkleTreeSizeLocal()(uint) {
	mtSize:= issuer.RevocationService.FetchMerkleTreeSizeLocal()
	if issuer.Debug==true {
		zap.S().Infoln("ISSUER- \t merkle tree size local: ", mtSize)
	}
	return mtSize
}
//func (entities *Issuer) verifyLocalTest(vc verifiable.Credential) {
//
//	zap.S().Infoln("ISSUER- ",entities.name, "***VERIFY*** vc:", vc.ID)
//	revokedStatus := true
//	for  _, vcID:= range entities.revokedVcIDs{
//		if vcID==vc.ID{
//			revokedStatus=false
//		}
//	}
//	zap.S().Infoln("ISSUER- \t vc id: ",vc.ID, "\t status: : ", revokedStatus)
//	if revokedStatus==true {
//		entities.getUpdatedMerkleProof(vc)
//	}
//	for _,rd := range entities.revocationProofs{
//		if (rd.VcId ==vc.ID){
//			var bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int
//			for i,v := range rd.BloomFilterIndexes{
//				bfIndexes[i]=v;
//			}
//			status,_ := entities.RevocationService.VerificationPhase1(bfIndexes)
//			zap.S().Infoln("ISSUER- \t Verification of VC: \t id: ",vc.ID, "\t phase 1 (bloomfilter) result : ",status)
//
//			rootHash, _ := entities.RevocationService.GetMerkleRoot()
//			//zap.S().Infoln("ISSUER- Verification test- merkle tree root: ", rootHash.Hex())
//			entities.RevocationService.LocalMTVerification(rootHash, rd)
//			zap.S().Infoln("\n\n********************************************************************************************************************************")
//
//			break;
//		}
//	}
//}

func (issuer *Issuer) CalculateResult(conf config.Config) {
	localMTSize := issuer.RevocationService.FetchMerkleTreeSizeLocal()
	dltMTSize := issuer.RevocationService.FetchMerkleTreeSizeInDLT()
	bfSize := issuer.RevocationService.FetchBloomFilterSizeInDLT(issuer.revokedVcIDs)
	issuer.Result.MerkleTreeSizeInDLT = int(dltMTSize)
	issuer.Result.MerkleTreeSizeTotal = int(localMTSize)
	issuer.Result.BloomFilterSize= int(bfSize)
}


func (issuer *Issuer) Reset(conf config.Config) {
	issuer.CredentialStore = []models.VerifiableCredential{}
	issuer.revokedVcIDs = []string{}
	issuer.revocationProofs = make(map[string]*revocation_service.RevocationData)
	issuer.AffectedVCIndexes = make(map[int]bool)
	rand.Seed(time.Now().UnixNano())
	issuer.credentialType="diploma"
	issuer.vcCounter = rand.Intn(100000)
	rs := revocation_service.CreateRevocationService(conf)
	issuer.setRevocationService(rs)
	issuer.Result = common.CreateResult()
	keyPair := signature.GenerateKeyPair()
	issuer.BbsKeyPair = make([]*signature.BBS, 1)
	issuer.BbsKeyPair[0] = keyPair
	issuer.bbs = bbs.NewBbs()
	publicKey1, _ := keyPair.PublicKey.Marshal()
	keys := make([][]byte, 1)
	keys[0]=publicKey1
	rs.AddPublicKeys(keys)
}

func (issuer *Issuer) SimulateRevocation(config config.Config){
	revocationBatchSize := int(config.RevocationBatchSize)


	totalRevokedVCs := int(config.ExpectedNumberofRevokedVCs)
	revokedVCs := make([]string, totalRevokedVCs)
	//totalVCs := int(config.ExpectedNumberOfTotalVCs)
	vcs := issuer.CredentialStore
	for batch:=0; batch<int(int64(math.Ceil(float64(totalRevokedVCs/revocationBatchSize)))); batch++ {
		revokedVCsInBatch := make([]string, 0)
		for i, counter := 0, 0; counter < revocationBatchSize; {

			i = 2
			for {
				vcID := fmt.Sprintf("%v", vcs[i].Metadata.Id)
				isalreadyRevoked := false
				for _, revokedId := range revokedVCs {
					if vcID == revokedId {
						isalreadyRevoked = true
						break
					}
				}
				if isalreadyRevoked == false {
					revokedVCsInBatch = append(revokedVCsInBatch, vcID)
					revokedVCs = append(revokedVCs, vcID)
					counter++
					break
				}
				rand.Seed(time.Now().UnixNano())
				i = rand.Intn(int(config.ExpectedNumberOfTotalVCs))
			}
		}
		indexes, amount, revocationTime := issuer.RevokeVCInBatches(config, revokedVCsInBatch)
		issuer.revokedVcIDs = append(issuer.revokedVcIDs, common.RevokedVC)
		issuer.Result.AffectedIndexes = issuer.Result.AffectedIndexes.Union(indexes)
		issuer.Result.AddRevocationCostPerBatch(amount)
		issuer.Result.AddRevocationTimePerBatch(revocationTime.Seconds())
		issuer.Result.AddRevocationTimeTotal(revocationTime.Seconds())
	}



	issuer.Result.RevocationBatchSize = revocationBatchSize

}

/*
CalculateNumberOfVCsWouldRetrieveWitnessFromDLT calculates how many valid vcs need to retrieve witness from dlt

First, it computes the list of valid vcs that are affected by the bloom filter
 */
func (issuer *Issuer) CalculateNumberOfVCsWouldRetrieveWitnessFromDLT(conf config.Config) (int, int){

	numberOfVCsRetrievingVCsFromDLT := 0
	numberOfFalsePositives := 0
	var vcIDs []string
	revokedVCIDs := make(map[string]bool)


	bf := techniques.CreateBloomFilter(conf.ExpectedNumberofRevokedVCs, conf.FalsePositiveRate)
	for i := 0; i < len(issuer.revokedVcIDs); i++ {
		revokedVCIDs[issuer.revokedVcIDs[i]] = true
		bf.RevokeInBloomFilter(issuer.revokedVcIDs[i])
	}



	for i:=0; i< int(conf.ExpectedNumberOfTotalVCs);i++{
		vcId := issuer.CredentialStore[i].GetId()
		if bf.CheckStatusInBloomFilter(vcId)==false{
		if revokedVCIDs[vcId]==false{
			numberOfFalsePositives++
				mtIndex := issuer.revocationProofs[vcId].MtIndex
				if issuer.AffectedVCIndexes[mtIndex]==false{
					numberOfVCsRetrievingVCsFromDLT++
					vcIDs = append(vcIDs, vcId)
				}
			}
		}
}
	zap.S().Infoln("VCs that would retrieve witness from DLTs: ", vcIDs)
	return numberOfFalsePositives, numberOfVCsRetrievingVCsFromDLT
}

func (issuer *Issuer) SetExperimentConfigs(conf *config.Config, exp config.Experiment){
conf.ExpectedNumberOfTotalVCs = uint(exp.TotalVCs)
conf.ExpectedNumberofRevokedVCs = uint(exp.RevokedVCs)
conf.FalsePositiveRate = exp.FalsePositiveRate
conf.MTHeight = uint(exp.MtHeight)
conf.MtLevelInDLT = uint(exp.MtLevelInDLT)
conf.RevocationBatchSize = uint(exp.RevocationBatchSize)
zap.S().Infoln("ISSUER - updated config with experiment config: ", exp.String())

}