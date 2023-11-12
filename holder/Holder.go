package holder

import (
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/issuer"
)

type Holder struct {
	name string
	credentialStore map[string]verifiable.Credential
	revocationDataStore map[string]issuer.RevocationData
}

func (holder Holder) ReceiveVC(vc verifiable.Credential, revocationData issuer.RevocationData)  {

}
