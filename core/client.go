package core

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*data types*/
const (
	DataTypeXML       = "xml"
	DataTypeJSON      = "json"
	DataTypeQuery     = "query"
	DataTypeForm      = "form_params"
	DataTypeFile      = "file"
	DataTypeMultipart = "multipart"
	DataTypeSecurity  = "security"
)

const defaultCa = `
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
-----END CERTIFICATE-----`

var client = &Client{
	Context: context.Background(),
}

/*Client Client */
type Client struct {
	context.Context
}

// ClientSetter ...
type ClientSetter interface {
	SetClient(*Client)
}

// ClientGetter ...
type ClientGetter interface {
	GetClient() *Client
}

// ClientGet ...
func ClientGet(v []interface{}) *Client {
	size := len(v)
	for i := 0; i < size; i++ {
		switch sv := v[i].(type) {
		case *Client:
			return sv
		case Client:
			return &sv
		}
	}

	return DefaultClient()
}

// ClientSet ...
func ClientSet(setter ClientSetter, v []interface{}) bool {
	size := len(v)
	for i := 0; i < size; i++ {
		switch sv := (v[i]).(type) {
		case *Client:
			setter.SetClient(sv)
			return true
		case Client:
			setter.SetClient(&sv)
			return true
		}
	}

	return false
}

// PostForm post form request
func PostForm(url string, query util.Map, form interface{}) Responder {
	return client.PostForm(url, query, form)
}

// PostForm post form request
func (c *Client) PostForm(url string, query util.Map, form interface{}) Responder {
	url = url + "?" + query.URLEncode()
	request := processForm(POST, url, form)
	client := buildTransport(NilConfig())
	return do(c.Context, client, request)
}

// PostJSON json post请求
func PostJSON(url string, query util.Map, json interface{}) Responder {
	return client.PostJSON(url, query, json)
}

// PostJSON json post请求
func (c *Client) PostJSON(url string, query util.Map, json interface{}) Responder {
	url = url + "?" + query.URLEncode()
	request := processJSON(POST, url, json)
	client := buildTransport(NilConfig())
	return do(c.Context, client, request)
}

// PostXML  xml post请求
func PostXML(url string, query util.Map, xml interface{}) Responder {
	return client.PostXML(url, query, xml)
}

// PostXML  xml post请求
func (c *Client) PostXML(url string, query util.Map, xml interface{}) Responder {
	url = url + "?" + query.URLEncode()
	request := processXML(POST, url, xml)
	client := buildTransport(NilConfig())
	return do(c.Context, client, request)
}

// Upload upload请求
func Upload(url string, query, multi util.Map) Responder {
	return client.Upload(url, query, multi)
}

// Upload upload请求
func (c *Client) Upload(url string, query, multi util.Map) Responder {
	url = url + "?" + query.URLEncode()
	request := processMultipart(POST, url, multi)
	client := buildTransport(NilConfig())
	return do(c.Context, client, request)
}

// Post post请求
func Post(url string, maps util.Map) Responder {
	return client.Post(url, maps)
}

// Post post请求
func (c *Client) Post(url string, maps util.Map) Responder {
	client := buildClient(maps)
	url = buildRequestURL(url, maps)
	req := buildRequest(POST, url, maps)
	return do(c.Context, client, req)
}

// Get get请求
func Get(url string, query util.Map) Responder {
	return client.Get(url, query)
}

// Get get请求
func (c *Client) Get(url string, query util.Map) Responder {
	url = url + "?" + query.URLEncode()
	request := processNothing(GET, url, nil)
	client := buildTransport(NilConfig())
	return do(c.Context, client, request)
}

//GetRaw get request result raw data
func (c *Client) GetRaw(url string, query util.Map) []byte {
	url = url + "?" + query.URLEncode()
	request := processNothing(GET, url, nil)
	client := buildTransport(NilConfig())
	return do(c.Context, client, request).Bytes()
}

// Request ...
func Request(url string, method string, option util.Map) Responder {
	log.Debug("Request|httpClient", url, method, option)
	return request(context.Background(), url, method, option)
}

