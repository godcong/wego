package wego

import (
	"bytes"
	"encoding/xml"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

// BodyType ...
type BodyType string

// BodyTypeJSON ...
const BodyTypeJSON BodyType = "json"

// BodyTypeXML ...
const BodyTypeXML BodyType = "xml"

// BodyTypeNone ...
const BodyTypeNone BodyType = "none"

// BodyTypeMultipart ...
const BodyTypeMultipart BodyType = "multipart"

// BodyTypeForm ...
const BodyTypeForm BodyType = "form"

// RequestBody ...
type RequestBody struct {
	BodyType       BodyType
	BodyInstance   interface{}
	RequestBuilder RequestBuilderFunc
}

// RequestBuilderFunc ...
type RequestBuilderFunc func(method, url string, i interface{}) (*http.Request, error)

//TODO
var buildMultipart = buildNothing
var buildForm = buildNothing

var builder = map[BodyType]RequestBuilderFunc{
	BodyTypeXML:       buildXML,
	BodyTypeJSON:      buildJSON,
	BodyTypeForm:      buildForm,
	BodyTypeMultipart: buildMultipart,
	BodyTypeNone:      buildNothing,
}

func buildXML(method, url string, i interface{}) (*http.Request, error) {
	request, e := http.NewRequest(method, url, xmlReader(i))
	if e != nil {
		return nil, e
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	return request, nil
}

// xmlReader ...
func xmlReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("toXMLReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("toXMLReader|[]byte", v)
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("toXMLReader|util.Map", string(v.ToXML()))
		reader = bytes.NewReader(v.ToXML())
	default:
		log.Debug("toXMLReader|default", v)
		if v0, e := xml.Marshal(v); e == nil {
			log.Debug("toXMLReader|v0", v0, e)
			reader = bytes.NewReader(v0)
		}
	}
	return reader
}

func buildJSON(method, url string, i interface{}) (*http.Request, error) {
	request, e := http.NewRequest(method, url, jsonReader(i))
	if e != nil {
		return nil, e
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	return request, nil
}

func buildNothing(method, url string, i interface{}) (*http.Request, error) {
	request, e := http.NewRequest(method, url, nil)
	if e != nil {
		return nil, e
	}
	return request, nil
}

// jsonReader ...
func jsonReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("jsonReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("jsonReader|[]byte", string(v))
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("jsonReader|util.Map", v.String())
		reader = bytes.NewReader(v.ToJSON())
	default:
		log.Debug("jsonReader|default", v)
		if v0, e := jsoniter.Marshal(v); e == nil {
			reader = bytes.NewReader(v0)
		}
	}
	return reader
}

// buildBody ...
func buildBody(v interface{}, tp BodyType) *RequestBody {
	build, b := builder[tp]
	if !b {
		build = buildNothing
	}
	return &RequestBody{
		BodyType:       tp,
		RequestBuilder: build,
		BodyInstance:   v,
	}
}

/*Requester Requester */
type Requester interface {
	ToMap() util.Map
	Bytes() []byte
	Error() error
	Unmarshal(v interface{}) error
	Result() (util.Map, error)
}

// Request ...
type Request struct {
	bytes []byte
	err   error
}

// xmlResponse ...
type xmlRequest struct {
	Request
	data util.Map
}

// XMLRequest ...
func XMLRequest(bytes []byte) Requester {
	return &xmlRequest{
		Request: Request{
			bytes: bytes,
		},
	}
}

// ToMap ...
func (r *xmlRequest) ToMap() util.Map {
	maps, e := r.Result()
	if e != nil {
		return nil
	}
	return maps
}

// Unmarshal ...
func (r *xmlRequest) Unmarshal(v interface{}) error {
	return xml.Unmarshal(r.bytes, v)
}

// Result ...
func (r *xmlRequest) Result() (util.Map, error) {
	if r.data != nil {
		return r.data, nil
	}
	r.data = make(util.Map)
	e := r.Unmarshal(&r.data)
	return r.data, e
}

// jsonResponse ...
type jsonRequest struct {
	Request
	data util.Map
}

// JSONRequest ...
func JSONRequest(bytes []byte) Requester {
	return &jsonRequest{
		Request: Request{
			bytes: bytes,
		},
	}
}

// ToMap ...
func (r *jsonRequest) ToMap() util.Map {
	maps, e := r.Result()
	if e != nil {
		return nil
	}
	return maps
}

// Unmarshal ...
func (r *jsonRequest) Unmarshal(v interface{}) error {
	return jsoniter.Unmarshal(r.bytes, v)
}

// Result ...
func (r *jsonRequest) Result() (util.Map, error) {
	r.data = make(util.Map)
	e := r.Unmarshal(&r.data)
	return r.data, e
}

// ToMap ...
func (r *Request) ToMap() util.Map {
	return nil
}

// Unmarshal ...
func (r *Request) Unmarshal(v interface{}) error {
	return r.err
}

// Result ...
func (r *Request) Result() (util.Map, error) {
	return nil, r.err
}

// ErrRequest ...
func ErrRequest(err error) Requester {
	return &Request{
		bytes: nil,
		err:   err,
	}
}

// Bytes ...
func (r *Request) Bytes() []byte {
	return r.bytes
}

// Error ...
func (r *Request) Error() error {
	return r.err
}

// ReadRequest ...
func ReadRequest(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(io.LimitReader(r.Body, math.MaxUint32))
}

// buildRequester ...
func buildRequester(req *http.Request) Requester {
	ct := req.Header.Get("Content-Type")
	body, err := ReadRequest(req)
	if err != nil {
		log.Error(body, err)
		return ErrResponse(err)
	}

	log.Debug("request:", string(body[:128]), len(body)) //max 128 char
	if strings.Index(ct, "xml") != -1 ||
		bytes.Index(body, []byte("<xml")) != -1 {
		return XMLResponse(body)
	}
	return JSONResponse(body)
	//return ErrResponse(xerrors.New("error with code " + req.Status))
}
