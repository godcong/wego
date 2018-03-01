package core

import (
	"flag"

	"github.com/godcong/wego/cache"
)

type Application struct {
	Config
	Client
	obj Map
}

var f = flag.String("f", "config.toml", "config file path")
var app *Application

func initLoader() *Tree {
	t := ConfigTree(*f)
	initSystem(t.GetTree("system"))
	return t
}

func newApplication(v ...interface{}) *Application {
	app := &Application{
		obj: Map{},
	}
	for _, value := range v {
		switch v := value.(type) {
		case Config:
			app.Register("config", v)
			app.Config = v
		}
	}
	if app.Get("config") == nil {
		app.Register("config", GetRootConfig())
		app.Config = GetRootConfig()
	}
	return app
}

func initApp(config Config) *Application {
	if app == nil {
		app = newApplication(config)
	}
	return app
}

func init() {
	c := cache.GetCache()
	flag.Parse()
	config := initLoader()
	if system.UseCache {
		CacheOn()
		c.Set("cache", config)
	}
	initLog(system)
	initApp(config)
}

func (a *Application) Get(name string) interface{} {
	if v, b := (*a).obj[name]; b {
		return v
	}
	return nil
}

func (a *Application) Register(name string, v interface{}) {
	(*a).obj[name] = v
}

func App() *Application {
	Debug("app:", app)
	return app
}

func (a *Application) InSandbox() bool {
	c := a.Get("config").(Config)
	return c.GetBool("sandbox")
}

func (a *Application) GetKey(s string) string {
	c := a.Get("sandbox").(*Sandbox)
	if a.InSandbox() {
		c.GetKey()
	}
	return c.Get("aes_key")

}

func (a *Application) Scheme(id string) string {
	c := a.Get("config").(Config)
	m := make(Map)
	m.Set("appid", c.Get("app_id"))
	m.Set("mch_id", c.Get("mch_id"))
	m.Set("time_stamp", Time(nil))
	m.Set("nonce_str", GenerateNonceStr())
	m.Set("product_id", id)
	m.Set("sign", GenerateSignature(m, c.Get("aes_key"), SIGN_TYPE_MD5))
	return BIZPAYURL + m.UrlEncode()
}

func (a *Application) HandleNotify(typ string, f func(interface{})) {

}

func (a *Application) SetSubMerchant(mchid, appid string) *Application {
	c := a.Get("config").(Config)
	c.Set("sub_mch_id", mchid)
	c.Set("sub_appid", appid)
	return a
}
