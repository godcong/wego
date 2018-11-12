package payment

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"io"
	"strings"
)

/*SignFunc sign函数定义 */
type SignFunc func(data, key string) string

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

func signIgnore(current string, s []string) bool {
	size := len(s)
	for j := 0; j < size; j++ {
		if current == s[j] {
			return true
		}
	}
	return false
}

// GenerateSignature2 ...
func GenerateSignature2(p util.Map, key string, ignore ...string) string {
	keys := p.SortKeys()
	var sign []string
	size := len(keys)

	for i := 0; i < size; i++ {
		if signIgnore(keys[i], ignore) {
			continue
		}
		v := strings.TrimSpace(p.GetString(keys[i]))

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
func GenerateSignature(m util.Map, key string, fn SignFunc) string {
	keys := m.SortKeys()
	var sign []string
	size := len(keys)
	for i := 0; i < size; i++ {
		if keys[i] == FieldSign {
			continue
		}
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

// ValidateSign ...
func ValidateSign(maps util.Map, key string) bool {
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
