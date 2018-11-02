package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

// TestJSSDK_BuildConfig ...
func TestJSSDK_BuildConfig(t *testing.T) {
	js := NewJSSDK(core.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	js.SetURL("https://mp.quick58.com")

	resp := js.BuildConfig([]string{"onMenuShareQQ", "onMenuShareWeibo"})
	resp = js.BuildConfig([]string{"onMenuShareQQ", "onMenuShareWeibo"})
	resp = js.BuildConfig([]string{"onMenuShareQQ", "onMenuShareWeibo"})
	t.Log(resp)
}
