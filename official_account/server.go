package official_account

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
)

type Server struct {
	message         *core.Message
	defaultCallback []core.MessageCallback
	callback        map[message.MsgType][]core.MessageCallback
}

func (s *Server) RegisterCallback(sc core.MessageCallback, types ...message.MsgType) {
	size := len(types)
	if size == 0 {
		s.defaultCallback = append(s.defaultCallback, sc)
		return
	}
	for _, t := range types {
		if callback, b := s.callback[t]; b {
			s.callback[t] = append(callback, sc)
		} else {
			s.callback[t] = []core.MessageCallback{sc}
		}
	}
}
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, e := ioutil.ReadAll(r.Body)
	if e != nil {
		core.Error(e)
		return
	}
	message := new(core.Message)
	e = xml.Unmarshal(body, message)
	if e != nil {
		core.Error(e)
		return
	}
	result := s.CallbackFunc(message)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	return
}

func (s *Server) CallbackFunc(message *core.Message) []byte {
	var result []byte
	for _, v := range s.defaultCallback {
		if r := v(message); r != nil {
			result = r
		}
	}

	if v0, b := s.callback[message.GetType()]; b {
		for _, v := range v0 {
			if r := v(message); r != nil {
				result = r
			}
		}
	}
	return result
}

func MessageProcess(msg *core.Message) string {
	switch msg.GetType() {
	case message.TypeImage:

	}
	return ""
}

func NewServer() *Server {
	return &Server{
		message:         nil,
		defaultCallback: []core.MessageCallback{},
		callback:        map[message.MsgType][]core.MessageCallback{},
	}
}
