package wego

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Request(url string, params Map, method string, options map[string]Map) Map
	RequestRaw(url string, params Map, method string, options map[string]Map) []byte
	SafeRequest(url string, params Map, method string, options map[string]Map) Map
	Link(string) string
}

type client struct {
	Config
	app     Application
	request *Request
}

func (c *client) Request(url string, params Map, method string, options map[string]Map) Map {
	params.Set("mch_id", c.Get("mch_id"))
	params.Set("nonce_str", GenerateUUID())
	params.Set("sub_mch_id", c.Get("sub_mch_id"))
	params.Set("sub_appid", c.Get("sub_appid"))

	params.Set("sign_type", SIGN_TYPE_MD5.String())

	params.Set("sign", GenerateSignature(params, c.Get("aes_key"), SIGN_TYPE_MD5))

	resp, err := c.request.PerformRequest(c.buildTransport(), c.Link(url), "post", options)
	if err != nil {
		Println(err)
		return nil
	}
	return XmlToMap(resp)
}

func (c *client) RequestRaw(url string, params Map, method string, options map[string]Map) []byte {
	m := c.Request(url, params, method, options)
	if m != nil {
		return m.ToJson()
	}
	return []byte(nil)
}

func (c *client) SafeRequest(url string, params Map, method string, options map[string]Map) Map {
	params.Set("mch_id", c.Get("mch_id"))
	params.Set("nonce_str", GenerateUUID())
	params.Set("sub_mch_id", c.Get("sub_mch_id"))
	params.Set("sub_appid", c.Get("sub_appid"))
	params.Set("sign_type", SIGN_TYPE_MD5.String())
	params.Set("sign", GenerateSignature(params, c.Get("aes_key"), SIGN_TYPE_MD5))

	resp, err := c.request.PerformRequest(c.buildSafeTransport(), c.Link(url), "post", options)
	if err != nil {
		Println(err)
		return nil
	}
	return XmlToMap(resp)
}

func (c *client) Link(url string) string {
	if c.GetBool("sandbox") {
		return DomainUrl() + SANDBOX_URL_SUFFIX + url
	}
	return DomainUrl() + url
}

func NewClient(application Application, request *Request) Client {
	return &client{
		Config:  application.Config(),
		app:     application,
		request: request,
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
