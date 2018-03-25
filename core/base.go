package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/godcong/wego/core/message"
	"github.com/satori/go.uuid"
)

const CUSTOM_HEADER = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>`

var (
	ErrorSignType  = errors.New("sign type error")
	ErrorParameter = errors.New("JsonApiParameters() check error")
	ErrorToken     = errors.New("EditAddressParameters() token is nil")
)

type CDATA = message.CDATA

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

// MapToString
func MapToString(data Map) string {
	var keys sort.StringSlice
	for k := range data {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	var sign []string

	for _, k := range keys {
		if k == FIELD_SIGN {
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

func ToUrlParams(data Map) string {
	keys := data.SortKeys()
	var sign []string
	for _, k := range keys {
		if k == FIELD_SIGN {
			continue
		}
		v := strings.TrimSpace(data.GetString(k))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}

	return strings.Join(sign, "&")

}

// MakeSignMD5 make sign with md5
func MakeSignMD5(data string) string {
	m := md5.New()
	io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// MakeSignHMACSHA256 make sign with hmac-sha256
func MakeSignHMACSHA256(data, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(data))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// MapToXml Convert MAP to XML
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
			Println(token)
		}

	}

	return m
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
	m := sha1.New()
	m.Write([]byte(s))
	return fmt.Sprintf("%x", m.Sum(nil))
}

func ParseNumber(v interface{}) float64 {

	switch v.(type) {
	case float64:
		return v.(float64)
	case float32:
		return float64(v.(float32))
	}

	return -1
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
	return -1
}

// GenerateSignature make sign from map data
func GenerateSignature(m Map, key string, signType SignType) string {
	keys := m.SortKeys()
	var sign []string

	for _, k := range keys {
		if k == FIELD_SIGN {
			continue
		}
		v := strings.TrimSpace(m.GetString(k))

		if len(v) > 0 {
			Debug(k, v)
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	if signType == SIGN_TYPE_HMACSHA256 {
		return MakeSignHMACSHA256(sb, key)
	} else {
		return MakeSignMD5(sb)
	}

}

// SandboxSignKey get wechat sandbox sign key
func SandboxSignKey(config Config) []byte {
	m := make(Map)
	m.Set("mch_id", config.Get("mch_id"))
	m.Set("nonce_str", GenerateNonceStr())
	sign := GenerateSignature(m, config.Get("aes_key"), SIGN_TYPE_MD5)
	m.Set("sign", sign)
	// _ = NewApplication(config)
	// return app.GetRequest().Request(SANDBOX_SIGNKEY_URL_SUFFIX, m)
	return []byte(nil)
}

func GetServerIp() string {
	adds, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, address := range adds {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return "127.0.0.1"
}

func GetClientIp(r *http.Request) string {
	ip := r.Header.Get("Remote_addr")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}
