package core

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"time"
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

// RequestBody ...
type RequestBody struct {
	BodyType string
	BodyInst interface{}
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
		client:   buildTransport(NilConfig()),
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

func buildTransport(config *Config, option ...util.Map) *http.Client {
	m := util.CombineMaps(util.Map{
		"time_out":   30,
		"keep_alive": 30,
	}, option)
	timeOut, _ := m.GetInt64("time_out")
	keepAlive, _ := m.GetInt64("keep_alive")

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

func buildSafeTransport(config *Config, option ...util.Map) *http.Client {
	if config == nil {
		panic("safe request must set config before use")
	}

	m := util.CombineMaps(util.Map{
		"time_out":   30,
		"keep_alive": 30,
	}, option)
	timeOut, _ := m.GetInt64("time_out")
	keepAlive, _ := m.GetInt64("keep_alive")

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
		return buildTransport(NilConfig())

	}
	//默认输出未配置client
	log.Debug("default client")

	return buildTransport(NilConfig())
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
