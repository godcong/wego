package core

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client interface {
	HttpClient() *http.Client
	SetHttpClient(client *http.Client) Client
	DataType() DataType
	SetDataType(dataType DataType) Client
	URL() string
	SetDomain(domain Domain) Client
	HttpGet(url string, m Map) *Response
	HttpPost(url string, m Map) *Response
	HttpPostJson(url string, m Map, query Map) *Response
	Request(url string, params Map, method string, options Map) *Response
	RequestRaw(url string, params Map, method string, options Map) *Response
	SafeRequest(url string, params Map, method string, options Map) *Response
	Link(string) string
}

type DataType string

const (
	DATA_TYPE_XML  DataType = "xml"
	DATA_TYPE_JSON DataType = "json"
)

type client struct {
	dataType    DataType
	domain      Domain
	app         *Application
	accessToken *accessToken
	request     *Request
	response    *Response
	client      *http.Client
	Config
}

func (c *client) SetDomain(domain Domain) Client {
	c.domain = domain
	return c
}

func (c *client) URL() string {
	return c.domain.URL()
}

func (c *client) HttpClient() *http.Client {
	return c.client
}

func (c *client) SetHttpClient(client *http.Client) Client {
	c.client = client
	return c
}

func (c *client) DataType() DataType {
	return c.dataType
}

func (c *client) SetDataType(dataType DataType) Client {
	c.dataType = dataType
	return c
}

func (c *client) HttpPostJson(url string, data Map, ops Map) *Response {
	ops = MapNilMake(ops)
	if c.dataType == DATA_TYPE_JSON {
		ops.Set(REQUEST_TYPE_JSON.String(), data)
	} else {
		ops.Set(REQUEST_TYPE_XML.String(), data)
	}

	return c.Request(url, nil, "post", ops)
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
	Debug("SafeRequest|httpClient", c.client)
	c.response = request(c, url, params, method, options)
	return c.response
}

func (c *client) Link(uri string) string {
	if c.GetBool("sandbox") {
		return c.URL() + SANDBOX_URL_SUFFIX + uri
	}
	return c.URL() + uri
}

func NewClient(config Config) Client {
	return &client{
		request:  DefaultRequest,
		Config:   config,
		dataType: DATA_TYPE_XML,
		domain:   NewDomain("default"),
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
		panic(err)
	}

	caFile, err := ioutil.ReadFile(config.Get("rootca_path"))
	if err != nil {
		panic(err)
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
	Debug("request", c, url, params, method, op)
	op = MapNilMake(op)
	if params != nil {
		params.Set("mch_id", c.Get("mch_id"))
		params.Set("nonce_str", GenerateUUID())
		params.Set("sub_mch_id", c.Get("sub_mch_id"))
		params.Set("sub_appid", c.Get("sub_appid"))
		params.Set("sign_type", SIGN_TYPE_MD5.String())
		params.Set("sign", GenerateSignature(params, c.Get("aes_key"), SIGN_TYPE_MD5))
	}

	data := toRequestData(c, params, op)

	if r := c.request.PerformRequest(url, method, data); r.Error() == nil {
		return ClientDo(c, r)
	} else {
		return ErrorResponse(r.Error())
	}
}

func toRequestData(client *client, p, op Map) *RequestData {
	data := client.request.RequestDataCopy()
	data.Query = processQuery(op.Get(REQUEST_TYPE_QUERY.String()))
	if client.DataType() == DATA_TYPE_JSON {
		data.SetHeaderJson()
		data.Body = bytes.NewReader(p.ToJson())
	}
	if client.DataType() == DATA_TYPE_XML {
		data.SetHeaderXml()
		data.Body = strings.NewReader(p.ToXml())
	}

	return data
}

func processFormParams(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case Map:
		return v.ToXml()
	}
	return ""
}
func processXml(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case Map:
		return v.ToXml()
	}
	return ""
}

func processJson(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case Map:
		return string(v.ToJson())
	}
	return ""
}

func processQuery(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case Map:
		return v.UrlEncode()
	}
	return ""
}
