package official_test

import (
	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core"
	"testing"
)

var config *core.Config

func init() {
	cfg, _ := core.LoadConfig("D:\\workspace\\project\\goproject\\wego\\config.toml")
	config = cfg.GetSubConfig("official_account.default")
}

func TestBase_GetCallbackIp(t *testing.T) {
	base := official.NewBase(config)
	rlt := base.GetCallbackIP()
	t.Log(rlt)
}

func TestBase_ClearQuota(t *testing.T) {
	base := official.NewBase(config)
	rlt := base.ClearQuota()
	t.Log(rlt)
}
