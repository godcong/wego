package core

import (
	"bytes"
	"log"
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

type RequestType string

const (
	REQUEST_TYPE_JSON        RequestType = "json"
	REQUEST_TYPE_QUERY                   = "query"
	REQUEST_TYPE_XML                     = "xml"
	REQUEST_TYPE_FORM_PARAMS             = "form_params"
	REQUEST_TYPE_FILE                    = "file"
	REQUEST_TYPE_MULTIPART               = "multipart"
	REQUEST_TYPE_STRING                  = "string"
	REQUEST_TYPE_HEADERS                 = "headers"
)

type Request struct {
	request *http.Request
	error   error
	options Map
}

var DefaultRequest = NewRequest()

func init() {
	log.Println("init")
}

func NewRequest() *Request {
	log.Println("request")
	r := Request{}
	r.options = Map{}
	return &r
}

func (r *Request) SetOptions(op Map) *Request {
	r.options = op
	return r
}

func (r *Request) GetOptions() Map {
	return r.options
}

func (r *Request) Error() error {
	return r.error
}

func (r *Request) PerformRequest(url string, method string, ops Map) *Request {
	var req *http.Request
	var err error
	method = strings.ToUpper(method)
	op := optionMerge(r, ops)

	reqData := ""
	if _, b := ops["json"]; b {
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

	header, b := op["headers"].(Map)
	if method == "POST" {
		if b {
			for k, v := range header {
				req.Header.Set(k, v.(string))
			}
		}
	}
	r.request = req
	r.error = err
	return r
}

func optionMerge(r *Request, options Map) Map {
	base := r.GetOptions()
	if options == nil {
		return base
	}

	for key, value := range options {
		base[key] = value
	}
	if v, b := options["json"]; b {
		base["headers"] = Map{"Content-Type": "application/json"}
		base["body"] = v
	}
	return base
}

func pickUpBody(typ RequestType, body interface{}) string {
	switch v := body.(type) {
	case Map:
		if typ == "json" {
			return string(v.ToJson())
		} else {

		}

	}
	return ""
}
