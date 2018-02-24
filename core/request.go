package core

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
	options Map
	//Transport     *http.Transport
}

func NewRequest() *Request {
	r := Request{}
	r.options = Map{
		"headers": Map{
			"Content-Type": "text/xml",
		},
	}
	return &r
}

func (r *Request) Options() Map {
	return r.options
}

func (r *Request) PerformRequest(transport *http.Transport, url string, method string, options Map) ([]byte, error) {
	var req *http.Request
	var err error
	op := optionCheck(r, options)
	method = strings.ToUpper(method)
	client := &http.Client{
		Transport: transport,
		//Timeout:   time.Duration(50000),
	}
	reqData := ""
	if _, b := options["json"]; b {
		switch v := op["body"].(type) {
		case string:
			reqData = v
		case Map:
			reqData = string(v.ToJson())
		}
	} else {
		body := op["body"].(Map)
		reqData = body.ToXml()
	}
	if method == "GET" {
		Info(method, url)
		req, err = http.NewRequest(method, url, nil)
	} else {
		Info(method, url)
		req, err = http.NewRequest(method, url, bytes.NewBufferString(reqData))
	}

	if err != nil {
		Error(err)
		return []byte(nil), err
	}
	header, b := op["headers"].(Map)
	if method == "POST" {
		if b {
			for k, v := range header {
				req.Header.Set(k, v.(string))
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		Error(err)
		return []byte(nil), err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func optionCheck(r *Request, options Map) Map {
	base := r.Options()
	if options == nil {
		return base
	}

	for key, value := range options {
		base[key] = value
	}
	if v, b := options["json"]; b {
		base["headers"] = Map{"Content-Type": "application/json; encoding=utf-8"}
		base["body"] = v
	}

	return base
}

func MakeOption(options Map) Map {
	if options == nil {
		return make(Map)
	}
	return options
}
