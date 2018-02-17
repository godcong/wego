package wego

import "encoding/json"

type Map map[string]string

func (m *Map) String() string {
	if v, e := json.Marshal(m); e == nil {
		return string(v)
	}
	return ""
}

func (m *Map) Set(s string, v string) *Map {
	(*m)[s] = v
	return m
}

func (m *Map) Get(s string) string {
	if v, b := (*m)[s]; b {
		return v
	}
	return ""
}
