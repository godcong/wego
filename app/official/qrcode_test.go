package official_test

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"

	"github.com/godcong/wego/app/official"
)

// TestQrCode_Create ...
func TestQrCode_Create(t *testing.T) {
	code := official.NewQrCode(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))

	resp := code.Create(
		&official.QrCodeAction{
			ExpireSeconds: 2592000,
			ActionName:    "QR_STR_SCENE",
			ActionInfo: official.QrCodeActionInfo{
				Scene: &official.QrCodeScene{
					// SceneID:  0,
					SceneStr: "igettheteickkka:///fdsaowkkkdfsaoowjkwodf",
				},
			},
		})
	t.Log(string(resp.Bytes()))
}

// TestQrCode_ShowQrCode ...
func TestQrCode_ShowQrCode(t *testing.T) {
	code := official.NewQrCode(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	resp := code.ShowQrCode("gQGd7zwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAybVNDTzBrTHdkeWkxVklIb3hxY3oAAgRsHrFaAwQAjScA")
	t.Log(string(resp.Bytes()))
	core.SaveTo(resp, "d:/text.jpg")
}
