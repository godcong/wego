package core

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

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
	httpResponse *http.Response
	httpClient   *http.Client
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

//HTTPResponse get client http Response
func (c *Client) HTTPResponse() *http.Response {
	return c.httpResponse
}

//SetHTTPResponse set client http Response
func (c *Client) SetHTTPResponse(httpResponse *http.Response) {
	c.httpResponse = httpResponse
}

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
	response, err := request(c, url, method, ops)

	log.Println(response.Header)
	log.Println(response.ContentLength)
	log.Println(response.Status)
	log.Println(response.StatusCode)
	if err != nil {
		return parseResponse(nil, err.Error())
	}
	b, err := ParseBody(response)
	if err != nil {
		return parseResponse(nil, err.Error())
	}
	return parseResponse(b, c.requestType)
}

func parseResponse(b []byte, t string) Response {
	if t == DataTypeXML {
		return &responseXML{
			Data: b,
		}
	} else if t == DataTypeJSON {
		return &responseJSON{
			Data: b,
		}
	}
	return &responseError{
		Data: b,
		Err:  errors.New(t),
	}
}

/*RequestRaw raw请求 */
func (c *Client) RequestRaw(url string, method string, ops util.Map) []byte {
	log.Debug("Request|httpClient", c.httpClient)
	c.httpClient = buildTransport(c.config)
	response, err := request(c, url, method, ops)
	if err != nil {
		return nil
	}
	b, err := ParseBody(response)
	if err != nil {
		return nil
	}
	return b
}

/*SafeRequest 安全请求 */
func (c *Client) SafeRequest(url string, method string, ops util.Map) Response {
	c.httpClient = buildSafeTransport(c.config)
	log.Debug("SafeRequest|httpClient", c.httpClient)
	response, err := request(c, url, method, ops)
	if err != nil {
		return parseResponse(nil, err.Error())
	}
	b, err := ParseBody(response)
	if err != nil {
		return parseResponse(nil, err.Error())
	}
	return parseResponse(b, c.requestType)
}

/*SafeRequestRaw 安全请求 */
func (c *Client) SafeRequestRaw(url string, method string, ops util.Map) []byte {
	c.httpClient = buildSafeTransport(c.config)
	log.Debug("SafeRequest|httpClient", c.httpClient)
	response, err := request(c, url, method, ops)
	if err != nil {
		return nil
	}
	b, err := ParseBody(response)
	if err != nil {
		return nil
	}
	return b
}

///*Link 拼接地址 */
//func (c *Client) Link(uri string) string {
//	if c.GetBool("sandbox") {
//		return c.domain.URL() + sandboxURLSuffix + uri
//	}
//	return c.domain.Link(uri)
//}

/*GetRequest get net response */
//func (c *Client) GetResponse() core.Response {
//return c.response
//}

/*GetRequest get net request */
//func (c *Client) GetRequest() *net.Request {
//	return c.request
//}

/*NewClient 创建一个client */
func NewClient(config *Config) *Client {
	return &Client{
		Context:      context.Background(),
		config:       config,
		requestType:  DataTypeXML,
		responseData: nil,
		httpRequest:  nil,
		httpResponse: nil,
		httpClient:   nil,
	}
}

func buildTransport(config *Config) *http.Client {
	_ = config
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

func request(client *Client, url string, method string, ops util.Map) (*http.Response, error) {
	method = strings.ToUpper(method)
	query := processQuery(ops.Get(DataTypeQuery))
	url = parseQuery(url, query)

	if client.httpRequest == nil {
		newRequest := requestData(client.requestType)
		client.httpRequest = newRequest(method, url, ops.Get(client.requestType))
	}

	defer func() {
		//clear request when done
		client.httpRequest = nil
	}()

	log.Debug("client|request", client, url, method, ops)
	response, err := http.DefaultClient.Do(client.httpRequest.WithContext(client.Context))
	if err != nil {
		log.Error("Client|Do", err)
		select {
		case <-client.Context.Done():
			return nil, err
		default:
			return nil, client.Context.Err()
		}

		return nil, err
	}
	return response, err
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

func processQuery(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case util.Map:
		return v.URLEncode()
	}
	return ""
}
