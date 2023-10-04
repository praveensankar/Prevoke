package blockchain

import "github.com/praveensankar/Revocation-Service/config"

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