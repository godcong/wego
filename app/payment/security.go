package payment

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Security struct {
	config.Config
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

func (s *Security) GetPublicKey() *net.Response {
	s.client.SetDataType(core.DataTypeXML)
	return s.client.SafeRequest(RISK_GETPUBLICKEY_URL_SUFFIX, util.Map{
		net.REQUEST_TYPE_XML.String(): s.preRequest(util.Map{"sign_type": "MD5"}),
	}, "post")
}
