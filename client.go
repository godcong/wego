package wego

import (
	"context"
	"crypto/tls"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
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

// RequestContent ...
type RequestContent struct {
	Method string
	URL    string
	Query  util.Map
	Body   *RequestBody
}

// Client ...
type Client struct {
	context   context.Context
	TLSConfig *tls.Config
	BodyType  BodyType
	//useSafe   bool
	//safeCert    *SafeCertProperty
	//accessToken *AccessToken
	//timeout     int64
	//keepAlive   int64
}

// UseSafe ...
func (obj *Client) UseSafe() bool {
	return obj.TLSConfig != nil
}

// Context ...
func (obj *Client) Context() context.Context {
	if obj.context == nil {
		i, _ := Context()
		return i
	}
	return obj.context
}

// NewClient ...
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		BodyType: BodyTypeXML,
	}
	client.parse(options...)
	return client
}

func (obj *Client) parse(options ...ClientOption) {
	if options == nil {
		return
	}

	for _, o := range options {
		o(obj)
	}
}

// Post ...
func (obj *Client) Post(ctx context.Context, url string, query util.Map, body interface{}) Responder {
	log.Debug("post ", url, body)
	return obj.do(ctx, &RequestContent{
		Method: POST,
		URL:    url,
		Query:  query,
		Body:   buildBody(body, obj.BodyType),
	})
}

// Get ...
func (obj *Client) Get(ctx context.Context, url string, query util.Map) Responder {
	log.Debug("get ", url)
	return obj.do(ctx, &RequestContent{
		Method: POST,
		URL:    url,
		Query:  query,
		Body:   buildBody(nil, obj.BodyType),
	})
}

// HTTPClient ...
func (obj *Client) HTTPClient() (*http.Client, error) {
	return buildHTTPClient(obj, obj.UseSafe())
}

// do ...
func (obj *Client) do(ctx context.Context, content *RequestContent) Responder {
	client, e := obj.HTTPClient()
	if e != nil {
		return ErrResponder(xerrors.Errorf("client build err:%+v", e))
	}
	request, e := content.BuildRequest()
	if e != nil {
		return ErrResponder(xerrors.Errorf("request build err:%+v", e))
	}
	response, e := client.Do(request.WithContext(ctx))
	if e != nil {
		return ErrResponder(xerrors.Errorf("response get err:%+v", e))
	}
	return BuildResponder(response)
}

// PostForm post form request
func PostForm(url string, query util.Map, form interface{}) Responder {
	log.Debug("post form:", url, query, form)
	client := &Client{
		BodyType: BodyTypeForm,
	}
	return client.Post(context.Background(), url, query, form)
}

// PostJSON json post请求
func PostJSON(url string, query util.Map, json interface{}) Responder {
	log.Debug("post json:", url, query, json)
	client := &Client{
		BodyType: BodyTypeJSON,
	}
	return client.Post(context.Background(), url, query, json)
}

// PostXML  xml post请求
func PostXML(url string, query util.Map, xml interface{}) Responder {
	log.Debug("post xml:", url, query, xml)
	client := &Client{
		BodyType: BodyTypeXML,
	}
	return client.Post(context.Background(), url, query, xml)
}

// Get get请求
func Get(url string, query util.Map) Responder {
	log.Println("get request:", url, query)
	client := NewClient()
	return client.Get(context.Background(), url, query)
}

// Upload upload请求
func Upload(url string, query, multi util.Map) Responder {
	client := NewClient()
	return client.do(context.Background(), &RequestContent{
		Method: POST,
		URL:    url,
		Query:  query,
		Body:   buildBody(multi, BodyTypeMultipart),
	})
}

// Context ...
func Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func buildTransport(client *Client) (*http.Transport, error) {
	return &http.Transport{
		Proxy:       nil,
		DialContext: (&net.Dialer{
			//Timeout:   client.TimeOut(),
			//KeepAlive: client.KeepAlive(),
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
	}, nil

}

func buildSafeTransport(client *Client) (*http.Transport, error) {
	return &http.Transport{
		DialContext: (&net.Dialer{
			//Timeout:   client.TimeOut(),
			//KeepAlive: client.KeepAlive(),
			//DualStack: true,
		}).DialContext,
		TLSClientConfig: client.TLSConfig,
		Proxy:           nil,
		//TLSHandshakeTimeout:   10 * time.Second,
		//ResponseHeaderTimeout: 10 * time.Second,
		//ExpectContinueTimeout: 1 * time.Second,
	}, nil
}

func buildHTTPClient(client *Client, isSafe bool) (cli *http.Client, e error) {
	cli = new(http.Client)
	//判断能否创建safe client
	fun := buildTransport
	if isSafe {
		fun = buildSafeTransport
	}
	cli.Transport, e = fun(client)
	return
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

// BuildRequest ...
func (c *RequestContent) BuildRequest() (*http.Request, error) {
	if c.Body == nil {
		return http.NewRequest(c.Method, c.URLQuery(), nil)
	}
	return c.Body.RequestBuilder(c.Method, c.URLQuery(), c.Body.BodyInstance)
}

// URLQuery ...
func (c *RequestContent) URLQuery() string {
	if c.Query == nil {
		return c.URL
	}
	return c.URL + "?" + c.Query.URLEncode()

}
