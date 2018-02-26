package core

import "net/http"

type Response struct {
	response *http.Response
	Content  []byte
	Type     string
	Error    error
}

type ResponseType string

const (
	RESPONSE_TYPE_JSON   ResponseType = "json"
	RESPONSE_TYPE_XML                 = "xml"
	RESPONSE_TYPE_ARRAY               = "array"
	RESPONSE_TYPE_STRUCT              = "struct"
	RESPONSE_TYPE_MAP                 = "map"
	RESPONSE_TYPE_RAW                 = "raw"
)
