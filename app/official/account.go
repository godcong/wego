package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Account Account*/
type Account struct {
	*core.Config
	Sub         util.Map
	client      *core.Client
	accessToken *core.AccessToken
}

func newOfficialAccount(config *core.Config, p util.Map) *Account {
	return &Account{
		Config: config,
		Sub:    p,
	}
}

//NewOfficialAccount return a official account
func NewOfficialAccount(config *core.Config, v ...interface{}) *Account {
	client := core.ClientGet(v)
	accessToken := newAccessToken(util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	})
	accessToken.SetClient(client)

	account := newOfficialAccount(config, util.Map{})
	account.SetClient(client)
	account.SetAccessToken(accessToken)
	return account
}

// AccessToken ...
func (a *Account) AccessToken() *core.AccessToken {
	return a.accessToken
}

// SetAccessToken ...
func (a *Account) SetAccessToken(accessToken *core.AccessToken) {
	a.accessToken = accessToken
}

// Client ...
func (a *Account) Client() *core.Client {
	return a.client
}

//SetClient set client replace the default client
func (a *Account) SetClient(client *core.Client) {
	a.client = client
}

/*Server Server*/
func (a *Account) Server() *Server {
	obj, b := a.Sub["Server"]
	if !b {
		obj = newServer(a)
		a.Sub["Server"] = obj
	}
	return obj.(*Server)
}

/*Base Base*/
func (a *Account) Base() *Base {
	obj, b := a.Sub["Base"]
	if !b {
		obj = newBase(a)
		a.Sub["Base"] = obj
	}
	return obj.(*Base)
}

/*Menu Menu*/
func (a *Account) Menu() *Menu {
	obj, b := a.Sub["Menu"]
	if !b {
		obj = newMenu(a)
		a.Sub["Menu"] = obj
	}
	return obj.(*Menu)
}

/*OAuth OAuth*/
func (a *Account) OAuth() *OAuth {
	obj, b := a.Sub["OAuth"]
	if !b {
		obj = newOAuth(a)
		a.Sub["OAuth"] = obj
	}
	return obj.(*OAuth)
}

//Link 拼接地址
func Link(url string) string {
	return core.Connect(core.DefaultConfig().GetStringD("domain.official_account.url", domain), url)
}
