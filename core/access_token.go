package core

import (
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/godcong/wego/cache"
)

type Token map[string]interface{}

type AccessToken struct {
	Config
	Client
	credentials Map
	token       string
}

const ACCESS_TOKEN_KEY = "access_token"
const ACCESS_TOKEN_EXPIRES_IN = "expires_in"

const ACCESS_TOKEN_SAFE_SECONDS = 500

func (a *AccessToken) getQuery() Map {
	panic("implement me")
}

func (a *AccessToken) sendRequest(s string) []byte {
	m0 := Map{
		"grant_type": "client_credential",
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
	}

	m := a.Request(API_WEIXIN_URL_SUFFIX+CGI_BIN_TOKEN_URL_SUFFIX+"?"+m0.UrlEncode(), nil, "get", nil)
	return m.ToJson()
}

func NewAccessToken(config Config, client Client) *AccessToken {
	return &AccessToken{
		Config: config,
		Client: client,
	}
}

func (a *AccessToken) Refresh() *AccessToken {
	a.getToken(true)
	return a
}

func (a *AccessToken) GetRefreshedToken() Token {
	return a.getToken(true)
}

func (a *AccessToken) GetToken() Token {
	return a.getToken(false)
}

func (a *AccessToken) GetTokenWithRefresh() Token {
	return a.getToken(true)
}

func (a *AccessToken) getToken(refresh bool) Token {
	key := a.getCacheKey()
	cache := cache.GetCache()

	if !refresh && cache.Has(key) {
		if v, b := cache.Get(key).(Token); b {
			return v
		}
	}

	token := a.RequestToken(a.getCredentials())
	if v := token.GetExpiresIn(); v != -1 {
		a.SetTokenWithLife(token.GetKey(), int(v))
	} else {
		a.SetToken(token.GetKey())
	}

	return token

}
func (a *AccessToken) RequestToken(credentials string) Token {
	response := a.sendRequest(credentials)
	m := Token{}
	json.Unmarshal(response, &m)
	return m
}

func (a *AccessToken) SetTokenWithLife(token string, lifeTime int) *AccessToken {
	return a.setToken(token, lifeTime)
}

func (a *AccessToken) SetToken(token string) *AccessToken {
	return a.setToken(token, 7200)
}

func (a *AccessToken) setToken(token string, lifeTime int) *AccessToken {
	cache.GetCache().SetWithTTL(a.getCacheKey(), Token{
		ACCESS_TOKEN_KEY: token,
		"expires_in":     lifeTime,
	}, lifeTime-ACCESS_TOKEN_SAFE_SECONDS)
	return a
}

func (a *AccessToken) getCredentials() string {
	c := md5.Sum(a.credentials.ToJson())
	return fmt.Sprintf("%x", c[:])
}

func (a *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + a.getCredentials()
}

func (t *Token) SetKey(s string) *Token {
	(*t)[ACCESS_TOKEN_KEY] = s
	return t
}

func (t *Token) GetKey() string {
	if v, b := (*t)[ACCESS_TOKEN_KEY]; b {
		return v.(string)
	}
	return ""
}

func (t *Token) SetExpiresIn(i int64) *Token {
	(*t)[ACCESS_TOKEN_EXPIRES_IN] = i
	return t
}

func (t *Token) GetExpiresIn() int64 {
	if i, b := (*t)[ACCESS_TOKEN_EXPIRES_IN]; b {
		return ParseInt(i)
	}
	return -1
}

func (t *Token) ToJson() string {
	v, e := json.Marshal(*t)
	if e != nil {
		return ""
	}
	return string(v)
}
