package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Base 基本功能组件 */
type Base struct {
	*Account
}

func newBase(account *Account) *Base {
	return &Base{
		Account: account,
	}
}

// NewBase 基础库
func NewBase(config *core.Config) *Base {
	return newBase(NewOfficialAccount(config))
}

/*ClearQuota 公众号的所有api调用（包括第三方帮其调用）次数进行清零
HTTP请求方式:POST
HTTP调用: https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
*/
func (b *Base) ClearQuota() core.Response {
	token := b.accessToken.GetToken()

	params := util.Map{
		"appid": b.Get("app_id"),
	}
	return b.client.PostJSON(Link(clearQuotaURLSuffix), token.KeyMap(), params)

}

/*GetCallbackIP 请求微信的服务器IP列表
HTTP请求方式: GET
HTTP调用:https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
*/
func (b *Base) GetCallbackIP() core.Response {
	token := b.accessToken.GetToken()
	return b.client.Get(Link(getCallbackIPURLSuffix), token.KeyMap())
}
