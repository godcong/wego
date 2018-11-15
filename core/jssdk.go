package core

import (
	"strings"

	"github.com/godcong/wego/util"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*Config
	ticket *Ticket
}

// Ticket ...
func (J *JSSDK) Ticket() *Ticket {
	return J.ticket
}

// SetTicket ...
func (J *JSSDK) SetTicket(ticket *Ticket) {
	J.ticket = ticket
}

func newJSSDK(config *Config) interface{} {
	return &JSSDK{
		Config: config,
	}
}

/*NewJSSDK NewJSSDK */
func NewJSSDK(config *Config) *JSSDK {
	jssdk := newJSSDK(config).(*JSSDK)
	jssdk.ticket = NewTicket(config)
	return jssdk
}

func (J *JSSDK) getURL() string {
	return GetServerIP()
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

	//m.Set("paySign", GenerateSignature(m, J.GetString("key"), MakeSignMD5))

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

	//m.Set("sign", GenerateSignature(m, J.GetString("aes_key"), MakeSignMD5))
	return m
}

// ShareAddressConfig ...
//参数:token
//类型:string或*core.AccessToken
func (J *JSSDK) ShareAddressConfig(v interface{}) util.Map {
	token := ""
	switch vv := v.(type) {
	case *AccessToken:
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
