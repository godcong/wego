package payment

import (
	"strconv"

	"github.com/godcong/wego/core"
)

type Refund struct {
	core.Config
	*Payment
}

func (r *Refund) refund(num string, total, refund int, options core.Map) *core.Response {
	options.NilSet("out_refund_no", num)
	options.NilSet("total_fee", strconv.Itoa(total))
	options.NilSet("refund_fee", strconv.Itoa(refund))
	options.NilSet("appid", r.Get("app_id"))

	return r.SafeRequest(REFUND_URL_SUFFIX, options)
}

func (r *Refund) ByOutTradeNumber(tradeNum, num string, total, refund int, options core.Map) core.Map {
	options.NilSet("out_trade_no", tradeNum)
	return r.refund(num, total, refund, options).ToMap()
}

func (r *Refund) ByTransactionId(tid, num string, total, refund int, options core.Map) core.Map {
	options.NilSet("transaction_id", tid)
	return r.refund(num, total, refund, options).ToMap()
}

func (r *Refund) query(m core.Map) *core.Response {
	m.Set("appid", r.Get("app_id"))
	return r.Request(REFUNDQUERY_URL_SUFFIX, m)
}

func (r *Refund) QueryByRefundId(id string) core.Map {
	return r.query(core.Map{"refund_id": id}).ToMap()
}

func (r *Refund) QueryByOutRefundNumber(id string) core.Map {
	return r.query(core.Map{"out_refund_no": id}).ToMap()
}

func (r *Refund) QueryByOutTradeNumber(id string) core.Map {
	return r.query(core.Map{"out_trade_no": id}).ToMap()
}

func (r *Refund) QueryByTransactionId(id string) core.Map {
	return r.query(core.Map{"transaction_id": id}).ToMap()
}
