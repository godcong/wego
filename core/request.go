package core

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Requester ...
type Requester interface {
	Do(ctx context.Context) Responder
}

// RequestBuildFunc ...
type RequestBuildFunc func(url, method string, i interface{}) *http.Request

// ErrNilRequestBody ...
var ErrNilRequestBody = errors.New("nil request body")

func buildRequestURL(url string, p util.Map) string {
	query := buildRequestQuery(p)
	if query == "" {
		return url
	}
	return url + "?" + query
}

// RequestToMap ...
func RequestToMap(r *http.Request) (util.Map, error) {
	m := make(util.Map)
	ct := r.Header.Get("Content-Type")
	body, err := ParseRequest(r)
	if err != nil {
		log.Error(body, err)
		return nil, err
	}
	if strings.Index(ct, "xml") != -1 ||
		bytes.Index(body, []byte("<xml")) != -1 {
		err = xml.Unmarshal(body, &m)
		if err != nil {
			return nil, err
		}

	} else if strings.Index(ct, "json") != -1 {
		err = json.Unmarshal(body, &m)
		if err != nil {
			return nil, err
		}
	} else {
		//other case
		return nil, nil
	}
	return m, nil
}

// ParseRequest ...
func ParseRequest(r *http.Request) ([]byte, error) {
	if r.Body != nil {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return body, nil
	}
	return nil, ErrNilRequestBody
}

func buildRequestQuery(p util.Map) string {
	query := p.GetD(DataTypeQuery, "")
	switch v := query.(type) {
	case string:
		return v
	case util.Map:
		return v.URLEncode()
	default:
		return ""
	}
}

type request struct {
	client   *http.Client
	function RequestBuildFunc
	method   string
	url      string
	body     interface{}
}

func BuildRequester() Requester {
	return nil
}

// Do ...
func (r *request) Do(ctx context.Context) Responder {
	log.Debug("Requester|Do", r.method, r.url, r.body)
	request := r.function(r.method, r.url, r.body)
	if request == nil {
		return Err(nil, errors.New("nil request"))
	}
	return do(ctx, r.client, request)
}

//type getRequester struct {
//	client *http.Client
//	method string
//	url    string
//}

// Request ...
//func (r *getRequester) Do(ctx context.Context) Responder {
//	request := processNothing(r.method, r.url, nil)
//	return do(ctx, r.client, request)
//}

func processNothing(method, url string, i interface{}) *http.Request {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Error(err)
		return nil
	}
	return request
}

//
//type postRequester struct {
//	client *http.Client
//	method string
//	url    string
//	body   util.Map
//}

//func (r *postRequester) Do(ctx context.Context) Responder {
//	request := buildRequester(r.method, r.url, r.body)
//	return do(ctx, r.client, request)
//}
//type multiRequester struct {
//	client *http.Client
//	method string
//	url    string
//	body   interface{}
//}
//
//func (r *multiRequester) Do(ctx context.Context) Responder {
//	request := processMultipart(r.method, r.url, r.body)
//	return do(ctx, r.client, request)
//}

// Request ...
//func (r *multiRequester) Request() *http.Request {
//	buf := bytes.Buffer{}
//	writer := multipart.NewWriter(&buf)
//	defer writer.Close()
//	log.Debug("processMultipart|i", r.body)
//	switch v := r.body.(type) {
//	case util.Map:
//		path := v.GetString("media")
//		fh, e := os.Open(path)
//		if e != nil {
//			log.Debug("processMultipart|e", e)
//			return nil
//		}
//		defer fh.Close()
//
//		fw, e := writer.CreateFormFile("media", path)
//		if e != nil {
//			log.Debug("processMultipart|e", e)
//			return nil
//		}
//
//		if _, e = io.Copy(fw, fh); e != nil {
//			log.Debug("processMultipart|e", e)
//			return nil
//		}
//		des := v.GetMap("description")
//		if des != nil {
//			writer.WriteField("description", string(des.ToJSON()))
//		}
//	}
//	request, err := http.NewRequest(r.method, r.url, &buf)
//	if err != nil {
//		return nil
//	}
//	request.Header.Set("Content-Type", writer.FormDataContentType())
//	return request
//}

