package wxpay_test

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/godcong/wopay/wxpay"
)

var out_trade_no = "201613091059590000003433-asd002"
var total_fee = "1"
var data wxpay.PayData = map[string]string{
	"first":  "1",
	"second": "2",
	"aecond": "3",
	"becond": "4",
}
var long_url = "weixin://wxpay/bizpayurl?pr=etxB4DY"

func TestPayData_Get(t *testing.T) {
	log.Println(data.Get("first"))
}

func TestPayData_Set(t *testing.T) {
	data.Set("third", "3")
	log.Println(data)
}

func TestPayData_IsExist(t *testing.T) {
	log.Println(data.IsExist("first"))
}

func TestUnifiedOrder(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("body", "腾讯充值中心-QQ会员充值")
	data.Set("out_trade_no", out_trade_no)
	data.Set("device_info", "")
	data.Set("fee_type", "CNY")
	data.Set("total_fee", "1")
	data.Set("spbill_create_ip", "123.12.12.123")
	data.Set("notify_url", "http://test.letiantian.me/wxpay/notify")
	data.Set("trade_type", "NATIVE")
	data.Set("product_id", "12")

	rdata, err := wxpay.UnifiedOrder(data)
	log.Println(rdata, err)

}

func TestCloseOrder(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("out_trade_no", out_trade_no)
	rdata, err := wxpay.CloseOrder(data)
	log.Println(rdata, err)
}

func TestQueryOrder(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("out_trade_no", out_trade_no)
	rdata, err := wxpay.QueryOrder(data)
	log.Println(rdata, err)
}

func TestReverseOrder(t *testing.T) {
	//cert, _ := ioutil.ReadFile(`D:\Godcong\Workspace\g7n3\src\github.com\godcong\wopay\wx\cert\apiclient_cert.p12`)

	data := make(wxpay.PayData)
	data.Set("out_trade_no", out_trade_no)
	rdata, err := wxpay.ReverseOrder(data)
	log.Println(rdata, err)
}

func TestRefund(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("out_trade_no", out_trade_no)
	data.Set("out_refund_no", out_trade_no)
	data.Set("total_fee", total_fee)
	data.Set("refund_fee", total_fee)
	data.Set("refund_fee_type", "CNY")
	data.Set("op_user_id", wxpay.PayConfigInstance().MchID())
	rdata, err := wxpay.Refund(data)
	log.Println(rdata, err)
}

func TestShortUrl(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("long_url", long_url)
	log.Println(wxpay.ShortUrl(data))
}

func TestUnifiedOrderSpeed(t *testing.T) {
	for i := 0; i < 100; i++ {
		sta := wxpay.CurrentTimeStampNS()
		out_trade_no = out_trade_no + strconv.FormatInt(int64(i), 10)
		TestUnifiedOrder(t)
		end := wxpay.CurrentTimeStampNS()
		log.Println(end - sta)
		time.Sleep(1000)
	}

}

func TestQueryRefund(t *testing.T) {
	data := wxpay.PayData{"out_refund_no": out_trade_no}
	r, e := wxpay.QueryRefund(data)
	if r.Get("return_msg") == "OK" {
		log.Println(r, e)
	}

}

func TestDownloadBill(t *testing.T) {
	data := wxpay.PayData{
		"bill_date": "20170913",
		"bill_type": "ALL",
	}
	r, e := wxpay.DownloadBill(data)

	log.Println(r, e)

}

//untest
func TestMicroPay(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("body", "腾讯充值中心-QQ会员充值")
	data.Set("out_trade_no", out_trade_no)
	data.Set("device_info", "")
	data.Set("fee_type", "CNY")
	data.Set("total_fee", "1")
	data.Set("spbill_create_ip", "123.12.12.123")
	data.Set("notify_url", "http://test.letiantian.me/wxpay/notify")
	data.Set("trade_type", "NATIVE")
	data.Set("product_id", "12")

	rdata, err := wxpay.MicroPay(data)
	log.Println(rdata, err)
}

//untest
func TestMicroPayWithPos(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("body", "腾讯充值中心-QQ会员充值")
	data.Set("out_trade_no", out_trade_no)
	data.Set("device_info", "")
	data.Set("fee_type", "CNY")
	data.Set("total_fee", "1")
	data.Set("spbill_create_ip", "123.12.12.123")
	data.Set("notify_url", "http://test.letiantian.me/wxpay/notify")
	data.Set("trade_type", "NATIVE")
	data.Set("product_id", "12")

	rdata, err := wxpay.MicroPayWithPos(data)
	log.Println(rdata, err)
}

//untest
func TestAuthCodeToOpenid(t *testing.T) {
	data := make(wxpay.PayData)
	data.Set("body", "腾讯充值中心-QQ会员充值")
	data.Set("out_trade_no", out_trade_no)
	data.Set("device_info", "")
	data.Set("fee_type", "CNY")
	data.Set("total_fee", "1")
	data.Set("spbill_create_ip", "123.12.12.123")
	data.Set("notify_url", "http://test.letiantian.me/wxpay/notify")
	data.Set("trade_type", "NATIVE")
	data.Set("product_id", "12")

	rdata, err := wxpay.AuthCodeToOpenid(data)
	log.Println(rdata, err)
}
