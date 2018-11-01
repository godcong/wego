package payment_test

import (
	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"testing"
)

// TestNewPayment ...
func TestNewPayment(t *testing.T) {
	cfg, _ := core.C(util.Map{
		"app_id": "wxxxxxxxxxxxxxxx",
		"mch_id": "150000000000",                 //商户ID
		"key":    "aTKnSUcTkbaaaaaaaaaaaaaaaaaa", //支付key

		"notify_url": "https://host.address/uri", //支付回调地址

		//如需使用敏感接口（如退款、发送红包等）需要配置 API 证书路径(登录商户平台下载 API 证书)
		"cert_path": "cert/apiclient_cert.pem", //支付证书地址
		"key_path":  "cert/apiclient_key.pem",  //支付证书地址

		//银行转账功能
		"rootca_path": "cert/rootca.pem",     //(可不填)
		"pubkey_path": "cert/publickey.pem",  //(可不填)部分支付使用（如：银行转账）
		"prikey_path": "cert/privatekey.pem", //(可不填)部分支付使用（如：银行转账）
	})

	payment := payment.NewPayment(cfg)
	m := make(util.Map)
	m.Set("body", "腾讯充值中心-QQ会员充值")
	m.Set("out_trade_no", "123456")
	m.Set("total_fee", "1")
	m.Set("trade_type", "NATIVE")
	r := payment.Order().Unify(m)
	if r.Error() != nil {
		t.Log(r)
	}
	log.Println(string(r.Bytes()))
	log.Println(r.ToMap())
}
