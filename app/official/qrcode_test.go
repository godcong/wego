package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

var q = official.NewQrCode(config)

func TestNewQrCode(t *testing.T) {

}

func TestQrCode_Create(t *testing.T) {
	resp := q.Create(
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

func TestQrCode_ShowQrCode(t *testing.T) {
	resp := q.ShowQrCode("gQGd7zwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAybVNDTzBrTHdkeWkxVklIb3hxY3oAAgRsHrFaAwQAjScA")
	t.Log(string(resp.Bytes()))
	//resp.ToFile("d:/test.jpg")
}
