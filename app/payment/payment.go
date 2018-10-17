package payment

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

/*Payment Payment */
type Payment struct {
	config  Config
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

var defaultConfig Config
var payment *Payment

func init() {
	defaultConfig = config.GetConfig("payment.default")
	app := core.App()

	payment = newPayment(app)
	app.Register("payment", payment)

}

func newPayment(application *core.Application) *Payment {
	client := core.NewClient(defaultConfig)
	token := core.NewAccessToken()
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

/*SetClient 设置client*/
func (p *Payment) SetClient(c *core.Client) *Payment {
	p.client = c
	return p
}

/*GetClient 获取client*/
func (p *Payment) GetClient() *core.Client {
	return p.client
}

/*Request 普通请求*/
func (p *Payment) Request(url string, params util.Map) *net.Response {
	m := util.Map{
		net.RequestTypeXML.String(): p.preRequest(params),
	}

	return p.client.Request(p.client.Link(url), m, "post")
}

/*RequestRaw raw请求*/
func (p *Payment) RequestRaw(url string, params util.Map) *net.Response {
	m := util.Map{
		net.RequestTypeXML.String(): p.preRequest(params),
	}

	return p.client.RequestRaw(p.client.Link(url), m, "post")
}

/*SafeRequest 安全请求*/
func (p *Payment) SafeRequest(url string, params util.Map) *net.Response {
	m := util.Map{
		net.RequestTypeXML.String(): p.preRequest(params),
	}

	return p.client.SafeRequest(p.client.Link(url), m, "post")
}

/*Pay 支付
接口地址
SDK下载:https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=11_1
https://api.mch.weixin.qq.com/pay/micropay
是否需要证书
不需要。
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
货币类型	fee_type	否	String(16)	CNY	符合ISO4217标准的三位字母代码，默认人民币：CNY，详见货币类型
终端IP	spbill_create_ip	是	String(16)	8.8.8.8	调用微信支付API的机器IP
订单优惠标记	goods_tag	否	String(32)	1234	订单优惠标记，代金券或立减优惠功能的参数，详见代金券或立减优惠
指定支付方式	limit_pay	否	String(32)	no_credit	no_credit--指定不能使用信用卡支付
交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
交易结束时间	time_expire	否	String(14)	20091227091010
订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。注意：最短失效时间间隔需大于1分钟
授权码	auth_code	是	String(128)	120061098828009406	扫码支付授权码，设备读取用户微信中的条码或者二维码信息
（注：用户刷卡条形码规则：18位纯数字，以10、11、12、13、14、15开头）
+场景信息	scene_info	否	String(256)
该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开
*/
func (p *Payment) Pay(params util.Map) util.Map {
	params.Set("appid", p.config.Get("app_id"))
	return p.client.Request(microPayURLSuffix, p.preRequest(params), "post").ToMap()
}

/*AuthCodeToOpenid 通过授权码查询公众号Openid
接口链接: https://api.mch.weixin.qq.com/tools/authcodetoopenid
通过授权码查询公众号Openid，调用查询后，该授权码只能由此商户号发起扣款，直至授权码更新。
*/
func (p *Payment) AuthCodeToOpenid(authCode string) util.Map {
	m := make(util.Map)
	m.Set("appid", p.config.Get("app_id"))
	m.Set("auth_code", authCode)
	return p.client.Request(authCodeToOpenidURLSuffix, p.preRequest(m), "post").ToMap()
}

/*Security Security */
func (p *Payment) Security() wego.Security {
	if p.security == nil {
		p.security = NewSecurity()
	}
	return p.security
}

/*Refund 获取Refund*/
func (p *Payment) Refund() wego.Refund {
	if p.refund == nil {
		p.refund = &Refund{
			Config:  p.config,
			Payment: p,
		}
	}
	return p.refund
}

/*AccessToken 获取AccessToken*/
func (p *Payment) AccessToken() *core.AccessToken {
	if p.token == nil {
		p.token = core.NewAccessToken()
	}
	return p.token
}

/*Order 获取order*/
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
		params.Set("sign_type", core.SignTypeMd5.String())
		params.Set("sign", core.GenerateSignature(params, p.client.Get("key"), core.MakeSignMD5))
	}
	log.Debug("preRequest", params)
	return params
}

/*Link link url*/
func (p *Payment) Link(url string) string {
	return p.client.Link(url)
}
