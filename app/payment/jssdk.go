package payment

import (
	"github.com/godcong/wego/core"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*core.JSSDK
}

func newJSSDK(p *Payment) interface{} {
	jssdk := &JSSDK{
		JSSDK: core.NewJSSDK(p.Config),
	}
	jssdk.CacheKey = jssdk.getCacheKey
	return jssdk
}

/*NewJSSDK NewJSSDK */
func NewJSSDK(config *core.Config) *JSSDK {
	jssdk := newJSSDK(NewPayment(config)).(*JSSDK)

	return jssdk
}

func (j *JSSDK) getURL() string {
	return core.GetServerIP()
}
func (j *JSSDK) getCacheKey() string {
	return "godcong.wego.payment.jssdk.ticket.jsapi" + j.GetString("app_id")
}
