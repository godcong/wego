package official

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"

	"github.com/godcong/wego/util"
)

/*CallbackValue CallbackValue */
type CallbackValue struct {
	Type  string
	Value interface{}
}

/*CallbackFunc func(w http.ResponseWriter, r *http.Request, val *CallbackValue) []byte*/
type CallbackFunc func(w http.ResponseWriter, r *http.Request, val *CallbackValue) []byte

/*OAuth OAuth */
type OAuth struct {
	*Account
	//response    core.Response
	callback    map[string]CallbackFunc
	authorize   string
	scopes      []string
	redirectURI string
}

func newOAuth(acc *Account) *OAuth {
	oauth := &OAuth{
		Account:  acc,
		callback: map[string]CallbackFunc{},
	}
	//oauth.client.SetRequestType(core.DataTypeJSON)
	oauth.scopes = oauth.GetStringArrayD("oauth.scopes", []string{snsapiBase})
	oauth.redirectURI = oauth.GetStringD("oauth.redirect_uri", defaultOauthRedirectURLSuffix)
	oauth.authorize = oauth.GetStringD("oauth.authorize", oauth2AuthorizeURLSuffix)
	return oauth
}

/*NewOAuth NewOAuth*/
func NewOAuth(config *core.Config) *OAuth {
	return newOAuth(NewOfficialAccount(config))
}

/*RegisterCodeCallback RegisterCodeCallback*/
func (o *OAuth) RegisterCodeCallback(callbackFunc CallbackFunc) *OAuth {
	return o.registerCallback("code", callbackFunc)
}

/*RegisterStateCallback RegisterStateCallback*/
func (o *OAuth) RegisterStateCallback(callbackFunc CallbackFunc) *OAuth {
	return o.registerCallback("state", callbackFunc)
}

/*RegisterAllCallback RegisterAllCallback*/
func (o *OAuth) RegisterAllCallback(callbackFunc CallbackFunc) *OAuth {
	return o.registerCallback("all", callbackFunc)
}

/*RegisterCallback RegisterCallback*/
func (o *OAuth) RegisterCallback(callbackFunc CallbackFunc) *OAuth {
	return o.registerCallback("all", callbackFunc)
}

/*RegisterInfoCallback RegisterInfoCallback*/
func (o *OAuth) RegisterInfoCallback(callbackFunc CallbackFunc) *OAuth {
	return o.registerCallback("info", callbackFunc)
}

func (o *OAuth) registerCallback(name string, callbackFunc CallbackFunc) *OAuth {
	o.callback[name] = callbackFunc
	return o
}

// ServeHTTP 监听授权服务
// 失败:
// {"errcode":40163,"errmsg":"code been used, hints: [ req_id: OsIKda0848th19 ]"}
// {"errcode":40029,"errmsg":"invalid code, hints: [ req_id: 5u8NWa0990th40 ]"}
// 成功:
// {"access_token":"7_0MSpG_WEPVwQki6eFQSFQbRwkEkTEhkvBjkuKTODS7_xe6vBOEsc88kcCu_781YvXXP2FwWC4M5m-B9WXs51rA","expires_in":7200,"refresh_token":"7_51Axvh89ev5cGH-WR4qPKb-rcPf2VQrMg25MNDs1899cHYb5UomPi4fnc1NAks07Vw5Bb0pTFvvritU-aQtxFg","openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc","scope":"snsapi_base"}]
func (o *OAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := o.hookAccessToken(w, r)
	log.Debug("ServeHTTP|token", token)
	if token != nil {
		info := o.hookUserInfo(w, r, token)
		log.Debug("ServeHTTP|info", info)
		return
	}

	uri := o.hookState(w, r)
	log.Debug("ServeHTTP|uri", uri)
	http.Redirect(w, r, uri, http.StatusFound)
	return
}

func (o *OAuth) hookState(w http.ResponseWriter, r *http.Request) string {
	v := CallbackValue{Type: "info", Value: nil}

	if f, b := o.callback["state"]; b {
		if rlt := f(w, r, &v); rlt != nil {
			return o.AuthCodeURL(string(rlt))
		}
	}
	return o.AuthCodeURL("")
}

func (o *OAuth) hookUserInfo(w http.ResponseWriter, r *http.Request, token *core.Token) *core.UserInfo {
	info := o.UserInfo(token)
	v := CallbackValue{Type: "info", Value: info}
	if a, b := o.callback["all"]; b {
		if rlt := a(w, r, &v); rlt != nil {
			w.Write(rlt)
		}
	}
	if a, b := o.callback["info"]; b {
		if rlt := a(w, r, &v); rlt != nil {
			w.Write(rlt)
		}
	}
	return info
}

func (o *OAuth) hookAccessToken(w http.ResponseWriter, r *http.Request) *core.Token {
	r.ParseForm()
	if code := r.Form.Get("code"); code != "" {
		token := o.AccessToken(code)
		v := CallbackValue{Type: "code", Value: token}
		if a, b := o.callback["all"]; b {
			if rlt := a(w, r, &v); rlt != nil {
				w.Write(rlt)
			}
		}
		if a, b := o.callback["code"]; b {
			if rlt := a(w, r, &v); rlt != nil {
				w.Write(rlt)
			}
		}
		return token
	}
	return nil
}

