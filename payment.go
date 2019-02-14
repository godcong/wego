package wego

import (
	"context"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

//Payment ...
type Payment struct {
	*PaymentConfig
	client *Client
	config *Config
	option *PaymentOption
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
func NewPayment(config *Config, opts ...*PaymentOption) *Payment {
	var opt *PaymentOption
	if opts != nil {
		opt = opts[0]
	}
	bt := BodyTypeXML
	return &Payment{
		client: NewClient(&ClientOption{
			//AccessToken: NewAccessToken(config.AccessToken.Credential()),
			BodyType: &bt,
		}),
		PaymentConfig: config.Payment,
		config:        config,
		option:        opt,
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

//DownloadFundFlow 下载资金账单
//资金账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式：20140603
//资金账户类型	account_type	是	String(8)	Basic
//账单的资金来源账户：
//Basic  基本账户
//Operation 运营账户
//Fees 手续费账户
//压缩账单	tar_type	否	String(8)	GZIP	非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
func (obj *Payment) DownloadFundFlow(bd string, at string, opts ...util.Map) Responder {
	m := util.CombineMaps(util.Map{
		"appid":        obj.AppID,
		"bill_date":    bd,
		"sign_type":    util.HMACSHA256,
		"account_type": at,
	}, opts)

	return obj.SafeRequest(payDownloadFundFlow, m)
}

/*ReverseByOutTradeNumber 通过out_trade_no撤销订单
接口地址
https://apihk.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:其它）
https://api.mch.weixin.qq.com/secapi/pay/reverse        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	类型	必填	示例值	描述
商户订单号	out_trade_no	String(32)	是	1217752501201407033233368018	商户系统内部的订单号,transaction_id、out_trade_no二选一，如果同时存在优先级:transaction_id> out_trade_no
*/
func (obj *Payment) ReverseByOutTradeNumber(num string) Responder {
	return obj.reverse(util.Map{"out_trade_no": num})
}

/*ReverseByTransactionID 通过transaction_id撤销订单
接口地址
https://apihk.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:其它）
https://api.mch.weixin.qq.com/secapi/pay/reverse        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	类型	必填	示例值	描述
微信订单号	transaction_id	String(32)	否	1217752501201407033233368018	微信的订单号，优先使用
*/
func (obj *Payment) ReverseByTransactionID(id string) Responder {
	return obj.reverse(util.Map{"transaction_id": id})
}

func (obj *Payment) reverse(m util.Map) Responder {
	return obj.SafeRequest(payReverse, m)
}

/*Unify 统一下单
字段名	变量名	必填	类型	示例值	描述
商品描述	body	是	String(128)	Ipad mini  16G  白色	商品或支付单简要描述
商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。 其他说明见商户订单号
标价金额	total_fee	是	Int	888	标价金额，单位为该币种最小计算单位，只能为整数，详见标价金额
交易类型	trade_type	是	String(16)	JSAPI	取值如下:JSAPI，NATIVE，APP，详细说明见参数规定
用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
*/
func (obj *Payment) Unify(m util.Map, opts ...util.Map) Responder {
	m = util.CombineMaps(m, opts)
	if !m.Has("spbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("spbill_create_ip", core.GetServerIP())
		}
		//TODO: getclientip with request
		//if obj.request != nil {
		//	m.Set("spbill_create_ip", core.GetClientIP(obj.request))
		//}
	}

	if !m.Has("notify_url") {
		m.Set("notify_url", obj.NotifyURL())
	}

	return obj.Request(payUnifiedOrder, m)
}

// Request 默认请求
func (obj *Payment) Request(url string, p util.Map) Responder {
	return PostXML(obj.RemoteURL(url), nil, obj.initPay(p))
}

// SafeRequest 安全请求
func (obj *Payment) SafeRequest(url string, p util.Map) Responder {
	bt := BodyTypeXML
	client := NewClient(&ClientOption{
		UseSafe:  true,
		SafeCert: obj.config.SafeCert,
		BodyType: &bt,
	})
	return client.Post(context.Background(), obj.RemoteURL(url), obj.initPay(p))
}

func (obj *Payment) initPay(p util.Map) util.Map {
	p.Set("appid", obj.AppID)
	p.Set("mch_id", obj.MchID)
	p.Set("nonce_str", util.GenerateUUID())
	if obj.SubMchID != "" {
		p.Set("sub_mch_id", obj.SubMchID)
	}
	if obj.SubAppID != "" {
		p.Set("sub_appid", obj.SubAppID)
	}

	if !p.Has("sign") {
		p.Set("sign", util.GenSign(p, obj.GetKey()))
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
	key := obj.k
	if obj.IsSandbox() {
		keyName := obj.Sandbox().getCacheKey()
		cachedKey := cache.Get(keyName)
		if cachedKey != nil {
			log.Println("cached:", cachedKey.(string))
			key = cachedKey.(string)
			return key
		}

		response := obj.Sandbox().SignKey().ToMap()
		if response.GetString("return_code") == "SUCCESS" {
			key = response.GetString("sandbox_signkey")
			log.Println("cache key:", keyName)
			cache.SetWithTTL(keyName, key, 3*24*3600)
		}
	}

	if 32 != len(key) {
		log.Error(fmt.Sprintf("%s should be 32 chars length.", key))
		return ""
	}
	return key

}

// Sandbox ...
func (obj *Payment) Sandbox() *SandboxProperty {
	if obj.option != nil && obj.option.Sandbox != nil {
		return obj.option.Sandbox
	}
	return &SandboxProperty{}
}

// RemoteURL ...
func (obj *Payment) RemoteURL(uri string) string {
	if obj.IsSandbox() {
		return util.URL(remotePayment(obj), sandboxNew, uri)
	}
	return util.URL(remotePayment(obj), uri)
}
func remotePayment(obj *Payment) string {
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
	return wegoLocal
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
