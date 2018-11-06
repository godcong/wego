package official

import (
	"github.com/godcong/wego/core"
)

// Current ...
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

//AutoReplyInfo
//http请求方式: GET（请使用https协议）
//https://api.weixin.qq.com/cgi-bin/get_current_autoreply_info?access_token=ACCESS_TOKEN
func (c *Current) AutoReplyInfo() core.Response {
	token := c.accessToken.GetToken()
	return c.client.Get(Link(getCurrentAutoReplyInfo), token.KeyMap())
}

//SelfMenuInfo
//http请求方式: GET（请使用https协议）
//https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN
func (c *Current) SelfMenuInfo() core.Response {
	token := c.accessToken.GetToken()
	return c.client.Get(Link(getCurrentSelfMenuInfo), token.KeyMap())
}
