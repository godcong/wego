package official

import (
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/util"
)

// NewAble ...
type NewAble func(account *Account) interface{}

var moduleLists = util.Map{
	"Base":   newBase,
	"JSSDK":  newJSSDK,
	"OAuth":  newOAuth,
	"Menu":   newMenu,
	"Ticket": newTicket,
}

/*Account Account*/
type Account struct {
	*core.Config
	Module      util.Map
	accessToken *core.AccessToken
}

func newOfficialAccount(config *core.Config, p util.Map) *Account {
	return &Account{
		Config: config,
		Module: p,
	}
}

//NewOfficialAccount return a official account
func NewOfficialAccount(config *core.Config, v ...interface{}) *Account {

	accessToken := newAccessToken(util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	})
	account := newOfficialAccount(config, util.Map{})
	account.SetAccessToken(accessToken)
	return account
}

func subInit(payment *Account, p util.Map) *Account {
	for k, v := range p {
		if vv, b := v.(NewAble); b {
			payment.Module[k] = vv(payment)
		}
	}
	return payment
}

// InitModule ...
func (a *Account) InitModule() *Account {
	return subInit(a, moduleLists)
}

// InitModuleExpect ...
func (a *Account) InitModuleExpect(except ...string) *Account {
	return subInit(a, moduleLists.Expect(except))
}

// InitModuleOnly ...
func (a *Account) InitModuleOnly(only ...string) *Account {
	return subInit(a, moduleLists.Only(only))
}

// AccessToken ...
func (a *Account) AccessToken() *core.AccessToken {
	return a.accessToken
}

// SetAccessToken ...
func (a *Account) SetAccessToken(accessToken *core.AccessToken) {
	a.accessToken = accessToken
}

/*Server Server*/
func (a *Account) Server() *Server {
	obj, b := a.Module["Server"]
	if !b {
		obj = newServer(a)
		a.Module["Server"] = obj
	}
	return obj.(*Server)
}

/*Base Base*/
func (a *Account) Base() *Base {
	obj, b := a.Module["Base"]
	if !b {
		obj = newBase(a)
		a.Module["Base"] = obj
	}
	return obj.(*Base)
}

/*Menu Menu*/
func (a *Account) Menu() *Menu {
	obj, b := a.Module["Menu"]
	if !b {
		obj = newMenu(a)
		a.Module["Menu"] = obj
	}
	return obj.(*Menu)
}

/*OAuth OAuth*/
func (a *Account) OAuth() *OAuth {
	obj, b := a.Module["OAuth"]
	if !b {
		obj = newOAuth(a)
		a.Module["OAuth"] = obj
	}
	return obj.(*OAuth)
}

// Ticket ...
func (a *Account) Ticket() *Ticket {
	obj, b := a.Module["Ticket"]
	if !b {
		obj = newTicket(a)
		a.Module["Ticket"] = obj
	}
	return obj.(*Ticket)
}

// JSSDK ...
func (a *Account) JSSDK() *JSSDK {
	obj, b := a.Module["JSSDK"]
	if !b {
		obj = newJSSDK(a)
		a.Module["JSSDK"] = obj
	}
	return obj.(*JSSDK)
}

// HandelMessageTypeNotify ...
func (a *Account) HandelMessageTypeNotify(msgType message.MsgType, f NotifyCallback) Notify {
	token := a.GetString("token")
	key := a.GetString("aes_key")
	id := a.GetString("app_id")
	return &messageTypeNotify{
		messageNotify: &messageNotify{
			Account:        a,
			NotifyCallback: f,
			bizMsg:         cipher.NewBizMsg(token, key, id),
		},
		msgType: msgType,
	}
}

// HandleTypeMessage ...
func (a *Account) HandleTypeMessage(f NotifyCallback) NotifyFunc {
	return a.HandleMessageNotify(f).ServeHTTP
}

// HandleMessageNotify ...
func (a *Account) HandleMessageNotify(f NotifyCallback) Notify {
	token := a.GetString("token")
	key := a.GetString("aes_key")
	id := a.GetString("app_id")
	return &messageNotify{
		Account:        a,
		NotifyCallback: f,
		bizMsg:         cipher.NewBizMsg(token, key, id),
	}
}

// HandleMessage ...
func (a *Account) HandleMessage(f NotifyCallback) NotifyFunc {
	return a.HandleMessageNotify(f).ServeHTTP
}

//Link 拼接地址
func Link(url string) string {
	return core.Connect(core.DefaultConfig().GetStringD("domain.official_account.url", domain), url)
}
