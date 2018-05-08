package payment

import (
	"encoding/xml"
	"net/http"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/util"
)

type Server struct {
	config.Config
	*Payment
	callback []core.PaymentCallback
}

type ActionResult struct {
	XMLName    xml.Name      `xml:"xml"`
	ReturnCode message.CDATA `xml:"return_code"`
	ReturnMsg  message.CDATA `xml:"return_msg"`
}

var ACTION_SUCCESS = ActionResult{
	ReturnCode: message.CDATA{
		Value: "SUCCESS",
	},
	ReturnMsg: message.CDATA{
		Value: "OK",
	},
}

var ACTION_FAIL = ActionResult{
	ReturnCode: message.CDATA{
		Value: "FAIL",
	},
	ReturnMsg: message.CDATA{
		Value: "OK",
	},
}

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

func (s *Server) ProcessCallback(p util.Map) *ActionResult {
	rlt := ACTION_SUCCESS
	if s.callback == nil {
		rlt.ReturnMsg = message.CDATA{
			Value: "UNPROCESSED",
		}

		return &rlt
	}

	for _, v := range s.callback {
		rlt.ReturnMsg = message.CDATA{
			Value: "PROCESSED",
		}
		if v(p) == false {
			rlt.ReturnCode = message.CDATA{
				Value: "FAIL",
			}
		}
	}
	return &rlt
}
