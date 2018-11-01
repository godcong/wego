package core_test

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

var config *core.Config

// TestBase_GetCallbackIP ...
func TestBase_GetCallbackIP(t *testing.T) {
	base := core.NewBase(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	resp := base.GetCallbackIP()
	t.Log(resp.Error())
	t.Log(resp.ToMap())
	t.Log(string(resp.Bytes()))
}

// TestBase_ClearQuota ...
func TestBase_ClearQuota(t *testing.T) {
	base := core.NewBase(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	resp := base.ClearQuota()
	t.Log(resp.Error())
	t.Log(resp.ToMap())
	t.Log(string(resp.Bytes()))
}
