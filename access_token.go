package wego

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"

	"github.com/godcong/wego/cache"
)

type AccessTokenInterface interface {
	GetToken() Token
	Refresh() AccessTokenInterface
	ApplyToRequest(RequestInterface, Map) RequestInterface
	//getCredentials() Map
	//getQuery() Map
	//sendRequest() []byte
}

type Token map[string]interface{}

type AccessToken struct {
	Config
	app         Application
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
	//https://api.weixin.qq.com/cgi-bin/token
	m := a.app.Client(a.Config).Request(CGI_BIN_TOKEN_URL_SUFFIX, Map{
		"json": s,
	}, "get", nil)
	return m.ToJson()
}

var acc AccessTokenInterface

func NewAccessToken(application Application, config Config) AccessTokenInterface {
	return &AccessToken{
		Config: config,
		app:    application,
	}
}

func (a *AccessToken) GetRefreshToken() string {
	panic("implement me")
}

func (a *AccessToken) Refresh() AccessTokenInterface {
	panic("implement me")
}

func (a *AccessToken) ApplyToRequest(RequestInterface, Map) RequestInterface {
	panic("implement me")
}

func (a *AccessToken) GetRefreshedToken(RequestInterface, Map) RequestInterface {
	panic("implement me")
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
		if v, b := cache.Get(key).(map[string]interface{}); b {
			return v
		}
	}

	token := a.RequestToken(a.getCredentials())
	if v := token.GetExpiresIn(); v != -1 {
		a.SetTokenWithLife(token.GetKey(), v)
	} else {
		a.SetToken(token.GetKey())
	}

	return token

}
func (a *AccessToken) RequestToken(credentials string) Token {
	response := a.sendRequest(credentials)
	log.Println(string(response))
	m := Token{}
	json.Unmarshal(response, &m)
	return m
}

func (a *AccessToken) SetTokenWithLife(token string, lifeTime int) AccessTokenInterface {
	return a.setToken(token, lifeTime)
}

func (a *AccessToken) SetToken(token string) AccessTokenInterface {
	return a.setToken(token, 7200)
}

func (a *AccessToken) setToken(token string, lifeTime int) AccessTokenInterface {
	cache.GetCache().SetWithTTL(a.getCacheKey(), map[string]interface{}{
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
	return "godcong.wego.access_token" + a.getCredentials()
}

func NewToken() *Token {
	return &Token{}
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

func (t *Token) SetExpiresIn(i int) *Token {
	(*t)[ACCESS_TOKEN_EXPIRES_IN] = i
	return t
}

func (t *Token) GetExpiresIn() int {
	if i, b := (*t)[ACCESS_TOKEN_EXPIRES_IN]; b {
		return i.(int)
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
