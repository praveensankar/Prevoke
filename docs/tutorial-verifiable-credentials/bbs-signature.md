# BBS Signature

**Prevoke** supports selective disclosure of the vc ID, one or more claims, Bloom Filter indexes, and Merkle Leaf value. It means that each Holder can decide to seletively disclose any of the above with the Verifiers. Hence, Holders never share the complete Verifiable Credentials with the Verifiers.



The following packages need to be imported:

```go
	"github.com/Revocation-Service/signature"
	"github.com/suutaku/go-bbs/pkg/bbs"
```



1. Organize the elements to be selectively disclosed and generate messages in the organized order. The following function generates messages in a specific order.&#x20;

<pre class="language-go"><code class="lang-go">/*
The order of messages for selective disclosure
bf index1, ..., bf indexn, mtLeafHash, vcId, claims
 */
<strong>func generateMessages( vcId string, claims DiplomaClaim,
</strong>	bfIndexes []string, mtLeafHash string) [][]byte{
	var messages [][]byte

	// 1) append the bf indexes to the messages
	for i:=0;i&#x3C; len(bfIndexes);i++ {
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
</code></pre>

2. Sign the ordered messages using BBS Signature and encode them as a proof.

```
bbsKeys := signature.GenerateKeyPair()
privateKey := bbsKeys.PrivateKey
publicKey, _ := bbsKeys.PublicKey.Marshal()

signature  := signature.Sign(privateKey, messages)
proof := models.Proof{Type: "bbs+", ProofValue: signature}
```

