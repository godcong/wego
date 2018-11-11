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

// ErrNilRequestBody ...
var ErrNilRequestBody = errors.New("nil request body")

func connectQuery(url, query string) string {
	if query == "" {
		return url
	}
	return url + "?" + query
}

func parseBody(body string) io.Reader {
	if body == "" {
		return nil
	}
	return bytes.NewBufferString(body)
}

func cloneHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
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
