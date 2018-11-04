package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Reverse Reverse */
type Reverse struct {
	*Payment
}

func newReverse(p *Payment) *Reverse {
	return &Reverse{
		Payment: p,
	}
}

/*NewReverse NewReverse */
func NewReverse(config *core.Config) *Reverse {
	return newReverse(NewPayment(config))
}

/*ByOutTradeNumber 通过out_trade_no撤销订单
接口地址
https://apihk.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:其它）
https://api.mch.weixin.qq.com/secapi/pay/reverse        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	类型	必填	示例值	描述
公众账号ID	appid	String(32)	是	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	String(32)	是	1900000109	微信支付分配的商户号
微信订单号	transaction_id	String(32)	否	1217752501201407033233368018	微信的订单号，优先使用
商户订单号	out_trade_no	String(32)	是	1217752501201407033233368018	商户系统内部的订单号,transaction_id、out_trade_no二选一，如果同时存在优先级:transaction_id> out_trade_no
随机字符串	nonce_str	String(32)	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	String(32)	是	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
*/
func (r *Reverse) ByOutTradeNumber(num string) core.Response {
	return r.reverse(util.Map{"out_trade_no": num})
}

/*ByTransactionID 通过transaction_id撤销订单
接口地址
https://apihk.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:东南亚）
https://apius.mch.weixin.qq.com/secapi/pay/reverse    （建议接入点:其它）
https://api.mch.weixin.qq.com/secapi/pay/reverse        （建议接入点:中国国内）
注:商户可根据实际请求情况选择最优域名进行访问，建议在接入时做好兼容，当访问其中一个域名出现异常时，可自动切换为其他域名。
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	类型	必填	示例值	描述
公众账号ID	appid	String(32)	是	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	String(32)	是	1900000109	微信支付分配的商户号
微信订单号	transaction_id	String(32)	否	1217752501201407033233368018	微信的订单号，优先使用
商户订单号	out_trade_no	String(32)	是	1217752501201407033233368018	商户系统内部的订单号,transaction_id、out_trade_no二选一，如果同时存在优先级:transaction_id> out_trade_no
随机字符串	nonce_str	String(32)	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	String(32)	是	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
*/
func (r *Reverse) ByTransactionID(id string) core.Response {
	return r.reverse(util.Map{"transaction_id": id})
}

func (r *Reverse) reverse(m util.Map) core.Response {
	m.Set("appid", r.Get("app_id"))
	return r.SafeRequest(payReverse, m)
}
