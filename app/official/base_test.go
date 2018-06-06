package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

func TestBase_GetCallbackIp(t *testing.T) {
	base := official.NewBase()
	rlt := base.GetCallbackIP()
	t.Log(rlt)
}

func TestBase_ClearQuota(t *testing.T) {
	base := official.NewBase()
	rlt := base.ClearQuota()
	t.Log(rlt)
}
