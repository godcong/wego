package wego

import "github.com/godcong/wego/rsa"

type Transfer interface {
	QueryBalanceOrder(string) Map
	ToBalance(Map) Map
	QueryBankCardOrder(string) Map
	ToBankCard(Map) Map
}

type transfer struct {
	Config
	Payment
}

func NewTransfer(application Application, config Config) Transfer {
	return &transfer{
		Config:  config,
		Payment: application.Payment(),
	}
}

func (t *transfer) QueryBalanceOrder(s string) Map {
	m := Map{
		"appid":            t.Get("app_id"),
		"mch_id":           t.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(GETTRANSFERINFO_URL_SUFFIX, m)
}

func (t *transfer) ToBalance(m Map) Map {
	m.Set("mch_id", "")
	m.Set("mchid", t.Get("mch_id"))
	m.Set("mch_appid", t.Get("app_id"))

	if !m.Has("spbill_create_ip") {
		m.Set("spbill_create_ip", GetServerIp())
	}
	return t.SafeRequest(PROMOTION_TRANSFERS_URL_SUFFIX, m)
}

func (t *transfer) QueryBankCardOrder(s string) Map {
	m := Map{
		"mch_id":           t.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(MMPAYSPTRANS_QUERY_BANK_URL_SUFFIX, m)
}

func (t *transfer) ToBankCard(m Map) Map {
	keys := []string{"bank_code", "partner_trade_no", "enc_bank_no", "enc_true_name", "amount"}
	for _, v := range keys {
		if !m.Has(v) {
			return Map{
				"return_code": "FAIL",
				"return_msg":  v + " is required.",
			}
		}
	}

	m.Set("enc_bank_no", rsa.Encrypt(t.Get("public_key"), m.Get("enc_bank_no")))
	m.Set("enc_true_name", rsa.Encrypt(t.Get("public_key"), m.Get("enc_true_name")))

	return t.SafeRequest(MMPAYSPTRANS_PAY_BANK_URL_SUFFIX, m)
}
