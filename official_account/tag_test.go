package official_account_test

import (
	"testing"

	"github.com/godcong/wego/official_account"
)

var t0 = official_account.NewTag()

func TestNewTag(t *testing.T) {
	resp := t0.Create("testtag")
	t.Log(resp.ToString())
}

func TestTag_Get(t *testing.T) {
	resp := t0.Get()
	t.Log(resp.ToString())
}

func TestTag_Update(t *testing.T) {
	resp := t0.Update(100, "changetag")
	t.Log(resp.ToString())
}

func TestTag_Delete(t *testing.T) {
	resp := t0.Delete(2)
	t.Log(resp.ToString())
}

func TestTag_UserTagGet(t *testing.T) {
	resp := t0.UserTagGet(2, "")
	t.Log(resp.ToString())
}
