package entities

import (
	"github.com/praveensankar/Revocation-Service/common"
	"net"
	"sync"
)

type  Manager struct {
	sync.RWMutex
	name string
	Type Entity
	Conn net.Conn
	issuerAddress string
	verifierAddress string
	holderAddress string
	Results []common.Results
}

