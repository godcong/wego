package payment_test

import (
	"github.com/godcong/wego/app/payment"
	"testing"

	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

var backValue = `<?xml version="1.0" encoding="UTF-8" standalone="no"?><xml><body><![CDATA[微信充值]]></body><out_trade_no>3208629970201259</out_trade_no><notify_url><![CDATA[https://mp.quick58.com/charge/callback]]></notify_url><sign_type><![CDATA[MD5]]></sign_type><mch_id>1498009232</mch_id><nonce_str><![CDATA[3ba6b031626611e8904800163e04155d]]></nonce_str><sign><![CDATA[35F7C1BA75C64E4B88558D46ED5E8E4A]]></sign><openid><![CDATA[oE_gl0Yr54fUjBhU5nBlP4hS2efo]]></openid><total_fee>200</total_fee><trade_type><![CDATA[JSAPI]]></trade_type><appid><![CDATA[wx1ad61aeef1903b93]]></appid></xml>`

// TestServer_ServeHTTP ...
func TestServer_ServeHTTP(t *testing.T) {
	m := util.XMLToMap([]byte(backValue))
	s := payment.GenerateSignature(m, "O9aVVkSTmgJK4qSibhSYpGQzRbZ2NQSJ", payment.MakeSignMD5)
	log.Println(s)
}
