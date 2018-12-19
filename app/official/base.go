package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*ClearQuota 公众号的所有api调用（包括第三方帮其调用）次数进行清零
HTTP请求方式:POST
HTTP调用: https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
*/
func (a *Account) ClearQuota() core.Responder {
	token := a.accessToken.GetToken()

	params := util.Map{
		"appid": a.Get("app_id"),
	}
	return core.PostJSON(Link(clearQuotaURLSuffix), token.KeyMap(), params)

}

/*GetCallbackIP 请求微信的服务器IP列表
HTTP请求方式: GET
HTTP调用:https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
*/
func (a *Account) GetCallbackIP() core.Responder {
	token := a.accessToken.GetToken()
	return core.Get(Link(getCallbackIPURLSuffix), token.KeyMap())
}
