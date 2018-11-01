package payment

import (
	"github.com/godcong/wego/core"
	"strconv"

	"github.com/godcong/wego/util"
)

/*Refund Refund */
type Refund struct {
	*Payment
}

func newRefund(p *Payment) *Refund {
	return &Refund{
		Payment: p,
	}
}

/*NewRefund NewRefund */
func NewRefund(config *core.Config) *Refund {
	return newRefund(NewPayment(config))
}

func (r *Refund) refund(num string, total, refund int, options util.Map) core.Response {
	options = util.MapNilMake(options)
	options.SetNil("out_refund_no", num)
	options.SetNil("total_fee", strconv.Itoa(total))
	options.SetNil("refund_fee", strconv.Itoa(refund))
	options.SetNil("appid", r.Get("app_id"))

	return r.SafeRequest(refundURLSuffix, options)
}

/*ByOutTradeNumber 按照out_trade_no发起退款
接口地址
接口链接:https://api.mch.weixin.qq.com/secapi/pay/refund
*/
func (r *Refund) ByOutTradeNumber(tradeNum, num string, total, refund int, options util.Map) core.Response {
	options = util.MapNilMake(options)
	options.SetNil("out_trade_no", tradeNum)
	return r.refund(num, total, refund, options)
}

/*ByTransactionID 按照transaction_id发起退款
接口地址
接口链接:https://api.mch.weixin.qq.com/secapi/pay/refund
*/
func (r *Refund) ByTransactionID(tid, num string, total, refund int, options util.Map) core.Response {
	options = util.MapNilMake(options)
	options.SetNil("transaction_id", tid)
	return r.refund(num, total, refund, options)
}

func (r *Refund) query(m util.Map) core.Response {
	m.Set("appid", r.Get("app_id"))
	return r.Request(refundQueryURLSuffix, m)
}

/*QueryByRefundID 按refund_id查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (r *Refund) QueryByRefundID(id string) core.Response {
	return r.query(util.Map{"refund_id": id})
}

/*QueryByOutRefundNumber 按out_refund_no查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (r *Refund) QueryByOutRefundNumber(id string) core.Response {
	return r.query(util.Map{"out_refund_no": id})
}

/*QueryByOutTradeNumber 按out_trade_no查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (r *Refund) QueryByOutTradeNumber(id string) core.Response {
	return r.query(util.Map{"out_trade_no": id})
}

/*QueryByTransactionID 按transaction_id查找退款订单
接口地址
接口链接:https://api.mch.weixin.qq.com/pay/refundquery
*/
func (r *Refund) QueryByTransactionID(id string) core.Response {
	return r.query(util.Map{"transaction_id": id})
}
