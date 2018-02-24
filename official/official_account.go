package official

import "github.com/godcong/wego/core"

//type OfficialAccount interface {
//	AccessToken() AccessTokenInterface
//}
//
//type officialAccount struct {
//	Config
//	app Application
//}
//
//type OfficialAccountAccessToken struct {
//}
//
//func NewOfficialAccount(application Application) OfficialAccount {
//	return &officialAccount{
//		Config: application.GetConfig("official.default"),
//		app:    application,
//	}
//}
//
//func (a *officialAccount) AccessToken() AccessTokenInterface {
//	return NewAccessToken(a.app, a.Config)
//}

type Base struct {
	core.Config
	core.Client
}

func (b *Base) ClearQuota() core.Map {
	params := core.Map{
		"appid": b.Get("app_id"),
	}
	return b.HttpPostJson(core.CLEAR_QUOTA_URL_SUFFIX, params, nil)
}

func (b *Base) GetValidIps() core.Map {
	params := core.Map{
		"appid": b.Get("app_id"),
	}
	return b.HttpPostJson(core.GETCALLBACKIP_URL_SUFFIX, params, nil)
}
