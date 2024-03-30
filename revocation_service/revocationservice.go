package revocation_service

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/contracts"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math"
	"math/big"
	"reflect"
)

type IRevocationService interface {
	IssueVC(vcID string) *RevocationData
	IssueVCsInBulk(vcIDs []string) ([]*RevocationData, int64)
	RevokeVC(vcID string) (int, int64, error)
	RevocationCostCalculator(bfIndexes []*big.Int, mtIndexes []*big.Int, mtValuesInBytes [][32]byte) (int64, error)
	RevokeVCInBatches(vcIDs []string) (map[string]int, int64, error)
	RetreiveUpdatedProof(vcID string) *techniques.MerkleProof
	VerificationPhase1(bfIndexes []*big.Int) (bool, error)
	VerificationPhase2(leafHash string, witnesses []*techniques.Witness) (bool, error)
	VerifyVC( _bfIndexes []*big.Int, data *RevocationData) (bool, error)
	GetMerkleRoot()(string, error)
	FetchMerkleTree() ([]string)
	PrintMerkleTree()
	LocalMTVerification(mtRoot string, data *RevocationData)
	AddPublicKeys(publicKeys [][]byte)
	FetchPublicKeys()([][]byte)
	FetchPublicKeysCached()([][]byte)
	FetchMerkleTreeSizeInDLT()(uint)
	FetchMerkleTreeSizeLocal()(uint)
	FindAncesstorInMerkleTree(index int)(int, string)
	FetchBloomFilterSizeInDLT(revokedVcIDs []string)(uint)
	GetLocalBloomFilter() *techniques.BloomFilter

}


type RevocationService struct{
	merkleTreeAcc *techniques.MerkleTreeAccumulator2
	VCToBigInts map[string]*big.Int
	vcCounter int64
	bloomFilter *techniques.BloomFilter
	MtLevelInDLT int
	mtHeight int
	NumberOfEntriesForMTInDLT int
	blockchainRPCEndpoint string
	account common.Address
	smartContractAddress common.Address
	privateKey string
	gasLimit uint64
	gasPrice *big.Int
	isCached  bool
	isPublicKeysCached bool
	cachedPublicKeys [][]byte
}



func CreateRevocationService(config config.Config) *RevocationService {
	rs := RevocationService{}
	rs.blockchainRPCEndpoint = config.BlockchainRpcEndpoint
	rs.merkleTreeAcc = techniques.CreateMerkleTreeAccumulator(config)
	rs.bloomFilter = techniques.CreateBloomFilter(config.ExpectedNumberofRevokedVCs, config.FalsePositiveRate)
	//rs.bloomFilter = techniques.CreateBloomFilter(10000, config.FalsePositiveRate)
	rs.smartContractAddress= common.HexToAddress(config.SmartContractAddress)
	rs.privateKey = config.PrivateKeys[0]
	rs.gasLimit = config.GasLimit
	rs.gasPrice = config.GasPrice
	rs.VCToBigInts = make(map[string]*big.Int)
	rs.MtLevelInDLT = int(config.MtLevelInDLT)
	rs.mtHeight = int(config.MTHeight)
	rs.account = common.HexToAddress(config.SenderAddress)
	rs.NumberOfEntriesForMTInDLT = 0
	for i := 0; i <= rs.MtLevelInDLT; i++ {
		rs.NumberOfEntriesForMTInDLT += int(math.Pow(2, float64(i)))
	}
	rs.vcCounter = 0
	rs.isCached=false
	rs.isPublicKeysCached=false
	rs.cachedPublicKeys=make([][]byte,0)
	return &rs
}

