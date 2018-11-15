package core

import (
	"github.com/godcong/wego/log"
)

/*Ticket Ticket */
type Ticket struct {
	*Config
	client      *Client
	accessToken *AccessToken
}

// Client ...
func (t *Ticket) Client() *Client {
	return t.client
}

// SetClient ...
func (t *Ticket) SetClient(client *Client) {
	t.client = client
}

// AccessToken ...
func (t *Ticket) AccessToken() *AccessToken {
	return t.accessToken
}

// SetAccessToken ...
func (t *Ticket) SetAccessToken(accessToken *AccessToken) {
	t.accessToken = accessToken
}

func newTicket(config *Config) *Ticket {
	return &Ticket{
		Config: config,
	}
}

/*NewTicket NewTicket */
func NewTicket(config *Config, v ...interface{}) *Ticket {
	client := ClientGet(v)
	accessToken := newAccessToken(ClientCredential(config))
	accessToken.SetClient(client)

	ticket := newTicket(config)
	ticket.SetClient(client)
	ticket.SetAccessToken(accessToken)

	return ticket
}

//Get 获取api_ticket
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=wx_card
func (t *Ticket) Get(typ string, refresh bool) Response {
	log.Debug("Ticket|Get", typ)
	p := t.accessToken.GetToken().KeyMap()
	p.Set("type", typ)
	resp := t.client.Get(
		Connect(APIWeixin, ticketGetTicket),
		p)
	return resp
}
