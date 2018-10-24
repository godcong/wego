package official_test

import (
	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"testing"
)

var config *core.Config

func init() {
	log.Println("load")
	cfg, _ := core.LoadConfig("D:\\workspace\\project\\goproject\\wego\\config.toml")
	config = cfg.GetSubConfig("official_account.default")
	cache.Set("config", cfg)
}

func TestBase_GetCallbackIp(t *testing.T) {
	base := official.NewBase(config)
	rlt := base.GetCallbackIP()
	t.Log(string(rlt.Bytes()))
	t.Log(rlt.ToMap())
}

func TestBase_ClearQuota(t *testing.T) {
	base := official.NewBase(config)
	rlt := base.ClearQuota()
	t.Log(string(rlt.Bytes()))
	t.Log(rlt.ToMap())
}
