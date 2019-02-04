package app

import (
	"bytes"
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

// Body ...
type Body struct {
	BodyType     BodyType
	BodyInstance interface{}
	Builder      RequestBuilder
}

// RequestBuilder ...
type RequestBuilder func(url, method string, i interface{}) *http.Request

var buildXML = buildNothing
var buildMultipart = buildNothing

var buildForm = buildNothing

var builder = map[BodyType]RequestBuilder{
	BodyTypeXML:       buildXML,
	BodyTypeJSON:      buildJSON,
	BodyTypeForm:      buildForm,
	BodyTypeMultipart: buildMultipart,
	BodyTypeNone:      buildNothing,
}

func buildJSON(method, url string, i interface{}) *http.Request {
	request, e := http.NewRequest(method, url, JSONReader(i))
	if e != nil {
		log.Error(e)
		return nil
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	return request
}

func buildNothing(method, url string, i interface{}) *http.Request {
	request, e := http.NewRequest(method, url, JSONReader(i))
	if e != nil {
		log.Error(e)
		return nil
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	return request
}

// JSONReader ...
func JSONReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("JSONReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("JSONReader|[]byte", string(v))
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("JSONReader|util.Map", v.String())
		reader = bytes.NewReader(v.ToJSON())
	default:
		log.Debug("JSONReader|default", v)
		if v0, e := jsoniter.Marshal(v); e == nil {
			reader = bytes.NewReader(v0)
		}
	}
	return reader
}

// NewBody ...
func NewBody(v interface{}, tps ...BodyType) *Body {
	tp := BodyTypeNone
	if tps != nil {
		tp = tps[0]
	}
	build, b := builder[tp]
	if !b {
		build = buildNothing
	}
	return &Body{
		BodyType:     tp,
		Builder:      build,
		BodyInstance: v,
	}
}
