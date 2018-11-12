package wego_test

import (
	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/core"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"testing"

	"github.com/godcong/wego"
	_ "github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/util"
)

var out_trade_no = "201813091059590000003433-asd003"

var long_url = "weixin://wxpay/bizpayurl?pr=etxB4DY"

var cfg = core.C(util.Map{
	//"sandbox":   true,
	"sandbox": true,
	"app_id":  "wx3c69535993f4651d",
	"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
	"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
})

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

}

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

	t.Log(payment.ValidateSign(m, "O9aVVkSTmgJK4qSibhSYpGQzRbZ2NQSJ"))
	rlt := payment.SUCCESS()
	t.Log(string(rlt.ToXML()))
}

// TestMerchant_AddSubMerchant ...
func TestMerchant_AddSubMerchant(t *testing.T) {
	obj := payment.NewMerchant(cfg)
	resp := obj.AddSubMerchant(util.Map{
		"page_index": 1,
		"page_size":  10,
	})

	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestMerchant_QuerySubMerchantByMerchantId ...
func TestMerchant_QuerySubMerchantByMerchantId(t *testing.T) {
	obj := payment.NewMerchant(cfg)
	resp := obj.QuerySubMerchantByMerchantID("123")

	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestMerchant_QuerySubMerchantByWeChatId ...
func TestMerchant_QuerySubMerchantByWeChatId(t *testing.T) {
	obj := payment.NewMerchant(cfg)
	resp := obj.QuerySubMerchantByWeChatID("123")

	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestBill_Download ...
func TestBill_Download(t *testing.T) {
	bill := payment.NewBill(cfg)
	resp := bill.Download("20181103")
	_ = core.SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())

}

// TestBill_BatchQueryComment ...
func TestBill_BatchQueryComment(t *testing.T) {
	bill := payment.NewBill(cfg)
	resp := bill.BatchQueryComment("20181101", "20181112", 0)
	_ = core.SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestBill_DownloadFundFlow ...
func TestBill_DownloadFundFlow(t *testing.T) {
	bill := payment.NewBill(cfg)
	resp := bill.DownloadFundFlow("20181109", "Operation")
	_ = core.SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestBase_Pay ...
func TestBase_Pay(t *testing.T) {
	base := payment.NewBase(core.C(util.Map{
		"sandbox": true,
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := base.Pay(util.Map{
		"body":         "image形象店-深圳腾大- QQ公仔",
		"out_trade_no": "1217752501201407033233368018",
		"total_fee":    "888",
		"auth_code":    "120061098828009406",
	})
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestBase_AuthCodeToOpenid ...
func TestBase_AuthCodeToOpenid(t *testing.T) {
	base := payment.NewBase(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"mch_id": "1516796851",
	}))
	resp := base.AuthCodeToOpenid("1212121")
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestNewTransfer ...
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
