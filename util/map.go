package util

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/godcong/wego/log"
	"net/url"
	"sort"
	"strings"
)

// ErrNilMap ...
var ErrNilMap = errors.New("nil map")

/*StringAble StringAble */
type StringAble interface {
	String() string
}

/*String String */
type String string

/*String String */
func (s String) String() string {
	return string(s)
}

/*ToString ToString */
func ToString(s string) String {
	return String(s)
}

/*Map Map */
type Map map[string]interface{}

/*String transfer map to JSON string */
func (m Map) String() string {
	return string(m.ToJSON())
}

/*MapNilMake if m is nil result a nil map */
func MapNilMake(m Map) Map {
	if m == nil {
		return make(Map)
	}
	return m
}

/*MapToMap transfer to map[string]interface{} to Map  */
func MapToMap(p map[string]interface{}) Map {
	return Map(p)
}

// MapsToMap ...
func MapsToMap(p Map, m []Map) Map {
	if m == nil {
		return p
	}

	for _, v := range m {
		p.join(v, true)
	}
	return p
}

/*Set set interface */
func (m Map) Set(key string, v interface{}) Map {
	return m.SetPath(strings.Split(key, "."), v)
}

// SetPath is the same as SetPath, but allows you to provide comment
// information to the key, that will be reused by Marshal().
func (m Map) SetPath(keys []string, v interface{}) Map {
	subtree := m
	for _, intermediateKey := range keys[:len(keys)-1] {
		nextTree, exists := subtree[intermediateKey]
		if !exists {
			nextTree = make(Map)
			subtree[intermediateKey] = nextTree // add new element here
		}
		switch node := nextTree.(type) {
		case Map:
			subtree = node
		case []Map:
			// go to most recent element
			if len(node) == 0 {
				// create element if it does not exist
				subtree[intermediateKey] = append(node, make(Map))
			}
			subtree = node[len(node)-1]
		}
	}
	subtree[keys[len(keys)-1]] = v
	return m
}

/*SetNil set interface if key is not exist */
func (m Map) SetNil(s string, v interface{}) Map {
	if !m.Has(s) {
		m.Set(s, v)
	}
	return m
}

/*SetHas set interface if key is exist */
func (m Map) SetHas(s string, v interface{}) Map {
	if m.Has(s) {
		m.Set(s, v)
	}
	return m
}

/*SetGet set value from map if key is exist */
func (m Map) SetGet(s string, v Map) Map {
	if v.Has(s) {
		m.Set(s, v[s])
	}
	return m
}

/*Get get interface from map with out default */
func (m Map) Get(key string) interface{} {
	if key == "" {
		return nil
	}
	return m.GetPath(strings.Split(key, "."))
}

/*GetD get interface from map with default */
func (m Map) GetD(s string, d interface{}) interface{} {
	if v := m.Get(s); v != nil {
		return v
	}
	return d
}

/*GetMap get map from map with out default */
func (m Map) GetMap(s string) Map {
	switch v := m.Get(s).(type) {
	case map[string]interface{}:
		return (Map)(v)
	case Map:
		return v
	default:
		return nil
	}
}

/*GetMapD get map from map with default */
func (m Map) GetMapD(s string, d Map) Map {
	if v := m.GetMap(s); v != nil {
		return v
	}
	return d
}

/*GetMapArray get map from map with out default */
func (m Map) GetMapArray(s string) []Map {
	switch v := m.Get(s).(type) {
	case []Map:
		return v
	case []map[string]interface{}:
		var sub []Map
		for _, mp := range v {
			sub = append(sub, (Map)(mp))
		}
		return sub
	default:
		return nil
	}
}

/*GetMapArrayD get map from map with default */
func (m Map) GetMapArrayD(s string, d []Map) []Map {
	if v := m.GetMapArray(s); v != nil {
		return v
	}
	return d
}

// GetArray ...
func (m Map) GetArray(s string) []interface{} {
	switch v := m.Get(s).(type) {
	case []interface{}:
		return v
	default:
		return nil
	}
}

