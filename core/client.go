package core

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type DataType string
type ContextKey struct{}

const (
	DATA_TYPE_XML       DataType = "xml"
	DATA_TYPE_JSON      DataType = "json"
	DATA_TYPE_MULTIPART DataType = "multipart"
)

type URL struct {
	token  *AccessToken
	client *Client
}

type Client struct {
	Config
	dataType DataType
	domain   *Domain
	app      *Application
	token    *AccessToken
	request  *Request
	response *Response
	client   *http.Client
}

var HTTPClient ContextKey

func (c *Client) SetDomain(domain *Domain) *Client {
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

func (c *Client) HttpPostJson(url string, query Map, json interface{}) *Response {
	c.dataType = DATA_TYPE_JSON
	p := Map{
		REQUEST_TYPE_JSON.String(): json,
	}
	if query != nil {
		p.Set(REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "post")
}

func (c *Client) HttpPostXml(url string, query Map, xml interface{}) *Response {
	c.dataType = DATA_TYPE_XML
	p := Map{
		REQUEST_TYPE_XML.String(): xml,
	}
	if query != nil {
		p.Set(REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "post")
}

func (c *Client) HttpUpload(url string, query, multi Map) *Response {
	c.dataType = DATA_TYPE_MULTIPART
	p := Map{
		REQUEST_TYPE_MULTIPART.String(): multi,
	}
	if query != nil {
		p.Set(REQUEST_TYPE_QUERY.String(), query)
	}

	return c.Request(url, p, "post")
}

func (c *Client) HttpGet(url string, query Map) *Response {
	p := Map{}
	if query != nil {
		p.Set(REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "get")
}

func (c *Client) HttpPost(url string, query Map, ops Map) *Response {
	p := Map{}
	if query != nil {
		p.Set(REQUEST_TYPE_QUERY.String(), query)
	}
	if ops != nil {
		p.ReplaceJoin(ops)
	}
	return c.Request(url, p, "post")
}

func (c *Client) Request(url string, ops Map, method string) *Response {
	Debug("Request|httpClient", c.client)
	c.client = buildTransport(c.Config)
	c.response = request(c, url, ops, method)
	return c.response
}

func (c *Client) RequestRaw(url string, ops Map, method string) *Response {
	return c.Request(url, ops, method)
}

func (c *Client) SafeRequest(url string, ops Map, method string) *Response {
	c.client = buildSafeTransport(c.Config)
	Debug("SafeRequest|httpClient", c.client)
	c.response = request(c, url, ops, method)
	return c.response
}

func (c *Client) Link(uri string) string {
	if c.GetBool("sandbox") {
		return c.URL() + SANDBOX_URL_SUFFIX + uri
	}
	return c.domain.Link(uri)
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
	Debug("buildSafeTransport", config.Get("cert_path"), config.Get("key_path"), config.Get("rootca_path"))
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

func request(c *Client, url string, ops Map, method string) *Response {
	Debug("client|request", c, url, ops, method)
	data := toRequestData(c, ops)

	if r := c.request.PerformRequest(url, method, data); r.Error() == nil {
		return Do(context.Background(), c, r)
	} else {
		return ErrorResponse(r.Error())
	}
}

func Do(ctx context.Context, client *Client, request *Request) *Response {
	var response Response

	r, err := client.client.Do(request.request.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
		Error("Client|Do", err)
		response.error = err
		return &response
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		Debug("Do|StatusCode", r.StatusCode)
	}

	response.responseData, response.error = ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	response.responseType = RESPONSE_TYPE_JSON
	if client.DataType() == RESPONSE_TYPE_XML {
		response.responseType = RESPONSE_TYPE_XML
	}
	Debug("ClientDo|response", response.responseType, response.error, response.responseMap)
	Debug("ClientDo|response|data", len(response.responseData))
	return &response
}

func toRequestData(client *Client, ops Map) *RequestData {
	data := client.request.RequestDataCopy()
	data.Query = processQuery(ops.Get(REQUEST_TYPE_QUERY.String()))
	data.Body = nil
	if client.DataType() == DATA_TYPE_JSON {
		data.SetHeaderJson()
		data.Body = processJson(ops.Get(REQUEST_TYPE_JSON.String()))
	}
	if client.DataType() == DATA_TYPE_XML {
		data.SetHeaderXml()
		data.Body = processXml(ops.Get(REQUEST_TYPE_XML.String()))
	}

	if client.DataType() == DATA_TYPE_MULTIPART {
		buf := bytes.Buffer{}
		writer := multipart.NewWriter(&buf)
		writer = processMultipart(writer, ops.Get(REQUEST_TYPE_MULTIPART.String()))
		data.Body = &buf
		data.Header.Set("Content-Type", writer.FormDataContentType())
		defer writer.Close()
	}

	return data
}
func processMultipart(w *multipart.Writer, i interface{}) *multipart.Writer {
	Debug("processMultipart|i", i)
	switch v := i.(type) {
	case Map:
		path := v.GetString("media")
		// Debug("processMultipart|name", name)

		// Debug("processMultipart|path", path)
		fh, e := os.Open(path)
		if e != nil {
			Debug("processMultipart|e", e)
			return w
		}
		defer fh.Close()

		fw, e := w.CreateFormFile("media", path)
		if e != nil {
			Debug("processMultipart|e", e)
			return w
		}

		if _, e = io.Copy(fw, fh); e != nil {
			Debug("processMultipart|e", e)
			return w
		}
		des := v.GetMap("description")
		if des != nil {
			w.WriteField("description", string(des.ToJson()))
		}

	}
	return w
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
		Debug("processXml|string", v)
		return strings.NewReader(v)
	case []byte:
		Debug("processXml|[]byte", v)
		return bytes.NewReader(v)
	case Map:
		Debug("processXml|Map", v.ToXml())
		return strings.NewReader(v.ToXml())
	default:
		Debug("processXml|default", v)
		if v0, e := xml.Marshal(v); e == nil {
			Debug("processXml|v0", v0, e)
			return bytes.NewReader(v0)
		}
		return nil
	}

}

func processJson(i interface{}) io.Reader {
	switch v := i.(type) {
	case string:
		Debug("processJson|string", v)
		return strings.NewReader(v)
	case []byte:
		Debug("processJson|[]byte", string(v))
		return bytes.NewReader(v)
	case Map:
		Debug("processJson|Map", v.String())
		return bytes.NewReader(v.ToJson())
	default:
		Debug("processJson|default", v)
		if v0, e := json.Marshal(v); e == nil {
			Debug("processJson|v0", string(v0), e)
			return bytes.NewReader(v0)
		}
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
		REQUEST_TYPE_QUERY.String(): token.KeyMap(),
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
