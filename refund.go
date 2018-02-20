package wego

import "strconv"

type Refund interface {
	Refund(string, int, int, Map) Map
	Query(Map) Map
}

type refund struct {
	Config
	Payment
}

func NewRefund(application Application, config Config) Refund {
	return &refund{
		Config:  config,
		Payment: application.Payment(),
	}
}

func (r *refund) Refund(num string, total, refund int, options Map) Map {
	options.NullSet("out_refund_no", num)
	options.NullSet("total_fee", strconv.Itoa(total))
	options.NullSet("refund_fee", strconv.Itoa(refund))
	options.NullSet("appid", r.Get("app_id"))

	return r.SafeRequest(REFUND_URL_SUFFIX, options)
}

//func (r *refund) safeRequest(url string, m Map) Map {
//	return r.client.SafeRequest(r.client.Link(url), m, "post", nil)
//}
//
//func (r *refund) request(url string, m Map) Map {
//	return r.client.Request(r.client.Link(url), m, "post", nil)
//}

func (r *refund) ByOutTradeNumber(tradeNum, num string, total, refund int, options Map) Map {
	options.NullSet("out_trade_no", tradeNum)
	return r.Refund(num, total, refund, options)
}

func (r *refund) ByTransactionId(tid, num string, total, refund int, options Map) Map {
	options.NullSet("transaction_id", tid)
	return r.Refund(num, total, refund, options)
}

func (r *refund) Query(m Map) Map {
	m.Set("appid", r.Get("app_id"))
	return r.Request(REFUNDQUERY_URL_SUFFIX, m)
}

func (r *refund) QueryByRefundId(id string) Map {
	m := make(Map)
	m.Set("refund_id", id)
	return r.Query(m)
}

func (r *refund) QueryByOutRefundNumber(id string) Map {
	m := make(Map)
	m.Set("out_refund_no", id)
	return r.Query(m)
}

func (r *refund) QueryByOutTradeNumber(id string) Map {
	m := make(Map)
	m.Set("out_trade_no", id)
	return r.Query(m)
}

func (r *refund) QueryByTransactionId(id string) Map {
	m := make(Map)
	m.Set("transaction_id", id)
	return r.Query(m)
}
