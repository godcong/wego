package wego

import (
	"github.com/godcong/wego/app/mini"
	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

//RegConfig config
const RegConfig = "config"

//RegClient client
const RegClient = "client"

//RegAccessToken access token
const RegAccessToken = "access_token"

var app *Application

func initSystem(config *core.Config) *System {
	var system System
	if !config.IsNil() {
		err := config.Unmarshal(&system)
		if err != nil {
			return &system
		}
	}
	return &System{
		Debug:    false,
		UseCache: false,
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

//NewApplication create an application instance with config.toml path
func NewApplication(path string) *Application {
	config, err := core.LoadConfig(path)
	if err != nil {
		panic(err)
	}

	return newApplication(config)
}

func newApplication(config *core.Config) *Application {
	app := &Application{
		objects: make(util.Map),
	}

	app.System = initSystem(config.GetSubConfig("system"))
	app.Register(RegConfig, config)
	return app
}

func init() {
	app = newApplication(core.DefaultConfig())
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
	if v, b := a.Get(RegClient); b {
		return v.(*core.Client)
	}
	client := core.DefaultClient()
	a.Register("client", client)
	return client
}

//Payment return a default Payment
func (a *Application) Payment(cfg string) *payment.Payment {
	return payment.NewPayment(a.Config().GetSubConfig(cfg))
}

//OfficialAccount return a default OfficialAccount
func (a *Application) OfficialAccount(cfg string) *official.Account {
	return official.NewOfficialAccount(a.Config().GetSubConfig(cfg))
}

//MiniProgram return a default MiniProgram
func (a *Application) MiniProgram(cfg string) *mini.Program {
	return mini.NewMiniProgram(a.Config().GetSubConfig(cfg))
}

// Config ...
func Config() *core.Config {
	return app.Config()
}

/*Get 获取注册的数据 */
func (a *Application) Get(name string) (interface{}, bool) {
	if v, b := a.objects[name]; b {
		return v, true
	}
	return nil, false
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

//New create an new instance
func (a *Application) New(name string, args ...interface{}) interface{} {
	return nil
}

/*App 获取App */
func App() *Application {
	log.Debug("app:", app)
	if app != nil {
		return app
	}
	return newApplication(core.DefaultConfig())
}

/*System 系统定义 */
type System struct {
	//debug = true
	Debug bool `toml:"debug"`
	//response_type = 'array'
	//ResponseType string `toml:"response_type"`
	//use_cache = true
	//DataType DataType `toml:"data_type"`
	UseSandbox bool `toml:"use_sandbox"`
	UseCache   bool `toml:"use_cache"`
	Log        log.Log
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
//	HttpGet(url string, m util.Map) core.Response
//	HttpPost(url string, m util.Map) core.Response
//	PostJSON(url string, m util.Map, query util.Map) core.Response
//	Request(url string, params util.Map, method string, options util.Map) core.Response
//	RequestRaw(url string, params util.Map, method string, options util.Map) core.Response
//	SafeRequest(url string, params util.Map, method string, options util.Map) core.Response
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
