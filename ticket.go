package wego

import (
	"github.com/godcong/wego/util"
)

// TicketRes ticket response data
type TicketRes struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

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
	return GetTicket(p)
}

// GetTicketRes ...
func (t *Ticket) GetTicketRes(s string) (*TicketRes, error) {
	var tr TicketRes
	ticket := t.Get(s)

	if err := ticket.Error(); err != nil {
		return nil, err
	}

	if err := ticket.Unmarshal(&tr); err != nil {
		return nil, err
	}

	return &tr, nil
}

//GetTicket 获取api_ticket
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card
// type: jsapi,wx_card
// access_token: #token#
func GetTicket(p util.Map) Responder {
	return Get(util.URL(apiWeixin, ticketGetTicket), p)
}
