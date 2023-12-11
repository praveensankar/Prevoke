package blockchain

import (
	"github.com/praveensankar/Revocation-Service/config"
	"sync"
)

func TestConnectionToBlockchain(config config.Config){

	// set up blockchain node
	node := BlockchainNode{rpcEndPoint: config.BlockchainRpcEndpoint}

	// get balance
	node.getBalance(config.SmartContractAddress)

	node.queryBlock()
	node.doTransaction(config.PrivateKey)
	//node.sendEther(config.privateKey,1, config.otherAccounts[0])
	//node.sendTransaction()
	node.sendLegacyTransaction(config.PrivateKey, config.OtherAccounts[0])
	_ = createNewKeyStore()


}



func TestSmartContract(my_config config.Config){
	ReadFromContract(my_config)



	var waitGroupforBlocksListener sync.WaitGroup

	waitGroupforBlocksListener.Add(1)
	go func(my_config config.Config) {
		defer waitGroupforBlocksListener.Done()
		SubscribeToEvents(my_config)
	}(my_config)


	waitGroupforBlocksListener.Wait()
}

func TestDeployment(conf config.Config){
	//DeployContract(conf)
}

func Test(conf config.Config){
	//TestDeployment(conf)
	TestSmartContract(conf)
}