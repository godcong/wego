package app

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

//Payment ...
type Payment struct {
	*PaymentProperty
	client   *Client
	property *Property
	option   *PaymentOption
}

// PaymentOption ...
type PaymentOption struct {
	RemoteAddress string
	LocalHost     string
	UseSandbox    bool
	Sandbox       *SandboxProperty
	NotifyURL     string
	RefundURL     string
}

// NewPayment ...
func NewPayment(property *Property, opts ...*PaymentOption) *Payment {
	var opt *PaymentOption
	if opts != nil {
		opt = opts[0]
	}
	bt := BodyTypeXML
	return &Payment{
		client: NewClient(&ClientOption{
			//AccessToken: NewAccessToken(property.AccessToken.Credential()),
			BodyType: &bt,
		}),
		PaymentProperty: property.Payment,
		property:        property,
		option:          opt,
	}
}

/*Pay 支付
接口地址
SDK下载:https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=11_1
https://api.mch.weixin.qq.com/pay/micropay
输入参数
名称	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
设备号	device_info	否	String(32)	013467007045764	终端设备号(商户自定义，如门店编号)
随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
商品描述	body	是	String(128)	image形象店-深圳腾大- QQ公仔	商品简单描述，该字段须严格按照规范传递，具体请见参数规定
商品详情	detail	否	String(6000)
单品优惠功能字段，需要接入详见单品优惠详细说明
附加数据	attach	否	String(127)	说明	附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
商户订单号	out_trade_no	是	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
订单金额	total_fee	是	Int	888	订单总金额，单位为分，只能为整数，详见支付金额
货币类型	fee_type	否	String(16)	CNY	符合ISO4217标准的三位字母代码，默认人民币:CNY，详见货币类型
终端IP	spbill_create_ip	是	String(16)	8.8.8.8	调用微信支付API的机器IP
订单优惠标记	goods_tag	否	String(32)	1234	订单优惠标记，代金券或立减优惠功能的参数，详见代金券或立减优惠
指定支付方式	limit_pay	否	String(32)	no_credit	no_credit--指定不能使用信用卡支付
交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
交易结束时间	time_expire	否	String(14)	20091227091010 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。注意:最短失效时间间隔需大于1分钟
授权码	auth_code	是	String(128)	120061098828009406	扫码支付授权码，设备读取用户微信中的条码或者二维码信息（注:用户刷卡条形码规则:18位纯数字，以10、11、12、13、14、15开头）
+场景信息	scene_info	否	String(256)
该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开
*/
func (obj *Payment) Pay(p util.Map) Responder {
	p.Set("appid", obj.AppID)

	//set notify callback
	notify := obj.NotifyURL()
	if !p.Has("notify_url") {
		p.Set("notify_url", notify)
	}

	return obj.Request(payMicroPay, p)
}

