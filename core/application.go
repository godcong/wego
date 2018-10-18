package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/pelletier/go-toml"
	"reflect"
)

const configPath = "config.toml"

//RegConfig config
const RegConfig = "config"

//RegClient client
const RegClient = "client"

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
func NewApplication(s ...string) *Application {
	if s != nil {
		return newApplication(s[0])
	}
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
	app.Register(RegConfig, config)

	app.Register(RegClient, client)

	app.System = initSystem(config.GetSubConfig("system"))

	return app
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
	b := a.Get(RegConfig, &config)
	if b {
		return &config
	}
	return Config(nil)
}

//GetClient get application client instance
func (a *Application) GetClient() *Client {
	client := Client{}
	b := a.Get(RegClient, &client)
	if b {
		return &client
	}
	return nil
}

//GetAccessToken get application access token instance
func (a *Application) GetAccessToken() *AccessToken {
	token := AccessToken{}
	b := a.Get(RegClient, &token)
	if b {
		return &token
	}
	return nil
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
	return initApp(configPath)
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
func (a *Application) Scheme(id string) string {
	cfg := a.GetConfig().GetSubConfig("official_account.default") //TODO: get used config
	m := make(util.Map)
	m.Set("appid", cfg.Get("app_id"))
	m.Set("mch_id", cfg.Get("mch_id"))
	m.Set("time_stamp", util.Time())
	m.Set("nonce_str", util.GenerateNonceStr())
	m.Set("product_id", id)
	m.Set("sign", GenerateSignature(m, cfg.Get("aes_key"), MakeSignMD5))
	return BizPayURL + m.URLEncode()
}

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
