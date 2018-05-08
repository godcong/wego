package core

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/godcong/wego/util"
)

// type Token Map
// {"access_token":"7_VnJy3IwXNobh33ctx2SGFs0VBUQwqJC3hixeK8XAr-Wf8z7pm86S1Fvk9J0tHbSoRjxBFmAIJ4asbnOdQWicag","expires_in":7200,"refresh_token":"7__-y8XCMD549OYO9PifRba08tUhIfDhKUNQsKGUPs0hNDhxli9nlm0DfaD-bwPyZx8aAvoUwNEslRP6ckl-BDnA","openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc","scope":"snsapi_base"}
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
	OpenId string `json:"openid"`

	// wechat scope
	Scope string `json:"scope"`
	// Raw optionally contains extra metadata from the server
	// when updating a token.
	Raw interface{}
}

func (t *Token) KeyMap() util.Map {
	m := make(util.Map)
	if t.AccessToken != "" {
		m.Set(ACCESS_TOKEN_KEY, t.AccessToken)
	}
	return m
}

func (t *Token) SetExpiresIn(ti time.Time) *Token {
	t.ExpiresIn = ti.Unix()
	return t
}

func (t *Token) GetExpiresIn() time.Time {
	return time.Unix(t.ExpiresIn, 0)
}

func (t *Token) ToJson() string {
	v, e := json.Marshal(t)
	if e != nil {
		return ""
	}
	return string(v)
}

func (t *Token) GetScopes() []string {
	return strings.Split(t.Scope, ",")
}

func (t *Token) SetScopes(s []string) *Token {
	strings.Join(s, ",")
	return t
}

func ParseToken(j string) (*Token, error) {
	t := new(Token)
	e := json.Unmarshal([]byte(j), t)
	if e != nil {
		return nil, e
	}
	return t, nil
}
