package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/godcong/wego/log"
	"io"
	"strings"
)

/*SignFunc sign函数定义 */
type SignFunc func(data, key string) string

// FieldSign ...
const FieldSign = "sign"

// FieldSignType ...
const FieldSignType = "sign_type"

// FieldLimit ...
const FieldLimit = "limit"

/*HMACSHA256 定义:HMAC-SHA256 */
const HMACSHA256 = "HMAC-SHA256"

/*MD5 定义:MD5 */
const MD5 = "MD5"

// MakeSignHMACSHA256 make sign with hmac-sha256
func MakeSignHMACSHA256(data, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(data))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// MakeSignMD5 make sign with md5
func MakeSignMD5(data, key string) string {
	m := md5.New()
	_, _ = io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// GenerateSignatureWithIgnore ...
func GenerateSignatureWithIgnore(p Map, key string, ignore []string) string {
	m := p.Expect(ignore)
	keys := m.SortKeys()
	var sign []string
	size := len(keys)

	for i := 0; i < size; i++ {
		v := strings.TrimSpace(m.GetString(keys[i]))
		if len(v) > 0 {
			log.Debug(keys[i], v)
			sign = append(sign, strings.Join([]string{keys[i], v}, "="))
		}
	}

	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	fn := MakeSignMD5
	if p.GetString("sign_type") == HMACSHA256 {
		fn = MakeSignHMACSHA256
	}

	return fn(sb, key)
}

// GenerateSignature make sign from map data
func GenerateSignature(p Map, key string, fn SignFunc) string {
	m := p.Expect([]string{FieldSign})
	keys := m.SortKeys()
	var sign []string
	size := len(keys)
	for i := 0; i < size; i++ {
		v := strings.TrimSpace(m.GetString(keys[i]))

		if len(v) > 0 {
			log.Debug(keys[i], v)
			sign = append(sign, strings.Join([]string{keys[i], v}, "="))
		}
	}

	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	return fn(sb, key)
}

// ValidateSign check the sign validate
func ValidateSign(maps Map, key string) bool {
	if !maps.Has("sign") {
		return false
	}
	sign := maps.GetString("sign")
	newSign := ""
	switch maps.GetString("sign_type") {
	case HMACSHA256:
		newSign = GenerateSignature(maps, key, MakeSignHMACSHA256)
	default:
		newSign = GenerateSignature(maps, key, MakeSignMD5)
	}
	log.Debug(sign, newSign)
	if strings.Compare(sign, newSign) == 0 {
		return true
	}

	return false
}
