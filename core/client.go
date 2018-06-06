package core

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

/*DataType DataType */
type DataType string
type ContextKey struct{}

const (
	DataTypeXML       DataType = "xml"
	DataTypeJSON      DataType = "json"
	DataTypeMultipart DataType = "multipart"
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
	//request  *net.Request
	//response *net.Response
	client *http.Client
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

func (c *Client) HTTPClient() *http.Client {
	return c.client
}

func (c *Client) SetHTTPClient(client *http.Client) *Client {
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

func (c *Client) HTTPPostJSON(url string, query util.Map, json interface{}) *net.Response {
	c.dataType = DataTypeJSON
	p := util.Map{
		net.REQUEST_TYPE_JSON.String(): json,
	}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "post")
}

func (c *Client) HTTPPostXML(url string, query util.Map, xml interface{}) *net.Response {
	c.dataType = DataTypeXML
	p := util.Map{
		net.REQUEST_TYPE_XML.String(): xml,
	}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "post")
}

func (c *Client) HTTPUpload(url string, query, multi util.Map) *net.Response {
	c.dataType = DataTypeMultipart
	p := util.Map{
		net.REQUEST_TYPE_MULTIPART.String(): multi,
	}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}

	return c.Request(url, p, "post")
}

func (c *Client) HTTPGet(url string, query util.Map) *net.Response {
	p := util.Map{}
	if query != nil {
		p.Set(net.REQUEST_TYPE_QUERY.String(), query)
	}
	return c.Request(url, p, "get")
}

func (c *Client) HTTPPost(url string, query util.Map, ops util.Map) *net.Response {
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
	return request(c, url, ops, method)
}

func (c *Client) RequestRaw(url string, ops util.Map, method string) *net.Response {
	return c.Request(url, ops, method)
}

func (c *Client) SafeRequest(url string, ops util.Map, method string) *net.Response {
	c.client = buildSafeTransport(c.Config)
	log.Debug("SafeRequest|httpClient", c.client)
	return request(c, url, ops, method)
}

func (c *Client) Link(uri string) string {
	if c.GetBool("sandbox") {
		return c.URL() + sandboxUrlSuffix + uri
	}
	return c.domain.Link(uri)
}

/*GetRequest get net response */
//func (c *Client) GetResponse() *net.Response {
//return c.response
//}

/*GetRequest get net request */
//func (c *Client) GetRequest() *net.Request {
//	return c.request
//}

/*DefaultClient DefaultClient */
func DefaultClient() *Client {
	return nil
}

/*NewClient 创建一个client */
func NewClient(config config.Config) *Client {
	log.Debug("NewClient|config", config)
	domain := NewDomain("default")
	if config == nil {
		config = defaultConfig
	}
	return &Client{
		//request:  net.DefaultRequest,
		Config:   config,
		dataType: DataTypeXML,
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
	r := net.PerformRequest(url, method, data)
	if r.Error() == nil {
		return Do(context.Background(), c, r)
	}
	return net.ErrorResponse(r.Error())
}

/*Do do request */
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
	response = net.ParseResponse(request.GetRequestType(), r)

	return response
}