// RequestWithContext ...
func RequestWithContext(ctx context.Context, url string, method string, option util.Map) Responder {
	log.Debug("RequestWithContext|httpClient", url, method, option)
	return request(ctx, url, method, option)
}

// RequestRaw ...
func RequestRaw(url string, method string, option util.Map) []byte {
	log.Debug("RequestRaw|httpClient", url, method, option)
	return request(context.Background(), url, method, option).Bytes()
}

/*Request 普通请求 */
func (c *Client) Request(url string, method string, option util.Map) Responder {
	log.Debug("ClientRequest|httpClient", url, method, option)
	return request(c.Context, url, method, option)
}

/*RequestRaw raw请求 */
func (c *Client) RequestRaw(url string, method string, option util.Map) []byte {
	log.Debug("ClientRequest|httpClient", url, method, option)
	return request(c.Context, url, method, option).Bytes()
}

func newClient(ctx context.Context) *Client {
	return &Client{
		Context: ctx,
	}
}

/*NewClient 创建一个client */
func NewClient(ctx context.Context) *Client {
	return newClient(ctx)
}

//DefaultClient result a client with default value
func DefaultClient() *Client {
	return client
}

func buildTransport(config *Config) *http.Client {
	timeOut := config.GetIntD("http.time_out", 30)
	keepAlive := config.GetIntD("http.keep_alive", 30)
	return &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(timeOut) * time.Second,
				KeepAlive: time.Duration(keepAlive) * time.Second,
				//DualStack: true,
			}).DialContext,
			//Dial:        nil,
			//DialTLS:     nil,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			//TLSHandshakeTimeout:    0,
			//DisableKeepAlives:      false,
			//DisableCompression:     false,
			//MaxIdleConns:           0,
			//MaxIdleConnsPerHost:    0,
			//MaxConnsPerHost:        0,
			//IdleConnTimeout:        0,
			//ResponseHeaderTimeout:  0,
			//ExpectContinueTimeout:  0,
			//TLSNextProto:           nil,
			//ProxyConnectHeader:     nil,
			//MaxResponseHeaderBytes: 0,
		},
		//CheckRedirect: nil,
		//Jar:           nil,
		//Timeout:       0,
	}

}

func buildSafeTransport(config *Config) *http.Client {
	if config == nil {
		panic("safe request must set config before use")
	}

	//if idx := config.Check("cert_path", "key_path"); idx != -1 {
	//	panic(fmt.Sprintf("the %d key was not found", idx))
	//}

	cert, err := tls.LoadX509KeyPair(config.GetString("cert_path"), config.GetString("key_path"))
	if err != nil {
		panic(err)
	}

	caFile, err := ioutil.ReadFile(config.GetString("rootca_path"))
	if err != nil {
		caFile = []byte(defaultCa)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caFile)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: true, //client端略过对证书的校验
	}
	tlsConfig.BuildNameToCertificate()
	timeOut := config.GetIntD("http.time_out", 30)
	keepAlive := config.GetIntD("http.keep_alive", 30)
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(timeOut) * time.Second,
				KeepAlive: time.Duration(keepAlive) * time.Second,
				//DualStack: true,
			}).DialContext,
			TLSClientConfig: tlsConfig,
			Proxy:           nil,
			//TLSHandshakeTimeout:   10 * time.Second,
			//ResponseHeaderTimeout: 10 * time.Second,
			//ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

func buildClient(maps util.Map) *http.Client {
	//检查是否包含security

	if maps.Has(DataTypeSecurity) {
		//判断能否创建safe client
		v, b := maps.Get(DataTypeSecurity).(*Config)
		//log.Debug("build client \n", v)
		if b && v.Check("cert_path", "key_path") == -1 {
			return buildSafeTransport(v)
		}
		return buildTransport(v)

	}
	//默认输出未配置client
	log.Debug("default client")

	return buildTransport(NilConfig())
}

func do(ctx context.Context, c *http.Client, r *http.Request) Responder {
	response, err := c.Do(r.WithContext(ctx))
	if err != nil {
		log.Error("Client|Do", err)
		return Err(nil, err)
	}
	{
		select {
		case <-ctx.Done():
			return Err(nil, ctx.Err())
		default:
			//return Err(nil, err)
		}
	}

	return CastToResponse(response)
}
