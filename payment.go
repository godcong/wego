package wego

//
//import (
//	"log"
//	"strings"
//	"time"
//	"github.com/godcong/wopay/util"
//)
//
//type PayData = util.PayData
//
//type Pay struct {
//	config     PayConfig
//	payRequest *PayRequest
//	signType   SignType
//	autoReport bool
//	useSanBox  bool
//	notifyUrl  string
//}
//
//type RequestFunc func(url string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error)
//
//func MicroPay(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.MicroPay(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func MicroPayWithPos(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.MicroPayWithPos(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
////UnifiedOrder
//func UnifiedOrder(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.UnifiedOrder(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
////CloseOrder
//func CloseOrder(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.CloseOrder(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
////QueryOrder
//func QueryOrder(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.QueryOrder(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
////ReverseOrder
//func ReverseOrder(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.ReverseOrder(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
////QueryRefund
//func QueryRefund(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.QueryRefund(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
////Refund
//func Refund(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.Refund(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
////ShortUrl
//func ShortUrl(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.ShortUrl(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func DownloadBill(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.DownloadBill(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func Report(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.Report(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func AuthCodeToOpenid(data PayData) (PayData, error) {
//	pay := NewPay(PayConfigInstance())
//	data, err := pay.AuthCodeToOpenid(data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//

type payment struct {
	Config
	request Request
	sandbox Sandbox
}

func (p *payment) InSandbox() bool {
	return p.Config.GetBool("sandbox")
}

func (p *payment) GetKey(s string) string {
	if s == SANDBOX_SIGNKEY_URL_SUFFIX {
		return p.Config.Get("aes_key")
	}
	if p.InSandbox() {
		p.sandbox.GetKey()
	}
	return p.Config.Get("aes_key")
}

//func (p *payment) GetRequest() Request {
//	return p.request
//}

func (p *payment) Link(url string) string {
	if p.GetBool("sandbox") {
		return DomainUrl() + SANDBOX_URL_SUFFIX + url
	}
	return DomainUrl() + url
}

type Payment interface {
	UnifiedOrder(m Map) (Map, error)
	Link(string) string
	GetKey(s string) string
	InSandbox() bool
}

func NewPayment(config Config) Payment {
	c := config
	if config == nil {
		c = GetConfig("payment.default")
	}
	return &payment{
		Config:  c,
		request: NewRequest(c),
		sandbox: NewSandbox(c),
	}
}

//func newPay(config Config) *payment {
//	p := payment{
//		Config: config,
//	}
//	pay.signType = SIGN_TYPE_HMACSHA256
//	if useSandbox {
//		pay.signType = SIGN_TYPE_MD5
//	}
//	pay.payRequest = NewPayRequest(config)
//	return &pay
//}

