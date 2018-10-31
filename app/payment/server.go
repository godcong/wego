package payment

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*Server Server */
type Server struct {
	*Payment
	mType    string
	callback []core.PaymentCallback
}

/*ActionResult ActionResult */
type ActionResult struct {
	XMLName    xml.Name      `xml:"xml"`
	ReturnCode message.CDATA `xml:"return_code"`
	ReturnMsg  message.CDATA `xml:"return_msg"`
}

/*ActionSuccess ActionSuccess */
var ActionSuccess = ActionResult{
	ReturnCode: message.CDATA{
		Value: "SUCCESS",
	},
	ReturnMsg: message.CDATA{
		Value: "OK",
	},
}

/*ActionFail ActionFail */
var ActionFail = ActionResult{
	ReturnCode: message.CDATA{
		Value: "FAIL",
	},
	ReturnMsg: message.CDATA{
		Value: "OK",
	},
}

var result = []byte(`<xml><return_code><![CDATA[FAIL]]></return_code><return_msg><![CDATA[OK]]></return_msg></xml>`)

func newServer(p *Payment) *Server {
	return &Server{
		mType:    "xml",
		Payment:  p,
		callback: nil,
	}
}

/*NewServer NewServer */
func NewServer(config *core.Config) *Server {
	return newServer(NewPayment(config))
}

/*ServeHTTP 服务监听 */
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var bodyBytes []byte
	var rlt message.Messager

	//var err error
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	w.WriteHeader(http.StatusOK)
	if len(bodyBytes) == 0 {
		return
	}

	m := util.XMLToMap(bodyBytes)
	if validateCallback(m, s.GetString("key")) {
		rlt = s.ProcessCallback(m)

	}

	if rlt == nil {
		w.Write(result)
		return
	}
	rltXML, err := rlt.ToXML()
	//错误返回,并记录log
	if err != nil {
		log.Error(err)
		w.Write(result)
		return
	}

	if s.mType == "xml" {
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = []string{"application/xml; charset=utf-8"}
		}
	} else {
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = []string{"application/json; charset=utf-8"}
		}
	}
	log.Debug(string(rltXML))
	w.Write(rltXML)

}

func validateCallback(p util.Map, key string) bool {
	st := p.GetString("sign_type")
	ft := MakeSignHMACSHA256
	if st == "MD5" {
		ft = MakeSignMD5
	}

	sign := GenerateSignature(p, key, ft)
	if sign == p.GetString("sign") {
		return true
	}
	return false
}

/*AddCallback add callback */
func (s *Server) AddCallback(pc core.PaymentCallback) *Server {
	if s.callback == nil {
		s.callback = []core.PaymentCallback{}
	}
	s.callback = append(s.callback, pc)
	return s
}

/*SetCallback set callback */
func (s *Server) SetCallback(pc []core.PaymentCallback) *Server {
	s.callback = pc
	return s
}

/*GetCallback get callback */
func (s *Server) GetCallback() []core.PaymentCallback {
	return s.callback
}

/*ProcessCallback process callback */
func (s *Server) ProcessCallback(p util.Map) message.Messager {
	rlt := ActionSuccess
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

/*ToXML transfer action result to xml */
func (r *ActionResult) ToXML() ([]byte, error) {
	return xml.Marshal(r)
}

/*ToJSON transfer action result to json */
func (r *ActionResult) ToJSON() ([]byte, error) {
	m := util.Map{
		"return_code": r.ReturnCode,
		"return_msg":  r.ReturnMsg,
	}
	return m.ToJSON(), nil
}