func (r RevocationService) getAuth()  *bind.TransactOpts{
	// step 1: connect to a blockchain node using RPC endpoint
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(r.privateKey)
	if err != nil {
		zap.S().Fatalln("REVOCATION SERVICE: auth error: ",err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		zap.S().Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		zap.S().Fatalln(err)
	}


	gasLimit := uint64(r.gasLimit)                // in units
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	zap.S().Fatalln(err)
	//}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = r.gasPrice

	return auth
}

func SwapEndianness(b []byte) []byte {
	o := make([]byte, len(b))
	for i := range b {
		o[len(b)-1-i] = b[i]
	}
	return o
}

/*
Issues VC to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	vc - verifiable credential

Output:
	RevocationData
*/
func (r *RevocationService) IssueVC(vcID string) (*RevocationData) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()


	mtIndex, leafHash := r.merkleTreeAcc.AddLeaf(vcID)
	mtIndexes, mtValues := r.merkleTreeAcc.GetEntriesInLevelOrder(r.NumberOfEntriesForMTInDLT)

	var mtValuesInBytes [][32]byte
	for i:=0; i< len(mtValues);i++{
		h,_ := hex.DecodeString(mtValues[i])
		byteRepr := [32]byte{}
		copy(byteRepr[:], h[:])
		mtValuesInBytes = append(mtValuesInBytes, byteRepr)
	}


	_, err =revocationService.IssueVC(auth, mtIndexes, mtValuesInBytes)
	if err != nil {
		zap.S().Fatalln("failed to issue vc", err)
	}

	//generate bloom filter indexes for the vc and give it to the holders
	bfIndexes := r.bloomFilter.GetIndexes(vcID)
	merkleProof := r.merkleTreeAcc.GetProof(vcID)

	revocationData := CreateRevocationData(vcID, mtIndex, bfIndexes, leafHash, merkleProof)
	//revocationData.PrintRevocationData()

	return revocationData
}
/*
Issues VCs in bulk to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	vcIDS - []string

Output:
	RevocationData - []RevocationData
 */
func (r *RevocationService) IssueVCsInBulk(vcIDs []string) ([]*RevocationData, int64) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()
	var revocationDataALl []*RevocationData
	for _, vcID := range vcIDs {
		mtIndex, leafHash := r.merkleTreeAcc.AddLeaf(vcID)
		bfIndexes := r.bloomFilter.GetIndexes(vcID)
		merkleProof := r.merkleTreeAcc.GetProof(vcID)
		revocationData := CreateRevocationData(vcID, mtIndex, bfIndexes, leafHash, merkleProof)
		revocationDataALl = append(revocationDataALl, revocationData)
	}

	mtIndexes, mtValues := r.merkleTreeAcc.GetEntriesInLevelOrder(r.NumberOfEntriesForMTInDLT)

	var mtValuesInBytes [][32]byte
	for i:=0; i< len(mtValues);i++{
		h,_ := hex.DecodeString(mtValues[i])
		byteRepr := [32]byte{}
		copy(byteRepr[:], h[:])
		mtValuesInBytes = append(mtValuesInBytes, byteRepr)
	}

	startBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	_, err =revocationService.IssueVC(auth, mtIndexes, mtValuesInBytes)
	endBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	gasUsed := (startBalance.Int64()-endBalance.Int64()) / r.gasPrice.Int64()

	if err != nil {
		zap.S().Fatalln("failed to issue vc", err)
	}
	//mt := r.FetchMerkleTree()
	//zap.S().Infoln("REVOCATION SERVICE - merkle tree: ", mt)
	return revocationDataALl, gasUsed
}

func (r RevocationService) RetreiveUpdatedProof(vcID string)  *techniques.MerkleProof{
	merkleProof := r.merkleTreeAcc.GetProof(vcID)
	return merkleProof
}

func (r RevocationService) FindAncesstorInMerkleTree(index int)(int, string){
	currentLevel := r.mtHeight
	parentIndex := index
	for i:=currentLevel; i>r.MtLevelInDLT; i-- {
		temp := int(math.Floor(float64((parentIndex - 1) / 2)))
		parentIndex = temp
	}
	_, values := r.merkleTreeAcc.GetEntriesInLevelOrder(r.NumberOfEntriesForMTInDLT)
	ancesstor := values[parentIndex]
	return parentIndex, ancesstor
}

