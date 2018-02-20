package wego

import (
	"log"
	"testing"
)

func TestSecurity_GetPublicKey(t *testing.T) {
	m := GetSecurity().GetPublicKey()
	log.Println(m.Get("pub_key"))
}
