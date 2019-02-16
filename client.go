package wego

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"golang.org/x/exp/xerrors"
	"io"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"time"
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

// Client ...
type Client struct {
	Method      string
	URL         string
	Query       util.Map
	Body        *RequestBody
	BodyType    BodyType
	safe        bool
	safeCert    *SafeCertConfig
	accessToken *AccessToken
	timeout     int64
	keepAlive   int64
}

// ClientOption ...
type ClientOption struct {
	UseSafe     bool
	SafeCert    *SafeCertConfig
	AccessToken *AccessToken
	BodyType    *BodyType
	Timeout     int64
	KeepAlive   int64
}

// NewClient ...
func NewClient(opts ...*ClientOption) *Client {
	client := &Client{
		BodyType: BodyTypeXML,
	}
	client.parse(opts)
	return client
}

func (obj *Client) parse(opts []*ClientOption) {
	if opts == nil {
		return
	}

	obj.safeCert = opts[0].SafeCert
	if opts[0].BodyType != nil {
		obj.BodyType = *opts[0].BodyType
	}
	obj.accessToken = opts[0].AccessToken
	obj.accessToken = opts[0].AccessToken
	obj.timeout = opts[0].Timeout
	obj.keepAlive = opts[0].KeepAlive
}

// IsSafe ...
func (obj *Client) IsSafe() bool {
	return obj.safe
}

// SetSafe ...
func (obj *Client) SetSafe(b bool) {
	obj.safe = b
}

// Post ...
func (obj *Client) Post(ctx context.Context, url string, query util.Map, body interface{}) Responder {
	log.Debug("post ", url, body)
	obj.Method = POST
	obj.Query = query
	obj.URL = url
	obj.Body = buildBody(body, obj.BodyType)
	return obj.do(ctx)
}

// Get ...
func (obj *Client) Get(ctx context.Context, url string, query util.Map) Responder {
	log.Debug("get ", url)
	obj.Method = GET
	obj.Query = query
	obj.URL = url
	obj.Body = buildBody(nil, obj.BodyType)
	return obj.do(ctx)
}

// CA ...
func (obj *Client) CA() []byte {
	if obj.safeCert != nil && obj.safeCert.RootCA != nil {
		return obj.safeCert.RootCA
	}
	return []byte(defaultCa)
}

func (obj *Client) certKey() ([]byte, []byte) {
	if obj.safeCert != nil &&
		obj.safeCert.Key != nil &&
		obj.safeCert.Cert != nil {
		return obj.safeCert.Cert, obj.safeCert.Key
	}
	return nil, nil
}

// HTTPClient ...
func (obj *Client) HTTPClient() (*http.Client, error) {
	return buildHTTPClient(obj)
}

// makeClient ...
func makeClient(method string, url string, query util.Map, body interface{}, opts ...*ClientOption) *Client {
	client := NewClient(opts...)
	client.Method = method
	client.URL = url
	client.Query = query
	//buildBody must after parse
	client.BuildBody(body)
	return client
}

// BuildBody ...
func (obj *Client) BuildBody(body interface{}) {
	obj.Body = buildBody(body, obj.BodyType)
}

// Request ...
func (obj *Client) Request() (*http.Request, error) {
	if obj.Body == nil {
		return http.NewRequest(obj.Method, obj.URLQuery(), nil)
	}
	return obj.Body.RequestBuilder(obj.Method, obj.URLQuery(), obj.Body.BodyInstance)
}

// do ...
func (obj *Client) do(ctx context.Context) Responder {
	client, e := obj.HTTPClient()
	if e != nil {
		return ErrResponse(xerrors.Errorf("client:%w", e))
	}
	request, e := obj.Request()
	if e != nil {
		return ErrResponse(xerrors.Errorf("request:%w", e))
	}
	response, e := client.Do(request.WithContext(ctx))
	if e != nil {
		return ErrResponse(xerrors.Errorf("response:%w", e))
	}
	return BuildResponder(response)
}

// URLQuery ...
func (obj *Client) URLQuery() string {
	if obj.Query == nil {
		return obj.URL
	}
	return obj.URL + "?" + obj.Query.URLEncode()

}

// PostForm post form request
func PostForm(url string, query util.Map, form interface{}) Responder {
	log.Debug("post form:", url, query, form)
	bt := BodyTypeForm
	client := makeClient(POST, url, query, form, &ClientOption{
		BodyType: &bt,
	})
	return client.do(context.Background())
}

// PostJSON json post请求
func PostJSON(url string, query util.Map, json interface{}) Responder {
	log.Debug("post json:", url, query, json)
	bt := BodyTypeJSON
	client := makeClient(POST, url, query, json, &ClientOption{
		BodyType: &bt,
	})
	return client.do(context.Background())
}

// PostXML  xml post请求
func PostXML(url string, query util.Map, xml interface{}) Responder {
	log.Debug("post xml:", url, query, xml)
	bt := BodyTypeXML
	client := makeClient(POST, url, query, xml, &ClientOption{
		BodyType: &bt,
	})
	return client.do(context.Background())
}

// Get get请求
func Get(url string, query util.Map) Responder {
	log.Println("get request:", url, query)
	bt := BodyTypeXML
	client := makeClient(GET, url, query, nil, &ClientOption{
		BodyType: &bt,
	})
	return client.do(context.Background())
}

// Upload upload请求
func Upload(url string, query, multi util.Map) Responder {
	bt := BodyTypeMultipart
	client := makeClient(POST, url, query, multi, &ClientOption{
		BodyType: &bt,
	})
	return client.do(context.Background())
}

// TimeOut ...
func (obj *Client) TimeOut() time.Duration {
	return time.Duration(util.MustInt64(obj.timeout, defaultTimeout)) * time.Second
}

// KeepAlive ...
func (obj *Client) KeepAlive() time.Duration {
	return time.Duration(util.MustInt64(obj.keepAlive, defaultKeepAlive)) * time.Second
}

func buildTransport(client *Client) (*http.Client, error) {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout:   client.TimeOut(),
				KeepAlive: client.KeepAlive(),
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
	}, nil

}

func buildSafeTransport(client *Client) (*http.Client, error) {
	cert, e := tls.X509KeyPair(client.certKey())
	if e != nil {
		log.Error(e)
		return nil, e
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(client.CA())
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: true, //client端略过对证书的校验
	}
	tlsConfig.BuildNameToCertificate()
	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   client.TimeOut(),
				KeepAlive: client.KeepAlive(),
				//DualStack: true,
			}).DialContext,
			TLSClientConfig: tlsConfig,
			Proxy:           nil,
			//TLSHandshakeTimeout:   10 * time.Second,
			//ResponseHeaderTimeout: 10 * time.Second,
			//ExpectContinueTimeout: 1 * time.Second,
		},
	}, nil
}

func buildHTTPClient(client *Client) (*http.Client, error) {
	//检查是否包含security
	if client.IsSafe() {
		//判断能否创建safe client
		return buildSafeTransport(client)
	}
	return buildTransport(client)

}

// BodyReader ...
type BodyReader interface {
	ToMap() util.Map
	Bytes() []byte
	Error() error
	Unmarshal(v interface{}) error
	Result() (util.Map, error)
}

/*readBody get response data */
func readBody(r io.ReadCloser) ([]byte, error) {
	return ioutil.ReadAll(io.LimitReader(r, math.MaxUint32))
}
