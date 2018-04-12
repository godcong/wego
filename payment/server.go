package payment

import (
	"net/http"

	"github.com/godcong/wego/core"
)

type Server struct {
	core.Config
	*Payment
	callback []core.PaymentCallback
}

var SUCCESS = core.Map{"return_code": "SUCCESS", "return_msg": "OK"}
var FAIL = core.Map{"return_code": "FAIL", "return_msg": "OK"}

func newServer(p *Payment) *Server {
	return &Server{
		Config:   defaultConfig,
		Payment:  payment,
		callback: nil,
	}
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.callback == nil {
		return
	}
}

func (s *Server) AddCallback(pc core.PaymentCallback) *Server {
	if s.callback == nil {
		s.callback = []core.PaymentCallback{}
	}
	s.callback = append(s.callback, pc)
	return s
}

func (s *Server) SetCallback(pc []core.PaymentCallback) *Server {
	s.callback = pc
	return s
}

func (s *Server) GetCallback() []core.PaymentCallback {
	return s.callback
}

func (s *Server) ProcessCallback(p core.Map) core.Map {
	rlt := core.Map{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
	}
	if s.callback == nil {
		rlt.Set("return_msg", "UNPROCESSED")
		return rlt
	}

	for _, v := range s.callback {
		rlt.Set("return_msg", "PROCESSED")
		if v(p) == false {
			rlt.Set("return_code", "FAIL")
		}
	}
	return rlt
}
