package wego

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
	"time"
)

// AccessTokenOption ...
type AccessTokenOption struct {
	RemoteHost string
	TokenKey   string
	TokenURL   string
}

/*AccessToken AccessToken */
type AccessToken struct {
	*AccessTokenConfig
	remoteHost string
	tokenKey   string
	tokenURL   string
}

/*AccessTokenSafeSeconds token安全时间 */
const AccessTokenSafeSeconds = 500

// RemoteHost ...
func (obj *AccessToken) RemoteHost() string {
	if obj != nil && obj.remoteHost != "" {
		return obj.remoteHost
	}
	return apiWeixin
}

// TokenURL ...
func (obj *AccessToken) TokenURL() string {
	return util.URL(obj.RemoteHost(), tokenURL(obj))
}
func tokenURL(obj *AccessToken) string {
	if obj != nil && obj.tokenURL != "" {
		return obj.tokenURL
	}
	return accessTokenURLSuffix
}

/*NewAccessToken NewAccessToken*/
func NewAccessToken(property *AccessTokenConfig, options ...*AccessTokenOption) *AccessToken {
	token := &AccessToken{
		AccessTokenConfig: property,
	}
	token.parse(options)
	return token
}

func (obj *AccessToken) parse(options []*AccessTokenOption) {
	if options == nil {
		return
	}
	obj.remoteHost = options[0].RemoteHost
	obj.tokenURL = options[0].TokenURL
	obj.tokenKey = options[0].TokenKey
}

/*Refresh 刷新AccessToken */
func (obj *AccessToken) Refresh() *AccessToken {
	log.Debug("AccessToken|Refresh")
	obj.getToken(true)
	return obj
}

/*GetRefreshToken 获取刷新token */
func (obj *AccessToken) GetRefreshToken() *Token {
	log.Debug("AccessToken|GetRefreshedToken")
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

	if !refresh && cache.Has(key) {
		if v, b := cache.Get(key).(*Token); b {
			if v.ExpiresIn > time.Now().Unix() {
				log.Infof("cached accessToken:%+v", v)
				return v
			}
		}
	}

	token, e := requestToken(obj.TokenURL(), obj.AccessTokenConfig)
	if e != nil {
		log.Error(e)
		return nil
	}

	log.Infof("accessToken:%+v", *token)
	if v := token.ExpiresIn; v != 0 {
		obj.SetTokenWithLife(token.AccessToken, v-AccessTokenSafeSeconds)
	} else {
		obj.SetToken(token.AccessToken)
	}
	return token
}

func requestToken(url string, credentials *AccessTokenConfig) (*Token, error) {
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
	cache.SetWithTTL(obj.getCacheKey(), &Token{
		AccessToken: token,
		ExpiresIn:   time.Now().Add(time.Duration(lifeTime)).Unix(),
	}, lifeTime)
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

func parseAccessToken(token interface{}) string {
	switch v := token.(type) {
	case AccessToken:
		return v.GetToken().AccessToken
	case *AccessToken:
		return v.GetToken().AccessToken
	case string:
		return v
	}
	return ""
}
