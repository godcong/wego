package official

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
)

/*Account Account*/
type Account struct {
	config          *core.Config
	client          *core.Client
	token           *core.AccessToken
	base            *Base
	menu            *Menu
	server          *Server
	customerService *CustomerService
}

func init() {

}

func newOfficialAccount(config *core.Config) *Account {
	client := core.NewClient(config)
	token := core.NewAccessToken(config, client)

	account := &Account{
		app:    application,
		Config: config,
		client: client,
		token:  token,
	}

	client.SetDomain(domain)
	client.SetDataType(core.DataTypeJSON)
	return account
}

/*Server Server*/
func (m *Account) Server() wego.Server {
	if m.server == nil {
		m.server = NewServer()
	}
	return m.server
}

/*Base Base*/
func (m *Account) Base() wego.Base {
	if m.base == nil {
		m.base = newBase(m)
	}
	return m.base
}

/*Menu Menu*/
func (m *Account) Menu() wego.Menu {
	if m.menu == nil {
		m.menu = newMenu(m)
	}
	return m.menu
}

/*AccessToken AccessToken*/
func (m *Account) AccessToken() *core.AccessToken {
	return m.token
}

//
//
//func (m *Account) Online() {
//
//}
//
//func (m *Account) Create(account, nickname string) {
//
//}
//
//func (m *Account) Update(account, nickname string) {
//
//}
//
//func (m *Account) Delete(account string) {
//
//}
//
//func (m *Account) Invite(account, wechatId string) {
//
//}
//
//func (m *Account) SetAvatar(account, path string) {
//
//}
//
//func (m *Account) Send(message util.Map) {
//
//}
//
//func (m *Account) Message(message util.Map) {
//
//}
