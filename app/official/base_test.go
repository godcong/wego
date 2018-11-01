package official_test

import (
	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

// TestBase_GetCallbackIp ...
func TestBase_GetCallbackIp(t *testing.T) {
	base := official.NewBase(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	rlt := base.GetCallbackIP()
	t.Log(string(rlt.Bytes()))
	t.Log(rlt.ToMap())
}

// TestBase_ClearQuota ...
func TestBase_ClearQuota(t *testing.T) {
	base := official.NewBase(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	rlt := base.ClearQuota()
	t.Log(string(rlt.Bytes()))
	t.Log(rlt.ToMap())
}
