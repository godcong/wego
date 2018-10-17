package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/pelletier/go-toml"
	"reflect"
)

const configPath = "config.toml"

var app *Application

func initSystem(v interface{}) *System {
	var system System
	if v != nil {
		err := (*toml.Tree)(v.(*Tree)).Unmarshal(&system)
		if err != nil {
			return &system
		}
	}
	return &System{
		Debug:        false,
		ResponseType: "array",
		UseCache:     false,
		Log: log.Log{
			Level: "debug",
			File:  "logs/wechat.log",
		},
	}

}

//NewApplication create an application instance
func NewApplication() *Application {
	return newApplication(configPath)
}

func newApplication(cpath string) *Application {
	app := &Application{
		objects: make(util.Map),
	}

	config, err := LoadConfig(cpath)
	if err != nil {
		panic(err)
	}
	app.Register("config", config)
	app.System = initSystem(config.GetSubConfig("system"))

	return app
}

//DefaultApplication returns an application with default config
func DefaultApplication() *Application {
	return initApp(configPath)
}

func initApp(cpath string) *Application {
	if app == nil {
		app = newApplication(cpath)
	}
	return app
}

func init() {
	initApp(configPath)
}

//GetConfig get application config interface
func (a *Application) GetConfig() Config {
	config := Tree{}
	b := a.Get("config", &config)
	if b {
		return &config
	}
	return Config(nil)
}

/*Get 获取注册的数据 */
func (a *Application) Get(name string, v interface{}) bool {
	if v0, b := a.objects[name]; b {
		return reflectSet(v, v0)
	}
	return false
}

func reflectSet(tar, src interface{}) bool {
	if src != nil && tar != nil {
		reflect.ValueOf(tar).Elem().Set(reflect.ValueOf(src).Elem())
		return true
	}
	return false
}

/*GetInterface 获取注册的interface */
func (a *Application) GetInterface(name string) (interface{}, bool) {
	if v0, b := a.objects[name]; b {
		return v0, true
	}
	return nil, false
}

/*Register 注册 */
func (a *Application) Register(name string, v interface{}) {
	a.objects[name] = v
}

/*App 获取App */
func App() *Application {
	log.Debug("app:", app)
	return app
}

/*InSandbox 是否沙箱环境 */
func (a *Application) InSandbox() bool {
	//c := a.Get("config").(Config)
	//return a.GetBool("payment.default.sandbox")

	return false
}

/*GetKey 获取沙箱key */
func (a *Application) GetKey(s string) string {
	sb := Sandbox{}
	b := a.Get("sandbox", &sb)
	if b && a.InSandbox() {
		sb.GetKey()
	}
	return sb.Get("aes_key")

}

/*Scheme 获取微信Scheme */
//func (a *Application) Scheme(id string) string {
//	//c := a.Config
//	m := make(util.Map)
//	m.Set("appid", a.Config.Get("app_id"))
//	m.Set("mch_id", a.Config.Get("mch_id"))
//	m.Set("time_stamp", util.Time())
//	m.Set("nonce_str", util.GenerateNonceStr())
//	m.Set("product_id", id)
//	m.Set("sign", GenerateSignature(m, a.Config.Get("aes_key"), MakeSignMD5))
//	return BizPayURL + m.URLEncode()
//}

//func (a *Application) HandleNotify(typ string, f func(interface{})) {
//
//}

/*SetSubMerchant 设置子商户id */
//func (a *Application) SetSubMerchant(mchid, appid string) *Application {
//	a.Config.Set("sub_mch_id", mchid)
//	a.Config.Set("sub_appid", appid)
//	return a
//}

/*CacheOn turn on cache */
func (s *System) CacheOn() {
	s.UseCache = true
}

/*CacheOff turn off cache */
func (s *System) CacheOff() {
	s.UseCache = false
}

/*CacheStatus return cache status */
func (s *System) CacheStatus() bool {
	return s.UseCache
}

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

/*Application Application */
type Application struct {
	*System
	//Client
	objects util.Map
}
