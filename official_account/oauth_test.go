package official_account_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/official_account"
)

func TestOAuth_AuthCodeURL(t *testing.T) {
	oauth := official_account.NewOAuth()
	rlt := oauth.AuthCodeURL("run")
	t.Log(rlt)
}

func TestOAuth_ServeHTTP(t *testing.T) {
	oauth := official_account.NewOAuth()
	oauth.RegisterCallback(func(w http.ResponseWriter, r *http.Request, token *core.Token) bool {
		core.Debug(*token)
		return false
	})
	ts := httptest.NewServer(http.HandlerFunc(oauth.ServeHTTP))
	defer ts.Close()
	//core.Debug("resp", oauth.GetResponse().ToString())

	resp, e := http.Get(ts.URL + "/oauth_redirect?code=001V9bZR0lce192bq5XR0RZmZR0V9bZt&state=run")

	b, e := ioutil.ReadAll(resp.Body)

	core.Info(string(b), e)

	//http://shop.commm.top/oauth_callback?code=0712iHTj2YqtYE0BBpRj2GQZTj22iHTr&state=run
}
