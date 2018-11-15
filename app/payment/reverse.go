package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Reverse Reverse */
type Reverse struct {
	*Payment
}

func newReverse(p *Payment) interface{} {
	return &Reverse{
		Payment: p,
	}
}

/*NewReverse NewReverse */
func NewReverse(config *core.Config) *Reverse {
	return newReverse(NewPayment(config)).(*Reverse)
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
商户订单号	out_trade_no	String(32)	是	1217752501201407033233368018	商户系统内部的订单号,transaction_id、out_trade_no二选一，如果同时存在优先级:transaction_id> out_trade_no
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
微信订单号	transaction_id	String(32)	否	1217752501201407033233368018	微信的订单号，优先使用
*/
func (r *Reverse) ByTransactionID(id string) core.Response {
	return r.reverse(util.Map{"transaction_id": id})
}

func (r *Reverse) reverse(m util.Map) core.Response {
	m.Set("appid", r.Get("app_id"))
	return r.SafeRequest(payReverse, m)
}
