package wego

import (
	"github.com/godcong/wego/core"
)

type Base interface {
	GetCallbackIp() core.Map
}

type OfficialAccount interface {
	Base() Base
}

func GetOfficialAccount() OfficialAccount {
	obj := GetApp().Get("official_account").(OfficialAccount)
	core.Debug("GetOfficialAccount|obj:", obj)
	return obj
}
