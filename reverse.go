package wego

type Reverse interface {
	ByOutTradeNumber(string) Map
	ByTransactionId(string) Map
}

type reverse struct {
	Config
	Payment
}

func NewReverse(application Application, config Config) Reverse {
	return &reverse{
		Config:  config,
		Payment: application.Payment(),
	}
}

func (r *reverse) ByOutTradeNumber(num string) Map {
	return r.reverse(Map{"out_trade_no": num})
}

func (r *reverse) ByTransactionId(id string) Map {
	return r.reverse(Map{"transaction_id": id})
}

func (r *reverse) reverse(m Map) Map {
	m.Set("appid", r.Get("app_id"))
	return r.SafeRequest(REVERSE_URL_SUFFIX, m)
}
