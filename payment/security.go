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
	s.client.SetDataType(core.DATA_TYPE_XML)
	return s.client.SafeRequest(RISK_GETPUBLICKEY_URL_SUFFIX, core.Map{
		core.REQUEST_TYPE_XML.String(): s.preRequest(core.Map{"sign_type": "MD5"}),
	}, "post")
}
