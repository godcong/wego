package payment

import (
	"crypto/md5"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
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
	str := s.GetString("app_id") + s.GetString("mch_id")
	key := cache.Get("godcong.wego.payment.sandbox." + string(md5.Sum([]byte(str))[:]))
	if key != nil {
		return key.(string)
	}

	//TODO

	return string(s.SandboxSignKey())
}

//func (s *Sandbox) GetCacheKey() string {
//	return ""
//}

/*SandboxSignKey 沙箱key */
func (s *Sandbox) SandboxSignKey() []byte {
	m := make(util.Map)
	m.Set("mch_id", s.Get("mch_id"))
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := core.GenerateSignature(m, s.GetString("aes_key"), core.MakeSignMD5)
	m.Set("sign", sign)
	resp := s.client.RequestRaw(Link(sandboxSignKeyURLSuffix), "post", m)

	return resp

}
