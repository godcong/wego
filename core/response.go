package core

import (
	"os"
)

type Response struct {
	responseType ResponseType
	responseData []byte
	responseMap  Map
	//response     *http.Response
	error error
}

type ResponseType string

const (
	RESPONSE_TYPE_JSON ResponseType = "json"
	RESPONSE_TYPE_XML               = "xml"
	//RESPONSE_TYPE_ARRAY               = "array"
	//RESPONSE_TYPE_STRUCT              = "struct"
	//RESPONSE_TYPE_MAP                 = "map"
	//RESPONSE_TYPE_RAW                 = "raw"
)

func (r *Response) Error() error {
	return r.error
}

func (r *Response) ToMap() Map {
	if r.responseType == RESPONSE_TYPE_XML {
		r.responseMap = XmlToMap(r.responseData)
	}
	if r.responseType == RESPONSE_TYPE_JSON {
		r.responseMap = JsonToMap(r.responseData)
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
	// if s := len(r.responseData); s > 2048 {
	// 	return strconv.Itoa(s)
	// }
	return string(r.responseData)
}

func (r *Response) ToFile(path string) {
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if e != nil {
		Debug("Response|ToFile", e)
		return
	}
	file.Write(r.ToBytes())
}

func ErrorResponse(err error) *Response {
	Debug("ErrorResponse|err", err)
	return &Response{
		error: err,
	}
}

func (r ResponseType) String() string {
	return string(r)
}
