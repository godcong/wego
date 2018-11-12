package payment_test

import (
	"errors"
	"github.com/godcong/wego"
	"github.com/godcong/wego/util"
	"net/http"
	"testing"
)

func dummyQuery(s string) util.Map {
	return nil
}

func dummySave(p util.Map) {

}

// TestScannedNotify_ServeHTTP ...
func TestScannedNotify_ServeHTTP(t *testing.T) {
	ScannedCallbackFunction := func(p util.Map) (maps util.Map, e error) {
		// 使用通知里的 "微信支付订单号" 或者 "商户订单号" 去自己的数据库找到订单
		order := dummyQuery(p.GetString("out_trade_no"))         //通过out_trade_no查询订单,dummyQuery为查询订单函数
		if order != nil || order.GetString("status") != "paid" { // 如果订单不存在 或者 订单已经支付过了
			return nil, nil // 告诉微信，我已经处理完了，订单没找到，别再通知我了
		}

		if rc := p.GetString("return_code"); rc == "SUCCESS" { // return_code 表示通信状态，不代表支付状态
			// 用户是否支付成功
			if p.GetString("result_code") == "SUCCESS" {
				order.Set("paid", util.Time()) // 更新支付时间为当前时间
				order.Set("status", "paid")

				// 用户支付失败
			} else if rc == "FAIL" {
				order.Set("status", "pay_failed")
			}
		} else {
			return nil, errors.New("失败，请稍后再试")
		}
		dummySave(order)
		return nil, nil
	}

	serve1 := wego.Payment().HandleScannedNotify(ScannedCallbackFunction).ServeHTTP

	serve2 := wego.Payment().HandleScanned(ScannedCallbackFunction)

	http.HandleFunc("/scanned/callback/address", serve1)
	http.HandleFunc("/scanned/callback/address2", serve2)

	t.Fatal(http.ListenAndServe(":8080", nil))
}
