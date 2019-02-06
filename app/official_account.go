package app

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

// OfficialAccount ...
type OfficialAccount struct {
	*core.Config
	accessToken *AccessToken
	prefix      string
}

/*ClearQuota 公众号的所有api调用（包括第三方帮其调用）次数进行清零
HTTP请求方式:POST
HTTP调用: https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
*/
func (a *OfficialAccount) ClearQuota() core.Responder {
	token := a.accessToken.GetToken()

	params := util.Map{
		"appid": a.Get("app_id"),
	}
	return core.PostJSON(util.URL(apiWeixin, clearQuota), token.KeyMap(), params)

}

/*GetCallbackIP 请求微信的服务器IP列表
HTTP请求方式: GET
HTTP调用:https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
*/
func (a *OfficialAccount) GetCallbackIP() core.Responder {
	token := a.accessToken.GetToken()
	return core.Get(util.URL(apiWeixin, getCallbackIP), token.KeyMap())
}
