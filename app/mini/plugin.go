package mini

import (
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Plugin Plugin */
type Plugin struct {
	*Program
	dc *cipher.DataCrypt
}

func newPlugin(program *Program) *Plugin {
	return &Plugin{
		Program: program,
		dc:      cipher.NewDataCrypt(program.GetString("app_id")),
	}
}

/*NewPlugin NewPlugin */
func NewPlugin(config *core.Config) *Plugin {
	return newPlugin(NewMiniProgram(config))
}

// Apply ...
func (p *Plugin) Apply(appID string) core.Responder {
	token := p.accessToken.GetToken()
	return core.PostJSON(Link(wxaPlugin), token.KeyMap(), util.Map{
		"action":       "apply",
		"plugin_appid": appID,
	})
}

// List ...
func (p *Plugin) List() core.Responder {
	token := p.accessToken.GetToken()
	return core.PostJSON(Link(wxaPlugin), token.KeyMap(), util.Map{
		"action": "list",
	})
}

// Unbind ...
func (p *Plugin) Unbind(appID string) core.Responder {
	token := p.accessToken.GetToken()
	return core.PostJSON(Link(wxaPlugin), token.KeyMap(), util.Map{
		"action":       "unbind",
		"plugin_appid": appID,
	})
}

// DevApplyList ...
func (p *Plugin) DevApplyList(appID string, page, num int) core.Responder {
	token := p.accessToken.GetToken()
	return core.PostJSON(Link(wxaDevPlugin), token.KeyMap(), util.Map{
		"action": "dev_apply_list",
		"page":   page,
		"num":    num,
	})
}

// DevDelete ...
func (p *Plugin) DevDelete() core.Responder {
	return p.devAction(util.Map{
		"action": "dev_delete",
	})
}

// DevRefuse ...
func (p *Plugin) DevRefuse(reason string) core.Responder {
	return p.devAction(util.Map{
		"action": "dev_refuse",
		"reason": reason,
	})
}

// DevAgree ...
func (p *Plugin) DevAgree(appID string) core.Responder {
	return p.devAction(util.Map{
		"action": "dev_agree",
		"appid":  appID,
	})
}

func (p *Plugin) devAction(maps util.Map) core.Responder {
	token := p.accessToken.GetToken()
	return core.PostJSON(Link(wxaDevPlugin), token.KeyMap(), maps)
}
