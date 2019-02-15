package wego

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/util"
)

// Sandbox ...
type Sandbox struct {
	*SandboxConfig
	subMchID string
	subAppID string
}

// NewSandbox ...
func NewSandbox(config *SandboxConfig, opts ...SandboxOption) *Sandbox {
	sandbox := &Sandbox{
		SandboxConfig: config,
	}
	return sandbox
}

func (obj *Sandbox) parse(opts []SandboxOption) {
	if opts == nil {
		return
	}
	obj.subAppID = opts[0].SubAppID
	obj.subMchID = opts[0].SubMchID
}

func (obj *Sandbox) getCacheKey() string {
	name := obj.AppID + "." + obj.MchID
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
	resp := PostXML(util.URL(apiMCHWeixin, sandboxNew, getSignKey), nil, m)
	return resp
}
