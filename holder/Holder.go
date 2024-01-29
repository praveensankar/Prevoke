package holder

import (
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/praveensankar/Revocation-Service/revocation_service"
)

type Holder struct {
	name string
	credentialStore map[string]verifiable.Credential
	revocationDataStore map[string]revocation_service.RevocationData
}

func (holder Holder) ReceiveVC(vc verifiable.Credential, revocationData revocation_service.RevocationData)  {

}
