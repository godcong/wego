package wego

type Application interface {
	Payment() Payment
	Client() Client
	Config() Config
	Scheme(id string) string
	GetKey(s string) string
	InSandbox() bool
	SetSubMerchant(mchid, appid string) Application
}

type application struct {
	config  Config
	sandbox Sandbox
	payment Payment
	order   Order
	client  Client
	request *Request
}

var app *application

func initApp(config Config) {
	if app == nil {
		app = newApplication(config)
	}
}

func GetOrder() Order {
	return app.Payment().Order()
}

func (a *application) Payment() Payment {
	if a.payment == nil {
		a.payment = NewPayment(a)
	}
	return a.payment
}

func (a *application) Request() *Request {
	if a.request == nil {
		a.request = NewRequest(a)
	}
	return a.request
}

func (a *application) Client() Client {
	if a.client == nil {
		a.client = NewClient(a, a.Request())
	}
	return a.client
}

func (a *application) SetSubMerchant(mchid, appid string) Application {
	a.config.Set("sub_mch_id", mchid)
	a.config.Set("sub_appid", appid)
	return a
}

//func (a *application) Unify(m Map) (Map, error) {
//	return a.order.Unify(m)
//}
//
//func (a *application) Close(no string) (Map, error) {
//	return a.order.Close(no)
//}
//
//func (a *application) Query(m Map) (Map, error) {
//	return a.order.Query(m)
//}
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

func (a *application) Config() Config {
	return a.config
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
