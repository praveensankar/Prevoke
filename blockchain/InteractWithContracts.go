package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/praveensankar/Revocation-Service/config"
	"github.com/praveensankar/Revocation-Service/contracts"
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

	revocationService, err := contracts.NewRevocationService(revocationServiceSmartContract, ethClient)
	if err != nil {
		zap.S().Infof("Failed to instantiate Storage contract: %v", err)
	}

	var status bool
	var field int64 = 30
	value, err := revocationService.BloomFilter(nil, big.NewInt(field))
	if value.Int64()==1{
		status=true
	}else {
		status = false
	}
	zap.S().Infof("\n status at field %d is %t", field, status)


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

