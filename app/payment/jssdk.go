package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"strings"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*core.JSSDK
}

func newJSSDK(p *Payment) interface{} {
	jssdk := &JSSDK{
		JSSDK: core.NewJSSDK(p.Config),
	}
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
