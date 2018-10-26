package payment_test

import (
	"github.com/godcong/wego/core"
	"testing"

	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/util"
)

func TestNewTransfer(t *testing.T) {
	var tran = payment.NewTransfer(core.DefaultConfig().GetSubConfig("payment.default"))
	m := util.Map{}
	// 商户企业付款单号 partner_trade_no
	// 收款方银行卡号 enc_bank_no
	// 收款方用户名 enc_true_name
	// 付款金额 amount
	m.Set("partner_trade_no", "1234")
	m.Set("enc_bank_no", "6217001210053551022")
	m.Set("enc_true_name", "蒋聪聪")
	m.Set("bank_code", "1003")
	m.Set("amount", "1000")
	m1 := tran.ToBankCard(m)
	t.Log(m1.ToMap())
}
