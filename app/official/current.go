package official

import (
	"github.com/godcong/wego/core"
)

type Current struct {
	*Account
}

func newCurrent(account *Account) *Current {
	return &Current{
		Account: account,
	}
}

//NewCurrent current
func NewCurrent(config *core.Config) *Current {
	return newCurrent(NewOfficialAccount(config))
}

//http请求方式: GET（请使用https协议）
//https://api.weixin.qq.com/cgi-bin/get_current_autoreply_info?access_token=ACCESS_TOKEN
func (c *Current) AutoReplyInfo() core.Response {
	token := c.accessToken.GetToken()

	return c.client.Get(Link(getCurrentAutoReplyInfo), token.KeyMap())
}
