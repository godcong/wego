package net

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

type Response struct {
	responseType ResponseType
	responseData []byte
	responseMap  util.Map
	//response     *http.Response
	error error
}

type ResponseType string

const (
	RESPONSE_TYPE_JSON ResponseType = "json"
	RESPONSE_TYPE_XML               = "xml"
	RESPONSE_TYPE_HTML              = "html"
	//RESPONSE_TYPE_ARRAY               = "array"
	//RESPONSE_TYPE_STRUCT              = "struct"
	//RESPONSE_TYPE_MAP                 = "map"
	//RESPONSE_TYPE_RAW                 = "raw"
)

func filterContent(content string) string {
	log.Debug("content", content)
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func RespType(reqType RequestType) ResponseType {
	log.Debug("respType", respType)
	switch reqType {
	//case CONTENT_TYPE_JSON:
	//	return RESPONSE_TYPE_JSON
	//case CONTENT_TYPE_HTML:
	//	return RESPONSE_TYPE_HTML
	//case CONTENT_TYPE_XML, CONTENT_TYPE_XML2:
	//	return RESPONSE_TYPE_XML
	//case CONTENT_TYPE_Plain:
	//case CONTENT_TYPE_POSTForm:
	//case CONTENT_TYPE_MultipartPOSTForm:
	//case CONTENT_TYPE_PROTOBUF:
	//case CONTENT_TYPE_MSGPACK:
	//case CONTENT_TYPE_MSGPACK2:
	case REQUEST_TYPE_JSON:
		return RESPONSE_TYPE_JSON
	case REQUEST_TYPE_QUERY:
	case REQUEST_TYPE_XML:
		return RESPONSE_TYPE_XML
	case REQUEST_TYPE_FORM_PARAMS:
	case REQUEST_TYPE_FILE:
	case REQUEST_TYPE_MULTIPART:
	case REQUEST_TYPE_STRING:
	case REQUEST_TYPE_HEADERS:
	case REQUEST_TYPE_CUSTOM:
	}
	return RESPONSE_TYPE_JSON
}

func ParseResponse(typ RequestType, r *http.Response) *Response {
	var resp Response
	resp.responseData, resp.error = ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	//resp.responseType = RespType(filterContent(r.Header.Get("Content-Type")))
	resp.responseType = RespType(typ)

	log.Debug("ClientDo|response", resp.responseType, resp.error, resp.responseMap)
	log.Debug("ClientDo|response|data", len(resp.responseData))
	return &resp
}

func (r *Response) Error() error {
	return r.error
}

func (r *Response) ToMap() util.Map {
	if r.responseType == RESPONSE_TYPE_XML {
		r.responseMap = util.XmlToMap(r.responseData)
	}
	if r.responseType == RESPONSE_TYPE_JSON {
		r.responseMap = util.JsonToMap(r.responseData)
	}

	return r.responseMap
}

func (r *Response) ToXml() string {
	if r.responseType == RESPONSE_TYPE_XML {
		return string(r.responseData)
	}
	return r.responseMap.ToXml()
}

func (r *Response) ToJson() []byte {
	if r.responseType == RESPONSE_TYPE_JSON {
		return r.responseData
	}
	return r.responseMap.ToJson()
}

func (r *Response) ToBytes() []byte {
	return r.responseData
}

func (r *Response) ToString() string {
	return string(r.responseData)
}

func (r *Response) ToFile(path string) {
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if e != nil {
		log.Debug("Response|ToFile", e)
		return
	}
	file.Write(r.ToBytes())
}

func ErrorResponse(err error) *Response {
	log.Debug("ErrorResponse|err", err)
	return &Response{
		error: err,
	}
}

func (r ResponseType) String() string {
	return string(r)
}
