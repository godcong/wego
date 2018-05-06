package core

import (
	"github.com/godcong/wego/core/config"
	"github.com/godcong/wego/core/util"
)

type Sandbox struct {
	config.Config
	client *Client
}

func NewSandbox(config config.Config) *Sandbox {
	return &Sandbox{
		Config: config,
		client: NewClient(config),
	}
}

func (s *Sandbox) GetKey() string {
	return string(s.SandboxSignKey())
}

func (s *Sandbox) GetCacheKey() string {
	return ""
}

func (s *Sandbox) SandboxSignKey() []byte {
	m := make(util.Map)
	m.Set("mch_id", s.Get("mch_id"))
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := GenerateSignature(m, s.Get("aes_key"), MakeSignMD5)
	m.Set("sign", sign)
	resp := s.client.Request(s.client.Link(SANDBOX_SIGNKEY_URL_SUFFIX), m, "post")

	return resp.ToBytes()

}
