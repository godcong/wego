package official_account_test

import (
	"testing"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/official_account"
)

func TestBase_GetCallbackIp(t *testing.T) {
	o := core.App().Get("official_account").(official_account.OfficialAccount)
}
