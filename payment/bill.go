package payment

import (
	"github.com/godcong/wego/core"
)

type Bill struct {
	core.Config
	*Payment
}

func newBill(p *Payment) *Bill {
	return &Bill{
		Config:  defaultConfig,
		Payment: p,
	}
}

func NewBill() *Bill {
	return newBill(payment)
}
func (b *Bill) Get(bd string, bt string, op core.Map) core.Map {
	m := make(core.Map)
	m.Set("appid", b.Config.Get("app_id"))
	m.Set("bill_date", bd)
	m.Set("bill_type", bt)
	m.Join(op)
	return b.RequestRaw(DOWNLOADBILL_URL_SUFFIX, m).ToMap()
}
