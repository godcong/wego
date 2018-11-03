package official_test

import (
	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

func TestCurrent_AutoReplyInfo(t *testing.T) {
	current := official.NewCurrent(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	rlt := current.AutoReplyInfo()
	t.Log(string(rlt.Bytes()))
	t.Log(rlt.ToMap())
}
