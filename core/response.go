package core

import (
	"encoding/json"
	"encoding/xml"
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
	m := make(util.Map)
	err := json.Unmarshal(r.Data, &m)
	if err != nil {
		return nil
	}
	return m
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
	m := make(util.Map)
	err := xml.Unmarshal(r.Data, &m)
	if err != nil {
		return nil
	}
	return m
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
