package mini

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/crypt"
	"github.com/godcong/wego/util"
)

/*Plugin Plugin */
type Plugin struct {
	*Program
	dc *crypt.DataCrypt
}

func newPlugin(program *Program) *Plugin {
	return &Plugin{
		Program: program,
		dc:      crypt.NewDataCrypt(program.GetString("app_id")),
	}
}

/*NewPlugin NewPlugin */
func NewPlugin(config *core.Config) *Plugin {
	return newPlugin(NewMiniProgram(config))
}

func (p *Plugin) Apply(appID string) core.Response {
	token := p.accessToken.GetToken()
	return p.client.PostJSON(Link(wxaPlugin), token.KeyMap(), util.Map{
		"action":       "apply",
		"plugin_appid": appID,
	})
}
func (p *Plugin) List() core.Response {
	token := p.accessToken.GetToken()
	return p.client.PostJSON(Link(wxaPlugin), token.KeyMap(), util.Map{
		"action": "list",
	})
}

func (p *Plugin) Unbind(appID string) core.Response {
	token := p.accessToken.GetToken()
	return p.client.PostJSON(Link(wxaPlugin), token.KeyMap(), util.Map{
		"action":       "unbind",
		"plugin_appid": appID,
	})
}

func (p *Plugin) DevApplyList(appID string, page, num int) core.Response {
	token := p.accessToken.GetToken()
	return p.client.PostJSON(Link(wxaDevPlugin), token.KeyMap(), util.Map{
		"action": "dev_apply_list",
		"page":   page,
		"num":    num,
	})
}

func (p *Plugin) DevRefuse(reason string) core.Response {
	return p.devAction(util.Map{
		"action": "dev_refuse",
		"reason": reason,
	})
}

func (p *Plugin) DevAgree(appID string) core.Response {
	return p.devAction(util.Map{
		"action": "dev_agree",
		"appid":  appID,
	})
}

func (p *Plugin) devAction(maps util.Map) core.Response {
	token := p.accessToken.GetToken()
	return p.client.PostJSON(Link(wxaDevPlugin), token.KeyMap(), maps)
}
