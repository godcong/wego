package payment

import (
	"fmt"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"strconv"
)

// NewAble ...
type NewAble func(payment *Payment) interface{}

var moduleLists = util.Map{
	"Bill":     newBill,
	"Coupon":   newCoupon,
	"JSSDK":    newJSSDK,
	"Merchant": newMerchant,
	"Order":    newOrder,
	"RedPack":  newRedPack,
	"Refund":   newRefund,
	"Reverse":  newReverse,
	"Sandbox":  newSandbox,
	"Security": newSecurity,
	"Transfer": newTransfer,
}

/*Payment Payment */
type Payment struct {
	*core.Config
	Module util.Map
	client *core.Client
}

func newPayment(config *core.Config, p util.Map) *Payment {
	payment := &Payment{
		Config: config,
		Module: p,
	}

	return payment
}

//NewPayment create an payment instance
func NewPayment(config *core.Config, v ...interface{}) *Payment {
	payment := newPayment(config, util.Map{})
	payment.SetClient(core.ClientGet(v))
	return payment
}

func subInit(payment *Payment, p util.Map) *Payment {
	for k, v := range p {
		if vv, b := v.(NewAble); b {
			payment.Module[k] = vv(payment)
		}
	}
	return payment
}

// InitModule ...
func (p *Payment) InitModule() *Payment {
	return subInit(p, moduleLists)
}

// InitModuleExpect ...
func (p *Payment) InitModuleExpect(except ...string) *Payment {
	return subInit(p, moduleLists.Expect(except))
}

// InitModuleOnly ...
func (p *Payment) InitModuleOnly(only ...string) *Payment {
	return subInit(p, moduleLists.Only(only))
}

//SetClient set client replace the default client
func (p *Payment) SetClient(client *core.Client) {
	p.client = client
}

//IsSandbox check is use sandbox
func (p *Payment) IsSandbox() bool {
	return p.GetBool("sandbox")
}

