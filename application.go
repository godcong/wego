package wego

import (
	"github.com/godcong/wego/log"
	"reflect"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

const configPath = "config.toml"

//RegConfig config
const RegConfig = "config"

//RegClient client
const RegClient = "client"

//RegAccessToken access token
const RegAccessToken = "access_token"

var app *Application

func initSystem(cfg *core.Config) *System {
	var system System
	if cfg != nil {
		err := cfg.Unmarshal(&system)
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

//DefaultApplication result an default application
func DefaultApplication() *Application {
	return app
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

	config, err := core.LoadConfig(cpath)
	if err != nil {
		panic(err)
	}
	app.Register(RegConfig, config)

	//app.Register(RegClient, client)

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

//Config get application config interface
func (a *Application) Config() *core.Config {
	v, b := a.Get(RegConfig)
	if b {
		return v.(*core.Config)
	}
	return (*core.Config)(nil)
}

//Client get application client instance
func (a *Application) Client() *core.Client {
	v, b := a.Get(RegClient)
	if b {
		return v.(*core.Client)
	}
	return nil
}

//AccessToken get application access token instance
func (a *Application) AccessToken() *core.AccessToken {
	v, b := a.Get(RegAccessToken)
	if b {
		return v.(*core.AccessToken)
	}
	return nil
}

/*Get 获取注册的数据 */
func (a *Application) Get(name string) (interface{}, bool) {
	if v0, b := a.objects[name]; b {
		//return reflectSet(v, v0)
		return v0, true
	}
	return nil, false
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

//
///*Application 基础应用*/
//type Application interface {
//	Get(name string) interface{}
//	Register(name string, v interface{})
//	Scheme(id string) string
//	GetKey(s string) string
//	InSandbox() bool
//	SetSubMerchant(mchid, appid string) *core.Application
//}
//
///*AccessToken 访问Token */
//type AccessToken interface {
//	GetToken() core.Token
//	GetTokenWithRefresh() core.Token
//	GetRefreshedToken() core.Token
//	Refresh() *core.AccessToken
//}
//
///*Client 客户端*/
//type Client interface {
//	HttpClient() *http.Client
//	SetHttpClient(client *http.Client) Client
//	DataType() core.DataType
//	SetDataType(dataType core.DataType) Client
//	URL() string
//	SetDomain(domain *core.Domain) Client
//	HttpGet(url string, m util.Map) *net.Response
//	HttpPost(url string, m util.Map) *net.Response
//	HttpPostJson(url string, m util.Map, query util.Map) *net.Response
//	Request(url string, params util.Map, method string, options util.Map) *net.Response
//	RequestRaw(url string, params util.Map, method string, options util.Map) *net.Response
//	SafeRequest(url string, params util.Map, method string, options util.Map) *net.Response
//	Link(string) string
//}

///*Domain 域名*/
//type Domain interface {
//	URL() string
//	Link(s string) string
//}

///*GetApp 获取Application */
//func GetApp() Application {
//	return core.App()
//}
