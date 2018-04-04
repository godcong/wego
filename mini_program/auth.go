package mini_program

import "github.com/godcong/wego/core"

type Auth struct {
	core.Config
	*MiniProgram
}

func newAuth(program *MiniProgram) *Auth {
	return &Auth{
		Config:      defaultConfig,
		MiniProgram: program,
	}
}

func NewAuth() *Auth {
	return newAuth(program)
}

// 成功:
// {"openid":"oE_gl0Yr54fUjBhU5nBlP4hS2efo","session_key":"UaPsfKqS9eJYxi1PCYYxuA=="}
// 失败:
// {"errcode":40163,"errmsg":"code been used, hints: [ req_id: gjUEFA0981th20 ]"}
func (a *Auth) Session(code string) core.Map {
	params := core.Map{
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	a.client.SetDomain(core.NewDomain("mini_program"))
	resp := a.client.HttpGet(a.client.Link(core.SNS_JSCODE2SESSION_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): params,
		})
	return resp.ToMap()
}

func (a *Auth) UserInfo(code, encrypted, iv string) core.Map {
	p := a.Session(code)
	if p.Has("errcode") {
		return nil
	}
	//sessionKey := p.GetString("session_key")
	//userInfo :=
	return nil
}
