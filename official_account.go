package wego

import (
	"bytes"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"strings"
)

// OfficialAccount ...
type OfficialAccount struct {
	*OfficialAccountConfig
	BodyType    BodyType
	client      *Client
	remoteHost  string
	redirectURI string
	localHost   string
}

// OfficialAccountOption ...
type OfficialAccountOption struct {
	BodyType   *BodyType
	RemoteHost string
	LocalHost  string
}

// NewOfficialAccount ...
func NewOfficialAccount(config *OfficialAccountConfig, opts ...*OfficialAccountOption) *OfficialAccount {
	officialAccount := &OfficialAccount{
		BodyType:              BodyTypeJSON,
		OfficialAccountConfig: config,
	}
	officialAccount.parse(opts)
	officialAccount.client = NewClient(&ClientOption{
		BodyType: &officialAccount.BodyType,
	})
	return officialAccount
}

func (obj *OfficialAccount) parse(opts []*OfficialAccountOption) {
	if opts == nil {
		return
	}
	if opts[0].BodyType != nil {
		obj.BodyType = *opts[0].BodyType
	}
	obj.redirectURI = opts[0].RedirectURI
	obj.remoteHost = opts[0].RemoteHost
	obj.localHost = opts[0].LocalHost
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

	v := util.Map{
		"appid":      obj.AppID,
		"secret":     obj.AppSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	log.Debug(obj.redirectURI)
	if obj.redirectURI != "" {
		if strings.Index(obj.redirectURI, "http") == 0 {
			v.Set("redirect_uri", obj.redirectURI)
		} else {
			//TODO:
			v.Set("redirect_uri", util.URL(obj.redirectURI))
		}
	}
	responder := PostJSON(
		oauth2AccessToken,
		v,
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

	if obj.Scopes != nil {
		p.Set("scope", obj.Scopes)
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
	log.Debug("RedirectURI", obj.redirectURI)
	if obj.redirectURI != "" {
		if strings.Index(obj.redirectURI, "http") == 0 {
			return obj.redirectURI
		}
		return util.URL(obj.localHost, obj.redirectURI)
	}
	return ""
}
