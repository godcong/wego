package payment

import (
	"net/http"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Order struct {
	config.Config
	*Payment
	request *http.Request
}

func newOrder(p *Payment) *Order {
	return &Order{
		Config:  defaultConfig,
		Payment: p,
	}
}

func NewOrder() *Order {
	return newOrder(payment)
}

//SetRequest to set a http request for Unify to get the client ip
func (o *Order) SetRequest(r *http.Request) *Order {
	o.request = r
	return o
}

/*
Unify 统一下单
	接口链接
	URL地址：https://api.mch.weixin.qq.com/pay/unifiedorder

	是否需要证书
	否

	接口请求参数参考：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
 	参数：
		util.Map 向wxpay post的请求数据
	返回值：
		*net.Response wxpay应答数据
*/
func (o *Order) Unify(m util.Map) *net.Response {
	if !m.Has("spbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("spbill_create_ip", core.GetServerIp())
		}
		//TODO: getclientip with request
		if o.request != nil {
			m.Set("spbill_create_ip", core.GetClientIp(o.request))
		}
	}

	m.Set("appid", o.Config.Get("app_id"))

	//$params['notify_url'] = $params['notify_url'] ?? $this->app['config']['notify_url'];
	if !m.Has("notify_url") {
		m.Set("notify_url", o.Config.Get("notify_url"))
	}
	resp := o.Request(UNIFIEDORDER_URL_SUFFIX, m)
	resp.CheckError()
	return resp
}

/**
* 作用：关闭订单
* 场景：公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @return PayData, error API返回数据
 */
func (o *Order) Close(no string) *net.Response {
	m := make(util.Map)
	m.Set("appid", o.Config.Get("app_id"))
	m.Set("out_trade_no", no)
	resp := o.Request(CLOSEORDER_URL_SUFFIX, m)
	resp.CheckError()
	return resp
}

/** QueryOrder
* 作用：查询订单
* 场景：刷卡支付、公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @param connectTimeoutMs 连接超时时间，单位是毫秒
* @param readTimeoutMs 读超时时间，单位是毫秒
* @return PayData, error API返回数据
 */
func (o *Order) query(m util.Map) *net.Response {
	m.Set("appid", o.Config.Get("app_id"))
	return o.Request(ORDERQUERY_URL_SUFFIX, m)
}

func (o *Order) QueryByTransactionId(id string) *net.Response {
	return o.query(util.Map{"transaction_id": id})
}

func (o *Order) QueryByOutTradeNumber(no string) *net.Response {
	return o.query(util.Map{"out_trade_no": no})
}