//func (pay *Pay) SetSandBox(useSandbox bool) *Pay {
//	pay.signType = SIGN_TYPE_HMACSHA256
//	if useSandbox {
//		pay.signType = SIGN_TYPE_MD5
//	}
//	pay.useSanBox = useSandbox
//	return pay
//}
//
//func (pay *Pay) ApplySandBox(url string) string {
//	if pay.useSanBox {
//		return SANDBOX_URL_SUFFIX + url
//	}
//	return url
//}
//
//func (pay *Pay) RequestWithoutCert(url string, data PayData) (string, error) {
//	msgUUID := data.Get("nonce_str")
//	reqBody, err := util.MapToXml(data)
//	if err != nil {
//		return "", err
//	}
//	resp, err := pay.payRequest.RequestWithoutCert(url, msgUUID, reqBody, pay.autoReport)
//	return resp, err
//}
//
//func (pay *Pay) RequestWithoutCertTimeout(url string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error) {
//	msgUUID := data.Get("nonce_str")
//	reqBody, err := util.MapToXml(data)
//	if err != nil {
//		return "", err
//	}
//	resp, err := pay.payRequest.RequestWithoutCertTimeout(url, msgUUID, reqBody, connectTimeoutMs, readTimeoutMs, pay.autoReport)
//	return resp, err
//}
//
//func (pay *Pay) RequestWithCert(url string, data PayData) (string, error) {
//	msgUUID := data.Get("nonce_str")
//	reqBody, err := util.MapToXml(data)
//	if err != nil {
//		return "", err
//	}
//	resp, err := pay.payRequest.RequestWithCert(url, msgUUID, reqBody, pay.autoReport)
//	return resp, err
//}
//
//func (pay *Pay) RequestWithCertTimeout(url string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error) {
//	msgUUID := data.Get("nonce_str")
//	reqBody, err := util.MapToXml(data)
//	if err != nil {
//		return "", err
//	}
//	resp, err := pay.payRequest.RequestWithCertTimeout(url, msgUUID, reqBody, connectTimeoutMs, readTimeoutMs, pay.autoReport)
//	return resp, err
//}
//
//func (pay *Pay) fillRequest(requestFunc RequestFunc, data PayData, suffix string) (string, error) {
//	return pay.fillRequestTimeout(requestFunc, suffix, data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
//func (pay *Pay) fillRequestTimeout(requestFunc RequestFunc, suffix string, data PayData, connectTimeoutMs, readTimeoutMs int) (string, error) {
//	usb := pay.ApplySandBox(suffix)
//	m, err := pay.FillRequestData(data)
//	if err != nil {
//		return "", err
//	}
//	return requestFunc(usb, m, connectTimeoutMs, readTimeoutMs)
//}
//perform
//func (pay *Pay) FillRequestData(data PayData) (PayData, error) {
//	data.Set("appid", pay.config.AppID())
//	data.Set("mch_id", pay.config.MchID())
//	data.Set("nonce_str", util.GenerateUUID())
//	data.Set("sign_type", pay.signType.ToString())
//	sign, e := GenerateSignature(data, pay.config.Key(), pay.signType)
//	if e != nil {
//		return nil, e
//	}
//	data.Set("sign", sign)
//	return data, nil
//}
//
///** UnifiedOrder
//* 作用：统一下单
//* 场景：公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) UnifiedOrder(data PayData) (PayData, error) {
//	return pay.unifiedOrder(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** UnifiedOrder
//* 作用：统一下单
//* 场景：公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) UnifiedOrderTimeout(data PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
//	return pay.unifiedOrder(data, connectTimeoutMs, readTimeoutMs)
//}
//
func (p *payment) UnifiedOrder(m Map) (Map, error) {
	//if (empty($params['spbill_create_ip'])) {
	//$params['spbill_create_ip'] = ('NATIVE' === $params['trade_type']) ? Support\get_server_ip() : Support\get_client_ip();
	//}
	if !m.Has("pbill_create_ip") {

	}

	m.Set("appid", p.Config.Get("app_id"))
	//$params['notify_url'] = $params['notify_url'] ?? $this->app['config']['notify_url'];
	if !m.Has("notify_url") {
		m.Set("notify_url", p.Config.Get("notify_url"))
	}

	resp, err := p.Request(p.Link(UNIFIEDORDER_URL_SUFFIX), m)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (p *payment) Request(url string, m Map) ([]byte, error) {
	m.Set("mch_id", p.Config.Get("mch_id"))
	m.Set("nonce_str", GenerateUUID())
	//m.Set("sub_mch_id", p.Config.Get("sub_mch_id"))
	//m.Set("sub_appid", p.Config.Get("sub_appid"))

	m.Set("sign_type", SIGN_TYPE_MD5.String())
	sign, e := GenerateSignature(m, p.Config.Get("aes_key"), SIGN_TYPE_MD5)
	if e != nil {
		return nil, e
	}
	m.Set("sign", sign)
	Println(m)
	return p.request.Request(url, m)
}

//
///**
//* 作用：关闭订单
//* 场景：公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) CloseOrder(data PayData) (PayData, error) {
//	return pay.closeOrder(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** CloseOrderTimeout
//* 作用：关闭订单
//* 场景：公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) CloseOrderTimeout(data PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
//	return pay.closeOrder(data, connectTimeoutMs, readTimeoutMs)
//}
//
//func (pay *Pay) closeOrder(data PayData, connectTimeoutMs, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, CLOSEORDER_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//	return util.XmlToMap(resp), nil
//}
//
///** QueryOrder
//* 作用：查询订单
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) QueryOrder(data PayData) (PayData, error) {
//	return pay.queryOrder(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** QueryOrder
//* 作用：查询订单
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) QueryOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.queryOrder(data, connectTimeoutMs, readTimeoutMs)
//}
//func (pay *Pay) queryOrder(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, ORDERQUERY_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//	return util.XmlToMap(resp), nil
//}
//
///** ReverseOrder
//* 作用：撤销订单
//* 场景：刷卡支付
//* @param data 向wxpay post的请求数据
//* @return API返回数据
// */
//func (pay *Pay) ReverseOrder(data PayData) (PayData, error) {
//	return pay.reverseOrder(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** ReverseOrderTimeout
//* 作用：撤销订单
//* 场景：刷卡支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return API返回数据
// */
//func (pay *Pay) ReverseOrderTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.reverseOrder(data, connectTimeoutMs, readTimeoutMs)
//}
//func (pay *Pay) reverseOrder(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithCertTimeout, REVERSE_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//
//	return util.XmlToMap(resp), nil
//}
//
///** Refund
//* 作用：申请退款
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* 其他：需要证书
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) Refund(data PayData) (PayData, error) {
//	return pay.refund(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** RefundTimeout
//* 作用：申请退款
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* 其他：需要证书
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) RefundTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.refund(data, connectTimeoutMs, readTimeoutMs)
//}
//func (pay *Pay) refund(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithCertTimeout, REFUND_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//
//	return util.XmlToMap(resp), nil
//}
//
///** ShortUrl
//* 作用：转换短链接
//* 场景：刷卡支付、扫码支付
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) ShortUrl(data PayData) (PayData, error) {
//	return pay.shortUrl(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** ShortUrlTimeout
//* 作用：转换短链接
//* 场景：刷卡支付、扫码支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) ShortUrlTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.shortUrl(data, connectTimeoutMs, readTimeoutMs)
//}
//
//func (pay *Pay) shortUrl(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithCertTimeout, SHORTURL_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//	return util.XmlToMap(resp), nil
//}
//
//func (pay *Pay) QueryRefund(data PayData) (PayData, error) {
//	return pay.queryRefund(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
//func (pay *Pay) QueryRefundTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.queryRefund(data, connectTimeoutMs, readTimeoutMs)
//}
//
//func (pay *Pay) queryRefund(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, REFUNDQUERY_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//
//	return util.XmlToMap(resp), nil
//}
//
///** DownloadBill
//* 作用：对账单下载
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* 其他：无论是否成功都返回Map。若成功，返回的Map中含有return_code、return_msg、data，
//*      其中return_code为`SUCCESS`，data为对账单数据。
//* @param data 向wxpay post的请求数据
//* @return PayData, error 经过封装的API返回数据
// */
//func (pay *Pay) DownloadBill(data PayData) (PayData, error) {
//	return pay.downloadBill(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** DownloadBillTimeout
//* 作用：对账单下载
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* 其他：无论是否成功都返回Map。若成功，返回的Map中含有return_code、return_msg、data，
//*      其中return_code为`SUCCESS`，data为对账单数据。
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error 经过封装的API返回数据
// */
//func (pay *Pay) DownloadBillTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.downloadBill(data, connectTimeoutMs, readTimeoutMs)
//}
//func (pay *Pay) downloadBill(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, DOWNLOADBILL_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//	var ret PayData
//	if strings.Index(resp, "<") == 0 {
//		ret = util.XmlToMap(resp)
//	} else {
//		ret = make(PayData)
//		ret.Set("return_code", SUCCESS)
//		ret.Set("return_msg", "ok")
//		ret.Set("data", resp)
//	}
//
//	return ret, nil
//}
//
///** Report
//* 作用：交易保障
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) Report(data PayData) (PayData, error) {
//	return pay.report(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** ReportTimeout
//* 作用：交易保障
//* 场景：刷卡支付、公共号支付、扫码支付、APP支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) ReportTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.report(data, connectTimeoutMs, readTimeoutMs)
//}
//func (pay *Pay) report(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, REPORT_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//	return util.XmlToMap(resp), nil
//}
//
///** AuthCodeToOpenid
//* 作用: 授权码查询OPENID接口
//* 场景：刷卡支付
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) AuthCodeToOpenid(data PayData) (PayData, error) {
//	return pay.authCodeToOpenid(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** AuthCodeToOpenidTimeout
//* 作用: 授权码查询OPENID接口
//* 场景：刷卡支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) AuthCodeToOpenidTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.authCodeToOpenid(data, connectTimeoutMs, readTimeoutMs)
//}
//
//func (pay *Pay) authCodeToOpenid(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, AUTHCODETOOPENID_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//	return util.XmlToMap(resp), nil
//}
//
///** MicroPay
//* 作用：提交刷卡支付
//* 场景：刷卡支付
//* @param data 向wxpay post的请求数据
//* @return PayData, error API返回数据
// */
//func (pay *Pay) MicroPay(data PayData) (PayData, error) {
//	return pay.microPay(data, pay.config.ConnectTimeoutMs(), pay.config.ReadTimeoutMs())
//}
//
///** MicroPayTimeout
//* 作用：提交刷卡支付
//* 场景：刷卡支付
//* @param data 向wxpay post的请求数据
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @param readTimeoutMs 读超时时间，单位是毫秒
//* @return PayData, error API返回数据
// */
//func (pay *Pay) MicroPayTimeout(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	return pay.microPay(data, connectTimeoutMs, readTimeoutMs)
//}
//
//func (pay *Pay) microPay(data PayData, connectTimeoutMs int, readTimeoutMs int) (PayData, error) {
//	resp, err := pay.fillRequestTimeout(pay.RequestWithoutCertTimeout, MICROPAY_URL_SUFFIX, data, connectTimeoutMs, readTimeoutMs)
//	if err != nil {
//		return nil, err
//	}
//	return util.XmlToMap(resp), nil
//}
//
///** MicroPayWithPos
//* 提交刷卡支付，针对软POS，尽可能做成功
//* 内置重试机制，最多60s
//* @param data
//* @return PayData, error
// */
//func (pay *Pay) MicroPayWithPos(data PayData) (PayData, error) {
//	return pay.microPayWithPosConnectTimeout(data, pay.config.ConnectTimeoutMs())
//}
//
///** MicroPayWithPosConnectTimeout
//* 提交刷卡支付，针对软POS，尽可能做成功
//* 内置重试机制，最多60s
//* @param data
//* @param connectTimeoutMs 连接超时时间，单位是毫秒
//* @return PayData, error
// */
//func (pay *Pay) MicroPayWithPosConnectTimeout(data PayData, connectTimeoutMs int) (PayData, error) {
//	return pay.microPayWithPosConnectTimeout(data, connectTimeoutMs)
//}
//
//func (pay *Pay) microPayWithPosConnectTimeout(data PayData, connectTimeoutMs int) (PayData, error) {
//	remainingTimeMs := 60 * 1000
//	var err error
//	var lastResult PayData
//	for {
//		startTimestampMs := util.CurrentTimeStampMS()
//		readTimeoutMs := remainingTimeMs - connectTimeoutMs
//		if readTimeoutMs > 1000 {
//			lastResult, err = pay.microPay(data, connectTimeoutMs, readTimeoutMs)
//			if err != nil {
//				goto ERROR
//			}
//			if lastResult.Get("return_code") == SUCCESS {
//				errCode := lastResult.Get("err_code")
//				if resultCode := lastResult.Get("result_code"); resultCode == SUCCESS {
//					break
//				}
//				// 看错误码，若支付结果未知，则重试提交刷卡支付
//				if errCode == SYSTEMERROR || errCode == BANKERROR || errCode == USERPAYING {
//					remainingTimeMs = remainingTimeMs - (int)(util.CurrentTimeStampMS()-startTimestampMs)
//					if remainingTimeMs <= 100 {
//						break
//					}
//					log.Println("microPayWithPos: try micropay again")
//					if remainingTimeMs > 5*1000 {
//						time.Sleep(5 * time.Second)
//					} else {
//						time.Sleep(time.Second)
//					}
//					continue
//
//				} else {
//					break
//				}
//
//			} else {
//				break
//			}
//
//		} else {
//			break
//		}
//	}
//
//	return lastResult, nil
//ERROR:
//	lastResult = nil
//	return lastResult, err
//}