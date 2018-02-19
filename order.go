package wego

type Order interface {
	Unify(m Map) (Map, error)
	Close(no string) (Map, error)
	Query(Map) (Map, error)
}

type order struct {
	Config
	app Application
}

func NewOrder(application Application) Order {
	return &order{
		Config: application.Config(),
		//client: nil,
		app: application,
	}
}

func (o *order) Unify(m Map) (Map, error) {

	if !m.Has("pbill_create_ip") {
		if m.Get("trade_type") == "NATIVE" {
			m.Set("pbill_create_ip", GetServerIp())
		}
		//TODO: getclientip with request
	}

	m.Set("appid", o.Get("app_id"))
	//$params['notify_url'] = $params['notify_url'] ?? $this->app['config']['notify_url'];
	if !m.Has("notify_url") {
		m.Set("notify_url", o.Get("notify_url"))
	}

	resp, err := o.app.Payment().Request(o.app.Link(UNIFIEDORDER_URL_SUFFIX), m)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (p *order) Request(url string, m Map) ([]byte, error) {
	m.Set("mch_id", p.Config.Get("mch_id"))
	m.Set("nonce_str", GenerateUUID())
	//m.Set("sub_mch_id", p.Config.Get("sub_mch_id"))
	//m.Set("sub_appid", p.Config.Get("sub_appid"))

	m.Set("sign_type", SIGN_TYPE_MD5.String())

	m.Set("sign", GenerateSignature(m, p.Config.Get("aes_key"), SIGN_TYPE_MD5))
	//Println(m)
	return p.request.Request(url, m)
}

/**
* 作用：关闭订单
* 场景：公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @return PayData, error API返回数据
 */
func (p *order) Close(no string) (Map, error) {
	m := make(Map)
	m.Set("appid", p.Config.Get("app_id"))
	m.Set("out_trade_no", no)
	resp, err := p.Request(p.Link(CLOSEORDER_URL_SUFFIX), m)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

/** QueryOrder
* 作用：查询订单
* 场景：刷卡支付、公共号支付、扫码支付、APP支付
* @param data 向wxpay post的请求数据
* @param connectTimeoutMs 连接超时时间，单位是毫秒
* @param readTimeoutMs 读超时时间，单位是毫秒
* @return PayData, error API返回数据
 */
func (o *order) Query(m Map) (Map, error) {
	m.Set("appid", o.Config.Get("app_id"))
	resp, err := p.Request(app.Link(ORDERQUERY_URL_SUFFIX), m)
	if err != nil {
		return nil, err
	}
	return XmlToMap(resp), nil
}

func (o *order) QueryByTransactionId(id string) (Map, error) {
	m := make(Map)
	m.Set("transaction_id", id)
	return o.Query(m)
}

func (o *order) QueryByOutTradeNumber(no string) (Map, error) {
	m := make(Map)
	m.Set("out_trade_no", no)
	return o.Query(m)
}
