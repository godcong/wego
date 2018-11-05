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

// SetClient SetClient
func (p *Program) SetClient(c *core.Client) *Program {
	p.client = c
	return p
}

// GetClient GetClient
func (p *Program) GetClient() *core.Client {
	return p.client
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

//Link 拼接地址
func Link(url string) string {
	return core.Connect(core.DefaultConfig().GetStringD("domain.official_account.url", domain), url)
}
