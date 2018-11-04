package payment

import (
	"net/http"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Order Order */
type Order struct {
	*Payment
	request *http.Request
}

func newOrder(payment *Payment) *Order {
	return &Order{
		Payment: payment,
	}
}

/*NewOrder NewOrder */
func NewOrder(config *core.Config) *Order {
	return newOrder(NewPayment(config))
}

//SetRequest to set a http request for Unify to get the client ip
func (o *Order) SetRequest(r *http.Request) *Order {
	o.request = r
	return o
}

/*Unify 统一下单
字段名	变量名	必填	类型	示例值	描述
商品描述	body	是	String(128)	Ipad mini  16G  白色	商品或支付单简要描述
商户订单号	out_trade_no	是	String(32)	20150806125346	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。 其他说明见商户订单号
标价金额	total_fee	是	Int	888	标价金额，单位为该币种最小计算单位，只能为整数，详见标价金额
交易类型	trade_type	是	String(16)	JSAPI	取值如下:JSAPI，NATIVE，APP，详细说明见参数规定
用户标识	openid	否	String(128)	oUpF8uMuAJO_M2pxb1Q9zNjWeS6o	trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
*/
func (o *Order) Unify(m util.Map) core.Response {
	if !m.Has("spbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("spbill_create_ip", core.GetServerIP())
		}
		//TODO: getclientip with request
		if o.request != nil {
			m.Set("spbill_create_ip", core.GetClientIP(o.request))
		}
	}

	m.Set("appid", o.GetString("app_id"))

	//$params['notify_url'] = $params['notify_url'] ?? $this->app['config']['notify_url'];
	if !m.Has("notify_url") {
		m.Set("notify_url", o.GetString("notify_url"))
	}
	return o.Request(payUnifiedOrder, m)
}

/*Close 关闭订单
字段名	变量名	必填	类型	示例值	描述
商户订单号	out_trade_no	是	String(32)	1217752501201407033233368018	商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
*/
func (o *Order) Close(no string) core.Response {
	m := make(util.Map)
	m.Set("appid", o.Get("app_id"))
	m.Set("out_trade_no", no)
	return o.Request(payCloseOrder, m)
}

/** QueryOrder 查询订单
* 场景:刷卡支付、公共号支付、扫码支付、APP支付
接口地址
https://apihk.mch.weixin.qq.com/pay/orderquery    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/pay/orderquery    （建议接入点:其它）
https://api.mch.weixin.qq.com/pay/orderquery        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
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
func (o *Order) query(m util.Map) core.Response {
	m.Set("appid", o.Get("app_id"))
	return o.Request(payOrderQuery, m)
}

/*QueryByTransactionID 通过transaction_id查询订单
场景:刷卡支付、公共号支付、扫码支付、APP支付
字段名	变量名	必填	类型	示例值	描述
微信订单号	transaction_id	String(32)	1009660380201506130728806387	微信的订单号，优先使用
*/
func (o *Order) QueryByTransactionID(id string) core.Response {
	return o.query(util.Map{"transaction_id": id})
}

/*QueryByOutTradeNumber 通过out_trade_no查询订单
场景:刷卡支付、公共号支付、扫码支付、APP支付
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wxd678efh567hg6787	微信分配的公众账号ID（企业号corpid即为此appId）
商户订单号	out_trade_no	String(32)	20150806125346	商户系统内部的订单号，当没提供transaction_id时需要传这个。
*/
func (o *Order) QueryByOutTradeNumber(no string) core.Response {
	return o.query(util.Map{"out_trade_no": no})
}
