package entities

import "net"

type Entity string
const (
	ISSUER Entity = "issuer"
	VERIFIER        = "verifier"
	HOLDER      = "holder"
	REVOCATIONSERVICE ="revocation service"
)

type IEntity interface {
	GetType() Entity
	GetConnection() net.Conn
}
