package wego

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
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

type RequestInterface interface {
}

type Request struct {
	client  Client
	app     Application
	options map[string]Map
	//Transport     *http.Transport
}

func NewRequest() *Request {
	r := Request{}
	r.options = map[string]Map{
		"header": {
			"Content-Type": "text/xml",
		},
	}
	return &r
}

func (r *Request) Options() map[string]Map {
	return r.options
}

func (r *Request) PerformRequest(transport *http.Transport, url string, method string, options map[string]Map) ([]byte, error) {
	op := optionCheck(r, options)
	client := &http.Client{
		Transport: transport,
		//Timeout:   time.Duration((connectTimeoutMs + readTimeoutMs) * 1000000),
	}
	body, b := op["body"]
	if !b {
		body = Map{}
	}
	reqData := body.ToXml()
	if _, b := options["json"]; b {
		reqData = string(body.ToJson())
	}
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBufferString(reqData))
	if err != nil {
		Println(err)
		return []byte(nil), err
	}
	header, b := op["header"]
	if b {
		for k, v := range header {
			//req.Header.Set("Content-Type", "text/xml")
			req.Header.Set(k, v)
		}
	}

	//req.Header.Set("User-Agent", "wxpay sdk go v1.0 ")
	resp, err := client.Do(req)
	if err != nil {
		Println(err)
		return []byte(nil), err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func optionCheck(r *Request, options map[string]Map) map[string]Map {
	base := r.Options()
	if options == nil {
		return base
	}
	for key, value := range options {
		base[key] = value
	}

	if v, b := base["json"]; b {
		base["headers"] = Map{"Content-Type": "application/json"}
		//b, _ := json.Marshal(v)
		base["body"] = v
	}
	delete(base, "json")

	return base
}

func MakeOption(options map[string]Map) map[string]Map {
	if options == nil {
		return make(map[string]Map)
	}
	return options
}
