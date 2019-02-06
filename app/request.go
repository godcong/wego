package app

import (
	"bytes"
	"encoding/xml"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"io"
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
