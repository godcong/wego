package official_account

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/core"
)

type OfficialAccount struct {
	app             *core.Application
	client          *core.Client
	token           *core.AccessToken
	base            *Base
	customerService *CustomerService
}

func (m *OfficialAccount) Base() wego.Base {
	return m.base
}

func DataType() core.DataType {
	return core.DATA_TYPE_XML
}

var defaultConfig core.Config

func init() {
	defaultConfig = core.GetConfig("official_account.default")

	app := core.App()
	app.Register("official_account", newOfficialAccount())
}

func newOfficialAccount() *OfficialAccount {
	app := core.App()
	client := core.NewClient(defaultConfig)
	token := core.NewAccessToken(defaultConfig, client)
	domain := core.NewDomain("official_account")
	base := newBase(client, token)
	account := &OfficialAccount{
		app:    app,
		client: client,
		token:  token,
		base:   base,
		//customerService: nil,
	}
	account.base = &Base{
		Config: defaultConfig,
		client: client,
		token:  token,
	}
	client.SetDomain(domain)
	client.SetDataType(core.DATA_TYPE_JSON)
	return account
}

func (m *OfficialAccount) Server() wego.Server {
	return NewServer()
}

func (m *OfficialAccount) Online() {

}

func (m *OfficialAccount) Create(account, nickname string) {

}

func (m *OfficialAccount) Update(account, nickname string) {

}

func (m *OfficialAccount) Delete(account string) {

}

func (m *OfficialAccount) Invite(account, wechatId string) {

}

func (m *OfficialAccount) SetAvatar(account, path string) {

}

func (m *OfficialAccount) Send(message core.Map) {

}

func (m *OfficialAccount) Message(message core.Map) {

}

//type OfficialAccountAccessToken struct {
//}
//
//func NewOfficialAccount(application Application) OfficialAccount {
//	return &officialAccount{
//		Config: application.GetConfig("official_account.default"),
//		app:    application,
//	}
//}
//
//func (a *officialAccount) accessToken() AccessTokenInterface {
//	return NewAccessToken(a.app, a.Config)
//}
