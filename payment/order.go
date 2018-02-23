package payment

import "github.com/godcong/wego/core"

type Order struct {
	core.Config
	*Payment
}

func (o *Order) Unify(m core.Map) core.Map {
	if !m.Has("spbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("spbill_create_ip", core.GetServerIp())
		}
		//TODO: getclientip with request
	}

	m.Set("appid", o.Get("app_id"))
	//$params['notify_url'] = $params['notify_url'] ?? $this->app['config']['notify_url'];
	if !m.Has("notify_url") {
		m.Set("notify_url", o.Get("notify_url"))
	}
	return o.Request(core.UNIFIEDORDER_URL_SUFFIX, m)
}

//func (o *order) request(url string, m Map) Map {
//	return o.Request(o.Link(url), m, "post", nil)
//}

/**
* 作用：关闭订单
* 场景：公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @return PayData, error API返回数据
 */
func (o *Order) Close(no string) core.Map {
	m := make(core.Map)
	m.Set("appid", o.Get("app_id"))
	m.Set("out_trade_no", no)
	return o.Request(core.CLOSEORDER_URL_SUFFIX, m)
}

/** QueryOrder
* 作用：查询订单
* 场景：刷卡支付、公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @param connectTimeoutMs 连接超时时间，单位是毫秒
* @param readTimeoutMs 读超时时间，单位是毫秒
* @return PayData, error API返回数据
 */
func (o *Order) query(m core.Map) core.Map {
	m.Set("appid", o.Get("app_id"))
	return o.Request(core.ORDERQUERY_URL_SUFFIX, m)
}

func (o *Order) QueryByTransactionId(id string) core.Map {
	return o.query(core.Map{"transaction_id": id})
}

func (o *Order) QueryByOutTradeNumber(no string) core.Map {
	return o.query(core.Map{"out_trade_no": no})
}
