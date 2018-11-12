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
	m := util.MapsToMap(util.Map{
		"appid":      s.Get("app_id"),
		"date_end":   dateEnd,
		"date_start": dateStart,
		"offset":     strconv.Itoa(offset),
		"limit":      strconv.Itoa(limit),
		"usetag":     strconv.Itoa(useTag),
	}, option)
	return s.Request(paySettlementquery, m)
}
