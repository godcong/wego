package payment

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Payment struct {
	config  config.Config
	client  *core.Client
	token   *core.AccessToken
	sandbox *core.Sandbox
	app     *core.Application

	sub      util.Map
	bill     *Bill
	redPack  *RedPack
	order    *Order
	refund   *Refund
	security *Security
	jssdk    *JSSDK
}

var defaultConfig config.Config
var payment *Payment

func init() {
	defaultConfig = config.GetConfig("payment.default")
	app := core.App()

	payment = newPayment(app)
	app.Register("payment", payment)

}

func newPayment(application *core.Application) *Payment {
	client := core.NewClient(defaultConfig)
	token := core.NewAccessToken(defaultConfig, client)
	domain := core.NewDomain("default")

	payment = &Payment{
		config: defaultConfig,
		app:    application,
		client: client,
		token:  token,
	}

	client.SetDomain(domain)
	client.SetDataType(core.DataTypeXML)
	return payment
}

func (p *Payment) SetClient(c *core.Client) *Payment {
	p.client = c
	return p
}

func (p *Payment) GetClient() *core.Client {
	return p.client
}

func (p *Payment) Request(url string, params util.Map) *net.Response {
	m := util.Map{
		net.REQUEST_TYPE_XML.String(): p.preRequest(params),
	}

	return p.client.Request(p.client.Link(url), m, "post")
}

func (p *Payment) RequestRaw(url string, params util.Map) *net.Response {
	m := util.Map{
		net.REQUEST_TYPE_XML.String(): p.preRequest(params),
	}

	return p.client.RequestRaw(p.client.Link(url), m, "post")
}

func (p *Payment) SafeRequest(url string, params util.Map) *net.Response {
	m := util.Map{
		net.REQUEST_TYPE_XML.String(): p.preRequest(params),
	}

	return p.client.SafeRequest(p.client.Link(url), m, "post")
}

func (p *Payment) Pay(params util.Map) util.Map {
	params.Set("appid", p.config.Get("app_id"))
	return p.client.Request(MICROPAY_URL_SUFFIX, p.preRequest(params), "post").ToMap()
}

func (p *Payment) AuthCodeToOpenid(authCode string) util.Map {
	m := make(util.Map)
	m.Set("appid", p.config.Get("app_id"))
	m.Set("auth_code", authCode)
	return p.client.Request(AUTHCODETOOPENID_URL_SUFFIX, p.preRequest(m), "post").ToMap()
}

func (p *Payment) Security() wego.Security {
	if p.security == nil {
		p.security = NewSecurity()
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

func (p *Payment) preRequest(params util.Map) util.Map {
	if params != nil {
		params.Set("mch_id", p.client.Get("mch_id"))
		params.Set("nonce_str", util.GenerateUUID())
		if v := p.client.Get("sub_mch_id"); v != "" {
			params.Set("sub_mch_id", v)
		}
		if v := p.client.Get("sub_appid"); v != "" {
			params.Set("sub_appid", v)
		}
		params.Set("sign_type", core.SIGN_TYPE_MD5.String())
		params.Set("sign", core.GenerateSignature(params, p.client.Get("key"), core.MakeSignMD5))
	}
	log.Debug("preRequest", params)
	return params
}

func (p *Payment) Link(url string) string {
	return p.client.Link(url)
}
