package core

import (
	"bytes"
	"io"
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
	REQUEST_TYPE_QUERY       RequestType = "query"
	REQUEST_TYPE_XML         RequestType = "xml"
	REQUEST_TYPE_FORM_PARAMS RequestType = "form_params"
	REQUEST_TYPE_FILE        RequestType = "file"
	REQUEST_TYPE_MULTIPART   RequestType = "multipart"
	REQUEST_TYPE_STRING      RequestType = "string"
	REQUEST_TYPE_HEADERS     RequestType = "headers"
)

type Request struct {
	request     *http.Request
	error       error
	requestData RequestData
}

type RequestData struct {
	Query       string
	Body        string
	Method      string
	requestType RequestType
	Header      http.Header
}

var DefaultRequest = NewRequest()

func init() {
	log.Println("init")
}

func NewRequest() *Request {
	log.Println("request")
	r := Request{
		request: nil,
		error:   nil,
		requestData: RequestData{
			Query:  "",
			Body:   "",
			Method: "",
			Header: http.Header{},
		},
	}
	return &r
}

func (r *Request) Error() error {
	return r.error
}

func (r *Request) HttpRequest() *http.Request {
	return r.request
}

func (r *Request) PerformRequest(url string, method string, ops Map) *Request {
	var req *http.Request
	var err error
	data := optionProcess(r, method, ops)

	req, err = http.NewRequest(data.Method, parseQuery(url, data.Query), parseBody(data.Body))
	Info(data)
	if err != nil {
		r.error = err
	}
	req.Header = data.Header
	r.request = req
	return r
}

func optionProcess(r *Request, method string, src Map) RequestData {
	base := r.requestData
	if src == nil {
		return base
	}

	for key, value := range src {

		switch (RequestType)(key) {
		case REQUEST_TYPE_JSON:
			base.Body = processJson(value)
			base.Header.Set("Content-Type", "application/json")
			base.requestType = REQUEST_TYPE_JSON
		case REQUEST_TYPE_QUERY:
			base.Query = processQuery(value)
		case REQUEST_TYPE_XML:
			base.requestType = REQUEST_TYPE_XML
			base.Body = processXml(value)
		case REQUEST_TYPE_FORM_PARAMS:
		case REQUEST_TYPE_FILE:
		case REQUEST_TYPE_MULTIPART:
		case REQUEST_TYPE_STRING:
		case REQUEST_TYPE_HEADERS:
		}
	}

	if UseUTF8() {
		base.Header.Add("Content-Type", "charset=utf-8")
	}

	base.Method = strings.ToUpper(method)

	return base
}
func processXml(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case Map:
		return v.ToXml()
	}
	return ""
}

func processJson(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case Map:
		log.Println("map", v)
		return string(v.ToJson())
	}
	return ""
}

func processQuery(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case Map:
		return v.UrlEncode()
	}
	return ""
}

//REQUEST_TYPE_JSON:        nil,
//REQUEST_TYPE_QUERY:       nil,
//REQUEST_TYPE_XML:         nil,
//REQUEST_TYPE_FORM_PARAMS: nil,
//REQUEST_TYPE_FILE:        nil,
//REQUEST_TYPE_MULTIPART:   nil,
//REQUEST_TYPE_STRING:      nil,
//REQUEST_TYPE_HEADERS:     nil,
func parseQuery(url, query string) string {
	if query == "" {
		return url
	}
	return url + "?" + query
}

func parseBody(body string) io.Reader {
	if body == "" {
		return nil
	}
	return bytes.NewBufferString(body)
}

func UseUTF8() bool {
	return true
}

func (r RequestType) String() string {
	return string(r)
}
