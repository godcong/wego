package core

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

/*Responder Responder */
type Responder interface {
	ToMap() util.Map
	Bytes() []byte
	Error() error
	Result() (util.Map, error)
}
type respMap struct {
	Data util.Map
}

// Result ...
func (r *respMap) Result() (util.Map, error) {
	return r.ToMap(), r.Error()
}

type respJSON struct {
	Data []byte
	Err  error
}

// Result ...
func (r *respJSON) Result() (util.Map, error) {
	m := make(util.Map)
	err := json.Unmarshal(r.Data, &m)
	if err != nil {
		r.Err = err
		return nil, err
	}
	return m, err
}

type respXML struct {
	Data []byte
	Err  error
}

// Result ...
func (r *respXML) Result() (util.Map, error) {
	m := make(util.Map)
	err := xml.Unmarshal(r.Data, &m)
	if err != nil {
		r.Err = err
		return nil, err
	}
	return m, nil
}

type respData struct {
	Data []byte
	Err  error
}

// Result ...
func (r *respData) Result() (util.Map, error) {
	return nil, r.Err
}

// ToMap ...
func (r *respData) ToMap() util.Map {
	return nil
}

// Bytes ...
func (r *respData) Bytes() []byte {
	return r.Data
}

// Error ...
func (r *respData) Error() error {
	return r.Err
}

type respError struct {
	Data []byte
	Err  error
}

//ToMap response to map
func (r *respMap) ToMap() util.Map {
	return r.Data
}

//Bytes response to bytes
func (r *respMap) Bytes() []byte {
	return r.Data.ToJSON()
}

//Error response error
func (r *respMap) Error() error {
	return nil
}

//ToMap response to map
func (r *respJSON) ToMap() util.Map {
	m, _ := r.Result()
	return m
}

//Bytes response to bytes
func (r *respJSON) Bytes() []byte {
	return r.Data
}

//Error response error
func (*respJSON) Error() error {
	return nil
}

// Result ...
func (r *respError) Result() (util.Map, error) {
	return nil, r.Err
}

//ToMap response to map
func (r *respError) ToMap() util.Map {
	return nil
}

//Bytes response to bytes
func (r *respError) Bytes() []byte {
	return r.Data
}

//Error response error
func (r *respError) Error() error {
	return r.Err
}

//ToMap response to map
func (r *respXML) ToMap() util.Map {
	m, _ := r.Result()
	return m
}

//Bytes response to bytes
func (r *respXML) Bytes() []byte {
	return r.Data
}

//Error response error
func (r *respXML) Error() error {
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
func Err(data []byte, err error) Responder {
	return &respError{
		Data: data,
		Err:  err,
	}
}

/*ParseResponse get response data */
func ParseResponse(r *http.Response) ([]byte, error) {
	return ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
}

// CastToResponse ...
func CastToResponse(resp *http.Response) Responder {
	ct := resp.Header.Get("Content-Type")
	body, err := ParseResponse(resp)
	if err != nil {
		log.Error(body, err)
		return Err(body, err)
	}

	log.Debug("response:", string(body[:256])) //max 256 char
	if resp.StatusCode == 200 {
		if strings.Index(ct, "xml") != -1 ||
			bytes.Index(body, []byte("<xml")) != -1 {
			return &respXML{
				Data: body,
			}
		}
		return &respJSON{
			Data: body,
		}
	}
	log.Error(body, "error with "+resp.Status)
	return Err(body, errors.New("error with "+resp.Status))
}
