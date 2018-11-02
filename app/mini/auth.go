package mini

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/crypt"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*Auth Auth */
type Auth struct {
	*core.Config
	*Program
	dc *crypt.DataCrypt
}

func newAuth(program *Program) *Auth {
	return &Auth{
		Config:  defaultConfig,
		Program: program,
		dc:      crypt.NewDataCrypt(defaultConfig.Get("app_id")),
	}
}

/*NewAuth NewAuth */
func NewAuth() *Auth {
	return newAuth(program)
}

/*Session 登录凭证校验
临时登录凭证校验接口是一个 HTTPS 接口，开发者服务器使用 临时登录凭证code 获取 session_key 和 openid 等。
注意:
会话密钥session_key 是对用户数据进行加密签名的密钥。为了应用自身的数据安全，开发者服务器不应该把会话密钥下发到小程序，也不应该对外提供这个密钥。
UnionID 只在满足一定条件的情况下返回。具体参看UnionID机制说明
临时登录凭证code只能使用一次
接口地址:
https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
请求参数
参数	必填	说明
appid	是	小程序唯一标识
secret	是	小程序的 app secret
js_code	是	登录时获取的 code
grant_type	是	填写为 authorization_code
在不满足UnionID下发条件的情况下，返回参数
参数	说明
openid	用户唯一标识
session_key	会话密钥
在满足UnionID下发条件的情况下，返回参数
参数	说明
openid	用户唯一标识
session_key	会话密钥
unionid	用户在开放平台的唯一标识符

成功:
{"openid":"oE_gl0Yr54fUjBhU5nBlP4hS2efo","session_key":"UaPsfKqS9eJYxi1PCYYxuA=="}
失败:
{"errcode":40163,"errmsg":"code been used, hints: [ req_id: gjUEFA0981th20 ]"}
*/
func (a *Auth) Session(code string) util.Map {
	params := util.Map{
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	resp := a.client.Get(
		Link(snsJscode2sessionURLSuffix),
		params,
	)
	return resp.ToMap()
}

/*UserInfoByCode 从Code获取用户信息 */
func (a *Auth) UserInfoByCode(code, encrypted, iv string) []byte {
	p := a.Session(code)
	if !p.Has("errcode") {
		return a.UserInfo(p.GetString("session_key"), encrypted, iv)
	}
	return nil
}

/*UserInfo 用户信息 */
func (a *Auth) UserInfo(key, encrypted, iv string) []byte {
	r, e := a.dc.Decrypt(encrypted, iv, key)
	if e != nil {
		log.Error(e)
		return nil
	}
	return r
}
