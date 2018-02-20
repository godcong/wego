package wego

import "strings"

type JSSDK interface {
	BridgeConfig(pid string) Map
	SdkConfig(pid string) Map
	AppConfig(pid string) Map
	ShareAddressConfig(accessToken interface{}) Map
}

type jssdk struct {
	Config
	app Application
}

func (j *jssdk) getUrl() string {
	return "http://y11e.com"
}

func (j *jssdk) BridgeConfig(pid string) Map {
	appid := j.Get("sub_appid")
	if appid == "" {
		appid = j.Get("app_id")
	}

	m := Map{
		"appId":     appid,
		"timeStamp": Time(),
		"nonceStr":  GenerateNonceStr(),
		"package":   strings.Join([]string{"prepay_id", pid}, "="),
		"signType":  "MD5",
	}

	m.Set("paySign", GenerateSignature(m, j.Get("aes_key"), SIGN_TYPE_MD5))

	return m
}

func (j *jssdk) SdkConfig(pid string) Map {
	config := j.BridgeConfig(pid)

	config.Set("timestamp", config.Get("timeStamp"))
	config.Delete("timeStamp")

	return config
}

func (j *jssdk) AppConfig(pid string) Map {
	m := Map{
		"appid":     j.Get("app_id"),
		"partnerid": j.Get("mch_id"),
		"prepayid":  pid,
		"noncestr":  GenerateNonceStr(),
		"timestamp": Time(),
		"package":   "Sign=WXPay",
	}

	m.Set("sign", GenerateSignature(m, j.Get("aes_key"), SIGN_TYPE_MD5))
	return m
}

func (j *jssdk) ShareAddressConfig(accessToken interface{}) Map {
	token := ""
	switch accessToken.(type) {
	case AccessTokenInterface:
		token = accessToken.(AccessTokenInterface).GetToken()
	case string:
		token = accessToken.(string)
	}
	m := Map{
		"appId":     j.Get("app_id"),
		"scope":     "jsapi_address",
		"timeStamp": Time(),
		"nonceStr":  GenerateNonceStr(),
		"signType":  "SHA1",
	}

	sm := Map{
		"appid": m.Get("appId"),
		//"url" : $this->getUrl(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": token,
	}

	sm.SortKeys()
	sm.ToUrlQuery()
	m.Set("addrSign", SHA1(sm.ToSortUrlQuery()))

	return m
}

func NewJSSDK(application Application, config Config) JSSDK {
	return &jssdk{
		Config: config,
		app:    application,
		//client: application.Client(),
	}
}
