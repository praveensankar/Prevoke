package issuer

import (
	"encoding/hex"
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
	issue(config config.Config, credential verifiable.Credential)
	revoke(config config.Config, credential verifiable.Credential)
	setRevocationService(rs IRevocationService)
}


type Issuer struct{
	name string
	credentialStore []verifiable.Credential
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
	issuer.revocationProofs = make(map[string]*RevocationData)
	rand.Seed(time.Now().UnixNano())
	issuer.vcCounter = rand.Intn(100000)
	rs := CreateRevocationService(config)
	issuer.setRevocationService(rs)
	zap.S().Infoln("ISSUER-","new issuer created: issuer name - ",issuer.name)
	return &issuer
}


func (issuer *Issuer) setRevocationService(rs IRevocationService) {
	issuer.RevocationService = rs
}

func (issuer *Issuer) generateDummyVC() *verifiable.Credential {
	// step 1 - issuer generates new VC
	issuer.vcCounter = issuer.vcCounter+1
	vcId := strconv.Itoa(issuer.vcCounter)
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	zap.S().Infoln("ISSUER- ", "generating dummy vc- \t id: ",vcId)
	vc := vc.CreateEmployementProofCredential(vcId)
	return vc
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


func (issuer *Issuer) issue(vc verifiable.Credential) {
	// when issuer issue new credentials, the credential is created
	issuer.AddCretentialToStore(vc)
	zap.S().Infoln("ISSUER- ",issuer.name, "***ISSUED*** vc:", vc.ID)
	revocationData := issuer.RevocationService.IssueVC(vc)
	issuer.RevocationService.PrintMerkleTree()



	issuer.AddRevocationProofForNewVC(revocationData)
	zap.S().Infoln("ISSUER- \t sending revocation data to holder: ", revocationData.PrintRevocationData)
	//Todo: send revocationData and vc to Holder

}

func (issuer *Issuer) getUpdatedMerkleProof(vc verifiable.Credential) *merkletree.Proof {

	merkleProof := issuer.RevocationService.RetreiveUpdatedProof(vc)
	zap.S().Infoln("ISSUER- ", issuer.name, "\t vc:", vc.ID, "\t updated merkle proof: ",merkleProof)
	issuer.UpdateMerkleProofInRevocationData(vc.ID, merkleProof)
	return merkleProof


}

func (issuer *Issuer) revoke(vc verifiable.Credential) {
	issuer.RevocationService.RevokeVC(vc)
	zap.S().Infoln("ISSUER-", issuer.name, "***REVOKED*** vc:", vc.ID)
}

func (issuer *Issuer) verifyTest(vc verifiable.Credential) {

	for _,rd := range issuer.revocationProofs{
		if (rd.vcId==vc.ID){
			var bfIndexes [techniques.NUMBER_OF_INDEXES_PER_ENTRY_IN_BLOOMFILTER]*big.Int
			for i,v := range rd.BloomFilterIndexes{
				bfIndexes[i]=v;
			}
			status,_ := issuer.RevocationService.VerificationPhase1(bfIndexes)
			zap.S().Infoln("ISSUER- \t Verification of VC: \t id: ",vc.ID, "\t phase 1 (bloomfilter) result : ",status)

			//h, _ := rd.MerkleTreeLeafValue.MarshalText()
			//merkleLeaf := [32]byte{}
			//copy(merkleLeaf[:], h)
			//
			//var merkleProof [][32]byte
			//for _, v := range rd.MerkleProof.AllSiblings(){
			//	h, _ := v.MarshalText()
			//	proof := [32]byte{}
			//	copy(proof[:], h)
			//	merkleProof = append(merkleProof, proof)
			//}
			mtRoot, _ := issuer.RevocationService.VerificationPhase2()

			byteRepr := [32]byte{}
			copy(byteRepr[:], mtRoot[:])
			hexString := hex.EncodeToString(mtRoot[:])
			rootHash,_ := merkletree.NewHashFromHex(hexString)
			//zap.S().Infoln("ISSUER- Verification test- merkle tree root: ", rootHash.Hex())
			issuer.RevocationService.LocalMTVerification(rootHash, rd)

			break;
		}
	}
}


