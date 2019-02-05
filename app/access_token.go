package app

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"golang.org/x/exp/xerrors"
	"time"
)

// GrantType ...
type GrantType string

// GrantTypeClient ...
const GrantTypeClient GrantType = "client_credential"

// AccessTokenCredential ...
type AccessTokenCredential struct {
	GrantType GrantType
	AppID     string
	Secret    string
}

// ToMap ...
func (obj *AccessTokenCredential) ToMap() util.Map {
	return util.Map{
		"grant_type": obj.GrantType,
		"app_id":     obj.AppID,
		"secret":     obj.Secret,
	}
}

// ToJSON ...
func (obj *AccessTokenCredential) ToJSON() []byte {
	bytes, err := jsoniter.Marshal(obj)
	if err != nil {
		return nil
	}
	return bytes
}

// AccessTokenOption ...
type AccessTokenOption struct {
	RemoteHost string
	TokenKey   string
	TokenURL   string
}

/*AccessToken AccessToken */
type AccessToken struct {
	credential *AccessTokenCredential
	option     *AccessTokenOption
}

/*AccessTokenSafeSeconds token安全时间 */
const AccessTokenSafeSeconds = 500

// RemoteHost ...
func (obj *AccessToken) RemoteHost() string {
	return accessTokenRemoteHost(obj)
}
func accessTokenRemoteHost(obj *AccessToken) string {
	if obj != nil && obj.option != nil && obj.option.RemoteHost != "" {
		return obj.option.RemoteHost
	}
	return apiWeixin
}

// TokenURL ...
func (obj *AccessToken) TokenURL() string {
	return util.URL(obj.RemoteHost(), tokenURL(obj))
}
func tokenURL(obj *AccessToken) string {
	if obj != nil && obj.option != nil && obj.option.TokenURL != "" {
		return obj.option.TokenURL
	}
	return accessTokenURLSuffix
}

func newAccessToken(credential *AccessTokenCredential, opts ...*AccessTokenOption) *AccessToken {
	var opt *AccessTokenOption
	if opts != nil {
		opt = opts[0]
	}
	return &AccessToken{
		credential: credential,
		option:     opt,
	}
}

/*NewAccessToken NewAccessToken*/
func NewAccessToken(credential *AccessTokenCredential, opts ...*AccessTokenOption) *AccessToken {
	return newAccessToken(credential, opts...)
}

/*Refresh 刷新AccessToken */
func (obj *AccessToken) Refresh() *AccessToken {
	log.Debug("AccessToken|Refresh")
	obj.getToken(true)
	return obj
}

/*GetRefreshedToken 获取刷新token */
func (obj *AccessToken) GetRefreshedToken() *core.Token {
	log.Debug("AccessToken|GetRefreshedToken")
	return obj.getToken(true)
}

/*GetToken 获取token */
func (obj *AccessToken) GetToken() *core.Token {
	log.Debug("AccessToken|GetToken")
	return obj.getToken(false)
}

// KeyMap ...
func (obj *AccessToken) KeyMap() util.Map {
	log.Debug("AccessToken|KeyMap")
	token := obj.getToken(false)
	return token.KeyMap()
}

/*GetTokenWithRefresh 重新获取token */
func (obj *AccessToken) GetTokenWithRefresh() *core.Token {
	log.Debug("AccessToken|GetTokenWithRefresh")
	return obj.getToken(true)
}

func (obj *AccessToken) getToken(refresh bool) *core.Token {
	key := obj.getCacheKey()

	if !refresh && cache.Has(key) {
		log.Debug("cached accessToken", key)
		if v, b := cache.Get(key).(*core.Token); b {
			if v.ExpiresIn > time.Now().Unix() {
				return v
			}
		}
	}

	token := requestToken(obj.TokenURL(), obj.credential)
	if token == nil {
		return nil
	}
	log.Debug("AccessToken|getToken", *token)
	if v := token.ExpiresIn; v != 0 {
		obj.SetTokenWithLife(token.AccessToken, v-AccessTokenSafeSeconds)
	} else {
		obj.SetToken(token.AccessToken)
	}
	return token
}

func requestToken(url string, credentials *AccessTokenCredential) *core.Token {
	var token core.Token
	e := Get(url, credentials.ToMap()).Unmarshal(&token)
	if e != nil {
		log.Error("requestToken error", e)
		return nil
	}
	return &token
}

/*SetTokenWithLife set string accessToken with life time */
func (obj *AccessToken) SetTokenWithLife(token string, lifeTime int64) *AccessToken {
	return obj.setToken(token, lifeTime)
}

/*SetToken set string accessToken */
func (obj *AccessToken) SetToken(token string) *AccessToken {
	return obj.setToken(token, 7200)
}

func (obj *AccessToken) setToken(token string, lifeTime int64) *AccessToken {
	cache.SetWithTTL(obj.getCacheKey(), &core.Token{
		AccessToken: token,
		ExpiresIn:   time.Now().Add(time.Duration(lifeTime)).Unix(),
	}, lifeTime)
	return obj
}

func (obj *AccessToken) getCredentials() string {
	c := md5.Sum(obj.credential.ToJSON())
	return fmt.Sprintf("%x", c[:])
}

func (obj *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + obj.getCredentials()
}

const accessTokenNil = "nil point  accessToken"
const tokenNil = "nil point token"

/*MustKeyMap get accessToken's key,value with map when nil or error return nil map */
func MustKeyMap(at *AccessToken) util.Map {
	m := util.Map{}
	if m, e := KeyMap(at); e != nil {
		return m
	}
	return m
}

/*KeyMap get accessToken's key,value with map */
func KeyMap(at *AccessToken) (util.Map, error) {
	if at == nil {
		return nil, xerrors.New(accessTokenNil)
	}
	if token := at.GetToken(); token != nil {
		return token.KeyMap(), nil
	}
	return nil, xerrors.New(tokenNil)
}