// UnifyResponse ...
type UnifyResponse struct {
	AppID      string `xml:"appid"`
	CodeURL    string `xml:"code_url"`
	DeviceInfo string `xml:"device_info"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	MchID      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	PrepayID   string `xml:"prepay_id"`
	ResultCode string `xml:"result_code"`
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Sign       string `xml:"sign"`
	TradeType  string `xml:"trade_type"`
}

/*Unify 统一下单
字段名	变量名	必填	类型	示例值	描述
商品描述	body	是	String(128)	Ipad mini  16G  白色	商品或支付单简要描述
商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。 其他说明见商户订单号
标价金额	total_fee	是	Int	888	标价金额，单位为该币种最小计算单位，只能为整数，详见标价金额
交易类型	trade_type	是	String(16)	JSAPI	取值如下:JSAPI，NATIVE，APP，详细说明见参数规定
用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
*/
func (obj *Payment) Unify(m util.Map) Responder {
	if !m.Has("spbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("spbill_create_ip", core.GetServerIP())
		}
		//TODO: getclientip with request
		//if obj.request != nil {
		//	m.Set("spbill_create_ip", core.GetClientIP(obj.request))
		//}
	}

	m.Set("appid", obj.AppID)

	//$params['notify_url'] = $params['notify_url'] ?? $this->app['config']['notify_url'];
	if !m.Has("notify_url") {
		m.Set("notify_url", obj.NotifyURL())
	}

	return obj.Request(payUnifiedOrder, m)
}

// Request 默认请求
func (obj *Payment) Request(uri string, p util.Map) Responder {
	return PostXML(obj.RemoteURL(uri), nil, obj.initPay(p))
}

// SafeRequest 安全请求
func (obj *Payment) SafeRequest(s string, p util.Map) Responder {
	//m := util.Map{
	//	core.DataTypeXML:      obj.initRequest(p),
	//	core.DataTypeSecurity: obj.Config,
	//}
	//return core.Request(core.POST, obj.Link(s), m)
	return nil
}

func (obj *Payment) initPay(p util.Map) util.Map {
	p.Set("mch_id", obj.MchID)
	p.Set("nonce_str", util.GenerateUUID())
	if obj.SubMchID != "" {
		p.Set("sub_mch_id", obj.SubMchID)
	}
	if obj.SubAppID != "" {
		p.Set("sub_appid", obj.SubAppID)
	}

	if !p.Has("sign") {
		p.Set("sign", util.GenSign(p, obj.GetKey(), util.SignMD5))
	}
	log.Debug("initPay end", p)
	return p
}

// IsSandbox ...
func (obj *Payment) IsSandbox() bool {
	if obj.option != nil {
		return obj.option.UseSandbox
	}
	return false
}

/*GetKey 沙箱key(string类型) */
func (obj *Payment) GetKey() string {
	key := obj.Key
	if obj.IsSandbox() {
		cachedKey := cache.Get(obj.getCacheKey())
		if cachedKey != nil {
			key = cachedKey.(string)
		}

		response := obj.sandboxSignKey().ToMap()
		if response.GetString("return_code") == "SUCCESS" {
			key = response.GetString("sandbox_signkey")
			cache.SetWithTTL(obj.getCacheKey(), key, 3*24*3600)
		}
	}

	if 32 != len(key) {
		log.Error(fmt.Sprintf("%s should be 32 chars length.", key))
		return ""
	}
	return key

}

func (obj *Payment) getCacheKey() string {
	name := obj.option.Sandbox.AppID + "." + obj.option.Sandbox.MchID
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

func (obj *Payment) sandboxSignKey() Responder {
	m := make(util.Map)
	m.Set("mch_id", obj.option.Sandbox.MchID)
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := util.GenSign(m, obj.option.Sandbox.Key, util.SignMD5)
	m.Set("sign", sign)
	resp := PostXML(obj.RemoteURL(sandboxSignKeyURLSuffix), nil, m)

	return resp

}

// RemoteURL ...
func (obj *Payment) RemoteURL(uri string) string {
	if obj.IsSandbox() {
		return util.URL(remote(obj), sandboxURLSuffix, uri)
	}
	return util.URL(remote(obj), uri)
}
func remote(obj *Payment) string {
	if obj != nil && obj.option != nil && obj.option.RemoteAddress != "" {
		return obj.option.RemoteAddress
	}
	return apiMCHWeixin
}

// LocalURL ...
func (obj *Payment) LocalURL() string {
	return local(obj)
}
func local(obj *Payment) string {
	if obj != nil && obj.option != nil && obj.option.LocalHost != "" {
		return obj.option.LocalHost
	}
	return localAddress
}

// NotifyURL ...
func (obj *Payment) NotifyURL() string {
	return util.URL(obj.LocalURL(), paymentNotifyURL(obj))
}
func paymentNotifyURL(obj *Payment) string {
	if obj != nil && obj.option != nil && obj.option.NotifyURL != "" {
		return obj.option.NotifyURL
	}
	return notifyCB
}

// RefundURL ...
func (obj *Payment) RefundURL() string {
	return util.URL(obj.LocalURL(), paymentRefundURL(obj))
}
func paymentRefundURL(obj *Payment) string {
	if obj != nil && obj.option != nil && obj.option.RefundURL != "" {
		return obj.option.RefundURL
	}
	return refundedCB
}
