package official

import (
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"time"
)

// JSSDK ...
type JSSDK struct {
	*Account
	url string
}

// URL ...
func (J *JSSDK) URL() string {
	return J.url
}

// SetURL ...
func (J *JSSDK) SetURL(url string) {
	J.url = url
}

func newJSSDK(account *Account) *JSSDK {
	return &JSSDK{
		Account: account,
	}
}

// NewJSSDK jssdk
func NewJSSDK(config *core.Config) *JSSDK {
	return newJSSDK(NewOfficialAccount(config))
}

// BuildConfig ...
func (J *JSSDK) BuildConfig(s []string, flags ...bool) util.Map {
	ticket := J.GetTicket()
	nonce := util.GenerateNonceStr()
	ts := util.Time()
	url := J.URL()
	m := util.Map{
		"appId":     J.Get("app_id"),
		"nonceStr":  nonce,
		"timestamp": ts,
		"url":       url,
		"jsApiList": s,
		"signature": getTicketSignature(ticket, nonce, ts, url),
	}
	return m
}

// GetTicket ...
func (J *JSSDK) GetTicket() string {
	if cache.Has(J.getCacheKey()) {
		maps := cache.Get(J.getCacheKey()).(util.Map)
		return maps.GetString("ticket")
	}

	resp := J.Account.Ticket().Get("jsapi")
	if resp.Error() != nil {
		return ""
	}
	m := resp.ToMap()
	ticket := m.GetString("ticket")
	expires, b := m.GetInt64("expires_in")
	if !b {
		return ""
	}
	t := time.Unix(expires, 0)
	cache.SetWithTTL(J.getCacheKey(), ticket, &t)
	return ticket

}

func (J *JSSDK) getCacheKey() string {
	return "godcong.wego.official.account.jssdk.ticket.jsapi" + J.GetString("app_id")
}

func getTicketSignature(ticket, nonce, ts, url string) string {
	return util.SHA1(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, ts, url))
}
