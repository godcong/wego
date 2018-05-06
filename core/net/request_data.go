package net

import (
	"io"
	"net/http"
)

type RequestData struct {
	Query  string
	Method string
	Header http.Header
	Body   io.Reader
}

func (r *RequestData) SetHeaderJson() *RequestData {
	r.Header.Set("Content-Type", "application/json")
	return r
}

func (r *RequestData) SetHeaderXml() *RequestData {
	r.Header.Set("Content-Type", "application/xml")
	return r
}
