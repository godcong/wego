package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
	"github.com/godcong/wego/core/util"
	_ "github.com/godcong/wego/payment"
)

var out_trade_no = "201813091059590000003433-asd003"

var long_url = "weixin://wxpay/bizpayurl?pr=etxB4DY"

func TestOrder_Query(t *testing.T) {
	//m := make(wego.Map)
	//m.Set("out_trade_no", out_trade_no)

	r := wego.GetPayment().Order()
	log.Println(r)
	// {"appid":"wx426b3015555a46be","attach":"","mch_id":"1900009851","nonce_str":"lJhbZ9dwP4Pd5aKm","out_trade_no":"201813091059590000003433-asd002","result_code":"SUCCESS","return_code":"SUCCESS","return_msg":"OK","sign":"2F60EDECAAC5F139A82570B6724AA941","trade_state":"CLOSED","trade_state_desc":"订单已关闭"}

}

func TestOrder_Close(t *testing.T) {
	r := wego.GetPayment().Order().Close(out_trade_no)
	log.Println(string(r.ToJson()))

}

func TestOrder_QueryByOutTradeNumber(t *testing.T) {
	r := wego.GetPayment().Order().QueryByOutTradeNumber(out_trade_no)
	log.Println(r)
}

func TestOrder_QueryByTransactionId(t *testing.T) {
	r := wego.GetPayment().Order().QueryByTransactionId("123")
	log.Println(r)
}

func TestOrder_Unify(t *testing.T) {
	m := make(util.Map)
	m.Set("body", "腾讯充值中心-QQ会员充值")
	m.Set("out_trade_no", out_trade_no+"4")
	//m.Set("device_info", "")
	////m.Set("fee_type", "CNY")
	m.Set("total_fee", "1")
	////m.Set("spbill_create_ip", "123.12.12.123")
	//m.Set("notify_url", "https://test.letiantian.me/wxpay/notify")
	m.Set("trade_type", "NATIVE")
	//m.Set("openid", "oLyBi0hSYhggnD-kOIms0IzZFqrc")
	//m.Set("openid", "oE_gl0Yr54fUjBhU5nBlP4hS2efo")

	////m.Set("product_id", "12")
	r := wego.GetPayment().Order().Unify(m)

	log.Println(string(r.ToString()))
	log.Println(r.ToMap())
	//{"appid":"wx426b3015555a46be","code_url":"weixin://wxpay/bizpayurl?pr=D3sNT8y","mch_id":"1900009851","nonce_str":"FRFByNNdrzRuEGkp","prepay_id":"wx20180220113507842dff20340601189342","result_code":"SUCCESS","return_code":"SUCCESS","return_msg":"OK","sign":"D398DA0644A14D0BC00A8B82D8D4ECDC","trade_type":"NATIVE"}

}

const rltRefund = `<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg><appid><![CDATA[wxbafed7010e0f4531]]></appid><mch_id><![CDATA[1497361732]]></mch_id><nonce_str><![CDATA[C7M5peUJulyD3ljQ]]></nonce_str><sign><![CDATA[AED925A15E9531F4DAF61C4EEA05B608]]></sign><result_code><![CDATA[SUCCESS]]></result_code><transaction_id><![CDATA[4200000059201803063688861057]]></transaction_id><out_trade_no><![CDATA[20180306145209635869487577]]></out_trade_no><out_refund_no><![CDATA[4200000059201803063688861057]]></out_refund_no><refund_id><![CDATA[50000106012018030603707835282]]></refund_id><refund_channel><![CDATA[]]></refund_channel><refund_fee>3</refund_fee><coupon_refund_fee>0</coupon_refund_fee><total_fee>3</total_fee><cash_fee>3</cash_fee><coupon_refund_count>0</coupon_refund_count><cash_refund_fee>3</cash_refund_fee></xml>]`

func TestRefund_Refund(t *testing.T) {
	//r := wego.GetPayment().Refund().ByOutTradeNumber(`20180313160643671522177497`, `1`, 30, 30, nil)
	r := wego.GetPayment().Refund().ByTransactionId(`4200000066201803138050731804`, `2`, 3, 3, nil)
	log.Println(r.ToMap())
	//{"appid":"wx426b3015555a46be","err_code":"ORDERNOTEXIST","err_code_des":"订单不存在","mch_id":"1900009851","nonce_str":"kSGYwLY4WNZvw91Y","result_code":"FAIL","return_code":"SUCCESS","return_msg":"OK","sign":"CC8F6CD5E5CADB15EEECEAA1DB4791FF"}
}

func TestRefund_Query(t *testing.T) {
	r := wego.GetPayment().Refund().QueryByOutTradeNumber(out_trade_no)
	log.Println(r.ToMap())
	// {"appid":"wx426b3015555a46be","err_code":"REFUNDNOTEXIST","err_code_des":"not exist","mch_id":"1900009851","nonce_str":"QBHv3JDrQ21HsOrG","result_code":"FAIL","return_code":"SUCCESS","return_msg":"OK","sign":"5EB4E68C6DA23B4E9EBEF07A069792BF"}

}

func TestSecurity_GetPublicKey(t *testing.T) {
	m := wego.GetSecurity().GetPublicKey()
	log.Println(m.ToMap())
}
