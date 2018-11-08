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
	"strings"
)

// NewAble ...
type NewAble func(payment *Payment) interface{}

var subLists = util.Map{
	"Bill": newBill,
}

/*Payment Payment */
type Payment struct {
	*core.Config
	client *core.Client
	Sub    util.Map
}

func newPayment(config *core.Config, p util.Map) *Payment {
	payment := &Payment{
		Config: config,
		Sub:    p,
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
			payment.Sub[k] = vv(payment)
		}
	}
	return payment
}

// SubInit ...
func (p *Payment) SubInit() *Payment {
	return subInit(p, subLists)
}

// SubExpectInit ...
func (p *Payment) SubExpectInit(except ...string) *Payment {
	return subInit(p, subLists.Expect(except))
}

// SubOnlyInit ...
func (p *Payment) SubOnlyInit(only ...string) *Payment {
	return subInit(p, subLists.Only(only))
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

//SetSubMerchant set Sub merchat
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
	obj, b := p.Sub["Base"]
	if !b {
		obj = newBase(p)
		p.Sub["Base"] = obj
	}
	return obj.(*Base)
}

// Reverse ...
func (p *Payment) Reverse() *Reverse {
	obj, b := p.Sub["Reverse"]
	if !b {
		obj = newReverse(p)
		p.Sub["Reverse"] = obj
	}
	return obj.(*Reverse)
}

// JSSDK ...
func (p *Payment) JSSDK() *JSSDK {
	obj, b := p.Sub["JSSDK"]
	if !b {
		obj = newJSSDK(p)
		//p.Sub["JSSDK"] = obj
	}
	return obj.(*JSSDK)
}

// RedPack ...
func (p *Payment) RedPack() *RedPack {
	obj, b := p.Sub["RedPack"]
	if !b {
		obj = newRedPack(p)
		//p.Sub["RedPack"] = obj
	}
	return obj.(*RedPack)
}

// Security ...
func (p *Payment) Security() *Security {
	obj, b := p.Sub["Security"]
	if !b {
		obj = newSecurity(p)
		//p.Sub["Security"] = obj
	}
	return obj.(*Security)
}

// Refund ...
func (p *Payment) Refund() *Refund {
	obj, b := p.Sub["Refund"]
	if !b {
		obj = newRefund(p)
		//p.Sub["Refund"] = obj
	}
	return obj.(*Refund)
}

// Order ...
func (p *Payment) Order() *Order {
	obj, b := p.Sub["Order"]
	if !b {
		obj = newOrder(p)
		//p.Sub["Order"] = obj
	}
	return obj.(*Order)
}

// Bill ...
func (p *Payment) Bill() *Bill {
	obj, b := p.Sub["Bill"]
	if !b {
		obj = newBill(p)
		//p.Sub["Bill"] = obj
	}
	return obj.(*Bill)
}

// Transfer ...
func (p *Payment) Transfer() *Transfer {
	obj, b := p.Sub["Transfer"]
	if !b {
		obj = newTransfer(p)
		//p.Sub["Transfer"] = obj
	}
	return obj.(*Transfer)
}

// Sandbox ...
func (p *Payment) Sandbox() *Sandbox {
	obj, b := p.Sub["Sandbox"]
	if !b {
		obj = newSandbox(p)
		//p.Sub["Sandbox"] = obj
	}
	return obj.(*Sandbox)
}

// Coupon ...
func (p *Payment) Coupon() *Coupon {
	obj, b := p.Sub["Coupon"]
	if !b {
		obj = newCoupon(p)
		//p.Sub["Coupon"] = obj
	}
	return obj.(*Coupon)
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
		maps.Set("sign_type", SignTypeMd5.String())
		maps.Set("sign", GenerateSignature(maps, p.GetKey(), MakeSignMD5))
		log.Debug("initRequest", maps)
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

/*SignType SignType */
type SignType int

/*sign types */
const (
	SignTypeMd5        SignType = iota
	SignTypeHmacSha256 SignType = iota
)

// String ...
func (t SignType) String() string {
	if t == SignTypeHmacSha256 {
		return HMACSHA256
	}
	return MD5
}

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
