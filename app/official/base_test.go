package official_test

import (
	"github.com/godcong/wego/app/official"
	"testing"
)

// TestBase_GetCallbackIp ...
func TestBase_GetCallbackIp(t *testing.T) {
	base := official.NewBase(config)
	rlt := base.GetCallbackIP()
	t.Log(string(rlt.Bytes()))
	t.Log(rlt.ToMap())
}

// TestBase_ClearQuota ...
func TestBase_ClearQuota(t *testing.T) {
	base := official.NewBase(config)
	rlt := base.ClearQuota()
	t.Log(string(rlt.Bytes()))
	t.Log(rlt.ToMap())
}
