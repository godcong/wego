package official

import (
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"time"
)

// JSSDK ...
type JSSDK struct {
	*Account
	url string
}

// URL ...
func (j *JSSDK) URL() string {
	return j.url
}

// SetURL ...
func (j *JSSDK) SetURL(url string) {
	j.url = url
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
func (j *JSSDK) BuildConfig(maps util.Map) util.Map {
	ticket := j.GetTicket()
	nonce := util.GenerateNonceStr()
	ts := util.Time()
	url := j.URL()
	m := util.Map{
		"appId":     j.Get("appId"),
		"nonceStr":  nonce,
		"timestamp": ts,
		"url":       url,
		"jsApiList": maps.Get("jsApiList"),
		"signature": getTicketSignature(ticket, nonce, ts, url),
	}
	return m
}

// GetTicket ...
func (j *JSSDK) GetTicket() string {
	if cache.Has(j.getCacheKey()) {
		return cache.Get(j.getCacheKey()).(string)
	}

	resp := j.Account.Ticket().Get("jsapi")
	if resp.Error() != nil {
		return ""
	}
	m := resp.ToMap()
	ticket := m.GetString("ticket")
	expires, b := m.GetInt64("expires_in")

	log.Println(expires, b)
	if !b {
		return ""
	}
	t := time.Unix(time.Now().Unix()+expires-500, 0)
	cache.SetWithTTL(j.getCacheKey(), ticket, &t)
	return ticket

}

func (j *JSSDK) getCacheKey() string {
	return "godcong.wego.official.account.jssdk.ticket.jsapi" + j.GetString("app_id")
}

func getTicketSignature(ticket, nonce, ts, url string) string {
	return util.SHA1(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, ts, url))
}
