package wego

import "strings"

type Bill interface {
	GetBill(string, string, Map) Map
}

type bill struct {
	Config
	Payment
}

func (b *bill) GetBill(dat string, typ string, op Map) Map {
	m := make(Map)
	m.Set("appid", b.Get("app_id"))
	m.Set("bill_date", dat)
	m.Set("bill_type", typ)
	m.Join(op)
	resp := b.RequestRaw(DOWNLOADBILL_URL_SUFFIX, m)
	if strings.Index(string(resp), "<xml>") == 0 {
		return XmlToMap(resp)
	}
	r := make(Map)
	r.Set("contents", string(resp))
	return r
}

func NewBill(application Application, config Config) Bill {
	return &bill{
		Config:  config,
		Payment: application.Payment(),
	}
}
