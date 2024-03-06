package entities

import "net"

type Entity string
const (
	ISSUER Entity = "issuer"
	VERIFIER        = "verifier"
	HOLDER      = "holder"
	REVOCATIONSERVICE ="revocation service"
	MANGER = "experiment manager"
)

type IEntity interface {
	GetType() Entity
	GetConnection() net.Conn
}
