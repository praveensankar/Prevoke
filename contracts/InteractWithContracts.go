package contracts

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/praveensankar/Revocation-Service/config"
	"go.uber.org/zap"
	"math/big"
	"time"
)


const key = `<<json object from keystore>>`



func ReadFromContract(config config.Config) {

	// step 1: connect to a blockchain node using RPC endpoint
	ethClient, err := ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		zap.S().Infof("Failed to connect to the Ethereum client: %v", err)
	}

	// step 2: convert smart contract to the requied format
	revocationServiceSmartContract := common.HexToAddress(config.SmartContractAddress)

	balance, err := ethClient.BalanceAt(context.Background(), revocationServiceSmartContract, nil)

	zap.S().Infof("%s has %d balance", revocationServiceSmartContract.String(), balance)

	revocationService, err := NewRevocationService(revocationServiceSmartContract, ethClient)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}

	var status bool
	var field int64 = 30
	status, err = revocationService.BloomFilter(nil, big.NewInt(field))
	zap.S().Infof("\n status at field %d is %t", field, status)


}


func WriteToContract(config config.Config){
	client, err :=  ethclient.Dial(config.BlockchainRpcEndpoint)
	if err != nil {
		zap.S().Fatalln(err)
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
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


	gasLimit := uint64(6721975)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		zap.S().Fatalln(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	revocationServiceSmartContract := common.HexToAddress(config.SmartContractAddress)
	revocationService, err := NewRevocationService(revocationServiceSmartContract, client)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}


	index1 := big.NewInt(515)
	index2 := big.NewInt(625)
	index3 := big.NewInt(735)

	indexes := [3]*big.Int{index1, index2, index3}

	status, err := revocationService.CheckRevocationStatusInBloomFilter(nil, indexes)
	zap.S().Infof("\n revocation status[%d, %d, %d]: %t", index1, index2, index3, status)

	zap.S().Infoln("\n Revoking in Bloom Filter")
	tx, err :=revocationService.RevokeInBloomFilter(auth, indexes)
	if err != nil {
		zap.S().Fatalln("failed to revoke", err)
	}
	zap.S().Infof("tx hash: %s\n", tx.Hash().Hex())

	status, err = revocationService.CheckRevocationStatusInBloomFilter(nil, indexes)
	zap.S().Infof("\n revocation status[%d, %d, %d]: %t", index1, index2, index3, status)

}


func SubscribeToEvents(config config.Config){
	client, err :=  ethclient.Dial(config.BlockchainWebSocketEndPoint)
	if err != nil {
		zap.S().Fatalln(err)
	}

	revocationServiceSmartContract := common.HexToAddress(config.SmartContractAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{revocationServiceSmartContract},
	}

	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		zap.S().Fatalln(err)
	}

	stopListener := time.NewTimer(5 * time.Second)
	for {
		select {
		case err := <-sub.Err():
			zap.S().Fatalln(err)
		case <-stopListener.C:
			zap.S().Infoln("\n stopping the listener")
			return
		case vLog := <-logs:
			zap.S().Infoln("new event ",vLog) // pointer to event log
		}
	}
}

