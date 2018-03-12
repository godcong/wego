package wego

import (
	"net/http"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
)

type Base interface {
	GetCallbackIp() core.Map
	ClearQuota() core.Map
}
type Menu interface {
}

type Server interface {
	RegisterCallback(sc core.ServerCallback, types ...message.MsgType)
	Monitor(w http.ResponseWriter, r *http.Request) error
}

type OfficialAccount interface {
	Base() Base
	Menu() Menu
	Server() Server
}

func GetOfficialAccount() OfficialAccount {
	obj := GetApp().Get("official_account").(OfficialAccount)
	core.Debug("GetOfficialAccount|obj:", obj)
	return obj
}
