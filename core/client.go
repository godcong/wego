package core

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

type Client interface {
	HttpClient() *http.Client
	HttpGet(url string, m Map) *Response
	HttpPost(url string, m Map) *Response
	HttpPostJson(url string, m Map, query Map) *Response
	Request(url string, params Map, method string, options Map) *Response
	RequestRaw(url string, params Map, method string, options Map) *Response
	SafeRequest(url string, params Map, method string, options Map) *Response
	Link(string) string
}

type client struct {
	Config
	app         *Application
	accessToken *accessToken
	request     *Request
	response    *Response
	client      *http.Client
	uri         string
}

func (c *client) HttpClient() *http.Client {
	return c.client
}

func (c *client) HttpPostJson(url string, data Map, query Map) *Response {
	return c.Request(url, nil, "post", Map{REQUEST_TYPE_QUERY.String(): query, REQUEST_TYPE_JSON.String(): data})
}

func (c *client) HttpGet(url string, m Map) *Response {
	return c.Request(url, nil, "get", Map{REQUEST_TYPE_QUERY.String(): m})
}

func (c *client) HttpPost(url string, m Map) *Response {

	return c.Request(url, nil, "post", Map{REQUEST_TYPE_FORM_PARAMS.String(): m})
}

func (c *client) Request(url string, params Map, method string, options Map) *Response {
	c.client = buildTransport(c.Config)
	resp := request(c, url, params, method, options)
	c.response = resp
	return resp
}

func (c *client) RequestRaw(url string, params Map, method string, options Map) *Response {
	return c.Request(url, params, method, options)
}

func (c *client) SafeRequest(url string, params Map, method string, options Map) *Response {
	c.client = buildSafeTransport(c.Config)
	resp := request(c, url, params, method, options)
	c.response = resp
	return resp
}

func (c *client) Link(url string) string {
	if c.GetBool("sandbox") {
		return DomainUrl() + SANDBOX_URL_SUFFIX + url
	}
	return DomainUrl() + url
}

func NewClient(config Config) Client {
	return &client{
		request: DefaultRequest,
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

func request(c *client, url string, params Map, method string, op Map) *Response {
	op = MapNilMake(op)
	if params != nil {
		params.Set("mch_id", c.Get("mch_id"))
		params.Set("nonce_str", GenerateUUID())
		params.Set("sub_mch_id", c.Get("sub_mch_id"))
		params.Set("sub_appid", c.Get("sub_appid"))
		params.Set("sign_type", SIGN_TYPE_MD5.String())
		params.Set("sign", GenerateSignature(params, c.Get("aes_key"), SIGN_TYPE_MD5))
	}
	op[REQUEST_TYPE_FORM_PARAMS.String()] = params

	if r := c.request.PerformRequest(url, method, op); r.Error() == nil {
		return ParseClient(c.HttpClient(), r)
	} else {
		return ErrorResponse(r.Error())
	}
}
