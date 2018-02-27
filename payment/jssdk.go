package payment

import (
	"strings"

	"github.com/godcong/wego/core"
)

type JSSDK struct {
	core.Config
	*Payment
}

func (j *JSSDK) getUrl() string {
	return core.GetServerIp()
}

func (j *JSSDK) BridgeConfig(pid string) core.Map {
	appid := j.Get("sub_appid")
	if appid == "" {
		appid = j.Get("app_id")
	}

	m := core.Map{
		"appId":     appid,
		"timeStamp": core.Time(),
		"nonceStr":  core.GenerateNonceStr(),
		"package":   strings.Join([]string{"prepay_id", pid}, "="),
		"signType":  "MD5",
	}

	m.Set("paySign", core.GenerateSignature(m, j.Get("aes_key"), core.SIGN_TYPE_MD5))

	return m
}

func (j *JSSDK) SdkConfig(pid string) core.Map {
	config := j.BridgeConfig(pid)

	config.Set("timestamp", config.Get("timeStamp"))
	config.Delete("timeStamp")

	return config
}

func (j *JSSDK) AppConfig(pid string) core.Map {
	m := core.Map{
		"appid":     j.Get("app_id"),
		"partnerid": j.Get("mch_id"),
		"prepayid":  pid,
		"noncestr":  core.GenerateNonceStr(),
		"timestamp": core.Time(),
		"package":   "Sign=WXPay",
	}

	m.Set("sign", core.GenerateSignature(m, j.Get("aes_key"), core.SIGN_TYPE_MD5))
	return m
}

func (j *JSSDK) ShareAddressConfig(accessToken interface{}) core.Map {
	token := ""
	switch v := accessToken.(type) {
	case core.AccessToken:
		token0 := v.GetToken()
		token = token0.ToJson()
	case string:
		token = accessToken.(string)
	}
	m := core.Map{
		"appId":     j.Get("app_id"),
		"scope":     "jsapi_address",
		"timeStamp": core.Time(),
		"nonceStr":  core.GenerateNonceStr(),
		"signType":  "SHA1",
	}

	sm := core.Map{
		"appid":       m.Get("appId"),
		"url":         j.getUrl(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": token,
	}

	m.Set("addrSign", core.SHA1(sm.UrlEncode()))

	return m
}
