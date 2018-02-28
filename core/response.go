package core

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	response     *http.Response
	responseData []byte
	responseType ResponseType
	error        error
}

type ResponseType string

const (
	RESPONSE_TYPE_JSON   ResponseType = "json"
	RESPONSE_TYPE_XML                 = "xml"
	RESPONSE_TYPE_ARRAY               = "array"
	RESPONSE_TYPE_STRUCT              = "struct"
	RESPONSE_TYPE_MAP                 = "map"
	RESPONSE_TYPE_RAW                 = "raw"
)

func (r *Response) Error() error {
	return r.error
}

func (r *Response) ToMap() Map {
	if r.responseType == RESPONSE_TYPE_XML {
		return XmlToMap(r.responseData)
	}
	if r.responseType == RESPONSE_TYPE_JSON {
		return JsonToMap(r.responseData)
	}
	return Map{}
}

func (r *Response) ToBytes() []byte {
	return r.responseData
}

func ParseClient(client *http.Client, request *Request) *Response {
	response := &Response{}

	response.response, response.error = client.Do(request.HttpRequest())
	if response.error != nil {
		return response
	}

	response.responseData, response.error = ioutil.ReadAll(response.response.Body)
	if request.GetRequestType() == REQUEST_TYPE_JSON {
		response.responseType = RESPONSE_TYPE_JSON
	} else {
		response.responseType = RESPONSE_TYPE_XML
	}
	return response
}

func ErrorResponse(err error) *Response {
	return &Response{
		error: err,
	}
}
