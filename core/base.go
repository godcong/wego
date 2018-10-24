package core

import (
	"github.com/godcong/wego/util"
)

/*Base 基础 */
type Base struct {
	config *Config
	client *Client
	token  *AccessToken
}

func newBase(config *Config) *Base {
	client := NewClient(config)
	return &Base{
		config: config,
		client: client,
		token:  NewAccessToken(config),
	}
}

//NewBase NewBase
func NewBase(config *Config) *Base {
	return newBase(config)
}

//ClearQuota  公众号的所有api调用（包括第三方帮其调用）次数进行清零
//公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零：
//HTTP请求：POST HTTP调用： https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
func (b *Base) ClearQuota() Response {
	token := b.token.GetToken()

	params := util.Map{
		"appid": b.config.Get("app_id"),
	}
	return b.client.PostJSON(APIWeixin+clearQuotaURLSuffix, params, util.Map{
		DataTypeQuery: token.KeyMap(),
	})

}

//GetCallbackIP 请求微信的服务器IP列表
//接口调用请求说明
//http请求方式: GEThttps://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
func (b *Base) GetCallbackIP() Response {
	token := b.token.GetToken()
	b.client.SetRequestType(DataTypeJSON)
	return b.client.Get(APIWeixin+getCallbackIPURLSuffix, token.KeyMap())
}
