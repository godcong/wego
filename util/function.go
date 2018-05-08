package util

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/godcong/wego/log"
	uuid "github.com/satori/go.uuid"
)

const CUSTOM_HEADER = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`

type CDATA struct {
	Value string `xml:",cdata"`
}

var (
	ErrorSignType  = errors.New("sign type error")
	ErrorParameter = errors.New("JsonApiParameters() check error")
	ErrorToken     = errors.New("EditAddressParameters() token is nil")
)

type RandomKind int

const (
	T_RAND_NUM      RandomKind = iota // 纯数字
	T_RAND_LOWER                      // 小写字母
	T_RAND_UPPER                      // 大写字母
	T_RAND_LOWERNUM                   // 数字、小写字母
	T_RAND_UPPERNUM                   // 数字、大写字母
	T_RAND_ALL                        // 数字、大小写字母
)

var (
	RandomString = map[RandomKind]string{
		T_RAND_NUM:      "0123456789",
		T_RAND_LOWER:    "abcdefghijklmnopqrstuvwxyz",
		T_RAND_UPPER:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		T_RAND_LOWERNUM: "0123456789abcdefghijklmnopqrstuvwxyz",
		T_RAND_UPPERNUM: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		T_RAND_ALL:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
)

func ParseNumber(v interface{}) float64 {
	switch v.(type) {
	case float64:
		return v.(float64)
	case float32:
		return float64(v.(float32))
	}
	return 0
}

func ParseInt(v interface{}) int64 {
	switch v0 := v.(type) {
	case int:
		return int64(v0)
	case int32:
		return int64(v0)
	case int64:
		return int64(v0)
	case uint:
		return int64(v0)
	case uint32:
		return int64(v0)
	case uint64:
		return int64(v0)
	default:
	}
	return 0
}

// util.MapToXml Convert MAP to XML
func MapToXml(m Map) (string, error) {
	return mapToXml(m, false)
}

func mapToXml(m Map, needHeader bool) (string, error) {

	buff := bytes.NewBuffer([]byte(CUSTOM_HEADER))
	if needHeader {
		buff.Write([]byte(xml.Header))
	}

	enc := xml.NewEncoder(buff)

	enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "xml"}})
	for k, v := range m {
		if v0, b := v.(string); b {
			if _, err := strconv.ParseInt(v0, 10, 0); err != nil {
				enc.EncodeElement(
					CDATA{Value: v0}, xml.StartElement{Name: xml.Name{Local: k}})
				continue
			}
		}
		enc.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: k}})
	}
	enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "xml"}})
	enc.Flush()
	return buff.String(), nil
}

// XmlToMap Convert XML to MAP
func XmlToMap(contentXml []byte) Map {
	return xmlToMap(contentXml, false)
}

func JsonToMap(content []byte) Map {
	m := Map{}
	json.Unmarshal(content, &m)
	return m
}

func xmlToMap(contentXml []byte, hasHeader bool) Map {
	m := make(Map)
	dec := xml.NewDecoder(bytes.NewReader(contentXml))
	ele, val := "", ""

	for t, err := dec.Token(); err == nil; t, err = dec.Token() {
		switch token := t.(type) {
		// 处理元素开始（标签）
		case xml.StartElement:
			ele = token.Name.Local
			// fmt.Printf("This is the sta: %s\n", ele)
			if strings.ToLower(ele) == "xml" {
				// xmlFlag = true
				continue
			}

			// 处理元素结束（标签）
		case xml.EndElement:
			name := token.Name.Local
			// fmt.Printf("This is the end: %s\n", name)
			if strings.ToLower(name) == "xml" {
				break
			}
			if ele == name && ele != "" {
				m.Set(ele, val)
				ele = ""
				val = ""
			}
			// 处理字符数据（这里就是元素的文本）
		case xml.CharData:
			// content := string(token)
			// fmt.Printf("This is the content: %v\n", content)
			val = string(token)
			// 异常处理(Log输出）
		default:
			log.Println(token)
		}

	}

	return m
}

func Time(t ...*time.Time) string {
	if t == nil {
		return strconv.Itoa(time.Now().Nanosecond())
	}
	return strconv.Itoa(t[0].Nanosecond())
}

func GenerateNonceStr() string {
	return GenerateUUID()
}

func GenerateUUID() string {
	s := uuid.NewV1().String()
	s = strings.Replace(s, "-", "", -1)
	run := ([]rune)(s)[:32]
	return string(run)
}

func In(source []string, v string) bool {
	for _, v0 := range source {
		if v0 == v {
			return true
		}
	}
	return false
}

// MapToString
func MapToString(data Map, skip []string) string {
	var keys sort.StringSlice
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	var sign []string

	for _, k := range keys {
		if In(skip, k) {
			continue
		}

		v := strings.TrimSpace(data.GetString(k))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	log.Println(strings.Join(sign, "&"))
	return strings.Join(sign, "&")
}

func ToUrlParams(data Map, skip []string) string {
	keys := data.SortKeys()
	var sign []string

	for _, k := range keys {
		if In(skip, k) {
			continue
		}
		v := strings.TrimSpace(data.GetString(k))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	return strings.Join(sign, "&")
}

// CurrentTimeStampMS get current time with millisecond
func CurrentTimeStampMS() int64 {
	return time.Now().UnixNano() / time.Millisecond.Nanoseconds()
}

// CurrentTimeStampNS get current time with nanoseconds
func CurrentTimeStampNS() int64 {
	return time.Now().UnixNano()
}

// CurrentTimeStamp get current time with unix
func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func CurrentTimeStampString() string {
	return strconv.FormatInt(CurrentTimeStamp(), 10)
}

func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func signatureSHA1(m Map) string {
	keys := m.SortKeys()
	var sign []string
	for _, k := range keys {
		if v := strings.TrimSpace(m.GetString(k)); v != "" {
			log.Debug(k, v)
			sign = append(sign, strings.Join([]string{k, v}, "="))
		} else if v := m.GetInt64(k); v != 0 {
			log.Debug(k, v)
			sign = append(sign, strings.Join([]string{k, strconv.FormatInt(v, 10)}, "="))
		}
	}
	sb := strings.Join(sign, "&")
	return SHA1(sb)
}

// 随机字符串
func GenerateRandomString2(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func GenerateRandomString(size int, kind ...RandomKind) string {

	bytes := RandomString[T_RAND_ALL]
	if kind != nil {
		if k, b := RandomString[kind[0]]; b == true {
			bytes = k
		}
	}
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
