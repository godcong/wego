package crypt

import (
	"encoding/base64"
	"testing"
)

func TestPrpCrypt_Encrypt(t *testing.T) {
	k, _ := base64.RawStdEncoding.DecodeString(encodingAesKey)
	prp := NewPrp([]byte(k))
	b, e := prp.Encrypt(text, appID)
	t.Log(string(b), e)

	b, e = prp.Decrypt(b, appID)
	t.Log(string(b), e)

	//t.Log(Base64Decode([]byte("TNwHN28RXXoyVxkMCUEqKuCL08eBpCKgWZTkWNVnGLu")))
	//v, _ := base64.RawStdEncoding.DecodeString("TNwHN28RXXoyVxkMCUEqKuCL08eBpCKgWZTkWNVnGLu")
	//t.Log(string(v), len(v))
}