//GetKey get key
func (p *Payment) GetKey() string {
	log.Debug(p.String())
	key := p.GetString("key")
	if p.IsSandbox() {
		key = p.Sandbox().GetKey()
		log.Info("sandbox", key)
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
	maps := make(util.Map)
	maps.Set("appid", p.Get("app_id"))
	maps.Set("mch_id", p.Get("mch_id"))
	maps.Set("time_stamp", util.Time())
	maps.Set("nonce_str", util.GenerateNonceStr())
	maps.Set("product_id", pid)
	maps.Set("sign", GenerateSignature(maps, p.GetKey(), MakeSignMD5))
	return bizPayURL + maps.URLEncode()
}

//SetSubMerchant set Module merchat
func (p *Payment) SetSubMerchant(mchID, appID string) *Payment {
	p.Set("sub_mch_id", mchID)
	p.Set("sub_appid", appID)
	return p
}

// Request 默认请求
func (p *Payment) Request(s string, maps util.Map) core.Response {

	m := util.Map{
		core.DataTypeXML: p.initRequest(maps),
	}
	return p.client.Request(p.Link(s), "post", m)
}

// RequestRaw Response转成[]byte
func (p *Payment) RequestRaw(s string, maps util.Map) []byte {
	return p.Request(s, maps).Bytes()
}

// SafeRequest 安全请求
func (p *Payment) SafeRequest(s string, maps util.Map) core.Response {
	m := util.Map{
		core.DataTypeXML:      p.initRequest(maps),
		core.DataTypeSecurity: p.Config,
	}
	return p.client.Request(p.Link(s), "post", m)
}

// Base ...
func (p *Payment) Base() *Base {
	obj, b := p.Module["Base"]
	if !b {
		obj = newBase(p)
		p.Module["Base"] = obj
	}
	return obj.(*Base)
}

// Reverse ...
func (p *Payment) Reverse() *Reverse {
	obj, b := p.Module["Reverse"]
	if !b {
		obj = newReverse(p)
		p.Module["Reverse"] = obj
	}
	return obj.(*Reverse)
}

// JSSDK ...
func (p *Payment) JSSDK() *JSSDK {
	obj, b := p.Module["JSSDK"]
	if !b {
		obj = newJSSDK(p)
		//p.Module["JSSDK"] = obj
	}
	return obj.(*JSSDK)
}

// RedPack ...
func (p *Payment) RedPack() *RedPack {
	obj, b := p.Module["RedPack"]
	if !b {
		obj = newRedPack(p)
		//p.Module["RedPack"] = obj
	}
	return obj.(*RedPack)
}

// Security ...
func (p *Payment) Security() *Security {
	obj, b := p.Module["Security"]
	if !b {
		obj = newSecurity(p)
		//p.Module["Security"] = obj
	}
	return obj.(*Security)
}

// Refund ...
func (p *Payment) Refund() *Refund {
	obj, b := p.Module["Refund"]
	if !b {
		obj = newRefund(p)
		//p.Module["Refund"] = obj
	}
	return obj.(*Refund)
}

// Order ...
func (p *Payment) Order() *Order {
	obj, b := p.Module["Order"]
	if !b {
		obj = newOrder(p)
		//p.Module["Order"] = obj
	}
	return obj.(*Order)
}

// Bill ...
func (p *Payment) Bill() *Bill {
	obj, b := p.Module["Bill"]
	if !b {
		obj = newBill(p)
		//p.Module["Bill"] = obj
	}
	return obj.(*Bill)
}

// Transfer ...
func (p *Payment) Transfer() *Transfer {
	obj, b := p.Module["Transfer"]
	if !b {
		obj = newTransfer(p)
		//p.Module["Transfer"] = obj
	}
	return obj.(*Transfer)
}

// Sandbox ...
func (p *Payment) Sandbox() *Sandbox {
	obj, b := p.Module["Sandbox"]
	if !b {
		obj = newSandbox(p)
		//p.Module["Sandbox"] = obj
	}
	return obj.(*Sandbox)
}

// Coupon ...
func (p *Payment) Coupon() *Coupon {
	obj, b := p.Module["Coupon"]
	if !b {
		obj = newCoupon(p)
		//p.Module["Coupon"] = obj
	}
	return obj.(*Coupon)
}

// HandleRefundedNotify ...
func (p *Payment) HandleRefundedNotify(f NotifyCallback) Notify {
	return &refundedNotify{
		Payment:        p,
		NotifyCallback: f,
	}
}

// HandleRefunded ...
func (p *Payment) HandleRefunded(f NotifyCallback) NotifyFunc {
	return p.HandleRefundedNotify(f).ServeHTTP
}

// HandleScannedNotify ...
func (p *Payment) HandleScannedNotify(f NotifyCallback) Notify {
	return &scannedNotify{
		Payment:        p,
		NotifyCallback: f,
	}
}

// HandleScanned ...
func (p *Payment) HandleScanned(f NotifyCallback) NotifyFunc {
	return p.HandleScannedNotify(f).ServeHTTP
}

// HandlePaidNotify ...
func (p *Payment) HandlePaidNotify(f NotifyCallback) Notify {
	return &paidNotify{
		Payment:        p,
		NotifyCallback: f,
	}
}

// HandlePaid ...
func (p *Payment) HandlePaid(f NotifyCallback) NotifyFunc {
	return p.HandlePaidNotify(f).ServeHTTP
}

func (p *Payment) initRequestWithIgnore(maps util.Map, ignore ...string) util.Map {
	if p == nil || maps == nil {
		return nil
	}

	maps.Set("mch_id", p.GetString("mch_id"))
	maps.Set("nonce_str", util.GenerateUUID())
	if p.Has("sub_mch_id") {
		maps.Set("sub_mch_id", p.GetString("sub_mch_id"))
	}
	if p.Has("sub_appid") {
		maps.Set("sub_appid", p.GetString("sub_appid"))
	}

	if !maps.Has("sign") {
		maps.Set("sign", GenerateSignatureWithIgnore(maps, p.GetKey(), ignore))
	}

	log.Debug("initRequest end", maps)
	return maps
}

func (p *Payment) initRequest(maps util.Map) util.Map {
	if p == nil || maps == nil {
		return nil
	}
	maps.Set("mch_id", p.GetString("mch_id"))
	maps.Set("nonce_str", util.GenerateUUID())
	if p.Has("sub_mch_id") {
		maps.Set("sub_mch_id", p.GetString("sub_mch_id"))
	}
	if p.Has("sub_appid") {
		maps.Set("sub_appid", p.GetString("sub_appid"))
	}

	if !maps.Has("sign") {
		maps.Set("sign", GenerateSignatureWithIgnore(maps, p.GetKey(), []string{FieldSign}))
	}

	log.Debug("initRequest end", maps)
	return maps
}

// SettlementQuery ...
func (p *Payment) SettlementQuery(useTag, offset, limit int, dateStart, dateEnd string, option ...util.Map) core.Response {
	m := util.MapsToMap(util.Map{
		"appid":      p.Get("app_id"),
		"date_end":   dateEnd,
		"date_start": dateStart,
		"offset":     strconv.Itoa(offset),
		"limit":      strconv.Itoa(limit),
		"usetag":     strconv.Itoa(useTag),
	}, option)
	return p.Request(paySettlementquery, m)
}

// Link connect domain url and url suffix
func (p *Payment) Link(url string) string {
	if p.IsSandbox() {
		return core.Connect(core.DefaultConfig().GetStringD("domain.payment.url", domain)+sandboxURLSuffix, url)
	}
	return core.Connect(core.DefaultConfig().GetStringD("domain.payment.url", domain), url)
}

//Link url
func Link(url string) string {
	return core.Connect(core.DefaultConfig().GetStringD("domain.payment.url", domain), url)
}
