package payment

import (
	"fmt"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*Payment Payment */
type Payment struct {
	*core.Config
	client *core.Client
	token  *core.AccessToken

	sub util.Map
}

func newPayment(config *core.Config, client *core.Client) *Payment {
	payment := &Payment{
		Config: config,
		client: client,
	}

	return payment
}

//NewPayment create an payment instance
func NewPayment(config *core.Config, client ...*core.Client) *Payment {
	if client == nil {
		return newPayment(config, core.DefaultClient())
	}
	return newPayment(config, client[0])
}

func (p *Payment) IsSandbox() bool {
	return p.GetBool("sandbox")
}

func (p *Payment) GetKey(s string) string {
	key := p.GetString("key")
	if p.IsSandbox() {
		key = p.Sandbox().GetKey()
	}

	if 32 != len(key) {
		log.Error(fmt.Sprintf("%s should be 32 chars length.", key))
		return ""
	}

	return key
}

//Scheme 获取微信Scheme
//参数: string product_id
//返回: string
func (p *Payment) Scheme(pid string) string {
	m := make(util.Map)
	m.Set("appid", p.Get("app_id"))
	m.Set("mch_id", p.Get("mch_id"))
	m.Set("time_stamp", util.Time())
	m.Set("nonce_str", util.GenerateNonceStr())
	m.Set("product_id", pid)
	m.Set("sign", core.GenerateSignature(m, p.GetString("aes_key"), core.MakeSignMD5))
	return BizPayURL + m.URLEncode()
}

/*Request 普通请求*/
func (p *Payment) Request(url string, params util.Map) core.Response {
	m := util.Map{
		core.DataTypeXML: p.initRequest(params),
	}

	return p.client.Request(Link(url), "post", m)
}

/*RequestRaw raw请求*/
func (p *Payment) RequestRaw(url string, params util.Map) []byte {
	m := util.Map{
		core.DataTypeXML: p.initRequest(params),
	}

	return p.client.RequestRaw(Link(url), "post", m)
}

/*SafeRequest 安全请求*/
func (p *Payment) SafeRequest(url string, params util.Map) core.Response {
	m := util.Map{
		core.DataTypeXML: p.initRequest(params),
	}

	return p.client.SafeRequest(Link(url), "post", m)
}

func (p *Payment) Base() *Base {
	obj, b := p.sub["Base"]
	if !b {
		obj = newBase(p)
		p.sub["Base"] = obj
	}
	return obj.(*Base)
}

//Reverse Reverse
func (p *Payment) Reverse() *Reverse {
	obj, b := p.sub["Reverse"]
	if !b {
		obj = newReverse(p)
		p.sub["Reverse"] = obj
	}
	return obj.(*Reverse)
}

//JSSDK JSSDK
func (p *Payment) JSSDK() *JSSDK {
	obj, b := p.sub["JSSDK"]
	if !b {
		obj = newJSSDK(p)
		p.sub["JSSDK"] = obj
	}
	return obj.(*JSSDK)
}

//RedPack get RedPack
func (p *Payment) RedPack() *RedPack {
	obj, b := p.sub["RedPack"]
	if !b {
		obj = newRedPack(p)
		p.sub["RedPack"] = obj
	}
	return obj.(*RedPack)
}

/*Security get Security */
func (p *Payment) Security() *Security {
	obj, b := p.sub["Security"]
	if !b {
		obj = newSecurity(p)
		p.sub["Security"] = obj
	}
	return obj.(*Security)
}

/*Refund get Refund*/
func (p *Payment) Refund() *Refund {
	obj, b := p.sub["Refund"]
	if !b {
		obj = newRefund(p)
		p.sub["Refund"] = obj
	}
	return obj.(*Refund)
}

/*Order get Order*/
func (p *Payment) Order() *Order {
	obj, b := p.sub["Order"]
	if !b {
		obj = newOrder(p)
		p.sub["Order"] = obj
	}
	return obj.(*Order)
}

/*Bill get Bill*/
func (p *Payment) Bill() *Bill {
	obj, b := p.sub["Bill"]
	if !b {
		obj = newBill(p)
		p.sub["Bill"] = obj
	}
	return obj.(*Bill)
}

/*Transfer get Transfer*/
func (p *Payment) Transfer() *Transfer {
	obj, b := p.sub["Transfer"]
	if !b {
		obj = newTransfer(p)
		p.sub["Transfer"] = obj
	}
	return obj.(*Transfer)
}

func (p *Payment) Sandbox() *Sandbox {
	obj, b := p.sub["Sandbox"]
	if !b {
		obj = newSandbox(p)
		p.sub["Sandbox"] = obj
	}
	return obj.(*Sandbox)
}

func (p *Payment) initRequest(params util.Map) util.Map {
	if params != nil {
		params.Set("mch_id", p.GetString("mch_id"))
		params.Set("nonce_str", util.GenerateUUID())
		if p.Has("sub_mch_id") {
			params.Set("sub_mch_id", p.GetString("sub_mch_id"))
		}
		if p.Has("sub_appid") {
			params.Set("sub_appid", p.GetString("sub_appid"))
		}
		params.Set("sign_type", core.SignTypeMd5.String())
		params.Set("sign", core.GenerateSignature(params, p.GetString("key"), core.MakeSignMD5))
	}
	log.Debug("initRequest", params)
	return params
}

func Link(url string, isSandbox ...bool) string {
	if isSandbox != nil && isSandbox[0] == true {
		core.Connect(core.DefaultConfig().GetStringD("domain.payment.url", domain)+sandboxURLSuffix, url)
	}
	return core.Connect(core.DefaultConfig().GetStringD("domain.payment.url", domain), url)
}
