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
	ts := httptest.NewServer(http.HandlerFunc(oauth.ServeHTTP))
	defer ts.Close()

	resp, e := http.Get(ts.URL + "/oauth_callback?code=0819yiqr16NBan0ml2tr1e1eqr19yiqp&state=run")

	b, e := ioutil.ReadAll(resp.Body)

	core.Info(string(b), e)

	//http://shop.commm.top/oauth_callback?code=0712iHTj2YqtYE0BBpRj2GQZTj22iHTr&state=run
}
