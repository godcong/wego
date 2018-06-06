package net

import (
	"errors"
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
	log.Debug("reqType", reqType)
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
	case RequestTypeJson:
		return RESPONSE_TYPE_JSON
	case RequestTypeQuery:
	case RequestTypeXml:
		return RESPONSE_TYPE_XML
	case RequestTypeFormParams:
	case RequestTypeFile:
	case RequestTypeMultipart:
	case RequestTypeString:
	case RequestTypeHeaders:
	case RequestTypeCustom:
	}
	return RESPONSE_TYPE_JSON
}

func ParseResponse(typ RequestType, r *http.Response) *Response {
	var resp Response
	resp.responseData, resp.error = ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))

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
		r.responseMap = util.XMLToMap(r.responseData)
	}
	if r.responseType == RESPONSE_TYPE_JSON {
		r.responseMap = util.JSONToMap(r.responseData)
	}

	return r.responseMap
}

/*ToXML transfer response data to xml*/
func (r *Response) ToXML() string {
	if r.responseType == RESPONSE_TYPE_XML {
		return string(r.responseData)
	}
	return r.responseMap.ToXML()
}

/*ToJSON transfer response data to json*/
func (r *Response) ToJSON() []byte {
	if r.responseType == RESPONSE_TYPE_JSON {
		return r.responseData
	}
	return r.responseMap.ToJSON()
}

/*ToBytes transfer response data to bytes*/
func (r *Response) ToBytes() []byte {
	return r.responseData
}

/*ToString transfer response data to string*/
func (r *Response) ToString() string {
	return string(r.responseData)
}

/*ToFile save response data to file with path */
func (r *Response) ToFile(path string) {
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if e != nil {
		log.Debug("Response|ToFile", e)
		return
	}
	file.Write(r.ToBytes())
}

/*CheckError check wechat result error */
func (r *Response) CheckError() error {
	if r.error != nil {
		return r.error
	}
	m := r.ToMap()
	if m.GetNumber("errcode") != 0 {
		r.error = errors.New(m.GetString("errmsg"))
	}
	return r.error
}

/*ErrorResponse return response with error */
func ErrorResponse(err error) *Response {
	log.Debug("ErrorResponse|err", err)
	return &Response{
		error: err,
	}
}

func (r ResponseType) String() string {
	return string(r)
}
