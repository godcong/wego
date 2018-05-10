package official_account_test

import (
	"testing"

	"github.com/godcong/wego/app/official_account"
)

func TestBase_GetCallbackIp(t *testing.T) {
	base := official_account.NewBase()
	rlt := base.GetCallbackIp()
	t.Log(rlt)
}

func TestBase_ClearQuota(t *testing.T) {
	base := official_account.NewBase()
	rlt := base.ClearQuota()
	t.Log(rlt)
}
