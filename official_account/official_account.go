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
	menu            *Menu
	server          *Server
	customerService *CustomerService
}

var defaultConfig core.Config
var account *OfficialAccount

func init() {
	core.Debug("OfficialAccount|init")
	defaultConfig = core.GetConfig("official_account.default")
	account = newOfficialAccount()

	app := core.App()
	app.Register("official_account", account)
}

func newOfficialAccount() *OfficialAccount {
	app := core.App()
	client := core.NewClient(defaultConfig)
	token := core.NewAccessToken(defaultConfig, client)
	domain := core.NewDomain("official_account")

	account := &OfficialAccount{
		app:    app,
		client: client,
		token:  token,
	}

	client.SetDomain(domain)
	client.SetDataType(core.DATA_TYPE_JSON)
	return account
}

func (m *OfficialAccount) Server() wego.Server {
	if m.server == nil {
		m.server = NewServer()
	}
	return m.server
}

func (m *OfficialAccount) Base() wego.Base {
	if m.base == nil {
		m.base = newBase(m)
	}
	return m.base
}

func (m *OfficialAccount) Menu() wego.Menu {
	if m.menu == nil {
		m.menu = newMenu(m)
	}
	return m.menu
}

//
//
//func (m *OfficialAccount) Online() {
//
//}
//
//func (m *OfficialAccount) Create(account, nickname string) {
//
//}
//
//func (m *OfficialAccount) Update(account, nickname string) {
//
//}
//
//func (m *OfficialAccount) Delete(account string) {
//
//}
//
//func (m *OfficialAccount) Invite(account, wechatId string) {
//
//}
//
//func (m *OfficialAccount) SetAvatar(account, path string) {
//
//}
//
//func (m *OfficialAccount) Send(message core.Map) {
//
//}
//
//func (m *OfficialAccount) Message(message core.Map) {
//
//}
