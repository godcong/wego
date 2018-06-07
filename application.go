package wego

import (
	"net/http"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

/*Application 基础应用*/
type Application interface {
	Get(name string) interface{}
	Register(name string, v interface{})
	Scheme(id string) string
	GetKey(s string) string
	InSandbox() bool
	SetSubMerchant(mchid, appid string) *core.Application
}

/*AccessToken 访问Token */
type AccessToken interface {
	GetToken() core.Token
	GetTokenWithRefresh() core.Token
	GetRefreshedToken() core.Token
	Refresh() *core.AccessToken
}

/*Client 客户端*/
type Client interface {
	HttpClient() *http.Client
	SetHttpClient(client *http.Client) Client
	DataType() core.DataType
	SetDataType(dataType core.DataType) Client
	URL() string
	SetDomain(domain *core.Domain) Client
	HttpGet(url string, m util.Map) *net.Response
	HttpPost(url string, m util.Map) *net.Response
	HttpPostJson(url string, m util.Map, query util.Map) *net.Response
	Request(url string, params util.Map, method string, options util.Map) *net.Response
	RequestRaw(url string, params util.Map, method string, options util.Map) *net.Response
	SafeRequest(url string, params util.Map, method string, options util.Map) *net.Response
	Link(string) string
}

/*Domain 域名*/
type Domain interface {
	URL() string
	Link(s string) string
}

/*GetApp 获取Application */
func GetApp() Application {
	return core.App()
}
