# Tutorial - Verifiable Credentials

**Verifiable Credentials:**

The W3C compliance data models for Verifiable Credentials and Verifiable Presentations are implemented in the package "models".

The models package defines the Verifiable Credentials as a structure:&#x20;

{% code fullWidth="false" %}
```go
type VerifiableCredential struct{
	Metadata Metadata
	Claims   Claims
	Proofs   []Proof
}
```
{% endcode %}

The VeriflableCredential struct consists of the following three structs:

```go
type Metadata struct{
	Contexts         interface{}
	Id               URI
	Types            []URI
	Issuer           URI
	IssuanceDate     TimeString
	ExpirationDate   TimeString
	CredentialStatus CredentialStatus `json:"credentialStatus"`
}

type Claims interface {}

type Proof struct {
	Type string
	ProofValue []byte
}

type CredentialStatus struct{
	Id     URI         `json:"id"`
	Method interface{} `json:"method"`
	BfIndexes []string `json:"bfIndexes"`
	MTLeafValue string `json:"mtLeaf"`
}
```

CredentialStatus encodes Bloom Filter indexes and Merkle Tree leaf value.





**Creating Verifiable Credentials (e.g. Diploma):**

1. Import models package to use Verifiable Credentials

```go
import "github.com/Revocation-Service/models"
```

2. Create a struct for claims&#x20;

```go
type DiplomaClaim struct{
	Id          models.URI
	StudentName string
	StudentId   models.URI
	University  string
	Degree string
	GraduationYear int
	Grade string
}

```

3. Instantiate the claims

```go
myDiplomaClaims := DiplomaClaim{
		Id:             vcID,
		StudentName:    "praveen",
		StudentId:      vcID,
		University:     "University",
		Degree:         "Doctor of Philosophy",
		GraduationYear: 2000,
		Grade:          "A",
}
```

4. Create Diploma VC

```go
diploma := models.VerifiableCredential{}
diploma.Claims = myDiplomaClaims
```



**Encoding Bloom Filter Indexes and Merkle Tree Leaf values in the Verifiable Credentials:**

```go
CredentialStatus := models.CredentialStatus{
				Id:     <id>,
				Method: "2-phase revocation",
				BfIndexes: ["1", "2", "3"],
				MTLeafValue: "0x123456",
			}
```