func processMultipart(method, url string, i interface{}) *http.Request {
	buf := bytes.Buffer{}
	writer := multipart.NewWriter(&buf)
	defer writer.Close()
	log.Debug("processMultipart|i", i)
	switch v := i.(type) {
	case util.Map:
		path := v.GetString("media")
		fh, err := os.Open(path)
		if err != nil {
			log.Error("processMultipart|err", err)
			return nil
		}
		defer func() {
			err = fh.Close()
			if err != nil {
				log.Error(err)
			}
		}()

		fw, err := writer.CreateFormFile("media", path)
		if err != nil {
			log.Error("processMultipart|err", err)
			return nil
		}

		if _, err = io.Copy(fw, fh); err != nil {
			log.Error("processMultipart|err", err)
			return nil
		}
		des := v.GetMap("description")
		if des != nil {
			_ = writer.WriteField("description", string(des.ToJSON()))
		}
	}
	request, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request
}

func toXMLReader(v interface{}) io.Reader {
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

func processXML(method, url string, i interface{}) *http.Request {
	log.Debug(method, url, i)
	request, err := http.NewRequest(method, url, toXMLReader(i))
	if err != nil {
		log.Error(err)
		return nil
	}
	request.Header["Content-Type"] = []string{"application/xml; charset=utf-8"}
	return request
}

func toJSONReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("toJSONReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("toJSONReader|[]byte", string(v))
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("toJSONReader|util.Map", v.String())
		reader = bytes.NewReader(v.ToJSON())
	default:
		log.Debug("toJSONReader|default", v)
		if v0, e := json.Marshal(v); e == nil {
			log.Debug("toJSONReader|v0", string(v0), e)
			reader = bytes.NewReader(v0)
		}
	}
	return reader
}

func processJSON(method, url string, i interface{}) *http.Request {
	request, err := http.NewRequest(method, url, toJSONReader(i))
	if err != nil {
		log.Error(err)
		return nil
	}
	request.Header["Content-Type"] = []string{"application/json; charset=utf-8"}
	return request
}

func toFormReader(v interface{}) io.Reader {
	var reader io.Reader
	switch v := v.(type) {
	case string:
		log.Debug("toFormReader|string", v)
		reader = strings.NewReader(v)
	case []byte:
		log.Debug("toFormReader|[]byte", string(v))
		reader = bytes.NewReader(v)
	case util.Map:
		log.Debug("toFormReader|util.Map", v.URLEncode())
		reader = strings.NewReader(v.URLEncode())
	case url.Values:
		log.Debug("toFormReader|util.Map", v.Encode())
		reader = strings.NewReader(v.Encode())
	default:
		//do nothing
	}
	return reader
}

func processForm(method, url string, i interface{}) *http.Request {
	request, err := http.NewRequest(method, url, toFormReader(i))
	if err != nil {
		log.Error(err)
		return nil
	}

	request.Header["Content-Type"] = []string{"application/x-www-form-urlencoded; charset=utf-8"}
	return request
}

func buildRequester(method, url string, m util.Map) Requester {
	request := &request{
		client:   buildClient(m),
		function: processNothing,
		method:   method,
		url:      buildRequestURL(url, m),
		body:     nil,
	}
	switch {
	case m.Has(DataTypeJSON):
		request.function = processJSON
		request.body = m.Get(DataTypeJSON)
	case m.Has(DataTypeXML):
		request.function = processXML
		request.body = m.Get(DataTypeXML)
	case m.Has(DataTypeForm):
		request.function = processForm
		request.body = m.Get(DataTypeForm)
	case m.Has(DataTypeMultipart):
		request.function = processMultipart
		request.body = m.Get(DataTypeMultipart)
	}

	return request
}