// GetArrayD ...
func (m Map) GetArrayD(s string, d []interface{}) []interface{} {
	switch v := m.Get(s).(type) {
	case []interface{}:
		return v
	default:
		return d
	}
}

/*GetBool get bool from map with out default */
func (m Map) GetBool(s string) bool {
	return m.GetBoolD(s, false)
}

/*GetBoolD get bool from map with default */
func (m Map) GetBoolD(s string, b bool) bool {
	if v, b := m.Get(s).(bool); b {
		return v
	}
	return b
}

/*GetNumber get float64 from map with out default */
func (m Map) GetNumber(s string) (float64, bool) {
	return ParseNumber(m.Get(s))
}

/*GetNumberD get float64 from map with default */
func (m Map) GetNumberD(s string, d float64) float64 {
	n, b := ParseNumber(m.Get(s))
	if b {
		return n
	}
	return d
}

/*GetInt64 get int64 from map with out default */
func (m Map) GetInt64(s string) (int64, bool) {
	return ParseInt(m.Get(s))
}

/*GetInt64D get int64 from map with default */
func (m Map) GetInt64D(s string, d int64) int64 {
	i, b := ParseInt(m.Get(s))
	if b {
		return i
	}
	return d
}

/*GetString get string from map with out default */
func (m Map) GetString(s string) string {
	if v, b := m.Get(s).(string); b {
		return v
	}
	return ""
}

/*GetStringD get string from map with default */
func (m Map) GetStringD(s string, d string) string {
	if v, b := m.Get(s).(string); b {
		return v
	}
	return d
}

/*GetBytes get bytes from map with default */
func (m Map) GetBytes(s string) []byte {
	if v, b := m.Get(s).([]byte); b {
		return v
	}
	return []byte(nil)
}

/*Delete delete if exist */
func (m Map) Delete(key string) bool {
	if key == "" {
		return false
	}
	return m.DeletePath(strings.Split(key, "."))
}

// DeletePath ...
func (m Map) DeletePath(keys []string) bool {
	if len(keys) == 0 {
		return false
	}
	subtree := m
	for _, intermediateKey := range keys[:len(keys)-1] {
		value, exists := subtree[intermediateKey]
		if !exists {
			return false
		}
		switch node := value.(type) {
		case Map:
			subtree = node
		case []Map:
			if len(node) == 0 {
				return false
			}
			subtree = node[len(node)-1]
		default:
			return false // cannot navigate through other node types
		}
	}
	// branch based on final node type
	if _, b := subtree[keys[len(keys)-1]]; !b {
		return false
	}
	delete(subtree, keys[len(keys)-1])
	return true
}

/*Has check if key exist */
func (m Map) Has(key string) bool {
	if key == "" {
		return false
	}
	return m.HasPath(strings.Split(key, "."))

}

// HasPath returns true if the given path of keys exists, false otherwise.
func (m Map) HasPath(keys []string) bool {
	return m.GetPath(keys) != nil
}

// GetPath returns the element in the tree indicated by 'keys'.
// If keys is of length zero, the current tree is returned.
func (m Map) GetPath(keys []string) interface{} {
	if len(keys) == 0 {
		return m
	}
	subtree := m
	for _, intermediateKey := range keys[:len(keys)-1] {
		value, exists := subtree[intermediateKey]
		if !exists {
			return nil
		}
		switch node := value.(type) {
		case Map:
			subtree = node
		case []Map:
			if len(node) == 0 {
				return nil
			}
			subtree = node[len(node)-1]
		default:
			return nil // cannot navigate through other node types
		}
	}
	// branch based on final node type
	return subtree[keys[len(keys)-1]]
}

/*SortKeys 排列key */
func (m Map) SortKeys() []string {
	var keys sort.StringSlice
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}

/*ToXML transfer map to XML */
func (m Map) ToXML() []byte {
	v, e := xml.Marshal(&m)
	if e != nil {
		log.Error(e)
		return []byte(nil)
	}
	return v
}

