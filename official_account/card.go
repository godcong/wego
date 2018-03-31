package official_account

import "github.com/godcong/wego/core"

type Card struct {
	core.Config
	*OfficialAccount
}

func newCard(officialAccount *OfficialAccount) *Card {
	return &Card{
		Config:          defaultConfig,
		OfficialAccount: officialAccount,
	}
}

func NewCard() *Card {
	return newCard(account)
}

// HTTP请求方式: POST
// URL:https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN
// func (c *Card) Create() {
// 	key := c.token.GetToken().KeyMap()
// 	resp := c.client.HttpPostJson(
// 		c.client.Link(url),
// 		core.Map{core.REQUEST_TYPE_QUERY.String(): key})
// 	return resp
// }
