package core

import (
	"io"
	"net/http"
)

/*RequestData RequestData */
type RequestData struct {
	Query  string
	Method string
	Header http.Header
	Body   io.Reader
}

/*SetJSONHeader set content type with json */
func (r *RequestData) SetJSONHeader() *RequestData {
	r.Header.Set("Content-Type", "application/json")
	return r
}

/*SetXMLHeader set content type with xml */
func (r *RequestData) SetXMLHeader() *RequestData {
	r.Header.Set("Content-Type", "application/xml")
	return r
}
