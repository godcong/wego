package payment

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
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
	//TODO
	log.Debug("TODO")
	key := cache.Get(s.cacheName())
	if key != nil {
		return key.(string)
	}

	return string(s.SandboxSignKey())
}

func (s *Sandbox) cacheName() string {
	name := s.GetString("app_id") + s.GetString("mch_id")
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

/*SandboxSignKey 沙箱key */
func (s *Sandbox) SandboxSignKey() []byte {
	m := make(util.Map)
	m.Set("mch_id", s.Get("mch_id"))
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := GenerateSignature(m, s.GetString("aes_key"), MakeSignMD5)
	m.Set("sign", sign)
	resp := s.client.RequestRaw(Link(sandboxSignKeyURLSuffix), "post", m)

	return resp

}
