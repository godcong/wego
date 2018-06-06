package payment

import (
	"net/http"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

/*Order Order */
type Order struct {
	config.Config
	*Payment
	request *http.Request
}

func newOrder(p *Payment) *Order {
	return &Order{
		Config:  defaultConfig,
		Payment: p,
	}
}

/*NewOrder NewOrder */
func NewOrder() *Order {
	return newOrder(payment)
}

//SetRequest to set a http request for Unify to get the client ip
func (o *Order) SetRequest(r *http.Request) *Order {
	o.request = r
	return o
}

/*Unify 统一下单
接口地址
https://apihk.mch.weixin.qq.com/pay/unifiedorder    （建议接入点：东南亚）
https://apius.mch.weixin.qq.com/pay/unifiedorder    （建议接入点：其它）
https://api.mch.weixin.qq.com/pay/unifiedorder        （建议接入点：中国国内）
注：商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
不需要
请求参数
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
设备号	device_info	否	String(32)	013467007045764	终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
商品描述	body	是	String(128)	Ipad mini  16G  白色	商品或支付单简要描述
版本号	version	否	String(32)	1.0	固定值 1.0
商品详情	detail	否	String(6000)
商品详细列表，使用Json格式，传输签名前请务必使用CDATA标签将JSON文本串保护起来。
goods_detail
└ goods_name String 必填 256 商品名称
└ quantity Int 必填4 商品数量
附加数据	attach	否	String(127)	深圳分店	附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。 其他说明见商户订单号
标价币种	fee_type	是	String(16)	GBP	符合ISO4217标准的三位字母代码详见标价币种
标价金额	total_fee	是	Int	888	标价金额，单位为该币种最小计算单位，只能为整数，详见标价金额
终端IP	spbill_create_ip	是	String(16)	123.12.12.123	APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
交易起始时间	time_start	否	String(14)	20091225091010	订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
交易结束时间	time_expire	否	String(14)	20091227091010
订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id。其他详见时间规则
建议：最短失效时间间隔大于1分钟
通知地址	notify_url	是	String(256)	http://www.weixin.qq.com/wxpay/pay.php	接收微信支付异步通知回调地址
交易类型	trade_type	是	String(16)	JSAPI	取值如下：JSAPI，NATIVE，APP，详细说明见参数规定
商品ID	product_id	否	String(32)	12235413214070356458058	trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
指定支付方式	limit_pay	否	String(32)	no_credit	no_credit--指定不能使用信用卡支付
用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
*/
func (o *Order) Unify(m util.Map) *net.Response {
	if !m.Has("spbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("spbill_create_ip", core.GetServerIP())
		}
		//TODO: getclientip with request
		if o.request != nil {
			m.Set("spbill_create_ip", core.GetClientIP(o.request))
		}
	}

	m.Set("appid", o.Config.Get("app_id"))

	//$params['notify_url'] = $params['notify_url'] ?? $this->app['config']['notify_url'];
	if !m.Has("notify_url") {
		m.Set("notify_url", o.Config.Get("notify_url"))
	}
	resp := o.Request(unifiedOrderURLSuffix, m)
	resp.CheckError()
	return resp
}

/*Close 关闭订单
接口地址
https://apihk.mch.weixin.qq.com/pay/closeorder    （建议接入点：东南亚）
https://apius.mch.weixin.qq.com/pay/closeorder    （建议接入点：其它）
https://api.mch.weixin.qq.com/pay/closeorder       （建议接入点：中国国内）
注：商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
不需要。
请求参数
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
商户订单号	out_trade_no	是	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
*/
func (o *Order) Close(no string) *net.Response {
	m := make(util.Map)
	m.Set("appid", o.Config.Get("app_id"))
	m.Set("out_trade_no", no)
	resp := o.Request(closeOrderURLSuffix, m)
	resp.CheckError()
	return resp
}

/** QueryOrder 查询订单
* 场景：刷卡支付、公共号支付、扫码支付、APP支付
接口地址
https://apihk.mch.weixin.qq.com/pay/orderquery    （建议接入点：东南亚）
https://apius.mch.weixin.qq.com/pay/orderquery    （建议接入点：其它）
https://api.mch.weixin.qq.com/pay/orderquery        （建议接入点：中国国内）
注：商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
不需要
请求参数
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	是	String(32)	1230000109	微信支付分配的商户号
微信订单号	transaction_id	二选一	String(32)	1009660380201506130728806387	微信的订单号，优先使用
商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部的订单号，当没提供transaction_id时需要传这个。
随机字符串	nonce_str	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
*/
func (o *Order) query(m util.Map) *net.Response {
	m.Set("appid", o.Config.Get("app_id"))
	return o.Request(orderQueryURLSuffix, m)
}

/*QueryByTransactionID 通过transaction_id查询订单 */
func (o *Order) QueryByTransactionID(id string) *net.Response {
	return o.query(util.Map{"transaction_id": id})
}

/*QueryByOutTradeNumber 通过out_trade_no查询订单 */
func (o *Order) QueryByOutTradeNumber(no string) *net.Response {
	return o.query(util.Map{"out_trade_no": no})
}
