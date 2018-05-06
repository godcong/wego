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

	"github.com/godcong/wego/core/config"
	"github.com/godcong/wego/core/log"
	"github.com/godcong/wego/core/net"
	"github.com/godcong/wego/core/util"
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
	config.Config
	dataType DataType
	domain   *Domain
	app      *Application
	token    *AccessToken
	request  *net.Request
	response *net.Response
	client   *http.Client
}

var defaultCa = []byte(`
-----BEGIN CERTIFICATE-----
MIIDIDCCAomgAwIBAgIENd70zzANBgkqhkiG9w0BAQUFADBOMQswCQYDVQQGEwJV
UzEQMA4GA1UEChMHRXF1aWZheDEtMCsGA1UECxMkRXF1aWZheCBTZWN1cmUgQ2Vy
dGlmaWNhdGUgQXV0aG9yaXR5MB4XDTk4MDgyMjE2NDE1MVoXDTE4MDgyMjE2NDE1
MVowTjELMAkGA1UEBhMCVVMxEDAOBgNVBAoTB0VxdWlmYXgxLTArBgNVBAsTJEVx
dWlmYXggU2VjdXJlIENlcnRpZmljYXRlIEF1dGhvcml0eTCBnzANBgkqhkiG9w0B
AQEFAAOBjQAwgYkCgYEAwV2xWGcIYu6gmi0fCG2RFGiYCh7+2gRvE4RiIcPRfM6f
BeC4AfBONOziipUEZKzxa1NfBbPLZ4C/QgKO/t0BCezhABRP/PvwDN1Dulsr4R+A
cJkVV5MW8Q+XarfCaCMczE1ZMKxRHjuvK9buY0V7xdlfUNLjUA86iOe/FP3gx7kC
AwEAAaOCAQkwggEFMHAGA1UdHwRpMGcwZaBjoGGkXzBdMQswCQYDVQQGEwJVUzEQ
MA4GA1UEChMHRXF1aWZheDEtMCsGA1UECxMkRXF1aWZheCBTZWN1cmUgQ2VydGlm
aWNhdGUgQXV0aG9yaXR5MQ0wCwYDVQQDEwRDUkwxMBoGA1UdEAQTMBGBDzIwMTgw
ODIyMTY0MTUxWjALBgNVHQ8EBAMCAQYwHwYDVR0jBBgwFoAUSOZo+SvSspXXR9gj
IBBPM5iQn9QwHQYDVR0OBBYEFEjmaPkr0rKV10fYIyAQTzOYkJ/UMAwGA1UdEwQF
MAMBAf8wGgYJKoZIhvZ9B0EABA0wCxsFVjMuMGMDAgbAMA0GCSqGSIb3DQEBBQUA
A4GBAFjOKer89961zgK5F7WF0bnj4JXMJTENAKaSbn+2kmOeUJXRmm/kEd5jhW6Y
7qj/WsjTVbJmcVfewCHrPSqnI0kBBIZCe/zuf6IWUrVnZ9NA2zsmWLIodz2uFHdh
1voqZiegDfqnc1zqcPGUIWVEX/r87yloqaKHee9570+sB3c4
-----END CERTIFICATE-----`)

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

func (c *Client) HttpPostJson(url string, query util.Map, json interface{}) *net.Response {
	c.dataType = DATA_TYPE_JSON
	p := util.Map{
		net.REQUEST_TYPE_JSON.String(): json,
	}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "post")
}

func (c *Client) HttpPostXml(url string, query util.Map, xml interface{}) *net.Response {
	c.dataType = DATA_TYPE_XML
	p := util.Map{
		net.REQUEST_TYPE_XML.String(): xml,
	}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "post")
}

func (c *Client) HttpUpload(url string, query, multi util.Map) *net.Response {
	c.dataType = DATA_TYPE_MULTIPART
	p := util.Map{
		net.REQUEST_TYPE_MULTIPART.String(): multi,
	}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}

	return c.Request(url, p, "post")
}

func (c *Client) HttpGet(url string, query util.Map) *net.Response {
	p := util.Map{}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "get")
}

func (c *Client) HttpPost(url string, query util.Map, ops util.Map) *net.Response {
	p := util.Map{}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}
	if ops != nil {
		p.ReplaceJoin(ops)
	}
	return c.Request(url, p, "post")
}

func (c *Client) Request(url string, ops util.Map, method string) *net.Response {
	log.Debug("Request|httpClient", c.client)
	c.client = buildTransport(c.Config)
	c.response = request(c, url, ops, method)
	return c.response
}

func (c *Client) RequestRaw(url string, ops util.Map, method string) *net.Response {
	return c.Request(url, ops, method)
}

func (c *Client) SafeRequest(url string, ops util.Map, method string) *net.Response {
	c.client = buildSafeTransport(c.Config)
	log.Debug("SafeRequest|httpClient", c.client)
	c.response = request(c, url, ops, method)
	return c.response
}

func (c *Client) Link(uri string) string {
	if c.GetBool("sandbox") {
		return c.URL() + SANDBOX_URL_SUFFIX + uri
	}
	return c.domain.Link(uri)
}

func (c *Client) GetResponse() *net.Response {
	return c.response
}

func (c *Client) GetRequest() *net.Request {
	return c.request
}

func DefaultClient() *Client {
	return nil
}

