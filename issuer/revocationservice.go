package issuer

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/contracts"
	"github.com/praveensankar/Revocation-Service/techniques"
	"go.uber.org/zap"
	"math/big"
	"context"
	"crypto/ecdsa"
)

type IRevocationService interface {
	IssueVC( _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error)
	RevokeVC( _bfIndexes [3]*big.Int, _mtIndexes []*big.Int, _mtValues [][32]byte) (*types.Transaction, error)
	VerifyVC( _bfIndexes [3]*big.Int, vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error)
}


type RevocationService struct{
	merkleTreeAcc *techniques.MerkleTreeAccumulator
	bloomFilter *techniques.BloomFilter
	blockchainRPCEndpoint string
	smartContractAddress common.Address
	privateKey string
	gasLimit uint64
}



func CreateRevocationService(config config.Config) *RevocationService{
	rs := RevocationService{}
	rs.blockchainRPCEndpoint = config.BlockchainRpcEndpoint
	rs.merkleTreeAcc = techniques.CreateMerkleTree()
	rs.bloomFilter = techniques.CreateBloomFilter(config.ExpectedNumberOfTotalVCs, config.FalsePositiveRate)
	rs.smartContractAddress= common.HexToAddress(config.SmartContractAddress)
	rs.privateKey = config.PrivateKey
	rs.gasLimit = config.GasLimit
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
		zap.S().Fatalln(err)
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
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		zap.S().Fatalln(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth
}

/*
Issues VC to holder. and updates the merkle tree both locally and in smart contract.

Inputs:
	_mtIndexes: merkle tree indexes
	_mtValues: merkle tree values
 */
func (r RevocationService) IssueVC(vc verifiable.Credential) (*types.Transaction, error) {
	client, err := ethclient.Dial(r.blockchainRPCEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}
	revocationService, err := contracts.NewRevocationService(r.smartContractAddress, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}
	auth := r.getAuth()

	//Todo: add merkle tree to the issuer or the revocation service.
	//Todo: generate the merkle tree indexes and values and send them to smart contract
	var _mtIndexes []*big.Int
	var _mtValues [][32]byte
	tx, err:=revocationService.IssueVC(auth, _mtIndexes, _mtValues)
	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}
	zap.S().Infof("tx hash: %s\n", tx.Hash().Hex())
	return tx, nil

}

func (r RevocationService) RevokeVC(vc verifiable.Credential) (*types.Transaction, error) {
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
	var _bfIndexes [3]*big.Int
	var _mtIndexes []*big.Int
	var _mtValues [][32]byte
	tx, err := revocationService.RevokeVC(auth, _bfIndexes, _mtIndexes, _mtValues)
	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}
	zap.S().Infof("tx hash: %s\n", tx.Hash().Hex())
	return tx, nil
}

func (r RevocationService) VerifyVC( _bfIndexes [3]*big.Int, vcLeaf [32]byte, proof [][32]byte) (*types.Transaction, error) {
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
	tx, err := revocationService.VerifyVC(auth, _bfIndexes, vcLeaf, proof)
	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}
	zap.S().Infof("tx hash: %s\n", tx.Hash().Hex())
	return tx, nil
}


