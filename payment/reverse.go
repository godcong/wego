package payment

import "github.com/godcong/wego/core"

type Reverse struct {
	core.Config
	*Payment
}

func newReverse(p *Payment) *Reverse {
	return &Reverse{
		Config:  defaultConfig,
		Payment: p,
	}
}

func NewReverse() *Reverse {
	return newReverse(payment)
}

func (r *Reverse) ByOutTradeNumber(num string) core.Map {
	return r.reverse(core.Map{"out_trade_no": num}).ToMap()
}

func (r *Reverse) ByTransactionId(id string) core.Map {
	return r.reverse(core.Map{"transaction_id": id}).ToMap()
}

func (r *Reverse) reverse(m core.Map) *core.Response {
	m.Set("appid", r.Config.Get("app_id"))
	return r.SafeRequest(REVERSE_URL_SUFFIX, m)
}