// returns old mt index and amount of gwei paid
func (r RevocationService) RevokeVC(vcID string) (int, int64, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()

	//Todo: retrieve the bloom filter indexes, merkle tree indexes and merkle tree values

	var bfIndexes []*big.Int
	for _, value := range r.bloomFilter.GetIndexes(vcID){
		bfIndexes = append(bfIndexes, value)
	}

	var mtIndexes []*big.Int
	var mtValues []string
	var parentIndex int

	vcIndex, _ := r.merkleTreeAcc.UpdateLeaf(vcID, "-1")

	currentLevel := r.mtHeight
	index := vcIndex
	for i:=currentLevel; i>=0; i--{
		parentIndex = int(math.Floor(float64((index - 1) / 2)))

		if i<=r.MtLevelInDLT{
			hashValue := r.merkleTreeAcc.Tree[index]
			mtIndexes = append(mtIndexes, big.NewInt(int64(index)))
			mtValues = append(mtValues, hashValue.Value)
		}
		index = parentIndex
	}

	var mtValuesInBytes [][32]byte
	for i:=0; i< len(mtValues);i++{
		h,_ := hex.DecodeString(mtValues[i])
		byteRepr := [32]byte{}
		copy(byteRepr[:], h[:])
		mtValuesInBytes = append(mtValuesInBytes, byteRepr)
	}


	startBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	_, err = revocationService.RevokeVC(auth, bfIndexes, mtIndexes, mtValuesInBytes)
	endBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	gasUsed := (startBalance.Int64()-endBalance.Int64()) / r.gasPrice.Int64()
	//gasUsed := (startBalance.Int64()-endBalance.Int64())/int64(math.Pow(10,9))
	//zap.S().Infoln("REVOCATION SERVICE- \t MT Accumulator levels in DLT: ",r.NumberOfEntriesForMTInDLT, "GAS USAGE in gwei: ", gasUsed)



	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}

	return vcIndex, gasUsed, nil
}

// returns old mt index and amount of gwei paid
func (r RevocationService) RevokeVCInBatches(vcIDs []string) (map[string]int, int64, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()

	//Todo: retrieve the bloom filter indexes, merkle tree indexes and merkle tree values

	var bfIndexes []*big.Int
	oldMTIndexes := make(map[string]int)
	mtTree := make(map[int]string)
	var mtIndexes []*big.Int
	var mtValues []string
	for i:=0; i<len(vcIDs);i++{
		for _, value := range r.bloomFilter.GetIndexes(vcIDs[i]) {
			bfIndexes = append(bfIndexes, value)
		}

		//oldMTIndex := r.VCToBigInts[vc.ID]
		vcIndex, _ := r.merkleTreeAcc.UpdateLeaf(vcIDs[i], "-1")
		oldMTIndexes[vcIDs[i]]=vcIndex

		currentLevel := r.mtHeight

		index := vcIndex
		for i:=currentLevel; i>=0; i--{
			parentIndex := int(math.Floor(float64((index - 1) / 2)))

			if i<=r.MtLevelInDLT{
				hashValue := r.merkleTreeAcc.Tree[index]
				mtTree[index] = hashValue.Value
			}
			index = parentIndex
		}

		//mtIndexes, mtValues = r.merkleTreeAcc.GetEntriesInLevelOrder(r.NumberOfEntriesForMTInDLT)
	}

	for index, value := range mtTree{
		mtIndexes = append(mtIndexes, big.NewInt(int64(index)))
		mtValues = append(mtValues, value)
	}

	var mtValuesInBytes [][32]byte
	for i:=0; i< len(mtValues);i++{
		h,_ := hex.DecodeString(mtValues[i])
		byteRepr := [32]byte{}
		copy(byteRepr[:], h[:])
		mtValuesInBytes = append(mtValuesInBytes, byteRepr)
	}



	//zap.S().Infoln("REVOCATION SERVICE- \t mt indexes: ", mtIndexes, "\t mt values: ",mtValues)
	//zap.S().Infoln("REVOCATION SERVICE- \t number of non-leaf nodes of MT accumulator stored in smart contract ",levelCounter)
	startBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	_, err = revocationService.RevokeVC(auth, bfIndexes, mtIndexes, mtValuesInBytes)
	endBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	//gasUsed := (startBalance.Int64()-endBalance.Int64())/int64(math.Pow(10,9))
	gasUsed := (startBalance.Int64()-endBalance.Int64()) / r.gasPrice.Int64()
	//zap.S().Infoln("REVOCATION SERVICE- \t bf indexes: ",bfIndexes, "\t mt indexes: ", mtIndexes, "\t mt values: ", mtValues, "\tGAS USAGE in gwei: ", gasUsed)



	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}


	return oldMTIndexes, gasUsed, nil
}

func (r RevocationService) VerificationPhase1(bfIndexes []*big.Int) (bool, error){
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	vcStatus, err := revocationService.VerificationPhase1(nil, bfIndexes)
	//zap.S().Errorln("REVOCATION SERVICE-  vc.IDverification phase 1: ",vcStatus)

	return vcStatus, err
}


