package official

import "github.com/godcong/wego/core"

//type OfficialAccount interface {
//	accessToken() AccessTokenInterface
//}
//
type OfficialAccount struct {
	core.Config
	client core.Client
	*Base
}

func DataType() core.DataType {
	return core.DATA_TYPE_XML
}

func init() {
	app := core.App()
	app.Register("official_account", newOfficialAccount())

}

func newOfficialAccount() *OfficialAccount {
	config := core.GetConfig("payment.default")
	official0 := &OfficialAccount{
		Config: config,
		client: core.NewClient(config),
		Base: &Base{
			Config: config,
			Client: core.NewClient(config),
		},
	}
	return official0
}

func (m *OfficialAccount) prefix(s string) string {
	return core.API_WEIXIN_URL_SUFFIX + s
}

func (m *OfficialAccount) List() {
	m.HttpGet(m.prefix(core.GETKFLIST_URL_SUFFIX), nil)
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
//		Config: application.GetConfig("official.default"),
//		app:    application,
//	}
//}
//
//func (a *officialAccount) accessToken() AccessTokenInterface {
//	return NewAccessToken(a.app, a.Config)
//}
