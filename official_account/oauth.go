package official_account

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/godcong/wego/core"
)

type CallbackFunc func(w http.ResponseWriter, r *http.Request, token *core.Token) bool

type OAuth struct {
	*OfficialAccount
	core.Config
	domain      *core.Domain
	response    *core.Response
	callback    CallbackFunc
	authorize   string
	scopes      string
	redirectUri string
}

func newOAuth(officialAccount *OfficialAccount) *OAuth {
	core.Debug("newOAuth", officialAccount)
	oauth := &OAuth{
		OfficialAccount: officialAccount,
	}

	oauth.Config = core.GetConfig("official_account.oauth")
	oauth.domain = core.DomainHost()
	oauth.scopes = oauth.GetD("scopes", SNSAPI_BASE)
	oauth.redirectUri = oauth.GetD("oauth_redirect_uri", DEFAULT_OAUTH_REDIRECT_URI_SUFFIX)
	oauth.authorize = oauth.GetD("redirect", OAUTH2_AUTHORIZE_URL_SUFFIX)
	return oauth
}

func NewOAuth() *OAuth {
	return newOAuth(account)
}

func (o *OAuth) RegisterCallback(callbackFunc CallbackFunc) *OAuth {
	o.callback = callbackFunc
	return o
}

func (o *OAuth) PrepareCallbackUrl() {
	//$callback = $app['config']->get('oauth.callback');
	//if (0 === stripos($callback, 'http')) {
	//return $callback;
	//}
	//$baseUrl = $app['request']->getSchemeAndHttpHost();
	//
	//return $baseUrl.'/'.ltrim($callback, '/');
}

//失败：
//{"errcode":40163,"errmsg":"code been used, hints: [ req_id: OsIKda0848th19 ]"}
//{"errcode":40029,"errmsg":"invalid code, hints: [ req_id: 5u8NWa0990th40 ]"}
//成功：
//{"access_token":"7_0MSpG_WEPVwQki6eFQSFQbRwkEkTEhkvBjkuKTODS7_xe6vBOEsc88kcCu_781YvXXP2FwWC4M5m-B9WXs51rA","expires_in":7200,"refresh_token":"7_51Axvh89ev5cGH-WR4qPKb-rcPf2VQrMg25MNDs1899cHYb5UomPi4fnc1NAks07Vw5Bb0pTFvvritU-aQtxFg","openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc","scope":"snsapi_base"}]
func (o *OAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if code := r.Form.Get("code"); code != "" {
		config := o.OfficialAccount.Config
		v := core.Map{
			"appid":      config.Get("app_id"),
			"secret":     config.Get("secret"),
			"code":       code,
			"grant_type": "authorization_code",
		}
		if o.redirectUri != "" {
			v.Set("redirect_uri", o.domain.Link(o.redirectUri))
		}
		response := o.client.HttpPost(
			o.client.Link(OAUTH2_ACCESS_TOKEN_URL_SUFFIX),
			core.Map{
				core.REQUEST_TYPE_QUERY.String(): v,
			},
		)
		core.Debug("ServeHTTP|response", response)
		var token core.Token
		e := json.Unmarshal(response.ToBytes(), &token)
		if e != nil {
			core.Debug("ServeHTTP|e", e)
		}
		if o.callback != nil {
			if b := o.callback(w, r, &token); b {
				return
			}
		}
		w.Write(response.ToJson())
		return
	}
	http.Redirect(w, r, o.AuthCodeURL(""), http.StatusFound)
}
func retrieveToken(ctx context.Context, auth *OAuth, values url.Values) (*core.Token, error) {

	return nil, nil
}

func (o *OAuth) AuthCodeURL(state string) string {
	core.Debug("AuthCodeURL|OfficialAccount", o.OfficialAccount)
	var buf bytes.Buffer
	buf.WriteString(o.authorize)
	v := url.Values{
		"response_type": {"code"},
		"appid":         {o.OfficialAccount.Get("app_id")},
	}
	if o.redirectUri != "" {
		v.Set("redirect_uri", o.domain.Link(o.redirectUri))
	}
	if o.scopes != "" {
		v.Set("scope", o.scopes)
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		v.Set("state", state)
	}

	if !strings.Contains(o.authorize, "?") {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
}

func (o *OAuth) GetResponse() *core.Response {
	return o.response
}

//https://api.weixin.qq.com/sns/oauth2/refresh_token
func (o *OAuth) RefreshToken() {

}
