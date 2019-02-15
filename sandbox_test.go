package wego

import "testing"

// TestSandbox_SignKey ...
func TestSandbox_SignKey(t *testing.T) {
	sandbox := NewSandbox(&SandboxConfig{
		AppID:  "",
		Secret: "",
		MchID:  "",
		Key:    "",
		option: &SandboxOption{
			SubMchID: "",
			SubAppID: "",
		},
	})

	key := sandbox.SignKey()

	if key.Error() != nil {
		t.Error(key.Error(), key.ToMap())
	}
	t.Log(key.Result())
}
