package main

func main()  {
	//testAries()
	config,_ := parseConfig()
	//testConnectionToBlockchain(config)
	testSmartContract(config)

	//testIssuer(config)
}
