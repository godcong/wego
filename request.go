package wego

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

//
//import (
//	"crypto/tls"
//
//	"net/http"
//
//	"bytes"
//
//	"crypto/x509"
//	"errors"
//	"io/ioutil"
//	"log"
//	"time"
//
//	"github.com/godcong/wopay/util"
//)
//
//type PayRequest struct {
//	config PayConfig
//}
//
//var (
//	ErrorNilDomain       = errors.New("PayConfig.PayDomain().getDomain() is empty or null")
//	ErrorLoadX509KeyPair = errors.New("LoadX509KeyPair() is empty to load")
//	ErrorReadRootCAFile  = errors.New("read rootca.pem file error")
//)
//
//func NewPayRequest(config PayConfig) *PayRequest {
//	return &PayRequest{config: config}
//}
var requestInst Request

type Request interface {
	SafeRequest(string, Map) ([]byte, error)
	Request(string, Map) ([]byte, error)
}

type requester struct {
	Config
	SafeTransport *http.Transport
	Transport     *http.Transport
}

func NewRequest(config Config) Request {
	r := &requester{
		Config: config,
	}
	r.BuildTransport()
	return r
}

func (r *requester) BuildTransport() {
	r.Transport = &http.Transport{
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

	cert, err := tls.LoadX509KeyPair(r.Get("cert_path"), r.Get("key_path"))
	if err != nil {
		Println(err)
		return
	}

	caFile, err := ioutil.ReadFile(r.Get("rootca_path"))
	if err != nil {
		Println(err)
		return
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caFile)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	}
	tlsConfig.BuildNameToCertificate()
	r.SafeTransport = &http.Transport{
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

func (r *requester) SafeRequest(url string, m Map) ([]byte, error) {
	return request(r.SafeTransport, url, m)
}

func (r *requester) Request(url string, m Map) ([]byte, error) {
	return request(r.Transport, url, m)
}

func request(transport *http.Transport, url string, m Map) ([]byte, error) {
	client := &http.Client{
		Transport: transport,
		//Timeout:   time.Duration((connectTimeoutMs + readTimeoutMs) * 1000000),
	}
	Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(m.ToXml()))
	if err != nil {
		Println(err)
		return []byte(nil), err
	}
	req.Header.Set("Content-Type", "text/xml")
	//req.Header.Set("User-Agent", "wxpay sdk go v1.0 ")

	resp, err := client.Do(req)
	if err != nil {
		Println(err)
		return []byte(nil), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}
