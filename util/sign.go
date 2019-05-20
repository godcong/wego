package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
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

// SignSHA256 make sign with hmac-sha256
func SignSHA256(data, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(data))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// SignMD5 make sign with md5
func SignMD5(data, key string) string {
	m := md5.New()
	_, _ = io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// GenSignWithIgnore ...
func GenSignWithIgnore(p Map, key string, ignore []string) string {
	return GenSign(p, key, ignore...)
}

// MapSignFunc ...
func MapSignFunc(p Map) SignFunc {
	if p.GetString("sign_type") == HMACSHA256 {
		return SignSHA256
	}
	return SignMD5
}

// GenSign make sign from map data
func GenSign(p Map, key string, ignores ...string) string {
	log.Debug("sign:", p, key, ignores)
	exp := append(ignores[:], FieldSign)
	m := p.Expect(exp)
	keys := m.SortKeys()
	var sign []string
	size := len(keys)
	for i := 0; i < size; i++ {
		v := strings.TrimSpace(m.GetString(keys[i]))
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{keys[i], v}, "="))
		}
	}

	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	return MapSignFunc(p)(sb, key)
}

// ValidateSign check the sign validate
func ValidateSign(p Map, key string) bool {
	if !p.Has("sign") {
		return false
	}
	sign := p.GetString("sign")
	newSign := GenSign(p, key)
	log.Debug(sign, newSign)

	if strings.Compare(sign, newSign) == 0 {
		return true
	}
	return false
}
