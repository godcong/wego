package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"io"
	"io/ioutil"
	"net/http"
)

/*Response Response */
type Response interface {
	ToMap() util.Map
	Bytes() []byte
	Error() error
}

type responseMap struct {
	Data util.Map
}

type responseJSON struct {
	Data []byte
}

type responseXML struct {
	Data []byte
}

type responseError struct {
	Err error
}

//ToMap response to map
func (r *responseMap) ToMap() util.Map {
	return r.Data
}

//Bytes response to bytes
func (r *responseMap) Bytes() []byte {
	return r.Data.ToJSON()
}

//Error response error
func (r *responseMap) Error() error {
	return nil
}

//ToMap response to map
func (r *responseJSON) ToMap() util.Map {
	return util.JSONToMap(r.Data)
}

//Bytes response to bytes
func (r *responseJSON) Bytes() []byte {
	return r.Data
}

//Error response error
func (*responseJSON) Error() error {
	return nil
}

//ToMap response to map
func (r *responseError) ToMap() util.Map {
	return nil
}

//Bytes response to bytes
func (r *responseError) Bytes() []byte {
	return nil
}

//Error response error
func (r *responseError) Error() error {
	return r.Err
}

//ToMap response to map
func (r *responseXML) ToMap() util.Map {
	return util.XMLToMap(r.Data)
}

//Bytes response to bytes
func (r *responseXML) Bytes() []byte {
	return r.Data
}

//Error response error
func (r *responseXML) Error() error {
	return nil
}

func filterContent(content string) string {
	log.Debug("content", content)
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

/*RespType RespType */
//func RespType(reqType RequestType) ResponseType {
//	log.Debug("reqType", reqType)
//	switch reqType {
//	//case CONTENT_TYPE_JSON:
//	//	return RESPONSE_TYPE_JSON
//	//case CONTENT_TYPE_HTML:
//	//	return RESPONSE_TYPE_HTML
//	//case CONTENT_TYPE_XML, CONTENT_TYPE_XML2:
//	//	return RESPONSE_TYPE_XML
//	//case CONTENT_TYPE_Plain:
//	//case CONTENT_TYPE_POSTForm:
//	//case CONTENT_TYPE_MultipartPOSTForm:
//	//case CONTENT_TYPE_PROTOBUF:
//	//case CONTENT_TYPE_MSGPACK:
//	//case CONTENT_TYPE_MSGPACK2:
//	case RequestTypeJSON:
//		return ResponseTypeJSON
//	case RequestTypeQuery:
//	case RequestTypeXML:
//		return ResponseTypeXML
//	case RequestTypeFormParams:
//	case RequestTypeFile:
//	case RequestTypeMultipart:
//	case RequestTypeString:
//	case RequestTypeHeaders:
//	case RequestTypeCustom:
//	}
//	return ResponseTypeJSON
//}

/*ParseBody get response data */
func ParseBody(r *http.Response) ([]byte, error) {
	return ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
}

/*BodyToMap transfer response body to map data */
func BodyToMap(b []byte, d string) util.Map {
	if d == DataTypeXML {
		return util.XMLToMap(b)
	} else if d == DataTypeJSON {
		return util.JSONToMap(b)
	} else {

	}
	return nil
}

///*ToXML transfer response data to xml */
//func (r *Response) ToXML() []byte {
//	if r.dataType == DataTypeXML {
//		return r.data
//	} else if r.dataType == DataTypeJSON {
//		return []byte(util.XMLToMap(r.data).ToXML())
//	}
//	return nil
//}
//
//func (r *Response) ToMap() util.Map {
//	if r.dataType == DataTypeJSON {
//		return util.JSONToMap(r.data)
//	} else if r.dataType == DataTypeXML {
//		return util.XMLToMap(r.data)
//	} else {
//
//	}
//
//	return nil
//}
//
///*ToJSON transfer response data to json */
//func (r *Response) ToJSON() []byte {
//	if r.dataType == DataTypeJSON {
//		return r.data
//	} else if
//	return r.responseMap.ToJSON()
//}

/*ToBytes transfer response data to bytes */
//func (r *Response) ToBytes() []byte {
//	return r.responseData
//}

/*ToString transfer response data to string */
//func (r *Response) ToString() string {
//	return string(r.responseData)
//}

/*ToFile save response data to file with path */
//func (r *Response) ToFile(path string) {
//	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
//	if e != nil {
//		log.Debug("Response|ToFile", e)
//		return
//	}
//	file.Write(r.ToBytes())
//}

/*CheckError check wechat result error */
//func (r *Response) CheckError() error {
//	if r.error != nil {
//		return r.error
//	}
//	m := r.ToMap()
//	if m.GetNumber("errcode") != 0 {
//		r.error = errors.New(m.GetString("errmsg"))
//	}
//	return r.error
//}

/*ErrorResponse return response with error */
//func ErrorResponse(err error) *Response {
//	log.Debug("ErrorResponse|err", err)
//	return &Response{
//		error: err,
//	}
//}

//func (r ResponseType) String() string {
//	return string(r)
//}
