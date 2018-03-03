package wego

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/official_account"
)

type OfficialAccount interface {
	Base() *official_account.Base
}

type Base interface {
	GetCallbackIp() core.Map
}

func GetOfficialAccount() OfficialAccount {
	obj := GetApp().Get("official_account").(*official_account.OfficialAccount)
	core.Debug("GetOfficialAccount|official_account:", obj)
	return obj
}
