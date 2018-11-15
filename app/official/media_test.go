package official_test

import (
	"github.com/godcong/wego/core"
	"testing"

	"github.com/godcong/wego/app/official"
)

func TestMedia_Upload(t *testing.T) {
	media := official.NewMedia(config)
	var resp core.Responder
	// resp = media.UploadImage(`test.jpg`)
	resp = media.UploadImage(`D:\temp\微信图片_20180516164809.jpg`)
	t.Log(string(resp.Bytes()))

	// resp = media.UploadVoice(`D:\temp\3.mp3`)
	// t.Log(resp.ToString())
	return

}

func TestMedia_Get(t *testing.T) {
	media := official.NewMedia(config)
	resp := media.Get("9fCk1Any5VcwmbJPzGztWMq3a1PsWv11KpgLTdM_YXgIlwdAUosdeSI_M6M7Qtwb")
	t.Log(string(resp.Bytes()))
}

func TestMedia_GetJssdk(t *testing.T) {
	media := official.NewMedia(config)
	resp := media.GetJssdk("JLqX5-WgxC5k7zu91j54HupFziaCqsfGfrIzOclTs_CHvvmPLJ5bdIZtBfI-pgYF")
	t.Log(string(resp.Bytes()))
}
