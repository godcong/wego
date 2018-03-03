package official_account

import (
	"github.com/godcong/wego/core"
)

//type OfficialAccount interface {
//	accessToken() AccessTokenInterface
//}
//
type OfficialAccount struct {
	core.Config
	client core.Client
	token  core.AccessToken
	app    *core.Application

	base            *Base
	customerService *CustomerService
}

func (m *OfficialAccount) Base() *Base {
	return m.base
}

func DataType() core.DataType {
	return core.DATA_TYPE_XML
}

func init() {
	app := core.App()
	app.Register("official_account", newOfficialAccount())

}

func newOfficialAccount() *OfficialAccount {
	config := core.GetConfig("official_account.default")
	official0 := &OfficialAccount{
		Config: config,
		client: core.NewClient(config),
		base: &Base{
			Config: config,
			Client: core.NewClient(config),
		},
	}
	official0.client.SetDomain(core.NewDomain("official_account"))
	official0.client.SetDataType(core.DATA_TYPE_JSON)
	official0.base = &Base{
		Config:          official0.Config,
		Client:          official0.client,
		AccessToken:     core.NewAccessToken(config, official0.client),
		OfficialAccount: official0,
	}

	return official0
}

func (m *OfficialAccount) prefix(s string) string {
	return core.API_WEIXIN_URL_SUFFIX + s
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
