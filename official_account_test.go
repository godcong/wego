package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestOfficialAccount(t *testing.T) {
	ip := wego.GetOfficialAccount().Base().GetCallbackIp()
	log.Println(ip)
}
