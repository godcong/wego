package wego

import (
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
)

/*Ticket Ticket */
type Ticket struct {
	*AccessToken
}

/*NewTicket NewTicket */
func NewTicket(accessToken *AccessToken) *Ticket {
	return &Ticket{
		AccessToken: accessToken,
	}
}

//Get 获取api_ticket
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card
// type: jsapi,wx_card
func (t *Ticket) Get(s string) Responder {
	log.Debug("Ticket|Get", s)
	p := t.KeyMap().Set("type", s)
	return Get(util.URL(apiWeixin, ticketGetTicket), p)
}

// GetTicket ...
func GetTicket(p util.Map) Responder {
	return Get(util.URL(apiWeixin, ticketGetTicket), p)
}
