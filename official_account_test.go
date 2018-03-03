package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
	"github.com/godcong/wego/official_account"
)

func TestOfficialAccount(t *testing.T) {
	o := wego.GetApp().Get("official_account").(*official_account.OfficialAccount)
	log.Println(o.GetValidIps())
	log.Println(o.ClearQuota())
	wego.GetO

}
