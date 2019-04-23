package wego

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/util"
	"strings"
)

// Sandbox ...
type Sandbox struct {
	*SandboxProperty
	subMchID string
	subAppID string
}

// NewSandbox ...
func NewSandbox(config *SandboxProperty, options ...SandboxOption) *Sandbox {
	sandbox := &Sandbox{
		SandboxProperty: config,
	}

	sandbox.parse(options...)
	return sandbox
}

func (obj *Sandbox) parse(options ...SandboxOption) {
	if options == nil {
		return
	}
	for _, o := range options {
		o(obj)
	}
}

func (obj *Sandbox) getCacheKey() string {
	name := strings.Join([]string{obj.AppID, obj.MchID}, ".")
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

// SignKey ...
func (obj *Sandbox) SignKey() Responder {
	m := make(util.Map)
	m.Set("mch_id", obj.MchID)
	m.Set("nonce_str", util.GenerateNonceStr())
	m.Set("sign", util.GenSign(m, obj.Key))
	if obj.subMchID != "" {
		m.Set("sub_mch_id", obj.subMchID)
	}
	if obj.subAppID != "" {
		m.Set("sub_appid", obj.subAppID)
	}
	return PostXML(util.URL(apiMCHWeixin, sandboxNew, getSignKey), nil, m)
}
