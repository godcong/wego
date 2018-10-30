package util

import "testing"

func TestMap_Set(t *testing.T) {
	m := make(Map)
	m.Set("one.two.three", "abc")

	t.Log(m.Get("one.two.three") == "abc")
}
