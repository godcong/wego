package app

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/godcong/wego/core"
	"time"

	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*AccessToken AccessToken */
type AccessToken struct {
	URL         string
	TokenKey    string
	credentials util.Map
	prefix      string
}

/*accessTokenKey 键值 */
const accessTokenKey = "access_token"
const accessTokenURLSuffix = "/cgi-bin/token"

/*AccessTokenSafeSeconds token安全时间 */
const AccessTokenSafeSeconds = 500

func (a *AccessToken) sendRequest(s string) []byte {
	return core.Get(core.Splice(a.prefix, a.URL), a.credentials).Bytes()
}

func newAccessToken(p util.Map) *AccessToken {
	return &AccessToken{
		URL:         accessTokenURLSuffix,
		TokenKey:    accessTokenKey,
		credentials: p,
		prefix:      core.APIWeixin,
	}
}

// ClientCredential ...
func ClientCredential(config *core.Config) util.Map {
	return util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	}
}

// CredentialGet ...
func CredentialGet(v []interface{}) util.Map {
	size := len(v)
	for i := 0; i < size; i++ {
		switch vv := v[i].(type) {
		case util.Map:
			return vv
		case map[string]interface{}:
			return (util.Map)(vv)
		}
	}

	return nil
}

/*NewAccessToken NewAccessToken*/
func NewAccessToken(v ...interface{}) *AccessToken {
	accessToken := newAccessToken(CredentialGet(v))
	return accessToken
}

// Credentials ...
func (a *AccessToken) Credentials() util.Map {
	return a.credentials
}

//SetCredentials set request credential
func (a *AccessToken) SetCredentials(p util.Map) *AccessToken {
	if idx := p.Check("grant_type", "appid", "secret"); idx != -1 {
		log.Error(fmt.Errorf("the %d key was not found", idx))
	}
	a.credentials = p
	return a
}

/*Refresh 刷新AccessToken */
func (a *AccessToken) Refresh() *AccessToken {
	log.Debug("AccessToken|Refresh")
	a.getToken(true)
	return a
}

/*GetRefreshedToken 获取刷新token */
func (a *AccessToken) GetRefreshedToken() *core.Token {
	log.Debug("AccessToken|GetRefreshedToken")
	return a.getToken(true)
}

/*GetToken 获取token */
func (a *AccessToken) GetToken() *core.Token {
	log.Debug("AccessToken|GetToken")
	return a.getToken(false)
}

// KeyMap ...
func (a *AccessToken) KeyMap() util.Map {
	log.Debug("AccessToken|KeyMap")
	token := a.getToken(false)
	return token.KeyMap()
}

/*GetTokenWithRefresh 重新获取token */
func (a *AccessToken) GetTokenWithRefresh() *core.Token {
	log.Debug("AccessToken|GetTokenWithRefresh")
	return a.getToken(true)
}

func (a *AccessToken) getToken(refresh bool) *core.Token {
	key := a.getCacheKey()

	if !refresh && cache.Has(key) {
		log.Debug("cached accessToken", key)
		if v, b := cache.Get(key).(*core.Token); b {
			if v.ExpiresIn > time.Now().Unix() {
				return v
			}
		}
	}

	token := a.RequestToken(a.getCredentials())
	if token == nil {
		return nil
	}
	log.Debug("AccessToken|getToken", *token)
	if v := token.ExpiresIn; v != 0 {
		a.SetTokenWithLife(token.AccessToken, v)
	} else {
		a.SetToken(token.AccessToken)
	}

	return token

}

/*RequestToken 请求获取token */
func (a *AccessToken) RequestToken(credentials string) *core.Token {
	var token core.Token
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

/*SetTokenWithLife set string accessToken with life time */
func (a *AccessToken) SetTokenWithLife(token string, lifeTime int64) *AccessToken {
	return a.setToken(token, lifeTime)
}

/*SetToken set string accessToken */
func (a *AccessToken) SetToken(token string) *AccessToken {
	return a.setToken(token, 7200)
}

func (a *AccessToken) setToken(token string, lifeTime int64) *AccessToken {
	cache.SetWithTTL(a.getCacheKey(), &core.Token{
		AccessToken: token,
		ExpiresIn:   time.Now().Add(time.Duration(lifeTime)).Unix(),
	}, lifeTime)
	return a
}

func (a *AccessToken) getCredentials() string {
	c := md5.Sum(a.credentials.ToJSON())
	return fmt.Sprintf("%x", c[:])
}

func (a *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + a.getCredentials()
}
