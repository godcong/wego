package official_account

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"

	"github.com/godcong/wego/core"
)

type OAuth struct {
	*OfficialAccount
	core.Config
	domain    *core.Domain
	response  *core.Response
	authorize string
	scopes    string
	callback  string
}

func newOAuth(officialAccount *OfficialAccount) *OAuth {
	core.Debug("newOAuth", officialAccount)
	oauth := &OAuth{
		OfficialAccount: officialAccount,
	}

	oauth.Config = core.GetConfig("official_account.oauth")
	oauth.domain = core.DomainHost()
	oauth.scopes = oauth.GetD("scopes", SNSAPI_BASE)
	oauth.callback = oauth.GetD("callback", DEFAULT_CALLBACK_URL_SUFFIX)
	oauth.authorize = oauth.GetD("redirect", OAUTH2_AUTHORIZE_URL_SUFFIX)
	return oauth
}

func NewOAuth() *OAuth {
	return newOAuth(account)
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
//成功：
//{"access_token":"7_0MSpG_WEPVwQki6eFQSFQbRwkEkTEhkvBjkuKTODS7_xe6vBOEsc88kcCu_781YvXXP2FwWC4M5m-B9WXs51rA","expires_in":7200,"refresh_token":"7_51Axvh89ev5cGH-WR4qPKb-rcPf2VQrMg25MNDs1899cHYb5UomPi4fnc1NAks07Vw5Bb0pTFvvritU-aQtxFg","openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc","scope":"snsapi_base"}]
func (o *OAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if code := r.Form.Get("code"); code != "" {
		config := o.OfficialAccount.Config
		o.response = o.client.HttpGet(
			o.client.Link(OAUTH2_ACCESS_TOKEN_URL_SUFFIX),
			core.Map{
				core.REQUEST_TYPE_QUERY.String(): core.Map{
					"appid":      config.Get("app_id"),
					"secret":     config.Get("secret"),
					"code":       code,
					"grant_type": "authorization_code",
				},
			},
		)
		core.Debug("OAuth|ServeHTTP|response", o.response)
		return
	}
	http.Redirect(w, r, o.AuthCodeURL(""), http.StatusFound)
}

func (o *OAuth) AuthCodeURL(state string) string {
	core.Debug("AuthCodeURL|OfficialAccount", o.OfficialAccount)
	var buf bytes.Buffer
	buf.WriteString(o.authorize)
	v := url.Values{
		"response_type": {"code"},
		"appid":         {o.OfficialAccount.Get("app_id")},
	}
	if o.callback != "" {
		v.Set("redirect_uri", o.domain.Link(o.callback))
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

//func NewOAuth(application Application, config Config) OAuth {
//	return &Oauth{
//		Config: config,
//		app:    application,
//		//client: application.Client(),
//	}
//}

//qq回调配置
//https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=310198347&redirect_uri=http%3A%2F%2Fwww.right.com.cn%2Fforum%2Fconnect.php%3Fmod%3Dlogin%26op%3Dcallback%26referer%3Dhttp%253A%252F%252Fwww.right.com.cn%252Fforum%252Fthread-147109-1-1.html&state=72a5eb8ae2eba26edc851175955d5094&scope=get_user_info%2Cadd_share%2Cadd_t%2Cadd_pic_t%2Cget_repost_list
