package mini_program

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/crypt"
)

type Auth struct {
	core.Config
	*MiniProgram
	dc *crypt.DataCrypt
}

func newAuth(program *MiniProgram) *Auth {
	return &Auth{
		Config:      defaultConfig,
		MiniProgram: program,
		dc:          crypt.NewDataCrypt(defaultConfig.Get("app_id")),
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

func (a *Auth) UserInfo(code, encrypted, iv string) []byte {
	p := a.Session(code)
	if p.Has("errcode") {
		core.Error(p)
		return nil
	}
	sessionKey := p.GetString("session_key")
	//mp := core.Map{}
	r, e := a.dc.Decrypt(encrypted, iv, sessionKey)
	if e != nil {
		core.Error(e)
	}
	return r
	//return mp.ParseJson(r)
}
