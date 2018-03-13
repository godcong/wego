package core

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type DataType string

const (
	DATA_TYPE_XML  DataType = "xml"
	DATA_TYPE_JSON DataType = "json"
)

type URL struct {
	token  *AccessToken
	client *Client
}

type Client struct {
	Config
	dataType DataType
	domain   Domain
	app      *Application
	token    *AccessToken
	request  *Request
	response *Response
	client   *http.Client
}

func (c *Client) SetDomain(domain Domain) *Client {
	c.domain = domain
	return c
}

func (c *Client) URL() string {
	return c.domain.URL()
}

func (c *Client) HttpClient() *http.Client {
	return c.client
}

func (c *Client) SetHttpClient(client *http.Client) *Client {
	c.client = client
	return c
}

func (c *Client) DataType() DataType {
	return c.dataType
}

func (c *Client) SetDataType(dataType DataType) *Client {
	c.dataType = dataType
	return c
}

func (c *Client) HttpPostJson(url string, data Map, ops Map) *Response {
	ops = MapNilMake(ops)
	c.dataType = DATA_TYPE_JSON
	ops.Set(REQUEST_TYPE_JSON.String(), data)
	return c.Request(url, nil, "post", ops)
}

func (c *Client) HttpGet(url string, ops Map) *Response {
	return c.Request(url, nil, "get", ops)
}

func (c *Client) HttpPost(url string, ops Map) *Response {
	return c.Request(url, nil, "post", ops)
}

func (c *Client) Request(url string, params Map, method string, ops Map) *Response {
	Debug("Request|httpClient", c.client)
	c.client = buildTransport(c.Config)
	c.response = request(c, url, params, method, ops)
	return c.response
}

func (c *Client) RequestRaw(url string, params Map, method string, ops Map) *Response {
	return c.Request(url, params, method, ops)
}

func (c *Client) SafeRequest(url string, params Map, method string, ops Map) *Response {
	c.client = buildSafeTransport(c.Config)
	Debug("SafeRequest|httpClient", c.client)
	c.response = request(c, url, params, method, ops)
	return c.response
}

func (c *Client) Link(uri string) string {
	if c.GetBool("sandbox") {
		return c.URL() + SANDBOX_URL_SUFFIX + uri
	}
	return c.URL() + uri
}

func (c *Client) GetResponse() *Response {
	return c.response
}

func (c *Client) GetRequest() *Request {
	return c.request
}

func DefaultClient() *Client {
	return nil
}

func NewClient(config Config) *Client {
	Debug("NewClient|config", config)
	domain := NewDomain("default")
	if config == nil {
		config = defaultConfig
	}
	return &Client{
		request:  DefaultRequest,
		Config:   config,
		dataType: DATA_TYPE_XML,
		domain:   domain,
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

func request(c *Client, url string, params Map, method string, op Map) *Response {
	Debug("client|request", c, url, params, method, op)
	op = MapNilMake(op)
	if params != nil {
		params.Set("mch_id", c.Get("mch_id"))
		params.Set("nonce_str", GenerateUUID())
		params.Set("sub_mch_id", c.Get("sub_mch_id"))
		params.Set("sub_appid", c.Get("sub_appid"))
		params.Set("sign_type", SIGN_TYPE_MD5.String())
		params.Set("sign", GenerateSignature(params, c.Get("key"), SIGN_TYPE_MD5))
	}

	data := toRequestData(c, params, op)

	if r := c.request.PerformRequest(url, method, data); r.Error() == nil {
		return Do(context.Background(), c, r)
	} else {
		return ErrorResponse(r.Error())
	}
}

func Do(ctx context.Context, client *Client, request *Request) *Response {
	response := &Response{}
	response.response, response.error = client.HttpClient().Do(request.HttpRequest().WithContext(ctx))
	if response.error != nil {
		select {
		case <-ctx.Done():
			response.error = ctx.Err()
		default:
		}
		return response
	}

	response.responseData, response.error = ioutil.ReadAll(response.response.Body)
	response.responseType = RESPONSE_TYPE_XML
	if client.DataType() == DATA_TYPE_JSON {
		response.responseType = RESPONSE_TYPE_JSON
	}
	Debug("ClientDo|response", *response)
	Debug("ClientDo|response|data", string(response.responseData))
	return response
}

func toRequestData(client *Client, params, ops Map) *RequestData {
	data := client.request.RequestDataCopy()
	data.Query = processQuery(ops.Get(REQUEST_TYPE_QUERY.String()))
	if client.DataType() == DATA_TYPE_JSON {
		data.SetHeaderJson()
		if params == nil {
			data.Body = processJson(ops.Get(REQUEST_TYPE_JSON.String()))
		} else {
			data.Body = bytes.NewReader(params.ToJson())
		}
	}
	if client.DataType() == DATA_TYPE_XML {
		data.SetHeaderXml()
		if params == nil {
			data.Body = processXml(ops.Get(REQUEST_TYPE_XML.String()))
		} else {
			data.Body = strings.NewReader(params.ToXml())
		}

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
func processXml(i interface{}) io.Reader {
	switch v := i.(type) {
	case string:
		return strings.NewReader(v)
	case []byte:
		return bytes.NewReader(v)
	case Map:
		return strings.NewReader(v.ToXml())
	default:
		return nil
	}

}

func processJson(i interface{}) io.Reader {
	switch v := i.(type) {
	case string:
		return strings.NewReader(v)
	case []byte:
		return bytes.NewReader(v)
	case Map:
		return bytes.NewReader(v.ToJson())
	default:
		return nil
	}

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

func (u *URL) ShortUrl(url string) Map {
	m := Map{
		"action":   "long2short",
		"long_url": url,
	}
	token := u.token.GetToken()
	ops := Map{
		REQUEST_TYPE_QUERY.String(): Map{
			"access_token": token.GetKey(),
		},
	}
	resp := u.client.HttpPostJson(u.client.Link(SHORTURL_URL_SUFFIX), m, ops)
	Debug("URL|ShortUrl", *resp)
	return resp.ToMap()
}

func NewURL(config Config, client *Client) *URL {
	return &URL{
		token:  NewAccessToken(config, client),
		client: client,
	}
}
