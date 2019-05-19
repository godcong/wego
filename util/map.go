package util

import (
	"bytes"
	"encoding/xml"
	"errors"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// MapAble ...
type MapAble interface {
	ToMap() Map
}

// XMLAble ...
type XMLAble interface {
	ToXML() []byte
}

// JSONAble ...
type JSONAble interface {
	ToJSON() []byte
}

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

// Map ...
type Map map[string]interface{}

/*String transfer map to JSON string */
func (m Map) String() string {
	return string(m.ToJSON())
}

// StructToMap ...
func StructToMap(s interface{}, p Map) error {
	var err error
	buf := bytes.NewBuffer(nil)
	enc := jsoniter.NewEncoder(buf)
	err = enc.Encode(s)
	if err != nil {
		return err
	}
	dec := jsoniter.NewDecoder(buf)
	err = dec.Decode(&p)
	if err != nil {
		return err
	}
	return nil
}

/*MapMake make new map only if m is nil result a new map with nothing */
func MapMake(m Map) Map {
	if m == nil {
		return make(Map)
	}
	return m
}

/*ToMap transfer to map[string]interface{} or MapAble to GMap  */
func ToMap(p interface{}) Map {
	switch v := p.(type) {
	case map[string]interface{}:
		return Map(v)
	case MapAble:
		return v.ToMap()
	}
	return nil
}

// CombineMaps ...
func CombineMaps(p Map, m ...Map) Map {
	if p == nil {
		p = make(Map)
	}
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
	v, e := jsoniter.Marshal(m)
	if e != nil {
		log.Error(e)
		return []byte(nil)
	}
	return v
}

/*ParseJSON parse JSON bytes to map */
func (m Map) ParseJSON(b []byte) Map {
	tmp := Map{}
	if e := jsoniter.Unmarshal(b, &tmp); e == nil {
		m.Join(tmp)
	}
	return m
}

// URLValues ...
func (m Map) URLValues() url.Values {
	val := url.Values{}
	for key, value := range m {
		m.Set(key, value)
	}
	return val
}

/*URLEncode transfer map to url encode */
func (m Map) URLEncode() string {
	var buf strings.Builder
	keys := m.SortKeys()
	size := len(keys)
	for i := 0; i < size; i++ {
		vs := m[keys[i]]
		keyEscaped := url.QueryEscape(keys[i])
		switch val := vs.(type) {
		case string:
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(val))
		case []string:
			for _, v := range val {
				if buf.Len() > 0 {
					buf.WriteByte('&')
				}
				buf.WriteString(keyEscaped)
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(v))
			}
		}
	}

	return buf.String()
}

func (m Map) join(source Map, replace bool) Map {
	for k, v := range source {
		if replace || !m.Has(k) {
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

/*SignatureSHA1 make sha1 from map */
func (m Map) SignatureSHA1() string {
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

// GoMap trans return a map[string]interface from Map
func (m Map) GoMap() map[string]interface{} {
	return (map[string]interface{})(m)
}

// ToMap implements MapAble
func (m Map) ToMap() Map {
	return m
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

func marshalXML(maps Map, e *xml.Encoder, start xml.StartElement) error {
	if maps == nil {
		return errors.New("map is nil")
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	for k, v := range maps {
		err := convertXML(k, v, e, xml.StartElement{Name: xml.Name{Local: k}})
		if err != nil {
			return err
		}
	}
	return e.EncodeToken(start.End())
}

func unmarshalXML(maps Map, d *xml.Decoder, start xml.StartElement, needCast bool) error {
	current := ""
	var data interface{}
	last := ""
	arrayTmp := make(Map)
	arrayTag := ""
	var ele []string

	for t, err := d.Token(); err == nil; t, err = d.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			if strings.ToLower(token.Name.Local) == "xml" ||
				strings.ToLower(token.Name.Local) == "root" {
				continue
			}
			ele = append(ele, token.Name.Local)
			current = strings.Join(ele, ".")
			//log.Debug("EndElement", current)
			//log.Debug("EndElement", last)
			//log.Debug("EndElement", arrayTag)
			if current == last {
				arrayTag = current
				tmp := maps.Get(arrayTag)
				switch tmp.(type) {
				case []interface{}:
					arrayTmp.Set(arrayTag, tmp)
				default:
					arrayTmp.Set(arrayTag, []interface{}{tmp})
				}
				maps.Delete(arrayTag)
			}
			log.Debug("StartElement", ele)
			// 处理元素结束（标签）
		case xml.EndElement:
			name := token.Name.Local
			// fmt.Printf("This is the end: %s\n", name)
			if strings.ToLower(name) == "xml" ||
				strings.ToLower(name) == "root" {
				break
			}
			last = strings.Join(ele, ".")
			//log.Debug("EndElement", current)
			//log.Debug("EndElement", last)
			//log.Debug("EndElement", arrayTag)

			if current == last {
				if data != nil {
					log.Debug("CharData", data)
					maps.Set(current, data)
				} else {
					//m.Set(current, nil)
				}
				data = nil
			}
			if last == arrayTag {
				arr := arrayTmp.GetArray(arrayTag)
				if arr != nil {
					if v := maps.Get(arrayTag); v != nil {
						maps.Set(arrayTag, append(arr, v))
					} else {
						maps.Set(arrayTag, arr)
					}
				} else {
					//exception doing
					maps.Set(arrayTag, []interface{}{maps.Get(arrayTag)})
				}
				arrayTmp.Delete(arrayTag)
				arrayTag = ""
			}

			ele = ele[:len(ele)-1]
			//log.Debug("EndElement", ele)
			// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			if needCast {
				data, err = strconv.Atoi(string(token))
				if err == nil {
					continue
				}

				data, err = strconv.ParseFloat(string(token), 64)
				if err == nil {
					continue
				}

				data, err = strconv.ParseBool(string(token))
				if err == nil {
					continue
				}
			}

			data = string(token)
			//log.Debug("CharData", data)
			// 异常处理(Log输出）
		default:
			log.Debug(token)
		}

	}

	return nil
}

func convertXML(k string, v interface{}, e *xml.Encoder, start xml.StartElement) error {
	var err error
	switch v1 := v.(type) {
	case Map:
		return marshalXML(v1, e, xml.StartElement{Name: xml.Name{Local: k}})
	case map[string]interface{}:
		return marshalXML(v1, e, xml.StartElement{Name: xml.Name{Local: k}})
	case string:
		if _, err := strconv.ParseInt(v1, 10, 0); err != nil {
			err = e.EncodeElement(
				CDATA{Value: v1}, xml.StartElement{Name: xml.Name{Local: k}})
			return err
		}
		err = e.EncodeElement(v1, xml.StartElement{Name: xml.Name{Local: k}})
		return err
	case float64:
		if v1 == float64(int64(v1)) {
			err = e.EncodeElement(int64(v1), xml.StartElement{Name: xml.Name{Local: k}})
			return err
		}
		err = e.EncodeElement(v1, xml.StartElement{Name: xml.Name{Local: k}})
		return err
	case bool:
		err = e.EncodeElement(v1, xml.StartElement{Name: xml.Name{Local: k}})
		return err
	case []interface{}:
		size := len(v1)
		for i := 0; i < size; i++ {
			err := convertXML(k, v1[i], e, xml.StartElement{Name: xml.Name{Local: k}})
			if err != nil {
				return err
			}
		}
		if len(v1) == 1 {
			return convertXML(k, "", e, xml.StartElement{Name: xml.Name{Local: k}})
		}
	default:
		log.Error(v1)
	}
	return nil
}
