package wego

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/util"
)

// Sandbox ...
type Sandbox struct {
	*SandboxConfig
	SubMchID string
	SubAppID string
}

// NewSandbox ...
func NewSandbox(config *SandboxConfig) *Sandbox {
	sandbox := &Sandbox{
		SandboxConfig: config,
	}
	return sandbox
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
	if obj.SubMchID != "" {
		m.Set("sub_mch_id", obj.SubMchID)
	}
	if obj.SubAppID != "" {
		m.Set("sub_appid", obj.SubAppID)
	}
	resp := PostXML(util.URL(apiMCHWeixin, sandboxNew, getSignKey), nil, m)
	return resp
}
