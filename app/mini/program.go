package mini

import (
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

// NewAble ...
type NewAble func(program *Program) interface{}

var subLists = util.Map{
	"AppCode": newAppcode,
}

/*Program Program */
type Program struct {
	*core.Config
	Sub         util.Map
	cipher      cipher.Cipher
	accessToken *core.AccessToken
}

func newMiniProgram(config *core.Config, p util.Map) *Program {
	return &Program{
		Config: config,
		Sub:    p,
	}
}

// NewMiniProgram ...
func NewMiniProgram(config *core.Config, v ...interface{}) *Program {

	accessToken := newAccessToken(util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	})
	account := newMiniProgram(config, util.Map{})
	account.SetAccessToken(accessToken)
	return account
}
func subInit(payment *Program, p util.Map) *Program {
	for k, v := range p {
		if vv, b := v.(NewAble); b {
			payment.Sub[k] = vv(payment)
		}
	}
	return payment
}

// SubInit ...
func (p *Program) SubInit() *Program {
	return subInit(p, subLists)
}

// SubExpectInit ...
func (p *Program) SubExpectInit(except ...string) *Program {
	return subInit(p, subLists.Expect(except))
}

// SubOnlyInit ...
func (p *Program) SubOnlyInit(only ...string) *Program {
	return subInit(p, subLists.Only(only))
}

// AccessToken ...
func (p *Program) AccessToken() *core.AccessToken {
	return p.accessToken
}

// SetAccessToken ...
func (p *Program) SetAccessToken(accessToken *core.AccessToken) {
	p.accessToken = accessToken
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
	return core.Splice(core.DefaultConfig().GetStringD("domain.mini_program.url", domain), url)
}
