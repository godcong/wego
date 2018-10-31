package core

import "testing"

// TestNewURL ...
func TestNewURL(t *testing.T) {
	url := NewURL(DefaultConfig().GetSubConfig("official_account.default"))
	t.Log(string(url.ShortURL("http://y11e.com/test").Bytes()))
}
