package app

import "testing"

var p = Property{
	OAuth:           nil,
	OpenPlatform:    nil,
	OfficialAccount: nil,
	MiniProgram:     nil,
	Payment: &PaymentProperty{
		AppID:      "",
		MchID:      "",
		Key:        "",
		CertPEM:    "",
		KeyPEM:     "",
		RootCaPEM:  "",
		PublicKey:  "",
		PrivateKey: "",
	},
}

// TestPayment_SandboxSignKey ...
func TestPayment_SandboxSignKey(t *testing.T) {
	payment := NewPayment(&p, &PaymentOption{
		Sandbox: SandboxProperty{
			UseSandbox: true,
		},
	})
	key := payment.SandboxSignKey().ToMap()
	t.Log(key)
	if !key.Has("return_code") || !key.Has("return_msg") {
		t.Error(key)
	}
}
