package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
)

/*Ticket Ticket */
type Ticket struct {
	*Account
}

func newTicket(acc *Account) *Ticket {
	return &Ticket{
		Account: acc,
	}
}

/*NewTicket NewTicket */
func NewTicket(config *core.Config) *Ticket {
	return newTicket(NewAccount(config))
}

//Get 获取api_ticket
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card
// 成功:
// {"errcode":0,"errmsg":"ok","ticket":"9KwiourQPRN3vx3Nn1c_iX9qGaI3Cf8dwVy7qqYeYKcd3BK4Zd_jSlol7E7baUfgOY0E2ybaw2OrlhkChKaS7w","expires_in":7200}
func (t *Ticket) Get(typ string) core.Response {
	log.Debug("Ticket|Get", typ)
	p := t.accessToken.GetToken().KeyMap()
	p.Set("type", typ)
	resp := t.client.Get(
		Link(getTicketURLSuffix),
		p)
	return resp
}
