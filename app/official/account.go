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

func newAccount(config *core.Config) *Account {

	account := &Account{
		config: config,
		client: core.DefaultClient(),
		token:  newAccessToken(config),
		sub:    util.Map{},
	}

	return account
}

//NewAccount return a official account
func NewAccount(config *core.Config) *Account {
	return newAccount(config)
}

//SetClient set client replace the default client
func (a *Account) SetClient(client *core.Client) *Account {
	a.client = client
	return a
}

/*Server Server*/
func (a *Account) Server() *Server {
	obj, b := a.sub["Server"]
	if !b {
		obj = newServer(a)
		a.sub["Server"] = obj
	}
	return obj.(*Server)
}

/*Base Base*/
func (a *Account) Base() *Base {
	obj, b := a.sub["Base"]
	if !b {
		obj = newBase(a)
		a.sub["Base"] = obj
	}
	return obj.(*Base)
}

/*Menu Menu*/
func (a *Account) Menu() *Menu {
	obj, b := a.sub["Menu"]
	if !b {
		obj = newMenu(a)
		a.sub["Menu"] = obj
	}
	return obj.(*Menu)
}

/*OAuth OAuth*/
func (a *Account) OAuth() *OAuth {
	obj, b := a.sub["OAuth"]
	if !b {
		obj = newOAuth(a)
		a.sub["OAuth"] = obj
	}
	return obj.(*OAuth)
}

//Link 拼接地址
func Link(url string) string {
	return core.Connect(core.DefaultConfig().GetStringD("domain.official_account.url", domain), url)
}
