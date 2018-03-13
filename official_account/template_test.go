package official_account_test

import (
	"testing"

	"github.com/godcong/wego/official_account"
)

func TestNewTemplate(t *testing.T) {
	t0 := official_account.NewTemplate()

	testTemplate_SetIndustry(t, t0)
	testTemplate_GetIndustry(t, t0)
	testTemplate_AddTemplate(t, t0)
}

func testTemplate_SetIndustry(t *testing.T, template *official_account.Template) {
	rlt := template.SetIndustry("1", "4")
	t.Log(rlt.ToString())
}

func testTemplate_GetIndustry(t *testing.T, template *official_account.Template) {
	rlt := template.GetIndustry()
	t.Log(rlt.ToString())
}

func testTemplate_AddTemplate(t *testing.T, template *official_account.Template) {
	rlt := template.AddTemplate("TM00015")
	t.Log(rlt.ToString())
}