func NewClient(config config.Config) *Client {
	log.Debug("NewClient|config", config)
	domain := NewDomain("default")
	if config == nil {
		config = defaultConfig
	}
	return &Client{
		request:  net.DefaultRequest,
		Config:   config,
		dataType: DATA_TYPE_XML,
		domain:   domain,
	}
}

func buildTransport(config config.Config) *http.Client {
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

func buildSafeTransport(config config.Config) *http.Client {
	log.Debug("buildSafeTransport", config.Get("cert_path"), config.Get("key_path"), config.Get("rootca_path"))
	cert, err := tls.LoadX509KeyPair(config.Get("cert_path"), config.Get("key_path"))
	if err != nil {
		panic(err)
	}

	caFile, err := ioutil.ReadFile(config.Get("rootca_path"))
	if err != nil {
		caFile = defaultCa
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

func request(c *Client, url string, ops util.Map, method string) *net.Response {
	log.Debug("client|request", c, url, ops, method)
	data := toRequestData(c, ops)

	if r := c.request.PerformRequest(url, method, data); r.Error() == nil {
		return Do(context.Background(), c, r)
	} else {
		return net.ErrorResponse(r.Error())
	}
}

func Do(ctx context.Context, client *Client, request *net.Request) *net.Response {
	var response *net.Response

	r, err := client.client.Do(request.HttpRequest().WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
		log.Error("Client|Do", err)
		return net.ErrorResponse(err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		log.Debug("Do|StatusCode", r.StatusCode)
	}
	response = net.ParseResponse(r)

	return response
}

func toRequestData(client *Client, ops util.Map) *net.RequestData {
	data := client.request.RequestDataCopy()
	data.Query = processQuery(ops.Get(net.REQUEST_TYPE_QUERY.String()))
	data.Body = nil
	if client.DataType() == DATA_TYPE_JSON {
		data.SetHeaderJson()
		data.Body = processJson(ops.Get(net.REQUEST_TYPE_JSON.String()))
	}
	if client.DataType() == DATA_TYPE_XML {
		data.SetHeaderXml()
		data.Body = processXml(ops.Get(net.REQUEST_TYPE_XML.String()))
	}

	if client.DataType() == DATA_TYPE_MULTIPART {
		buf := bytes.Buffer{}
		writer := multipart.NewWriter(&buf)
		writer = processMultipart(writer, ops.Get(net.REQUEST_TYPE_MULTIPART.String()))
		data.Body = &buf
		data.Header.Set("Content-Type", writer.FormDataContentType())
		defer writer.Close()
	}

	return data
}
func processMultipart(w *multipart.Writer, i interface{}) *multipart.Writer {
	log.Debug("processMultipart|i", i)
	switch v := i.(type) {
	case util.Map:
		path := v.GetString("media")
		// log.Debug("processMultipart|name", name)

		// log.Debug("processMultipart|path", path)
		fh, e := os.Open(path)
		if e != nil {
			log.Debug("processMultipart|e", e)
			return w
		}
		defer fh.Close()

		fw, e := w.CreateFormFile("media", path)
		if e != nil {
			log.Debug("processMultipart|e", e)
			return w
		}

		if _, e = io.Copy(fw, fh); e != nil {
			log.Debug("processMultipart|e", e)
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
	case util.Map:
		return v.ToXml()
	}
	return ""
}
func processXml(i interface{}) io.Reader {
	switch v := i.(type) {
	case string:
		log.Debug("processXml|string", v)
		return strings.NewReader(v)
	case []byte:
		log.Debug("processXml|[]byte", v)
		return bytes.NewReader(v)
	case util.Map:
		log.Debug("processXml|util.Map", v.ToXml())
		return strings.NewReader(v.ToXml())
	default:
		log.Debug("processXml|default", v)
		if v0, e := xml.Marshal(v); e == nil {
			log.Debug("processXml|v0", v0, e)
			return bytes.NewReader(v0)
		}
		return nil
	}

}

func processJson(i interface{}) io.Reader {
	switch v := i.(type) {
	case string:
		log.Debug("processJson|string", v)
		return strings.NewReader(v)
	case []byte:
		log.Debug("processJson|[]byte", string(v))
		return bytes.NewReader(v)
	case util.Map:
		log.Debug("processJson|util.Map", v.String())
		return bytes.NewReader(v.ToJson())
	default:
		log.Debug("processJson|default", v)
		if v0, e := json.Marshal(v); e == nil {
			log.Debug("processJson|v0", string(v0), e)
			return bytes.NewReader(v0)
		}
		return nil
	}
}

func processQuery(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case util.Map:
		return v.UrlEncode()
	}
	return ""
}

func (u *URL) ShortUrl(url string) util.Map {
	m := util.Map{
		"action":   "long2short",
		"long_url": url,
	}
	token := u.token.GetToken()
	ops := util.Map{
		net.REQUEST_TYPE_QUERY.String(): token.KeyMap(),
	}
	resp := u.client.HttpPostJson(u.client.Link(SHORTURL_URL_SUFFIX), m, ops)
	log.Debug("URL|ShortUrl", *resp)
	return resp.ToMap()
}

func NewURL(config config.Config, client *Client) *URL {
	return &URL{
		token:  NewAccessToken(config, client),
		client: client,
	}
}
