package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/juju/errors"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"hash/crc32"
	"io"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*CustomHeader xml header*/
const CustomHeader = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`

/*CDATA xml cdata defines */
type CDATA struct {
	XMLName xml.Name
	Value   string `xml:",cdata"`
}

/* error types */
var (
	ErrorSignType  = errors.New("sign type error")
	ErrorParameter = errors.New("JsonApiParameters() check error")
	ErrorToken     = errors.New("EditAddressParameters() token is nil")
)

/*RandomKind RandomKind */
type RandomKind int

/*random kinds */
const (
	RandomNum      RandomKind = iota // 纯数字
	RandomLower                      // 小写字母
	RandomUpper                      // 大写字母
	RandomLowerNum                   // 数字、小写字母
	RandomUpperNum                   // 数字、大写字母
	RandomAll                        // 数字、大小写字母
)

/*RandomString defines */
var (
	RandomString = map[RandomKind]string{
		RandomNum:      "0123456789",
		RandomLower:    "abcdefghijklmnopqrstuvwxyz",
		RandomUpper:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomLowerNum: "0123456789abcdefghijklmnopqrstuvwxyz",
		RandomUpperNum: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		RandomAll:      "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
)

/*ParseNumber parse interface to number */
func ParseNumber(v interface{}) (float64, bool) {
	switch v0 := v.(type) {
	case float64:
		return v0, true
	case float32:
		return float64(v0), true
	}
	return 0, false
}

/*ParseInt parse interface to int64 */
func ParseInt(v interface{}) (int64, bool) {
	switch v0 := v.(type) {
	case int:
		return int64(v0), true
	case int32:
		return int64(v0), true
	case int64:
		return int64(v0), true
	case uint:
		return int64(v0), true
	case uint32:
		return int64(v0), true
	case uint64:
		return int64(v0), true
	case float64:
		return int64(v0), true
	case float32:
		return int64(v0), true
	default:
	}
	return 0, false
}

/*ParseString parse interface to string */
func ParseString(v interface{}) (string, bool) {
	switch v0 := v.(type) {
	case string:
		return v0, true
	case []byte:
		return string(v0), true
	case bytes.Buffer:
		return v0.String(), true
	default:
	}
	return "", false
}

/*Time get time string */
func Time(t ...time.Time) string {
	if t == nil {
		return strconv.Itoa(time.Now().Nanosecond())
	}
	return strconv.Itoa(t[0].Nanosecond())
}

/*GenerateNonceStr GenerateNonceStr */
func GenerateNonceStr() string {
	return GenerateUUID()
}

/*GenerateUUID GenerateUUID */
func GenerateUUID() string {
	s := uuid.NewV1().String()
	s = strings.Replace(s, "-", "", -1)
	run := ([]rune)(s)[:32]
	return string(run)
}

/*In check v is in source */
func In(source []string, v string) bool {
	size := len(source)
	for i := 0; i < size; i++ {
		if source[i] == v {
			return true
		}
	}

	return false
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

// CurrentTimeStampString get current time to string
func CurrentTimeStampString() string {
	return strconv.FormatInt(CurrentTimeStamp(), 10)
}

// GenSHA1 transfer strings to sha1
func GenSHA1(text ...string) string {
	sort.Strings(text)
	s := strings.Join(text, "")
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

// CRC32 ...
func CRC32(data string) string {
	ieee := crc32.NewIEEE()
	_, _ = io.WriteString(ieee, data)
	return fmt.Sprintf("%X", ieee.Sum32())
}

// GenMD5 transfer strings to md5
func GenMD5(data string) string {
	m := md5.New()
	_, _ = io.WriteString(m, data)
	return fmt.Sprintf("%x", m.Sum(nil))
}

// GenSHA256 ...
func GenSHA256(data []byte, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write(data)
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

func signatureSHA1(m Map) string {
	keys := m.SortKeys()
	var sign []string
	size := len(keys)
	for i := 0; i < size; i++ {
		if v := strings.TrimSpace(m.GetString(keys[i])); v != "" {
			log.Debug(keys[i], v)
			sign = append(sign, strings.Join([]string{keys[i], v}, "="))
		} else if v, b := m.GetInt64(keys[i]); b {
			log.Debug(keys[i], v)
			sign = append(sign, strings.Join([]string{keys[i], strconv.FormatInt(v, 10)}, "="))
		}
	}

	sb := strings.Join(sign, "&")
	return GenSHA1(sb)
}

//GenerateRandomString2 随机字符串
func GenerateRandomString2(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

//GenerateRandomString 随机字符串
func GenerateRandomString(size int, kind ...RandomKind) string {
	bytes := RandomString[RandomAll]
	if kind != nil {
		if k, b := RandomString[kind[0]]; b == true {
			bytes = k
		}
	}
	var result []byte
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

// MustInt64 ...
func MustInt64(src, def int64) int64 {
	if src != 0 {
		return src
	}
	return def
}

/*MapSortSplice MapSortSplice */
func MapSortSplice(data Map, ignore []string) string {
	var sign []string
	m := data.Expect(ignore)
	keys := m.SortKeys()
	size := len(keys)

	for i := 0; i < size; i++ {
		v := strings.TrimSpace(m.GetString(keys[i]))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{keys[i], v}, "="))
		}
	}

	log.Debug(strings.Join(sign, "&"))
	return strings.Join(sign, "&")
}

/*ToURLParams map to url params */
func ToURLParams(data Map, ignore []string) string {
	var sign []string
	m := data.Expect(ignore)
	keys := m.SortKeys()
	size := len(keys)
	for i := 0; i < size; i++ {
		v := strings.TrimSpace(m.GetString(keys[i]))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{keys[i], v}, "="))
		}
	}
	return strings.Join(sign, "&")
}

// AnyToMap convert interface to map
func AnyToMap(v interface{}) (Map, error) {
	m := Map{}
	b, e := jsoniter.Marshal(v)
	if e != nil {
		return m, e
	}
	e = jsoniter.Unmarshal(b, &m)
	return m, e
}

/*XMLToMap Convert XML to MAP */
func XMLToMap(xml []byte) Map {
	m, err := xmlToMap(xml, false)
	if err != nil {
		return nil
	}
	return m
}

/*MapToXML Convert MAP to XML */
func MapToXML(m Map) ([]byte, error) {
	return mapToXML(m, false)
}

/*JSONToMap Convert JSON to MAP */
func JSONToMap(xml []byte) Map {
	m := Map{}
	err := jsoniter.Unmarshal(xml, &m)
	if err != nil {
		log.Error(err)
	}
	return m
}

func mapToXML(maps Map, needHeader bool) ([]byte, error) {

	buff := bytes.NewBuffer([]byte(CustomHeader))
	if needHeader {
		buff.Write([]byte(xml.Header))
	}

	enc := xml.NewEncoder(buff)
	err := marshalXML(maps, enc, xml.StartElement{Name: xml.Name{Local: "xml"}})
	if err != nil {
		return nil, err
	}
	err = enc.Flush()
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
func xmlToMap(contentXML []byte, hasHeader bool) (Map, error) {
	m := make(Map)
	dec := xml.NewDecoder(bytes.NewReader(contentXML))
	err := unmarshalXML(m, dec, xml.StartElement{Name: xml.Name{Local: "xml"}}, true)
	if err != nil {
		return nil, xerrors.Errorf("xml to map:%w", err)
	}

	return m, nil
}

// MustString ...
func MustString(v, def string) string {
	if v == "" {
		return def
	}
	return v
}

// MustInt ...
func MustInt(v string, def int) int {
	i, err := strconv.Atoi(v)
	if err == nil {
		return i
	}
	return def
}
