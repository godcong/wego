package payment

import "github.com/godcong/wego/core"

type Security struct {
	core.Config
	*Payment
}

func (s *Security) GetPublicKey() core.Map {
	return s.GetClient().SafeRequest(core.RISK_GETPUBLICKEY_URL_SUFFIX, core.Map{"sign_type": "MD5"}, "post", nil)
}
