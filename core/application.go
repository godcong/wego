package core

import (
	"flag"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	toml "github.com/pelletier/go-toml"
)

/*Application Application */
type Application struct {
	config.Config
	Client
	obj util.Map
}

var f = flag.String("f", "config.toml", "config file path")
var app *Application
var defaultConfig config.Config

/*System 系统定义 */
type System struct {
	//debug = true
	Debug bool `toml:"debug"`
	//response_type = 'array'
	ResponseType string `toml:"response_type"`
	//use_cache = true
	//DataType DataType `toml:"data_type"`

	UseCache bool `toml:"use_cache"`
	Log      log.Log
}

var system System
var useCache = false

func initLoader() *config.Tree {
	t, err := config.GetConfigTree(*f)
	if err != nil {
		t = &config.Tree{}
	}
	initSystem(t.GetTree("system"))
	return t
}

func initSystem(v interface{}) {
	if v == nil {
		system = System{
			Debug:        false,
			ResponseType: "array",
			UseCache:     false,
			Log: log.Log{
				Level: "debug",
				File:  "logs/wechat.log",
			},
		}
		return
	}
	v.(*toml.Tree).Unmarshal(&system)
	return
}

func newApplication(v ...interface{}) *Application {
	app := &Application{
		obj: util.Map{},
	}
	for _, value := range v {
		switch v := value.(type) {
		case config.Config:
			app.Register("config", v)
			app.Config = v
		}
	}
	if app.Get("config") == nil {
		app.Config = config.GetRootConfig()
		app.Register("config", app.Config)
	}
	return app
}

func initApp(config config.Config) *Application {
	if app == nil {
		app = newApplication(config)
	}
	return app
}

func init() {
	//c := cache.GetCache()
	flag.Parse()
	defaultConfig = initLoader()
	if !system.UseCache {
		config.CacheOff()
		//c.Set("cache", defaultConfig)
	}
	log.InitLog(system.Log, system.Debug)
	initApp(defaultConfig)
}

/*GetSystemConfig GetSystemConfig */
func GetSystemConfig() System {
	return system
}

/*Get 获取注册的interface */
func (a *Application) Get(name string) interface{} {
	if v, b := (*a).obj[name]; b {
		return v
	}
	return nil
}

/*Register 注册 */
func (a *Application) Register(name string, v interface{}) {
	a.obj[name] = v
}

/*App 获取App */
func App() *Application {
	log.Debug("app:", app)
	return app
}

/*InSandbox 是否沙箱环境 */
func (a *Application) InSandbox() bool {
	//c := a.Get("config").(Config)
	return a.GetBool("payment.default.sandbox")
}

/*GetKey 获取沙箱key */
func (a *Application) GetKey(s string) string {
	b := a.Get("sandbox").(*Sandbox)
	if a.InSandbox() {
		b.GetKey()
	}
	return b.Get("aes_key")

}

/*Scheme 获取微信Scheme */
func (a *Application) Scheme(id string) string {
	//c := a.Config
	m := make(util.Map)
	m.Set("appid", a.Config.Get("app_id"))
	m.Set("mch_id", a.Config.Get("mch_id"))
	m.Set("time_stamp", util.Time(nil))
	m.Set("nonce_str", util.GenerateNonceStr())
	m.Set("product_id", id)
	m.Set("sign", GenerateSignature(m, a.Config.Get("aes_key"), MakeSignMD5))
	return BizPayURL + m.URLEncode()
}

//func (a *Application) HandleNotify(typ string, f func(interface{})) {
//
//}

/*SetSubMerchant 设置子商户id */
func (a *Application) SetSubMerchant(mchid, appid string) *Application {
	a.Config.Set("sub_mch_id", mchid)
	a.Config.Set("sub_appid", appid)
	return a
}
