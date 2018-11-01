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
	"net/url"
	"os"
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
	DataTypeForm      = "form"
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
	for _, val := range v {
		switch sv := val.(type) {
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
	for _, val := range v {
		switch sv := val.(type) {
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
func (c *Client) PostForm(url string, query util.Map, form interface{}) Response {
	maps := util.Map{
		DataTypeForm: form,
	}
	if query != nil {
		maps.Set(DataTypeQuery, query)
	}
	return c.Post(url, maps)
}

/*PostJSON json post请求 */
func (c *Client) PostJSON(url string, query util.Map, json interface{}) Response {
	maps := util.Map{
		DataTypeJSON: json,
	}
	if query != nil {
		maps.Set(DataTypeQuery, query)
	}
	return c.Post(url, maps)
}

/*PostXML xml post请求 */
func (c *Client) PostXML(url string, query util.Map, xml interface{}) Response {
	maps := util.Map{
		DataTypeXML: xml,
	}
	if query != nil {
		maps.Set(DataTypeQuery, query)
	}
	return c.Post(url, maps)
}

/*Post post请求 */
func (c *Client) Post(url string, maps util.Map) Response {
	return c.Request(url, "post", maps)
}

/*Upload upload请求 */
func (c *Client) Upload(url string, query, multi util.Map) Response {
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

/*Request 普通请求 */
func (c *Client) Request(url string, method string, ops util.Map) Response {
	log.Debug("Request|httpClient", url, method, ops)
	return request(c.Context, url, method, ops)
}

func castToResponse(resp *http.Response) Response {
	ct := resp.Header.Get("Content-Type")
	body, err := ParseBody(resp)
	if err != nil {
		log.Error(body, err)
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
	log.Error(body, "error with "+resp.Status)
	return Err(body, errors.New("error with "+resp.Status))
}

/*RequestRaw raw请求 */
func (c *Client) RequestRaw(url string, method string, ops util.Map) []byte {
	log.Debug("Request|httpClient", url, method, ops)
	return request(c.Context, url, method, ops).Bytes()
}

///*SafeRequest 安全请求 */
//func (c *Client) SafeRequest(url string, method string, ops util.Map) Response {
//	log.Debug("SafeRequest|httpClient", url, method, ops)
//	return request(c.Context, url, method, ops)
//}
//
///*SafeRequestRaw 安全请求 */
//func (c *Client) SafeRequestRaw(url string, method string, ops util.Map) []byte {
//	log.Debug("SafeRequest|httpClient", url, method, ops)
//	return request(c.Context, url, method, ops).Bytes()
//}

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

	if idx := config.Check("cert_path", "key_path"); idx != -1 {
		panic(fmt.Sprintf("the %d key was not found", idx))
	}

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
	//c.httpRequest = nil
	//c.httpClient = nil
	//c.httpResponse = nil
}

func request(context context.Context, url string, method string, ops util.Map) Response {
	method = strings.ToUpper(method)
	query := buildHTTPQuery(ops.Get(DataTypeQuery))
	client := buildClient(ops)
	url = connectQuery(url, query)
	req := requestData(method, url, ops)

	log.Debug("client|request", client, url, method, ops)
	response, err := client.Do(req.WithContext(context))
	if err != nil {
		log.Error("Client|Do", err)
		return Err(nil, err)
	}
	{
		select {
		case <-context.Done():
			return Err(nil, context.Err())
		default:
			//return Err(nil, err)
		}
	}
	return castToResponse(response)
}
func buildClient(maps util.Map) *http.Client {
	//检查是否包含security
	if maps.Has(DataTypeSecurity) {
		//判断是否创建safeclient
		if v, b := maps.Get(DataTypeSecurity).(*Config); b {
			if v.Check("cert_path", "key_path") != -1 {
				return buildTransport(v)
			}
			return buildSafeTransport(v)
		}
	}
	return buildTransport(C(util.Map{}))
}

func requestData(method, url string, m util.Map) *http.Request {
	function := processNothing
	var data interface{}

	if m.Has(DataTypeJSON) {
		function = processJSON
		data = m.Get(DataTypeJSON)
	} else if m.Has(DataTypeXML) {
		function = processXML
		data = m.Get(DataTypeXML)
	} else if m.Has(DataTypeForm) {
		function = processForm
		data = m.Get(DataTypeForm)
	} else if m.Has(DataTypeMultipart) {
		function = processMultipart
		data = m.Get(DataTypeMultipart)
	} else {

	}

	return function(method, url, data)
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

func toXMLReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("toXMLReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("toXMLReader|[]byte", v)
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("toXMLReader|util.Map", v.ToXML())
		reader = strings.NewReader(v.ToXML())
	default:
		log.Debug("toXMLReader|default", v)
		if v0, e := xml.Marshal(v); e == nil {
			log.Debug("toXMLReader|v0", v0, e)
			reader = bytes.NewReader(v0)
		}
	}
	return reader
}

func processXML(method, url string, i interface{}) *http.Request {

	request, err := http.NewRequest(method, url, toXMLReader(i))
	if err != nil {
		return nil
	}
	request.Header["Content-Type"] = []string{"application/xml; charset=utf-8"}
	return request
}

func toJSONReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("toJSONReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("toJSONReader|[]byte", string(v))
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("toJSONReader|util.Map", v.String())
		reader = bytes.NewReader(v.ToJSON())
	default:
		log.Debug("toJSONReader|default", v)
		if v0, e := json.Marshal(v); e == nil {
			log.Debug("toJSONReader|v0", string(v0), e)
			reader = bytes.NewReader(v0)
		}
	}
	return reader
}

func processJSON(method, url string, i interface{}) *http.Request {
	request, err := http.NewRequest(method, url, toJSONReader(i))
	if err != nil {
		return nil
	}
	request.Header["Content-Type"] = []string{"application/json; charset=utf-8"}
	return request
}

func toFormReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("toFormReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("toFormReader|[]byte", string(v))
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("toFormReader|util.Map", v.URLEncode())
		reader = strings.NewReader(v.URLEncode())
	case url.Values:
		log.Debug("toFormReader|util.Map", v.Encode())
		reader = strings.NewReader(v.Encode())
	default:
		//do nothing
	}
	return reader
}

func processForm(method, url string, i interface{}) *http.Request {

	request, err := http.NewRequest(method, url, toFormReader(i))
	if err != nil {
		return nil
	}

	request.Header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=utf-8"}
	return request
}

func buildHTTPQuery(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case util.Map:
		return v.URLEncode()
	}
	return ""
}
