package payment

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/core"
)

type Payment struct {
	config  core.Config
	client  *core.Client
	token   *core.AccessToken
	sandbox *core.Sandbox
	app     *core.Application

	sub      core.Map
	bill     *Bill
	redPack  *RedPack
	order    *Order
	refund   *Refund
	security *Security
	jssdk    *JSSDK
}

var defaultConfig core.Config
var payment *Payment

func init() {
	defaultConfig = core.GetConfig(core.DeployJoin("payment", "default"))
	app := core.App()

	payment = newPayment(app)
	app.Register("payment", payment)

}

func newPayment(application *core.Application) *Payment {
	client := core.NewClient(defaultConfig)
	token := core.NewAccessToken(defaultConfig, client)
	domain := core.NewDomain("payment")

	payment = &Payment{
		app:    application,
		client: client,
		token:  token,
	}

	client.SetDomain(domain)
	client.SetDataType(core.DATA_TYPE_XML)
	return payment
}

func (p *Payment) SetClient(c *core.Client) *Payment {
	p.client = c
	return p
}

func (p *Payment) GetClient() *core.Client {
	return p.client
}

func (p *Payment) Request(url string, params core.Map) *core.Response {

	return p.client.Request(p.client.Link(url), p.preRequest(params), "post", nil)
}

func (p *Payment) RequestRaw(url string, params core.Map) *core.Response {
	return p.client.RequestRaw(p.client.Link(url), p.preRequest(params), "post", nil)
}

func (p *Payment) SafeRequest(url string, params core.Map) *core.Response {
	return p.client.SafeRequest(p.client.Link(url), p.preRequest(params), "post", nil)
}

func (p *Payment) Pay(params core.Map) core.Map {
	params.Set("appid", p.config.Get("app_id"))
	return p.client.Request(MICROPAY_URL_SUFFIX, p.preRequest(params), "post", nil).ToMap()
}

func (p *Payment) AuthCodeToOpenid(authCode string) core.Map {
	m := make(core.Map)
	m.Set("appid", p.config.Get("app_id"))
	m.Set("auth_code", authCode)
	return p.client.Request(AUTHCODETOOPENID_URL_SUFFIX, p.preRequest(m), "post", nil).ToMap()
}

func (p *Payment) Security() wego.Security {
	if p.security == nil {
		p.security = &Security{
			Config:  p.config,
			Payment: p,
		}
	}
	return p.security
}

func (p *Payment) Refund() wego.Refund {
	if p.refund == nil {
		p.refund = &Refund{
			Config:  p.config,
			Payment: p,
		}
	}
	return p.refund
}

func (p *Payment) AccessToken() *core.AccessToken {
	if p.token == nil {
		p.token = core.NewAccessToken(p.config, p.client)
	}
	return p.token
}

func (p *Payment) Order() wego.Order {
	if p.order == nil {
		p.order = &Order{
			Config:  p.config,
			Payment: p,
		}
	}
	return p.order
}

func (p *Payment) preRequest(params core.Map) core.Map {
	if params != nil {
		params.Set("mch_id", p.client.Get("mch_id"))
		params.Set("nonce_str", core.GenerateUUID())
		params.Set("sub_mch_id", p.client.Get("sub_mch_id"))
		params.Set("sub_appid", p.client.Get("sub_appid"))
		params.Set("sign_type", core.SIGN_TYPE_MD5.String())
		params.Set("sign", core.GenerateSignature(params, p.client.Get("key"), core.SIGN_TYPE_MD5))
	}
	return params
}

func (p *Payment) Link(url string) string {
	if p.config.GetBool("Sandbox") {
		return p.client.URL() + core.SANDBOX_URL_SUFFIX + url
	}
	return p.client.URL() + url
}
