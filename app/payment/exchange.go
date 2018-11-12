package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Exchange Exchange */
type Exchange struct {
	*Payment
}

func newExchange(p *Payment) interface{} {
	return &Exchange{
		Payment: p,
	}
}

/*NewExchange NewExchange */
func NewExchange(config *core.Config) *Exchange {
	return newExchange(NewPayment(config)).(*Exchange)
}

// QueryRate ...
func (s *Exchange) QueryRate(feeType, date string, option ...util.Map) core.Response {

	m := util.MapsToMap(util.Map{
		"appid":    s.Get("app_id"),
		"fee_type": feeType,
		"date":     date,
		"mch_id":   s.Get("mch_id"),
	}, option)

	m.Set("sign", GenerateSignatureWithIgnore(m, s.GetKey(), nil))
	return s.client.Request(s.Link(payQueryexchagerate), "post", util.Map{
		core.DataTypeXML: m,
	})
}
