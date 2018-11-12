package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"strconv"
)

/*Settlement Settlement */
type Settlement struct {
	*Payment
}

func newSettlement(p *Payment) interface{} {
	return &Settlement{
		Payment: p,
	}
}

/*NewSettlement NewSettlement */
func NewSettlement(config *core.Config) *Settlement {
	return newSettlement(NewPayment(config)).(*Settlement)
}

// Query ...
func (s *Settlement) Query(useTag, offset, limit int, dateStart, dateEnd string, option ...util.Map) core.Response {
	m := util.MapsToMap(option)
	m.Set("appid", s.Get("app_id"))
	m.Set("date_end", dateEnd)
	m.Set("date_start", dateStart)
	m.Set("offset", strconv.Itoa(offset))
	m.Set("limit", strconv.Itoa(limit))
	m.Set("usetag", strconv.Itoa(useTag))

	return s.Request(paySettlementquery, m)
}
