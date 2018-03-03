package official_account

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
)

type Server struct {
	message         *core.Message
	defaultCallback []ServerCallback
	callback        map[message.MsgType][]ServerCallback
}

type ServerCallback func(message *core.Message) string

func (s *Server) RegisterCallback(sc ServerCallback, types ...message.MsgType) {
	size := len(types)
	if size == 0 {
		s.defaultCallback = append(s.defaultCallback, sc)
		return
	}
	for _, t := range types {
		if callback, b := s.callback[t]; b {
			s.callback[t] = append(callback, sc)
		} else {
			s.callback[t] = []ServerCallback{sc}
		}
	}
}

func (s *Server) Callback(message *core.Message) string {
	result := ""
	for _, v := range s.defaultCallback {
		result = v(message)
	}

	if v0, b := s.callback[message.GetType()]; b {
		for _, v := range v0 {
			result = v(message)
		}
	}
	return result
}

func MessageProcess(msg *core.Message) string {
	switch msg.GetType() {
	case message.MESSAGE_TYPE_IMAGE:

	}
	return ""
}

func NewServer() Server {
	return Server{
		message:         nil,
		defaultCallback: []ServerCallback{},
		callback:        map[message.MsgType][]ServerCallback{},
	}
}
