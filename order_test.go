package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestOrder_Query(t *testing.T) {
	//m := make(wego.Map)
	//m.Set("out_trade_no", out_trade_no)
	r := wego.GetOrder().QueryByOutTradeNumber(out_trade_no)
	log.Println(string(r.ToJson()))
	// {"appid":"wx426b3015555a46be","attach":"","mch_id":"1900009851","nonce_str":"lJhbZ9dwP4Pd5aKm","out_trade_no":"201813091059590000003433-asd002","result_code":"SUCCESS","return_code":"SUCCESS","return_msg":"OK","sign":"2F60EDECAAC5F139A82570B6724AA941","trade_state":"CLOSED","trade_state_desc":"订单已关闭"}

}

func TestOrder_Close(t *testing.T) {
	r := wego.GetOrder().Close(out_trade_no)
	log.Println(string(r.ToJson()))

}

func TestOrder_QueryByOutTradeNumber(t *testing.T) {
	r := wego.GetOrder().QueryByOutTradeNumber(out_trade_no)
	log.Println(r)
}

func TestOrder_QueryByTransactionId(t *testing.T) {
	r := wego.GetOrder().QueryByTransactionId("123")
	log.Println(r)
}

func TestOrder_Unify(t *testing.T) {
	m := make(wego.Map)
	m.Set("body", "腾讯充值中心-QQ会员充值")
	m.Set("out_trade_no", out_trade_no)
	//m.Set("device_info", "")
	////m.Set("fee_type", "CNY")
	m.Set("total_fee", "1")
	////m.Set("spbill_create_ip", "123.12.12.123")
	//m.Set("notify_url", "https://test.letiantian.me/wxpay/notify")
	m.Set("trade_type", "NATIVE")
	////m.Set("product_id", "12")
	r := wego.GetOrder().Unify(m)
	log.Println(string(r.ToJson()))
	//{"appid":"wx426b3015555a46be","code_url":"weixin://wxpay/bizpayurl?pr=D3sNT8y","mch_id":"1900009851","nonce_str":"FRFByNNdrzRuEGkp","prepay_id":"wx20180220113507842dff20340601189342","result_code":"SUCCESS","return_code":"SUCCESS","return_msg":"OK","sign":"D398DA0644A14D0BC00A8B82D8D4ECDC","trade_type":"NATIVE"}

}