/*ParseXML parse XML bytes to map */
func (m Map) ParseXML(b []byte) {
	m.Join(XMLToMap(b))
}

/*ToJSON transfer map to JSON */
func (m Map) ToJSON() []byte {
	v, e := json.Marshal(m)
	if e != nil {
		log.Error(e)
		return []byte(nil)
	}
	return v
}

/*ParseJSON parse JSON bytes to map */
func (m Map) ParseJSON(b []byte) Map {
	tmp := Map{}
	if e := json.Unmarshal(b, &tmp); e == nil {
		m.Join(tmp)
	}
	return m
}

/*URLEncode transfer map to url encode */
func (m Map) URLEncode() string {
	var buf strings.Builder
	keys := m.SortKeys()
	size := len(keys)
	for i := 0; i < size; i++ {
		vs := m[keys[i]]
		keyEscaped := url.QueryEscape(keys[i])
		if v, b := vs.(string); b {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}

	return buf.String()
}

func (m Map) join(source Map, replace bool) Map {
	for k, v := range source {
		if !m.Has(k) || replace {
			m.Set(k, v)
		}
	}
	return m
}

// Append ...
func (m Map) Append(p Map) Map {
	for k, v := range p {
		if m.Has(k) {
			m.Set(k, []interface{}{m.Get(k), v})
		} else {
			m.Set(k, v)
		}
	}
	return m
}

/*ReplaceJoin insert map s to m with replace */
func (m Map) ReplaceJoin(s Map) Map {
	return m.join(s, true)
}

/*Join insert map s to m with out replace */
func (m Map) Join(s Map) Map {
	return m.join(s, false)
}

/*Only get map with keys */
func (m Map) Only(keys []string) Map {
	p := Map{}
	size := len(keys)
	for i := 0; i < size; i++ {
		p.Set(keys[i], m.Get(keys[i]))
	}

	return p
}

/*Expect get map expect keys */
func (m Map) Expect(keys []string) Map {
	p := m.Clone()
	size := len(keys)
	for i := 0; i < size; i++ {
		p.Delete(keys[i])
	}

	return p
}

/*Clone copy a map */
func (m Map) Clone() Map {
	v := deepCopy(m)
	return (v).(Map)
}

func deepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(Map); ok {
		newMap := make(Map)
		for k, v := range valueMap {
			newMap[k] = deepCopy(v)
		}
		return newMap
	} else if valueSlice, ok := value.([]Map); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = deepCopy(v)
		}
		return newSlice
	}

	return value
}

/*URLToSHA1 make sha1 from map */
func (m Map) URLToSHA1() string {
	return signatureSHA1(m)
}

//Range range all maps
func (m Map) Range(f func(key string, value interface{}) bool) {
	for k, v := range m {
		if !f(k, v) {
			return
		}
	}
}

//Check check all input keys
//return -1 if all is exist
//return index when not found
func (m Map) Check(s ...string) int {
	size := len(s)
	for i := 0; i < size; i++ {
		if !m.Has(s[i]) {
			return i
		}
	}

	return -1
}

// Map trans Map to map[string]interface
func (m Map) Map() map[string]interface{} {
	return (map[string]interface{})(m)
}

// MarshalXML ...
func (m Map) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return ErrNilMap
	}
	if start.Name.Local == "root" {
		return marshalXML(m, e, xml.StartElement{Name: xml.Name{Local: "root"}})
	}
	return marshalXML(m, e, xml.StartElement{Name: xml.Name{Local: "xml"}})
}

// UnmarshalXML ...
func (m Map) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if start.Name.Local == "root" {
		return unmarshalXML(m, d, xml.StartElement{Name: xml.Name{Local: "root"}}, false)
	}
	return unmarshalXML(m, d, xml.StartElement{Name: xml.Name{Local: "xml"}}, false)
}

// InterfaceToMap ...
func InterfaceToMap(v interface{}, m *Map) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, m)
}
