package official_account

import "github.com/godcong/wego/core"

type Server struct {
	message         *core.Message
	defaultCallback []ServerCallback
	callback        map[core.MessageType][]ServerCallback
}

type ServerCallback func(message *core.Message) string

func (s *Server) RegisterCallback(sc ServerCallback, types ...core.MessageType) {
	size := len(types)
	if size == 0 {
		s.defaultCallback = append(s.defaultCallback, sc)
	}
}

func (s *Server) Callback() {

}

func MessageProcess(message *core.Message) string {
	switch message.GetType() {
	case core.MESSAGE_TYPE_IMAGE:

	}
	return ""
}

func NewServer() Server {
	return Server{
		message:         nil,
		defaultCallback: []ServerCallback{},
		callback:        map[core.MessageType][]ServerCallback{},
	}
}
