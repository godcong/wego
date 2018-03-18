package official_account_test

import (
	"testing"

	"github.com/godcong/wego/official_account"
)

var u = official_account.NewUser()

func TestNewUser(t *testing.T) {

}

func TestUser_UpdateRemark(t *testing.T) {
	resp := u.UpdateRemark("oLyBi0tDnybg0WFkhKsn5HRetX1I", "nishi123")
	t.Log(resp.ToString())
}

func TestUser_UserInfo(t *testing.T) {
	resp := u.UserInfo("oLyBi0tDnybg0WFkhKsn5HRetX1I", "zh_CN")
	t.Log(*resp)
}

func TestUser_BatchGet(t *testing.T) {
	resp := u.BatchGet([]string{"oLyBi0tDnybg0WFkhKsn5HRetX1I", "oLyBi0lCK5rQPuo0_cHJrjQ4J9XE"}, "")
	t.Log(*resp[0], *resp[1])
}
