package payment

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Reverse struct {
	config.Config
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

func (r *Reverse) ByOutTradeNumber(num string) *net.Response {
	return r.reverse(util.Map{"out_trade_no": num})
}

func (r *Reverse) ByTransactionId(id string) *net.Response {
	return r.reverse(util.Map{"transaction_id": id})
}

func (r *Reverse) reverse(m util.Map) *net.Response {
	m.Set("appid", r.Config.Get("app_id"))
	return r.SafeRequest(REVERSE_URL_SUFFIX, m)
}
