package core

import (
	"github.com/godcong/wego/util"
)

/*Sandbox 沙箱 */
type Sandbox struct {
	Config
	client *Client
}

/*NewSandbox NewSandbox */
func NewSandbox(config Config) *Sandbox {
	return &Sandbox{
		Config: config,
		client: NewClient(config),
	}
}

/*GetKey 沙箱key(string类型) */
func (s *Sandbox) GetKey() string {
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
	sign := GenerateSignature(m, s.Get("aes_key"), MakeSignMD5)
	m.Set("sign", sign)
	resp := s.client.Request(s.client.domain.Link(sandboxSignKeyURLSuffix), m, "post")

	return resp.ToBytes()

}
