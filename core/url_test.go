package core_test

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

// TestNewURL ...
func TestNewURL(t *testing.T) {
	url := core.NewURL(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))

	resp := url.ShortURL("http://y11e.com/test")
	if resp.Error() != nil {
		t.Error(resp.Error())
	}
	if v, b := resp.ToMap().GetInt64("errcode"); b && v != 0 {
		t.Error("wrong code", v, b)
	}
	t.Log(resp.ToMap())
}
