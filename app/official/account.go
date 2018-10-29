package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

const domain = "https://api.weixin.qq.com"

/*Account Account*/
type Account struct {
	config *core.Config
	client *core.Client
	token  *core.AccessToken
	sub    util.Map
}

func init() {

}

func newAccount(config *core.Config) *Account {
	token := core.NewAccessToken(config)

	account := &Account{
		config: config,
		//client: client,
		token: token,
		sub:   util.Map{},
	}

	return account
}

//NewAccount return a official account
func NewAccount(config *core.Config) *Account {
	return newAccount(config)
}

/*Server Server*/
func (m *Account) Server() *Server {
	obj, b := m.sub["Server"]
	if !b {
		obj = newServer(m)
		m.sub["Server"] = obj
	}
	return obj.(*Server)
}

/*Base Base*/
func (m *Account) Base() *Base {
	obj, b := m.sub["Base"]
	if !b {
		obj = newBase(m)
		m.sub["Base"] = obj
	}
	return obj.(*Base)
}

/*Menu Menu*/
func (m *Account) Menu() *Menu {
	obj, b := m.sub["Menu"]
	if !b {
		obj = newMenu(m)
		m.sub["Menu"] = obj
	}
	return obj.(*Menu)
}

/*OAuth OAuth*/
func (m *Account) OAuth() *OAuth {
	obj, b := m.sub["OAuth"]
	if !b {
		obj = newOAuth(m)
		m.sub["OAuth"] = obj
	}
	return obj.(*OAuth)
}

//Link 拼接地址
func Link(url string) string {
	return core.Connect(core.DefaultConfig().GetStringD("domain.official_account.url", domain), url)
}
