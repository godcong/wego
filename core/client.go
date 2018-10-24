package core

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*data types*/
const (
	DataTypeXML       = "xml"
	DataTypeJSON      = "json"
	DataTypeQuery     = "query"
	DataTypeMultipart = "multipart"
)

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

/*Client Client */
type Client struct {
	context.Context
	config       *Config
	requestType  string
	responseData []byte
	httpRequest  *http.Request
	//httpResponse *http.Response
	httpClient *http.Client
	//app      *Application
	//token    *AccessToken
	//request  *net.Request
	//client *http.Client
}

//Config get client config
func (c *Client) Config() *Config {
	return c.config
}

//SetConfig set client config
func (c *Client) SetConfig(config *Config) {
	c.config = config
}

//RequestType get client request type
func (c *Client) RequestType() string {
	return c.requestType
}

//SetRequestType set client request type
func (c *Client) SetRequestType(requestType string) {
	c.requestType = requestType
}

//ResponseData get client response data
func (c *Client) ResponseData() []byte {
	return c.responseData
}

//SetResponseData set client response data
func (c *Client) SetResponseData(responseData []byte) {
	c.responseData = responseData
}

//HTTPClient get client http Client
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

//SetHTTPClient set client http Client
func (c *Client) SetHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

////HTTPResponse get client http Response
//func (c *Client) HTTPResponse() *http.Response {
//	return c.httpResponse
//}
//
////SetHTTPResponse set client http Response
//func (c *Client) SetHTTPResponse(httpResponse *http.Response) {
//	c.httpResponse = httpResponse
//}

//HTTPRequest get client http Request
func (c *Client) HTTPRequest() *http.Request {
	return c.httpRequest
}

//SetHTTPRequest set client http Request
func (c *Client) SetHTTPRequest(httpRequest *http.Request) {
	c.httpRequest = httpRequest
}

/*PostJSON json post请求 */
func (c *Client) PostJSON(url string, query util.Map, json interface{}) Response {
	c.requestType = DataTypeJSON
	p := util.Map{
		DataTypeJSON: json,
	}
	if query != nil {
		p.Set(DataTypeQuery, query)
	}
	return c.Request(url, "post", p)
}

/*PostXML xml post请求 */
func (c *Client) PostXML(url string, query util.Map, xml interface{}) Response {
	c.requestType = DataTypeXML
	p := util.Map{
		DataTypeXML: xml,
	}
	if query != nil {
		p.Set(DataTypeQuery, query)
	}
	return c.Request(url, "post", p)
}

/*Upload upload请求 */
func (c *Client) Upload(url string, query, multi util.Map) Response {
	c.requestType = DataTypeMultipart
	p := util.Map{
		DataTypeMultipart: multi,
	}
	if query != nil {
		p.Set(DataTypeQuery, query)
	}

	return c.Request(url, "post", p)
}

/*Get get请求 */
func (c *Client) Get(url string, query util.Map) Response {
	p := util.Map{}
	if query != nil {
		p.Set(DataTypeQuery, query)
	}
	return c.Request(url, "get", p)
}

//GetRaw get request result raw data
func (c *Client) GetRaw(url string, query util.Map) []byte {
	p := util.Map{}
	if query != nil {
		p.Set(DataTypeQuery, query)
	}
	return c.RequestRaw(url, "get", p)
}

/*Post post请求 */
func (c *Client) Post(url string, query util.Map, ops util.Map) Response {
	p := util.Map{}
	if query != nil {
		p.Set(DataTypeQuery, query)
	}
	if ops != nil {
		p.ReplaceJoin(ops)
	}
	return c.Request(url, "post", p)
}

/*Request 普通请求 */
func (c *Client) Request(url string, method string, ops util.Map) Response {
	log.Debug("Request|httpClient", c.httpClient)
	c.httpClient = buildTransport(c.config)
	return request(c, url, method, ops)
}

func castToResponse(resp *http.Response) Response {
	ct := resp.Header.Get("Content-Type")
	body, err := ParseBody(resp)
	body = removeControlCharacters(body) //Learn From Easywechat
	if err != nil {
		return Err(body, err)
	}

	log.Println("header:", ct)
	if resp.StatusCode == 200 {
		if strings.Index(ct, "xml") != -1 ||
			bytes.Index(body, []byte("<xml")) != -1 {
			return &responseXML{
				Data: body,
			}
		}
		return &responseJSON{
			Data: body,
		}
	}

	return Err(body, errors.New("error with "+resp.Status))
}

func removeControlCharacters(body []byte) []byte {
	pat := "/[\x00-\x1F\x80-\x9F]/u"
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	return re.ReplaceAll(body, []byte(""))
}

/*RequestRaw raw请求 */
func (c *Client) RequestRaw(url string, method string, ops util.Map) []byte {
	log.Debug("Request|httpClient", c.httpClient)
	c.httpClient = buildTransport(c.config)
	return request(c, url, method, ops).Bytes()
}

