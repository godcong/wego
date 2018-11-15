package core

import (
	"github.com/godcong/wego/util"
)

/*Base 基础 */
type Base struct {
	*Config
	accessToken *AccessToken
}

// AccessToken ...
func (b *Base) AccessToken() *AccessToken {
	return b.accessToken
}

// SetAccessToken ...
func (b *Base) SetAccessToken(accessToken *AccessToken) {
	b.accessToken = accessToken
}

func newBase(config *Config) *Base {
	return &Base{
		Config: config,
	}
}

//NewBase NewBase
//Deprecated: Base is moved to official_account.Base
func NewBase(config *Config, v ...interface{}) *Base {
	accessToken := newAccessToken(ClientCredential(config))

	base := newBase(config)
	base.SetAccessToken(accessToken)

	return base
}

//ClearQuota  公众号的所有api调用（包括第三方帮其调用）次数进行清零
//Deprecated: ClearQuota is moved to official_account.Base
//公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零:
//HTTP请求:POST HTTP调用: https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
func (b *Base) ClearQuota() Responder {
	token := b.accessToken.GetToken()

	params := util.Map{"appid": b.GetString("app_id")}

	return PostJSON(APIWeixin+clearQuotaURLSuffix, token.KeyMap(), params)

}

//GetCallbackIP 请求微信的服务器IP列表
//Deprecated: GetCallbackIP is moved to official_account.Base
//接口调用请求说明
//http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
func (b *Base) GetCallbackIP() Responder {
	token := b.accessToken.GetToken()
	return Get(APIWeixin+getCallbackIPURLSuffix, token.KeyMap())
}
