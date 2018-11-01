package official_test

import (
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
)

var config = Config()

// Config ...
func Config() *core.Config {
	log.Println("load")
	cfg, _ := core.LoadConfig("D:\\workspace\\project\\goproject\\wego\\config.toml")
	config := cfg.GetSubConfig("official_account.default")
	cache.Set("config", cfg)
	return config
}
