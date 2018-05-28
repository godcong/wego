package core

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

type AccessToken struct {
	config.Config
	client      *Client
	credentials util.Map
	token       string
}

const ACCESS_TOKEN_KEY = "access_token"
const ACCESS_TOKEN_EXPIRES_IN = "expires_in"

const ACCESS_TOKEN_SAFE_SECONDS = 500

func (a *AccessToken) getQuery() util.Map {
	panic("implement me")
}

func (a *AccessToken) sendRequest(s string) []byte {
	m := util.Map{
		"grant_type": "client_credential",
		"appid":      a.Get("app_id"),
		"secret":     a.Get("secret"),
	}

	resp := a.client.HttpGet(a.client.Link(CGI_BIN_TOKEN_URL_SUFFIX), m)
	log.Debug("AccessToken|sendRequest", resp.ToString())
	if resp.CheckError() != nil {
		return nil
	}
	return resp.ToBytes()

}

func NewAccessToken(config config.Config, client *Client) *AccessToken {
	return &AccessToken{
		Config: config,
		client: client,
	}
}

func (a *AccessToken) Refresh() *AccessToken {
	log.Debug("AccessToken|Refresh")
	a.getToken(true)
	return a
}

func (a *AccessToken) GetRefreshedToken() *Token {
	log.Debug("AccessToken|GetRefreshedToken")
	return a.getToken(true)
}

func (a *AccessToken) GetToken() *Token {
	log.Debug("AccessToken|GetToken")
	return a.getToken(false)
}

func (a *AccessToken) GetTokenWithRefresh() *Token {
	log.Debug("AccessToken|GetTokenWithRefresh")
	return a.getToken(true)
}

func (a *AccessToken) getToken(refresh bool) *Token {
	key := a.getCacheKey()
	cache := cache.GetCache()

	if !refresh && cache.Has(key) {
		log.Debug("cached token", key)
		if v, b := cache.Get(key).(*Token); b {
			if v.ExpiresIn > time.Now().Unix() {
				return v
			}
		}
	}

	token := a.RequestToken(a.getCredentials())
	if token == nil {
		return nil
	}
	log.Debug("AccessToken|getToken", token)
	if v := token.ExpiresIn; v != 0 {
		a.SetTokenWithLife(token.AccessToken, time.Unix(v, 0))
	} else {
		a.SetToken(token.AccessToken)
	}

	return token

}
func (a *AccessToken) RequestToken(credentials string) *Token {
	var token Token
	tokenByte := a.sendRequest(credentials)
	if tokenByte == nil {
		return nil
	}

	err := json.Unmarshal(tokenByte, &token)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &token
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
		ExpiresIn:   time.Now().Unix() + lifeTime.Unix(),
	}, time.Now())
	return a
}

func (a *AccessToken) getCredentials() string {
	c := md5.Sum(a.credentials.ToJson())
	return fmt.Sprintf("%x", c[:])
}

func (a *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + a.getCredentials()
}
