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

// GenerateSignature make sign from map data
func GenerateSignature(m util.Map, key string, fn SignFunc) string {
	keys := m.SortKeys()
	var sign []string

	for _, k := range keys {
		if k == FieldSign {
			continue
		}
		v := strings.TrimSpace(m.GetString(k))

		if len(v) > 0 {
			log.Debug(k, v)
			sign = append(sign, strings.Join([]string{k, v}, "="))
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
