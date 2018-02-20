package wego

type Refund interface {
	Refund(refundNumber, totalFee, refundFee string, options Map) Map
}

type refund struct {
	Config
	client Client
	app    Application
}

func NewRefund(application Application) Refund {
	return &refund{
		Config: application.Config().GetConfig("payment.default"),
		app:    application,
	}
}

func (r *refund) Refund(num, total, refund string, options Map) Map {
	if !options.Has("out_refund_no") {
		options.Set("out_refund_no", num)
	}

	if !options.Has("total_fee") {
		options.Set("total_fee", total)
	}

	if !options.Has("refund_fee") {
		options.Set("refund_fee", refund)
	}

	if !options.Has("appid") {
		options.Set("appid", r.Get("app_id"))
	}

	return r.safeRequest(REFUND_URL_SUFFIX, options)
}

func (r *refund) safeRequest(url string, m Map) Map {
	return r.client.SafeRequest(r.client.Link(url), m, "post", nil)
}
