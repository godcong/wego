package core

import (
	"bytes"
	"io"
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
//type RequestType string

/*request types */
//const (
//RequestTypeJSON       RequestType = "json"
//RequestTypeQuery      RequestType = "query"
//RequestTypeXML        RequestType = "xml"
//RequestTypeFormParams RequestType = "form_params"
//RequestTypeFile       RequestType = "file"
//RequestTypeMultipart  RequestType = "multipart"
//RequestTypeString     RequestType = "string"
//RequestTypeHeaders    RequestType = "headers"
//RequestTypeCustom     RequestType = "custom"
//)

///*Request Request */
//type Request struct {
//	requestType RequestType
//	//requestData *RequestData
//	httpRequest *http.Request
//	custom      func(*Request) string
//	error error
//}

/*NewRequest NewRequest */
//func NewRequest() *Request {
//	r := Request{
//		httpRequest: nil,
//		error:       nil,
//		//requestData: &RequestData{
//		//	Query:  "",
//		//	Body:   nil,
//		//	Method: "",
//		//	Header: http.Header{},
//		//},
//	}
//	return &r
//}
//
///*SetCustomCallback SetCustomCallback*/
//func (r *Request) SetCustomCallback(f func(*Request) string) *Request {
//	r.custom = f
//	return r
//}
//
///*GetRequestType GetRequestType */
//func (r *Request) GetRequestType() RequestType {
//	return r.requestType
//}
//
//func (r *Request) Error() error {
//	return r.error
//}
//
///*NewRequestData NewRequestData */
//func NewRequestData() *RequestData {
//	data := &RequestData{
//		Query:  "",
//		Body:   nil,
//		Method: "",
//		Header: http.Header{},
//	}
//	//data.Header = cloneHeader(r.requestData.Header)
//	return data
//}

/*httpRequest get http request*/
//func (r *Request) httpRequest() *http.Request {
//	return r.httpRequest
//}

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

/*UseUTF8 return true */
func UseUTF8() bool {
	return true
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
