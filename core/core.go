package core

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/util"
)

type SignType int

const (
	SIGN_TYPE_MD5        SignType = iota
	SIGN_TYPE_HMACSHA256 SignType = iota
)

func (t SignType) String() string {
	if t == SIGN_TYPE_HMACSHA256 {
		return HMACSHA256
	}
	return MD5
}

// SandboxSignKey get wechat sandbox sign key
func SandboxSignKey(config config.Config) []byte {
	m := make(util.Map)
	m.Set("mch_id", config.Get("mch_id"))
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := GenerateSignature(m, config.Get("aes_key"), MakeSignMD5)
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
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && ip != "127.0.0.1" {
		return ip
	}
	ip = r.Header.Get("X-Forwarded-For")
	if ip == "" {
		return "127.0.0.1"
	}
	return ip
}

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
	io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// GenerateSignature make sign from map data
func GenerateSignature(m util.Map, key string, fn SignFunc) string {
	m0 := m.Clone()
	m0.Set("key", key)

	if fn == nil {
		fn = MakeSignMD5
	}
	return fn(util.MapToString(m0, []string{FIELD_SIGN}), key)

}
