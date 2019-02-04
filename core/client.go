package core

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"net/http"
	"net/url"
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

// ClientOption ...
type ClientOption struct {
	RequestType string
	UseSafe     bool
	Cert        []byte
	Key         []byte
	RootCA      []byte
	Timeout     int64
	KeepAlive   int64
}

// Body ...
type RequestBody struct {
	BodyType string
	BodyInst interface{}
}

// Client ...
type Client struct {
	Method string
	URL    string
	Query  url.Values
	Body   interface{}
	Option *ClientOption
}

// Request ...
func (client *Client) Request() Responder {
	return (client).Do(context.Background())
}

// NewClient ...
func NewClient(opts ...*ClientOption) *Client {
	return &Client{
		Method: "",
		URL:    "",
		Query:  nil,
		Body:   nil,
		Option: nil,
	}
}

// PostForm post form request
func PostForm(url string, query util.Map, form interface{}) Responder {
	url = url + "?" + query.URLEncode()
	request := &request{
		client:   buildTransport(NilConfig()),
		function: processForm,
		method:   POST,
		url:      url,
		body:     form,
	}
	return request.Do(context.Background())
}

// PostJSON json post请求
func PostJSON(url string, query util.Map, json interface{}) Responder {
	url = url + "?" + query.URLEncode()
	request := &request{
		client:   buildTransport(NilConfig()),
		function: processJSON,
		method:   POST,
		url:      url,
		body:     json,
	}
	return request.Do(context.Background())
}

// PostXML  xml post请求
func PostXML(url string, query util.Map, xml interface{}) Responder {
	url = url + "?" + query.URLEncode()
	request := &request{
		client:   buildTransport(NilConfig()),
		function: processXML,
		method:   POST,
		url:      url,
		body:     xml,
	}
	return request.Do(context.Background())
}

// Upload upload请求
func Upload(url string, query, multi util.Map) Responder {
	url = url + "?" + query.URLEncode()
	request := &request{
		client:   buildTransport(NilConfig()),
		function: processMultipart,
		method:   POST,
		url:      url,
		body:     multi,
	}
	return request.Do(context.Background())
}

// Post post请求
func Post(url string, maps util.Map) Responder {
	return Request(POST, url, maps)
}

// Get get请求
func Get(url string, query util.Map) Responder {
	url = url + "?" + query.URLEncode()
	request := &request{
		client:   buildTransport(),
		function: processNothing,
		method:   GET,
		url:      url,
		body:     nil,
	}
	return request.Do(context.Background())
}

// GetRaw get请求 返回[]byte
func GetRaw(url string, query util.Map) []byte {
	return Get(url, query).Bytes()
}

// Request ...
func Request(method string, url string, option util.Map) Responder {
	log.Debug("Request|httpClient", url, method, option)
	return RequestWithContext(context.Background(), method, url, option)
}

// RequestWithContext ...
func RequestWithContext(ctx context.Context, method string, url string, option util.Map) Responder {
	log.Debug("RequestWithContext|httpClient", url, method, option)
	return buildRequester(method, url, option).Do(ctx)
}

// RequestRaw ...
func RequestRaw(method, url string, option util.Map) []byte {
	log.Debug("RequestRaw|httpClient", url, method, option)
	return RequestWithContext(context.Background(), method, url, option).Bytes()
}

// TimeOut ...
func TimeOut(src int64) time.Duration {
	return time.Duration(util.MustInt64(src, 30)) * time.Second
}

// KeepAlive ...
func KeepAlive(src int64) time.Duration {
	return time.Duration(util.MustInt64(src, 30)) * time.Second
}

func buildTransport(client *Client) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout:   TimeOut(client.Option.Timeout),
				KeepAlive: KeepAlive(client.Option.KeepAlive),
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

func buildSafeTransport(client *Client) *http.Client {
	cert, err := tls.X509KeyPair(client.Option.Key, client.Option.Cert)
	if err != nil {
		panic(err)
	}

	caFile := client.Option.RootCA
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
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   TimeOut(client.Option.Timeout),
				KeepAlive: KeepAlive(client.Option.KeepAlive),
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

func buildHTTPClient(client *Client) *http.Client {
	//检查是否包含security
	if client.Option.UseSafe {
		//判断能否创建safe client
		return buildSafeTransport(client)
	}
	return buildTransport(client)

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
