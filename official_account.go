package wego

import (
	"bytes"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"strings"
)

// OfficialAccount ...
type OfficialAccount struct {
	*OfficialAccountProperty
	BodyType   BodyType
	client     *Client
	remoteHost string
	localHost  string
}

// OfficialAccountOption ...
type OfficialAccountOption struct {
	BodyType   *BodyType
	RemoteHost string
	LocalHost  string
}

// NewOfficialAccount ...
func NewOfficialAccount(config *OfficialAccountProperty, options ...*OfficialAccountOption) *OfficialAccount {
	officialAccount := &OfficialAccount{
		BodyType:                BodyTypeJSON,
		OfficialAccountProperty: config,
	}
	officialAccount.parse(options)
	officialAccount.client = NewClient(&ClientOption{
		BodyType: &officialAccount.BodyType,
	})
	return officialAccount
}

func (obj *OfficialAccount) parse(options []*OfficialAccountOption) {
	if options == nil {
		return
	}
	if options[0].BodyType != nil {
		obj.BodyType = *options[0].BodyType
	}
	obj.remoteHost = options[0].RemoteHost
	obj.localHost = options[0].LocalHost
}

// HandleAuthorizeNotify ...
func (obj *OfficialAccount) HandleAuthorizeNotify(hooks ...interface{}) ServeHTTPFunc {
	return obj.HandleAuthorize(hooks...).ServeHTTP
}

// HandleAuthorize ...
func (obj *OfficialAccount) HandleAuthorize(hooks ...interface{}) Notifier {
	notify := &authorizeNotify{
		OfficialAccount: obj,
	}
	for _, hook := range hooks {
		switch h := hook.(type) {
		case TokenHook:
			notify.TokenHook = h
		case UserHook:
			notify.UserHook = h
		case StateHook:
			notify.StateHook = h
		}
	}
	return notify
}

// GetUserInfo ...
func (obj *OfficialAccount) GetUserInfo(token *core.Token) (*core.WechatUserInfo, error) {
	var info core.WechatUserInfo
	var e error
	p := util.Map{
		"access_token": token.AccessToken,
		"openid":       token.OpenID,
		"lang":         "zh_CN",
	}
	responder := Get(
		snsUserinfo,
		p,
	)
	log.Debug("WechatUserInfo|responder", string(responder.Bytes()), responder.Error())
	e = responder.Error()
	if e != nil {
		return &info, e
	}

	e = responder.Unmarshal(&info)
	if e != nil {
		return &info, e
	}
	return &info, nil
}

// Oauth2AuthorizeToken ...
func (obj *OfficialAccount) Oauth2AuthorizeToken(code string) (*core.Token, error) {
	var token core.Token
	var e error

	p := util.Map{
		"appid":      obj.AppID,
		"secret":     obj.AppSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}

	uri := obj.RedirectURI()
	if uri != "" {
		p.Set("redirect_uri", uri)
	}

	responder := PostJSON(
		oauth2AccessToken,
		p,
		nil,
	)
	e = responder.Error()
	log.Debug("GetAuthorizeToken|response", string(responder.Bytes()), e)
	if e != nil {
		return &token, e
	}

	e = responder.Unmarshal(&token)
	if e != nil {
		return &token, e
	}
	return &token, nil
}

/*AuthCodeURL 生成授权地址URL*/
func (obj *OfficialAccount) AuthCodeURL(state string) string {
	log.Debug("AuthCodeURL", state)
	var buf bytes.Buffer
	buf.WriteString(oauth2Authorize)
	p := util.Map{
		"response_type": "code",
		"appid":         obj.AppID,
	}

	uri := obj.RedirectURI()
	if uri != "" {
		p.Set("redirect_uri", uri)
	}

	if obj.OAuth.Scopes != nil {
		p.Set("scope", obj.OAuth.Scopes)
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		p.Set("state", state)
	}
	buf.WriteByte('?')
	buf.WriteString(p.URLEncode())
	return buf.String()
}

// RedirectURI ...
func (obj *OfficialAccount) RedirectURI() string {
	log.Debug("RedirectURI", obj.OAuth.RedirectURI)
	if obj.OAuth.RedirectURI != "" {
		if strings.Index(obj.OAuth.RedirectURI, "http") == 0 {
			return obj.OAuth.RedirectURI
		}
		return util.URL(obj.localHost, obj.OAuth.RedirectURI)
	}
	return ""
}
