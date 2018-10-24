package official_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
)

func TestOAuth_AuthCodeURL(t *testing.T) {
	oauth := official.NewOAuth(config)
	rlt := oauth.AuthCodeURL("run")
	t.Log(rlt)
}

func TestOAuth_ServeHTTP(t *testing.T) {
	oauth := official.NewOAuth(config)
	oauth.RegisterCodeCallback(func(w http.ResponseWriter, r *http.Request, val *official.CallbackValue) []byte {
		log.Debug(val)

		return nil
	})

	oauth.RegisterStateCallback(func(w http.ResponseWriter, r *http.Request, val *official.CallbackValue) []byte {
		log.Debug(val)
		return nil
	})

	oauth.RegisterAllCallback(func(w http.ResponseWriter, r *http.Request, val *official.CallbackValue) []byte {
		log.Debug(val)
		return nil
	})

	oauth.RegisterInfoCallback(func(w http.ResponseWriter, r *http.Request, val *official.CallbackValue) []byte {
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
	oauth := official.NewOAuth(config)
	token := oauth.AccessToken("011e3u1S0YsI082tiH0S0UEK1S0e3u1A")
	//{15_wrO-m7suC_CjAuogtANDIL6iQk9KSQaAIBq2WuhINwHj_Eii6UNDm-3Y-L6Yzz2fLARsuV07f193vWAVdkE6EQ 15_kB8iaR-T_YirCLd6NQYNJiVRO7eAuGo3rYnYSf8FISxAs2Ifpml59Oox0fpi6WsaXZT7LYRXy-J8wUG2g04GDQ 7200 oJ9y41a_Sv0XSliN-dOEoobF-B3U snsapi_userinfo <nil>}
	t.Log(fmt.Sprintf("%+v", token))
	testOAuth_Validate(t, token)
}

func TestOAuth_RefreshToken(t *testing.T) {
	oauth := official.NewOAuth(config)
	token := oauth.RefreshToken("15_kB8iaR-T_YirCLd6NQYNJiVRO7eAuGo3rYnYSf8FISxAs2Ifpml59Oox0fpi6WsaXZT7LYRXy-J8wUG2g04GDQ")
	//&{AccessToken:15_wrO-m7suC_CjAuogtANDIL6iQk9KSQaAIBq2WuhINwHj_Eii6UNDm-3Y-L6Yzz2fLARsuV07f193vWAVdkE6EQ RefreshToken:15_kB8iaR-T_YirCLd6NQYNJiVRO7eAuGo3rYnYSf8FISxAs2Ifpml59Oox0fpi6WsaXZT7LYRXy-J8wUG2g04GDQ ExpiresIn:7200 OpenID:oJ9y41a_Sv0XSliN-dOEoobF-B3U Scope:snsapi_base,snsapi_userinfo, Raw:<nil>}
	t.Log(fmt.Sprintf("%+v", token))
	testOAuth_Validate(t, token)
}

func TestOAuth_UserInfo(t *testing.T) {
	oauth := official.NewOAuth(config)
	token := oauth.RefreshToken("15_kB8iaR-T_YirCLd6NQYNJiVRO7eAuGo3rYnYSf8FISxAs2Ifpml59Oox0fpi6WsaXZT7LYRXy-J8wUG2g04GDQ")
	//{"openid":"oJ9y41a_Sv0XSliN-dOEoobF-B3U","access_token":"15_wrO-m7suC_CjAuogtANDIL6iQk9KSQaAIBq2WuhINwHj_Eii6UNDm-3Y-L6Yzz2fLARsuV07f193vWAVdkE6EQ","expires_in":7200,"refresh_token":"15_kB8iaR-T_YirCLd6NQYNJiVRO7eAuGo3rYnYSf8FISxAs2Ifpml59Oox0fpi6WsaXZT7LYRXy-J8wUG2g04GDQ","scope":"snsapi_base,snsapi_userinfo,"}
	rlt := oauth.UserInfo(token)
	t.Log(fmt.Sprintf("%+v", rlt))
}

func testOAuth_Validate(t *testing.T, token *core.Token) {
	oauth := official.NewOAuth(config)
	rlt := oauth.Validate(token)
	t.Log(rlt)
}
