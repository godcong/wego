package payment

import "github.com/godcong/wego/core"

type Reverse struct {
	core.Config
	*Payment
}

func (r *Reverse) ByOutTradeNumber(num string) core.Map {
	return r.reverse(core.Map{"out_trade_no": num})
}

func (r *Reverse) ByTransactionId(id string) core.Map {
	return r.reverse(core.Map{"transaction_id": id})
}

func (r *Reverse) reverse(m core.Map) core.Map {
	m.Set("appid", r.Get("app_id"))
	return r.SafeRequest(REVERSE_URL_SUFFIX, m)
}
