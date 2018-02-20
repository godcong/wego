package wego

type Security interface {
	GetPublicKey() Map
}

type security struct {
	Config
	Payment
}

func NewSecurity(application Application, config Config) Security {
	return &security{
		Config:  config,
		Payment: application.Payment(),
	}
}

func (s *security) GetPublicKey() Map {
	return s.Payment.Client().SafeRequest(RISK_GETPUBLICKEY_URL_SUFFIX, Map{"sign_type": "MD5"}, "post", nil)
}
