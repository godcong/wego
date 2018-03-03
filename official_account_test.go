package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
	"github.com/godcong/wego/official"
)

func TestOfficialAccount(t *testing.T) {
	o := wego.GetApp().Get("official_account").(*official.OfficialAccount)
	log.Println(ip)
	wego.GetO

}
