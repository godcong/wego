package mini_program

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/crypt"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

type Auth struct {
	config.Config
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
func (a *Auth) Session(code string) util.Map {
	params := util.Map{
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	a.client.SetDomain(core.NewDomain("mini_program"))
	resp := a.client.HttpGet(
		a.client.Link(core.SNS_JSCODE2SESSION_URL_SUFFIX),
		params,
	)
	return resp.ToMap()
}

func (a *Auth) UserInfoByCode(code, encrypted, iv string) []byte {
	p := a.Session(code)
	if !p.Has("errcode") {
		return a.UserInfo(p.GetString("session_key"), encrypted, iv)
	}
	return nil
}

func (a *Auth) UserInfo(key, encrypted, iv string) []byte {
	r, e := a.dc.Decrypt(encrypted, iv, key)
	if e != nil {
		log.Error(e)
		return nil
	}
	return r
}
