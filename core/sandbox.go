package core

type Sandbox struct {
	Config
	client *Client
}

func NewSandbox(config Config) *Sandbox {
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
	m := make(Map)
	m.Set("mch_id", s.Get("mch_id"))
	m.Set("nonce_str", GenerateNonceStr())
	sign := GenerateSignature(m, s.Get("aes_key"), SIGN_TYPE_MD5)
	m.Set("sign", sign)
	resp := s.client.Request(s.client.Link(SANDBOX_SIGNKEY_URL_SUFFIX), m, "post")

	return resp.ToBytes()

}
