package util

import "testing"

var urlTest = []struct {
	Prefix  string
	URI     string
	Success string
}{
	{
		Prefix:  "http://localhost",
		URI:     "/localaddress",
		Success: "http://localhost/localaddress",
	},
	{
		Prefix:  "http://localhost/",
		URI:     "localaddress",
		Success: "http://localhost/localaddress",
	},
	{
		Prefix:  "http://localhost",
		URI:     "localaddress",
		Success: "http://localhost/localaddress",
	},
	{
		Prefix:  "http://localhost/",
		URI:     "/localaddress",
		Success: "http://localhost/localaddress",
	},
}

// TestURL ...
func TestURL(t *testing.T) {
	for _, val := range urlTest {
		res := URL(val.Prefix, val.URI)
		if res != val.Success {
			t.Error(val.Prefix, val.URI, res, val.Success)
		}
	}
}
