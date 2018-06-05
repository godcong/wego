package official

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
)

/*Ticket Ticket */
type Ticket struct {
	config.Config
	*Account
}

func newTicket(officialAccount *Account) *Ticket {
	return &Ticket{
		Config:  defaultConfig,
		Account: officialAccount,
	}
}

/*NewTicket NewTicket */
func NewTicket() *Ticket {
	return newTicket(account)
}

//Get 获取api_ticket
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card
// 成功:
// {"errcode":0,"errmsg":"ok","ticket":"9KwiourQPRN3vx3Nn1c_iX9qGaI3Cf8dwVy7qqYeYKcd3BK4Zd_jSlol7E7baUfgOY0E2ybaw2OrlhkChKaS7w","expires_in":7200}
func (t *Ticket) Get(typ string) *net.Response {
	log.Debug("Ticket|Get", typ)
	p := t.token.GetToken().KeyMap()
	p.Set("type", typ)
	resp := t.client.HTTPGet(
		t.client.Link(getTicketURLSuffix),
		p)
	return resp
}