func (r RevocationService) VerificationPhase2(leafHash string, witnesses []*techniques.Witness)(bool, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)


	mtRoot, err := revocationService.VerificationPhase2(nil)
	if err!=nil{
		zap.S().Errorln("REVOCATION SERVICE- error verification phase 2: ",err)
	}


	mtRootInHex := hex.EncodeToString(mtRoot[:])

	status := r.merkleTreeAcc.VerifyProof(leafHash, witnesses, mtRootInHex)

	//zap.S().Errorln("REVOCATION SERVICE-  verification phase 2: ",status)
	return status, nil
}



func (r RevocationService) GetMerkleRoot()(string, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	mtRoot, err := revocationService.VerificationPhase2(nil)
	mtRootInHex := hex.EncodeToString(mtRoot[:])
	if err!=nil{
		zap.S().Errorln("REVOCATION SERVICE- error verification phase 2: ",err)
	}

	zap.S().Errorln("REVOCATION SERVICE- merkle root: ",mtRootInHex)
	return mtRootInHex, nil
}

func (r RevocationService) FetchMerkleTree() ([]string){
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	mtSize := big.NewInt(int64(r.NumberOfEntriesForMTInDLT))
	mtValues , err := revocationService.RetrieveMerkleTree(nil,mtSize)

	var mtValuesInHex []string

	for i:=0; i< len(mtValues);i++{
		mtValuesInHex = append(mtValuesInHex, hex.EncodeToString(mtValues[i][:]))
	}
	if err!=nil{
		zap.S().Errorln("REVOCATION SERVICE- error retrieving merkle tree - ", err)
	}
	//zap.S().Errorln("REVOCATION SERVICE- merkle values: ",GetShortString(mtValues))
	return mtValuesInHex;
}

func (r RevocationService) VerifyVC( _bfIndexes []*big.Int, data *RevocationData) (bool, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}


	//Todo: this function should be moved to the verifiers. The parameters should be shared to the holders.
	status, err := revocationService.VerificationPhase1(nil, _bfIndexes)
	zap.S().Infoln("\n\n********************************************************************************************************************************")
	zap.S().Infoln("REVOCATION SERVICE- ", "***VERIFY*** vc:",data.VcId)
	zap.S().Errorln("REVOCATION SERVICE-  verification phase 1: ",status)
	if status==true{
		return status, err
	} else{
		mtRoot, err := revocationService.VerificationPhase2(nil)
		if err!=nil{
			zap.S().Errorln("REVOCATION SERVICE- error verification phase 2: ",err)
		}
		mtRootInHex := hex.EncodeToString(mtRoot[:])
		status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRootInHex)
		zap.S().Errorln("REVOCATION SERVICE-  verification phase 2: ",status)
	}
	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}
	zap.S().Infoln("\n\n********************************************************************************************************************************")

	return status, err
}

func (r RevocationService) PrintMerkleTree(){
	zap.S().Infoln("REVOCATION SERVICE - local merkle tree: ")
	r.merkleTreeAcc.PrintTree()

	zap.S().Infoln("REVOCATION SERVICE - merkle tree stored in dlt: ", GetShortString(r.FetchMerkleTree()))

}


func (r RevocationService) LocalMTVerification(mtRoot string, data *RevocationData) {

	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification - \t vc id: ",data.VcId, "\t root: ",mtRoot,
		"\t leaf value: ", data.MerkleTreeLeafValue, "\t proof: ",data.MerkleProof.OrderedWitnesses)
	r.merkleTreeAcc.PrintTree()
	status := r.merkleTreeAcc.VerifyProof(data.MerkleProof.LeafHash, data.MerkleProof.OrderedWitnesses, mtRoot)
	zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification : ", status)
	//statusLocal := r.merkleTreeAcc.VerifyProof(data.merkleTreeIndex, data.MerkleProof)
	//zap.S().Infoln("REVOCATION SERVICE- ", "\t local MT verification local : ", statusLocal)
}

/*
AddPublicKeys adds the entities's public keys to the smart contract

Input:
	public Keys - []string
 */
func (r RevocationService) AddPublicKeys(publicKeys [][]byte) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}

	auth := r.getAuth()

	//Todo: this function should be moved to the verifiers. The parameters should be shared to the holders.
	_, err = revocationService.AddPublicKeys(auth, publicKeys)
	if err != nil {
		zap.S().Infof("Error adding public keys: %v", err)
	}

}

