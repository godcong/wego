package core

import (
	"bytes"
	"io"
	"net/http"
)

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
