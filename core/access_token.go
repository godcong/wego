package core

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*AccessToken AccessToken */
type AccessToken struct {
	config      *Config
	client      *Client
	credentials util.Map
}

/*AccessTokenKey 键值 */
const AccessTokenKey = "access_token"

/*AccessTokenExpiresIn 过期 */
const AccessTokenExpiresIn = "expires_in"

/*AccessTokenSafeSeconds token安全时间 */
const AccessTokenSafeSeconds = 500

func (a *AccessToken) getQuery() util.Map {
	panic("implement me")
}

func (a *AccessToken) sendRequest(s string) []byte {
	m := util.Map{
		"grant_type": "client_credential",
		"appid":      a.config.Get("app_id"),
		"secret":     a.config.Get("secret"),
	}
	resp := a.client.GetRaw(APIWeixin+tokenURLSuffix, m)
	return resp
}

func newAccessToken(config *Config, client *Client) *AccessToken {
	//client := NewClient(config)
	return &AccessToken{
		config:      config,
		client:      client,
		credentials: util.Map{},
	}
}

/*NewAccessToken NewAccessToken*/
func NewAccessToken(config *Config, client ...*Client) *AccessToken {
	if client == nil {
		return newAccessToken(config, NewClient())
	}
	return newAccessToken(config, client[0])
}

/*Refresh 刷新AccessToken */
func (a *AccessToken) Refresh() *AccessToken {
	log.Debug("AccessToken|Refresh")
	a.getToken(true)
	return a
}

/*GetRefreshedToken 获取刷新token */
func (a *AccessToken) GetRefreshedToken() *Token {
	log.Debug("AccessToken|GetRefreshedToken")
	return a.getToken(true)
}

/*GetToken 获取token */
func (a *AccessToken) GetToken() *Token {
	log.Debug("AccessToken|GetToken")
	return a.getToken(false)
}

/*GetTokenWithRefresh 重新获取token */
func (a *AccessToken) GetTokenWithRefresh() *Token {
	log.Debug("AccessToken|GetTokenWithRefresh")
	return a.getToken(true)
}

func (a *AccessToken) getToken(refresh bool) *Token {
	key := a.getCacheKey()

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

/*RequestToken 请求获取token */
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

/*SetTokenWithLife set string token with life time */
func (a *AccessToken) SetTokenWithLife(token string, lifeTime time.Time) *AccessToken {
	return a.setToken(token, lifeTime)
}

/*SetToken set string token */
func (a *AccessToken) SetToken(token string) *AccessToken {
	return a.setToken(token, time.Unix(7200, 0))
}

func (a *AccessToken) setToken(token string, lifeTime time.Time) *AccessToken {
	cache.SetWithTTL(a.getCacheKey(), &Token{
		AccessToken: token,
		ExpiresIn:   time.Now().Unix() + lifeTime.Unix(),
	}, time.Now())
	return a
}

func (a *AccessToken) getCredentials() string {
	c := md5.Sum(a.credentials.ToJSON())
	return fmt.Sprintf("%x", c[:])
}

func (a *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + a.getCredentials()
}