/*
FetchPublicKeys retrieves the entities's public keys from the smart contract

Output:
	public Keys - []string
*/
func (r RevocationService) FetchPublicKeys()([][]byte) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}


	//Todo: this function should be moved to the verifiers. The parameters should be shared to the holders.
	publicKeys, err := revocationService.RetrievePublicKeys(nil)
	if err != nil {
		zap.S().Infof("Error adding public keys: %v", err)
	}

	return publicKeys
}

/*
FetchPublicKeys retrieves the entities's public keys from the smart contract

Output:
	public Keys - []string
*/
func (r *RevocationService) FetchPublicKeysCached()([][]byte) {
	if r.isPublicKeysCached==false{
		publicKeys := r.FetchPublicKeys()
		r.cachedPublicKeys = append(r.cachedPublicKeys, publicKeys...)
		r.isPublicKeysCached=true
	}

	return r.cachedPublicKeys
}
func  GetShortString(inputs []string) []string{

	var res []string
	for _, input := range inputs {
		if len(input) > 0 {
			input = input[:techniques.SHORT_STRING_SIZE] + ".."
		}
		output := fmt.Sprintf("%s",input)
		res = append(res, output )
	}
	return res
}


/*
FetchMerkleTreeSize retrieves the actual size of merkle tree stored in the smart contract

Output:
	merkle tree size (in bytes) - uint
*/
func (r RevocationService) FetchMerkleTreeSizeInDLT()(uint) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}


	//Todo: this function should be moved to the verifiers. The parameters should be shared to the holders.
	mtLength := big.NewInt(int64(r.NumberOfEntriesForMTInDLT))
	mtSize, err := revocationService.GetMerkleTreeSize(nil, mtLength)
	if err != nil {
		zap.S().Infof("Error adding public keys: %v", err)
	}

	return uint(mtSize.Uint64())
}

/*
FetchMerkleTreeSize retrieves the actual size of merkle tree stored local

Output:
	merkle tree size (in bytes) - uint
*/
func (r RevocationService) FetchMerkleTreeSizeLocal()(uint) {
	n := 0
	for i := 0; i <= r.mtHeight; i++ {
		n  += int(math.Pow(2, float64(i)))
	}
	_, mtValues := r.merkleTreeAcc.GetEntriesInLevelOrder(n)
	size := 0
	for _, value := range mtValues{
		size = size + int(uint(reflect.TypeOf(value).Size()))
	}
	return uint(size)
}



/*
FetchBloomFilterSizeInDLT calculates the bloom filter size stored in the smartcontract locally

In the smart contract, each bfIndex is assigned a 256-bit bucket.
Therefore, each bfIndex is divided by 256 to identify the bucket and then the corresponding
bit in the bucket is set. Each bucket size is 356 bits (32 bytes).

Here, this function calculates how many buckets are used to store the bf indexes corresponding to the
revoked vc ids.

BF Size: number of buckets * 32

Output:
	bloom filter size (in bytes) - uint
*/
func (r RevocationService) FetchBloomFilterSizeInDLT(revokedVcIDs []string)(uint) {

	var bfIndexes []uint
	for i:=0; i<len(revokedVcIDs);i++ {
		for _, value := range r.bloomFilter.GetIndexes(revokedVcIDs[i]) {
			bfIndexes = append(bfIndexes, uint(value.Uint64()))
		}
	}
	buckets := mapset.NewSet()

	for i:=0; i<len(bfIndexes);i++{
		bucket := bfIndexes[i] >> 8

		if buckets.Contains(bucket) {
			continue
		} else {
			buckets.Add(bucket)
		}
	}

	bfSize := buckets.Cardinality() * 32
	return uint(bfSize)

}
func (r RevocationService)  GetLocalBloomFilter() *techniques.BloomFilter{
	return r.bloomFilter
}

/*
RevocationCostCalculator calculates the revocation cost by calling the revoke function in the smart contract

Inputs:
1) bfIndexes - bloom filter indexes
2) mtIndexes - merkle tree indexes
3) mtValues - merkle tree values

Outputs:
1) gasUsed
2) error message
 */
func (r RevocationService) RevocationCostCalculator(bfIndexes []*big.Int, mtIndexes []*big.Int, mtValuesInBytes [][32]byte) (int64, error){
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()

	startBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	_, err = revocationService.RevokeVC(auth, bfIndexes, mtIndexes, mtValuesInBytes)
	endBalance, err := client.BalanceAt(context.Background(), r.account, nil)
	gasUsed := (startBalance.Int64()-endBalance.Int64()) / r.gasPrice.Int64()
	return gasUsed, err
}
