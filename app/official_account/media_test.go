package official_account_test

import (
	"testing"

	"github.com/godcong/wego/app/official_account"
	"github.com/godcong/wego/net"
)

func TestMedia_Upload(t *testing.T) {
	media := official_account.NewMedia()
	var resp *net.Response
	// resp = media.UploadImage(`test.jpg`)
	resp = media.UploadImage(`D:\temp\微信图片_20180516164809.jpg`)
	t.Log(resp.ToString())

	// resp = media.UploadVoice(`D:\temp\3.mp3`)
	// t.Log(resp.ToString())
	return

}

func TestMedia_Get(t *testing.T) {
	media := official_account.NewMedia()
	resp := media.Get("9fCk1Any5VcwmbJPzGztWMq3a1PsWv11KpgLTdM_YXgIlwdAUosdeSI_M6M7Qtwb")
	t.Log(resp.ToString())
	return
}

func TestMedia_GetJssdk(t *testing.T) {
	media := official_account.NewMedia()
	resp := media.GetJssdk("JLqX5-WgxC5k7zu91j54HupFziaCqsfGfrIzOclTs_CHvvmPLJ5bdIZtBfI-pgYF")
	t.Log(resp.ToString())
	return
}
