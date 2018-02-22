package wego

import (
	"flag"

	"github.com/godcong/wego/cache"
	"github.com/pelletier/go-toml"
)

type Application interface {
	Payment() Payment
	MiniProgram() MiniProgram
	Cache() cache.Cache
	Client(config Config) Client
	GetConfig(s string) Config
	Scheme(id string) string
	GetKey(s string) string
	InSandbox() bool
	SetSubMerchant(mchid, appid string) Application
}

type application struct {
	cache   cache.Cache
	config  Config
	sandbox Sandbox
	payment Payment
	mini    MiniProgram
	order   Order
	request *Request
}

var f = flag.String("f", "config.toml", "config file path")

var system System

var app *application

func init() {
	c := cache.GetCache()
	flag.Parse()
	c.Set("cache_path", *f)
	config := initLoader()
	if system.UseCache {
		CacheOn()
		c.Set("cache", config)
	}
	initLog(system)
	initApp(config)
}

func initLoader() *Tree {
	t := ConfigTree(*f)
	t.GetTree("system").(*toml.Tree).Unmarshal(&system)
	return t
}

func initApp(config Config) Application {
	if app == nil {
		app = newApplication(config)
	}
	return app
}

func GetApplication() Application {
	return app
}

func GetSecurity() Security {
	return app.Payment().Security()
}

func GetOrder() Order {
	return app.Payment().Order()
}

func GetRefund() Refund {
	return app.Payment().Refund()
}

func GetBill() Bill {
	return app.Payment().Bill()
}

func (a *application) MiniProgram() MiniProgram {
	if a.mini == nil {
		a.mini = NewMiniProgram(a)
	}
	return a.mini
}

func (a *application) Cache() cache.Cache {
	return cache.GetCache()
}

func (a *application) Payment() Payment {
	if a.payment == nil {
		a.payment = NewPayment(a)
	}
	return a.payment
}

func (a *application) Request() *Request {
	if a.request == nil {
		a.request = NewRequest()
	}
	return a.request
}

func (a *application) Client(config Config) Client {
	return NewClient(a, config, a.Request())
}

func (a *application) SetSubMerchant(mchid, appid string) Application {
	a.config.Set("sub_mch_id", mchid)
	a.config.Set("sub_appid", appid)
	return a
}

func newApplication(v ...interface{}) *application {
	app := &application{}
	for _, value := range v {
		switch value.(type) {
		case Config:
			app.config = (value).(Config)
		}
	}
	if app.config == nil {
		app.config = GetRootConfig()
	}
	return app
}

func NewApplication(v ...interface{}) Application {
	return newApplication(v)
}

func (a *application) GetConfig(s string) Config {
	if s == "" {
		return a.config
	}
	return a.config.GetConfig(s)
}

func (a *application) InSandbox() bool {
	return a.config.GetBool("sandbox")
}

func (a *application) GetKey(s string) string {
	if a.InSandbox() {
		a.sandbox.GetKey()
	}
	return a.config.Get("aes_key")
}

func (a *application) Scheme(id string) string {
	m := make(Map)
	m.Set("appid", a.config.Get("app_id"))
	m.Set("mch_id", a.config.Get("mch_id"))
	m.Set("time_stamp", Time(nil))
	m.Set("nonce_str", GenerateNonceStr())
	m.Set("product_id", id)
	m.Set("sign", GenerateSignature(m, a.config.Get("aes_key"), SIGN_TYPE_MD5))
	return BIZPAYURL + m.ToUrlQuery()
}

func (a *application) HandleNotify(typ string, f func(interface{})) {

}
