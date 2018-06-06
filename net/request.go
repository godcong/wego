package net

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/godcong/wego/log"
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
//	"github.com/godcong/wopay/tool"
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

/*RequestType RequestType */
type RequestType string

/*request types */
const (
	RequestTypeJson       RequestType = "json"
	RequestTypeQuery      RequestType = "query"
	RequestTypeXml        RequestType = "xml"
	RequestTypeFormParams RequestType = "form_params"
	RequestTypeFile       RequestType = "file"
	RequestTypeMultipart  RequestType = "multipart"
	RequestTypeString     RequestType = "string"
	RequestTypeHeaders    RequestType = "headers"
	RequestTypeCustom     RequestType = "custom"
)

type Request struct {
	requestType RequestType
	//requestData *RequestData
	httpRequest *http.Request
	custom      func(*Request) string

	error error
}

/*NewRequest NewRequest */
func NewRequest() *Request {
	r := Request{
		httpRequest: nil,
		error:       nil,
		//requestData: &RequestData{
		//	Query:  "",
		//	Body:   nil,
		//	Method: "",
		//	Header: http.Header{},
		//},
	}
	return &r
}

/*SetCustomCallback SetCustomCallback*/
func (r *Request) SetCustomCallback(f func(*Request) string) *Request {
	r.custom = f
	return r
}

func (r *Request) GetRequestType() RequestType {
	return r.requestType
}

func (r *Request) Error() error {
	return r.error
}

func NewRequestData() *RequestData {
	data := &RequestData{
		Query:  "",
		Body:   nil,
		Method: "",
		Header: http.Header{},
	}
	//data.Header = cloneHeader(r.requestData.Header)
	return data
}
func ReqType(reqType string) RequestType {
	log.Debug("reqType", reqType)
	switch reqType {
	case ContentTypeJson:
		return RequestTypeJson
	case ContentTypeHtml:
		//return REQUEST_TYPE_HTML
	case ContentTypeXml, ContentTypeXml2:
		return RequestTypeXml
	case ContentTypePlain:
	case ContentTypePostForm:
	case ContentTypeMultipartPostForm:
	case ContentTypeProtoBuf:
	case ContentTypeMsgPack:
	case ContentTypeMsgPack2:
	}
	return RequestTypeJson
}

func (r *Request) HTTPRequest() *http.Request {
	return r.httpRequest
}

func PerformRequest(url string, method string, data *RequestData) *Request {
	request := NewRequest()
	var req *http.Request
	var err error
	data = dataProcess(request, method, data)
	url = parseQuery(url, data.Query)
	log.Debug("PerformRequest|url", url)
	log.Debug("PerformRequest|data", data.Header, data.Method, data.Query)
	// b, e := ioutil.ReadAll(data.Body)
	// Debug("PerformRequest|data.Body", b, e)

	req, err = http.NewRequest(data.Method, url, data.Body)
	if err != nil {
		request.error = err
	}

	for k, v := range data.Header {
		req.Header[k] = v
	}

	log.Debug(req.Header)
	log.Debug("PerformRequest|req", req)
	request.httpRequest = req
	request.requestType = ReqType(filterContent(data.Header.Get("Content-Type")))
	return request
}

func dataProcess(r *Request, method string, src *RequestData) *RequestData {
	if src == nil {
		src = NewRequestData()
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
