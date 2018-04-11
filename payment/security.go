package payment

import (
	"github.com/godcong/wego/core"
)

type Security struct {
	core.Config
	*Payment
}

func newSecurity(pay *Payment) *Security {
	return &Security{
		Config:  defaultConfig,
		Payment: pay,
	}
}

func NewSecurity() *Security {
	return newSecurity(payment)
}

func (s *Security) GetPublicKey() *core.Response {
	m := s.preRequest(core.Map{"sign_type": "MD5"})
	s.client.SetDataType(core.DATA_TYPE_XML)
	return s.client.SafeRequest(RISK_GETPUBLICKEY_URL_SUFFIX, m, "post", nil)
}
