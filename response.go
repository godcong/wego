package wego

import (
	"bytes"
	"encoding/xml"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"golang.org/x/text/transform"
	"golang.org/x/xerrors"
	"net/http"
	"os"
	"strings"
)

/*Responder Responder */
type Responder interface {
	Type() BodyType
	BodyReader
}

// Response ...
type Response struct {
	bytes []byte
	err   error
}

// Type ...
func (r *Response) Type() BodyType {
	return BodyTypeNone
}

// xmlResponse ...
type xmlResponse struct {
	Response
	data util.Map
}

// Type ...
func (r *xmlResponse) Type() BodyType {
	return BodyTypeXML
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

// Type ...
func (r *jsonResponse) Type() BodyType {
	return BodyTypeJSON
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

// ResponseWriter ...
func ResponseWriter(w http.ResponseWriter, responder Responder) error {
	w.WriteHeader(http.StatusOK)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		switch responder.Type() {
		case BodyTypeXML:
			header["Content-Type"] = []string{"application/xml; charset=utf-8"}
		case BodyTypeJSON:
			header["Content-Type"] = []string{"application/json; charset=utf-8"}
		}
	}
	b := responder.Bytes()
	if b == nil {
		return nil
	}

	_, err := w.Write(b)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// BuildResponder ...
func BuildResponder(resp *http.Response) Responder {
	ct := resp.Header.Get("Content-Type")
	body, err := readBody(resp.Body)
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
