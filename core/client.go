package core

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client interface {
	Request(url string, params Map, method string, options Map) Map
	RequestRaw(url string, params Map, method string, options Map) []byte
	SafeRequest(url string, params Map, method string, options Map) Map
	Link(string) string
}

type client struct {
	Config
	request *Request
}

func (c *client) Request(url string, params Map, method string, options Map) Map {
	resp := request(c, c.buildTransport(), url, params, method, options)
	if strings.Index(string(resp), "<xml>") == 0 {
		return XmlToMap(resp)
	}
	return JsonToMap(resp)
}

func (c *client) RequestRaw(url string, params Map, method string, options Map) []byte {
	return request(c, c.buildTransport(), url, params, method, options)
}

func (c *client) SafeRequest(url string, params Map, method string, options Map) Map {
	resp := request(c, c.buildSafeTransport(), url, params, method, options)
	return XmlToMap(resp)
}

func (c *client) Link(url string) string {
	if c.GetBool("sandbox") {
		return DomainUrl() + SANDBOX_URL_SUFFIX + url
	}
	return DomainUrl() + url
}

func NewClient(request *Request, config Config) Client {
	return &client{
		request: request,
		Config:  config,
	}
}

func (c *client) buildTransport() *http.Transport {
	return &http.Transport{
		//Dial: (&net.Dialer{
		//	Timeout:   30 * time.Second,
		//	KeepAlive: 30 * time.Second,
		//}).Dial,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		//TLSHandshakeTimeout:   10 * time.Second,
		//ResponseHeaderTimeout: 10 * time.Second,
		//ExpectContinueTimeout: 1 * time.Second,
	}

}

func (c *client) buildSafeTransport() *http.Transport {
	cert, err := tls.LoadX509KeyPair(c.Get("cert_path"), c.Get("key_path"))
	if err != nil {
		Println(err)
		return nil
	}

	caFile, err := ioutil.ReadFile(c.Get("rootca_path"))
	if err != nil {
		Println(err)
		return nil
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caFile)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	}
	tlsConfig.BuildNameToCertificate()
	return &http.Transport{
		//Dial: (&net.Dialer{
		//	Timeout:   30 * time.Second,
		//	KeepAlive: 30 * time.Second,
		//}).Dial,
		TLSClientConfig: tlsConfig,
		//TLSHandshakeTimeout:   10 * time.Second,
		//ResponseHeaderTimeout: 10 * time.Second,
		//ExpectContinueTimeout: 1 * time.Second,
	}
}

func request(c *client, transport *http.Transport, url string, params Map, method string, op Map) []byte {
	op = MakeOption(op)
	if params != nil {
		params.Set("mch_id", c.Get("mch_id"))
		params.Set("nonce_str", GenerateUUID())
		params.Set("sub_mch_id", c.Get("sub_mch_id"))
		params.Set("sub_appid", c.Get("sub_appid"))
		params.Set("sign_type", SIGN_TYPE_MD5.String())
		params.Set("sign", GenerateSignature(params, c.Get("aes_key"), SIGN_TYPE_MD5))
	}
	op["body"] = params
	resp, err := c.request.PerformRequest(transport, url, method, op)
	if err != nil {
		Error(err)
		return []byte(nil)
	}
	return resp
}
