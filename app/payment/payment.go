package payment

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"io"
	"net/http"
	"strings"
)

// NewAble ...
type NewAble func(payment *Payment) interface{}

var moduleLists = util.Map{
	"Bill":     newBill,
	"Coupon":   newCoupon,
	"JSSDK":    newJSSDK,
	"Merchant": newMerchant,
	"Notify":   newNotify,
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

/*Request 普通请求*/
func (p *Payment) Request(s string, maps util.Map) core.Response {
	m := util.Map{
		core.DataTypeXML: p.initRequest(maps),
	}
	return p.client.Request(p.Link(s), "post", m)
}

/*RequestRaw raw请求*/
func (p *Payment) RequestRaw(s string, maps util.Map) []byte {
	return p.Request(s, maps).Bytes()
}

/*SafeRequest 安全请求*/
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

// Notify ...
func (p *Payment) Notify() *Notify {
	obj, b := p.Module["Notify"]
	if !b {
		obj = newNotify(p)
		//p.Module["JSSDK"] = obj
	}
	return obj.(*Notify)
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

// HandleRefunded ...
func (p *Payment) HandleRefunded(f func(w http.ResponseWriter, req *http.Request)) {

}

// HandlePaid ...
func (p *Payment) HandlePaid(f func(w http.ResponseWriter, req *http.Request)) {

}

// HandleScanned ...
func (p *Payment) HandleScanned(f func(w http.ResponseWriter, req *http.Request)) {

}

func (p *Payment) initRequest(maps util.Map) util.Map {
	if maps != nil {
		maps.Set("mch_id", p.GetString("mch_id"))
		maps.Set("nonce_str", util.GenerateUUID())
		if p.Has("sub_mch_id") {
			maps.Set("sub_mch_id", p.GetString("sub_mch_id"))
		}
		if p.Has("sub_appid") {
			maps.Set("sub_appid", p.GetString("sub_appid"))
		}
		if maps.Has("sign_type") {
			switch maps.GetString("sign_type") {
			case MD5:
				maps.Set("sign", GenerateSignature(maps, p.GetKey(), MakeSignMD5))
			case HMACSHA256:
				maps.Set("sign", GenerateSignature(maps, p.GetKey(), MakeSignHMACSHA256))
			default:
				log.Error("wrong sign_type")
			}
		} else {
			maps.Set("sign_type", MD5)
			maps.Set("sign", GenerateSignature(maps, p.GetKey(), MakeSignMD5))
		}

		log.Debug("initRequest end", maps)
	}

	return maps
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

/*SignFunc sign函数定义 */
type SignFunc func(data, key string) string

// MakeSignHMACSHA256 make sign with hmac-sha256
func MakeSignHMACSHA256(data, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(data))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// MakeSignMD5 make sign with md5
func MakeSignMD5(data, key string) string {
	m := md5.New()
	_, _ = io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// GenerateSignature make sign from map data
func GenerateSignature(m util.Map, key string, fn SignFunc) string {
	keys := m.SortKeys()
	var sign []string

	for _, k := range keys {
		if k == FieldSign {
			continue
		}
		v := strings.TrimSpace(m.GetString(k))

		if len(v) > 0 {
			log.Debug(k, v)
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}

	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	return fn(sb, key)
}
