package payment_test

import (
	"testing"

	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/util"
)

var tran = payment.NewTransfer()

func TestNewTransfer(t *testing.T) {
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
	t.Log(m1.ToXml())
}
