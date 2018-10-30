package util

import "testing"

func TestMap_Set(t *testing.T) {
	m := make(Map)
	m.Set("one.two.three", "abc")

	t.Log(m.Get("one.two.three") == "abc")
}

func TestMap_Delete(t *testing.T) {
	m := make(Map)
	m.Set("one.two.three", "abc")
	if !m.Has("one.two") {
		t.Error("one.two")
	}
	m.Set("one.two.ab", "ddd")
	if m.GetString("one.two.three") != "abc" {
		t.Error("one.two.three")
	}

	if !m.Delete("one.two.ab") {
		t.Error("one.two.ab")
	}
}
