package wxpay_test

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/godcong/wopay/wxpay"
)

func TestPay_RequestWithoutCert(t *testing.T) {
	//log.Println(time.Now().Unix())
	//log.Println(time.Now().UnixNano() / 1000000)
	//log.Println(time.Now().UnixNano() / time.Millisecond.Nanoseconds())
	//log.Println(time.Now().UnixNano() / time.Second.Nanoseconds())
}

//1503909263813204
//1349333576093
//1503909402343
func TestNewPayRequest(t *testing.T) {
	reqBody := "<xml><body>测试商家-商品类目</body><trade_type>NATIVE</trade_type><mch_id>11473623</mch_id><sign_type>HMAC-SHA256</sign_type><nonce_str>b1089cb0231011e7b7e1484520356fdc</nonce_str><detail /><fee_type>CNY</fee_type><device_info>WEB</device_info><out_trade_no>20161909105959000000111108</out_trade_no><total_fee>1</total_fee><appid>wxab8acb865bb1637e</appid><notify_url>http://test.letiantian.com/wxpay/notify</notify_url><sign>78F24E555374B988277D18633BF2D4CA23A6EAF06FEE0CF1E50EA4EADEEC41A3</sign><spbill_create_ip>123.12.12.123</spbill_create_ip></xml>"

	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", strings.NewReader(reqBody))
	if err != nil {
		return
	}
	req.Host = "api.mch.weixin.qq.com"

	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Println(string(str))

}

func TestPayRequest_RequestOnce(t *testing.T) {
	reqBody := "<xml><body>测试商家-商品类目</body><trade_type>NATIVE</trade_type><mch_id>11473623</mch_id><sign_type>HMAC-SHA256</sign_type><nonce_str>b1089cb0231011e7b7e1484520356fdc</nonce_str><detail /><fee_type>CNY</fee_type><device_info>WEB</device_info><out_trade_no>20161909105959000000111108</out_trade_no><total_fee>1</total_fee><appid>wxab8acb865bb1637e</appid><notify_url>http://test.letiantian.com/wxpay/notify</notify_url><sign>78F24E555374B988277D18633BF2D4CA23A6EAF06FEE0CF1E50EA4EADEEC41A3</sign><spbill_create_ip>123.12.12.123</spbill_create_ip></xml>"

	resp, err := wxpay.NewPayRequest(wxpay.NewPayConfig()).RequestOnce(wxpay.DOMAIN_API, wxpay.UNIFIEDORDER_URL_SUFFIX, wxpay.GenerateUUID(), reqBody, 0, 0, false)
	if err != nil {
		return
	}
	log.Println(resp)
}
