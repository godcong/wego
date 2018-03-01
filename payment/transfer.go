package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/rsa"
)

type Transfer struct {
	core.Config
	Payment
}

func (t *Transfer) QueryBalanceOrder(s string) core.Map {
	m := core.Map{
		"appid":            t.Get("app_id"),
		"mch_id":           t.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(GETTRANSFERINFO_URL_SUFFIX, m)
}

func (t *Transfer) ToBalance(m core.Map) core.Map {
	m.Set("mch_id", "")
	m.Set("mchid", t.Get("mch_id"))
	m.Set("mch_appid", t.Get("app_id"))

	if !m.Has("spbill_create_ip") {
		m.Set("spbill_create_ip", core.GetServerIp())
	}
	return t.SafeRequest(PROMOTION_TRANSFERS_URL_SUFFIX, m)
}

func (t *Transfer) QueryBankCardOrder(s string) core.Map {
	m := core.Map{
		"mch_id":           t.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(MMPAYSPTRANS_QUERY_BANK_URL_SUFFIX, m)
}

func (t *Transfer) ToBankCard(m core.Map) core.Map {
	keys := []string{"bank_code", "partner_trade_no", "enc_bank_no", "enc_true_name", "amount"}
	for _, v := range keys {
		if !m.Has(v) {
			return core.Map{
				"return_code": "FAIL",
				"return_msg":  v + " is required.",
			}
		}
	}

	m.Set("enc_bank_no", rsa.Encrypt(t.Get("pubkey_path"), m.GetString("enc_bank_no")))
	m.Set("enc_true_name", rsa.Encrypt(t.Get("pubkey_path"), m.GetString("enc_true_name")))

	return t.SafeRequest(MMPAYSPTRANS_PAY_BANK_URL_SUFFIX, m)
}
