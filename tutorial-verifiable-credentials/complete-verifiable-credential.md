# Complete Verifiable Credential

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

```go
func generateMessages( vcId string, claims DiplomaClaim,
	bfIndexes []string, mtLeafHash string) [][]byte{
	var messages [][]byte

	// 1) append the bf indexes to the messages
	for i:=0;i< len(bfIndexes);i++ {
		messages = append(messages, []byte(bfIndexes[i]))
		//zap.S().Infoln("BBS - bf index: ", []byte(bfIndexes[i]))
	}
	// 2) append the mt leaf hash to the messages
	messages = append(messages, []byte(mtLeafHash))
	//zap.S().Infoln("BBS - mt leaf hash: ", []byte(mtLeafHash))

	// 2) append vc id to the messages
	messages = append(messages, []byte(vcId))
	//zap.S().Infoln("BBS - vc  id: ", []byte(vcId))

	// 4) append the claims. All the claims are appended together. However, it is trivial to implement
	messages = append(messages, []byte(fmt.Sprintf("%s", claims.Grade)))
	//zap.S().Infoln("BBS - grade: ",[]byte(fmt.Sprintf("%s", claims.Grade)))
	messages = append(messages, []byte(claims.Degree))
	messages = append(messages, []byte(fmt.Sprintf("%v", claims.Id)))
	messages = append(messages, []byte(fmt.Sprintf("%v", claims.StudentId)))

	messages = append(messages, []byte(claims.University))
	messages = append(messages, []byte(claims.StudentName))
	messages = append(messages, []byte(fmt.Sprintf("%d", claims.GraduationYear)))
	return messages
}

bbsKeys := signature.GenerateKeyPair()
privateKey := bbsKeys.PrivateKey
publicKey, _ := bbsKeys.PublicKey.Marshal()

signature  := signature.Sign(privateKey, messages)
proof := models.Proof{Type: "bbs+", ProofValue: signature}
var proofs []models.Proof
proofs = append(proofs, proof)
```

```go
myDiploma := models.VerifiableCredential{
		Metadata: models.Metadata{
			Contexts:       []string{"http:test.com/diploma"},
			Id:             vcId,
			Types:          []models.URI{"Diploma"},
			Issuer:         "university of oslo",
			IssuanceDate:   models.TimeString(time.Now()),
			ExpirationDate: models.TimeString(time.Now().Add(1000*time.Hour)),
			CredentialStatus: models.CredentialStatus{
				Id:     diplomaClaims.Id,
				Method: "2-phase revocation",
				BfIndexes: bfIndexes[:],
				MTLeafValue: mtLeafHash,
			},
		},
		Claims: diplomaClaims,
		Proofs: proofs,
	}
```
