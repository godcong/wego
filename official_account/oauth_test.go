package official_account_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/official_account"
)

func TestOAuth_AuthCodeURL(t *testing.T) {
	oauth := official_account.NewOAuth()
	rlt := oauth.AuthCodeURL("run")
	t.Log(rlt)
}

func TestOAuth_ServeHTTP(t *testing.T) {
	oauth := official_account.NewOAuth()
	oauth.RegisterCodeCallback(func(w http.ResponseWriter, r *http.Request, val *official_account.CallbackValue) []byte {
		log.Debug(val)

		return nil
	})

	oauth.RegisterStateCallback(func(w http.ResponseWriter, r *http.Request, val *official_account.CallbackValue) []byte {
		log.Debug(val)
		return nil
	})

	oauth.RegisterAllCallback(func(w http.ResponseWriter, r *http.Request, val *official_account.CallbackValue) []byte {
		log.Debug(val)
		return nil
	})

	oauth.RegisterInfoCallback(func(w http.ResponseWriter, r *http.Request, val *official_account.CallbackValue) []byte {
		log.Debug(val.Type, *(val.Value.(*core.UserInfo)))
		return nil
	})
	ts := httptest.NewServer(http.HandlerFunc(oauth.ServeHTTP))
	defer ts.Close()
	resp, e := http.Get(ts.URL + "/oauth_callback?code=061FbdY41MrphL15MHX41U21Y41FbdYe&state=run")
	b, e := ioutil.ReadAll(resp.Body)
	log.Info(string(b), e)
}

func TestOAuth_AccessToken(t *testing.T) {
	oauth := official_account.NewOAuth()
	token := oauth.AccessToken("012QfESl1absJl0S73Tl1R7DSl1QfES4")
	t.Log(*token)
}

func TestOAuth_RefreshToken(t *testing.T) {
	oauth := official_account.NewOAuth()
	token := oauth.RefreshToken("7_Ug1inUynfYtLPvPPRmlSlRGLhHq9Y1YyH0PO9dLjxTpnJl7XERCc6_qTmaaj5Y-_tPHI2ib8m8fB2Tq_Epjb7w")
	t.Log(*token)
	testOAuth_Validate(t, oauth)
}

func TestOAuth_UserInfo(t *testing.T) {
	oauth := official_account.NewOAuth()
	token := oauth.RefreshToken("7_Ug1inUynfYtLPvPPRmlSlRGLhHq9Y1YyH0PO9dLjxTpnJl7XERCc6_qTmaaj5Y-_tPHI2ib8m8fB2Tq_Epjb7w")
	rlt := oauth.UserInfo(token)
	t.Log(*rlt)
}

func testOAuth_Validate(t *testing.T, auth *official_account.OAuth) {
	token := auth.RefreshToken("7_Ug1inUynfYtLPvPPRmlSlRGLhHq9Y1YyH0PO9dLjxTpnJl7XERCc6_qTmaaj5Y-_tPHI2ib8m8fB2Tq_Epjb7w")
	rlt := auth.Validate(token)
	t.Log(rlt)
}
