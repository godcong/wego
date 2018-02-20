package wego

import (
	"encoding/json"
	"sort"
	"strings"
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

func (m *Map) NullSet(s string, v string) *Map {
	if !m.Has(s) {
		(*m)[s] = v
	}
	return m
}

func (m *Map) Get(s string) string {
	if v, b := (*m)[s]; b {
		return v
	}
	return ""
}

func (m *Map) Delete(s string) {
	delete(*m, s)
}

func (m *Map) Has(s string) bool {
	_, b := (*m)[s]
	return b
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

func (m *Map) ToJson() []byte {
	v, e := json.Marshal(*m)
	if e != nil {
		return []byte(nil)
	}
	return v
}

func (m *Map) ToUrlQuery() string {
	var varr []string
	for k, v := range *m {
		varr = append(varr, strings.Join([]string{k, v}, "="))
	}
	return strings.Join(varr, "&")
}

func (m *Map) ToSortUrlQuery() string {
	var varr []string
	keys := m.SortKeys()
	for _, key := range keys {
		varr = append(varr, strings.Join([]string{key, (*m)[key]}, "="))
	}
	return strings.Join(varr, "&")
}

func (m *Map) join(source Map, replace bool) *Map {
	for k, v := range source {
		if replace {
			m.Set(k, v)
			continue
		}
		if !m.Has(k) {
			m.Set(k, v)
		}
	}
	return m
}

func (m *Map) ReplaceJoin(s Map) *Map {
	return m.join(s, true)
}

func (m *Map) Join(s Map) *Map {
	return m.join(s, false)
}

func (m *Map) SaveAs(p string, f string) {

}
