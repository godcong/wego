package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
	_ "github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/util"
)

var out_trade_no = "201813091059590000003433-asd003"

var long_url = "weixin://wxpay/bizpayurl?pr=etxB4DY"

// TestOrder_Query ...
func TestOrder_Query(t *testing.T) {

}

// TestOrder_Close ...
func TestOrder_Close(t *testing.T) {
	r := wego.PaymentOrder().Close(out_trade_no + "5")
	log.Println(string(r.Bytes()))

}

// TestOrder_QueryByOutTradeNumber ...
func TestOrder_QueryByOutTradeNumber(t *testing.T) {
	r := wego.PaymentOrder().QueryByOutTradeNumber(out_trade_no + "5")
	log.Println(string(r.Bytes()))
}

// TestOrder_QueryByTransactionId ...
func TestOrder_QueryByTransactionId(t *testing.T) {
	r := wego.PaymentOrder().QueryByTransactionID(out_trade_no + "5")
	log.Println(string(r.Bytes()))
}

// TestOrder_Unify ...
func TestOrder_Unify(t *testing.T) {
	m := make(util.Map)
	m.Set("body", "腾讯充值中心-QQ会员充值")
	m.Set("out_trade_no", out_trade_no+"6")
	//m.Set("device_info", "")
	////m.Set("fee_type", "CNY")
	m.Set("total_fee", "1")
	////m.Set("spbill_create_ip", "123.12.12.123")
	//m.Set("notify_url", "https://test.letiantian.me/wxpay/notify")
	m.Set("trade_type", "NATIVE")
	//m.Set("openid", "oLyBi0hSYhggnD-kOIms0IzZFqrc")
	//m.Set("openid", "oE_gl0Yr54fUjBhU5nBlP4hS2efo")

	////m.Set("product_id", "12")
	r := wego.Payment().Order().Unify(m)
	if r.Error() != nil {
		t.Log(r)
	}
	log.Println(string(r.Bytes()))
	log.Println(r.ToMap())
	//order := payment.NewOrder()
	//resp := order.Unify(m)
	//log.Println(resp.ToMap())
	//{"appid":"wx426b3015555a46be","code_url":"weixin://wxpay/bizpayurl?pr=D3sNT8y","mch_id":"1900009851","nonce_str":"FRFByNNdrzRuEGkp","prepay_id":"wx20180220113507842dff20340601189342","result_code":"SUCCESS","return_code":"SUCCESS","return_msg":"OK","sign":"D398DA0644A14D0BC00A8B82D8D4ECDC","trade_type":"NATIVE"}

}

const rltRefund = `<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg><appid><![CDATA[wxbafed7010e0f4531]]></appid><mch_id><![CDATA[1497361732]]></mch_id><nonce_str><![CDATA[C7M5peUJulyD3ljQ]]></nonce_str><sign><![CDATA[AED925A15E9531F4DAF61C4EEA05B608]]></sign><result_code><![CDATA[SUCCESS]]></result_code><transaction_id><![CDATA[4200000059201803063688861057]]></transaction_id><out_trade_no><![CDATA[20180306145209635869487577]]></out_trade_no><out_refund_no><![CDATA[4200000059201803063688861057]]></out_refund_no><refund_id><![CDATA[50000106012018030603707835282]]></refund_id><refund_channel><![CDATA[]]></refund_channel><refund_fee>3</refund_fee><coupon_refund_fee>0</coupon_refund_fee><total_fee>3</total_fee><cash_fee>3</cash_fee><coupon_refund_count>0</coupon_refund_count><cash_refund_fee>3</cash_refund_fee></xml>]`

// TestRefund_Refund ...
func TestRefund_Refund(t *testing.T) {

}

// TestRefund_Query ...
func TestRefund_Query(t *testing.T) {

}

// TestSecurity_GetPublicKey ...
func TestSecurity_GetPublicKey(t *testing.T) {
	//m := wego.GetSecurity().GetPublicKey()
	//log.Println(m.ToMap())
}
