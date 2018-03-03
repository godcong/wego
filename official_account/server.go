package official_account

import "github.com/godcong/wego/core"

type AccountServer struct {
	message *core.Message
}

type ServerCallback func(message *core.Message) string

func (a *AccountServer) RegisterCallback(s ServerCallback) {

}

func (a *AccountServer) Callback() {

}

func MessageProecess(message *core.Message) string {
	switch message.GetType() {
	case core.MESSAGE_TYPE_IMAGE:

	}
	return ""
}

func New() {
	core.NewMessage()
}
