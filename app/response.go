package app

import (
	"bytes"
	"encoding/xml"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"golang.org/x/exp/xerrors"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

/*Responder Responder */
type Responder interface {
	ToMap() util.Map
	Bytes() []byte
	Error() error
	Unmarshal(v interface{}) error
	Result() (util.Map, error)
}

// Response ...
type Response struct {
	bytes []byte
	err   error
}

// xmlResponse ...
type xmlResponse struct {
	Response
	maps util.Map
}

// XMLResponse ...
func XMLResponse(bytes []byte) Responder {
	return &xmlResponse{
		Response: Response{
			bytes: bytes,
		},
	}
}

// ToMap ...
func (r *xmlResponse) ToMap() util.Map {
	maps, e := r.Result()
	if e != nil {
		return nil
	}
	return maps
}

// Unmarshal ...
func (r *xmlResponse) Unmarshal(v interface{}) error {
	return xml.Unmarshal(r.bytes, v)
}

// Result ...
func (r *xmlResponse) Result() (util.Map, error) {
	if r.maps != nil {
		return r.maps, nil
	}
	r.maps = make(util.Map)
	e := r.Unmarshal(&r.maps)
	return r.maps, e
}

// jsonResponse ...
type jsonResponse struct {
	Response
	maps util.Map
}

// JSONResponse ...
func JSONResponse(bytes []byte) Responder {
	return &jsonResponse{
		Response: Response{
			bytes: bytes,
		},
	}
}

// ToMap ...
func (r *jsonResponse) ToMap() util.Map {
	maps, e := r.Result()
	if e != nil {
		return nil
	}
	return maps
}

// Unmarshal ...
func (r *jsonResponse) Unmarshal(v interface{}) error {
	return jsoniter.Unmarshal(r.bytes, v)
}

// Result ...
func (r *jsonResponse) Result() (util.Map, error) {
	r.maps = make(util.Map)
	e := r.Unmarshal(&r.maps)
	return r.maps, e
}

// ToMap ...
func (r *Response) ToMap() util.Map {
	return nil
}

// Unmarshal ...
func (r *Response) Unmarshal(v interface{}) error {
	return r.err
}

// Result ...
func (r *Response) Result() (util.Map, error) {
	return nil, r.err
}

// ErrResponse ...
func ErrResponse(err error) Responder {
	return &Response{
		bytes: nil,
		err:   err,
	}
}

// Bytes ...
func (r *Response) Bytes() []byte {
	return r.bytes
}

// Error ...
func (r *Response) Error() error {
	return r.err
}

/*ReadResponse get response data */
func ReadResponse(r *http.Response) ([]byte, error) {
	return ioutil.ReadAll(io.LimitReader(r.Body, math.MaxUint32))
}

// buildResponder ...
func buildResponder(resp *http.Response) Responder {
	ct := resp.Header.Get("Content-Type")
	body, err := ReadResponse(resp)
	if err != nil {
		log.Error(body, err)
		return ErrResponse(err)
	}

	log.Debug("response:", string(body[:128]), len(body)) //max 128 char
	if resp.StatusCode == 200 {
		if strings.Index(ct, "xml") != -1 ||
			bytes.Index(body, []byte("<xml")) != -1 {
			return XMLResponse(body)
		}
		return JSONResponse(body)
	}
	log.Error("error with " + resp.Status)
	return ErrResponse(xerrors.New("error with code " + resp.Status))
}