func toRequestData(client *Client, ops util.Map) *net.RequestData {
	data := net.NewRequestData()
	data.Query = processQuery(ops.Get(net.REQUEST_TYPE_QUERY.String()))
	data.Body = nil
	if client.DataType() == DataTypeJSON {
		data.SetHeaderJson()
		data.Body = processJSON(ops.Get(net.REQUEST_TYPE_JSON.String()))
	}
	if client.DataType() == DataTypeXML {
		data.SetHeaderXml()
		data.Body = processXML(ops.Get(net.REQUEST_TYPE_XML.String()))
	}

	if client.DataType() == DataTypeMultipart {
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
func processXML(i interface{}) io.Reader {
	switch v := i.(type) {
	case string:
		log.Debug("processXML|string", v)
		return strings.NewReader(v)
	case []byte:
		log.Debug("processXML|[]byte", v)
		return bytes.NewReader(v)
	case util.Map:
		log.Debug("processXML|util.Map", v.ToXml())
		return strings.NewReader(v.ToXml())
	default:
		log.Debug("processXML|default", v)
		if v0, e := xml.Marshal(v); e == nil {
			log.Debug("processXML|v0", v0, e)
			return bytes.NewReader(v0)
		}
		return nil
	}

}

func processJSON(i interface{}) io.Reader {
	switch v := i.(type) {
	case string:
		log.Debug("processJSON|string", v)
		return strings.NewReader(v)
	case []byte:
		log.Debug("processJSON|[]byte", string(v))
		return bytes.NewReader(v)
	case util.Map:
		log.Debug("processJSON|util.Map", v.String())
		return bytes.NewReader(v.ToJson())
	default:
		log.Debug("processJSON|default", v)
		if v0, e := json.Marshal(v); e == nil {
			log.Debug("processJSON|v0", string(v0), e)
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

/*ShortURL 转换短链接
https://apihk.mch.weixin.qq.com/tools/shorturl    （建议接入点：东南亚）
https://apius.mch.weixin.qq.com/tools/shorturl    （建议接入点：其它）
https://api.mch.weixin.qq.com/tools/shorturl        （建议接入点：中国国内）
是否需要证书
否
请求参数
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID（企业号corpid即为此appId）
商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
URL链接	long_url	是	String(512、	weixin：//wxpay/bizpayurl?sign=XXXXX&appid=XXXXX&mch_id=XXXXX&product_id=XXXXXX&time_stamp=XXXXXX&nonce_str=XXXXX	需要转换的URL，签名用原串，传输需URLencode
随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
签名类型	sign_type	否	String(32)	HMAC-SHA256	签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
返回结果
字段名	变量名	必填	类型	示例值	描述
返回状态码	return_code	是	String(16)	SUCCESS
SUCCESS/FAIL
此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
返回信息	return_msg	否	String(128)	签名失败
返回信息，如非空，为错误原因签名失败
参数格式校验错误
以下字段在return_code为SUCCESS的时候有返回
字段名	变量名	必填	类型	示例值	描述
公众账号ID	appid	是	String(32)	wx8888888888888888	微信分配的公众账号ID
商户号	mch_id	是	String(32)	1900000109	微信支付分配的商户号
随机字符串	nonce_str	是	String(32)	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	随机字符串，不长于32位。推荐随机数生成算法
签名	sign	是	String(32)	C380BEC2BFD727A4B6845133519F3AD6	签名，详见签名生成算法
业务结果	result_code	是	String(16)	SUCCESS	SUCCESS/FAIL
错误代码	err_code	否	String(32)	SYSTEMERROR
SYSTEMERROR—系统错误
URLFORMATERROR—URL格式错误
URL链接	short_url	是	String(64)	weixin：//wxpay/s/XXXXXX	转换后的URL
错误码
名称	描述	原因	解决方案
SIGNERROR	签名错误	参数签名结果不正确	请检查签名参数和方法是否都符合签名算法要求
REQUIRE_POST_METHOD	请使用post方法	未使用post传递参数	请检查请求参数是否通过post方法提交
APPID_NOT_EXIST	APPID不存在	参数中缺少APPID	请检查APPID是否正确
MCHID_NOT_EXIST	MCHID不存在	参数中缺少MCHID	请检查MCHID是否正确
APPID_MCHID_NOT_MATCH	appid和mch_id不匹配	appid和mch_id不匹配	请确认appid和mch_id是否匹配
LACK_PARAMS	缺少参数	缺少必要的请求参数	请检查参数是否齐全
XML_FORMAT_ERROR	XML格式错误	XML格式错误	请检查XML参数格式是否正确
POST_DATA_EMPTY	post数据为空	post数据不能为空	请检查post数据是否为空
*/
func (u *URL) ShortURL(url string) util.Map {
	m := util.Map{
		"action":   "long2short",
		"long_url": url,
	}
	token := u.token.GetToken()
	ops := util.Map{
		net.REQUEST_TYPE_QUERY.String(): token.KeyMap(),
	}
	resp := u.client.HTTPPostJSON(u.client.domain.Link(shortURLSuffix), m, ops)
	log.Debug("URL|ShortURL", *resp)
	return resp.ToMap()
}

/*NewURL NewURL*/
func NewURL(config config.Config, client *Client) *URL {
	return &URL{
		token:  NewAccessToken(config, client),
		client: client,
	}
}
