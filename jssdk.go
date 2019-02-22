package wego

import (
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/util"
	"strings"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*JSSDKConfig
	accessToken *AccessToken
	ticket      *Ticket
	URL         string
	CacheKey    func() string
}

// Ticket ...
func (j *JSSDK) Ticket() *Ticket {
	return j.ticket
}

// SetTicket ...
func (j *JSSDK) SetTicket(ticket *Ticket) {
	j.ticket = ticket
}

func newJSSDK(config *Config) interface{} {
	jssdk := &JSSDK{
		Config: config,
		ticket: nil,
	}
	jssdk.CacheKey = jssdk.getCacheKey
	return jssdk
}

/*NewJSSDK NewJSSDK */
func NewJSSDK(config *JSSDKConfig) *JSSDK {
	jssdk := &JSSDK{
		JSSDKConfig: config,
		accessToken: NewAccessToken(config.AccessToken),
		ticket:      NewTicket(),
	}
}

func (j *JSSDK) getURL() string {
	return GetServerIP()
}

/*BridgeConfig bridge 设置 */
func (j *JSSDK) BridgeConfig(pid string) util.Map {
	appID := j.DeepGet("sub_appid", "app_id")

	m := util.Map{
		"appId":     appID,
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"package":   strings.Join([]string{"prepay_id", pid}, "="),
		"signType":  "MD5",
	}

	m.Set("paySign", util.GenSign(m, j.GetString("key")))

	return m
}

/*SdkConfig sdk 设置 */
func (j *JSSDK) SdkConfig(pid string) util.Map {
	config := j.BridgeConfig(pid)

	config.Set("timestamp", config.Get("timeStamp"))
	config.Delete("timeStamp")

	return config
}

/*AppConfig app 设置 */
func (j *JSSDK) AppConfig(pid string) util.Map {
	m := util.Map{
		"appid":     j.Get("app_id"),
		"partnerid": j.Get("mch_id"),
		"prepayid":  pid,
		"noncestr":  util.GenerateNonceStr(),
		"timestamp": util.Time(),
		"package":   "Sign=WXPay",
	}

	m.Set("sign", util.GenSign(m, j.GetString("aes_key")))
	return m
}

// ShareAddressConfig ...
//	参数:token
//	类型:string或*core.AccessToken
func (j *JSSDK) ShareAddressConfig(v interface{}) util.Map {
	token := ""
	switch vv := v.(type) {
	case *AccessToken:
		token = vv.GetToken().ToJSON()
	case string:
		token = vv
	}

	m := util.Map{
		"appId":     j.Get("app_id"),
		"scope":     "jsapi_address",
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"signType":  "SHA1",
	}

	signMsg := util.Map{
		"appid":       m.Get("appId"),
		"url":         j.getURL(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": token,
	}

	m.Set("addrSign", util.GenSHA1(signMsg.URLEncode()))

	return m
}

// BuildConfig ...
func (j *JSSDK) BuildConfig(maps util.Map) util.Map {
	ticket := j.GetTicket("jsapi", false)
	nonce := util.GenerateNonceStr()
	ts := util.Time()
	url := maps.GetStringD("url", j.URL)
	appID := maps.GetStringD("app_id", j.GetString("app_id"))
	m := util.Map{
		"appID":     appID,
		"nonceStr":  nonce,
		"timestamp": ts,
		"url":       url,
		"jsApiList": maps.Get("jsApiList"),
		"signature": getTicketSignature(ticket, nonce, ts, url),
	}
	return m
}

// GetTicket ...
func (j *JSSDK) GetTicket(genre string, refresh bool) string {
	if !refresh && cache.Has(j.getCacheKey()) {
		return cache.Get(j.getCacheKey()).(string)
	}

	resp := j.Ticket().Get(genre)
	if resp.Error() != nil {
		return ""
	}
	m := resp.ToMap()
	ticket := m.GetString("ticket")
	expires, b := m.GetInt64("expires_in")
	log.Debug(ticket, expires, b)
	if !b {
		return ""
	}

	cache.SetWithTTL(j.getCacheKey(), ticket, expires-500)
	return ticket

}

func (j *JSSDK) getCacheKey() string {
	return "godcong.wego.jssdk.ticket.jsapi" + j.GetString("app_id")
}

func getTicketSignature(ticket, nonce, ts, url string) string {
	return util.GenSHA1(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, ts, url))
}
