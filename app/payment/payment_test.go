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
	cfg := core.C(util.Map{
		"app_id": "wxxxxxxxxxxxxxxx",
		"mch_id": "150000000000",                 //商户ID
		"key":    "aTKnSUcTkbaaaaaaaaaaaaaaaaaa", //支付key

		"notify_url": "https://host.address/uri", //支付回调地址

		//如需使用敏感接口（如退款、发送红包等）需要配置 API 证书路径(登录商户平台下载 API 证书)
		"cert_path": "cert/apiclient_cert.pem", //支付证书地址
		"key_path":  "cert/apiclient_key.pem",  //支付证书地址

		//银行转账功能
		"rootca_path": "cert/rootca.pem",     //(可不填)
		"pubkey_path": "cert/publickey.pem",  //(可不填)部分支付使用（如:银行转账）
		"prikey_path": "cert/privatekey.pem", //(可不填)部分支付使用（如:银行转账）
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

// TestValidateSign ...
func TestValidateSign(t *testing.T) {
	data := `<xml><mch_id>1498009232</mch_id><transaction_id>4200000155201805096015992498</transaction_id><cash_fee>200</cash_fee><fee_type>CNY</fee_type><sign>BE9EA07614C09FA73A683071877D9DDB</sign><time_end>20180509175821</time_end><out_trade_no>8195400821515968</out_trade_no><result_code>SUCCESS</result_code><nonce_str>7cda1edf536f11e88cb200163e04155d</nonce_str><return_code>SUCCESS</return_code><total_fee>200</total_fee><appid>wx1ad61aeef1903b93</appid><bank_type>CMB_DEBIT</bank_type><trade_type>JSAPI</trade_type><is_subscribe>N</is_subscribe><openid>oE_gl0bQ7iJ2g3OBMQPWRiBSoiks</openid></xml>`
	m := util.XMLToMap([]byte(data))
	//m := make(util.Map)
	//xml.Unmarshal([]byte(data), &m)
	log.Debug(m)
	t.Log(payment.ValidateSign(m, "O9aVVkSTmgJK4qSibhSYpGQzRbZ2NQSJ"))
	rlt := payment.SUCCESS()
	t.Log(string(rlt.ToXML()))
}
