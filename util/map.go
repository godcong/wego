package util

import (
	"encoding/json"
	"net/url"
	"sort"
)

type StringAble interface {
	String() string
}

type String string

func (s *String) String() string {
	return string(*s)
}

func ToString(s string) String {
	return String(s)
}

type Map map[string]interface{}

func (m *Map) String() string {
	return string(m.ToJson())
}

func MapNilMake(m Map) Map {
	if m == nil {
		return make(Map)
	}
	return m
}

func (m *Map) Set(s string, v interface{}) *Map {
	(*m)[s] = v
	return m
}

func (m *Map) NilSet(s string, v interface{}) *Map {
	if !m.Has(s) {
		m.Set(s, v)
	}
	return m
}

func (m *Map) HasSet(s string, v interface{}) *Map {
	if m.Has(s) {
		m.Set(s, v)
	}
	return m
}

func (m *Map) Get(s string) interface{} {
	if v, b := (*m)[s]; b {
		return v
	}
	return nil
}

func (m *Map) GetD(s string, d interface{}) interface{} {
	if v, b := (*m)[s]; b {
		return v
	}
	return d
}

func (m *Map) GetMap(s string) Map {
	if v, b := m.Get(s).(map[string]interface{}); b {
		return v
	}

	// if v, b := m.Get(s).(Map); b {
	// 	return v
	// }
	return nil
}

func (m *Map) GetMapD(s string, d Map) Map {
	if v := m.GetMap(s); v != nil {
		return v
	}
	return d
}

func (m *Map) GetNumber(s string) float64 {
	return ParseNumber(m.Get(s))
}

func (m *Map) GetNumberD(s string, i float64) float64 {
	n := m.GetNumber(s)
	if n != 0 {
		return n
	}
	return i
}

func (m *Map) GetInt64(s string) int64 {
	return ParseInt(m.Get(s))
}

func (m *Map) GetString(s string) string {
	if v, b := m.Get(s).(string); b {
		return v
	}
	return ""
}

func (m *Map) GetBytes(s string) []byte {
	if v, b := m.Get(s).([]byte); b {
		return v
	}
	return []byte(nil)
}

func (m *Map) GetStringD(s string, d string) string {
	if v, b := m.Get(s).(string); b {
		return v
	}
	return d
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

func (m *Map) ParseXml(b []byte) {
	m.Join(XmlToMap(b))
}

func (m *Map) ToJson() []byte {
	v, e := json.Marshal(*m)
	if e != nil {
		return []byte(nil)
	}
	return v
}

func (m *Map) ParseJson(b []byte) Map {
	tmp := Map{}
	if e := json.Unmarshal(b, &tmp); e == nil {
		m.Join(tmp)
	}
	return *m
}

func (m *Map) UrlEncode() string {
	url := url.Values{}
	for key, v := range *m {
		if v0, b := v.(string); b {
			url.Add(key, v0)
		}
	}
	return url.Encode()
}

func (m *Map) join(source Map, replace bool) *Map {
	for k, v := range source {
		if _, b := (*m)[k]; replace || !b {
			(*m)[k] = v
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

func (m *Map) Only(columns []string) Map {
	p := Map{}
	for _, v := range columns {
		p.Set(v, m.Get(v))
	}
	return p
}

func (m *Map) Clone() Map {
	m0 := make(Map)
	for k, v := range *m {
		m0[k] = v
	}
	return m0
}

func (m *Map) UrlSHA1() string {
	return signatureSHA1(*m)
}
