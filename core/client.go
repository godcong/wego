package core

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

type Client interface {
	HttpGet(url string, m Map) Map
	HttpPost(url string, m Map) Map
	HttpPostJson(url string, m Map, query Map) Map
	Request(url string, params Map, method string, options Map) Map
	RequestRaw(url string, params Map, method string, options Map) []byte
	SafeRequest(url string, params Map, method string, options Map) Map
	Link(string) string
}

type client struct {
	Config
	app         *Application
	accessToken *AccessToken
	request     *Request
	response    *Response
	client      *http.Client
	transport   *http.Transport
	uri         string
}

func (c *client) HttpPostJson(url string, data Map, query Map) Map {
	return c.Request(url, nil, "post", Map{"query": query, "json": data})
}

func (c *client) HttpGet(url string, m Map) Map {
	return c.Request(url, nil, "get", Map{"query": m})
}

func (c *client) HttpPost(url string, m Map) Map {
	return c.Request(url, nil, "post", Map{"form_params": m})
}

func (c *client) Request(url string, params Map, method string, options Map) Map {
	panic("todo")
	//resp := request(c, buildTransport(c.Config), url, params, method, options)
	//if strings.Index(string(resp), "<xml>") == 0 {
	//	return XmlToMap(resp)
	//}
	//return JsonToMap(resp)
}

func (c *client) RequestRaw(url string, params Map, method string, options Map) []byte {
	panic("todo")

	//return request(c, buildTransport(c.Config), url, params, method, options)
}

func (c *client) SafeRequest(url string, params Map, method string, options Map) Map {
	panic("todo")

	//resp := request(c, buildSafeTransport(c.Config), url, params, method, options)
	//return XmlToMap(resp)
}

func (c *client) Link(url string) string {
	if c.GetBool("sandbox") {
		return DomainUrl() + SANDBOX_URL_SUFFIX + url
	}
	return DomainUrl() + url
}

func NewClient(request *Request, config Config) Client {
	return &client{
		request: request,
		Config:  config,
	}
}

func buildTransport(config Config) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			//Dial: (&net.Dialer{
			//	Timeout:   30 * time.Second,
			//	KeepAlive: 30 * time.Second,
			//}).Dial,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Proxy: nil,
			//TLSHandshakeTimeout:   10 * time.Second,
			//ResponseHeaderTimeout: 10 * time.Second,
			//ExpectContinueTimeout: 1 * time.Second,
		},
		//CheckRedirect: nil,
		//Jar:           nil,
		//Timeout:       0,
	}

}

func buildSafeTransport(config Config) *http.Client {
	cert, err := tls.LoadX509KeyPair(config.Get("cert_path"), config.Get("key_path"))
	if err != nil {
		Println(err)
		return nil
	}

	caFile, err := ioutil.ReadFile(config.Get("rootca_path"))
	if err != nil {
		Println(err)
		return nil
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caFile)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	}
	tlsConfig.BuildNameToCertificate()
	return &http.Client{
		Transport: &http.Transport{
			//Dial: (&net.Dialer{
			//	Timeout:   30 * time.Second,
			//	KeepAlive: 30 * time.Second,
			//}).Dial,
			TLSClientConfig: tlsConfig,
			Proxy:           nil,
			//TLSHandshakeTimeout:   10 * time.Second,
			//ResponseHeaderTimeout: 10 * time.Second,
			//ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

func request(c *client, transport *http.Transport, url string, params Map, method string, op Map) *Response {
	op = MapNilMake(op)
	if params != nil {
		params.Set("mch_id", c.Get("mch_id"))
		params.Set("nonce_str", GenerateUUID())
		params.Set("sub_mch_id", c.Get("sub_mch_id"))
		params.Set("sub_appid", c.Get("sub_appid"))
		params.Set("sign_type", SIGN_TYPE_MD5.String())
		params.Set("sign", GenerateSignature(params, c.Get("aes_key"), SIGN_TYPE_MD5))
	}
	op["body"] = params
	if r := c.request.PerformRequest(url, method, op); r.Error() == nil {

	}
	return &Response{}
}