/*AuthCodeURL 生成授权地址URL*/
func (o *OAuth) AuthCodeURL(state string) string {
	log.Debug("AuthCodeURL|Account", o.Account)
	var buf bytes.Buffer
	buf.WriteString(o.authorize)
	v := url.Values{
		"response_type": {"code"},
		"appid":         {o.GetString("app_id")},
	}
	if o.redirectURI != "" {
		log.Println(o.redirectURI)
		if strings.Index(o.redirectURI, "http") == 0 {
			v.Set("redirect_uri", o.redirectURI)
		} else {
			v.Set("redirect_uri", core.Link(o.redirectURI, "host"))
		}
	}
	if o.scopes != nil {
		v["scope"] = o.scopes
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		v.Set("state", state)
	}

	if !strings.Contains(o.authorize, "?") {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

//RefreshToken 刷新Token
// https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN
// 成功:
// {"openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc","access_token":"7_EVGE1V1XzagA0PXMPFUbLApiA4BCGO5oVSxkDRbZ-aiTfwpP32DSNxsdFBN0AuERGrEtCBuBfNzTpTv_mYi-NQ","expires_in":7200,"refresh_token":"7_XxwLIQsmfEHnuVsw91q8fK1WWRcq37z2-rTTlMjrouJussoQff77jE9043qtiIQMr8CJuBWc3hmMGONJbB_EQQ","scope":"snsapi_base,snsapi_userinfo,"}
func (o *OAuth) RefreshToken(refresh string) *core.Token {
	v := util.Map{
		"appid":         o.Get("app_id"),
		"grant_type":    "refresh_token",
		"refresh_token": refresh,
	}
	if o.redirectURI != "" {
		v.Set("redirect_uri", core.Link(o.redirectURI, "host"))
	}
	response := core.PostJSON(
		Link(oauth2RefreshTokenURLSuffix),
		v,
		nil,
	)
	log.Debug("AccessToken|response", string(response.Bytes()), response.Error())
	if response.Error() != nil {
		return nil
	}

	return unmarshalToken(response.Bytes())
}

/*AccessToken AccessToken*/
func (o *OAuth) AccessToken(code string) *core.Token {
	v := util.Map{
		"appid":      o.Get("app_id"),
		"secret":     o.Get("secret"),
		"code":       code,
		"grant_type": "authorization_code",
	}
	if o.redirectURI != "" {
		v.Set("redirect_uri", core.Link(o.redirectURI, "host"))
	}
	response := core.PostJSON(
		Link(oauth2AccessTokenURLSuffix),
		v,
		nil,
	)
	log.Debug("AccessToken|response", string(response.Bytes()), response.Error())
	if response.Error() != nil {
		return nil
	}

	return unmarshalToken(response.Bytes())
}

//UserInfo 用户信息
// http:GET（请使用https协议） https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
// lang: zh_CN 简体，zh_TW 繁体，en
// 成功:
// {"openid":"OPENID","nickname":NICKNAME,"sex":"1","province":"PROVINCE""city":"CITY","country":"COUNTRY","headimgurl":"http:thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46","privilege":["PRIVILEGE1""PRIVILEGE2"],"unionid":"o6_bmasdasdsad6_2sgVt7hMZOPfL"}
// 失败:
// {"errcode":41001,"errmsg":"access_token missing, hints: [ req_id: 8mfAmA0205s158 ]"}
func (o *OAuth) UserInfo(token *core.Token) *core.UserInfo {
	p := util.Map{
		"access_token": token.AccessToken,
		"openid":       token.OpenID,
		"lang":         "zh_CN",
	}
	response := core.Get(
		Link(oauth2UserinfoURLSuffix),
		p,
	)
	log.Debug("UserInfo|response", string(response.Bytes()), response.Error())
	if response.Error() != nil {
		return nil
	}

	return unmarshalUserInfo(response.Bytes())
}

//Validate 验证
// 成功:
// true
// 失败:
// false
func (o *OAuth) Validate(token *core.Token) bool {
	p := util.Map{
		"access_token": token.AccessToken,
		"openid":       token.OpenID,
	}
	response := core.Get(
		Link(oauth2AuthURLSuffix),
		p)
	log.Debug("Validate|response", string(response.Bytes()), response.Error())
	if response.Error() == nil {
		m := response.ToMap()
		v, b := m.GetInt64("errcode")
		if b && v == 0 &&
			m.GetString("errmsg") == "ok" {
			return true
		}
	}

	return false
}

func unmarshalToken(data []byte) *core.Token {
	var token core.Token
	err := json.Unmarshal(data, &token)
	if err != nil {
		return nil
	}
	return &token
}

func unmarshalUserInfo(data []byte) *core.UserInfo {
	var info core.UserInfo
	err := json.Unmarshal(data, &info)
	if err != nil {
		log.Debug(err)
		return nil
	}
	return &info
}
