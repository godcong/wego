package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

var u = official.NewUser(config)

func TestNewUser(t *testing.T) {

}

func TestUser_UpdateRemark(t *testing.T) {
	resp := u.UpdateRemark("oLyBi0tDnybg0WFkhKsn5HRetX1I", "nishi123")
	t.Log(string(resp.Bytes()))
}

func TestUser_UserInfo(t *testing.T) {
	resp := u.UserInfo("oLyBi0tDnybg0WFkhKsn5HRetX1I", "zh_CN")
	t.Log(*resp)
}

func TestUser_BatchGet(t *testing.T) {
	resp := u.BatchGet([]string{"oLyBi0tDnybg0WFkhKsn5HRetX1I", "oLyBi0lCK5rQPuo0_cHJrjQ4J9XE"}, "")
	t.Log(*resp[0], *resp[1])
}

func TestUser_Get(t *testing.T) {
	resp := u.Get("")
	resp1 := u.Get("oLyBi0tDnybg0WFkhKsn5HRetX1I")
	t.Log(string(resp.Bytes()), string(resp1.Bytes()))
}
