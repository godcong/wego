package official_account

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/config"
	"github.com/godcong/wego/core/log"
	"github.com/godcong/wego/core/net"
	"github.com/godcong/wego/core/util"
)

type Ticket struct {
	config.Config
	*OfficialAccount
}

func newTicket(officialAccount *OfficialAccount) *Ticket {
	return &Ticket{
		Config:          defaultConfig,
		OfficialAccount: officialAccount,
	}
}

func NewTicket() *Ticket {
	return newTicket(account)
}

// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card
// 成功:
// {"errcode":0,"errmsg":"ok","ticket":"9KwiourQPRN3vx3Nn1c_iX9qGaI3Cf8dwVy7qqYeYKcd3BK4Zd_jSlol7E7baUfgOY0E2ybaw2OrlhkChKaS7w","expires_in":7200}

func (t *Ticket) Get(typ string) *core.Response {
	log.Debug("Ticket|Get", typ)
	p := t.token.GetToken().KeyMap()
	p.Set("type", typ)
	resp := t.client.HttpGet(
		t.client.Link(GETTICKET_URL_SUFFIX),
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		},
	)
	return resp
}
