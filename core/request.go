package core

import (
	"bytes"
	"io"
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
	REQUEST_TYPE_CUSTOM      RequestType = "custom"
)

type Request struct {
	requestType RequestType
	requestData *RequestData
	request     *http.Request
	custom      func(*Request) string

	error error
}

var DefaultRequest = NewRequest()

func NewRequest() *Request {
	r := Request{
		request: nil,
		error:   nil,
		requestData: &RequestData{
			Query:  "",
			Body:   nil,
			Method: "",
			Header: http.Header{},
		},
	}
	return &r
}

func (r *Request) SetCustomCallback(f func(*Request) string) *Request {
	r.custom = f
	return r
}

//
//func (r *Request) GetRequestType() RequestType {
//	return r.requestData.requestType
//}

func (r *Request) Error() error {
	return r.error
}

func (r *Request) RequestDataCopy() *RequestData {
	data := *r.requestData
	data.Header = cloneHeader(r.requestData.Header)
	return &data
}

func (r *Request) HttpRequest() *http.Request {
	return r.request
}

func (r *Request) PerformRequest(url string, method string, data *RequestData) *Request {
	var req *http.Request
	var err error
	data = dataProcess(r, method, data)
	url = parseQuery(url, data.Query)
	Debug("PerformRequest|data", *data)
	Debug("PerformRequest|url", url)
	req, err = http.NewRequest(data.Method, url, data.Body)
	if err != nil {
		r.error = err
	}
	req.Header = data.Header
	Debug("PerformRequest|req", req)
	r.request = req
	return r
}

func dataProcess(r *Request, method string, src *RequestData) *RequestData {
	if src == nil {
		src = r.RequestDataCopy()
	}
	src.Method = strings.ToUpper(method)

	if src.Header.Get("Content-Type") == "" {
		src.Header.Set("Content-Type", "application/json")
	}

	if UseUTF8() {
		src.Header.Add("Content-Type", "charset=utf-8")
	}

	return src
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

func cloneHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}
