package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"golang.org/x/exp/xerrors"
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
	Method string
	URL    string
	Body   *RequestBody
	Option *ClientOption
}

// ClientOption ...
type ClientOption struct {
	UseSafe     bool
	UseToken    bool
	SafeCert    *SafeCertProperty
	AccessToken *AccessToken
	BodyType    *BodyType
	Query       util.Map
	Timeout     int64
	KeepAlive   int64
}

// NewClient ...
func NewClient(opts ...*ClientOption) *Client {
	var opt *ClientOption
	if opts != nil {
		opt = opts[0]
	}
	return &Client{
		Option: opt,
	}
}

// Post ...
func (c *Client) Post(ctx context.Context, url string, body interface{}) Responder {
	c.Method = POST
	c.URL = url
	c.Body = buildBody(body, c.BodyType())
	return c.do(ctx)
}

// Get ...
func (c *Client) Get(ctx context.Context, url string) Responder {
	c.Method = POST
	c.URL = url
	c.Body = buildBody(nil, c.BodyType())
	return c.do(ctx)
}

// BodyType ...
func (c *Client) BodyType() BodyType {
	return bodyType(c.Option)
}
func bodyType(opt *ClientOption) BodyType {
	if opt != nil && opt.BodyType != nil {
		return *opt.BodyType
	}
	return BodyTypeNone
}

// CA ...
func (c *Client) CA() []byte {
	if c.Option != nil && c.Option.SafeCert != nil && c.Option.SafeCert.RootCA != nil {
		return c.Option.SafeCert.RootCA
	}
	return []byte(defaultCa)
}

func (c *Client) certKey() ([]byte, []byte) {
	if c.Option != nil &&
		c.Option.SafeCert != nil &&
		c.Option.SafeCert.Key != nil &&
		c.Option.SafeCert.Cert != nil {
		return c.Option.SafeCert.Cert, c.Option.SafeCert.Key
	}
	return nil, nil
}

// HTTPClient ...
func (c *Client) HTTPClient() (*http.Client, error) {
	return buildHTTPClient(c)
}

// makeClient ...
func makeClient(method string, url string, body interface{}, opts ...*ClientOption) *Client {
	var opt *ClientOption
	if opts != nil {
		opt = opts[0]
	}
	return &Client{
		Method: method,
		URL:    url,
		Body:   buildBody(body, bodyType(opt)),
		Option: opt,
	}
}

// Request ...
func (c *Client) Request() (*http.Request, error) {
	if c.Body == nil {
		return http.NewRequest(c.Method, c.RemoteURL(), nil)
	}
	return c.Body.RequestBuilder(c.Method, c.RemoteURL(), c.Body.BodyInstance)
}

// do ...
func (c *Client) do(ctx context.Context) Responder {
	log.Debugf("%+v\n", c)
	client, e := c.HTTPClient()
	if e != nil {
		return ErrResponse(xerrors.Errorf("client:%w", e))
	}
	request, e := c.Request()
	if e != nil {
		return ErrResponse(xerrors.Errorf("request:%w", e))
	}
	response, e := client.Do(request.WithContext(ctx))
	if e != nil {
		return ErrResponse(xerrors.Errorf("response:%w", e))
	}
	return buildResponder(response)
}

// RemoteURL ...
func (c *Client) RemoteURL() string {
	if c.Option != nil && c.Option.Query != nil {
		return c.URL + "?" + c.Option.Query.URLEncode()
	}
	return c.URL
}

// PostForm post form request
func PostForm(url string, query util.Map, form interface{}) Responder {
	log.Debug("post form:", url, query, form)
	bt := BodyTypeForm
	client := makeClient(POST, url, form, &ClientOption{
		BodyType: &bt,
		Query:    query,
	})
	return client.do(context.Background())
}

// PostJSON json post请求
func PostJSON(url string, query util.Map, json interface{}) Responder {
	log.Debug("post json:", url, query, json)
	bt := BodyTypeJSON
	client := makeClient(POST, url, json, &ClientOption{
		BodyType: &bt,
		Query:    query,
	})
	return client.do(context.Background())
}

// PostXML  xml post请求
func PostXML(url string, query util.Map, xml interface{}) Responder {
	log.Debug("post xml:", url, query, xml)
	bt := BodyTypeXML
	client := makeClient(POST, url, xml, &ClientOption{
		BodyType: &bt,
		Query:    query,
	})
	return client.do(context.Background())
}

// Get get请求
func Get(url string, query util.Map) Responder {
	log.Println("get request:", url, query)
	client := makeClient(GET, url, nil, &ClientOption{
		Query: query,
	})
	return client.do(context.Background())
}

// Upload upload请求
func Upload(url string, query, multi util.Map) Responder {
	bt := BodyTypeMultipart
	client := makeClient(POST, url, multi, &ClientOption{
		BodyType: &bt,
		Query:    query,
	})
	return client.do(context.Background())
}

// TimeOut ...
func TimeOut(src int64) time.Duration {
	return time.Duration(util.MustInt64(src, 30)) * time.Second
}

// KeepAlive ...
func KeepAlive(src int64) time.Duration {
	return time.Duration(util.MustInt64(src, 30)) * time.Second
}

func buildTransport(client *Client) (*http.Client, error) {
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
	}, nil
}

func buildHTTPClient(client *Client) (*http.Client, error) {
	//检查是否包含security
	if client.Option.UseSafe {
		//判断能否创建safe client
		return buildSafeTransport(client)
	}
	return buildTransport(client)

}
