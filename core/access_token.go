package core

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godcong/wego/cache"
)

type AccessToken struct {
	Config
	client      *Client
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
	//m := a.client.Request(a.client.Link(CGI_BIN_TOKEN_URL_SUFFIX), nil, "get", Map{
	//	REQUEST_TYPE_QUERY.String(): m0,
	//})
	m := a.client.HttpGet(a.client.Link(CGI_BIN_TOKEN_URL_SUFFIX), m0)

	Debug("AccessToken|sendRequest", m.ToString())
	return m.ToBytes()
}

func NewAccessToken(config Config, client *Client) *AccessToken {
	return &AccessToken{
		Config: config,
		client: client,
	}
}

func (a *AccessToken) Refresh() *AccessToken {
	Debug("AccessToken|Refresh")
	a.getToken(true)
	return a
}

func (a *AccessToken) GetRefreshedToken() *Token {
	Debug("AccessToken|GetRefreshedToken")
	return a.getToken(true)
}

func (a *AccessToken) GetToken() *Token {
	Debug("AccessToken|GetToken")
	return a.getToken(false)
}

func (a *AccessToken) GetTokenWithRefresh() *Token {
	Debug("AccessToken|GetTokenWithRefresh")
	return a.getToken(true)
}

func (a *AccessToken) getToken(refresh bool) *Token {

	key := a.getCacheKey()
	cache := cache.GetCache()

	if !refresh && cache.Has(key) {
		if v, b := cache.Get(key).(Token); b {
			return &v
		}
	}

	token := a.RequestToken(a.getCredentials())
	Debug("AccessToken|getToken", token)
	if v := token.ExpiresIn; v != 0 {
		a.SetTokenWithLife(token.AccessToken, time.Unix(v, 0))
	} else {
		a.SetToken(token.AccessToken)
	}

	return &token

}
func (a *AccessToken) RequestToken(credentials string) Token {
	response := a.sendRequest(credentials)
	m := Token{}
	err := json.Unmarshal(response, &m)
	if err != nil {
		Error(err)
	}
	return m
}

func (a *AccessToken) SetTokenWithLife(token string, lifeTime time.Time) *AccessToken {
	return a.setToken(token, lifeTime)
}

func (a *AccessToken) SetToken(token string) *AccessToken {
	return a.setToken(token, time.Unix(7200, 0))
}

func (a *AccessToken) setToken(token string, lifeTime time.Time) *AccessToken {
	cache.GetCache().SetWithTTL(a.getCacheKey(), &Token{
		AccessToken: token,
		ExpiresIn:   lifeTime.Unix(),
	}, lifeTime.Add(time.Duration(-ACCESS_TOKEN_SAFE_SECONDS)))
	return a
}

func (a *AccessToken) getCredentials() string {
	c := md5.Sum(a.credentials.ToJson())
	return fmt.Sprintf("%x", c[:])
}

func (a *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + a.getCredentials()
}
