package payment

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"time"
)

/*Sandbox 沙箱 */
type Sandbox struct {
	*Payment
}

func newSandbox(payment *Payment) *Sandbox {
	return &Sandbox{
		Payment: payment,
	}
}

/*NewSandbox NewSandbox */
func NewSandbox(config *core.Config) *Sandbox {
	return newSandbox(NewPayment(config))
}

/*GetKey 沙箱key(string类型) */
func (s *Sandbox) GetKey() string {
	key := cache.Get(s.getCacheKey())
	if key != nil {
		return key.(string)
	}

	response := s.SandboxSignKey().ToMap()


	if response.GetString("return_code") == "SUCCESS" {
		key := response.GetString("sandbox_signkey")
		ttl := time.Unix(24*3600, 0)
		cache.SetWithTTL(s.getCacheKey(), key, &ttl)
		return key
	}
	return ""
}

func (s *Sandbox) getCacheKey() string {
	name := s.GetString("app_id") + s.GetString("mch_id")
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

/*SandboxSignKey 沙箱key */
func (s *Sandbox) SandboxSignKey() core.Response {
	m := make(util.Map)
	m.Set("mch_id", s.Get("mch_id"))
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := GenerateSignature(m, s.GetString("key"), MakeSignMD5)
	m.Set("sign", sign)
	resp := s.client.PostXML(Link(sandboxSignKeyURLSuffix), nil, m)

	return resp

}
