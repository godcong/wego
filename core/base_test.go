package core_test

import (
	"github.com/godcong/wego/core"
	"testing"
)

var config *core.Config

func init() {
	cfg, _ := core.LoadConfig("D:\\workspace\\project\\goproject\\wego\\config.toml")
	config = cfg.GetSubConfig("official_account.default")
}

func TestBase_GetCallbackIP(t *testing.T) {
	base := core.NewBase(config)
	resp := base.GetCallbackIP()

	t.Log(resp.Error())
	t.Log(resp.ToMap())
	t.Log(string(resp.Bytes()))

}

func TestURL_ShortURL(t *testing.T) {
	resp := core.NewURL(config).ShortURL("weixin://wxpay/bizpayurl?pr=etxB4DY")
	t.Log(resp.Error())
	t.Log(resp.ToMap())
	t.Log(string(resp.Bytes()))
}
