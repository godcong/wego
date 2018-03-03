package mini_program

import "github.com/godcong/wego/core"

type Auth struct {
	core.Config
	*MiniProgram
}

func (a *Auth) Session(code string) core.Map {
	params := core.Map{
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	resp := a.GetClient().HttpGet(a.prefix(core.SNS_JSCODE2SESSION_URL_SUFFIX), params)
	return resp.ToMap()
}
