package official_account_test

import (
	"testing"

	"github.com/godcong/wego/app/official_account"
)

var q = official_account.NewQrCode()

func TestNewQrCode(t *testing.T) {

}

func TestQrCode_Create(t *testing.T) {
	resp := q.Create(
		&official_account.QrCodeAction{
			ExpireSeconds: 2592000,
			ActionName:    "QR_STR_SCENE",
			ActionInfo: official_account.QrCodeActionInfo{
				Scene: &official_account.QrCodeScene{
					// SceneId:  0,
					SceneStr: "igettheteickkka:///fdsaowkkkdfsaoowjkwodf",
				},
			},
		})
	t.Log(resp.ToString())
}

func TestQrCode_ShowQrCode(t *testing.T) {
	resp := q.ShowQrCode("gQGd7zwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAybVNDTzBrTHdkeWkxVklIb3hxY3oAAgRsHrFaAwQAjScA")
	t.Log(resp.ToString())
	resp.ToFile("d:/test.jpg")
}
