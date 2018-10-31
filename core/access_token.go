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
	client      *Client
	URI         string
	TokenKey    string
	Credentials util.Map
}

/*AccessTokenKey 键值 */
const AccessTokenKey = "access_token"

/*AccessTokenExpiresIn 过期 */
const AccessTokenExpiresIn = "expires_in"

/*AccessTokenSafeSeconds token安全时间 */
const AccessTokenSafeSeconds = 500

func (a *AccessToken) sendRequest(s string) []byte {
	client := a.client
	if client == nil {
		client = DefaultClient()
	}
	resp := client.GetRaw(Connect(APIWeixin, tokenURLSuffix), a.Credentials)
	return resp
}

func newAccessToken() *AccessToken {
	//client := NewClient(config)
	return &AccessToken{
		TokenKey: AccessTokenKey,
	}
}

/*NewAccessToken NewAccessToken*/
func NewAccessToken() *AccessToken {
	return newAccessToken()
}

//SetCredentials set request credential
func (a *AccessToken) SetCredentials(p util.Map) *AccessToken {
	if idx := p.Check("grant_type", "appid", "secret"); idx != -1 {
		log.Error(fmt.Errorf("the %d key was not found", idx))
	}
	a.Credentials = p
	return a
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
	life := time.Now().Add(time.Duration(lifeTime.Unix()) - AccessTokenSafeSeconds)
	cache.SetWithTTL(a.getCacheKey(), &Token{
		AccessToken: token,
		ExpiresIn:   lifeTime.Unix(),
	}, &life)
	return a
}

func (a *AccessToken) getCredentials() string {
	c := md5.Sum(a.Credentials.ToJSON())
	return fmt.Sprintf("%x", c[:])
}

func (a *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + a.getCredentials()
}
