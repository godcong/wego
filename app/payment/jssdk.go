package payment

import (
	"strings"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*Payment
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

func (j *JSSDK) getURL() string {
	return core.GetServerIP()
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

	m.Set("paySign", GenerateSignature(m, j.GetString("key"), MakeSignMD5))

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

	m.Set("sign", GenerateSignature(m, j.GetString("aes_key"), MakeSignMD5))
	return m
}

// ShareAddressConfig ...
//参数:token
//类型:string或*core.AccessToken
func (j *JSSDK) ShareAddressConfig(v interface{}) util.Map {
	token := ""
	switch vv := v.(type) {
	case *core.AccessToken:
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

	sm := util.Map{
		"appid":       m.Get("appId"),
		"url":         j.getURL(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": token,
	}

	m.Set("addrSign", util.SHA1(sm.URLEncode()))

	return m
}