/*SafeRequest 安全请求 */
func (c *Client) SafeRequest(url string, method string, ops util.Map) Response {
	c.httpClient = buildSafeTransport(c.config)
	log.Debug("SafeRequest|httpClient", c.httpClient)
	return request(c, url, method, ops)
}

/*SafeRequestRaw 安全请求 */
func (c *Client) SafeRequestRaw(url string, method string, ops util.Map) []byte {
	c.httpClient = buildSafeTransport(c.config)
	log.Debug("SafeRequest|httpClient", c.httpClient)
	return request(c, url, method, ops).Bytes()

}

/*NewClient 创建一个client */
func NewClient(config *Config) *Client {
	return &Client{
		Context:      context.Background(),
		config:       config,
		requestType:  DataTypeXML,
		responseData: nil,
		httpRequest:  nil,
		//httpResponse: nil,
		httpClient: nil,
	}
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
	if idx := config.Check("cert_path", "key_path"); idx != 0 {
		panic(fmt.Sprintf("the %d key was not found", idx))
	}

	cert, err := tls.LoadX509KeyPair(config.GetString("cert_path"), config.GetString("key_path"))
	if err != nil {
		panic(err)
	}

	caFile, err := ioutil.ReadFile(config.GetString("rootca_path"))
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

func (c *Client) clear() {
	c.httpRequest = nil
	c.httpClient = nil
	//c.httpResponse = nil
}

func request(client *Client, url string, method string, ops util.Map) Response {
	method = strings.ToUpper(method)
	query := buildHttpQuery(ops.Get(DataTypeQuery))
	url = parseQuery(url, query)

	if client.httpRequest == nil {
		newRequest := requestData(client.requestType)
		client.httpRequest = newRequest(method, url, ops.Get(client.requestType))
	}

	defer client.clear()

	log.Debug("client|request", client, url, method, ops)
	response, err := client.httpClient.Do(client.httpRequest.WithContext(client.Context))
	if err != nil {
		log.Error("Client|Do", err)
		return Err(nil, err)
	}
	{
		select {
		case <-client.Context.Done():
			return Err(nil, client.Context.Err())
		default:
			//return Err(nil, err)
		}
	}
	return castToResponse(response)
}

func requestData(dt string) func(string, string, interface{}) *http.Request {
	if dt == DataTypeJSON {
		return processJSON
	} else if dt == DataTypeXML {
		return processXML
	} else if dt == DataTypeMultipart {
		return processMultipart
	} else {

	}
	return processNothing
}

func processNothing(method, url string, i interface{}) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}
	return request
}

func processMultipart(method, url string, i interface{}) *http.Request {
	buf := bytes.Buffer{}
	writer := multipart.NewWriter(&buf)
	defer writer.Close()
	log.Debug("processMultipart|i", i)
	switch v := i.(type) {
	case util.Map:
		path := v.GetString("media")
		fh, e := os.Open(path)
		if e != nil {
			log.Debug("processMultipart|e", e)
			return nil
		}
		defer fh.Close()

		fw, e := writer.CreateFormFile("media", path)
		if e != nil {
			log.Debug("processMultipart|e", e)
			return nil
		}

		if _, e = io.Copy(fw, fh); e != nil {
			log.Debug("processMultipart|e", e)
			return nil
		}
		des := v.GetMap("description")
		if des != nil {
			writer.WriteField("description", string(des.ToJSON()))
		}
	}
	request, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request
}

func processXML(method, url string, i interface{}) *http.Request {
	var reader io.Reader
	switch v := i.(type) {
	case string:
		log.Debug("processXML|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("processXML|[]byte", v)
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("processXML|util.Map", v.ToXML())
		reader = strings.NewReader(v.ToXML())
	default:
		log.Debug("processXML|default", v)
		if v0, e := xml.Marshal(v); e == nil {
			log.Debug("processXML|v0", v0, e)
			reader = bytes.NewReader(v0)
		}
	}

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil
	}
	request.Header["Content-Type"] = []string{"application/xml; charset=utf-8"}
	return request
}

func processJSON(method, url string, i interface{}) *http.Request {
	var reader io.Reader
	switch v := i.(type) {
	case string:
		log.Debug("processJSON|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("processJSON|[]byte", string(v))
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("processJSON|util.Map", v.String())
		reader = bytes.NewReader(v.ToJSON())
	default:
		log.Debug("processJSON|default", v)
		if v0, e := json.Marshal(v); e == nil {
			log.Debug("processJSON|v0", string(v0), e)
			reader = bytes.NewReader(v0)
		}
	}

	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil
	}
	request.Header["Content-Type"] = []string{"application/json; charset=utf-8"}
	return request
}

func buildHttpQuery(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case util.Map:
		return v.URLEncode()
	}
	return ""
}
