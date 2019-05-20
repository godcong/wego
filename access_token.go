package wego

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/util"
	"golang.org/x/xerrors"
	"strings"
)

/*AccessToken GetToken */
type AccessToken struct {
	*AccessTokenProperty
	remoteURL string
	tokenKey  string
	tokenURL  string
}

/*AccessTokenSafeSeconds token安全时间 */
const AccessTokenSafeSeconds = 500

// RemoteURL ...
func (obj *AccessToken) RemoteURL() string {
	if obj != nil && obj.remoteURL != "" {
		return obj.remoteURL
	}
	return apiWeixin
}

// TokenURL ...
func (obj *AccessToken) TokenURL() string {
	return util.URL(obj.RemoteURL(), tokenURL(obj))
}
func tokenURL(obj *AccessToken) string {
	if obj != nil && obj.tokenURL != "" {
		return obj.tokenURL
	}
	return accessToken
}

/*NewAccessToken NewAccessToken*/
func NewAccessToken(property *AccessTokenProperty, options ...AccessTokenOption) *AccessToken {
	token := &AccessToken{
		AccessTokenProperty: property,
	}
	token.parse(options...)
	return token
}

func (obj *AccessToken) parse(options ...AccessTokenOption) {
	if options == nil {
		return
	}
	for _, o := range options {
		o(obj)
	}
}

/*Refresh 刷新AccessToken */
func (obj *AccessToken) Refresh() *AccessToken {
	log.Debug("GetToken|Refresh")
	obj.getToken(true)
	return obj
}

/*GetRefreshToken 获取刷新token */
func (obj *AccessToken) GetRefreshToken() *Token {
	log.Debug("GetToken|GetRefreshedToken")
	return obj.getToken(true)
}

/*GetToken 获取token */
func (obj *AccessToken) GetToken() *Token {
	return obj.getToken(false)
}

// KeyMap ...
func (obj *AccessToken) KeyMap() util.Map {
	return MustKeyMap(obj)
}

func (obj *AccessToken) getToken(refresh bool) *Token {
	key := obj.getCacheKey()
	log.Infof("cached key:%t,%s", cache.Has(key), key)
	if !refresh && cache.Has(key) {
		if v, b := cache.Get(key).(string); b {
			token, e := ParseToken(v)
			if e != nil {
				log.Error("parse token error")
				return nil
			}
			log.Infof("cached accessToken:%+v", v)
			return token
		}
	}

	token, e := requestToken(obj.TokenURL(), obj.AccessTokenProperty)
	if e != nil {
		log.Error(e)
		return nil
	}

	log.Infof("accessToken:%+v", *token)
	if v := token.ExpiresIn; v != 0 {
		obj.SetTokenWithLife(token.ToJSON(), v-AccessTokenSafeSeconds)
	} else {
		obj.SetToken(token.ToJSON())
	}
	return token
}

func requestToken(url string, credentials *AccessTokenProperty) (*Token, error) {
	var t Token
	var e error
	token := Get(url, credentials.ToMap())
	if e := token.Error(); e != nil {
		return nil, e
	}
	e = token.Unmarshal(&t)
	if e != nil {
		return nil, e
	}
	return &t, nil
}

/*SetTokenWithLife set string accessToken with life time */
func (obj *AccessToken) SetTokenWithLife(token string, lifeTime int64) *AccessToken {
	return obj.setToken(token, lifeTime)
}

/*SetToken set string accessToken */
func (obj *AccessToken) SetToken(token string) *AccessToken {
	return obj.setToken(token, 7200-AccessTokenSafeSeconds)
}

func (obj *AccessToken) setToken(token string, lifeTime int64) *AccessToken {
	cache.SetWithTTL(obj.getCacheKey(), token, lifeTime)
	return obj
}

func (obj *AccessToken) getCredentials() string {
	cred := strings.Join([]string{obj.GrantType, obj.AppID, obj.AppSecret}, ".")
	c := md5.Sum([]byte(cred))
	return fmt.Sprintf("%x", c[:])
}

func (obj *AccessToken) getCacheKey() string {
	return "godcong.wego.access_token." + obj.getCredentials()
}

const accessTokenNil = "nil point accessToken"
const tokenNil = "nil point token"

/*MustKeyMap get accessToken's key,value with map when nil or error return nil map */
func MustKeyMap(at *AccessToken) util.Map {
	if m, e := KeyMap(at); e == nil {
		return m
	}
	return util.Map{}
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

func parseAccessToken(token interface{}) string {
	switch v := token.(type) {
	case Token:
		return v.AccessToken
	case *Token:
		return v.AccessToken
	case string:
		return v
	}
	return ""
}
