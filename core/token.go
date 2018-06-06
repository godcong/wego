package core

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/godcong/wego/util"
)

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

	// wechat openid
	OpenID string `json:"openid"`

	// wechat scope
	Scope string `json:"scope"`
	// Raw optionally contains extra metadata from the server
	// when updating a token.
	Raw interface{}
}

const accessTokenNil = "nil point access token"
const tokenNil = "nil point token"

/*MustKeyMap get token's key,value with map when nil or error return nil map */
func MustKeyMap(at *AccessToken) util.Map {
	m := util.Map{}
	if m, e := KeyMap(at); e != nil {
		return m
	}
	return m
}

/*KeyMap get token's key,value with map */
func KeyMap(at *AccessToken) (util.Map, error) {
	if at == nil {
		return nil, errors.New(accessTokenNil)
	}
	if token := at.GetToken(); token != nil {
		return token.KeyMap(), nil
	}
	return nil, errors.New(tokenNil)
}

/*KeyMap get token's key,value with map */
func (t *Token) KeyMap() util.Map {
	m := make(util.Map)
	if t.AccessToken != "" {
		m.Set(AccessTokenKey, t.AccessToken)
	}
	return m
}

/*SetExpiresIn set expires time */
func (t *Token) SetExpiresIn(ti time.Time) *Token {
	t.ExpiresIn = ti.Unix()
	return t
}

/*GetExpiresIn get expires time */
func (t *Token) GetExpiresIn() time.Time {
	return time.Unix(t.ExpiresIn, 0)
}

/*ToJSON transfer token to json*/
func (t *Token) ToJSON() string {
	v, e := json.Marshal(t)
	if e != nil {
		return ""
	}
	return string(v)
}

/*GetScopes get token scopes for get token*/
func (t *Token) GetScopes() []string {
	return strings.Split(t.Scope, ",")
}

/*SetScopes set token scopes for get token*/
func (t *Token) SetScopes(s []string) *Token {
	strings.Join(s, ",")
	return t
}

/*ParseToken parse token from string*/
func ParseToken(j string) (*Token, error) {
	t := new(Token)
	e := json.Unmarshal([]byte(j), t)
	if e != nil {
		return nil, e
	}
	return t, nil
}
