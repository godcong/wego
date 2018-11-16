package payment

import (
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/log"
	"strings"
	"time"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*Payment
	jssdk *core.JSSDK
}

func newJSSDK(p *Payment) interface{} {
	return &JSSDK{
		Payment: p,
	}
}

/*NewJSSDK NewJSSDK */
func NewJSSDK(config *core.Config) *JSSDK {
	return newJSSDK(NewPayment(config)).(*JSSDK)
}

func (J *JSSDK) getURL() string {
	return core.GetServerIP()
}

/*BridgeConfig bridge 设置 */
func (J *JSSDK) BridgeConfig(pid string) util.Map {
	appID := J.DeepGet("sub_appid", "app_id")

	m := util.Map{
		"appId":     appID,
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"package":   strings.Join([]string{"prepay_id", pid}, "="),
		"signType":  "MD5",
	}

	m.Set("paySign", util.GenerateSignature(m, J.GetString("key"), util.MakeSignMD5))

	return m
}

/*SdkConfig sdk 设置 */
func (J *JSSDK) SdkConfig(pid string) util.Map {
	config := J.BridgeConfig(pid)

	config.Set("timestamp", config.Get("timeStamp"))
	config.Delete("timeStamp")

	return config
}

/*AppConfig app 设置 */
func (J *JSSDK) AppConfig(pid string) util.Map {
	m := util.Map{
		"appid":     J.Get("app_id"),
		"partnerid": J.Get("mch_id"),
		"prepayid":  pid,
		"noncestr":  util.GenerateNonceStr(),
		"timestamp": util.Time(),
		"package":   "Sign=WXPay",
	}

	m.Set("sign", util.GenerateSignature(m, J.GetString("aes_key"), util.MakeSignMD5))
	return m
}

// ShareAddressConfig ...
//参数:token
//类型:string或*core.AccessToken
func (J *JSSDK) ShareAddressConfig(v interface{}) util.Map {
	token := ""
	switch vv := v.(type) {
	case *core.AccessToken:
		token = vv.GetToken().ToJSON()
	case string:
		token = vv
	}

	m := util.Map{
		"appId":     J.Get("app_id"),
		"scope":     "jsapi_address",
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"signType":  "SHA1",
	}

	signMsg := util.Map{
		"appid":       m.Get("appId"),
		"url":         J.getURL(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": token,
	}

	m.Set("addrSign", util.SHA1(signMsg.URLEncode()))

	return m
}

// BuildConfig ...
func (J *JSSDK) BuildConfig(maps util.Map) util.Map {
	ticket := J.GetTicket()
	nonce := util.GenerateNonceStr()
	ts := util.Time()
	url := maps.GetString("url")
	m := util.Map{
		"appId":     J.Get("appId"),
		"nonceStr":  nonce,
		"timestamp": ts,
		"url":       url,
		"jsApiList": maps.Get("jsApiList"),
		"signature": getTicketSignature(ticket, nonce, ts, url),
	}
	return m
}

// GetTicket ...
func (J *JSSDK) GetTicket() string {
	if cache.Has(J.getCacheKey()) {
		return cache.Get(J.getCacheKey()).(string)
	}

	resp := J.jssdk.Ticket().Get("jsapi", false)
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
	cache.SetWithTTL(J.getCacheKey(), ticket, &t)
	return ticket

}

func (J *JSSDK) getCacheKey() string {
	return "godcong.wego.official.account.jssdk.ticket.jsapi" + J.GetString("app_id")
}

func getTicketSignature(ticket, nonce, ts, url string) string {
	return util.SHA1(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, ts, url))
}
