package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestRefund_Refund(t *testing.T) {
	m := make(wego.Map)
	m.Set("out_trade_no", out_trade_no)
	r := wego.GetRefund().Refund(out_trade_no, 1, 1, m)
	log.Println(string(r.ToJson()))
	//{"appid":"wx426b3015555a46be","err_code":"ORDERNOTEXIST","err_code_des":"订单不存在","mch_id":"1900009851","nonce_str":"kSGYwLY4WNZvw91Y","result_code":"FAIL","return_code":"SUCCESS","return_msg":"OK","sign":"CC8F6CD5E5CADB15EEECEAA1DB4791FF"}
}

func TestRefund_Query(t *testing.T) {
	m := make(wego.Map)
	m.Set("out_trade_no", out_trade_no)
	r := wego.GetRefund().Query(m)
	log.Println(string(r.ToJson()))
	// {"appid":"wx426b3015555a46be","err_code":"REFUNDNOTEXIST","err_code_des":"not exist","mch_id":"1900009851","nonce_str":"QBHv3JDrQ21HsOrG","result_code":"FAIL","return_code":"SUCCESS","return_msg":"OK","sign":"5EB4E68C6DA23B4E9EBEF07A069792BF"}

}
