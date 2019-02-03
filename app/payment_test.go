package app

import "testing"

var p = Property{}

// TestPayment_SandboxSignKey ...
func TestPayment_SandboxSignKey(t *testing.T) {
	payment := NewPayment(&p, &PaymentOption{
		UsePayment: true,
		Host:       "",
	})
	key := payment.SandboxSignKey().ToMap()
	t.Log(key)
	if !key.Has("return_code") || !key.Has("return_msg") {
		t.Error(key)
	}
}
