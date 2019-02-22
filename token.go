package wego

import (
	"github.com/json-iterator/go"
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
	// AccessToken is the accessToken that authorizes and authenticates
	// the requests.
	AccessToken string `json:"access_token"`

	// RefreshToken is a accessToken that's used by the application
	// (as opposed to the user) to refresh the access accessToken
	// if it expires.
	RefreshToken string `json:"refresh_token"`

	// Expiry is the optional expiration time of the access accessToken.
	//
	// If zero, TokenSource implementations will reuse the same
	// accessToken forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	ExpiresIn int64 `json:"expires_in"`

	// wechat openid
	OpenID string `json:"openid"`

	// wechat scope
	Scope string `json:"scope"`
	// Raw optionally contains extra metadata from the server
	// when updating a accessToken.
	Raw interface{}
}

/*KeyMap get accessToken's key,value with map */
func (t *Token) KeyMap() util.Map {
	if t.AccessToken == "" {
		return nil
	}
	return util.Map{
		accessTokenKey: t.AccessToken,
	}
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

/*GetScopes get accessToken scopes for get accessToken*/
func (t *Token) GetScopes() []string {
	return strings.Split(t.Scope, ",")
}

/*SetScopes set accessToken scopes for get accessToken*/
func (t *Token) SetScopes(s []string) *Token {
	strings.Join(s, ",")
	return t
}

/*ToJSON transfer accessToken to json*/
func (t *Token) ToJSON() string {
	v, e := jsoniter.Marshal(t)
	if e != nil {
		return ""
	}
	return string(v)
}

/*ParseToken parse accessToken from string*/
func ParseToken(src string) (*Token, error) {
	var t Token
	e := jsoniter.Unmarshal([]byte(src), &t)
	if e != nil {
		return nil, e
	}
	return &t, nil
}
