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

func ParseClient(client *http.Client, request *Request) *Response {
	response := &Response{}

	response.response, response.error = client.Do(request.HttpRequest())
	if response.error != nil {
		return response
	}

	response.responseData, response.error = ioutil.ReadAll(response.response.Body)
	return response
}

func (r *Response) Error() error {
	return r.error
}

func ErrorResponse(err error) *Response {
	return &Response{
		error: err,
	}
}
