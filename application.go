package wego

type Application interface {
	Payment
	Sandbox
	Order
	Config() Config
	InSandbox() bool
	SetSubMerchant(mchid, appid string) Application
}

type application struct {
	config  Config
	sandbox Sandbox
	payment Payment
	order   Order
}

func (a *application) SetSubMerchant(mchid, appid string) Application {
	a.config.Set("sub_mch_id", mchid)
	a.config.Set("sub_appid", appid)
	return a
}

func (a *application) Unify(m Map) (Map, error) {
	return a.order.Unify(m)
}

func (a *application) Close(no string) (Map, error) {
	return a.order.Close(no)
}

func (a *application) Query(m Map) (Map, error) {
	return a.order.Query(m)
}

func NewApplication(config Config) Application {
	app := &application{
		config: config,
	}
	app.order = NewOrder(app)
	app.payment = NewPayment(app)
	app.sandbox = NewSandbox()
	return app
}

func (a *application) Config() Config {
	return a.config
}

func (a *application) InSandbox() bool {
	return a.config.GetBool("sandbox")
}

func (a *application) GetKey() string {
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
