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
	return p.client.PostJSON(Link(wxaPlugin), nil, util.Map{
		"action":       "apply",
		"plugin_appid": appID,
	})
}
