package wego

import (
	"encoding/json"
	"sort"
)

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

func (m *Map) SortKeys() []string {
	var keys sort.StringSlice
	for k := range *m {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

func (m *Map) ToXml() string {
	if v, e := MapToXml(*m); e == nil {
		return v
	}
	return ""

}
