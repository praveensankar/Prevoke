package contracts

import (
	"github.com/praveensankar/Revocation-Service/config"
	"sync"
)

func TestSmartContract(my_config config.Config){
	ReadFromContract(my_config)
	WriteToContract(my_config)


	var waitGroupforBlocksListener sync.WaitGroup

	waitGroupforBlocksListener.Add(1)
	go func(my_config config.Config) {
		defer waitGroupforBlocksListener.Done()
		SubscribeToEvents(my_config)
	}(my_config)


	waitGroupforBlocksListener.Wait()
}
