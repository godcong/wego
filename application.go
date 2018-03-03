package wego

import (
	"github.com/godcong/wego/core"
)

type Application interface {
	Get(name string) interface{}
	Register(name string, v interface{})
	Scheme(id string) string
	GetKey(s string) string
	InSandbox() bool
	SetSubMerchant(mchid, appid string) *core.Application
}

//type AccessToken interface {
//	GetToken() core.Token
//	Refresh() core.AccessToken
//ApplyToRequest(RequestInterface, Map) RequestInterface
//getCredentials() Map
//getQuery() Map
//sendRequest() []byte
//}

func GetApp() Application {
	return core.App()
}

//
//func GetBill() Bill {
//	return app.Payment().Bill()
//}
//
//func (a *application) MiniProgram() MiniProgram {
//	if a.mini_program == nil {
//		a.mini_program = NewMiniProgram(a)
//	}
//	return a.mini_program
//}
//
//func (a *application) Cache() cache.Cache {
//	return cache.GetCache()
//}
//
//func (a *application) Payment() Payment {
//	if a.payment == nil {
//		a.payment = NewPayment(a)
//	}
//	return a.payment
//}
//
//func (a *application) Request() *Request {
//	if a.request == nil {
//		a.request = NewRequest()
//	}
//	return a.request
//}
//
//func (a *application) Client(config core.Config) Client {
//	return NewClient(a, config, a.Request())
//}
//

//
//func NewApplication(v ...interface{}) Application {
//	return newApplication(v)
//}
//
//func (a *application) GetConfig(s string) core.Config {
//	if s == "" {
//		return a.config
//	}
//	return a.config.GetConfig(s)
//}
