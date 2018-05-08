package payment

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/crypt"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Transfer struct {
	config.Config
	*Payment
}

func newTransfer(pay *Payment) *Transfer {
	return &Transfer{
		Config:  defaultConfig,
		Payment: pay,
	}
}

func NewTransfer() *Transfer {
	return newTransfer(payment)
}

func (t *Transfer) QueryBalanceOrder(s string) *net.Response {
	m := util.Map{
		"appid":            t.Config.Get("app_id"),
		"mch_id":           t.Config.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(GETTRANSFERINFO_URL_SUFFIX, m)
}

func (t *Transfer) ToBalance(m util.Map) *net.Response {
	m.Set("mch_id", "")
	m.Set("mchid", t.Config.Get("mch_id"))
	m.Set("mch_appid", t.Config.Get("app_id"))

	if !m.Has("spbill_create_ip") {
		m.Set("spbill_create_ip", core.GetServerIp())
	}
	return t.SafeRequest(PROMOTION_TRANSFERS_URL_SUFFIX, m)
}

func (t *Transfer) QueryBankCardOrder(s string) *net.Response {
	m := util.Map{
		"mch_id":           t.Config.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(MMPAYSPTRANS_QUERY_BANK_URL_SUFFIX, m)
}

func (t *Transfer) ToBankCard(m util.Map) *net.Response {
	keys := []string{"bank_code", "partner_trade_no", "enc_bank_no", "enc_true_name", "amount"}
	for _, v := range keys {
		if !m.Has(v) {
			log.Error(v + " is required.")
			return nil
		}
	}
	m.Set("mch_id", t.client.Get("mch_id"))
	m.Set("nonce_str", util.GenerateUUID())

	m.Set("enc_bank_no", crypt.Encrypt(t.Get("pubkey_path"), m.GetString("enc_bank_no")))
	m.Set("enc_true_name", crypt.Encrypt(t.Get("pubkey_path"), m.GetString("enc_true_name")))
	m.Set("sign", core.GenerateSignature(m, t.client.Get("key"), core.MakeSignMD5))
	return t.client.SafeRequest(t.client.Link(MMPAYSPTRANS_PAY_BANK_URL_SUFFIX), util.Map{
		net.REQUEST_TYPE_XML.String(): m,
	}, "post")
}
