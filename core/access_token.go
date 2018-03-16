package core

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godcong/wego/cache"
)

//type Token Map
//{"access_token":"7_VnJy3IwXNobh33ctx2SGFs0VBUQwqJC3hixeK8XAr-Wf8z7pm86S1Fvk9J0tHbSoRjxBFmAIJ4asbnOdQWicag","expires_in":7200,"refresh_token":"7__-y8XCMD549OYO9PifRba08tUhIfDhKUNQsKGUPs0hNDhxli9nlm0DfaD-bwPyZx8aAvoUwNEslRP6ckl-BDnA","openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc","scope":"snsapi_base"}
// Token represents the credentials used to authorize
// the requests to access protected resources on the OAuth 2.0
// provider's backend.
//
// This type is a mirror of oauth2.Token and exists to break
// an otherwise-circular dependency. Other internal packages
// should convert this Token into an oauth2.Token before use.
type Token struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string `json:"access_token"`

	// RefreshToken is a token that's used by the application
	// (as opposed to the user) to refresh the access token
	// if it expires.
	RefreshToken string `json:"refresh_token"`

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	ExpiresIn int64 `json:"expires_in"`

	//wechat openid
	OpenId string `json:"openid"`

	//wechat scope
	Scope string `json:"scope"`

	// Raw optionally contains extra metadata from the server
	// when updating a token.
	Raw interface{}
}

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

	m := a.client.Request(API_WEIXIN_URL_SUFFIX+CGI_BIN_TOKEN_URL_SUFFIX+"?"+m0.UrlEncode(), nil, "get", nil)
	return m.ToBytes()
	//return m
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

func (a *AccessToken) GetRefreshedToken() Token {
	Debug("AccessToken|GetRefreshedToken")
	return a.getToken(true)
}

func (a *AccessToken) GetToken() Token {
	Debug("AccessToken|GetToken")
	return a.getToken(false)
}

func (a *AccessToken) GetTokenWithRefresh() Token {
	Debug("AccessToken|GetTokenWithRefresh")
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
	Debug("AccessToken|getToken", token)
	if v := token.ExpiresIn; v != 0 {
		a.SetTokenWithLife(token.AccessToken, time.Unix(v, 0))
	} else {
		a.SetToken(token.AccessToken)
	}

	return token

}
func (a *AccessToken) RequestToken(credentials string) Token {
	response := a.sendRequest(credentials)
	m := Token{}
	json.Unmarshal(response, &m)
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

//func (t Token) SetKey(s string) Token {
//	(t)[ACCESS_TOKEN_KEY] = s
//	return t
//}

//func (t Token) GetKey() string {
//	if v, b := (t)[ACCESS_TOKEN_KEY]; b {
//		return v.(string)
//	}
//	return ""
//}

func (t Token) KeyMap() Map {
	m := make(Map)
	if t.AccessToken != "" {
		m.Set(ACCESS_TOKEN_KEY, t.AccessToken)
	}
	return m
}

//func (t Token) SetExpiresIn(i int64) Token {
//	(t)[ACCESS_TOKEN_EXPIRES_IN] = i
//	return t
//}

//func (t Token) GetExpiresIn() time.Time {
//	return t.ExpiresIn
//}

func (t Token) ToJson() string {
	v, e := json.Marshal(t)
	if e != nil {
		return ""
	}
	return string(v)
}
