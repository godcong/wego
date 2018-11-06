package mini

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Program Program */
type Program struct {
	*core.Config
	Sub         util.Map
	client      *core.Client
	accessToken *core.AccessToken
}

// Client ...
func (p *Program) Client() *core.Client {
	return p.client
}

// SetClient ...
func (p *Program) SetClient(client *core.Client) {
	p.client = client
}

// AccessToken ...
func (p *Program) AccessToken() *core.AccessToken {
	return p.accessToken
}

// SetAccessToken ...
func (p *Program) SetAccessToken(accessToken *core.AccessToken) {
	p.accessToken = accessToken
}

// NewMiniProgram ...
func NewMiniProgram(config *core.Config, v ...interface{}) *Program {
	client := core.ClientGet(v)
	accessToken := newAccessToken(util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	})
	accessToken.SetClient(client)

	account := newMiniProgram(config, util.Map{})
	account.SetClient(client)
	account.SetAccessToken(accessToken)
	return account
}

func newMiniProgram(config *core.Config, p util.Map) *Program {
	return &Program{
		Config: config,
		Sub:    p,
	}
}

// Auth ...
func (p *Program) Auth() *Auth {
	obj, b := p.Sub["Auth"]
	if !b {
		obj = newAuth(p)
		p.Sub["Auth"] = obj
	}
	return obj.(*Auth)
}

// Message ...
func (p *Program) Message() *Message {
	obj, b := p.Sub["Message"]
	if !b {
		obj = newMessage(p)
		p.Sub["Message"] = obj
	}
	return obj.(*Message)
}

// Template ...
func (p *Program) Template() *Template {
	obj, b := p.Sub["Template"]
	if !b {
		obj = newTemplate(p)
		p.Sub["Template"] = obj
	}
	return obj.(*Template)
}

//Link 拼接地址
func Link(url string) string {
	return core.Connect(core.DefaultConfig().GetStringD("domain.mini_program.url", domain), url)
}
