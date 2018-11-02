package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
	Data []byte
	Err  error
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
	return r.Data
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

//Err return if response has error
func Err(data []byte, err error) Response {
	return &responseError{
		Data: data,
		Err:  err,
	}
}

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

// SaveTo ...
func SaveTo(response Response, path string) error {
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if e != nil {
		log.Debug("Response|ToFile", e)
		return e
	}

	_, e = file.Write(response.Bytes())
	if e != nil {
		return e
	}
	return nil
}

// SaveEncodingTo ...
func SaveEncodingTo(response Response, path string, t transform.Transformer) error {
	file, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if e != nil {
		log.Debug("Response|ToFile", e)
		return e
	}

	writer := transform.NewWriter(file, t)
	_, e = writer.Write(response.Bytes())
	if e != nil {
		return e
	}
	return nil
}

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
