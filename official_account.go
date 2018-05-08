package wego

import (
	"net/http"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

type Base interface {
	GetCallbackIp() util.Map
	ClearQuota() util.Map
}
type Menu interface {
}

type Server interface {
	RegisterCallback(sc core.MessageCallback, types ...message.MsgType)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type OfficialAccount interface {
	Base() Base
	Menu() Menu
	Server() Server
}

func GetOfficialAccount() OfficialAccount {
	obj := GetApp().Get("official_account").(OfficialAccount)
	log.Debug("GetOfficialAccount|obj:", obj)
	return obj
}
