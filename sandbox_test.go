package wego

import "testing"

// TestSandbox_SignKey ...
func TestSandbox_SignKey(t *testing.T) {
	sandbox := NewSandbox(&SandboxProperty{
		AppID:     "",
		AppSecret: "",
		MchID:     "",
		Key:       "",
	}, SandboxSubID("mch", "app"))

	key := sandbox.SignKey()

	if key.Error() != nil {
		t.Error(key.Error(), key.ToMap())
	}
	t.Log(key.Result())
}
