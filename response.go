package wego

import (
	"bytes"
	"encoding/xml"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"golang.org/x/exp/xerrors"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
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
	data util.Map
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
	if r.data != nil {
		return r.data, nil
	}
	r.data = make(util.Map)
	e := r.Unmarshal(&r.data)
	return r.data, e
}

// jsonResponse ...
type jsonResponse struct {
	Response
	data util.Map
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
	r.data = make(util.Map)
	e := r.Unmarshal(&r.data)
	return r.data, e
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

// SaveTo ...
func SaveTo(response Responder, path string) error {
	var err error
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Debug("Responder|ToFile", err)
		return err
	}
	defer func() {
		err = file.Close()
	}()
	_, err = file.Write(response.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// SaveEncodingTo ...
func SaveEncodingTo(response Responder, path string, t transform.Transformer) error {
	var err error
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_SYNC, os.ModePerm)
	if err != nil {
		log.Debug("Responder|ToFile", err)
		return err
	}
	defer func() {
		err = file.Close()
	}()
	writer := transform.NewWriter(file, t)
	_, err = writer.Write(response.Bytes())
	if err != nil {
		return err
	}
	defer func() {
		err = writer.Close()
	}()
	return nil
}
