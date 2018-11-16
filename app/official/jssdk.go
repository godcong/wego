package official

import (
	"github.com/godcong/wego/core"
)

// JSSDK ...
type JSSDK struct {
	*core.JSSDK
	URL string
}

func newJSSDK(account *Account) *JSSDK {
	jssdk := &JSSDK{
		JSSDK: core.NewJSSDK(account.Config),
	}
	jssdk.CacheKey = jssdk.getCacheKey
	return jssdk
}

// NewJSSDK jssdk
func NewJSSDK(config *core.Config) *JSSDK {
	return newJSSDK(NewOfficialAccount(config))
}

func (j *JSSDK) getCacheKey() string {
	return "godcong.wego.official.account.jssdk.ticket.jsapi" + j.GetString("app_id")
}
